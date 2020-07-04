package info

import "github.com/gliderlabs/ssh"

type agentSessionInfo struct {
	Session   ssh.Session
	PublicKey string
	Port      int
}

var sessionInfo map[string]agentSessionInfo

func init() {
	sessionInfo = make(map[string]agentSessionInfo)
}

func AddAgentInfo(user string, key string, port int, session ssh.Session) {
	sessionInfo[user] = agentSessionInfo{
		Session:   session,
		PublicKey: key,
		Port:      port,
	}
}

func AddAgentInfoKey(user string, key string) {
	sessionInfo[user] = agentSessionInfo{
		PublicKey: key,
	}
}

func GetAgentInfo(user string) *agentSessionInfo {
	if v, ok := sessionInfo[user]; ok {
		return &v
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

type AgentProxyConnectionInfo struct {
}
