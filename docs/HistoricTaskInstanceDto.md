# HistoricTaskInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The task id. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition the task belongs to. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition the task belongs to. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance the task belongs to. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The id of the execution the task belongs to. | [optional] 
**CaseDefinitionKey** | Pointer to **NullableString** | The key of the case definition the task belongs to. | [optional] 
**CaseDefinitionId** | Pointer to **NullableString** | The id of the case definition the task belongs to. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | The id of the case instance the task belongs to. | [optional] 
**CaseExecutionId** | Pointer to **NullableString** | The id of the case execution the task belongs to. | [optional] 
**ActivityInstanceId** | Pointer to **NullableString** | The id of the activity that this object is an instance of. | [optional] 
**Name** | Pointer to **NullableString** | The task name. | [optional] 
**Description** | Pointer to **NullableString** | The task&#39;s description. | [optional] 
**DeleteReason** | Pointer to **NullableString** | The task&#39;s delete reason. | [optional] 
**Owner** | Pointer to **NullableString** | The owner&#39;s id. | [optional] 
**Assignee** | Pointer to **NullableString** | The assignee&#39;s id. | [optional] 
**StartTime** | Pointer to **NullableTime** | The time the task was started. Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**EndTime** | Pointer to **NullableTime** | The time the task ended. Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**Duration** | Pointer to **NullableInt64** | The time the task took to finish (in milliseconds). | [optional] 
**TaskDefinitionKey** | Pointer to **NullableString** | The task&#39;s key. | [optional] 
**Priority** | Pointer to **NullableInt32** | The task&#39;s priority. | [optional] 
**Due** | Pointer to **NullableTime** | The task&#39;s due date. Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**ParentTaskId** | Pointer to **NullableString** | The id of the parent task, if this task is a subtask. | [optional] 
**FollowUp** | Pointer to **NullableTime** | The follow-up date for the task. Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the task instance. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the task should be removed by the History Cleanup job. Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process containing this task. | [optional] 

## Methods

### NewHistoricTaskInstanceDto

`func NewHistoricTaskInstanceDto() *HistoricTaskInstanceDto`

NewHistoricTaskInstanceDto instantiates a new HistoricTaskInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricTaskInstanceDtoWithDefaults

`func NewHistoricTaskInstanceDtoWithDefaults() *HistoricTaskInstanceDto`

NewHistoricTaskInstanceDtoWithDefaults instantiates a new HistoricTaskInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricTaskInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricTaskInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricTaskInstanceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricTaskInstanceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricTaskInstanceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricTaskInstanceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricTaskInstanceDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricTaskInstanceDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricTaskInstanceDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricTaskInstanceDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricTaskInstanceDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricTaskInstanceDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricTaskInstanceDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricTaskInstanceDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricTaskInstanceDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricTaskInstanceDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricTaskInstanceDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricTaskInstanceDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricTaskInstanceDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricTaskInstanceDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricTaskInstanceDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricTaskInstanceDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricTaskInstanceDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricTaskInstanceDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetExecutionId

`func (o *HistoricTaskInstanceDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HistoricTaskInstanceDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HistoricTaskInstanceDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HistoricTaskInstanceDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HistoricTaskInstanceDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HistoricTaskInstanceDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetCaseDefinitionKey

`func (o *HistoricTaskInstanceDto) GetCaseDefinitionKey() string`

GetCaseDefinitionKey returns the CaseDefinitionKey field if non-nil, zero value otherwise.

### GetCaseDefinitionKeyOk

`func (o *HistoricTaskInstanceDto) GetCaseDefinitionKeyOk() (*string, bool)`

GetCaseDefinitionKeyOk returns a tuple with the CaseDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionKey

`func (o *HistoricTaskInstanceDto) SetCaseDefinitionKey(v string)`

SetCaseDefinitionKey sets CaseDefinitionKey field to given value.

### HasCaseDefinitionKey

`func (o *HistoricTaskInstanceDto) HasCaseDefinitionKey() bool`

HasCaseDefinitionKey returns a boolean if a field has been set.

### SetCaseDefinitionKeyNil

`func (o *HistoricTaskInstanceDto) SetCaseDefinitionKeyNil(b bool)`

 SetCaseDefinitionKeyNil sets the value for CaseDefinitionKey to be an explicit nil

### UnsetCaseDefinitionKey
`func (o *HistoricTaskInstanceDto) UnsetCaseDefinitionKey()`

UnsetCaseDefinitionKey ensures that no value is present for CaseDefinitionKey, not even an explicit nil
### GetCaseDefinitionId

`func (o *HistoricTaskInstanceDto) GetCaseDefinitionId() string`

GetCaseDefinitionId returns the CaseDefinitionId field if non-nil, zero value otherwise.

### GetCaseDefinitionIdOk

`func (o *HistoricTaskInstanceDto) GetCaseDefinitionIdOk() (*string, bool)`

GetCaseDefinitionIdOk returns a tuple with the CaseDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionId

`func (o *HistoricTaskInstanceDto) SetCaseDefinitionId(v string)`

SetCaseDefinitionId sets CaseDefinitionId field to given value.

### HasCaseDefinitionId

`func (o *HistoricTaskInstanceDto) HasCaseDefinitionId() bool`

HasCaseDefinitionId returns a boolean if a field has been set.

### SetCaseDefinitionIdNil

`func (o *HistoricTaskInstanceDto) SetCaseDefinitionIdNil(b bool)`

 SetCaseDefinitionIdNil sets the value for CaseDefinitionId to be an explicit nil

### UnsetCaseDefinitionId
`func (o *HistoricTaskInstanceDto) UnsetCaseDefinitionId()`

UnsetCaseDefinitionId ensures that no value is present for CaseDefinitionId, not even an explicit nil
### GetCaseInstanceId

`func (o *HistoricTaskInstanceDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *HistoricTaskInstanceDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *HistoricTaskInstanceDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *HistoricTaskInstanceDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *HistoricTaskInstanceDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *HistoricTaskInstanceDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetCaseExecutionId

`func (o *HistoricTaskInstanceDto) GetCaseExecutionId() string`

GetCaseExecutionId returns the CaseExecutionId field if non-nil, zero value otherwise.

### GetCaseExecutionIdOk

`func (o *HistoricTaskInstanceDto) GetCaseExecutionIdOk() (*string, bool)`

GetCaseExecutionIdOk returns a tuple with the CaseExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExecutionId

`func (o *HistoricTaskInstanceDto) SetCaseExecutionId(v string)`

SetCaseExecutionId sets CaseExecutionId field to given value.

### HasCaseExecutionId

`func (o *HistoricTaskInstanceDto) HasCaseExecutionId() bool`

HasCaseExecutionId returns a boolean if a field has been set.

### SetCaseExecutionIdNil

`func (o *HistoricTaskInstanceDto) SetCaseExecutionIdNil(b bool)`

 SetCaseExecutionIdNil sets the value for CaseExecutionId to be an explicit nil

### UnsetCaseExecutionId
`func (o *HistoricTaskInstanceDto) UnsetCaseExecutionId()`

UnsetCaseExecutionId ensures that no value is present for CaseExecutionId, not even an explicit nil
### GetActivityInstanceId

`func (o *HistoricTaskInstanceDto) GetActivityInstanceId() string`

GetActivityInstanceId returns the ActivityInstanceId field if non-nil, zero value otherwise.

### GetActivityInstanceIdOk

`func (o *HistoricTaskInstanceDto) GetActivityInstanceIdOk() (*string, bool)`

GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceId

`func (o *HistoricTaskInstanceDto) SetActivityInstanceId(v string)`

SetActivityInstanceId sets ActivityInstanceId field to given value.

### HasActivityInstanceId

`func (o *HistoricTaskInstanceDto) HasActivityInstanceId() bool`

HasActivityInstanceId returns a boolean if a field has been set.

### SetActivityInstanceIdNil

`func (o *HistoricTaskInstanceDto) SetActivityInstanceIdNil(b bool)`

 SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil

### UnsetActivityInstanceId
`func (o *HistoricTaskInstanceDto) UnsetActivityInstanceId()`

UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
### GetName

`func (o *HistoricTaskInstanceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HistoricTaskInstanceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HistoricTaskInstanceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HistoricTaskInstanceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HistoricTaskInstanceDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HistoricTaskInstanceDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *HistoricTaskInstanceDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HistoricTaskInstanceDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HistoricTaskInstanceDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HistoricTaskInstanceDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HistoricTaskInstanceDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HistoricTaskInstanceDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDeleteReason

`func (o *HistoricTaskInstanceDto) GetDeleteReason() string`

GetDeleteReason returns the DeleteReason field if non-nil, zero value otherwise.

### GetDeleteReasonOk

`func (o *HistoricTaskInstanceDto) GetDeleteReasonOk() (*string, bool)`

GetDeleteReasonOk returns a tuple with the DeleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteReason

`func (o *HistoricTaskInstanceDto) SetDeleteReason(v string)`

SetDeleteReason sets DeleteReason field to given value.

### HasDeleteReason

`func (o *HistoricTaskInstanceDto) HasDeleteReason() bool`

HasDeleteReason returns a boolean if a field has been set.

### SetDeleteReasonNil

`func (o *HistoricTaskInstanceDto) SetDeleteReasonNil(b bool)`

 SetDeleteReasonNil sets the value for DeleteReason to be an explicit nil

### UnsetDeleteReason
`func (o *HistoricTaskInstanceDto) UnsetDeleteReason()`

UnsetDeleteReason ensures that no value is present for DeleteReason, not even an explicit nil
### GetOwner

`func (o *HistoricTaskInstanceDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *HistoricTaskInstanceDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *HistoricTaskInstanceDto) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *HistoricTaskInstanceDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *HistoricTaskInstanceDto) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *HistoricTaskInstanceDto) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetAssignee

`func (o *HistoricTaskInstanceDto) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *HistoricTaskInstanceDto) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *HistoricTaskInstanceDto) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *HistoricTaskInstanceDto) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *HistoricTaskInstanceDto) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *HistoricTaskInstanceDto) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetStartTime

`func (o *HistoricTaskInstanceDto) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HistoricTaskInstanceDto) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HistoricTaskInstanceDto) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HistoricTaskInstanceDto) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *HistoricTaskInstanceDto) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *HistoricTaskInstanceDto) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *HistoricTaskInstanceDto) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HistoricTaskInstanceDto) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HistoricTaskInstanceDto) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *HistoricTaskInstanceDto) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *HistoricTaskInstanceDto) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *HistoricTaskInstanceDto) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetDuration

`func (o *HistoricTaskInstanceDto) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *HistoricTaskInstanceDto) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *HistoricTaskInstanceDto) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *HistoricTaskInstanceDto) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *HistoricTaskInstanceDto) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *HistoricTaskInstanceDto) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetTaskDefinitionKey

`func (o *HistoricTaskInstanceDto) GetTaskDefinitionKey() string`

GetTaskDefinitionKey returns the TaskDefinitionKey field if non-nil, zero value otherwise.

### GetTaskDefinitionKeyOk

`func (o *HistoricTaskInstanceDto) GetTaskDefinitionKeyOk() (*string, bool)`

GetTaskDefinitionKeyOk returns a tuple with the TaskDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionKey

`func (o *HistoricTaskInstanceDto) SetTaskDefinitionKey(v string)`

SetTaskDefinitionKey sets TaskDefinitionKey field to given value.

### HasTaskDefinitionKey

`func (o *HistoricTaskInstanceDto) HasTaskDefinitionKey() bool`

HasTaskDefinitionKey returns a boolean if a field has been set.

### SetTaskDefinitionKeyNil

`func (o *HistoricTaskInstanceDto) SetTaskDefinitionKeyNil(b bool)`

 SetTaskDefinitionKeyNil sets the value for TaskDefinitionKey to be an explicit nil

### UnsetTaskDefinitionKey
`func (o *HistoricTaskInstanceDto) UnsetTaskDefinitionKey()`

UnsetTaskDefinitionKey ensures that no value is present for TaskDefinitionKey, not even an explicit nil
### GetPriority

`func (o *HistoricTaskInstanceDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *HistoricTaskInstanceDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *HistoricTaskInstanceDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *HistoricTaskInstanceDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *HistoricTaskInstanceDto) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *HistoricTaskInstanceDto) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetDue

`func (o *HistoricTaskInstanceDto) GetDue() time.Time`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *HistoricTaskInstanceDto) GetDueOk() (*time.Time, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *HistoricTaskInstanceDto) SetDue(v time.Time)`

SetDue sets Due field to given value.

### HasDue

`func (o *HistoricTaskInstanceDto) HasDue() bool`

HasDue returns a boolean if a field has been set.

### SetDueNil

`func (o *HistoricTaskInstanceDto) SetDueNil(b bool)`

 SetDueNil sets the value for Due to be an explicit nil

### UnsetDue
`func (o *HistoricTaskInstanceDto) UnsetDue()`

UnsetDue ensures that no value is present for Due, not even an explicit nil
### GetParentTaskId

`func (o *HistoricTaskInstanceDto) GetParentTaskId() string`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *HistoricTaskInstanceDto) GetParentTaskIdOk() (*string, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *HistoricTaskInstanceDto) SetParentTaskId(v string)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *HistoricTaskInstanceDto) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### SetParentTaskIdNil

`func (o *HistoricTaskInstanceDto) SetParentTaskIdNil(b bool)`

 SetParentTaskIdNil sets the value for ParentTaskId to be an explicit nil

### UnsetParentTaskId
`func (o *HistoricTaskInstanceDto) UnsetParentTaskId()`

UnsetParentTaskId ensures that no value is present for ParentTaskId, not even an explicit nil
### GetFollowUp

`func (o *HistoricTaskInstanceDto) GetFollowUp() time.Time`

GetFollowUp returns the FollowUp field if non-nil, zero value otherwise.

### GetFollowUpOk

`func (o *HistoricTaskInstanceDto) GetFollowUpOk() (*time.Time, bool)`

GetFollowUpOk returns a tuple with the FollowUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUp

`func (o *HistoricTaskInstanceDto) SetFollowUp(v time.Time)`

SetFollowUp sets FollowUp field to given value.

### HasFollowUp

`func (o *HistoricTaskInstanceDto) HasFollowUp() bool`

HasFollowUp returns a boolean if a field has been set.

### SetFollowUpNil

`func (o *HistoricTaskInstanceDto) SetFollowUpNil(b bool)`

 SetFollowUpNil sets the value for FollowUp to be an explicit nil

### UnsetFollowUp
`func (o *HistoricTaskInstanceDto) UnsetFollowUp()`

UnsetFollowUp ensures that no value is present for FollowUp, not even an explicit nil
### GetTenantId

`func (o *HistoricTaskInstanceDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricTaskInstanceDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricTaskInstanceDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricTaskInstanceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricTaskInstanceDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricTaskInstanceDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetRemovalTime

`func (o *HistoricTaskInstanceDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricTaskInstanceDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricTaskInstanceDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricTaskInstanceDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricTaskInstanceDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricTaskInstanceDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricTaskInstanceDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricTaskInstanceDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricTaskInstanceDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricTaskInstanceDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricTaskInstanceDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricTaskInstanceDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


