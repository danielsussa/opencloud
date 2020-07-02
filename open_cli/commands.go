package main

import (
	"flag"
	"fmt"
	"log"
)

type flags struct {
	Agent *string
	Key *string
	Command *string
}

func loadAllFlags()flags{
	agent := flag.String("agent", "", "")
	key := flag.String("key", "", "")
	command := flag.String("command", "", "")
	flag.Parse()
	return flags{
		Agent:    agent,
		Key:     key,
		Command: command,
	}
}

type newAgentCommand struct {
	Key string
	Agent string
}

func(newAgent newAgentCommand)Text()string{
	return fmt.Sprintf("%s %s %s","new_agent" ,newAgent.Key, newAgent.Agent)
}

type pingCommand struct {
	Agent string
}

func(ping pingCommand)Text()string{
	return fmt.Sprintf("%s %s","ping" ,ping.Agent)
}

func(f flags) returnCommand()string{
	switch *f.Command {
	case "new_agent":
		return newAgentCommand{Key:  *f.Key, Agent: *f.Agent}.Text()
	case "ping":
		return pingCommand{Agent: *f.Agent}.Text()

	}
	log.Fatal("cannot find command")
	return ""
}
