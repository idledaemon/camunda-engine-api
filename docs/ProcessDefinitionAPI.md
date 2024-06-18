# \ProcessDefinitionAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcessDefinition**](ProcessDefinitionAPI.md#DeleteProcessDefinition) | **Delete** /process-definition/{id} | Delete
[**DeleteProcessDefinitionsByKey**](ProcessDefinitionAPI.md#DeleteProcessDefinitionsByKey) | **Delete** /process-definition/key/{key} | Delete By Key
[**DeleteProcessDefinitionsByKeyAndTenantId**](ProcessDefinitionAPI.md#DeleteProcessDefinitionsByKeyAndTenantId) | **Delete** /process-definition/key/{key}/tenant-id/{tenant-id} | Delete By Key
[**GetActivityStatistics**](ProcessDefinitionAPI.md#GetActivityStatistics) | **Get** /process-definition/{id}/statistics | Get Activity Instance Statistics
[**GetActivityStatisticsByProcessDefinitionKey**](ProcessDefinitionAPI.md#GetActivityStatisticsByProcessDefinitionKey) | **Get** /process-definition/key/{key}/statistics | Get Activity Instance Statistics
[**GetActivityStatisticsByProcessDefinitionKeyAndTenantId**](ProcessDefinitionAPI.md#GetActivityStatisticsByProcessDefinitionKeyAndTenantId) | **Get** /process-definition/key/{key}/tenant-id/{tenant-id}/statistics | Get Activity Instance Statistics
[**GetDeployedStartForm**](ProcessDefinitionAPI.md#GetDeployedStartForm) | **Get** /process-definition/{id}/deployed-start-form | Get Deployed Start Form
[**GetDeployedStartFormByKey**](ProcessDefinitionAPI.md#GetDeployedStartFormByKey) | **Get** /process-definition/key/{key}/deployed-start-form | Get Deployed Start Form
[**GetDeployedStartFormByKeyAndTenantId**](ProcessDefinitionAPI.md#GetDeployedStartFormByKeyAndTenantId) | **Get** /process-definition/key/{key}/tenant-id/{tenant-id}/deployed-start-form | Get Deployed Start Form
[**GetLatestProcessDefinitionByTenantId**](ProcessDefinitionAPI.md#GetLatestProcessDefinitionByTenantId) | **Get** /process-definition/key/{key}/tenant-id/{tenant-id} | Get
[**GetProcessDefinition**](ProcessDefinitionAPI.md#GetProcessDefinition) | **Get** /process-definition/{id} | Get
[**GetProcessDefinitionBpmn20Xml**](ProcessDefinitionAPI.md#GetProcessDefinitionBpmn20Xml) | **Get** /process-definition/{id}/xml | Get XML
[**GetProcessDefinitionBpmn20XmlByKey**](ProcessDefinitionAPI.md#GetProcessDefinitionBpmn20XmlByKey) | **Get** /process-definition/key/{key}/xml | Get XML
[**GetProcessDefinitionBpmn20XmlByKeyAndTenantId**](ProcessDefinitionAPI.md#GetProcessDefinitionBpmn20XmlByKeyAndTenantId) | **Get** /process-definition/key/{key}/tenant-id/{tenant-id}/xml | Get XML
[**GetProcessDefinitionByKey**](ProcessDefinitionAPI.md#GetProcessDefinitionByKey) | **Get** /process-definition/key/{key} | Get
[**GetProcessDefinitionDiagram**](ProcessDefinitionAPI.md#GetProcessDefinitionDiagram) | **Get** /process-definition/{id}/diagram | Get Diagram
[**GetProcessDefinitionDiagramByKey**](ProcessDefinitionAPI.md#GetProcessDefinitionDiagramByKey) | **Get** /process-definition/key/{key}/diagram | Get Diagram
[**GetProcessDefinitionDiagramByKeyAndTenantId**](ProcessDefinitionAPI.md#GetProcessDefinitionDiagramByKeyAndTenantId) | **Get** /process-definition/key/{key}/tenant-id/{tenant-id}/diagram | Get Diagram
[**GetProcessDefinitionStatistics**](ProcessDefinitionAPI.md#GetProcessDefinitionStatistics) | **Get** /process-definition/statistics | Get Process Instance Statistics
[**GetProcessDefinitions**](ProcessDefinitionAPI.md#GetProcessDefinitions) | **Get** /process-definition | Get List
[**GetProcessDefinitionsCount**](ProcessDefinitionAPI.md#GetProcessDefinitionsCount) | **Get** /process-definition/count | Get List Count
[**GetRenderedStartForm**](ProcessDefinitionAPI.md#GetRenderedStartForm) | **Get** /process-definition/{id}/rendered-form | Get Rendered Start Form
[**GetRenderedStartFormByKey**](ProcessDefinitionAPI.md#GetRenderedStartFormByKey) | **Get** /process-definition/key/{key}/rendered-form | Get Rendered Start Form
[**GetRenderedStartFormByKeyAndTenantId**](ProcessDefinitionAPI.md#GetRenderedStartFormByKeyAndTenantId) | **Get** /process-definition/key/{key}/tenant-id/{tenant-id}/rendered-form | Get Rendered Start Form
[**GetStartForm**](ProcessDefinitionAPI.md#GetStartForm) | **Get** /process-definition/{id}/startForm | Get Start Form Key
[**GetStartFormByKey**](ProcessDefinitionAPI.md#GetStartFormByKey) | **Get** /process-definition/key/{key}/startForm | Get Start Form Key
[**GetStartFormByKeyAndTenantId**](ProcessDefinitionAPI.md#GetStartFormByKeyAndTenantId) | **Get** /process-definition/key/{key}/tenant-id/{tenant-id}/startForm | Get Start Form Key
[**GetStartFormVariables**](ProcessDefinitionAPI.md#GetStartFormVariables) | **Get** /process-definition/{id}/form-variables | Get Start Form Variables
[**GetStartFormVariablesByKey**](ProcessDefinitionAPI.md#GetStartFormVariablesByKey) | **Get** /process-definition/key/{key}/form-variables | Get Start Form Variables
[**GetStartFormVariablesByKeyAndTenantId**](ProcessDefinitionAPI.md#GetStartFormVariablesByKeyAndTenantId) | **Get** /process-definition/key/{key}/tenant-id/{tenant-id}/form-variables | Get Start Form Variables
[**GetStaticCalledProcessDefinitions**](ProcessDefinitionAPI.md#GetStaticCalledProcessDefinitions) | **Get** /process-definition/{id}/static-called-process-definitions | Get Static Called Process Definitions
[**RestartProcessInstance**](ProcessDefinitionAPI.md#RestartProcessInstance) | **Post** /process-definition/{id}/restart | Restart Process Instance
[**RestartProcessInstanceAsyncOperation**](ProcessDefinitionAPI.md#RestartProcessInstanceAsyncOperation) | **Post** /process-definition/{id}/restart-async | Restart Process Instance Async
[**StartProcessInstance**](ProcessDefinitionAPI.md#StartProcessInstance) | **Post** /process-definition/{id}/start | Start Instance
[**StartProcessInstanceByKey**](ProcessDefinitionAPI.md#StartProcessInstanceByKey) | **Post** /process-definition/key/{key}/start | Start Instance
[**StartProcessInstanceByKeyAndTenantId**](ProcessDefinitionAPI.md#StartProcessInstanceByKeyAndTenantId) | **Post** /process-definition/key/{key}/tenant-id/{tenant-id}/start | Start Instance
[**SubmitForm**](ProcessDefinitionAPI.md#SubmitForm) | **Post** /process-definition/{id}/submit-form | Submit Start Form
[**SubmitFormByKey**](ProcessDefinitionAPI.md#SubmitFormByKey) | **Post** /process-definition/key/{key}/submit-form | Submit Start Form
[**SubmitFormByKeyAndTenantId**](ProcessDefinitionAPI.md#SubmitFormByKeyAndTenantId) | **Post** /process-definition/key/{key}/tenant-id/{tenant-id}/submit-form | Submit Start Form
[**UpdateHistoryTimeToLiveByProcessDefinitionId**](ProcessDefinitionAPI.md#UpdateHistoryTimeToLiveByProcessDefinitionId) | **Put** /process-definition/{id}/history-time-to-live | Update History Time to Live
[**UpdateHistoryTimeToLiveByProcessDefinitionKey**](ProcessDefinitionAPI.md#UpdateHistoryTimeToLiveByProcessDefinitionKey) | **Put** /process-definition/key/{key}/history-time-to-live | Update History Time to Live
[**UpdateHistoryTimeToLiveByProcessDefinitionKeyAndTenantId**](ProcessDefinitionAPI.md#UpdateHistoryTimeToLiveByProcessDefinitionKeyAndTenantId) | **Put** /process-definition/key/{key}/tenant-id/{tenant-id}/history-time-to-live | Update History Time to Live
[**UpdateProcessDefinitionSuspensionState**](ProcessDefinitionAPI.md#UpdateProcessDefinitionSuspensionState) | **Put** /process-definition/suspended | Activate/Suspend By Key
[**UpdateProcessDefinitionSuspensionStateById**](ProcessDefinitionAPI.md#UpdateProcessDefinitionSuspensionStateById) | **Put** /process-definition/{id}/suspended | Activate/Suspend By Id
[**UpdateProcessDefinitionSuspensionStateByKey**](ProcessDefinitionAPI.md#UpdateProcessDefinitionSuspensionStateByKey) | **Put** /process-definition/key/{key}/suspended | Activate/Suspend by Id
[**UpdateProcessDefinitionSuspensionStateByKeyAndTenantId**](ProcessDefinitionAPI.md#UpdateProcessDefinitionSuspensionStateByKeyAndTenantId) | **Put** /process-definition/key/{key}/tenant-id/{tenant-id}/suspended | Activate/Suspend by Id



## DeleteProcessDefinition

> DeleteProcessDefinition(ctx, id).Cascade(cascade).SkipCustomListeners(skipCustomListeners).SkipIoMappings(skipIoMappings).Execute()

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
	id := "id_example" // string | The id of the process definition to be deleted.
	cascade := true // bool | `true`, if all process instances, historic process instances and jobs for this process definition should be deleted. (optional)
	skipCustomListeners := true // bool | `true`, if only the built-in ExecutionListeners should be notified with the end event. (optional) (default to false)
	skipIoMappings := true // bool | A boolean value to control whether input/output mappings should be executed during deletion. `true`, if input/output mappings should not be invoked. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessDefinitionAPI.DeleteProcessDefinition(context.Background(), id).Cascade(cascade).SkipCustomListeners(skipCustomListeners).SkipIoMappings(skipIoMappings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.DeleteProcessDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcessDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **bool** | &#x60;true&#x60;, if all process instances, historic process instances and jobs for this process definition should be deleted. | 
 **skipCustomListeners** | **bool** | &#x60;true&#x60;, if only the built-in ExecutionListeners should be notified with the end event. | [default to false]
 **skipIoMappings** | **bool** | A boolean value to control whether input/output mappings should be executed during deletion. &#x60;true&#x60;, if input/output mappings should not be invoked. | [default to false]

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


## DeleteProcessDefinitionsByKey

> DeleteProcessDefinitionsByKey(ctx, key).Cascade(cascade).SkipCustomListeners(skipCustomListeners).SkipIoMappings(skipIoMappings).Execute()

Delete By Key



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
	key := "key_example" // string | The key of the process definitions to be deleted.
	cascade := true // bool | `true`, if all process instances, historic process instances and jobs for this process definition should be deleted. (optional)
	skipCustomListeners := true // bool | `true`, if only the built-in ExecutionListeners should be notified with the end event. (optional) (default to false)
	skipIoMappings := true // bool | A boolean value to control whether input/output mappings should be executed during deletion. `true`, if input/output mappings should not be invoked. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessDefinitionAPI.DeleteProcessDefinitionsByKey(context.Background(), key).Cascade(cascade).SkipCustomListeners(skipCustomListeners).SkipIoMappings(skipIoMappings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.DeleteProcessDefinitionsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definitions to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcessDefinitionsByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **bool** | &#x60;true&#x60;, if all process instances, historic process instances and jobs for this process definition should be deleted. | 
 **skipCustomListeners** | **bool** | &#x60;true&#x60;, if only the built-in ExecutionListeners should be notified with the end event. | [default to false]
 **skipIoMappings** | **bool** | A boolean value to control whether input/output mappings should be executed during deletion. &#x60;true&#x60;, if input/output mappings should not be invoked. | [default to false]

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


## DeleteProcessDefinitionsByKeyAndTenantId

> DeleteProcessDefinitionsByKeyAndTenantId(ctx, key, tenantId).Cascade(cascade).SkipCustomListeners(skipCustomListeners).SkipIoMappings(skipIoMappings).Execute()

Delete By Key



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
	key := "key_example" // string | The key of the process definitions to be deleted.
	tenantId := "tenantId_example" // string | The id of the tenant the process definitions belong to.
	cascade := true // bool | `true`, if all process instances, historic process instances and jobs for this process definition should be deleted. (optional)
	skipCustomListeners := true // bool | `true`, if only the built-in ExecutionListeners should be notified with the end event. (optional) (default to false)
	skipIoMappings := true // bool | A boolean value to control whether input/output mappings should be executed during deletion. `true`, if input/output mappings should not be invoked. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessDefinitionAPI.DeleteProcessDefinitionsByKeyAndTenantId(context.Background(), key, tenantId).Cascade(cascade).SkipCustomListeners(skipCustomListeners).SkipIoMappings(skipIoMappings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.DeleteProcessDefinitionsByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definitions to be deleted. | 
**tenantId** | **string** | The id of the tenant the process definitions belong to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcessDefinitionsByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cascade** | **bool** | &#x60;true&#x60;, if all process instances, historic process instances and jobs for this process definition should be deleted. | 
 **skipCustomListeners** | **bool** | &#x60;true&#x60;, if only the built-in ExecutionListeners should be notified with the end event. | [default to false]
 **skipIoMappings** | **bool** | A boolean value to control whether input/output mappings should be executed during deletion. &#x60;true&#x60;, if input/output mappings should not be invoked. | [default to false]

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


## GetActivityStatistics

> []ActivityStatisticsResultDto GetActivityStatistics(ctx, id).FailedJobs(failedJobs).Incidents(incidents).IncidentsForType(incidentsForType).Execute()

Get Activity Instance Statistics



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
	failedJobs := true // bool | Whether to include the number of failed jobs in the result or not. Valid values are `true` or `false`. (optional)
	incidents := true // bool | Valid values for this property are `true` or `false`. If this property has been set to `true` the result will include the corresponding number of incidents for each occurred incident type. If it is set to `false`, the incidents will not be included in the result. Cannot be used in combination with `incidentsForType`. (optional)
	incidentsForType := "incidentsForType_example" // string | If this property has been set with any incident type (i.e., a string value) the result will only include the number of incidents for the assigned incident type. Cannot be used in combination with `incidents`. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetActivityStatistics(context.Background(), id).FailedJobs(failedJobs).Incidents(incidents).IncidentsForType(incidentsForType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetActivityStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityStatistics`: []ActivityStatisticsResultDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetActivityStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **failedJobs** | **bool** | Whether to include the number of failed jobs in the result or not. Valid values are &#x60;true&#x60; or &#x60;false&#x60;. | 
 **incidents** | **bool** | Valid values for this property are &#x60;true&#x60; or &#x60;false&#x60;. If this property has been set to &#x60;true&#x60; the result will include the corresponding number of incidents for each occurred incident type. If it is set to &#x60;false&#x60;, the incidents will not be included in the result. Cannot be used in combination with &#x60;incidentsForType&#x60;. | 
 **incidentsForType** | **string** | If this property has been set with any incident type (i.e., a string value) the result will only include the number of incidents for the assigned incident type. Cannot be used in combination with &#x60;incidents&#x60;. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 

### Return type

[**[]ActivityStatisticsResultDto**](ActivityStatisticsResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityStatisticsByProcessDefinitionKey

> []ActivityStatisticsResultDto GetActivityStatisticsByProcessDefinitionKey(ctx, key).FailedJobs(failedJobs).Incidents(incidents).IncidentsForType(incidentsForType).Execute()

Get Activity Instance Statistics



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.
	failedJobs := true // bool | Whether to include the number of failed jobs in the result or not. Valid values are `true` or `false`. (optional)
	incidents := true // bool | Valid values for this property are `true` or `false`. If this property has been set to `true` the result will include the corresponding number of incidents for each occurred incident type. If it is set to `false`, the incidents will not be included in the result. Cannot be used in combination with `incidentsForType`. (optional)
	incidentsForType := "incidentsForType_example" // string | If this property has been set with any incident type (i.e., a string value) the result will only include the number of incidents for the assigned incident type. Cannot be used in combination with `incidents`. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetActivityStatisticsByProcessDefinitionKey(context.Background(), key).FailedJobs(failedJobs).Incidents(incidents).IncidentsForType(incidentsForType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetActivityStatisticsByProcessDefinitionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityStatisticsByProcessDefinitionKey`: []ActivityStatisticsResultDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetActivityStatisticsByProcessDefinitionKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityStatisticsByProcessDefinitionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **failedJobs** | **bool** | Whether to include the number of failed jobs in the result or not. Valid values are &#x60;true&#x60; or &#x60;false&#x60;. | 
 **incidents** | **bool** | Valid values for this property are &#x60;true&#x60; or &#x60;false&#x60;. If this property has been set to &#x60;true&#x60; the result will include the corresponding number of incidents for each occurred incident type. If it is set to &#x60;false&#x60;, the incidents will not be included in the result. Cannot be used in combination with &#x60;incidentsForType&#x60;. | 
 **incidentsForType** | **string** | If this property has been set with any incident type (i.e., a string value) the result will only include the number of incidents for the assigned incident type. Cannot be used in combination with &#x60;incidents&#x60;. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 

### Return type

[**[]ActivityStatisticsResultDto**](ActivityStatisticsResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityStatisticsByProcessDefinitionKeyAndTenantId

> []ActivityStatisticsResultDto GetActivityStatisticsByProcessDefinitionKeyAndTenantId(ctx, key, tenantId).FailedJobs(failedJobs).Incidents(incidents).IncidentsForType(incidentsForType).Execute()

Get Activity Instance Statistics



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant the process definition belongs to.
	failedJobs := true // bool | Whether to include the number of failed jobs in the result or not. Valid values are `true` or `false`. (optional)
	incidents := true // bool | Valid values for this property are `true` or `false`. If this property has been set to `true` the result will include the corresponding number of incidents for each occurred incident type. If it is set to `false`, the incidents will not be included in the result. Cannot be used in combination with `incidentsForType`. (optional)
	incidentsForType := "incidentsForType_example" // string | If this property has been set with any incident type (i.e., a string value) the result will only include the number of incidents for the assigned incident type. Cannot be used in combination with `incidents`. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetActivityStatisticsByProcessDefinitionKeyAndTenantId(context.Background(), key, tenantId).FailedJobs(failedJobs).Incidents(incidents).IncidentsForType(incidentsForType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetActivityStatisticsByProcessDefinitionKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityStatisticsByProcessDefinitionKeyAndTenantId`: []ActivityStatisticsResultDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetActivityStatisticsByProcessDefinitionKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 
**tenantId** | **string** | The id of the tenant the process definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityStatisticsByProcessDefinitionKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **failedJobs** | **bool** | Whether to include the number of failed jobs in the result or not. Valid values are &#x60;true&#x60; or &#x60;false&#x60;. | 
 **incidents** | **bool** | Valid values for this property are &#x60;true&#x60; or &#x60;false&#x60;. If this property has been set to &#x60;true&#x60; the result will include the corresponding number of incidents for each occurred incident type. If it is set to &#x60;false&#x60;, the incidents will not be included in the result. Cannot be used in combination with &#x60;incidentsForType&#x60;. | 
 **incidentsForType** | **string** | If this property has been set with any incident type (i.e., a string value) the result will only include the number of incidents for the assigned incident type. Cannot be used in combination with &#x60;incidents&#x60;. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 

### Return type

[**[]ActivityStatisticsResultDto**](ActivityStatisticsResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployedStartForm

> *os.File GetDeployedStartForm(ctx, id).Execute()

Get Deployed Start Form



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
	id := "id_example" // string | The id of the process definition to get the deployed start form for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetDeployedStartForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetDeployedStartForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployedStartForm`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetDeployedStartForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition to get the deployed start form for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployedStartFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xhtml+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployedStartFormByKey

> *os.File GetDeployedStartFormByKey(ctx, key).Execute()

Get Deployed Start Form



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetDeployedStartFormByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetDeployedStartFormByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployedStartFormByKey`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetDeployedStartFormByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployedStartFormByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xhtml+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployedStartFormByKeyAndTenantId

> *os.File GetDeployedStartFormByKeyAndTenantId(ctx, key, tenantId).Execute()

Get Deployed Start Form



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant the process definitions belong to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetDeployedStartFormByKeyAndTenantId(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetDeployedStartFormByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployedStartFormByKeyAndTenantId`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetDeployedStartFormByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 
**tenantId** | **string** | The id of the tenant the process definitions belong to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployedStartFormByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xhtml+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestProcessDefinitionByTenantId

> ProcessDefinitionDto GetLatestProcessDefinitionByTenantId(ctx, key, tenantId).Execute()

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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant the process definition belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetLatestProcessDefinitionByTenantId(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetLatestProcessDefinitionByTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestProcessDefinitionByTenantId`: ProcessDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetLatestProcessDefinitionByTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 
**tenantId** | **string** | The id of the tenant the process definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestProcessDefinitionByTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProcessDefinitionDto**](ProcessDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessDefinition

> ProcessDefinitionDto GetProcessDefinition(ctx, id).Execute()

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
	id := "id_example" // string | The id of the process definition to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetProcessDefinition(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetProcessDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDefinition`: ProcessDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetProcessDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessDefinitionDto**](ProcessDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessDefinitionBpmn20Xml

> ProcessDefinitionDiagramDto GetProcessDefinitionBpmn20Xml(ctx, id).Execute()

Get XML



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
	resp, r, err := apiClient.ProcessDefinitionAPI.GetProcessDefinitionBpmn20Xml(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetProcessDefinitionBpmn20Xml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDefinitionBpmn20Xml`: ProcessDefinitionDiagramDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetProcessDefinitionBpmn20Xml`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDefinitionBpmn20XmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessDefinitionDiagramDto**](ProcessDefinitionDiagramDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessDefinitionBpmn20XmlByKey

> ProcessDefinitionDiagramDto GetProcessDefinitionBpmn20XmlByKey(ctx, key).Execute()

Get XML



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) whose XML should be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetProcessDefinitionBpmn20XmlByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetProcessDefinitionBpmn20XmlByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDefinitionBpmn20XmlByKey`: ProcessDefinitionDiagramDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetProcessDefinitionBpmn20XmlByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) whose XML should be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDefinitionBpmn20XmlByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessDefinitionDiagramDto**](ProcessDefinitionDiagramDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessDefinitionBpmn20XmlByKeyAndTenantId

> ProcessDefinitionDiagramDto GetProcessDefinitionBpmn20XmlByKeyAndTenantId(ctx, key, tenantId).Execute()

Get XML



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) whose XML should be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant the process definition belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetProcessDefinitionBpmn20XmlByKeyAndTenantId(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetProcessDefinitionBpmn20XmlByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDefinitionBpmn20XmlByKeyAndTenantId`: ProcessDefinitionDiagramDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetProcessDefinitionBpmn20XmlByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) whose XML should be retrieved. | 
**tenantId** | **string** | The id of the tenant the process definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDefinitionBpmn20XmlByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProcessDefinitionDiagramDto**](ProcessDefinitionDiagramDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessDefinitionByKey

> ProcessDefinitionDto GetProcessDefinitionByKey(ctx, key).Execute()

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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetProcessDefinitionByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetProcessDefinitionByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDefinitionByKey`: ProcessDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetProcessDefinitionByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDefinitionByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessDefinitionDto**](ProcessDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessDefinitionDiagram

> *os.File GetProcessDefinitionDiagram(ctx, id).Execute()

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
	resp, r, err := apiClient.ProcessDefinitionAPI.GetProcessDefinitionDiagram(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetProcessDefinitionDiagram``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDefinitionDiagram`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetProcessDefinitionDiagram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDefinitionDiagramRequest struct via the builder pattern


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


## GetProcessDefinitionDiagramByKey

> *os.File GetProcessDefinitionDiagramByKey(ctx, key).Execute()

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
	key := "key_example" // string | The key of the process definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetProcessDefinitionDiagramByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetProcessDefinitionDiagramByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDefinitionDiagramByKey`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetProcessDefinitionDiagramByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDefinitionDiagramByKeyRequest struct via the builder pattern


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


## GetProcessDefinitionDiagramByKeyAndTenantId

> *os.File GetProcessDefinitionDiagramByKeyAndTenantId(ctx, key, tenantId).Execute()

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
	key := "key_example" // string | The key of the process definition.
	tenantId := "tenantId_example" // string | The id of the tenant the process definition belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetProcessDefinitionDiagramByKeyAndTenantId(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetProcessDefinitionDiagramByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDefinitionDiagramByKeyAndTenantId`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetProcessDefinitionDiagramByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition. | 
**tenantId** | **string** | The id of the tenant the process definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDefinitionDiagramByKeyAndTenantIdRequest struct via the builder pattern


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


## GetProcessDefinitionStatistics

> []ProcessDefinitionStatisticsResultDto GetProcessDefinitionStatistics(ctx).FailedJobs(failedJobs).Incidents(incidents).IncidentsForType(incidentsForType).RootIncidents(rootIncidents).Execute()

Get Process Instance Statistics



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
	failedJobs := true // bool | Whether to include the number of failed jobs in the result or not. Valid values are `true` or `false`. (optional)
	incidents := true // bool | Valid values for this property are `true` or `false`. If this property has been set to `true` the result will include the corresponding number of incidents for each occurred incident type. If it is set to `false`, the incidents will not be included in the result. Cannot be used in combination with `incidentsForType`. (optional)
	incidentsForType := "incidentsForType_example" // string | If this property has been set with any incident type (i.e., a string value) the result will only include the number of incidents for the assigned incident type. Cannot be used in combination with `incidents`. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	rootIncidents := true // bool | Valid values for this property are `true` or `false`. If this property has been set to `true` the result will include the corresponding number of root incidents for each occurred incident type. If it is set to `false`, the incidents will not be included in the result. Cannot be used in combination with `incidentsForType` or `incidents`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetProcessDefinitionStatistics(context.Background()).FailedJobs(failedJobs).Incidents(incidents).IncidentsForType(incidentsForType).RootIncidents(rootIncidents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetProcessDefinitionStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDefinitionStatistics`: []ProcessDefinitionStatisticsResultDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetProcessDefinitionStatistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDefinitionStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **failedJobs** | **bool** | Whether to include the number of failed jobs in the result or not. Valid values are &#x60;true&#x60; or &#x60;false&#x60;. | 
 **incidents** | **bool** | Valid values for this property are &#x60;true&#x60; or &#x60;false&#x60;. If this property has been set to &#x60;true&#x60; the result will include the corresponding number of incidents for each occurred incident type. If it is set to &#x60;false&#x60;, the incidents will not be included in the result. Cannot be used in combination with &#x60;incidentsForType&#x60;. | 
 **incidentsForType** | **string** | If this property has been set with any incident type (i.e., a string value) the result will only include the number of incidents for the assigned incident type. Cannot be used in combination with &#x60;incidents&#x60;. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **rootIncidents** | **bool** | Valid values for this property are &#x60;true&#x60; or &#x60;false&#x60;. If this property has been set to &#x60;true&#x60; the result will include the corresponding number of root incidents for each occurred incident type. If it is set to &#x60;false&#x60;, the incidents will not be included in the result. Cannot be used in combination with &#x60;incidentsForType&#x60; or &#x60;incidents&#x60;. | 

### Return type

[**[]ProcessDefinitionStatisticsResultDto**](ProcessDefinitionStatisticsResultDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessDefinitions

> []ProcessDefinitionDto GetProcessDefinitions(ctx).ProcessDefinitionId(processDefinitionId).ProcessDefinitionIdIn(processDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).DeployedAfter(deployedAfter).DeployedAt(deployedAt).Key(key).KeysIn(keysIn).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).StartableBy(startableBy).Active(active).Suspended(suspended).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeProcessDefinitionsWithoutTenantId(includeProcessDefinitionsWithoutTenantId).VersionTag(versionTag).VersionTagLike(versionTagLike).WithoutVersionTag(withoutVersionTag).StartableInTasklist(startableInTasklist).NotStartableInTasklist(notStartableInTasklist).StartablePermissionCheck(startablePermissionCheck).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

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
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	processDefinitionIdIn := "processDefinitionIdIn_example" // string | Filter by a comma-separated list of process definition ids. (optional)
	name := "name_example" // string | Filter by process definition name. (optional)
	nameLike := "nameLike_example" // string | Filter by process definition names that the parameter is a substring of. (optional)
	deploymentId := "deploymentId_example" // string | Filter by the deployment the id belongs to. (optional)
	deployedAfter := time.Now() // time.Time | Filter by the deploy time of the deployment the process definition belongs to. Only selects process definitions that have been deployed after (exclusive) a specific time. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.546+0200`. (optional)
	deployedAt := time.Now() // time.Time | Filter by the deploy time of the deployment the process definition belongs to. Only selects process definitions that have been deployed at a specific time (exact match). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.546+0200`. (optional)
	key := "key_example" // string | Filter by process definition key, i.e., the id in the BPMN 2.0 XML. Exact match. (optional)
	keysIn := "keysIn_example" // string | Filter by a comma-separated list of process definition keys. (optional)
	keyLike := "keyLike_example" // string | Filter by process definition keys that the parameter is a substring of. (optional)
	category := "category_example" // string | Filter by process definition category. Exact match. (optional)
	categoryLike := "categoryLike_example" // string | Filter by process definition categories that the parameter is a substring of. (optional)
	version := int32(56) // int32 | Filter by process definition version. (optional)
	latestVersion := true // bool | Only include those process definitions that are latest versions. Value may only be `true`, as `false` is the default behavior. (optional)
	resourceName := "resourceName_example" // string | Filter by the name of the process definition resource. Exact match. (optional)
	resourceNameLike := "resourceNameLike_example" // string | Filter by names of those process definition resources that the parameter is a substring of. (optional)
	startableBy := "startableBy_example" // string | Filter by a user name who is allowed to start the process. (optional)
	active := true // bool | Only include active process definitions. Value may only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | Only include suspended process definitions. Value may only be `true`, as `false` is the default behavior. (optional)
	incidentId := "incidentId_example" // string | Filter by the incident id. (optional)
	incidentType := "incidentType_example" // string | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentMessage := "incidentMessage_example" // string | Filter by the incident message. Exact match. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Filter by the incident message that the parameter is a substring of. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A process definition must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include process definitions which belong to no tenant. Value may only be true, as false is the default behavior. (optional)
	includeProcessDefinitionsWithoutTenantId := true // bool | Include process definitions which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional)
	versionTag := "versionTag_example" // string | Filter by the version tag. (optional)
	versionTagLike := "versionTagLike_example" // string | Filter by the version tag that the parameter is a substring of. (optional)
	withoutVersionTag := true // bool | Only include process definitions without a `versionTag`. (optional)
	startableInTasklist := true // bool | Filter by process definitions which are startable in Tasklist.. (optional)
	notStartableInTasklist := true // bool | Filter by process definitions which are not startable in Tasklist. (optional)
	startablePermissionCheck := true // bool | Filter by process definitions which the user is allowed to start in Tasklist. If the user doesn't have these permissions the result will be empty list. The permissions are: * `CREATE` permission for all Process instances * `CREATE_INSTANCE` and `READ` permission on Process definition level (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetProcessDefinitions(context.Background()).ProcessDefinitionId(processDefinitionId).ProcessDefinitionIdIn(processDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).DeployedAfter(deployedAfter).DeployedAt(deployedAt).Key(key).KeysIn(keysIn).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).StartableBy(startableBy).Active(active).Suspended(suspended).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeProcessDefinitionsWithoutTenantId(includeProcessDefinitionsWithoutTenantId).VersionTag(versionTag).VersionTagLike(versionTagLike).WithoutVersionTag(withoutVersionTag).StartableInTasklist(startableInTasklist).NotStartableInTasklist(notStartableInTasklist).StartablePermissionCheck(startablePermissionCheck).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetProcessDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDefinitions`: []ProcessDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetProcessDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **processDefinitionIdIn** | **string** | Filter by a comma-separated list of process definition ids. | 
 **name** | **string** | Filter by process definition name. | 
 **nameLike** | **string** | Filter by process definition names that the parameter is a substring of. | 
 **deploymentId** | **string** | Filter by the deployment the id belongs to. | 
 **deployedAfter** | **time.Time** | Filter by the deploy time of the deployment the process definition belongs to. Only selects process definitions that have been deployed after (exclusive) a specific time. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.546+0200&#x60;. | 
 **deployedAt** | **time.Time** | Filter by the deploy time of the deployment the process definition belongs to. Only selects process definitions that have been deployed at a specific time (exact match). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.546+0200&#x60;. | 
 **key** | **string** | Filter by process definition key, i.e., the id in the BPMN 2.0 XML. Exact match. | 
 **keysIn** | **string** | Filter by a comma-separated list of process definition keys. | 
 **keyLike** | **string** | Filter by process definition keys that the parameter is a substring of. | 
 **category** | **string** | Filter by process definition category. Exact match. | 
 **categoryLike** | **string** | Filter by process definition categories that the parameter is a substring of. | 
 **version** | **int32** | Filter by process definition version. | 
 **latestVersion** | **bool** | Only include those process definitions that are latest versions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **resourceName** | **string** | Filter by the name of the process definition resource. Exact match. | 
 **resourceNameLike** | **string** | Filter by names of those process definition resources that the parameter is a substring of. | 
 **startableBy** | **string** | Filter by a user name who is allowed to start the process. | 
 **active** | **bool** | Only include active process definitions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | Only include suspended process definitions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **incidentId** | **string** | Filter by the incident id. | 
 **incidentType** | **string** | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentMessage** | **string** | Filter by the incident message. Exact match. | 
 **incidentMessageLike** | **string** | Filter by the incident message that the parameter is a substring of. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A process definition must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include process definitions which belong to no tenant. Value may only be true, as false is the default behavior. | 
 **includeProcessDefinitionsWithoutTenantId** | **bool** | Include process definitions which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **versionTag** | **string** | Filter by the version tag. | 
 **versionTagLike** | **string** | Filter by the version tag that the parameter is a substring of. | 
 **withoutVersionTag** | **bool** | Only include process definitions without a &#x60;versionTag&#x60;. | 
 **startableInTasklist** | **bool** | Filter by process definitions which are startable in Tasklist.. | 
 **notStartableInTasklist** | **bool** | Filter by process definitions which are not startable in Tasklist. | 
 **startablePermissionCheck** | **bool** | Filter by process definitions which the user is allowed to start in Tasklist. If the user doesn&#39;t have these permissions the result will be empty list. The permissions are: * &#x60;CREATE&#x60; permission for all Process instances * &#x60;CREATE_INSTANCE&#x60; and &#x60;READ&#x60; permission on Process definition level | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]ProcessDefinitionDto**](ProcessDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessDefinitionsCount

> CountResultDto GetProcessDefinitionsCount(ctx).ProcessDefinitionId(processDefinitionId).ProcessDefinitionIdIn(processDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).DeployedAfter(deployedAfter).DeployedAt(deployedAt).Key(key).KeysIn(keysIn).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).StartableBy(startableBy).Active(active).Suspended(suspended).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeProcessDefinitionsWithoutTenantId(includeProcessDefinitionsWithoutTenantId).VersionTag(versionTag).VersionTagLike(versionTagLike).WithoutVersionTag(withoutVersionTag).StartableInTasklist(startableInTasklist).NotStartableInTasklist(notStartableInTasklist).StartablePermissionCheck(startablePermissionCheck).Execute()

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
	processDefinitionId := "processDefinitionId_example" // string | Filter by process definition id. (optional)
	processDefinitionIdIn := "processDefinitionIdIn_example" // string | Filter by a comma-separated list of process definition ids. (optional)
	name := "name_example" // string | Filter by process definition name. (optional)
	nameLike := "nameLike_example" // string | Filter by process definition names that the parameter is a substring of. (optional)
	deploymentId := "deploymentId_example" // string | Filter by the deployment the id belongs to. (optional)
	deployedAfter := time.Now() // time.Time | Filter by the deploy time of the deployment the process definition belongs to. Only selects process definitions that have been deployed after (exclusive) a specific time. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.546+0200`. (optional)
	deployedAt := time.Now() // time.Time | Filter by the deploy time of the deployment the process definition belongs to. Only selects process definitions that have been deployed at a specific time (exact match). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.546+0200`. (optional)
	key := "key_example" // string | Filter by process definition key, i.e., the id in the BPMN 2.0 XML. Exact match. (optional)
	keysIn := "keysIn_example" // string | Filter by a comma-separated list of process definition keys. (optional)
	keyLike := "keyLike_example" // string | Filter by process definition keys that the parameter is a substring of. (optional)
	category := "category_example" // string | Filter by process definition category. Exact match. (optional)
	categoryLike := "categoryLike_example" // string | Filter by process definition categories that the parameter is a substring of. (optional)
	version := int32(56) // int32 | Filter by process definition version. (optional)
	latestVersion := true // bool | Only include those process definitions that are latest versions. Value may only be `true`, as `false` is the default behavior. (optional)
	resourceName := "resourceName_example" // string | Filter by the name of the process definition resource. Exact match. (optional)
	resourceNameLike := "resourceNameLike_example" // string | Filter by names of those process definition resources that the parameter is a substring of. (optional)
	startableBy := "startableBy_example" // string | Filter by a user name who is allowed to start the process. (optional)
	active := true // bool | Only include active process definitions. Value may only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | Only include suspended process definitions. Value may only be `true`, as `false` is the default behavior. (optional)
	incidentId := "incidentId_example" // string | Filter by the incident id. (optional)
	incidentType := "incidentType_example" // string | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentMessage := "incidentMessage_example" // string | Filter by the incident message. Exact match. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Filter by the incident message that the parameter is a substring of. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. A process definition must have one of the given tenant ids. (optional)
	withoutTenantId := true // bool | Only include process definitions which belong to no tenant. Value may only be true, as false is the default behavior. (optional)
	includeProcessDefinitionsWithoutTenantId := true // bool | Include process definitions which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior. (optional)
	versionTag := "versionTag_example" // string | Filter by the version tag. (optional)
	versionTagLike := "versionTagLike_example" // string | Filter by the version tag that the parameter is a substring of. (optional)
	withoutVersionTag := true // bool | Only include process definitions without a `versionTag`. (optional)
	startableInTasklist := true // bool | Filter by process definitions which are startable in Tasklist.. (optional)
	notStartableInTasklist := true // bool | Filter by process definitions which are not startable in Tasklist. (optional)
	startablePermissionCheck := true // bool | Filter by process definitions which the user is allowed to start in Tasklist. If the user doesn't have these permissions the result will be empty list. The permissions are: * `CREATE` permission for all Process instances * `CREATE_INSTANCE` and `READ` permission on Process definition level (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetProcessDefinitionsCount(context.Background()).ProcessDefinitionId(processDefinitionId).ProcessDefinitionIdIn(processDefinitionIdIn).Name(name).NameLike(nameLike).DeploymentId(deploymentId).DeployedAfter(deployedAfter).DeployedAt(deployedAt).Key(key).KeysIn(keysIn).KeyLike(keyLike).Category(category).CategoryLike(categoryLike).Version(version).LatestVersion(latestVersion).ResourceName(resourceName).ResourceNameLike(resourceNameLike).StartableBy(startableBy).Active(active).Suspended(suspended).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).WithoutTenantId(withoutTenantId).IncludeProcessDefinitionsWithoutTenantId(includeProcessDefinitionsWithoutTenantId).VersionTag(versionTag).VersionTagLike(versionTagLike).WithoutVersionTag(withoutVersionTag).StartableInTasklist(startableInTasklist).NotStartableInTasklist(notStartableInTasklist).StartablePermissionCheck(startablePermissionCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetProcessDefinitionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessDefinitionsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetProcessDefinitionsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessDefinitionsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processDefinitionId** | **string** | Filter by process definition id. | 
 **processDefinitionIdIn** | **string** | Filter by a comma-separated list of process definition ids. | 
 **name** | **string** | Filter by process definition name. | 
 **nameLike** | **string** | Filter by process definition names that the parameter is a substring of. | 
 **deploymentId** | **string** | Filter by the deployment the id belongs to. | 
 **deployedAfter** | **time.Time** | Filter by the deploy time of the deployment the process definition belongs to. Only selects process definitions that have been deployed after (exclusive) a specific time. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.546+0200&#x60;. | 
 **deployedAt** | **time.Time** | Filter by the deploy time of the deployment the process definition belongs to. Only selects process definitions that have been deployed at a specific time (exact match). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.546+0200&#x60;. | 
 **key** | **string** | Filter by process definition key, i.e., the id in the BPMN 2.0 XML. Exact match. | 
 **keysIn** | **string** | Filter by a comma-separated list of process definition keys. | 
 **keyLike** | **string** | Filter by process definition keys that the parameter is a substring of. | 
 **category** | **string** | Filter by process definition category. Exact match. | 
 **categoryLike** | **string** | Filter by process definition categories that the parameter is a substring of. | 
 **version** | **int32** | Filter by process definition version. | 
 **latestVersion** | **bool** | Only include those process definitions that are latest versions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **resourceName** | **string** | Filter by the name of the process definition resource. Exact match. | 
 **resourceNameLike** | **string** | Filter by names of those process definition resources that the parameter is a substring of. | 
 **startableBy** | **string** | Filter by a user name who is allowed to start the process. | 
 **active** | **bool** | Only include active process definitions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | Only include suspended process definitions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **incidentId** | **string** | Filter by the incident id. | 
 **incidentType** | **string** | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentMessage** | **string** | Filter by the incident message. Exact match. | 
 **incidentMessageLike** | **string** | Filter by the incident message that the parameter is a substring of. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. A process definition must have one of the given tenant ids. | 
 **withoutTenantId** | **bool** | Only include process definitions which belong to no tenant. Value may only be true, as false is the default behavior. | 
 **includeProcessDefinitionsWithoutTenantId** | **bool** | Include process definitions which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **versionTag** | **string** | Filter by the version tag. | 
 **versionTagLike** | **string** | Filter by the version tag that the parameter is a substring of. | 
 **withoutVersionTag** | **bool** | Only include process definitions without a &#x60;versionTag&#x60;. | 
 **startableInTasklist** | **bool** | Filter by process definitions which are startable in Tasklist.. | 
 **notStartableInTasklist** | **bool** | Filter by process definitions which are not startable in Tasklist. | 
 **startablePermissionCheck** | **bool** | Filter by process definitions which the user is allowed to start in Tasklist. If the user doesn&#39;t have these permissions the result will be empty list. The permissions are: * &#x60;CREATE&#x60; permission for all Process instances * &#x60;CREATE_INSTANCE&#x60; and &#x60;READ&#x60; permission on Process definition level | 

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


## GetRenderedStartForm

> *os.File GetRenderedStartForm(ctx, id).Execute()

Get Rendered Start Form



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
	id := "id_example" // string | The id of the process definition to get the rendered start form for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetRenderedStartForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetRenderedStartForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRenderedStartForm`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetRenderedStartForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition to get the rendered start form for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRenderedStartFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xhtml+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRenderedStartFormByKey

> *os.File GetRenderedStartFormByKey(ctx, key).Execute()

Get Rendered Start Form



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetRenderedStartFormByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetRenderedStartFormByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRenderedStartFormByKey`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetRenderedStartFormByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRenderedStartFormByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xhtml+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRenderedStartFormByKeyAndTenantId

> *os.File GetRenderedStartFormByKeyAndTenantId(ctx, key, tenantId).Execute()

Get Rendered Start Form



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant the process definition belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetRenderedStartFormByKeyAndTenantId(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetRenderedStartFormByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRenderedStartFormByKeyAndTenantId`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetRenderedStartFormByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 
**tenantId** | **string** | The id of the tenant the process definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRenderedStartFormByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xhtml+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStartForm

> FormDto GetStartForm(ctx, id).Execute()

Get Start Form Key



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
	id := "id_example" // string | The id of the process definition to get the start form key for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetStartForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetStartForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStartForm`: FormDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetStartForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition to get the start form key for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStartFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormDto**](FormDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStartFormByKey

> FormDto GetStartFormByKey(ctx, key).Execute()

Get Start Form Key



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) for which the form key is to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetStartFormByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetStartFormByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStartFormByKey`: FormDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetStartFormByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) for which the form key is to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStartFormByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormDto**](FormDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStartFormByKeyAndTenantId

> FormDto GetStartFormByKeyAndTenantId(ctx, key, tenantId).Execute()

Get Start Form Key



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) for which the form key is to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant the process definition belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetStartFormByKeyAndTenantId(context.Background(), key, tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetStartFormByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStartFormByKeyAndTenantId`: FormDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetStartFormByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) for which the form key is to be retrieved. | 
**tenantId** | **string** | The id of the tenant the process definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStartFormByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FormDto**](FormDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStartFormVariables

> map[string]VariableValueDto GetStartFormVariables(ctx, id).VariableNames(variableNames).DeserializeValues(deserializeValues).Execute()

Get Start Form Variables



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
	id := "id_example" // string | The id of the process definition to retrieve the variables for.
	variableNames := "variableNames_example" // string | A comma-separated list of variable names. Allows restricting the list of requested variables to the variable names in the list. It is best practice to restrict the list of variables to the variables actually required by the form in order to minimize fetching of data. If the query parameter is ommitted all variables are fetched. If the query parameter contains non-existent variable names, the variable names are ignored. (optional)
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](http://jackson.codehaus.org/) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note**: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetStartFormVariables(context.Background(), id).VariableNames(variableNames).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetStartFormVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStartFormVariables`: map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetStartFormVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition to retrieve the variables for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStartFormVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableNames** | **string** | A comma-separated list of variable names. Allows restricting the list of requested variables to the variable names in the list. It is best practice to restrict the list of variables to the variables actually required by the form in order to minimize fetching of data. If the query parameter is ommitted all variables are fetched. If the query parameter contains non-existent variable names, the variable names are ignored. | 
 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](http://jackson.codehaus.org/) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note**: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. | [default to true]

### Return type

[**map[string]VariableValueDto**](VariableValueDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStartFormVariablesByKey

> map[string]VariableValueDto GetStartFormVariablesByKey(ctx, key).VariableNames(variableNames).DeserializeValues(deserializeValues).Execute()

Get Start Form Variables



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.
	variableNames := "variableNames_example" // string | A comma-separated list of variable names. Allows restricting the list of requested variables to the variable names in the list. It is best practice to restrict the list of variables to the variables actually required by the form in order to minimize fetching of data. If the query parameter is ommitted all variables are fetched. If the query parameter contains non-existent variable names, the variable names are ignored. (optional)
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](http://jackson.codehaus.org/) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note**: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetStartFormVariablesByKey(context.Background(), key).VariableNames(variableNames).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetStartFormVariablesByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStartFormVariablesByKey`: map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetStartFormVariablesByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStartFormVariablesByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableNames** | **string** | A comma-separated list of variable names. Allows restricting the list of requested variables to the variable names in the list. It is best practice to restrict the list of variables to the variables actually required by the form in order to minimize fetching of data. If the query parameter is ommitted all variables are fetched. If the query parameter contains non-existent variable names, the variable names are ignored. | 
 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](http://jackson.codehaus.org/) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note**: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. | [default to true]

### Return type

[**map[string]VariableValueDto**](VariableValueDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStartFormVariablesByKeyAndTenantId

> map[string]VariableValueDto GetStartFormVariablesByKeyAndTenantId(ctx, key, tenantId).VariableNames(variableNames).DeserializeValues(deserializeValues).Execute()

Get Start Form Variables



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant the process definition belongs to.
	variableNames := "variableNames_example" // string | A comma-separated list of variable names. Allows restricting the list of requested variables to the variable names in the list. It is best practice to restrict the list of variables to the variables actually required by the form in order to minimize fetching of data. If the query parameter is ommitted all variables are fetched. If the query parameter contains non-existent variable names, the variable names are ignored. (optional)
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](http://jackson.codehaus.org/) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note**: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.GetStartFormVariablesByKeyAndTenantId(context.Background(), key, tenantId).VariableNames(variableNames).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetStartFormVariablesByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStartFormVariablesByKeyAndTenantId`: map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetStartFormVariablesByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 
**tenantId** | **string** | The id of the tenant the process definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStartFormVariablesByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **variableNames** | **string** | A comma-separated list of variable names. Allows restricting the list of requested variables to the variable names in the list. It is best practice to restrict the list of variables to the variables actually required by the form in order to minimize fetching of data. If the query parameter is ommitted all variables are fetched. If the query parameter contains non-existent variable names, the variable names are ignored. | 
 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default true).  If set to true, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](http://jackson.codehaus.org/) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to false, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note**: While true is the default value for reasons of backward compatibility, we recommend setting this parameter to false when developing web applications that are independent of the Java process applications deployed to the engine. | [default to true]

### Return type

[**map[string]VariableValueDto**](VariableValueDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStaticCalledProcessDefinitions

> []CalledProcessDefinitionDto GetStaticCalledProcessDefinitions(ctx, id).Execute()

Get Static Called Process Definitions



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
	resp, r, err := apiClient.ProcessDefinitionAPI.GetStaticCalledProcessDefinitions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.GetStaticCalledProcessDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStaticCalledProcessDefinitions`: []CalledProcessDefinitionDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.GetStaticCalledProcessDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStaticCalledProcessDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CalledProcessDefinitionDto**](CalledProcessDefinitionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartProcessInstance

> RestartProcessInstance(ctx, id).RestartProcessInstanceDto(restartProcessInstanceDto).Execute()

Restart Process Instance



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
	id := "id_example" // string | The id of the process definition of the process instances to restart.
	restartProcessInstanceDto := *openapiclient.NewRestartProcessInstanceDto() // RestartProcessInstanceDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessDefinitionAPI.RestartProcessInstance(context.Background(), id).RestartProcessInstanceDto(restartProcessInstanceDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.RestartProcessInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition of the process instances to restart. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartProcessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restartProcessInstanceDto** | [**RestartProcessInstanceDto**](RestartProcessInstanceDto.md) |  | 

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


## RestartProcessInstanceAsyncOperation

> BatchDto RestartProcessInstanceAsyncOperation(ctx, id).RestartProcessInstanceDto(restartProcessInstanceDto).Execute()

Restart Process Instance Async



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
	id := "id_example" // string | The id of the process definition of the process instances to restart.
	restartProcessInstanceDto := *openapiclient.NewRestartProcessInstanceDto() // RestartProcessInstanceDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.RestartProcessInstanceAsyncOperation(context.Background(), id).RestartProcessInstanceDto(restartProcessInstanceDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.RestartProcessInstanceAsyncOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestartProcessInstanceAsyncOperation`: BatchDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.RestartProcessInstanceAsyncOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition of the process instances to restart. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartProcessInstanceAsyncOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restartProcessInstanceDto** | [**RestartProcessInstanceDto**](RestartProcessInstanceDto.md) |  | 

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


## StartProcessInstance

> ProcessInstanceWithVariablesDto StartProcessInstance(ctx, id).StartProcessInstanceDto(startProcessInstanceDto).Execute()

Start Instance



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
	id := "id_example" // string | The id of the process definition to be retrieved.
	startProcessInstanceDto := *openapiclient.NewStartProcessInstanceDto() // StartProcessInstanceDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.StartProcessInstance(context.Background(), id).StartProcessInstanceDto(startProcessInstanceDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.StartProcessInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartProcessInstance`: ProcessInstanceWithVariablesDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.StartProcessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartProcessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startProcessInstanceDto** | [**StartProcessInstanceDto**](StartProcessInstanceDto.md) |  | 

### Return type

[**ProcessInstanceWithVariablesDto**](ProcessInstanceWithVariablesDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartProcessInstanceByKey

> ProcessInstanceWithVariablesDto StartProcessInstanceByKey(ctx, key).StartProcessInstanceDto(startProcessInstanceDto).Execute()

Start Instance



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.
	startProcessInstanceDto := *openapiclient.NewStartProcessInstanceDto() // StartProcessInstanceDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.StartProcessInstanceByKey(context.Background(), key).StartProcessInstanceDto(startProcessInstanceDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.StartProcessInstanceByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartProcessInstanceByKey`: ProcessInstanceWithVariablesDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.StartProcessInstanceByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartProcessInstanceByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startProcessInstanceDto** | [**StartProcessInstanceDto**](StartProcessInstanceDto.md) |  | 

### Return type

[**ProcessInstanceWithVariablesDto**](ProcessInstanceWithVariablesDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartProcessInstanceByKeyAndTenantId

> ProcessInstanceWithVariablesDto StartProcessInstanceByKeyAndTenantId(ctx, key, tenantId).StartProcessInstanceDto(startProcessInstanceDto).Execute()

Start Instance



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be retrieved.
	tenantId := "tenantId_example" // string | The id of the tenant the process definition belongs to.
	startProcessInstanceDto := *openapiclient.NewStartProcessInstanceDto() // StartProcessInstanceDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.StartProcessInstanceByKeyAndTenantId(context.Background(), key, tenantId).StartProcessInstanceDto(startProcessInstanceDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.StartProcessInstanceByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartProcessInstanceByKeyAndTenantId`: ProcessInstanceWithVariablesDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.StartProcessInstanceByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be retrieved. | 
**tenantId** | **string** | The id of the tenant the process definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartProcessInstanceByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startProcessInstanceDto** | [**StartProcessInstanceDto**](StartProcessInstanceDto.md) |  | 

### Return type

[**ProcessInstanceWithVariablesDto**](ProcessInstanceWithVariablesDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitForm

> ProcessInstanceDto SubmitForm(ctx, id).StartProcessInstanceFormDto(startProcessInstanceFormDto).Execute()

Submit Start Form



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
	id := "id_example" // string | The id of the process definition to submit the form for.
	startProcessInstanceFormDto := *openapiclient.NewStartProcessInstanceFormDto() // StartProcessInstanceFormDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.SubmitForm(context.Background(), id).StartProcessInstanceFormDto(startProcessInstanceFormDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.SubmitForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitForm`: ProcessInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.SubmitForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition to submit the form for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startProcessInstanceFormDto** | [**StartProcessInstanceFormDto**](StartProcessInstanceFormDto.md) |  | 

### Return type

[**ProcessInstanceDto**](ProcessInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitFormByKey

> ProcessInstanceDto SubmitFormByKey(ctx, key).StartProcessInstanceFormDto(startProcessInstanceFormDto).Execute()

Submit Start Form



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
	key := "key_example" // string | The key of the process definition to submit the form for.
	startProcessInstanceFormDto := *openapiclient.NewStartProcessInstanceFormDto() // StartProcessInstanceFormDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.SubmitFormByKey(context.Background(), key).StartProcessInstanceFormDto(startProcessInstanceFormDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.SubmitFormByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitFormByKey`: ProcessInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.SubmitFormByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition to submit the form for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitFormByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startProcessInstanceFormDto** | [**StartProcessInstanceFormDto**](StartProcessInstanceFormDto.md) |  | 

### Return type

[**ProcessInstanceDto**](ProcessInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitFormByKeyAndTenantId

> ProcessInstanceDto SubmitFormByKeyAndTenantId(ctx, key, tenantId).StartProcessInstanceFormDto(startProcessInstanceFormDto).Execute()

Submit Start Form



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
	key := "key_example" // string | The key of the process definition to submit the form for.
	tenantId := "tenantId_example" // string | The id of the tenant the process definition belongs to.
	startProcessInstanceFormDto := *openapiclient.NewStartProcessInstanceFormDto() // StartProcessInstanceFormDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessDefinitionAPI.SubmitFormByKeyAndTenantId(context.Background(), key, tenantId).StartProcessInstanceFormDto(startProcessInstanceFormDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.SubmitFormByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitFormByKeyAndTenantId`: ProcessInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ProcessDefinitionAPI.SubmitFormByKeyAndTenantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition to submit the form for. | 
**tenantId** | **string** | The id of the tenant the process definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitFormByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startProcessInstanceFormDto** | [**StartProcessInstanceFormDto**](StartProcessInstanceFormDto.md) |  | 

### Return type

[**ProcessInstanceDto**](ProcessInstanceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHistoryTimeToLiveByProcessDefinitionId

> UpdateHistoryTimeToLiveByProcessDefinitionId(ctx, id).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()

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
	id := "id_example" // string | The id of the process definition to change history time to live.
	historyTimeToLiveDto := *openapiclient.NewHistoryTimeToLiveDto() // HistoryTimeToLiveDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessDefinitionAPI.UpdateHistoryTimeToLiveByProcessDefinitionId(context.Background(), id).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.UpdateHistoryTimeToLiveByProcessDefinitionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition to change history time to live. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHistoryTimeToLiveByProcessDefinitionIdRequest struct via the builder pattern


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


## UpdateHistoryTimeToLiveByProcessDefinitionKey

> UpdateHistoryTimeToLiveByProcessDefinitionKey(ctx, key).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()

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
	key := "key_example" // string | The key of the process definition to change history time to live.
	historyTimeToLiveDto := *openapiclient.NewHistoryTimeToLiveDto() // HistoryTimeToLiveDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessDefinitionAPI.UpdateHistoryTimeToLiveByProcessDefinitionKey(context.Background(), key).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.UpdateHistoryTimeToLiveByProcessDefinitionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition to change history time to live. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHistoryTimeToLiveByProcessDefinitionKeyRequest struct via the builder pattern


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


## UpdateHistoryTimeToLiveByProcessDefinitionKeyAndTenantId

> UpdateHistoryTimeToLiveByProcessDefinitionKeyAndTenantId(ctx, key, tenantId).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()

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
	key := "key_example" // string | The key of the process definition to change history time to live.
	tenantId := "tenantId_example" // string | The id of the tenant the process definition belongs to.
	historyTimeToLiveDto := *openapiclient.NewHistoryTimeToLiveDto() // HistoryTimeToLiveDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessDefinitionAPI.UpdateHistoryTimeToLiveByProcessDefinitionKeyAndTenantId(context.Background(), key, tenantId).HistoryTimeToLiveDto(historyTimeToLiveDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.UpdateHistoryTimeToLiveByProcessDefinitionKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition to change history time to live. | 
**tenantId** | **string** | The id of the tenant the process definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHistoryTimeToLiveByProcessDefinitionKeyAndTenantIdRequest struct via the builder pattern


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


## UpdateProcessDefinitionSuspensionState

> UpdateProcessDefinitionSuspensionState(ctx).ProcessDefinitionSuspensionStateDto(processDefinitionSuspensionStateDto).Execute()

Activate/Suspend By Key



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
	processDefinitionSuspensionStateDto := *openapiclient.NewProcessDefinitionSuspensionStateDto() // ProcessDefinitionSuspensionStateDto | **Note**: Unallowed property is `processDefinitionId`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessDefinitionAPI.UpdateProcessDefinitionSuspensionState(context.Background()).ProcessDefinitionSuspensionStateDto(processDefinitionSuspensionStateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.UpdateProcessDefinitionSuspensionState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProcessDefinitionSuspensionStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processDefinitionSuspensionStateDto** | [**ProcessDefinitionSuspensionStateDto**](ProcessDefinitionSuspensionStateDto.md) | **Note**: Unallowed property is &#x60;processDefinitionId&#x60;. | 

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


## UpdateProcessDefinitionSuspensionStateById

> UpdateProcessDefinitionSuspensionStateById(ctx, id).ProcessDefinitionSuspensionStateDto(processDefinitionSuspensionStateDto).Execute()

Activate/Suspend By Id



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
	id := "id_example" // string | The id of the process definition to activate or suspend.
	processDefinitionSuspensionStateDto := *openapiclient.NewProcessDefinitionSuspensionStateDto() // ProcessDefinitionSuspensionStateDto | **Note**: Unallowed properties are `processDefinitionId` and `processDefinitionKey`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessDefinitionAPI.UpdateProcessDefinitionSuspensionStateById(context.Background(), id).ProcessDefinitionSuspensionStateDto(processDefinitionSuspensionStateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.UpdateProcessDefinitionSuspensionStateById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the process definition to activate or suspend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProcessDefinitionSuspensionStateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **processDefinitionSuspensionStateDto** | [**ProcessDefinitionSuspensionStateDto**](ProcessDefinitionSuspensionStateDto.md) | **Note**: Unallowed properties are &#x60;processDefinitionId&#x60; and &#x60;processDefinitionKey&#x60;. | 

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


## UpdateProcessDefinitionSuspensionStateByKey

> UpdateProcessDefinitionSuspensionStateByKey(ctx, key).ProcessDefinitionSuspensionStateDto(processDefinitionSuspensionStateDto).Execute()

Activate/Suspend by Id



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be activated/suspended.
	processDefinitionSuspensionStateDto := *openapiclient.NewProcessDefinitionSuspensionStateDto() // ProcessDefinitionSuspensionStateDto | **Note**: Unallowed properties are `processDefinitionId` and `processDefinitionKey`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessDefinitionAPI.UpdateProcessDefinitionSuspensionStateByKey(context.Background(), key).ProcessDefinitionSuspensionStateDto(processDefinitionSuspensionStateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.UpdateProcessDefinitionSuspensionStateByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be activated/suspended. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProcessDefinitionSuspensionStateByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **processDefinitionSuspensionStateDto** | [**ProcessDefinitionSuspensionStateDto**](ProcessDefinitionSuspensionStateDto.md) | **Note**: Unallowed properties are &#x60;processDefinitionId&#x60; and &#x60;processDefinitionKey&#x60;. | 

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


## UpdateProcessDefinitionSuspensionStateByKeyAndTenantId

> UpdateProcessDefinitionSuspensionStateByKeyAndTenantId(ctx, key, tenantId).ProcessDefinitionSuspensionStateDto(processDefinitionSuspensionStateDto).Execute()

Activate/Suspend by Id



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
	key := "key_example" // string | The key of the process definition (the latest version thereof) to be activated/suspended.
	tenantId := "tenantId_example" // string | The id of the tenant the process definition belongs to.
	processDefinitionSuspensionStateDto := *openapiclient.NewProcessDefinitionSuspensionStateDto() // ProcessDefinitionSuspensionStateDto | **Note**: Unallowed properties are `processDefinitionId` and `processDefinitionKey`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessDefinitionAPI.UpdateProcessDefinitionSuspensionStateByKeyAndTenantId(context.Background(), key, tenantId).ProcessDefinitionSuspensionStateDto(processDefinitionSuspensionStateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessDefinitionAPI.UpdateProcessDefinitionSuspensionStateByKeyAndTenantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the process definition (the latest version thereof) to be activated/suspended. | 
**tenantId** | **string** | The id of the tenant the process definition belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProcessDefinitionSuspensionStateByKeyAndTenantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **processDefinitionSuspensionStateDto** | [**ProcessDefinitionSuspensionStateDto**](ProcessDefinitionSuspensionStateDto.md) | **Note**: Unallowed properties are &#x60;processDefinitionId&#x60; and &#x60;processDefinitionKey&#x60;. | 

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

