package service

import (
	. "micro-batching/api"
)

func Process(jobs []Job) {
	for i := range jobs {
		jobs[i].Status = SUBMITTED
	}
}
