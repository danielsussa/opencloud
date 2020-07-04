package command

import (
	"flag"
	"github.com/danielsussa/opencloud/shared"
	"log"
)

type ApiCommand interface {
	Request() string
	Response(strArr []string) string
}


type flags struct {
	Port    *string
	Agent   *string
	Key     *string
	Command *string
}

func ReturnCommand() ApiCommand {
	flags := loadFlags()
	switch *flags.Command {
	case shared.NEW_AGENT:
		return newAgentCommand{flags: flags}
	case shared.PING:
		return pingCommand{flags: flags}
	case shared.ADD_REVERSE_PROXY:
		return addReverseProxy{flags: flags}

	}
	log.Fatal("cannot find command")
	return nil
}

func loadFlags() flags {
	agent := flag.String("agent", "", "")
	key := flag.String("key", "", "")
	port := flag.String("port", "", "")
	command := flag.String("command", "", "")
	flag.Parse()
	return flags{
		Agent:   agent,
		Port:    port,
		Key:     key,
		Command: command,
	}
}
