package main

import (
	"net/http"
	"time"
)

func apertura() error {
	_, err := http.Get(endpoint + "/LED=ON1")
	if err != nil {
		return err
	}

	time.Sleep(100 * time.Millisecond)

	_, err = http.Get(endpoint + "/LED=OFF1")

	return err
}
