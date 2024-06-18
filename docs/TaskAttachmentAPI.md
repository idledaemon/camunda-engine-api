# \TaskAttachmentAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAttachment**](TaskAttachmentAPI.md#AddAttachment) | **Post** /task/{id}/attachment/create | Create
[**DeleteAttachment**](TaskAttachmentAPI.md#DeleteAttachment) | **Delete** /task/{id}/attachment/{attachmentId} | Delete
[**GetAttachment**](TaskAttachmentAPI.md#GetAttachment) | **Get** /task/{id}/attachment/{attachmentId} | Get
[**GetAttachmentData**](TaskAttachmentAPI.md#GetAttachmentData) | **Get** /task/{id}/attachment/{attachmentId}/data | Get (Binary)
[**GetAttachments**](TaskAttachmentAPI.md#GetAttachments) | **Get** /task/{id}/attachment | Get List



## AddAttachment

> AttachmentDto AddAttachment(ctx, id).AttachmentName(attachmentName).AttachmentDescription(attachmentDescription).AttachmentType(attachmentType).Url(url).Content(content).Execute()

Create



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
	id := "id_example" // string | The id of the task to add the attachment to.
	attachmentName := "attachmentName_example" // string | The name of the attachment. (optional)
	attachmentDescription := "attachmentDescription_example" // string | The description of the attachment. (optional)
	attachmentType := "attachmentType_example" // string | The type of the attachment. (optional)
	url := "url_example" // string | The url to the remote content of the attachment. (optional)
	content := os.NewFile(1234, "some_file") // *os.File | The content of the attachment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAttachmentAPI.AddAttachment(context.Background(), id).AttachmentName(attachmentName).AttachmentDescription(attachmentDescription).AttachmentType(attachmentType).Url(url).Content(content).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAttachmentAPI.AddAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAttachment`: AttachmentDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAttachmentAPI.AddAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to add the attachment to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachmentName** | **string** | The name of the attachment. | 
 **attachmentDescription** | **string** | The description of the attachment. | 
 **attachmentType** | **string** | The type of the attachment. | 
 **url** | **string** | The url to the remote content of the attachment. | 
 **content** | ***os.File** | The content of the attachment. | 

### Return type

[**AttachmentDto**](AttachmentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttachment

> DeleteAttachment(ctx, id, attachmentId).Execute()

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
	id := "id_example" // string | The id of the task.
	attachmentId := "attachmentId_example" // string | The id of the attachment to be removed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskAttachmentAPI.DeleteAttachment(context.Background(), id, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAttachmentAPI.DeleteAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task. | 
**attachmentId** | **string** | The id of the attachment to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentRequest struct via the builder pattern


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


## GetAttachment

> AttachmentDto GetAttachment(ctx, id, attachmentId).Execute()

Get



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
	id := "id_example" // string | The id of the task.
	attachmentId := "attachmentId_example" // string | The id of the attachment to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAttachmentAPI.GetAttachment(context.Background(), id, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAttachmentAPI.GetAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachment`: AttachmentDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAttachmentAPI.GetAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task. | 
**attachmentId** | **string** | The id of the attachment to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttachmentDto**](AttachmentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachmentData

> *os.File GetAttachmentData(ctx, id, attachmentId).Execute()

Get (Binary)



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
	id := "id_example" // string | The id of the task.
	attachmentId := "attachmentId_example" // string | The id of the attachment to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAttachmentAPI.GetAttachmentData(context.Background(), id, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAttachmentAPI.GetAttachmentData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentData`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TaskAttachmentAPI.GetAttachmentData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task. | 
**attachmentId** | **string** | The id of the attachment to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentDataRequest struct via the builder pattern


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


## GetAttachments

> []AttachmentDto GetAttachments(ctx, id).Execute()

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
	id := "id_example" // string | The id of the task to retrieve the attachments for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAttachmentAPI.GetAttachments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAttachmentAPI.GetAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachments`: []AttachmentDto
	fmt.Fprintf(os.Stdout, "Response from `TaskAttachmentAPI.GetAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the attachments for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AttachmentDto**](AttachmentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

