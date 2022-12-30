package main

import (
	"net/http"
)

func apertura() error {
	_, err := http.Get(endpoint)
	if err != nil {
		return err
	}

	_, err = http.Get(endpoint)
	return err
}
