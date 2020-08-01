package flags

import (
	"flag"
)

type flags struct {
	DataPath   *string
	ClientPort *string
	AgentPort  *string
	TlsPublic  *string
	TlsPrivate *string
	TlsPort    *string
}

var currFlags flags

func GetDataPath() string {
	return *currFlags.DataPath + "/data.json"
}

func GetTlsPrivate() string {
	return *currFlags.TlsPrivate
}

func GetTlsPort() string {
	return *currFlags.TlsPort
}

func GetTlsPublic() string {
	return *currFlags.TlsPublic
}

func GetClientPort() string {
	return *currFlags.ClientPort
}

func GetQgentPort() string {
	return *currFlags.AgentPort
}

func init() {
	dataPath := flag.String("dataPath", ".data", "")
	clientPort := flag.String("clientPort", ":9999", "")
	agentPort := flag.String("agentPort", ":2222", "")

	tlsPrivate := flag.String("tlsPrivate", "", "")
	tlsPublic := flag.String("tlsPublic", "", "")
	tlsPort := flag.String("tlsPort", ":443", "")

	flag.Parse()
	currFlags = flags{
		DataPath:   dataPath,
		ClientPort: clientPort,
		AgentPort:  agentPort,
		TlsPrivate: tlsPrivate,
		TlsPublic:  tlsPublic,
		TlsPort:    tlsPort,
	}
}
