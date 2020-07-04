package main

import (
	"flag"
	"fmt"
	"github.com/danielsussa/opencloud/shared"
	"log"
)

type flags struct {
	Agent   *string
	Key     *string
	Command *string
}

func (f flags) returnCommandRequest() string {
	switch *f.Command {
	case shared.NEW_AGENT:
		return newAgentCommandRequest{Key: *f.Key, Agent: *f.Agent}.Text()
	case shared.PING:
		return pingCommandRequest{Agent: *f.Agent}.Text()

	}
	log.Fatal("cannot find command")
	return ""
}

func loadAllFlags() flags {
	agent := flag.String("agent", "", "")
	key := flag.String("key", "", "")
	command := flag.String("command", "", "")
	flag.Parse()
	return flags{
		Agent:   agent,
		Key:     key,
		Command: command,
	}
}

type newAgentCommandRequest struct {
	Key   string
	Agent string
}

func (newAgent newAgentCommandRequest) Text() string {
	return fmt.Sprintf("%s %s %s", "new_agent", newAgent.Key, newAgent.Agent)
}

type pingCommandRequest struct {
	Agent string
}

func (ping pingCommandRequest) Text() string {
	return fmt.Sprintf("%s %s", "ping", ping.Agent)
}
