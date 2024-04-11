package service

import (
	"github.com/stretchr/testify/assert"
	. "micro-batching/api"
	"testing"
)

func TestTake(t *testing.T) {
	var b Batching

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