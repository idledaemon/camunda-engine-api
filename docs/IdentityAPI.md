# \IdentityAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckPassword**](IdentityAPI.md#CheckPassword) | **Post** /identity/password-policy | Validate Password
[**GetGroupInfo**](IdentityAPI.md#GetGroupInfo) | **Get** /identity/groups | Get a User&#39;s Groups
[**GetPasswordPolicy**](IdentityAPI.md#GetPasswordPolicy) | **Get** /identity/password-policy | Get Password Policy
[**VerifyUser**](IdentityAPI.md#VerifyUser) | **Post** /identity/verify | Verify User



## CheckPassword

> CheckPasswordPolicyResultDto CheckPassword(ctx).PasswordPolicyRequestDto(passwordPolicyRequestDto).Execute()

Validate Password



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
	passwordPolicyRequestDto := *openapiclient.NewPasswordPolicyRequestDto() // PasswordPolicyRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.CheckPassword(context.Background()).PasswordPolicyRequestDto(passwordPolicyRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.CheckPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckPassword`: CheckPasswordPolicyResultDto
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.CheckPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordPolicyRequestDto** | [**PasswordPolicyRequestDto**](PasswordPolicyRequestDto.md) |  | 

### Return type

[**CheckPasswordPolicyResultDto**](CheckPasswordPolicyResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupInfo

> IdentityServiceGroupInfoDto GetGroupInfo(ctx).UserId(userId).Execute()

Get a User's Groups



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
	userId := "userId_example" // string | The id of the user to get the groups for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.GetGroupInfo(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.GetGroupInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupInfo`: IdentityServiceGroupInfoDto
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.GetGroupInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The id of the user to get the groups for. | 

### Return type

[**IdentityServiceGroupInfoDto**](IdentityServiceGroupInfoDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordPolicy

> PasswordPolicyDto GetPasswordPolicy(ctx).Execute()

Get Password Policy



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
	resp, r, err := apiClient.IdentityAPI.GetPasswordPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.GetPasswordPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPasswordPolicy`: PasswordPolicyDto
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.GetPasswordPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordPolicyRequest struct via the builder pattern


### Return type

[**PasswordPolicyDto**](PasswordPolicyDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyUser

> AuthenticationResult VerifyUser(ctx).BasicUserCredentialsDto(basicUserCredentialsDto).Execute()

Verify User



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
	basicUserCredentialsDto := *openapiclient.NewBasicUserCredentialsDto() // BasicUserCredentialsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.VerifyUser(context.Background()).BasicUserCredentialsDto(basicUserCredentialsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.VerifyUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyUser`: AuthenticationResult
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.VerifyUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **basicUserCredentialsDto** | [**BasicUserCredentialsDto**](BasicUserCredentialsDto.md) |  | 

### Return type

[**AuthenticationResult**](AuthenticationResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

