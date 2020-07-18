package data

import "errors"

type dataWeb struct {
	Port int
}

type Web map[string]*dataWeb

func (w *Web) SetDomain(domain string, port int) {
	dw := &dataWeb{Port: port}
	(*w)[domain] = dw
}

func (w Web) GetPort(domain string) (int, error) {
	if _, ok := w[domain]; !ok {
		return 0, errors.New("domain not found")
	}
	return w[domain].Port, nil
}
