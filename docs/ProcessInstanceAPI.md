# \ProcessInstanceAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorrelateMessageAsyncOperation**](ProcessInstanceAPI.md#CorrelateMessageAsyncOperation) | **Post** /process-instance/message-async | Correlate Message Async (POST)
[**DeleteAsyncHistoricQueryBased**](ProcessInstanceAPI.md#DeleteAsyncHistoricQueryBased) | **Post** /process-instance/delete-historic-query-based | Delete Async Historic Query Based (POST)
[**DeleteProcessInstance**](ProcessInstanceAPI.md#DeleteProcessInstance) | **Delete** /process-instance/{id} | Delete
[**DeleteProcessInstanceVariable**](ProcessInstanceAPI.md#DeleteProcessInstanceVariable) | **Delete** /process-instance/{id}/variables/{varName} | Delete Process Variable
[**DeleteProcessInstancesAsyncOperation**](ProcessInstanceAPI.md#DeleteProcessInstancesAsyncOperation) | **Post** /process-instance/delete | Delete Async (POST)
[**GetActivityInstanceTree**](ProcessInstanceAPI.md#GetActivityInstanceTree) | **Get** /process-instance/{id}/activity-instances | Get Activity Instance
[**GetProcessInstance**](ProcessInstanceAPI.md#GetProcessInstance) | **Get** /process-instance/{id} | Get Process Instance
[**GetProcessInstanceComments**](ProcessInstanceAPI.md#GetProcessInstanceComments) | **Get** /process-instance/{id}/comment | Get Process Instance Comments
[**GetProcessInstanceVariable**](ProcessInstanceAPI.md#GetProcessInstanceVariable) | **Get** /process-instance/{id}/variables/{varName} | Get Process Variable
[**GetProcessInstanceVariableBinary**](ProcessInstanceAPI.md#GetProcessInstanceVariableBinary) | **Get** /process-instance/{id}/variables/{varName}/data | Get Process Variable (Binary)
[**GetProcessInstanceVariables**](ProcessInstanceAPI.md#GetProcessInstanceVariables) | **Get** /process-instance/{id}/variables | Get Process Variables
[**GetProcessInstances**](ProcessInstanceAPI.md#GetProcessInstances) | **Get** /process-instance | Get List
[**GetProcessInstancesCount**](ProcessInstanceAPI.md#GetProcessInstancesCount) | **Get** /process-instance/count | Get List Count
[**ModifyProcessInstance**](ProcessInstanceAPI.md#ModifyProcessInstance) | **Post** /process-instance/{id}/modification | Modify Process Instance Execution State
[**ModifyProcessInstanceAsyncOperation**](ProcessInstanceAPI.md#ModifyProcessInstanceAsyncOperation) | **Post** /process-instance/{id}/modification-async | Modify Process Instance Execution State Async
[**ModifyProcessInstanceVariables**](ProcessInstanceAPI.md#ModifyProcessInstanceVariables) | **Post** /process-instance/{id}/variables | Update/Delete Process Variables
[**QueryProcessInstances**](ProcessInstanceAPI.md#QueryProcessInstances) | **Post** /process-instance | Get List (POST)
[**QueryProcessInstancesCount**](ProcessInstanceAPI.md#QueryProcessInstancesCount) | **Post** /process-instance/count | Get List Count (POST)
[**SetProcessInstanceVariable**](ProcessInstanceAPI.md#SetProcessInstanceVariable) | **Put** /process-instance/{id}/variables/{varName} | Update Process Variable
[**SetProcessInstanceVariableBinary**](ProcessInstanceAPI.md#SetProcessInstanceVariableBinary) | **Post** /process-instance/{id}/variables/{varName}/data | Update Process Variable (Binary)
[**SetRetriesByProcess**](ProcessInstanceAPI.md#SetRetriesByProcess) | **Post** /process-instance/job-retries | Set Job Retries Async (POST)
[**SetRetriesByProcessHistoricQueryBased**](ProcessInstanceAPI.md#SetRetriesByProcessHistoricQueryBased) | **Post** /process-instance/job-retries-historic-query-based | Set Job Retries Async Historic Query Based (POST)
[**SetVariablesAsyncOperation**](ProcessInstanceAPI.md#SetVariablesAsyncOperation) | **Post** /process-instance/variables-async | Set Variables Async (POST)
[**UpdateSuspensionState**](ProcessInstanceAPI.md#UpdateSuspensionState) | **Put** /process-instance/suspended | Activate/Suspend In Group
[**UpdateSuspensionStateAsyncOperation**](ProcessInstanceAPI.md#UpdateSuspensionStateAsyncOperation) | **Post** /process-instance/suspended-async | Activate/Suspend In Batch
[**UpdateSuspensionStateById**](ProcessInstanceAPI.md#UpdateSuspensionStateById) | **Put** /process-instance/{id}/suspended | Activate/Suspend Process Instance By Id



## CorrelateMessageAsyncOperation

> BatchDto CorrelateMessageAsyncOperation(ctx).CorrelationMessageAsyncDto(correlationMessageAsyncDto).Execute()

Correlate Message Async (POST)



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
	correlationMessageAsyncDto := *openapiclient.NewCorrelationMessageAsyncDto() // CorrelationMessageAsyncDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.CorrelateMessageAsyncOperation(context.Background()).CorrelationMessageAsyncDto(correlationMessageAsyncDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.CorrelateMessageAsyncOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CorrelateMessageAsyncOperation`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.CorrelateMessageAsyncOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorrelateMessageAsyncOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **correlationMessageAsyncDto** | [**CorrelationMessageAsyncDto**](CorrelationMessageAsyncDto.md) |  | 

### Return type

[**BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAsyncHistoricQueryBased

> BatchDto DeleteAsyncHistoricQueryBased(ctx).DeleteProcessInstancesDto(deleteProcessInstancesDto).Execute()

Delete Async Historic Query Based (POST)



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
	deleteProcessInstancesDto := *openapiclient.NewDeleteProcessInstancesDto() // DeleteProcessInstancesDto | **Unallowed property**: `processInstanceQuery` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.DeleteAsyncHistoricQueryBased(context.Background()).DeleteProcessInstancesDto(deleteProcessInstancesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.DeleteAsyncHistoricQueryBased``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAsyncHistoricQueryBased`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.DeleteAsyncHistoricQueryBased`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAsyncHistoricQueryBasedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteProcessInstancesDto** | [**DeleteProcessInstancesDto**](DeleteProcessInstancesDto.md) | **Unallowed property**: &#x60;processInstanceQuery&#x60; | 

### Return type

[**BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProcessInstance

> DeleteProcessInstance(ctx, id).SkipCustomListeners(skipCustomListeners).SkipIoMappings(skipIoMappings).SkipSubprocesses(skipSubprocesses).FailIfNotExists(failIfNotExists).Execute()

Delete



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
	id := "id_example" // string | The id of the process instance to be deleted.
	skipCustomListeners := true // bool | If set to true, the custom listeners will be skipped. (optional) (default to false)
	skipIoMappings := true // bool | If set to true, the input/output mappings will be skipped. (optional) (default to false)
	skipSubprocesses := true // bool | If set to true, subprocesses related to deleted processes will be skipped. (optional) (default to false)
	failIfNotExists := true // bool | If set to false, the request will still be successful if the process id is not found. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessInstanceAPI.DeleteProcessInstance(context.Background(), id).SkipCustomListeners(skipCustomListeners).SkipIoMappings(skipIoMappings).SkipSubprocesses(skipSubprocesses).FailIfNotExists(failIfNotExists).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.DeleteProcessInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCustomListeners** | **bool** | If set to true, the custom listeners will be skipped. | [default to false]
 **skipIoMappings** | **bool** | If set to true, the input/output mappings will be skipped. | [default to false]
 **skipSubprocesses** | **bool** | If set to true, subprocesses related to deleted processes will be skipped. | [default to false]
 **failIfNotExists** | **bool** | If set to false, the request will still be successful if the process id is not found. | [default to true]

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


## DeleteProcessInstanceVariable

> DeleteProcessInstanceVariable(ctx, id, varName).Execute()

Delete Process Variable



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
	id := "id_example" // string | The id of the process instance to delete the variable from.
	varName := "varName_example" // string | The name of the variable to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessInstanceAPI.DeleteProcessInstanceVariable(context.Background(), id, varName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.DeleteProcessInstanceVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to delete the variable from. | 
**varName** | **string** | The name of the variable to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcessInstanceVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProcessInstancesAsyncOperation

> BatchDto DeleteProcessInstancesAsyncOperation(ctx).DeleteProcessInstancesDto(deleteProcessInstancesDto).Execute()

Delete Async (POST)



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
	deleteProcessInstancesDto := *openapiclient.NewDeleteProcessInstancesDto() // DeleteProcessInstancesDto | **Unallowed property**: `historicProcessInstanceQuery` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.DeleteProcessInstancesAsyncOperation(context.Background()).DeleteProcessInstancesDto(deleteProcessInstancesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.DeleteProcessInstancesAsyncOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProcessInstancesAsyncOperation`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.DeleteProcessInstancesAsyncOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcessInstancesAsyncOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteProcessInstancesDto** | [**DeleteProcessInstancesDto**](DeleteProcessInstancesDto.md) | **Unallowed property**: &#x60;historicProcessInstanceQuery&#x60; | 

### Return type

[**BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityInstanceTree

> ActivityInstanceDto GetActivityInstanceTree(ctx, id).Execute()

Get Activity Instance



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
	id := "id_example" // string | The id of the process instance for which the activity instance should be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.GetActivityInstanceTree(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.GetActivityInstanceTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityInstanceTree`: ActivityInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.GetActivityInstanceTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance for which the activity instance should be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityInstanceTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivityInstanceDto**](ActivityInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessInstance

> ProcessInstanceDto GetProcessInstance(ctx, id).Execute()

Get Process Instance



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
	id := "id_example" // string | The id of the process instance to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.GetProcessInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.GetProcessInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessInstance`: ProcessInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.GetProcessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessInstanceDto**](ProcessInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessInstanceComments

> []CommentDto GetProcessInstanceComments(ctx, id).Execute()

Get Process Instance Comments



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
	id := "id_example" // string | The id of the process instance to retrieve the comments for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.GetProcessInstanceComments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.GetProcessInstanceComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessInstanceComments`: []CommentDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.GetProcessInstanceComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to retrieve the comments for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessInstanceCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommentDto**](CommentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessInstanceVariable

> VariableValueDto GetProcessInstanceVariable(ctx, id, varName).DeserializeValue(deserializeValue).Execute()

Get Process Variable



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
	id := "id_example" // string | The id of the process instance to retrieve the variable for.
	varName := "varName_example" // string | The name of the variable to retrieve.
	deserializeValue := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  Note: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.GetProcessInstanceVariable(context.Background(), id, varName).DeserializeValue(deserializeValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.GetProcessInstanceVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessInstanceVariable`: VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.GetProcessInstanceVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to retrieve the variable for. | 
**varName** | **string** | The name of the variable to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessInstanceVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deserializeValue** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  Note: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. | [default to true]

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


## GetProcessInstanceVariableBinary

> *os.File GetProcessInstanceVariableBinary(ctx, id, varName).Execute()

Get Process Variable (Binary)



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
	id := "id_example" // string | The id of the process instance to retrieve the variable for.
	varName := "varName_example" // string | The name of the variable to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.GetProcessInstanceVariableBinary(context.Background(), id, varName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.GetProcessInstanceVariableBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessInstanceVariableBinary`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.GetProcessInstanceVariableBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to retrieve the variable for. | 
**varName** | **string** | The name of the variable to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessInstanceVariableBinaryRequest struct via the builder pattern


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


## GetProcessInstanceVariables

> map[string]VariableValueDto GetProcessInstanceVariables(ctx, id).DeserializeValues(deserializeValues).Execute()

Get Process Variables



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
	id := "id_example" // string | The id of the process instance to retrieve the variables from.
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  Note: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.GetProcessInstanceVariables(context.Background(), id).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.GetProcessInstanceVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessInstanceVariables`: map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.GetProcessInstanceVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to retrieve the variables from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessInstanceVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  Note: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. | [default to true]

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


## GetProcessInstances

> []ProcessInstanceDto GetProcessInstances(ctx).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).ProcessInstanceIds(processInstanceIds).BusinessKey(businessKey).BusinessKeyLike(businessKeyLike).CaseInstanceId(caseInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionKeyNotIn(processDefinitionKeyNotIn).DeploymentId(deploymentId).SuperProcessInstance(superProcessInstance).SubProcessInstance(subProcessInstance).SuperCaseInstance(superCaseInstance).SubCaseInstance(subCaseInstance).Active(active).Suspended(suspended).WithIncident(withIncident).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).ProcessDefinitionWithoutTenantId(processDefinitionWithoutTenantId).ActivityIdIn(activityIdIn).RootProcessInstances(rootProcessInstances).LeafProcessInstances(leafProcessInstances).Variables(variables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).Execute()

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
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)
	processInstanceIds := "processInstanceIds_example" // string | Filter by a comma-separated list of process instance ids. (optional)
	businessKey := "businessKey_example" // string | Filter by process instance business key. (optional)
	businessKeyLike := "businessKeyLike_example" // string | Filter by process instance business key that the parameter is a substring of. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by case instance id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the deployment the id belongs to. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by the key of the process definition the instances run on. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Filter by a comma-separated list of process definition keys. A process instance must have one of the given process definition keys. (optional)
	processDefinitionKeyNotIn := "processDefinitionKeyNotIn_example" // string | Exclude instances by a comma-separated list of process definition keys. A process instance must not have one of the given process definition keys. (optional)
	deploymentId := "deploymentId_example" // string | Filter by the deployment the id belongs to. (optional)
	superProcessInstance := "superProcessInstance_example" // string | Restrict query to all process instances that are sub process instances of the given process instance. Takes a process instance id. (optional)
	subProcessInstance := "subProcessInstance_example" // string | Restrict query to all process instances that have the given process instance as a sub process instance. Takes a process instance id. (optional)
	superCaseInstance := "superCaseInstance_example" // string | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. (optional)
	subCaseInstance := "subCaseInstance_example" // string | Restrict query to all process instances that have the given case instance as a sub case instance. Takes a case instance id. (optional)
	active := true // bool | Only include active process instances. Value may only be true, as false is the default behavior. (optional) (default to false)
	suspended := true // bool | Only include suspended process instances. Value may only be true, as false is the default behavior. (optional) (default to false)
	withIncident := true // bool | Filter by presence of incidents. Selects only process instances that have an incident. (optional) (default to false)
	incidentId := "incidentId_example" // string | Filter by the incident id. (optional)
	incidentType := "incidentType_example" // string | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentMessage := "incidentMessage_example" // string | Filter by the incident message. Exact match. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Filter by the incident message that the parameter is a substring of. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A process instance must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include process instances which belong to no tenant. (optional) (default to false)
	processDefinitionWithoutTenantId := true // bool | Only include process instances which process definition has no tenant id. (optional) (default to false)
	activityIdIn := "activityIdIn_example" // string | Filter by a comma-separated list of activity ids. A process instance must currently wait in a leaf activity with one of the given activity ids. (optional)
	rootProcessInstances := true // bool | Restrict the query to all process instances that are top level process instances. (optional) (default to false)
	leafProcessInstances := true // bool | Restrict the query to all process instances that are leaf instances. (i.e. don't have any sub instances). (optional) (default to false)
	variables := "variables_example" // string | Only include process instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note**: Values are always treated as String objects on server side.  Valid `operator` values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`. `key` and `value` may not contain underscore or comma characters. (optional)
	variableNamesIgnoreCase := true // bool | Match all variable names in this query case-insensitively. If set to true variableName and variablename are treated as equal. (optional) (default to false)
	variableValuesIgnoreCase := true // bool | Match all variable values in this query case-insensitively. If set to true variableValue and variablevalue are treated as equal. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.GetProcessInstances(context.Background()).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).ProcessInstanceIds(processInstanceIds).BusinessKey(businessKey).BusinessKeyLike(businessKeyLike).CaseInstanceId(caseInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionKeyNotIn(processDefinitionKeyNotIn).DeploymentId(deploymentId).SuperProcessInstance(superProcessInstance).SubProcessInstance(subProcessInstance).SuperCaseInstance(superCaseInstance).SubCaseInstance(subCaseInstance).Active(active).Suspended(suspended).WithIncident(withIncident).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).ProcessDefinitionWithoutTenantId(processDefinitionWithoutTenantId).ActivityIdIn(activityIdIn).RootProcessInstances(rootProcessInstances).LeafProcessInstances(leafProcessInstances).Variables(variables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.GetProcessInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessInstances`: []ProcessInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.GetProcessInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **processInstanceIds** | **string** | Filter by a comma-separated list of process instance ids. | 
 **businessKey** | **string** | Filter by process instance business key. | 
 **businessKeyLike** | **string** | Filter by process instance business key that the parameter is a substring of. | 
 **caseInstanceId** | **string** | Filter by case instance id. | 
 **processDefinitionId** | **string** | Filter by the deployment the id belongs to. | 
 **processDefinitionKey** | **string** | Filter by the key of the process definition the instances run on. | 
 **processDefinitionKeyIn** | **string** | Filter by a comma-separated list of process definition keys. A process instance must have one of the given process definition keys. | 
 **processDefinitionKeyNotIn** | **string** | Exclude instances by a comma-separated list of process definition keys. A process instance must not have one of the given process definition keys. | 
 **deploymentId** | **string** | Filter by the deployment the id belongs to. | 
 **superProcessInstance** | **string** | Restrict query to all process instances that are sub process instances of the given process instance. Takes a process instance id. | 
 **subProcessInstance** | **string** | Restrict query to all process instances that have the given process instance as a sub process instance. Takes a process instance id. | 
 **superCaseInstance** | **string** | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. | 
 **subCaseInstance** | **string** | Restrict query to all process instances that have the given case instance as a sub case instance. Takes a case instance id. | 
 **active** | **bool** | Only include active process instances. Value may only be true, as false is the default behavior. | [default to false]
 **suspended** | **bool** | Only include suspended process instances. Value may only be true, as false is the default behavior. | [default to false]
 **withIncident** | **bool** | Filter by presence of incidents. Selects only process instances that have an incident. | [default to false]
 **incidentId** | **string** | Filter by the incident id. | 
 **incidentType** | **string** | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentMessage** | **string** | Filter by the incident message. Exact match. | 
 **incidentMessageLike** | **string** | Filter by the incident message that the parameter is a substring of. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A process instance must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include process instances which belong to no tenant. | [default to false]
 **processDefinitionWithoutTenantId** | **bool** | Only include process instances which process definition has no tenant id. | [default to false]
 **activityIdIn** | **string** | Filter by a comma-separated list of activity ids. A process instance must currently wait in a leaf activity with one of the given activity ids. | 
 **rootProcessInstances** | **bool** | Restrict the query to all process instances that are top level process instances. | [default to false]
 **leafProcessInstances** | **bool** | Restrict the query to all process instances that are leaf instances. (i.e. don&#39;t have any sub instances). | [default to false]
 **variables** | **string** | Only include process instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note**: Values are always treated as String objects on server side.  Valid &#x60;operator&#x60; values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **variableNamesIgnoreCase** | **bool** | Match all variable names in this query case-insensitively. If set to true variableName and variablename are treated as equal. | [default to false]
 **variableValuesIgnoreCase** | **bool** | Match all variable values in this query case-insensitively. If set to true variableValue and variablevalue are treated as equal. | [default to false]

### Return type

[**[]ProcessInstanceDto**](ProcessInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessInstancesCount

> CountResultDto GetProcessInstancesCount(ctx).ProcessInstanceIds(processInstanceIds).BusinessKey(businessKey).BusinessKeyLike(businessKeyLike).CaseInstanceId(caseInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionKeyNotIn(processDefinitionKeyNotIn).DeploymentId(deploymentId).SuperProcessInstance(superProcessInstance).SubProcessInstance(subProcessInstance).SuperCaseInstance(superCaseInstance).SubCaseInstance(subCaseInstance).Active(active).Suspended(suspended).WithIncident(withIncident).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).ProcessDefinitionWithoutTenantId(processDefinitionWithoutTenantId).ActivityIdIn(activityIdIn).RootProcessInstances(rootProcessInstances).LeafProcessInstances(leafProcessInstances).Variables(variables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).Execute()

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
	processInstanceIds := "processInstanceIds_example" // string | Filter by a comma-separated list of process instance ids. (optional)
	businessKey := "businessKey_example" // string | Filter by process instance business key. (optional)
	businessKeyLike := "businessKeyLike_example" // string | Filter by process instance business key that the parameter is a substring of. (optional)
	caseInstanceId := "caseInstanceId_example" // string | Filter by case instance id. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the deployment the id belongs to. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by the key of the process definition the instances run on. (optional)
	processDefinitionKeyIn := "processDefinitionKeyIn_example" // string | Filter by a comma-separated list of process definition keys. A process instance must have one of the given process definition keys. (optional)
	processDefinitionKeyNotIn := "processDefinitionKeyNotIn_example" // string | Exclude instances by a comma-separated list of process definition keys. A process instance must not have one of the given process definition keys. (optional)
	deploymentId := "deploymentId_example" // string | Filter by the deployment the id belongs to. (optional)
	superProcessInstance := "superProcessInstance_example" // string | Restrict query to all process instances that are sub process instances of the given process instance. Takes a process instance id. (optional)
	subProcessInstance := "subProcessInstance_example" // string | Restrict query to all process instances that have the given process instance as a sub process instance. Takes a process instance id. (optional)
	superCaseInstance := "superCaseInstance_example" // string | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. (optional)
	subCaseInstance := "subCaseInstance_example" // string | Restrict query to all process instances that have the given case instance as a sub case instance. Takes a case instance id. (optional)
	active := true // bool | Only include active process instances. Value may only be true, as false is the default behavior. (optional) (default to false)
	suspended := true // bool | Only include suspended process instances. Value may only be true, as false is the default behavior. (optional) (default to false)
	withIncident := true // bool | Filter by presence of incidents. Selects only process instances that have an incident. (optional) (default to false)
	incidentId := "incidentId_example" // string | Filter by the incident id. (optional)
	incidentType := "incidentType_example" // string | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentMessage := "incidentMessage_example" // string | Filter by the incident message. Exact match. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Filter by the incident message that the parameter is a substring of. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A process instance must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include process instances which belong to no tenant. (optional) (default to false)
	processDefinitionWithoutTenantId := true // bool | Only include process instances which process definition has no tenant id. (optional) (default to false)
	activityIdIn := "activityIdIn_example" // string | Filter by a comma-separated list of activity ids. A process instance must currently wait in a leaf activity with one of the given activity ids. (optional)
	rootProcessInstances := true // bool | Restrict the query to all process instances that are top level process instances. (optional) (default to false)
	leafProcessInstances := true // bool | Restrict the query to all process instances that are leaf instances. (i.e. don't have any sub instances). (optional) (default to false)
	variables := "variables_example" // string | Only include process instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value.  **Note**: Values are always treated as String objects on server side.  Valid `operator` values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`. `key` and `value` may not contain underscore or comma characters. (optional)
	variableNamesIgnoreCase := true // bool | Match all variable names in this query case-insensitively. If set to true variableName and variablename are treated as equal. (optional) (default to false)
	variableValuesIgnoreCase := true // bool | Match all variable values in this query case-insensitively. If set to true variableValue and variablevalue are treated as equal. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.GetProcessInstancesCount(context.Background()).ProcessInstanceIds(processInstanceIds).BusinessKey(businessKey).BusinessKeyLike(businessKeyLike).CaseInstanceId(caseInstanceId).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessDefinitionKeyIn(processDefinitionKeyIn).ProcessDefinitionKeyNotIn(processDefinitionKeyNotIn).DeploymentId(deploymentId).SuperProcessInstance(superProcessInstance).SubProcessInstance(subProcessInstance).SuperCaseInstance(superCaseInstance).SubCaseInstance(subCaseInstance).Active(active).Suspended(suspended).WithIncident(withIncident).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).ProcessDefinitionWithoutTenantId(processDefinitionWithoutTenantId).ActivityIdIn(activityIdIn).RootProcessInstances(rootProcessInstances).LeafProcessInstances(leafProcessInstances).Variables(variables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.GetProcessInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.GetProcessInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processInstanceIds** | **string** | Filter by a comma-separated list of process instance ids. | 
 **businessKey** | **string** | Filter by process instance business key. | 
 **businessKeyLike** | **string** | Filter by process instance business key that the parameter is a substring of. | 
 **caseInstanceId** | **string** | Filter by case instance id. | 
 **processDefinitionId** | **string** | Filter by the deployment the id belongs to. | 
 **processDefinitionKey** | **string** | Filter by the key of the process definition the instances run on. | 
 **processDefinitionKeyIn** | **string** | Filter by a comma-separated list of process definition keys. A process instance must have one of the given process definition keys. | 
 **processDefinitionKeyNotIn** | **string** | Exclude instances by a comma-separated list of process definition keys. A process instance must not have one of the given process definition keys. | 
 **deploymentId** | **string** | Filter by the deployment the id belongs to. | 
 **superProcessInstance** | **string** | Restrict query to all process instances that are sub process instances of the given process instance. Takes a process instance id. | 
 **subProcessInstance** | **string** | Restrict query to all process instances that have the given process instance as a sub process instance. Takes a process instance id. | 
 **superCaseInstance** | **string** | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. | 
 **subCaseInstance** | **string** | Restrict query to all process instances that have the given case instance as a sub case instance. Takes a case instance id. | 
 **active** | **bool** | Only include active process instances. Value may only be true, as false is the default behavior. | [default to false]
 **suspended** | **bool** | Only include suspended process instances. Value may only be true, as false is the default behavior. | [default to false]
 **withIncident** | **bool** | Filter by presence of incidents. Selects only process instances that have an incident. | [default to false]
 **incidentId** | **string** | Filter by the incident id. | 
 **incidentType** | **string** | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentMessage** | **string** | Filter by the incident message. Exact match. | 
 **incidentMessageLike** | **string** | Filter by the incident message that the parameter is a substring of. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A process instance must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include process instances which belong to no tenant. | [default to false]
 **processDefinitionWithoutTenantId** | **bool** | Only include process instances which process definition has no tenant id. | [default to false]
 **activityIdIn** | **string** | Filter by a comma-separated list of activity ids. A process instance must currently wait in a leaf activity with one of the given activity ids. | 
 **rootProcessInstances** | **bool** | Restrict the query to all process instances that are top level process instances. | [default to false]
 **leafProcessInstances** | **bool** | Restrict the query to all process instances that are leaf instances. (i.e. don&#39;t have any sub instances). | [default to false]
 **variables** | **string** | Only include process instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value.  **Note**: Values are always treated as String objects on server side.  Valid &#x60;operator&#x60; values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **variableNamesIgnoreCase** | **bool** | Match all variable names in this query case-insensitively. If set to true variableName and variablename are treated as equal. | [default to false]
 **variableValuesIgnoreCase** | **bool** | Match all variable values in this query case-insensitively. If set to true variableValue and variablevalue are treated as equal. | [default to false]

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


## ModifyProcessInstance

> ModifyProcessInstance(ctx, id).ProcessInstanceModificationDto(processInstanceModificationDto).Execute()

Modify Process Instance Execution State



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
	id := "id_example" // string | The id of the process instance to modify.
	processInstanceModificationDto := *openapiclient.NewProcessInstanceModificationDto() // ProcessInstanceModificationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessInstanceAPI.ModifyProcessInstance(context.Background(), id).ProcessInstanceModificationDto(processInstanceModificationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.ModifyProcessInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to modify. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyProcessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **processInstanceModificationDto** | [**ProcessInstanceModificationDto**](ProcessInstanceModificationDto.md) |  | 

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


## ModifyProcessInstanceAsyncOperation

> BatchDto ModifyProcessInstanceAsyncOperation(ctx, id).ProcessInstanceModificationDto(processInstanceModificationDto).Execute()

Modify Process Instance Execution State Async



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
	id := "id_example" // string | The id of the process instance to modify.
	processInstanceModificationDto := *openapiclient.NewProcessInstanceModificationDto() // ProcessInstanceModificationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.ModifyProcessInstanceAsyncOperation(context.Background(), id).ProcessInstanceModificationDto(processInstanceModificationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.ModifyProcessInstanceAsyncOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyProcessInstanceAsyncOperation`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.ModifyProcessInstanceAsyncOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to modify. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyProcessInstanceAsyncOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **processInstanceModificationDto** | [**ProcessInstanceModificationDto**](ProcessInstanceModificationDto.md) |  | 

### Return type

[**BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyProcessInstanceVariables

> ModifyProcessInstanceVariables(ctx, id).PatchVariablesDto(patchVariablesDto).Execute()

Update/Delete Process Variables



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
	id := "id_example" // string | The id of the process instance to set variables for.
	patchVariablesDto := *openapiclient.NewPatchVariablesDto() // PatchVariablesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessInstanceAPI.ModifyProcessInstanceVariables(context.Background(), id).PatchVariablesDto(patchVariablesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.ModifyProcessInstanceVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to set variables for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyProcessInstanceVariablesRequest struct via the builder pattern


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


## QueryProcessInstances

> []ProcessInstanceDto QueryProcessInstances(ctx).FirstResult(firstResult).MaxResults(maxResults).ProcessInstanceQueryDto(processInstanceQueryDto).Execute()

Get List (POST)



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
	processInstanceQueryDto := *openapiclient.NewProcessInstanceQueryDto() // ProcessInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.QueryProcessInstances(context.Background()).FirstResult(firstResult).MaxResults(maxResults).ProcessInstanceQueryDto(processInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.QueryProcessInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryProcessInstances`: []ProcessInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.QueryProcessInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryProcessInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **processInstanceQueryDto** | [**ProcessInstanceQueryDto**](ProcessInstanceQueryDto.md) |  | 

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


## QueryProcessInstancesCount

> CountResultDto QueryProcessInstancesCount(ctx).ProcessInstanceQueryDto(processInstanceQueryDto).Execute()

Get List Count (POST)



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
	processInstanceQueryDto := *openapiclient.NewProcessInstanceQueryDto() // ProcessInstanceQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.QueryProcessInstancesCount(context.Background()).ProcessInstanceQueryDto(processInstanceQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.QueryProcessInstancesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryProcessInstancesCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.QueryProcessInstancesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryProcessInstancesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processInstanceQueryDto** | [**ProcessInstanceQueryDto**](ProcessInstanceQueryDto.md) |  | 

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


## SetProcessInstanceVariable

> SetProcessInstanceVariable(ctx, id, varName).VariableValueDto(variableValueDto).Execute()

Update Process Variable



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
	id := "id_example" // string | The id of the process instance to set the variable for.
	varName := "varName_example" // string | The name of the variable to set.
	variableValueDto := *openapiclient.NewVariableValueDto() // VariableValueDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessInstanceAPI.SetProcessInstanceVariable(context.Background(), id, varName).VariableValueDto(variableValueDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.SetProcessInstanceVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to set the variable for. | 
**varName** | **string** | The name of the variable to set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProcessInstanceVariableRequest struct via the builder pattern


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


## SetProcessInstanceVariableBinary

> SetProcessInstanceVariableBinary(ctx, id, varName).Data(data).ValueType(valueType).Execute()

Update Process Variable (Binary)



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
	id := "id_example" // string | The id of the process instance to retrieve the variable for.
	varName := "varName_example" // string | The name of the variable to retrieve.
	data := os.NewFile(1234, "some_file") // *os.File | The binary data to be set. For File variables, this multipart can contain the filename, binary value and MIME type of the file variable to be set Only the filename is mandatory. (optional)
	valueType := "valueType_example" // string | The name of the variable type. Either Bytes for a byte array variable or File for a file variable. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessInstanceAPI.SetProcessInstanceVariableBinary(context.Background(), id, varName).Data(data).ValueType(valueType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.SetProcessInstanceVariableBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to retrieve the variable for. | 
**varName** | **string** | The name of the variable to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProcessInstanceVariableBinaryRequest struct via the builder pattern


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


## SetRetriesByProcess

> BatchDto SetRetriesByProcess(ctx).SetJobRetriesByProcessDto(setJobRetriesByProcessDto).Execute()

Set Job Retries Async (POST)



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
	setJobRetriesByProcessDto := *openapiclient.NewSetJobRetriesByProcessDto() // SetJobRetriesByProcessDto | Please note that if both processInstances and processInstanceQuery are provided, then the resulting execution will be performed on the union of these sets. **Unallowed property**: `historicProcessInstanceQuery` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.SetRetriesByProcess(context.Background()).SetJobRetriesByProcessDto(setJobRetriesByProcessDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.SetRetriesByProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRetriesByProcess`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.SetRetriesByProcess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRetriesByProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setJobRetriesByProcessDto** | [**SetJobRetriesByProcessDto**](SetJobRetriesByProcessDto.md) | Please note that if both processInstances and processInstanceQuery are provided, then the resulting execution will be performed on the union of these sets. **Unallowed property**: &#x60;historicProcessInstanceQuery&#x60; | 

### Return type

[**BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRetriesByProcessHistoricQueryBased

> BatchDto SetRetriesByProcessHistoricQueryBased(ctx).SetJobRetriesByProcessDto(setJobRetriesByProcessDto).Execute()

Set Job Retries Async Historic Query Based (POST)



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
	setJobRetriesByProcessDto := *openapiclient.NewSetJobRetriesByProcessDto() // SetJobRetriesByProcessDto | Please note that if both processInstances and historicProcessInstanceQuery are provided, then the resulting execution will be performed on the union of these sets. **Unallowed property**: `processInstanceQuery` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.SetRetriesByProcessHistoricQueryBased(context.Background()).SetJobRetriesByProcessDto(setJobRetriesByProcessDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.SetRetriesByProcessHistoricQueryBased``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRetriesByProcessHistoricQueryBased`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.SetRetriesByProcessHistoricQueryBased`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRetriesByProcessHistoricQueryBasedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setJobRetriesByProcessDto** | [**SetJobRetriesByProcessDto**](SetJobRetriesByProcessDto.md) | Please note that if both processInstances and historicProcessInstanceQuery are provided, then the resulting execution will be performed on the union of these sets. **Unallowed property**: &#x60;processInstanceQuery&#x60; | 

### Return type

[**BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVariablesAsyncOperation

> BatchDto SetVariablesAsyncOperation(ctx).SetVariablesAsyncDto(setVariablesAsyncDto).Execute()

Set Variables Async (POST)



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
	setVariablesAsyncDto := *openapiclient.NewSetVariablesAsyncDto() // SetVariablesAsyncDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.SetVariablesAsyncOperation(context.Background()).SetVariablesAsyncDto(setVariablesAsyncDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.SetVariablesAsyncOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVariablesAsyncOperation`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.SetVariablesAsyncOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetVariablesAsyncOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setVariablesAsyncDto** | [**SetVariablesAsyncDto**](SetVariablesAsyncDto.md) |  | 

### Return type

[**BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSuspensionState

> UpdateSuspensionState(ctx).ProcessInstanceSuspensionStateDto(processInstanceSuspensionStateDto).Execute()

Activate/Suspend In Group



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
	processInstanceSuspensionStateDto := *openapiclient.NewProcessInstanceSuspensionStateDto() // ProcessInstanceSuspensionStateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessInstanceAPI.UpdateSuspensionState(context.Background()).ProcessInstanceSuspensionStateDto(processInstanceSuspensionStateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.UpdateSuspensionState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSuspensionStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processInstanceSuspensionStateDto** | [**ProcessInstanceSuspensionStateDto**](ProcessInstanceSuspensionStateDto.md) |  | 

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


## UpdateSuspensionStateAsyncOperation

> BatchDto UpdateSuspensionStateAsyncOperation(ctx).ProcessInstanceSuspensionStateAsyncDto(processInstanceSuspensionStateAsyncDto).Execute()

Activate/Suspend In Batch



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
	processInstanceSuspensionStateAsyncDto := *openapiclient.NewProcessInstanceSuspensionStateAsyncDto() // ProcessInstanceSuspensionStateAsyncDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessInstanceAPI.UpdateSuspensionStateAsyncOperation(context.Background()).ProcessInstanceSuspensionStateAsyncDto(processInstanceSuspensionStateAsyncDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.UpdateSuspensionStateAsyncOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSuspensionStateAsyncOperation`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessInstanceAPI.UpdateSuspensionStateAsyncOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSuspensionStateAsyncOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processInstanceSuspensionStateAsyncDto** | [**ProcessInstanceSuspensionStateAsyncDto**](ProcessInstanceSuspensionStateAsyncDto.md) |  | 

### Return type

[**BatchDto**](BatchDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSuspensionStateById

> UpdateSuspensionStateById(ctx, id).SuspensionStateDto(suspensionStateDto).Execute()

Activate/Suspend Process Instance By Id



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
	id := "id_example" // string | The id of the process instance to activate or suspend.
	suspensionStateDto := *openapiclient.NewSuspensionStateDto() // SuspensionStateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessInstanceAPI.UpdateSuspensionStateById(context.Background(), id).SuspensionStateDto(suspensionStateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessInstanceAPI.UpdateSuspensionStateById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process instance to activate or suspend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSuspensionStateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suspensionStateDto** | [**SuspensionStateDto**](SuspensionStateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

