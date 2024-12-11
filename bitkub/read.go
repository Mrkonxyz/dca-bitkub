package bitkub

import (
	"io"
	"log"
)

func ReadResponse(r io.Reader) []byte {
	body, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	return body
}
