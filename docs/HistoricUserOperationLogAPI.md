# \HistoricUserOperationLogAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearAnnotationUserOperationLog**](HistoricUserOperationLogAPI.md#ClearAnnotationUserOperationLog) | **Put** /history/user-operation/{operationId}/clear-annotation | Clear Annotation of an User Operation Log (Historic)
[**QueryUserOperationCount**](HistoricUserOperationLogAPI.md#QueryUserOperationCount) | **Get** /history/user-operation/count | Get User Operation Log Count
[**QueryUserOperationEntries**](HistoricUserOperationLogAPI.md#QueryUserOperationEntries) | **Get** /history/user-operation | Get User Operation Log (Historic)
[**SetAnnotationUserOperationLog**](HistoricUserOperationLogAPI.md#SetAnnotationUserOperationLog) | **Put** /history/user-operation/{operationId}/set-annotation | Set Annotation to an User Operation Log (Historic)



## ClearAnnotationUserOperationLog

> ClearAnnotationUserOperationLog(ctx, operationId).Execute()

Clear Annotation of an User Operation Log (Historic)



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
	operationId := "operationId_example" // string | The operation id of the operation log to be updated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HistoricUserOperationLogAPI.ClearAnnotationUserOperationLog(context.Background(), operationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricUserOperationLogAPI.ClearAnnotationUserOperationLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | The operation id of the operation log to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearAnnotationUserOperationLogRequest struct via the builder pattern


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


## QueryUserOperationCount

> CountResultDto QueryUserOperationCount(ctx).DeploymentId(deploymentId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).ExecutionId(executionId).CaseDefinitionId(caseDefinitionId).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).TaskId(taskId).ExternalTaskId(externalTaskId).BatchId(batchId).JobId(jobId).JobDefinitionId(jobDefinitionId).UserId(userId).OperationId(operationId).OperationType(operationType).EntityType(entityType).EntityTypeIn(entityTypeIn).Category(category).CategoryIn(categoryIn).Property(property).AfterTimestamp(afterTimestamp).BeforeTimestamp(beforeTimestamp).Execute()

Get User Operation Log Count



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
	deploymentId := "deploymentId_example" // string | Filter by deployment id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by process definition key. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	executionId := "executionId_example" // string | Filter by execution id. (optional)
	caseDefinitionId := "caseDefinitionId_example" // string | Filter by case definition id. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by case instance id. (optional)
	caseExecutionId := "caseExecutionId_example" // string | Filter by case execution id. (optional)
	taskId := "taskId_example" // string | Only include operations on this task. (optional)
	externalTaskId := "externalTaskId_example" // string | Only include operations on this external task. (optional)
	batchId := "batchId_example" // string | Only include operations on this batch. (optional)
	jobId := "jobId_example" // string | Filter by job id. (optional)
	jobDefinitionId := "jobDefinitionId_example" // string | Filter by job definition id. (optional)
	userId := "userId_example" // string | Only include operations of this user. (optional)
	operationId := "operationId_example" // string | Filter by the id of the operation. This allows fetching of multiple entries which are part of a composite operation. (optional)
	operationType := "operationType_example" // string | Filter by the type of the operation like `Claim` or `Delegate`. See the [Javadoc](https://docs.camunda.org/manual/7.21/reference/javadoc/?org/camunda/bpm/engine/history/UserOperationLogEntry.html) for a list of available operation types. (optional)
	entityType := "entityType_example" // string | Filter by the type of the entity that was affected by this operation, possible values are `Task`, `Attachment` or `IdentityLink`. (optional)
	entityTypeIn := "entityTypeIn_example" // string | Filter by a comma-separated list of types of the entities that was affected by this operation, possible values are `Task`, `Attachment` or `IdentityLink`. (optional)
	category := "category_example" // string | Filter by the category that this operation is associated with, possible values are `TaskWorker`, `Admin` or `Operator`. (optional)
	categoryIn := "categoryIn_example" // string | Filter by a comma-separated list of categories that this operation is associated with, possible values are `TaskWorker`, `Admin` or `Operator`. (optional)
	property := "property_example" // string | Only include operations that changed this property, e.g., `owner` or `assignee`. (optional)
	afterTimestamp := time.Now() // time.Time | Restrict to entries that were created after the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., 2013-01-23T14:42:45.000+0200. (optional)
	beforeTimestamp := time.Now() // time.Time | Restrict to entries that were created before the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., 2013-01-23T14:42:45.000+0200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricUserOperationLogAPI.QueryUserOperationCount(context.Background()).DeploymentId(deploymentId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).ExecutionId(executionId).CaseDefinitionId(caseDefinitionId).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).TaskId(taskId).ExternalTaskId(externalTaskId).BatchId(batchId).JobId(jobId).JobDefinitionId(jobDefinitionId).UserId(userId).OperationId(operationId).OperationType(operationType).EntityType(entityType).EntityTypeIn(entityTypeIn).Category(category).CategoryIn(categoryIn).Property(property).AfterTimestamp(afterTimestamp).BeforeTimestamp(beforeTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricUserOperationLogAPI.QueryUserOperationCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryUserOperationCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricUserOperationLogAPI.QueryUserOperationCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryUserOperationCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentId** | **string** | Filter by deployment id. | 
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **processDefinitionKey** | **string** | Filter by process definition key. | 
 **processInstanceId** | **string** | Filter by process instance id. | 
 **executionId** | **string** | Filter by execution id. | 
 **caseDefinitionId** | **string** | Filter by case definition id. | 
 **caseInstanceId** | **string** | Filter by case instance id. | 
 **caseExecutionId** | **string** | Filter by case execution id. | 
 **taskId** | **string** | Only include operations on this task. | 
 **externalTaskId** | **string** | Only include operations on this external task. | 
 **batchId** | **string** | Only include operations on this batch. | 
 **jobId** | **string** | Filter by job id. | 
 **jobDefinitionId** | **string** | Filter by job definition id. | 
 **userId** | **string** | Only include operations of this user. | 
 **operationId** | **string** | Filter by the id of the operation. This allows fetching of multiple entries which are part of a composite operation. | 
 **operationType** | **string** | Filter by the type of the operation like &#x60;Claim&#x60; or &#x60;Delegate&#x60;. See the [Javadoc](https://docs.camunda.org/manual/7.21/reference/javadoc/?org/camunda/bpm/engine/history/UserOperationLogEntry.html) for a list of available operation types. | 
 **entityType** | **string** | Filter by the type of the entity that was affected by this operation, possible values are &#x60;Task&#x60;, &#x60;Attachment&#x60; or &#x60;IdentityLink&#x60;. | 
 **entityTypeIn** | **string** | Filter by a comma-separated list of types of the entities that was affected by this operation, possible values are &#x60;Task&#x60;, &#x60;Attachment&#x60; or &#x60;IdentityLink&#x60;. | 
 **category** | **string** | Filter by the category that this operation is associated with, possible values are &#x60;TaskWorker&#x60;, &#x60;Admin&#x60; or &#x60;Operator&#x60;. | 
 **categoryIn** | **string** | Filter by a comma-separated list of categories that this operation is associated with, possible values are &#x60;TaskWorker&#x60;, &#x60;Admin&#x60; or &#x60;Operator&#x60;. | 
 **property** | **string** | Only include operations that changed this property, e.g., &#x60;owner&#x60; or &#x60;assignee&#x60;. | 
 **afterTimestamp** | **time.Time** | Restrict to entries that were created after the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200. | 
 **beforeTimestamp** | **time.Time** | Restrict to entries that were created before the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200. | 

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


## QueryUserOperationEntries

> []UserOperationLogEntryDto QueryUserOperationEntries(ctx).DeploymentId(deploymentId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).ExecutionId(executionId).CaseDefinitionId(caseDefinitionId).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).TaskId(taskId).ExternalTaskId(externalTaskId).BatchId(batchId).JobId(jobId).JobDefinitionId(jobDefinitionId).UserId(userId).OperationId(operationId).OperationType(operationType).EntityType(entityType).EntityTypeIn(entityTypeIn).Category(category).CategoryIn(categoryIn).Property(property).AfterTimestamp(afterTimestamp).BeforeTimestamp(beforeTimestamp).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get User Operation Log (Historic)



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
	deploymentId := "deploymentId_example" // string | Filter by deployment id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by process definition key. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	executionId := "executionId_example" // string | Filter by execution id. (optional)
	caseDefinitionId := "caseDefinitionId_example" // string | Filter by case definition id. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by case instance id. (optional)
	caseExecutionId := "caseExecutionId_example" // string | Filter by case execution id. (optional)
	taskId := "taskId_example" // string | Only include operations on this task. (optional)
	externalTaskId := "externalTaskId_example" // string | Only include operations on this external task. (optional)
	batchId := "batchId_example" // string | Only include operations on this batch. (optional)
	jobId := "jobId_example" // string | Filter by job id. (optional)
	jobDefinitionId := "jobDefinitionId_example" // string | Filter by job definition id. (optional)
	userId := "userId_example" // string | Only include operations of this user. (optional)
	operationId := "operationId_example" // string | Filter by the id of the operation. This allows fetching of multiple entries which are part of a composite operation. (optional)
	operationType := "operationType_example" // string | Filter by the type of the operation like `Claim` or `Delegate`. See the [Javadoc](https://docs.camunda.org/manual/7.21/reference/javadoc/?org/camunda/bpm/engine/history/UserOperationLogEntry.html) for a list of available operation types. (optional)
	entityType := "entityType_example" // string | Filter by the type of the entity that was affected by this operation, possible values are `Task`, `Attachment` or `IdentityLink`. (optional)
	entityTypeIn := "entityTypeIn_example" // string | Filter by a comma-separated list of types of the entities that was affected by this operation, possible values are `Task`, `Attachment` or `IdentityLink`. (optional)
	category := "category_example" // string | Filter by the category that this operation is associated with, possible values are `TaskWorker`, `Admin` or `Operator`. (optional)
	categoryIn := "categoryIn_example" // string | Filter by a comma-separated list of categories that this operation is associated with, possible values are `TaskWorker`, `Admin` or `Operator`. (optional)
	property := "property_example" // string | Only include operations that changed this property, e.g., `owner` or `assignee`. (optional)
	afterTimestamp := time.Now() // time.Time | Restrict to entries that were created after the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., 2013-01-23T14:42:45.000+0200. (optional)
	beforeTimestamp := time.Now() // time.Time | Restrict to entries that were created before the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., 2013-01-23T14:42:45.000+0200. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricUserOperationLogAPI.QueryUserOperationEntries(context.Background()).DeploymentId(deploymentId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).ExecutionId(executionId).CaseDefinitionId(caseDefinitionId).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).TaskId(taskId).ExternalTaskId(externalTaskId).BatchId(batchId).JobId(jobId).JobDefinitionId(jobDefinitionId).UserId(userId).OperationId(operationId).OperationType(operationType).EntityType(entityType).EntityTypeIn(entityTypeIn).Category(category).CategoryIn(categoryIn).Property(property).AfterTimestamp(afterTimestamp).BeforeTimestamp(beforeTimestamp).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricUserOperationLogAPI.QueryUserOperationEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryUserOperationEntries`: []UserOperationLogEntryDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricUserOperationLogAPI.QueryUserOperationEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryUserOperationEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentId** | **string** | Filter by deployment id. | 
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **processDefinitionKey** | **string** | Filter by process definition key. | 
 **processInstanceId** | **string** | Filter by process instance id. | 
 **executionId** | **string** | Filter by execution id. | 
 **caseDefinitionId** | **string** | Filter by case definition id. | 
 **caseInstanceId** | **string** | Filter by case instance id. | 
 **caseExecutionId** | **string** | Filter by case execution id. | 
 **taskId** | **string** | Only include operations on this task. | 
 **externalTaskId** | **string** | Only include operations on this external task. | 
 **batchId** | **string** | Only include operations on this batch. | 
 **jobId** | **string** | Filter by job id. | 
 **jobDefinitionId** | **string** | Filter by job definition id. | 
 **userId** | **string** | Only include operations of this user. | 
 **operationId** | **string** | Filter by the id of the operation. This allows fetching of multiple entries which are part of a composite operation. | 
 **operationType** | **string** | Filter by the type of the operation like &#x60;Claim&#x60; or &#x60;Delegate&#x60;. See the [Javadoc](https://docs.camunda.org/manual/7.21/reference/javadoc/?org/camunda/bpm/engine/history/UserOperationLogEntry.html) for a list of available operation types. | 
 **entityType** | **string** | Filter by the type of the entity that was affected by this operation, possible values are &#x60;Task&#x60;, &#x60;Attachment&#x60; or &#x60;IdentityLink&#x60;. | 
 **entityTypeIn** | **string** | Filter by a comma-separated list of types of the entities that was affected by this operation, possible values are &#x60;Task&#x60;, &#x60;Attachment&#x60; or &#x60;IdentityLink&#x60;. | 
 **category** | **string** | Filter by the category that this operation is associated with, possible values are &#x60;TaskWorker&#x60;, &#x60;Admin&#x60; or &#x60;Operator&#x60;. | 
 **categoryIn** | **string** | Filter by a comma-separated list of categories that this operation is associated with, possible values are &#x60;TaskWorker&#x60;, &#x60;Admin&#x60; or &#x60;Operator&#x60;. | 
 **property** | **string** | Only include operations that changed this property, e.g., &#x60;owner&#x60; or &#x60;assignee&#x60;. | 
 **afterTimestamp** | **time.Time** | Restrict to entries that were created after the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200. | 
 **beforeTimestamp** | **time.Time** | Restrict to entries that were created before the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]UserOperationLogEntryDto**](UserOperationLogEntryDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAnnotationUserOperationLog

> SetAnnotationUserOperationLog(ctx, operationId).AnnotationDto(annotationDto).Execute()

Set Annotation to an User Operation Log (Historic)



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
	operationId := "operationId_example" // string | The operation id of the operation log to be updated.
	annotationDto := *openapiclient.NewAnnotationDto() // AnnotationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HistoricUserOperationLogAPI.SetAnnotationUserOperationLog(context.Background(), operationId).AnnotationDto(annotationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricUserOperationLogAPI.SetAnnotationUserOperationLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | The operation id of the operation log to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAnnotationUserOperationLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **annotationDto** | [**AnnotationDto**](AnnotationDto.md) |  | 

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

