# VariableInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **interface{}** | Can be any value - string, number, boolean, array or object.  **Note**: Not every endpoint supports every type. | [optional] 
**Type** | Pointer to **NullableString** | The value type of the variable. | [optional] 
**ValueInfo** | Pointer to **map[string]interface{}** | A JSON object containing additional, value-type-dependent properties. For serialized variables of type Object, the following properties can be provided:  * &#x60;objectTypeName&#x60;: A string representation of the object&#39;s type name. * &#x60;serializationDataFormat&#x60;: The serialization format used to store the variable.  For serialized variables of type File, the following properties can be provided:  * &#x60;filename&#x60;: The name of the file. This is not the variable name but the name that will be used when downloading the file again. * &#x60;mimetype&#x60;: The MIME type of the file that is being uploaded. * &#x60;encoding&#x60;: The encoding of the file that is being uploaded.  The following property can be provided for all value types:  * &#x60;transient&#x60;: Indicates whether the variable should be transient or not. See [documentation](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables#transient-variables) for more informations. (Not applicable for &#x60;decision-definition&#x60;, &#x60; /process-instance/variables-async&#x60;, and &#x60;/migration/executeAsync&#x60; endpoints) | [optional] 
**Id** | Pointer to **NullableString** | The id of the variable instance. | [optional] 
**Name** | Pointer to **NullableString** | The name of the variable instance. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition that this variable instance belongs to. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance that this variable instance belongs to. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The id of the execution that this variable instance belongs to. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | The id of the case instance that this variable instance belongs to. | [optional] 
**CaseExecutionId** | Pointer to **NullableString** | The id of the case execution that this variable instance belongs to. | [optional] 
**TaskId** | Pointer to **NullableString** | The id of the task that this variable instance belongs to. | [optional] 
**BatchId** | Pointer to **NullableString** | The id of the batch that this variable instance belongs to.&lt; | [optional] 
**ActivityInstanceId** | Pointer to **NullableString** | The id of the activity instance that this variable instance belongs to. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant that this variable instance belongs to. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | An error message in case a Java Serialized Object could not be de-serialized. | [optional] 

## Methods

### NewVariableInstanceDto

`func NewVariableInstanceDto() *VariableInstanceDto`

NewVariableInstanceDto instantiates a new VariableInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableInstanceDtoWithDefaults

`func NewVariableInstanceDtoWithDefaults() *VariableInstanceDto`

NewVariableInstanceDtoWithDefaults instantiates a new VariableInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *VariableInstanceDto) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableInstanceDto) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableInstanceDto) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableInstanceDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *VariableInstanceDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *VariableInstanceDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetType

`func (o *VariableInstanceDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VariableInstanceDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VariableInstanceDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VariableInstanceDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *VariableInstanceDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VariableInstanceDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValueInfo

`func (o *VariableInstanceDto) GetValueInfo() map[string]interface{}`

GetValueInfo returns the ValueInfo field if non-nil, zero value otherwise.

### GetValueInfoOk

`func (o *VariableInstanceDto) GetValueInfoOk() (*map[string]interface{}, bool)`

GetValueInfoOk returns a tuple with the ValueInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInfo

`func (o *VariableInstanceDto) SetValueInfo(v map[string]interface{})`

SetValueInfo sets ValueInfo field to given value.

### HasValueInfo

`func (o *VariableInstanceDto) HasValueInfo() bool`

HasValueInfo returns a boolean if a field has been set.

### GetId

`func (o *VariableInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariableInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariableInstanceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VariableInstanceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VariableInstanceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VariableInstanceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *VariableInstanceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableInstanceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableInstanceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VariableInstanceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VariableInstanceDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VariableInstanceDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProcessDefinitionId

`func (o *VariableInstanceDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *VariableInstanceDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *VariableInstanceDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *VariableInstanceDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *VariableInstanceDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *VariableInstanceDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessInstanceId

`func (o *VariableInstanceDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *VariableInstanceDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *VariableInstanceDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *VariableInstanceDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *VariableInstanceDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *VariableInstanceDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetExecutionId

`func (o *VariableInstanceDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *VariableInstanceDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *VariableInstanceDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *VariableInstanceDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *VariableInstanceDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *VariableInstanceDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetCaseInstanceId

`func (o *VariableInstanceDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *VariableInstanceDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *VariableInstanceDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *VariableInstanceDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *VariableInstanceDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *VariableInstanceDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetCaseExecutionId

`func (o *VariableInstanceDto) GetCaseExecutionId() string`

GetCaseExecutionId returns the CaseExecutionId field if non-nil, zero value otherwise.

### GetCaseExecutionIdOk

`func (o *VariableInstanceDto) GetCaseExecutionIdOk() (*string, bool)`

GetCaseExecutionIdOk returns a tuple with the CaseExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExecutionId

`func (o *VariableInstanceDto) SetCaseExecutionId(v string)`

SetCaseExecutionId sets CaseExecutionId field to given value.

### HasCaseExecutionId

`func (o *VariableInstanceDto) HasCaseExecutionId() bool`

HasCaseExecutionId returns a boolean if a field has been set.

### SetCaseExecutionIdNil

`func (o *VariableInstanceDto) SetCaseExecutionIdNil(b bool)`

 SetCaseExecutionIdNil sets the value for CaseExecutionId to be an explicit nil

### UnsetCaseExecutionId
`func (o *VariableInstanceDto) UnsetCaseExecutionId()`

UnsetCaseExecutionId ensures that no value is present for CaseExecutionId, not even an explicit nil
### GetTaskId

`func (o *VariableInstanceDto) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *VariableInstanceDto) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *VariableInstanceDto) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *VariableInstanceDto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *VariableInstanceDto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *VariableInstanceDto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetBatchId

`func (o *VariableInstanceDto) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *VariableInstanceDto) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *VariableInstanceDto) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *VariableInstanceDto) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchIdNil

`func (o *VariableInstanceDto) SetBatchIdNil(b bool)`

 SetBatchIdNil sets the value for BatchId to be an explicit nil

### UnsetBatchId
`func (o *VariableInstanceDto) UnsetBatchId()`

UnsetBatchId ensures that no value is present for BatchId, not even an explicit nil
### GetActivityInstanceId

`func (o *VariableInstanceDto) GetActivityInstanceId() string`

GetActivityInstanceId returns the ActivityInstanceId field if non-nil, zero value otherwise.

### GetActivityInstanceIdOk

`func (o *VariableInstanceDto) GetActivityInstanceIdOk() (*string, bool)`

GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceId

`func (o *VariableInstanceDto) SetActivityInstanceId(v string)`

SetActivityInstanceId sets ActivityInstanceId field to given value.

### HasActivityInstanceId

`func (o *VariableInstanceDto) HasActivityInstanceId() bool`

HasActivityInstanceId returns a boolean if a field has been set.

### SetActivityInstanceIdNil

`func (o *VariableInstanceDto) SetActivityInstanceIdNil(b bool)`

 SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil

### UnsetActivityInstanceId
`func (o *VariableInstanceDto) UnsetActivityInstanceId()`

UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
### GetTenantId

`func (o *VariableInstanceDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *VariableInstanceDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *VariableInstanceDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *VariableInstanceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *VariableInstanceDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *VariableInstanceDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetErrorMessage

`func (o *VariableInstanceDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *VariableInstanceDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *VariableInstanceDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *VariableInstanceDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *VariableInstanceDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *VariableInstanceDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


