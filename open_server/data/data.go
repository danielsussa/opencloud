package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type (
	data struct {
		Agents map[string]dataAgent
	}
	dataAgent struct {
		PublicKey    string
		ReverseProxy map[string]reverseProxyData
	}
	reverseProxyData struct {
		Port int
	}
)

var currentData *data

func GetData() *data {
	return currentData
}

func InitData(path string) {
	if currentData != nil {
		return
	}
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		currentData = &data{}
		SaveData(path)
		return
	}
	json.Unmarshal(dat, &currentData)

}

func SaveData(path string) {
	b, err := json.Marshal(currentData)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(path, b, 0644)
}

// load data
func init() {

}
