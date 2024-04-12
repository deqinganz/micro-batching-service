package service

import (
	"container/ring"
	"fmt"
	"github.com/stretchr/testify/assert"
	. "micro-batching/api"
	"micro-batching/internal/config"
	"testing"
)

func TestTake(t *testing.T) {
	b := NewBatchingWithConfig(config.RunConfig{
		QueueSize: 1,
	})

	expectedUserName := "name"
	params := JobRequest_Params{}
	err := params.FromUpdateUserInfoParams(UpdateUserInfoParams{
		UserId: "12",
		Name:   &expectedUserName,
	})
	assert.NoError(t, err)
	jobRequest := JobRequest{
		Type:   "Type",
		Params: params,
	}

	job := b.Take(jobRequest)

	assert.Equal(t, QUEUED, job.Status)
	assert.NotNil(t, job.Id)

	updateUserInfo, err := job.Params.AsUpdateUserInfoParams()
	assert.Equal(t, "12", updateUserInfo.UserId)
	assert.Equal(t, expectedUserName, *updateUserInfo.Name)
}

func TestJobInfo(t *testing.T) {
	b := NewBatchingWithConfig(config.RunConfig{
		QueueSize: 1,
	})

	expectedUserName := "name"
	params := JobRequest_Params{}
	err := params.FromUpdateUserInfoParams(UpdateUserInfoParams{
		UserId: "12",
		Name:   &expectedUserName,
	})
	assert.NoError(t, err)
	jobRequest := JobRequest{
		Type:   "Type",
		Params: params,
	}

	job := b.Take(jobRequest)

	jobInfo, err := b.JobInfo(job.Id)
	assert.NoError(t, err)

	assert.Equal(t, job, jobInfo)
}

func TestA(t *testing.T) {
	// Create a new ring of size 6
	r := ring.New(6)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
	fmt.Println("hh")

	// Unlink three elements from r, starting from r.Next()
	r = r.Prev()
	r.Unlink(3)
	r = r.Next()

	r.Value = 6
	r = r.Next()

	// Iterate through the remaining ring and print its contents
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})
}
