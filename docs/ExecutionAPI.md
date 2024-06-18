# \ExecutionAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIncident**](ExecutionAPI.md#CreateIncident) | **Post** /execution/{id}/create-incident | Create Incident
[**DeleteLocalExecutionVariable**](ExecutionAPI.md#DeleteLocalExecutionVariable) | **Delete** /execution/{id}/localVariables/{varName} | Delete Local Execution Variable
[**GetExecution**](ExecutionAPI.md#GetExecution) | **Get** /execution/{id} | Get Execution
[**GetExecutions**](ExecutionAPI.md#GetExecutions) | **Get** /execution | Get Executions
[**GetExecutionsCount**](ExecutionAPI.md#GetExecutionsCount) | **Get** /execution/count | Get Execution Count
[**GetLocalExecutionVariable**](ExecutionAPI.md#GetLocalExecutionVariable) | **Get** /execution/{id}/localVariables/{varName} | Get Local Execution Variable
[**GetLocalExecutionVariableBinary**](ExecutionAPI.md#GetLocalExecutionVariableBinary) | **Get** /execution/{id}/localVariables/{varName}/data | Get Local Execution Variable (Binary)
[**GetLocalExecutionVariables**](ExecutionAPI.md#GetLocalExecutionVariables) | **Get** /execution/{id}/localVariables | Get Local Execution Variables
[**GetMessageEventSubscription**](ExecutionAPI.md#GetMessageEventSubscription) | **Get** /execution/{id}/messageSubscriptions/{messageName} | Get Message Event Subscription
[**ModifyLocalExecutionVariables**](ExecutionAPI.md#ModifyLocalExecutionVariables) | **Post** /execution/{id}/localVariables | Update/Delete Local Execution Variables
[**PutLocalExecutionVariable**](ExecutionAPI.md#PutLocalExecutionVariable) | **Put** /execution/{id}/localVariables/{varName} | Put Local Execution Variable
[**QueryExecutions**](ExecutionAPI.md#QueryExecutions) | **Post** /execution | Get Executions (POST)
[**QueryExecutionsCount**](ExecutionAPI.md#QueryExecutionsCount) | **Post** /execution/count | Get Execution Count (POST)
[**SetLocalExecutionVariableBinary**](ExecutionAPI.md#SetLocalExecutionVariableBinary) | **Post** /execution/{id}/localVariables/{varName}/data | Post Local Execution Variable (Binary)
[**SignalExecution**](ExecutionAPI.md#SignalExecution) | **Post** /execution/{id}/signal | Trigger Execution
[**TriggerEvent**](ExecutionAPI.md#TriggerEvent) | **Post** /execution/{id}/messageSubscriptions/{messageName}/trigger | Trigger Message Event Subscription



## CreateIncident

> IncidentDto CreateIncident(ctx, id).CreateIncidentDto(createIncidentDto).Execute()

Create Incident



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
	id := "id_example" // string | The id of the execution to create a new incident for.
	createIncidentDto := *openapiclient.NewCreateIncidentDto() // CreateIncidentDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.CreateIncident(context.Background(), id).CreateIncidentDto(createIncidentDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.CreateIncident``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIncident`: IncidentDto
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.CreateIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution to create a new incident for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIncidentDto** | [**CreateIncidentDto**](CreateIncidentDto.md) |  | 

### Return type

[**IncidentDto**](IncidentDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLocalExecutionVariable

> DeleteLocalExecutionVariable(ctx, id, varName).Execute()

Delete Local Execution Variable



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
	id := "id_example" // string | The id of the execution to delete the variable from.
	varName := "varName_example" // string | The name of the variable to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutionAPI.DeleteLocalExecutionVariable(context.Background(), id, varName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.DeleteLocalExecutionVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution to delete the variable from. | 
**varName** | **string** | The name of the variable to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocalExecutionVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecution

> ExecutionDto GetExecution(ctx, id).Execute()

Get Execution



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
	id := "id_example" // string | The id of the execution to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.GetExecution(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.GetExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecution`: ExecutionDto
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.GetExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExecutionDto**](ExecutionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutions

> []ExecutionDto GetExecutions(ctx).BusinessKey(businessKey).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).ActivityId(activityId).SignalEventSubscriptionName(signalEventSubscriptionName).MessageEventSubscriptionName(messageEventSubscriptionName).Active(active).Suspended(suspended).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).Variables(variables).ProcessVariables(processVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()

Get Executions



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
	businessKey := "businessKey_example" // string | Filter by the business key of the process instances the executions belong to. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the process definition the executions run on. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by the key of the process definition the executions run on. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by the id of the process instance the execution belongs to. (optional)
	activityId := "activityId_example" // string | Filter by the id of the activity the execution currently executes. (optional)
	signalEventSubscriptionName := "signalEventSubscriptionName_example" // string | Select only those executions that expect a signal of the given name. (optional)
	messageEventSubscriptionName := "messageEventSubscriptionName_example" // string | Select only those executions that expect a message of the given name. (optional)
	active := true // bool | Only include active executions. Value may only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | Only include suspended executions. Value may only be `true`, as `false` is the default behavior. (optional)
	incidentId := "incidentId_example" // string | Filter by the incident id. (optional)
	incidentType := "incidentType_example" // string | Filter by the incident type. See the [User Guide](/manual/develop/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentMessage := "incidentMessage_example" // string | Filter by the incident message. Exact match. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Filter by the incident message that the parameter is a substring of. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. An execution must have one of the given tenant ids. (optional)
	variables := "variables_example" // string | Only include executions that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value. **Note:** Values are always treated as `String` objects on server side.  Valid operator values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`. `key` and `value` may not contain underscore or comma characters. (optional)
	processVariables := "processVariables_example" // string | Only include executions that belong to a process instance with variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value. **Note:** Values are always treated as `String` objects on server side.  Valid operator values are: `eq` - equal to; `neq` - not equal to. `key` and `value` may not contain underscore or comma characters. (optional)
	variableNamesIgnoreCase := true // bool | Match all variable names provided in `variables` and `processVariables` case- insensitively. If set to `true` **variableName** and **variablename** are treated as equal. (optional)
	variableValuesIgnoreCase := true // bool | Match all variable values provided in `variables` and `processVariables` case- insensitively. If set to `true` **variableValue** and **variablevalue** are treated as equal. (optional)
	sortBy := "sortBy_example" // string | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. (optional)
	sortOrder := "sortOrder_example" // string | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. (optional)
	firstResult := int32(56) // int32 | Pagination of results. Specifies the index of the first result to return. (optional)
	maxResults := int32(56) // int32 | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.GetExecutions(context.Background()).BusinessKey(businessKey).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).ActivityId(activityId).SignalEventSubscriptionName(signalEventSubscriptionName).MessageEventSubscriptionName(messageEventSubscriptionName).Active(active).Suspended(suspended).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).Variables(variables).ProcessVariables(processVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).SortBy(sortBy).SortOrder(sortOrder).FirstResult(firstResult).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.GetExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecutions`: []ExecutionDto
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.GetExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessKey** | **string** | Filter by the business key of the process instances the executions belong to. | 
 **processDefinitionId** | **string** | Filter by the process definition the executions run on. | 
 **processDefinitionKey** | **string** | Filter by the key of the process definition the executions run on. | 
 **processInstanceId** | **string** | Filter by the id of the process instance the execution belongs to. | 
 **activityId** | **string** | Filter by the id of the activity the execution currently executes. | 
 **signalEventSubscriptionName** | **string** | Select only those executions that expect a signal of the given name. | 
 **messageEventSubscriptionName** | **string** | Select only those executions that expect a message of the given name. | 
 **active** | **bool** | Only include active executions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | Only include suspended executions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **incidentId** | **string** | Filter by the incident id. | 
 **incidentType** | **string** | Filter by the incident type. See the [User Guide](/manual/develop/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentMessage** | **string** | Filter by the incident message. Exact match. | 
 **incidentMessageLike** | **string** | Filter by the incident message that the parameter is a substring of. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. An execution must have one of the given tenant ids. | 
 **variables** | **string** | Only include executions that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. **Note:** Values are always treated as &#x60;String&#x60; objects on server side.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **processVariables** | **string** | Only include executions that belong to a process instance with variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. **Note:** Values are always treated as &#x60;String&#x60; objects on server side.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **variableNamesIgnoreCase** | **bool** | Match all variable names provided in &#x60;variables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | 
 **variableValuesIgnoreCase** | **bool** | Match all variable values provided in &#x60;variables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | 
 **sortBy** | **string** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | 
 **sortOrder** | **string** | Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter. | 
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 

### Return type

[**[]ExecutionDto**](ExecutionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutionsCount

> CountResultDto GetExecutionsCount(ctx).BusinessKey(businessKey).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).ActivityId(activityId).SignalEventSubscriptionName(signalEventSubscriptionName).MessageEventSubscriptionName(messageEventSubscriptionName).Active(active).Suspended(suspended).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).Variables(variables).ProcessVariables(processVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).Execute()

Get Execution Count



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
	businessKey := "businessKey_example" // string | Filter by the business key of the process instances the executions belong to. (optional)
	processDefinitionId := "processDefinitionId_example" // string | Filter by the process definition the executions run on. (optional)
	processDefinitionKey := "processDefinitionKey_example" // string | Filter by the key of the process definition the executions run on. (optional)
	processInstanceId := "processInstanceId_example" // string | Filter by the id of the process instance the execution belongs to. (optional)
	activityId := "activityId_example" // string | Filter by the id of the activity the execution currently executes. (optional)
	signalEventSubscriptionName := "signalEventSubscriptionName_example" // string | Select only those executions that expect a signal of the given name. (optional)
	messageEventSubscriptionName := "messageEventSubscriptionName_example" // string | Select only those executions that expect a message of the given name. (optional)
	active := true // bool | Only include active executions. Value may only be `true`, as `false` is the default behavior. (optional)
	suspended := true // bool | Only include suspended executions. Value may only be `true`, as `false` is the default behavior. (optional)
	incidentId := "incidentId_example" // string | Filter by the incident id. (optional)
	incidentType := "incidentType_example" // string | Filter by the incident type. See the [User Guide](/manual/develop/user-guide/process-engine/incidents/#incident-types) for a list of incident types. (optional)
	incidentMessage := "incidentMessage_example" // string | Filter by the incident message. Exact match. (optional)
	incidentMessageLike := "incidentMessageLike_example" // string | Filter by the incident message that the parameter is a substring of. (optional)
	tenantIdIn := "tenantIdIn_example" // string | Filter by a comma-separated list of tenant ids. An execution must have one of the given tenant ids. (optional)
	variables := "variables_example" // string | Only include executions that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value. **Note:** Values are always treated as `String` objects on server side.  Valid operator values are: `eq` - equal to; `neq` - not equal to; `gt` - greater than; `gteq` - greater than or equal to; `lt` - lower than; `lteq` - lower than or equal to; `like`. `key` and `value` may not contain underscore or comma characters. (optional)
	processVariables := "processVariables_example" // string | Only include executions that belong to a process instance with variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form `key_operator_value`. `key` is the variable name, `operator` is the comparison operator to be used and `value` the variable value. **Note:** Values are always treated as `String` objects on server side.  Valid operator values are: `eq` - equal to; `neq` - not equal to. `key` and `value` may not contain underscore or comma characters. (optional)
	variableNamesIgnoreCase := true // bool | Match all variable names provided in `variables` and `processVariables` case- insensitively. If set to `true` **variableName** and **variablename** are treated as equal. (optional)
	variableValuesIgnoreCase := true // bool | Match all variable values provided in `variables` and `processVariables` case- insensitively. If set to `true` **variableValue** and **variablevalue** are treated as equal. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.GetExecutionsCount(context.Background()).BusinessKey(businessKey).ProcessDefinitionId(processDefinitionId).ProcessDefinitionKey(processDefinitionKey).ProcessInstanceId(processInstanceId).ActivityId(activityId).SignalEventSubscriptionName(signalEventSubscriptionName).MessageEventSubscriptionName(messageEventSubscriptionName).Active(active).Suspended(suspended).IncidentId(incidentId).IncidentType(incidentType).IncidentMessage(incidentMessage).IncidentMessageLike(incidentMessageLike).TenantIdIn(tenantIdIn).Variables(variables).ProcessVariables(processVariables).VariableNamesIgnoreCase(variableNamesIgnoreCase).VariableValuesIgnoreCase(variableValuesIgnoreCase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.GetExecutionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecutionsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.GetExecutionsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessKey** | **string** | Filter by the business key of the process instances the executions belong to. | 
 **processDefinitionId** | **string** | Filter by the process definition the executions run on. | 
 **processDefinitionKey** | **string** | Filter by the key of the process definition the executions run on. | 
 **processInstanceId** | **string** | Filter by the id of the process instance the execution belongs to. | 
 **activityId** | **string** | Filter by the id of the activity the execution currently executes. | 
 **signalEventSubscriptionName** | **string** | Select only those executions that expect a signal of the given name. | 
 **messageEventSubscriptionName** | **string** | Select only those executions that expect a message of the given name. | 
 **active** | **bool** | Only include active executions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **suspended** | **bool** | Only include suspended executions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | 
 **incidentId** | **string** | Filter by the incident id. | 
 **incidentType** | **string** | Filter by the incident type. See the [User Guide](/manual/develop/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | 
 **incidentMessage** | **string** | Filter by the incident message. Exact match. | 
 **incidentMessageLike** | **string** | Filter by the incident message that the parameter is a substring of. | 
 **tenantIdIn** | **string** | Filter by a comma-separated list of tenant ids. An execution must have one of the given tenant ids. | 
 **variables** | **string** | Only include executions that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. **Note:** Values are always treated as &#x60;String&#x60; objects on server side.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **processVariables** | **string** | Only include executions that belong to a process instance with variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. **Note:** Values are always treated as &#x60;String&#x60; objects on server side.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | 
 **variableNamesIgnoreCase** | **bool** | Match all variable names provided in &#x60;variables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | 
 **variableValuesIgnoreCase** | **bool** | Match all variable values provided in &#x60;variables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | 

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


## GetLocalExecutionVariable

> VariableValueDto GetLocalExecutionVariable(ctx, id, varName).DeserializeValue(deserializeValue).Execute()

Get Local Execution Variable



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
	id := "id_example" // string | The id of the execution to retrieve the variable from.
	varName := "varName_example" // string | The name of the variable to get.
	deserializeValue := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath. If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.GetLocalExecutionVariable(context.Background(), id, varName).DeserializeValue(deserializeValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.GetLocalExecutionVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalExecutionVariable`: VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.GetLocalExecutionVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution to retrieve the variable from. | 
**varName** | **string** | The name of the variable to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalExecutionVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deserializeValue** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath. If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | 

### Return type

[**VariableValueDto**](VariableValueDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalExecutionVariableBinary

> *os.File GetLocalExecutionVariableBinary(ctx, id, varName).Execute()

Get Local Execution Variable (Binary)



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
	id := "id_example" // string | The id of the execution to retrieve the variable from.
	varName := "varName_example" // string | The name of the variable to get.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.GetLocalExecutionVariableBinary(context.Background(), id, varName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.GetLocalExecutionVariableBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalExecutionVariableBinary`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.GetLocalExecutionVariableBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution to retrieve the variable from. | 
**varName** | **string** | The name of the variable to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalExecutionVariableBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalExecutionVariables

> map[string]VariableValueDto GetLocalExecutionVariables(ctx, id).DeserializeValues(deserializeValues).Execute()

Get Local Execution Variables



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
	id := "id_example" // string | The id of the execution to retrieve the variables from.
	deserializeValues := true // bool | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default `true`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While `true` is the default value for reasons of backward compatibility, we recommend setting this parameter to `false` when developing web applications that are independent of the Java process applications deployed to the engine. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.GetLocalExecutionVariables(context.Background(), id).DeserializeValues(deserializeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.GetLocalExecutionVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalExecutionVariables`: map[string]VariableValueDto
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.GetLocalExecutionVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution to retrieve the variables from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalExecutionVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deserializeValues** | **bool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;true&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.  **Note:** While &#x60;true&#x60; is the default value for reasons of backward compatibility, we recommend setting this parameter to &#x60;false&#x60; when developing web applications that are independent of the Java process applications deployed to the engine. | 

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


## GetMessageEventSubscription

> EventSubscriptionDto GetMessageEventSubscription(ctx, id, messageName).Execute()

Get Message Event Subscription



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
	id := "id_example" // string | The id of the execution that holds the subscription.
	messageName := "messageName_example" // string | The name of the message that the subscription corresponds to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.GetMessageEventSubscription(context.Background(), id, messageName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.GetMessageEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageEventSubscription`: EventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.GetMessageEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution that holds the subscription. | 
**messageName** | **string** | The name of the message that the subscription corresponds to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EventSubscriptionDto**](EventSubscriptionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyLocalExecutionVariables

> ModifyLocalExecutionVariables(ctx, id).PatchVariablesDto(patchVariablesDto).Execute()

Update/Delete Local Execution Variables



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
	id := "id_example" // string | The id of the execution to set variables for.
	patchVariablesDto := *openapiclient.NewPatchVariablesDto() // PatchVariablesDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutionAPI.ModifyLocalExecutionVariables(context.Background(), id).PatchVariablesDto(patchVariablesDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.ModifyLocalExecutionVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution to set variables for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLocalExecutionVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchVariablesDto** | [**PatchVariablesDto**](PatchVariablesDto.md) |  | 

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


## PutLocalExecutionVariable

> PutLocalExecutionVariable(ctx, id, varName).VariableValueDto(variableValueDto).Execute()

Put Local Execution Variable



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
	id := "id_example" // string | The id of the execution to set the variable for.
	varName := "varName_example" // string | The name of the variable to set.
	variableValueDto := *openapiclient.NewVariableValueDto() // VariableValueDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutionAPI.PutLocalExecutionVariable(context.Background(), id, varName).VariableValueDto(variableValueDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.PutLocalExecutionVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution to set the variable for. | 
**varName** | **string** | The name of the variable to set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLocalExecutionVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **variableValueDto** | [**VariableValueDto**](VariableValueDto.md) |  | 

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


## QueryExecutions

> []ExecutionDto QueryExecutions(ctx).FirstResult(firstResult).MaxResults(maxResults).ExecutionQueryDto(executionQueryDto).Execute()

Get Executions (POST)



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
	executionQueryDto := *openapiclient.NewExecutionQueryDto() // ExecutionQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.QueryExecutions(context.Background()).FirstResult(firstResult).MaxResults(maxResults).ExecutionQueryDto(executionQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.QueryExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryExecutions`: []ExecutionDto
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.QueryExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstResult** | **int32** | Pagination of results. Specifies the index of the first result to return. | 
 **maxResults** | **int32** | Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left. | 
 **executionQueryDto** | [**ExecutionQueryDto**](ExecutionQueryDto.md) |  | 

### Return type

[**[]ExecutionDto**](ExecutionDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryExecutionsCount

> CountResultDto QueryExecutionsCount(ctx).ExecutionQueryDto(executionQueryDto).Execute()

Get Execution Count (POST)



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
	executionQueryDto := *openapiclient.NewExecutionQueryDto() // ExecutionQueryDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.QueryExecutionsCount(context.Background()).ExecutionQueryDto(executionQueryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.QueryExecutionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryExecutionsCount`: CountResultDto
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.QueryExecutionsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryExecutionsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionQueryDto** | [**ExecutionQueryDto**](ExecutionQueryDto.md) |  | 

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


## SetLocalExecutionVariableBinary

> SetLocalExecutionVariableBinary(ctx, id, varName).Data(data).ValueType(valueType).Execute()

Post Local Execution Variable (Binary)



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
	id := "id_example" // string | The id of the execution to set the variable for.
	varName := "varName_example" // string | The name of the variable to set.
	data := os.NewFile(1234, "some_file") // *os.File | The binary data to be set. For File variables, this multipart can contain the filename, binary value and MIME type of the file variable to be set Only the filename is mandatory. (optional)
	valueType := "valueType_example" // string | The name of the variable type. Either Bytes for a byte array variable or File for a file variable. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutionAPI.SetLocalExecutionVariableBinary(context.Background(), id, varName).Data(data).ValueType(valueType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.SetLocalExecutionVariableBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution to set the variable for. | 
**varName** | **string** | The name of the variable to set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLocalExecutionVariableBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | ***os.File** | The binary data to be set. For File variables, this multipart can contain the filename, binary value and MIME type of the file variable to be set Only the filename is mandatory. | 
 **valueType** | **string** | The name of the variable type. Either Bytes for a byte array variable or File for a file variable. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignalExecution

> SignalExecution(ctx, id).ExecutionTriggerDto(executionTriggerDto).Execute()

Trigger Execution



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
	id := "id_example" // string | The id of the execution to signal.
	executionTriggerDto := *openapiclient.NewExecutionTriggerDto() // ExecutionTriggerDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutionAPI.SignalExecution(context.Background(), id).ExecutionTriggerDto(executionTriggerDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.SignalExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution to signal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignalExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **executionTriggerDto** | [**ExecutionTriggerDto**](ExecutionTriggerDto.md) |  | 

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


## TriggerEvent

> TriggerEvent(ctx, id, messageName).ExecutionTriggerDto(executionTriggerDto).Execute()

Trigger Message Event Subscription



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
	id := "id_example" // string | The id of the execution to submit the message to.
	messageName := "messageName_example" // string | The name of the message that the addressed subscription corresponds to.
	executionTriggerDto := *openapiclient.NewExecutionTriggerDto() // ExecutionTriggerDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutionAPI.TriggerEvent(context.Background(), id, messageName).ExecutionTriggerDto(executionTriggerDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.TriggerEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the execution to submit the message to. | 
**messageName** | **string** | The name of the message that the addressed subscription corresponds to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **executionTriggerDto** | [**ExecutionTriggerDto**](ExecutionTriggerDto.md) |  | 

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

