# \HistoricDecisionInstanceAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAsync**](HistoricDecisionInstanceAPI.md#DeleteAsync) | **Post** /history/decision-instance/delete | Delete Async (POST)
[**GetHistoricDecisionInstance**](HistoricDecisionInstanceAPI.md#GetHistoricDecisionInstance) | **Get** /history/decision-instance/{id} | Get Historic Decision Instance
[**GetHistoricDecisionInstances**](HistoricDecisionInstanceAPI.md#GetHistoricDecisionInstances) | **Get** /history/decision-instance | Get Historic Decision Instances
[**GetHistoricDecisionInstancesCount**](HistoricDecisionInstanceAPI.md#GetHistoricDecisionInstancesCount) | **Get** /history/decision-instance/count | Get Historic Decision Instance Count
[**SetRemovalTimeAsyncHistoricDecisionInstance**](HistoricDecisionInstanceAPI.md#SetRemovalTimeAsyncHistoricDecisionInstance) | **Post** /history/decision-instance/set-removal-time | Set Removal Time Async (POST)



## DeleteAsync

> BatchDto DeleteAsync(ctx).DeleteHistoricDecisionInstancesDto(deleteHistoricDecisionInstancesDto).Execute()

Delete Async (POST)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	deleteHistoricDecisionInstancesDto := *openapiclient.NewDeleteHistoricDecisionInstancesDto() // DeleteHistoricDecisionInstancesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDecisionInstanceAPI.DeleteAsync(context.Background()).DeleteHistoricDecisionInstancesDto(deleteHistoricDecisionInstancesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDecisionInstanceAPI.DeleteAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAsync`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDecisionInstanceAPI.DeleteAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteHistoricDecisionInstancesDto** | [**DeleteHistoricDecisionInstancesDto**](DeleteHistoricDecisionInstancesDto.md) |  | 

### Return type

[**BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricDecisionInstance

> HistoricDecisionInstanceDto GetHistoricDecisionInstance(ctx, id).IncludeInputs(includeInputs).IncludeOutputs(includeOutputs).DisableBinaryFetching(disableBinaryFetching).DisableCustomObjectDeserialization(disableCustomObjectDeserialization).Execute()

Get Historic Decision Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of the historic decision instance to be retrieved.
	includeInputs := true // bool | Include input values in the result. Value may only be `true`, as `false` is the default behavior. (optional)
	includeOutputs := true // bool | Include output values in the result. Value may only be `true`, as `false` is the default behavior. (optional)
	disableBinaryFetching := true // bool | Disables fetching of byte array input and output values. Value may only be `true`, as `false` is the default behavior. (optional)
	disableCustomObjectDeserialization := true // bool | Disables deserialization of input and output values that are custom objects. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDecisionInstanceAPI.GetHistoricDecisionInstance(context.Background(), id).IncludeInputs(includeInputs).IncludeOutputs(includeOutputs).DisableBinaryFetching(disableBinaryFetching).DisableCustomObjectDeserialization(disableCustomObjectDeserialization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDecisionInstanceAPI.GetHistoricDecisionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricDecisionInstance`: HistoricDecisionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDecisionInstanceAPI.GetHistoricDecisionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the historic decision instance to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricDecisionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeInputs** | **bool** | Include input values in the result. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeOutputs** | **bool** | Include output values in the result. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **disableBinaryFetching** | **bool** | Disables fetching of byte array input and output values. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **disableCustomObjectDeserialization** | **bool** | Disables deserialization of input and output values that are custom objects. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

### Return type

[**HistoricDecisionInstanceDto**](HistoricDecisionInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricDecisionInstances

> []HistoricDecisionInstanceDto GetHistoricDecisionInstances(ctx).DecisionInstanceId(decisionInstanceId).DecisionInstanceIdIn(decisionInstanceIdIn).DecisionDefinitionId(decisionDefinitionId).DecisionDefinitionIdIn(decisionDefinitionIdIn).DecisionDefinitionKey(decisionDefinitionKey).DecisionDefinitionKeyIn(decisionDefinitionKeyIn).DecisionDefinitionName(decisionDefinitionName).DecisionDefinitionNameLike(decisionDefinitionNameLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseInstanceId(caseInstanceId).ActivityIdIn(activityIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).EvaluatedBefore(evaluatedBefore).EvaluatedAfter(evaluatedAfter).UserId(userId).RootDecisionInstanceId(rootDecisionInstanceId).RootDecisionInstancesOnly(rootDecisionInstancesOnly).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionKey(decisionRequirementsDefinitionKey).IncludeInputs(includeInputs).IncludeOutputs(includeOutputs).DisableBinaryFetching(disableBinaryFetching).DisableCustomObjectDeserialization(disableCustomObjectDeserialization).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Historic Decision Instances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	decisionInstanceId := "decisionInstanceId_example" // string | Filter by decision instance id. (optional)
	decisionInstanceIdIn := "decisionInstanceIdIn_example" // string | Filter by decision instance ids. Must be a comma-separated list of decision instance ids. (optional)
	decisionDefinitionId := "decisionDefinitionId_example" // string | Filter by the decision definition the instances belongs to. (optional)
	decisionDefinitionIdIn := "decisionDefinitionIdIn_example" // string | Filter by the decision definitions the instances belongs to. Must be a comma-separated list of decision definition ids. (optional)
	decisionDefinitionKey := "decisionDefinitionKey_example" // string | Filter by the key of the decision definition the instances belongs to. (optional)
	decisionDefinitionKeyIn := "decisionDefinitionKeyIn_example" // string | Filter by the keys of the decision definition the instances belongs to. Must be a comma- separated list of decision definition keys. (optional)
	decisionDefinitionName := "decisionDefinitionName_example" // string | Filter by the name of the decision definition the instances belongs to. (optional)
	decisionDefinitionNameLike := "decisionDefinitionNameLike_example" // string | Filter by the name of the decision definition the instances belongs to, that the parameter is a substring of. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the process definition the instances belongs to. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by the key of the process definition the instances belongs to. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by the process instance the instances belongs to. (optional)
	caseDefinitionId := "caseDefinitionId_example" // string | Filter by the case definition the instances belongs to. (optional)
	caseDefinitionKey := "caseDefinitionKey_example" // string | Filter by the key of the case definition the instances belongs to. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by the case instance the instances belongs to. (optional)
	activityIdIn := "activityIdIn_example" // string | Filter by the activity ids the instances belongs to. Must be a comma-separated list of acitvity ids. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Filter by the activity instance ids the instances belongs to. Must be a comma-separated list of acitvity instance ids. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A historic decision instance must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic decision instances that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	evaluatedBefore := time.Now() // time.Time | Restrict to instances that were evaluated before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM- dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	evaluatedAfter := time.Now() // time.Time | Restrict to instances that were evaluated after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM- dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	userId := "userId_example" // string | Restrict to instances that were evaluated by the given user. (optional)
	rootDecisionInstanceId := "rootDecisionInstanceId_example" // string | Restrict to instances that have a given root decision instance id. This also includes the decision instance with the given id. (optional)
	rootDecisionInstancesOnly := true // bool | Restrict to instances those are the root decision instance of an evaluation. Value may only be `true`, as `false` is the default behavior. (optional)
	decisionRequirementsDefinitionId := "decisionRequirementsDefinitionId_example" // string | Filter by the decision requirements definition the instances belongs to. (optional)
	decisionRequirementsDefinitionKey := "decisionRequirementsDefinitionKey_example" // string | Filter by the key of the decision requirements definition the instances belongs to. (optional)
	includeInputs := true // bool | Include input values in the result. Value may only be `true`, as `false` is the default behavior. (optional)
	includeOutputs := true // bool | Include output values in the result. Value may only be `true`, as `false` is the default behavior. (optional)
	disableBinaryFetching := true // bool | Disables fetching of byte array input and output values. Value may only be `true`, as `false` is the default behavior. (optional)
	disableCustomObjectDeserialization := true // bool | Disables deserialization of input and output values that are custom objects. Value may only be `true`, as `false` is the default behavior. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDecisionInstanceAPI.GetHistoricDecisionInstances(context.Background()).DecisionInstanceId(decisionInstanceId).DecisionInstanceIdIn(decisionInstanceIdIn).DecisionDefinitionId(decisionDefinitionId).DecisionDefinitionIdIn(decisionDefinitionIdIn).DecisionDefinitionKey(decisionDefinitionKey).DecisionDefinitionKeyIn(decisionDefinitionKeyIn).DecisionDefinitionName(decisionDefinitionName).DecisionDefinitionNameLike(decisionDefinitionNameLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseInstanceId(caseInstanceId).ActivityIdIn(activityIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).EvaluatedBefore(evaluatedBefore).EvaluatedAfter(evaluatedAfter).UserId(userId).RootDecisionInstanceId(rootDecisionInstanceId).RootDecisionInstancesOnly(rootDecisionInstancesOnly).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionKey(decisionRequirementsDefinitionKey).IncludeInputs(includeInputs).IncludeOutputs(includeOutputs).DisableBinaryFetching(disableBinaryFetching).DisableCustomObjectDeserialization(disableCustomObjectDeserialization).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDecisionInstanceAPI.GetHistoricDecisionInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricDecisionInstances`: []HistoricDecisionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDecisionInstanceAPI.GetHistoricDecisionInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricDecisionInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decisionInstanceId** | **string** | Filter by decision instance id. | 
 **decisionInstanceIdIn** | **string** | Filter by decision instance ids. Must be a comma-separated list of decision instance ids. | 
 **decisionDefinitionId** | **string** | Filter by the decision definition the instances belongs to. | 
 **decisionDefinitionIdIn** | **string** | Filter by the decision definitions the instances belongs to. Must be a comma-separated list of decision definition ids. | 
 **decisionDefinitionKey** | **string** | Filter by the key of the decision definition the instances belongs to. | 
 **decisionDefinitionKeyIn** | **string** | Filter by the keys of the decision definition the instances belongs to. Must be a comma- separated list of decision definition keys. | 
 **decisionDefinitionName** | **string** | Filter by the name of the decision definition the instances belongs to. | 
 **decisionDefinitionNameLike** | **string** | Filter by the name of the decision definition the instances belongs to, that the parameter is a substring of. | 
 **processDefinitionId** | **string** | Filter by the process definition the instances belongs to. | 
 **processDefinitionKey** | **string** | Filter by the key of the process definition the instances belongs to. | 
 **processInstanceId** | **string** | Filter by the process instance the instances belongs to. | 
 **caseDefinitionId** | **string** | Filter by the case definition the instances belongs to. | 
 **caseDefinitionKey** | **string** | Filter by the key of the case definition the instances belongs to. | 
 **caseInstanceId** | **string** | Filter by the case instance the instances belongs to. | 
 **activityIdIn** | **string** | Filter by the activity ids the instances belongs to. Must be a comma-separated list of acitvity ids. | 
 **activityInstanceIdIn** | **string** | Filter by the activity instance ids the instances belongs to. Must be a comma-separated list of acitvity instance ids. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A historic decision instance must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic decision instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **evaluatedBefore** | **time.Time** | Restrict to instances that were evaluated before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM- dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **evaluatedAfter** | **time.Time** | Restrict to instances that were evaluated after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM- dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **userId** | **string** | Restrict to instances that were evaluated by the given user. | 
 **rootDecisionInstanceId** | **string** | Restrict to instances that have a given root decision instance id. This also includes the decision instance with the given id. | 
 **rootDecisionInstancesOnly** | **bool** | Restrict to instances those are the root decision instance of an evaluation. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **decisionRequirementsDefinitionId** | **string** | Filter by the decision requirements definition the instances belongs to. | 
 **decisionRequirementsDefinitionKey** | **string** | Filter by the key of the decision requirements definition the instances belongs to. | 
 **includeInputs** | **bool** | Include input values in the result. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeOutputs** | **bool** | Include output values in the result. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **disableBinaryFetching** | **bool** | Disables fetching of byte array input and output values. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **disableCustomObjectDeserialization** | **bool** | Disables deserialization of input and output values that are custom objects. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]HistoricDecisionInstanceDto**](HistoricDecisionInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricDecisionInstancesCount

> CountResultDto GetHistoricDecisionInstancesCount(ctx).DecisionInstanceId(decisionInstanceId).DecisionInstanceIdIn(decisionInstanceIdIn).DecisionDefinitionId(decisionDefinitionId).DecisionDefinitionIdIn(decisionDefinitionIdIn).DecisionDefinitionKey(decisionDefinitionKey).DecisionDefinitionKeyIn(decisionDefinitionKeyIn).DecisionDefinitionName(decisionDefinitionName).DecisionDefinitionNameLike(decisionDefinitionNameLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseInstanceId(caseInstanceId).ActivityIdIn(activityIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).EvaluatedBefore(evaluatedBefore).EvaluatedAfter(evaluatedAfter).UserId(userId).RootDecisionInstanceId(rootDecisionInstanceId).RootDecisionInstancesOnly(rootDecisionInstancesOnly).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionKey(decisionRequirementsDefinitionKey).Execute()

Get Historic Decision Instance Count



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	decisionInstanceId := "decisionInstanceId_example" // string | Filter by decision instance id. (optional)
	decisionInstanceIdIn := "decisionInstanceIdIn_example" // string | Filter by decision instance ids. Must be a comma-separated list of decision instance ids. (optional)
	decisionDefinitionId := "decisionDefinitionId_example" // string | Filter by the decision definition the instances belongs to. (optional)
	decisionDefinitionIdIn := "decisionDefinitionIdIn_example" // string | Filter by the decision definitions the instances belongs to. Must be a comma-separated list of decision definition ids. (optional)
	decisionDefinitionKey := "decisionDefinitionKey_example" // string | Filter by the key of the decision definition the instances belongs to. (optional)
	decisionDefinitionKeyIn := "decisionDefinitionKeyIn_example" // string | Filter by the keys of the decision definition the instances belongs to. Must be a comma- separated list of decision definition keys. (optional)
	decisionDefinitionName := "decisionDefinitionName_example" // string | Filter by the name of the decision definition the instances belongs to. (optional)
	decisionDefinitionNameLike := "decisionDefinitionNameLike_example" // string | Filter by the name of the decision definition the instances belongs to, that the parameter is a substring of. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the process definition the instances belongs to. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by the key of the process definition the instances belongs to. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by the process instance the instances belongs to. (optional)
	caseDefinitionId := "caseDefinitionId_example" // string | Filter by the case definition the instances belongs to. (optional)
	caseDefinitionKey := "caseDefinitionKey_example" // string | Filter by the key of the case definition the instances belongs to. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by the case instance the instances belongs to. (optional)
	activityIdIn := "activityIdIn_example" // string | Filter by the activity ids the instances belongs to. Must be a comma-separated list of acitvity ids. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Filter by the activity instance ids the instances belongs to. Must be a comma-separated list of acitvity instance ids. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A historic decision instance must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic decision instances that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	evaluatedBefore := time.Now() // time.Time | Restrict to instances that were evaluated before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM- dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	evaluatedAfter := time.Now() // time.Time | Restrict to instances that were evaluated after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM- dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	userId := "userId_example" // string | Restrict to instances that were evaluated by the given user. (optional)
	rootDecisionInstanceId := "rootDecisionInstanceId_example" // string | Restrict to instances that have a given root decision instance id. This also includes the decision instance with the given id. (optional)
	rootDecisionInstancesOnly := true // bool | Restrict to instances those are the root decision instance of an evaluation. Value may only be `true`, as `false` is the default behavior. (optional)
	decisionRequirementsDefinitionId := "decisionRequirementsDefinitionId_example" // string | Filter by the decision requirements definition the instances belongs to. (optional)
	decisionRequirementsDefinitionKey := "decisionRequirementsDefinitionKey_example" // string | Filter by the key of the decision requirements definition the instances belongs to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDecisionInstanceAPI.GetHistoricDecisionInstancesCount(context.Background()).DecisionInstanceId(decisionInstanceId).DecisionInstanceIdIn(decisionInstanceIdIn).DecisionDefinitionId(decisionDefinitionId).DecisionDefinitionIdIn(decisionDefinitionIdIn).DecisionDefinitionKey(decisionDefinitionKey).DecisionDefinitionKeyIn(decisionDefinitionKeyIn).DecisionDefinitionName(decisionDefinitionName).DecisionDefinitionNameLike(decisionDefinitionNameLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseInstanceId(caseInstanceId).ActivityIdIn(activityIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).EvaluatedBefore(evaluatedBefore).EvaluatedAfter(evaluatedAfter).UserId(userId).RootDecisionInstanceId(rootDecisionInstanceId).RootDecisionInstancesOnly(rootDecisionInstancesOnly).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionKey(decisionRequirementsDefinitionKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDecisionInstanceAPI.GetHistoricDecisionInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricDecisionInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDecisionInstanceAPI.GetHistoricDecisionInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricDecisionInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decisionInstanceId** | **string** | Filter by decision instance id. | 
 **decisionInstanceIdIn** | **string** | Filter by decision instance ids. Must be a comma-separated list of decision instance ids. | 
 **decisionDefinitionId** | **string** | Filter by the decision definition the instances belongs to. | 
 **decisionDefinitionIdIn** | **string** | Filter by the decision definitions the instances belongs to. Must be a comma-separated list of decision definition ids. | 
 **decisionDefinitionKey** | **string** | Filter by the key of the decision definition the instances belongs to. | 
 **decisionDefinitionKeyIn** | **string** | Filter by the keys of the decision definition the instances belongs to. Must be a comma- separated list of decision definition keys. | 
 **decisionDefinitionName** | **string** | Filter by the name of the decision definition the instances belongs to. | 
 **decisionDefinitionNameLike** | **string** | Filter by the name of the decision definition the instances belongs to, that the parameter is a substring of. | 
 **processDefinitionId** | **string** | Filter by the process definition the instances belongs to. | 
 **processDefinitionKey** | **string** | Filter by the key of the process definition the instances belongs to. | 
 **processInstanceId** | **string** | Filter by the process instance the instances belongs to. | 
 **caseDefinitionId** | **string** | Filter by the case definition the instances belongs to. | 
 **caseDefinitionKey** | **string** | Filter by the key of the case definition the instances belongs to. | 
 **caseInstanceId** | **string** | Filter by the case instance the instances belongs to. | 
 **activityIdIn** | **string** | Filter by the activity ids the instances belongs to. Must be a comma-separated list of acitvity ids. | 
 **activityInstanceIdIn** | **string** | Filter by the activity instance ids the instances belongs to. Must be a comma-separated list of acitvity instance ids. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A historic decision instance must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic decision instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **evaluatedBefore** | **time.Time** | Restrict to instances that were evaluated before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM- dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **evaluatedAfter** | **time.Time** | Restrict to instances that were evaluated after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM- dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **userId** | **string** | Restrict to instances that were evaluated by the given user. | 
 **rootDecisionInstanceId** | **string** | Restrict to instances that have a given root decision instance id. This also includes the decision instance with the given id. | 
 **rootDecisionInstancesOnly** | **bool** | Restrict to instances those are the root decision instance of an evaluation. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **decisionRequirementsDefinitionId** | **string** | Filter by the decision requirements definition the instances belongs to. | 
 **decisionRequirementsDefinitionKey** | **string** | Filter by the key of the decision requirements definition the instances belongs to. | 

### Return type

[**CountResultDto**](CountResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRemovalTimeAsyncHistoricDecisionInstance

> BatchDto SetRemovalTimeAsyncHistoricDecisionInstance(ctx).SetRemovalTimeToHistoricDecisionInstancesDto(setRemovalTimeToHistoricDecisionInstancesDto).Execute()

Set Removal Time Async (POST)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	setRemovalTimeToHistoricDecisionInstancesDto := *openapiclient.NewSetRemovalTimeToHistoricDecisionInstancesDto() // SetRemovalTimeToHistoricDecisionInstancesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDecisionInstanceAPI.SetRemovalTimeAsyncHistoricDecisionInstance(context.Background()).SetRemovalTimeToHistoricDecisionInstancesDto(setRemovalTimeToHistoricDecisionInstancesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDecisionInstanceAPI.SetRemovalTimeAsyncHistoricDecisionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRemovalTimeAsyncHistoricDecisionInstance`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDecisionInstanceAPI.SetRemovalTimeAsyncHistoricDecisionInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRemovalTimeAsyncHistoricDecisionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setRemovalTimeToHistoricDecisionInstancesDto** | [**SetRemovalTimeToHistoricDecisionInstancesDto**](SetRemovalTimeToHistoricDecisionInstancesDto.md) |  | 

### Return type

[**BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

