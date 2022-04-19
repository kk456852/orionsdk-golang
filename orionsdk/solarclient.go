package orionsdk

import (
	"fmt"
	_ "net/http"
)


type SwisClient struct {
	Url 	string
	Session string
}

func newSwisClient(hostname, username, password string) *SwisClient {
	url := fmt.Sprintf("https://%s:17778/SolarWinds/InformationService/v3/Json/", hostname)
	return &SwisClient {
		Url: url,
		Session: "",
	}
}

// request 请求
func (c *SwisClient) _req(method, frag string, data map[string]interface{}) ([]byte, error) {
	return nil, nil
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