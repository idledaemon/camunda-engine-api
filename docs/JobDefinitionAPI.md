# \JobDefinitionAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJobDefinition**](JobDefinitionAPI.md#GetJobDefinition) | **Get** /job-definition/{id} | Get Job Definition
[**GetJobDefinitions**](JobDefinitionAPI.md#GetJobDefinitions) | **Get** /job-definition | Get Job Definitions
[**GetJobDefinitionsCount**](JobDefinitionAPI.md#GetJobDefinitionsCount) | **Get** /job-definition/count | Get Job Definition Count
[**QueryJobDefinitions**](JobDefinitionAPI.md#QueryJobDefinitions) | **Post** /job-definition | Get Job Definitions (POST)
[**QueryJobDefinitionsCount**](JobDefinitionAPI.md#QueryJobDefinitionsCount) | **Post** /job-definition/count | Get Job Definition Count (POST)
[**SetJobPriorityJobDefinition**](JobDefinitionAPI.md#SetJobPriorityJobDefinition) | **Put** /job-definition/{id}/jobPriority | Set Job Definition Priority by Id
[**SetJobRetriesJobDefinition**](JobDefinitionAPI.md#SetJobRetriesJobDefinition) | **Put** /job-definition/{id}/retries | Set Job Retries By Job Definition Id
[**UpdateSuspensionStateJobDefinition**](JobDefinitionAPI.md#UpdateSuspensionStateJobDefinition) | **Put** /job-definition/{id}/suspended | Activate/Suspend Job Definition By Id
[**UpdateSuspensionStateJobDefinitions**](JobDefinitionAPI.md#UpdateSuspensionStateJobDefinitions) | **Put** /job-definition/suspended | Activate/Suspend Job Definitions



## GetJobDefinition

> JobDefinitionDto GetJobDefinition(ctx, id).Execute()

Get Job Definition



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
	id := "id_example" // string | The id of the job definition to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobDefinitionAPI.GetJobDefinition(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDefinitionAPI.GetJobDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobDefinition`: JobDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `JobDefinitionAPI.GetJobDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job definition to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobDefinitionDto**](JobDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobDefinitions

> []JobDefinitionDto GetJobDefinitions(ctx).JobDefinitionId(jobDefinitionId).ActivityIdIn(activityIdIn).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).JobType(jobType).JobConfiguration(jobConfiguration).Active(active).Suspended(suspended).WithOverridingJobPriority(withOverridingJobPriority).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeJobDefinitionsWithoutTenantId(includeJobDefinitionsWithoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Job Definitions



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
	jobDefinitionId := "jobDefinitionId_example" // string | Filter by job definition id. (optional)
	activityIdIn := "activityIdIn_example" // string | Only include job definitions which belong to one of the passed and comma-separated activity ids. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Only include job definitions which exist for the given process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Only include job definitions which exist for the given process definition key. (optional)
	jobType := "jobType_example" // string | Only include job definitions which exist for the given job type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job types. (optional)
	jobConfiguration := "jobConfiguration_example" // string | Only include job definitions which exist for the given job configuration. For example: for timer jobs it is the timer configuration. (optional)
	active := true // bool | Only include active job definitions. Value may only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | Only include suspended job definitions. Value may only be `true`, as `false` is the default behavior. (optional)
	withOverridingJobPriority := true // bool | Only include job definitions that have an overriding job priority defined. The only effective value is `true`. If set to `false`, this filter is not applied. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include job definitions which belong to one of the passed and comma-separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include job definitions which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	includeJobDefinitionsWithoutTenantId := true // bool | Include job definitions which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobDefinitionAPI.GetJobDefinitions(context.Background()).JobDefinitionId(jobDefinitionId).ActivityIdIn(activityIdIn).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).JobType(jobType).JobConfiguration(jobConfiguration).Active(active).Suspended(suspended).WithOverridingJobPriority(withOverridingJobPriority).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeJobDefinitionsWithoutTenantId(includeJobDefinitionsWithoutTenantId).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDefinitionAPI.GetJobDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobDefinitions`: []JobDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `JobDefinitionAPI.GetJobDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobDefinitionId** | **string** | Filter by job definition id. | 
 **activityIdIn** | **string** | Only include job definitions which belong to one of the passed and comma-separated activity ids. | 
 **processDefinitionId** | **string** | Only include job definitions which exist for the given process definition id. | 
 **processDefinitionKey** | **string** | Only include job definitions which exist for the given process definition key. | 
 **jobType** | **string** | Only include job definitions which exist for the given job type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job types. | 
 **jobConfiguration** | **string** | Only include job definitions which exist for the given job configuration. For example: for timer jobs it is the timer configuration. | 
 **active** | **bool** | Only include active job definitions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | Only include suspended job definitions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **withOverridingJobPriority** | **bool** | Only include job definitions that have an overriding job priority defined. The only effective value is &#x60;true&#x60;. If set to &#x60;false&#x60;, this filter is not applied. | 
 **tenantIdIn** | **string** | Only include job definitions which belong to one of the passed and comma-separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include job definitions which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeJobDefinitionsWithoutTenantId** | **bool** | Include job definitions which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]JobDefinitionDto**](JobDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobDefinitionsCount

> CountResultDto GetJobDefinitionsCount(ctx).JobDefinitionId(jobDefinitionId).ActivityIdIn(activityIdIn).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).JobType(jobType).JobConfiguration(jobConfiguration).Active(active).Suspended(suspended).WithOverridingJobPriority(withOverridingJobPriority).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeJobDefinitionsWithoutTenantId(includeJobDefinitionsWithoutTenantId).Execute()

Get Job Definition Count



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
	jobDefinitionId := "jobDefinitionId_example" // string | Filter by job definition id. (optional)
	activityIdIn := "activityIdIn_example" // string | Only include job definitions which belong to one of the passed and comma-separated activity ids. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Only include job definitions which exist for the given process definition id. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Only include job definitions which exist for the given process definition key. (optional)
	jobType := "jobType_example" // string | Only include job definitions which exist for the given job type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job types. (optional)
	jobConfiguration := "jobConfiguration_example" // string | Only include job definitions which exist for the given job configuration. For example: for timer jobs it is the timer configuration. (optional)
	active := true // bool | Only include active job definitions. Value may only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | Only include suspended job definitions. Value may only be `true`, as `false` is the default behavior. (optional)
	withOverridingJobPriority := true // bool | Only include job definitions that have an overriding job priority defined. The only effective value is `true`. If set to `false`, this filter is not applied. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Only include job definitions which belong to one of the passed and comma-separated tenant ids. (optional)
	withoutTenantId := true // bool | Only include job definitions which belong to no tenant. Value may only be `true`, as `false` is the default behavior. (optional)
	includeJobDefinitionsWithoutTenantId := true // bool | Include job definitions which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobDefinitionAPI.GetJobDefinitionsCount(context.Background()).JobDefinitionId(jobDefinitionId).ActivityIdIn(activityIdIn).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).JobType(jobType).JobConfiguration(jobConfiguration).Active(active).Suspended(suspended).WithOverridingJobPriority(withOverridingJobPriority).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeJobDefinitionsWithoutTenantId(includeJobDefinitionsWithoutTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDefinitionAPI.GetJobDefinitionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobDefinitionsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `JobDefinitionAPI.GetJobDefinitionsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobDefinitionsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobDefinitionId** | **string** | Filter by job definition id. | 
 **activityIdIn** | **string** | Only include job definitions which belong to one of the passed and comma-separated activity ids. | 
 **processDefinitionId** | **string** | Only include job definitions which exist for the given process definition id. | 
 **processDefinitionKey** | **string** | Only include job definitions which exist for the given process definition key. | 
 **jobType** | **string** | Only include job definitions which exist for the given job type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job types. | 
 **jobConfiguration** | **string** | Only include job definitions which exist for the given job configuration. For example: for timer jobs it is the timer configuration. | 
 **active** | **bool** | Only include active job definitions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | Only include suspended job definitions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **withOverridingJobPriority** | **bool** | Only include job definitions that have an overriding job priority defined. The only effective value is &#x60;true&#x60;. If set to &#x60;false&#x60;, this filter is not applied. | 
 **tenantIdIn** | **string** | Only include job definitions which belong to one of the passed and comma-separated tenant ids. | 
 **withoutTenantId** | **bool** | Only include job definitions which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **includeJobDefinitionsWithoutTenantId** | **bool** | Include job definitions which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 

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


## QueryJobDefinitions

> []JobDefinitionDto QueryJobDefinitions(ctx).FirstResult(firstResult).MaxResults(maxResults).JobDefinitionQueryDto(jobDefinitionQueryDto).Execute()

Get Job Definitions (POST)



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
	jobDefinitionQueryDto := *openapiclient.NewJobDefinitionQueryDto() // JobDefinitionQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobDefinitionAPI.QueryJobDefinitions(context.Background()).FirstResult(firstResult).MaxResults(maxResults).JobDefinitionQueryDto(jobDefinitionQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDefinitionAPI.QueryJobDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryJobDefinitions`: []JobDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `JobDefinitionAPI.QueryJobDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryJobDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **jobDefinitionQueryDto** | [**JobDefinitionQueryDto**](JobDefinitionQueryDto.md) |  | 

### Return type

[**[]JobDefinitionDto**](JobDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryJobDefinitionsCount

> CountResultDto QueryJobDefinitionsCount(ctx).JobDefinitionQueryDto(jobDefinitionQueryDto).Execute()

Get Job Definition Count (POST)



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
	jobDefinitionQueryDto := *openapiclient.NewJobDefinitionQueryDto() // JobDefinitionQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobDefinitionAPI.QueryJobDefinitionsCount(context.Background()).JobDefinitionQueryDto(jobDefinitionQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDefinitionAPI.QueryJobDefinitionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryJobDefinitionsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `JobDefinitionAPI.QueryJobDefinitionsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryJobDefinitionsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobDefinitionQueryDto** | [**JobDefinitionQueryDto**](JobDefinitionQueryDto.md) |  | 

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


## SetJobPriorityJobDefinition

> SetJobPriorityJobDefinition(ctx, id).JobDefinitionPriorityDto(jobDefinitionPriorityDto).Execute()

Set Job Definition Priority by Id



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
	id := "id_example" // string | The id of the job definition to be updated.
	jobDefinitionPriorityDto := *openapiclient.NewJobDefinitionPriorityDto() // JobDefinitionPriorityDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobDefinitionAPI.SetJobPriorityJobDefinition(context.Background(), id).JobDefinitionPriorityDto(jobDefinitionPriorityDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDefinitionAPI.SetJobPriorityJobDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job definition to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetJobPriorityJobDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobDefinitionPriorityDto** | [**JobDefinitionPriorityDto**](JobDefinitionPriorityDto.md) |  | 

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


## SetJobRetriesJobDefinition

> SetJobRetriesJobDefinition(ctx, id).RetriesDto(retriesDto).Execute()

Set Job Retries By Job Definition Id



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
	id := "id_example" // string | The id of the job definition to be updated.
	retriesDto := *openapiclient.NewRetriesDto() // RetriesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobDefinitionAPI.SetJobRetriesJobDefinition(context.Background(), id).RetriesDto(retriesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDefinitionAPI.SetJobRetriesJobDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job definition to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetJobRetriesJobDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **retriesDto** | [**RetriesDto**](RetriesDto.md) |  | 

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


## UpdateSuspensionStateJobDefinition

> UpdateSuspensionStateJobDefinition(ctx, id).JobDefinitionSuspensionStateDto(jobDefinitionSuspensionStateDto).Execute()

Activate/Suspend Job Definition By Id



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
	id := "id_example" // string | The id of the job definition to activate or suspend.
	jobDefinitionSuspensionStateDto := *openapiclient.NewJobDefinitionSuspensionStateDto() // JobDefinitionSuspensionStateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobDefinitionAPI.UpdateSuspensionStateJobDefinition(context.Background(), id).JobDefinitionSuspensionStateDto(jobDefinitionSuspensionStateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDefinitionAPI.UpdateSuspensionStateJobDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the job definition to activate or suspend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSuspensionStateJobDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobDefinitionSuspensionStateDto** | [**JobDefinitionSuspensionStateDto**](JobDefinitionSuspensionStateDto.md) |  | 

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


## UpdateSuspensionStateJobDefinitions

> UpdateSuspensionStateJobDefinitions(ctx).JobDefinitionsSuspensionStateDto(jobDefinitionsSuspensionStateDto).Execute()

Activate/Suspend Job Definitions



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
	jobDefinitionsSuspensionStateDto := *openapiclient.NewJobDefinitionsSuspensionStateDto() // JobDefinitionsSuspensionStateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobDefinitionAPI.UpdateSuspensionStateJobDefinitions(context.Background()).JobDefinitionsSuspensionStateDto(jobDefinitionsSuspensionStateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDefinitionAPI.UpdateSuspensionStateJobDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSuspensionStateJobDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobDefinitionsSuspensionStateDto** | [**JobDefinitionsSuspensionStateDto**](JobDefinitionsSuspensionStateDto.md) |  | 

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

