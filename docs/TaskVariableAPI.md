# \TaskVariableAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaskVariable**](TaskVariableAPI.md#DeleteTaskVariable) | **Delete** /task/{id}/variables/{varName} | Delete Task Variable
[**GetTaskVariable**](TaskVariableAPI.md#GetTaskVariable) | **Get** /task/{id}/variables/{varName} | Get Task Variable
[**GetTaskVariableBinary**](TaskVariableAPI.md#GetTaskVariableBinary) | **Get** /task/{id}/variables/{varName}/data | Get Task Variable (Binary)
[**GetTaskVariables**](TaskVariableAPI.md#GetTaskVariables) | **Get** /task/{id}/variables | Get Task Variables
[**ModifyTaskVariables**](TaskVariableAPI.md#ModifyTaskVariables) | **Post** /task/{id}/variables | Update/Delete Task Variables
[**PutTaskVariable**](TaskVariableAPI.md#PutTaskVariable) | **Put** /task/{id}/variables/{varName} | Update Task Variable
[**SetBinaryTaskVariable**](TaskVariableAPI.md#SetBinaryTaskVariable) | **Post** /task/{id}/variables/{varName}/data | Update Task Variable (Binary)



## DeleteTaskVariable

> DeleteTaskVariable(ctx, id, varName).Execute()

Delete Task Variable



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
	id := "id_example" // string | The id of the task to delete the variable from.
	varName := "varName_example" // string | The name of the variable to be removed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskVariableAPI.DeleteTaskVariable(context.Background(), id, varName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskVariableAPI.DeleteTaskVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to delete the variable from. | 
**varName** | **string** | The name of the variable to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskVariableRequest struct via the builder pattern


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


## GetTaskVariable

> VariableValueDto GetTaskVariable(ctx, id, varName).DeserializeValue(deserializeValue).Execute()

Get Task Variable



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
	id := "id_example" // string | The id of the task to retrieve the variable from.
	varName := "varName_example" // string | The name of the variable to get.
	deserializeValue := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on the server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  Note: While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskVariableAPI.GetTaskVariable(context.Background(), id, varName).DeserializeValue(deserializeValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskVariableAPI.GetTaskVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskVariable`: VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `TaskVariableAPI.GetTaskVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the variable from. | 
**varName** | **string** | The name of the variable to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deserializeValue** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on the server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  Note: While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | [default to true]

### Return type

[**VariableValueDto**](VariableValueDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskVariableBinary

> *os.File GetTaskVariableBinary(ctx, id, varName).Execute()

Get Task Variable (Binary)



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
	id := "id_example" // string | The id of the task to retrieve the variable for.
	varName := "varName_example" // string | The name of the variable to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskVariableAPI.GetTaskVariableBinary(context.Background(), id, varName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskVariableAPI.GetTaskVariableBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskVariableBinary`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TaskVariableAPI.GetTaskVariableBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the variable for. | 
**varName** | **string** | The name of the variable to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskVariableBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskVariables

> map[string]VariableValueDto GetTaskVariables(ctx, id).DeserializeValues(deserializeValues).Execute()

Get Task Variables



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
	id := "id_example" // string | The id of the task to retrieve the variables from.
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on the server side (default `true`). If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  Note: While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskVariableAPI.GetTaskVariables(context.Background(), id).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskVariableAPI.GetTaskVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskVariables`: map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `TaskVariableAPI.GetTaskVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the variables from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on the server side (default &#x60;true&#x60;). If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  Note: While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | [default to true]

### Return type

[**map[string]VariableValueDto**](VariableValueDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTaskVariables

> ModifyTaskVariables(ctx, id).PatchVariablesDto(patchVariablesDto).Execute()

Update/Delete Task Variables



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
	id := "id_example" // string | The id of the task to set variables for.
	patchVariablesDto := *openapiclient.NewPatchVariablesDto() // PatchVariablesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskVariableAPI.ModifyTaskVariables(context.Background(), id).PatchVariablesDto(patchVariablesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskVariableAPI.ModifyTaskVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to set variables for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTaskVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchVariablesDto** | [**PatchVariablesDto**](PatchVariablesDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTaskVariable

> PutTaskVariable(ctx, id, varName).VariableValueDto(variableValueDto).Execute()

Update Task Variable



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
	id := "id_example" // string | The id of the task to set the variable for.
	varName := "varName_example" // string | The name of the variable to set.
	variableValueDto := *openapiclient.NewVariableValueDto() // VariableValueDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskVariableAPI.PutTaskVariable(context.Background(), id, varName).VariableValueDto(variableValueDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskVariableAPI.PutTaskVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to set the variable for. | 
**varName** | **string** | The name of the variable to set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTaskVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **variableValueDto** | [**VariableValueDto**](VariableValueDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBinaryTaskVariable

> SetBinaryTaskVariable(ctx, id, varName).Data(data).ValueType(valueType).Execute()

Update Task Variable (Binary)



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
	id := "id_example" // string | The id of the task to retrieve the variable for.
	varName := "varName_example" // string | The name of the variable to retrieve.
	data := os.NewFile(1234, "some_file") // *os.File | The binary data to be set. For File variables, this multipart can contain the filename, binary value and MIME type of the file variable to be set Only the filename is mandatory. (optional)
	valueType := "valueType_example" // string | The name of the variable type. Either Bytes for a byte array variable or File for a file variable. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskVariableAPI.SetBinaryTaskVariable(context.Background(), id, varName).Data(data).ValueType(valueType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskVariableAPI.SetBinaryTaskVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the variable for. | 
**varName** | **string** | The name of the variable to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetBinaryTaskVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | ***os.File** | The binary data to be set. For File variables, this multipart can contain the filename, binary value and MIME type of the file variable to be set Only the filename is mandatory. | 
 **valueType** | **string** | The name of the variable type. Either Bytes for a byte array variable or File for a file variable. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

