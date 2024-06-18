# \UserAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvailableOperations**](UserAPI.md#AvailableOperations) | **Options** /user | Options
[**AvailableUserOperations**](UserAPI.md#AvailableUserOperations) | **Options** /user/{id} | Options
[**CreateUser**](UserAPI.md#CreateUser) | **Post** /user/create | Create
[**DeleteUser**](UserAPI.md#DeleteUser) | **Delete** /user/{id} | Delete
[**GetUserCount**](UserAPI.md#GetUserCount) | **Get** /user/count | Get List Count
[**GetUserProfile**](UserAPI.md#GetUserProfile) | **Get** /user/{id}/profile | Get Profile
[**GetUsers**](UserAPI.md#GetUsers) | **Get** /user | Get List
[**UnlockUser**](UserAPI.md#UnlockUser) | **Post** /user/{id}/unlock | Unlock User
[**UpdateCredentials**](UserAPI.md#UpdateCredentials) | **Put** /user/{id}/credentials | Update Credentials
[**UpdateProfile**](UserAPI.md#UpdateProfile) | **Put** /user/{id}/profile | Update User Profile



## AvailableOperations

> ResourceOptionsDto AvailableOperations(ctx).Execute()

Options



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
	resp, r, err := apiClient.UserAPI.AvailableOperations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.AvailableOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableOperations`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.AvailableOperations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableOperationsRequest struct via the builder pattern


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


## AvailableUserOperations

> ResourceOptionsDto AvailableUserOperations(ctx, id).Execute()

Options



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
	id := "id_example" // string | The id of the user to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.AvailableUserOperations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.AvailableUserOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableUserOperations`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.AvailableUserOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the user to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableUserOperationsRequest struct via the builder pattern


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


## CreateUser

> CreateUser(ctx).UserDto(userDto).Execute()

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
	userDto := *openapiclient.NewUserDto() // UserDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.CreateUser(context.Background()).UserDto(userDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userDto** | [**UserDto**](UserDto.md) |  | 

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


## DeleteUser

> DeleteUser(ctx, id).Execute()

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
	id := "id_example" // string | The id of the user to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the user to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## GetUserCount

> CountResultDto GetUserCount(ctx).Id(id).IdIn(idIn).FirstName(firstName).FirstNameLike(firstNameLike).LastName(lastName).LastNameLike(lastNameLike).Email(email).EmailLike(emailLike).MemberOfGroup(memberOfGroup).MemberOfTenant(memberOfTenant).PotentialStarter(potentialStarter).Execute()

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
	id := "id_example" // string | Filter by user id (optional)
	idIn := "idIn_example" // string | Filter by a comma-separated list of user ids. (optional)
	firstName := "firstName_example" // string | Filter by the first name of the user. Exact match. (optional)
	firstNameLike := "firstNameLike_example" // string | Filter by the first name that the parameter is a substring of. (optional)
	lastName := "lastName_example" // string | Filter by the last name of the user. Exact match. (optional)
	lastNameLike := "lastNameLike_example" // string | Filter by the last name that the parameter is a substring of. (optional)
	email := "email_example" // string | Filter by the email of the user. Exact match. (optional)
	emailLike := "emailLike_example" // string | Filter by the email that the parameter is a substring of. (optional)
	memberOfGroup := "memberOfGroup_example" // string | Filter for users which are members of the given group. (optional)
	memberOfTenant := "memberOfTenant_example" // string | Filter for users which are members of the given tenant. (optional)
	potentialStarter := "potentialStarter_example" // string | Only select Users that are potential starter for the given process definition. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserCount(context.Background()).Id(id).IdIn(idIn).FirstName(firstName).FirstNameLike(firstNameLike).LastName(lastName).LastNameLike(lastNameLike).Email(email).EmailLike(emailLike).MemberOfGroup(memberOfGroup).MemberOfTenant(memberOfTenant).PotentialStarter(potentialStarter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter by user id | 
 **idIn** | **string** | Filter by a comma-separated list of user ids. | 
 **firstName** | **string** | Filter by the first name of the user. Exact match. | 
 **firstNameLike** | **string** | Filter by the first name that the parameter is a substring of. | 
 **lastName** | **string** | Filter by the last name of the user. Exact match. | 
 **lastNameLike** | **string** | Filter by the last name that the parameter is a substring of. | 
 **email** | **string** | Filter by the email of the user. Exact match. | 
 **emailLike** | **string** | Filter by the email that the parameter is a substring of. | 
 **memberOfGroup** | **string** | Filter for users which are members of the given group. | 
 **memberOfTenant** | **string** | Filter for users which are members of the given tenant. | 
 **potentialStarter** | **string** | Only select Users that are potential starter for the given process definition. | 

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


## GetUserProfile

> UserProfileDto GetUserProfile(ctx, id).Execute()

Get Profile



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
	id := "id_example" // string | The id of the user to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserProfile`: UserProfileDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the user to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserProfileDto**](UserProfileDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> []UserProfileDto GetUsers(ctx).Id(id).IdIn(idIn).FirstName(firstName).FirstNameLike(firstNameLike).LastName(lastName).LastNameLike(lastNameLike).Email(email).EmailLike(emailLike).MemberOfGroup(memberOfGroup).MemberOfTenant(memberOfTenant).PotentialStarter(potentialStarter).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

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
	id := "id_example" // string | Filter by user id (optional)
	idIn := "idIn_example" // string | Filter by a comma-separated list of user ids. (optional)
	firstName := "firstName_example" // string | Filter by the first name of the user. Exact match. (optional)
	firstNameLike := "firstNameLike_example" // string | Filter by the first name that the parameter is a substring of. (optional)
	lastName := "lastName_example" // string | Filter by the last name of the user. Exact match. (optional)
	lastNameLike := "lastNameLike_example" // string | Filter by the last name that the parameter is a substring of. (optional)
	email := "email_example" // string | Filter by the email of the user. Exact match. (optional)
	emailLike := "emailLike_example" // string | Filter by the email that the parameter is a substring of. (optional)
	memberOfGroup := "memberOfGroup_example" // string | Filter for users which are members of the given group. (optional)
	memberOfTenant := "memberOfTenant_example" // string | Filter for users which are members of the given tenant. (optional)
	potentialStarter := "potentialStarter_example" // string | Only select Users that are potential starter for the given process definition. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUsers(context.Background()).Id(id).IdIn(idIn).FirstName(firstName).FirstNameLike(firstNameLike).LastName(lastName).LastNameLike(lastNameLike).Email(email).EmailLike(emailLike).MemberOfGroup(memberOfGroup).MemberOfTenant(memberOfTenant).PotentialStarter(potentialStarter).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: []UserProfileDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter by user id | 
 **idIn** | **string** | Filter by a comma-separated list of user ids. | 
 **firstName** | **string** | Filter by the first name of the user. Exact match. | 
 **firstNameLike** | **string** | Filter by the first name that the parameter is a substring of. | 
 **lastName** | **string** | Filter by the last name of the user. Exact match. | 
 **lastNameLike** | **string** | Filter by the last name that the parameter is a substring of. | 
 **email** | **string** | Filter by the email of the user. Exact match. | 
 **emailLike** | **string** | Filter by the email that the parameter is a substring of. | 
 **memberOfGroup** | **string** | Filter for users which are members of the given group. | 
 **memberOfTenant** | **string** | Filter for users which are members of the given tenant. | 
 **potentialStarter** | **string** | Only select Users that are potential starter for the given process definition. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]UserProfileDto**](UserProfileDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockUser

> UnlockUser(ctx, id).Execute()

Unlock User



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
	id := "id_example" // string | The id of the user to be unlocked.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UnlockUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UnlockUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the user to be unlocked. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockUserRequest struct via the builder pattern


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


## UpdateCredentials

> UpdateCredentials(ctx, id).UserCredentialsDto(userCredentialsDto).Execute()

Update Credentials



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
	id := "id_example" // string | The id of the user to be updated.
	userCredentialsDto := *openapiclient.NewUserCredentialsDto() // UserCredentialsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateCredentials(context.Background(), id).UserCredentialsDto(userCredentialsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the user to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userCredentialsDto** | [**UserCredentialsDto**](UserCredentialsDto.md) |  | 

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


## UpdateProfile

> UpdateProfile(ctx, id).UserProfileDto(userProfileDto).Execute()

Update User Profile



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
	id := "id_example" // string | The id of the user.
	userProfileDto := *openapiclient.NewUserProfileDto() // UserProfileDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateProfile(context.Background(), id).UserProfileDto(userProfileDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userProfileDto** | [**UserProfileDto**](UserProfileDto.md) |  | 

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

