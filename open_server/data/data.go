package data

import (
	"encoding/json"
	"errors"
	"github.com/danielsussa/freeport"
	"github.com/danielsussa/opencloud/open_server/flags"
	"io/ioutil"
	"log"
)

type data struct {
	Agents map[string]*dataAgent
	Web    Web // key=host(dubdomain.domain.com.br)
}

var loadedData *data
var currentData *data

func GetData() *data {
	return currentData
}

func SaveAgentProfile(agent string) {
	currAgentData := currentData.Agents[agent]
	loadedData.Agents[agent] = currAgentData
	saveData(loadedData)
}

func (d *data) NewAgent(user string, key string) error {
	port, err := GetNewFreeNotAllocatedPort(currentData)
	if err != nil {
		return err
	}

	agentData := &dataAgent{
		PublicKey:    key,
		ReverseProxy: nil,
		Port:         port,
	}
	d.Agents[user] = agentData
	return nil
}

func GetNewFreeNotAllocatedPort(d *data) (int, error) {
	port := 0
	var err error
	maxTries := 8
	for {
		port, err = freeport.GetFreePort()
		if err != nil {
			maxTries--
		}
		if isPortAllocated(d, port) {
			maxTries--
		}
		if maxTries < 0 {
			return 0, errors.New("cannot generate free port")
		}
		break
	}
	return port, nil
}

func isPortAllocated(d *data, port int) bool {
	for _, agent := range d.Agents {
		if agent.Port == port {
			return true
		}
		for _, rp := range agent.ReverseProxy {
			if rp.RemotePort == port {
				return true
			}
		}
	}
	return false
}

func AnyUserHasKey(key string) bool {
	for _, agent := range currentData.Agents {
		if agent.PublicKey == key {
			return true
		}
	}
	return false
}

func GetAgentData(agent string) *dataAgent {
	if _, ok := currentData.Agents[agent]; !ok {
		return nil
	}
	return currentData.Agents[agent]
}

func InitFromConfig(d *data) {
	currentData = d
}

func InitFromFile() {
	if currentData != nil {
		return
	}
	dat, err := ioutil.ReadFile(flags.GetDataPath())
	if err != nil {
		agentMap := make(map[string]*dataAgent)
		currentData = &data{
			Agents: agentMap,
		}
		saveData(currentData)
		loadedData = currentData
		return
	}
	json.Unmarshal(dat, &currentData)
	json.Unmarshal(dat, &loadedData)
}

func saveData(d *data) {
	b, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(flags.GetDataPath(), b, 0644)
}
