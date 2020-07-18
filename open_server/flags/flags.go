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
}

var currFlags flags

func GetDataPath() string {
	return *currFlags.DataPath + "/data.json"
}

func GetTlsPrivate() string {
	return *currFlags.TlsPrivate
}

func GetTlsPublic() string {
	return *currFlags.TlsPublic
}

func init() {
	dataPath := flag.String("dataPath", ".data", "")
	clientPort := flag.String("clientPort", ":9999", "")
	agentPort := flag.String("agentPort", ":2222", "")

	tlsPrivate := flag.String("tlsPrivate", "", "")
	tlsPublic := flag.String("tlsPublic", "", "")

	flag.Parse()
	currFlags = flags{
		DataPath:   dataPath,
		ClientPort: clientPort,
		AgentPort:  agentPort,
		TlsPrivate: tlsPrivate,
		TlsPublic:  tlsPublic,
	}
}
