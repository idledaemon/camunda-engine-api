# \FilterAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFilter**](FilterAPI.md#CreateFilter) | **Post** /filter/create | Create Filter
[**DeleteFilter**](FilterAPI.md#DeleteFilter) | **Delete** /filter/{id} | Delete Filter
[**ExecuteFilterCount**](FilterAPI.md#ExecuteFilterCount) | **Get** /filter/{id}/count | Execute Filter Count
[**ExecuteFilterList**](FilterAPI.md#ExecuteFilterList) | **Get** /filter/{id}/list | Execute Filter List
[**ExecuteFilterSingleResult**](FilterAPI.md#ExecuteFilterSingleResult) | **Get** /filter/{id}/singleResult | Execute Filter Single Result
[**FilterResourceOptions**](FilterAPI.md#FilterResourceOptions) | **Options** /filter | Filter Resource Options
[**FilterResourceOptionsSingle**](FilterAPI.md#FilterResourceOptionsSingle) | **Options** /filter/{id} | Filter Resource Options
[**GetFilterCount**](FilterAPI.md#GetFilterCount) | **Get** /filter/count | Get Filter Count
[**GetFilterList**](FilterAPI.md#GetFilterList) | **Get** /filter | Get Filters
[**GetSingleFilter**](FilterAPI.md#GetSingleFilter) | **Get** /filter/{id} | Get Single Filter
[**PostExecuteFilterCount**](FilterAPI.md#PostExecuteFilterCount) | **Post** /filter/{id}/count | Execute Filter Count (POST)
[**PostExecuteFilterList**](FilterAPI.md#PostExecuteFilterList) | **Post** /filter/{id}/list | Execute Filter List (POST)
[**PostExecuteFilterSingleResult**](FilterAPI.md#PostExecuteFilterSingleResult) | **Post** /filter/{id}/singleResult | Execute Filter Single Result (POST)
[**UpdateFilter**](FilterAPI.md#UpdateFilter) | **Put** /filter/{id} | Update Filter



## CreateFilter

> FilterDto CreateFilter(ctx).CreateFilterDto(createFilterDto).Execute()

Create Filter



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
	createFilterDto := *openapiclient.NewCreateFilterDto() // CreateFilterDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.CreateFilter(context.Background()).CreateFilterDto(createFilterDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.CreateFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFilter`: FilterDto
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.CreateFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFilterDto** | [**CreateFilterDto**](CreateFilterDto.md) |  | 

### Return type

[**FilterDto**](FilterDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFilter

> DeleteFilter(ctx, id).Execute()

Delete Filter



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
	id := "id_example" // string | The id of the filter to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilterAPI.DeleteFilter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.DeleteFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFilterRequest struct via the builder pattern


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


## ExecuteFilterCount

> CountResultDto ExecuteFilterCount(ctx, id).Execute()

Execute Filter Count



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
	id := "id_example" // string | The id of the filter to execute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.ExecuteFilterCount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.ExecuteFilterCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteFilterCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.ExecuteFilterCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to execute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteFilterCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ExecuteFilterList

> []map[string]interface{} ExecuteFilterList(ctx, id).FirstResult(firstResult).MaxResults(maxResults).Execute()

Execute Filter List



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
	id := "id_example" // string | The id of the filter to execute.
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.ExecuteFilterList(context.Background(), id).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.ExecuteFilterList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteFilterList`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.ExecuteFilterList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to execute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteFilterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteFilterSingleResult

> map[string]interface{} ExecuteFilterSingleResult(ctx, id).Execute()

Execute Filter Single Result



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
	id := "id_example" // string | The id of the filter to execute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.ExecuteFilterSingleResult(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.ExecuteFilterSingleResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteFilterSingleResult`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.ExecuteFilterSingleResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to execute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteFilterSingleResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilterResourceOptions

> ResourceOptionsDto FilterResourceOptions(ctx).Execute()

Filter Resource Options



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
	resp, r, err := apiClient.FilterAPI.FilterResourceOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.FilterResourceOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterResourceOptions`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.FilterResourceOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFilterResourceOptionsRequest struct via the builder pattern


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


## FilterResourceOptionsSingle

> ResourceOptionsDto FilterResourceOptionsSingle(ctx, id).Execute()

Filter Resource Options



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
	id := "id_example" // string | The id of the filter to be checked.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.FilterResourceOptionsSingle(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.FilterResourceOptionsSingle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterResourceOptionsSingle`: ResourceOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.FilterResourceOptionsSingle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to be checked. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilterResourceOptionsSingleRequest struct via the builder pattern


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


## GetFilterCount

> CountResultDto GetFilterCount(ctx).FilterId(filterId).ResourceType(resourceType).Name(name).NameLike(nameLike).Owner(owner).Execute()

Get Filter Count



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
	filterId := "filterId_example" // string | Filter by the id of the filter. (optional)
	resourceType := "resourceType_example" // string | Filter by the resource type of the filter, e.g., `Task`. (optional)
	name := "name_example" // string | Filter by the name of the filter. (optional)
	nameLike := "nameLike_example" // string | Filter by the name that the parameter is a substring of. (optional)
	owner := "owner_example" // string | Filter by the user id of the owner of the filter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.GetFilterCount(context.Background()).FilterId(filterId).ResourceType(resourceType).Name(name).NameLike(nameLike).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.GetFilterCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilterCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.GetFilterCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilterCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterId** | **string** | Filter by the id of the filter. | 
 **resourceType** | **string** | Filter by the resource type of the filter, e.g., &#x60;Task&#x60;. | 
 **name** | **string** | Filter by the name of the filter. | 
 **nameLike** | **string** | Filter by the name that the parameter is a substring of. | 
 **owner** | **string** | Filter by the user id of the owner of the filter. | 

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


## GetFilterList

> []FilterDto GetFilterList(ctx).FilterId(filterId).ResourceType(resourceType).Name(name).NameLike(nameLike).Owner(owner).ItemCount(itemCount).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Filters



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
	filterId := "filterId_example" // string | Filter by the id of the filter. (optional)
	resourceType := "resourceType_example" // string | Filter by the resource type of the filter, e.g., `Task`. (optional)
	name := "name_example" // string | Filter by the name of the filter. (optional)
	nameLike := "nameLike_example" // string | Filter by the name that the parameter is a substring of. (optional)
	owner := "owner_example" // string | Filter by the user id of the owner of the filter. (optional)
	itemCount := true // bool | If set to `true`, each filter result will contain an `itemCount` property with the number of items matched by the filter itself. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.GetFilterList(context.Background()).FilterId(filterId).ResourceType(resourceType).Name(name).NameLike(nameLike).Owner(owner).ItemCount(itemCount).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.GetFilterList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilterList`: []FilterDto
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.GetFilterList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterId** | **string** | Filter by the id of the filter. | 
 **resourceType** | **string** | Filter by the resource type of the filter, e.g., &#x60;Task&#x60;. | 
 **name** | **string** | Filter by the name of the filter. | 
 **nameLike** | **string** | Filter by the name that the parameter is a substring of. | 
 **owner** | **string** | Filter by the user id of the owner of the filter. | 
 **itemCount** | **bool** | If set to &#x60;true&#x60;, each filter result will contain an &#x60;itemCount&#x60; property with the number of items matched by the filter itself. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]FilterDto**](FilterDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleFilter

> FilterDto GetSingleFilter(ctx, id).ItemCount(itemCount).Execute()

Get Single Filter



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
	id := "id_example" // string | The id of the filter to be retrieved.
	itemCount := true // bool | If set to `true`, each filter result will contain an `itemCount` property with the number of items matched by the filter itself. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.GetSingleFilter(context.Background(), id).ItemCount(itemCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.GetSingleFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleFilter`: FilterDto
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.GetSingleFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemCount** | **bool** | If set to &#x60;true&#x60;, each filter result will contain an &#x60;itemCount&#x60; property with the number of items matched by the filter itself. | 

### Return type

[**FilterDto**](FilterDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExecuteFilterCount

> CountResultDto PostExecuteFilterCount(ctx, id).Body(body).Execute()

Execute Filter Count (POST)



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
	id := "id_example" // string | The id of the filter to execute.
	body := map[string]interface{}{ ... } // map[string]interface{} | A JSON object which corresponds to the type of the saved query of the filter, i.e., if the resource type of the filter is Task the body should form a valid task query corresponding to the Task resource. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.PostExecuteFilterCount(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.PostExecuteFilterCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExecuteFilterCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.PostExecuteFilterCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to execute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostExecuteFilterCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | A JSON object which corresponds to the type of the saved query of the filter, i.e., if the resource type of the filter is Task the body should form a valid task query corresponding to the Task resource. | 

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


## PostExecuteFilterList

> []map[string]interface{} PostExecuteFilterList(ctx, id).FirstResult(firstResult).MaxResults(maxResults).Body(body).Execute()

Execute Filter List (POST)



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
	id := "id_example" // string | The id of the filter to execute.
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} | A JSON object which corresponds to the type of the saved query of the filter, i.e., if the resource type of the filter is Task the body should form a valid task query corresponding to the Task resource. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.PostExecuteFilterList(context.Background(), id).FirstResult(firstResult).MaxResults(maxResults).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.PostExecuteFilterList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExecuteFilterList`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.PostExecuteFilterList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to execute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostExecuteFilterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **body** | **map[string]interface{}** | A JSON object which corresponds to the type of the saved query of the filter, i.e., if the resource type of the filter is Task the body should form a valid task query corresponding to the Task resource. | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExecuteFilterSingleResult

> map[string]interface{} PostExecuteFilterSingleResult(ctx, id).Body(body).Execute()

Execute Filter Single Result (POST)



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
	id := "id_example" // string | The id of the filter to execute.
	body := map[string]interface{}{ ... } // map[string]interface{} | A JSON object which corresponds to the type of the saved query of the filter, i.e., if the resource type of the filter is Task the body should form a valid task query corresponding to the Task resource. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.PostExecuteFilterSingleResult(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.PostExecuteFilterSingleResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExecuteFilterSingleResult`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.PostExecuteFilterSingleResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to execute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostExecuteFilterSingleResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | A JSON object which corresponds to the type of the saved query of the filter, i.e., if the resource type of the filter is Task the body should form a valid task query corresponding to the Task resource. | 

### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFilter

> UpdateFilter(ctx, id).CreateFilterDto(createFilterDto).Execute()

Update Filter



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
	id := "id_example" // string | The id of the filter to be updated.
	createFilterDto := *openapiclient.NewCreateFilterDto() // CreateFilterDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilterAPI.UpdateFilter(context.Background(), id).CreateFilterDto(createFilterDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.UpdateFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFilterDto** | [**CreateFilterDto**](CreateFilterDto.md) |  | 

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

