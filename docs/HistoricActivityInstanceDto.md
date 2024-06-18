# HistoricActivityInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the activity instance. | [optional] 
**ParentActivityInstanceId** | Pointer to **NullableString** | The id of the parent activity instance, for example a sub process instance. | [optional] 
**ActivityId** | Pointer to **NullableString** | The id of the activity that this object is an instance of. | [optional] 
**ActivityName** | Pointer to **NullableString** | The name of the activity that this object is an instance of. | [optional] 
**ActivityType** | Pointer to **NullableString** | The type of the activity that this object is an instance of. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition that this activity instance belongs to. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition that this activity instance belongs to. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance that this activity instance belongs to. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The id of the execution that executed this activity instance. | [optional] 
**TaskId** | Pointer to **NullableString** | The id of the task that is associated to this activity instance. Is only set if the activity is a user task. | [optional] 
**Assignee** | Pointer to **NullableString** | The assignee of the task that is associated to this activity instance. Is only set if the activity is a user task. | [optional] 
**CalledProcessInstanceId** | Pointer to **NullableString** | The id of the called process instance. Is only set if the activity is a call activity and the called instance a process instance. | [optional] 
**CalledCaseInstanceId** | Pointer to **NullableString** | The id of the called case instance. Is only set if the activity is a call activity and the called instance a case instance. | [optional] 
**StartTime** | Pointer to **NullableTime** | The time the instance was started. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**EndTime** | Pointer to **NullableTime** | The time the instance ended. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**DurationInMillis** | Pointer to **NullableInt64** | The time the instance took to finish (in milliseconds). | [optional] 
**Canceled** | Pointer to **NullableBool** | If &#x60;true&#x60;, this activity instance is canceled. | [optional] 
**CompleteScope** | Pointer to **NullableBool** | If &#x60;true&#x60;, this activity instance did complete a BPMN 2.0 scope. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the activity instance. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the activity instance should be removed by the History Cleanup job. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process containing this activity instance. | [optional] 

## Methods

### NewHistoricActivityInstanceDto

`func NewHistoricActivityInstanceDto() *HistoricActivityInstanceDto`

NewHistoricActivityInstanceDto instantiates a new HistoricActivityInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricActivityInstanceDtoWithDefaults

`func NewHistoricActivityInstanceDtoWithDefaults() *HistoricActivityInstanceDto`

NewHistoricActivityInstanceDtoWithDefaults instantiates a new HistoricActivityInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricActivityInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricActivityInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricActivityInstanceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricActivityInstanceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricActivityInstanceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricActivityInstanceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetParentActivityInstanceId

`func (o *HistoricActivityInstanceDto) GetParentActivityInstanceId() string`

GetParentActivityInstanceId returns the ParentActivityInstanceId field if non-nil, zero value otherwise.

### GetParentActivityInstanceIdOk

`func (o *HistoricActivityInstanceDto) GetParentActivityInstanceIdOk() (*string, bool)`

GetParentActivityInstanceIdOk returns a tuple with the ParentActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentActivityInstanceId

`func (o *HistoricActivityInstanceDto) SetParentActivityInstanceId(v string)`

SetParentActivityInstanceId sets ParentActivityInstanceId field to given value.

### HasParentActivityInstanceId

`func (o *HistoricActivityInstanceDto) HasParentActivityInstanceId() bool`

HasParentActivityInstanceId returns a boolean if a field has been set.

### SetParentActivityInstanceIdNil

`func (o *HistoricActivityInstanceDto) SetParentActivityInstanceIdNil(b bool)`

 SetParentActivityInstanceIdNil sets the value for ParentActivityInstanceId to be an explicit nil

### UnsetParentActivityInstanceId
`func (o *HistoricActivityInstanceDto) UnsetParentActivityInstanceId()`

UnsetParentActivityInstanceId ensures that no value is present for ParentActivityInstanceId, not even an explicit nil
### GetActivityId

`func (o *HistoricActivityInstanceDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *HistoricActivityInstanceDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *HistoricActivityInstanceDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *HistoricActivityInstanceDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *HistoricActivityInstanceDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *HistoricActivityInstanceDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetActivityName

`func (o *HistoricActivityInstanceDto) GetActivityName() string`

GetActivityName returns the ActivityName field if non-nil, zero value otherwise.

### GetActivityNameOk

`func (o *HistoricActivityInstanceDto) GetActivityNameOk() (*string, bool)`

GetActivityNameOk returns a tuple with the ActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityName

`func (o *HistoricActivityInstanceDto) SetActivityName(v string)`

SetActivityName sets ActivityName field to given value.

### HasActivityName

`func (o *HistoricActivityInstanceDto) HasActivityName() bool`

HasActivityName returns a boolean if a field has been set.

### SetActivityNameNil

`func (o *HistoricActivityInstanceDto) SetActivityNameNil(b bool)`

 SetActivityNameNil sets the value for ActivityName to be an explicit nil

### UnsetActivityName
`func (o *HistoricActivityInstanceDto) UnsetActivityName()`

UnsetActivityName ensures that no value is present for ActivityName, not even an explicit nil
### GetActivityType

`func (o *HistoricActivityInstanceDto) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *HistoricActivityInstanceDto) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *HistoricActivityInstanceDto) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *HistoricActivityInstanceDto) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### SetActivityTypeNil

`func (o *HistoricActivityInstanceDto) SetActivityTypeNil(b bool)`

 SetActivityTypeNil sets the value for ActivityType to be an explicit nil

### UnsetActivityType
`func (o *HistoricActivityInstanceDto) UnsetActivityType()`

UnsetActivityType ensures that no value is present for ActivityType, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricActivityInstanceDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricActivityInstanceDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricActivityInstanceDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricActivityInstanceDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricActivityInstanceDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricActivityInstanceDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricActivityInstanceDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricActivityInstanceDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricActivityInstanceDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricActivityInstanceDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricActivityInstanceDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricActivityInstanceDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricActivityInstanceDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricActivityInstanceDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricActivityInstanceDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricActivityInstanceDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricActivityInstanceDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricActivityInstanceDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetExecutionId

`func (o *HistoricActivityInstanceDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HistoricActivityInstanceDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HistoricActivityInstanceDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HistoricActivityInstanceDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HistoricActivityInstanceDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HistoricActivityInstanceDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetTaskId

`func (o *HistoricActivityInstanceDto) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *HistoricActivityInstanceDto) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *HistoricActivityInstanceDto) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *HistoricActivityInstanceDto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *HistoricActivityInstanceDto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *HistoricActivityInstanceDto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetAssignee

`func (o *HistoricActivityInstanceDto) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *HistoricActivityInstanceDto) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *HistoricActivityInstanceDto) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *HistoricActivityInstanceDto) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *HistoricActivityInstanceDto) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *HistoricActivityInstanceDto) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetCalledProcessInstanceId

`func (o *HistoricActivityInstanceDto) GetCalledProcessInstanceId() string`

GetCalledProcessInstanceId returns the CalledProcessInstanceId field if non-nil, zero value otherwise.

### GetCalledProcessInstanceIdOk

`func (o *HistoricActivityInstanceDto) GetCalledProcessInstanceIdOk() (*string, bool)`

GetCalledProcessInstanceIdOk returns a tuple with the CalledProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalledProcessInstanceId

`func (o *HistoricActivityInstanceDto) SetCalledProcessInstanceId(v string)`

SetCalledProcessInstanceId sets CalledProcessInstanceId field to given value.

### HasCalledProcessInstanceId

`func (o *HistoricActivityInstanceDto) HasCalledProcessInstanceId() bool`

HasCalledProcessInstanceId returns a boolean if a field has been set.

### SetCalledProcessInstanceIdNil

`func (o *HistoricActivityInstanceDto) SetCalledProcessInstanceIdNil(b bool)`

 SetCalledProcessInstanceIdNil sets the value for CalledProcessInstanceId to be an explicit nil

### UnsetCalledProcessInstanceId
`func (o *HistoricActivityInstanceDto) UnsetCalledProcessInstanceId()`

UnsetCalledProcessInstanceId ensures that no value is present for CalledProcessInstanceId, not even an explicit nil
### GetCalledCaseInstanceId

`func (o *HistoricActivityInstanceDto) GetCalledCaseInstanceId() string`

GetCalledCaseInstanceId returns the CalledCaseInstanceId field if non-nil, zero value otherwise.

### GetCalledCaseInstanceIdOk

`func (o *HistoricActivityInstanceDto) GetCalledCaseInstanceIdOk() (*string, bool)`

GetCalledCaseInstanceIdOk returns a tuple with the CalledCaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalledCaseInstanceId

`func (o *HistoricActivityInstanceDto) SetCalledCaseInstanceId(v string)`

SetCalledCaseInstanceId sets CalledCaseInstanceId field to given value.

### HasCalledCaseInstanceId

`func (o *HistoricActivityInstanceDto) HasCalledCaseInstanceId() bool`

HasCalledCaseInstanceId returns a boolean if a field has been set.

### SetCalledCaseInstanceIdNil

`func (o *HistoricActivityInstanceDto) SetCalledCaseInstanceIdNil(b bool)`

 SetCalledCaseInstanceIdNil sets the value for CalledCaseInstanceId to be an explicit nil

### UnsetCalledCaseInstanceId
`func (o *HistoricActivityInstanceDto) UnsetCalledCaseInstanceId()`

UnsetCalledCaseInstanceId ensures that no value is present for CalledCaseInstanceId, not even an explicit nil
### GetStartTime

`func (o *HistoricActivityInstanceDto) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HistoricActivityInstanceDto) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HistoricActivityInstanceDto) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HistoricActivityInstanceDto) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *HistoricActivityInstanceDto) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *HistoricActivityInstanceDto) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *HistoricActivityInstanceDto) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HistoricActivityInstanceDto) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HistoricActivityInstanceDto) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *HistoricActivityInstanceDto) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *HistoricActivityInstanceDto) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *HistoricActivityInstanceDto) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetDurationInMillis

`func (o *HistoricActivityInstanceDto) GetDurationInMillis() int64`

GetDurationInMillis returns the DurationInMillis field if non-nil, zero value otherwise.

### GetDurationInMillisOk

`func (o *HistoricActivityInstanceDto) GetDurationInMillisOk() (*int64, bool)`

GetDurationInMillisOk returns a tuple with the DurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMillis

`func (o *HistoricActivityInstanceDto) SetDurationInMillis(v int64)`

SetDurationInMillis sets DurationInMillis field to given value.

### HasDurationInMillis

`func (o *HistoricActivityInstanceDto) HasDurationInMillis() bool`

HasDurationInMillis returns a boolean if a field has been set.

### SetDurationInMillisNil

`func (o *HistoricActivityInstanceDto) SetDurationInMillisNil(b bool)`

 SetDurationInMillisNil sets the value for DurationInMillis to be an explicit nil

### UnsetDurationInMillis
`func (o *HistoricActivityInstanceDto) UnsetDurationInMillis()`

UnsetDurationInMillis ensures that no value is present for DurationInMillis, not even an explicit nil
### GetCanceled

`func (o *HistoricActivityInstanceDto) GetCanceled() bool`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *HistoricActivityInstanceDto) GetCanceledOk() (*bool, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *HistoricActivityInstanceDto) SetCanceled(v bool)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *HistoricActivityInstanceDto) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### SetCanceledNil

`func (o *HistoricActivityInstanceDto) SetCanceledNil(b bool)`

 SetCanceledNil sets the value for Canceled to be an explicit nil

### UnsetCanceled
`func (o *HistoricActivityInstanceDto) UnsetCanceled()`

UnsetCanceled ensures that no value is present for Canceled, not even an explicit nil
### GetCompleteScope

`func (o *HistoricActivityInstanceDto) GetCompleteScope() bool`

GetCompleteScope returns the CompleteScope field if non-nil, zero value otherwise.

### GetCompleteScopeOk

`func (o *HistoricActivityInstanceDto) GetCompleteScopeOk() (*bool, bool)`

GetCompleteScopeOk returns a tuple with the CompleteScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteScope

`func (o *HistoricActivityInstanceDto) SetCompleteScope(v bool)`

SetCompleteScope sets CompleteScope field to given value.

### HasCompleteScope

`func (o *HistoricActivityInstanceDto) HasCompleteScope() bool`

HasCompleteScope returns a boolean if a field has been set.

### SetCompleteScopeNil

`func (o *HistoricActivityInstanceDto) SetCompleteScopeNil(b bool)`

 SetCompleteScopeNil sets the value for CompleteScope to be an explicit nil

### UnsetCompleteScope
`func (o *HistoricActivityInstanceDto) UnsetCompleteScope()`

UnsetCompleteScope ensures that no value is present for CompleteScope, not even an explicit nil
### GetTenantId

`func (o *HistoricActivityInstanceDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricActivityInstanceDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricActivityInstanceDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricActivityInstanceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricActivityInstanceDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricActivityInstanceDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetRemovalTime

`func (o *HistoricActivityInstanceDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricActivityInstanceDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricActivityInstanceDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricActivityInstanceDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricActivityInstanceDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricActivityInstanceDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricActivityInstanceDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricActivityInstanceDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricActivityInstanceDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricActivityInstanceDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricActivityInstanceDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricActivityInstanceDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


