/*
Camunda Platform REST API

Testing DecisionDefinitionAPIService

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

func Test_camundarestgo_DecisionDefinitionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DecisionDefinitionAPIService EvaluateDecisionById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.EvaluateDecisionById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService EvaluateDecisionByKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.EvaluateDecisionByKey(context.Background(), key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService EvaluateDecisionByKeyAndTenant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string
		var tenantId string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.EvaluateDecisionByKeyAndTenant(context.Background(), key, tenantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService GetDecisionDefinitionById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService GetDecisionDefinitionByKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionByKey(context.Background(), key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService GetDecisionDefinitionByKeyAndTenantId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string
		var tenantId string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionByKeyAndTenantId(context.Background(), key, tenantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService GetDecisionDefinitionDiagram", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDiagram(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService GetDecisionDefinitionDiagramByKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDiagramByKey(context.Background(), key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService GetDecisionDefinitionDiagramByKeyAndTenant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string
		var tenantId string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDiagramByKeyAndTenant(context.Background(), key, tenantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService GetDecisionDefinitionDmnXmlById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService GetDecisionDefinitionDmnXmlByKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlByKey(context.Background(), key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService GetDecisionDefinitionDmnXmlByKeyAndTenant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string
		var tenantId string

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlByKeyAndTenant(context.Background(), key, tenantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService GetDecisionDefinitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService GetDecisionDefinitionsCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionsCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService UpdateHistoryTimeToLiveByDecisionDefinitionId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.DecisionDefinitionAPI.UpdateHistoryTimeToLiveByDecisionDefinitionId(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService UpdateHistoryTimeToLiveByDecisionDefinitionKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string

		httpRes, err := apiClient.DecisionDefinitionAPI.UpdateHistoryTimeToLiveByDecisionDefinitionKey(context.Background(), key).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DecisionDefinitionAPIService UpdateHistoryTimeToLiveByDecisionDefinitionKeyAndTenant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string
		var tenantId string

		httpRes, err := apiClient.DecisionDefinitionAPI.UpdateHistoryTimeToLiveByDecisionDefinitionKeyAndTenant(context.Background(), key, tenantId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
