# \MigrationAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteMigrationPlan**](MigrationAPI.md#ExecuteMigrationPlan) | **Post** /migration/execute | Execute Migration Plan
[**ExecuteMigrationPlanAsync**](MigrationAPI.md#ExecuteMigrationPlanAsync) | **Post** /migration/executeAsync | Execute Migration Plan Async (Batch)
[**GenerateMigrationPlan**](MigrationAPI.md#GenerateMigrationPlan) | **Post** /migration/generate | Generate Migration Plan
[**ValidateMigrationPlan**](MigrationAPI.md#ValidateMigrationPlan) | **Post** /migration/validate | Validate Migration Plan



## ExecuteMigrationPlan

> ExecuteMigrationPlan(ctx).MigrationExecutionDto(migrationExecutionDto).Execute()

Execute Migration Plan



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
	migrationExecutionDto := *openapiclient.NewMigrationExecutionDto() // MigrationExecutionDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MigrationAPI.ExecuteMigrationPlan(context.Background()).MigrationExecutionDto(migrationExecutionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.ExecuteMigrationPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteMigrationPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationExecutionDto** | [**MigrationExecutionDto**](MigrationExecutionDto.md) |  | 

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


## ExecuteMigrationPlanAsync

> BatchDto ExecuteMigrationPlanAsync(ctx).MigrationExecutionDto(migrationExecutionDto).Execute()

Execute Migration Plan Async (Batch)



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
	migrationExecutionDto := *openapiclient.NewMigrationExecutionDto() // MigrationExecutionDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrationAPI.ExecuteMigrationPlanAsync(context.Background()).MigrationExecutionDto(migrationExecutionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.ExecuteMigrationPlanAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteMigrationPlanAsync`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.ExecuteMigrationPlanAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteMigrationPlanAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationExecutionDto** | [**MigrationExecutionDto**](MigrationExecutionDto.md) |  | 

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


## GenerateMigrationPlan

> MigrationPlanDto GenerateMigrationPlan(ctx).MigrationPlanGenerationDto(migrationPlanGenerationDto).Execute()

Generate Migration Plan



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
	migrationPlanGenerationDto := *openapiclient.NewMigrationPlanGenerationDto() // MigrationPlanGenerationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrationAPI.GenerateMigrationPlan(context.Background()).MigrationPlanGenerationDto(migrationPlanGenerationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.GenerateMigrationPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateMigrationPlan`: MigrationPlanDto
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.GenerateMigrationPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateMigrationPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationPlanGenerationDto** | [**MigrationPlanGenerationDto**](MigrationPlanGenerationDto.md) |  | 

### Return type

[**MigrationPlanDto**](MigrationPlanDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateMigrationPlan

> MigrationPlanReportDto ValidateMigrationPlan(ctx).MigrationPlanDto(migrationPlanDto).Execute()

Validate Migration Plan



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
	migrationPlanDto := *openapiclient.NewMigrationPlanDto() // MigrationPlanDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrationAPI.ValidateMigrationPlan(context.Background()).MigrationPlanDto(migrationPlanDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.ValidateMigrationPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateMigrationPlan`: MigrationPlanReportDto
	fmt.Fprintf(os.Stdout, "Response from `MigrationAPI.ValidateMigrationPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateMigrationPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationPlanDto** | [**MigrationPlanDto**](MigrationPlanDto.md) |  | 

### Return type

[**MigrationPlanReportDto**](MigrationPlanReportDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

