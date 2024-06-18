# HistoricExternalTaskLogQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogId** | Pointer to **NullableString** | Filter by historic external task log id. | [optional] 
**ExternalTaskId** | Pointer to **NullableString** | Filter by external task id. | [optional] 
**TopicName** | Pointer to **NullableString** | Filter by an external task topic. | [optional] 
**WorkerId** | Pointer to **NullableString** | Filter by the id of the worker that the task was most recently locked by. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Filter by external task exception message. | [optional] 
**ActivityIdIn** | Pointer to **[]string** | Only include historic external task logs which belong to one of the passed activity ids. | [optional] 
**ActivityInstanceIdIn** | Pointer to **[]string** | Only include historic external task logs which belong to one of the passed activity instance ids. | [optional] 
**ExecutionIdIn** | Pointer to **[]string** | Only include historic external task logs which belong to one of the passed execution ids. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | Filter by process instance id. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter by process definition id. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Filter by process definition key. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Only include historic external task log entries which belong to one of the passed and comma-separated tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include historic external task log entries that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**PriorityLowerThanOrEquals** | Pointer to **NullableInt64** | Only include logs for which the associated external task had a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | [optional] 
**PriorityHigherThanOrEquals** | Pointer to **NullableInt64** | Only include logs for which the associated external task had a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | [optional] 
**CreationLog** | Pointer to **NullableBool** | Only include creation logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**FailureLog** | Pointer to **NullableBool** | Only include failure logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**SuccessLog** | Pointer to **NullableBool** | Only include success logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**DeletionLog** | Pointer to **NullableBool** | Only include deletion logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Sorting** | Pointer to [**[]HistoricExternalTaskLogQueryDtoSortingInner**](HistoricExternalTaskLogQueryDtoSortingInner.md) | An array of criteria to sort the result by. Each element of the array is                        an object that specifies one ordering. The position in the array                        identifies the rank of an ordering, i.e., whether it is primary, secondary,                        etc. Sorting has no effect for &#x60;count&#x60; endpoints. | [optional] 

## Methods

### NewHistoricExternalTaskLogQueryDto

`func NewHistoricExternalTaskLogQueryDto() *HistoricExternalTaskLogQueryDto`

NewHistoricExternalTaskLogQueryDto instantiates a new HistoricExternalTaskLogQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricExternalTaskLogQueryDtoWithDefaults

`func NewHistoricExternalTaskLogQueryDtoWithDefaults() *HistoricExternalTaskLogQueryDto`

NewHistoricExternalTaskLogQueryDtoWithDefaults instantiates a new HistoricExternalTaskLogQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogId

`func (o *HistoricExternalTaskLogQueryDto) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *HistoricExternalTaskLogQueryDto) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *HistoricExternalTaskLogQueryDto) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *HistoricExternalTaskLogQueryDto) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *HistoricExternalTaskLogQueryDto) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *HistoricExternalTaskLogQueryDto) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetExternalTaskId

`func (o *HistoricExternalTaskLogQueryDto) GetExternalTaskId() string`

GetExternalTaskId returns the ExternalTaskId field if non-nil, zero value otherwise.

### GetExternalTaskIdOk

`func (o *HistoricExternalTaskLogQueryDto) GetExternalTaskIdOk() (*string, bool)`

GetExternalTaskIdOk returns a tuple with the ExternalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTaskId

`func (o *HistoricExternalTaskLogQueryDto) SetExternalTaskId(v string)`

SetExternalTaskId sets ExternalTaskId field to given value.

### HasExternalTaskId

`func (o *HistoricExternalTaskLogQueryDto) HasExternalTaskId() bool`

HasExternalTaskId returns a boolean if a field has been set.

### SetExternalTaskIdNil

`func (o *HistoricExternalTaskLogQueryDto) SetExternalTaskIdNil(b bool)`

 SetExternalTaskIdNil sets the value for ExternalTaskId to be an explicit nil

### UnsetExternalTaskId
`func (o *HistoricExternalTaskLogQueryDto) UnsetExternalTaskId()`

UnsetExternalTaskId ensures that no value is present for ExternalTaskId, not even an explicit nil
### GetTopicName

`func (o *HistoricExternalTaskLogQueryDto) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *HistoricExternalTaskLogQueryDto) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *HistoricExternalTaskLogQueryDto) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *HistoricExternalTaskLogQueryDto) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### SetTopicNameNil

`func (o *HistoricExternalTaskLogQueryDto) SetTopicNameNil(b bool)`

 SetTopicNameNil sets the value for TopicName to be an explicit nil

### UnsetTopicName
`func (o *HistoricExternalTaskLogQueryDto) UnsetTopicName()`

UnsetTopicName ensures that no value is present for TopicName, not even an explicit nil
### GetWorkerId

`func (o *HistoricExternalTaskLogQueryDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *HistoricExternalTaskLogQueryDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *HistoricExternalTaskLogQueryDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *HistoricExternalTaskLogQueryDto) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### SetWorkerIdNil

`func (o *HistoricExternalTaskLogQueryDto) SetWorkerIdNil(b bool)`

 SetWorkerIdNil sets the value for WorkerId to be an explicit nil

### UnsetWorkerId
`func (o *HistoricExternalTaskLogQueryDto) UnsetWorkerId()`

UnsetWorkerId ensures that no value is present for WorkerId, not even an explicit nil
### GetErrorMessage

`func (o *HistoricExternalTaskLogQueryDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *HistoricExternalTaskLogQueryDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *HistoricExternalTaskLogQueryDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *HistoricExternalTaskLogQueryDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *HistoricExternalTaskLogQueryDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *HistoricExternalTaskLogQueryDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetActivityIdIn

`func (o *HistoricExternalTaskLogQueryDto) GetActivityIdIn() []string`

GetActivityIdIn returns the ActivityIdIn field if non-nil, zero value otherwise.

### GetActivityIdInOk

`func (o *HistoricExternalTaskLogQueryDto) GetActivityIdInOk() (*[]string, bool)`

GetActivityIdInOk returns a tuple with the ActivityIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIdIn

`func (o *HistoricExternalTaskLogQueryDto) SetActivityIdIn(v []string)`

SetActivityIdIn sets ActivityIdIn field to given value.

### HasActivityIdIn

`func (o *HistoricExternalTaskLogQueryDto) HasActivityIdIn() bool`

HasActivityIdIn returns a boolean if a field has been set.

### SetActivityIdInNil

`func (o *HistoricExternalTaskLogQueryDto) SetActivityIdInNil(b bool)`

 SetActivityIdInNil sets the value for ActivityIdIn to be an explicit nil

### UnsetActivityIdIn
`func (o *HistoricExternalTaskLogQueryDto) UnsetActivityIdIn()`

UnsetActivityIdIn ensures that no value is present for ActivityIdIn, not even an explicit nil
### GetActivityInstanceIdIn

`func (o *HistoricExternalTaskLogQueryDto) GetActivityInstanceIdIn() []string`

GetActivityInstanceIdIn returns the ActivityInstanceIdIn field if non-nil, zero value otherwise.

### GetActivityInstanceIdInOk

`func (o *HistoricExternalTaskLogQueryDto) GetActivityInstanceIdInOk() (*[]string, bool)`

GetActivityInstanceIdInOk returns a tuple with the ActivityInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceIdIn

`func (o *HistoricExternalTaskLogQueryDto) SetActivityInstanceIdIn(v []string)`

SetActivityInstanceIdIn sets ActivityInstanceIdIn field to given value.

### HasActivityInstanceIdIn

`func (o *HistoricExternalTaskLogQueryDto) HasActivityInstanceIdIn() bool`

HasActivityInstanceIdIn returns a boolean if a field has been set.

### SetActivityInstanceIdInNil

`func (o *HistoricExternalTaskLogQueryDto) SetActivityInstanceIdInNil(b bool)`

 SetActivityInstanceIdInNil sets the value for ActivityInstanceIdIn to be an explicit nil

### UnsetActivityInstanceIdIn
`func (o *HistoricExternalTaskLogQueryDto) UnsetActivityInstanceIdIn()`

UnsetActivityInstanceIdIn ensures that no value is present for ActivityInstanceIdIn, not even an explicit nil
### GetExecutionIdIn

`func (o *HistoricExternalTaskLogQueryDto) GetExecutionIdIn() []string`

GetExecutionIdIn returns the ExecutionIdIn field if non-nil, zero value otherwise.

### GetExecutionIdInOk

`func (o *HistoricExternalTaskLogQueryDto) GetExecutionIdInOk() (*[]string, bool)`

GetExecutionIdInOk returns a tuple with the ExecutionIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionIdIn

`func (o *HistoricExternalTaskLogQueryDto) SetExecutionIdIn(v []string)`

SetExecutionIdIn sets ExecutionIdIn field to given value.

### HasExecutionIdIn

`func (o *HistoricExternalTaskLogQueryDto) HasExecutionIdIn() bool`

HasExecutionIdIn returns a boolean if a field has been set.

### SetExecutionIdInNil

`func (o *HistoricExternalTaskLogQueryDto) SetExecutionIdInNil(b bool)`

 SetExecutionIdInNil sets the value for ExecutionIdIn to be an explicit nil

### UnsetExecutionIdIn
`func (o *HistoricExternalTaskLogQueryDto) UnsetExecutionIdIn()`

UnsetExecutionIdIn ensures that no value is present for ExecutionIdIn, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricExternalTaskLogQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricExternalTaskLogQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricExternalTaskLogQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricExternalTaskLogQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricExternalTaskLogQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricExternalTaskLogQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricExternalTaskLogQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricExternalTaskLogQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricExternalTaskLogQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricExternalTaskLogQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricExternalTaskLogQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricExternalTaskLogQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricExternalTaskLogQueryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricExternalTaskLogQueryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricExternalTaskLogQueryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricExternalTaskLogQueryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricExternalTaskLogQueryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricExternalTaskLogQueryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetTenantIdIn

`func (o *HistoricExternalTaskLogQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *HistoricExternalTaskLogQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *HistoricExternalTaskLogQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *HistoricExternalTaskLogQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *HistoricExternalTaskLogQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *HistoricExternalTaskLogQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *HistoricExternalTaskLogQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *HistoricExternalTaskLogQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *HistoricExternalTaskLogQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *HistoricExternalTaskLogQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *HistoricExternalTaskLogQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *HistoricExternalTaskLogQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetPriorityLowerThanOrEquals

`func (o *HistoricExternalTaskLogQueryDto) GetPriorityLowerThanOrEquals() int64`

GetPriorityLowerThanOrEquals returns the PriorityLowerThanOrEquals field if non-nil, zero value otherwise.

### GetPriorityLowerThanOrEqualsOk

`func (o *HistoricExternalTaskLogQueryDto) GetPriorityLowerThanOrEqualsOk() (*int64, bool)`

GetPriorityLowerThanOrEqualsOk returns a tuple with the PriorityLowerThanOrEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLowerThanOrEquals

`func (o *HistoricExternalTaskLogQueryDto) SetPriorityLowerThanOrEquals(v int64)`

SetPriorityLowerThanOrEquals sets PriorityLowerThanOrEquals field to given value.

### HasPriorityLowerThanOrEquals

`func (o *HistoricExternalTaskLogQueryDto) HasPriorityLowerThanOrEquals() bool`

HasPriorityLowerThanOrEquals returns a boolean if a field has been set.

### SetPriorityLowerThanOrEqualsNil

`func (o *HistoricExternalTaskLogQueryDto) SetPriorityLowerThanOrEqualsNil(b bool)`

 SetPriorityLowerThanOrEqualsNil sets the value for PriorityLowerThanOrEquals to be an explicit nil

### UnsetPriorityLowerThanOrEquals
`func (o *HistoricExternalTaskLogQueryDto) UnsetPriorityLowerThanOrEquals()`

UnsetPriorityLowerThanOrEquals ensures that no value is present for PriorityLowerThanOrEquals, not even an explicit nil
### GetPriorityHigherThanOrEquals

`func (o *HistoricExternalTaskLogQueryDto) GetPriorityHigherThanOrEquals() int64`

GetPriorityHigherThanOrEquals returns the PriorityHigherThanOrEquals field if non-nil, zero value otherwise.

### GetPriorityHigherThanOrEqualsOk

`func (o *HistoricExternalTaskLogQueryDto) GetPriorityHigherThanOrEqualsOk() (*int64, bool)`

GetPriorityHigherThanOrEqualsOk returns a tuple with the PriorityHigherThanOrEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityHigherThanOrEquals

`func (o *HistoricExternalTaskLogQueryDto) SetPriorityHigherThanOrEquals(v int64)`

SetPriorityHigherThanOrEquals sets PriorityHigherThanOrEquals field to given value.

### HasPriorityHigherThanOrEquals

`func (o *HistoricExternalTaskLogQueryDto) HasPriorityHigherThanOrEquals() bool`

HasPriorityHigherThanOrEquals returns a boolean if a field has been set.

### SetPriorityHigherThanOrEqualsNil

`func (o *HistoricExternalTaskLogQueryDto) SetPriorityHigherThanOrEqualsNil(b bool)`

 SetPriorityHigherThanOrEqualsNil sets the value for PriorityHigherThanOrEquals to be an explicit nil

### UnsetPriorityHigherThanOrEquals
`func (o *HistoricExternalTaskLogQueryDto) UnsetPriorityHigherThanOrEquals()`

UnsetPriorityHigherThanOrEquals ensures that no value is present for PriorityHigherThanOrEquals, not even an explicit nil
### GetCreationLog

`func (o *HistoricExternalTaskLogQueryDto) GetCreationLog() bool`

GetCreationLog returns the CreationLog field if non-nil, zero value otherwise.

### GetCreationLogOk

`func (o *HistoricExternalTaskLogQueryDto) GetCreationLogOk() (*bool, bool)`

GetCreationLogOk returns a tuple with the CreationLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationLog

`func (o *HistoricExternalTaskLogQueryDto) SetCreationLog(v bool)`

SetCreationLog sets CreationLog field to given value.

### HasCreationLog

`func (o *HistoricExternalTaskLogQueryDto) HasCreationLog() bool`

HasCreationLog returns a boolean if a field has been set.

### SetCreationLogNil

`func (o *HistoricExternalTaskLogQueryDto) SetCreationLogNil(b bool)`

 SetCreationLogNil sets the value for CreationLog to be an explicit nil

### UnsetCreationLog
`func (o *HistoricExternalTaskLogQueryDto) UnsetCreationLog()`

UnsetCreationLog ensures that no value is present for CreationLog, not even an explicit nil
### GetFailureLog

`func (o *HistoricExternalTaskLogQueryDto) GetFailureLog() bool`

GetFailureLog returns the FailureLog field if non-nil, zero value otherwise.

### GetFailureLogOk

`func (o *HistoricExternalTaskLogQueryDto) GetFailureLogOk() (*bool, bool)`

GetFailureLogOk returns a tuple with the FailureLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLog

`func (o *HistoricExternalTaskLogQueryDto) SetFailureLog(v bool)`

SetFailureLog sets FailureLog field to given value.

### HasFailureLog

`func (o *HistoricExternalTaskLogQueryDto) HasFailureLog() bool`

HasFailureLog returns a boolean if a field has been set.

### SetFailureLogNil

`func (o *HistoricExternalTaskLogQueryDto) SetFailureLogNil(b bool)`

 SetFailureLogNil sets the value for FailureLog to be an explicit nil

### UnsetFailureLog
`func (o *HistoricExternalTaskLogQueryDto) UnsetFailureLog()`

UnsetFailureLog ensures that no value is present for FailureLog, not even an explicit nil
### GetSuccessLog

`func (o *HistoricExternalTaskLogQueryDto) GetSuccessLog() bool`

GetSuccessLog returns the SuccessLog field if non-nil, zero value otherwise.

### GetSuccessLogOk

`func (o *HistoricExternalTaskLogQueryDto) GetSuccessLogOk() (*bool, bool)`

GetSuccessLogOk returns a tuple with the SuccessLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessLog

`func (o *HistoricExternalTaskLogQueryDto) SetSuccessLog(v bool)`

SetSuccessLog sets SuccessLog field to given value.

### HasSuccessLog

`func (o *HistoricExternalTaskLogQueryDto) HasSuccessLog() bool`

HasSuccessLog returns a boolean if a field has been set.

### SetSuccessLogNil

`func (o *HistoricExternalTaskLogQueryDto) SetSuccessLogNil(b bool)`

 SetSuccessLogNil sets the value for SuccessLog to be an explicit nil

### UnsetSuccessLog
`func (o *HistoricExternalTaskLogQueryDto) UnsetSuccessLog()`

UnsetSuccessLog ensures that no value is present for SuccessLog, not even an explicit nil
### GetDeletionLog

`func (o *HistoricExternalTaskLogQueryDto) GetDeletionLog() bool`

GetDeletionLog returns the DeletionLog field if non-nil, zero value otherwise.

### GetDeletionLogOk

`func (o *HistoricExternalTaskLogQueryDto) GetDeletionLogOk() (*bool, bool)`

GetDeletionLogOk returns a tuple with the DeletionLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionLog

`func (o *HistoricExternalTaskLogQueryDto) SetDeletionLog(v bool)`

SetDeletionLog sets DeletionLog field to given value.

### HasDeletionLog

`func (o *HistoricExternalTaskLogQueryDto) HasDeletionLog() bool`

HasDeletionLog returns a boolean if a field has been set.

### SetDeletionLogNil

`func (o *HistoricExternalTaskLogQueryDto) SetDeletionLogNil(b bool)`

 SetDeletionLogNil sets the value for DeletionLog to be an explicit nil

### UnsetDeletionLog
`func (o *HistoricExternalTaskLogQueryDto) UnsetDeletionLog()`

UnsetDeletionLog ensures that no value is present for DeletionLog, not even an explicit nil
### GetSorting

`func (o *HistoricExternalTaskLogQueryDto) GetSorting() []HistoricExternalTaskLogQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *HistoricExternalTaskLogQueryDto) GetSortingOk() (*[]HistoricExternalTaskLogQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *HistoricExternalTaskLogQueryDto) SetSorting(v []HistoricExternalTaskLogQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *HistoricExternalTaskLogQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *HistoricExternalTaskLogQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *HistoricExternalTaskLogQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


