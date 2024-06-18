# \BatchAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBatch**](BatchAPI.md#DeleteBatch) | **Delete** /batch/{id} | Delete
[**GetBatch**](BatchAPI.md#GetBatch) | **Get** /batch/{id} | Get
[**GetBatchStatistics**](BatchAPI.md#GetBatchStatistics) | **Get** /batch/statistics | Get Statistics
[**GetBatchStatisticsCount**](BatchAPI.md#GetBatchStatisticsCount) | **Get** /batch/statistics/count | Get Statistics Count
[**GetBatches**](BatchAPI.md#GetBatches) | **Get** /batch | Get List
[**GetBatchesCount**](BatchAPI.md#GetBatchesCount) | **Get** /batch/count | Get List Count
[**UpdateBatchSuspensionState**](BatchAPI.md#UpdateBatchSuspensionState) | **Put** /batch/{id}/suspended | Activate/Suspend



## DeleteBatch

> DeleteBatch(ctx, id).Cascade(cascade).Execute()

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
	id := "id_example" // string | The id of the batch to be deleted.
	cascade := true // bool | `true`, if the historic batch and historic job logs for this batch should also be deleted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchAPI.DeleteBatch(context.Background(), id).Cascade(cascade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.DeleteBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the batch to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **bool** | &#x60;true&#x60;, if the historic batch and historic job logs for this batch should also be deleted. | 

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


## GetBatch

> BatchDto GetBatch(ctx, id).Execute()

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
	id := "id_example" // string | The id of the batch to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.GetBatch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.GetBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatch`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.GetBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the batch to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchStatistics

> []BatchStatisticsDto GetBatchStatistics(ctx).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).BatchId(batchId).Type_(type_).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Suspended(suspended).CreatedBy(createdBy).StartedBefore(startedBefore).StartedAfter(startedAfter).WithFailures(withFailures).WithoutFailures(withoutFailures).Execute()

Get Statistics



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
	batchId := "batchId_example" // string | Filter by batch id. (optional)
	type_ := "type__example" // string | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of `Strings`. A batch matches if it has one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include batches which belong to no tenant. Value can effectively only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | A `Boolean` value which indicates whether only active or suspended batches should be included. When the value is set to `true`, only suspended batches will be returned and when the value is set to `false`, only active batches will be returned. (optional)
	createdBy := "createdBy_example" // string | Only include batches that were started by this user id. (optional)
	startedBefore := time.Now() // time.Time | Only include batches that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Only include batches that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	withFailures := true // bool | Only include batches having jobs with failures. Value can only be `true`. (optional)
	withoutFailures := true // bool | Only include batches having jobs without failures. Value can only be `true`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.GetBatchStatistics(context.Background()).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).BatchId(batchId).Type_(type_).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Suspended(suspended).CreatedBy(createdBy).StartedBefore(startedBefore).StartedAfter(startedAfter).WithFailures(withFailures).WithoutFailures(withoutFailures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.GetBatchStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchStatistics`: []BatchStatisticsDto
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.GetBatchStatistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **batchId** | **string** | Filter by batch id. | 
 **type_** | **string** | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of &#x60;Strings&#x60;. A batch matches if it has one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include batches which belong to no tenant. Value can effectively only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | A &#x60;Boolean&#x60; value which indicates whether only active or suspended batches should be included. When the value is set to &#x60;true&#x60;, only suspended batches will be returned and when the value is set to &#x60;false&#x60;, only active batches will be returned. | 
 **createdBy** | **string** | Only include batches that were started by this user id. | 
 **startedBefore** | **time.Time** | Only include batches that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Only include batches that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **withFailures** | **bool** | Only include batches having jobs with failures. Value can only be &#x60;true&#x60;. | 
 **withoutFailures** | **bool** | Only include batches having jobs without failures. Value can only be &#x60;true&#x60;. | 

### Return type

[**[]BatchStatisticsDto**](BatchStatisticsDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchStatisticsCount

> CountResultDto GetBatchStatisticsCount(ctx).BatchId(batchId).Type_(type_).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Suspended(suspended).CreatedBy(createdBy).StartedBefore(startedBefore).StartedAfter(startedAfter).WithFailures(withFailures).WithoutFailures(withoutFailures).Execute()

Get Statistics Count



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
	batchId := "batchId_example" // string | Filter by batch id. (optional)
	type_ := "type__example" // string | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of `Strings`. A batch matches if it has one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include batches which belong to no tenant. Value can effectively only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | A `Boolean` value which indicates whether only active or suspended batches should be included. When the value is set to `true`, only suspended batches will be returned and when the value is set to `false`, only active batches will be returned. (optional)
	createdBy := "createdBy_example" // string | Only include batches that were started by this user id. (optional)
	startedBefore := time.Now() // time.Time | Only include batches that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Only include batches that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	withFailures := true // bool | Only include batches having jobs with failures. Value can only be `true`. (optional)
	withoutFailures := true // bool | Only include batches having jobs without failures. Value can only be `true`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.GetBatchStatisticsCount(context.Background()).BatchId(batchId).Type_(type_).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Suspended(suspended).CreatedBy(createdBy).StartedBefore(startedBefore).StartedAfter(startedAfter).WithFailures(withFailures).WithoutFailures(withoutFailures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.GetBatchStatisticsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchStatisticsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.GetBatchStatisticsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchStatisticsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchId** | **string** | Filter by batch id. | 
 **type_** | **string** | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of &#x60;Strings&#x60;. A batch matches if it has one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include batches which belong to no tenant. Value can effectively only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | A &#x60;Boolean&#x60; value which indicates whether only active or suspended batches should be included. When the value is set to &#x60;true&#x60;, only suspended batches will be returned and when the value is set to &#x60;false&#x60;, only active batches will be returned. | 
 **createdBy** | **string** | Only include batches that were started by this user id. | 
 **startedBefore** | **time.Time** | Only include batches that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Only include batches that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **withFailures** | **bool** | Only include batches having jobs with failures. Value can only be &#x60;true&#x60;. | 
 **withoutFailures** | **bool** | Only include batches having jobs without failures. Value can only be &#x60;true&#x60;. | 

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


## GetBatches

> []BatchDto GetBatches(ctx).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).BatchId(batchId).Type_(type_).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Suspended(suspended).CreatedBy(createdBy).StartedBefore(startedBefore).StartedAfter(startedAfter).WithFailures(withFailures).WithoutFailures(withoutFailures).Execute()

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
	batchId := "batchId_example" // string | Filter by batch id. (optional)
	type_ := "type__example" // string | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of `Strings`. A batch matches if it has one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include batches which belong to no tenant. Value can effectively only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | A `Boolean` value which indicates whether only active or suspended batches should be included. When the value is set to `true`, only suspended batches will be returned and when the value is set to `false`, only active batches will be returned. (optional)
	createdBy := "createdBy_example" // string | Only include batches that were started by this user id. (optional)
	startedBefore := time.Now() // time.Time | Only include batches that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Only include batches that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	withFailures := true // bool | Only include batches having jobs with failures. Value can only be `true`. (optional)
	withoutFailures := true // bool | Only include batches having jobs without failures. Value can only be `true`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.GetBatches(context.Background()).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).BatchId(batchId).Type_(type_).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Suspended(suspended).CreatedBy(createdBy).StartedBefore(startedBefore).StartedAfter(startedAfter).WithFailures(withFailures).WithoutFailures(withoutFailures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.GetBatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatches`: []BatchDto
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.GetBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **batchId** | **string** | Filter by batch id. | 
 **type_** | **string** | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of &#x60;Strings&#x60;. A batch matches if it has one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include batches which belong to no tenant. Value can effectively only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | A &#x60;Boolean&#x60; value which indicates whether only active or suspended batches should be included. When the value is set to &#x60;true&#x60;, only suspended batches will be returned and when the value is set to &#x60;false&#x60;, only active batches will be returned. | 
 **createdBy** | **string** | Only include batches that were started by this user id. | 
 **startedBefore** | **time.Time** | Only include batches that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Only include batches that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **withFailures** | **bool** | Only include batches having jobs with failures. Value can only be &#x60;true&#x60;. | 
 **withoutFailures** | **bool** | Only include batches having jobs without failures. Value can only be &#x60;true&#x60;. | 

### Return type

[**[]BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchesCount

> CountResultDto GetBatchesCount(ctx).BatchId(batchId).Type_(type_).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Suspended(suspended).CreatedBy(createdBy).StartedBefore(startedBefore).StartedAfter(startedAfter).WithFailures(withFailures).WithoutFailures(withoutFailures).Execute()

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
	batchId := "batchId_example" // string | Filter by batch id. (optional)
	type_ := "type__example" // string | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of `Strings`. A batch matches if it has one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include batches which belong to no tenant. Value can effectively only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | A `Boolean` value which indicates whether only active or suspended batches should be included. When the value is set to `true`, only suspended batches will be returned and when the value is set to `false`, only active batches will be returned. (optional)
	createdBy := "createdBy_example" // string | Only include batches that were started by this user id. (optional)
	startedBefore := time.Now() // time.Time | Only include batches that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	startedAfter := time.Now() // time.Time | Only include batches that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	withFailures := true // bool | Only include batches having jobs with failures. Value can only be `true`. (optional)
	withoutFailures := true // bool | Only include batches having jobs without failures. Value can only be `true`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.GetBatchesCount(context.Background()).BatchId(batchId).Type_(type_).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Suspended(suspended).CreatedBy(createdBy).StartedBefore(startedBefore).StartedAfter(startedAfter).WithFailures(withFailures).WithoutFailures(withoutFailures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.GetBatchesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.GetBatchesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchId** | **string** | Filter by batch id. | 
 **type_** | **string** | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of &#x60;Strings&#x60;. A batch matches if it has one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include batches which belong to no tenant. Value can effectively only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | A &#x60;Boolean&#x60; value which indicates whether only active or suspended batches should be included. When the value is set to &#x60;true&#x60;, only suspended batches will be returned and when the value is set to &#x60;false&#x60;, only active batches will be returned. | 
 **createdBy** | **string** | Only include batches that were started by this user id. | 
 **startedBefore** | **time.Time** | Only include batches that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **startedAfter** | **time.Time** | Only include batches that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **withFailures** | **bool** | Only include batches having jobs with failures. Value can only be &#x60;true&#x60;. | 
 **withoutFailures** | **bool** | Only include batches having jobs without failures. Value can only be &#x60;true&#x60;. | 

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


## UpdateBatchSuspensionState

> UpdateBatchSuspensionState(ctx, id).SuspensionStateDto(suspensionStateDto).Execute()

Activate/Suspend



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
	id := "id_example" // string | The id of the batch to activate or suspend.
	suspensionStateDto := *openapiclient.NewSuspensionStateDto() // SuspensionStateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchAPI.UpdateBatchSuspensionState(context.Background(), id).SuspensionStateDto(suspensionStateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.UpdateBatchSuspensionState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the batch to activate or suspend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBatchSuspensionStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suspensionStateDto** | [**SuspensionStateDto**](SuspensionStateDto.md) |  | 

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

