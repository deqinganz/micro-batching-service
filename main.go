package main

import (
	"log"
	"micro-batching/internal/http"
	"micro-batching/internal/service"
	"os"
)

func main() {
	batching, err := service.NewBatching()
	if err != nil {
		path, _ := os.Getwd()
		log.Fatalf("creating batching service failed: %v in %s", err, path)
		return
	}

	http.SetupHandler(&batching)

	//numberOfSeconds := batching.GetFrequency().Frequency // invoke every n seconds
	//
	//// create a scheduler
	//s, err := gocron.NewScheduler()
	//if err != nil {
	//	// handle error
	//}
	//
	//// add a job to the scheduler
	//_, err = s.NewJob(
	//	gocron.DurationJob(
	//		time.Duration(numberOfSeconds/numberOfSeconds)*time.Second,
	//	),
	//	gocron.NewTask(
	//		func() {
	//			batching.Post()
	//		},
	//	),
	//)
	//if err != nil {
	//	// handle error
	//}
	//
	//s.Start() // block until you are ready to shut down
	//select {
	//case <-time.After(10 * time.Second):
	//}
	//
	//s.Shutdown()
}
