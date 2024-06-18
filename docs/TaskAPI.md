# \TaskAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Claim**](TaskAPI.md#Claim) | **Post** /task/{id}/claim | Claim
[**Complete**](TaskAPI.md#Complete) | **Post** /task/{id}/complete | Complete
[**CreateTask**](TaskAPI.md#CreateTask) | **Post** /task/create | Create
[**DelegateTask**](TaskAPI.md#DelegateTask) | **Post** /task/{id}/delegate | Delegate
[**DeleteTask**](TaskAPI.md#DeleteTask) | **Delete** /task/{id} | Delete
[**GetDeployedForm**](TaskAPI.md#GetDeployedForm) | **Get** /task/{id}/deployed-form | Get Deployed Form
[**GetForm**](TaskAPI.md#GetForm) | **Get** /task/{id}/form | Get Form Key
[**GetFormVariables**](TaskAPI.md#GetFormVariables) | **Get** /task/{id}/form-variables | Get Task Form Variables
[**GetRenderedForm**](TaskAPI.md#GetRenderedForm) | **Get** /task/{id}/rendered-form | Get Rendered Form
[**GetTask**](TaskAPI.md#GetTask) | **Get** /task/{id} | Get
[**GetTaskCountByCandidateGroup**](TaskAPI.md#GetTaskCountByCandidateGroup) | **Get** /task/report/candidate-group-count | Get Task Count By Candidate Group
[**GetTasks**](TaskAPI.md#GetTasks) | **Get** /task | Get List
[**GetTasksCount**](TaskAPI.md#GetTasksCount) | **Get** /task/count | Get List Count
[**HandleBpmnError**](TaskAPI.md#HandleBpmnError) | **Post** /task/{id}/bpmnError | Handle BPMN Error
[**HandleEscalation**](TaskAPI.md#HandleEscalation) | **Post** /task/{id}/bpmnEscalation | Handle BPMN Escalation
[**QueryTasks**](TaskAPI.md#QueryTasks) | **Post** /task | Get List (POST)
[**QueryTasksCount**](TaskAPI.md#QueryTasksCount) | **Post** /task/count | Get List Count (POST)
[**Resolve**](TaskAPI.md#Resolve) | **Post** /task/{id}/resolve | Resolve
[**SetAssignee**](TaskAPI.md#SetAssignee) | **Post** /task/{id}/assignee | Set Assignee
[**Submit**](TaskAPI.md#Submit) | **Post** /task/{id}/submit-form | Submit Form
[**Unclaim**](TaskAPI.md#Unclaim) | **Post** /task/{id}/unclaim | Unclaim
[**UpdateTask**](TaskAPI.md#UpdateTask) | **Put** /task/{id} | Update



## Claim

> Claim(ctx, id).UserIdDto(userIdDto).Execute()

Claim



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
	id := "id_example" // string | The id of the task to claim.
	userIdDto := *openapiclient.NewUserIdDto() // UserIdDto | Provide the id of the user that claims the task. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.Claim(context.Background(), id).UserIdDto(userIdDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.Claim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to claim. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userIdDto** | [**UserIdDto**](UserIdDto.md) | Provide the id of the user that claims the task. | 

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


## Complete

> map[string]VariableValueDto Complete(ctx, id).CompleteTaskDto(completeTaskDto).Execute()

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
	completeTaskDto := *openapiclient.NewCompleteTaskDto() // CompleteTaskDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.Complete(context.Background(), id).CompleteTaskDto(completeTaskDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.Complete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Complete`: map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.Complete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to complete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **completeTaskDto** | [**CompleteTaskDto**](CompleteTaskDto.md) |  | 

### Return type

[**map[string]VariableValueDto**](VariableValueDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTask

> CreateTask(ctx).TaskDto(taskDto).Execute()

Create



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
	taskDto := *openapiclient.NewTaskDto() // TaskDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.CreateTask(context.Background()).TaskDto(taskDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.CreateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskDto** | [**TaskDto**](TaskDto.md) |  | 

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


## DelegateTask

> DelegateTask(ctx, id).UserIdDto(userIdDto).Execute()

Delegate



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
	id := "id_example" // string | The id of the task to delegate.
	userIdDto := *openapiclient.NewUserIdDto() // UserIdDto | Provide the id of the user that the task should be delegated to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.DelegateTask(context.Background(), id).UserIdDto(userIdDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.DelegateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to delegate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelegateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userIdDto** | [**UserIdDto**](UserIdDto.md) | Provide the id of the user that the task should be delegated to. | 

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


## DeleteTask

> DeleteTask(ctx, id).Execute()

Delete



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
	id := "id_example" // string | The id of the task to be removed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.DeleteTask(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.DeleteTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskRequest struct via the builder pattern


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


## GetDeployedForm

> *os.File GetDeployedForm(ctx, id).Execute()

Get Deployed Form



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
	id := "id_example" // string | The id of the task to get the deployed form for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetDeployedForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetDeployedForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployedForm`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetDeployedForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to get the deployed form for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployedFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xhtml+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForm

> FormDto GetForm(ctx, id).Execute()

Get Form Key



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
	id := "id_example" // string | The id of the task to retrieve the form for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForm`: FormDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the form for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormDto**](FormDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFormVariables

> map[string]VariableValueDto GetFormVariables(ctx, id).VariableNames(variableNames).DeserializeValues(deserializeValues).Execute()

Get Task Form Variables



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
	id := "id_example" // string | The id of the task to retrieve the variables for.
	variableNames := "variableNames_example" // string | A comma-separated list of variable names. Allows restricting the list of requested variables to the variable names in the list. It is best practice to restrict the list of variables to the variables actually required by the form in order to minimize fetching of data. If the query parameter is ommitted all variables are fetched. If the query parameter contains non-existent variable names, the variable names are ignored. (optional)
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](http://jackson.codehaus.org/) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  Note: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetFormVariables(context.Background(), id).VariableNames(variableNames).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetFormVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormVariables`: map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetFormVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the variables for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableNames** | **string** | A comma-separated list of variable names. Allows restricting the list of requested variables to the variable names in the list. It is best practice to restrict the list of variables to the variables actually required by the form in order to minimize fetching of data. If the query parameter is ommitted all variables are fetched. If the query parameter contains non-existent variable names, the variable names are ignored. | 
 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](http://jackson.codehaus.org/) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  Note: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. | [default to true]

### Return type

[**map[string]VariableValueDto**](VariableValueDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRenderedForm

> *os.File GetRenderedForm(ctx, id).Execute()

Get Rendered Form



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
	id := "id_example" // string | The id of the task to get the rendered form for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetRenderedForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetRenderedForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRenderedForm`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetRenderedForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to get the rendered form for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRenderedFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xhtml+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> TaskDto GetTask(ctx, id).Execute()

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
	id := "id_example" // string | The id of the task to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTask(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTask`: TaskDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskDto**](TaskDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskCountByCandidateGroup

> []TaskCountByCandidateGroupResultDto GetTaskCountByCandidateGroup(ctx).Execute()

Get Task Count By Candidate Group



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTaskCountByCandidateGroup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTaskCountByCandidateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskCountByCandidateGroup`: []TaskCountByCandidateGroupResultDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTaskCountByCandidateGroup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskCountByCandidateGroupRequest struct via the builder pattern


### Return type

[**[]TaskCountByCandidateGroupResultDto**](TaskCountByCandidateGroupResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/csv, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> []TaskDto GetTasks(ctx).TaskId(taskId).TaskIdIn(taskIdIn).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyExpression(processInstanceBusinessKeyExpression).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).ProcessInstanceBusinessKeyLikeExpression(processInstanceBusinessKeyLikeExpression).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionName(processDefinitionName).ProcessDefinitionNameLike(processDefinitionNameLike).ExecutionId(executionId).CaseInstanceId(caseInstanceId).CaseInstanceBusinessKey(caseInstanceBusinessKey).CaseInstanceBusinessKeyLike(caseInstanceBusinessKeyLike).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseDefinitionName(caseDefinitionName).CaseDefinitionNameLike(caseDefinitionNameLike).CaseExecutionId(caseExecutionId).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Assignee(assignee).AssigneeExpression(assigneeExpression).AssigneeLike(assigneeLike).AssigneeLikeExpression(assigneeLikeExpression).AssigneeIn(assigneeIn).AssigneeNotIn(assigneeNotIn).Owner(owner).OwnerExpression(ownerExpression).CandidateGroup(candidateGroup).CandidateGroupExpression(candidateGroupExpression).CandidateUser(candidateUser).CandidateUserExpression(candidateUserExpression).IncludeAssignedTasks(includeAssignedTasks).InvolvedUser(involvedUser).InvolvedUserExpression(involvedUserExpression).Assigned(assigned).Unassigned(unassigned).TaskDefinitionKey(taskDefinitionKey).TaskDefinitionKeyIn(taskDefinitionKeyIn).TaskDefinitionKeyLike(taskDefinitionKeyLike).Name(name).NameNotEqual(nameNotEqual).NameLike(nameLike).NameNotLike(nameNotLike).Description(description).DescriptionLike(descriptionLike).Priority(priority).MaxPriority(maxPriority).MinPriority(minPriority).DueDate(dueDate).DueDateExpression(dueDateExpression).DueAfter(dueAfter).DueAfterExpression(dueAfterExpression).DueBefore(dueBefore).DueBeforeExpression(dueBeforeExpression).WithoutDueDate(withoutDueDate).FollowUpDate(followUpDate).FollowUpDateExpression(followUpDateExpression).FollowUpAfter(followUpAfter).FollowUpAfterExpression(followUpAfterExpression).FollowUpBefore(followUpBefore).FollowUpBeforeExpression(followUpBeforeExpression).FollowUpBeforeOrNotExistent(followUpBeforeOrNotExistent).FollowUpBeforeOrNotExistentExpression(followUpBeforeOrNotExistentExpression).CreatedOn(createdOn).CreatedOnExpression(createdOnExpression).CreatedAfter(createdAfter).CreatedAfterExpression(createdAfterExpression).CreatedBefore(createdBefore).CreatedBeforeExpression(createdBeforeExpression).UpdatedAfter(updatedAfter).UpdatedAfterExpression(updatedAfterExpression).DelegationState(delegationState).CandidateGroups(candidateGroups).CandidateGroupsExpression(candidateGroupsExpression).WithCandidateGroups(withCandidateGroups).WithoutCandidateGroups(withoutCandidateGroups).WithCandidateUsers(withCandidateUsers).WithoutCandidateUsers(withoutCandidateUsers).Active(active).Suspended(suspended).TaskVariables(taskVariables).ProcessVariables(processVariables).CaseInstanceVariables(caseInstanceVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).ParentTaskId(parentTaskId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get List



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
	taskId := "taskId_example" // string | Restrict to task with the given id. (optional)
	taskIdIn := "taskIdIn_example" // string | Restrict to tasks with any of the given ids. (optional)
	processInstanceId := "processInstanceId_example" // string | Restrict to tasks that belong to process instances with the given id. (optional)
	processInstanceIdIn := "processInstanceIdIn_example" // string | Restrict to tasks that belong to process instances with the given ids. (optional)
	processInstanceBusinessKey := "processInstanceBusinessKey_example" // string | Restrict to tasks that belong to process instances with the given business key. (optional)
	processInstanceBusinessKeyExpression := "processInstanceBusinessKeyExpression_example" // string | Restrict to tasks that belong to process instances with the given business key which  is described by an expression. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. (optional)
	processInstanceBusinessKeyIn := "processInstanceBusinessKeyIn_example" // string | Restrict to tasks that belong to process instances with one of the give business keys.  The keys need to be in a comma-separated list. (optional)
	processInstanceBusinessKeyLike := "processInstanceBusinessKeyLike_example" // string | Restrict to tasks that have a process instance business key that has the parameter  value as a substring. (optional)
	processInstanceBusinessKeyLikeExpression := "processInstanceBusinessKeyLikeExpression_example" // string | Restrict to tasks that have a process instance business key that has the parameter  value as a substring and is described by an expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Restrict to tasks that belong to a process definition with the given id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Restrict to tasks that belong to a process definition with the given key. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Restrict to tasks that belong to a process definition with one of the given keys. The  keys need to be in a comma-separated list. (optional)
	processDefinitionName := "processDefinitionName_example" // string | Restrict to tasks that belong to a process definition with the given name. (optional)
	processDefinitionNameLike := "processDefinitionNameLike_example" // string | Restrict to tasks that have a process definition name that has the parameter value as  a substring. (optional)
	executionId := "executionId_example" // string | Restrict to tasks that belong to an execution with the given id. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Restrict to tasks that belong to case instances with the given id. (optional)
	caseInstanceBusinessKey := "caseInstanceBusinessKey_example" // string | Restrict to tasks that belong to case instances with the given business key. (optional)
	caseInstanceBusinessKeyLike := "caseInstanceBusinessKeyLike_example" // string | Restrict to tasks that have a case instance business key that has the parameter value  as a substring. (optional)
	caseDefinitionId := "caseDefinitionId_example" // string | Restrict to tasks that belong to a case definition with the given id. (optional)
	caseDefinitionKey := "caseDefinitionKey_example" // string | Restrict to tasks that belong to a case definition with the given key. (optional)
	caseDefinitionName := "caseDefinitionName_example" // string | Restrict to tasks that belong to a case definition with the given name. (optional)
	caseDefinitionNameLike := "caseDefinitionNameLike_example" // string | Restrict to tasks that have a case definition name that has the parameter value as a  substring. (optional)
	caseExecutionId := "caseExecutionId_example" // string | Restrict to tasks that belong to a case execution with the given id. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Only include tasks which belong to one of the passed and comma-separated activity  instance ids. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include tasks which belong to one of the passed and comma-separated  tenant ids. (optional)
	withoutTenantId := true // bool | Only include tasks which belong to no tenant. Value may only be `true`,  as `false` is the default behavior. (optional) (default to false)
	assignee := "assignee_example" // string | Restrict to tasks that the given user is assigned to. (optional)
	assigneeExpression := "assigneeExpression_example" // string | Restrict to tasks that the user described by the given expression is assigned to.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	assigneeLike := "assigneeLike_example" // string | Restrict to tasks that have an assignee that has the parameter  value as a substring. (optional)
	assigneeLikeExpression := "assigneeLikeExpression_example" // string | Restrict to tasks that have an assignee that has the parameter value described by the  given expression as a substring. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	assigneeIn := "assigneeIn_example" // string | Only include tasks which are assigned to one of the passed and  comma-separated user ids. (optional)
	assigneeNotIn := "assigneeNotIn_example" // string | Only include tasks which are not assigned to one of the passed and comma-separated user ids. (optional)
	owner := "owner_example" // string | Restrict to tasks that the given user owns. (optional)
	ownerExpression := "ownerExpression_example" // string | Restrict to tasks that the user described by the given expression owns. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	candidateGroup := "candidateGroup_example" // string | Only include tasks that are offered to the given group. (optional)
	candidateGroupExpression := "candidateGroupExpression_example" // string | Only include tasks that are offered to the group described by the given expression.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	candidateUser := "candidateUser_example" // string | Only include tasks that are offered to the given user or to one of his groups. (optional)
	candidateUserExpression := "candidateUserExpression_example" // string | Only include tasks that are offered to the user described by the given expression.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	includeAssignedTasks := true // bool | Also include tasks that are assigned to users in candidate queries. Default is to only  include tasks that are not assigned to any user if you query by candidate user or group(s). (optional) (default to false)
	involvedUser := "involvedUser_example" // string | Only include tasks that the given user is involved in. A user is involved in a task if  an identity link exists between task and user (e.g., the user is the assignee). (optional)
	involvedUserExpression := "involvedUserExpression_example" // string | Only include tasks that the user described by the given expression is involved in. A user is involved in a task if an identity link exists between task and user (e.g., the user is the assignee). See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. (optional)
	assigned := true // bool | If set to `true`, restricts the query to all tasks that are assigned. (optional) (default to false)
	unassigned := true // bool | If set to `true`, restricts the query to all tasks that are unassigned. (optional) (default to false)
	taskDefinitionKey := "taskDefinitionKey_example" // string | Restrict to tasks that have the given key. (optional)
	taskDefinitionKeyIn := "taskDefinitionKeyIn_example" // string | Restrict to tasks that have one of the given keys. The keys need to be in a comma-separated list. (optional)
	taskDefinitionKeyLike := "taskDefinitionKeyLike_example" // string | Restrict to tasks that have a key that has the parameter value as a substring. (optional)
	name := "name_example" // string | Restrict to tasks that have the given name. (optional)
	nameNotEqual := "nameNotEqual_example" // string | Restrict to tasks that do not have the given name. (optional)
	nameLike := "nameLike_example" // string | Restrict to tasks that have a name with the given parameter value as substring. (optional)
	nameNotLike := "nameNotLike_example" // string | Restrict to tasks that do not have a name with the given parameter value as substring. (optional)
	description := "description_example" // string | Restrict to tasks that have the given description. (optional)
	descriptionLike := "descriptionLike_example" // string | Restrict to tasks that have a description that has the parameter value as a substring. (optional)
	priority := int32(56) // int32 | Restrict to tasks that have the given priority. (optional)
	maxPriority := int32(56) // int32 | Restrict to tasks that have a lower or equal priority. (optional)
	minPriority := int32(56) // int32 | Restrict to tasks that have a higher or equal priority. (optional)
	dueDate := "dueDate_example" // string | Restrict to tasks that are due on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.546+0200`. (optional)
	dueDateExpression := "dueDateExpression_example" // string | Restrict to tasks that are due on the date described by the given expression. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	dueAfter := "dueAfter_example" // string | Restrict to tasks that are due after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.435+0200`. (optional)
	dueAfterExpression := "dueAfterExpression_example" // string | Restrict to tasks that are due after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	dueBefore := "dueBefore_example" // string | Restrict to tasks that are due before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.243+0200`. (optional)
	dueBeforeExpression := "dueBeforeExpression_example" // string | Restrict to tasks that are due before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	withoutDueDate := true // bool | Only include tasks which have no due date. Value may only be `true`,  as `false` is the default behavior. (optional) (default to false)
	followUpDate := "followUpDate_example" // string | Restrict to tasks that have a followUp date on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.342+0200`. (optional)
	followUpDateExpression := "followUpDateExpression_example" // string | Restrict to tasks that have a followUp date on the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	followUpAfter := "followUpAfter_example" // string | Restrict to tasks that have a followUp date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.542+0200`. (optional)
	followUpAfterExpression := "followUpAfterExpression_example" // string | Restrict to tasks that have a followUp date after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	followUpBefore := "followUpBefore_example" // string | Restrict to tasks that have a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.234+0200`. (optional)
	followUpBeforeExpression := "followUpBeforeExpression_example" // string | Restrict to tasks that have a followUp date before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	followUpBeforeOrNotExistent := "followUpBeforeOrNotExistent_example" // string | Restrict to tasks that have no followUp date or a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.432+0200`. The typical use case is to query all `active` tasks for a user for a given date. (optional)
	followUpBeforeOrNotExistentExpression := "followUpBeforeOrNotExistentExpression_example" // string | Restrict to tasks that have no followUp date or a followUp date before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	createdOn := "createdOn_example" // string | Restrict to tasks that were created on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.324+0200`. (optional)
	createdOnExpression := "createdOnExpression_example" // string | Restrict to tasks that were created on the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	createdAfter := "createdAfter_example" // string | Restrict to tasks that were created after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.342+0200`. (optional)
	createdAfterExpression := "createdAfterExpression_example" // string | Restrict to tasks that were created after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	createdBefore := "createdBefore_example" // string | Restrict to tasks that were created before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.332+0200`. (optional)
	createdBeforeExpression := "createdBeforeExpression_example" // string | Restrict to tasks that were created before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	updatedAfter := "updatedAfter_example" // string | Restrict to tasks that were updated after the given date. Every action that fires  a [task update event](https://docs.camunda.org/manual/7.21/user-guide/process-engine/delegation-code/#task-listener-event-lifecycle) is considered as updating the task. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.332+0200`. (optional)
	updatedAfterExpression := "updatedAfterExpression_example" // string | Restrict to tasks that were updated after the date described by the given expression. Every action that fires  a [task update event](https://docs.camunda.org/manual/7.21/user-guide/process-engine/delegation-code/#task-listener-event-lifecycle) is considered as updating the task. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	delegationState := "delegationState_example" // string | Restrict to tasks that are in the given delegation state. Valid values are `PENDING` and `RESOLVED`. (optional)
	candidateGroups := "candidateGroups_example" // string | Restrict to tasks that are offered to any of the given candidate groups. Takes a comma-separated list of group names, so for example `developers,support,sales`. (optional)
	candidateGroupsExpression := "candidateGroupsExpression_example" // string | Restrict to tasks that are offered to any of the candidate groups described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to `java.util.List` of Strings. (optional)
	withCandidateGroups := true // bool | Only include tasks which have a candidate group. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	withoutCandidateGroups := true // bool | Only include tasks which have no candidate group. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	withCandidateUsers := true // bool | Only include tasks which have a candidate user. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	withoutCandidateUsers := true // bool | Only include tasks which have no candidate users. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	active := true // bool | Only include active tasks. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	suspended := true // bool | Only include suspended tasks. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	taskVariables := "taskVariables_example" // string | Only include tasks that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note**: Values are always treated as String objects on server side.  Valid `operator` values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`. `key` and `value` may not contain underscore or comma characters. (optional)
	processVariables := "processVariables_example" // string | Only include tasks that belong to process instances that have variables with certain  values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note**: Values are always treated as String objects on server side.  Valid `operator` values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`; `notLike`. `key` and `value` may not contain underscore or comma characters. (optional)
	caseInstanceVariables := "caseInstanceVariables_example" // string | Only include tasks that belong to case instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note**: Values are always treated as String objects on server side.  Valid `operator` values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`. `key` and `value` may not contain underscore or comma characters. (optional)
	variableNamesIgnoreCase := true // bool | Match all variable names in this query case-insensitively. If set `variableName` and `variablename` are treated as equal. (optional) (default to false)
	variableValuesIgnoreCase := true // bool | Match all variable values in this query case-insensitively. If set `variableValue` and `variablevalue` are treated as equal. (optional) (default to false)
	parentTaskId := "parentTaskId_example" // string | Restrict query to all tasks that are sub tasks of the given task. Takes a task id. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTasks(context.Background()).TaskId(taskId).TaskIdIn(taskIdIn).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyExpression(processInstanceBusinessKeyExpression).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).ProcessInstanceBusinessKeyLikeExpression(processInstanceBusinessKeyLikeExpression).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionName(processDefinitionName).ProcessDefinitionNameLike(processDefinitionNameLike).ExecutionId(executionId).CaseInstanceId(caseInstanceId).CaseInstanceBusinessKey(caseInstanceBusinessKey).CaseInstanceBusinessKeyLike(caseInstanceBusinessKeyLike).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseDefinitionName(caseDefinitionName).CaseDefinitionNameLike(caseDefinitionNameLike).CaseExecutionId(caseExecutionId).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Assignee(assignee).AssigneeExpression(assigneeExpression).AssigneeLike(assigneeLike).AssigneeLikeExpression(assigneeLikeExpression).AssigneeIn(assigneeIn).AssigneeNotIn(assigneeNotIn).Owner(owner).OwnerExpression(ownerExpression).CandidateGroup(candidateGroup).CandidateGroupExpression(candidateGroupExpression).CandidateUser(candidateUser).CandidateUserExpression(candidateUserExpression).IncludeAssignedTasks(includeAssignedTasks).InvolvedUser(involvedUser).InvolvedUserExpression(involvedUserExpression).Assigned(assigned).Unassigned(unassigned).TaskDefinitionKey(taskDefinitionKey).TaskDefinitionKeyIn(taskDefinitionKeyIn).TaskDefinitionKeyLike(taskDefinitionKeyLike).Name(name).NameNotEqual(nameNotEqual).NameLike(nameLike).NameNotLike(nameNotLike).Description(description).DescriptionLike(descriptionLike).Priority(priority).MaxPriority(maxPriority).MinPriority(minPriority).DueDate(dueDate).DueDateExpression(dueDateExpression).DueAfter(dueAfter).DueAfterExpression(dueAfterExpression).DueBefore(dueBefore).DueBeforeExpression(dueBeforeExpression).WithoutDueDate(withoutDueDate).FollowUpDate(followUpDate).FollowUpDateExpression(followUpDateExpression).FollowUpAfter(followUpAfter).FollowUpAfterExpression(followUpAfterExpression).FollowUpBefore(followUpBefore).FollowUpBeforeExpression(followUpBeforeExpression).FollowUpBeforeOrNotExistent(followUpBeforeOrNotExistent).FollowUpBeforeOrNotExistentExpression(followUpBeforeOrNotExistentExpression).CreatedOn(createdOn).CreatedOnExpression(createdOnExpression).CreatedAfter(createdAfter).CreatedAfterExpression(createdAfterExpression).CreatedBefore(createdBefore).CreatedBeforeExpression(createdBeforeExpression).UpdatedAfter(updatedAfter).UpdatedAfterExpression(updatedAfterExpression).DelegationState(delegationState).CandidateGroups(candidateGroups).CandidateGroupsExpression(candidateGroupsExpression).WithCandidateGroups(withCandidateGroups).WithoutCandidateGroups(withoutCandidateGroups).WithCandidateUsers(withCandidateUsers).WithoutCandidateUsers(withoutCandidateUsers).Active(active).Suspended(suspended).TaskVariables(taskVariables).ProcessVariables(processVariables).CaseInstanceVariables(caseInstanceVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).ParentTaskId(parentTaskId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTasks`: []TaskDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **string** | Restrict to task with the given id. | 
 **taskIdIn** | **string** | Restrict to tasks with any of the given ids. | 
 **processInstanceId** | **string** | Restrict to tasks that belong to process instances with the given id. | 
 **processInstanceIdIn** | **string** | Restrict to tasks that belong to process instances with the given ids. | 
 **processInstanceBusinessKey** | **string** | Restrict to tasks that belong to process instances with the given business key. | 
 **processInstanceBusinessKeyExpression** | **string** | Restrict to tasks that belong to process instances with the given business key which  is described by an expression. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. | 
 **processInstanceBusinessKeyIn** | **string** | Restrict to tasks that belong to process instances with one of the give business keys.  The keys need to be in a comma-separated list. | 
 **processInstanceBusinessKeyLike** | **string** | Restrict to tasks that have a process instance business key that has the parameter  value as a substring. | 
 **processInstanceBusinessKeyLikeExpression** | **string** | Restrict to tasks that have a process instance business key that has the parameter  value as a substring and is described by an expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **processDefinitionId** | **string** | Restrict to tasks that belong to a process definition with the given id. | 
 **processDefinitionKey** | **string** | Restrict to tasks that belong to a process definition with the given key. | 
 **processDefinitionKeyIn** | **string** | Restrict to tasks that belong to a process definition with one of the given keys. The  keys need to be in a comma-separated list. | 
 **processDefinitionName** | **string** | Restrict to tasks that belong to a process definition with the given name. | 
 **processDefinitionNameLike** | **string** | Restrict to tasks that have a process definition name that has the parameter value as  a substring. | 
 **executionId** | **string** | Restrict to tasks that belong to an execution with the given id. | 
 **caseInstanceId** | **string** | Restrict to tasks that belong to case instances with the given id. | 
 **caseInstanceBusinessKey** | **string** | Restrict to tasks that belong to case instances with the given business key. | 
 **caseInstanceBusinessKeyLike** | **string** | Restrict to tasks that have a case instance business key that has the parameter value  as a substring. | 
 **caseDefinitionId** | **string** | Restrict to tasks that belong to a case definition with the given id. | 
 **caseDefinitionKey** | **string** | Restrict to tasks that belong to a case definition with the given key. | 
 **caseDefinitionName** | **string** | Restrict to tasks that belong to a case definition with the given name. | 
 **caseDefinitionNameLike** | **string** | Restrict to tasks that have a case definition name that has the parameter value as a  substring. | 
 **caseExecutionId** | **string** | Restrict to tasks that belong to a case execution with the given id. | 
 **activityInstanceIdIn** | **string** | Only include tasks which belong to one of the passed and comma-separated activity  instance ids. | 
 **tenantIdIn** | **string** | Only include tasks which belong to one of the passed and comma-separated  tenant ids. | 
 **withoutTenantId** | **bool** | Only include tasks which belong to no tenant. Value may only be &#x60;true&#x60;,  as &#x60;false&#x60; is the default behavior. | [default to false]
 **assignee** | **string** | Restrict to tasks that the given user is assigned to. | 
 **assigneeExpression** | **string** | Restrict to tasks that the user described by the given expression is assigned to.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **assigneeLike** | **string** | Restrict to tasks that have an assignee that has the parameter  value as a substring. | 
 **assigneeLikeExpression** | **string** | Restrict to tasks that have an assignee that has the parameter value described by the  given expression as a substring. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **assigneeIn** | **string** | Only include tasks which are assigned to one of the passed and  comma-separated user ids. | 
 **assigneeNotIn** | **string** | Only include tasks which are not assigned to one of the passed and comma-separated user ids. | 
 **owner** | **string** | Restrict to tasks that the given user owns. | 
 **ownerExpression** | **string** | Restrict to tasks that the user described by the given expression owns. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **candidateGroup** | **string** | Only include tasks that are offered to the given group. | 
 **candidateGroupExpression** | **string** | Only include tasks that are offered to the group described by the given expression.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **candidateUser** | **string** | Only include tasks that are offered to the given user or to one of his groups. | 
 **candidateUserExpression** | **string** | Only include tasks that are offered to the user described by the given expression.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **includeAssignedTasks** | **bool** | Also include tasks that are assigned to users in candidate queries. Default is to only  include tasks that are not assigned to any user if you query by candidate user or group(s). | [default to false]
 **involvedUser** | **string** | Only include tasks that the given user is involved in. A user is involved in a task if  an identity link exists between task and user (e.g., the user is the assignee). | 
 **involvedUserExpression** | **string** | Only include tasks that the user described by the given expression is involved in. A user is involved in a task if an identity link exists between task and user (e.g., the user is the assignee). See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. | 
 **assigned** | **bool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are assigned. | [default to false]
 **unassigned** | **bool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are unassigned. | [default to false]
 **taskDefinitionKey** | **string** | Restrict to tasks that have the given key. | 
 **taskDefinitionKeyIn** | **string** | Restrict to tasks that have one of the given keys. The keys need to be in a comma-separated list. | 
 **taskDefinitionKeyLike** | **string** | Restrict to tasks that have a key that has the parameter value as a substring. | 
 **name** | **string** | Restrict to tasks that have the given name. | 
 **nameNotEqual** | **string** | Restrict to tasks that do not have the given name. | 
 **nameLike** | **string** | Restrict to tasks that have a name with the given parameter value as substring. | 
 **nameNotLike** | **string** | Restrict to tasks that do not have a name with the given parameter value as substring. | 
 **description** | **string** | Restrict to tasks that have the given description. | 
 **descriptionLike** | **string** | Restrict to tasks that have a description that has the parameter value as a substring. | 
 **priority** | **int32** | Restrict to tasks that have the given priority. | 
 **maxPriority** | **int32** | Restrict to tasks that have a lower or equal priority. | 
 **minPriority** | **int32** | Restrict to tasks that have a higher or equal priority. | 
 **dueDate** | **string** | Restrict to tasks that are due on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.546+0200&#x60;. | 
 **dueDateExpression** | **string** | Restrict to tasks that are due on the date described by the given expression. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **dueAfter** | **string** | Restrict to tasks that are due after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.435+0200&#x60;. | 
 **dueAfterExpression** | **string** | Restrict to tasks that are due after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **dueBefore** | **string** | Restrict to tasks that are due before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.243+0200&#x60;. | 
 **dueBeforeExpression** | **string** | Restrict to tasks that are due before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **withoutDueDate** | **bool** | Only include tasks which have no due date. Value may only be &#x60;true&#x60;,  as &#x60;false&#x60; is the default behavior. | [default to false]
 **followUpDate** | **string** | Restrict to tasks that have a followUp date on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.342+0200&#x60;. | 
 **followUpDateExpression** | **string** | Restrict to tasks that have a followUp date on the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **followUpAfter** | **string** | Restrict to tasks that have a followUp date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.542+0200&#x60;. | 
 **followUpAfterExpression** | **string** | Restrict to tasks that have a followUp date after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **followUpBefore** | **string** | Restrict to tasks that have a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.234+0200&#x60;. | 
 **followUpBeforeExpression** | **string** | Restrict to tasks that have a followUp date before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **followUpBeforeOrNotExistent** | **string** | Restrict to tasks that have no followUp date or a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.432+0200&#x60;. The typical use case is to query all &#x60;active&#x60; tasks for a user for a given date. | 
 **followUpBeforeOrNotExistentExpression** | **string** | Restrict to tasks that have no followUp date or a followUp date before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **createdOn** | **string** | Restrict to tasks that were created on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.324+0200&#x60;. | 
 **createdOnExpression** | **string** | Restrict to tasks that were created on the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **createdAfter** | **string** | Restrict to tasks that were created after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.342+0200&#x60;. | 
 **createdAfterExpression** | **string** | Restrict to tasks that were created after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **createdBefore** | **string** | Restrict to tasks that were created before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.332+0200&#x60;. | 
 **createdBeforeExpression** | **string** | Restrict to tasks that were created before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **updatedAfter** | **string** | Restrict to tasks that were updated after the given date. Every action that fires  a [task update event](https://docs.camunda.org/manual/7.21/user-guide/process-engine/delegation-code/#task-listener-event-lifecycle) is considered as updating the task. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.332+0200&#x60;. | 
 **updatedAfterExpression** | **string** | Restrict to tasks that were updated after the date described by the given expression. Every action that fires  a [task update event](https://docs.camunda.org/manual/7.21/user-guide/process-engine/delegation-code/#task-listener-event-lifecycle) is considered as updating the task. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **delegationState** | **string** | Restrict to tasks that are in the given delegation state. Valid values are &#x60;PENDING&#x60; and &#x60;RESOLVED&#x60;. | 
 **candidateGroups** | **string** | Restrict to tasks that are offered to any of the given candidate groups. Takes a comma-separated list of group names, so for example &#x60;developers,support,sales&#x60;. | 
 **candidateGroupsExpression** | **string** | Restrict to tasks that are offered to any of the candidate groups described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to &#x60;java.util.List&#x60; of Strings. | 
 **withCandidateGroups** | **bool** | Only include tasks which have a candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **withoutCandidateGroups** | **bool** | Only include tasks which have no candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **withCandidateUsers** | **bool** | Only include tasks which have a candidate user. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **withoutCandidateUsers** | **bool** | Only include tasks which have no candidate users. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **active** | **bool** | Only include active tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **suspended** | **bool** | Only include suspended tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **taskVariables** | **string** | Only include tasks that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note**: Values are always treated as String objects on server side.  Valid &#x60;operator&#x60; values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **processVariables** | **string** | Only include tasks that belong to process instances that have variables with certain  values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note**: Values are always treated as String objects on server side.  Valid &#x60;operator&#x60; values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;; &#x60;notLike&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **caseInstanceVariables** | **string** | Only include tasks that belong to case instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note**: Values are always treated as String objects on server side.  Valid &#x60;operator&#x60; values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **variableNamesIgnoreCase** | **bool** | Match all variable names in this query case-insensitively. If set &#x60;variableName&#x60; and &#x60;variablename&#x60; are treated as equal. | [default to false]
 **variableValuesIgnoreCase** | **bool** | Match all variable values in this query case-insensitively. If set &#x60;variableValue&#x60; and &#x60;variablevalue&#x60; are treated as equal. | [default to false]
 **parentTaskId** | **string** | Restrict query to all tasks that are sub tasks of the given task. Takes a task id. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]TaskDto**](TaskDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksCount

> CountResultDto GetTasksCount(ctx).TaskId(taskId).TaskIdIn(taskIdIn).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyExpression(processInstanceBusinessKeyExpression).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).ProcessInstanceBusinessKeyLikeExpression(processInstanceBusinessKeyLikeExpression).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionName(processDefinitionName).ProcessDefinitionNameLike(processDefinitionNameLike).ExecutionId(executionId).CaseInstanceId(caseInstanceId).CaseInstanceBusinessKey(caseInstanceBusinessKey).CaseInstanceBusinessKeyLike(caseInstanceBusinessKeyLike).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseDefinitionName(caseDefinitionName).CaseDefinitionNameLike(caseDefinitionNameLike).CaseExecutionId(caseExecutionId).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Assignee(assignee).AssigneeExpression(assigneeExpression).AssigneeLike(assigneeLike).AssigneeLikeExpression(assigneeLikeExpression).AssigneeIn(assigneeIn).AssigneeNotIn(assigneeNotIn).Owner(owner).OwnerExpression(ownerExpression).CandidateGroup(candidateGroup).CandidateGroupExpression(candidateGroupExpression).CandidateUser(candidateUser).CandidateUserExpression(candidateUserExpression).IncludeAssignedTasks(includeAssignedTasks).InvolvedUser(involvedUser).InvolvedUserExpression(involvedUserExpression).Assigned(assigned).Unassigned(unassigned).TaskDefinitionKey(taskDefinitionKey).TaskDefinitionKeyIn(taskDefinitionKeyIn).TaskDefinitionKeyLike(taskDefinitionKeyLike).Name(name).NameNotEqual(nameNotEqual).NameLike(nameLike).NameNotLike(nameNotLike).Description(description).DescriptionLike(descriptionLike).Priority(priority).MaxPriority(maxPriority).MinPriority(minPriority).DueDate(dueDate).DueDateExpression(dueDateExpression).DueAfter(dueAfter).DueAfterExpression(dueAfterExpression).DueBefore(dueBefore).DueBeforeExpression(dueBeforeExpression).WithoutDueDate(withoutDueDate).FollowUpDate(followUpDate).FollowUpDateExpression(followUpDateExpression).FollowUpAfter(followUpAfter).FollowUpAfterExpression(followUpAfterExpression).FollowUpBefore(followUpBefore).FollowUpBeforeExpression(followUpBeforeExpression).FollowUpBeforeOrNotExistent(followUpBeforeOrNotExistent).FollowUpBeforeOrNotExistentExpression(followUpBeforeOrNotExistentExpression).CreatedOn(createdOn).CreatedOnExpression(createdOnExpression).CreatedAfter(createdAfter).CreatedAfterExpression(createdAfterExpression).CreatedBefore(createdBefore).CreatedBeforeExpression(createdBeforeExpression).UpdatedAfter(updatedAfter).UpdatedAfterExpression(updatedAfterExpression).DelegationState(delegationState).CandidateGroups(candidateGroups).CandidateGroupsExpression(candidateGroupsExpression).WithCandidateGroups(withCandidateGroups).WithoutCandidateGroups(withoutCandidateGroups).WithCandidateUsers(withCandidateUsers).WithoutCandidateUsers(withoutCandidateUsers).Active(active).Suspended(suspended).TaskVariables(taskVariables).ProcessVariables(processVariables).CaseInstanceVariables(caseInstanceVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).ParentTaskId(parentTaskId).Execute()

Get List Count



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
	taskId := "taskId_example" // string | Restrict to task with the given id. (optional)
	taskIdIn := "taskIdIn_example" // string | Restrict to tasks with any of the given ids. (optional)
	processInstanceId := "processInstanceId_example" // string | Restrict to tasks that belong to process instances with the given id. (optional)
	processInstanceIdIn := "processInstanceIdIn_example" // string | Restrict to tasks that belong to process instances with the given ids. (optional)
	processInstanceBusinessKey := "processInstanceBusinessKey_example" // string | Restrict to tasks that belong to process instances with the given business key. (optional)
	processInstanceBusinessKeyExpression := "processInstanceBusinessKeyExpression_example" // string | Restrict to tasks that belong to process instances with the given business key which  is described by an expression. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. (optional)
	processInstanceBusinessKeyIn := "processInstanceBusinessKeyIn_example" // string | Restrict to tasks that belong to process instances with one of the give business keys.  The keys need to be in a comma-separated list. (optional)
	processInstanceBusinessKeyLike := "processInstanceBusinessKeyLike_example" // string | Restrict to tasks that have a process instance business key that has the parameter  value as a substring. (optional)
	processInstanceBusinessKeyLikeExpression := "processInstanceBusinessKeyLikeExpression_example" // string | Restrict to tasks that have a process instance business key that has the parameter  value as a substring and is described by an expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Restrict to tasks that belong to a process definition with the given id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Restrict to tasks that belong to a process definition with the given key. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Restrict to tasks that belong to a process definition with one of the given keys. The  keys need to be in a comma-separated list. (optional)
	processDefinitionName := "processDefinitionName_example" // string | Restrict to tasks that belong to a process definition with the given name. (optional)
	processDefinitionNameLike := "processDefinitionNameLike_example" // string | Restrict to tasks that have a process definition name that has the parameter value as  a substring. (optional)
	executionId := "executionId_example" // string | Restrict to tasks that belong to an execution with the given id. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Restrict to tasks that belong to case instances with the given id. (optional)
	caseInstanceBusinessKey := "caseInstanceBusinessKey_example" // string | Restrict to tasks that belong to case instances with the given business key. (optional)
	caseInstanceBusinessKeyLike := "caseInstanceBusinessKeyLike_example" // string | Restrict to tasks that have a case instance business key that has the parameter value  as a substring. (optional)
	caseDefinitionId := "caseDefinitionId_example" // string | Restrict to tasks that belong to a case definition with the given id. (optional)
	caseDefinitionKey := "caseDefinitionKey_example" // string | Restrict to tasks that belong to a case definition with the given key. (optional)
	caseDefinitionName := "caseDefinitionName_example" // string | Restrict to tasks that belong to a case definition with the given name. (optional)
	caseDefinitionNameLike := "caseDefinitionNameLike_example" // string | Restrict to tasks that have a case definition name that has the parameter value as a  substring. (optional)
	caseExecutionId := "caseExecutionId_example" // string | Restrict to tasks that belong to a case execution with the given id. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Only include tasks which belong to one of the passed and comma-separated activity  instance ids. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include tasks which belong to one of the passed and comma-separated  tenant ids. (optional)
	withoutTenantId := true // bool | Only include tasks which belong to no tenant. Value may only be `true`,  as `false` is the default behavior. (optional) (default to false)
	assignee := "assignee_example" // string | Restrict to tasks that the given user is assigned to. (optional)
	assigneeExpression := "assigneeExpression_example" // string | Restrict to tasks that the user described by the given expression is assigned to.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	assigneeLike := "assigneeLike_example" // string | Restrict to tasks that have an assignee that has the parameter  value as a substring. (optional)
	assigneeLikeExpression := "assigneeLikeExpression_example" // string | Restrict to tasks that have an assignee that has the parameter value described by the  given expression as a substring. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	assigneeIn := "assigneeIn_example" // string | Only include tasks which are assigned to one of the passed and  comma-separated user ids. (optional)
	assigneeNotIn := "assigneeNotIn_example" // string | Only include tasks which are not assigned to one of the passed and comma-separated user ids. (optional)
	owner := "owner_example" // string | Restrict to tasks that the given user owns. (optional)
	ownerExpression := "ownerExpression_example" // string | Restrict to tasks that the user described by the given expression owns. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	candidateGroup := "candidateGroup_example" // string | Only include tasks that are offered to the given group. (optional)
	candidateGroupExpression := "candidateGroupExpression_example" // string | Only include tasks that are offered to the group described by the given expression.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	candidateUser := "candidateUser_example" // string | Only include tasks that are offered to the given user or to one of his groups. (optional)
	candidateUserExpression := "candidateUserExpression_example" // string | Only include tasks that are offered to the user described by the given expression.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. (optional)
	includeAssignedTasks := true // bool | Also include tasks that are assigned to users in candidate queries. Default is to only  include tasks that are not assigned to any user if you query by candidate user or group(s). (optional) (default to false)
	involvedUser := "involvedUser_example" // string | Only include tasks that the given user is involved in. A user is involved in a task if  an identity link exists between task and user (e.g., the user is the assignee). (optional)
	involvedUserExpression := "involvedUserExpression_example" // string | Only include tasks that the user described by the given expression is involved in. A user is involved in a task if an identity link exists between task and user (e.g., the user is the assignee). See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. (optional)
	assigned := true // bool | If set to `true`, restricts the query to all tasks that are assigned. (optional) (default to false)
	unassigned := true // bool | If set to `true`, restricts the query to all tasks that are unassigned. (optional) (default to false)
	taskDefinitionKey := "taskDefinitionKey_example" // string | Restrict to tasks that have the given key. (optional)
	taskDefinitionKeyIn := "taskDefinitionKeyIn_example" // string | Restrict to tasks that have one of the given keys. The keys need to be in a comma-separated list. (optional)
	taskDefinitionKeyLike := "taskDefinitionKeyLike_example" // string | Restrict to tasks that have a key that has the parameter value as a substring. (optional)
	name := "name_example" // string | Restrict to tasks that have the given name. (optional)
	nameNotEqual := "nameNotEqual_example" // string | Restrict to tasks that do not have the given name. (optional)
	nameLike := "nameLike_example" // string | Restrict to tasks that have a name with the given parameter value as substring. (optional)
	nameNotLike := "nameNotLike_example" // string | Restrict to tasks that do not have a name with the given parameter value as substring. (optional)
	description := "description_example" // string | Restrict to tasks that have the given description. (optional)
	descriptionLike := "descriptionLike_example" // string | Restrict to tasks that have a description that has the parameter value as a substring. (optional)
	priority := int32(56) // int32 | Restrict to tasks that have the given priority. (optional)
	maxPriority := int32(56) // int32 | Restrict to tasks that have a lower or equal priority. (optional)
	minPriority := int32(56) // int32 | Restrict to tasks that have a higher or equal priority. (optional)
	dueDate := "dueDate_example" // string | Restrict to tasks that are due on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.546+0200`. (optional)
	dueDateExpression := "dueDateExpression_example" // string | Restrict to tasks that are due on the date described by the given expression. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	dueAfter := "dueAfter_example" // string | Restrict to tasks that are due after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.435+0200`. (optional)
	dueAfterExpression := "dueAfterExpression_example" // string | Restrict to tasks that are due after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	dueBefore := "dueBefore_example" // string | Restrict to tasks that are due before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.243+0200`. (optional)
	dueBeforeExpression := "dueBeforeExpression_example" // string | Restrict to tasks that are due before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	withoutDueDate := true // bool | Only include tasks which have no due date. Value may only be `true`,  as `false` is the default behavior. (optional) (default to false)
	followUpDate := "followUpDate_example" // string | Restrict to tasks that have a followUp date on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.342+0200`. (optional)
	followUpDateExpression := "followUpDateExpression_example" // string | Restrict to tasks that have a followUp date on the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	followUpAfter := "followUpAfter_example" // string | Restrict to tasks that have a followUp date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.542+0200`. (optional)
	followUpAfterExpression := "followUpAfterExpression_example" // string | Restrict to tasks that have a followUp date after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	followUpBefore := "followUpBefore_example" // string | Restrict to tasks that have a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.234+0200`. (optional)
	followUpBeforeExpression := "followUpBeforeExpression_example" // string | Restrict to tasks that have a followUp date before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	followUpBeforeOrNotExistent := "followUpBeforeOrNotExistent_example" // string | Restrict to tasks that have no followUp date or a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.432+0200`. The typical use case is to query all `active` tasks for a user for a given date. (optional)
	followUpBeforeOrNotExistentExpression := "followUpBeforeOrNotExistentExpression_example" // string | Restrict to tasks that have no followUp date or a followUp date before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	createdOn := "createdOn_example" // string | Restrict to tasks that were created on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.324+0200`. (optional)
	createdOnExpression := "createdOnExpression_example" // string | Restrict to tasks that were created on the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	createdAfter := "createdAfter_example" // string | Restrict to tasks that were created after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.342+0200`. (optional)
	createdAfterExpression := "createdAfterExpression_example" // string | Restrict to tasks that were created after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	createdBefore := "createdBefore_example" // string | Restrict to tasks that were created before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.332+0200`. (optional)
	createdBeforeExpression := "createdBeforeExpression_example" // string | Restrict to tasks that were created before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	updatedAfter := "updatedAfter_example" // string | Restrict to tasks that were updated after the given date. Every action that fires  a [task update event](https://docs.camunda.org/manual/7.21/user-guide/process-engine/delegation-code/#task-listener-event-lifecycle) is considered as updating the task. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.332+0200`. (optional)
	updatedAfterExpression := "updatedAfterExpression_example" // string | Restrict to tasks that were updated after the date described by the given expression. Every action that fires  a [task update event](https://docs.camunda.org/manual/7.21/user-guide/process-engine/delegation-code/#task-listener-event-lifecycle) is considered as updating the task. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a `java.util.Date` or `org.joda.time.DateTime` object. (optional)
	delegationState := "delegationState_example" // string | Restrict to tasks that are in the given delegation state. Valid values are `PENDING` and `RESOLVED`. (optional)
	candidateGroups := "candidateGroups_example" // string | Restrict to tasks that are offered to any of the given candidate groups. Takes a comma-separated list of group names, so for example `developers,support,sales`. (optional)
	candidateGroupsExpression := "candidateGroupsExpression_example" // string | Restrict to tasks that are offered to any of the candidate groups described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to `java.util.List` of Strings. (optional)
	withCandidateGroups := true // bool | Only include tasks which have a candidate group. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	withoutCandidateGroups := true // bool | Only include tasks which have no candidate group. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	withCandidateUsers := true // bool | Only include tasks which have a candidate user. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	withoutCandidateUsers := true // bool | Only include tasks which have no candidate users. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	active := true // bool | Only include active tasks. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	suspended := true // bool | Only include suspended tasks. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	taskVariables := "taskVariables_example" // string | Only include tasks that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note**: Values are always treated as String objects on server side.  Valid `operator` values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`. `key` and `value` may not contain underscore or comma characters. (optional)
	processVariables := "processVariables_example" // string | Only include tasks that belong to process instances that have variables with certain  values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note**: Values are always treated as String objects on server side.  Valid `operator` values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`; `notLike`. `key` and `value` may not contain underscore or comma characters. (optional)
	caseInstanceVariables := "caseInstanceVariables_example" // string | Only include tasks that belong to case instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note**: Values are always treated as String objects on server side.  Valid `operator` values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`. `key` and `value` may not contain underscore or comma characters. (optional)
	variableNamesIgnoreCase := true // bool | Match all variable names in this query case-insensitively. If set `variableName` and `variablename` are treated as equal. (optional) (default to false)
	variableValuesIgnoreCase := true // bool | Match all variable values in this query case-insensitively. If set `variableValue` and `variablevalue` are treated as equal. (optional) (default to false)
	parentTaskId := "parentTaskId_example" // string | Restrict query to all tasks that are sub tasks of the given task. Takes a task id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTasksCount(context.Background()).TaskId(taskId).TaskIdIn(taskIdIn).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyExpression(processInstanceBusinessKeyExpression).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).ProcessInstanceBusinessKeyLikeExpression(processInstanceBusinessKeyLikeExpression).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionName(processDefinitionName).ProcessDefinitionNameLike(processDefinitionNameLike).ExecutionId(executionId).CaseInstanceId(caseInstanceId).CaseInstanceBusinessKey(caseInstanceBusinessKey).CaseInstanceBusinessKeyLike(caseInstanceBusinessKeyLike).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseDefinitionName(caseDefinitionName).CaseDefinitionNameLike(caseDefinitionNameLike).CaseExecutionId(caseExecutionId).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Assignee(assignee).AssigneeExpression(assigneeExpression).AssigneeLike(assigneeLike).AssigneeLikeExpression(assigneeLikeExpression).AssigneeIn(assigneeIn).AssigneeNotIn(assigneeNotIn).Owner(owner).OwnerExpression(ownerExpression).CandidateGroup(candidateGroup).CandidateGroupExpression(candidateGroupExpression).CandidateUser(candidateUser).CandidateUserExpression(candidateUserExpression).IncludeAssignedTasks(includeAssignedTasks).InvolvedUser(involvedUser).InvolvedUserExpression(involvedUserExpression).Assigned(assigned).Unassigned(unassigned).TaskDefinitionKey(taskDefinitionKey).TaskDefinitionKeyIn(taskDefinitionKeyIn).TaskDefinitionKeyLike(taskDefinitionKeyLike).Name(name).NameNotEqual(nameNotEqual).NameLike(nameLike).NameNotLike(nameNotLike).Description(description).DescriptionLike(descriptionLike).Priority(priority).MaxPriority(maxPriority).MinPriority(minPriority).DueDate(dueDate).DueDateExpression(dueDateExpression).DueAfter(dueAfter).DueAfterExpression(dueAfterExpression).DueBefore(dueBefore).DueBeforeExpression(dueBeforeExpression).WithoutDueDate(withoutDueDate).FollowUpDate(followUpDate).FollowUpDateExpression(followUpDateExpression).FollowUpAfter(followUpAfter).FollowUpAfterExpression(followUpAfterExpression).FollowUpBefore(followUpBefore).FollowUpBeforeExpression(followUpBeforeExpression).FollowUpBeforeOrNotExistent(followUpBeforeOrNotExistent).FollowUpBeforeOrNotExistentExpression(followUpBeforeOrNotExistentExpression).CreatedOn(createdOn).CreatedOnExpression(createdOnExpression).CreatedAfter(createdAfter).CreatedAfterExpression(createdAfterExpression).CreatedBefore(createdBefore).CreatedBeforeExpression(createdBeforeExpression).UpdatedAfter(updatedAfter).UpdatedAfterExpression(updatedAfterExpression).DelegationState(delegationState).CandidateGroups(candidateGroups).CandidateGroupsExpression(candidateGroupsExpression).WithCandidateGroups(withCandidateGroups).WithoutCandidateGroups(withoutCandidateGroups).WithCandidateUsers(withCandidateUsers).WithoutCandidateUsers(withoutCandidateUsers).Active(active).Suspended(suspended).TaskVariables(taskVariables).ProcessVariables(processVariables).CaseInstanceVariables(caseInstanceVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).ParentTaskId(parentTaskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTasksCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTasksCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTasksCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **string** | Restrict to task with the given id. | 
 **taskIdIn** | **string** | Restrict to tasks with any of the given ids. | 
 **processInstanceId** | **string** | Restrict to tasks that belong to process instances with the given id. | 
 **processInstanceIdIn** | **string** | Restrict to tasks that belong to process instances with the given ids. | 
 **processInstanceBusinessKey** | **string** | Restrict to tasks that belong to process instances with the given business key. | 
 **processInstanceBusinessKeyExpression** | **string** | Restrict to tasks that belong to process instances with the given business key which  is described by an expression. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. | 
 **processInstanceBusinessKeyIn** | **string** | Restrict to tasks that belong to process instances with one of the give business keys.  The keys need to be in a comma-separated list. | 
 **processInstanceBusinessKeyLike** | **string** | Restrict to tasks that have a process instance business key that has the parameter  value as a substring. | 
 **processInstanceBusinessKeyLikeExpression** | **string** | Restrict to tasks that have a process instance business key that has the parameter  value as a substring and is described by an expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **processDefinitionId** | **string** | Restrict to tasks that belong to a process definition with the given id. | 
 **processDefinitionKey** | **string** | Restrict to tasks that belong to a process definition with the given key. | 
 **processDefinitionKeyIn** | **string** | Restrict to tasks that belong to a process definition with one of the given keys. The  keys need to be in a comma-separated list. | 
 **processDefinitionName** | **string** | Restrict to tasks that belong to a process definition with the given name. | 
 **processDefinitionNameLike** | **string** | Restrict to tasks that have a process definition name that has the parameter value as  a substring. | 
 **executionId** | **string** | Restrict to tasks that belong to an execution with the given id. | 
 **caseInstanceId** | **string** | Restrict to tasks that belong to case instances with the given id. | 
 **caseInstanceBusinessKey** | **string** | Restrict to tasks that belong to case instances with the given business key. | 
 **caseInstanceBusinessKeyLike** | **string** | Restrict to tasks that have a case instance business key that has the parameter value  as a substring. | 
 **caseDefinitionId** | **string** | Restrict to tasks that belong to a case definition with the given id. | 
 **caseDefinitionKey** | **string** | Restrict to tasks that belong to a case definition with the given key. | 
 **caseDefinitionName** | **string** | Restrict to tasks that belong to a case definition with the given name. | 
 **caseDefinitionNameLike** | **string** | Restrict to tasks that have a case definition name that has the parameter value as a  substring. | 
 **caseExecutionId** | **string** | Restrict to tasks that belong to a case execution with the given id. | 
 **activityInstanceIdIn** | **string** | Only include tasks which belong to one of the passed and comma-separated activity  instance ids. | 
 **tenantIdIn** | **string** | Only include tasks which belong to one of the passed and comma-separated  tenant ids. | 
 **withoutTenantId** | **bool** | Only include tasks which belong to no tenant. Value may only be &#x60;true&#x60;,  as &#x60;false&#x60; is the default behavior. | [default to false]
 **assignee** | **string** | Restrict to tasks that the given user is assigned to. | 
 **assigneeExpression** | **string** | Restrict to tasks that the user described by the given expression is assigned to.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **assigneeLike** | **string** | Restrict to tasks that have an assignee that has the parameter  value as a substring. | 
 **assigneeLikeExpression** | **string** | Restrict to tasks that have an assignee that has the parameter value described by the  given expression as a substring. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **assigneeIn** | **string** | Only include tasks which are assigned to one of the passed and  comma-separated user ids. | 
 **assigneeNotIn** | **string** | Only include tasks which are not assigned to one of the passed and comma-separated user ids. | 
 **owner** | **string** | Restrict to tasks that the given user owns. | 
 **ownerExpression** | **string** | Restrict to tasks that the user described by the given expression owns. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **candidateGroup** | **string** | Only include tasks that are offered to the given group. | 
 **candidateGroupExpression** | **string** | Only include tasks that are offered to the group described by the given expression.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **candidateUser** | **string** | Only include tasks that are offered to the given user or to one of his groups. | 
 **candidateUserExpression** | **string** | Only include tasks that are offered to the user described by the given expression.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | 
 **includeAssignedTasks** | **bool** | Also include tasks that are assigned to users in candidate queries. Default is to only  include tasks that are not assigned to any user if you query by candidate user or group(s). | [default to false]
 **involvedUser** | **string** | Only include tasks that the given user is involved in. A user is involved in a task if  an identity link exists between task and user (e.g., the user is the assignee). | 
 **involvedUserExpression** | **string** | Only include tasks that the user described by the given expression is involved in. A user is involved in a task if an identity link exists between task and user (e.g., the user is the assignee). See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. | 
 **assigned** | **bool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are assigned. | [default to false]
 **unassigned** | **bool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are unassigned. | [default to false]
 **taskDefinitionKey** | **string** | Restrict to tasks that have the given key. | 
 **taskDefinitionKeyIn** | **string** | Restrict to tasks that have one of the given keys. The keys need to be in a comma-separated list. | 
 **taskDefinitionKeyLike** | **string** | Restrict to tasks that have a key that has the parameter value as a substring. | 
 **name** | **string** | Restrict to tasks that have the given name. | 
 **nameNotEqual** | **string** | Restrict to tasks that do not have the given name. | 
 **nameLike** | **string** | Restrict to tasks that have a name with the given parameter value as substring. | 
 **nameNotLike** | **string** | Restrict to tasks that do not have a name with the given parameter value as substring. | 
 **description** | **string** | Restrict to tasks that have the given description. | 
 **descriptionLike** | **string** | Restrict to tasks that have a description that has the parameter value as a substring. | 
 **priority** | **int32** | Restrict to tasks that have the given priority. | 
 **maxPriority** | **int32** | Restrict to tasks that have a lower or equal priority. | 
 **minPriority** | **int32** | Restrict to tasks that have a higher or equal priority. | 
 **dueDate** | **string** | Restrict to tasks that are due on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.546+0200&#x60;. | 
 **dueDateExpression** | **string** | Restrict to tasks that are due on the date described by the given expression. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **dueAfter** | **string** | Restrict to tasks that are due after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.435+0200&#x60;. | 
 **dueAfterExpression** | **string** | Restrict to tasks that are due after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **dueBefore** | **string** | Restrict to tasks that are due before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.243+0200&#x60;. | 
 **dueBeforeExpression** | **string** | Restrict to tasks that are due before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **withoutDueDate** | **bool** | Only include tasks which have no due date. Value may only be &#x60;true&#x60;,  as &#x60;false&#x60; is the default behavior. | [default to false]
 **followUpDate** | **string** | Restrict to tasks that have a followUp date on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.342+0200&#x60;. | 
 **followUpDateExpression** | **string** | Restrict to tasks that have a followUp date on the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **followUpAfter** | **string** | Restrict to tasks that have a followUp date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.542+0200&#x60;. | 
 **followUpAfterExpression** | **string** | Restrict to tasks that have a followUp date after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **followUpBefore** | **string** | Restrict to tasks that have a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.234+0200&#x60;. | 
 **followUpBeforeExpression** | **string** | Restrict to tasks that have a followUp date before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **followUpBeforeOrNotExistent** | **string** | Restrict to tasks that have no followUp date or a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.432+0200&#x60;. The typical use case is to query all &#x60;active&#x60; tasks for a user for a given date. | 
 **followUpBeforeOrNotExistentExpression** | **string** | Restrict to tasks that have no followUp date or a followUp date before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **createdOn** | **string** | Restrict to tasks that were created on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.324+0200&#x60;. | 
 **createdOnExpression** | **string** | Restrict to tasks that were created on the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **createdAfter** | **string** | Restrict to tasks that were created after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.342+0200&#x60;. | 
 **createdAfterExpression** | **string** | Restrict to tasks that were created after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **createdBefore** | **string** | Restrict to tasks that were created before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.332+0200&#x60;. | 
 **createdBeforeExpression** | **string** | Restrict to tasks that were created before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **updatedAfter** | **string** | Restrict to tasks that were updated after the given date. Every action that fires  a [task update event](https://docs.camunda.org/manual/7.21/user-guide/process-engine/delegation-code/#task-listener-event-lifecycle) is considered as updating the task. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.332+0200&#x60;. | 
 **updatedAfterExpression** | **string** | Restrict to tasks that were updated after the date described by the given expression. Every action that fires  a [task update event](https://docs.camunda.org/manual/7.21/user-guide/process-engine/delegation-code/#task-listener-event-lifecycle) is considered as updating the task. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | 
 **delegationState** | **string** | Restrict to tasks that are in the given delegation state. Valid values are &#x60;PENDING&#x60; and &#x60;RESOLVED&#x60;. | 
 **candidateGroups** | **string** | Restrict to tasks that are offered to any of the given candidate groups. Takes a comma-separated list of group names, so for example &#x60;developers,support,sales&#x60;. | 
 **candidateGroupsExpression** | **string** | Restrict to tasks that are offered to any of the candidate groups described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to &#x60;java.util.List&#x60; of Strings. | 
 **withCandidateGroups** | **bool** | Only include tasks which have a candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **withoutCandidateGroups** | **bool** | Only include tasks which have no candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **withCandidateUsers** | **bool** | Only include tasks which have a candidate user. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **withoutCandidateUsers** | **bool** | Only include tasks which have no candidate users. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **active** | **bool** | Only include active tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **suspended** | **bool** | Only include suspended tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **taskVariables** | **string** | Only include tasks that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note**: Values are always treated as String objects on server side.  Valid &#x60;operator&#x60; values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **processVariables** | **string** | Only include tasks that belong to process instances that have variables with certain  values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note**: Values are always treated as String objects on server side.  Valid &#x60;operator&#x60; values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;; &#x60;notLike&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **caseInstanceVariables** | **string** | Only include tasks that belong to case instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note**: Values are always treated as String objects on server side.  Valid &#x60;operator&#x60; values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **variableNamesIgnoreCase** | **bool** | Match all variable names in this query case-insensitively. If set &#x60;variableName&#x60; and &#x60;variablename&#x60; are treated as equal. | [default to false]
 **variableValuesIgnoreCase** | **bool** | Match all variable values in this query case-insensitively. If set &#x60;variableValue&#x60; and &#x60;variablevalue&#x60; are treated as equal. | [default to false]
 **parentTaskId** | **string** | Restrict query to all tasks that are sub tasks of the given task. Takes a task id. | 

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


## HandleBpmnError

> HandleBpmnError(ctx, id).TaskBpmnErrorDto(taskBpmnErrorDto).Execute()

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
	id := "id_example" // string | The id of the task a BPMN error is reported for.
	taskBpmnErrorDto := *openapiclient.NewTaskBpmnErrorDto() // TaskBpmnErrorDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.HandleBpmnError(context.Background(), id).TaskBpmnErrorDto(taskBpmnErrorDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.HandleBpmnError``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task a BPMN error is reported for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleBpmnErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskBpmnErrorDto** | [**TaskBpmnErrorDto**](TaskBpmnErrorDto.md) |  | 

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


## HandleEscalation

> HandleEscalation(ctx, id).TaskEscalationDto(taskEscalationDto).Execute()

Handle BPMN Escalation



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
	id := "id_example" // string | The id of the task in which context a BPMN escalation is reported.
	taskEscalationDto := *openapiclient.NewTaskEscalationDto() // TaskEscalationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.HandleEscalation(context.Background(), id).TaskEscalationDto(taskEscalationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.HandleEscalation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task in which context a BPMN escalation is reported. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleEscalationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskEscalationDto** | [**TaskEscalationDto**](TaskEscalationDto.md) |  | 

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


## QueryTasks

> []TaskDto QueryTasks(ctx).FirstResult(firstResult).MaxResults(maxResults).TaskQueryDto(taskQueryDto).Execute()

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
	taskQueryDto := *openapiclient.NewTaskQueryDto() // TaskQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.QueryTasks(context.Background()).FirstResult(firstResult).MaxResults(maxResults).TaskQueryDto(taskQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.QueryTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryTasks`: []TaskDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.QueryTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **taskQueryDto** | [**TaskQueryDto**](TaskQueryDto.md) |  | 

### Return type

[**[]TaskDto**](TaskDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTasksCount

> CountResultDto QueryTasksCount(ctx).TaskQueryDto(taskQueryDto).Execute()

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
	taskQueryDto := *openapiclient.NewTaskQueryDto() // TaskQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.QueryTasksCount(context.Background()).TaskQueryDto(taskQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.QueryTasksCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryTasksCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.QueryTasksCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTasksCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskQueryDto** | [**TaskQueryDto**](TaskQueryDto.md) |  | 

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


## Resolve

> Resolve(ctx, id).CompleteTaskDto(completeTaskDto).Execute()

Resolve



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
	id := "id_example" // string | The id of the task to resolve.
	completeTaskDto := *openapiclient.NewCompleteTaskDto() // CompleteTaskDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.Resolve(context.Background(), id).CompleteTaskDto(completeTaskDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.Resolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to resolve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **completeTaskDto** | [**CompleteTaskDto**](CompleteTaskDto.md) |  | 

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


## SetAssignee

> SetAssignee(ctx, id).UserIdDto(userIdDto).Execute()

Set Assignee



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
	id := "id_example" // string | The id of the task to set the assignee for.
	userIdDto := *openapiclient.NewUserIdDto() // UserIdDto | Provide the id of the user that will be the assignee of the task. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.SetAssignee(context.Background(), id).UserIdDto(userIdDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.SetAssignee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to set the assignee for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAssigneeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userIdDto** | [**UserIdDto**](UserIdDto.md) | Provide the id of the user that will be the assignee of the task. | 

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


## Submit

> map[string]VariableValueDto Submit(ctx, id).CompleteTaskDto(completeTaskDto).Execute()

Submit Form



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
	id := "id_example" // string | The id of the task to submit the form for.
	completeTaskDto := *openapiclient.NewCompleteTaskDto() // CompleteTaskDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.Submit(context.Background(), id).CompleteTaskDto(completeTaskDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.Submit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Submit`: map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.Submit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to submit the form for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **completeTaskDto** | [**CompleteTaskDto**](CompleteTaskDto.md) |  | 

### Return type

[**map[string]VariableValueDto**](VariableValueDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unclaim

> Unclaim(ctx, id).Execute()

Unclaim



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
	id := "id_example" // string | The id of the task to unclaim.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.Unclaim(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.Unclaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to unclaim. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnclaimRequest struct via the builder pattern


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


## UpdateTask

> UpdateTask(ctx, id).TaskDto(taskDto).Execute()

Update



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
	id := "id_example" // string | The id of the task to be updated.
	taskDto := *openapiclient.NewTaskDto() // TaskDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAPI.UpdateTask(context.Background(), id).TaskDto(taskDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.UpdateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskDto** | [**TaskDto**](TaskDto.md) |  | 

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

