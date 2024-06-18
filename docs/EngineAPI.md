# \EngineAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProcessEngineNames**](EngineAPI.md#GetProcessEngineNames) | **Get** /engine | Get List



## GetProcessEngineNames

> []ProcessEngineDto GetProcessEngineNames(ctx).Execute()

Get List



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
	resp, r, err := apiClient.EngineAPI.GetProcessEngineNames(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.GetProcessEngineNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessEngineNames`: []ProcessEngineDto
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.GetProcessEngineNames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessEngineNamesRequest struct via the builder pattern


### Return type

[**[]ProcessEngineDto**](ProcessEngineDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

