# \HistoryCleanupAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanupAsync**](HistoryCleanupAPI.md#CleanupAsync) | **Post** /history/cleanup | Clean up history (POST)
[**FindCleanupJob**](HistoryCleanupAPI.md#FindCleanupJob) | **Get** /history/cleanup/job | Find clean up history job (GET)
[**FindCleanupJobs**](HistoryCleanupAPI.md#FindCleanupJobs) | **Get** /history/cleanup/jobs | Find clean up history jobs (GET)
[**GetHistoryCleanupConfiguration**](HistoryCleanupAPI.md#GetHistoryCleanupConfiguration) | **Get** /history/cleanup/configuration | Get History Cleanup Configuration



## CleanupAsync

> JobDto CleanupAsync(ctx).ImmediatelyDue(immediatelyDue).Execute()

Clean up history (POST)



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
	immediatelyDue := true // bool | When true the job will be scheduled for nearest future. When `false`, the job will be scheduled for next batch window start time. Default is `true`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoryCleanupAPI.CleanupAsync(context.Background()).ImmediatelyDue(immediatelyDue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryCleanupAPI.CleanupAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CleanupAsync`: JobDto
	fmt.Fprintf(os.Stdout, "Response from `HistoryCleanupAPI.CleanupAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCleanupAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **immediatelyDue** | **bool** | When true the job will be scheduled for nearest future. When &#x60;false&#x60;, the job will be scheduled for next batch window start time. Default is &#x60;true&#x60;. | 

### Return type

[**JobDto**](JobDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCleanupJob

> JobDto FindCleanupJob(ctx).Execute()

Find clean up history job (GET)



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
	resp, r, err := apiClient.HistoryCleanupAPI.FindCleanupJob(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryCleanupAPI.FindCleanupJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindCleanupJob`: JobDto
	fmt.Fprintf(os.Stdout, "Response from `HistoryCleanupAPI.FindCleanupJob`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindCleanupJobRequest struct via the builder pattern


### Return type

[**JobDto**](JobDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCleanupJobs

> []JobDto FindCleanupJobs(ctx).Execute()

Find clean up history jobs (GET)



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
	resp, r, err := apiClient.HistoryCleanupAPI.FindCleanupJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryCleanupAPI.FindCleanupJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindCleanupJobs`: []JobDto
	fmt.Fprintf(os.Stdout, "Response from `HistoryCleanupAPI.FindCleanupJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindCleanupJobsRequest struct via the builder pattern


### Return type

[**[]JobDto**](JobDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoryCleanupConfiguration

> HistoryCleanupConfigurationDto GetHistoryCleanupConfiguration(ctx).Execute()

Get History Cleanup Configuration



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
	resp, r, err := apiClient.HistoryCleanupAPI.GetHistoryCleanupConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryCleanupAPI.GetHistoryCleanupConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoryCleanupConfiguration`: HistoryCleanupConfigurationDto
	fmt.Fprintf(os.Stdout, "Response from `HistoryCleanupAPI.GetHistoryCleanupConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryCleanupConfigurationRequest struct via the builder pattern


### Return type

[**HistoryCleanupConfigurationDto**](HistoryCleanupConfigurationDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

