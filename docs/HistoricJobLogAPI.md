# \HistoricJobLogAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricJobLog**](HistoricJobLogAPI.md#GetHistoricJobLog) | **Get** /history/job-log/{id} | Get Job Log
[**GetHistoricJobLogs**](HistoricJobLogAPI.md#GetHistoricJobLogs) | **Get** /history/job-log | Get Job Logs
[**GetHistoricJobLogsCount**](HistoricJobLogAPI.md#GetHistoricJobLogsCount) | **Get** /history/job-log/count | Get Job Log Count
[**GetStacktraceHistoricJobLog**](HistoricJobLogAPI.md#GetStacktraceHistoricJobLog) | **Get** /history/job-log/{id}/stacktrace | Get Job Log Exception Stacktrace
[**QueryHistoricJobLogs**](HistoricJobLogAPI.md#QueryHistoricJobLogs) | **Post** /history/job-log | Get Job Logs (POST)
[**QueryHistoricJobLogsCount**](HistoricJobLogAPI.md#QueryHistoricJobLogsCount) | **Post** /history/job-log/count | Get Job Log Count (POST)



## GetHistoricJobLog

> HistoricJobLogDto GetHistoricJobLog(ctx, id).Execute()

Get Job Log



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
	resp, r, err := apiClient.HistoricJobLogAPI.GetHistoricJobLog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricJobLogAPI.GetHistoricJobLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricJobLog`: HistoricJobLogDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricJobLogAPI.GetHistoricJobLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the log entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricJobLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HistoricJobLogDto**](HistoricJobLogDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricJobLogs

> []HistoricJobLogDto GetHistoricJobLogs(ctx).LogId(logId).JobId(jobId).JobExceptionMessage(jobExceptionMessage).JobDefinitionId(jobDefinitionId).JobDefinitionType(jobDefinitionType).JobDefinitionConfiguration(jobDefinitionConfiguration).ActivityIdIn(activityIdIn).FailedActivityIdIn(failedActivityIdIn).ExecutionIdIn(executionIdIn).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).DeploymentId(deploymentId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Hostname(hostname).JobPriorityLowerThanOrEquals(jobPriorityLowerThanOrEquals).JobPriorityHigherThanOrEquals(jobPriorityHigherThanOrEquals).CreationLog(creationLog).FailureLog(failureLog).SuccessLog(successLog).DeletionLog(deletionLog).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Job Logs



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
	logId := "logId_example" // string | Filter by historic job log id. (optional)
	jobId := "jobId_example" // string | Filter by job id. (optional)
	jobExceptionMessage := "jobExceptionMessage_example" // string | Filter by job exception message. (optional)
	jobDefinitionId := "jobDefinitionId_example" // string | Filter by job definition id. (optional)
	jobDefinitionType := "jobDefinitionType_example" // string | Filter by job definition type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job definition types. (optional)
	jobDefinitionConfiguration := "jobDefinitionConfiguration_example" // string | Filter by job definition configuration. (optional)
	activityIdIn := "activityIdIn_example" // string | Only include historic job logs which belong to one of the passed activity ids. (optional)
	failedActivityIdIn := "failedActivityIdIn_example" // string | Only include historic job logs which belong to failures of one of the passed activity ids. (optional)
	executionIdIn := "executionIdIn_example" // string | Only include historic job logs which belong to one of the passed execution ids. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by process definition key. (optional)
	deploymentId := "deploymentId_example" // string | Filter by deployment id. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include historic job log entries which belong to one of the passed and comma- separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic job log entries that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	hostname := "hostname_example" // string | Filter by hostname. (optional)
	jobPriorityLowerThanOrEquals := int64(789) // int64 | Only include logs for which the associated job had a priority lower than or equal to the given value. Value must be a valid `long` value. (optional)
	jobPriorityHigherThanOrEquals := int64(789) // int64 | Only include logs for which the associated job had a priority higher than or equal to the given value. Value must be a valid `long` value. (optional)
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
	resp, r, err := apiClient.HistoricJobLogAPI.GetHistoricJobLogs(context.Background()).LogId(logId).JobId(jobId).JobExceptionMessage(jobExceptionMessage).JobDefinitionId(jobDefinitionId).JobDefinitionType(jobDefinitionType).JobDefinitionConfiguration(jobDefinitionConfiguration).ActivityIdIn(activityIdIn).FailedActivityIdIn(failedActivityIdIn).ExecutionIdIn(executionIdIn).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).DeploymentId(deploymentId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Hostname(hostname).JobPriorityLowerThanOrEquals(jobPriorityLowerThanOrEquals).JobPriorityHigherThanOrEquals(jobPriorityHigherThanOrEquals).CreationLog(creationLog).FailureLog(failureLog).SuccessLog(successLog).DeletionLog(deletionLog).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricJobLogAPI.GetHistoricJobLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricJobLogs`: []HistoricJobLogDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricJobLogAPI.GetHistoricJobLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricJobLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logId** | **string** | Filter by historic job log id. | 
 **jobId** | **string** | Filter by job id. | 
 **jobExceptionMessage** | **string** | Filter by job exception message. | 
 **jobDefinitionId** | **string** | Filter by job definition id. | 
 **jobDefinitionType** | **string** | Filter by job definition type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job definition types. | 
 **jobDefinitionConfiguration** | **string** | Filter by job definition configuration. | 
 **activityIdIn** | **string** | Only include historic job logs which belong to one of the passed activity ids. | 
 **failedActivityIdIn** | **string** | Only include historic job logs which belong to failures of one of the passed activity ids. | 
 **executionIdIn** | **string** | Only include historic job logs which belong to one of the passed execution ids. | 
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **processDefinitionKey** | **string** | Filter by process definition key. | 
 **deploymentId** | **string** | Filter by deployment id. | 
 **tenantIdIn** | **string** | Only include historic job log entries which belong to one of the passed and comma- separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic job log entries that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **hostname** | **string** | Filter by hostname. | 
 **jobPriorityLowerThanOrEquals** | **int64** | Only include logs for which the associated job had a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **jobPriorityHigherThanOrEquals** | **int64** | Only include logs for which the associated job had a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **creationLog** | **bool** | Only include creation logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **failureLog** | **bool** | Only include failure logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **successLog** | **bool** | Only include success logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **deletionLog** | **bool** | Only include deletion logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]HistoricJobLogDto**](HistoricJobLogDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricJobLogsCount

> CountResultDto GetHistoricJobLogsCount(ctx).LogId(logId).JobId(jobId).JobExceptionMessage(jobExceptionMessage).JobDefinitionId(jobDefinitionId).JobDefinitionType(jobDefinitionType).JobDefinitionConfiguration(jobDefinitionConfiguration).ActivityIdIn(activityIdIn).FailedActivityIdIn(failedActivityIdIn).ExecutionIdIn(executionIdIn).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).DeploymentId(deploymentId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Hostname(hostname).JobPriorityLowerThanOrEquals(jobPriorityLowerThanOrEquals).JobPriorityHigherThanOrEquals(jobPriorityHigherThanOrEquals).CreationLog(creationLog).FailureLog(failureLog).SuccessLog(successLog).DeletionLog(deletionLog).Execute()

Get Job Log Count



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
	logId := "logId_example" // string | Filter by historic job log id. (optional)
	jobId := "jobId_example" // string | Filter by job id. (optional)
	jobExceptionMessage := "jobExceptionMessage_example" // string | Filter by job exception message. (optional)
	jobDefinitionId := "jobDefinitionId_example" // string | Filter by job definition id. (optional)
	jobDefinitionType := "jobDefinitionType_example" // string | Filter by job definition type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job definition types. (optional)
	jobDefinitionConfiguration := "jobDefinitionConfiguration_example" // string | Filter by job definition configuration. (optional)
	activityIdIn := "activityIdIn_example" // string | Only include historic job logs which belong to one of the passed activity ids. (optional)
	failedActivityIdIn := "failedActivityIdIn_example" // string | Only include historic job logs which belong to failures of one of the passed activity ids. (optional)
	executionIdIn := "executionIdIn_example" // string | Only include historic job logs which belong to one of the passed execution ids. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by process definition key. (optional)
	deploymentId := "deploymentId_example" // string | Filter by deployment id. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include historic job log entries which belong to one of the passed and comma- separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic job log entries that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	hostname := "hostname_example" // string | Filter by hostname. (optional)
	jobPriorityLowerThanOrEquals := int64(789) // int64 | Only include logs for which the associated job had a priority lower than or equal to the given value. Value must be a valid `long` value. (optional)
	jobPriorityHigherThanOrEquals := int64(789) // int64 | Only include logs for which the associated job had a priority higher than or equal to the given value. Value must be a valid `long` value. (optional)
	creationLog := true // bool | Only include creation logs. Value may only be `true`, as `false` is the default behavior. (optional)
	failureLog := true // bool | Only include failure logs. Value may only be `true`, as `false` is the default behavior. (optional)
	successLog := true // bool | Only include success logs. Value may only be `true`, as `false` is the default behavior. (optional)
	deletionLog := true // bool | Only include deletion logs. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricJobLogAPI.GetHistoricJobLogsCount(context.Background()).LogId(logId).JobId(jobId).JobExceptionMessage(jobExceptionMessage).JobDefinitionId(jobDefinitionId).JobDefinitionType(jobDefinitionType).JobDefinitionConfiguration(jobDefinitionConfiguration).ActivityIdIn(activityIdIn).FailedActivityIdIn(failedActivityIdIn).ExecutionIdIn(executionIdIn).ProcessInstanceId(processInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).DeploymentId(deploymentId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Hostname(hostname).JobPriorityLowerThanOrEquals(jobPriorityLowerThanOrEquals).JobPriorityHigherThanOrEquals(jobPriorityHigherThanOrEquals).CreationLog(creationLog).FailureLog(failureLog).SuccessLog(successLog).DeletionLog(deletionLog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricJobLogAPI.GetHistoricJobLogsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricJobLogsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricJobLogAPI.GetHistoricJobLogsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricJobLogsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logId** | **string** | Filter by historic job log id. | 
 **jobId** | **string** | Filter by job id. | 
 **jobExceptionMessage** | **string** | Filter by job exception message. | 
 **jobDefinitionId** | **string** | Filter by job definition id. | 
 **jobDefinitionType** | **string** | Filter by job definition type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job definition types. | 
 **jobDefinitionConfiguration** | **string** | Filter by job definition configuration. | 
 **activityIdIn** | **string** | Only include historic job logs which belong to one of the passed activity ids. | 
 **failedActivityIdIn** | **string** | Only include historic job logs which belong to failures of one of the passed activity ids. | 
 **executionIdIn** | **string** | Only include historic job logs which belong to one of the passed execution ids. | 
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **processDefinitionKey** | **string** | Filter by process definition key. | 
 **deploymentId** | **string** | Filter by deployment id. | 
 **tenantIdIn** | **string** | Only include historic job log entries which belong to one of the passed and comma- separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic job log entries that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **hostname** | **string** | Filter by hostname. | 
 **jobPriorityLowerThanOrEquals** | **int64** | Only include logs for which the associated job had a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
 **jobPriorityHigherThanOrEquals** | **int64** | Only include logs for which the associated job had a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | 
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


## GetStacktraceHistoricJobLog

> interface{} GetStacktraceHistoricJobLog(ctx, id).Execute()

Get Job Log Exception Stacktrace



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
	id := "id_example" // string | The id of the historic job log to get the exception stacktrace for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricJobLogAPI.GetStacktraceHistoricJobLog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricJobLogAPI.GetStacktraceHistoricJobLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStacktraceHistoricJobLog`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `HistoricJobLogAPI.GetStacktraceHistoricJobLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the historic job log to get the exception stacktrace for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStacktraceHistoricJobLogRequest struct via the builder pattern


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


## QueryHistoricJobLogs

> []HistoricJobLogDto QueryHistoricJobLogs(ctx).FirstResult(firstResult).MaxResults(maxResults).HistoricJobLogQueryDto(historicJobLogQueryDto).Execute()

Get Job Logs (POST)



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
	historicJobLogQueryDto := *openapiclient.NewHistoricJobLogQueryDto() // HistoricJobLogQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricJobLogAPI.QueryHistoricJobLogs(context.Background()).FirstResult(firstResult).MaxResults(maxResults).HistoricJobLogQueryDto(historicJobLogQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricJobLogAPI.QueryHistoricJobLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricJobLogs`: []HistoricJobLogDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricJobLogAPI.QueryHistoricJobLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricJobLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **historicJobLogQueryDto** | [**HistoricJobLogQueryDto**](HistoricJobLogQueryDto.md) |  | 

### Return type

[**[]HistoricJobLogDto**](HistoricJobLogDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryHistoricJobLogsCount

> CountResultDto QueryHistoricJobLogsCount(ctx).HistoricJobLogQueryDto(historicJobLogQueryDto).Execute()

Get Job Log Count (POST)



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
	historicJobLogQueryDto := *openapiclient.NewHistoricJobLogQueryDto() // HistoricJobLogQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricJobLogAPI.QueryHistoricJobLogsCount(context.Background()).HistoricJobLogQueryDto(historicJobLogQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricJobLogAPI.QueryHistoricJobLogsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricJobLogsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricJobLogAPI.QueryHistoricJobLogsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricJobLogsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **historicJobLogQueryDto** | [**HistoricJobLogQueryDto**](HistoricJobLogQueryDto.md) |  | 

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

