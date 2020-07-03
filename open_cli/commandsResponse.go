package main

import (
	"fmt"
	"strings"
)

type pingCommandResponse struct {
	str []string
}

func (ping pingCommandResponse) Text() string {
	return fmt.Sprintf("sucessfull ping to host")
}

type newAgentCommandResponse struct {
	str []string
}

func (ping newAgentCommandResponse) Text() string {
	return fmt.Sprintf("sucessfull creating new agent")
}

func returnCommandResponse(str string) string {
	strArr := strings.Split(str, " ")
	switch strArr[0] {
	case "ping":
		return pingCommandResponse{str: strArr}.Text()
	case "new_agent":
		return newAgentCommandResponse{str: strArr}.Text()
	}
	return fmt.Sprintf("error on command: %s", str)
}
