# \HistoricBatchAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHistoricBatch**](HistoricBatchAPI.md#DeleteHistoricBatch) | **Delete** /history/batch/{id} | Delete Historic Batch
[**GetCleanableHistoricBatchesReport**](HistoricBatchAPI.md#GetCleanableHistoricBatchesReport) | **Get** /history/batch/cleanable-batch-report | Get Cleanable Batch Report
[**GetCleanableHistoricBatchesReportCount**](HistoricBatchAPI.md#GetCleanableHistoricBatchesReportCount) | **Get** /history/batch/cleanable-batch-report/count | Get Cleanable Batch Report Count
[**GetHistoricBatch**](HistoricBatchAPI.md#GetHistoricBatch) | **Get** /history/batch/{id} | Get Historic Batch
[**GetHistoricBatches**](HistoricBatchAPI.md#GetHistoricBatches) | **Get** /history/batch | Get Historic Batches
[**GetHistoricBatchesCount**](HistoricBatchAPI.md#GetHistoricBatchesCount) | **Get** /history/batch/count | Get Historic Batch Count
[**SetRemovalTimeAsyncHistoricBatch**](HistoricBatchAPI.md#SetRemovalTimeAsyncHistoricBatch) | **Post** /history/batch/set-removal-time | Set Removal Time Async (POST)



## DeleteHistoricBatch

> DeleteHistoricBatch(ctx, id).Execute()

Delete Historic Batch



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HistoricBatchAPI.DeleteHistoricBatch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricBatchAPI.DeleteHistoricBatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteHistoricBatchRequest struct via the builder pattern


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


## GetCleanableHistoricBatchesReport

> []CleanableHistoricBatchReportResultDto GetCleanableHistoricBatchesReport(ctx).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Cleanable Batch Report



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
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricBatchAPI.GetCleanableHistoricBatchesReport(context.Background()).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricBatchAPI.GetCleanableHistoricBatchesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCleanableHistoricBatchesReport`: []CleanableHistoricBatchReportResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricBatchAPI.GetCleanableHistoricBatchesReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCleanableHistoricBatchesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]CleanableHistoricBatchReportResultDto**](CleanableHistoricBatchReportResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCleanableHistoricBatchesReportCount

> CountResultDto GetCleanableHistoricBatchesReportCount(ctx).Execute()

Get Cleanable Batch Report Count



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
	resp, r, err := apiClient.HistoricBatchAPI.GetCleanableHistoricBatchesReportCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricBatchAPI.GetCleanableHistoricBatchesReportCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCleanableHistoricBatchesReportCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricBatchAPI.GetCleanableHistoricBatchesReportCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCleanableHistoricBatchesReportCountRequest struct via the builder pattern


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


## GetHistoricBatch

> HistoricBatchDto GetHistoricBatch(ctx, id).Execute()

Get Historic Batch



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
	id := "id_example" // string | The id of the historic batch to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricBatchAPI.GetHistoricBatch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricBatchAPI.GetHistoricBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricBatch`: HistoricBatchDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricBatchAPI.GetHistoricBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the historic batch to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HistoricBatchDto**](HistoricBatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricBatches

> []HistoricBatchDto GetHistoricBatches(ctx).BatchId(batchId).Type_(type_).Completed(completed).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Historic Batches



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
	batchId := "batchId_example" // string | Filter by batch id. (optional)
	type_ := "type__example" // string | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. (optional)
	completed := true // bool |  Filter completed or not completed batches. If the value is `true`, only completed batches, i.e., end time is set, are returned. Otherwise, if the value is `false`, only running batches, i.e., end time is null, are returned. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A batch matches if it has one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include batches which belong to no tenant. Value can effectively only be `true`, as `false` is the default behavior. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricBatchAPI.GetHistoricBatches(context.Background()).BatchId(batchId).Type_(type_).Completed(completed).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricBatchAPI.GetHistoricBatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricBatches`: []HistoricBatchDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricBatchAPI.GetHistoricBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchId** | **string** | Filter by batch id. | 
 **type_** | **string** | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. | 
 **completed** | **bool** |  Filter completed or not completed batches. If the value is &#x60;true&#x60;, only completed batches, i.e., end time is set, are returned. Otherwise, if the value is &#x60;false&#x60;, only running batches, i.e., end time is null, are returned. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A batch matches if it has one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include batches which belong to no tenant. Value can effectively only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]HistoricBatchDto**](HistoricBatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricBatchesCount

> CountResultDto GetHistoricBatchesCount(ctx).BatchId(batchId).Type_(type_).Completed(completed).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Execute()

Get Historic Batch Count



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
	batchId := "batchId_example" // string | Filter by batch id. (optional)
	type_ := "type__example" // string | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. (optional)
	completed := true // bool |  Filter completed or not completed batches. If the value is `true`, only completed batches, i.e., end time is set, are returned. Otherwise, if the value is `false`, only running batches, i.e., end time is null, are returned. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A batch matches if it has one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include batches which belong to no tenant. Value can effectively only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricBatchAPI.GetHistoricBatchesCount(context.Background()).BatchId(batchId).Type_(type_).Completed(completed).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricBatchAPI.GetHistoricBatchesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricBatchesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricBatchAPI.GetHistoricBatchesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricBatchesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchId** | **string** | Filter by batch id. | 
 **type_** | **string** | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. | 
 **completed** | **bool** |  Filter completed or not completed batches. If the value is &#x60;true&#x60;, only completed batches, i.e., end time is set, are returned. Otherwise, if the value is &#x60;false&#x60;, only running batches, i.e., end time is null, are returned. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A batch matches if it has one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include batches which belong to no tenant. Value can effectively only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

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


## SetRemovalTimeAsyncHistoricBatch

> BatchDto SetRemovalTimeAsyncHistoricBatch(ctx).SetRemovalTimeToHistoricBatchesDto(setRemovalTimeToHistoricBatchesDto).Execute()

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
	setRemovalTimeToHistoricBatchesDto := *openapiclient.NewSetRemovalTimeToHistoricBatchesDto() // SetRemovalTimeToHistoricBatchesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricBatchAPI.SetRemovalTimeAsyncHistoricBatch(context.Background()).SetRemovalTimeToHistoricBatchesDto(setRemovalTimeToHistoricBatchesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricBatchAPI.SetRemovalTimeAsyncHistoricBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRemovalTimeAsyncHistoricBatch`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricBatchAPI.SetRemovalTimeAsyncHistoricBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRemovalTimeAsyncHistoricBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setRemovalTimeToHistoricBatchesDto** | [**SetRemovalTimeToHistoricBatchesDto**](SetRemovalTimeToHistoricBatchesDto.md) |  | 

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

