# \HistoricProcessInstanceAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHistoricProcessInstance**](HistoricProcessInstanceAPI.md#DeleteHistoricProcessInstance) | **Delete** /history/process-instance/{id} | Delete
[**DeleteHistoricProcessInstancesAsync**](HistoricProcessInstanceAPI.md#DeleteHistoricProcessInstancesAsync) | **Post** /history/process-instance/delete | Delete Async (POST)
[**DeleteHistoricVariableInstancesOfHistoricProcessInstance**](HistoricProcessInstanceAPI.md#DeleteHistoricVariableInstancesOfHistoricProcessInstance) | **Delete** /history/process-instance/{id}/variable-instances | Delete Variable Instances
[**GetHistoricProcessInstance**](HistoricProcessInstanceAPI.md#GetHistoricProcessInstance) | **Get** /history/process-instance/{id} | Get
[**GetHistoricProcessInstanceDurationReport**](HistoricProcessInstanceAPI.md#GetHistoricProcessInstanceDurationReport) | **Get** /history/process-instance/report | Get Duration Report
[**GetHistoricProcessInstances**](HistoricProcessInstanceAPI.md#GetHistoricProcessInstances) | **Get** /history/process-instance | Get List
[**GetHistoricProcessInstancesCount**](HistoricProcessInstanceAPI.md#GetHistoricProcessInstancesCount) | **Get** /history/process-instance/count | Get List Count
[**QueryHistoricProcessInstances**](HistoricProcessInstanceAPI.md#QueryHistoricProcessInstances) | **Post** /history/process-instance | Get List (POST)
[**QueryHistoricProcessInstancesCount**](HistoricProcessInstanceAPI.md#QueryHistoricProcessInstancesCount) | **Post** /history/process-instance/count | Get List Count (POST)
[**SetRemovalTimeAsync**](HistoricProcessInstanceAPI.md#SetRemovalTimeAsync) | **Post** /history/process-instance/set-removal-time | Set Removal Time Async (POST)



## DeleteHistoricProcessInstance

> DeleteHistoricProcessInstance(ctx, id).FailIfNotExists(failIfNotExists).Execute()

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
	id := "id_example" // string | The id of the historic process instance to be deleted.
	failIfNotExists := true // bool | If set to `false`, the request will still be successful if the process id is not found. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HistoricProcessInstanceAPI.DeleteHistoricProcessInstance(context.Background(), id).FailIfNotExists(failIfNotExists).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessInstanceAPI.DeleteHistoricProcessInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the historic process instance to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHistoricProcessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **failIfNotExists** | **bool** | If set to &#x60;false&#x60;, the request will still be successful if the process id is not found. | 

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


## DeleteHistoricProcessInstancesAsync

> BatchDto DeleteHistoricProcessInstancesAsync(ctx).DeleteHistoricProcessInstancesDto(deleteHistoricProcessInstancesDto).Execute()

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
	deleteHistoricProcessInstancesDto := *openapiclient.NewDeleteHistoricProcessInstancesDto() // DeleteHistoricProcessInstancesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricProcessInstanceAPI.DeleteHistoricProcessInstancesAsync(context.Background()).DeleteHistoricProcessInstancesDto(deleteHistoricProcessInstancesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessInstanceAPI.DeleteHistoricProcessInstancesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteHistoricProcessInstancesAsync`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricProcessInstanceAPI.DeleteHistoricProcessInstancesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHistoricProcessInstancesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteHistoricProcessInstancesDto** | [**DeleteHistoricProcessInstancesDto**](DeleteHistoricProcessInstancesDto.md) |  | 

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


## DeleteHistoricVariableInstancesOfHistoricProcessInstance

> DeleteHistoricVariableInstancesOfHistoricProcessInstance(ctx, id).Execute()

Delete Variable Instances



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
	id := "id_example" // string | The id of the process instance for which all historic variables are to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HistoricProcessInstanceAPI.DeleteHistoricVariableInstancesOfHistoricProcessInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessInstanceAPI.DeleteHistoricVariableInstancesOfHistoricProcessInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance for which all historic variables are to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHistoricVariableInstancesOfHistoricProcessInstanceRequest struct via the builder pattern


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


## GetHistoricProcessInstance

> HistoricProcessInstanceDto GetHistoricProcessInstance(ctx, id).Execute()

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
	id := "id_example" // string | The id of the historic process instance to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricProcessInstanceAPI.GetHistoricProcessInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessInstanceAPI.GetHistoricProcessInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricProcessInstance`: HistoricProcessInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricProcessInstanceAPI.GetHistoricProcessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the historic process instance to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricProcessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HistoricProcessInstanceDto**](HistoricProcessInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricProcessInstanceDurationReport

> []DurationReportResultDto GetHistoricProcessInstanceDurationReport(ctx).ReportType(reportType).PeriodUnit(periodUnit).ProcessDefinitionIdIn(processDefinitionIdIn).ProcessDefinitionKeyIn(processDefinitionKeyIn).StartedBefore(startedBefore).StartedAfter(startedAfter).Execute()

Get Duration Report



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
	reportType := "reportType_example" // string | **Mandatory.** Specifies the type of the report to retrieve. To retrieve a report about the duration of process instances, the value must be set to `duration`.
	periodUnit := "periodUnit_example" // string | **Mandatory.** Specifies the granularity of the report. Valid values are `month` and `quarter`.
	processDefinitionIdIn := "processDefinitionIdIn_example" // string | Filter by process definition ids. Must be a comma-separated list of process definition ids. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Filter by process definition keys. Must be a comma-separated list of process definition keys. (optional)
	startedBefore := time.Now() // time.Time | Restrict to instances that were started before the given date. By [default](), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2016-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2016-01-23T14:42:45.000+0200`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricProcessInstanceAPI.GetHistoricProcessInstanceDurationReport(context.Background()).ReportType(reportType).PeriodUnit(periodUnit).ProcessDefinitionIdIn(processDefinitionIdIn).ProcessDefinitionKeyIn(processDefinitionKeyIn).StartedBefore(startedBefore).StartedAfter(startedAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessInstanceAPI.GetHistoricProcessInstanceDurationReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricProcessInstanceDurationReport`: []DurationReportResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricProcessInstanceAPI.GetHistoricProcessInstanceDurationReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricProcessInstanceDurationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportType** | **string** | **Mandatory.** Specifies the type of the report to retrieve. To retrieve a report about the duration of process instances, the value must be set to &#x60;duration&#x60;. | 
 **periodUnit** | **string** | **Mandatory.** Specifies the granularity of the report. Valid values are &#x60;month&#x60; and &#x60;quarter&#x60;. | 
 **processDefinitionIdIn** | **string** | Filter by process definition ids. Must be a comma-separated list of process definition ids. | 
 **processDefinitionKeyIn** | **string** | Filter by process definition keys. Must be a comma-separated list of process definition keys. | 
 **startedBefore** | **time.Time** | Restrict to instances that were started before the given date. By [default](), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2016-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2016-01-23T14:42:45.000+0200&#x60;. | 

### Return type

[**[]DurationReportResultDto**](DurationReportResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/csv, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricProcessInstances

> []HistoricProcessInstanceDto GetHistoricProcessInstances(ctx).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).ProcessInstanceId(processInstanceId).ProcessInstanceIds(processInstanceIds).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionName(processDefinitionName).ProcessDefinitionNameLike(processDefinitionNameLike).ProcessDefinitionKeyNotIn(processDefinitionKeyNotIn).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).RootProcessInstances(rootProcessInstances).Finished(finished).Unfinished(unfinished).WithIncidents(withIncidents).WithRootIncidents(withRootIncidents).IncidentType(incidentType).IncidentStatus(incidentStatus).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).ExecutedActivityAfter(executedActivityAfter).ExecutedActivityBefore(executedActivityBefore).ExecutedJobAfter(executedJobAfter).ExecutedJobBefore(executedJobBefore).StartedBy(startedBy).SuperProcessInstanceId(superProcessInstanceId).SubProcessInstanceId(subProcessInstanceId).SuperCaseInstanceId(superCaseInstanceId).SubCaseInstanceId(subCaseInstanceId).CaseInstanceId(caseInstanceId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).ExecutedActivityIdIn(executedActivityIdIn).ActiveActivityIdIn(activeActivityIdIn).Active(active).Suspended(suspended).Completed(completed).ExternallyTerminated(externallyTerminated).InternallyTerminated(internallyTerminated).Variables(variables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).Execute()

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
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processInstanceIds := "processInstanceIds_example" // string | Filter by process instance ids. Filter by a comma-separated list of `Strings`. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the process definition the instances run on. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by the key of the process definition the instances run on. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Filter by a list of process definition keys. A process instance must have one of the given process definition keys. Filter by a comma-separated list of `Strings`. (optional)
	processDefinitionName := "processDefinitionName_example" // string | Filter by the name of the process definition the instances run on. (optional)
	processDefinitionNameLike := "processDefinitionNameLike_example" // string | Filter by process definition names that the parameter is a substring of. (optional)
	processDefinitionKeyNotIn := "processDefinitionKeyNotIn_example" // string | Exclude instances that belong to a set of process definitions. Filter by a comma-separated list of `Strings`. (optional)
	processInstanceBusinessKey := "processInstanceBusinessKey_example" // string | Filter by process instance business key. (optional)
	processInstanceBusinessKeyIn := "processInstanceBusinessKeyIn_example" // string | Filter by a list of business keys. A process instance must have one of the given business keys. Filter by a comma-separated list of `Strings` (optional)
	processInstanceBusinessKeyLike := "processInstanceBusinessKeyLike_example" // string | Filter by process instance business key that the parameter is a substring of. (optional)
	rootProcessInstances := true // bool | Restrict the query to all process instances that are top level process instances. (optional)
	finished := true // bool | Only include finished process instances. This flag includes all process instances that are completed or terminated. Value may only be `true`, as `false` is the default behavior. (optional)
	unfinished := true // bool | Only include unfinished process instances. Value may only be `true`, as `false` is the default behavior. (optional)
	withIncidents := true // bool | Only include process instances which have an incident. Value may only be `true`, as `false` is the default behavior. (optional)
	withRootIncidents := true // bool | Only include process instances which have a root incident. Value may only be `true`, as `false` is the default behavior. (optional)
	incidentType := "incidentType_example" // string | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentStatus := "incidentStatus_example" // string | Only include process instances which have an incident in status either open or resolved. To get all process instances, use the query parameter withIncidents. (optional)
	incidentMessage := "incidentMessage_example" // string | Filter by the incident message. Exact match. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Filter by the incident message that the parameter is a substring of. (optional)
	startedBefore := time.Now() // time.Time | Restrict to instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedBefore := time.Now() // time.Time | Restrict to instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedAfter := time.Now() // time.Time | Restrict to instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	executedActivityAfter := time.Now() // time.Time | Restrict to instances that executed an activity after the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	executedActivityBefore := time.Now() // time.Time | Restrict to instances that executed an activity before the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	executedJobAfter := time.Now() // time.Time | Restrict to instances that executed an job after the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	executedJobBefore := time.Now() // time.Time | Restrict to instances that executed an job before the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedBy := "startedBy_example" // string | Only include process instances that were started by the given user. (optional)
	superProcessInstanceId := "superProcessInstanceId_example" // string | Restrict query to all process instances that are sub process instances of the given process instance. Takes a process instance id. (optional)
	subProcessInstanceId := "subProcessInstanceId_example" // string | Restrict query to one process instance that has a sub process instance with the given id. (optional)
	superCaseInstanceId := "superCaseInstanceId_example" // string | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. (optional)
	subCaseInstanceId := "subCaseInstanceId_example" // string | Restrict query to one process instance that has a sub case instance with the given id. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a list of tenant ids. A process instance must have one of the given tenant ids. Filter by a comma-separated list of `Strings` (optional)
	withoutTenantId := true // bool | Only include historic process instances which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	executedActivityIdIn := "executedActivityIdIn_example" // string | Restrict to instances that executed an activity with one of given ids. Filter by a comma-separated list of `Strings` (optional)
	activeActivityIdIn := "activeActivityIdIn_example" // string | Restrict to instances that have an active activity with one of given ids. Filter by a comma-separated list of `Strings` (optional)
	active := true // bool | Restrict to instances that are active. (optional)
	suspended := true // bool | Restrict to instances that are suspended. (optional)
	completed := true // bool | Restrict to instances that are completed. (optional)
	externallyTerminated := true // bool | Restrict to instances that are externallyTerminated. (optional)
	internallyTerminated := true // bool | Restrict to instances that are internallyTerminated. (optional)
	variables := "variables_example" // string | Only include process instances that have/had variables with certain values. Variable filtering expressions are comma-separated and are structured as follows: A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note:** Values are always treated as String objects on server side.  Valid operator values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`.  Key and value may not contain underscore or comma characters.  (optional)
	variableNamesIgnoreCase := true // bool | Match all variable names provided in variables case-insensitively. If set to `true` variableName and variablename are treated as equal. (optional)
	variableValuesIgnoreCase := true // bool | Match all variable values provided in variables case-insensitively. If set to `true` variableValue and variablevalue are treated as equal. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricProcessInstanceAPI.GetHistoricProcessInstances(context.Background()).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).ProcessInstanceId(processInstanceId).ProcessInstanceIds(processInstanceIds).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionName(processDefinitionName).ProcessDefinitionNameLike(processDefinitionNameLike).ProcessDefinitionKeyNotIn(processDefinitionKeyNotIn).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).RootProcessInstances(rootProcessInstances).Finished(finished).Unfinished(unfinished).WithIncidents(withIncidents).WithRootIncidents(withRootIncidents).IncidentType(incidentType).IncidentStatus(incidentStatus).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).ExecutedActivityAfter(executedActivityAfter).ExecutedActivityBefore(executedActivityBefore).ExecutedJobAfter(executedJobAfter).ExecutedJobBefore(executedJobBefore).StartedBy(startedBy).SuperProcessInstanceId(superProcessInstanceId).SubProcessInstanceId(subProcessInstanceId).SuperCaseInstanceId(superCaseInstanceId).SubCaseInstanceId(subCaseInstanceId).CaseInstanceId(caseInstanceId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).ExecutedActivityIdIn(executedActivityIdIn).ActiveActivityIdIn(activeActivityIdIn).Active(active).Suspended(suspended).Completed(completed).ExternallyTerminated(externallyTerminated).InternallyTerminated(internallyTerminated).Variables(variables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessInstanceAPI.GetHistoricProcessInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricProcessInstances`: []HistoricProcessInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricProcessInstanceAPI.GetHistoricProcessInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricProcessInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processInstanceIds** | **string** | Filter by process instance ids. Filter by a comma-separated list of &#x60;Strings&#x60;. | 
 **processDefinitionId** | **string** | Filter by the process definition the instances run on. | 
 **processDefinitionKey** | **string** | Filter by the key of the process definition the instances run on. | 
 **processDefinitionKeyIn** | **string** | Filter by a list of process definition keys. A process instance must have one of the given process definition keys. Filter by a comma-separated list of &#x60;Strings&#x60;. | 
 **processDefinitionName** | **string** | Filter by the name of the process definition the instances run on. | 
 **processDefinitionNameLike** | **string** | Filter by process definition names that the parameter is a substring of. | 
 **processDefinitionKeyNotIn** | **string** | Exclude instances that belong to a set of process definitions. Filter by a comma-separated list of &#x60;Strings&#x60;. | 
 **processInstanceBusinessKey** | **string** | Filter by process instance business key. | 
 **processInstanceBusinessKeyIn** | **string** | Filter by a list of business keys. A process instance must have one of the given business keys. Filter by a comma-separated list of &#x60;Strings&#x60; | 
 **processInstanceBusinessKeyLike** | **string** | Filter by process instance business key that the parameter is a substring of. | 
 **rootProcessInstances** | **bool** | Restrict the query to all process instances that are top level process instances. | 
 **finished** | **bool** | Only include finished process instances. This flag includes all process instances that are completed or terminated. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **unfinished** | **bool** | Only include unfinished process instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **withIncidents** | **bool** | Only include process instances which have an incident. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **withRootIncidents** | **bool** | Only include process instances which have a root incident. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **incidentType** | **string** | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentStatus** | **string** | Only include process instances which have an incident in status either open or resolved. To get all process instances, use the query parameter withIncidents. | 
 **incidentMessage** | **string** | Filter by the incident message. Exact match. | 
 **incidentMessageLike** | **string** | Filter by the incident message that the parameter is a substring of. | 
 **startedBefore** | **time.Time** | Restrict to instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedBefore** | **time.Time** | Restrict to instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedAfter** | **time.Time** | Restrict to instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **executedActivityAfter** | **time.Time** | Restrict to instances that executed an activity after the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **executedActivityBefore** | **time.Time** | Restrict to instances that executed an activity before the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **executedJobAfter** | **time.Time** | Restrict to instances that executed an job after the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **executedJobBefore** | **time.Time** | Restrict to instances that executed an job before the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedBy** | **string** | Only include process instances that were started by the given user. | 
 **superProcessInstanceId** | **string** | Restrict query to all process instances that are sub process instances of the given process instance. Takes a process instance id. | 
 **subProcessInstanceId** | **string** | Restrict query to one process instance that has a sub process instance with the given id. | 
 **superCaseInstanceId** | **string** | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. | 
 **subCaseInstanceId** | **string** | Restrict query to one process instance that has a sub case instance with the given id. | 
 **caseInstanceId** | **string** | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. | 
 **tenantIdIn** | **string** | Filter by a list of tenant ids. A process instance must have one of the given tenant ids. Filter by a comma-separated list of &#x60;Strings&#x60; | 
 **withoutTenantId** | **bool** | Only include historic process instances which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **executedActivityIdIn** | **string** | Restrict to instances that executed an activity with one of given ids. Filter by a comma-separated list of &#x60;Strings&#x60; | 
 **activeActivityIdIn** | **string** | Restrict to instances that have an active activity with one of given ids. Filter by a comma-separated list of &#x60;Strings&#x60; | 
 **active** | **bool** | Restrict to instances that are active. | 
 **suspended** | **bool** | Restrict to instances that are suspended. | 
 **completed** | **bool** | Restrict to instances that are completed. | 
 **externallyTerminated** | **bool** | Restrict to instances that are externallyTerminated. | 
 **internallyTerminated** | **bool** | Restrict to instances that are internallyTerminated. | 
 **variables** | **string** | Only include process instances that have/had variables with certain values. Variable filtering expressions are comma-separated and are structured as follows: A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note:** Values are always treated as String objects on server side.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;.  Key and value may not contain underscore or comma characters.  | 
 **variableNamesIgnoreCase** | **bool** | Match all variable names provided in variables case-insensitively. If set to &#x60;true&#x60; variableName and variablename are treated as equal. | 
 **variableValuesIgnoreCase** | **bool** | Match all variable values provided in variables case-insensitively. If set to &#x60;true&#x60; variableValue and variablevalue are treated as equal. | 

### Return type

[**[]HistoricProcessInstanceDto**](HistoricProcessInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricProcessInstancesCount

> CountResultDto GetHistoricProcessInstancesCount(ctx).ProcessInstanceId(processInstanceId).ProcessInstanceIds(processInstanceIds).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionName(processDefinitionName).ProcessDefinitionNameLike(processDefinitionNameLike).ProcessDefinitionKeyNotIn(processDefinitionKeyNotIn).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).RootProcessInstances(rootProcessInstances).Finished(finished).Unfinished(unfinished).WithIncidents(withIncidents).WithRootIncidents(withRootIncidents).IncidentType(incidentType).IncidentStatus(incidentStatus).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).ExecutedActivityAfter(executedActivityAfter).ExecutedActivityBefore(executedActivityBefore).ExecutedJobAfter(executedJobAfter).ExecutedJobBefore(executedJobBefore).StartedBy(startedBy).SuperProcessInstanceId(superProcessInstanceId).SubProcessInstanceId(subProcessInstanceId).SuperCaseInstanceId(superCaseInstanceId).SubCaseInstanceId(subCaseInstanceId).CaseInstanceId(caseInstanceId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).ExecutedActivityIdIn(executedActivityIdIn).ActiveActivityIdIn(activeActivityIdIn).Active(active).Suspended(suspended).Completed(completed).ExternallyTerminated(externallyTerminated).InternallyTerminated(internallyTerminated).Variables(variables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).Execute()

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
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processInstanceIds := "processInstanceIds_example" // string | Filter by process instance ids. Filter by a comma-separated list of `Strings`. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the process definition the instances run on. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by the key of the process definition the instances run on. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Filter by a list of process definition keys. A process instance must have one of the given process definition keys. Filter by a comma-separated list of `Strings`. (optional)
	processDefinitionName := "processDefinitionName_example" // string | Filter by the name of the process definition the instances run on. (optional)
	processDefinitionNameLike := "processDefinitionNameLike_example" // string | Filter by process definition names that the parameter is a substring of. (optional)
	processDefinitionKeyNotIn := "processDefinitionKeyNotIn_example" // string | Exclude instances that belong to a set of process definitions. Filter by a comma-separated list of `Strings`. (optional)
	processInstanceBusinessKey := "processInstanceBusinessKey_example" // string | Filter by process instance business key. (optional)
	processInstanceBusinessKeyIn := "processInstanceBusinessKeyIn_example" // string | Filter by a list of business keys. A process instance must have one of the given business keys. Filter by a comma-separated list of `Strings` (optional)
	processInstanceBusinessKeyLike := "processInstanceBusinessKeyLike_example" // string | Filter by process instance business key that the parameter is a substring of. (optional)
	rootProcessInstances := true // bool | Restrict the query to all process instances that are top level process instances. (optional)
	finished := true // bool | Only include finished process instances. This flag includes all process instances that are completed or terminated. Value may only be `true`, as `false` is the default behavior. (optional)
	unfinished := true // bool | Only include unfinished process instances. Value may only be `true`, as `false` is the default behavior. (optional)
	withIncidents := true // bool | Only include process instances which have an incident. Value may only be `true`, as `false` is the default behavior. (optional)
	withRootIncidents := true // bool | Only include process instances which have a root incident. Value may only be `true`, as `false` is the default behavior. (optional)
	incidentType := "incidentType_example" // string | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentStatus := "incidentStatus_example" // string | Only include process instances which have an incident in status either open or resolved. To get all process instances, use the query parameter withIncidents. (optional)
	incidentMessage := "incidentMessage_example" // string | Filter by the incident message. Exact match. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Filter by the incident message that the parameter is a substring of. (optional)
	startedBefore := time.Now() // time.Time | Restrict to instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedBefore := time.Now() // time.Time | Restrict to instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedAfter := time.Now() // time.Time | Restrict to instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	executedActivityAfter := time.Now() // time.Time | Restrict to instances that executed an activity after the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	executedActivityBefore := time.Now() // time.Time | Restrict to instances that executed an activity before the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	executedJobAfter := time.Now() // time.Time | Restrict to instances that executed an job after the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	executedJobBefore := time.Now() // time.Time | Restrict to instances that executed an job before the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedBy := "startedBy_example" // string | Only include process instances that were started by the given user. (optional)
	superProcessInstanceId := "superProcessInstanceId_example" // string | Restrict query to all process instances that are sub process instances of the given process instance. Takes a process instance id. (optional)
	subProcessInstanceId := "subProcessInstanceId_example" // string | Restrict query to one process instance that has a sub process instance with the given id. (optional)
	superCaseInstanceId := "superCaseInstanceId_example" // string | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. (optional)
	subCaseInstanceId := "subCaseInstanceId_example" // string | Restrict query to one process instance that has a sub case instance with the given id. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a list of tenant ids. A process instance must have one of the given tenant ids. Filter by a comma-separated list of `Strings` (optional)
	withoutTenantId := true // bool | Only include historic process instances which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	executedActivityIdIn := "executedActivityIdIn_example" // string | Restrict to instances that executed an activity with one of given ids. Filter by a comma-separated list of `Strings` (optional)
	activeActivityIdIn := "activeActivityIdIn_example" // string | Restrict to instances that have an active activity with one of given ids. Filter by a comma-separated list of `Strings` (optional)
	active := true // bool | Restrict to instances that are active. (optional)
	suspended := true // bool | Restrict to instances that are suspended. (optional)
	completed := true // bool | Restrict to instances that are completed. (optional)
	externallyTerminated := true // bool | Restrict to instances that are externallyTerminated. (optional)
	internallyTerminated := true // bool | Restrict to instances that are internallyTerminated. (optional)
	variables := "variables_example" // string | Only include process instances that have/had variables with certain values. Variable filtering expressions are comma-separated and are structured as follows: A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note:** Values are always treated as String objects on server side.  Valid operator values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`.  Key and value may not contain underscore or comma characters.  (optional)
	variableNamesIgnoreCase := true // bool | Match all variable names provided in variables case-insensitively. If set to `true` variableName and variablename are treated as equal. (optional)
	variableValuesIgnoreCase := true // bool | Match all variable values provided in variables case-insensitively. If set to `true` variableValue and variablevalue are treated as equal. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricProcessInstanceAPI.GetHistoricProcessInstancesCount(context.Background()).ProcessInstanceId(processInstanceId).ProcessInstanceIds(processInstanceIds).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionName(processDefinitionName).ProcessDefinitionNameLike(processDefinitionNameLike).ProcessDefinitionKeyNotIn(processDefinitionKeyNotIn).ProcessInstanceBusinessKey(processInstanceBusinessKey).ProcessInstanceBusinessKeyIn(processInstanceBusinessKeyIn).ProcessInstanceBusinessKeyLike(processInstanceBusinessKeyLike).RootProcessInstances(rootProcessInstances).Finished(finished).Unfinished(unfinished).WithIncidents(withIncidents).WithRootIncidents(withRootIncidents).IncidentType(incidentType).IncidentStatus(incidentStatus).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).ExecutedActivityAfter(executedActivityAfter).ExecutedActivityBefore(executedActivityBefore).ExecutedJobAfter(executedJobAfter).ExecutedJobBefore(executedJobBefore).StartedBy(startedBy).SuperProcessInstanceId(superProcessInstanceId).SubProcessInstanceId(subProcessInstanceId).SuperCaseInstanceId(superCaseInstanceId).SubCaseInstanceId(subCaseInstanceId).CaseInstanceId(caseInstanceId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).ExecutedActivityIdIn(executedActivityIdIn).ActiveActivityIdIn(activeActivityIdIn).Active(active).Suspended(suspended).Completed(completed).ExternallyTerminated(externallyTerminated).InternallyTerminated(internallyTerminated).Variables(variables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessInstanceAPI.GetHistoricProcessInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricProcessInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricProcessInstanceAPI.GetHistoricProcessInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricProcessInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processInstanceIds** | **string** | Filter by process instance ids. Filter by a comma-separated list of &#x60;Strings&#x60;. | 
 **processDefinitionId** | **string** | Filter by the process definition the instances run on. | 
 **processDefinitionKey** | **string** | Filter by the key of the process definition the instances run on. | 
 **processDefinitionKeyIn** | **string** | Filter by a list of process definition keys. A process instance must have one of the given process definition keys. Filter by a comma-separated list of &#x60;Strings&#x60;. | 
 **processDefinitionName** | **string** | Filter by the name of the process definition the instances run on. | 
 **processDefinitionNameLike** | **string** | Filter by process definition names that the parameter is a substring of. | 
 **processDefinitionKeyNotIn** | **string** | Exclude instances that belong to a set of process definitions. Filter by a comma-separated list of &#x60;Strings&#x60;. | 
 **processInstanceBusinessKey** | **string** | Filter by process instance business key. | 
 **processInstanceBusinessKeyIn** | **string** | Filter by a list of business keys. A process instance must have one of the given business keys. Filter by a comma-separated list of &#x60;Strings&#x60; | 
 **processInstanceBusinessKeyLike** | **string** | Filter by process instance business key that the parameter is a substring of. | 
 **rootProcessInstances** | **bool** | Restrict the query to all process instances that are top level process instances. | 
 **finished** | **bool** | Only include finished process instances. This flag includes all process instances that are completed or terminated. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **unfinished** | **bool** | Only include unfinished process instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **withIncidents** | **bool** | Only include process instances which have an incident. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **withRootIncidents** | **bool** | Only include process instances which have a root incident. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **incidentType** | **string** | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentStatus** | **string** | Only include process instances which have an incident in status either open or resolved. To get all process instances, use the query parameter withIncidents. | 
 **incidentMessage** | **string** | Filter by the incident message. Exact match. | 
 **incidentMessageLike** | **string** | Filter by the incident message that the parameter is a substring of. | 
 **startedBefore** | **time.Time** | Restrict to instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedBefore** | **time.Time** | Restrict to instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedAfter** | **time.Time** | Restrict to instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **executedActivityAfter** | **time.Time** | Restrict to instances that executed an activity after the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **executedActivityBefore** | **time.Time** | Restrict to instances that executed an activity before the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **executedJobAfter** | **time.Time** | Restrict to instances that executed an job after the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **executedJobBefore** | **time.Time** | Restrict to instances that executed an job before the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedBy** | **string** | Only include process instances that were started by the given user. | 
 **superProcessInstanceId** | **string** | Restrict query to all process instances that are sub process instances of the given process instance. Takes a process instance id. | 
 **subProcessInstanceId** | **string** | Restrict query to one process instance that has a sub process instance with the given id. | 
 **superCaseInstanceId** | **string** | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. | 
 **subCaseInstanceId** | **string** | Restrict query to one process instance that has a sub case instance with the given id. | 
 **caseInstanceId** | **string** | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. | 
 **tenantIdIn** | **string** | Filter by a list of tenant ids. A process instance must have one of the given tenant ids. Filter by a comma-separated list of &#x60;Strings&#x60; | 
 **withoutTenantId** | **bool** | Only include historic process instances which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **executedActivityIdIn** | **string** | Restrict to instances that executed an activity with one of given ids. Filter by a comma-separated list of &#x60;Strings&#x60; | 
 **activeActivityIdIn** | **string** | Restrict to instances that have an active activity with one of given ids. Filter by a comma-separated list of &#x60;Strings&#x60; | 
 **active** | **bool** | Restrict to instances that are active. | 
 **suspended** | **bool** | Restrict to instances that are suspended. | 
 **completed** | **bool** | Restrict to instances that are completed. | 
 **externallyTerminated** | **bool** | Restrict to instances that are externallyTerminated. | 
 **internallyTerminated** | **bool** | Restrict to instances that are internallyTerminated. | 
 **variables** | **string** | Only include process instances that have/had variables with certain values. Variable filtering expressions are comma-separated and are structured as follows: A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note:** Values are always treated as String objects on server side.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;.  Key and value may not contain underscore or comma characters.  | 
 **variableNamesIgnoreCase** | **bool** | Match all variable names provided in variables case-insensitively. If set to &#x60;true&#x60; variableName and variablename are treated as equal. | 
 **variableValuesIgnoreCase** | **bool** | Match all variable values provided in variables case-insensitively. If set to &#x60;true&#x60; variableValue and variablevalue are treated as equal. | 

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


## QueryHistoricProcessInstances

> []HistoricProcessInstanceDto QueryHistoricProcessInstances(ctx).FirstResult(firstResult).MaxResults(maxResults).HistoricProcessInstanceQueryDto(historicProcessInstanceQueryDto).Execute()

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
	historicProcessInstanceQueryDto := *openapiclient.NewHistoricProcessInstanceQueryDto() // HistoricProcessInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricProcessInstanceAPI.QueryHistoricProcessInstances(context.Background()).FirstResult(firstResult).MaxResults(maxResults).HistoricProcessInstanceQueryDto(historicProcessInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessInstanceAPI.QueryHistoricProcessInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricProcessInstances`: []HistoricProcessInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricProcessInstanceAPI.QueryHistoricProcessInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricProcessInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **historicProcessInstanceQueryDto** | [**HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) |  | 

### Return type

[**[]HistoricProcessInstanceDto**](HistoricProcessInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryHistoricProcessInstancesCount

> CountResultDto QueryHistoricProcessInstancesCount(ctx).HistoricProcessInstanceQueryDto(historicProcessInstanceQueryDto).Execute()

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
	historicProcessInstanceQueryDto := *openapiclient.NewHistoricProcessInstanceQueryDto() // HistoricProcessInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricProcessInstanceAPI.QueryHistoricProcessInstancesCount(context.Background()).HistoricProcessInstanceQueryDto(historicProcessInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessInstanceAPI.QueryHistoricProcessInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricProcessInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricProcessInstanceAPI.QueryHistoricProcessInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricProcessInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **historicProcessInstanceQueryDto** | [**HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) |  | 

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


## SetRemovalTimeAsync

> BatchDto SetRemovalTimeAsync(ctx).SetRemovalTimeToHistoricProcessInstancesDto(setRemovalTimeToHistoricProcessInstancesDto).Execute()

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
	setRemovalTimeToHistoricProcessInstancesDto := *openapiclient.NewSetRemovalTimeToHistoricProcessInstancesDto() // SetRemovalTimeToHistoricProcessInstancesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricProcessInstanceAPI.SetRemovalTimeAsync(context.Background()).SetRemovalTimeToHistoricProcessInstancesDto(setRemovalTimeToHistoricProcessInstancesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessInstanceAPI.SetRemovalTimeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRemovalTimeAsync`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricProcessInstanceAPI.SetRemovalTimeAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRemovalTimeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setRemovalTimeToHistoricProcessInstancesDto** | [**SetRemovalTimeToHistoricProcessInstancesDto**](SetRemovalTimeToHistoricProcessInstancesDto.md) |  | 

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

