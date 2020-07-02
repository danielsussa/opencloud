package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type fileConfig struct {
	Host string
	AgentName string
}

func loadConfig()(fc fileConfig){
	file, err := ioutil.ReadFile("config/config.json")
	if err != nil{
		fc = makeConfig()
	}
	json.Unmarshal(file, &fc)
	return fc
}

func makeConfig()(fc fileConfig){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter server host (ex: localhost:9090): ")
	host, _ := reader.ReadString('\n')
	fc.Host = strings.Trim(host,"\n")

	fmt.Print("Enter agent name (ex: rasp-1): ")
	agentName, _ := reader.ReadString('\n')
	fc.AgentName = strings.Trim(agentName,"\n")

	b,_ := json.Marshal(fc)
	err := ioutil.WriteFile("config/config.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return fc
}

