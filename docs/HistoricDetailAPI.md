# \HistoricDetailAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricDetails**](HistoricDetailAPI.md#GetHistoricDetails) | **Get** /history/detail | Get Historic Details
[**GetHistoricDetailsCount**](HistoricDetailAPI.md#GetHistoricDetailsCount) | **Get** /history/detail/count | Get Historic Detail Count
[**HistoricDetail**](HistoricDetailAPI.md#HistoricDetail) | **Get** /history/detail/{id} | Get Historic Detail
[**HistoricDetailBinary**](HistoricDetailAPI.md#HistoricDetailBinary) | **Get** /history/detail/{id}/data | Get Historic Detail (Binary)
[**QueryHistoricDetails**](HistoricDetailAPI.md#QueryHistoricDetails) | **Post** /history/detail | Get Historic Details (POST)



## GetHistoricDetails

> []HistoricDetailDto GetHistoricDetails(ctx).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ExecutionId(executionId).TaskId(taskId).ActivityInstanceId(activityInstanceId).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).VariableInstanceId(variableInstanceId).VariableTypeIn(variableTypeIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).UserOperationId(userOperationId).FormFields(formFields).VariableUpdates(variableUpdates).ExcludeTaskDetails(excludeTaskDetails).Initial(initial).OccurredBefore(occurredBefore).OccurredAfter(occurredAfter).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).Execute()

Get Historic Details



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
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processInstanceIdIn := "processInstanceIdIn_example" // string | Only include historic details which belong to one of the passed comma-separated process instance ids. (optional)
	executionId := "executionId_example" // string | Filter by execution id. (optional)
	taskId := "taskId_example" // string | Filter by task id. (optional)
	activityInstanceId := "activityInstanceId_example" // string | Filter by activity instance id. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by case instance id. (optional)
	caseExecutionId := "caseExecutionId_example" // string | Filter by case execution id. (optional)
	variableInstanceId := "variableInstanceId_example" // string | Filter by variable instance id. (optional)
	variableTypeIn := "variableTypeIn_example" // string | Only include historic details where the variable updates belong to one of the passed comma-separated list of variable types. A list of all supported variable types can be found [here](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#supported-variable-values). **Note:** All non-primitive variables are associated with the type `serializable`. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic details that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	userOperationId := "userOperationId_example" // string | Filter by a user operation id. (optional)
	formFields := true // bool | Only include `HistoricFormFields`. Value may only be `true`, as `false` is the default behavior. (optional)
	variableUpdates := true // bool | Only include `HistoricVariableUpdates`. Value may only be `true`, as `false` is the default behavior. (optional)
	excludeTaskDetails := true // bool | Excludes all task-related `HistoricDetails`, so only items which have no task id set will be selected. When this parameter is used together with `taskId`, this call is ignored and task details are not excluded. Value may only be `true`, as `false` is the default behavior. (optional)
	initial := true // bool | Restrict to historic variable updates that contain only initial variable values. Value may only be `true`, as `false` is the default behavior. (optional)
	occurredBefore := time.Now() // time.Time | Restrict to historic details that occured before the given date (including the date). Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., 2013-01-23T14:42:45.000+0200. (optional)
	occurredAfter := time.Now() // time.Time | Restrict to historic details that occured after the given date (including the date). Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., 2013-01-23T14:42:45.000+0200. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDetailAPI.GetHistoricDetails(context.Background()).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ExecutionId(executionId).TaskId(taskId).ActivityInstanceId(activityInstanceId).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).VariableInstanceId(variableInstanceId).VariableTypeIn(variableTypeIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).UserOperationId(userOperationId).FormFields(formFields).VariableUpdates(variableUpdates).ExcludeTaskDetails(excludeTaskDetails).Initial(initial).OccurredBefore(occurredBefore).OccurredAfter(occurredAfter).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDetailAPI.GetHistoricDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricDetails`: []HistoricDetailDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDetailAPI.GetHistoricDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processInstanceIdIn** | **string** | Only include historic details which belong to one of the passed comma-separated process instance ids. | 
 **executionId** | **string** | Filter by execution id. | 
 **taskId** | **string** | Filter by task id. | 
 **activityInstanceId** | **string** | Filter by activity instance id. | 
 **caseInstanceId** | **string** | Filter by case instance id. | 
 **caseExecutionId** | **string** | Filter by case execution id. | 
 **variableInstanceId** | **string** | Filter by variable instance id. | 
 **variableTypeIn** | **string** | Only include historic details where the variable updates belong to one of the passed comma-separated list of variable types. A list of all supported variable types can be found [here](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#supported-variable-values). **Note:** All non-primitive variables are associated with the type &#x60;serializable&#x60;. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic details that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **userOperationId** | **string** | Filter by a user operation id. | 
 **formFields** | **bool** | Only include &#x60;HistoricFormFields&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **variableUpdates** | **bool** | Only include &#x60;HistoricVariableUpdates&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **excludeTaskDetails** | **bool** | Excludes all task-related &#x60;HistoricDetails&#x60;, so only items which have no task id set will be selected. When this parameter is used together with &#x60;taskId&#x60;, this call is ignored and task details are not excluded. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **initial** | **bool** | Restrict to historic variable updates that contain only initial variable values. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **occurredBefore** | **time.Time** | Restrict to historic details that occured before the given date (including the date). Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200. | 
 **occurredAfter** | **time.Time** | Restrict to historic details that occured after the given date (including the date). Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | 

### Return type

[**[]HistoricDetailDto**](HistoricDetailDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricDetailsCount

> CountResultDto GetHistoricDetailsCount(ctx).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ExecutionId(executionId).TaskId(taskId).ActivityInstanceId(activityInstanceId).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).VariableInstanceId(variableInstanceId).VariableTypeIn(variableTypeIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).UserOperationId(userOperationId).FormFields(formFields).VariableUpdates(variableUpdates).ExcludeTaskDetails(excludeTaskDetails).Initial(initial).OccurredBefore(occurredBefore).OccurredAfter(occurredAfter).Execute()

Get Historic Detail Count



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
	processInstanceId := "processInstanceId_example" // string | Filter by process instance id. (optional)
	processInstanceIdIn := "processInstanceIdIn_example" // string | Only include historic details which belong to one of the passed comma-separated process instance ids. (optional)
	executionId := "executionId_example" // string | Filter by execution id. (optional)
	taskId := "taskId_example" // string | Filter by task id. (optional)
	activityInstanceId := "activityInstanceId_example" // string | Filter by activity instance id. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by case instance id. (optional)
	caseExecutionId := "caseExecutionId_example" // string | Filter by case execution id. (optional)
	variableInstanceId := "variableInstanceId_example" // string | Filter by variable instance id. (optional)
	variableTypeIn := "variableTypeIn_example" // string | Only include historic details where the variable updates belong to one of the passed comma-separated list of variable types. A list of all supported variable types can be found [here](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#supported-variable-values). **Note:** All non-primitive variables are associated with the type `serializable`. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic details that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	userOperationId := "userOperationId_example" // string | Filter by a user operation id. (optional)
	formFields := true // bool | Only include `HistoricFormFields`. Value may only be `true`, as `false` is the default behavior. (optional)
	variableUpdates := true // bool | Only include `HistoricVariableUpdates`. Value may only be `true`, as `false` is the default behavior. (optional)
	excludeTaskDetails := true // bool | Excludes all task-related `HistoricDetails`, so only items which have no task id set will be selected. When this parameter is used together with `taskId`, this call is ignored and task details are not excluded. Value may only be `true`, as `false` is the default behavior. (optional)
	initial := true // bool | Restrict to historic variable updates that contain only initial variable values. Value may only be `true`, as `false` is the default behavior. (optional)
	occurredBefore := time.Now() // time.Time | Restrict to historic details that occured before the given date (including the date). Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., 2013-01-23T14:42:45.000+0200. (optional)
	occurredAfter := time.Now() // time.Time | Restrict to historic details that occured after the given date (including the date). Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., 2013-01-23T14:42:45.000+0200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDetailAPI.GetHistoricDetailsCount(context.Background()).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ExecutionId(executionId).TaskId(taskId).ActivityInstanceId(activityInstanceId).CaseInstanceId(caseInstanceId).CaseExecutionId(caseExecutionId).VariableInstanceId(variableInstanceId).VariableTypeIn(variableTypeIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).UserOperationId(userOperationId).FormFields(formFields).VariableUpdates(variableUpdates).ExcludeTaskDetails(excludeTaskDetails).Initial(initial).OccurredBefore(occurredBefore).OccurredAfter(occurredAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDetailAPI.GetHistoricDetailsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricDetailsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDetailAPI.GetHistoricDetailsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricDetailsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processInstanceId** | **string** | Filter by process instance id. | 
 **processInstanceIdIn** | **string** | Only include historic details which belong to one of the passed comma-separated process instance ids. | 
 **executionId** | **string** | Filter by execution id. | 
 **taskId** | **string** | Filter by task id. | 
 **activityInstanceId** | **string** | Filter by activity instance id. | 
 **caseInstanceId** | **string** | Filter by case instance id. | 
 **caseExecutionId** | **string** | Filter by case execution id. | 
 **variableInstanceId** | **string** | Filter by variable instance id. | 
 **variableTypeIn** | **string** | Only include historic details where the variable updates belong to one of the passed comma-separated list of variable types. A list of all supported variable types can be found [here](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#supported-variable-values). **Note:** All non-primitive variables are associated with the type &#x60;serializable&#x60;. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic details that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **userOperationId** | **string** | Filter by a user operation id. | 
 **formFields** | **bool** | Only include &#x60;HistoricFormFields&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **variableUpdates** | **bool** | Only include &#x60;HistoricVariableUpdates&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **excludeTaskDetails** | **bool** | Excludes all task-related &#x60;HistoricDetails&#x60;, so only items which have no task id set will be selected. When this parameter is used together with &#x60;taskId&#x60;, this call is ignored and task details are not excluded. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **initial** | **bool** | Restrict to historic variable updates that contain only initial variable values. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **occurredBefore** | **time.Time** | Restrict to historic details that occured before the given date (including the date). Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200. | 
 **occurredAfter** | **time.Time** | Restrict to historic details that occured after the given date (including the date). Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200. | 

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


## HistoricDetail

> HistoricDetailDto HistoricDetail(ctx, id).DeserializeValue(deserializeValue).Execute()

Get Historic Detail



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
	id := "id_example" // string | The id of the detail.
	deserializeValue := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDetailAPI.HistoricDetail(context.Background(), id).DeserializeValue(deserializeValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDetailAPI.HistoricDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HistoricDetail`: HistoricDetailDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDetailAPI.HistoricDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the detail. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHistoricDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deserializeValue** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | 

### Return type

[**HistoricDetailDto**](HistoricDetailDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HistoricDetailBinary

> *os.File HistoricDetailBinary(ctx, id).Execute()

Get Historic Detail (Binary)



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
	id := "id_example" // string | The id of the historic variable update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDetailAPI.HistoricDetailBinary(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDetailAPI.HistoricDetailBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HistoricDetailBinary`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `HistoricDetailAPI.HistoricDetailBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the historic variable update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHistoricDetailBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryHistoricDetails

> []HistoricDetailDto QueryHistoricDetails(ctx).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).HistoricDetailQueryDto(historicDetailQueryDto).Execute()

Get Historic Details (POST)



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
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional)
	historicDetailQueryDto := *openapiclient.NewHistoricDetailQueryDto() // HistoricDetailQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricDetailAPI.QueryHistoricDetails(context.Background()).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).HistoricDetailQueryDto(historicDetailQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricDetailAPI.QueryHistoricDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricDetails`: []HistoricDetailDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricDetailAPI.QueryHistoricDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | 
 **historicDetailQueryDto** | [**HistoricDetailQueryDto**](HistoricDetailQueryDto.md) |  | 

### Return type

[**[]HistoricDetailDto**](HistoricDetailDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

