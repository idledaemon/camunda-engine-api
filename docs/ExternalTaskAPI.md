# \ExternalTaskAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteExternalTaskResource**](ExternalTaskAPI.md#CompleteExternalTaskResource) | **Post** /external-task/{id}/complete | Complete
[**ExtendLock**](ExternalTaskAPI.md#ExtendLock) | **Post** /external-task/{id}/extendLock | Extend Lock
[**FetchAndLock**](ExternalTaskAPI.md#FetchAndLock) | **Post** /external-task/fetchAndLock | Fetch and Lock
[**GetExternalTask**](ExternalTaskAPI.md#GetExternalTask) | **Get** /external-task/{id} | Get
[**GetExternalTaskErrorDetails**](ExternalTaskAPI.md#GetExternalTaskErrorDetails) | **Get** /external-task/{id}/errorDetails | Get Error Details
[**GetExternalTasks**](ExternalTaskAPI.md#GetExternalTasks) | **Get** /external-task | Get List
[**GetExternalTasksCount**](ExternalTaskAPI.md#GetExternalTasksCount) | **Get** /external-task/count | Get List Count
[**GetTopicNames**](ExternalTaskAPI.md#GetTopicNames) | **Get** /external-task/topic-names | Get External Task Topic Names
[**HandleExternalTaskBpmnError**](ExternalTaskAPI.md#HandleExternalTaskBpmnError) | **Post** /external-task/{id}/bpmnError | Handle BPMN Error
[**HandleFailure**](ExternalTaskAPI.md#HandleFailure) | **Post** /external-task/{id}/failure | Handle Failure
[**Lock**](ExternalTaskAPI.md#Lock) | **Post** /external-task/{id}/lock | 
[**QueryExternalTasks**](ExternalTaskAPI.md#QueryExternalTasks) | **Post** /external-task | Get List (POST)
[**QueryExternalTasksCount**](ExternalTaskAPI.md#QueryExternalTasksCount) | **Post** /external-task/count | Get List Count (POST)
[**SetExternalTaskResourcePriority**](ExternalTaskAPI.md#SetExternalTaskResourcePriority) | **Put** /external-task/{id}/priority | Set Priority
[**SetExternalTaskResourceRetries**](ExternalTaskAPI.md#SetExternalTaskResourceRetries) | **Put** /external-task/{id}/retries | Set Retries
[**SetExternalTaskRetries**](ExternalTaskAPI.md#SetExternalTaskRetries) | **Put** /external-task/retries | Set Retries Sync
[**SetExternalTaskRetriesAsyncOperation**](ExternalTaskAPI.md#SetExternalTaskRetriesAsyncOperation) | **Post** /external-task/retries-async | Set Retries Async
[**Unlock**](ExternalTaskAPI.md#Unlock) | **Post** /external-task/{id}/unlock | Unlock



## CompleteExternalTaskResource

> CompleteExternalTaskResource(ctx, id).CompleteExternalTaskDto(completeExternalTaskDto).Execute()

Complete



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
	id := "id_example" // string | The id of the task to complete.
	completeExternalTaskDto := *openapiclient.NewCompleteExternalTaskDto() // CompleteExternalTaskDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalTaskAPI.CompleteExternalTaskResource(context.Background(), id).CompleteExternalTaskDto(completeExternalTaskDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.CompleteExternalTaskResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to complete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteExternalTaskResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **completeExternalTaskDto** | [**CompleteExternalTaskDto**](CompleteExternalTaskDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendLock

> ExtendLock(ctx, id).ExtendLockOnExternalTaskDto(extendLockOnExternalTaskDto).Execute()

Extend Lock



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
	id := "id_example" // string | The id of the external task.
	extendLockOnExternalTaskDto := *openapiclient.NewExtendLockOnExternalTaskDto() // ExtendLockOnExternalTaskDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalTaskAPI.ExtendLock(context.Background(), id).ExtendLockOnExternalTaskDto(extendLockOnExternalTaskDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.ExtendLock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the external task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extendLockOnExternalTaskDto** | [**ExtendLockOnExternalTaskDto**](ExtendLockOnExternalTaskDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAndLock

> []LockedExternalTaskDto FetchAndLock(ctx).FetchExternalTasksDto(fetchExternalTasksDto).Execute()

Fetch and Lock



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
	fetchExternalTasksDto := *openapiclient.NewFetchExternalTasksDto("WorkerId_example", NullableInt32(123)) // FetchExternalTasksDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTaskAPI.FetchAndLock(context.Background()).FetchExternalTasksDto(fetchExternalTasksDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.FetchAndLock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchAndLock`: []LockedExternalTaskDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalTaskAPI.FetchAndLock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchAndLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchExternalTasksDto** | [**FetchExternalTasksDto**](FetchExternalTasksDto.md) |  | 

### Return type

[**[]LockedExternalTaskDto**](LockedExternalTaskDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalTask

> ExternalTaskDto GetExternalTask(ctx, id).Execute()

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
	id := "id_example" // string | The id of the external task to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTaskAPI.GetExternalTask(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.GetExternalTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalTask`: ExternalTaskDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalTaskAPI.GetExternalTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the external task to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalTaskDto**](ExternalTaskDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalTaskErrorDetails

> string GetExternalTaskErrorDetails(ctx, id).Execute()

Get Error Details



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
	id := "id_example" // string | The id of the external task for which the error details should be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTaskAPI.GetExternalTaskErrorDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.GetExternalTaskErrorDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalTaskErrorDetails`: string
	fmt.Fprintf(os.Stdout, "Response from `ExternalTaskAPI.GetExternalTaskErrorDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the external task for which the error details should be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalTaskErrorDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalTasks

> []ExternalTaskDto GetExternalTasks(ctx).ExternalTaskId(externalTaskId).ExternalTaskIdIn(externalTaskIdIn).TopicName(topicName).WorkerId(workerId).Locked(locked).NotLocked(notLocked).WithRetriesLeft(withRetriesLeft).NoRetriesLeft(noRetriesLeft).LockExpirationAfter(lockExpirationAfter).LockExpirationBefore(lockExpirationBefore).ActivityId(activityId).ActivityIdIn(activityIdIn).ExecutionId(executionId).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessDefinitionId(processDefinitionId).TenantIdIn(tenantIdIn).Active(active).Suspended(suspended).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

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
	externalTaskId := "externalTaskId_example" // string | Filter by an external task's id. (optional)
	externalTaskIdIn := "externalTaskIdIn_example" // string | Filter by the comma-separated list of external task ids. (optional)
	topicName := "topicName_example" // string | Filter by an external task topic. (optional)
	workerId := "workerId_example" // string | Filter by the id of the worker that the task was most recently locked by. (optional)
	locked := true // bool | Only include external tasks that are currently locked (i.e., they have a lock time and it has not expired). Value may only be `true`, as `false` matches any external task. (optional)
	notLocked := true // bool | Only include external tasks that are currently not locked (i.e., they have no lock or it has expired). Value may only be `true`, as `false` matches any external task. (optional)
	withRetriesLeft := true // bool | Only include external tasks that have a positive (&gt; 0) number of retries (or `null`). Value may only be `true`, as `false` matches any external task. (optional)
	noRetriesLeft := true // bool | Only include external tasks that have 0 retries. Value may only be `true`, as `false` matches any external task. (optional)
	lockExpirationAfter := time.Now() // time.Time | Restrict to external tasks that have a lock that expires after a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	lockExpirationBefore := time.Now() // time.Time | Restrict to external tasks that have a lock that expires before a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	activityId := "activityId_example" // string | Filter by the id of the activity that an external task is created for. (optional)
	activityIdIn := "activityIdIn_example" // string | Filter by the comma-separated list of ids of the activities that an external task is created for. (optional)
	executionId := "executionId_example" // string | Filter by the id of the execution that an external task belongs to. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by the id of the process instance that an external task belongs to. (optional)
	processInstanceIdIn := "processInstanceIdIn_example" // string | Filter by a comma-separated list of process instance ids that an external task may belong to. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the id of the process definition that an external task belongs to. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. An external task must have one of the given tenant ids. (optional)
	active := true // bool | Only include active tasks. Value may only be `true`, as `false` matches any external task. (optional)
	suspended := true // bool | Only include suspended tasks. Value may only be `true`, as `false` matches any external task. (optional)
	priorityHigherThanOrEquals := int64(789) // int64 | Only include jobs with a priority higher than or equal to the given value. Value must be a valid `long` value. (optional)
	priorityLowerThanOrEquals := int64(789) // int64 | Only include jobs with a priority lower than or equal to the given value. Value must be a valid `long` value. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTaskAPI.GetExternalTasks(context.Background()).ExternalTaskId(externalTaskId).ExternalTaskIdIn(externalTaskIdIn).TopicName(topicName).WorkerId(workerId).Locked(locked).NotLocked(notLocked).WithRetriesLeft(withRetriesLeft).NoRetriesLeft(noRetriesLeft).LockExpirationAfter(lockExpirationAfter).LockExpirationBefore(lockExpirationBefore).ActivityId(activityId).ActivityIdIn(activityIdIn).ExecutionId(executionId).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessDefinitionId(processDefinitionId).TenantIdIn(tenantIdIn).Active(active).Suspended(suspended).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.GetExternalTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalTasks`: []ExternalTaskDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalTaskAPI.GetExternalTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalTaskId** | **string** | Filter by an external task&#39;s id. | 
 **externalTaskIdIn** | **string** | Filter by the comma-separated list of external task ids. | 
 **topicName** | **string** | Filter by an external task topic. | 
 **workerId** | **string** | Filter by the id of the worker that the task was most recently locked by. | 
 **locked** | **bool** | Only include external tasks that are currently locked (i.e., they have a lock time and it has not expired). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **notLocked** | **bool** | Only include external tasks that are currently not locked (i.e., they have no lock or it has expired). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **withRetriesLeft** | **bool** | Only include external tasks that have a positive (&amp;gt; 0) number of retries (or &#x60;null&#x60;). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **noRetriesLeft** | **bool** | Only include external tasks that have 0 retries. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **lockExpirationAfter** | **time.Time** | Restrict to external tasks that have a lock that expires after a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **lockExpirationBefore** | **time.Time** | Restrict to external tasks that have a lock that expires before a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **activityId** | **string** | Filter by the id of the activity that an external task is created for. | 
 **activityIdIn** | **string** | Filter by the comma-separated list of ids of the activities that an external task is created for. | 
 **executionId** | **string** | Filter by the id of the execution that an external task belongs to. | 
 **processInstanceId** | **string** | Filter by the id of the process instance that an external task belongs to. | 
 **processInstanceIdIn** | **string** | Filter by a comma-separated list of process instance ids that an external task may belong to. | 
 **processDefinitionId** | **string** | Filter by the id of the process definition that an external task belongs to. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. An external task must have one of the given tenant ids. | 
 **active** | **bool** | Only include active tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **suspended** | **bool** | Only include suspended tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **priorityHigherThanOrEquals** | **int64** | Only include jobs with a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **priorityLowerThanOrEquals** | **int64** | Only include jobs with a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]ExternalTaskDto**](ExternalTaskDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalTasksCount

> CountResultDto GetExternalTasksCount(ctx).ExternalTaskId(externalTaskId).ExternalTaskIdIn(externalTaskIdIn).TopicName(topicName).WorkerId(workerId).Locked(locked).NotLocked(notLocked).WithRetriesLeft(withRetriesLeft).NoRetriesLeft(noRetriesLeft).LockExpirationAfter(lockExpirationAfter).LockExpirationBefore(lockExpirationBefore).ActivityId(activityId).ActivityIdIn(activityIdIn).ExecutionId(executionId).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessDefinitionId(processDefinitionId).TenantIdIn(tenantIdIn).Active(active).Suspended(suspended).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).Execute()

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
	externalTaskId := "externalTaskId_example" // string | Filter by an external task's id. (optional)
	externalTaskIdIn := "externalTaskIdIn_example" // string | Filter by the comma-separated list of external task ids. (optional)
	topicName := "topicName_example" // string | Filter by an external task topic. (optional)
	workerId := "workerId_example" // string | Filter by the id of the worker that the task was most recently locked by. (optional)
	locked := true // bool | Only include external tasks that are currently locked (i.e., they have a lock time and it has not expired). Value may only be `true`, as `false` matches any external task. (optional)
	notLocked := true // bool | Only include external tasks that are currently not locked (i.e., they have no lock or it has expired). Value may only be `true`, as `false` matches any external task. (optional)
	withRetriesLeft := true // bool | Only include external tasks that have a positive (&gt; 0) number of retries (or `null`). Value may only be `true`, as `false` matches any external task. (optional)
	noRetriesLeft := true // bool | Only include external tasks that have 0 retries. Value may only be `true`, as `false` matches any external task. (optional)
	lockExpirationAfter := time.Now() // time.Time | Restrict to external tasks that have a lock that expires after a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	lockExpirationBefore := time.Now() // time.Time | Restrict to external tasks that have a lock that expires before a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	activityId := "activityId_example" // string | Filter by the id of the activity that an external task is created for. (optional)
	activityIdIn := "activityIdIn_example" // string | Filter by the comma-separated list of ids of the activities that an external task is created for. (optional)
	executionId := "executionId_example" // string | Filter by the id of the execution that an external task belongs to. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by the id of the process instance that an external task belongs to. (optional)
	processInstanceIdIn := "processInstanceIdIn_example" // string | Filter by a comma-separated list of process instance ids that an external task may belong to. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the id of the process definition that an external task belongs to. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. An external task must have one of the given tenant ids. (optional)
	active := true // bool | Only include active tasks. Value may only be `true`, as `false` matches any external task. (optional)
	suspended := true // bool | Only include suspended tasks. Value may only be `true`, as `false` matches any external task. (optional)
	priorityHigherThanOrEquals := int64(789) // int64 | Only include jobs with a priority higher than or equal to the given value. Value must be a valid `long` value. (optional)
	priorityLowerThanOrEquals := int64(789) // int64 | Only include jobs with a priority lower than or equal to the given value. Value must be a valid `long` value. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTaskAPI.GetExternalTasksCount(context.Background()).ExternalTaskId(externalTaskId).ExternalTaskIdIn(externalTaskIdIn).TopicName(topicName).WorkerId(workerId).Locked(locked).NotLocked(notLocked).WithRetriesLeft(withRetriesLeft).NoRetriesLeft(noRetriesLeft).LockExpirationAfter(lockExpirationAfter).LockExpirationBefore(lockExpirationBefore).ActivityId(activityId).ActivityIdIn(activityIdIn).ExecutionId(executionId).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessDefinitionId(processDefinitionId).TenantIdIn(tenantIdIn).Active(active).Suspended(suspended).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.GetExternalTasksCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalTasksCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalTaskAPI.GetExternalTasksCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalTasksCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalTaskId** | **string** | Filter by an external task&#39;s id. | 
 **externalTaskIdIn** | **string** | Filter by the comma-separated list of external task ids. | 
 **topicName** | **string** | Filter by an external task topic. | 
 **workerId** | **string** | Filter by the id of the worker that the task was most recently locked by. | 
 **locked** | **bool** | Only include external tasks that are currently locked (i.e., they have a lock time and it has not expired). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **notLocked** | **bool** | Only include external tasks that are currently not locked (i.e., they have no lock or it has expired). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **withRetriesLeft** | **bool** | Only include external tasks that have a positive (&amp;gt; 0) number of retries (or &#x60;null&#x60;). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **noRetriesLeft** | **bool** | Only include external tasks that have 0 retries. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **lockExpirationAfter** | **time.Time** | Restrict to external tasks that have a lock that expires after a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **lockExpirationBefore** | **time.Time** | Restrict to external tasks that have a lock that expires before a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **activityId** | **string** | Filter by the id of the activity that an external task is created for. | 
 **activityIdIn** | **string** | Filter by the comma-separated list of ids of the activities that an external task is created for. | 
 **executionId** | **string** | Filter by the id of the execution that an external task belongs to. | 
 **processInstanceId** | **string** | Filter by the id of the process instance that an external task belongs to. | 
 **processInstanceIdIn** | **string** | Filter by a comma-separated list of process instance ids that an external task may belong to. | 
 **processDefinitionId** | **string** | Filter by the id of the process definition that an external task belongs to. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. An external task must have one of the given tenant ids. | 
 **active** | **bool** | Only include active tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **suspended** | **bool** | Only include suspended tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **priorityHigherThanOrEquals** | **int64** | Only include jobs with a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **priorityLowerThanOrEquals** | **int64** | Only include jobs with a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 

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


## GetTopicNames

> []string GetTopicNames(ctx).WithLockedTasks(withLockedTasks).WithUnlockedTasks(withUnlockedTasks).WithRetriesLeft(withRetriesLeft).Execute()

Get External Task Topic Names



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
	withLockedTasks := true // bool | Only include external tasks that are currently locked (i.e., they have a lock time and it has not expired). Value may only be `true`, as `false` matches any external task. (optional)
	withUnlockedTasks := true // bool | Only include external tasks that are currently not locked (i.e., they have no lock or it has expired). Value may only be `true`, as `false` matches any external task. (optional)
	withRetriesLeft := true // bool | Only include external tasks that have a positive (&gt; 0) number of retries (or `null`). Value may only be `true`, as `false` matches any external task. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTaskAPI.GetTopicNames(context.Background()).WithLockedTasks(withLockedTasks).WithUnlockedTasks(withUnlockedTasks).WithRetriesLeft(withRetriesLeft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.GetTopicNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopicNames`: []string
	fmt.Fprintf(os.Stdout, "Response from `ExternalTaskAPI.GetTopicNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withLockedTasks** | **bool** | Only include external tasks that are currently locked (i.e., they have a lock time and it has not expired). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **withUnlockedTasks** | **bool** | Only include external tasks that are currently not locked (i.e., they have no lock or it has expired). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 
 **withRetriesLeft** | **bool** | Only include external tasks that have a positive (&amp;gt; 0) number of retries (or &#x60;null&#x60;). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleExternalTaskBpmnError

> HandleExternalTaskBpmnError(ctx, id).ExternalTaskBpmnError(externalTaskBpmnError).Execute()

Handle BPMN Error



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
	id := "id_example" // string | The id of the external task in which context a BPMN error is reported.
	externalTaskBpmnError := *openapiclient.NewExternalTaskBpmnError() // ExternalTaskBpmnError |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalTaskAPI.HandleExternalTaskBpmnError(context.Background(), id).ExternalTaskBpmnError(externalTaskBpmnError).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.HandleExternalTaskBpmnError``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the external task in which context a BPMN error is reported. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleExternalTaskBpmnErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalTaskBpmnError** | [**ExternalTaskBpmnError**](ExternalTaskBpmnError.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleFailure

> HandleFailure(ctx, id).ExternalTaskFailureDto(externalTaskFailureDto).Execute()

Handle Failure



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
	id := "id_example" // string | The id of the external task to report a failure for.
	externalTaskFailureDto := *openapiclient.NewExternalTaskFailureDto() // ExternalTaskFailureDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalTaskAPI.HandleFailure(context.Background(), id).ExternalTaskFailureDto(externalTaskFailureDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.HandleFailure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the external task to report a failure for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleFailureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalTaskFailureDto** | [**ExternalTaskFailureDto**](ExternalTaskFailureDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Lock

> Lock(ctx, id).LockExternalTaskDto(lockExternalTaskDto).Execute()





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
	id := "id_example" // string | The id of the external task.
	lockExternalTaskDto := *openapiclient.NewLockExternalTaskDto() // LockExternalTaskDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalTaskAPI.Lock(context.Background(), id).LockExternalTaskDto(lockExternalTaskDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.Lock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the external task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockExternalTaskDto** | [**LockExternalTaskDto**](LockExternalTaskDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryExternalTasks

> []ExternalTaskDto QueryExternalTasks(ctx).FirstResult(firstResult).MaxResults(maxResults).ExternalTaskQueryDto(externalTaskQueryDto).Execute()

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
	externalTaskQueryDto := *openapiclient.NewExternalTaskQueryDto() // ExternalTaskQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTaskAPI.QueryExternalTasks(context.Background()).FirstResult(firstResult).MaxResults(maxResults).ExternalTaskQueryDto(externalTaskQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.QueryExternalTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryExternalTasks`: []ExternalTaskDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalTaskAPI.QueryExternalTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryExternalTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **externalTaskQueryDto** | [**ExternalTaskQueryDto**](ExternalTaskQueryDto.md) |  | 

### Return type

[**[]ExternalTaskDto**](ExternalTaskDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryExternalTasksCount

> CountResultDto QueryExternalTasksCount(ctx).ExternalTaskQueryDto(externalTaskQueryDto).Execute()

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
	externalTaskQueryDto := *openapiclient.NewExternalTaskQueryDto() // ExternalTaskQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTaskAPI.QueryExternalTasksCount(context.Background()).ExternalTaskQueryDto(externalTaskQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.QueryExternalTasksCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryExternalTasksCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalTaskAPI.QueryExternalTasksCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryExternalTasksCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalTaskQueryDto** | [**ExternalTaskQueryDto**](ExternalTaskQueryDto.md) |  | 

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


## SetExternalTaskResourcePriority

> SetExternalTaskResourcePriority(ctx, id).PriorityDto(priorityDto).Execute()

Set Priority



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
	id := "id_example" // string | The id of the external task to set the priority for.
	priorityDto := *openapiclient.NewPriorityDto() // PriorityDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalTaskAPI.SetExternalTaskResourcePriority(context.Background(), id).PriorityDto(priorityDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.SetExternalTaskResourcePriority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the external task to set the priority for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetExternalTaskResourcePriorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **priorityDto** | [**PriorityDto**](PriorityDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetExternalTaskResourceRetries

> SetExternalTaskResourceRetries(ctx, id).RetriesDto(retriesDto).Execute()

Set Retries



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
	id := "id_example" // string | The id of the external task to set the number of retries for.
	retriesDto := *openapiclient.NewRetriesDto() // RetriesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalTaskAPI.SetExternalTaskResourceRetries(context.Background(), id).RetriesDto(retriesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.SetExternalTaskResourceRetries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the external task to set the number of retries for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetExternalTaskResourceRetriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **retriesDto** | [**RetriesDto**](RetriesDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetExternalTaskRetries

> SetExternalTaskRetries(ctx).SetRetriesForExternalTasksDto(setRetriesForExternalTasksDto).Execute()

Set Retries Sync



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
	setRetriesForExternalTasksDto := *openapiclient.NewSetRetriesForExternalTasksDto() // SetRetriesForExternalTasksDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalTaskAPI.SetExternalTaskRetries(context.Background()).SetRetriesForExternalTasksDto(setRetriesForExternalTasksDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.SetExternalTaskRetries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetExternalTaskRetriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setRetriesForExternalTasksDto** | [**SetRetriesForExternalTasksDto**](SetRetriesForExternalTasksDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetExternalTaskRetriesAsyncOperation

> BatchDto SetExternalTaskRetriesAsyncOperation(ctx).SetRetriesForExternalTasksDto(setRetriesForExternalTasksDto).Execute()

Set Retries Async



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
	setRetriesForExternalTasksDto := *openapiclient.NewSetRetriesForExternalTasksDto() // SetRetriesForExternalTasksDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTaskAPI.SetExternalTaskRetriesAsyncOperation(context.Background()).SetRetriesForExternalTasksDto(setRetriesForExternalTasksDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.SetExternalTaskRetriesAsyncOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetExternalTaskRetriesAsyncOperation`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalTaskAPI.SetExternalTaskRetriesAsyncOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetExternalTaskRetriesAsyncOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setRetriesForExternalTasksDto** | [**SetRetriesForExternalTasksDto**](SetRetriesForExternalTasksDto.md) |  | 

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


## Unlock

> Unlock(ctx, id).Execute()

Unlock



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
	id := "id_example" // string | The id of the external task to unlock.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalTaskAPI.Unlock(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaskAPI.Unlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the external task to unlock. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

