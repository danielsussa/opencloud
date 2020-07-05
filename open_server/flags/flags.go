package flags

import (
	"flag"
)

type flags struct {
	DataPath   *string
	ClientPort *string
	AgentPort  *string
}

var currFlags flags

func GetDataPath() string {
	return *currFlags.DataPath + "/data.json"
}

func init() {
	dataPath := flag.String("dataPath", ".data", "")
	clientPort := flag.String("clientPort", ":9999", "")
	agentPort := flag.String("agentPort", ":2222", "")
	flag.Parse()
	currFlags = flags{
		DataPath:   dataPath,
		ClientPort: clientPort,
		AgentPort:  agentPort,
	}
}
