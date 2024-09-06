package go_specs_greet

import (
	"io"
	"net/http"
)

type Driver struct {
	BaseUrl string
}

func (d Driver) Greet() (string, error) {
	res, err := http.Get(d.BaseUrl + "/greet")
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
