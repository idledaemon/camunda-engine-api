# VariableInstanceQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariableName** | Pointer to **NullableString** | Filter by variable instance name. | [optional] 
**VariableNameLike** | Pointer to **NullableString** | Filter by the variable instance name. The parameter can include the wildcard &#x60;%&#x60; to express like-strategy such as: starts with (&#x60;%&#x60;name), ends with (name&#x60;%&#x60;) or contains (&#x60;%&#x60;name&#x60;%&#x60;). | [optional] 
**ProcessInstanceIdIn** | Pointer to **[]string** | Only include variable instances which belong to one of the passed  process instance ids. | [optional] 
**ExecutionIdIn** | Pointer to **[]string** | Only include variable instances which belong to one of the passed  execution ids. | [optional] 
**CaseInstanceIdIn** | Pointer to **[]string** | Only include variable instances which belong to one of the passed  case instance ids. | [optional] 
**CaseExecutionIdIn** | Pointer to **[]string** | Only include variable instances which belong to one of the passed  case execution ids. | [optional] 
**TaskIdIn** | Pointer to **[]string** | Only include variable instances which belong to one of the passed  task ids. | [optional] 
**BatchIdIn** | Pointer to **[]string** | Only include variable instances which belong to one of the passed  batch ids. | [optional] 
**ActivityInstanceIdIn** | Pointer to **[]string** | Only include variable instances which belong to one of the passed  activity instance ids. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Only include variable instances which belong to one of the passed  tenant ids. | [optional] 
**VariableValues** | Pointer to [**[]VariableQueryParameterDto**](VariableQueryParameterDto.md) | An array to only include variable instances that have the certain values. The array consists of objects with the three properties &#x60;name&#x60;, &#x60;operator&#x60; and &#x60;value&#x60;. &#x60;name (String)&#x60; is the variable name, &#x60;operator (String)&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. &#x60;value&#x60; may be &#x60;String&#x60;, &#x60;Number&#x60; or &#x60;Boolean&#x60;.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60; | [optional] 
**VariableNamesIgnoreCase** | Pointer to **NullableBool** | Match all variable names provided in &#x60;variableValues&#x60; case-insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | [optional] 
**VariableValuesIgnoreCase** | Pointer to **NullableBool** | Match all variable values provided in &#x60;variableValues&#x60; case-insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | [optional] 
**VariableScopeIdIn** | Pointer to **[]string** | Only include variable instances which belong to one of passed scope ids. | [optional] 
**Sorting** | Pointer to [**[]VariableInstanceQueryDtoSortingInner**](VariableInstanceQueryDtoSortingInner.md) | An array of criteria to sort the result by. Each element of the array is an object that specifies one ordering.                       The position in the array identifies the rank of an ordering, i.e., whether it is primary, secondary, etc.                       Sorting has no effect for &#x60;count&#x60; endpoints | [optional] 

## Methods

### NewVariableInstanceQueryDto

`func NewVariableInstanceQueryDto() *VariableInstanceQueryDto`

NewVariableInstanceQueryDto instantiates a new VariableInstanceQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableInstanceQueryDtoWithDefaults

`func NewVariableInstanceQueryDtoWithDefaults() *VariableInstanceQueryDto`

NewVariableInstanceQueryDtoWithDefaults instantiates a new VariableInstanceQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariableName

`func (o *VariableInstanceQueryDto) GetVariableName() string`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *VariableInstanceQueryDto) GetVariableNameOk() (*string, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *VariableInstanceQueryDto) SetVariableName(v string)`

SetVariableName sets VariableName field to given value.

### HasVariableName

`func (o *VariableInstanceQueryDto) HasVariableName() bool`

HasVariableName returns a boolean if a field has been set.

### SetVariableNameNil

`func (o *VariableInstanceQueryDto) SetVariableNameNil(b bool)`

 SetVariableNameNil sets the value for VariableName to be an explicit nil

### UnsetVariableName
`func (o *VariableInstanceQueryDto) UnsetVariableName()`

UnsetVariableName ensures that no value is present for VariableName, not even an explicit nil
### GetVariableNameLike

`func (o *VariableInstanceQueryDto) GetVariableNameLike() string`

GetVariableNameLike returns the VariableNameLike field if non-nil, zero value otherwise.

### GetVariableNameLikeOk

`func (o *VariableInstanceQueryDto) GetVariableNameLikeOk() (*string, bool)`

GetVariableNameLikeOk returns a tuple with the VariableNameLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNameLike

`func (o *VariableInstanceQueryDto) SetVariableNameLike(v string)`

SetVariableNameLike sets VariableNameLike field to given value.

### HasVariableNameLike

`func (o *VariableInstanceQueryDto) HasVariableNameLike() bool`

HasVariableNameLike returns a boolean if a field has been set.

### SetVariableNameLikeNil

`func (o *VariableInstanceQueryDto) SetVariableNameLikeNil(b bool)`

 SetVariableNameLikeNil sets the value for VariableNameLike to be an explicit nil

### UnsetVariableNameLike
`func (o *VariableInstanceQueryDto) UnsetVariableNameLike()`

UnsetVariableNameLike ensures that no value is present for VariableNameLike, not even an explicit nil
### GetProcessInstanceIdIn

`func (o *VariableInstanceQueryDto) GetProcessInstanceIdIn() []string`

GetProcessInstanceIdIn returns the ProcessInstanceIdIn field if non-nil, zero value otherwise.

### GetProcessInstanceIdInOk

`func (o *VariableInstanceQueryDto) GetProcessInstanceIdInOk() (*[]string, bool)`

GetProcessInstanceIdInOk returns a tuple with the ProcessInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIdIn

`func (o *VariableInstanceQueryDto) SetProcessInstanceIdIn(v []string)`

SetProcessInstanceIdIn sets ProcessInstanceIdIn field to given value.

### HasProcessInstanceIdIn

`func (o *VariableInstanceQueryDto) HasProcessInstanceIdIn() bool`

HasProcessInstanceIdIn returns a boolean if a field has been set.

### SetProcessInstanceIdInNil

`func (o *VariableInstanceQueryDto) SetProcessInstanceIdInNil(b bool)`

 SetProcessInstanceIdInNil sets the value for ProcessInstanceIdIn to be an explicit nil

### UnsetProcessInstanceIdIn
`func (o *VariableInstanceQueryDto) UnsetProcessInstanceIdIn()`

UnsetProcessInstanceIdIn ensures that no value is present for ProcessInstanceIdIn, not even an explicit nil
### GetExecutionIdIn

`func (o *VariableInstanceQueryDto) GetExecutionIdIn() []string`

GetExecutionIdIn returns the ExecutionIdIn field if non-nil, zero value otherwise.

### GetExecutionIdInOk

`func (o *VariableInstanceQueryDto) GetExecutionIdInOk() (*[]string, bool)`

GetExecutionIdInOk returns a tuple with the ExecutionIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionIdIn

`func (o *VariableInstanceQueryDto) SetExecutionIdIn(v []string)`

SetExecutionIdIn sets ExecutionIdIn field to given value.

### HasExecutionIdIn

`func (o *VariableInstanceQueryDto) HasExecutionIdIn() bool`

HasExecutionIdIn returns a boolean if a field has been set.

### SetExecutionIdInNil

`func (o *VariableInstanceQueryDto) SetExecutionIdInNil(b bool)`

 SetExecutionIdInNil sets the value for ExecutionIdIn to be an explicit nil

### UnsetExecutionIdIn
`func (o *VariableInstanceQueryDto) UnsetExecutionIdIn()`

UnsetExecutionIdIn ensures that no value is present for ExecutionIdIn, not even an explicit nil
### GetCaseInstanceIdIn

`func (o *VariableInstanceQueryDto) GetCaseInstanceIdIn() []string`

GetCaseInstanceIdIn returns the CaseInstanceIdIn field if non-nil, zero value otherwise.

### GetCaseInstanceIdInOk

`func (o *VariableInstanceQueryDto) GetCaseInstanceIdInOk() (*[]string, bool)`

GetCaseInstanceIdInOk returns a tuple with the CaseInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceIdIn

`func (o *VariableInstanceQueryDto) SetCaseInstanceIdIn(v []string)`

SetCaseInstanceIdIn sets CaseInstanceIdIn field to given value.

### HasCaseInstanceIdIn

`func (o *VariableInstanceQueryDto) HasCaseInstanceIdIn() bool`

HasCaseInstanceIdIn returns a boolean if a field has been set.

### SetCaseInstanceIdInNil

`func (o *VariableInstanceQueryDto) SetCaseInstanceIdInNil(b bool)`

 SetCaseInstanceIdInNil sets the value for CaseInstanceIdIn to be an explicit nil

### UnsetCaseInstanceIdIn
`func (o *VariableInstanceQueryDto) UnsetCaseInstanceIdIn()`

UnsetCaseInstanceIdIn ensures that no value is present for CaseInstanceIdIn, not even an explicit nil
### GetCaseExecutionIdIn

`func (o *VariableInstanceQueryDto) GetCaseExecutionIdIn() []string`

GetCaseExecutionIdIn returns the CaseExecutionIdIn field if non-nil, zero value otherwise.

### GetCaseExecutionIdInOk

`func (o *VariableInstanceQueryDto) GetCaseExecutionIdInOk() (*[]string, bool)`

GetCaseExecutionIdInOk returns a tuple with the CaseExecutionIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExecutionIdIn

`func (o *VariableInstanceQueryDto) SetCaseExecutionIdIn(v []string)`

SetCaseExecutionIdIn sets CaseExecutionIdIn field to given value.

### HasCaseExecutionIdIn

`func (o *VariableInstanceQueryDto) HasCaseExecutionIdIn() bool`

HasCaseExecutionIdIn returns a boolean if a field has been set.

### SetCaseExecutionIdInNil

`func (o *VariableInstanceQueryDto) SetCaseExecutionIdInNil(b bool)`

 SetCaseExecutionIdInNil sets the value for CaseExecutionIdIn to be an explicit nil

### UnsetCaseExecutionIdIn
`func (o *VariableInstanceQueryDto) UnsetCaseExecutionIdIn()`

UnsetCaseExecutionIdIn ensures that no value is present for CaseExecutionIdIn, not even an explicit nil
### GetTaskIdIn

`func (o *VariableInstanceQueryDto) GetTaskIdIn() []string`

GetTaskIdIn returns the TaskIdIn field if non-nil, zero value otherwise.

### GetTaskIdInOk

`func (o *VariableInstanceQueryDto) GetTaskIdInOk() (*[]string, bool)`

GetTaskIdInOk returns a tuple with the TaskIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIdIn

`func (o *VariableInstanceQueryDto) SetTaskIdIn(v []string)`

SetTaskIdIn sets TaskIdIn field to given value.

### HasTaskIdIn

`func (o *VariableInstanceQueryDto) HasTaskIdIn() bool`

HasTaskIdIn returns a boolean if a field has been set.

### SetTaskIdInNil

`func (o *VariableInstanceQueryDto) SetTaskIdInNil(b bool)`

 SetTaskIdInNil sets the value for TaskIdIn to be an explicit nil

### UnsetTaskIdIn
`func (o *VariableInstanceQueryDto) UnsetTaskIdIn()`

UnsetTaskIdIn ensures that no value is present for TaskIdIn, not even an explicit nil
### GetBatchIdIn

`func (o *VariableInstanceQueryDto) GetBatchIdIn() []string`

GetBatchIdIn returns the BatchIdIn field if non-nil, zero value otherwise.

### GetBatchIdInOk

`func (o *VariableInstanceQueryDto) GetBatchIdInOk() (*[]string, bool)`

GetBatchIdInOk returns a tuple with the BatchIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchIdIn

`func (o *VariableInstanceQueryDto) SetBatchIdIn(v []string)`

SetBatchIdIn sets BatchIdIn field to given value.

### HasBatchIdIn

`func (o *VariableInstanceQueryDto) HasBatchIdIn() bool`

HasBatchIdIn returns a boolean if a field has been set.

### SetBatchIdInNil

`func (o *VariableInstanceQueryDto) SetBatchIdInNil(b bool)`

 SetBatchIdInNil sets the value for BatchIdIn to be an explicit nil

### UnsetBatchIdIn
`func (o *VariableInstanceQueryDto) UnsetBatchIdIn()`

UnsetBatchIdIn ensures that no value is present for BatchIdIn, not even an explicit nil
### GetActivityInstanceIdIn

`func (o *VariableInstanceQueryDto) GetActivityInstanceIdIn() []string`

GetActivityInstanceIdIn returns the ActivityInstanceIdIn field if non-nil, zero value otherwise.

### GetActivityInstanceIdInOk

`func (o *VariableInstanceQueryDto) GetActivityInstanceIdInOk() (*[]string, bool)`

GetActivityInstanceIdInOk returns a tuple with the ActivityInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceIdIn

`func (o *VariableInstanceQueryDto) SetActivityInstanceIdIn(v []string)`

SetActivityInstanceIdIn sets ActivityInstanceIdIn field to given value.

### HasActivityInstanceIdIn

`func (o *VariableInstanceQueryDto) HasActivityInstanceIdIn() bool`

HasActivityInstanceIdIn returns a boolean if a field has been set.

### SetActivityInstanceIdInNil

`func (o *VariableInstanceQueryDto) SetActivityInstanceIdInNil(b bool)`

 SetActivityInstanceIdInNil sets the value for ActivityInstanceIdIn to be an explicit nil

### UnsetActivityInstanceIdIn
`func (o *VariableInstanceQueryDto) UnsetActivityInstanceIdIn()`

UnsetActivityInstanceIdIn ensures that no value is present for ActivityInstanceIdIn, not even an explicit nil
### GetTenantIdIn

`func (o *VariableInstanceQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *VariableInstanceQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *VariableInstanceQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *VariableInstanceQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *VariableInstanceQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *VariableInstanceQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetVariableValues

`func (o *VariableInstanceQueryDto) GetVariableValues() []VariableQueryParameterDto`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *VariableInstanceQueryDto) GetVariableValuesOk() (*[]VariableQueryParameterDto, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *VariableInstanceQueryDto) SetVariableValues(v []VariableQueryParameterDto)`

SetVariableValues sets VariableValues field to given value.

### HasVariableValues

`func (o *VariableInstanceQueryDto) HasVariableValues() bool`

HasVariableValues returns a boolean if a field has been set.

### SetVariableValuesNil

`func (o *VariableInstanceQueryDto) SetVariableValuesNil(b bool)`

 SetVariableValuesNil sets the value for VariableValues to be an explicit nil

### UnsetVariableValues
`func (o *VariableInstanceQueryDto) UnsetVariableValues()`

UnsetVariableValues ensures that no value is present for VariableValues, not even an explicit nil
### GetVariableNamesIgnoreCase

`func (o *VariableInstanceQueryDto) GetVariableNamesIgnoreCase() bool`

GetVariableNamesIgnoreCase returns the VariableNamesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableNamesIgnoreCaseOk

`func (o *VariableInstanceQueryDto) GetVariableNamesIgnoreCaseOk() (*bool, bool)`

GetVariableNamesIgnoreCaseOk returns a tuple with the VariableNamesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNamesIgnoreCase

`func (o *VariableInstanceQueryDto) SetVariableNamesIgnoreCase(v bool)`

SetVariableNamesIgnoreCase sets VariableNamesIgnoreCase field to given value.

### HasVariableNamesIgnoreCase

`func (o *VariableInstanceQueryDto) HasVariableNamesIgnoreCase() bool`

HasVariableNamesIgnoreCase returns a boolean if a field has been set.

### SetVariableNamesIgnoreCaseNil

`func (o *VariableInstanceQueryDto) SetVariableNamesIgnoreCaseNil(b bool)`

 SetVariableNamesIgnoreCaseNil sets the value for VariableNamesIgnoreCase to be an explicit nil

### UnsetVariableNamesIgnoreCase
`func (o *VariableInstanceQueryDto) UnsetVariableNamesIgnoreCase()`

UnsetVariableNamesIgnoreCase ensures that no value is present for VariableNamesIgnoreCase, not even an explicit nil
### GetVariableValuesIgnoreCase

`func (o *VariableInstanceQueryDto) GetVariableValuesIgnoreCase() bool`

GetVariableValuesIgnoreCase returns the VariableValuesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableValuesIgnoreCaseOk

`func (o *VariableInstanceQueryDto) GetVariableValuesIgnoreCaseOk() (*bool, bool)`

GetVariableValuesIgnoreCaseOk returns a tuple with the VariableValuesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValuesIgnoreCase

`func (o *VariableInstanceQueryDto) SetVariableValuesIgnoreCase(v bool)`

SetVariableValuesIgnoreCase sets VariableValuesIgnoreCase field to given value.

### HasVariableValuesIgnoreCase

`func (o *VariableInstanceQueryDto) HasVariableValuesIgnoreCase() bool`

HasVariableValuesIgnoreCase returns a boolean if a field has been set.

### SetVariableValuesIgnoreCaseNil

`func (o *VariableInstanceQueryDto) SetVariableValuesIgnoreCaseNil(b bool)`

 SetVariableValuesIgnoreCaseNil sets the value for VariableValuesIgnoreCase to be an explicit nil

### UnsetVariableValuesIgnoreCase
`func (o *VariableInstanceQueryDto) UnsetVariableValuesIgnoreCase()`

UnsetVariableValuesIgnoreCase ensures that no value is present for VariableValuesIgnoreCase, not even an explicit nil
### GetVariableScopeIdIn

`func (o *VariableInstanceQueryDto) GetVariableScopeIdIn() []string`

GetVariableScopeIdIn returns the VariableScopeIdIn field if non-nil, zero value otherwise.

### GetVariableScopeIdInOk

`func (o *VariableInstanceQueryDto) GetVariableScopeIdInOk() (*[]string, bool)`

GetVariableScopeIdInOk returns a tuple with the VariableScopeIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableScopeIdIn

`func (o *VariableInstanceQueryDto) SetVariableScopeIdIn(v []string)`

SetVariableScopeIdIn sets VariableScopeIdIn field to given value.

### HasVariableScopeIdIn

`func (o *VariableInstanceQueryDto) HasVariableScopeIdIn() bool`

HasVariableScopeIdIn returns a boolean if a field has been set.

### SetVariableScopeIdInNil

`func (o *VariableInstanceQueryDto) SetVariableScopeIdInNil(b bool)`

 SetVariableScopeIdInNil sets the value for VariableScopeIdIn to be an explicit nil

### UnsetVariableScopeIdIn
`func (o *VariableInstanceQueryDto) UnsetVariableScopeIdIn()`

UnsetVariableScopeIdIn ensures that no value is present for VariableScopeIdIn, not even an explicit nil
### GetSorting

`func (o *VariableInstanceQueryDto) GetSorting() []VariableInstanceQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *VariableInstanceQueryDto) GetSortingOk() (*[]VariableInstanceQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *VariableInstanceQueryDto) SetSorting(v []VariableInstanceQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *VariableInstanceQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *VariableInstanceQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *VariableInstanceQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


