package orionsdk

type SolarWinds struct {
	swis *SwisClient
}


func newSolarWinds(npmServer, username, password string) *SolarWinds {
	return &SolarWinds{
		newSwisClient(npmServer, username, password),
	}
}


func (s *SolarWinds) doesNodesExist(nodeName string) bool {
	if s.getNodes(nodeName) == nil {
		return false
	} else{
		return true
	}
}

func (s *SolarWinds) getNodes(nodeName string ) interface{} {
	nodeId, _ := s.swis.query("SELECT NodeID, Caption FROM Orion.Nodes WHERE Caption = @caption", []string{nodeName})
	if nodeId == nil {
		return nil
	}else {
		return nodeId
	}
}

func (s *SolarWinds) getNodeUri(nodeName string) interface{}{
	nodeUri, _ := s.swis.query("SELECT Caption, Uri FROM Orion.Nodes WHERE Caption = @caption", []string{nodeName})
	if nodeUri == nil {
		return nil
	}else {
		return nodeUri
	}
}
