# \ModificationAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteModification**](ModificationAPI.md#ExecuteModification) | **Post** /modification/execute | Execute Modification
[**ExecuteModificationAsync**](ModificationAPI.md#ExecuteModificationAsync) | **Post** /modification/executeAsync | Execute Modification Async (Batch)



## ExecuteModification

> ExecuteModification(ctx).ModificationDto(modificationDto).Execute()

Execute Modification



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
	modificationDto := *openapiclient.NewModificationDto() // ModificationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ModificationAPI.ExecuteModification(context.Background()).ModificationDto(modificationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModificationAPI.ExecuteModification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteModificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modificationDto** | [**ModificationDto**](ModificationDto.md) |  | 

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


## ExecuteModificationAsync

> BatchDto ExecuteModificationAsync(ctx).ModificationDto(modificationDto).Execute()

Execute Modification Async (Batch)



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
	modificationDto := *openapiclient.NewModificationDto() // ModificationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModificationAPI.ExecuteModificationAsync(context.Background()).ModificationDto(modificationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModificationAPI.ExecuteModificationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteModificationAsync`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `ModificationAPI.ExecuteModificationAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteModificationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modificationDto** | [**ModificationDto**](ModificationDto.md) |  | 

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

