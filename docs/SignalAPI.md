# \SignalAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThrowSignal**](SignalAPI.md#ThrowSignal) | **Post** /signal | Event



## ThrowSignal

> ThrowSignal(ctx).SignalDto(signalDto).Execute()

Event



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
	signalDto := *openapiclient.NewSignalDto() // SignalDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SignalAPI.ThrowSignal(context.Background()).SignalDto(signalDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalAPI.ThrowSignal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThrowSignalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signalDto** | [**SignalDto**](SignalDto.md) |  | 

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

