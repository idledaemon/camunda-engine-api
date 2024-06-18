# ExternalTaskQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalTaskId** | Pointer to **NullableString** | Filter by an external task&#39;s id. | [optional] 
**ExternalTaskIdIn** | Pointer to **[]string** | Filter by the comma-separated list of external task ids. | [optional] 
**TopicName** | Pointer to **NullableString** | Filter by an external task topic. | [optional] 
**WorkerId** | Pointer to **NullableString** | Filter by the id of the worker that the task was most recently locked by. | [optional] 
**Locked** | Pointer to **NullableBool** | Only include external tasks that are currently locked (i.e., they have a lock time and it has not expired). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | [optional] 
**NotLocked** | Pointer to **NullableBool** | Only include external tasks that are currently not locked (i.e., they have no lock or it has expired). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | [optional] 
**WithRetriesLeft** | Pointer to **NullableBool** | Only include external tasks that have a positive (&amp;gt; 0) number of retries (or &#x60;null&#x60;). Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | [optional] 
**NoRetriesLeft** | Pointer to **NullableBool** | Only include external tasks that have 0 retries. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | [optional] 
**LockExpirationAfter** | Pointer to **NullableTime** | Restrict to external tasks that have a lock that expires after a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**LockExpirationBefore** | Pointer to **NullableTime** | Restrict to external tasks that have a lock that expires before a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**ActivityId** | Pointer to **NullableString** | Filter by the id of the activity that an external task is created for. | [optional] 
**ActivityIdIn** | Pointer to **[]string** | Filter by the comma-separated list of ids of the activities that an external task is created for. | [optional] 
**ExecutionId** | Pointer to **NullableString** | Filter by the id of the execution that an external task belongs to. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | Filter by the id of the process instance that an external task belongs to. | [optional] 
**ProcessInstanceIdIn** | Pointer to **[]string** | Filter by a comma-separated list of process instance ids that an external task may belong to. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter by the id of the process definition that an external task belongs to. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Filter by a comma-separated list of tenant ids. An external task must have one of the given tenant ids. | [optional] 
**Active** | Pointer to **NullableBool** | Only include active tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | [optional] 
**Suspended** | Pointer to **NullableBool** | Only include suspended tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; matches any external task. | [optional] 
**PriorityHigherThanOrEquals** | Pointer to **NullableInt64** | Only include jobs with a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | [optional] 
**PriorityLowerThanOrEquals** | Pointer to **NullableInt64** | Only include jobs with a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | [optional] 
**Sorting** | Pointer to [**[]ExternalTaskQueryDtoSortingInner**](ExternalTaskQueryDtoSortingInner.md) | A JSON array of criteria to sort the result by. Each element of the array is a JSON object that                     specifies one ordering. The position in the array identifies the rank of an ordering, i.e., whether                     it is primary, secondary, etc. The ordering objects have the following properties:                      **Note:** The &#x60;sorting&#x60; properties will not be applied to the External Task count query. | [optional] 

## Methods

### NewExternalTaskQueryDto

`func NewExternalTaskQueryDto() *ExternalTaskQueryDto`

NewExternalTaskQueryDto instantiates a new ExternalTaskQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalTaskQueryDtoWithDefaults

`func NewExternalTaskQueryDtoWithDefaults() *ExternalTaskQueryDto`

NewExternalTaskQueryDtoWithDefaults instantiates a new ExternalTaskQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalTaskId

`func (o *ExternalTaskQueryDto) GetExternalTaskId() string`

GetExternalTaskId returns the ExternalTaskId field if non-nil, zero value otherwise.

### GetExternalTaskIdOk

`func (o *ExternalTaskQueryDto) GetExternalTaskIdOk() (*string, bool)`

GetExternalTaskIdOk returns a tuple with the ExternalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTaskId

`func (o *ExternalTaskQueryDto) SetExternalTaskId(v string)`

SetExternalTaskId sets ExternalTaskId field to given value.

### HasExternalTaskId

`func (o *ExternalTaskQueryDto) HasExternalTaskId() bool`

HasExternalTaskId returns a boolean if a field has been set.

### SetExternalTaskIdNil

`func (o *ExternalTaskQueryDto) SetExternalTaskIdNil(b bool)`

 SetExternalTaskIdNil sets the value for ExternalTaskId to be an explicit nil

### UnsetExternalTaskId
`func (o *ExternalTaskQueryDto) UnsetExternalTaskId()`

UnsetExternalTaskId ensures that no value is present for ExternalTaskId, not even an explicit nil
### GetExternalTaskIdIn

`func (o *ExternalTaskQueryDto) GetExternalTaskIdIn() []string`

GetExternalTaskIdIn returns the ExternalTaskIdIn field if non-nil, zero value otherwise.

### GetExternalTaskIdInOk

`func (o *ExternalTaskQueryDto) GetExternalTaskIdInOk() (*[]string, bool)`

GetExternalTaskIdInOk returns a tuple with the ExternalTaskIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTaskIdIn

`func (o *ExternalTaskQueryDto) SetExternalTaskIdIn(v []string)`

SetExternalTaskIdIn sets ExternalTaskIdIn field to given value.

### HasExternalTaskIdIn

`func (o *ExternalTaskQueryDto) HasExternalTaskIdIn() bool`

HasExternalTaskIdIn returns a boolean if a field has been set.

### SetExternalTaskIdInNil

`func (o *ExternalTaskQueryDto) SetExternalTaskIdInNil(b bool)`

 SetExternalTaskIdInNil sets the value for ExternalTaskIdIn to be an explicit nil

### UnsetExternalTaskIdIn
`func (o *ExternalTaskQueryDto) UnsetExternalTaskIdIn()`

UnsetExternalTaskIdIn ensures that no value is present for ExternalTaskIdIn, not even an explicit nil
### GetTopicName

`func (o *ExternalTaskQueryDto) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ExternalTaskQueryDto) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ExternalTaskQueryDto) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *ExternalTaskQueryDto) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### SetTopicNameNil

`func (o *ExternalTaskQueryDto) SetTopicNameNil(b bool)`

 SetTopicNameNil sets the value for TopicName to be an explicit nil

### UnsetTopicName
`func (o *ExternalTaskQueryDto) UnsetTopicName()`

UnsetTopicName ensures that no value is present for TopicName, not even an explicit nil
### GetWorkerId

`func (o *ExternalTaskQueryDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ExternalTaskQueryDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ExternalTaskQueryDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *ExternalTaskQueryDto) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### SetWorkerIdNil

`func (o *ExternalTaskQueryDto) SetWorkerIdNil(b bool)`

 SetWorkerIdNil sets the value for WorkerId to be an explicit nil

### UnsetWorkerId
`func (o *ExternalTaskQueryDto) UnsetWorkerId()`

UnsetWorkerId ensures that no value is present for WorkerId, not even an explicit nil
### GetLocked

`func (o *ExternalTaskQueryDto) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ExternalTaskQueryDto) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ExternalTaskQueryDto) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ExternalTaskQueryDto) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *ExternalTaskQueryDto) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *ExternalTaskQueryDto) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetNotLocked

`func (o *ExternalTaskQueryDto) GetNotLocked() bool`

GetNotLocked returns the NotLocked field if non-nil, zero value otherwise.

### GetNotLockedOk

`func (o *ExternalTaskQueryDto) GetNotLockedOk() (*bool, bool)`

GetNotLockedOk returns a tuple with the NotLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotLocked

`func (o *ExternalTaskQueryDto) SetNotLocked(v bool)`

SetNotLocked sets NotLocked field to given value.

### HasNotLocked

`func (o *ExternalTaskQueryDto) HasNotLocked() bool`

HasNotLocked returns a boolean if a field has been set.

### SetNotLockedNil

`func (o *ExternalTaskQueryDto) SetNotLockedNil(b bool)`

 SetNotLockedNil sets the value for NotLocked to be an explicit nil

### UnsetNotLocked
`func (o *ExternalTaskQueryDto) UnsetNotLocked()`

UnsetNotLocked ensures that no value is present for NotLocked, not even an explicit nil
### GetWithRetriesLeft

`func (o *ExternalTaskQueryDto) GetWithRetriesLeft() bool`

GetWithRetriesLeft returns the WithRetriesLeft field if non-nil, zero value otherwise.

### GetWithRetriesLeftOk

`func (o *ExternalTaskQueryDto) GetWithRetriesLeftOk() (*bool, bool)`

GetWithRetriesLeftOk returns a tuple with the WithRetriesLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithRetriesLeft

`func (o *ExternalTaskQueryDto) SetWithRetriesLeft(v bool)`

SetWithRetriesLeft sets WithRetriesLeft field to given value.

### HasWithRetriesLeft

`func (o *ExternalTaskQueryDto) HasWithRetriesLeft() bool`

HasWithRetriesLeft returns a boolean if a field has been set.

### SetWithRetriesLeftNil

`func (o *ExternalTaskQueryDto) SetWithRetriesLeftNil(b bool)`

 SetWithRetriesLeftNil sets the value for WithRetriesLeft to be an explicit nil

### UnsetWithRetriesLeft
`func (o *ExternalTaskQueryDto) UnsetWithRetriesLeft()`

UnsetWithRetriesLeft ensures that no value is present for WithRetriesLeft, not even an explicit nil
### GetNoRetriesLeft

`func (o *ExternalTaskQueryDto) GetNoRetriesLeft() bool`

GetNoRetriesLeft returns the NoRetriesLeft field if non-nil, zero value otherwise.

### GetNoRetriesLeftOk

`func (o *ExternalTaskQueryDto) GetNoRetriesLeftOk() (*bool, bool)`

GetNoRetriesLeftOk returns a tuple with the NoRetriesLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRetriesLeft

`func (o *ExternalTaskQueryDto) SetNoRetriesLeft(v bool)`

SetNoRetriesLeft sets NoRetriesLeft field to given value.

### HasNoRetriesLeft

`func (o *ExternalTaskQueryDto) HasNoRetriesLeft() bool`

HasNoRetriesLeft returns a boolean if a field has been set.

### SetNoRetriesLeftNil

`func (o *ExternalTaskQueryDto) SetNoRetriesLeftNil(b bool)`

 SetNoRetriesLeftNil sets the value for NoRetriesLeft to be an explicit nil

### UnsetNoRetriesLeft
`func (o *ExternalTaskQueryDto) UnsetNoRetriesLeft()`

UnsetNoRetriesLeft ensures that no value is present for NoRetriesLeft, not even an explicit nil
### GetLockExpirationAfter

`func (o *ExternalTaskQueryDto) GetLockExpirationAfter() time.Time`

GetLockExpirationAfter returns the LockExpirationAfter field if non-nil, zero value otherwise.

### GetLockExpirationAfterOk

`func (o *ExternalTaskQueryDto) GetLockExpirationAfterOk() (*time.Time, bool)`

GetLockExpirationAfterOk returns a tuple with the LockExpirationAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockExpirationAfter

`func (o *ExternalTaskQueryDto) SetLockExpirationAfter(v time.Time)`

SetLockExpirationAfter sets LockExpirationAfter field to given value.

### HasLockExpirationAfter

`func (o *ExternalTaskQueryDto) HasLockExpirationAfter() bool`

HasLockExpirationAfter returns a boolean if a field has been set.

### SetLockExpirationAfterNil

`func (o *ExternalTaskQueryDto) SetLockExpirationAfterNil(b bool)`

 SetLockExpirationAfterNil sets the value for LockExpirationAfter to be an explicit nil

### UnsetLockExpirationAfter
`func (o *ExternalTaskQueryDto) UnsetLockExpirationAfter()`

UnsetLockExpirationAfter ensures that no value is present for LockExpirationAfter, not even an explicit nil
### GetLockExpirationBefore

`func (o *ExternalTaskQueryDto) GetLockExpirationBefore() time.Time`

GetLockExpirationBefore returns the LockExpirationBefore field if non-nil, zero value otherwise.

### GetLockExpirationBeforeOk

`func (o *ExternalTaskQueryDto) GetLockExpirationBeforeOk() (*time.Time, bool)`

GetLockExpirationBeforeOk returns a tuple with the LockExpirationBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockExpirationBefore

`func (o *ExternalTaskQueryDto) SetLockExpirationBefore(v time.Time)`

SetLockExpirationBefore sets LockExpirationBefore field to given value.

### HasLockExpirationBefore

`func (o *ExternalTaskQueryDto) HasLockExpirationBefore() bool`

HasLockExpirationBefore returns a boolean if a field has been set.

### SetLockExpirationBeforeNil

`func (o *ExternalTaskQueryDto) SetLockExpirationBeforeNil(b bool)`

 SetLockExpirationBeforeNil sets the value for LockExpirationBefore to be an explicit nil

### UnsetLockExpirationBefore
`func (o *ExternalTaskQueryDto) UnsetLockExpirationBefore()`

UnsetLockExpirationBefore ensures that no value is present for LockExpirationBefore, not even an explicit nil
### GetActivityId

`func (o *ExternalTaskQueryDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *ExternalTaskQueryDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *ExternalTaskQueryDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *ExternalTaskQueryDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *ExternalTaskQueryDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *ExternalTaskQueryDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetActivityIdIn

`func (o *ExternalTaskQueryDto) GetActivityIdIn() []string`

GetActivityIdIn returns the ActivityIdIn field if non-nil, zero value otherwise.

### GetActivityIdInOk

`func (o *ExternalTaskQueryDto) GetActivityIdInOk() (*[]string, bool)`

GetActivityIdInOk returns a tuple with the ActivityIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIdIn

`func (o *ExternalTaskQueryDto) SetActivityIdIn(v []string)`

SetActivityIdIn sets ActivityIdIn field to given value.

### HasActivityIdIn

`func (o *ExternalTaskQueryDto) HasActivityIdIn() bool`

HasActivityIdIn returns a boolean if a field has been set.

### SetActivityIdInNil

`func (o *ExternalTaskQueryDto) SetActivityIdInNil(b bool)`

 SetActivityIdInNil sets the value for ActivityIdIn to be an explicit nil

### UnsetActivityIdIn
`func (o *ExternalTaskQueryDto) UnsetActivityIdIn()`

UnsetActivityIdIn ensures that no value is present for ActivityIdIn, not even an explicit nil
### GetExecutionId

`func (o *ExternalTaskQueryDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ExternalTaskQueryDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ExternalTaskQueryDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ExternalTaskQueryDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *ExternalTaskQueryDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *ExternalTaskQueryDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetProcessInstanceId

`func (o *ExternalTaskQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *ExternalTaskQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *ExternalTaskQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *ExternalTaskQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *ExternalTaskQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *ExternalTaskQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessInstanceIdIn

`func (o *ExternalTaskQueryDto) GetProcessInstanceIdIn() []string`

GetProcessInstanceIdIn returns the ProcessInstanceIdIn field if non-nil, zero value otherwise.

### GetProcessInstanceIdInOk

`func (o *ExternalTaskQueryDto) GetProcessInstanceIdInOk() (*[]string, bool)`

GetProcessInstanceIdInOk returns a tuple with the ProcessInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIdIn

`func (o *ExternalTaskQueryDto) SetProcessInstanceIdIn(v []string)`

SetProcessInstanceIdIn sets ProcessInstanceIdIn field to given value.

### HasProcessInstanceIdIn

`func (o *ExternalTaskQueryDto) HasProcessInstanceIdIn() bool`

HasProcessInstanceIdIn returns a boolean if a field has been set.

### SetProcessInstanceIdInNil

`func (o *ExternalTaskQueryDto) SetProcessInstanceIdInNil(b bool)`

 SetProcessInstanceIdInNil sets the value for ProcessInstanceIdIn to be an explicit nil

### UnsetProcessInstanceIdIn
`func (o *ExternalTaskQueryDto) UnsetProcessInstanceIdIn()`

UnsetProcessInstanceIdIn ensures that no value is present for ProcessInstanceIdIn, not even an explicit nil
### GetProcessDefinitionId

`func (o *ExternalTaskQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *ExternalTaskQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *ExternalTaskQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *ExternalTaskQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *ExternalTaskQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *ExternalTaskQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetTenantIdIn

`func (o *ExternalTaskQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *ExternalTaskQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *ExternalTaskQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *ExternalTaskQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *ExternalTaskQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *ExternalTaskQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetActive

`func (o *ExternalTaskQueryDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ExternalTaskQueryDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ExternalTaskQueryDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ExternalTaskQueryDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *ExternalTaskQueryDto) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *ExternalTaskQueryDto) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetSuspended

`func (o *ExternalTaskQueryDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ExternalTaskQueryDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ExternalTaskQueryDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ExternalTaskQueryDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *ExternalTaskQueryDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *ExternalTaskQueryDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetPriorityHigherThanOrEquals

`func (o *ExternalTaskQueryDto) GetPriorityHigherThanOrEquals() int64`

GetPriorityHigherThanOrEquals returns the PriorityHigherThanOrEquals field if non-nil, zero value otherwise.

### GetPriorityHigherThanOrEqualsOk

`func (o *ExternalTaskQueryDto) GetPriorityHigherThanOrEqualsOk() (*int64, bool)`

GetPriorityHigherThanOrEqualsOk returns a tuple with the PriorityHigherThanOrEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityHigherThanOrEquals

`func (o *ExternalTaskQueryDto) SetPriorityHigherThanOrEquals(v int64)`

SetPriorityHigherThanOrEquals sets PriorityHigherThanOrEquals field to given value.

### HasPriorityHigherThanOrEquals

`func (o *ExternalTaskQueryDto) HasPriorityHigherThanOrEquals() bool`

HasPriorityHigherThanOrEquals returns a boolean if a field has been set.

### SetPriorityHigherThanOrEqualsNil

`func (o *ExternalTaskQueryDto) SetPriorityHigherThanOrEqualsNil(b bool)`

 SetPriorityHigherThanOrEqualsNil sets the value for PriorityHigherThanOrEquals to be an explicit nil

### UnsetPriorityHigherThanOrEquals
`func (o *ExternalTaskQueryDto) UnsetPriorityHigherThanOrEquals()`

UnsetPriorityHigherThanOrEquals ensures that no value is present for PriorityHigherThanOrEquals, not even an explicit nil
### GetPriorityLowerThanOrEquals

`func (o *ExternalTaskQueryDto) GetPriorityLowerThanOrEquals() int64`

GetPriorityLowerThanOrEquals returns the PriorityLowerThanOrEquals field if non-nil, zero value otherwise.

### GetPriorityLowerThanOrEqualsOk

`func (o *ExternalTaskQueryDto) GetPriorityLowerThanOrEqualsOk() (*int64, bool)`

GetPriorityLowerThanOrEqualsOk returns a tuple with the PriorityLowerThanOrEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLowerThanOrEquals

`func (o *ExternalTaskQueryDto) SetPriorityLowerThanOrEquals(v int64)`

SetPriorityLowerThanOrEquals sets PriorityLowerThanOrEquals field to given value.

### HasPriorityLowerThanOrEquals

`func (o *ExternalTaskQueryDto) HasPriorityLowerThanOrEquals() bool`

HasPriorityLowerThanOrEquals returns a boolean if a field has been set.

### SetPriorityLowerThanOrEqualsNil

`func (o *ExternalTaskQueryDto) SetPriorityLowerThanOrEqualsNil(b bool)`

 SetPriorityLowerThanOrEqualsNil sets the value for PriorityLowerThanOrEquals to be an explicit nil

### UnsetPriorityLowerThanOrEquals
`func (o *ExternalTaskQueryDto) UnsetPriorityLowerThanOrEquals()`

UnsetPriorityLowerThanOrEquals ensures that no value is present for PriorityLowerThanOrEquals, not even an explicit nil
### GetSorting

`func (o *ExternalTaskQueryDto) GetSorting() []ExternalTaskQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *ExternalTaskQueryDto) GetSortingOk() (*[]ExternalTaskQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *ExternalTaskQueryDto) SetSorting(v []ExternalTaskQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *ExternalTaskQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *ExternalTaskQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *ExternalTaskQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


