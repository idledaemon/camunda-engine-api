# \TaskIdentityLinkAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIdentityLink**](TaskIdentityLinkAPI.md#AddIdentityLink) | **Post** /task/{id}/identity-links | Add
[**DeleteIdentityLink**](TaskIdentityLinkAPI.md#DeleteIdentityLink) | **Post** /task/{id}/identity-links/delete | Delete
[**GetIdentityLinks**](TaskIdentityLinkAPI.md#GetIdentityLinks) | **Get** /task/{id}/identity-links | Get List



## AddIdentityLink

> AddIdentityLink(ctx, id).IdentityLinkDto(identityLinkDto).Execute()

Add



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
	id := "id_example" // string | The id of the task to add a link to.
	identityLinkDto := *openapiclient.NewIdentityLinkDto("Type_example") // IdentityLinkDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskIdentityLinkAPI.AddIdentityLink(context.Background(), id).IdentityLinkDto(identityLinkDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskIdentityLinkAPI.AddIdentityLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to add a link to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddIdentityLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityLinkDto** | [**IdentityLinkDto**](IdentityLinkDto.md) |  | 

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


## DeleteIdentityLink

> DeleteIdentityLink(ctx, id).IdentityLinkDto(identityLinkDto).Execute()

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
	id := "id_example" // string | The id of the task to remove a link from.
	identityLinkDto := *openapiclient.NewIdentityLinkDto("Type_example") // IdentityLinkDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskIdentityLinkAPI.DeleteIdentityLink(context.Background(), id).IdentityLinkDto(identityLinkDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskIdentityLinkAPI.DeleteIdentityLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to remove a link from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityLinkDto** | [**IdentityLinkDto**](IdentityLinkDto.md) |  | 

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


## GetIdentityLinks

> []IdentityLinkDto GetIdentityLinks(ctx, id).Type_(type_).Execute()

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
	id := "id_example" // string | The id of the task to retrieve the identity links for.
	type_ := "type__example" // string | Filter by the type of links to include. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskIdentityLinkAPI.GetIdentityLinks(context.Background(), id).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskIdentityLinkAPI.GetIdentityLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityLinks`: []IdentityLinkDto
	fmt.Fprintf(os.Stdout, "Response from `TaskIdentityLinkAPI.GetIdentityLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the task to retrieve the identity links for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Filter by the type of links to include. | 

### Return type

[**[]IdentityLinkDto**](IdentityLinkDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

