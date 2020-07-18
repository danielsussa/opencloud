package data

import "github.com/danielsussa/freeport"

type dataAgent struct {
	Port         int
	PublicKey    string
	ReverseProxy map[string]reverseProxyData
}

type reverseProxyData struct {
	LocalPort  int
	RemotePort int
}

func (a dataAgent) HasReverseProxy(port int) bool {
	for _, rp := range a.ReverseProxy {
		if rp.LocalPort == port {
			return true
		}
	}
	return false
}

func (a *dataAgent) AddReverseProxy(name string, localPort, remotePort int) error {
	if a.ReverseProxy == nil {
		a.ReverseProxy = make(map[string]reverseProxyData)
	}
	a.ReverseProxy[name] = reverseProxyData{LocalPort: localPort, RemotePort: remotePort}
	return nil
}

func (a *dataAgent) DeleteReverseProxy(name string) {
	delete(a.ReverseProxy, name)
}

func (a *dataAgent) GetPort() (int, error) {
	if a.Port != 0 && freeport.CheckPortIsFree(a.Port) {
		return a.Port, nil
	}
	port, err := GetNewFreeNotAllocatedPort(currentData)
	if err != nil {
		return 0, err
	}
	a.Port = port
	return port, nil
}
