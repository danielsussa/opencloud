package main

import (
	"flag"
	"fmt"
	"log"
)

type flags struct {
	Name *string
	Key *string
	Command *string
}

func loadAllFlags()flags{
	name := flag.String("name", "", "")
	key := flag.String("key", "", "")
	command := flag.String("command", "", "")
	flag.Parse()
	return flags{
		Name:    name,
		Key:     key,
		Command: command,
	}
}

type newAgent struct {
	Key string
	Name string
}

func(newAgent newAgent)Text()string{
	return fmt.Sprintf("%s %s %s","new_agent" ,newAgent.Key, newAgent.Name)
}

func(f flags) returnCommand()string{
	switch *f.Command {
	case "new_agent":
		return newAgent{Key:  *f.Key, Name: *f.Name}.Text()
	}
	log.Fatal("cannot find command")
	return ""
}
