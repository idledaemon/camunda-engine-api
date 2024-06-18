/*
Camunda Platform REST API

Testing HistoricUserOperationLogAPIService

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

func Test_camundarestgo_HistoricUserOperationLogAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HistoricUserOperationLogAPIService ClearAnnotationUserOperationLog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		httpRes, err := apiClient.HistoricUserOperationLogAPI.ClearAnnotationUserOperationLog(context.Background(), operationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HistoricUserOperationLogAPIService QueryUserOperationCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HistoricUserOperationLogAPI.QueryUserOperationCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HistoricUserOperationLogAPIService QueryUserOperationEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HistoricUserOperationLogAPI.QueryUserOperationEntries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HistoricUserOperationLogAPIService SetAnnotationUserOperationLog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var operationId string

		httpRes, err := apiClient.HistoricUserOperationLogAPI.SetAnnotationUserOperationLog(context.Background(), operationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
