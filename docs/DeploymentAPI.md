# \DeploymentAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeployment**](DeploymentAPI.md#CreateDeployment) | **Post** /deployment/create | Create
[**DeleteDeployment**](DeploymentAPI.md#DeleteDeployment) | **Delete** /deployment/{id} | Delete
[**GetDeployment**](DeploymentAPI.md#GetDeployment) | **Get** /deployment/{id} | Get
[**GetDeploymentResource**](DeploymentAPI.md#GetDeploymentResource) | **Get** /deployment/{id}/resources/{resourceId} | Get Resource
[**GetDeploymentResourceData**](DeploymentAPI.md#GetDeploymentResourceData) | **Get** /deployment/{id}/resources/{resourceId}/data | Get Resource (Binary)
[**GetDeploymentResources**](DeploymentAPI.md#GetDeploymentResources) | **Get** /deployment/{id}/resources | Get Resources
[**GetDeployments**](DeploymentAPI.md#GetDeployments) | **Get** /deployment | Get List
[**GetDeploymentsCount**](DeploymentAPI.md#GetDeploymentsCount) | **Get** /deployment/count | Get List Count
[**GetRegisteredDeployments**](DeploymentAPI.md#GetRegisteredDeployments) | **Get** /deployment/registered | Get Registered Deployments
[**Redeploy**](DeploymentAPI.md#Redeploy) | **Post** /deployment/{id}/redeploy | Redeploy



## CreateDeployment

> DeploymentWithDefinitionsDto CreateDeployment(ctx).TenantId(tenantId).DeploymentSource(deploymentSource).DeployChangedOnly(deployChangedOnly).EnableDuplicateFiltering(enableDuplicateFiltering).DeploymentName(deploymentName).DeploymentActivationTime(deploymentActivationTime).Data(data).Execute()

Create



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
	tenantId := "tenantId_example" // string | The tenant id for the deployment to be created. (optional)
	deploymentSource := "deploymentSource_example" // string | The source for the deployment to be created. (optional)
	deployChangedOnly := true // bool | A flag indicating whether the process engine should perform duplicate checking on a per-resource basis. If set to true, only those resources that have actually changed are deployed. Checks are made against resources included previous deployments of the same name and only against the latest versions of those resources. If set to true, the option enable-duplicate-filtering is overridden and set to true. (optional) (default to false)
	enableDuplicateFiltering := true // bool | A flag indicating whether the process engine should perform duplicate checking for the deployment or not. This allows you to check if a deployment with the same name and the same resouces already exists and if true, not create a new deployment but instead return the existing deployment. The default value is false. (optional) (default to false)
	deploymentName := "deploymentName_example" // string | The name for the deployment to be created. (optional)
	deploymentActivationTime := time.Now() // time.Time | Sets the date on which the process definitions contained in this deployment will be activated. This means that all process definitions will be deployed as usual, but they will be suspended from the start until the given activation date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	data := os.NewFile(1234, "some_file") // *os.File | The binary data to create the deployment resource. It is possible to have more than one form part with different form part names for the binary data to create a deployment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentAPI.CreateDeployment(context.Background()).TenantId(tenantId).DeploymentSource(deploymentSource).DeployChangedOnly(deployChangedOnly).EnableDuplicateFiltering(enableDuplicateFiltering).DeploymentName(deploymentName).DeploymentActivationTime(deploymentActivationTime).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.CreateDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeployment`: DeploymentWithDefinitionsDto
	fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.CreateDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | The tenant id for the deployment to be created. | 
 **deploymentSource** | **string** | The source for the deployment to be created. | 
 **deployChangedOnly** | **bool** | A flag indicating whether the process engine should perform duplicate checking on a per-resource basis. If set to true, only those resources that have actually changed are deployed. Checks are made against resources included previous deployments of the same name and only against the latest versions of those resources. If set to true, the option enable-duplicate-filtering is overridden and set to true. | [default to false]
 **enableDuplicateFiltering** | **bool** | A flag indicating whether the process engine should perform duplicate checking for the deployment or not. This allows you to check if a deployment with the same name and the same resouces already exists and if true, not create a new deployment but instead return the existing deployment. The default value is false. | [default to false]
 **deploymentName** | **string** | The name for the deployment to be created. | 
 **deploymentActivationTime** | **time.Time** | Sets the date on which the process definitions contained in this deployment will be activated. This means that all process definitions will be deployed as usual, but they will be suspended from the start until the given activation date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **data** | ***os.File** | The binary data to create the deployment resource. It is possible to have more than one form part with different form part names for the binary data to create a deployment. | 

### Return type

[**DeploymentWithDefinitionsDto**](DeploymentWithDefinitionsDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeployment

> DeleteDeployment(ctx, id).Cascade(cascade).SkipCustomListeners(skipCustomListeners).SkipIoMappings(skipIoMappings).Execute()

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
	id := "id_example" // string | The id of the deployment to be deleted.
	cascade := true // bool | `true`, if all process instances, historic process instances and jobs for this deployment should be deleted. (optional) (default to false)
	skipCustomListeners := true // bool | `true`, if only the built-in ExecutionListeners should be notified with the end event. (optional) (default to false)
	skipIoMappings := true // bool | `true`, if all input/output mappings should not be invoked. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentAPI.DeleteDeployment(context.Background(), id).Cascade(cascade).SkipCustomListeners(skipCustomListeners).SkipIoMappings(skipIoMappings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.DeleteDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the deployment to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **bool** | &#x60;true&#x60;, if all process instances, historic process instances and jobs for this deployment should be deleted. | [default to false]
 **skipCustomListeners** | **bool** | &#x60;true&#x60;, if only the built-in ExecutionListeners should be notified with the end event. | [default to false]
 **skipIoMappings** | **bool** | &#x60;true&#x60;, if all input/output mappings should not be invoked. | [default to false]

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


## GetDeployment

> DeploymentDto GetDeployment(ctx, id).Execute()

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
	id := "id_example" // string | The id of the deployment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentAPI.GetDeployment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployment`: DeploymentDto
	fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentDto**](DeploymentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentResource

> DeploymentResourceDto GetDeploymentResource(ctx, id, resourceId).Execute()

Get Resource



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
	id := "id_example" // string | The id of the deployment
	resourceId := "resourceId_example" // string | The id of the deployment resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentAPI.GetDeploymentResource(context.Background(), id, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetDeploymentResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentResource`: DeploymentResourceDto
	fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetDeploymentResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the deployment | 
**resourceId** | **string** | The id of the deployment resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeploymentResourceDto**](DeploymentResourceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentResourceData

> *os.File GetDeploymentResourceData(ctx, id, resourceId).Execute()

Get Resource (Binary)



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
	id := "id_example" // string | The id of the deployment.
	resourceId := "resourceId_example" // string | The id of the deployment resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentAPI.GetDeploymentResourceData(context.Background(), id, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetDeploymentResourceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentResourceData`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetDeploymentResourceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the deployment. | 
**resourceId** | **string** | The id of the deployment resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentResourceDataRequest struct via the builder pattern


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


## GetDeploymentResources

> []DeploymentResourceDto GetDeploymentResources(ctx, id).Execute()

Get Resources



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
	id := "id_example" // string | The id of the deployment to retrieve the deployment resources for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentAPI.GetDeploymentResources(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetDeploymentResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentResources`: []DeploymentResourceDto
	fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetDeploymentResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the deployment to retrieve the deployment resources for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeploymentResourceDto**](DeploymentResourceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployments

> []DeploymentDto GetDeployments(ctx).Id(id).Name(name).NameLike(nameLike).Source(source).WithoutSource(withoutSource).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDeploymentsWithoutTenantId(includeDeploymentsWithoutTenantId).After(after).Before(before).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

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
	id := "id_example" // string | Filter by deployment id (optional)
	name := "name_example" // string | Filter by the deployment name. Exact match. (optional)
	nameLike := "nameLike_example" // string | Filter by the deployment name that the parameter is a substring of. The parameter can include the wildcard `%` to express like-strategy such as: starts with (`%`name), ends with (name`%`) or contains (`%`name`%`). (optional)
	source := "source_example" // string | Filter by the deployment source. (optional)
	withoutSource := true // bool | Filter by the deployment source whereby source is equal to `null`. (optional) (default to false)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A deployment must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include deployments which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	includeDeploymentsWithoutTenantId := true // bool | Include deployments which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	after := time.Now() // time.Time | Restricts to all deployments after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	before := time.Now() // time.Time | Restricts to all deployments before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentAPI.GetDeployments(context.Background()).Id(id).Name(name).NameLike(nameLike).Source(source).WithoutSource(withoutSource).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDeploymentsWithoutTenantId(includeDeploymentsWithoutTenantId).After(after).Before(before).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployments`: []DeploymentDto
	fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter by deployment id | 
 **name** | **string** | Filter by the deployment name. Exact match. | 
 **nameLike** | **string** | Filter by the deployment name that the parameter is a substring of. The parameter can include the wildcard &#x60;%&#x60; to express like-strategy such as: starts with (&#x60;%&#x60;name), ends with (name&#x60;%&#x60;) or contains (&#x60;%&#x60;name&#x60;%&#x60;). | 
 **source** | **string** | Filter by the deployment source. | 
 **withoutSource** | **bool** | Filter by the deployment source whereby source is equal to &#x60;null&#x60;. | [default to false]
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A deployment must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include deployments which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **includeDeploymentsWithoutTenantId** | **bool** | Include deployments which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **after** | **time.Time** | Restricts to all deployments after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **before** | **time.Time** | Restricts to all deployments before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]DeploymentDto**](DeploymentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentsCount

> CountResultDto GetDeploymentsCount(ctx).Id(id).Name(name).NameLike(nameLike).Source(source).WithoutSource(withoutSource).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDeploymentsWithoutTenantId(includeDeploymentsWithoutTenantId).After(after).Before(before).Execute()

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
	id := "id_example" // string | Filter by deployment id (optional)
	name := "name_example" // string | Filter by the deployment name. Exact match. (optional)
	nameLike := "nameLike_example" // string | Filter by the deployment name that the parameter is a substring of. The parameter can include the wildcard `%` to express like-strategy such as: starts with (`%`name), ends with (name`%`) or contains (`%`name`%`). (optional)
	source := "source_example" // string | Filter by the deployment source. (optional)
	withoutSource := true // bool | Filter by the deployment source whereby source is equal to `null`. (optional) (default to false)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A deployment must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include deployments which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	includeDeploymentsWithoutTenantId := true // bool | Include deployments which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional) (default to false)
	after := time.Now() // time.Time | Restricts to all deployments after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)
	before := time.Now() // time.Time | Restricts to all deployments before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentAPI.GetDeploymentsCount(context.Background()).Id(id).Name(name).NameLike(nameLike).Source(source).WithoutSource(withoutSource).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeDeploymentsWithoutTenantId(includeDeploymentsWithoutTenantId).After(after).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetDeploymentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetDeploymentsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter by deployment id | 
 **name** | **string** | Filter by the deployment name. Exact match. | 
 **nameLike** | **string** | Filter by the deployment name that the parameter is a substring of. The parameter can include the wildcard &#x60;%&#x60; to express like-strategy such as: starts with (&#x60;%&#x60;name), ends with (name&#x60;%&#x60;) or contains (&#x60;%&#x60;name&#x60;%&#x60;). | 
 **source** | **string** | Filter by the deployment source. | 
 **withoutSource** | **bool** | Filter by the deployment source whereby source is equal to &#x60;null&#x60;. | [default to false]
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A deployment must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include deployments which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **includeDeploymentsWithoutTenantId** | **bool** | Include deployments which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [default to false]
 **after** | **time.Time** | Restricts to all deployments after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 
 **before** | **time.Time** | Restricts to all deployments before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | 

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


## GetRegisteredDeployments

> []string GetRegisteredDeployments(ctx).Execute()

Get Registered Deployments



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
	resp, r, err := apiClient.DeploymentAPI.GetRegisteredDeployments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetRegisteredDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegisteredDeployments`: []string
	fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetRegisteredDeployments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredDeploymentsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Redeploy

> DeploymentWithDefinitionsDto Redeploy(ctx, id).RedeploymentDto(redeploymentDto).Execute()

Redeploy



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
	id := "id_example" // string | The id of the deployment to re-deploy.
	redeploymentDto := *openapiclient.NewRedeploymentDto() // RedeploymentDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentAPI.Redeploy(context.Background(), id).RedeploymentDto(redeploymentDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.Redeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Redeploy`: DeploymentWithDefinitionsDto
	fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.Redeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the deployment to re-deploy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redeploymentDto** | [**RedeploymentDto**](RedeploymentDto.md) |  | 

### Return type

[**DeploymentWithDefinitionsDto**](DeploymentWithDefinitionsDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

