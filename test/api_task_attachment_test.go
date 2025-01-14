/*
Camunda Platform REST API

Testing TaskAttachmentAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package camundarestgo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_camundarestgo_TaskAttachmentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaskAttachmentAPIService AddAttachment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TaskAttachmentAPI.AddAttachment(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskAttachmentAPIService DeleteAttachment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var attachmentId string

		httpRes, err := apiClient.TaskAttachmentAPI.DeleteAttachment(context.Background(), id, attachmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskAttachmentAPIService GetAttachment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var attachmentId string

		resp, httpRes, err := apiClient.TaskAttachmentAPI.GetAttachment(context.Background(), id, attachmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskAttachmentAPIService GetAttachmentData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var attachmentId string

		resp, httpRes, err := apiClient.TaskAttachmentAPI.GetAttachmentData(context.Background(), id, attachmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskAttachmentAPIService GetAttachments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TaskAttachmentAPI.GetAttachments(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
