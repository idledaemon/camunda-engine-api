# JobQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **NullableString** | Filter by job id. | [optional] 
**JobIds** | Pointer to **[]string** | Filter by a  list of job ids. | [optional] 
**JobDefinitionId** | Pointer to **NullableString** | Only select jobs which exist for the given job definition. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | Only select jobs which exist for the given process instance. | [optional] 
**ProcessInstanceIds** | Pointer to **[]string** | Only select jobs which exist for the given  list of process instance ids. | [optional] 
**ExecutionId** | Pointer to **NullableString** | Only select jobs which exist for the given execution. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter by the id of the process definition the jobs run on. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Filter by the key of the process definition the jobs run on. | [optional] 
**ActivityId** | Pointer to **NullableString** | Only select jobs which exist for an activity with the given id. | [optional] 
**WithRetriesLeft** | Pointer to **NullableBool** | Only select jobs which have retries left. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Executable** | Pointer to **NullableBool** | Only select jobs which are executable, i.e., retries &gt; 0 and due date is &#x60;null&#x60; or due date is in the past. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Timers** | Pointer to **NullableBool** | Only select jobs that are timers. Cannot be used together with &#x60;messages&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Messages** | Pointer to **NullableBool** | Only select jobs that are messages. Cannot be used together with &#x60;timers&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**DueDates** | Pointer to [**[]JobConditionQueryParameterDto**](JobConditionQueryParameterDto.md) | Only select jobs where the due date is lower or higher than the given date.  | [optional] 
**CreateTimes** | Pointer to [**[]JobConditionQueryParameterDto**](JobConditionQueryParameterDto.md) | Only select jobs created before or after the given date.  | [optional] 
**WithException** | Pointer to **NullableBool** | Only select jobs that failed due to an exception. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**ExceptionMessage** | Pointer to **NullableString** | Only select jobs that failed due to an exception with the given message. | [optional] 
**FailedActivityId** | Pointer to **NullableString** | Only select jobs that failed due to an exception at an activity with the given id. | [optional] 
**NoRetriesLeft** | Pointer to **NullableBool** | Only select jobs which have no retries left. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Active** | Pointer to **NullableBool** | Only include active jobs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Suspended** | Pointer to **NullableBool** | Only include suspended jobs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**PriorityLowerThanOrEquals** | Pointer to **NullableInt64** | Only include jobs with a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | [optional] 
**PriorityHigherThanOrEquals** | Pointer to **NullableInt64** | Only include jobs with a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Only include jobs which belong to one of the passed  tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include jobs which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**IncludeJobsWithoutTenantId** | Pointer to **NullableBool** | Include jobs which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Sorting** | Pointer to [**[]JobQueryDtoSortingInner**](JobQueryDtoSortingInner.md) | An array of criteria to sort the result by. Each element of the array is                        an object that specifies one ordering. The position in the array                        identifies the rank of an ordering, i.e., whether it is primary, secondary,                        etc. Does not have an effect for the &#x60;count&#x60; endpoint. | [optional] 

## Methods

### NewJobQueryDto

`func NewJobQueryDto() *JobQueryDto`

NewJobQueryDto instantiates a new JobQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobQueryDtoWithDefaults

`func NewJobQueryDtoWithDefaults() *JobQueryDto`

NewJobQueryDtoWithDefaults instantiates a new JobQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobQueryDto) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobQueryDto) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobQueryDto) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *JobQueryDto) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *JobQueryDto) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *JobQueryDto) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobIds

`func (o *JobQueryDto) GetJobIds() []string`

GetJobIds returns the JobIds field if non-nil, zero value otherwise.

### GetJobIdsOk

`func (o *JobQueryDto) GetJobIdsOk() (*[]string, bool)`

GetJobIdsOk returns a tuple with the JobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIds

`func (o *JobQueryDto) SetJobIds(v []string)`

SetJobIds sets JobIds field to given value.

### HasJobIds

`func (o *JobQueryDto) HasJobIds() bool`

HasJobIds returns a boolean if a field has been set.

### SetJobIdsNil

`func (o *JobQueryDto) SetJobIdsNil(b bool)`

 SetJobIdsNil sets the value for JobIds to be an explicit nil

### UnsetJobIds
`func (o *JobQueryDto) UnsetJobIds()`

UnsetJobIds ensures that no value is present for JobIds, not even an explicit nil
### GetJobDefinitionId

`func (o *JobQueryDto) GetJobDefinitionId() string`

GetJobDefinitionId returns the JobDefinitionId field if non-nil, zero value otherwise.

### GetJobDefinitionIdOk

`func (o *JobQueryDto) GetJobDefinitionIdOk() (*string, bool)`

GetJobDefinitionIdOk returns a tuple with the JobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionId

`func (o *JobQueryDto) SetJobDefinitionId(v string)`

SetJobDefinitionId sets JobDefinitionId field to given value.

### HasJobDefinitionId

`func (o *JobQueryDto) HasJobDefinitionId() bool`

HasJobDefinitionId returns a boolean if a field has been set.

### SetJobDefinitionIdNil

`func (o *JobQueryDto) SetJobDefinitionIdNil(b bool)`

 SetJobDefinitionIdNil sets the value for JobDefinitionId to be an explicit nil

### UnsetJobDefinitionId
`func (o *JobQueryDto) UnsetJobDefinitionId()`

UnsetJobDefinitionId ensures that no value is present for JobDefinitionId, not even an explicit nil
### GetProcessInstanceId

`func (o *JobQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *JobQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *JobQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *JobQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *JobQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *JobQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessInstanceIds

`func (o *JobQueryDto) GetProcessInstanceIds() []string`

GetProcessInstanceIds returns the ProcessInstanceIds field if non-nil, zero value otherwise.

### GetProcessInstanceIdsOk

`func (o *JobQueryDto) GetProcessInstanceIdsOk() (*[]string, bool)`

GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIds

`func (o *JobQueryDto) SetProcessInstanceIds(v []string)`

SetProcessInstanceIds sets ProcessInstanceIds field to given value.

### HasProcessInstanceIds

`func (o *JobQueryDto) HasProcessInstanceIds() bool`

HasProcessInstanceIds returns a boolean if a field has been set.

### SetProcessInstanceIdsNil

`func (o *JobQueryDto) SetProcessInstanceIdsNil(b bool)`

 SetProcessInstanceIdsNil sets the value for ProcessInstanceIds to be an explicit nil

### UnsetProcessInstanceIds
`func (o *JobQueryDto) UnsetProcessInstanceIds()`

UnsetProcessInstanceIds ensures that no value is present for ProcessInstanceIds, not even an explicit nil
### GetExecutionId

`func (o *JobQueryDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *JobQueryDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *JobQueryDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *JobQueryDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *JobQueryDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *JobQueryDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetProcessDefinitionId

`func (o *JobQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *JobQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *JobQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *JobQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *JobQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *JobQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *JobQueryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *JobQueryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *JobQueryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *JobQueryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *JobQueryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *JobQueryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetActivityId

`func (o *JobQueryDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *JobQueryDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *JobQueryDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *JobQueryDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *JobQueryDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *JobQueryDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetWithRetriesLeft

`func (o *JobQueryDto) GetWithRetriesLeft() bool`

GetWithRetriesLeft returns the WithRetriesLeft field if non-nil, zero value otherwise.

### GetWithRetriesLeftOk

`func (o *JobQueryDto) GetWithRetriesLeftOk() (*bool, bool)`

GetWithRetriesLeftOk returns a tuple with the WithRetriesLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithRetriesLeft

`func (o *JobQueryDto) SetWithRetriesLeft(v bool)`

SetWithRetriesLeft sets WithRetriesLeft field to given value.

### HasWithRetriesLeft

`func (o *JobQueryDto) HasWithRetriesLeft() bool`

HasWithRetriesLeft returns a boolean if a field has been set.

### SetWithRetriesLeftNil

`func (o *JobQueryDto) SetWithRetriesLeftNil(b bool)`

 SetWithRetriesLeftNil sets the value for WithRetriesLeft to be an explicit nil

### UnsetWithRetriesLeft
`func (o *JobQueryDto) UnsetWithRetriesLeft()`

UnsetWithRetriesLeft ensures that no value is present for WithRetriesLeft, not even an explicit nil
### GetExecutable

`func (o *JobQueryDto) GetExecutable() bool`

GetExecutable returns the Executable field if non-nil, zero value otherwise.

### GetExecutableOk

`func (o *JobQueryDto) GetExecutableOk() (*bool, bool)`

GetExecutableOk returns a tuple with the Executable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutable

`func (o *JobQueryDto) SetExecutable(v bool)`

SetExecutable sets Executable field to given value.

### HasExecutable

`func (o *JobQueryDto) HasExecutable() bool`

HasExecutable returns a boolean if a field has been set.

### SetExecutableNil

`func (o *JobQueryDto) SetExecutableNil(b bool)`

 SetExecutableNil sets the value for Executable to be an explicit nil

### UnsetExecutable
`func (o *JobQueryDto) UnsetExecutable()`

UnsetExecutable ensures that no value is present for Executable, not even an explicit nil
### GetTimers

`func (o *JobQueryDto) GetTimers() bool`

GetTimers returns the Timers field if non-nil, zero value otherwise.

### GetTimersOk

`func (o *JobQueryDto) GetTimersOk() (*bool, bool)`

GetTimersOk returns a tuple with the Timers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimers

`func (o *JobQueryDto) SetTimers(v bool)`

SetTimers sets Timers field to given value.

### HasTimers

`func (o *JobQueryDto) HasTimers() bool`

HasTimers returns a boolean if a field has been set.

### SetTimersNil

`func (o *JobQueryDto) SetTimersNil(b bool)`

 SetTimersNil sets the value for Timers to be an explicit nil

### UnsetTimers
`func (o *JobQueryDto) UnsetTimers()`

UnsetTimers ensures that no value is present for Timers, not even an explicit nil
### GetMessages

`func (o *JobQueryDto) GetMessages() bool`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *JobQueryDto) GetMessagesOk() (*bool, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *JobQueryDto) SetMessages(v bool)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *JobQueryDto) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *JobQueryDto) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *JobQueryDto) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetDueDates

`func (o *JobQueryDto) GetDueDates() []JobConditionQueryParameterDto`

GetDueDates returns the DueDates field if non-nil, zero value otherwise.

### GetDueDatesOk

`func (o *JobQueryDto) GetDueDatesOk() (*[]JobConditionQueryParameterDto, bool)`

GetDueDatesOk returns a tuple with the DueDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDates

`func (o *JobQueryDto) SetDueDates(v []JobConditionQueryParameterDto)`

SetDueDates sets DueDates field to given value.

### HasDueDates

`func (o *JobQueryDto) HasDueDates() bool`

HasDueDates returns a boolean if a field has been set.

### SetDueDatesNil

`func (o *JobQueryDto) SetDueDatesNil(b bool)`

 SetDueDatesNil sets the value for DueDates to be an explicit nil

### UnsetDueDates
`func (o *JobQueryDto) UnsetDueDates()`

UnsetDueDates ensures that no value is present for DueDates, not even an explicit nil
### GetCreateTimes

`func (o *JobQueryDto) GetCreateTimes() []JobConditionQueryParameterDto`

GetCreateTimes returns the CreateTimes field if non-nil, zero value otherwise.

### GetCreateTimesOk

`func (o *JobQueryDto) GetCreateTimesOk() (*[]JobConditionQueryParameterDto, bool)`

GetCreateTimesOk returns a tuple with the CreateTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimes

`func (o *JobQueryDto) SetCreateTimes(v []JobConditionQueryParameterDto)`

SetCreateTimes sets CreateTimes field to given value.

### HasCreateTimes

`func (o *JobQueryDto) HasCreateTimes() bool`

HasCreateTimes returns a boolean if a field has been set.

### SetCreateTimesNil

`func (o *JobQueryDto) SetCreateTimesNil(b bool)`

 SetCreateTimesNil sets the value for CreateTimes to be an explicit nil

### UnsetCreateTimes
`func (o *JobQueryDto) UnsetCreateTimes()`

UnsetCreateTimes ensures that no value is present for CreateTimes, not even an explicit nil
### GetWithException

`func (o *JobQueryDto) GetWithException() bool`

GetWithException returns the WithException field if non-nil, zero value otherwise.

### GetWithExceptionOk

`func (o *JobQueryDto) GetWithExceptionOk() (*bool, bool)`

GetWithExceptionOk returns a tuple with the WithException field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithException

`func (o *JobQueryDto) SetWithException(v bool)`

SetWithException sets WithException field to given value.

### HasWithException

`func (o *JobQueryDto) HasWithException() bool`

HasWithException returns a boolean if a field has been set.

### SetWithExceptionNil

`func (o *JobQueryDto) SetWithExceptionNil(b bool)`

 SetWithExceptionNil sets the value for WithException to be an explicit nil

### UnsetWithException
`func (o *JobQueryDto) UnsetWithException()`

UnsetWithException ensures that no value is present for WithException, not even an explicit nil
### GetExceptionMessage

`func (o *JobQueryDto) GetExceptionMessage() string`

GetExceptionMessage returns the ExceptionMessage field if non-nil, zero value otherwise.

### GetExceptionMessageOk

`func (o *JobQueryDto) GetExceptionMessageOk() (*string, bool)`

GetExceptionMessageOk returns a tuple with the ExceptionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionMessage

`func (o *JobQueryDto) SetExceptionMessage(v string)`

SetExceptionMessage sets ExceptionMessage field to given value.

### HasExceptionMessage

`func (o *JobQueryDto) HasExceptionMessage() bool`

HasExceptionMessage returns a boolean if a field has been set.

### SetExceptionMessageNil

`func (o *JobQueryDto) SetExceptionMessageNil(b bool)`

 SetExceptionMessageNil sets the value for ExceptionMessage to be an explicit nil

### UnsetExceptionMessage
`func (o *JobQueryDto) UnsetExceptionMessage()`

UnsetExceptionMessage ensures that no value is present for ExceptionMessage, not even an explicit nil
### GetFailedActivityId

`func (o *JobQueryDto) GetFailedActivityId() string`

GetFailedActivityId returns the FailedActivityId field if non-nil, zero value otherwise.

### GetFailedActivityIdOk

`func (o *JobQueryDto) GetFailedActivityIdOk() (*string, bool)`

GetFailedActivityIdOk returns a tuple with the FailedActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedActivityId

`func (o *JobQueryDto) SetFailedActivityId(v string)`

SetFailedActivityId sets FailedActivityId field to given value.

### HasFailedActivityId

`func (o *JobQueryDto) HasFailedActivityId() bool`

HasFailedActivityId returns a boolean if a field has been set.

### SetFailedActivityIdNil

`func (o *JobQueryDto) SetFailedActivityIdNil(b bool)`

 SetFailedActivityIdNil sets the value for FailedActivityId to be an explicit nil

### UnsetFailedActivityId
`func (o *JobQueryDto) UnsetFailedActivityId()`

UnsetFailedActivityId ensures that no value is present for FailedActivityId, not even an explicit nil
### GetNoRetriesLeft

`func (o *JobQueryDto) GetNoRetriesLeft() bool`

GetNoRetriesLeft returns the NoRetriesLeft field if non-nil, zero value otherwise.

### GetNoRetriesLeftOk

`func (o *JobQueryDto) GetNoRetriesLeftOk() (*bool, bool)`

GetNoRetriesLeftOk returns a tuple with the NoRetriesLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRetriesLeft

`func (o *JobQueryDto) SetNoRetriesLeft(v bool)`

SetNoRetriesLeft sets NoRetriesLeft field to given value.

### HasNoRetriesLeft

`func (o *JobQueryDto) HasNoRetriesLeft() bool`

HasNoRetriesLeft returns a boolean if a field has been set.

### SetNoRetriesLeftNil

`func (o *JobQueryDto) SetNoRetriesLeftNil(b bool)`

 SetNoRetriesLeftNil sets the value for NoRetriesLeft to be an explicit nil

### UnsetNoRetriesLeft
`func (o *JobQueryDto) UnsetNoRetriesLeft()`

UnsetNoRetriesLeft ensures that no value is present for NoRetriesLeft, not even an explicit nil
### GetActive

`func (o *JobQueryDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *JobQueryDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *JobQueryDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *JobQueryDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *JobQueryDto) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *JobQueryDto) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetSuspended

`func (o *JobQueryDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *JobQueryDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *JobQueryDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *JobQueryDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *JobQueryDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *JobQueryDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetPriorityLowerThanOrEquals

`func (o *JobQueryDto) GetPriorityLowerThanOrEquals() int64`

GetPriorityLowerThanOrEquals returns the PriorityLowerThanOrEquals field if non-nil, zero value otherwise.

### GetPriorityLowerThanOrEqualsOk

`func (o *JobQueryDto) GetPriorityLowerThanOrEqualsOk() (*int64, bool)`

GetPriorityLowerThanOrEqualsOk returns a tuple with the PriorityLowerThanOrEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLowerThanOrEquals

`func (o *JobQueryDto) SetPriorityLowerThanOrEquals(v int64)`

SetPriorityLowerThanOrEquals sets PriorityLowerThanOrEquals field to given value.

### HasPriorityLowerThanOrEquals

`func (o *JobQueryDto) HasPriorityLowerThanOrEquals() bool`

HasPriorityLowerThanOrEquals returns a boolean if a field has been set.

### SetPriorityLowerThanOrEqualsNil

`func (o *JobQueryDto) SetPriorityLowerThanOrEqualsNil(b bool)`

 SetPriorityLowerThanOrEqualsNil sets the value for PriorityLowerThanOrEquals to be an explicit nil

### UnsetPriorityLowerThanOrEquals
`func (o *JobQueryDto) UnsetPriorityLowerThanOrEquals()`

UnsetPriorityLowerThanOrEquals ensures that no value is present for PriorityLowerThanOrEquals, not even an explicit nil
### GetPriorityHigherThanOrEquals

`func (o *JobQueryDto) GetPriorityHigherThanOrEquals() int64`

GetPriorityHigherThanOrEquals returns the PriorityHigherThanOrEquals field if non-nil, zero value otherwise.

### GetPriorityHigherThanOrEqualsOk

`func (o *JobQueryDto) GetPriorityHigherThanOrEqualsOk() (*int64, bool)`

GetPriorityHigherThanOrEqualsOk returns a tuple with the PriorityHigherThanOrEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityHigherThanOrEquals

`func (o *JobQueryDto) SetPriorityHigherThanOrEquals(v int64)`

SetPriorityHigherThanOrEquals sets PriorityHigherThanOrEquals field to given value.

### HasPriorityHigherThanOrEquals

`func (o *JobQueryDto) HasPriorityHigherThanOrEquals() bool`

HasPriorityHigherThanOrEquals returns a boolean if a field has been set.

### SetPriorityHigherThanOrEqualsNil

`func (o *JobQueryDto) SetPriorityHigherThanOrEqualsNil(b bool)`

 SetPriorityHigherThanOrEqualsNil sets the value for PriorityHigherThanOrEquals to be an explicit nil

### UnsetPriorityHigherThanOrEquals
`func (o *JobQueryDto) UnsetPriorityHigherThanOrEquals()`

UnsetPriorityHigherThanOrEquals ensures that no value is present for PriorityHigherThanOrEquals, not even an explicit nil
### GetTenantIdIn

`func (o *JobQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *JobQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *JobQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *JobQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *JobQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *JobQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *JobQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *JobQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *JobQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *JobQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *JobQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *JobQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetIncludeJobsWithoutTenantId

`func (o *JobQueryDto) GetIncludeJobsWithoutTenantId() bool`

GetIncludeJobsWithoutTenantId returns the IncludeJobsWithoutTenantId field if non-nil, zero value otherwise.

### GetIncludeJobsWithoutTenantIdOk

`func (o *JobQueryDto) GetIncludeJobsWithoutTenantIdOk() (*bool, bool)`

GetIncludeJobsWithoutTenantIdOk returns a tuple with the IncludeJobsWithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeJobsWithoutTenantId

`func (o *JobQueryDto) SetIncludeJobsWithoutTenantId(v bool)`

SetIncludeJobsWithoutTenantId sets IncludeJobsWithoutTenantId field to given value.

### HasIncludeJobsWithoutTenantId

`func (o *JobQueryDto) HasIncludeJobsWithoutTenantId() bool`

HasIncludeJobsWithoutTenantId returns a boolean if a field has been set.

### SetIncludeJobsWithoutTenantIdNil

`func (o *JobQueryDto) SetIncludeJobsWithoutTenantIdNil(b bool)`

 SetIncludeJobsWithoutTenantIdNil sets the value for IncludeJobsWithoutTenantId to be an explicit nil

### UnsetIncludeJobsWithoutTenantId
`func (o *JobQueryDto) UnsetIncludeJobsWithoutTenantId()`

UnsetIncludeJobsWithoutTenantId ensures that no value is present for IncludeJobsWithoutTenantId, not even an explicit nil
### GetSorting

`func (o *JobQueryDto) GetSorting() []JobQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *JobQueryDto) GetSortingOk() (*[]JobQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *JobQueryDto) SetSorting(v []JobQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *JobQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *JobQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *JobQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


