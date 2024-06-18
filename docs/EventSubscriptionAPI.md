# \EventSubscriptionAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventSubscriptions**](EventSubscriptionAPI.md#GetEventSubscriptions) | **Get** /event-subscription | Get List
[**GetEventSubscriptionsCount**](EventSubscriptionAPI.md#GetEventSubscriptionsCount) | **Get** /event-subscription/count | Get List Count



## GetEventSubscriptions

> []EventSubscriptionDto GetEventSubscriptions(ctx).EventSubscriptionId(eventSubscriptionId).EventName(eventName).EventType(eventType).ExecutionId(executionId).ProcessInstanceId(processInstanceId).ActivityId(activityId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeEventSubscriptionsWithoutTenantId(includeEventSubscriptionsWithoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

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
	eventSubscriptionId := "eventSubscriptionId_example" // string | Only select subscription with the given id. (optional)
	eventName := "eventName_example" // string | Only select subscriptions for events with the given name. (optional)
	eventType := "eventType_example" // string | Only select subscriptions for events with the given type. Valid values: `message`, `signal`, `compensate` and `conditional`. (optional)
	executionId := "executionId_example" // string | Only select subscriptions that belong to an execution with the given id. (optional)
	processInstanceId := "processInstanceId_example" // string | Only select subscriptions that belong to a process instance with the given id. (optional)
	activityId := "activityId_example" // string | Only select subscriptions that belong to an activity with the given id. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. Only select subscriptions that belong to one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only select subscriptions which have no tenant id. Value may only be `true`, as `false` is the default behavior. (optional)
	includeEventSubscriptionsWithoutTenantId := true // bool | Select event subscriptions which have no tenant id. Can be used in combination with tenantIdIn parameter. Value may only be `true`, as `false` is the default behavior. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventSubscriptionAPI.GetEventSubscriptions(context.Background()).EventSubscriptionId(eventSubscriptionId).EventName(eventName).EventType(eventType).ExecutionId(executionId).ProcessInstanceId(processInstanceId).ActivityId(activityId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeEventSubscriptionsWithoutTenantId(includeEventSubscriptionsWithoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventSubscriptionAPI.GetEventSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventSubscriptions`: []EventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `EventSubscriptionAPI.GetEventSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventSubscriptionId** | **string** | Only select subscription with the given id. | 
 **eventName** | **string** | Only select subscriptions for events with the given name. | 
 **eventType** | **string** | Only select subscriptions for events with the given type. Valid values: &#x60;message&#x60;, &#x60;signal&#x60;, &#x60;compensate&#x60; and &#x60;conditional&#x60;. | 
 **executionId** | **string** | Only select subscriptions that belong to an execution with the given id. | 
 **processInstanceId** | **string** | Only select subscriptions that belong to a process instance with the given id. | 
 **activityId** | **string** | Only select subscriptions that belong to an activity with the given id. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. Only select subscriptions that belong to one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only select subscriptions which have no tenant id. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeEventSubscriptionsWithoutTenantId** | **bool** | Select event subscriptions which have no tenant id. Can be used in combination with tenantIdIn parameter. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]EventSubscriptionDto**](EventSubscriptionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSubscriptionsCount

> CountResultDto GetEventSubscriptionsCount(ctx).EventSubscriptionId(eventSubscriptionId).EventName(eventName).EventType(eventType).ExecutionId(executionId).ProcessInstanceId(processInstanceId).ActivityId(activityId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeEventSubscriptionsWithoutTenantId(includeEventSubscriptionsWithoutTenantId).Execute()

Get List Count



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
	eventSubscriptionId := "eventSubscriptionId_example" // string | Only select subscription with the given id. (optional)
	eventName := "eventName_example" // string | Only select subscriptions for events with the given name. (optional)
	eventType := "eventType_example" // string | Only select subscriptions for events with the given type. Valid values: `message`, `signal`, `compensate` and `conditional`. (optional)
	executionId := "executionId_example" // string | Only select subscriptions that belong to an execution with the given id. (optional)
	processInstanceId := "processInstanceId_example" // string | Only select subscriptions that belong to a process instance with the given id. (optional)
	activityId := "activityId_example" // string | Only select subscriptions that belong to an activity with the given id. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. Only select subscriptions that belong to one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only select subscriptions which have no tenant id. Value may only be `true`, as `false` is the default behavior. (optional)
	includeEventSubscriptionsWithoutTenantId := true // bool | Select event subscriptions which have no tenant id. Can be used in combination with tenantIdIn parameter. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventSubscriptionAPI.GetEventSubscriptionsCount(context.Background()).EventSubscriptionId(eventSubscriptionId).EventName(eventName).EventType(eventType).ExecutionId(executionId).ProcessInstanceId(processInstanceId).ActivityId(activityId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeEventSubscriptionsWithoutTenantId(includeEventSubscriptionsWithoutTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventSubscriptionAPI.GetEventSubscriptionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventSubscriptionsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `EventSubscriptionAPI.GetEventSubscriptionsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventSubscriptionsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventSubscriptionId** | **string** | Only select subscription with the given id. | 
 **eventName** | **string** | Only select subscriptions for events with the given name. | 
 **eventType** | **string** | Only select subscriptions for events with the given type. Valid values: &#x60;message&#x60;, &#x60;signal&#x60;, &#x60;compensate&#x60; and &#x60;conditional&#x60;. | 
 **executionId** | **string** | Only select subscriptions that belong to an execution with the given id. | 
 **processInstanceId** | **string** | Only select subscriptions that belong to a process instance with the given id. | 
 **activityId** | **string** | Only select subscriptions that belong to an activity with the given id. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. Only select subscriptions that belong to one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only select subscriptions which have no tenant id. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeEventSubscriptionsWithoutTenantId** | **bool** | Select event subscriptions which have no tenant id. Can be used in combination with tenantIdIn parameter. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

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

