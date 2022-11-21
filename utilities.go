package main

import (
	"net/http"
)

func apertura() error {
	_, err := http.Get(endpoint + "/open")

	return err
}
