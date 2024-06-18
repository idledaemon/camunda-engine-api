# \MetricsAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaskMetrics**](MetricsAPI.md#DeleteTaskMetrics) | **Delete** /metrics/task-worker | Delete Task Worker Metrics
[**GetMetrics**](MetricsAPI.md#GetMetrics) | **Get** /metrics/{metrics-name}/sum | Get Sum
[**Interval**](MetricsAPI.md#Interval) | **Get** /metrics | Get Metrics in Interval



## DeleteTaskMetrics

> DeleteTaskMetrics(ctx).Date(date).Execute()

Delete Task Worker Metrics



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
	date := time.Now() // time.Time | The date prior to which all task worker metrics should be deleted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MetricsAPI.DeleteTaskMetrics(context.Background()).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.DeleteTaskMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** | The date prior to which all task worker metrics should be deleted. | 

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


## GetMetrics

> MetricsResultDto GetMetrics(ctx, metricsName).StartDate(startDate).EndDate(endDate).Execute()

Get Sum



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
	metricsName := "metricsName_example" // string | The name of the metric.
	startDate := time.Now() // time.Time | The start date (inclusive). (optional)
	endDate := time.Now() // time.Time | The end date (exclusive). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetrics(context.Background(), metricsName).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetrics`: MetricsResultDto
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricsName** | **string** | The name of the metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | The start date (inclusive). | 
 **endDate** | **time.Time** | The end date (exclusive). | 

### Return type

[**MetricsResultDto**](MetricsResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Interval

> []MetricsIntervalResultDto Interval(ctx).Name(name).Reporter(reporter).StartDate(startDate).EndDate(endDate).FirstResult(firstResult).MaxResults(maxResults).Interval(interval).AggregateByReporter(aggregateByReporter).Execute()

Get Metrics in Interval



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
	name := "name_example" // string | The name of the metric. (optional)
	reporter := "reporter_example" // string | The name of the reporter (host), on which the metrics was logged. This will have value provided by the [hostname configuration property](https://docs.camunda.org/manual/7.21/reference/deployment-descriptors/tags/process-engine/#hostname). (optional)
	startDate := time.Now() // time.Time | The start date (inclusive). (optional)
	endDate := time.Now() // time.Time | The end date (exclusive). (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)
	interval := "interval_example" // string | The interval for which the metrics should be aggregated. Time unit is seconds. Default: The interval is set to 15 minutes (900 seconds). (optional) (default to "900")
	aggregateByReporter := "aggregateByReporter_example" // string | Aggregate metrics by reporter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.Interval(context.Background()).Name(name).Reporter(reporter).StartDate(startDate).EndDate(endDate).FirstResult(firstResult).MaxResults(maxResults).Interval(interval).AggregateByReporter(aggregateByReporter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.Interval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Interval`: []MetricsIntervalResultDto
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.Interval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the metric. | 
 **reporter** | **string** | The name of the reporter (host), on which the metrics was logged. This will have value provided by the [hostname configuration property](https://docs.camunda.org/manual/7.21/reference/deployment-descriptors/tags/process-engine/#hostname). | 
 **startDate** | **time.Time** | The start date (inclusive). | 
 **endDate** | **time.Time** | The end date (exclusive). | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **interval** | **string** | The interval for which the metrics should be aggregated. Time unit is seconds. Default: The interval is set to 15 minutes (900 seconds). | [default to &quot;900&quot;]
 **aggregateByReporter** | **string** | Aggregate metrics by reporter. | 

### Return type

[**[]MetricsIntervalResultDto**](MetricsIntervalResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

