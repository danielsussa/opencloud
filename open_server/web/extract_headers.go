package web

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type HttpInfo struct {
	host   string
	method string
	path   string
}

func (http HttpInfo) ExtractHostSulfix(sulfix string) string {
	return http.host
}

func (http HttpInfo) Host() string {
	return http.host
}

func ExtractHttpInfo(conn net.Conn) *HttpInfo {
	var info HttpInfo
	r := bufio.NewReader(conn)
	firstLine := true
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		if firstLine {
			fl := strings.Split(msg, " ")
			info.method = fl[0]
			info.path = fl[1]

		}
		firstLine = false
		if strings.Contains(msg, "Host: ") {
			info.host = msg[6 : len(msg)-2]
			break
		}
	}
	return &info
}
