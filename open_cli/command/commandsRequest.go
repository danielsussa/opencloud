package command

import (
	"flag"
	"github.com/danielsussa/opencloud/shared"
	"log"
)

type flags struct {
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

	}
	log.Fatal("cannot find command")
	return nil
}

func loadFlags() flags {
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
