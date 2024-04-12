package service

import (
	"fmt"
	"log"
	. "micro-batching/api"
)

type BatchProcessor struct{}

// Process is one method of BatchProcessor which should be implemented externally
func (b *BatchProcessor) Process(jobs []Job) {
	for i := range jobs {
		jobs[i].Status = SUBMITTED
	}

	var output string
	for _, job := range jobs {
		output = output + fmt.Sprintf("[%s %s \"%s\"] ", job.Id.String()[:8], job.Type, *job.Name)
	}
	log.Printf("Processed jobs: %s", output)
}
