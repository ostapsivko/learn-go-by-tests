package httpserver

import (
	"io"
	"net/http"
)

type Driver struct {
	BaseUrl string
	Client  *http.Client
}

func (d Driver) Greet(name string) (string, error) {
	res, err := d.Client.Get(d.BaseUrl + "/greet?name=" + name)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	greeting, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(greeting), nil
}

func (d Driver) Curse(name string) (string, error) {
	res, err := d.Client.Get(d.BaseUrl + "/curse?name=" + name)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	curse, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(curse), nil
}
