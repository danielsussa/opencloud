package internal

import "golang.org/x/crypto/ssh"

var channelMap map[string]ssh.Channel

func GetChannelByAgent(agent string)ssh.Channel{
	return channelMap[agent]
}
