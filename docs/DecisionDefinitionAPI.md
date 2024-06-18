# \DecisionDefinitionAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EvaluateDecisionById**](DecisionDefinitionAPI.md#EvaluateDecisionById) | **Post** /decision-definition/{id}/evaluate | Evaluate By Id
[**EvaluateDecisionByKey**](DecisionDefinitionAPI.md#EvaluateDecisionByKey) | **Post** /decision-definition/key/{key}/evaluate | Evaluate By Key
[**EvaluateDecisionByKeyAndTenant**](DecisionDefinitionAPI.md#EvaluateDecisionByKeyAndTenant) | **Post** /decision-definition/key/{key}/tenant-id/{tenant-id}/evaluate | Evaluate By Key And Tenant
[**GetDecisionDefinitionById**](DecisionDefinitionAPI.md#GetDecisionDefinitionById) | **Get** /decision-definition/{id} | Get Decision Definition By Id
[**GetDecisionDefinitionByKey**](DecisionDefinitionAPI.md#GetDecisionDefinitionByKey) | **Get** /decision-definition/key/{key} | Get Decision Definition By Key
[**GetDecisionDefinitionByKeyAndTenantId**](DecisionDefinitionAPI.md#GetDecisionDefinitionByKeyAndTenantId) | **Get** /decision-definition/key/{key}/tenant-id/{tenant-id} | Get Decision Definition By Key And Tenant Id
[**GetDecisionDefinitionDiagram**](DecisionDefinitionAPI.md#GetDecisionDefinitionDiagram) | **Get** /decision-definition/{id}/diagram | Get Diagram
[**GetDecisionDefinitionDiagramByKey**](DecisionDefinitionAPI.md#GetDecisionDefinitionDiagramByKey) | **Get** /decision-definition/key/{key}/diagram | Get Diagram By Key
[**GetDecisionDefinitionDiagramByKeyAndTenant**](DecisionDefinitionAPI.md#GetDecisionDefinitionDiagramByKeyAndTenant) | **Get** /decision-definition/key/{key}/tenant-id/{tenant-id}/diagram | Get Diagram By Key And Tenant
[**GetDecisionDefinitionDmnXmlById**](DecisionDefinitionAPI.md#GetDecisionDefinitionDmnXmlById) | **Get** /decision-definition/{id}/xml | Get XML By Id
[**GetDecisionDefinitionDmnXmlByKey**](DecisionDefinitionAPI.md#GetDecisionDefinitionDmnXmlByKey) | **Get** /decision-definition/key/{key}/xml | Get XML By Key
[**GetDecisionDefinitionDmnXmlByKeyAndTenant**](DecisionDefinitionAPI.md#GetDecisionDefinitionDmnXmlByKeyAndTenant) | **Get** /decision-definition/key/{key}/tenant-id/{tenant-id}/xml | Get XML By Key and Tenant
[**GetDecisionDefinitions**](DecisionDefinitionAPI.md#GetDecisionDefinitions) | **Get** /decision-definition | Get List
[**GetDecisionDefinitionsCount**](DecisionDefinitionAPI.md#GetDecisionDefinitionsCount) | **Get** /decision-definition/count | Get List Count
[**UpdateHistoryTimeToLiveByDecisionDefinitionId**](DecisionDefinitionAPI.md#UpdateHistoryTimeToLiveByDecisionDefinitionId) | **Put** /decision-definition/{id}/history-time-to-live | Update History Time to Live
[**UpdateHistoryTimeToLiveByDecisionDefinitionKey**](DecisionDefinitionAPI.md#UpdateHistoryTimeToLiveByDecisionDefinitionKey) | **Put** /decision-definition/key/{key}/history-time-to-live | Update History Time to Live By Key
[**UpdateHistoryTimeToLiveByDecisionDefinitionKeyAndTenant**](DecisionDefinitionAPI.md#UpdateHistoryTimeToLiveByDecisionDefinitionKeyAndTenant) | **Put** /decision-definition/key/{key}/tenant-id/{tenant-id}/history-time-to-live | Update History Time to Live By Key And Tenant



## EvaluateDecisionById

> []map[string]VariableValueDto EvaluateDecisionById(ctx, id).EvaluateDecisionDto(evaluateDecisionDto).Execute()

Evaluate By Id



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
	id := "id_example" // string | The id of the decision definition to be evaluated.
	evaluateDecisionDto := *openapiclient.NewEvaluateDecisionDto() // EvaluateDecisionDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.EvaluateDecisionById(context.Background(), id).EvaluateDecisionDto(evaluateDecisionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.EvaluateDecisionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateDecisionById`: []map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.EvaluateDecisionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the decision definition to be evaluated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateDecisionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **evaluateDecisionDto** | [**EvaluateDecisionDto**](EvaluateDecisionDto.md) |  | 

### Return type

[**[]map[string]VariableValueDto**](map.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvaluateDecisionByKey

> []map[string]VariableValueDto EvaluateDecisionByKey(ctx, key).EvaluateDecisionDto(evaluateDecisionDto).Execute()

Evaluate By Key



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
	key := "key_example" // string | The key of the decision definition (the latest version thereof) to be evaluated.
	evaluateDecisionDto := *openapiclient.NewEvaluateDecisionDto() // EvaluateDecisionDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.EvaluateDecisionByKey(context.Background(), key).EvaluateDecisionDto(evaluateDecisionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.EvaluateDecisionByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateDecisionByKey`: []map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.EvaluateDecisionByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision definition (the latest version thereof) to be evaluated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateDecisionByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **evaluateDecisionDto** | [**EvaluateDecisionDto**](EvaluateDecisionDto.md) |  | 

### Return type

[**[]map[string]VariableValueDto**](map.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvaluateDecisionByKeyAndTenant

> []map[string]VariableValueDto EvaluateDecisionByKeyAndTenant(ctx, key, tenantId).EvaluateDecisionDto(evaluateDecisionDto).Execute()

Evaluate By Key And Tenant



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
	key := "key_example" // string | The key of the decision definition (the latest version thereof) to be evaluated.
	tenantId := "tenantId_example" // string | The id of the tenant the decision definition belongs to.
	evaluateDecisionDto := *openapiclient.NewEvaluateDecisionDto() // EvaluateDecisionDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.EvaluateDecisionByKeyAndTenant(context.Background(), key, tenantId).EvaluateDecisionDto(evaluateDecisionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.EvaluateDecisionByKeyAndTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateDecisionByKeyAndTenant`: []map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.EvaluateDecisionByKeyAndTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision definition (the latest version thereof) to be evaluated. | 
**tenantId** | **string** | The id of the tenant the decision definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateDecisionByKeyAndTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **evaluateDecisionDto** | [**EvaluateDecisionDto**](EvaluateDecisionDto.md) |  | 

### Return type

[**[]map[string]VariableValueDto**](map.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionDefinitionById

> DecisionDefinitionDto GetDecisionDefinitionById(ctx, id).Execute()

Get Decision Definition By Id



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
	id := "id_example" // string | The id of the decision definition to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.GetDecisionDefinitionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionDefinitionById`: DecisionDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.GetDecisionDefinitionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the decision definition to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionDefinitionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DecisionDefinitionDto**](DecisionDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionDefinitionByKey

> DecisionDefinitionDto GetDecisionDefinitionByKey(ctx, key).Execute()

Get Decision Definition By Key



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
	key := "key_example" // string | The key of the decision definition (the latest version thereof) to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.GetDecisionDefinitionByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionDefinitionByKey`: DecisionDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.GetDecisionDefinitionByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision definition (the latest version thereof) to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionDefinitionByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DecisionDefinitionDto**](DecisionDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionDefinitionByKeyAndTenantId

> DecisionDefinitionDto GetDecisionDefinitionByKeyAndTenantId(ctx, key, tenantId).Execute()

Get Decision Definition By Key And Tenant Id



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
	key := "key_example" // string | The key of the decision definition (the latest version thereof) to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant the decision definition belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionByKeyAndTenantId(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.GetDecisionDefinitionByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionDefinitionByKeyAndTenantId`: DecisionDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.GetDecisionDefinitionByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision definition (the latest version thereof) to be retrieved. | 
**tenantId** | **string** | The id of the tenant the decision definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionDefinitionByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DecisionDefinitionDto**](DecisionDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionDefinitionDiagram

> *os.File GetDecisionDefinitionDiagram(ctx, id).Execute()

Get Diagram



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
	id := "id_example" // string | The id of the process definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDiagram(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.GetDecisionDefinitionDiagram``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionDefinitionDiagram`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.GetDecisionDefinitionDiagram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionDefinitionDiagramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionDefinitionDiagramByKey

> *os.File GetDecisionDefinitionDiagramByKey(ctx, key).Execute()

Get Diagram By Key



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
	key := "key_example" // string | The key of the decision definition (the latest version thereof) to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDiagramByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.GetDecisionDefinitionDiagramByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionDefinitionDiagramByKey`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.GetDecisionDefinitionDiagramByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision definition (the latest version thereof) to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionDefinitionDiagramByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionDefinitionDiagramByKeyAndTenant

> *os.File GetDecisionDefinitionDiagramByKeyAndTenant(ctx, key, tenantId).Execute()

Get Diagram By Key And Tenant



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
	key := "key_example" // string | The key of the decision definition (the latest version thereof) to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant the decision definition belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDiagramByKeyAndTenant(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.GetDecisionDefinitionDiagramByKeyAndTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionDefinitionDiagramByKeyAndTenant`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.GetDecisionDefinitionDiagramByKeyAndTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision definition (the latest version thereof) to be retrieved. | 
**tenantId** | **string** | The id of the tenant the decision definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionDefinitionDiagramByKeyAndTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionDefinitionDmnXmlById

> DecisionDefinitionDiagramDto GetDecisionDefinitionDmnXmlById(ctx, id).Execute()

Get XML By Id



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
	id := "id_example" // string | The id of the decision definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionDefinitionDmnXmlById`: DecisionDefinitionDiagramDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the decision definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionDefinitionDmnXmlByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DecisionDefinitionDiagramDto**](DecisionDefinitionDiagramDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionDefinitionDmnXmlByKey

> DecisionDefinitionDiagramDto GetDecisionDefinitionDmnXmlByKey(ctx, key).Execute()

Get XML By Key



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
	key := "key_example" // string | The key of the decision definition (the latest version thereof).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionDefinitionDmnXmlByKey`: DecisionDefinitionDiagramDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision definition (the latest version thereof). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionDefinitionDmnXmlByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DecisionDefinitionDiagramDto**](DecisionDefinitionDiagramDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionDefinitionDmnXmlByKeyAndTenant

> DecisionDefinitionDiagramDto GetDecisionDefinitionDmnXmlByKeyAndTenant(ctx, key, tenantId).Execute()

Get XML By Key and Tenant



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
	key := "key_example" // string | The key of the decision definition (the latest version thereof).
	tenantId := "tenantId_example" // string | The id of the tenant the decision definition belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlByKeyAndTenant(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlByKeyAndTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionDefinitionDmnXmlByKeyAndTenant`: DecisionDefinitionDiagramDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.GetDecisionDefinitionDmnXmlByKeyAndTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision definition (the latest version thereof). | 
**tenantId** | **string** | The id of the tenant the decision definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionDefinitionDmnXmlByKeyAndTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DecisionDefinitionDiagramDto**](DecisionDefinitionDiagramDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionDefinitions

> []DecisionDefinitionDto GetDecisionDefinitions(ctx).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).DecisionDefinitionId(decisionDefinitionId).DecisionDefinitionIdIn(decisionDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).DeployedAfter(deployedAfter).DeployedAt(deployedAt).Key(key).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionKey(decisionRequirementsDefinitionKey).WithoutDecisionRequirementsDefinition(withoutDecisionRequirementsDefinition).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDecisionDefinitionsWithoutTenantId(includeDecisionDefinitionsWithoutTenantId).VersionTag(versionTag).VersionTagLike(versionTagLike).Execute()

Get List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)
	decisionDefinitionId := "decisionDefinitionId_example" // string | Filter by decision definition id. (optional)
	decisionDefinitionIdIn := "decisionDefinitionIdIn_example" // string | Filter by decision definition ids. (optional)
	name := "name_example" // string | Filter by decision definition name. (optional)
	nameLike := "nameLike_example" // string | Filter by decision definition names that the parameter is a substring of. (optional)
	deploymentId := "deploymentId_example" // string | Filter by the deployment the id belongs to. (optional)
	deployedAfter := time.Now() // time.Time | Filter by the deploy time of the deployment the decision definition belongs to. Only selects decision definitions that have been deployed after (exclusive) a specific time. (optional)
	deployedAt := time.Now() // time.Time | Filter by the deploy time of the deployment the decision definition belongs to. Only selects decision definitions that have been deployed at a specific time (exact match). (optional)
	key := "key_example" // string | Filter by decision definition key, i.e., the id in the DMN 1.0 XML. Exact match. (optional)
	keyLike := "keyLike_example" // string | Filter by decision definition keys that the parameter is a substring of. (optional)
	category := "category_example" // string | Filter by decision definition category. Exact match. (optional)
	categoryLike := "categoryLike_example" // string | Filter by decision definition categories that the parameter is a substring of. (optional)
	version := int32(56) // int32 | Filter by decision definition version. (optional)
	latestVersion := true // bool | Only include those decision definitions that are latest versions. Value may only be `true`, as `false` is the default behavior. (optional)
	resourceName := "resourceName_example" // string | Filter by the name of the decision definition resource. Exact match. (optional)
	resourceNameLike := "resourceNameLike_example" // string | Filter by names of those decision definition resources that the parameter is a substring of. (optional)
	decisionRequirementsDefinitionId := "decisionRequirementsDefinitionId_example" // string | Filter by the id of the decision requirements definition this decision definition belongs to. (optional)
	decisionRequirementsDefinitionKey := "decisionRequirementsDefinitionKey_example" // string | Filter by the key of the decision requirements definition this decision definition belongs to. (optional)
	withoutDecisionRequirementsDefinition := true // bool | Only include decision definitions which does not belongs to any decision requirements definition. Value may only be `true`, as `false` is the default behavior. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of `Strings`. A decision definition must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include decision definitions which belong to no tenant. Value can effectively only be `true`, as `false` is the default behavior. (optional)
	includeDecisionDefinitionsWithoutTenantId := true // bool | Include decision definitions which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional)
	versionTag := "versionTag_example" // string | Filter by the version tag. (optional)
	versionTagLike := "versionTagLike_example" // string | Filter by the version tags of those decision definition resources that the parameter is a substring of. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitions(context.Background()).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).DecisionDefinitionId(decisionDefinitionId).DecisionDefinitionIdIn(decisionDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).DeployedAfter(deployedAfter).DeployedAt(deployedAt).Key(key).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionKey(decisionRequirementsDefinitionKey).WithoutDecisionRequirementsDefinition(withoutDecisionRequirementsDefinition).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDecisionDefinitionsWithoutTenantId(includeDecisionDefinitionsWithoutTenantId).VersionTag(versionTag).VersionTagLike(versionTagLike).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.GetDecisionDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionDefinitions`: []DecisionDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.GetDecisionDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **decisionDefinitionId** | **string** | Filter by decision definition id. | 
 **decisionDefinitionIdIn** | **string** | Filter by decision definition ids. | 
 **name** | **string** | Filter by decision definition name. | 
 **nameLike** | **string** | Filter by decision definition names that the parameter is a substring of. | 
 **deploymentId** | **string** | Filter by the deployment the id belongs to. | 
 **deployedAfter** | **time.Time** | Filter by the deploy time of the deployment the decision definition belongs to. Only selects decision definitions that have been deployed after (exclusive) a specific time. | 
 **deployedAt** | **time.Time** | Filter by the deploy time of the deployment the decision definition belongs to. Only selects decision definitions that have been deployed at a specific time (exact match). | 
 **key** | **string** | Filter by decision definition key, i.e., the id in the DMN 1.0 XML. Exact match. | 
 **keyLike** | **string** | Filter by decision definition keys that the parameter is a substring of. | 
 **category** | **string** | Filter by decision definition category. Exact match. | 
 **categoryLike** | **string** | Filter by decision definition categories that the parameter is a substring of. | 
 **version** | **int32** | Filter by decision definition version. | 
 **latestVersion** | **bool** | Only include those decision definitions that are latest versions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **resourceName** | **string** | Filter by the name of the decision definition resource. Exact match. | 
 **resourceNameLike** | **string** | Filter by names of those decision definition resources that the parameter is a substring of. | 
 **decisionRequirementsDefinitionId** | **string** | Filter by the id of the decision requirements definition this decision definition belongs to. | 
 **decisionRequirementsDefinitionKey** | **string** | Filter by the key of the decision requirements definition this decision definition belongs to. | 
 **withoutDecisionRequirementsDefinition** | **bool** | Only include decision definitions which does not belongs to any decision requirements definition. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of &#x60;Strings&#x60;. A decision definition must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include decision definitions which belong to no tenant. Value can effectively only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeDecisionDefinitionsWithoutTenantId** | **bool** | Include decision definitions which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **versionTag** | **string** | Filter by the version tag. | 
 **versionTagLike** | **string** | Filter by the version tags of those decision definition resources that the parameter is a substring of. | 

### Return type

[**[]DecisionDefinitionDto**](DecisionDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionDefinitionsCount

> CountResultDto GetDecisionDefinitionsCount(ctx).DecisionDefinitionId(decisionDefinitionId).DecisionDefinitionIdIn(decisionDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).DeployedAfter(deployedAfter).DeployedAt(deployedAt).Key(key).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionKey(decisionRequirementsDefinitionKey).WithoutDecisionRequirementsDefinition(withoutDecisionRequirementsDefinition).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDecisionDefinitionsWithoutTenantId(includeDecisionDefinitionsWithoutTenantId).VersionTag(versionTag).VersionTagLike(versionTagLike).Execute()

Get List Count



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	decisionDefinitionId := "decisionDefinitionId_example" // string | Filter by decision definition id. (optional)
	decisionDefinitionIdIn := "decisionDefinitionIdIn_example" // string | Filter by decision definition ids. (optional)
	name := "name_example" // string | Filter by decision definition name. (optional)
	nameLike := "nameLike_example" // string | Filter by decision definition names that the parameter is a substring of. (optional)
	deploymentId := "deploymentId_example" // string | Filter by the deployment the id belongs to. (optional)
	deployedAfter := time.Now() // time.Time | Filter by the deploy time of the deployment the decision definition belongs to. Only selects decision definitions that have been deployed after (exclusive) a specific time. (optional)
	deployedAt := time.Now() // time.Time | Filter by the deploy time of the deployment the decision definition belongs to. Only selects decision definitions that have been deployed at a specific time (exact match). (optional)
	key := "key_example" // string | Filter by decision definition key, i.e., the id in the DMN 1.0 XML. Exact match. (optional)
	keyLike := "keyLike_example" // string | Filter by decision definition keys that the parameter is a substring of. (optional)
	category := "category_example" // string | Filter by decision definition category. Exact match. (optional)
	categoryLike := "categoryLike_example" // string | Filter by decision definition categories that the parameter is a substring of. (optional)
	version := int32(56) // int32 | Filter by decision definition version. (optional)
	latestVersion := true // bool | Only include those decision definitions that are latest versions. Value may only be `true`, as `false` is the default behavior. (optional)
	resourceName := "resourceName_example" // string | Filter by the name of the decision definition resource. Exact match. (optional)
	resourceNameLike := "resourceNameLike_example" // string | Filter by names of those decision definition resources that the parameter is a substring of. (optional)
	decisionRequirementsDefinitionId := "decisionRequirementsDefinitionId_example" // string | Filter by the id of the decision requirements definition this decision definition belongs to. (optional)
	decisionRequirementsDefinitionKey := "decisionRequirementsDefinitionKey_example" // string | Filter by the key of the decision requirements definition this decision definition belongs to. (optional)
	withoutDecisionRequirementsDefinition := true // bool | Only include decision definitions which does not belongs to any decision requirements definition. Value may only be `true`, as `false` is the default behavior. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of `Strings`. A decision definition must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include decision definitions which belong to no tenant. Value can effectively only be `true`, as `false` is the default behavior. (optional)
	includeDecisionDefinitionsWithoutTenantId := true // bool | Include decision definitions which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional)
	versionTag := "versionTag_example" // string | Filter by the version tag. (optional)
	versionTagLike := "versionTagLike_example" // string | Filter by the version tags of those decision definition resources that the parameter is a substring of. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionDefinitionAPI.GetDecisionDefinitionsCount(context.Background()).DecisionDefinitionId(decisionDefinitionId).DecisionDefinitionIdIn(decisionDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).DeployedAfter(deployedAfter).DeployedAt(deployedAt).Key(key).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionKey(decisionRequirementsDefinitionKey).WithoutDecisionRequirementsDefinition(withoutDecisionRequirementsDefinition).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDecisionDefinitionsWithoutTenantId(includeDecisionDefinitionsWithoutTenantId).VersionTag(versionTag).VersionTagLike(versionTagLike).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.GetDecisionDefinitionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionDefinitionsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionDefinitionAPI.GetDecisionDefinitionsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionDefinitionsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decisionDefinitionId** | **string** | Filter by decision definition id. | 
 **decisionDefinitionIdIn** | **string** | Filter by decision definition ids. | 
 **name** | **string** | Filter by decision definition name. | 
 **nameLike** | **string** | Filter by decision definition names that the parameter is a substring of. | 
 **deploymentId** | **string** | Filter by the deployment the id belongs to. | 
 **deployedAfter** | **time.Time** | Filter by the deploy time of the deployment the decision definition belongs to. Only selects decision definitions that have been deployed after (exclusive) a specific time. | 
 **deployedAt** | **time.Time** | Filter by the deploy time of the deployment the decision definition belongs to. Only selects decision definitions that have been deployed at a specific time (exact match). | 
 **key** | **string** | Filter by decision definition key, i.e., the id in the DMN 1.0 XML. Exact match. | 
 **keyLike** | **string** | Filter by decision definition keys that the parameter is a substring of. | 
 **category** | **string** | Filter by decision definition category. Exact match. | 
 **categoryLike** | **string** | Filter by decision definition categories that the parameter is a substring of. | 
 **version** | **int32** | Filter by decision definition version. | 
 **latestVersion** | **bool** | Only include those decision definitions that are latest versions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **resourceName** | **string** | Filter by the name of the decision definition resource. Exact match. | 
 **resourceNameLike** | **string** | Filter by names of those decision definition resources that the parameter is a substring of. | 
 **decisionRequirementsDefinitionId** | **string** | Filter by the id of the decision requirements definition this decision definition belongs to. | 
 **decisionRequirementsDefinitionKey** | **string** | Filter by the key of the decision requirements definition this decision definition belongs to. | 
 **withoutDecisionRequirementsDefinition** | **bool** | Only include decision definitions which does not belongs to any decision requirements definition. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of &#x60;Strings&#x60;. A decision definition must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include decision definitions which belong to no tenant. Value can effectively only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeDecisionDefinitionsWithoutTenantId** | **bool** | Include decision definitions which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **versionTag** | **string** | Filter by the version tag. | 
 **versionTagLike** | **string** | Filter by the version tags of those decision definition resources that the parameter is a substring of. | 

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


## UpdateHistoryTimeToLiveByDecisionDefinitionId

> UpdateHistoryTimeToLiveByDecisionDefinitionId(ctx, id).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()

Update History Time to Live



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
	id := "id_example" // string | The id of the decision definition to change history time to live.
	historyTimeToLiveDto := *openapiclient.NewHistoryTimeToLiveDto() // HistoryTimeToLiveDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DecisionDefinitionAPI.UpdateHistoryTimeToLiveByDecisionDefinitionId(context.Background(), id).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.UpdateHistoryTimeToLiveByDecisionDefinitionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the decision definition to change history time to live. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHistoryTimeToLiveByDecisionDefinitionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **historyTimeToLiveDto** | [**HistoryTimeToLiveDto**](HistoryTimeToLiveDto.md) |  | 

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


## UpdateHistoryTimeToLiveByDecisionDefinitionKey

> UpdateHistoryTimeToLiveByDecisionDefinitionKey(ctx, key).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()

Update History Time to Live By Key



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
	key := "key_example" // string | The key of the decision definitions to change history time to live.
	historyTimeToLiveDto := *openapiclient.NewHistoryTimeToLiveDto() // HistoryTimeToLiveDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DecisionDefinitionAPI.UpdateHistoryTimeToLiveByDecisionDefinitionKey(context.Background(), key).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.UpdateHistoryTimeToLiveByDecisionDefinitionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision definitions to change history time to live. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHistoryTimeToLiveByDecisionDefinitionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **historyTimeToLiveDto** | [**HistoryTimeToLiveDto**](HistoryTimeToLiveDto.md) |  | 

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


## UpdateHistoryTimeToLiveByDecisionDefinitionKeyAndTenant

> UpdateHistoryTimeToLiveByDecisionDefinitionKeyAndTenant(ctx, key, tenantId).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()

Update History Time to Live By Key And Tenant



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
	key := "key_example" // string | The key of the decision definitions to change history time to live.
	tenantId := "tenantId_example" // string | The id of the tenant the decision definition belongs to.
	historyTimeToLiveDto := *openapiclient.NewHistoryTimeToLiveDto() // HistoryTimeToLiveDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DecisionDefinitionAPI.UpdateHistoryTimeToLiveByDecisionDefinitionKeyAndTenant(context.Background(), key, tenantId).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionDefinitionAPI.UpdateHistoryTimeToLiveByDecisionDefinitionKeyAndTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision definitions to change history time to live. | 
**tenantId** | **string** | The id of the tenant the decision definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHistoryTimeToLiveByDecisionDefinitionKeyAndTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **historyTimeToLiveDto** | [**HistoryTimeToLiveDto**](HistoryTimeToLiveDto.md) |  | 

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

