package main

import (
	"github.com/deqinganz/micro-batching"
	"log"
	"os"

	"micro-batching-service/internal"
)

func main() {
	b, err := batching.NewBatching()
	if err != nil {
		path, _ := os.Getwd()
		log.Fatalf("creating batching service failed: %v in %s", err, path)
		return
	}

	internal.SetupHandler(&b)
}
