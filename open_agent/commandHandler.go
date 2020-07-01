package main

type ApiCommandRequest struct {
	CommandType CommandType
	Command     ApiCommand
}

type ApiCommand interface {
	Kind() CommandType
}

type CommandType string

const (
	ADD_REVERSE_PROXY CommandType = "ADD_REVERSE_PROXY"
)

type ApiCommandHandler func(apiCommandRequest ApiCommandRequest, c OpenAgent)

func commandHandler(apiCommandRequest ApiCommandRequest, c OpenAgent) {
	switch apiCommandRequest.CommandType {
	case ADD_REVERSE_PROXY:
		c.AddSshReverseProxy(c, apiCommandRequest.Command.(ReverseProxyInfo))
	}
}
