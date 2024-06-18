# HistoricVariableInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **interface{}** | Can be any value - string, number, boolean, array or object.  **Note**: Not every endpoint supports every type. | [optional] 
**Type** | Pointer to **NullableString** | The value type of the variable. | [optional] 
**ValueInfo** | Pointer to **map[string]interface{}** | A JSON object containing additional, value-type-dependent properties. For serialized variables of type Object, the following properties can be provided:  * &#x60;objectTypeName&#x60;: A string representation of the object&#39;s type name. * &#x60;serializationDataFormat&#x60;: The serialization format used to store the variable.  For serialized variables of type File, the following properties can be provided:  * &#x60;filename&#x60;: The name of the file. This is not the variable name but the name that will be used when downloading the file again. * &#x60;mimetype&#x60;: The MIME type of the file that is being uploaded. * &#x60;encoding&#x60;: The encoding of the file that is being uploaded.  The following property can be provided for all value types:  * &#x60;transient&#x60;: Indicates whether the variable should be transient or not. See [documentation](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables#transient-variables) for more informations. (Not applicable for &#x60;decision-definition&#x60;, &#x60; /process-instance/variables-async&#x60;, and &#x60;/migration/executeAsync&#x60; endpoints) | [optional] 
**Id** | Pointer to **NullableString** | The id of the variable instance. | [optional] 
**Name** | Pointer to **NullableString** | The name of the variable instance. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition the variable instance belongs to. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition the variable instance belongs to. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The process instance id the variable instance belongs to. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The execution id the variable instance belongs to. | [optional] 
**ActivityInstanceId** | Pointer to **NullableString** | The id of the activity instance in which the variable is valid. | [optional] 
**CaseDefinitionKey** | Pointer to **NullableString** | The key of the case definition the variable instance belongs to. | [optional] 
**CaseDefinitionId** | Pointer to **NullableString** | The id of the case definition the variable instance belongs to. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | The case instance id the variable instance belongs to. | [optional] 
**CaseExecutionId** | Pointer to **NullableString** | The case execution id the variable instance belongs to. | [optional] 
**TaskId** | Pointer to **NullableString** | The id of the task the variable instance belongs to. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant that this variable instance belongs to. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | An error message in case a Java Serialized Object could not be de-serialized. | [optional] 
**State** | Pointer to **NullableString** | The current state of the variable. Can be &#39;CREATED&#39; or &#39;DELETED&#39;. | [optional] 
**CreateTime** | Pointer to **NullableTime** | The time the variable was inserted. [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the variable should be removed by the History Cleanup job. [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process containing this variable. | [optional] 

## Methods

### NewHistoricVariableInstanceDto

`func NewHistoricVariableInstanceDto() *HistoricVariableInstanceDto`

NewHistoricVariableInstanceDto instantiates a new HistoricVariableInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricVariableInstanceDtoWithDefaults

`func NewHistoricVariableInstanceDtoWithDefaults() *HistoricVariableInstanceDto`

NewHistoricVariableInstanceDtoWithDefaults instantiates a new HistoricVariableInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *HistoricVariableInstanceDto) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HistoricVariableInstanceDto) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HistoricVariableInstanceDto) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *HistoricVariableInstanceDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *HistoricVariableInstanceDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *HistoricVariableInstanceDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetType

`func (o *HistoricVariableInstanceDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoricVariableInstanceDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoricVariableInstanceDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HistoricVariableInstanceDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HistoricVariableInstanceDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HistoricVariableInstanceDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValueInfo

`func (o *HistoricVariableInstanceDto) GetValueInfo() map[string]interface{}`

GetValueInfo returns the ValueInfo field if non-nil, zero value otherwise.

### GetValueInfoOk

`func (o *HistoricVariableInstanceDto) GetValueInfoOk() (*map[string]interface{}, bool)`

GetValueInfoOk returns a tuple with the ValueInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInfo

`func (o *HistoricVariableInstanceDto) SetValueInfo(v map[string]interface{})`

SetValueInfo sets ValueInfo field to given value.

### HasValueInfo

`func (o *HistoricVariableInstanceDto) HasValueInfo() bool`

HasValueInfo returns a boolean if a field has been set.

### GetId

`func (o *HistoricVariableInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricVariableInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricVariableInstanceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricVariableInstanceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricVariableInstanceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricVariableInstanceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HistoricVariableInstanceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HistoricVariableInstanceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HistoricVariableInstanceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HistoricVariableInstanceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HistoricVariableInstanceDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HistoricVariableInstanceDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricVariableInstanceDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricVariableInstanceDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricVariableInstanceDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricVariableInstanceDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricVariableInstanceDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricVariableInstanceDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricVariableInstanceDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricVariableInstanceDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricVariableInstanceDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricVariableInstanceDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricVariableInstanceDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricVariableInstanceDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricVariableInstanceDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricVariableInstanceDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricVariableInstanceDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricVariableInstanceDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricVariableInstanceDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricVariableInstanceDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetExecutionId

`func (o *HistoricVariableInstanceDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HistoricVariableInstanceDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HistoricVariableInstanceDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HistoricVariableInstanceDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HistoricVariableInstanceDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HistoricVariableInstanceDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetActivityInstanceId

`func (o *HistoricVariableInstanceDto) GetActivityInstanceId() string`

GetActivityInstanceId returns the ActivityInstanceId field if non-nil, zero value otherwise.

### GetActivityInstanceIdOk

`func (o *HistoricVariableInstanceDto) GetActivityInstanceIdOk() (*string, bool)`

GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceId

`func (o *HistoricVariableInstanceDto) SetActivityInstanceId(v string)`

SetActivityInstanceId sets ActivityInstanceId field to given value.

### HasActivityInstanceId

`func (o *HistoricVariableInstanceDto) HasActivityInstanceId() bool`

HasActivityInstanceId returns a boolean if a field has been set.

### SetActivityInstanceIdNil

`func (o *HistoricVariableInstanceDto) SetActivityInstanceIdNil(b bool)`

 SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil

### UnsetActivityInstanceId
`func (o *HistoricVariableInstanceDto) UnsetActivityInstanceId()`

UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
### GetCaseDefinitionKey

`func (o *HistoricVariableInstanceDto) GetCaseDefinitionKey() string`

GetCaseDefinitionKey returns the CaseDefinitionKey field if non-nil, zero value otherwise.

### GetCaseDefinitionKeyOk

`func (o *HistoricVariableInstanceDto) GetCaseDefinitionKeyOk() (*string, bool)`

GetCaseDefinitionKeyOk returns a tuple with the CaseDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionKey

`func (o *HistoricVariableInstanceDto) SetCaseDefinitionKey(v string)`

SetCaseDefinitionKey sets CaseDefinitionKey field to given value.

### HasCaseDefinitionKey

`func (o *HistoricVariableInstanceDto) HasCaseDefinitionKey() bool`

HasCaseDefinitionKey returns a boolean if a field has been set.

### SetCaseDefinitionKeyNil

`func (o *HistoricVariableInstanceDto) SetCaseDefinitionKeyNil(b bool)`

 SetCaseDefinitionKeyNil sets the value for CaseDefinitionKey to be an explicit nil

### UnsetCaseDefinitionKey
`func (o *HistoricVariableInstanceDto) UnsetCaseDefinitionKey()`

UnsetCaseDefinitionKey ensures that no value is present for CaseDefinitionKey, not even an explicit nil
### GetCaseDefinitionId

`func (o *HistoricVariableInstanceDto) GetCaseDefinitionId() string`

GetCaseDefinitionId returns the CaseDefinitionId field if non-nil, zero value otherwise.

### GetCaseDefinitionIdOk

`func (o *HistoricVariableInstanceDto) GetCaseDefinitionIdOk() (*string, bool)`

GetCaseDefinitionIdOk returns a tuple with the CaseDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionId

`func (o *HistoricVariableInstanceDto) SetCaseDefinitionId(v string)`

SetCaseDefinitionId sets CaseDefinitionId field to given value.

### HasCaseDefinitionId

`func (o *HistoricVariableInstanceDto) HasCaseDefinitionId() bool`

HasCaseDefinitionId returns a boolean if a field has been set.

### SetCaseDefinitionIdNil

`func (o *HistoricVariableInstanceDto) SetCaseDefinitionIdNil(b bool)`

 SetCaseDefinitionIdNil sets the value for CaseDefinitionId to be an explicit nil

### UnsetCaseDefinitionId
`func (o *HistoricVariableInstanceDto) UnsetCaseDefinitionId()`

UnsetCaseDefinitionId ensures that no value is present for CaseDefinitionId, not even an explicit nil
### GetCaseInstanceId

`func (o *HistoricVariableInstanceDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *HistoricVariableInstanceDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *HistoricVariableInstanceDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *HistoricVariableInstanceDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *HistoricVariableInstanceDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *HistoricVariableInstanceDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetCaseExecutionId

`func (o *HistoricVariableInstanceDto) GetCaseExecutionId() string`

GetCaseExecutionId returns the CaseExecutionId field if non-nil, zero value otherwise.

### GetCaseExecutionIdOk

`func (o *HistoricVariableInstanceDto) GetCaseExecutionIdOk() (*string, bool)`

GetCaseExecutionIdOk returns a tuple with the CaseExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExecutionId

`func (o *HistoricVariableInstanceDto) SetCaseExecutionId(v string)`

SetCaseExecutionId sets CaseExecutionId field to given value.

### HasCaseExecutionId

`func (o *HistoricVariableInstanceDto) HasCaseExecutionId() bool`

HasCaseExecutionId returns a boolean if a field has been set.

### SetCaseExecutionIdNil

`func (o *HistoricVariableInstanceDto) SetCaseExecutionIdNil(b bool)`

 SetCaseExecutionIdNil sets the value for CaseExecutionId to be an explicit nil

### UnsetCaseExecutionId
`func (o *HistoricVariableInstanceDto) UnsetCaseExecutionId()`

UnsetCaseExecutionId ensures that no value is present for CaseExecutionId, not even an explicit nil
### GetTaskId

`func (o *HistoricVariableInstanceDto) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *HistoricVariableInstanceDto) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *HistoricVariableInstanceDto) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *HistoricVariableInstanceDto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *HistoricVariableInstanceDto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *HistoricVariableInstanceDto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetTenantId

`func (o *HistoricVariableInstanceDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricVariableInstanceDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricVariableInstanceDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricVariableInstanceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricVariableInstanceDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricVariableInstanceDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetErrorMessage

`func (o *HistoricVariableInstanceDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *HistoricVariableInstanceDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *HistoricVariableInstanceDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *HistoricVariableInstanceDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *HistoricVariableInstanceDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *HistoricVariableInstanceDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetState

`func (o *HistoricVariableInstanceDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HistoricVariableInstanceDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HistoricVariableInstanceDto) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HistoricVariableInstanceDto) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *HistoricVariableInstanceDto) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *HistoricVariableInstanceDto) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCreateTime

`func (o *HistoricVariableInstanceDto) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HistoricVariableInstanceDto) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HistoricVariableInstanceDto) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HistoricVariableInstanceDto) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### SetCreateTimeNil

`func (o *HistoricVariableInstanceDto) SetCreateTimeNil(b bool)`

 SetCreateTimeNil sets the value for CreateTime to be an explicit nil

### UnsetCreateTime
`func (o *HistoricVariableInstanceDto) UnsetCreateTime()`

UnsetCreateTime ensures that no value is present for CreateTime, not even an explicit nil
### GetRemovalTime

`func (o *HistoricVariableInstanceDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricVariableInstanceDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricVariableInstanceDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricVariableInstanceDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricVariableInstanceDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricVariableInstanceDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricVariableInstanceDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricVariableInstanceDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricVariableInstanceDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricVariableInstanceDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricVariableInstanceDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricVariableInstanceDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


