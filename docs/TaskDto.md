# TaskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The task id. | [optional] 
**Name** | Pointer to **NullableString** | The task name. | [optional] 
**Assignee** | Pointer to **NullableString** | The assignee&#39;s id. | [optional] 
**Owner** | Pointer to **NullableString** | The owner&#39;s id. | [optional] 
**Created** | Pointer to **NullableTime** | The date the task was created on. [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**LastUpdated** | Pointer to **NullableTime** | The date the task was last updated. Every action that fires a [task update event](https://docs.camunda.org/manual/7.21/user-guide/process-engine/delegation-code/#task-listener-event-lifecycle) will update this property. [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**Due** | Pointer to **NullableTime** | The task&#39;s due date. [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**FollowUp** | Pointer to **NullableTime** | The follow-up date for the task. [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**DelegationState** | Pointer to **NullableString** | The task&#39;s delegation state. Possible values are &#x60;PENDING&#x60; and &#x60;RESOLVED&#x60;. | [optional] 
**Description** | Pointer to **NullableString** | The task&#39;s description. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The id of the execution the task belongs to. | [optional] 
**ParentTaskId** | Pointer to **NullableString** | The id the parent task, if this task is a subtask. | [optional] 
**Priority** | Pointer to **NullableInt32** | The task&#39;s priority. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition the task belongs to. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance the task belongs to. | [optional] 
**CaseExecutionId** | Pointer to **NullableString** | The id of the case execution the task belongs to. | [optional] 
**CaseDefinitionId** | Pointer to **NullableString** | The id of the case definition the task belongs to. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | The id of the case instance the task belongs to. | [optional] 
**TaskDefinitionKey** | Pointer to **NullableString** | The task&#39;s key. | [optional] 
**Suspended** | Pointer to **NullableBool** | Whether the task belongs to a process instance that is suspended. | [optional] 
**FormKey** | Pointer to **NullableString** | If not &#x60;null&#x60;, the form key for the task. | [optional] 
**CamundaFormRef** | Pointer to [**CamundaFormRef**](CamundaFormRef.md) |  | [optional] 
**TenantId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the tenant id of the task. | [optional] 

## Methods

### NewTaskDto

`func NewTaskDto() *TaskDto`

NewTaskDto instantiates a new TaskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskDtoWithDefaults

`func NewTaskDtoWithDefaults() *TaskDto`

NewTaskDtoWithDefaults instantiates a new TaskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaskDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaskDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TaskDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TaskDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TaskDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAssignee

`func (o *TaskDto) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *TaskDto) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *TaskDto) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *TaskDto) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *TaskDto) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *TaskDto) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetOwner

`func (o *TaskDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TaskDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TaskDto) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TaskDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *TaskDto) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *TaskDto) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetCreated

`func (o *TaskDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TaskDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TaskDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TaskDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *TaskDto) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *TaskDto) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *TaskDto) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TaskDto) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TaskDto) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *TaskDto) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *TaskDto) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *TaskDto) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetDue

`func (o *TaskDto) GetDue() time.Time`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *TaskDto) GetDueOk() (*time.Time, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *TaskDto) SetDue(v time.Time)`

SetDue sets Due field to given value.

### HasDue

`func (o *TaskDto) HasDue() bool`

HasDue returns a boolean if a field has been set.

### SetDueNil

`func (o *TaskDto) SetDueNil(b bool)`

 SetDueNil sets the value for Due to be an explicit nil

### UnsetDue
`func (o *TaskDto) UnsetDue()`

UnsetDue ensures that no value is present for Due, not even an explicit nil
### GetFollowUp

`func (o *TaskDto) GetFollowUp() time.Time`

GetFollowUp returns the FollowUp field if non-nil, zero value otherwise.

### GetFollowUpOk

`func (o *TaskDto) GetFollowUpOk() (*time.Time, bool)`

GetFollowUpOk returns a tuple with the FollowUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUp

`func (o *TaskDto) SetFollowUp(v time.Time)`

SetFollowUp sets FollowUp field to given value.

### HasFollowUp

`func (o *TaskDto) HasFollowUp() bool`

HasFollowUp returns a boolean if a field has been set.

### SetFollowUpNil

`func (o *TaskDto) SetFollowUpNil(b bool)`

 SetFollowUpNil sets the value for FollowUp to be an explicit nil

### UnsetFollowUp
`func (o *TaskDto) UnsetFollowUp()`

UnsetFollowUp ensures that no value is present for FollowUp, not even an explicit nil
### GetDelegationState

`func (o *TaskDto) GetDelegationState() string`

GetDelegationState returns the DelegationState field if non-nil, zero value otherwise.

### GetDelegationStateOk

`func (o *TaskDto) GetDelegationStateOk() (*string, bool)`

GetDelegationStateOk returns a tuple with the DelegationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationState

`func (o *TaskDto) SetDelegationState(v string)`

SetDelegationState sets DelegationState field to given value.

### HasDelegationState

`func (o *TaskDto) HasDelegationState() bool`

HasDelegationState returns a boolean if a field has been set.

### SetDelegationStateNil

`func (o *TaskDto) SetDelegationStateNil(b bool)`

 SetDelegationStateNil sets the value for DelegationState to be an explicit nil

### UnsetDelegationState
`func (o *TaskDto) UnsetDelegationState()`

UnsetDelegationState ensures that no value is present for DelegationState, not even an explicit nil
### GetDescription

`func (o *TaskDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TaskDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TaskDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExecutionId

`func (o *TaskDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *TaskDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *TaskDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *TaskDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *TaskDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *TaskDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetParentTaskId

`func (o *TaskDto) GetParentTaskId() string`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *TaskDto) GetParentTaskIdOk() (*string, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *TaskDto) SetParentTaskId(v string)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *TaskDto) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### SetParentTaskIdNil

`func (o *TaskDto) SetParentTaskIdNil(b bool)`

 SetParentTaskIdNil sets the value for ParentTaskId to be an explicit nil

### UnsetParentTaskId
`func (o *TaskDto) UnsetParentTaskId()`

UnsetParentTaskId ensures that no value is present for ParentTaskId, not even an explicit nil
### GetPriority

`func (o *TaskDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaskDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *TaskDto) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *TaskDto) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetProcessDefinitionId

`func (o *TaskDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *TaskDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *TaskDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *TaskDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *TaskDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *TaskDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessInstanceId

`func (o *TaskDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *TaskDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *TaskDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *TaskDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *TaskDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *TaskDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetCaseExecutionId

`func (o *TaskDto) GetCaseExecutionId() string`

GetCaseExecutionId returns the CaseExecutionId field if non-nil, zero value otherwise.

### GetCaseExecutionIdOk

`func (o *TaskDto) GetCaseExecutionIdOk() (*string, bool)`

GetCaseExecutionIdOk returns a tuple with the CaseExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExecutionId

`func (o *TaskDto) SetCaseExecutionId(v string)`

SetCaseExecutionId sets CaseExecutionId field to given value.

### HasCaseExecutionId

`func (o *TaskDto) HasCaseExecutionId() bool`

HasCaseExecutionId returns a boolean if a field has been set.

### SetCaseExecutionIdNil

`func (o *TaskDto) SetCaseExecutionIdNil(b bool)`

 SetCaseExecutionIdNil sets the value for CaseExecutionId to be an explicit nil

### UnsetCaseExecutionId
`func (o *TaskDto) UnsetCaseExecutionId()`

UnsetCaseExecutionId ensures that no value is present for CaseExecutionId, not even an explicit nil
### GetCaseDefinitionId

`func (o *TaskDto) GetCaseDefinitionId() string`

GetCaseDefinitionId returns the CaseDefinitionId field if non-nil, zero value otherwise.

### GetCaseDefinitionIdOk

`func (o *TaskDto) GetCaseDefinitionIdOk() (*string, bool)`

GetCaseDefinitionIdOk returns a tuple with the CaseDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionId

`func (o *TaskDto) SetCaseDefinitionId(v string)`

SetCaseDefinitionId sets CaseDefinitionId field to given value.

### HasCaseDefinitionId

`func (o *TaskDto) HasCaseDefinitionId() bool`

HasCaseDefinitionId returns a boolean if a field has been set.

### SetCaseDefinitionIdNil

`func (o *TaskDto) SetCaseDefinitionIdNil(b bool)`

 SetCaseDefinitionIdNil sets the value for CaseDefinitionId to be an explicit nil

### UnsetCaseDefinitionId
`func (o *TaskDto) UnsetCaseDefinitionId()`

UnsetCaseDefinitionId ensures that no value is present for CaseDefinitionId, not even an explicit nil
### GetCaseInstanceId

`func (o *TaskDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *TaskDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *TaskDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *TaskDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *TaskDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *TaskDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetTaskDefinitionKey

`func (o *TaskDto) GetTaskDefinitionKey() string`

GetTaskDefinitionKey returns the TaskDefinitionKey field if non-nil, zero value otherwise.

### GetTaskDefinitionKeyOk

`func (o *TaskDto) GetTaskDefinitionKeyOk() (*string, bool)`

GetTaskDefinitionKeyOk returns a tuple with the TaskDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionKey

`func (o *TaskDto) SetTaskDefinitionKey(v string)`

SetTaskDefinitionKey sets TaskDefinitionKey field to given value.

### HasTaskDefinitionKey

`func (o *TaskDto) HasTaskDefinitionKey() bool`

HasTaskDefinitionKey returns a boolean if a field has been set.

### SetTaskDefinitionKeyNil

`func (o *TaskDto) SetTaskDefinitionKeyNil(b bool)`

 SetTaskDefinitionKeyNil sets the value for TaskDefinitionKey to be an explicit nil

### UnsetTaskDefinitionKey
`func (o *TaskDto) UnsetTaskDefinitionKey()`

UnsetTaskDefinitionKey ensures that no value is present for TaskDefinitionKey, not even an explicit nil
### GetSuspended

`func (o *TaskDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *TaskDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *TaskDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *TaskDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *TaskDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *TaskDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetFormKey

`func (o *TaskDto) GetFormKey() string`

GetFormKey returns the FormKey field if non-nil, zero value otherwise.

### GetFormKeyOk

`func (o *TaskDto) GetFormKeyOk() (*string, bool)`

GetFormKeyOk returns a tuple with the FormKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormKey

`func (o *TaskDto) SetFormKey(v string)`

SetFormKey sets FormKey field to given value.

### HasFormKey

`func (o *TaskDto) HasFormKey() bool`

HasFormKey returns a boolean if a field has been set.

### SetFormKeyNil

`func (o *TaskDto) SetFormKeyNil(b bool)`

 SetFormKeyNil sets the value for FormKey to be an explicit nil

### UnsetFormKey
`func (o *TaskDto) UnsetFormKey()`

UnsetFormKey ensures that no value is present for FormKey, not even an explicit nil
### GetCamundaFormRef

`func (o *TaskDto) GetCamundaFormRef() CamundaFormRef`

GetCamundaFormRef returns the CamundaFormRef field if non-nil, zero value otherwise.

### GetCamundaFormRefOk

`func (o *TaskDto) GetCamundaFormRefOk() (*CamundaFormRef, bool)`

GetCamundaFormRefOk returns a tuple with the CamundaFormRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCamundaFormRef

`func (o *TaskDto) SetCamundaFormRef(v CamundaFormRef)`

SetCamundaFormRef sets CamundaFormRef field to given value.

### HasCamundaFormRef

`func (o *TaskDto) HasCamundaFormRef() bool`

HasCamundaFormRef returns a boolean if a field has been set.

### GetTenantId

`func (o *TaskDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TaskDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TaskDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TaskDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TaskDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TaskDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


