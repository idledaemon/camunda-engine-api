# \VariableInstanceAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVariableInstance**](VariableInstanceAPI.md#GetVariableInstance) | **Get** /variable-instance/{id} | Get Variable Instance
[**GetVariableInstanceBinary**](VariableInstanceAPI.md#GetVariableInstanceBinary) | **Get** /variable-instance/{id}/data | Get Variable Instance (Binary)
[**GetVariableInstances**](VariableInstanceAPI.md#GetVariableInstances) | **Get** /variable-instance | Get Variable Instances
[**GetVariableInstancesCount**](VariableInstanceAPI.md#GetVariableInstancesCount) | **Get** /variable-instance/count | Get Variable Instance Count
[**QueryVariableInstances**](VariableInstanceAPI.md#QueryVariableInstances) | **Post** /variable-instance | Get Variable Instances (POST)
[**QueryVariableInstancesCount**](VariableInstanceAPI.md#QueryVariableInstancesCount) | **Post** /variable-instance/count | Get Variable Instance Count (POST)



## GetVariableInstance

> VariableInstanceDto GetVariableInstance(ctx, id).DeserializeValue(deserializeValue).Execute()

Get Variable Instance



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
	id := "id_example" // string | The id of the variable instance.
	deserializeValue := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:**  While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariableInstanceAPI.GetVariableInstance(context.Background(), id).DeserializeValue(deserializeValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariableInstanceAPI.GetVariableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVariableInstance`: VariableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `VariableInstanceAPI.GetVariableInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the variable instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deserializeValue** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:**  While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | 

### Return type

[**VariableInstanceDto**](VariableInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariableInstanceBinary

> *os.File GetVariableInstanceBinary(ctx, id).Execute()

Get Variable Instance (Binary)



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
	id := "id_example" // string | The id of the variable instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariableInstanceAPI.GetVariableInstanceBinary(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariableInstanceAPI.GetVariableInstanceBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVariableInstanceBinary`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `VariableInstanceAPI.GetVariableInstanceBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the variable instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableInstanceBinaryRequest struct via the builder pattern


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


## GetVariableInstances

> []VariableInstanceDto GetVariableInstances(ctx).VariableName(variableName).VariableNameLike(variableNameLike).ProcessInstanceIdIn(processInstanceIdIn).ExecutionIdIn(executionIdIn).CaseInstanceIdIn(caseInstanceIdIn).CaseExecutionIdIn(caseExecutionIdIn).TaskIdIn(taskIdIn).BatchIdIn(batchIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).VariableValues(variableValues).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).VariableScopeIdIn(variableScopeIdIn).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).Execute()

Get Variable Instances



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
	variableName := "variableName_example" // string | Filter by variable instance name. (optional)
	variableNameLike := "variableNameLike_example" // string | Filter by the variable instance name. The parameter can include the wildcard `%` to express like-strategy such as: starts with (`%`name), ends with (name`%`) or contains (`%`name`%`). (optional)
	processInstanceIdIn := "processInstanceIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated process instance ids. (optional)
	executionIdIn := "executionIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated execution ids. (optional)
	caseInstanceIdIn := "caseInstanceIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated case instance ids. (optional)
	caseExecutionIdIn := "caseExecutionIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated case execution ids. (optional)
	taskIdIn := "taskIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated task ids. (optional)
	batchIdIn := "batchIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated batch ids. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated activity instance ids. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated tenant ids. (optional)
	variableValues := "variableValues_example" // string | Only include variable instances that have the certain values. Value filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note:** Values are always treated as `String` objects on server side.  Valid operator values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`. `key` and `value` may not contain underscore or comma characters. (optional)
	variableNamesIgnoreCase := true // bool | Match all variable names provided in `variableValues` case-insensitively. If set to `true` **variableName** and **variablename** are treated as equal. (optional)
	variableValuesIgnoreCase := true // bool | Match all variable values provided in `variableValues` case-insensitively. If set to `true` **variableValue** and **variablevalue** are treated as equal. (optional)
	variableScopeIdIn := "variableScopeIdIn_example" // string | Only include variable instances which belong to one of passed scope ids. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariableInstanceAPI.GetVariableInstances(context.Background()).VariableName(variableName).VariableNameLike(variableNameLike).ProcessInstanceIdIn(processInstanceIdIn).ExecutionIdIn(executionIdIn).CaseInstanceIdIn(caseInstanceIdIn).CaseExecutionIdIn(caseExecutionIdIn).TaskIdIn(taskIdIn).BatchIdIn(batchIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).VariableValues(variableValues).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).VariableScopeIdIn(variableScopeIdIn).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariableInstanceAPI.GetVariableInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVariableInstances`: []VariableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `VariableInstanceAPI.GetVariableInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableName** | **string** | Filter by variable instance name. | 
 **variableNameLike** | **string** | Filter by the variable instance name. The parameter can include the wildcard &#x60;%&#x60; to express like-strategy such as: starts with (&#x60;%&#x60;name), ends with (name&#x60;%&#x60;) or contains (&#x60;%&#x60;name&#x60;%&#x60;). | 
 **processInstanceIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated process instance ids. | 
 **executionIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated execution ids. | 
 **caseInstanceIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated case instance ids. | 
 **caseExecutionIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated case execution ids. | 
 **taskIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated task ids. | 
 **batchIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated batch ids. | 
 **activityInstanceIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated activity instance ids. | 
 **tenantIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated tenant ids. | 
 **variableValues** | **string** | Only include variable instances that have the certain values. Value filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note:** Values are always treated as &#x60;String&#x60; objects on server side.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **variableNamesIgnoreCase** | **bool** | Match all variable names provided in &#x60;variableValues&#x60; case-insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | 
 **variableValuesIgnoreCase** | **bool** | Match all variable values provided in &#x60;variableValues&#x60; case-insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | 
 **variableScopeIdIn** | **string** | Only include variable instances which belong to one of passed scope ids. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | 

### Return type

[**[]VariableInstanceDto**](VariableInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariableInstancesCount

> CountResultDto GetVariableInstancesCount(ctx).VariableName(variableName).VariableNameLike(variableNameLike).ProcessInstanceIdIn(processInstanceIdIn).ExecutionIdIn(executionIdIn).CaseInstanceIdIn(caseInstanceIdIn).CaseExecutionIdIn(caseExecutionIdIn).TaskIdIn(taskIdIn).BatchIdIn(batchIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).VariableValues(variableValues).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).VariableScopeIdIn(variableScopeIdIn).Execute()

Get Variable Instance Count



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
	variableName := "variableName_example" // string | Filter by variable instance name. (optional)
	variableNameLike := "variableNameLike_example" // string | Filter by the variable instance name. The parameter can include the wildcard `%` to express like-strategy such as: starts with (`%`name), ends with (name`%`) or contains (`%`name`%`). (optional)
	processInstanceIdIn := "processInstanceIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated process instance ids. (optional)
	executionIdIn := "executionIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated execution ids. (optional)
	caseInstanceIdIn := "caseInstanceIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated case instance ids. (optional)
	caseExecutionIdIn := "caseExecutionIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated case execution ids. (optional)
	taskIdIn := "taskIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated task ids. (optional)
	batchIdIn := "batchIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated batch ids. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated activity instance ids. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include variable instances which belong to one of the passed and comma-separated tenant ids. (optional)
	variableValues := "variableValues_example" // string | Only include variable instances that have the certain values. Value filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note:** Values are always treated as `String` objects on server side.  Valid operator values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`. `key` and `value` may not contain underscore or comma characters. (optional)
	variableNamesIgnoreCase := true // bool | Match all variable names provided in `variableValues` case-insensitively. If set to `true` **variableName** and **variablename** are treated as equal. (optional)
	variableValuesIgnoreCase := true // bool | Match all variable values provided in `variableValues` case-insensitively. If set to `true` **variableValue** and **variablevalue** are treated as equal. (optional)
	variableScopeIdIn := "variableScopeIdIn_example" // string | Only include variable instances which belong to one of passed scope ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariableInstanceAPI.GetVariableInstancesCount(context.Background()).VariableName(variableName).VariableNameLike(variableNameLike).ProcessInstanceIdIn(processInstanceIdIn).ExecutionIdIn(executionIdIn).CaseInstanceIdIn(caseInstanceIdIn).CaseExecutionIdIn(caseExecutionIdIn).TaskIdIn(taskIdIn).BatchIdIn(batchIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).VariableValues(variableValues).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).VariableScopeIdIn(variableScopeIdIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariableInstanceAPI.GetVariableInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVariableInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `VariableInstanceAPI.GetVariableInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableName** | **string** | Filter by variable instance name. | 
 **variableNameLike** | **string** | Filter by the variable instance name. The parameter can include the wildcard &#x60;%&#x60; to express like-strategy such as: starts with (&#x60;%&#x60;name), ends with (name&#x60;%&#x60;) or contains (&#x60;%&#x60;name&#x60;%&#x60;). | 
 **processInstanceIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated process instance ids. | 
 **executionIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated execution ids. | 
 **caseInstanceIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated case instance ids. | 
 **caseExecutionIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated case execution ids. | 
 **taskIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated task ids. | 
 **batchIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated batch ids. | 
 **activityInstanceIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated activity instance ids. | 
 **tenantIdIn** | **string** | Only include variable instances which belong to one of the passed and comma-separated tenant ids. | 
 **variableValues** | **string** | Only include variable instances that have the certain values. Value filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note:** Values are always treated as &#x60;String&#x60; objects on server side.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **variableNamesIgnoreCase** | **bool** | Match all variable names provided in &#x60;variableValues&#x60; case-insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | 
 **variableValuesIgnoreCase** | **bool** | Match all variable values provided in &#x60;variableValues&#x60; case-insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | 
 **variableScopeIdIn** | **string** | Only include variable instances which belong to one of passed scope ids. | 

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


## QueryVariableInstances

> []VariableInstanceDto QueryVariableInstances(ctx).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).VariableInstanceQueryDto(variableInstanceQueryDto).Execute()

Get Variable Instances (POST)



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
	variableInstanceQueryDto := *openapiclient.NewVariableInstanceQueryDto() // VariableInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariableInstanceAPI.QueryVariableInstances(context.Background()).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).VariableInstanceQueryDto(variableInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariableInstanceAPI.QueryVariableInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryVariableInstances`: []VariableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `VariableInstanceAPI.QueryVariableInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryVariableInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | 
 **variableInstanceQueryDto** | [**VariableInstanceQueryDto**](VariableInstanceQueryDto.md) |  | 

### Return type

[**[]VariableInstanceDto**](VariableInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryVariableInstancesCount

> CountResultDto QueryVariableInstancesCount(ctx).VariableInstanceQueryDto(variableInstanceQueryDto).Execute()

Get Variable Instance Count (POST)



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
	variableInstanceQueryDto := *openapiclient.NewVariableInstanceQueryDto() // VariableInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VariableInstanceAPI.QueryVariableInstancesCount(context.Background()).VariableInstanceQueryDto(variableInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariableInstanceAPI.QueryVariableInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryVariableInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `VariableInstanceAPI.QueryVariableInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryVariableInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableInstanceQueryDto** | [**VariableInstanceQueryDto**](VariableInstanceQueryDto.md) |  | 

### Return type

[**CountResultDto**](CountResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

