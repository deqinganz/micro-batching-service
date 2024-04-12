package main

import (
	"log"
	"micro-batching/internal/http"
	"micro-batching/internal/service"
	"os"
	"time"
)

func main() {
	batching, err := service.NewBatching()
	if err != nil {
		path, _ := os.Getwd()
		log.Fatalf("creating batching service failed: %v in %s", err, path)
		return
	}

	go http.SetupHandler(&batching)

	numberOfSeconds := batching.GetFrequency().Frequency // invoke every n seconds
	for range time.Tick(time.Duration(numberOfSeconds) * time.Second) {
		batching.Post()
	}
}
