# \HistoricActivityInstanceAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricActivityInstance**](HistoricActivityInstanceAPI.md#GetHistoricActivityInstance) | **Get** /history/activity-instance/{id} | Get
[**GetHistoricActivityInstances**](HistoricActivityInstanceAPI.md#GetHistoricActivityInstances) | **Get** /history/activity-instance | Get List
[**GetHistoricActivityInstancesCount**](HistoricActivityInstanceAPI.md#GetHistoricActivityInstancesCount) | **Get** /history/activity-instance/count | Get List Count
[**QueryHistoricActivityInstances**](HistoricActivityInstanceAPI.md#QueryHistoricActivityInstances) | **Post** /history/activity-instance | Get List (POST)
[**QueryHistoricActivityInstancesCount**](HistoricActivityInstanceAPI.md#QueryHistoricActivityInstancesCount) | **Post** /history/activity-instance/count | Get List Count (POST)



## GetHistoricActivityInstance

> HistoricActivityInstanceDto GetHistoricActivityInstance(ctx, id).Execute()

Get



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
	id := "id_example" // string | The id of the historic activity instance to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricActivityInstanceAPI.GetHistoricActivityInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricActivityInstanceAPI.GetHistoricActivityInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricActivityInstance`: HistoricActivityInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricActivityInstanceAPI.GetHistoricActivityInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the historic activity instance to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricActivityInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HistoricActivityInstanceDto**](HistoricActivityInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricActivityInstances

> []HistoricActivityInstanceDto GetHistoricActivityInstances(ctx).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).ActivityInstanceId(activityInstanceId).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ExecutionId(executionId).ActivityId(activityId).ActivityName(activityName).ActivityType(activityType).TaskAssignee(taskAssignee).Finished(finished).Unfinished(unfinished).Canceled(canceled).CompleteScope(completeScope).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Execute()

Get List



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
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)
	activityInstanceId := "activityInstanceId_example" // string | Filter by activity instance id. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	executionId := "executionId_example" // string | Filter by the id of the execution that executed the activity instance. (optional)
	activityId := "activityId_example" // string | Filter by the activity id (according to BPMN 2.0 XML). (optional)
	activityName := "activityName_example" // string | Filter by the activity name (according to BPMN 2.0 XML). (optional)
	activityType := "activityType_example" // string | Filter by activity type. (optional)
	taskAssignee := "taskAssignee_example" // string | Only include activity instances that are user tasks and assigned to a given user. (optional)
	finished := true // bool | Only include finished activity instances. Value may only be `true`, as `false` behaves the same as when the property is not set. (optional)
	unfinished := true // bool | Only include unfinished activity instances. Value may only be `true`, as `false` behaves the same as when the property is not set. (optional)
	canceled := true // bool | Only include canceled activity instances. Value may only be `true`, as `false` behaves the same as when the property is not set. (optional)
	completeScope := true // bool | Only include activity instances which completed a scope. Value may only be `true`, as `false` behaves the same as when the property is not set. (optional)
	startedBefore := time.Now() // time.Time | Restrict to instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedBefore := time.Now() // time.Time | Restrict to instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedAfter := time.Now() // time.Time | Restrict to instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of ids. An activity instance must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic activity instances that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricActivityInstanceAPI.GetHistoricActivityInstances(context.Background()).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).ActivityInstanceId(activityInstanceId).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ExecutionId(executionId).ActivityId(activityId).ActivityName(activityName).ActivityType(activityType).TaskAssignee(taskAssignee).Finished(finished).Unfinished(unfinished).Canceled(canceled).CompleteScope(completeScope).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricActivityInstanceAPI.GetHistoricActivityInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricActivityInstances`: []HistoricActivityInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricActivityInstanceAPI.GetHistoricActivityInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricActivityInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **activityInstanceId** | **string** | Filter by activity instance id. | 
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **executionId** | **string** | Filter by the id of the execution that executed the activity instance. | 
 **activityId** | **string** | Filter by the activity id (according to BPMN 2.0 XML). | 
 **activityName** | **string** | Filter by the activity name (according to BPMN 2.0 XML). | 
 **activityType** | **string** | Filter by activity type. | 
 **taskAssignee** | **string** | Only include activity instances that are user tasks and assigned to a given user. | 
 **finished** | **bool** | Only include finished activity instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | 
 **unfinished** | **bool** | Only include unfinished activity instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | 
 **canceled** | **bool** | Only include canceled activity instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | 
 **completeScope** | **bool** | Only include activity instances which completed a scope. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | 
 **startedBefore** | **time.Time** | Restrict to instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedBefore** | **time.Time** | Restrict to instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedAfter** | **time.Time** | Restrict to instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of ids. An activity instance must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic activity instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

### Return type

[**[]HistoricActivityInstanceDto**](HistoricActivityInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricActivityInstancesCount

> CountResultDto GetHistoricActivityInstancesCount(ctx).ActivityInstanceId(activityInstanceId).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ExecutionId(executionId).ActivityId(activityId).ActivityName(activityName).ActivityType(activityType).TaskAssignee(taskAssignee).Finished(finished).Unfinished(unfinished).Canceled(canceled).CompleteScope(completeScope).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Execute()

Get List Count



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
	activityInstanceId := "activityInstanceId_example" // string | Filter by activity instance id. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	executionId := "executionId_example" // string | Filter by the id of the execution that executed the activity instance. (optional)
	activityId := "activityId_example" // string | Filter by the activity id (according to BPMN 2.0 XML). (optional)
	activityName := "activityName_example" // string | Filter by the activity name (according to BPMN 2.0 XML). (optional)
	activityType := "activityType_example" // string | Filter by activity type. (optional)
	taskAssignee := "taskAssignee_example" // string | Only include activity instances that are user tasks and assigned to a given user. (optional)
	finished := true // bool | Only include finished activity instances. Value may only be `true`, as `false` behaves the same as when the property is not set. (optional)
	unfinished := true // bool | Only include unfinished activity instances. Value may only be `true`, as `false` behaves the same as when the property is not set. (optional)
	canceled := true // bool | Only include canceled activity instances. Value may only be `true`, as `false` behaves the same as when the property is not set. (optional)
	completeScope := true // bool | Only include activity instances which completed a scope. Value may only be `true`, as `false` behaves the same as when the property is not set. (optional)
	startedBefore := time.Now() // time.Time | Restrict to instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedBefore := time.Now() // time.Time | Restrict to instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedAfter := time.Now() // time.Time | Restrict to instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of ids. An activity instance must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic activity instances that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricActivityInstanceAPI.GetHistoricActivityInstancesCount(context.Background()).ActivityInstanceId(activityInstanceId).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ExecutionId(executionId).ActivityId(activityId).ActivityName(activityName).ActivityType(activityType).TaskAssignee(taskAssignee).Finished(finished).Unfinished(unfinished).Canceled(canceled).CompleteScope(completeScope).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricActivityInstanceAPI.GetHistoricActivityInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricActivityInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricActivityInstanceAPI.GetHistoricActivityInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricActivityInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityInstanceId** | **string** | Filter by activity instance id. | 
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **executionId** | **string** | Filter by the id of the execution that executed the activity instance. | 
 **activityId** | **string** | Filter by the activity id (according to BPMN 2.0 XML). | 
 **activityName** | **string** | Filter by the activity name (according to BPMN 2.0 XML). | 
 **activityType** | **string** | Filter by activity type. | 
 **taskAssignee** | **string** | Only include activity instances that are user tasks and assigned to a given user. | 
 **finished** | **bool** | Only include finished activity instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | 
 **unfinished** | **bool** | Only include unfinished activity instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | 
 **canceled** | **bool** | Only include canceled activity instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | 
 **completeScope** | **bool** | Only include activity instances which completed a scope. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | 
 **startedBefore** | **time.Time** | Restrict to instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedBefore** | **time.Time** | Restrict to instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedAfter** | **time.Time** | Restrict to instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of ids. An activity instance must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic activity instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

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


## QueryHistoricActivityInstances

> []HistoricActivityInstanceDto QueryHistoricActivityInstances(ctx).FirstResult(firstResult).MaxResults(maxResults).HistoricActivityInstanceQueryDto(historicActivityInstanceQueryDto).Execute()

Get List (POST)



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
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)
	historicActivityInstanceQueryDto := *openapiclient.NewHistoricActivityInstanceQueryDto() // HistoricActivityInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricActivityInstanceAPI.QueryHistoricActivityInstances(context.Background()).FirstResult(firstResult).MaxResults(maxResults).HistoricActivityInstanceQueryDto(historicActivityInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricActivityInstanceAPI.QueryHistoricActivityInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricActivityInstances`: []HistoricActivityInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricActivityInstanceAPI.QueryHistoricActivityInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricActivityInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **historicActivityInstanceQueryDto** | [**HistoricActivityInstanceQueryDto**](HistoricActivityInstanceQueryDto.md) |  | 

### Return type

[**[]HistoricActivityInstanceDto**](HistoricActivityInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryHistoricActivityInstancesCount

> CountResultDto QueryHistoricActivityInstancesCount(ctx).HistoricActivityInstanceQueryDto(historicActivityInstanceQueryDto).Execute()

Get List Count (POST)



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
	historicActivityInstanceQueryDto := *openapiclient.NewHistoricActivityInstanceQueryDto() // HistoricActivityInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricActivityInstanceAPI.QueryHistoricActivityInstancesCount(context.Background()).HistoricActivityInstanceQueryDto(historicActivityInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricActivityInstanceAPI.QueryHistoricActivityInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricActivityInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricActivityInstanceAPI.QueryHistoricActivityInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricActivityInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **historicActivityInstanceQueryDto** | [**HistoricActivityInstanceQueryDto**](HistoricActivityInstanceQueryDto.md) |  | 

### Return type

[**CountResultDto**](CountResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

