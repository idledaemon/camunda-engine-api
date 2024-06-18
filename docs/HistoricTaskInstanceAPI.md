# \HistoricTaskInstanceAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricTaskInstanceReport**](HistoricTaskInstanceAPI.md#GetHistoricTaskInstanceReport) | **Get** /history/task/report | Get Task Report (Historic)
[**GetHistoricTaskInstances**](HistoricTaskInstanceAPI.md#GetHistoricTaskInstances) | **Get** /history/task | Get Tasks (Historic)
[**GetHistoricTaskInstancesCount**](HistoricTaskInstanceAPI.md#GetHistoricTaskInstancesCount) | **Get** /history/task/count | Get Task Count
[**QueryHistoricTaskInstances**](HistoricTaskInstanceAPI.md#QueryHistoricTaskInstances) | **Post** /history/task | Get Tasks (Historic) (POST)
[**QueryHistoricTaskInstancesCount**](HistoricTaskInstanceAPI.md#QueryHistoricTaskInstancesCount) | **Post** /history/task/count | Get Task Count (POST)



## GetHistoricTaskInstanceReport

> []HistoricTaskInstanceReportResultDto GetHistoricTaskInstanceReport(ctx).ReportType(reportType).PeriodUnit(periodUnit).CompletedBefore(completedBefore).CompletedAfter(completedAfter).GroupBy(groupBy).Execute()

Get Task Report (Historic)



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
	reportType := "reportType_example" // string | **Mandatory.** Specifies the kind of the report to execute. To retrieve a report about the duration of process instances the value must be set to `duration`. For a report of the completed tasks in a specific timespan the value must be set to `count`. (optional)
	periodUnit := "periodUnit_example" // string | When the report type is set to `duration`, this parameter is **mandatory**. Specifies the granularity of the report. Valid values are `month` and `quarter`. (optional)
	completedBefore := time.Now() // time.Time | Restrict to tasks that were completed before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	completedAfter := time.Now() // time.Time | Restrict to tasks that were completed after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	groupBy := "groupBy_example" // string | When the report type is set to `count`, this parameter is **mandatory**. Groups the tasks report by a given criterion. Valid values are `taskName` and `processDefinition`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricTaskInstanceAPI.GetHistoricTaskInstanceReport(context.Background()).ReportType(reportType).PeriodUnit(periodUnit).CompletedBefore(completedBefore).CompletedAfter(completedAfter).GroupBy(groupBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricTaskInstanceAPI.GetHistoricTaskInstanceReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricTaskInstanceReport`: []HistoricTaskInstanceReportResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricTaskInstanceAPI.GetHistoricTaskInstanceReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricTaskInstanceReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportType** | **string** | **Mandatory.** Specifies the kind of the report to execute. To retrieve a report about the duration of process instances the value must be set to &#x60;duration&#x60;. For a report of the completed tasks in a specific timespan the value must be set to &#x60;count&#x60;. | 
 **periodUnit** | **string** | When the report type is set to &#x60;duration&#x60;, this parameter is **mandatory**. Specifies the granularity of the report. Valid values are &#x60;month&#x60; and &#x60;quarter&#x60;. | 
 **completedBefore** | **time.Time** | Restrict to tasks that were completed before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **completedAfter** | **time.Time** | Restrict to tasks that were completed after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **groupBy** | **string** | When the report type is set to &#x60;count&#x60;, this parameter is **mandatory**. Groups the tasks report by a given criterion. Valid values are &#x60;taskName&#x60; and &#x60;processDefinition&#x60;. | 

### Return type

[**[]HistoricTaskInstanceReportResultDto**](HistoricTaskInstanceReportResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricTaskInstances

> []HistoricTaskInstanceDto GetHistoricTaskInstances(ctx).TaskId(taskId).TaskParentTaskId(taskParentTaskId).ProcessInstanceId(processInstanceId).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).ExecutionId(executionId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionName(processDefinitionName).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseDefinitionName(caseDefinitionName).ActivityInstanceIdIn(activityInstanceIdIn).TaskName(taskName).TaskNameLike(taskNameLike).TaskDescription(taskDescription).TaskDescriptionLike(taskDescriptionLike).TaskDefinitionKey(taskDefinitionKey).TaskDefinitionKeyIn(taskDefinitionKeyIn).TaskDeleteReason(taskDeleteReason).TaskDeleteReasonLike(taskDeleteReasonLike).TaskAssignee(taskAssignee).TaskAssigneeLike(taskAssigneeLike).TaskOwner(taskOwner).TaskOwnerLike(taskOwnerLike).TaskPriority(taskPriority).Assigned(assigned).Unassigned(unassigned).Finished(finished).Unfinished(unfinished).ProcessFinished(processFinished).ProcessUnfinished(processUnfinished).TaskDueDate(taskDueDate).TaskDueDateBefore(taskDueDateBefore).TaskDueDateAfter(taskDueDateAfter).WithoutTaskDueDate(withoutTaskDueDate).TaskFollowUpDate(taskFollowUpDate).TaskFollowUpDateBefore(taskFollowUpDateBefore).TaskFollowUpDateAfter(taskFollowUpDateAfter).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).TaskVariables(taskVariables).ProcessVariables(processVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).TaskInvolvedUser(taskInvolvedUser).TaskInvolvedGroup(taskInvolvedGroup).TaskHadCandidateUser(taskHadCandidateUser).TaskHadCandidateGroup(taskHadCandidateGroup).WithCandidateGroups(withCandidateGroups).WithoutCandidateGroups(withoutCandidateGroups).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Tasks (Historic)



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
	taskId := "taskId_example" // string | Filter by task id. (optional)
	taskParentTaskId := "taskParentTaskId_example" // string | Filter by parent task id. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processInstanceBusinessKey := "processInstanceBusinessKey_example" // string | Filter by process instance business key. (optional)
	processInstanceBusinessKeyIn := "processInstanceBusinessKeyIn_example" // string | Filter by process instances with one of the give business keys. The keys need to be in a comma-separated list. (optional)
	processInstanceBusinessKeyLike := "processInstanceBusinessKeyLike_example" // string | Filter by  process instance business key that has the parameter value as a substring. (optional)
	executionId := "executionId_example" // string | Filter by the id of the execution that executed the task. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Restrict to tasks that belong to a process definition with the given key. (optional)
	processDefinitionName := "processDefinitionName_example" // string | Restrict to tasks that belong to a process definition with the given name. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by case instance id. (optional)
	caseExecutionId := "caseExecutionId_example" // string | Filter by the id of the case execution that executed the task. (optional)
	caseDefinitionId := "caseDefinitionId_example" // string | Filter by case definition id. (optional)
	caseDefinitionKey := "caseDefinitionKey_example" // string | Restrict to tasks that belong to a case definition with the given key. (optional)
	caseDefinitionName := "caseDefinitionName_example" // string | Restrict to tasks that belong to a case definition with the given name. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Only include tasks which belong to one of the passed comma-separated activity instance ids. (optional)
	taskName := "taskName_example" // string | Restrict to tasks that have the given name. (optional)
	taskNameLike := "taskNameLike_example" // string | Restrict to tasks that have a name with the given parameter value as substring. (optional)
	taskDescription := "taskDescription_example" // string | Restrict to tasks that have the given description. (optional)
	taskDescriptionLike := "taskDescriptionLike_example" // string | Restrict to tasks that have a description that has the parameter value as a substring. (optional)
	taskDefinitionKey := "taskDefinitionKey_example" // string | Restrict to tasks that have the given key. (optional)
	taskDefinitionKeyIn := "taskDefinitionKeyIn_example" // string | Restrict to tasks that have one of the passed comma-separated task definition keys. (optional)
	taskDeleteReason := "taskDeleteReason_example" // string | Restrict to tasks that have the given delete reason. (optional)
	taskDeleteReasonLike := "taskDeleteReasonLike_example" // string | Restrict to tasks that have a delete reason that has the parameter value as a substring. (optional)
	taskAssignee := "taskAssignee_example" // string | Restrict to tasks that the given user is assigned to. (optional)
	taskAssigneeLike := "taskAssigneeLike_example" // string | Restrict to tasks that are assigned to users with the parameter value as a substring. (optional)
	taskOwner := "taskOwner_example" // string | Restrict to tasks that the given user owns. (optional)
	taskOwnerLike := "taskOwnerLike_example" // string | Restrict to tasks that are owned by users with the parameter value as a substring. (optional)
	taskPriority := int32(56) // int32 | Restrict to tasks that have the given priority. (optional)
	assigned := true // bool | If set to `true`, restricts the query to all tasks that are assigned. (optional)
	unassigned := true // bool | If set to `true`, restricts the query to all tasks that are unassigned. (optional)
	finished := true // bool | Only include finished tasks. Value may only be `true`, as `false` is the default behavior. (optional)
	unfinished := true // bool | Only include unfinished tasks. Value may only be `true`, as `false` is the default behavior. (optional)
	processFinished := true // bool | Only include tasks of finished processes. Value may only be `true`, as `false` is the default behavior. (optional)
	processUnfinished := true // bool | Only include tasks of unfinished processes. Value may only be `true`, as `false` is the default behavior. (optional)
	taskDueDate := time.Now() // time.Time | Restrict to tasks that are due on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	taskDueDateBefore := time.Now() // time.Time | Restrict to tasks that are due before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	taskDueDateAfter := time.Now() // time.Time | Restrict to tasks that are due after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	withoutTaskDueDate := true // bool | Only include tasks which have no due date. Value may only be `true`, as `false` is the default behavior. (optional)
	taskFollowUpDate := time.Now() // time.Time | Restrict to tasks that have a followUp date on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	taskFollowUpDateBefore := time.Now() // time.Time | Restrict to tasks that have a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	taskFollowUpDateAfter := time.Now() // time.Time | Restrict to tasks that have a followUp date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedBefore := time.Now() // time.Time | Restrict to tasks that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Restrict to tasks that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedBefore := time.Now() // time.Time | Restrict to tasks that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedAfter := time.Now() // time.Time | Restrict to tasks that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A task instance must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic task instances that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	taskVariables := "taskVariables_example" // string | Only include tasks that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value. **Note:** Values are always treated as `String` objects on server side.   Valid operator values are: * `eq` - equal to; * `neq` - not equal to; * `gt` - greater than; * `gteq` - greater than or equal to; * `lt` - lower than; * `lteq` - lower than or equal to; * `like`.  `key` and `value` may not contain underscore or comma characters. (optional)
	processVariables := "processVariables_example" // string | Only include tasks that belong to process instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value. **Note:** Values are always treated as `String` objects on server side.   Valid operator values are: * `eq` - equal to; * `neq` - not equal to; * `gt` - greater than; * `gteq` - greater than or equal to; * `lt` - lower than; * `lteq` - lower than or equal to; * `like`; * `notLike`.  `key` and `value` may not contain underscore or comma characters. (optional)
	variableNamesIgnoreCase := true // bool | Match the variable name provided in `taskVariables` and `processVariables` case- insensitively. If set to `true` **variableName** and **variablename** are treated as equal. (optional)
	variableValuesIgnoreCase := true // bool | Match the variable value provided in `taskVariables` and `processVariables` case- insensitively. If set to `true` **variableValue** and **variablevalue** are treated as equal. (optional)
	taskInvolvedUser := "taskInvolvedUser_example" // string | Restrict to tasks with a historic identity link to the given user. (optional)
	taskInvolvedGroup := "taskInvolvedGroup_example" // string | Restrict to tasks with a historic identity link to the given group. (optional)
	taskHadCandidateUser := "taskHadCandidateUser_example" // string | Restrict to tasks with a historic identity link to the given candidate user. (optional)
	taskHadCandidateGroup := "taskHadCandidateGroup_example" // string | Restrict to tasks with a historic identity link to the given candidate group. (optional)
	withCandidateGroups := true // bool | Only include tasks which have a candidate group. Value may only be `true`, as `false` is the default behavior. (optional)
	withoutCandidateGroups := true // bool | Only include tasks which have no candidate group. Value may only be `true`, as `false` is the default behavior. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricTaskInstanceAPI.GetHistoricTaskInstances(context.Background()).TaskId(taskId).TaskParentTaskId(taskParentTaskId).ProcessInstanceId(processInstanceId).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).ExecutionId(executionId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionName(processDefinitionName).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseDefinitionName(caseDefinitionName).ActivityInstanceIdIn(activityInstanceIdIn).TaskName(taskName).TaskNameLike(taskNameLike).TaskDescription(taskDescription).TaskDescriptionLike(taskDescriptionLike).TaskDefinitionKey(taskDefinitionKey).TaskDefinitionKeyIn(taskDefinitionKeyIn).TaskDeleteReason(taskDeleteReason).TaskDeleteReasonLike(taskDeleteReasonLike).TaskAssignee(taskAssignee).TaskAssigneeLike(taskAssigneeLike).TaskOwner(taskOwner).TaskOwnerLike(taskOwnerLike).TaskPriority(taskPriority).Assigned(assigned).Unassigned(unassigned).Finished(finished).Unfinished(unfinished).ProcessFinished(processFinished).ProcessUnfinished(processUnfinished).TaskDueDate(taskDueDate).TaskDueDateBefore(taskDueDateBefore).TaskDueDateAfter(taskDueDateAfter).WithoutTaskDueDate(withoutTaskDueDate).TaskFollowUpDate(taskFollowUpDate).TaskFollowUpDateBefore(taskFollowUpDateBefore).TaskFollowUpDateAfter(taskFollowUpDateAfter).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).TaskVariables(taskVariables).ProcessVariables(processVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).TaskInvolvedUser(taskInvolvedUser).TaskInvolvedGroup(taskInvolvedGroup).TaskHadCandidateUser(taskHadCandidateUser).TaskHadCandidateGroup(taskHadCandidateGroup).WithCandidateGroups(withCandidateGroups).WithoutCandidateGroups(withoutCandidateGroups).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricTaskInstanceAPI.GetHistoricTaskInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricTaskInstances`: []HistoricTaskInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricTaskInstanceAPI.GetHistoricTaskInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricTaskInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **string** | Filter by task id. | 
 **taskParentTaskId** | **string** | Filter by parent task id. | 
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processInstanceBusinessKey** | **string** | Filter by process instance business key. | 
 **processInstanceBusinessKeyIn** | **string** | Filter by process instances with one of the give business keys. The keys need to be in a comma-separated list. | 
 **processInstanceBusinessKeyLike** | **string** | Filter by  process instance business key that has the parameter value as a substring. | 
 **executionId** | **string** | Filter by the id of the execution that executed the task. | 
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **processDefinitionKey** | **string** | Restrict to tasks that belong to a process definition with the given key. | 
 **processDefinitionName** | **string** | Restrict to tasks that belong to a process definition with the given name. | 
 **caseInstanceId** | **string** | Filter by case instance id. | 
 **caseExecutionId** | **string** | Filter by the id of the case execution that executed the task. | 
 **caseDefinitionId** | **string** | Filter by case definition id. | 
 **caseDefinitionKey** | **string** | Restrict to tasks that belong to a case definition with the given key. | 
 **caseDefinitionName** | **string** | Restrict to tasks that belong to a case definition with the given name. | 
 **activityInstanceIdIn** | **string** | Only include tasks which belong to one of the passed comma-separated activity instance ids. | 
 **taskName** | **string** | Restrict to tasks that have the given name. | 
 **taskNameLike** | **string** | Restrict to tasks that have a name with the given parameter value as substring. | 
 **taskDescription** | **string** | Restrict to tasks that have the given description. | 
 **taskDescriptionLike** | **string** | Restrict to tasks that have a description that has the parameter value as a substring. | 
 **taskDefinitionKey** | **string** | Restrict to tasks that have the given key. | 
 **taskDefinitionKeyIn** | **string** | Restrict to tasks that have one of the passed comma-separated task definition keys. | 
 **taskDeleteReason** | **string** | Restrict to tasks that have the given delete reason. | 
 **taskDeleteReasonLike** | **string** | Restrict to tasks that have a delete reason that has the parameter value as a substring. | 
 **taskAssignee** | **string** | Restrict to tasks that the given user is assigned to. | 
 **taskAssigneeLike** | **string** | Restrict to tasks that are assigned to users with the parameter value as a substring. | 
 **taskOwner** | **string** | Restrict to tasks that the given user owns. | 
 **taskOwnerLike** | **string** | Restrict to tasks that are owned by users with the parameter value as a substring. | 
 **taskPriority** | **int32** | Restrict to tasks that have the given priority. | 
 **assigned** | **bool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are assigned. | 
 **unassigned** | **bool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are unassigned. | 
 **finished** | **bool** | Only include finished tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **unfinished** | **bool** | Only include unfinished tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **processFinished** | **bool** | Only include tasks of finished processes. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **processUnfinished** | **bool** | Only include tasks of unfinished processes. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **taskDueDate** | **time.Time** | Restrict to tasks that are due on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **taskDueDateBefore** | **time.Time** | Restrict to tasks that are due before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **taskDueDateAfter** | **time.Time** | Restrict to tasks that are due after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **withoutTaskDueDate** | **bool** | Only include tasks which have no due date. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **taskFollowUpDate** | **time.Time** | Restrict to tasks that have a followUp date on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **taskFollowUpDateBefore** | **time.Time** | Restrict to tasks that have a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **taskFollowUpDateAfter** | **time.Time** | Restrict to tasks that have a followUp date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedBefore** | **time.Time** | Restrict to tasks that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Restrict to tasks that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedBefore** | **time.Time** | Restrict to tasks that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedAfter** | **time.Time** | Restrict to tasks that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A task instance must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic task instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **taskVariables** | **string** | Only include tasks that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. **Note:** Values are always treated as &#x60;String&#x60; objects on server side.   Valid operator values are: * &#x60;eq&#x60; - equal to; * &#x60;neq&#x60; - not equal to; * &#x60;gt&#x60; - greater than; * &#x60;gteq&#x60; - greater than or equal to; * &#x60;lt&#x60; - lower than; * &#x60;lteq&#x60; - lower than or equal to; * &#x60;like&#x60;.  &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **processVariables** | **string** | Only include tasks that belong to process instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. **Note:** Values are always treated as &#x60;String&#x60; objects on server side.   Valid operator values are: * &#x60;eq&#x60; - equal to; * &#x60;neq&#x60; - not equal to; * &#x60;gt&#x60; - greater than; * &#x60;gteq&#x60; - greater than or equal to; * &#x60;lt&#x60; - lower than; * &#x60;lteq&#x60; - lower than or equal to; * &#x60;like&#x60;; * &#x60;notLike&#x60;.  &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **variableNamesIgnoreCase** | **bool** | Match the variable name provided in &#x60;taskVariables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | 
 **variableValuesIgnoreCase** | **bool** | Match the variable value provided in &#x60;taskVariables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | 
 **taskInvolvedUser** | **string** | Restrict to tasks with a historic identity link to the given user. | 
 **taskInvolvedGroup** | **string** | Restrict to tasks with a historic identity link to the given group. | 
 **taskHadCandidateUser** | **string** | Restrict to tasks with a historic identity link to the given candidate user. | 
 **taskHadCandidateGroup** | **string** | Restrict to tasks with a historic identity link to the given candidate group. | 
 **withCandidateGroups** | **bool** | Only include tasks which have a candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **withoutCandidateGroups** | **bool** | Only include tasks which have no candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]HistoricTaskInstanceDto**](HistoricTaskInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricTaskInstancesCount

> CountResultDto GetHistoricTaskInstancesCount(ctx).TaskId(taskId).TaskParentTaskId(taskParentTaskId).ProcessInstanceId(processInstanceId).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).ExecutionId(executionId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionName(processDefinitionName).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseDefinitionName(caseDefinitionName).ActivityInstanceIdIn(activityInstanceIdIn).TaskName(taskName).TaskNameLike(taskNameLike).TaskDescription(taskDescription).TaskDescriptionLike(taskDescriptionLike).TaskDefinitionKey(taskDefinitionKey).TaskDefinitionKeyIn(taskDefinitionKeyIn).TaskDeleteReason(taskDeleteReason).TaskDeleteReasonLike(taskDeleteReasonLike).TaskAssignee(taskAssignee).TaskAssigneeLike(taskAssigneeLike).TaskOwner(taskOwner).TaskOwnerLike(taskOwnerLike).TaskPriority(taskPriority).Assigned(assigned).Unassigned(unassigned).Finished(finished).Unfinished(unfinished).ProcessFinished(processFinished).ProcessUnfinished(processUnfinished).TaskDueDate(taskDueDate).TaskDueDateBefore(taskDueDateBefore).TaskDueDateAfter(taskDueDateAfter).WithoutTaskDueDate(withoutTaskDueDate).TaskFollowUpDate(taskFollowUpDate).TaskFollowUpDateBefore(taskFollowUpDateBefore).TaskFollowUpDateAfter(taskFollowUpDateAfter).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).TaskVariables(taskVariables).ProcessVariables(processVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).TaskInvolvedUser(taskInvolvedUser).TaskInvolvedGroup(taskInvolvedGroup).TaskHadCandidateUser(taskHadCandidateUser).TaskHadCandidateGroup(taskHadCandidateGroup).WithCandidateGroups(withCandidateGroups).WithoutCandidateGroups(withoutCandidateGroups).Execute()

Get Task Count



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
	taskId := "taskId_example" // string | Filter by task id. (optional)
	taskParentTaskId := "taskParentTaskId_example" // string | Filter by parent task id. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processInstanceBusinessKey := "processInstanceBusinessKey_example" // string | Filter by process instance business key. (optional)
	processInstanceBusinessKeyIn := "processInstanceBusinessKeyIn_example" // string | Filter by process instances with one of the give business keys. The keys need to be in a comma-separated list. (optional)
	processInstanceBusinessKeyLike := "processInstanceBusinessKeyLike_example" // string | Filter by  process instance business key that has the parameter value as a substring. (optional)
	executionId := "executionId_example" // string | Filter by the id of the execution that executed the task. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Restrict to tasks that belong to a process definition with the given key. (optional)
	processDefinitionName := "processDefinitionName_example" // string | Restrict to tasks that belong to a process definition with the given name. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by case instance id. (optional)
	caseExecutionId := "caseExecutionId_example" // string | Filter by the id of the case execution that executed the task. (optional)
	caseDefinitionId := "caseDefinitionId_example" // string | Filter by case definition id. (optional)
	caseDefinitionKey := "caseDefinitionKey_example" // string | Restrict to tasks that belong to a case definition with the given key. (optional)
	caseDefinitionName := "caseDefinitionName_example" // string | Restrict to tasks that belong to a case definition with the given name. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Only include tasks which belong to one of the passed comma-separated activity instance ids. (optional)
	taskName := "taskName_example" // string | Restrict to tasks that have the given name. (optional)
	taskNameLike := "taskNameLike_example" // string | Restrict to tasks that have a name with the given parameter value as substring. (optional)
	taskDescription := "taskDescription_example" // string | Restrict to tasks that have the given description. (optional)
	taskDescriptionLike := "taskDescriptionLike_example" // string | Restrict to tasks that have a description that has the parameter value as a substring. (optional)
	taskDefinitionKey := "taskDefinitionKey_example" // string | Restrict to tasks that have the given key. (optional)
	taskDefinitionKeyIn := "taskDefinitionKeyIn_example" // string | Restrict to tasks that have one of the passed comma-separated task definition keys. (optional)
	taskDeleteReason := "taskDeleteReason_example" // string | Restrict to tasks that have the given delete reason. (optional)
	taskDeleteReasonLike := "taskDeleteReasonLike_example" // string | Restrict to tasks that have a delete reason that has the parameter value as a substring. (optional)
	taskAssignee := "taskAssignee_example" // string | Restrict to tasks that the given user is assigned to. (optional)
	taskAssigneeLike := "taskAssigneeLike_example" // string | Restrict to tasks that are assigned to users with the parameter value as a substring. (optional)
	taskOwner := "taskOwner_example" // string | Restrict to tasks that the given user owns. (optional)
	taskOwnerLike := "taskOwnerLike_example" // string | Restrict to tasks that are owned by users with the parameter value as a substring. (optional)
	taskPriority := int32(56) // int32 | Restrict to tasks that have the given priority. (optional)
	assigned := true // bool | If set to `true`, restricts the query to all tasks that are assigned. (optional)
	unassigned := true // bool | If set to `true`, restricts the query to all tasks that are unassigned. (optional)
	finished := true // bool | Only include finished tasks. Value may only be `true`, as `false` is the default behavior. (optional)
	unfinished := true // bool | Only include unfinished tasks. Value may only be `true`, as `false` is the default behavior. (optional)
	processFinished := true // bool | Only include tasks of finished processes. Value may only be `true`, as `false` is the default behavior. (optional)
	processUnfinished := true // bool | Only include tasks of unfinished processes. Value may only be `true`, as `false` is the default behavior. (optional)
	taskDueDate := time.Now() // time.Time | Restrict to tasks that are due on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	taskDueDateBefore := time.Now() // time.Time | Restrict to tasks that are due before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	taskDueDateAfter := time.Now() // time.Time | Restrict to tasks that are due after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	withoutTaskDueDate := true // bool | Only include tasks which have no due date. Value may only be `true`, as `false` is the default behavior. (optional)
	taskFollowUpDate := time.Now() // time.Time | Restrict to tasks that have a followUp date on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	taskFollowUpDateBefore := time.Now() // time.Time | Restrict to tasks that have a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	taskFollowUpDateAfter := time.Now() // time.Time | Restrict to tasks that have a followUp date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedBefore := time.Now() // time.Time | Restrict to tasks that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Restrict to tasks that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedBefore := time.Now() // time.Time | Restrict to tasks that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedAfter := time.Now() // time.Time | Restrict to tasks that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A task instance must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic task instances that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	taskVariables := "taskVariables_example" // string | Only include tasks that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value. **Note:** Values are always treated as `String` objects on server side.   Valid operator values are: * `eq` - equal to; * `neq` - not equal to; * `gt` - greater than; * `gteq` - greater than or equal to; * `lt` - lower than; * `lteq` - lower than or equal to; * `like`.  `key` and `value` may not contain underscore or comma characters. (optional)
	processVariables := "processVariables_example" // string | Only include tasks that belong to process instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value. **Note:** Values are always treated as `String` objects on server side.   Valid operator values are: * `eq` - equal to; * `neq` - not equal to; * `gt` - greater than; * `gteq` - greater than or equal to; * `lt` - lower than; * `lteq` - lower than or equal to; * `like`; * `notLike`.  `key` and `value` may not contain underscore or comma characters. (optional)
	variableNamesIgnoreCase := true // bool | Match the variable name provided in `taskVariables` and `processVariables` case- insensitively. If set to `true` **variableName** and **variablename** are treated as equal. (optional)
	variableValuesIgnoreCase := true // bool | Match the variable value provided in `taskVariables` and `processVariables` case- insensitively. If set to `true` **variableValue** and **variablevalue** are treated as equal. (optional)
	taskInvolvedUser := "taskInvolvedUser_example" // string | Restrict to tasks with a historic identity link to the given user. (optional)
	taskInvolvedGroup := "taskInvolvedGroup_example" // string | Restrict to tasks with a historic identity link to the given group. (optional)
	taskHadCandidateUser := "taskHadCandidateUser_example" // string | Restrict to tasks with a historic identity link to the given candidate user. (optional)
	taskHadCandidateGroup := "taskHadCandidateGroup_example" // string | Restrict to tasks with a historic identity link to the given candidate group. (optional)
	withCandidateGroups := true // bool | Only include tasks which have a candidate group. Value may only be `true`, as `false` is the default behavior. (optional)
	withoutCandidateGroups := true // bool | Only include tasks which have no candidate group. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricTaskInstanceAPI.GetHistoricTaskInstancesCount(context.Background()).TaskId(taskId).TaskParentTaskId(taskParentTaskId).ProcessInstanceId(processInstanceId).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).ExecutionId(executionId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionName(processDefinitionName).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).CaseDefinitionId(caseDefinitionId).CaseDefinitionKey(caseDefinitionKey).CaseDefinitionName(caseDefinitionName).ActivityInstanceIdIn(activityInstanceIdIn).TaskName(taskName).TaskNameLike(taskNameLike).TaskDescription(taskDescription).TaskDescriptionLike(taskDescriptionLike).TaskDefinitionKey(taskDefinitionKey).TaskDefinitionKeyIn(taskDefinitionKeyIn).TaskDeleteReason(taskDeleteReason).TaskDeleteReasonLike(taskDeleteReasonLike).TaskAssignee(taskAssignee).TaskAssigneeLike(taskAssigneeLike).TaskOwner(taskOwner).TaskOwnerLike(taskOwnerLike).TaskPriority(taskPriority).Assigned(assigned).Unassigned(unassigned).Finished(finished).Unfinished(unfinished).ProcessFinished(processFinished).ProcessUnfinished(processUnfinished).TaskDueDate(taskDueDate).TaskDueDateBefore(taskDueDateBefore).TaskDueDateAfter(taskDueDateAfter).WithoutTaskDueDate(withoutTaskDueDate).TaskFollowUpDate(taskFollowUpDate).TaskFollowUpDateBefore(taskFollowUpDateBefore).TaskFollowUpDateAfter(taskFollowUpDateAfter).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).TaskVariables(taskVariables).ProcessVariables(processVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).TaskInvolvedUser(taskInvolvedUser).TaskInvolvedGroup(taskInvolvedGroup).TaskHadCandidateUser(taskHadCandidateUser).TaskHadCandidateGroup(taskHadCandidateGroup).WithCandidateGroups(withCandidateGroups).WithoutCandidateGroups(withoutCandidateGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricTaskInstanceAPI.GetHistoricTaskInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricTaskInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricTaskInstanceAPI.GetHistoricTaskInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricTaskInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **string** | Filter by task id. | 
 **taskParentTaskId** | **string** | Filter by parent task id. | 
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processInstanceBusinessKey** | **string** | Filter by process instance business key. | 
 **processInstanceBusinessKeyIn** | **string** | Filter by process instances with one of the give business keys. The keys need to be in a comma-separated list. | 
 **processInstanceBusinessKeyLike** | **string** | Filter by  process instance business key that has the parameter value as a substring. | 
 **executionId** | **string** | Filter by the id of the execution that executed the task. | 
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **processDefinitionKey** | **string** | Restrict to tasks that belong to a process definition with the given key. | 
 **processDefinitionName** | **string** | Restrict to tasks that belong to a process definition with the given name. | 
 **caseInstanceId** | **string** | Filter by case instance id. | 
 **caseExecutionId** | **string** | Filter by the id of the case execution that executed the task. | 
 **caseDefinitionId** | **string** | Filter by case definition id. | 
 **caseDefinitionKey** | **string** | Restrict to tasks that belong to a case definition with the given key. | 
 **caseDefinitionName** | **string** | Restrict to tasks that belong to a case definition with the given name. | 
 **activityInstanceIdIn** | **string** | Only include tasks which belong to one of the passed comma-separated activity instance ids. | 
 **taskName** | **string** | Restrict to tasks that have the given name. | 
 **taskNameLike** | **string** | Restrict to tasks that have a name with the given parameter value as substring. | 
 **taskDescription** | **string** | Restrict to tasks that have the given description. | 
 **taskDescriptionLike** | **string** | Restrict to tasks that have a description that has the parameter value as a substring. | 
 **taskDefinitionKey** | **string** | Restrict to tasks that have the given key. | 
 **taskDefinitionKeyIn** | **string** | Restrict to tasks that have one of the passed comma-separated task definition keys. | 
 **taskDeleteReason** | **string** | Restrict to tasks that have the given delete reason. | 
 **taskDeleteReasonLike** | **string** | Restrict to tasks that have a delete reason that has the parameter value as a substring. | 
 **taskAssignee** | **string** | Restrict to tasks that the given user is assigned to. | 
 **taskAssigneeLike** | **string** | Restrict to tasks that are assigned to users with the parameter value as a substring. | 
 **taskOwner** | **string** | Restrict to tasks that the given user owns. | 
 **taskOwnerLike** | **string** | Restrict to tasks that are owned by users with the parameter value as a substring. | 
 **taskPriority** | **int32** | Restrict to tasks that have the given priority. | 
 **assigned** | **bool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are assigned. | 
 **unassigned** | **bool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are unassigned. | 
 **finished** | **bool** | Only include finished tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **unfinished** | **bool** | Only include unfinished tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **processFinished** | **bool** | Only include tasks of finished processes. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **processUnfinished** | **bool** | Only include tasks of unfinished processes. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **taskDueDate** | **time.Time** | Restrict to tasks that are due on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **taskDueDateBefore** | **time.Time** | Restrict to tasks that are due before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **taskDueDateAfter** | **time.Time** | Restrict to tasks that are due after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **withoutTaskDueDate** | **bool** | Only include tasks which have no due date. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **taskFollowUpDate** | **time.Time** | Restrict to tasks that have a followUp date on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **taskFollowUpDateBefore** | **time.Time** | Restrict to tasks that have a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **taskFollowUpDateAfter** | **time.Time** | Restrict to tasks that have a followUp date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedBefore** | **time.Time** | Restrict to tasks that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Restrict to tasks that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedBefore** | **time.Time** | Restrict to tasks that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedAfter** | **time.Time** | Restrict to tasks that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A task instance must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic task instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **taskVariables** | **string** | Only include tasks that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. **Note:** Values are always treated as &#x60;String&#x60; objects on server side.   Valid operator values are: * &#x60;eq&#x60; - equal to; * &#x60;neq&#x60; - not equal to; * &#x60;gt&#x60; - greater than; * &#x60;gteq&#x60; - greater than or equal to; * &#x60;lt&#x60; - lower than; * &#x60;lteq&#x60; - lower than or equal to; * &#x60;like&#x60;.  &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **processVariables** | **string** | Only include tasks that belong to process instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. **Note:** Values are always treated as &#x60;String&#x60; objects on server side.   Valid operator values are: * &#x60;eq&#x60; - equal to; * &#x60;neq&#x60; - not equal to; * &#x60;gt&#x60; - greater than; * &#x60;gteq&#x60; - greater than or equal to; * &#x60;lt&#x60; - lower than; * &#x60;lteq&#x60; - lower than or equal to; * &#x60;like&#x60;; * &#x60;notLike&#x60;.  &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **variableNamesIgnoreCase** | **bool** | Match the variable name provided in &#x60;taskVariables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | 
 **variableValuesIgnoreCase** | **bool** | Match the variable value provided in &#x60;taskVariables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | 
 **taskInvolvedUser** | **string** | Restrict to tasks with a historic identity link to the given user. | 
 **taskInvolvedGroup** | **string** | Restrict to tasks with a historic identity link to the given group. | 
 **taskHadCandidateUser** | **string** | Restrict to tasks with a historic identity link to the given candidate user. | 
 **taskHadCandidateGroup** | **string** | Restrict to tasks with a historic identity link to the given candidate group. | 
 **withCandidateGroups** | **bool** | Only include tasks which have a candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **withoutCandidateGroups** | **bool** | Only include tasks which have no candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

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


## QueryHistoricTaskInstances

> []HistoricTaskInstanceDto QueryHistoricTaskInstances(ctx).FirstResult(firstResult).MaxResults(maxResults).HistoricTaskInstanceQueryDto(historicTaskInstanceQueryDto).Execute()

Get Tasks (Historic) (POST)



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
	historicTaskInstanceQueryDto := *openapiclient.NewHistoricTaskInstanceQueryDto() // HistoricTaskInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricTaskInstanceAPI.QueryHistoricTaskInstances(context.Background()).FirstResult(firstResult).MaxResults(maxResults).HistoricTaskInstanceQueryDto(historicTaskInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricTaskInstanceAPI.QueryHistoricTaskInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricTaskInstances`: []HistoricTaskInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricTaskInstanceAPI.QueryHistoricTaskInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricTaskInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **historicTaskInstanceQueryDto** | [**HistoricTaskInstanceQueryDto**](HistoricTaskInstanceQueryDto.md) |  | 

### Return type

[**[]HistoricTaskInstanceDto**](HistoricTaskInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryHistoricTaskInstancesCount

> CountResultDto QueryHistoricTaskInstancesCount(ctx).HistoricTaskInstanceQueryDto(historicTaskInstanceQueryDto).Execute()

Get Task Count (POST)



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
	historicTaskInstanceQueryDto := *openapiclient.NewHistoricTaskInstanceQueryDto() // HistoricTaskInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricTaskInstanceAPI.QueryHistoricTaskInstancesCount(context.Background()).HistoricTaskInstanceQueryDto(historicTaskInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricTaskInstanceAPI.QueryHistoricTaskInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricTaskInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricTaskInstanceAPI.QueryHistoricTaskInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricTaskInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **historicTaskInstanceQueryDto** | [**HistoricTaskInstanceQueryDto**](HistoricTaskInstanceQueryDto.md) |  | 

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

