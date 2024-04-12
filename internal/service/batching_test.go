package service

import (
	"github.com/stretchr/testify/assert"
	. "micro-batching/api"
	"micro-batching/internal/config"
	"testing"
)

func TestTake(t *testing.T) {
	b := NewBatchingWithConfig(config.RunConfig{})

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
	b := NewBatchingWithConfig(config.RunConfig{})

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
