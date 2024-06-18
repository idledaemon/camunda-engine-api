# HistoricVariableInstanceQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariableName** | Pointer to **NullableString** | Filter by variable name. | [optional] 
**VariableNameLike** | Pointer to **NullableString** | Restrict to variables with a name like the parameter. | [optional] 
**VariableValue** | Pointer to **map[string]interface{}** | Filter by variable value. May be &#x60;String&#x60;, &#x60;Number&#x60; or &#x60;Boolean&#x60;. | [optional] 
**VariableNamesIgnoreCase** | Pointer to **NullableBool** | Match the variable name provided in &#x60;variableName&#x60; and &#x60;variableNameLike&#x60; case- insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | [optional] 
**VariableValuesIgnoreCase** | Pointer to **NullableBool** | Match the variable value provided in &#x60;variableValue&#x60; case-insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | [optional] 
**VariableTypeIn** | Pointer to **[]string** | Only include historic variable instances which belong to one of the passed and comma- separated variable types. A list of all supported variable types can be found [here](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#supported-variable-values). **Note:** All non-primitive variables are associated with the type &#39;serializable&#39;. | [optional] 
**IncludeDeleted** | Pointer to **NullableBool** | Include variables that has already been deleted during the execution. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | Filter by the process instance the variable belongs to. | [optional] 
**ProcessInstanceIdIn** | Pointer to **[]string** | Only include historic variable instances which belong to one of the passed  process instance ids. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter by the process definition the variable belongs to. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Filter by a key of the process definition the variable belongs to. | [optional] 
**ExecutionIdIn** | Pointer to **[]string** | Only include historic variable instances which belong to one of the passed and  execution ids. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | Filter by the case instance the variable belongs to. | [optional] 
**CaseExecutionIdIn** | Pointer to **[]string** | Only include historic variable instances which belong to one of the passed and  case execution ids. | [optional] 
**CaseActivityIdIn** | Pointer to **[]string** | Only include historic variable instances which belong to one of the passed and  case activity ids. | [optional] 
**TaskIdIn** | Pointer to **[]string** | Only include historic variable instances which belong to one of the passed and  task ids. | [optional] 
**ActivityInstanceIdIn** | Pointer to **[]string** | Only include historic variable instances which belong to one of the passed and  activity instance ids. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Only include historic variable instances which belong to one of the passed and comma- separated tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include historic variable instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**VariableNameIn** | Pointer to **[]string** | Only include historic variable instances which belong to one of the passed  variable names. | [optional] 
**Sorting** | Pointer to [**[]HistoricVariableInstanceQueryDtoSortingInner**](HistoricVariableInstanceQueryDtoSortingInner.md) | An array of criteria to sort the result by. Each element of the array is                      an object that specifies one ordering. The position in the array                      identifies the rank of an ordering, i.e., whether it is primary, secondary,                      etc. Sorting has no effect for &#x60;count&#x60; endpoints | [optional] 

## Methods

### NewHistoricVariableInstanceQueryDto

`func NewHistoricVariableInstanceQueryDto() *HistoricVariableInstanceQueryDto`

NewHistoricVariableInstanceQueryDto instantiates a new HistoricVariableInstanceQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricVariableInstanceQueryDtoWithDefaults

`func NewHistoricVariableInstanceQueryDtoWithDefaults() *HistoricVariableInstanceQueryDto`

NewHistoricVariableInstanceQueryDtoWithDefaults instantiates a new HistoricVariableInstanceQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariableName

`func (o *HistoricVariableInstanceQueryDto) GetVariableName() string`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *HistoricVariableInstanceQueryDto) GetVariableNameOk() (*string, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *HistoricVariableInstanceQueryDto) SetVariableName(v string)`

SetVariableName sets VariableName field to given value.

### HasVariableName

`func (o *HistoricVariableInstanceQueryDto) HasVariableName() bool`

HasVariableName returns a boolean if a field has been set.

### SetVariableNameNil

`func (o *HistoricVariableInstanceQueryDto) SetVariableNameNil(b bool)`

 SetVariableNameNil sets the value for VariableName to be an explicit nil

### UnsetVariableName
`func (o *HistoricVariableInstanceQueryDto) UnsetVariableName()`

UnsetVariableName ensures that no value is present for VariableName, not even an explicit nil
### GetVariableNameLike

`func (o *HistoricVariableInstanceQueryDto) GetVariableNameLike() string`

GetVariableNameLike returns the VariableNameLike field if non-nil, zero value otherwise.

### GetVariableNameLikeOk

`func (o *HistoricVariableInstanceQueryDto) GetVariableNameLikeOk() (*string, bool)`

GetVariableNameLikeOk returns a tuple with the VariableNameLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNameLike

`func (o *HistoricVariableInstanceQueryDto) SetVariableNameLike(v string)`

SetVariableNameLike sets VariableNameLike field to given value.

### HasVariableNameLike

`func (o *HistoricVariableInstanceQueryDto) HasVariableNameLike() bool`

HasVariableNameLike returns a boolean if a field has been set.

### SetVariableNameLikeNil

`func (o *HistoricVariableInstanceQueryDto) SetVariableNameLikeNil(b bool)`

 SetVariableNameLikeNil sets the value for VariableNameLike to be an explicit nil

### UnsetVariableNameLike
`func (o *HistoricVariableInstanceQueryDto) UnsetVariableNameLike()`

UnsetVariableNameLike ensures that no value is present for VariableNameLike, not even an explicit nil
### GetVariableValue

`func (o *HistoricVariableInstanceQueryDto) GetVariableValue() map[string]interface{}`

GetVariableValue returns the VariableValue field if non-nil, zero value otherwise.

### GetVariableValueOk

`func (o *HistoricVariableInstanceQueryDto) GetVariableValueOk() (*map[string]interface{}, bool)`

GetVariableValueOk returns a tuple with the VariableValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValue

`func (o *HistoricVariableInstanceQueryDto) SetVariableValue(v map[string]interface{})`

SetVariableValue sets VariableValue field to given value.

### HasVariableValue

`func (o *HistoricVariableInstanceQueryDto) HasVariableValue() bool`

HasVariableValue returns a boolean if a field has been set.

### GetVariableNamesIgnoreCase

`func (o *HistoricVariableInstanceQueryDto) GetVariableNamesIgnoreCase() bool`

GetVariableNamesIgnoreCase returns the VariableNamesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableNamesIgnoreCaseOk

`func (o *HistoricVariableInstanceQueryDto) GetVariableNamesIgnoreCaseOk() (*bool, bool)`

GetVariableNamesIgnoreCaseOk returns a tuple with the VariableNamesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNamesIgnoreCase

`func (o *HistoricVariableInstanceQueryDto) SetVariableNamesIgnoreCase(v bool)`

SetVariableNamesIgnoreCase sets VariableNamesIgnoreCase field to given value.

### HasVariableNamesIgnoreCase

`func (o *HistoricVariableInstanceQueryDto) HasVariableNamesIgnoreCase() bool`

HasVariableNamesIgnoreCase returns a boolean if a field has been set.

### SetVariableNamesIgnoreCaseNil

`func (o *HistoricVariableInstanceQueryDto) SetVariableNamesIgnoreCaseNil(b bool)`

 SetVariableNamesIgnoreCaseNil sets the value for VariableNamesIgnoreCase to be an explicit nil

### UnsetVariableNamesIgnoreCase
`func (o *HistoricVariableInstanceQueryDto) UnsetVariableNamesIgnoreCase()`

UnsetVariableNamesIgnoreCase ensures that no value is present for VariableNamesIgnoreCase, not even an explicit nil
### GetVariableValuesIgnoreCase

`func (o *HistoricVariableInstanceQueryDto) GetVariableValuesIgnoreCase() bool`

GetVariableValuesIgnoreCase returns the VariableValuesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableValuesIgnoreCaseOk

`func (o *HistoricVariableInstanceQueryDto) GetVariableValuesIgnoreCaseOk() (*bool, bool)`

GetVariableValuesIgnoreCaseOk returns a tuple with the VariableValuesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValuesIgnoreCase

`func (o *HistoricVariableInstanceQueryDto) SetVariableValuesIgnoreCase(v bool)`

SetVariableValuesIgnoreCase sets VariableValuesIgnoreCase field to given value.

### HasVariableValuesIgnoreCase

`func (o *HistoricVariableInstanceQueryDto) HasVariableValuesIgnoreCase() bool`

HasVariableValuesIgnoreCase returns a boolean if a field has been set.

### SetVariableValuesIgnoreCaseNil

`func (o *HistoricVariableInstanceQueryDto) SetVariableValuesIgnoreCaseNil(b bool)`

 SetVariableValuesIgnoreCaseNil sets the value for VariableValuesIgnoreCase to be an explicit nil

### UnsetVariableValuesIgnoreCase
`func (o *HistoricVariableInstanceQueryDto) UnsetVariableValuesIgnoreCase()`

UnsetVariableValuesIgnoreCase ensures that no value is present for VariableValuesIgnoreCase, not even an explicit nil
### GetVariableTypeIn

`func (o *HistoricVariableInstanceQueryDto) GetVariableTypeIn() []string`

GetVariableTypeIn returns the VariableTypeIn field if non-nil, zero value otherwise.

### GetVariableTypeInOk

`func (o *HistoricVariableInstanceQueryDto) GetVariableTypeInOk() (*[]string, bool)`

GetVariableTypeInOk returns a tuple with the VariableTypeIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableTypeIn

`func (o *HistoricVariableInstanceQueryDto) SetVariableTypeIn(v []string)`

SetVariableTypeIn sets VariableTypeIn field to given value.

### HasVariableTypeIn

`func (o *HistoricVariableInstanceQueryDto) HasVariableTypeIn() bool`

HasVariableTypeIn returns a boolean if a field has been set.

### SetVariableTypeInNil

`func (o *HistoricVariableInstanceQueryDto) SetVariableTypeInNil(b bool)`

 SetVariableTypeInNil sets the value for VariableTypeIn to be an explicit nil

### UnsetVariableTypeIn
`func (o *HistoricVariableInstanceQueryDto) UnsetVariableTypeIn()`

UnsetVariableTypeIn ensures that no value is present for VariableTypeIn, not even an explicit nil
### GetIncludeDeleted

`func (o *HistoricVariableInstanceQueryDto) GetIncludeDeleted() bool`

GetIncludeDeleted returns the IncludeDeleted field if non-nil, zero value otherwise.

### GetIncludeDeletedOk

`func (o *HistoricVariableInstanceQueryDto) GetIncludeDeletedOk() (*bool, bool)`

GetIncludeDeletedOk returns a tuple with the IncludeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDeleted

`func (o *HistoricVariableInstanceQueryDto) SetIncludeDeleted(v bool)`

SetIncludeDeleted sets IncludeDeleted field to given value.

### HasIncludeDeleted

`func (o *HistoricVariableInstanceQueryDto) HasIncludeDeleted() bool`

HasIncludeDeleted returns a boolean if a field has been set.

### SetIncludeDeletedNil

`func (o *HistoricVariableInstanceQueryDto) SetIncludeDeletedNil(b bool)`

 SetIncludeDeletedNil sets the value for IncludeDeleted to be an explicit nil

### UnsetIncludeDeleted
`func (o *HistoricVariableInstanceQueryDto) UnsetIncludeDeleted()`

UnsetIncludeDeleted ensures that no value is present for IncludeDeleted, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricVariableInstanceQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricVariableInstanceQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricVariableInstanceQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricVariableInstanceQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricVariableInstanceQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricVariableInstanceQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessInstanceIdIn

`func (o *HistoricVariableInstanceQueryDto) GetProcessInstanceIdIn() []string`

GetProcessInstanceIdIn returns the ProcessInstanceIdIn field if non-nil, zero value otherwise.

### GetProcessInstanceIdInOk

`func (o *HistoricVariableInstanceQueryDto) GetProcessInstanceIdInOk() (*[]string, bool)`

GetProcessInstanceIdInOk returns a tuple with the ProcessInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIdIn

`func (o *HistoricVariableInstanceQueryDto) SetProcessInstanceIdIn(v []string)`

SetProcessInstanceIdIn sets ProcessInstanceIdIn field to given value.

### HasProcessInstanceIdIn

`func (o *HistoricVariableInstanceQueryDto) HasProcessInstanceIdIn() bool`

HasProcessInstanceIdIn returns a boolean if a field has been set.

### SetProcessInstanceIdInNil

`func (o *HistoricVariableInstanceQueryDto) SetProcessInstanceIdInNil(b bool)`

 SetProcessInstanceIdInNil sets the value for ProcessInstanceIdIn to be an explicit nil

### UnsetProcessInstanceIdIn
`func (o *HistoricVariableInstanceQueryDto) UnsetProcessInstanceIdIn()`

UnsetProcessInstanceIdIn ensures that no value is present for ProcessInstanceIdIn, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricVariableInstanceQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricVariableInstanceQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricVariableInstanceQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricVariableInstanceQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricVariableInstanceQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricVariableInstanceQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricVariableInstanceQueryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricVariableInstanceQueryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricVariableInstanceQueryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricVariableInstanceQueryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricVariableInstanceQueryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricVariableInstanceQueryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetExecutionIdIn

`func (o *HistoricVariableInstanceQueryDto) GetExecutionIdIn() []string`

GetExecutionIdIn returns the ExecutionIdIn field if non-nil, zero value otherwise.

### GetExecutionIdInOk

`func (o *HistoricVariableInstanceQueryDto) GetExecutionIdInOk() (*[]string, bool)`

GetExecutionIdInOk returns a tuple with the ExecutionIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionIdIn

`func (o *HistoricVariableInstanceQueryDto) SetExecutionIdIn(v []string)`

SetExecutionIdIn sets ExecutionIdIn field to given value.

### HasExecutionIdIn

`func (o *HistoricVariableInstanceQueryDto) HasExecutionIdIn() bool`

HasExecutionIdIn returns a boolean if a field has been set.

### SetExecutionIdInNil

`func (o *HistoricVariableInstanceQueryDto) SetExecutionIdInNil(b bool)`

 SetExecutionIdInNil sets the value for ExecutionIdIn to be an explicit nil

### UnsetExecutionIdIn
`func (o *HistoricVariableInstanceQueryDto) UnsetExecutionIdIn()`

UnsetExecutionIdIn ensures that no value is present for ExecutionIdIn, not even an explicit nil
### GetCaseInstanceId

`func (o *HistoricVariableInstanceQueryDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *HistoricVariableInstanceQueryDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *HistoricVariableInstanceQueryDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *HistoricVariableInstanceQueryDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *HistoricVariableInstanceQueryDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *HistoricVariableInstanceQueryDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetCaseExecutionIdIn

`func (o *HistoricVariableInstanceQueryDto) GetCaseExecutionIdIn() []string`

GetCaseExecutionIdIn returns the CaseExecutionIdIn field if non-nil, zero value otherwise.

### GetCaseExecutionIdInOk

`func (o *HistoricVariableInstanceQueryDto) GetCaseExecutionIdInOk() (*[]string, bool)`

GetCaseExecutionIdInOk returns a tuple with the CaseExecutionIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExecutionIdIn

`func (o *HistoricVariableInstanceQueryDto) SetCaseExecutionIdIn(v []string)`

SetCaseExecutionIdIn sets CaseExecutionIdIn field to given value.

### HasCaseExecutionIdIn

`func (o *HistoricVariableInstanceQueryDto) HasCaseExecutionIdIn() bool`

HasCaseExecutionIdIn returns a boolean if a field has been set.

### SetCaseExecutionIdInNil

`func (o *HistoricVariableInstanceQueryDto) SetCaseExecutionIdInNil(b bool)`

 SetCaseExecutionIdInNil sets the value for CaseExecutionIdIn to be an explicit nil

### UnsetCaseExecutionIdIn
`func (o *HistoricVariableInstanceQueryDto) UnsetCaseExecutionIdIn()`

UnsetCaseExecutionIdIn ensures that no value is present for CaseExecutionIdIn, not even an explicit nil
### GetCaseActivityIdIn

`func (o *HistoricVariableInstanceQueryDto) GetCaseActivityIdIn() []string`

GetCaseActivityIdIn returns the CaseActivityIdIn field if non-nil, zero value otherwise.

### GetCaseActivityIdInOk

`func (o *HistoricVariableInstanceQueryDto) GetCaseActivityIdInOk() (*[]string, bool)`

GetCaseActivityIdInOk returns a tuple with the CaseActivityIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseActivityIdIn

`func (o *HistoricVariableInstanceQueryDto) SetCaseActivityIdIn(v []string)`

SetCaseActivityIdIn sets CaseActivityIdIn field to given value.

### HasCaseActivityIdIn

`func (o *HistoricVariableInstanceQueryDto) HasCaseActivityIdIn() bool`

HasCaseActivityIdIn returns a boolean if a field has been set.

### SetCaseActivityIdInNil

`func (o *HistoricVariableInstanceQueryDto) SetCaseActivityIdInNil(b bool)`

 SetCaseActivityIdInNil sets the value for CaseActivityIdIn to be an explicit nil

### UnsetCaseActivityIdIn
`func (o *HistoricVariableInstanceQueryDto) UnsetCaseActivityIdIn()`

UnsetCaseActivityIdIn ensures that no value is present for CaseActivityIdIn, not even an explicit nil
### GetTaskIdIn

`func (o *HistoricVariableInstanceQueryDto) GetTaskIdIn() []string`

GetTaskIdIn returns the TaskIdIn field if non-nil, zero value otherwise.

### GetTaskIdInOk

`func (o *HistoricVariableInstanceQueryDto) GetTaskIdInOk() (*[]string, bool)`

GetTaskIdInOk returns a tuple with the TaskIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIdIn

`func (o *HistoricVariableInstanceQueryDto) SetTaskIdIn(v []string)`

SetTaskIdIn sets TaskIdIn field to given value.

### HasTaskIdIn

`func (o *HistoricVariableInstanceQueryDto) HasTaskIdIn() bool`

HasTaskIdIn returns a boolean if a field has been set.

### SetTaskIdInNil

`func (o *HistoricVariableInstanceQueryDto) SetTaskIdInNil(b bool)`

 SetTaskIdInNil sets the value for TaskIdIn to be an explicit nil

### UnsetTaskIdIn
`func (o *HistoricVariableInstanceQueryDto) UnsetTaskIdIn()`

UnsetTaskIdIn ensures that no value is present for TaskIdIn, not even an explicit nil
### GetActivityInstanceIdIn

`func (o *HistoricVariableInstanceQueryDto) GetActivityInstanceIdIn() []string`

GetActivityInstanceIdIn returns the ActivityInstanceIdIn field if non-nil, zero value otherwise.

### GetActivityInstanceIdInOk

`func (o *HistoricVariableInstanceQueryDto) GetActivityInstanceIdInOk() (*[]string, bool)`

GetActivityInstanceIdInOk returns a tuple with the ActivityInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceIdIn

`func (o *HistoricVariableInstanceQueryDto) SetActivityInstanceIdIn(v []string)`

SetActivityInstanceIdIn sets ActivityInstanceIdIn field to given value.

### HasActivityInstanceIdIn

`func (o *HistoricVariableInstanceQueryDto) HasActivityInstanceIdIn() bool`

HasActivityInstanceIdIn returns a boolean if a field has been set.

### SetActivityInstanceIdInNil

`func (o *HistoricVariableInstanceQueryDto) SetActivityInstanceIdInNil(b bool)`

 SetActivityInstanceIdInNil sets the value for ActivityInstanceIdIn to be an explicit nil

### UnsetActivityInstanceIdIn
`func (o *HistoricVariableInstanceQueryDto) UnsetActivityInstanceIdIn()`

UnsetActivityInstanceIdIn ensures that no value is present for ActivityInstanceIdIn, not even an explicit nil
### GetTenantIdIn

`func (o *HistoricVariableInstanceQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *HistoricVariableInstanceQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *HistoricVariableInstanceQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *HistoricVariableInstanceQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *HistoricVariableInstanceQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *HistoricVariableInstanceQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *HistoricVariableInstanceQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *HistoricVariableInstanceQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *HistoricVariableInstanceQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *HistoricVariableInstanceQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *HistoricVariableInstanceQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *HistoricVariableInstanceQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetVariableNameIn

`func (o *HistoricVariableInstanceQueryDto) GetVariableNameIn() []string`

GetVariableNameIn returns the VariableNameIn field if non-nil, zero value otherwise.

### GetVariableNameInOk

`func (o *HistoricVariableInstanceQueryDto) GetVariableNameInOk() (*[]string, bool)`

GetVariableNameInOk returns a tuple with the VariableNameIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNameIn

`func (o *HistoricVariableInstanceQueryDto) SetVariableNameIn(v []string)`

SetVariableNameIn sets VariableNameIn field to given value.

### HasVariableNameIn

`func (o *HistoricVariableInstanceQueryDto) HasVariableNameIn() bool`

HasVariableNameIn returns a boolean if a field has been set.

### SetVariableNameInNil

`func (o *HistoricVariableInstanceQueryDto) SetVariableNameInNil(b bool)`

 SetVariableNameInNil sets the value for VariableNameIn to be an explicit nil

### UnsetVariableNameIn
`func (o *HistoricVariableInstanceQueryDto) UnsetVariableNameIn()`

UnsetVariableNameIn ensures that no value is present for VariableNameIn, not even an explicit nil
### GetSorting

`func (o *HistoricVariableInstanceQueryDto) GetSorting() []HistoricVariableInstanceQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *HistoricVariableInstanceQueryDto) GetSortingOk() (*[]HistoricVariableInstanceQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *HistoricVariableInstanceQueryDto) SetSorting(v []HistoricVariableInstanceQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *HistoricVariableInstanceQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *HistoricVariableInstanceQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *HistoricVariableInstanceQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


