/*
Camunda Platform REST API

Testing HistoricTaskInstanceAPIService

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

func Test_camundarestgo_HistoricTaskInstanceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HistoricTaskInstanceAPIService GetHistoricTaskInstanceReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HistoricTaskInstanceAPI.GetHistoricTaskInstanceReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HistoricTaskInstanceAPIService GetHistoricTaskInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HistoricTaskInstanceAPI.GetHistoricTaskInstances(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HistoricTaskInstanceAPIService GetHistoricTaskInstancesCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HistoricTaskInstanceAPI.GetHistoricTaskInstancesCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HistoricTaskInstanceAPIService QueryHistoricTaskInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HistoricTaskInstanceAPI.QueryHistoricTaskInstances(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HistoricTaskInstanceAPIService QueryHistoricTaskInstancesCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HistoricTaskInstanceAPI.QueryHistoricTaskInstancesCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
