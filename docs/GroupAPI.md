# \GroupAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvailableGroupInstanceOperations**](GroupAPI.md#AvailableGroupInstanceOperations) | **Options** /group/{id} | Group Resource Instance Options
[**AvailableGroupMembersOperations**](GroupAPI.md#AvailableGroupMembersOperations) | **Options** /group/{id}/members | Group Membership Resource Options
[**AvailableGroupOperations**](GroupAPI.md#AvailableGroupOperations) | **Options** /group | Group Resource Options
[**CreateGroup**](GroupAPI.md#CreateGroup) | **Post** /group/create | Create Group
[**CreateGroupMember**](GroupAPI.md#CreateGroupMember) | **Put** /group/{id}/members/{userId} | Create Group Member
[**DeleteGroup**](GroupAPI.md#DeleteGroup) | **Delete** /group/{id} | Delete Group
[**DeleteGroupMember**](GroupAPI.md#DeleteGroupMember) | **Delete** /group/{id}/members/{userId} | Delete a Group Member
[**GetGroup**](GroupAPI.md#GetGroup) | **Get** /group/{id} | Get Group
[**GetGroupCount**](GroupAPI.md#GetGroupCount) | **Get** /group/count | Get List Count
[**GetQueryGroups**](GroupAPI.md#GetQueryGroups) | **Get** /group | Get List
[**PostQueryGroups**](GroupAPI.md#PostQueryGroups) | **Post** /group | Get List (POST)
[**QueryGroupCount**](GroupAPI.md#QueryGroupCount) | **Post** /group/count | Get List Count (POST)
[**UpdateGroup**](GroupAPI.md#UpdateGroup) | **Put** /group/{id} | Update Group



## AvailableGroupInstanceOperations

> ResourceOptionsDto AvailableGroupInstanceOperations(ctx, id).Execute()

Group Resource Instance Options



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
	id := "id_example" // string | The id of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.AvailableGroupInstanceOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.AvailableGroupInstanceOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableGroupInstanceOperations`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.AvailableGroupInstanceOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableGroupInstanceOperationsRequest struct via the builder pattern


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


## AvailableGroupMembersOperations

> ResourceOptionsDto AvailableGroupMembersOperations(ctx, id).Execute()

Group Membership Resource Options



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
	id := "id_example" // string | The id of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.AvailableGroupMembersOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.AvailableGroupMembersOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableGroupMembersOperations`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.AvailableGroupMembersOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableGroupMembersOperationsRequest struct via the builder pattern


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


## AvailableGroupOperations

> ResourceOptionsDto AvailableGroupOperations(ctx).Execute()

Group Resource Options



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
	resp, r, err := apiClient.GroupAPI.AvailableGroupOperations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.AvailableGroupOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableGroupOperations`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.AvailableGroupOperations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableGroupOperationsRequest struct via the builder pattern


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


## CreateGroup

> CreateGroup(ctx).GroupDto(groupDto).Execute()

Create Group



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
	groupDto := *openapiclient.NewGroupDto() // GroupDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.CreateGroup(context.Background()).GroupDto(groupDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupDto** | [**GroupDto**](GroupDto.md) |  | 

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


## CreateGroupMember

> CreateGroupMember(ctx, id, userId).Execute()

Create Group Member



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
	id := "id_example" // string | The id of the group.
	userId := "userId_example" // string | The id of user to add to the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.CreateGroupMember(context.Background(), id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.CreateGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the group. | 
**userId** | **string** | The id of user to add to the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupMemberRequest struct via the builder pattern


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


## DeleteGroup

> DeleteGroup(ctx, id).Execute()

Delete Group



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
	id := "id_example" // string | The id of the group to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.DeleteGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the group to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## DeleteGroupMember

> DeleteGroupMember(ctx, id, userId).Execute()

Delete a Group Member



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
	id := "id_example" // string | The id of the group.
	userId := "userId_example" // string | The id of user to remove from the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.DeleteGroupMember(context.Background(), id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DeleteGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the group. | 
**userId** | **string** | The id of user to remove from the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupMemberRequest struct via the builder pattern


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


## GetGroup

> GroupDto GetGroup(ctx, id).Execute()

Get Group



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
	id := "id_example" // string | The id of the group to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GetGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroup`: GroupDto
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the group to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupDto**](GroupDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupCount

> CountResultDto GetGroupCount(ctx).Id(id).IdIn(idIn).Name(name).NameLike(nameLike).Type_(type_).Member(member).MemberOfTenant(memberOfTenant).Execute()

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
	id := "id_example" // string | Filter by the id of the group. (optional)
	idIn := "idIn_example" // string | Filter by a comma seperated list of group ids. (optional)
	name := "name_example" // string | Filter by the name of the group. (optional)
	nameLike := "nameLike_example" // string | Filter by the name that the parameter is a substring of. (optional)
	type_ := "type__example" // string | Filter by the type of the group. (optional)
	member := "member_example" // string | Only retrieve groups where the given user id is a member of. (optional)
	memberOfTenant := "memberOfTenant_example" // string | Only retrieve groups which are members of the given tenant. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GetGroupCount(context.Background()).Id(id).IdIn(idIn).Name(name).NameLike(nameLike).Type_(type_).Member(member).MemberOfTenant(memberOfTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetGroupCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetGroupCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter by the id of the group. | 
 **idIn** | **string** | Filter by a comma seperated list of group ids. | 
 **name** | **string** | Filter by the name of the group. | 
 **nameLike** | **string** | Filter by the name that the parameter is a substring of. | 
 **type_** | **string** | Filter by the type of the group. | 
 **member** | **string** | Only retrieve groups where the given user id is a member of. | 
 **memberOfTenant** | **string** | Only retrieve groups which are members of the given tenant. | 

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


## GetQueryGroups

> []GroupDto GetQueryGroups(ctx).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Id(id).IdIn(idIn).Name(name).NameLike(nameLike).Type_(type_).Member(member).MemberOfTenant(memberOfTenant).Execute()

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
	id := "id_example" // string | Filter by the id of the group. (optional)
	idIn := "idIn_example" // string | Filter by a comma seperated list of group ids. (optional)
	name := "name_example" // string | Filter by the name of the group. (optional)
	nameLike := "nameLike_example" // string | Filter by the name that the parameter is a substring of. (optional)
	type_ := "type__example" // string | Filter by the type of the group. (optional)
	member := "member_example" // string | Only retrieve groups where the given user id is a member of. (optional)
	memberOfTenant := "memberOfTenant_example" // string | Only retrieve groups which are members of the given tenant. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GetQueryGroups(context.Background()).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Id(id).IdIn(idIn).Name(name).NameLike(nameLike).Type_(type_).Member(member).MemberOfTenant(memberOfTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetQueryGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueryGroups`: []GroupDto
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetQueryGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **id** | **string** | Filter by the id of the group. | 
 **idIn** | **string** | Filter by a comma seperated list of group ids. | 
 **name** | **string** | Filter by the name of the group. | 
 **nameLike** | **string** | Filter by the name that the parameter is a substring of. | 
 **type_** | **string** | Filter by the type of the group. | 
 **member** | **string** | Only retrieve groups where the given user id is a member of. | 
 **memberOfTenant** | **string** | Only retrieve groups which are members of the given tenant. | 

### Return type

[**[]GroupDto**](GroupDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQueryGroups

> []GroupDto PostQueryGroups(ctx).FirstResult(firstResult).MaxResults(maxResults).GroupQueryDto(groupQueryDto).Execute()

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
	groupQueryDto := *openapiclient.NewGroupQueryDto() // GroupQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.PostQueryGroups(context.Background()).FirstResult(firstResult).MaxResults(maxResults).GroupQueryDto(groupQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.PostQueryGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQueryGroups`: []GroupDto
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.PostQueryGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostQueryGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **groupQueryDto** | [**GroupQueryDto**](GroupQueryDto.md) |  | 

### Return type

[**[]GroupDto**](GroupDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryGroupCount

> CountResultDto QueryGroupCount(ctx).GroupQueryDto(groupQueryDto).Execute()

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
	groupQueryDto := *openapiclient.NewGroupQueryDto() // GroupQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.QueryGroupCount(context.Background()).GroupQueryDto(groupQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.QueryGroupCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryGroupCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.QueryGroupCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryGroupCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupQueryDto** | [**GroupQueryDto**](GroupQueryDto.md) |  | 

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


## UpdateGroup

> UpdateGroup(ctx, id).GroupDto(groupDto).Execute()

Update Group



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
	id := "id_example" // string | The id of the group.
	groupDto := *openapiclient.NewGroupDto() // GroupDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.UpdateGroup(context.Background(), id).GroupDto(groupDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.UpdateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupDto** | [**GroupDto**](GroupDto.md) |  | 

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

