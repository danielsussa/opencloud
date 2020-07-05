package sessionInfo

import "github.com/gliderlabs/ssh"

var sessionInfo map[string]*agentSessionInfo

func init() {
	sessionInfo = make(map[string]*agentSessionInfo)
}

type agentSessionInfo struct {
	Session   ssh.Session
	PublicKey string
	Port      int

	reverseProxtMap map[int]reverseProxyInfo // (localPort) -> remotePort
}

type reverseProxyInfo struct {
	remotePort int
	name       string
}

func (s *agentSessionInfo) HasReverseProxy(localPort int) bool {
	if _, ok := s.reverseProxtMap[localPort]; ok {
		return true
	}
	return false
}

func (s *agentSessionInfo) AddReverseProxy(name string, localPort, remotePort int) {
	if s.reverseProxtMap == nil {
		s.reverseProxtMap = make(map[int]reverseProxyInfo)
	}
	s.reverseProxtMap[localPort] = reverseProxyInfo{
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
