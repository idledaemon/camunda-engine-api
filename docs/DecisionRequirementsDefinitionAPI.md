# \DecisionRequirementsDefinitionAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDecisionRequirementsDefinitionById**](DecisionRequirementsDefinitionAPI.md#GetDecisionRequirementsDefinitionById) | **Get** /decision-requirements-definition/{id} | Get Decision Requirements Definition by ID
[**GetDecisionRequirementsDefinitionByKey**](DecisionRequirementsDefinitionAPI.md#GetDecisionRequirementsDefinitionByKey) | **Get** /decision-requirements-definition/key/{key} | Get Decision Requirements Definition by Key
[**GetDecisionRequirementsDefinitionByKeyAndTenantId**](DecisionRequirementsDefinitionAPI.md#GetDecisionRequirementsDefinitionByKeyAndTenantId) | **Get** /decision-requirements-definition/key/{key}/tenant-id/{tenant-id} | Get Decision Requirements Definition by Key and Tenant ID
[**GetDecisionRequirementsDefinitionDiagramById**](DecisionRequirementsDefinitionAPI.md#GetDecisionRequirementsDefinitionDiagramById) | **Get** /decision-requirements-definition/{id}/diagram | Get Decision Requirements Diagram by ID
[**GetDecisionRequirementsDefinitionDiagramByKey**](DecisionRequirementsDefinitionAPI.md#GetDecisionRequirementsDefinitionDiagramByKey) | **Get** /decision-requirements-definition/key/{key}/diagram | Get Decision Requirements Diagram by Key
[**GetDecisionRequirementsDefinitionDiagramByKeyAndTenantId**](DecisionRequirementsDefinitionAPI.md#GetDecisionRequirementsDefinitionDiagramByKeyAndTenantId) | **Get** /decision-requirements-definition/key/{key}/tenant-id/{tenant-id}/diagram | Get Decision Requirements Diagram by Key and Tenant ID
[**GetDecisionRequirementsDefinitionDmnXmlById**](DecisionRequirementsDefinitionAPI.md#GetDecisionRequirementsDefinitionDmnXmlById) | **Get** /decision-requirements-definition/{id}/xml | Get DMN XML by ID
[**GetDecisionRequirementsDefinitionDmnXmlByKey**](DecisionRequirementsDefinitionAPI.md#GetDecisionRequirementsDefinitionDmnXmlByKey) | **Get** /decision-requirements-definition/key/{key}/xml | Get DMN XML by Key
[**GetDecisionRequirementsDefinitionDmnXmlByKeyAndTenantId**](DecisionRequirementsDefinitionAPI.md#GetDecisionRequirementsDefinitionDmnXmlByKeyAndTenantId) | **Get** /decision-requirements-definition/key/{key}/tenant-id/{tenant-id}/xml | Get DMN XML by Key and Tenant ID
[**GetDecisionRequirementsDefinitions**](DecisionRequirementsDefinitionAPI.md#GetDecisionRequirementsDefinitions) | **Get** /decision-requirements-definition | Get Decision Requirements Definitions
[**GetDecisionRequirementsDefinitionsCount**](DecisionRequirementsDefinitionAPI.md#GetDecisionRequirementsDefinitionsCount) | **Get** /decision-requirements-definition/count | Get Decision Requirements Definition Count



## GetDecisionRequirementsDefinitionById

> DecisionRequirementsDefinitionDto GetDecisionRequirementsDefinitionById(ctx, id).Execute()

Get Decision Requirements Definition by ID



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
	id := "id_example" // string | The id of the decision requirements definition to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionRequirementsDefinitionById`: DecisionRequirementsDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the decision requirements definition to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionRequirementsDefinitionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DecisionRequirementsDefinitionDto**](DecisionRequirementsDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionRequirementsDefinitionByKey

> DecisionRequirementsDefinitionDto GetDecisionRequirementsDefinitionByKey(ctx, key).Execute()

Get Decision Requirements Definition by Key



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
	key := "key_example" // string | The key of the decision requirements definition (the latest version thereof) to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionRequirementsDefinitionByKey`: DecisionRequirementsDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision requirements definition (the latest version thereof) to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionRequirementsDefinitionByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DecisionRequirementsDefinitionDto**](DecisionRequirementsDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionRequirementsDefinitionByKeyAndTenantId

> DecisionRequirementsDefinitionDto GetDecisionRequirementsDefinitionByKeyAndTenantId(ctx, key, tenantId).Execute()

Get Decision Requirements Definition by Key and Tenant ID



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
	key := "key_example" // string | The key of the decision requirements definition (the latest version thereof) to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant to which the decision requirements definition belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionByKeyAndTenantId(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionRequirementsDefinitionByKeyAndTenantId`: DecisionRequirementsDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision requirements definition (the latest version thereof) to be retrieved. | 
**tenantId** | **string** | The id of the tenant to which the decision requirements definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionRequirementsDefinitionByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DecisionRequirementsDefinitionDto**](DecisionRequirementsDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionRequirementsDefinitionDiagramById

> *os.File GetDecisionRequirementsDefinitionDiagramById(ctx, id).Execute()

Get Decision Requirements Diagram by ID



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
	id := "id_example" // string | The id of the decision requirements definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDiagramById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDiagramById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionRequirementsDefinitionDiagramById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDiagramById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the decision requirements definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionRequirementsDefinitionDiagramByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionRequirementsDefinitionDiagramByKey

> *os.File GetDecisionRequirementsDefinitionDiagramByKey(ctx, key).Execute()

Get Decision Requirements Diagram by Key



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
	key := "key_example" // string | The key of the decision requirements definition (the latest version thereof) to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDiagramByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDiagramByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionRequirementsDefinitionDiagramByKey`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDiagramByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision requirements definition (the latest version thereof) to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionRequirementsDefinitionDiagramByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionRequirementsDefinitionDiagramByKeyAndTenantId

> *os.File GetDecisionRequirementsDefinitionDiagramByKeyAndTenantId(ctx, key, tenantId).Execute()

Get Decision Requirements Diagram by Key and Tenant ID



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
	key := "key_example" // string | The key of the decision requirements definition (the latest version thereof) to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant to which the decision requirements definition belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDiagramByKeyAndTenantId(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDiagramByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionRequirementsDefinitionDiagramByKeyAndTenantId`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDiagramByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision requirements definition (the latest version thereof) to be retrieved. | 
**tenantId** | **string** | The id of the tenant to which the decision requirements definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionRequirementsDefinitionDiagramByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionRequirementsDefinitionDmnXmlById

> DecisionRequirementsDefinitionXmlDto GetDecisionRequirementsDefinitionDmnXmlById(ctx, id).Execute()

Get DMN XML by ID



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
	id := "id_example" // string | The id of the decision requirements definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDmnXmlById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDmnXmlById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionRequirementsDefinitionDmnXmlById`: DecisionRequirementsDefinitionXmlDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDmnXmlById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the decision requirements definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionRequirementsDefinitionDmnXmlByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DecisionRequirementsDefinitionXmlDto**](DecisionRequirementsDefinitionXmlDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionRequirementsDefinitionDmnXmlByKey

> DecisionRequirementsDefinitionXmlDto GetDecisionRequirementsDefinitionDmnXmlByKey(ctx, key).Execute()

Get DMN XML by Key



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
	key := "key_example" // string | The key of the decision requirements definition (the latest version thereof) to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDmnXmlByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDmnXmlByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionRequirementsDefinitionDmnXmlByKey`: DecisionRequirementsDefinitionXmlDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDmnXmlByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision requirements definition (the latest version thereof) to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionRequirementsDefinitionDmnXmlByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DecisionRequirementsDefinitionXmlDto**](DecisionRequirementsDefinitionXmlDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionRequirementsDefinitionDmnXmlByKeyAndTenantId

> DecisionRequirementsDefinitionXmlDto GetDecisionRequirementsDefinitionDmnXmlByKeyAndTenantId(ctx, key, tenantId).Execute()

Get DMN XML by Key and Tenant ID



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
	key := "key_example" // string | The key of the decision requirements definition (the latest version thereof) to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant to which the decision requirements definition belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDmnXmlByKeyAndTenantId(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDmnXmlByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionRequirementsDefinitionDmnXmlByKeyAndTenantId`: DecisionRequirementsDefinitionXmlDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionDmnXmlByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the decision requirements definition (the latest version thereof) to be retrieved. | 
**tenantId** | **string** | The id of the tenant to which the decision requirements definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionRequirementsDefinitionDmnXmlByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DecisionRequirementsDefinitionXmlDto**](DecisionRequirementsDefinitionXmlDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionRequirementsDefinitions

> []DecisionRequirementsDefinitionDto GetDecisionRequirementsDefinitions(ctx).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionIdIn(decisionRequirementsDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).Key(key).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDecisionRequirementsDefinitionsWithoutTenantId(includeDecisionRequirementsDefinitionsWithoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Decision Requirements Definitions



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
	decisionRequirementsDefinitionId := "decisionRequirementsDefinitionId_example" // string | Filter by decision requirements definition id. (optional)
	decisionRequirementsDefinitionIdIn := "decisionRequirementsDefinitionIdIn_example" // string | Filter by decision requirements definition ids. (optional)
	name := "name_example" // string | Filter by decision requirements definition name. (optional)
	nameLike := "nameLike_example" // string | Filter by decision requirements definition names that the parameter is a substring of. (optional)
	deploymentId := "deploymentId_example" // string | Filter by the id of the deployment a decision requirement definition belongs to. (optional)
	key := "key_example" // string | Filter by decision requirements definition key, i.e., the id in the DMN 1.3 XML. Exact match. (optional)
	keyLike := "keyLike_example" // string | Filter by decision requirements definition keys that the parameter is a substring of. (optional)
	category := "category_example" // string | Filter by decision requirements definition category. Exact match. (optional)
	categoryLike := "categoryLike_example" // string | Filter by decision requirements definition categories that the parameter is a substring of. (optional)
	version := int32(56) // int32 | Filter by decision requirements definition version. (optional)
	latestVersion := true // bool | Only include those decision requirements definitions that are latest versions. Value may only be `true`, as `false` is the default behavior. (optional)
	resourceName := "resourceName_example" // string | Filter by the name of the decision requirements definition resource. Exact match. (optional)
	resourceNameLike := "resourceNameLike_example" // string | Filter by names of those decision requirements definition resources that the parameter is a substring of. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A decision requirements definition must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include decision requirements definitions which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	includeDecisionRequirementsDefinitionsWithoutTenantId := true // bool | Include decision requirements definitions which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitions(context.Background()).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionIdIn(decisionRequirementsDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).Key(key).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDecisionRequirementsDefinitionsWithoutTenantId(includeDecisionRequirementsDefinitionsWithoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionRequirementsDefinitions`: []DecisionRequirementsDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionRequirementsDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decisionRequirementsDefinitionId** | **string** | Filter by decision requirements definition id. | 
 **decisionRequirementsDefinitionIdIn** | **string** | Filter by decision requirements definition ids. | 
 **name** | **string** | Filter by decision requirements definition name. | 
 **nameLike** | **string** | Filter by decision requirements definition names that the parameter is a substring of. | 
 **deploymentId** | **string** | Filter by the id of the deployment a decision requirement definition belongs to. | 
 **key** | **string** | Filter by decision requirements definition key, i.e., the id in the DMN 1.3 XML. Exact match. | 
 **keyLike** | **string** | Filter by decision requirements definition keys that the parameter is a substring of. | 
 **category** | **string** | Filter by decision requirements definition category. Exact match. | 
 **categoryLike** | **string** | Filter by decision requirements definition categories that the parameter is a substring of. | 
 **version** | **int32** | Filter by decision requirements definition version. | 
 **latestVersion** | **bool** | Only include those decision requirements definitions that are latest versions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **resourceName** | **string** | Filter by the name of the decision requirements definition resource. Exact match. | 
 **resourceNameLike** | **string** | Filter by names of those decision requirements definition resources that the parameter is a substring of. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A decision requirements definition must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include decision requirements definitions which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeDecisionRequirementsDefinitionsWithoutTenantId** | **bool** | Include decision requirements definitions which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]DecisionRequirementsDefinitionDto**](DecisionRequirementsDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionRequirementsDefinitionsCount

> CountResultDto GetDecisionRequirementsDefinitionsCount(ctx).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionIdIn(decisionRequirementsDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).Key(key).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDecisionRequirementsDefinitionsWithoutTenantId(includeDecisionRequirementsDefinitionsWithoutTenantId).Execute()

Get Decision Requirements Definition Count



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
	decisionRequirementsDefinitionId := "decisionRequirementsDefinitionId_example" // string | Filter by decision requirements definition id. (optional)
	decisionRequirementsDefinitionIdIn := "decisionRequirementsDefinitionIdIn_example" // string | Filter by decision requirements definition ids. (optional)
	name := "name_example" // string | Filter by decision requirements definition name. (optional)
	nameLike := "nameLike_example" // string | Filter by decision requirements definition names that the parameter is a substring of. (optional)
	deploymentId := "deploymentId_example" // string | Filter by the id of the deployment a decision requirement definition belongs to. (optional)
	key := "key_example" // string | Filter by decision requirements definition key, i.e., the id in the DMN 1.3 XML. Exact match. (optional)
	keyLike := "keyLike_example" // string | Filter by decision requirements definition keys that the parameter is a substring of. (optional)
	category := "category_example" // string | Filter by decision requirements definition category. Exact match. (optional)
	categoryLike := "categoryLike_example" // string | Filter by decision requirements definition categories that the parameter is a substring of. (optional)
	version := int32(56) // int32 | Filter by decision requirements definition version. (optional)
	latestVersion := true // bool | Only include those decision requirements definitions that are latest versions. Value may only be `true`, as `false` is the default behavior. (optional)
	resourceName := "resourceName_example" // string | Filter by the name of the decision requirements definition resource. Exact match. (optional)
	resourceNameLike := "resourceNameLike_example" // string | Filter by names of those decision requirements definition resources that the parameter is a substring of. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A decision requirements definition must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include decision requirements definitions which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	includeDecisionRequirementsDefinitionsWithoutTenantId := true // bool | Include decision requirements definitions which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionsCount(context.Background()).DecisionRequirementsDefinitionId(decisionRequirementsDefinitionId).DecisionRequirementsDefinitionIdIn(decisionRequirementsDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).Key(key).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDecisionRequirementsDefinitionsWithoutTenantId(includeDecisionRequirementsDefinitionsWithoutTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecisionRequirementsDefinitionsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `DecisionRequirementsDefinitionAPI.GetDecisionRequirementsDefinitionsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionRequirementsDefinitionsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decisionRequirementsDefinitionId** | **string** | Filter by decision requirements definition id. | 
 **decisionRequirementsDefinitionIdIn** | **string** | Filter by decision requirements definition ids. | 
 **name** | **string** | Filter by decision requirements definition name. | 
 **nameLike** | **string** | Filter by decision requirements definition names that the parameter is a substring of. | 
 **deploymentId** | **string** | Filter by the id of the deployment a decision requirement definition belongs to. | 
 **key** | **string** | Filter by decision requirements definition key, i.e., the id in the DMN 1.3 XML. Exact match. | 
 **keyLike** | **string** | Filter by decision requirements definition keys that the parameter is a substring of. | 
 **category** | **string** | Filter by decision requirements definition category. Exact match. | 
 **categoryLike** | **string** | Filter by decision requirements definition categories that the parameter is a substring of. | 
 **version** | **int32** | Filter by decision requirements definition version. | 
 **latestVersion** | **bool** | Only include those decision requirements definitions that are latest versions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **resourceName** | **string** | Filter by the name of the decision requirements definition resource. Exact match. | 
 **resourceNameLike** | **string** | Filter by names of those decision requirements definition resources that the parameter is a substring of. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A decision requirements definition must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include decision requirements definitions which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeDecisionRequirementsDefinitionsWithoutTenantId** | **bool** | Include decision requirements definitions which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

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

