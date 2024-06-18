# \IncidentAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearIncidentAnnotation**](IncidentAPI.md#ClearIncidentAnnotation) | **Delete** /incident/{id}/annotation | Clear Incident Annotation
[**GetIncident**](IncidentAPI.md#GetIncident) | **Get** /incident/{id} | Get Incident
[**GetIncidents**](IncidentAPI.md#GetIncidents) | **Get** /incident | Get List
[**GetIncidentsCount**](IncidentAPI.md#GetIncidentsCount) | **Get** /incident/count | Get List Count
[**ResolveIncident**](IncidentAPI.md#ResolveIncident) | **Delete** /incident/{id} | Resolve Incident
[**SetIncidentAnnotation**](IncidentAPI.md#SetIncidentAnnotation) | **Put** /incident/{id}/annotation | Set Incident Annotation



## ClearIncidentAnnotation

> ClearIncidentAnnotation(ctx, id).Execute()

Clear Incident Annotation



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
	id := "id_example" // string | The id of the incident to clear the annotation at.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IncidentAPI.ClearIncidentAnnotation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentAPI.ClearIncidentAnnotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the incident to clear the annotation at. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearIncidentAnnotationRequest struct via the builder pattern


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


## GetIncident

> IncidentDto GetIncident(ctx, id).Execute()

Get Incident



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
	id := "id_example" // string | The id of the incident to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentAPI.GetIncident(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentAPI.GetIncident``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncident`: IncidentDto
	fmt.Fprintf(os.Stdout, "Response from `IncidentAPI.GetIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the incident to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IncidentDto**](IncidentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidents

> []IncidentDto GetIncidents(ctx).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessInstanceId(processInstanceId).ExecutionId(executionId).IncidentTimestampBefore(incidentTimestampBefore).IncidentTimestampAfter(incidentTimestampAfter).ActivityId(activityId).FailedActivityId(failedActivityId).CauseIncidentId(causeIncidentId).RootCauseIncidentId(rootCauseIncidentId).Configuration(configuration).TenantIdIn(tenantIdIn).JobDefinitionIdIn(jobDefinitionIdIn).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

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
	incidentId := "incidentId_example" // string | Restricts to incidents that have the given id. (optional)
	incidentType := "incidentType_example" // string | Restricts to incidents that belong to the given incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentMessage := "incidentMessage_example" // string | Restricts to incidents that have the given incident message. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Restricts to incidents that incidents message is a substring of the given value. The string can include the wildcard character '%' to express like-strategy: starts with (`string%`), ends with (`%string`) or contains (`%string%`). (optional)
	processDefinitionId := "processDefinitionId_example" // string | Restricts to incidents that belong to a process definition with the given id. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Restricts to incidents that belong to a process definition with the given keys. Must be a comma-separated list. (optional)
	processInstanceId := "processInstanceId_example" // string | Restricts to incidents that belong to a process instance with the given id. (optional)
	executionId := "executionId_example" // string | Restricts to incidents that belong to an execution with the given id. (optional)
	incidentTimestampBefore := time.Now() // time.Time | Restricts to incidents that have an incidentTimestamp date before the given date.  By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	incidentTimestampAfter := time.Now() // time.Time | Restricts to incidents that have an incidentTimestamp date after the given date.  By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	activityId := "activityId_example" // string | Restricts to incidents that belong to an activity with the given id. (optional)
	failedActivityId := "failedActivityId_example" // string | Restricts to incidents that were created due to the failure of an activity with the given id. (optional)
	causeIncidentId := "causeIncidentId_example" // string | Restricts to incidents that have the given incident id as cause incident. (optional)
	rootCauseIncidentId := "rootCauseIncidentId_example" // string | Restricts to incidents that have the given incident id as root cause incident. (optional)
	configuration := "configuration_example" // string | Restricts to incidents that have the given parameter set as configuration. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Restricts to incidents that have one of the given comma-separated tenant ids. (optional)
	jobDefinitionIdIn := "jobDefinitionIdIn_example" // string | Restricts to incidents that have one of the given comma-separated job definition ids. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentAPI.GetIncidents(context.Background()).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessInstanceId(processInstanceId).ExecutionId(executionId).IncidentTimestampBefore(incidentTimestampBefore).IncidentTimestampAfter(incidentTimestampAfter).ActivityId(activityId).FailedActivityId(failedActivityId).CauseIncidentId(causeIncidentId).RootCauseIncidentId(rootCauseIncidentId).Configuration(configuration).TenantIdIn(tenantIdIn).JobDefinitionIdIn(jobDefinitionIdIn).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentAPI.GetIncidents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncidents`: []IncidentDto
	fmt.Fprintf(os.Stdout, "Response from `IncidentAPI.GetIncidents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incidentId** | **string** | Restricts to incidents that have the given id. | 
 **incidentType** | **string** | Restricts to incidents that belong to the given incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentMessage** | **string** | Restricts to incidents that have the given incident message. | 
 **incidentMessageLike** | **string** | Restricts to incidents that incidents message is a substring of the given value. The string can include the wildcard character &#39;%&#39; to express like-strategy: starts with (&#x60;string%&#x60;), ends with (&#x60;%string&#x60;) or contains (&#x60;%string%&#x60;). | 
 **processDefinitionId** | **string** | Restricts to incidents that belong to a process definition with the given id. | 
 **processDefinitionKeyIn** | **string** | Restricts to incidents that belong to a process definition with the given keys. Must be a comma-separated list. | 
 **processInstanceId** | **string** | Restricts to incidents that belong to a process instance with the given id. | 
 **executionId** | **string** | Restricts to incidents that belong to an execution with the given id. | 
 **incidentTimestampBefore** | **time.Time** | Restricts to incidents that have an incidentTimestamp date before the given date.  By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **incidentTimestampAfter** | **time.Time** | Restricts to incidents that have an incidentTimestamp date after the given date.  By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **activityId** | **string** | Restricts to incidents that belong to an activity with the given id. | 
 **failedActivityId** | **string** | Restricts to incidents that were created due to the failure of an activity with the given id. | 
 **causeIncidentId** | **string** | Restricts to incidents that have the given incident id as cause incident. | 
 **rootCauseIncidentId** | **string** | Restricts to incidents that have the given incident id as root cause incident. | 
 **configuration** | **string** | Restricts to incidents that have the given parameter set as configuration. | 
 **tenantIdIn** | **string** | Restricts to incidents that have one of the given comma-separated tenant ids. | 
 **jobDefinitionIdIn** | **string** | Restricts to incidents that have one of the given comma-separated job definition ids. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]IncidentDto**](IncidentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentsCount

> CountResultDto GetIncidentsCount(ctx).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessInstanceId(processInstanceId).ExecutionId(executionId).IncidentTimestampBefore(incidentTimestampBefore).IncidentTimestampAfter(incidentTimestampAfter).ActivityId(activityId).FailedActivityId(failedActivityId).CauseIncidentId(causeIncidentId).RootCauseIncidentId(rootCauseIncidentId).Configuration(configuration).TenantIdIn(tenantIdIn).JobDefinitionIdIn(jobDefinitionIdIn).Execute()

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
	incidentId := "incidentId_example" // string | Restricts to incidents that have the given id. (optional)
	incidentType := "incidentType_example" // string | Restricts to incidents that belong to the given incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentMessage := "incidentMessage_example" // string | Restricts to incidents that have the given incident message. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Restricts to incidents that incidents message is a substring of the given value. The string can include the wildcard character '%' to express like-strategy: starts with (`string%`), ends with (`%string`) or contains (`%string%`). (optional)
	processDefinitionId := "processDefinitionId_example" // string | Restricts to incidents that belong to a process definition with the given id. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Restricts to incidents that belong to a process definition with the given keys. Must be a comma-separated list. (optional)
	processInstanceId := "processInstanceId_example" // string | Restricts to incidents that belong to a process instance with the given id. (optional)
	executionId := "executionId_example" // string | Restricts to incidents that belong to an execution with the given id. (optional)
	incidentTimestampBefore := time.Now() // time.Time | Restricts to incidents that have an incidentTimestamp date before the given date.  By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	incidentTimestampAfter := time.Now() // time.Time | Restricts to incidents that have an incidentTimestamp date after the given date.  By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	activityId := "activityId_example" // string | Restricts to incidents that belong to an activity with the given id. (optional)
	failedActivityId := "failedActivityId_example" // string | Restricts to incidents that were created due to the failure of an activity with the given id. (optional)
	causeIncidentId := "causeIncidentId_example" // string | Restricts to incidents that have the given incident id as cause incident. (optional)
	rootCauseIncidentId := "rootCauseIncidentId_example" // string | Restricts to incidents that have the given incident id as root cause incident. (optional)
	configuration := "configuration_example" // string | Restricts to incidents that have the given parameter set as configuration. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Restricts to incidents that have one of the given comma-separated tenant ids. (optional)
	jobDefinitionIdIn := "jobDefinitionIdIn_example" // string | Restricts to incidents that have one of the given comma-separated job definition ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentAPI.GetIncidentsCount(context.Background()).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessInstanceId(processInstanceId).ExecutionId(executionId).IncidentTimestampBefore(incidentTimestampBefore).IncidentTimestampAfter(incidentTimestampAfter).ActivityId(activityId).FailedActivityId(failedActivityId).CauseIncidentId(causeIncidentId).RootCauseIncidentId(rootCauseIncidentId).Configuration(configuration).TenantIdIn(tenantIdIn).JobDefinitionIdIn(jobDefinitionIdIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentAPI.GetIncidentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncidentsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `IncidentAPI.GetIncidentsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incidentId** | **string** | Restricts to incidents that have the given id. | 
 **incidentType** | **string** | Restricts to incidents that belong to the given incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentMessage** | **string** | Restricts to incidents that have the given incident message. | 
 **incidentMessageLike** | **string** | Restricts to incidents that incidents message is a substring of the given value. The string can include the wildcard character &#39;%&#39; to express like-strategy: starts with (&#x60;string%&#x60;), ends with (&#x60;%string&#x60;) or contains (&#x60;%string%&#x60;). | 
 **processDefinitionId** | **string** | Restricts to incidents that belong to a process definition with the given id. | 
 **processDefinitionKeyIn** | **string** | Restricts to incidents that belong to a process definition with the given keys. Must be a comma-separated list. | 
 **processInstanceId** | **string** | Restricts to incidents that belong to a process instance with the given id. | 
 **executionId** | **string** | Restricts to incidents that belong to an execution with the given id. | 
 **incidentTimestampBefore** | **time.Time** | Restricts to incidents that have an incidentTimestamp date before the given date.  By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **incidentTimestampAfter** | **time.Time** | Restricts to incidents that have an incidentTimestamp date after the given date.  By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **activityId** | **string** | Restricts to incidents that belong to an activity with the given id. | 
 **failedActivityId** | **string** | Restricts to incidents that were created due to the failure of an activity with the given id. | 
 **causeIncidentId** | **string** | Restricts to incidents that have the given incident id as cause incident. | 
 **rootCauseIncidentId** | **string** | Restricts to incidents that have the given incident id as root cause incident. | 
 **configuration** | **string** | Restricts to incidents that have the given parameter set as configuration. | 
 **tenantIdIn** | **string** | Restricts to incidents that have one of the given comma-separated tenant ids. | 
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


## ResolveIncident

> ResolveIncident(ctx, id).Execute()

Resolve Incident



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
	id := "id_example" // string | The id of the incident to be resolved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IncidentAPI.ResolveIncident(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentAPI.ResolveIncident``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the incident to be resolved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveIncidentRequest struct via the builder pattern


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


## SetIncidentAnnotation

> SetIncidentAnnotation(ctx, id).AnnotationDto(annotationDto).Execute()

Set Incident Annotation



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
	id := "id_example" // string | The id of the incident to clear the annotation at.
	annotationDto := *openapiclient.NewAnnotationDto() // AnnotationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IncidentAPI.SetIncidentAnnotation(context.Background(), id).AnnotationDto(annotationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentAPI.SetIncidentAnnotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the incident to clear the annotation at. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetIncidentAnnotationRequest struct via the builder pattern


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

