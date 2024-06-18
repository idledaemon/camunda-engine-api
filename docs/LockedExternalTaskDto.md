# LockedExternalTaskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | Pointer to **NullableString** | The id of the activity that this external task belongs to. | [optional] 
**ActivityInstanceId** | Pointer to **NullableString** | The id of the activity instance that the external task belongs to. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | The full error message submitted with the latest reported failure executing this task;&#x60;null&#x60; if no failure was reported previously or if no error message was submitted | [optional] 
**ErrorDetails** | Pointer to **NullableString** | The error details submitted with the latest reported failure executing this task.&#x60;null&#x60; if no failure was reported previously or if no error details was submitted | [optional] 
**ExecutionId** | Pointer to **NullableString** | The id of the execution that the external task belongs to. | [optional] 
**Id** | Pointer to **NullableString** | The id of the external task. | [optional] 
**LockExpirationTime** | Pointer to **NullableTime** | The date that the task&#39;s most recent lock expires or has expired. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition the external task is defined in. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition the external task is defined in. | [optional] 
**ProcessDefinitionVersionTag** | Pointer to **NullableString** | The version tag of the process definition the external task is defined in. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance the external task belongs to. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant the external task belongs to. | [optional] 
**Retries** | Pointer to **NullableInt32** | The number of retries the task currently has left. | [optional] 
**Suspended** | Pointer to **NullableBool** | Whether the process instance the external task belongs to is suspended. | [optional] 
**WorkerId** | Pointer to **NullableString** | The id of the worker that posesses or posessed the most recent lock. | [optional] 
**Priority** | Pointer to **NullableInt64** | The priority of the external task. | [optional] 
**TopicName** | Pointer to **NullableString** | The topic name of the external task. | [optional] 
**BusinessKey** | Pointer to **NullableString** | The business key of the process instance the external task belongs to. | [optional] 
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A JSON object containing a property for each of the requested variables. The key is the variable name, the value is a JSON object of serialized variable values with the following properties: | [optional] 

## Methods

### NewLockedExternalTaskDto

`func NewLockedExternalTaskDto() *LockedExternalTaskDto`

NewLockedExternalTaskDto instantiates a new LockedExternalTaskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockedExternalTaskDtoWithDefaults

`func NewLockedExternalTaskDtoWithDefaults() *LockedExternalTaskDto`

NewLockedExternalTaskDtoWithDefaults instantiates a new LockedExternalTaskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *LockedExternalTaskDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *LockedExternalTaskDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *LockedExternalTaskDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *LockedExternalTaskDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *LockedExternalTaskDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *LockedExternalTaskDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetActivityInstanceId

`func (o *LockedExternalTaskDto) GetActivityInstanceId() string`

GetActivityInstanceId returns the ActivityInstanceId field if non-nil, zero value otherwise.

### GetActivityInstanceIdOk

`func (o *LockedExternalTaskDto) GetActivityInstanceIdOk() (*string, bool)`

GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceId

`func (o *LockedExternalTaskDto) SetActivityInstanceId(v string)`

SetActivityInstanceId sets ActivityInstanceId field to given value.

### HasActivityInstanceId

`func (o *LockedExternalTaskDto) HasActivityInstanceId() bool`

HasActivityInstanceId returns a boolean if a field has been set.

### SetActivityInstanceIdNil

`func (o *LockedExternalTaskDto) SetActivityInstanceIdNil(b bool)`

 SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil

### UnsetActivityInstanceId
`func (o *LockedExternalTaskDto) UnsetActivityInstanceId()`

UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
### GetErrorMessage

`func (o *LockedExternalTaskDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *LockedExternalTaskDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *LockedExternalTaskDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *LockedExternalTaskDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *LockedExternalTaskDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *LockedExternalTaskDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetErrorDetails

`func (o *LockedExternalTaskDto) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *LockedExternalTaskDto) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *LockedExternalTaskDto) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *LockedExternalTaskDto) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### SetErrorDetailsNil

`func (o *LockedExternalTaskDto) SetErrorDetailsNil(b bool)`

 SetErrorDetailsNil sets the value for ErrorDetails to be an explicit nil

### UnsetErrorDetails
`func (o *LockedExternalTaskDto) UnsetErrorDetails()`

UnsetErrorDetails ensures that no value is present for ErrorDetails, not even an explicit nil
### GetExecutionId

`func (o *LockedExternalTaskDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *LockedExternalTaskDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *LockedExternalTaskDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *LockedExternalTaskDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *LockedExternalTaskDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *LockedExternalTaskDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetId

`func (o *LockedExternalTaskDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LockedExternalTaskDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LockedExternalTaskDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LockedExternalTaskDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LockedExternalTaskDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LockedExternalTaskDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLockExpirationTime

`func (o *LockedExternalTaskDto) GetLockExpirationTime() time.Time`

GetLockExpirationTime returns the LockExpirationTime field if non-nil, zero value otherwise.

### GetLockExpirationTimeOk

`func (o *LockedExternalTaskDto) GetLockExpirationTimeOk() (*time.Time, bool)`

GetLockExpirationTimeOk returns a tuple with the LockExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockExpirationTime

`func (o *LockedExternalTaskDto) SetLockExpirationTime(v time.Time)`

SetLockExpirationTime sets LockExpirationTime field to given value.

### HasLockExpirationTime

`func (o *LockedExternalTaskDto) HasLockExpirationTime() bool`

HasLockExpirationTime returns a boolean if a field has been set.

### SetLockExpirationTimeNil

`func (o *LockedExternalTaskDto) SetLockExpirationTimeNil(b bool)`

 SetLockExpirationTimeNil sets the value for LockExpirationTime to be an explicit nil

### UnsetLockExpirationTime
`func (o *LockedExternalTaskDto) UnsetLockExpirationTime()`

UnsetLockExpirationTime ensures that no value is present for LockExpirationTime, not even an explicit nil
### GetProcessDefinitionId

`func (o *LockedExternalTaskDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *LockedExternalTaskDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *LockedExternalTaskDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *LockedExternalTaskDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *LockedExternalTaskDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *LockedExternalTaskDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *LockedExternalTaskDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *LockedExternalTaskDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *LockedExternalTaskDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *LockedExternalTaskDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *LockedExternalTaskDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *LockedExternalTaskDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionVersionTag

`func (o *LockedExternalTaskDto) GetProcessDefinitionVersionTag() string`

GetProcessDefinitionVersionTag returns the ProcessDefinitionVersionTag field if non-nil, zero value otherwise.

### GetProcessDefinitionVersionTagOk

`func (o *LockedExternalTaskDto) GetProcessDefinitionVersionTagOk() (*string, bool)`

GetProcessDefinitionVersionTagOk returns a tuple with the ProcessDefinitionVersionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionVersionTag

`func (o *LockedExternalTaskDto) SetProcessDefinitionVersionTag(v string)`

SetProcessDefinitionVersionTag sets ProcessDefinitionVersionTag field to given value.

### HasProcessDefinitionVersionTag

`func (o *LockedExternalTaskDto) HasProcessDefinitionVersionTag() bool`

HasProcessDefinitionVersionTag returns a boolean if a field has been set.

### SetProcessDefinitionVersionTagNil

`func (o *LockedExternalTaskDto) SetProcessDefinitionVersionTagNil(b bool)`

 SetProcessDefinitionVersionTagNil sets the value for ProcessDefinitionVersionTag to be an explicit nil

### UnsetProcessDefinitionVersionTag
`func (o *LockedExternalTaskDto) UnsetProcessDefinitionVersionTag()`

UnsetProcessDefinitionVersionTag ensures that no value is present for ProcessDefinitionVersionTag, not even an explicit nil
### GetProcessInstanceId

`func (o *LockedExternalTaskDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *LockedExternalTaskDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *LockedExternalTaskDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *LockedExternalTaskDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *LockedExternalTaskDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *LockedExternalTaskDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetTenantId

`func (o *LockedExternalTaskDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LockedExternalTaskDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LockedExternalTaskDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LockedExternalTaskDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LockedExternalTaskDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LockedExternalTaskDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetRetries

`func (o *LockedExternalTaskDto) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *LockedExternalTaskDto) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *LockedExternalTaskDto) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *LockedExternalTaskDto) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *LockedExternalTaskDto) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *LockedExternalTaskDto) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetSuspended

`func (o *LockedExternalTaskDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *LockedExternalTaskDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *LockedExternalTaskDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *LockedExternalTaskDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *LockedExternalTaskDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *LockedExternalTaskDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetWorkerId

`func (o *LockedExternalTaskDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *LockedExternalTaskDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *LockedExternalTaskDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *LockedExternalTaskDto) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### SetWorkerIdNil

`func (o *LockedExternalTaskDto) SetWorkerIdNil(b bool)`

 SetWorkerIdNil sets the value for WorkerId to be an explicit nil

### UnsetWorkerId
`func (o *LockedExternalTaskDto) UnsetWorkerId()`

UnsetWorkerId ensures that no value is present for WorkerId, not even an explicit nil
### GetPriority

`func (o *LockedExternalTaskDto) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *LockedExternalTaskDto) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *LockedExternalTaskDto) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *LockedExternalTaskDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *LockedExternalTaskDto) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *LockedExternalTaskDto) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetTopicName

`func (o *LockedExternalTaskDto) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *LockedExternalTaskDto) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *LockedExternalTaskDto) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *LockedExternalTaskDto) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### SetTopicNameNil

`func (o *LockedExternalTaskDto) SetTopicNameNil(b bool)`

 SetTopicNameNil sets the value for TopicName to be an explicit nil

### UnsetTopicName
`func (o *LockedExternalTaskDto) UnsetTopicName()`

UnsetTopicName ensures that no value is present for TopicName, not even an explicit nil
### GetBusinessKey

`func (o *LockedExternalTaskDto) GetBusinessKey() string`

GetBusinessKey returns the BusinessKey field if non-nil, zero value otherwise.

### GetBusinessKeyOk

`func (o *LockedExternalTaskDto) GetBusinessKeyOk() (*string, bool)`

GetBusinessKeyOk returns a tuple with the BusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKey

`func (o *LockedExternalTaskDto) SetBusinessKey(v string)`

SetBusinessKey sets BusinessKey field to given value.

### HasBusinessKey

`func (o *LockedExternalTaskDto) HasBusinessKey() bool`

HasBusinessKey returns a boolean if a field has been set.

### SetBusinessKeyNil

`func (o *LockedExternalTaskDto) SetBusinessKeyNil(b bool)`

 SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil

### UnsetBusinessKey
`func (o *LockedExternalTaskDto) UnsetBusinessKey()`

UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
### GetVariables

`func (o *LockedExternalTaskDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *LockedExternalTaskDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *LockedExternalTaskDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *LockedExternalTaskDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *LockedExternalTaskDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *LockedExternalTaskDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


