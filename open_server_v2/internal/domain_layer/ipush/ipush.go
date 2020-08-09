package ipush

type Push struct {
	SendPingCommand func(agent string)error
}

var push Push

func Get()Push {
	return push
}
