package service

import (
	. "micro-batching/api"
)

// Process is one method of BatchProcessor which should be implemented externally
func Process(jobs []Job) {
	for i := range jobs {
		jobs[i].Status = SUBMITTED
	}
}
