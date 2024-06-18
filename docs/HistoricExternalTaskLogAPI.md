# \HistoricExternalTaskLogAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetErrorDetailsHistoricExternalTaskLog**](HistoricExternalTaskLogAPI.md#GetErrorDetailsHistoricExternalTaskLog) | **Get** /history/external-task-log/{id}/error-details | Get External Task Log Error Details
[**GetHistoricExternalTaskLog**](HistoricExternalTaskLogAPI.md#GetHistoricExternalTaskLog) | **Get** /history/external-task-log/{id} | Get External Task Log
[**GetHistoricExternalTaskLogs**](HistoricExternalTaskLogAPI.md#GetHistoricExternalTaskLogs) | **Get** /history/external-task-log | Get External Task Logs
[**GetHistoricExternalTaskLogsCount**](HistoricExternalTaskLogAPI.md#GetHistoricExternalTaskLogsCount) | **Get** /history/external-task-log/count | Get External Task Log Count
[**QueryHistoricExternalTaskLogs**](HistoricExternalTaskLogAPI.md#QueryHistoricExternalTaskLogs) | **Post** /history/external-task-log | Get External Task Logs (POST)
[**QueryHistoricExternalTaskLogsCount**](HistoricExternalTaskLogAPI.md#QueryHistoricExternalTaskLogsCount) | **Post** /history/external-task-log/count | Get External Task Log Count (POST)



## GetErrorDetailsHistoricExternalTaskLog

> interface{} GetErrorDetailsHistoricExternalTaskLog(ctx, id).Execute()

Get External Task Log Error Details



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
	id := "id_example" // string | The id of the historic external task log to get the error details for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricExternalTaskLogAPI.GetErrorDetailsHistoricExternalTaskLog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricExternalTaskLogAPI.GetErrorDetailsHistoricExternalTaskLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetErrorDetailsHistoricExternalTaskLog`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `HistoricExternalTaskLogAPI.GetErrorDetailsHistoricExternalTaskLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the historic external task log to get the error details for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetErrorDetailsHistoricExternalTaskLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricExternalTaskLog

> HistoricExternalTaskLogDto GetHistoricExternalTaskLog(ctx, id).Execute()

Get External Task Log



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
	id := "id_example" // string | The id of the log entry.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricExternalTaskLogAPI.GetHistoricExternalTaskLog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricExternalTaskLogAPI.GetHistoricExternalTaskLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricExternalTaskLog`: HistoricExternalTaskLogDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricExternalTaskLogAPI.GetHistoricExternalTaskLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the log entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricExternalTaskLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HistoricExternalTaskLogDto**](HistoricExternalTaskLogDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricExternalTaskLogs

> []HistoricExternalTaskLogDto GetHistoricExternalTaskLogs(ctx).LogId(logId).ExternalTaskId(externalTaskId).TopicName(topicName).WorkerId(workerId).ErrorMessage(errorMessage).ActivityIdIn(activityIdIn).ActivityInstanceIdIn(activityInstanceIdIn).ExecutionIdIn(executionIdIn).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).CreationLog(creationLog).FailureLog(failureLog).SuccessLog(successLog).DeletionLog(deletionLog).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get External Task Logs



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
	logId := "logId_example" // string | Filter by historic external task log id. (optional)
	externalTaskId := "externalTaskId_example" // string | Filter by external task id. (optional)
	topicName := "topicName_example" // string | Filter by an external task topic. (optional)
	workerId := "workerId_example" // string | Filter by the id of the worker that the task was most recently locked by. (optional)
	errorMessage := "errorMessage_example" // string | Filter by external task exception message. (optional)
	activityIdIn := "activityIdIn_example" // string | Only include historic external task logs which belong to one of the passed activity ids. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Only include historic external task logs which belong to one of the passed activity instance ids. (optional)
	executionIdIn := "executionIdIn_example" // string | Only include historic external task logs which belong to one of the passed execution ids. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by process definition key. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include historic external task log entries which belong to one of the passed and comma-separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic external task log entries that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	priorityLowerThanOrEquals := int64(789) // int64 | Only include logs for which the associated external task had a priority lower than or equal to the given value. Value must be a valid `long` value. (optional)
	priorityHigherThanOrEquals := int64(789) // int64 | Only include logs for which the associated external task had a priority higher than or equal to the given value. Value must be a valid `long` value. (optional)
	creationLog := true // bool | Only include creation logs. Value may only be `true`, as `false` is the default behavior. (optional)
	failureLog := true // bool | Only include failure logs. Value may only be `true`, as `false` is the default behavior. (optional)
	successLog := true // bool | Only include success logs. Value may only be `true`, as `false` is the default behavior. (optional)
	deletionLog := true // bool | Only include deletion logs. Value may only be `true`, as `false` is the default behavior. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricExternalTaskLogAPI.GetHistoricExternalTaskLogs(context.Background()).LogId(logId).ExternalTaskId(externalTaskId).TopicName(topicName).WorkerId(workerId).ErrorMessage(errorMessage).ActivityIdIn(activityIdIn).ActivityInstanceIdIn(activityInstanceIdIn).ExecutionIdIn(executionIdIn).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).CreationLog(creationLog).FailureLog(failureLog).SuccessLog(successLog).DeletionLog(deletionLog).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricExternalTaskLogAPI.GetHistoricExternalTaskLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricExternalTaskLogs`: []HistoricExternalTaskLogDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricExternalTaskLogAPI.GetHistoricExternalTaskLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricExternalTaskLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logId** | **string** | Filter by historic external task log id. | 
 **externalTaskId** | **string** | Filter by external task id. | 
 **topicName** | **string** | Filter by an external task topic. | 
 **workerId** | **string** | Filter by the id of the worker that the task was most recently locked by. | 
 **errorMessage** | **string** | Filter by external task exception message. | 
 **activityIdIn** | **string** | Only include historic external task logs which belong to one of the passed activity ids. | 
 **activityInstanceIdIn** | **string** | Only include historic external task logs which belong to one of the passed activity instance ids. | 
 **executionIdIn** | **string** | Only include historic external task logs which belong to one of the passed execution ids. | 
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **processDefinitionKey** | **string** | Filter by process definition key. | 
 **tenantIdIn** | **string** | Only include historic external task log entries which belong to one of the passed and comma-separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic external task log entries that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **priorityLowerThanOrEquals** | **int64** | Only include logs for which the associated external task had a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **priorityHigherThanOrEquals** | **int64** | Only include logs for which the associated external task had a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **creationLog** | **bool** | Only include creation logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **failureLog** | **bool** | Only include failure logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **successLog** | **bool** | Only include success logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **deletionLog** | **bool** | Only include deletion logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]HistoricExternalTaskLogDto**](HistoricExternalTaskLogDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricExternalTaskLogsCount

> CountResultDto GetHistoricExternalTaskLogsCount(ctx).LogId(logId).ExternalTaskId(externalTaskId).TopicName(topicName).WorkerId(workerId).ErrorMessage(errorMessage).ActivityIdIn(activityIdIn).ActivityInstanceIdIn(activityInstanceIdIn).ExecutionIdIn(executionIdIn).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).CreationLog(creationLog).FailureLog(failureLog).SuccessLog(successLog).DeletionLog(deletionLog).Execute()

Get External Task Log Count



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
	logId := "logId_example" // string | Filter by historic external task log id. (optional)
	externalTaskId := "externalTaskId_example" // string | Filter by external task id. (optional)
	topicName := "topicName_example" // string | Filter by an external task topic. (optional)
	workerId := "workerId_example" // string | Filter by the id of the worker that the task was most recently locked by. (optional)
	errorMessage := "errorMessage_example" // string | Filter by external task exception message. (optional)
	activityIdIn := "activityIdIn_example" // string | Only include historic external task logs which belong to one of the passed activity ids. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Only include historic external task logs which belong to one of the passed activity instance ids. (optional)
	executionIdIn := "executionIdIn_example" // string | Only include historic external task logs which belong to one of the passed execution ids. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by process definition key. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include historic external task log entries which belong to one of the passed and comma-separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic external task log entries that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	priorityLowerThanOrEquals := int64(789) // int64 | Only include logs for which the associated external task had a priority lower than or equal to the given value. Value must be a valid `long` value. (optional)
	priorityHigherThanOrEquals := int64(789) // int64 | Only include logs for which the associated external task had a priority higher than or equal to the given value. Value must be a valid `long` value. (optional)
	creationLog := true // bool | Only include creation logs. Value may only be `true`, as `false` is the default behavior. (optional)
	failureLog := true // bool | Only include failure logs. Value may only be `true`, as `false` is the default behavior. (optional)
	successLog := true // bool | Only include success logs. Value may only be `true`, as `false` is the default behavior. (optional)
	deletionLog := true // bool | Only include deletion logs. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricExternalTaskLogAPI.GetHistoricExternalTaskLogsCount(context.Background()).LogId(logId).ExternalTaskId(externalTaskId).TopicName(topicName).WorkerId(workerId).ErrorMessage(errorMessage).ActivityIdIn(activityIdIn).ActivityInstanceIdIn(activityInstanceIdIn).ExecutionIdIn(executionIdIn).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).CreationLog(creationLog).FailureLog(failureLog).SuccessLog(successLog).DeletionLog(deletionLog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricExternalTaskLogAPI.GetHistoricExternalTaskLogsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricExternalTaskLogsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricExternalTaskLogAPI.GetHistoricExternalTaskLogsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricExternalTaskLogsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logId** | **string** | Filter by historic external task log id. | 
 **externalTaskId** | **string** | Filter by external task id. | 
 **topicName** | **string** | Filter by an external task topic. | 
 **workerId** | **string** | Filter by the id of the worker that the task was most recently locked by. | 
 **errorMessage** | **string** | Filter by external task exception message. | 
 **activityIdIn** | **string** | Only include historic external task logs which belong to one of the passed activity ids. | 
 **activityInstanceIdIn** | **string** | Only include historic external task logs which belong to one of the passed activity instance ids. | 
 **executionIdIn** | **string** | Only include historic external task logs which belong to one of the passed execution ids. | 
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **processDefinitionKey** | **string** | Filter by process definition key. | 
 **tenantIdIn** | **string** | Only include historic external task log entries which belong to one of the passed and comma-separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic external task log entries that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **priorityLowerThanOrEquals** | **int64** | Only include logs for which the associated external task had a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **priorityHigherThanOrEquals** | **int64** | Only include logs for which the associated external task had a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **creationLog** | **bool** | Only include creation logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **failureLog** | **bool** | Only include failure logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **successLog** | **bool** | Only include success logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **deletionLog** | **bool** | Only include deletion logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

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


## QueryHistoricExternalTaskLogs

> []HistoricExternalTaskLogDto QueryHistoricExternalTaskLogs(ctx).HistoricExternalTaskLogQueryDto(historicExternalTaskLogQueryDto).Execute()

Get External Task Logs (POST)



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
	historicExternalTaskLogQueryDto := *openapiclient.NewHistoricExternalTaskLogQueryDto() // HistoricExternalTaskLogQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricExternalTaskLogAPI.QueryHistoricExternalTaskLogs(context.Background()).HistoricExternalTaskLogQueryDto(historicExternalTaskLogQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricExternalTaskLogAPI.QueryHistoricExternalTaskLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricExternalTaskLogs`: []HistoricExternalTaskLogDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricExternalTaskLogAPI.QueryHistoricExternalTaskLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricExternalTaskLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **historicExternalTaskLogQueryDto** | [**HistoricExternalTaskLogQueryDto**](HistoricExternalTaskLogQueryDto.md) |  | 

### Return type

[**[]HistoricExternalTaskLogDto**](HistoricExternalTaskLogDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryHistoricExternalTaskLogsCount

> CountResultDto QueryHistoricExternalTaskLogsCount(ctx).HistoricExternalTaskLogQueryDto(historicExternalTaskLogQueryDto).Execute()

Get External Task Log Count (POST)



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
	historicExternalTaskLogQueryDto := *openapiclient.NewHistoricExternalTaskLogQueryDto() // HistoricExternalTaskLogQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricExternalTaskLogAPI.QueryHistoricExternalTaskLogsCount(context.Background()).HistoricExternalTaskLogQueryDto(historicExternalTaskLogQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricExternalTaskLogAPI.QueryHistoricExternalTaskLogsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricExternalTaskLogsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricExternalTaskLogAPI.QueryHistoricExternalTaskLogsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricExternalTaskLogsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **historicExternalTaskLogQueryDto** | [**HistoricExternalTaskLogQueryDto**](HistoricExternalTaskLogQueryDto.md) |  | 

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

