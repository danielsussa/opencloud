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

type ApiCommandHandler func(apiCommandRequest ApiCommandRequest, c Client)

func commandHandler(apiCommandRequest ApiCommandRequest, c Client) {
	switch apiCommandRequest.CommandType {
	case ADD_REVERSE_PROXY:
		c.AddSshReverseProxy(apiCommandRequest.Command.(ReverseProxyInfo), c)
	}
}
