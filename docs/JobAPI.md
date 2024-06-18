# \JobAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteJob**](JobAPI.md#DeleteJob) | **Delete** /job/{id} | Delete Job
[**ExecuteJob**](JobAPI.md#ExecuteJob) | **Post** /job/{id}/execute | Execute Job
[**GetJob**](JobAPI.md#GetJob) | **Get** /job/{id} | Get Job
[**GetJobs**](JobAPI.md#GetJobs) | **Get** /job | Get Jobs
[**GetJobsCount**](JobAPI.md#GetJobsCount) | **Get** /job/count | Get Job Count
[**GetStacktrace**](JobAPI.md#GetStacktrace) | **Get** /job/{id}/stacktrace | Get Exception Stacktrace
[**QueryJobs**](JobAPI.md#QueryJobs) | **Post** /job | Get Jobs (POST)
[**QueryJobsCount**](JobAPI.md#QueryJobsCount) | **Post** /job/count | Get Job Count (POST)
[**RecalculateDuedate**](JobAPI.md#RecalculateDuedate) | **Post** /job/{id}/duedate/recalculate | Recalculate Job Due Date
[**SetJobDuedate**](JobAPI.md#SetJobDuedate) | **Put** /job/{id}/duedate | Set Job Due Date
[**SetJobPriority**](JobAPI.md#SetJobPriority) | **Put** /job/{id}/priority | Set Job Priority
[**SetJobRetries**](JobAPI.md#SetJobRetries) | **Put** /job/{id}/retries | Set Job Retries
[**SetJobRetriesAsyncOperation**](JobAPI.md#SetJobRetriesAsyncOperation) | **Post** /job/retries | Set Job Retries Async (POST)
[**UpdateJobSuspensionState**](JobAPI.md#UpdateJobSuspensionState) | **Put** /job/{id}/suspended | Activate/Suspend Job By Id
[**UpdateSuspensionStateBy**](JobAPI.md#UpdateSuspensionStateBy) | **Put** /job/suspended | Activate/Suspend Jobs



## DeleteJob

> DeleteJob(ctx, id).Execute()

Delete Job



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
	id := "id_example" // string | The id of the job to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.DeleteJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobRequest struct via the builder pattern


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


## ExecuteJob

> ExecuteJob(ctx, id).Execute()

Execute Job



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
	id := "id_example" // string | The id of the job to be executed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.ExecuteJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.ExecuteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job to be executed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteJobRequest struct via the builder pattern


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


## GetJob

> JobDto GetJob(ctx, id).Execute()

Get Job



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
	id := "id_example" // string | The id of the job to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJob`: JobDto
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobDto**](JobDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobs

> []JobDto GetJobs(ctx).JobId(jobId).JobIds(jobIds).JobDefinitionId(jobDefinitionId).ProcessInstanceId(processInstanceId).ProcessInstanceIds(processInstanceIds).ExecutionId(executionId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ActivityId(activityId).WithRetriesLeft(withRetriesLeft).Executable(executable).Timers(timers).Messages(messages).DueDates(dueDates).CreateTimes(createTimes).WithException(withException).ExceptionMessage(exceptionMessage).FailedActivityId(failedActivityId).NoRetriesLeft(noRetriesLeft).Active(active).Suspended(suspended).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeJobsWithoutTenantId(includeJobsWithoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Jobs



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
	jobId := "jobId_example" // string | Filter by job id. (optional)
	jobIds := "jobIds_example" // string | Filter by a comma-separated list of job ids. (optional)
	jobDefinitionId := "jobDefinitionId_example" // string | Only select jobs which exist for the given job definition. (optional)
	processInstanceId := "processInstanceId_example" // string | Only select jobs which exist for the given process instance. (optional)
	processInstanceIds := "processInstanceIds_example" // string | Only select jobs which exist for the given comma-separated list of process instance ids. (optional)
	executionId := "executionId_example" // string | Only select jobs which exist for the given execution. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the id of the process definition the jobs run on. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by the key of the process definition the jobs run on. (optional)
	activityId := "activityId_example" // string | Only select jobs which exist for an activity with the given id. (optional)
	withRetriesLeft := true // bool | Only select jobs which have retries left. Value may only be `true`, as `false` is the default behavior. (optional)
	executable := true // bool | Only select jobs which are executable, i.e., retries > 0 and due date is `null` or due date is in the past. Value may only be `true`, as `false` is the default behavior. (optional)
	timers := true // bool | Only select jobs that are timers. Cannot be used together with `messages`. Value may only be `true`, as `false` is the default behavior. (optional)
	messages := true // bool | Only select jobs that are messages. Cannot be used together with `timers`. Value may only be `true`, as `false` is the default behavior. (optional)
	dueDates := "dueDates_example" // string | Only select jobs where the due date is lower or higher than the given date. Due date expressions are comma-separated and are structured as follows:  A valid condition value has the form `operator_value`. `operator` is the comparison operator to be used and `value` the date value as string.  Valid operator values are: `gt` - greater than; `lt` - lower than. `value` may not contain underscore or comma characters. (optional)
	createTimes := "createTimes_example" // string | Only select jobs created before or after the given date.  Create time expressions are comma-separated and are structured as follows:  A valid condition value has the form `operator_value`. `operator` is the comparison operator to be used and `value` the date value as string.  Valid operator values are: `gt` - greater than; `lt` - lower than. `value` may not contain underscore or comma characters. (optional)
	withException := true // bool | Only select jobs that failed due to an exception. Value may only be `true`, as `false` is the default behavior. (optional)
	exceptionMessage := "exceptionMessage_example" // string | Only select jobs that failed due to an exception with the given message. (optional)
	failedActivityId := "failedActivityId_example" // string | Only select jobs that failed due to an exception at an activity with the given id. (optional)
	noRetriesLeft := true // bool | Only select jobs which have no retries left. Value may only be `true`, as `false` is the default behavior. (optional)
	active := true // bool | Only include active jobs. Value may only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | Only include suspended jobs. Value may only be `true`, as `false` is the default behavior. (optional)
	priorityLowerThanOrEquals := int64(789) // int64 | Only include jobs with a priority lower than or equal to the given value. Value must be a valid `long` value. (optional)
	priorityHigherThanOrEquals := int64(789) // int64 | Only include jobs with a priority higher than or equal to the given value. Value must be a valid `long` value. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include jobs which belong to one of the passed comma-separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include jobs which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	includeJobsWithoutTenantId := true // bool | Include jobs which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobs(context.Background()).JobId(jobId).JobIds(jobIds).JobDefinitionId(jobDefinitionId).ProcessInstanceId(processInstanceId).ProcessInstanceIds(processInstanceIds).ExecutionId(executionId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ActivityId(activityId).WithRetriesLeft(withRetriesLeft).Executable(executable).Timers(timers).Messages(messages).DueDates(dueDates).CreateTimes(createTimes).WithException(withException).ExceptionMessage(exceptionMessage).FailedActivityId(failedActivityId).NoRetriesLeft(noRetriesLeft).Active(active).Suspended(suspended).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeJobsWithoutTenantId(includeJobsWithoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobs`: []JobDto
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string** | Filter by job id. | 
 **jobIds** | **string** | Filter by a comma-separated list of job ids. | 
 **jobDefinitionId** | **string** | Only select jobs which exist for the given job definition. | 
 **processInstanceId** | **string** | Only select jobs which exist for the given process instance. | 
 **processInstanceIds** | **string** | Only select jobs which exist for the given comma-separated list of process instance ids. | 
 **executionId** | **string** | Only select jobs which exist for the given execution. | 
 **processDefinitionId** | **string** | Filter by the id of the process definition the jobs run on. | 
 **processDefinitionKey** | **string** | Filter by the key of the process definition the jobs run on. | 
 **activityId** | **string** | Only select jobs which exist for an activity with the given id. | 
 **withRetriesLeft** | **bool** | Only select jobs which have retries left. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **executable** | **bool** | Only select jobs which are executable, i.e., retries &gt; 0 and due date is &#x60;null&#x60; or due date is in the past. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **timers** | **bool** | Only select jobs that are timers. Cannot be used together with &#x60;messages&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **messages** | **bool** | Only select jobs that are messages. Cannot be used together with &#x60;timers&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **dueDates** | **string** | Only select jobs where the due date is lower or higher than the given date. Due date expressions are comma-separated and are structured as follows:  A valid condition value has the form &#x60;operator_value&#x60;. &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the date value as string.  Valid operator values are: &#x60;gt&#x60; - greater than; &#x60;lt&#x60; - lower than. &#x60;value&#x60; may not contain underscore or comma characters. | 
 **createTimes** | **string** | Only select jobs created before or after the given date.  Create time expressions are comma-separated and are structured as follows:  A valid condition value has the form &#x60;operator_value&#x60;. &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the date value as string.  Valid operator values are: &#x60;gt&#x60; - greater than; &#x60;lt&#x60; - lower than. &#x60;value&#x60; may not contain underscore or comma characters. | 
 **withException** | **bool** | Only select jobs that failed due to an exception. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **exceptionMessage** | **string** | Only select jobs that failed due to an exception with the given message. | 
 **failedActivityId** | **string** | Only select jobs that failed due to an exception at an activity with the given id. | 
 **noRetriesLeft** | **bool** | Only select jobs which have no retries left. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **active** | **bool** | Only include active jobs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | Only include suspended jobs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **priorityLowerThanOrEquals** | **int64** | Only include jobs with a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **priorityHigherThanOrEquals** | **int64** | Only include jobs with a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **tenantIdIn** | **string** | Only include jobs which belong to one of the passed comma-separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include jobs which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeJobsWithoutTenantId** | **bool** | Include jobs which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]JobDto**](JobDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobsCount

> CountResultDto GetJobsCount(ctx).JobId(jobId).JobIds(jobIds).JobDefinitionId(jobDefinitionId).ProcessInstanceId(processInstanceId).ProcessInstanceIds(processInstanceIds).ExecutionId(executionId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ActivityId(activityId).WithRetriesLeft(withRetriesLeft).Executable(executable).Timers(timers).Messages(messages).DueDates(dueDates).CreateTimes(createTimes).WithException(withException).ExceptionMessage(exceptionMessage).FailedActivityId(failedActivityId).NoRetriesLeft(noRetriesLeft).Active(active).Suspended(suspended).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeJobsWithoutTenantId(includeJobsWithoutTenantId).Execute()

Get Job Count



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
	jobId := "jobId_example" // string | Filter by job id. (optional)
	jobIds := "jobIds_example" // string | Filter by a comma-separated list of job ids. (optional)
	jobDefinitionId := "jobDefinitionId_example" // string | Only select jobs which exist for the given job definition. (optional)
	processInstanceId := "processInstanceId_example" // string | Only select jobs which exist for the given process instance. (optional)
	processInstanceIds := "processInstanceIds_example" // string | Only select jobs which exist for the given comma-separated list of process instance ids. (optional)
	executionId := "executionId_example" // string | Only select jobs which exist for the given execution. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the id of the process definition the jobs run on. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by the key of the process definition the jobs run on. (optional)
	activityId := "activityId_example" // string | Only select jobs which exist for an activity with the given id. (optional)
	withRetriesLeft := true // bool | Only select jobs which have retries left. Value may only be `true`, as `false` is the default behavior. (optional)
	executable := true // bool | Only select jobs which are executable, i.e., retries > 0 and due date is `null` or due date is in the past. Value may only be `true`, as `false` is the default behavior. (optional)
	timers := true // bool | Only select jobs that are timers. Cannot be used together with `messages`. Value may only be `true`, as `false` is the default behavior. (optional)
	messages := true // bool | Only select jobs that are messages. Cannot be used together with `timers`. Value may only be `true`, as `false` is the default behavior. (optional)
	dueDates := "dueDates_example" // string | Only select jobs where the due date is lower or higher than the given date. Due date expressions are comma-separated and are structured as follows:  A valid condition value has the form `operator_value`. `operator` is the comparison operator to be used and `value` the date value as string.  Valid operator values are: `gt` - greater than; `lt` - lower than. `value` may not contain underscore or comma characters. (optional)
	createTimes := "createTimes_example" // string | Only select jobs created before or after the given date.  Create time expressions are comma-separated and are structured as follows:  A valid condition value has the form `operator_value`. `operator` is the comparison operator to be used and `value` the date value as string.  Valid operator values are: `gt` - greater than; `lt` - lower than. `value` may not contain underscore or comma characters. (optional)
	withException := true // bool | Only select jobs that failed due to an exception. Value may only be `true`, as `false` is the default behavior. (optional)
	exceptionMessage := "exceptionMessage_example" // string | Only select jobs that failed due to an exception with the given message. (optional)
	failedActivityId := "failedActivityId_example" // string | Only select jobs that failed due to an exception at an activity with the given id. (optional)
	noRetriesLeft := true // bool | Only select jobs which have no retries left. Value may only be `true`, as `false` is the default behavior. (optional)
	active := true // bool | Only include active jobs. Value may only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | Only include suspended jobs. Value may only be `true`, as `false` is the default behavior. (optional)
	priorityLowerThanOrEquals := int64(789) // int64 | Only include jobs with a priority lower than or equal to the given value. Value must be a valid `long` value. (optional)
	priorityHigherThanOrEquals := int64(789) // int64 | Only include jobs with a priority higher than or equal to the given value. Value must be a valid `long` value. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include jobs which belong to one of the passed comma-separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include jobs which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	includeJobsWithoutTenantId := true // bool | Include jobs which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobsCount(context.Background()).JobId(jobId).JobIds(jobIds).JobDefinitionId(jobDefinitionId).ProcessInstanceId(processInstanceId).ProcessInstanceIds(processInstanceIds).ExecutionId(executionId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ActivityId(activityId).WithRetriesLeft(withRetriesLeft).Executable(executable).Timers(timers).Messages(messages).DueDates(dueDates).CreateTimes(createTimes).WithException(withException).ExceptionMessage(exceptionMessage).FailedActivityId(failedActivityId).NoRetriesLeft(noRetriesLeft).Active(active).Suspended(suspended).PriorityLowerThanOrEquals(priorityLowerThanOrEquals).PriorityHigherThanOrEquals(priorityHigherThanOrEquals).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeJobsWithoutTenantId(includeJobsWithoutTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string** | Filter by job id. | 
 **jobIds** | **string** | Filter by a comma-separated list of job ids. | 
 **jobDefinitionId** | **string** | Only select jobs which exist for the given job definition. | 
 **processInstanceId** | **string** | Only select jobs which exist for the given process instance. | 
 **processInstanceIds** | **string** | Only select jobs which exist for the given comma-separated list of process instance ids. | 
 **executionId** | **string** | Only select jobs which exist for the given execution. | 
 **processDefinitionId** | **string** | Filter by the id of the process definition the jobs run on. | 
 **processDefinitionKey** | **string** | Filter by the key of the process definition the jobs run on. | 
 **activityId** | **string** | Only select jobs which exist for an activity with the given id. | 
 **withRetriesLeft** | **bool** | Only select jobs which have retries left. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **executable** | **bool** | Only select jobs which are executable, i.e., retries &gt; 0 and due date is &#x60;null&#x60; or due date is in the past. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **timers** | **bool** | Only select jobs that are timers. Cannot be used together with &#x60;messages&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **messages** | **bool** | Only select jobs that are messages. Cannot be used together with &#x60;timers&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **dueDates** | **string** | Only select jobs where the due date is lower or higher than the given date. Due date expressions are comma-separated and are structured as follows:  A valid condition value has the form &#x60;operator_value&#x60;. &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the date value as string.  Valid operator values are: &#x60;gt&#x60; - greater than; &#x60;lt&#x60; - lower than. &#x60;value&#x60; may not contain underscore or comma characters. | 
 **createTimes** | **string** | Only select jobs created before or after the given date.  Create time expressions are comma-separated and are structured as follows:  A valid condition value has the form &#x60;operator_value&#x60;. &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the date value as string.  Valid operator values are: &#x60;gt&#x60; - greater than; &#x60;lt&#x60; - lower than. &#x60;value&#x60; may not contain underscore or comma characters. | 
 **withException** | **bool** | Only select jobs that failed due to an exception. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **exceptionMessage** | **string** | Only select jobs that failed due to an exception with the given message. | 
 **failedActivityId** | **string** | Only select jobs that failed due to an exception at an activity with the given id. | 
 **noRetriesLeft** | **bool** | Only select jobs which have no retries left. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **active** | **bool** | Only include active jobs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | Only include suspended jobs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **priorityLowerThanOrEquals** | **int64** | Only include jobs with a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **priorityHigherThanOrEquals** | **int64** | Only include jobs with a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **tenantIdIn** | **string** | Only include jobs which belong to one of the passed comma-separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include jobs which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeJobsWithoutTenantId** | **bool** | Include jobs which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

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


## GetStacktrace

> interface{} GetStacktrace(ctx, id).Execute()

Get Exception Stacktrace



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
	id := "id_example" // string | The id of the job to get the exception stacktrace for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetStacktrace(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetStacktrace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStacktrace`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetStacktrace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job to get the exception stacktrace for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStacktraceRequest struct via the builder pattern


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


## QueryJobs

> []JobDto QueryJobs(ctx).FirstResult(firstResult).MaxResults(maxResults).JobQueryDto(jobQueryDto).Execute()

Get Jobs (POST)



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
	jobQueryDto := *openapiclient.NewJobQueryDto() // JobQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.QueryJobs(context.Background()).FirstResult(firstResult).MaxResults(maxResults).JobQueryDto(jobQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.QueryJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryJobs`: []JobDto
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.QueryJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **jobQueryDto** | [**JobQueryDto**](JobQueryDto.md) |  | 

### Return type

[**[]JobDto**](JobDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryJobsCount

> CountResultDto QueryJobsCount(ctx).JobQueryDto(jobQueryDto).Execute()

Get Job Count (POST)



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
	jobQueryDto := *openapiclient.NewJobQueryDto() // JobQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.QueryJobsCount(context.Background()).JobQueryDto(jobQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.QueryJobsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryJobsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.QueryJobsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryJobsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobQueryDto** | [**JobQueryDto**](JobQueryDto.md) |  | 

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


## RecalculateDuedate

> RecalculateDuedate(ctx, id).CreationDateBased(creationDateBased).Execute()

Recalculate Job Due Date



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
	id := "id_example" // string | The id of the job to be updated.
	creationDateBased := true // bool | Recalculate the due date based on the creation date of the job or the current date. Value may only be `false`, as `true` is the default behavior.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.RecalculateDuedate(context.Background(), id).CreationDateBased(creationDateBased).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.RecalculateDuedate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecalculateDuedateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creationDateBased** | **bool** | Recalculate the due date based on the creation date of the job or the current date. Value may only be &#x60;false&#x60;, as &#x60;true&#x60; is the default behavior.  | 

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


## SetJobDuedate

> SetJobDuedate(ctx, id).JobDuedateDto(jobDuedateDto).Execute()

Set Job Due Date



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
	id := "id_example" // string | The id of the job to be updated.
	jobDuedateDto := *openapiclient.NewJobDuedateDto() // JobDuedateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.SetJobDuedate(context.Background(), id).JobDuedateDto(jobDuedateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.SetJobDuedate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetJobDuedateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobDuedateDto** | [**JobDuedateDto**](JobDuedateDto.md) |  | 

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


## SetJobPriority

> SetJobPriority(ctx, id).PriorityDto(priorityDto).Execute()

Set Job Priority



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
	id := "id_example" // string | The id of the job to be updated.
	priorityDto := *openapiclient.NewPriorityDto() // PriorityDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.SetJobPriority(context.Background(), id).PriorityDto(priorityDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.SetJobPriority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetJobPriorityRequest struct via the builder pattern


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


## SetJobRetries

> SetJobRetries(ctx, id).JobRetriesDto(jobRetriesDto).Execute()

Set Job Retries



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
	id := "id_example" // string | The id of the job to be updated.
	jobRetriesDto := *openapiclient.NewJobRetriesDto() // JobRetriesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.SetJobRetries(context.Background(), id).JobRetriesDto(jobRetriesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.SetJobRetries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetJobRetriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobRetriesDto** | [**JobRetriesDto**](JobRetriesDto.md) |  | 

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


## SetJobRetriesAsyncOperation

> BatchDto SetJobRetriesAsyncOperation(ctx).SetJobRetriesDto(setJobRetriesDto).Execute()

Set Job Retries Async (POST)



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
	setJobRetriesDto := *openapiclient.NewSetJobRetriesDto() // SetJobRetriesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.SetJobRetriesAsyncOperation(context.Background()).SetJobRetriesDto(setJobRetriesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.SetJobRetriesAsyncOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetJobRetriesAsyncOperation`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.SetJobRetriesAsyncOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetJobRetriesAsyncOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setJobRetriesDto** | [**SetJobRetriesDto**](SetJobRetriesDto.md) |  | 

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


## UpdateJobSuspensionState

> UpdateJobSuspensionState(ctx, id).SuspensionStateDto(suspensionStateDto).Execute()

Activate/Suspend Job By Id



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
	id := "id_example" // string | The id of the job to activate or suspend.
	suspensionStateDto := *openapiclient.NewSuspensionStateDto() // SuspensionStateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.UpdateJobSuspensionState(context.Background(), id).SuspensionStateDto(suspensionStateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.UpdateJobSuspensionState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job to activate or suspend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJobSuspensionStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suspensionStateDto** | [**SuspensionStateDto**](SuspensionStateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSuspensionStateBy

> UpdateSuspensionStateBy(ctx).JobSuspensionStateDto(jobSuspensionStateDto).Execute()

Activate/Suspend Jobs



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
	jobSuspensionStateDto := *openapiclient.NewJobSuspensionStateDto() // JobSuspensionStateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.UpdateSuspensionStateBy(context.Background()).JobSuspensionStateDto(jobSuspensionStateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.UpdateSuspensionStateBy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSuspensionStateByRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobSuspensionStateDto** | [**JobSuspensionStateDto**](JobSuspensionStateDto.md) |  | 

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

