package main

import (
	"bytes"
	"net/http"
)

func apertura(e endpoint) error {
	// Adds the entity_id to the request
	req, err := http.NewRequest("POST", e.URL, bytes.NewBufferString("{\"entity_id\": \""+e.ID+"\"}"))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+HAToken)
	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	return err
}
