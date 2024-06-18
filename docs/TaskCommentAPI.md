# \TaskCommentAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComment**](TaskCommentAPI.md#CreateComment) | **Post** /task/{id}/comment/create | Create
[**GetComment**](TaskCommentAPI.md#GetComment) | **Get** /task/{id}/comment/{commentId} | Get
[**GetComments**](TaskCommentAPI.md#GetComments) | **Get** /task/{id}/comment | Get List



## CreateComment

> CommentDto CreateComment(ctx, id).CommentDto(commentDto).Execute()

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
	id := "id_example" // string | The id of the task to add the comment to.
	commentDto := *openapiclient.NewCommentDto() // CommentDto | **Note:** Only the `message` and `processInstanceId` properties will be used. Every other property passed to this endpoint will be ignored. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCommentAPI.CreateComment(context.Background(), id).CommentDto(commentDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCommentAPI.CreateComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateComment`: CommentDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCommentAPI.CreateComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to add the comment to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commentDto** | [**CommentDto**](CommentDto.md) | **Note:** Only the &#x60;message&#x60; and &#x60;processInstanceId&#x60; properties will be used. Every other property passed to this endpoint will be ignored. | 

### Return type

[**CommentDto**](CommentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComment

> CommentDto GetComment(ctx, id, commentId).Execute()

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
	commentId := "commentId_example" // string | The id of the comment to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCommentAPI.GetComment(context.Background(), id, commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCommentAPI.GetComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComment`: CommentDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCommentAPI.GetComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task. | 
**commentId** | **string** | The id of the comment to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CommentDto**](CommentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComments

> []CommentDto GetComments(ctx, id).Execute()

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
	id := "id_example" // string | The id of the task to retrieve the comments for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCommentAPI.GetComments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCommentAPI.GetComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComments`: []CommentDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCommentAPI.GetComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the comments for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentsRequest struct via the builder pattern


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

