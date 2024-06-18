# \HistoricDecisionDefinitionAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCleanableHistoricDecisionInstanceReport**](HistoricDecisionDefinitionAPI.md#GetCleanableHistoricDecisionInstanceReport) | **Get** /history/decision-definition/cleanable-decision-instance-report | Get Cleanable Decision Instance Report
[**GetCleanableHistoricDecisionInstanceReportCount**](HistoricDecisionDefinitionAPI.md#GetCleanableHistoricDecisionInstanceReportCount) | **Get** /history/decision-definition/cleanable-decision-instance-report/count | Get Cleanable Decision Instance Report Count



## GetCleanableHistoricDecisionInstanceReport

> []CleanableHistoricDecisionInstanceReportResultDto GetCleanableHistoricDecisionInstanceReport(ctx).DecisionDefinitionIdIn(decisionDefinitionIdIn).DecisionDefinitionKeyIn(decisionDefinitionKeyIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Compact(compact).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Cleanable Decision Instance Report



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
	decisionDefinitionIdIn := "decisionDefinitionIdIn_example" // string | Filter by decision definition ids. Must be a comma-separated list of decision definition ids. (optional)
	decisionDefinitionKeyIn := "decisionDefinitionKeyIn_example" // string | Filter by decision definition keys. Must be a comma-separated list of decision definition keys. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A decision definition must have one of the given tenant  ids. (optional)
	withoutTenantId := true // bool | Only include decision definitions which belong to no tenant. Value may only be `true`, as `false`  is the default behavior. (optional)
	compact := true // bool | Only include decision instances which have more than zero finished instances. Value may only be `true`,  as `false` is the default behavior. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDecisionDefinitionAPI.GetCleanableHistoricDecisionInstanceReport(context.Background()).DecisionDefinitionIdIn(decisionDefinitionIdIn).DecisionDefinitionKeyIn(decisionDefinitionKeyIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Compact(compact).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDecisionDefinitionAPI.GetCleanableHistoricDecisionInstanceReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCleanableHistoricDecisionInstanceReport`: []CleanableHistoricDecisionInstanceReportResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDecisionDefinitionAPI.GetCleanableHistoricDecisionInstanceReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCleanableHistoricDecisionInstanceReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decisionDefinitionIdIn** | **string** | Filter by decision definition ids. Must be a comma-separated list of decision definition ids. | 
 **decisionDefinitionKeyIn** | **string** | Filter by decision definition keys. Must be a comma-separated list of decision definition keys. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A decision definition must have one of the given tenant  ids. | 
 **withoutTenantId** | **bool** | Only include decision definitions which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60;  is the default behavior. | 
 **compact** | **bool** | Only include decision instances which have more than zero finished instances. Value may only be &#x60;true&#x60;,  as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]CleanableHistoricDecisionInstanceReportResultDto**](CleanableHistoricDecisionInstanceReportResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCleanableHistoricDecisionInstanceReportCount

> CountResultDto GetCleanableHistoricDecisionInstanceReportCount(ctx).DecisionDefinitionIdIn(decisionDefinitionIdIn).DecisionDefinitionKeyIn(decisionDefinitionKeyIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Compact(compact).Execute()

Get Cleanable Decision Instance Report Count



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
	decisionDefinitionIdIn := "decisionDefinitionIdIn_example" // string | Filter by decision definition ids. Must be a comma-separated list of decision definition ids. (optional)
	decisionDefinitionKeyIn := "decisionDefinitionKeyIn_example" // string | Filter by decision definition keys. Must be a comma-separated list of decision definition keys. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A decision definition must have one of the given tenant  ids. (optional)
	withoutTenantId := true // bool | Only include decision definitions which belong to no tenant. Value may only be `true`, as `false`  is the default behavior. (optional)
	compact := true // bool | Only include decision instances which have more than zero finished instances. Value may only be `true`,  as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDecisionDefinitionAPI.GetCleanableHistoricDecisionInstanceReportCount(context.Background()).DecisionDefinitionIdIn(decisionDefinitionIdIn).DecisionDefinitionKeyIn(decisionDefinitionKeyIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Compact(compact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDecisionDefinitionAPI.GetCleanableHistoricDecisionInstanceReportCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCleanableHistoricDecisionInstanceReportCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDecisionDefinitionAPI.GetCleanableHistoricDecisionInstanceReportCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCleanableHistoricDecisionInstanceReportCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decisionDefinitionIdIn** | **string** | Filter by decision definition ids. Must be a comma-separated list of decision definition ids. | 
 **decisionDefinitionKeyIn** | **string** | Filter by decision definition keys. Must be a comma-separated list of decision definition keys. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A decision definition must have one of the given tenant  ids. | 
 **withoutTenantId** | **bool** | Only include decision definitions which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60;  is the default behavior. | 
 **compact** | **bool** | Only include decision instances which have more than zero finished instances. Value may only be &#x60;true&#x60;,  as &#x60;false&#x60; is the default behavior. | 

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

