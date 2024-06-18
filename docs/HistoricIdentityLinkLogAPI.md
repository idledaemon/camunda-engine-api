# \HistoricIdentityLinkLogAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricIdentityLinks**](HistoricIdentityLinkLogAPI.md#GetHistoricIdentityLinks) | **Get** /history/identity-link-log | Get Identity Link Logs
[**GetHistoricIdentityLinksCount**](HistoricIdentityLinkLogAPI.md#GetHistoricIdentityLinksCount) | **Get** /history/identity-link-log/count | Get Identity Link Log Count



## GetHistoricIdentityLinks

> []HistoricIdentityLinkLogDto GetHistoricIdentityLinks(ctx).Type_(type_).UserId(userId).GroupId(groupId).DateBefore(dateBefore).DateAfter(dateAfter).TaskId(taskId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).OperationType(operationType).AssignerId(assignerId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Identity Link Logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	type_ := "type__example" // string | Restricts to identity links that have the given type (candidate/assignee/owner). (optional)
	userId := "userId_example" // string | Restricts to identity links that have the given user id. (optional)
	groupId := "groupId_example" // string | Restricts to identity links that have the given group id. (optional)
	dateBefore := time.Now() // time.Time | Restricts to identity links that have the time before the given time. (optional)
	dateAfter := time.Now() // time.Time | Restricts to identity links that have the time after the given time. (optional)
	taskId := "taskId_example" // string | Restricts to identity links that have the given task id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Restricts to identity links that have the given process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Restricts to identity links that have the given process definition key. (optional)
	operationType := "operationType_example" // string | Restricts to identity links that have the given operationType (add/delete). (optional)
	assignerId := "assignerId_example" // string | Restricts to identity links that have the given assigner id. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic identity links that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricIdentityLinkLogAPI.GetHistoricIdentityLinks(context.Background()).Type_(type_).UserId(userId).GroupId(groupId).DateBefore(dateBefore).DateAfter(dateAfter).TaskId(taskId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).OperationType(operationType).AssignerId(assignerId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricIdentityLinkLogAPI.GetHistoricIdentityLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricIdentityLinks`: []HistoricIdentityLinkLogDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricIdentityLinkLogAPI.GetHistoricIdentityLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricIdentityLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Restricts to identity links that have the given type (candidate/assignee/owner). | 
 **userId** | **string** | Restricts to identity links that have the given user id. | 
 **groupId** | **string** | Restricts to identity links that have the given group id. | 
 **dateBefore** | **time.Time** | Restricts to identity links that have the time before the given time. | 
 **dateAfter** | **time.Time** | Restricts to identity links that have the time after the given time. | 
 **taskId** | **string** | Restricts to identity links that have the given task id. | 
 **processDefinitionId** | **string** | Restricts to identity links that have the given process definition id. | 
 **processDefinitionKey** | **string** | Restricts to identity links that have the given process definition key. | 
 **operationType** | **string** | Restricts to identity links that have the given operationType (add/delete). | 
 **assignerId** | **string** | Restricts to identity links that have the given assigner id. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic identity links that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]HistoricIdentityLinkLogDto**](HistoricIdentityLinkLogDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricIdentityLinksCount

> CountResultDto GetHistoricIdentityLinksCount(ctx).Type_(type_).UserId(userId).GroupId(groupId).DateBefore(dateBefore).DateAfter(dateAfter).TaskId(taskId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).OperationType(operationType).AssignerId(assignerId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Execute()

Get Identity Link Log Count



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	type_ := "type__example" // string | Restricts to identity links that have the given type (candidate/assignee/owner). (optional)
	userId := "userId_example" // string | Restricts to identity links that have the given user id. (optional)
	groupId := "groupId_example" // string | Restricts to identity links that have the given group id. (optional)
	dateBefore := time.Now() // time.Time | Restricts to identity links that have the time before the given time. (optional)
	dateAfter := time.Now() // time.Time | Restricts to identity links that have the time after the given time. (optional)
	taskId := "taskId_example" // string | Restricts to identity links that have the given task id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Restricts to identity links that have the given process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Restricts to identity links that have the given process definition key. (optional)
	operationType := "operationType_example" // string | Restricts to identity links that have the given operationType (add/delete). (optional)
	assignerId := "assignerId_example" // string | Restricts to identity links that have the given assigner id. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic identity links that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricIdentityLinkLogAPI.GetHistoricIdentityLinksCount(context.Background()).Type_(type_).UserId(userId).GroupId(groupId).DateBefore(dateBefore).DateAfter(dateAfter).TaskId(taskId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).OperationType(operationType).AssignerId(assignerId).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricIdentityLinkLogAPI.GetHistoricIdentityLinksCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricIdentityLinksCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricIdentityLinkLogAPI.GetHistoricIdentityLinksCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricIdentityLinksCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Restricts to identity links that have the given type (candidate/assignee/owner). | 
 **userId** | **string** | Restricts to identity links that have the given user id. | 
 **groupId** | **string** | Restricts to identity links that have the given group id. | 
 **dateBefore** | **time.Time** | Restricts to identity links that have the time before the given time. | 
 **dateAfter** | **time.Time** | Restricts to identity links that have the time after the given time. | 
 **taskId** | **string** | Restricts to identity links that have the given task id. | 
 **processDefinitionId** | **string** | Restricts to identity links that have the given process definition id. | 
 **processDefinitionKey** | **string** | Restricts to identity links that have the given process definition key. | 
 **operationType** | **string** | Restricts to identity links that have the given operationType (add/delete). | 
 **assignerId** | **string** | Restricts to identity links that have the given assigner id. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic identity links that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

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

