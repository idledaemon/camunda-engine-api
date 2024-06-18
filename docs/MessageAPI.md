# \MessageAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliverMessage**](MessageAPI.md#DeliverMessage) | **Post** /message | Correlate



## DeliverMessage

> []MessageCorrelationResultWithVariableDto DeliverMessage(ctx).CorrelationMessageDto(correlationMessageDto).Execute()

Correlate



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
	correlationMessageDto := *openapiclient.NewCorrelationMessageDto() // CorrelationMessageDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.DeliverMessage(context.Background()).CorrelationMessageDto(correlationMessageDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.DeliverMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeliverMessage`: []MessageCorrelationResultWithVariableDto
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.DeliverMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliverMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **correlationMessageDto** | [**CorrelationMessageDto**](CorrelationMessageDto.md) |  | 

### Return type

[**[]MessageCorrelationResultWithVariableDto**](MessageCorrelationResultWithVariableDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

