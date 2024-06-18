# \HistoricIncidentAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricIncidents**](HistoricIncidentAPI.md#GetHistoricIncidents) | **Get** /history/incident | Get Incidents
[**GetHistoricIncidentsCount**](HistoricIncidentAPI.md#GetHistoricIncidentsCount) | **Get** /history/incident/count | Get Incident Count



## GetHistoricIncidents

> []HistoricIncidentDto GetHistoricIncidents(ctx).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessInstanceId(processInstanceId).ExecutionId(executionId).CreateTimeBefore(createTimeBefore).CreateTimeAfter(createTimeAfter).EndTimeBefore(endTimeBefore).EndTimeAfter(endTimeAfter).ActivityId(activityId).FailedActivityId(failedActivityId).CauseIncidentId(causeIncidentId).RootCauseIncidentId(rootCauseIncidentId).Configuration(configuration).HistoryConfiguration(historyConfiguration).Open(open).Resolved(resolved).Deleted(deleted).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).JobDefinitionIdIn(jobDefinitionIdIn).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Incidents



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
	incidentId := "incidentId_example" // string | Restricts to incidents that have the given id. (optional)
	incidentType := "incidentType_example" // string | Restricts to incidents that belong to the given incident type. See the [User Guide](/manual/develop/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentMessage := "incidentMessage_example" // string | Restricts to incidents that have the given incident message. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Restricts to incidents that incidents message is a substring of the given value. The string can include the wildcard character '%' to express like-strategy: starts with (string%), ends with (%string) or contains (%string%).  (optional)
	processDefinitionId := "processDefinitionId_example" // string | Restricts to incidents that belong to a process definition with the given id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Restricts to incidents that have the given processDefinitionKey. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Restricts to incidents that have one of the given process definition keys. (optional)
	processInstanceId := "processInstanceId_example" // string | Restricts to incidents that belong to a process instance with the given id. (optional)
	executionId := "executionId_example" // string | Restricts to incidents that belong to an execution with the given id. (optional)
	createTimeBefore := time.Now() // time.Time | Restricts to incidents that have a createTime date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	createTimeAfter := time.Now() // time.Time | Restricts to incidents that have a createTime date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	endTimeBefore := time.Now() // time.Time | Restricts to incidents that have an endTimeBefore date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	endTimeAfter := time.Now() // time.Time | Restricts to incidents that have an endTimeAfter date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	activityId := "activityId_example" // string | Restricts to incidents that belong to an activity with the given id. (optional)
	failedActivityId := "failedActivityId_example" // string | Restricts to incidents that were created due to the failure of an activity with the given id. (optional)
	causeIncidentId := "causeIncidentId_example" // string | Restricts to incidents that have the given incident id as cause incident. (optional)
	rootCauseIncidentId := "rootCauseIncidentId_example" // string | Restricts to incidents that have the given incident id as root cause incident. (optional)
	configuration := "configuration_example" // string | Restricts to incidents that have the given parameter set as configuration. (optional)
	historyConfiguration := "historyConfiguration_example" // string | Restricts to incidents that have the given parameter set as history configuration. (optional)
	open := true // bool | Restricts to incidents that are open. (optional)
	resolved := true // bool | Restricts to incidents that are resolved. (optional)
	deleted := true // bool | Restricts to incidents that are deleted. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Restricts to incidents that have one of the given comma-separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic incidents that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	jobDefinitionIdIn := "jobDefinitionIdIn_example" // string | Restricts to incidents that have one of the given comma-separated job definition ids. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricIncidentAPI.GetHistoricIncidents(context.Background()).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessInstanceId(processInstanceId).ExecutionId(executionId).CreateTimeBefore(createTimeBefore).CreateTimeAfter(createTimeAfter).EndTimeBefore(endTimeBefore).EndTimeAfter(endTimeAfter).ActivityId(activityId).FailedActivityId(failedActivityId).CauseIncidentId(causeIncidentId).RootCauseIncidentId(rootCauseIncidentId).Configuration(configuration).HistoryConfiguration(historyConfiguration).Open(open).Resolved(resolved).Deleted(deleted).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).JobDefinitionIdIn(jobDefinitionIdIn).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricIncidentAPI.GetHistoricIncidents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricIncidents`: []HistoricIncidentDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricIncidentAPI.GetHistoricIncidents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incidentId** | **string** | Restricts to incidents that have the given id. | 
 **incidentType** | **string** | Restricts to incidents that belong to the given incident type. See the [User Guide](/manual/develop/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentMessage** | **string** | Restricts to incidents that have the given incident message. | 
 **incidentMessageLike** | **string** | Restricts to incidents that incidents message is a substring of the given value. The string can include the wildcard character &#39;%&#39; to express like-strategy: starts with (string%), ends with (%string) or contains (%string%).  | 
 **processDefinitionId** | **string** | Restricts to incidents that belong to a process definition with the given id. | 
 **processDefinitionKey** | **string** | Restricts to incidents that have the given processDefinitionKey. | 
 **processDefinitionKeyIn** | **string** | Restricts to incidents that have one of the given process definition keys. | 
 **processInstanceId** | **string** | Restricts to incidents that belong to a process instance with the given id. | 
 **executionId** | **string** | Restricts to incidents that belong to an execution with the given id. | 
 **createTimeBefore** | **time.Time** | Restricts to incidents that have a createTime date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **createTimeAfter** | **time.Time** | Restricts to incidents that have a createTime date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **endTimeBefore** | **time.Time** | Restricts to incidents that have an endTimeBefore date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **endTimeAfter** | **time.Time** | Restricts to incidents that have an endTimeAfter date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **activityId** | **string** | Restricts to incidents that belong to an activity with the given id. | 
 **failedActivityId** | **string** | Restricts to incidents that were created due to the failure of an activity with the given id. | 
 **causeIncidentId** | **string** | Restricts to incidents that have the given incident id as cause incident. | 
 **rootCauseIncidentId** | **string** | Restricts to incidents that have the given incident id as root cause incident. | 
 **configuration** | **string** | Restricts to incidents that have the given parameter set as configuration. | 
 **historyConfiguration** | **string** | Restricts to incidents that have the given parameter set as history configuration. | 
 **open** | **bool** | Restricts to incidents that are open. | 
 **resolved** | **bool** | Restricts to incidents that are resolved. | 
 **deleted** | **bool** | Restricts to incidents that are deleted. | 
 **tenantIdIn** | **string** | Restricts to incidents that have one of the given comma-separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic incidents that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **jobDefinitionIdIn** | **string** | Restricts to incidents that have one of the given comma-separated job definition ids. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]HistoricIncidentDto**](HistoricIncidentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricIncidentsCount

> CountResultDto GetHistoricIncidentsCount(ctx).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessInstanceId(processInstanceId).ExecutionId(executionId).CreateTimeBefore(createTimeBefore).CreateTimeAfter(createTimeAfter).EndTimeBefore(endTimeBefore).EndTimeAfter(endTimeAfter).ActivityId(activityId).FailedActivityId(failedActivityId).CauseIncidentId(causeIncidentId).RootCauseIncidentId(rootCauseIncidentId).Configuration(configuration).HistoryConfiguration(historyConfiguration).Open(open).Resolved(resolved).Deleted(deleted).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).JobDefinitionIdIn(jobDefinitionIdIn).Execute()

Get Incident Count



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
	incidentId := "incidentId_example" // string | Restricts to incidents that have the given id. (optional)
	incidentType := "incidentType_example" // string | Restricts to incidents that belong to the given incident type. See the [User Guide](/manual/develop/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentMessage := "incidentMessage_example" // string | Restricts to incidents that have the given incident message. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Restricts to incidents that incidents message is a substring of the given value. The string can include the wildcard character '%' to express like-strategy: starts with (string%), ends with (%string) or contains (%string%).  (optional)
	processDefinitionId := "processDefinitionId_example" // string | Restricts to incidents that belong to a process definition with the given id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Restricts to incidents that have the given processDefinitionKey. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Restricts to incidents that have one of the given process definition keys. (optional)
	processInstanceId := "processInstanceId_example" // string | Restricts to incidents that belong to a process instance with the given id. (optional)
	executionId := "executionId_example" // string | Restricts to incidents that belong to an execution with the given id. (optional)
	createTimeBefore := time.Now() // time.Time | Restricts to incidents that have a createTime date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	createTimeAfter := time.Now() // time.Time | Restricts to incidents that have a createTime date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	endTimeBefore := time.Now() // time.Time | Restricts to incidents that have an endTimeBefore date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	endTimeAfter := time.Now() // time.Time | Restricts to incidents that have an endTimeAfter date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	activityId := "activityId_example" // string | Restricts to incidents that belong to an activity with the given id. (optional)
	failedActivityId := "failedActivityId_example" // string | Restricts to incidents that were created due to the failure of an activity with the given id. (optional)
	causeIncidentId := "causeIncidentId_example" // string | Restricts to incidents that have the given incident id as cause incident. (optional)
	rootCauseIncidentId := "rootCauseIncidentId_example" // string | Restricts to incidents that have the given incident id as root cause incident. (optional)
	configuration := "configuration_example" // string | Restricts to incidents that have the given parameter set as configuration. (optional)
	historyConfiguration := "historyConfiguration_example" // string | Restricts to incidents that have the given parameter set as history configuration. (optional)
	open := true // bool | Restricts to incidents that are open. (optional)
	resolved := true // bool | Restricts to incidents that are resolved. (optional)
	deleted := true // bool | Restricts to incidents that are deleted. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Restricts to incidents that have one of the given comma-separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic incidents that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	jobDefinitionIdIn := "jobDefinitionIdIn_example" // string | Restricts to incidents that have one of the given comma-separated job definition ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricIncidentAPI.GetHistoricIncidentsCount(context.Background()).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessInstanceId(processInstanceId).ExecutionId(executionId).CreateTimeBefore(createTimeBefore).CreateTimeAfter(createTimeAfter).EndTimeBefore(endTimeBefore).EndTimeAfter(endTimeAfter).ActivityId(activityId).FailedActivityId(failedActivityId).CauseIncidentId(causeIncidentId).RootCauseIncidentId(rootCauseIncidentId).Configuration(configuration).HistoryConfiguration(historyConfiguration).Open(open).Resolved(resolved).Deleted(deleted).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).JobDefinitionIdIn(jobDefinitionIdIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricIncidentAPI.GetHistoricIncidentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricIncidentsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricIncidentAPI.GetHistoricIncidentsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricIncidentsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incidentId** | **string** | Restricts to incidents that have the given id. | 
 **incidentType** | **string** | Restricts to incidents that belong to the given incident type. See the [User Guide](/manual/develop/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentMessage** | **string** | Restricts to incidents that have the given incident message. | 
 **incidentMessageLike** | **string** | Restricts to incidents that incidents message is a substring of the given value. The string can include the wildcard character &#39;%&#39; to express like-strategy: starts with (string%), ends with (%string) or contains (%string%).  | 
 **processDefinitionId** | **string** | Restricts to incidents that belong to a process definition with the given id. | 
 **processDefinitionKey** | **string** | Restricts to incidents that have the given processDefinitionKey. | 
 **processDefinitionKeyIn** | **string** | Restricts to incidents that have one of the given process definition keys. | 
 **processInstanceId** | **string** | Restricts to incidents that belong to a process instance with the given id. | 
 **executionId** | **string** | Restricts to incidents that belong to an execution with the given id. | 
 **createTimeBefore** | **time.Time** | Restricts to incidents that have a createTime date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **createTimeAfter** | **time.Time** | Restricts to incidents that have a createTime date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **endTimeBefore** | **time.Time** | Restricts to incidents that have an endTimeBefore date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **endTimeAfter** | **time.Time** | Restricts to incidents that have an endTimeAfter date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **activityId** | **string** | Restricts to incidents that belong to an activity with the given id. | 
 **failedActivityId** | **string** | Restricts to incidents that were created due to the failure of an activity with the given id. | 
 **causeIncidentId** | **string** | Restricts to incidents that have the given incident id as cause incident. | 
 **rootCauseIncidentId** | **string** | Restricts to incidents that have the given incident id as root cause incident. | 
 **configuration** | **string** | Restricts to incidents that have the given parameter set as configuration. | 
 **historyConfiguration** | **string** | Restricts to incidents that have the given parameter set as history configuration. | 
 **open** | **bool** | Restricts to incidents that are open. | 
 **resolved** | **bool** | Restricts to incidents that are resolved. | 
 **deleted** | **bool** | Restricts to incidents that are deleted. | 
 **tenantIdIn** | **string** | Restricts to incidents that have one of the given comma-separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic incidents that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **jobDefinitionIdIn** | **string** | Restricts to incidents that have one of the given comma-separated job definition ids. | 

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

