# \TenantAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvailableTenantGroupMembersOperations**](TenantAPI.md#AvailableTenantGroupMembersOperations) | **Options** /tenant/{id}/group-members | Tenant Group Membership Resource Options
[**AvailableTenantInstanceOperations**](TenantAPI.md#AvailableTenantInstanceOperations) | **Options** /tenant/{id} | Tenant Resource Options
[**AvailableTenantResourceOperations**](TenantAPI.md#AvailableTenantResourceOperations) | **Options** /tenant | Tenant Resource Options
[**AvailableTenantUserMembersOperations**](TenantAPI.md#AvailableTenantUserMembersOperations) | **Options** /tenant/{id}/user-members | Tenant User Membership Resource Options
[**CreateGroupMembership**](TenantAPI.md#CreateGroupMembership) | **Put** /tenant/{id}/group-members/{groupId} | Create Tenant Group Membership
[**CreateTenant**](TenantAPI.md#CreateTenant) | **Post** /tenant/create | Create Tenant
[**CreateUserMembership**](TenantAPI.md#CreateUserMembership) | **Put** /tenant/{id}/user-members/{userId} | Create Tenant User Membership
[**DeleteGroupMembership**](TenantAPI.md#DeleteGroupMembership) | **Delete** /tenant/{id}/group-members/{groupId} | Create Tenant Group Membership
[**DeleteTenant**](TenantAPI.md#DeleteTenant) | **Delete** /tenant/{id} | Delete Tenant
[**DeleteUserMembership**](TenantAPI.md#DeleteUserMembership) | **Delete** /tenant/{id}/user-members/{userId} | Delete a Tenant User Membership
[**GetTenant**](TenantAPI.md#GetTenant) | **Get** /tenant/{id} | Get Tenant
[**GetTenantCount**](TenantAPI.md#GetTenantCount) | **Get** /tenant/count | Get Tenant Count
[**QueryTenants**](TenantAPI.md#QueryTenants) | **Get** /tenant | Get Tenants
[**UpdateTenant**](TenantAPI.md#UpdateTenant) | **Put** /tenant/{id} | Update Tenant



## AvailableTenantGroupMembersOperations

> ResourceOptionsDto AvailableTenantGroupMembersOperations(ctx, id).Execute()

Tenant Group Membership Resource Options



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
	id := "id_example" // string | The id of the tenant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.AvailableTenantGroupMembersOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.AvailableTenantGroupMembersOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableTenantGroupMembersOperations`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.AvailableTenantGroupMembersOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the tenant | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableTenantGroupMembersOperationsRequest struct via the builder pattern


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


## AvailableTenantInstanceOperations

> ResourceOptionsDto AvailableTenantInstanceOperations(ctx, id).Execute()

Tenant Resource Options



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
	id := "id_example" // string | The id of the tenant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.AvailableTenantInstanceOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.AvailableTenantInstanceOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableTenantInstanceOperations`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.AvailableTenantInstanceOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the tenant | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableTenantInstanceOperationsRequest struct via the builder pattern


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


## AvailableTenantResourceOperations

> ResourceOptionsDto AvailableTenantResourceOperations(ctx).Execute()

Tenant Resource Options



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
	resp, r, err := apiClient.TenantAPI.AvailableTenantResourceOperations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.AvailableTenantResourceOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableTenantResourceOperations`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.AvailableTenantResourceOperations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableTenantResourceOperationsRequest struct via the builder pattern


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


## AvailableTenantUserMembersOperations

> ResourceOptionsDto AvailableTenantUserMembersOperations(ctx, id).Execute()

Tenant User Membership Resource Options



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
	id := "id_example" // string | The id of the tenant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.AvailableTenantUserMembersOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.AvailableTenantUserMembersOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableTenantUserMembersOperations`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.AvailableTenantUserMembersOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the tenant | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableTenantUserMembersOperationsRequest struct via the builder pattern


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


## CreateGroupMembership

> CreateGroupMembership(ctx, id, groupId).Execute()

Create Tenant Group Membership



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
	id := "id_example" // string | The id of the tenant.
	groupId := "groupId_example" // string | The id of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.CreateGroupMembership(context.Background(), id, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.CreateGroupMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the tenant. | 
**groupId** | **string** | The id of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupMembershipRequest struct via the builder pattern


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


## CreateTenant

> CreateTenant(ctx).TenantDto(tenantDto).Execute()

Create Tenant



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
	tenantDto := *openapiclient.NewTenantDto() // TenantDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.CreateTenant(context.Background()).TenantDto(tenantDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.CreateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantDto** | [**TenantDto**](TenantDto.md) |  | 

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


## CreateUserMembership

> CreateUserMembership(ctx, id, userId).Execute()

Create Tenant User Membership



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
	id := "id_example" // string | The id of the tenant.
	userId := "userId_example" // string | The id of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.CreateUserMembership(context.Background(), id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.CreateUserMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the tenant. | 
**userId** | **string** | The id of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserMembershipRequest struct via the builder pattern


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


## DeleteGroupMembership

> DeleteGroupMembership(ctx, id, groupId).Execute()

Create Tenant Group Membership



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
	id := "id_example" // string | The id of the tenant.
	groupId := "groupId_example" // string | The id of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.DeleteGroupMembership(context.Background(), id, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.DeleteGroupMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the tenant. | 
**groupId** | **string** | The id of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupMembershipRequest struct via the builder pattern


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


## DeleteTenant

> DeleteTenant(ctx, id).Execute()

Delete Tenant



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
	id := "id_example" // string | The id of the tenant to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.DeleteTenant(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.DeleteTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the tenant to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


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


## DeleteUserMembership

> DeleteUserMembership(ctx, id, userId).Execute()

Delete a Tenant User Membership



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
	id := "id_example" // string | The id of the tenant.
	userId := "userId_example" // string | The id of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.DeleteUserMembership(context.Background(), id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.DeleteUserMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the tenant. | 
**userId** | **string** | The id of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserMembershipRequest struct via the builder pattern


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


## GetTenant

> TenantDto GetTenant(ctx, id).Execute()

Get Tenant



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
	id := "id_example" // string | The id of the tenant to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.GetTenant(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenant`: TenantDto
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the tenant to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TenantDto**](TenantDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantCount

> CountResultDto GetTenantCount(ctx).Id(id).Name(name).NameLike(nameLike).UserMember(userMember).GroupMember(groupMember).IncludingGroupsOfUser(includingGroupsOfUser).Execute()

Get Tenant Count



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
	id := "id_example" // string | Filter by the id of the tenant. (optional)
	name := "name_example" // string | Filter by the name of the tenant. (optional)
	nameLike := "nameLike_example" // string | Filter by the name that the parameter is a substring of. (optional)
	userMember := "userMember_example" // string | Select only tenants where the given user is a member of. (optional)
	groupMember := "groupMember_example" // string | Select only tenants where the given group is a member of. (optional)
	includingGroupsOfUser := true // bool | Select only tenants where the user or one of his groups is a member of. Can only be used in combination with the `userMember` parameter. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.GetTenantCount(context.Background()).Id(id).Name(name).NameLike(nameLike).UserMember(userMember).GroupMember(groupMember).IncludingGroupsOfUser(includingGroupsOfUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetTenantCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetTenantCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter by the id of the tenant. | 
 **name** | **string** | Filter by the name of the tenant. | 
 **nameLike** | **string** | Filter by the name that the parameter is a substring of. | 
 **userMember** | **string** | Select only tenants where the given user is a member of. | 
 **groupMember** | **string** | Select only tenants where the given group is a member of. | 
 **includingGroupsOfUser** | **bool** | Select only tenants where the user or one of his groups is a member of. Can only be used in combination with the &#x60;userMember&#x60; parameter. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

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


## QueryTenants

> []TenantDto QueryTenants(ctx).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Id(id).Name(name).NameLike(nameLike).UserMember(userMember).GroupMember(groupMember).IncludingGroupsOfUser(includingGroupsOfUser).Execute()

Get Tenants



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
	id := "id_example" // string | Filter by the id of the tenant. (optional)
	name := "name_example" // string | Filter by the name of the tenant. (optional)
	nameLike := "nameLike_example" // string | Filter by the name that the parameter is a substring of. (optional)
	userMember := "userMember_example" // string | Select only tenants where the given user is a member of. (optional)
	groupMember := "groupMember_example" // string | Select only tenants where the given group is a member of. (optional)
	includingGroupsOfUser := true // bool | Select only tenants where the user or one of his groups is a member of. Can only be used in combination with the `userMember` parameter. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.QueryTenants(context.Background()).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Id(id).Name(name).NameLike(nameLike).UserMember(userMember).GroupMember(groupMember).IncludingGroupsOfUser(includingGroupsOfUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.QueryTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryTenants`: []TenantDto
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.QueryTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **id** | **string** | Filter by the id of the tenant. | 
 **name** | **string** | Filter by the name of the tenant. | 
 **nameLike** | **string** | Filter by the name that the parameter is a substring of. | 
 **userMember** | **string** | Select only tenants where the given user is a member of. | 
 **groupMember** | **string** | Select only tenants where the given group is a member of. | 
 **includingGroupsOfUser** | **bool** | Select only tenants where the user or one of his groups is a member of. Can only be used in combination with the &#x60;userMember&#x60; parameter. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

### Return type

[**[]TenantDto**](TenantDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenant

> UpdateTenant(ctx, id).TenantDto(tenantDto).Execute()

Update Tenant



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
	id := "id_example" // string | The id of the tenant.
	tenantDto := *openapiclient.NewTenantDto() // TenantDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.UpdateTenant(context.Background(), id).TenantDto(tenantDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.UpdateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantDto** | [**TenantDto**](TenantDto.md) |  | 

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

