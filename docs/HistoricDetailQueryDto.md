# HistoricDetailQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessInstanceId** | Pointer to **NullableString** | Filter by process instance id. | [optional] 
**ProcessInstanceIdIn** | Pointer to **[]string** | Only include historic details which belong to one of the passed  process instance ids. | [optional] 
**ExecutionId** | Pointer to **NullableString** | Filter by execution id. | [optional] 
**TaskId** | Pointer to **NullableString** | Filter by task id. | [optional] 
**ActivityInstanceId** | Pointer to **NullableString** | Filter by activity instance id. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | Filter by case instance id. | [optional] 
**CaseExecutionId** | Pointer to **NullableString** | Filter by case execution id. | [optional] 
**VariableInstanceId** | Pointer to **NullableString** | Filter by variable instance id. | [optional] 
**VariableTypeIn** | Pointer to **[]string** | Only include historic details where the variable updates belong to one of the passed  list of variable types. A list of all supported variable types can be found [here](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#supported-variable-values). **Note:** All non-primitive variables are associated with the type &#x60;serializable&#x60;. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Filter by a  list of tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include historic details that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**UserOperationId** | Pointer to **NullableString** | Filter by a user operation id. | [optional] 
**FormFields** | Pointer to **NullableBool** | Only include &#x60;HistoricFormFields&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**VariableUpdates** | Pointer to **NullableBool** | Only include &#x60;HistoricVariableUpdates&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**ExcludeTaskDetails** | Pointer to **NullableBool** | Excludes all task-related &#x60;HistoricDetails&#x60;, so only items which have no task id set will be selected. When this parameter is used together with &#x60;taskId&#x60;, this call is ignored and task details are not excluded. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Initial** | Pointer to **NullableBool** | Restrict to historic variable updates that contain only initial variable values. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**OccurredBefore** | Pointer to **NullableTime** | Restrict to historic details that occured before the given date (including the date). Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200. | [optional] 
**OccurredAfter** | Pointer to **NullableTime** | Restrict to historic details that occured after the given date (including the date). Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200. | [optional] 
**Sorting** | Pointer to [**[]HistoricDetailQueryDtoSortingInner**](HistoricDetailQueryDtoSortingInner.md) | A JSON array of criteria to sort the result by. Each element of the array is                     a JSON object that specifies one ordering. The position in the array                     identifies the rank of an ordering, i.e., whether it is primary, secondary,                     etc. Does not have an effect for the &#x60;count&#x60; endpoint. | [optional] 

## Methods

### NewHistoricDetailQueryDto

`func NewHistoricDetailQueryDto() *HistoricDetailQueryDto`

NewHistoricDetailQueryDto instantiates a new HistoricDetailQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricDetailQueryDtoWithDefaults

`func NewHistoricDetailQueryDtoWithDefaults() *HistoricDetailQueryDto`

NewHistoricDetailQueryDtoWithDefaults instantiates a new HistoricDetailQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessInstanceId

`func (o *HistoricDetailQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricDetailQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricDetailQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricDetailQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricDetailQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricDetailQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessInstanceIdIn

`func (o *HistoricDetailQueryDto) GetProcessInstanceIdIn() []string`

GetProcessInstanceIdIn returns the ProcessInstanceIdIn field if non-nil, zero value otherwise.

### GetProcessInstanceIdInOk

`func (o *HistoricDetailQueryDto) GetProcessInstanceIdInOk() (*[]string, bool)`

GetProcessInstanceIdInOk returns a tuple with the ProcessInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIdIn

`func (o *HistoricDetailQueryDto) SetProcessInstanceIdIn(v []string)`

SetProcessInstanceIdIn sets ProcessInstanceIdIn field to given value.

### HasProcessInstanceIdIn

`func (o *HistoricDetailQueryDto) HasProcessInstanceIdIn() bool`

HasProcessInstanceIdIn returns a boolean if a field has been set.

### SetProcessInstanceIdInNil

`func (o *HistoricDetailQueryDto) SetProcessInstanceIdInNil(b bool)`

 SetProcessInstanceIdInNil sets the value for ProcessInstanceIdIn to be an explicit nil

### UnsetProcessInstanceIdIn
`func (o *HistoricDetailQueryDto) UnsetProcessInstanceIdIn()`

UnsetProcessInstanceIdIn ensures that no value is present for ProcessInstanceIdIn, not even an explicit nil
### GetExecutionId

`func (o *HistoricDetailQueryDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HistoricDetailQueryDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HistoricDetailQueryDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HistoricDetailQueryDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HistoricDetailQueryDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HistoricDetailQueryDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetTaskId

`func (o *HistoricDetailQueryDto) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *HistoricDetailQueryDto) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *HistoricDetailQueryDto) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *HistoricDetailQueryDto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *HistoricDetailQueryDto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *HistoricDetailQueryDto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetActivityInstanceId

`func (o *HistoricDetailQueryDto) GetActivityInstanceId() string`

GetActivityInstanceId returns the ActivityInstanceId field if non-nil, zero value otherwise.

### GetActivityInstanceIdOk

`func (o *HistoricDetailQueryDto) GetActivityInstanceIdOk() (*string, bool)`

GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceId

`func (o *HistoricDetailQueryDto) SetActivityInstanceId(v string)`

SetActivityInstanceId sets ActivityInstanceId field to given value.

### HasActivityInstanceId

`func (o *HistoricDetailQueryDto) HasActivityInstanceId() bool`

HasActivityInstanceId returns a boolean if a field has been set.

### SetActivityInstanceIdNil

`func (o *HistoricDetailQueryDto) SetActivityInstanceIdNil(b bool)`

 SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil

### UnsetActivityInstanceId
`func (o *HistoricDetailQueryDto) UnsetActivityInstanceId()`

UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
### GetCaseInstanceId

`func (o *HistoricDetailQueryDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *HistoricDetailQueryDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *HistoricDetailQueryDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *HistoricDetailQueryDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *HistoricDetailQueryDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *HistoricDetailQueryDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetCaseExecutionId

`func (o *HistoricDetailQueryDto) GetCaseExecutionId() string`

GetCaseExecutionId returns the CaseExecutionId field if non-nil, zero value otherwise.

### GetCaseExecutionIdOk

`func (o *HistoricDetailQueryDto) GetCaseExecutionIdOk() (*string, bool)`

GetCaseExecutionIdOk returns a tuple with the CaseExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExecutionId

`func (o *HistoricDetailQueryDto) SetCaseExecutionId(v string)`

SetCaseExecutionId sets CaseExecutionId field to given value.

### HasCaseExecutionId

`func (o *HistoricDetailQueryDto) HasCaseExecutionId() bool`

HasCaseExecutionId returns a boolean if a field has been set.

### SetCaseExecutionIdNil

`func (o *HistoricDetailQueryDto) SetCaseExecutionIdNil(b bool)`

 SetCaseExecutionIdNil sets the value for CaseExecutionId to be an explicit nil

### UnsetCaseExecutionId
`func (o *HistoricDetailQueryDto) UnsetCaseExecutionId()`

UnsetCaseExecutionId ensures that no value is present for CaseExecutionId, not even an explicit nil
### GetVariableInstanceId

`func (o *HistoricDetailQueryDto) GetVariableInstanceId() string`

GetVariableInstanceId returns the VariableInstanceId field if non-nil, zero value otherwise.

### GetVariableInstanceIdOk

`func (o *HistoricDetailQueryDto) GetVariableInstanceIdOk() (*string, bool)`

GetVariableInstanceIdOk returns a tuple with the VariableInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableInstanceId

`func (o *HistoricDetailQueryDto) SetVariableInstanceId(v string)`

SetVariableInstanceId sets VariableInstanceId field to given value.

### HasVariableInstanceId

`func (o *HistoricDetailQueryDto) HasVariableInstanceId() bool`

HasVariableInstanceId returns a boolean if a field has been set.

### SetVariableInstanceIdNil

`func (o *HistoricDetailQueryDto) SetVariableInstanceIdNil(b bool)`

 SetVariableInstanceIdNil sets the value for VariableInstanceId to be an explicit nil

### UnsetVariableInstanceId
`func (o *HistoricDetailQueryDto) UnsetVariableInstanceId()`

UnsetVariableInstanceId ensures that no value is present for VariableInstanceId, not even an explicit nil
### GetVariableTypeIn

`func (o *HistoricDetailQueryDto) GetVariableTypeIn() []string`

GetVariableTypeIn returns the VariableTypeIn field if non-nil, zero value otherwise.

### GetVariableTypeInOk

`func (o *HistoricDetailQueryDto) GetVariableTypeInOk() (*[]string, bool)`

GetVariableTypeInOk returns a tuple with the VariableTypeIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableTypeIn

`func (o *HistoricDetailQueryDto) SetVariableTypeIn(v []string)`

SetVariableTypeIn sets VariableTypeIn field to given value.

### HasVariableTypeIn

`func (o *HistoricDetailQueryDto) HasVariableTypeIn() bool`

HasVariableTypeIn returns a boolean if a field has been set.

### SetVariableTypeInNil

`func (o *HistoricDetailQueryDto) SetVariableTypeInNil(b bool)`

 SetVariableTypeInNil sets the value for VariableTypeIn to be an explicit nil

### UnsetVariableTypeIn
`func (o *HistoricDetailQueryDto) UnsetVariableTypeIn()`

UnsetVariableTypeIn ensures that no value is present for VariableTypeIn, not even an explicit nil
### GetTenantIdIn

`func (o *HistoricDetailQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *HistoricDetailQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *HistoricDetailQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *HistoricDetailQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *HistoricDetailQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *HistoricDetailQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *HistoricDetailQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *HistoricDetailQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *HistoricDetailQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *HistoricDetailQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *HistoricDetailQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *HistoricDetailQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetUserOperationId

`func (o *HistoricDetailQueryDto) GetUserOperationId() string`

GetUserOperationId returns the UserOperationId field if non-nil, zero value otherwise.

### GetUserOperationIdOk

`func (o *HistoricDetailQueryDto) GetUserOperationIdOk() (*string, bool)`

GetUserOperationIdOk returns a tuple with the UserOperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOperationId

`func (o *HistoricDetailQueryDto) SetUserOperationId(v string)`

SetUserOperationId sets UserOperationId field to given value.

### HasUserOperationId

`func (o *HistoricDetailQueryDto) HasUserOperationId() bool`

HasUserOperationId returns a boolean if a field has been set.

### SetUserOperationIdNil

`func (o *HistoricDetailQueryDto) SetUserOperationIdNil(b bool)`

 SetUserOperationIdNil sets the value for UserOperationId to be an explicit nil

### UnsetUserOperationId
`func (o *HistoricDetailQueryDto) UnsetUserOperationId()`

UnsetUserOperationId ensures that no value is present for UserOperationId, not even an explicit nil
### GetFormFields

`func (o *HistoricDetailQueryDto) GetFormFields() bool`

GetFormFields returns the FormFields field if non-nil, zero value otherwise.

### GetFormFieldsOk

`func (o *HistoricDetailQueryDto) GetFormFieldsOk() (*bool, bool)`

GetFormFieldsOk returns a tuple with the FormFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFields

`func (o *HistoricDetailQueryDto) SetFormFields(v bool)`

SetFormFields sets FormFields field to given value.

### HasFormFields

`func (o *HistoricDetailQueryDto) HasFormFields() bool`

HasFormFields returns a boolean if a field has been set.

### SetFormFieldsNil

`func (o *HistoricDetailQueryDto) SetFormFieldsNil(b bool)`

 SetFormFieldsNil sets the value for FormFields to be an explicit nil

### UnsetFormFields
`func (o *HistoricDetailQueryDto) UnsetFormFields()`

UnsetFormFields ensures that no value is present for FormFields, not even an explicit nil
### GetVariableUpdates

`func (o *HistoricDetailQueryDto) GetVariableUpdates() bool`

GetVariableUpdates returns the VariableUpdates field if non-nil, zero value otherwise.

### GetVariableUpdatesOk

`func (o *HistoricDetailQueryDto) GetVariableUpdatesOk() (*bool, bool)`

GetVariableUpdatesOk returns a tuple with the VariableUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableUpdates

`func (o *HistoricDetailQueryDto) SetVariableUpdates(v bool)`

SetVariableUpdates sets VariableUpdates field to given value.

### HasVariableUpdates

`func (o *HistoricDetailQueryDto) HasVariableUpdates() bool`

HasVariableUpdates returns a boolean if a field has been set.

### SetVariableUpdatesNil

`func (o *HistoricDetailQueryDto) SetVariableUpdatesNil(b bool)`

 SetVariableUpdatesNil sets the value for VariableUpdates to be an explicit nil

### UnsetVariableUpdates
`func (o *HistoricDetailQueryDto) UnsetVariableUpdates()`

UnsetVariableUpdates ensures that no value is present for VariableUpdates, not even an explicit nil
### GetExcludeTaskDetails

`func (o *HistoricDetailQueryDto) GetExcludeTaskDetails() bool`

GetExcludeTaskDetails returns the ExcludeTaskDetails field if non-nil, zero value otherwise.

### GetExcludeTaskDetailsOk

`func (o *HistoricDetailQueryDto) GetExcludeTaskDetailsOk() (*bool, bool)`

GetExcludeTaskDetailsOk returns a tuple with the ExcludeTaskDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTaskDetails

`func (o *HistoricDetailQueryDto) SetExcludeTaskDetails(v bool)`

SetExcludeTaskDetails sets ExcludeTaskDetails field to given value.

### HasExcludeTaskDetails

`func (o *HistoricDetailQueryDto) HasExcludeTaskDetails() bool`

HasExcludeTaskDetails returns a boolean if a field has been set.

### SetExcludeTaskDetailsNil

`func (o *HistoricDetailQueryDto) SetExcludeTaskDetailsNil(b bool)`

 SetExcludeTaskDetailsNil sets the value for ExcludeTaskDetails to be an explicit nil

### UnsetExcludeTaskDetails
`func (o *HistoricDetailQueryDto) UnsetExcludeTaskDetails()`

UnsetExcludeTaskDetails ensures that no value is present for ExcludeTaskDetails, not even an explicit nil
### GetInitial

`func (o *HistoricDetailQueryDto) GetInitial() bool`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *HistoricDetailQueryDto) GetInitialOk() (*bool, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *HistoricDetailQueryDto) SetInitial(v bool)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *HistoricDetailQueryDto) HasInitial() bool`

HasInitial returns a boolean if a field has been set.

### SetInitialNil

`func (o *HistoricDetailQueryDto) SetInitialNil(b bool)`

 SetInitialNil sets the value for Initial to be an explicit nil

### UnsetInitial
`func (o *HistoricDetailQueryDto) UnsetInitial()`

UnsetInitial ensures that no value is present for Initial, not even an explicit nil
### GetOccurredBefore

`func (o *HistoricDetailQueryDto) GetOccurredBefore() time.Time`

GetOccurredBefore returns the OccurredBefore field if non-nil, zero value otherwise.

### GetOccurredBeforeOk

`func (o *HistoricDetailQueryDto) GetOccurredBeforeOk() (*time.Time, bool)`

GetOccurredBeforeOk returns a tuple with the OccurredBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredBefore

`func (o *HistoricDetailQueryDto) SetOccurredBefore(v time.Time)`

SetOccurredBefore sets OccurredBefore field to given value.

### HasOccurredBefore

`func (o *HistoricDetailQueryDto) HasOccurredBefore() bool`

HasOccurredBefore returns a boolean if a field has been set.

### SetOccurredBeforeNil

`func (o *HistoricDetailQueryDto) SetOccurredBeforeNil(b bool)`

 SetOccurredBeforeNil sets the value for OccurredBefore to be an explicit nil

### UnsetOccurredBefore
`func (o *HistoricDetailQueryDto) UnsetOccurredBefore()`

UnsetOccurredBefore ensures that no value is present for OccurredBefore, not even an explicit nil
### GetOccurredAfter

`func (o *HistoricDetailQueryDto) GetOccurredAfter() time.Time`

GetOccurredAfter returns the OccurredAfter field if non-nil, zero value otherwise.

### GetOccurredAfterOk

`func (o *HistoricDetailQueryDto) GetOccurredAfterOk() (*time.Time, bool)`

GetOccurredAfterOk returns a tuple with the OccurredAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAfter

`func (o *HistoricDetailQueryDto) SetOccurredAfter(v time.Time)`

SetOccurredAfter sets OccurredAfter field to given value.

### HasOccurredAfter

`func (o *HistoricDetailQueryDto) HasOccurredAfter() bool`

HasOccurredAfter returns a boolean if a field has been set.

### SetOccurredAfterNil

`func (o *HistoricDetailQueryDto) SetOccurredAfterNil(b bool)`

 SetOccurredAfterNil sets the value for OccurredAfter to be an explicit nil

### UnsetOccurredAfter
`func (o *HistoricDetailQueryDto) UnsetOccurredAfter()`

UnsetOccurredAfter ensures that no value is present for OccurredAfter, not even an explicit nil
### GetSorting

`func (o *HistoricDetailQueryDto) GetSorting() []HistoricDetailQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *HistoricDetailQueryDto) GetSortingOk() (*[]HistoricDetailQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *HistoricDetailQueryDto) SetSorting(v []HistoricDetailQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *HistoricDetailQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *HistoricDetailQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *HistoricDetailQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


