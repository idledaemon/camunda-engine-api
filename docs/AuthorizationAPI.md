# \AuthorizationAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvailableOperationsAuthorization**](AuthorizationAPI.md#AvailableOperationsAuthorization) | **Options** /authorization | Authorization Resource Options
[**AvailableOperationsAuthorizationInstance**](AuthorizationAPI.md#AvailableOperationsAuthorizationInstance) | **Options** /authorization/{id} | Authorization Resource Options
[**CreateAuthorization**](AuthorizationAPI.md#CreateAuthorization) | **Post** /authorization/create | Create a New Authorization
[**DeleteAuthorization**](AuthorizationAPI.md#DeleteAuthorization) | **Delete** /authorization/{id} | Delete Authorization
[**GetAuthorization**](AuthorizationAPI.md#GetAuthorization) | **Get** /authorization/{id} | Get Authorization
[**GetAuthorizationCount**](AuthorizationAPI.md#GetAuthorizationCount) | **Get** /authorization/count | Get Authorization Count
[**IsUserAuthorized**](AuthorizationAPI.md#IsUserAuthorized) | **Get** /authorization/check | Perform an Authorization Check
[**QueryAuthorizations**](AuthorizationAPI.md#QueryAuthorizations) | **Get** /authorization | Get Authorizations
[**UpdateAuthorization**](AuthorizationAPI.md#UpdateAuthorization) | **Put** /authorization/{id} | Update an Authorization



## AvailableOperationsAuthorization

> ResourceOptionsDto AvailableOperationsAuthorization(ctx).Execute()

Authorization Resource Options



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.AvailableOperationsAuthorization(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.AvailableOperationsAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableOperationsAuthorization`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.AvailableOperationsAuthorization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableOperationsAuthorizationRequest struct via the builder pattern


### Return type

[**ResourceOptionsDto**](ResourceOptionsDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AvailableOperationsAuthorizationInstance

> ResourceOptionsDto AvailableOperationsAuthorizationInstance(ctx, id).Execute()

Authorization Resource Options



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
	id := "id_example" // string | The id of the authorization to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.AvailableOperationsAuthorizationInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.AvailableOperationsAuthorizationInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableOperationsAuthorizationInstance`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.AvailableOperationsAuthorizationInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the authorization to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableOperationsAuthorizationInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceOptionsDto**](ResourceOptionsDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthorization

> AuthorizationDto CreateAuthorization(ctx).AuthorizationCreateDto(authorizationCreateDto).Execute()

Create a New Authorization



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
	authorizationCreateDto := *openapiclient.NewAuthorizationCreateDto() // AuthorizationCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.CreateAuthorization(context.Background()).AuthorizationCreateDto(authorizationCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.CreateAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAuthorization`: AuthorizationDto
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.CreateAuthorization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationCreateDto** | [**AuthorizationCreateDto**](AuthorizationCreateDto.md) |  | 

### Return type

[**AuthorizationDto**](AuthorizationDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthorization

> DeleteAuthorization(ctx, id).Execute()

Delete Authorization



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
	id := "id_example" // string | The id of the authorization to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthorizationAPI.DeleteAuthorization(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.DeleteAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the authorization to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationRequest struct via the builder pattern


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


## GetAuthorization

> AuthorizationDto GetAuthorization(ctx, id).Execute()

Get Authorization



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
	id := "id_example" // string | The id of the authorization to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.GetAuthorization(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.GetAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthorization`: AuthorizationDto
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.GetAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the authorization to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorizationDto**](AuthorizationDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationCount

> CountResultDto GetAuthorizationCount(ctx).Id(id).Type_(type_).UserIdIn(userIdIn).GroupIdIn(groupIdIn).ResourceType(resourceType).ResourceId(resourceId).Execute()

Get Authorization Count



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
	id := "id_example" // string | Filter by the id of the authorization. (optional)
	type_ := int32(56) // int32 | Filter by authorization type. (0=global, 1=grant, 2=revoke). See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#authorization-type) for more information about authorization types. (optional)
	userIdIn := "userIdIn_example" // string | Filter by a comma-separated list of userIds. (optional)
	groupIdIn := "groupIdIn_example" // string | Filter by a comma-separated list of groupIds. (optional)
	resourceType := int32(56) // int32 | Filter by an integer representation of the resource type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#resources) for a list of integer representations of resource types. (optional)
	resourceId := "resourceId_example" // string | Filter by resource id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.GetAuthorizationCount(context.Background()).Id(id).Type_(type_).UserIdIn(userIdIn).GroupIdIn(groupIdIn).ResourceType(resourceType).ResourceId(resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.GetAuthorizationCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthorizationCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.GetAuthorizationCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter by the id of the authorization. | 
 **type_** | **int32** | Filter by authorization type. (0&#x3D;global, 1&#x3D;grant, 2&#x3D;revoke). See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#authorization-type) for more information about authorization types. | 
 **userIdIn** | **string** | Filter by a comma-separated list of userIds. | 
 **groupIdIn** | **string** | Filter by a comma-separated list of groupIds. | 
 **resourceType** | **int32** | Filter by an integer representation of the resource type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#resources) for a list of integer representations of resource types. | 
 **resourceId** | **string** | Filter by resource id. | 

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


## IsUserAuthorized

> AuthorizationCheckResultDto IsUserAuthorized(ctx).PermissionName(permissionName).ResourceName(resourceName).ResourceType(resourceType).ResourceId(resourceId).UserId(userId).Execute()

Perform an Authorization Check



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
	permissionName := "permissionName_example" // string | String value representing the permission name to check for.
	resourceName := "resourceName_example" // string | String value for the name of the resource to check permissions for.
	resourceType := int32(56) // int32 | An integer representing the resource type to check permissions for. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#resources) for a list of integer representations of resource types.
	resourceId := "resourceId_example" // string | The id of the resource to check permissions for. If left blank, a check for global permissions on the resource is performed. (optional)
	userId := "userId_example" // string | The id of the user to check permissions for. The currently authenticated user must have a READ permission for the Authorization resource. If `userId` is blank, a check for the currently authenticated user is performed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.IsUserAuthorized(context.Background()).PermissionName(permissionName).ResourceName(resourceName).ResourceType(resourceType).ResourceId(resourceId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.IsUserAuthorized``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsUserAuthorized`: AuthorizationCheckResultDto
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.IsUserAuthorized`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIsUserAuthorizedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionName** | **string** | String value representing the permission name to check for. | 
 **resourceName** | **string** | String value for the name of the resource to check permissions for. | 
 **resourceType** | **int32** | An integer representing the resource type to check permissions for. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#resources) for a list of integer representations of resource types. | 
 **resourceId** | **string** | The id of the resource to check permissions for. If left blank, a check for global permissions on the resource is performed. | 
 **userId** | **string** | The id of the user to check permissions for. The currently authenticated user must have a READ permission for the Authorization resource. If &#x60;userId&#x60; is blank, a check for the currently authenticated user is performed. | 

### Return type

[**AuthorizationCheckResultDto**](AuthorizationCheckResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryAuthorizations

> []AuthorizationDto QueryAuthorizations(ctx).Id(id).Type_(type_).UserIdIn(userIdIn).GroupIdIn(groupIdIn).ResourceType(resourceType).ResourceId(resourceId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Authorizations



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
	id := "id_example" // string | Filter by the id of the authorization. (optional)
	type_ := int32(56) // int32 | Filter by authorization type. (0=global, 1=grant, 2=revoke). See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#authorization-type) for more information about authorization types. (optional)
	userIdIn := "userIdIn_example" // string | Filter by a comma-separated list of userIds. (optional)
	groupIdIn := "groupIdIn_example" // string | Filter by a comma-separated list of groupIds. (optional)
	resourceType := int32(56) // int32 | Filter by an integer representation of the resource type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#resources) for a list of integer representations of resource types. (optional)
	resourceId := "resourceId_example" // string | Filter by resource id. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.QueryAuthorizations(context.Background()).Id(id).Type_(type_).UserIdIn(userIdIn).GroupIdIn(groupIdIn).ResourceType(resourceType).ResourceId(resourceId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.QueryAuthorizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryAuthorizations`: []AuthorizationDto
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.QueryAuthorizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryAuthorizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter by the id of the authorization. | 
 **type_** | **int32** | Filter by authorization type. (0&#x3D;global, 1&#x3D;grant, 2&#x3D;revoke). See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#authorization-type) for more information about authorization types. | 
 **userIdIn** | **string** | Filter by a comma-separated list of userIds. | 
 **groupIdIn** | **string** | Filter by a comma-separated list of groupIds. | 
 **resourceType** | **int32** | Filter by an integer representation of the resource type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#resources) for a list of integer representations of resource types. | 
 **resourceId** | **string** | Filter by resource id. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]AuthorizationDto**](AuthorizationDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthorization

> UpdateAuthorization(ctx, id).AuthorizationUpdateDto(authorizationUpdateDto).Execute()

Update an Authorization



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
	id := "id_example" // string | The id of the authorization to be updated.
	authorizationUpdateDto := *openapiclient.NewAuthorizationUpdateDto() // AuthorizationUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthorizationAPI.UpdateAuthorization(context.Background(), id).AuthorizationUpdateDto(authorizationUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.UpdateAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the authorization to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationUpdateDto** | [**AuthorizationUpdateDto**](AuthorizationUpdateDto.md) |  | 

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

