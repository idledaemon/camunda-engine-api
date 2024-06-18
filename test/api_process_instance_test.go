/*
Camunda Platform REST API

Testing ProcessInstanceAPIService

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

func Test_camundarestgo_ProcessInstanceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProcessInstanceAPIService CorrelateMessageAsyncOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProcessInstanceAPI.CorrelateMessageAsyncOperation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService DeleteAsyncHistoricQueryBased", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProcessInstanceAPI.DeleteAsyncHistoricQueryBased(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService DeleteProcessInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ProcessInstanceAPI.DeleteProcessInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService DeleteProcessInstanceVariable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var varName string

		httpRes, err := apiClient.ProcessInstanceAPI.DeleteProcessInstanceVariable(context.Background(), id, varName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService DeleteProcessInstancesAsyncOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProcessInstanceAPI.DeleteProcessInstancesAsyncOperation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService GetActivityInstanceTree", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ProcessInstanceAPI.GetActivityInstanceTree(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService GetProcessInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ProcessInstanceAPI.GetProcessInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService GetProcessInstanceComments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ProcessInstanceAPI.GetProcessInstanceComments(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService GetProcessInstanceVariable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var varName string

		resp, httpRes, err := apiClient.ProcessInstanceAPI.GetProcessInstanceVariable(context.Background(), id, varName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService GetProcessInstanceVariableBinary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var varName string

		resp, httpRes, err := apiClient.ProcessInstanceAPI.GetProcessInstanceVariableBinary(context.Background(), id, varName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService GetProcessInstanceVariables", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ProcessInstanceAPI.GetProcessInstanceVariables(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService GetProcessInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProcessInstanceAPI.GetProcessInstances(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService GetProcessInstancesCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProcessInstanceAPI.GetProcessInstancesCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService ModifyProcessInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ProcessInstanceAPI.ModifyProcessInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService ModifyProcessInstanceAsyncOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ProcessInstanceAPI.ModifyProcessInstanceAsyncOperation(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService ModifyProcessInstanceVariables", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ProcessInstanceAPI.ModifyProcessInstanceVariables(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService QueryProcessInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProcessInstanceAPI.QueryProcessInstances(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService QueryProcessInstancesCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProcessInstanceAPI.QueryProcessInstancesCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService SetProcessInstanceVariable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var varName string

		httpRes, err := apiClient.ProcessInstanceAPI.SetProcessInstanceVariable(context.Background(), id, varName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService SetProcessInstanceVariableBinary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var varName string

		httpRes, err := apiClient.ProcessInstanceAPI.SetProcessInstanceVariableBinary(context.Background(), id, varName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService SetRetriesByProcess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProcessInstanceAPI.SetRetriesByProcess(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService SetRetriesByProcessHistoricQueryBased", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProcessInstanceAPI.SetRetriesByProcessHistoricQueryBased(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService SetVariablesAsyncOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProcessInstanceAPI.SetVariablesAsyncOperation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService UpdateSuspensionState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProcessInstanceAPI.UpdateSuspensionState(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService UpdateSuspensionStateAsyncOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProcessInstanceAPI.UpdateSuspensionStateAsyncOperation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProcessInstanceAPIService UpdateSuspensionStateById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ProcessInstanceAPI.UpdateSuspensionStateById(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
