# \TaskLocalVariableAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaskLocalVariable**](TaskLocalVariableAPI.md#DeleteTaskLocalVariable) | **Delete** /task/{id}/localVariables/{varName} | Delete Local Task Variable
[**GetTaskLocalVariable**](TaskLocalVariableAPI.md#GetTaskLocalVariable) | **Get** /task/{id}/localVariables/{varName} | Get Local Task Variable
[**GetTaskLocalVariableBinary**](TaskLocalVariableAPI.md#GetTaskLocalVariableBinary) | **Get** /task/{id}/localVariables/{varName}/data | Get Local Task Variable (Binary)
[**GetTaskLocalVariables**](TaskLocalVariableAPI.md#GetTaskLocalVariables) | **Get** /task/{id}/localVariables | Get Local Task Variables
[**ModifyTaskLocalVariables**](TaskLocalVariableAPI.md#ModifyTaskLocalVariables) | **Post** /task/{id}/localVariables | Update/Delete Local Task Variables
[**PutTaskLocalVariable**](TaskLocalVariableAPI.md#PutTaskLocalVariable) | **Put** /task/{id}/localVariables/{varName} | Update Local Task Variable
[**SetBinaryTaskLocalVariable**](TaskLocalVariableAPI.md#SetBinaryTaskLocalVariable) | **Post** /task/{id}/localVariables/{varName}/data | Update Local Task Variable (Binary)



## DeleteTaskLocalVariable

> DeleteTaskLocalVariable(ctx, id, varName).Execute()

Delete Local Task Variable



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
	r, err := apiClient.TaskLocalVariableAPI.DeleteTaskLocalVariable(context.Background(), id, varName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLocalVariableAPI.DeleteTaskLocalVariable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTaskLocalVariableRequest struct via the builder pattern


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


## GetTaskLocalVariable

> VariableValueDto GetTaskLocalVariable(ctx, id, varName).DeserializeValue(deserializeValue).Execute()

Get Local Task Variable



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
	varName := "varName_example" // string | The name of the variable to get
	deserializeValue := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on the server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  Note: While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskLocalVariableAPI.GetTaskLocalVariable(context.Background(), id, varName).DeserializeValue(deserializeValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLocalVariableAPI.GetTaskLocalVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskLocalVariable`: VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `TaskLocalVariableAPI.GetTaskLocalVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the variable from. | 
**varName** | **string** | The name of the variable to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskLocalVariableRequest struct via the builder pattern


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


## GetTaskLocalVariableBinary

> *os.File GetTaskLocalVariableBinary(ctx, id, varName).Execute()

Get Local Task Variable (Binary)



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
	resp, r, err := apiClient.TaskLocalVariableAPI.GetTaskLocalVariableBinary(context.Background(), id, varName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLocalVariableAPI.GetTaskLocalVariableBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskLocalVariableBinary`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TaskLocalVariableAPI.GetTaskLocalVariableBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the variable for. | 
**varName** | **string** | The name of the variable to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskLocalVariableBinaryRequest struct via the builder pattern


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


## GetTaskLocalVariables

> map[string]VariableValueDto GetTaskLocalVariables(ctx, id).DeserializeValues(deserializeValues).Execute()

Get Local Task Variables



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
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on the server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskLocalVariableAPI.GetTaskLocalVariables(context.Background(), id).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLocalVariableAPI.GetTaskLocalVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskLocalVariables`: map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `TaskLocalVariableAPI.GetTaskLocalVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the variables from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskLocalVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on the server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | [default to true]

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


## ModifyTaskLocalVariables

> ModifyTaskLocalVariables(ctx, id).PatchVariablesDto(patchVariablesDto).Execute()

Update/Delete Local Task Variables



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
	r, err := apiClient.TaskLocalVariableAPI.ModifyTaskLocalVariables(context.Background(), id).PatchVariablesDto(patchVariablesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLocalVariableAPI.ModifyTaskLocalVariables``: %v\n", err)
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

Other parameters are passed through a pointer to a apiModifyTaskLocalVariablesRequest struct via the builder pattern


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


## PutTaskLocalVariable

> PutTaskLocalVariable(ctx, id, varName).VariableValueDto(variableValueDto).Execute()

Update Local Task Variable



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
	r, err := apiClient.TaskLocalVariableAPI.PutTaskLocalVariable(context.Background(), id, varName).VariableValueDto(variableValueDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLocalVariableAPI.PutTaskLocalVariable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutTaskLocalVariableRequest struct via the builder pattern


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


## SetBinaryTaskLocalVariable

> SetBinaryTaskLocalVariable(ctx, id, varName).Data(data).ValueType(valueType).Execute()

Update Local Task Variable (Binary)



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
	r, err := apiClient.TaskLocalVariableAPI.SetBinaryTaskLocalVariable(context.Background(), id, varName).Data(data).ValueType(valueType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLocalVariableAPI.SetBinaryTaskLocalVariable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetBinaryTaskLocalVariableRequest struct via the builder pattern


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

