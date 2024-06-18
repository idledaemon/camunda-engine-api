# ExternalTaskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | Pointer to **NullableString** | The id of the activity that this external task belongs to. | [optional] 
**ActivityInstanceId** | Pointer to **NullableString** | The id of the activity instance that the external task belongs to. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | The full error message submitted with the latest reported failure executing this task; &#x60;null&#x60; if no failure was reported previously or if no error message was submitted | [optional] 
**ExecutionId** | Pointer to **NullableString** | The id of the execution that the external task belongs to. | [optional] 
**Id** | Pointer to **NullableString** | The id of the external task. | [optional] 
**LockExpirationTime** | Pointer to **NullableTime** | The date that the task&#39;s most recent lock expires or has expired. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition the external task is defined in. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition the external task is defined in. | [optional] 
**ProcessDefinitionVersionTag** | Pointer to **NullableString** | The version tag of the process definition the external task is defined in. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance the external task belongs to. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant the external task belongs to. | [optional] 
**Retries** | Pointer to **NullableInt32** | The number of retries the task currently has left. | [optional] 
**Suspended** | Pointer to **NullableBool** | A flag indicating whether the external task is suspended or not. | [optional] 
**WorkerId** | Pointer to **NullableString** | The id of the worker that posesses or posessed the most recent lock. | [optional] 
**TopicName** | Pointer to **NullableString** | The topic name of the external task. | [optional] 
**Priority** | Pointer to **NullableInt64** | The priority of the external task. | [optional] 
**BusinessKey** | Pointer to **NullableString** | The business key of the process instance the external task belongs to. | [optional] 

## Methods

### NewExternalTaskDto

`func NewExternalTaskDto() *ExternalTaskDto`

NewExternalTaskDto instantiates a new ExternalTaskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalTaskDtoWithDefaults

`func NewExternalTaskDtoWithDefaults() *ExternalTaskDto`

NewExternalTaskDtoWithDefaults instantiates a new ExternalTaskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *ExternalTaskDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *ExternalTaskDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *ExternalTaskDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *ExternalTaskDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *ExternalTaskDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *ExternalTaskDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetActivityInstanceId

`func (o *ExternalTaskDto) GetActivityInstanceId() string`

GetActivityInstanceId returns the ActivityInstanceId field if non-nil, zero value otherwise.

### GetActivityInstanceIdOk

`func (o *ExternalTaskDto) GetActivityInstanceIdOk() (*string, bool)`

GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceId

`func (o *ExternalTaskDto) SetActivityInstanceId(v string)`

SetActivityInstanceId sets ActivityInstanceId field to given value.

### HasActivityInstanceId

`func (o *ExternalTaskDto) HasActivityInstanceId() bool`

HasActivityInstanceId returns a boolean if a field has been set.

### SetActivityInstanceIdNil

`func (o *ExternalTaskDto) SetActivityInstanceIdNil(b bool)`

 SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil

### UnsetActivityInstanceId
`func (o *ExternalTaskDto) UnsetActivityInstanceId()`

UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
### GetErrorMessage

`func (o *ExternalTaskDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ExternalTaskDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ExternalTaskDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ExternalTaskDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ExternalTaskDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ExternalTaskDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetExecutionId

`func (o *ExternalTaskDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ExternalTaskDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ExternalTaskDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ExternalTaskDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *ExternalTaskDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *ExternalTaskDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetId

`func (o *ExternalTaskDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalTaskDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalTaskDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalTaskDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExternalTaskDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExternalTaskDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLockExpirationTime

`func (o *ExternalTaskDto) GetLockExpirationTime() time.Time`

GetLockExpirationTime returns the LockExpirationTime field if non-nil, zero value otherwise.

### GetLockExpirationTimeOk

`func (o *ExternalTaskDto) GetLockExpirationTimeOk() (*time.Time, bool)`

GetLockExpirationTimeOk returns a tuple with the LockExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockExpirationTime

`func (o *ExternalTaskDto) SetLockExpirationTime(v time.Time)`

SetLockExpirationTime sets LockExpirationTime field to given value.

### HasLockExpirationTime

`func (o *ExternalTaskDto) HasLockExpirationTime() bool`

HasLockExpirationTime returns a boolean if a field has been set.

### SetLockExpirationTimeNil

`func (o *ExternalTaskDto) SetLockExpirationTimeNil(b bool)`

 SetLockExpirationTimeNil sets the value for LockExpirationTime to be an explicit nil

### UnsetLockExpirationTime
`func (o *ExternalTaskDto) UnsetLockExpirationTime()`

UnsetLockExpirationTime ensures that no value is present for LockExpirationTime, not even an explicit nil
### GetProcessDefinitionId

`func (o *ExternalTaskDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *ExternalTaskDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *ExternalTaskDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *ExternalTaskDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *ExternalTaskDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *ExternalTaskDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *ExternalTaskDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *ExternalTaskDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *ExternalTaskDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *ExternalTaskDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *ExternalTaskDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *ExternalTaskDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionVersionTag

`func (o *ExternalTaskDto) GetProcessDefinitionVersionTag() string`

GetProcessDefinitionVersionTag returns the ProcessDefinitionVersionTag field if non-nil, zero value otherwise.

### GetProcessDefinitionVersionTagOk

`func (o *ExternalTaskDto) GetProcessDefinitionVersionTagOk() (*string, bool)`

GetProcessDefinitionVersionTagOk returns a tuple with the ProcessDefinitionVersionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionVersionTag

`func (o *ExternalTaskDto) SetProcessDefinitionVersionTag(v string)`

SetProcessDefinitionVersionTag sets ProcessDefinitionVersionTag field to given value.

### HasProcessDefinitionVersionTag

`func (o *ExternalTaskDto) HasProcessDefinitionVersionTag() bool`

HasProcessDefinitionVersionTag returns a boolean if a field has been set.

### SetProcessDefinitionVersionTagNil

`func (o *ExternalTaskDto) SetProcessDefinitionVersionTagNil(b bool)`

 SetProcessDefinitionVersionTagNil sets the value for ProcessDefinitionVersionTag to be an explicit nil

### UnsetProcessDefinitionVersionTag
`func (o *ExternalTaskDto) UnsetProcessDefinitionVersionTag()`

UnsetProcessDefinitionVersionTag ensures that no value is present for ProcessDefinitionVersionTag, not even an explicit nil
### GetProcessInstanceId

`func (o *ExternalTaskDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *ExternalTaskDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *ExternalTaskDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *ExternalTaskDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *ExternalTaskDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *ExternalTaskDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetTenantId

`func (o *ExternalTaskDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExternalTaskDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExternalTaskDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ExternalTaskDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ExternalTaskDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ExternalTaskDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetRetries

`func (o *ExternalTaskDto) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *ExternalTaskDto) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *ExternalTaskDto) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *ExternalTaskDto) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *ExternalTaskDto) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *ExternalTaskDto) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetSuspended

`func (o *ExternalTaskDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ExternalTaskDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ExternalTaskDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ExternalTaskDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *ExternalTaskDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *ExternalTaskDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetWorkerId

`func (o *ExternalTaskDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ExternalTaskDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ExternalTaskDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *ExternalTaskDto) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### SetWorkerIdNil

`func (o *ExternalTaskDto) SetWorkerIdNil(b bool)`

 SetWorkerIdNil sets the value for WorkerId to be an explicit nil

### UnsetWorkerId
`func (o *ExternalTaskDto) UnsetWorkerId()`

UnsetWorkerId ensures that no value is present for WorkerId, not even an explicit nil
### GetTopicName

`func (o *ExternalTaskDto) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ExternalTaskDto) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ExternalTaskDto) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *ExternalTaskDto) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### SetTopicNameNil

`func (o *ExternalTaskDto) SetTopicNameNil(b bool)`

 SetTopicNameNil sets the value for TopicName to be an explicit nil

### UnsetTopicName
`func (o *ExternalTaskDto) UnsetTopicName()`

UnsetTopicName ensures that no value is present for TopicName, not even an explicit nil
### GetPriority

`func (o *ExternalTaskDto) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ExternalTaskDto) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ExternalTaskDto) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ExternalTaskDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *ExternalTaskDto) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *ExternalTaskDto) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetBusinessKey

`func (o *ExternalTaskDto) GetBusinessKey() string`

GetBusinessKey returns the BusinessKey field if non-nil, zero value otherwise.

### GetBusinessKeyOk

`func (o *ExternalTaskDto) GetBusinessKeyOk() (*string, bool)`

GetBusinessKeyOk returns a tuple with the BusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKey

`func (o *ExternalTaskDto) SetBusinessKey(v string)`

SetBusinessKey sets BusinessKey field to given value.

### HasBusinessKey

`func (o *ExternalTaskDto) HasBusinessKey() bool`

HasBusinessKey returns a boolean if a field has been set.

### SetBusinessKeyNil

`func (o *ExternalTaskDto) SetBusinessKeyNil(b bool)`

 SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil

### UnsetBusinessKey
`func (o *ExternalTaskDto) UnsetBusinessKey()`

UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


