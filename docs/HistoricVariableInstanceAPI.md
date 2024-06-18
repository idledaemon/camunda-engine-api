# \HistoricVariableInstanceAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHistoricVariableInstance**](HistoricVariableInstanceAPI.md#DeleteHistoricVariableInstance) | **Delete** /history/variable-instance/{id} | Delete Variable Instance
[**GetHistoricVariableInstance**](HistoricVariableInstanceAPI.md#GetHistoricVariableInstance) | **Get** /history/variable-instance/{id} | Get Variable Instance
[**GetHistoricVariableInstanceBinary**](HistoricVariableInstanceAPI.md#GetHistoricVariableInstanceBinary) | **Get** /history/variable-instance/{id}/data | Get Variable Instance (Binary)
[**GetHistoricVariableInstances**](HistoricVariableInstanceAPI.md#GetHistoricVariableInstances) | **Get** /history/variable-instance | Get Variable Instances
[**GetHistoricVariableInstancesCount**](HistoricVariableInstanceAPI.md#GetHistoricVariableInstancesCount) | **Get** /history/variable-instance/count | Get Variable Instance Count
[**QueryHistoricVariableInstances**](HistoricVariableInstanceAPI.md#QueryHistoricVariableInstances) | **Post** /history/variable-instance | Get Variable Instances (POST)
[**QueryHistoricVariableInstancesCount**](HistoricVariableInstanceAPI.md#QueryHistoricVariableInstancesCount) | **Post** /history/variable-instance/count | Get Variable Instance Count (POST)



## DeleteHistoricVariableInstance

> DeleteHistoricVariableInstance(ctx, id).Execute()

Delete Variable Instance



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
	r, err := apiClient.HistoricVariableInstanceAPI.DeleteHistoricVariableInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricVariableInstanceAPI.DeleteHistoricVariableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the variable instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHistoricVariableInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricVariableInstance

> HistoricVariableInstanceDto GetHistoricVariableInstance(ctx, id).DeserializeValues(deserializeValues).Execute()

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
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricVariableInstanceAPI.GetHistoricVariableInstance(context.Background(), id).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricVariableInstanceAPI.GetHistoricVariableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricVariableInstance`: HistoricVariableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricVariableInstanceAPI.GetHistoricVariableInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the variable instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricVariableInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | 

### Return type

[**HistoricVariableInstanceDto**](HistoricVariableInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricVariableInstanceBinary

> *os.File GetHistoricVariableInstanceBinary(ctx, id).Execute()

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
	resp, r, err := apiClient.HistoricVariableInstanceAPI.GetHistoricVariableInstanceBinary(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricVariableInstanceAPI.GetHistoricVariableInstanceBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricVariableInstanceBinary`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `HistoricVariableInstanceAPI.GetHistoricVariableInstanceBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the variable instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricVariableInstanceBinaryRequest struct via the builder pattern


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


## GetHistoricVariableInstances

> []HistoricVariableInstanceDto GetHistoricVariableInstances(ctx).VariableName(variableName).VariableNameLike(variableNameLike).VariableValue(variableValue).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).VariableTypeIn(variableTypeIn).IncludeDeleted(includeDeleted).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ExecutionIdIn(executionIdIn).CaseInstanceId(caseInstanceId).CaseExecutionIdIn(caseExecutionIdIn).CaseActivityIdIn(caseActivityIdIn).TaskIdIn(taskIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).VariableNameIn(variableNameIn).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).Execute()

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
	variableName := "variableName_example" // string | Filter by variable name. (optional)
	variableNameLike := "variableNameLike_example" // string | Restrict to variables with a name like the parameter. (optional)
	variableValue := map[string]interface{}{ ... } // map[string]interface{} | Filter by variable value. Is treated as a `String` object on server side. (optional)
	variableNamesIgnoreCase := true // bool | Match the variable name provided in `variableName` and `variableNameLike` case- insensitively. If set to `true` **variableName** and **variablename** are treated as equal. (optional)
	variableValuesIgnoreCase := true // bool | Match the variable value provided in `variableValue` case-insensitively. If set to `true` **variableValue** and **variablevalue** are treated as equal. (optional)
	variableTypeIn := "variableTypeIn_example" // string | Only include historic variable instances which belong to one of the passed and comma- separated variable types. A list of all supported variable types can be found [here](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#supported-variable-values). **Note:** All non-primitive variables are associated with the type 'serializable'. (optional)
	includeDeleted := true // bool | Include variables that has already been deleted during the execution. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by the process instance the variable belongs to. (optional)
	processInstanceIdIn := "processInstanceIdIn_example" // string | Only include historic variable instances which belong to one of the passed and comma-separated process instance ids. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the process definition the variable belongs to. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by a key of the process definition the variable belongs to. (optional)
	executionIdIn := "executionIdIn_example" // string | Only include historic variable instances which belong to one of the passed and and comma-separated execution ids. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by the case instance the variable belongs to. (optional)
	caseExecutionIdIn := "caseExecutionIdIn_example" // string | Only include historic variable instances which belong to one of the passed and and comma-separated case execution ids. (optional)
	caseActivityIdIn := "caseActivityIdIn_example" // string | Only include historic variable instances which belong to one of the passed and and comma-separated case activity ids. (optional)
	taskIdIn := "taskIdIn_example" // string | Only include historic variable instances which belong to one of the passed and and comma-separated task ids. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Only include historic variable instances which belong to one of the passed and and comma-separated activity instance ids. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include historic variable instances which belong to one of the passed and comma- separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic variable instances that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	variableNameIn := "variableNameIn_example" // string | Only include historic variable instances which belong to one of the passed and comma-separated variable names. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricVariableInstanceAPI.GetHistoricVariableInstances(context.Background()).VariableName(variableName).VariableNameLike(variableNameLike).VariableValue(variableValue).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).VariableTypeIn(variableTypeIn).IncludeDeleted(includeDeleted).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ExecutionIdIn(executionIdIn).CaseInstanceId(caseInstanceId).CaseExecutionIdIn(caseExecutionIdIn).CaseActivityIdIn(caseActivityIdIn).TaskIdIn(taskIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).VariableNameIn(variableNameIn).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricVariableInstanceAPI.GetHistoricVariableInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricVariableInstances`: []HistoricVariableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricVariableInstanceAPI.GetHistoricVariableInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricVariableInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableName** | **string** | Filter by variable name. | 
 **variableNameLike** | **string** | Restrict to variables with a name like the parameter. | 
 **variableValue** | [**map[string]interface{}**](map[string]interface{}.md) | Filter by variable value. Is treated as a &#x60;String&#x60; object on server side. | 
 **variableNamesIgnoreCase** | **bool** | Match the variable name provided in &#x60;variableName&#x60; and &#x60;variableNameLike&#x60; case- insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | 
 **variableValuesIgnoreCase** | **bool** | Match the variable value provided in &#x60;variableValue&#x60; case-insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | 
 **variableTypeIn** | **string** | Only include historic variable instances which belong to one of the passed and comma- separated variable types. A list of all supported variable types can be found [here](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#supported-variable-values). **Note:** All non-primitive variables are associated with the type &#39;serializable&#39;. | 
 **includeDeleted** | **bool** | Include variables that has already been deleted during the execution. | 
 **processInstanceId** | **string** | Filter by the process instance the variable belongs to. | 
 **processInstanceIdIn** | **string** | Only include historic variable instances which belong to one of the passed and comma-separated process instance ids. | 
 **processDefinitionId** | **string** | Filter by the process definition the variable belongs to. | 
 **processDefinitionKey** | **string** | Filter by a key of the process definition the variable belongs to. | 
 **executionIdIn** | **string** | Only include historic variable instances which belong to one of the passed and and comma-separated execution ids. | 
 **caseInstanceId** | **string** | Filter by the case instance the variable belongs to. | 
 **caseExecutionIdIn** | **string** | Only include historic variable instances which belong to one of the passed and and comma-separated case execution ids. | 
 **caseActivityIdIn** | **string** | Only include historic variable instances which belong to one of the passed and and comma-separated case activity ids. | 
 **taskIdIn** | **string** | Only include historic variable instances which belong to one of the passed and and comma-separated task ids. | 
 **activityInstanceIdIn** | **string** | Only include historic variable instances which belong to one of the passed and and comma-separated activity instance ids. | 
 **tenantIdIn** | **string** | Only include historic variable instances which belong to one of the passed and comma- separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic variable instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **variableNameIn** | **string** | Only include historic variable instances which belong to one of the passed and comma-separated variable names. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | 

### Return type

[**[]HistoricVariableInstanceDto**](HistoricVariableInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricVariableInstancesCount

> CountResultDto GetHistoricVariableInstancesCount(ctx).VariableName(variableName).VariableNameLike(variableNameLike).VariableValue(variableValue).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).VariableTypeIn(variableTypeIn).IncludeDeleted(includeDeleted).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ExecutionIdIn(executionIdIn).CaseInstanceId(caseInstanceId).CaseExecutionIdIn(caseExecutionIdIn).CaseActivityIdIn(caseActivityIdIn).TaskIdIn(taskIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).VariableNameIn(variableNameIn).Execute()

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
	variableName := "variableName_example" // string | Filter by variable name. (optional)
	variableNameLike := "variableNameLike_example" // string | Restrict to variables with a name like the parameter. (optional)
	variableValue := map[string]interface{}{ ... } // map[string]interface{} | Filter by variable value. Is treated as a `String` object on server side. (optional)
	variableNamesIgnoreCase := true // bool | Match the variable name provided in `variableName` and `variableNameLike` case- insensitively. If set to `true` **variableName** and **variablename** are treated as equal. (optional)
	variableValuesIgnoreCase := true // bool | Match the variable value provided in `variableValue` case-insensitively. If set to `true` **variableValue** and **variablevalue** are treated as equal. (optional)
	variableTypeIn := "variableTypeIn_example" // string | Only include historic variable instances which belong to one of the passed and comma- separated variable types. A list of all supported variable types can be found [here](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#supported-variable-values). **Note:** All non-primitive variables are associated with the type 'serializable'. (optional)
	includeDeleted := true // bool | Include variables that has already been deleted during the execution. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by the process instance the variable belongs to. (optional)
	processInstanceIdIn := "processInstanceIdIn_example" // string | Only include historic variable instances which belong to one of the passed and comma-separated process instance ids. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the process definition the variable belongs to. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by a key of the process definition the variable belongs to. (optional)
	executionIdIn := "executionIdIn_example" // string | Only include historic variable instances which belong to one of the passed and and comma-separated execution ids. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by the case instance the variable belongs to. (optional)
	caseExecutionIdIn := "caseExecutionIdIn_example" // string | Only include historic variable instances which belong to one of the passed and and comma-separated case execution ids. (optional)
	caseActivityIdIn := "caseActivityIdIn_example" // string | Only include historic variable instances which belong to one of the passed and and comma-separated case activity ids. (optional)
	taskIdIn := "taskIdIn_example" // string | Only include historic variable instances which belong to one of the passed and and comma-separated task ids. (optional)
	activityInstanceIdIn := "activityInstanceIdIn_example" // string | Only include historic variable instances which belong to one of the passed and and comma-separated activity instance ids. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include historic variable instances which belong to one of the passed and comma- separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include historic variable instances that belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	variableNameIn := "variableNameIn_example" // string | Only include historic variable instances which belong to one of the passed and comma-separated variable names. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricVariableInstanceAPI.GetHistoricVariableInstancesCount(context.Background()).VariableName(variableName).VariableNameLike(variableNameLike).VariableValue(variableValue).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).VariableTypeIn(variableTypeIn).IncludeDeleted(includeDeleted).ProcessInstanceId(processInstanceId).ProcessInstanceIdIn(processInstanceIdIn).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ExecutionIdIn(executionIdIn).CaseInstanceId(caseInstanceId).CaseExecutionIdIn(caseExecutionIdIn).CaseActivityIdIn(caseActivityIdIn).TaskIdIn(taskIdIn).ActivityInstanceIdIn(activityInstanceIdIn).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).VariableNameIn(variableNameIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricVariableInstanceAPI.GetHistoricVariableInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricVariableInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricVariableInstanceAPI.GetHistoricVariableInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricVariableInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableName** | **string** | Filter by variable name. | 
 **variableNameLike** | **string** | Restrict to variables with a name like the parameter. | 
 **variableValue** | [**map[string]interface{}**](map[string]interface{}.md) | Filter by variable value. Is treated as a &#x60;String&#x60; object on server side. | 
 **variableNamesIgnoreCase** | **bool** | Match the variable name provided in &#x60;variableName&#x60; and &#x60;variableNameLike&#x60; case- insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | 
 **variableValuesIgnoreCase** | **bool** | Match the variable value provided in &#x60;variableValue&#x60; case-insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | 
 **variableTypeIn** | **string** | Only include historic variable instances which belong to one of the passed and comma- separated variable types. A list of all supported variable types can be found [here](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#supported-variable-values). **Note:** All non-primitive variables are associated with the type &#39;serializable&#39;. | 
 **includeDeleted** | **bool** | Include variables that has already been deleted during the execution. | 
 **processInstanceId** | **string** | Filter by the process instance the variable belongs to. | 
 **processInstanceIdIn** | **string** | Only include historic variable instances which belong to one of the passed and comma-separated process instance ids. | 
 **processDefinitionId** | **string** | Filter by the process definition the variable belongs to. | 
 **processDefinitionKey** | **string** | Filter by a key of the process definition the variable belongs to. | 
 **executionIdIn** | **string** | Only include historic variable instances which belong to one of the passed and and comma-separated execution ids. | 
 **caseInstanceId** | **string** | Filter by the case instance the variable belongs to. | 
 **caseExecutionIdIn** | **string** | Only include historic variable instances which belong to one of the passed and and comma-separated case execution ids. | 
 **caseActivityIdIn** | **string** | Only include historic variable instances which belong to one of the passed and and comma-separated case activity ids. | 
 **taskIdIn** | **string** | Only include historic variable instances which belong to one of the passed and and comma-separated task ids. | 
 **activityInstanceIdIn** | **string** | Only include historic variable instances which belong to one of the passed and and comma-separated activity instance ids. | 
 **tenantIdIn** | **string** | Only include historic variable instances which belong to one of the passed and comma- separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include historic variable instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **variableNameIn** | **string** | Only include historic variable instances which belong to one of the passed and comma-separated variable names. | 

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


## QueryHistoricVariableInstances

> []HistoricVariableInstanceDto QueryHistoricVariableInstances(ctx).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).HistoricVariableInstanceQueryDto(historicVariableInstanceQueryDto).Execute()

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
	historicVariableInstanceQueryDto := *openapiclient.NewHistoricVariableInstanceQueryDto() // HistoricVariableInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricVariableInstanceAPI.QueryHistoricVariableInstances(context.Background()).FirstResult(firstResult).MaxResults(maxResults).DeserializeValues(deserializeValues).HistoricVariableInstanceQueryDto(historicVariableInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricVariableInstanceAPI.QueryHistoricVariableInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricVariableInstances`: []HistoricVariableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricVariableInstanceAPI.QueryHistoricVariableInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricVariableInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | 
 **historicVariableInstanceQueryDto** | [**HistoricVariableInstanceQueryDto**](HistoricVariableInstanceQueryDto.md) |  | 

### Return type

[**[]HistoricVariableInstanceDto**](HistoricVariableInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryHistoricVariableInstancesCount

> CountResultDto QueryHistoricVariableInstancesCount(ctx).HistoricVariableInstanceQueryDto(historicVariableInstanceQueryDto).Execute()

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
	historicVariableInstanceQueryDto := *openapiclient.NewHistoricVariableInstanceQueryDto() // HistoricVariableInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricVariableInstanceAPI.QueryHistoricVariableInstancesCount(context.Background()).HistoricVariableInstanceQueryDto(historicVariableInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricVariableInstanceAPI.QueryHistoricVariableInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryHistoricVariableInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `HistoricVariableInstanceAPI.QueryHistoricVariableInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryHistoricVariableInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **historicVariableInstanceQueryDto** | [**HistoricVariableInstanceQueryDto**](HistoricVariableInstanceQueryDto.md) |  | 

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

