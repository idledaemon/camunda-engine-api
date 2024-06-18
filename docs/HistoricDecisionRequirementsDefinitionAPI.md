# \HistoricDecisionRequirementsDefinitionAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDecisionStatistics**](HistoricDecisionRequirementsDefinitionAPI.md#GetDecisionStatistics) | **Get** /history/decision-requirements-definition/{id}/statistics | Get DRD Statistics



## GetDecisionStatistics

> []HistoricDecisionInstanceStatisticsDto GetDecisionStatistics(ctx, id).DecisionInstanceId(decisionInstanceId).Execute()

Get DRD Statistics



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
	id := "id_example" // string | The id of the decision requirements definition.
	decisionInstanceId := "decisionInstanceId_example" // string | Restrict query results to be based only on specific evaluation instance of a given decision requirements definition. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDecisionRequirementsDefinitionAPI.GetDecisionStatistics(context.Background(), id).DecisionInstanceId(decisionInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDecisionRequirementsDefinitionAPI.GetDecisionStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionStatistics`: []HistoricDecisionInstanceStatisticsDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDecisionRequirementsDefinitionAPI.GetDecisionStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the decision requirements definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decisionInstanceId** | **string** | Restrict query results to be based only on specific evaluation instance of a given decision requirements definition. | 

### Return type

[**[]HistoricDecisionInstanceStatisticsDto**](HistoricDecisionInstanceStatisticsDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

