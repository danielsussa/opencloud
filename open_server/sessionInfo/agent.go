package sessionInfo

import (
	"github.com/danielsussa/opencloud/open_server/data"
	"github.com/gliderlabs/ssh"
)

// user -> agentSessionInfo
var sessionInfo map[string]*agentSessionInfo

func init() {
	sessionInfo = make(map[string]*agentSessionInfo)
}

type agentSessionInfo struct {
	Session   ssh.Session
	PublicKey string
	Port      int

	reverseProxyMap map[int]reverseProxyInfo // (localPort) -> remotePort
}

type reverseProxyInfo struct {
	remotePort int
	name       string
}

func LoadAgentData() {
	currentData := data.GetData()
	for name, agent := range currentData.Agents {
		sessionInfo[name] = &agentSessionInfo{
			PublicKey: agent.PublicKey,
		}

		rpMap := make(map[int]reverseProxyInfo)
		for rpName, rp := range agent.ReverseProxy {
			rpMap[rp.Port] = reverseProxyInfo{
				remotePort: 0,
				name:       rpName,
			}
		}
		sessionInfo[name].reverseProxyMap = rpMap
	}
}

func (s *agentSessionInfo) HasReverseProxy(localPort int) bool {
	if _, ok := s.reverseProxyMap[localPort]; ok {
		return true
	}
	return false
}

func (s *agentSessionInfo) DeleteReverseProxy(name string) {
	for k, val := range s.reverseProxyMap {
		if val.name == name {
			delete(s.reverseProxyMap, k)
		}
	}
}

func (s *agentSessionInfo) AddReverseProxy(name string, localPort, remotePort int) {
	if s.reverseProxyMap == nil {
		s.reverseProxyMap = make(map[int]reverseProxyInfo)
	}
	s.reverseProxyMap[localPort] = reverseProxyInfo{
		remotePort: remotePort,
		name:       name,
	}
}

func AddAgentInfo(user string, key string, port int, session ssh.Session) {
	sessionInfo[user] = &agentSessionInfo{
		Session:   session,
		PublicKey: key,
		Port:      port,
	}
}

func AddAgentInfoKey(user string, key string) {
	sessionInfo[user] = &agentSessionInfo{
		PublicKey: key,
	}
}

func GetAgentInfo(user string) *agentSessionInfo {
	if v, ok := sessionInfo[user]; ok {
		return v
	}
	return nil
}

func HasKey(key string) bool {
	for _, info := range sessionInfo {
		if info.PublicKey == key {
			return true
		}
	}
	return false
}
