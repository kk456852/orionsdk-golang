package orionsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)


type SwisClient struct {
	Url 		string
	UserName 	string
	Password 	string
	httpClient  *http.Client
}

func newSwisClient(hostname, username, password string) *SwisClient {
	url := fmt.Sprintf("https://%s:17778/SolarWinds/InformationService/v3/Json/", hostname)
	return &SwisClient {
		Url: url,
		UserName: username,
		Password: password,
		httpClient: NewHttpClient(),
	}
}

func NewHttpClient() *http.Client {
	// cookiejar.New source code return jar, nil
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	return client
}

// request 请求
func (c *SwisClient) _req(method, frag string, data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Json marshal error: %v", err)
		return nil, err
	}
	httpReq, createReqErr := http.NewRequest(method, c.Url + frag, bytes.NewBuffer(jsonData))
	if createReqErr != nil {

	}
	httpReq.SetBasicAuth(c.UserName, c.Password)
	httpReq.Header.Add("Content-Type", "application/json")
	httpRes, err := c.httpClient.Do(httpReq)
	if err != nil || 400 <= httpRes.StatusCode  || httpRes.StatusCode < 600 {
		log.Printf("Request failed cause: %v, response body is %v" , err, httpRes.Body)
		return nil, err
	}
	defer httpRes.Body.Close()
	result, err := ioutil.ReadAll(httpRes.Body)
	if err != nil{
		log.Printf("Read http response body failed: %v", err)
		return nil, err
	}
	return result, nil
}

func (c *SwisClient) query(query string, params []string) ([]byte, error) {
	return c._req("POST", "QUERY", map[string]interface{}{
		"query" 	 : query,
		"parameters" : params,
	})
}

func (c *SwisClient) invoke(entity, verb string, args map[string]interface{}) ([]byte, error) {
	frag := fmt.Sprintf("Invoke/%s/%s", entity, verb)
	return c._req("POST", frag, args)
}

func (c *SwisClient) create(entity string, properties map[string]interface{}) ([]byte, error) {
	return c._req("POST", "Create/" + entity, properties)
}

func (c *SwisClient) read(uri string) ([]byte, error) {
	return c._req("GET", uri, nil)
}

func (c *SwisClient) update(uri string, properties map[string]interface{}) ([]byte, error) {
	return c._req("POST", uri, properties)
}

func (c *SwisClient) bulkUpdate(uris string, properties map[string]interface{}) ([]byte, error) {
	return c._req("POST", "BulkUpdate", map[string]interface{}{
		"uris" 		 : uris,
		"properties" : properties,
	})
}

func (c *SwisClient) delete(uri string) ([]byte, error) {
	return c._req("DELETE", uri, nil)
}

func (c *SwisClient) bulkDelete(uris string) ([]byte, error) {
	return c._req("POST", "BulkDelete", map[string]interface{}{
		"uris" : uris,
	})
}