# \HistoricProcessDefinitionAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCleanableHistoricProcessInstanceReport**](HistoricProcessDefinitionAPI.md#GetCleanableHistoricProcessInstanceReport) | **Get** /history/process-definition/cleanable-process-instance-report | Get Cleanable Process Instance Report
[**GetCleanableHistoricProcessInstanceReportCount**](HistoricProcessDefinitionAPI.md#GetCleanableHistoricProcessInstanceReportCount) | **Get** /history/process-definition/cleanable-process-instance-report/count | Get Cleanable Process Instance Report Count
[**GetHistoricActivityStatistics**](HistoricProcessDefinitionAPI.md#GetHistoricActivityStatistics) | **Get** /history/process-definition/{id}/statistics | Get Historic Activity Statistics



## GetCleanableHistoricProcessInstanceReport

> []CleanableHistoricProcessInstanceReportResultDto GetCleanableHistoricProcessInstanceReport(ctx).ProcessDefinitionIdIn(processDefinitionIdIn).ProcessDefinitionKeyIn(processDefinitionKeyIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Compact(compact).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Cleanable Process Instance Report



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
	processDefinitionIdIn := "processDefinitionIdIn_example" // string | Filter by process definition ids. Must be a comma-separated list of process definition ids. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Filter by process definition keys. Must be a comma-separated list of process definition keys. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A process definition must have one of the given  tenant ids. (optional)
	withoutTenantId := true // bool | Only include process definitions which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	compact := true // bool | Only include process instances which have more than zero finished instances. Value may only be `true`, as `false` is the default behavior. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricProcessDefinitionAPI.GetCleanableHistoricProcessInstanceReport(context.Background()).ProcessDefinitionIdIn(processDefinitionIdIn).ProcessDefinitionKeyIn(processDefinitionKeyIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Compact(compact).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessDefinitionAPI.GetCleanableHistoricProcessInstanceReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCleanableHistoricProcessInstanceReport`: []CleanableHistoricProcessInstanceReportResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricProcessDefinitionAPI.GetCleanableHistoricProcessInstanceReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCleanableHistoricProcessInstanceReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processDefinitionIdIn** | **string** | Filter by process definition ids. Must be a comma-separated list of process definition ids. | 
 **processDefinitionKeyIn** | **string** | Filter by process definition keys. Must be a comma-separated list of process definition keys. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A process definition must have one of the given  tenant ids. | 
 **withoutTenantId** | **bool** | Only include process definitions which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **compact** | **bool** | Only include process instances which have more than zero finished instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]CleanableHistoricProcessInstanceReportResultDto**](CleanableHistoricProcessInstanceReportResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCleanableHistoricProcessInstanceReportCount

> CountResultDto GetCleanableHistoricProcessInstanceReportCount(ctx).ProcessDefinitionIdIn(processDefinitionIdIn).ProcessDefinitionKeyIn(processDefinitionKeyIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Compact(compact).Execute()

Get Cleanable Process Instance Report Count



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
	processDefinitionIdIn := "processDefinitionIdIn_example" // string | Filter by process definition ids. Must be a comma-separated list of process definition ids. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Filter by process definition keys. Must be a comma-separated list of process definition keys. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A process definition must have one of the given  tenant ids. (optional)
	withoutTenantId := true // bool | Only include process definitions which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	compact := true // bool | Only include process instances which have more than zero finished instances. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricProcessDefinitionAPI.GetCleanableHistoricProcessInstanceReportCount(context.Background()).ProcessDefinitionIdIn(processDefinitionIdIn).ProcessDefinitionKeyIn(processDefinitionKeyIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Compact(compact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessDefinitionAPI.GetCleanableHistoricProcessInstanceReportCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCleanableHistoricProcessInstanceReportCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricProcessDefinitionAPI.GetCleanableHistoricProcessInstanceReportCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCleanableHistoricProcessInstanceReportCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processDefinitionIdIn** | **string** | Filter by process definition ids. Must be a comma-separated list of process definition ids. | 
 **processDefinitionKeyIn** | **string** | Filter by process definition keys. Must be a comma-separated list of process definition keys. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A process definition must have one of the given  tenant ids. | 
 **withoutTenantId** | **bool** | Only include process definitions which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **compact** | **bool** | Only include process instances which have more than zero finished instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

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


## GetHistoricActivityStatistics

> []HistoricActivityStatisticsDto GetHistoricActivityStatistics(ctx, id).Canceled(canceled).Finished(finished).CompleteScope(completeScope).Incidents(incidents).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).ProcessInstanceIdIn(processInstanceIdIn).SortBy(sortBy).SortOrder(sortOrder).Execute()

Get Historic Activity Statistics



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
	id := "id_example" // string | The id of the process definition.
	canceled := true // bool | Whether to include the number of canceled activity instances in the result or not. Valid values are `true` or `false`. Default: `false`. (optional)
	finished := true // bool | Whether to include the number of finished activity instances in the result or not. Valid values are `true` or `false`. Default: `false`. (optional)
	completeScope := true // bool | Whether to include the number of activity instances which completed a scope in the result or not. Valid values are `true` or `false`. Default: `false`. (optional)
	incidents := true // bool | Whether to include the number of incidents. Valid values are `true` or `false`. Default: `false`. (optional)
	startedBefore := time.Now() // time.Time | Restrict to process instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/),  the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`,  e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Restrict to process instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/),  the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`,  e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedBefore := time.Now() // time.Time | Restrict to process instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/),  the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`,  e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	finishedAfter := time.Now() // time.Time | Restrict to process instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/),  the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`,  e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	processInstanceIdIn := "processInstanceIdIn_example" // string | Restrict to process instances with the given IDs. The IDs must be provided as a comma- separated list. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricProcessDefinitionAPI.GetHistoricActivityStatistics(context.Background(), id).Canceled(canceled).Finished(finished).CompleteScope(completeScope).Incidents(incidents).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).ProcessInstanceIdIn(processInstanceIdIn).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricProcessDefinitionAPI.GetHistoricActivityStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricActivityStatistics`: []HistoricActivityStatisticsDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricProcessDefinitionAPI.GetHistoricActivityStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricActivityStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **canceled** | **bool** | Whether to include the number of canceled activity instances in the result or not. Valid values are &#x60;true&#x60; or &#x60;false&#x60;. Default: &#x60;false&#x60;. | 
 **finished** | **bool** | Whether to include the number of finished activity instances in the result or not. Valid values are &#x60;true&#x60; or &#x60;false&#x60;. Default: &#x60;false&#x60;. | 
 **completeScope** | **bool** | Whether to include the number of activity instances which completed a scope in the result or not. Valid values are &#x60;true&#x60; or &#x60;false&#x60;. Default: &#x60;false&#x60;. | 
 **incidents** | **bool** | Whether to include the number of incidents. Valid values are &#x60;true&#x60; or &#x60;false&#x60;. Default: &#x60;false&#x60;. | 
 **startedBefore** | **time.Time** | Restrict to process instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/),  the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;,  e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Restrict to process instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/),  the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;,  e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedBefore** | **time.Time** | Restrict to process instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/),  the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;,  e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **finishedAfter** | **time.Time** | Restrict to process instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/),  the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;,  e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **processInstanceIdIn** | **string** | Restrict to process instances with the given IDs. The IDs must be provided as a comma- separated list. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 

### Return type

[**[]HistoricActivityStatisticsDto**](HistoricActivityStatisticsDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

