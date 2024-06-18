# \ConditionAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EvaluateCondition**](ConditionAPI.md#EvaluateCondition) | **Post** /condition | Evaluate



## EvaluateCondition

> []ProcessInstanceDto EvaluateCondition(ctx).EvaluationConditionDto(evaluationConditionDto).Execute()

Evaluate



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
	evaluationConditionDto := *openapiclient.NewEvaluationConditionDto() // EvaluationConditionDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConditionAPI.EvaluateCondition(context.Background()).EvaluationConditionDto(evaluationConditionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConditionAPI.EvaluateCondition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateCondition`: []ProcessInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ConditionAPI.EvaluateCondition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evaluationConditionDto** | [**EvaluationConditionDto**](EvaluationConditionDto.md) |  | 

### Return type

[**[]ProcessInstanceDto**](ProcessInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

