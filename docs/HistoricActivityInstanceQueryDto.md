# HistoricActivityInstanceQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityInstanceId** | Pointer to **NullableString** | Filter by activity instance id. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | Filter by process instance id. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter by process definition id. | [optional] 
**ExecutionId** | Pointer to **NullableString** | Filter by the id of the execution that executed the activity instance. | [optional] 
**ActivityId** | Pointer to **NullableString** | Filter by the activity id (according to BPMN 2.0 XML). | [optional] 
**ActivityName** | Pointer to **NullableString** | Filter by the activity name (according to BPMN 2.0 XML). | [optional] 
**ActivityType** | Pointer to **NullableString** | Filter by activity type. | [optional] 
**TaskAssignee** | Pointer to **NullableString** | Only include activity instances that are user tasks and assigned to a given user. | [optional] 
**Finished** | Pointer to **NullableBool** | Only include finished activity instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | [optional] 
**Unfinished** | Pointer to **NullableBool** | Only include unfinished activity instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | [optional] 
**Canceled** | Pointer to **NullableBool** | Only include canceled activity instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | [optional] 
**CompleteScope** | Pointer to **NullableBool** | Only include activity instances which completed a scope. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; behaves the same as when the property is not set. | [optional] 
**StartedBefore** | Pointer to **NullableTime** | Restrict to instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**StartedAfter** | Pointer to **NullableTime** | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**FinishedBefore** | Pointer to **NullableTime** | Restrict to instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**FinishedAfter** | Pointer to **NullableTime** | Restrict to instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Must be a JSON array of Strings. An activity instance must have one of the given tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include historic activity instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Sorting** | Pointer to [**[]HistoricActivityInstanceQueryDtoSortingInner**](HistoricActivityInstanceQueryDtoSortingInner.md) | Apply sorting of the result | [optional] 

## Methods

### NewHistoricActivityInstanceQueryDto

`func NewHistoricActivityInstanceQueryDto() *HistoricActivityInstanceQueryDto`

NewHistoricActivityInstanceQueryDto instantiates a new HistoricActivityInstanceQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricActivityInstanceQueryDtoWithDefaults

`func NewHistoricActivityInstanceQueryDtoWithDefaults() *HistoricActivityInstanceQueryDto`

NewHistoricActivityInstanceQueryDtoWithDefaults instantiates a new HistoricActivityInstanceQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityInstanceId

`func (o *HistoricActivityInstanceQueryDto) GetActivityInstanceId() string`

GetActivityInstanceId returns the ActivityInstanceId field if non-nil, zero value otherwise.

### GetActivityInstanceIdOk

`func (o *HistoricActivityInstanceQueryDto) GetActivityInstanceIdOk() (*string, bool)`

GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceId

`func (o *HistoricActivityInstanceQueryDto) SetActivityInstanceId(v string)`

SetActivityInstanceId sets ActivityInstanceId field to given value.

### HasActivityInstanceId

`func (o *HistoricActivityInstanceQueryDto) HasActivityInstanceId() bool`

HasActivityInstanceId returns a boolean if a field has been set.

### SetActivityInstanceIdNil

`func (o *HistoricActivityInstanceQueryDto) SetActivityInstanceIdNil(b bool)`

 SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil

### UnsetActivityInstanceId
`func (o *HistoricActivityInstanceQueryDto) UnsetActivityInstanceId()`

UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricActivityInstanceQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricActivityInstanceQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricActivityInstanceQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricActivityInstanceQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricActivityInstanceQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricActivityInstanceQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricActivityInstanceQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricActivityInstanceQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricActivityInstanceQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricActivityInstanceQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricActivityInstanceQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricActivityInstanceQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetExecutionId

`func (o *HistoricActivityInstanceQueryDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HistoricActivityInstanceQueryDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HistoricActivityInstanceQueryDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HistoricActivityInstanceQueryDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HistoricActivityInstanceQueryDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HistoricActivityInstanceQueryDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetActivityId

`func (o *HistoricActivityInstanceQueryDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *HistoricActivityInstanceQueryDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *HistoricActivityInstanceQueryDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *HistoricActivityInstanceQueryDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *HistoricActivityInstanceQueryDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *HistoricActivityInstanceQueryDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetActivityName

`func (o *HistoricActivityInstanceQueryDto) GetActivityName() string`

GetActivityName returns the ActivityName field if non-nil, zero value otherwise.

### GetActivityNameOk

`func (o *HistoricActivityInstanceQueryDto) GetActivityNameOk() (*string, bool)`

GetActivityNameOk returns a tuple with the ActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityName

`func (o *HistoricActivityInstanceQueryDto) SetActivityName(v string)`

SetActivityName sets ActivityName field to given value.

### HasActivityName

`func (o *HistoricActivityInstanceQueryDto) HasActivityName() bool`

HasActivityName returns a boolean if a field has been set.

### SetActivityNameNil

`func (o *HistoricActivityInstanceQueryDto) SetActivityNameNil(b bool)`

 SetActivityNameNil sets the value for ActivityName to be an explicit nil

### UnsetActivityName
`func (o *HistoricActivityInstanceQueryDto) UnsetActivityName()`

UnsetActivityName ensures that no value is present for ActivityName, not even an explicit nil
### GetActivityType

`func (o *HistoricActivityInstanceQueryDto) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *HistoricActivityInstanceQueryDto) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *HistoricActivityInstanceQueryDto) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *HistoricActivityInstanceQueryDto) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### SetActivityTypeNil

`func (o *HistoricActivityInstanceQueryDto) SetActivityTypeNil(b bool)`

 SetActivityTypeNil sets the value for ActivityType to be an explicit nil

### UnsetActivityType
`func (o *HistoricActivityInstanceQueryDto) UnsetActivityType()`

UnsetActivityType ensures that no value is present for ActivityType, not even an explicit nil
### GetTaskAssignee

`func (o *HistoricActivityInstanceQueryDto) GetTaskAssignee() string`

GetTaskAssignee returns the TaskAssignee field if non-nil, zero value otherwise.

### GetTaskAssigneeOk

`func (o *HistoricActivityInstanceQueryDto) GetTaskAssigneeOk() (*string, bool)`

GetTaskAssigneeOk returns a tuple with the TaskAssignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAssignee

`func (o *HistoricActivityInstanceQueryDto) SetTaskAssignee(v string)`

SetTaskAssignee sets TaskAssignee field to given value.

### HasTaskAssignee

`func (o *HistoricActivityInstanceQueryDto) HasTaskAssignee() bool`

HasTaskAssignee returns a boolean if a field has been set.

### SetTaskAssigneeNil

`func (o *HistoricActivityInstanceQueryDto) SetTaskAssigneeNil(b bool)`

 SetTaskAssigneeNil sets the value for TaskAssignee to be an explicit nil

### UnsetTaskAssignee
`func (o *HistoricActivityInstanceQueryDto) UnsetTaskAssignee()`

UnsetTaskAssignee ensures that no value is present for TaskAssignee, not even an explicit nil
### GetFinished

`func (o *HistoricActivityInstanceQueryDto) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *HistoricActivityInstanceQueryDto) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *HistoricActivityInstanceQueryDto) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *HistoricActivityInstanceQueryDto) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *HistoricActivityInstanceQueryDto) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *HistoricActivityInstanceQueryDto) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetUnfinished

`func (o *HistoricActivityInstanceQueryDto) GetUnfinished() bool`

GetUnfinished returns the Unfinished field if non-nil, zero value otherwise.

### GetUnfinishedOk

`func (o *HistoricActivityInstanceQueryDto) GetUnfinishedOk() (*bool, bool)`

GetUnfinishedOk returns a tuple with the Unfinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinished

`func (o *HistoricActivityInstanceQueryDto) SetUnfinished(v bool)`

SetUnfinished sets Unfinished field to given value.

### HasUnfinished

`func (o *HistoricActivityInstanceQueryDto) HasUnfinished() bool`

HasUnfinished returns a boolean if a field has been set.

### SetUnfinishedNil

`func (o *HistoricActivityInstanceQueryDto) SetUnfinishedNil(b bool)`

 SetUnfinishedNil sets the value for Unfinished to be an explicit nil

### UnsetUnfinished
`func (o *HistoricActivityInstanceQueryDto) UnsetUnfinished()`

UnsetUnfinished ensures that no value is present for Unfinished, not even an explicit nil
### GetCanceled

`func (o *HistoricActivityInstanceQueryDto) GetCanceled() bool`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *HistoricActivityInstanceQueryDto) GetCanceledOk() (*bool, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *HistoricActivityInstanceQueryDto) SetCanceled(v bool)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *HistoricActivityInstanceQueryDto) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### SetCanceledNil

`func (o *HistoricActivityInstanceQueryDto) SetCanceledNil(b bool)`

 SetCanceledNil sets the value for Canceled to be an explicit nil

### UnsetCanceled
`func (o *HistoricActivityInstanceQueryDto) UnsetCanceled()`

UnsetCanceled ensures that no value is present for Canceled, not even an explicit nil
### GetCompleteScope

`func (o *HistoricActivityInstanceQueryDto) GetCompleteScope() bool`

GetCompleteScope returns the CompleteScope field if non-nil, zero value otherwise.

### GetCompleteScopeOk

`func (o *HistoricActivityInstanceQueryDto) GetCompleteScopeOk() (*bool, bool)`

GetCompleteScopeOk returns a tuple with the CompleteScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteScope

`func (o *HistoricActivityInstanceQueryDto) SetCompleteScope(v bool)`

SetCompleteScope sets CompleteScope field to given value.

### HasCompleteScope

`func (o *HistoricActivityInstanceQueryDto) HasCompleteScope() bool`

HasCompleteScope returns a boolean if a field has been set.

### SetCompleteScopeNil

`func (o *HistoricActivityInstanceQueryDto) SetCompleteScopeNil(b bool)`

 SetCompleteScopeNil sets the value for CompleteScope to be an explicit nil

### UnsetCompleteScope
`func (o *HistoricActivityInstanceQueryDto) UnsetCompleteScope()`

UnsetCompleteScope ensures that no value is present for CompleteScope, not even an explicit nil
### GetStartedBefore

`func (o *HistoricActivityInstanceQueryDto) GetStartedBefore() time.Time`

GetStartedBefore returns the StartedBefore field if non-nil, zero value otherwise.

### GetStartedBeforeOk

`func (o *HistoricActivityInstanceQueryDto) GetStartedBeforeOk() (*time.Time, bool)`

GetStartedBeforeOk returns a tuple with the StartedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBefore

`func (o *HistoricActivityInstanceQueryDto) SetStartedBefore(v time.Time)`

SetStartedBefore sets StartedBefore field to given value.

### HasStartedBefore

`func (o *HistoricActivityInstanceQueryDto) HasStartedBefore() bool`

HasStartedBefore returns a boolean if a field has been set.

### SetStartedBeforeNil

`func (o *HistoricActivityInstanceQueryDto) SetStartedBeforeNil(b bool)`

 SetStartedBeforeNil sets the value for StartedBefore to be an explicit nil

### UnsetStartedBefore
`func (o *HistoricActivityInstanceQueryDto) UnsetStartedBefore()`

UnsetStartedBefore ensures that no value is present for StartedBefore, not even an explicit nil
### GetStartedAfter

`func (o *HistoricActivityInstanceQueryDto) GetStartedAfter() time.Time`

GetStartedAfter returns the StartedAfter field if non-nil, zero value otherwise.

### GetStartedAfterOk

`func (o *HistoricActivityInstanceQueryDto) GetStartedAfterOk() (*time.Time, bool)`

GetStartedAfterOk returns a tuple with the StartedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAfter

`func (o *HistoricActivityInstanceQueryDto) SetStartedAfter(v time.Time)`

SetStartedAfter sets StartedAfter field to given value.

### HasStartedAfter

`func (o *HistoricActivityInstanceQueryDto) HasStartedAfter() bool`

HasStartedAfter returns a boolean if a field has been set.

### SetStartedAfterNil

`func (o *HistoricActivityInstanceQueryDto) SetStartedAfterNil(b bool)`

 SetStartedAfterNil sets the value for StartedAfter to be an explicit nil

### UnsetStartedAfter
`func (o *HistoricActivityInstanceQueryDto) UnsetStartedAfter()`

UnsetStartedAfter ensures that no value is present for StartedAfter, not even an explicit nil
### GetFinishedBefore

`func (o *HistoricActivityInstanceQueryDto) GetFinishedBefore() time.Time`

GetFinishedBefore returns the FinishedBefore field if non-nil, zero value otherwise.

### GetFinishedBeforeOk

`func (o *HistoricActivityInstanceQueryDto) GetFinishedBeforeOk() (*time.Time, bool)`

GetFinishedBeforeOk returns a tuple with the FinishedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedBefore

`func (o *HistoricActivityInstanceQueryDto) SetFinishedBefore(v time.Time)`

SetFinishedBefore sets FinishedBefore field to given value.

### HasFinishedBefore

`func (o *HistoricActivityInstanceQueryDto) HasFinishedBefore() bool`

HasFinishedBefore returns a boolean if a field has been set.

### SetFinishedBeforeNil

`func (o *HistoricActivityInstanceQueryDto) SetFinishedBeforeNil(b bool)`

 SetFinishedBeforeNil sets the value for FinishedBefore to be an explicit nil

### UnsetFinishedBefore
`func (o *HistoricActivityInstanceQueryDto) UnsetFinishedBefore()`

UnsetFinishedBefore ensures that no value is present for FinishedBefore, not even an explicit nil
### GetFinishedAfter

`func (o *HistoricActivityInstanceQueryDto) GetFinishedAfter() time.Time`

GetFinishedAfter returns the FinishedAfter field if non-nil, zero value otherwise.

### GetFinishedAfterOk

`func (o *HistoricActivityInstanceQueryDto) GetFinishedAfterOk() (*time.Time, bool)`

GetFinishedAfterOk returns a tuple with the FinishedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAfter

`func (o *HistoricActivityInstanceQueryDto) SetFinishedAfter(v time.Time)`

SetFinishedAfter sets FinishedAfter field to given value.

### HasFinishedAfter

`func (o *HistoricActivityInstanceQueryDto) HasFinishedAfter() bool`

HasFinishedAfter returns a boolean if a field has been set.

### SetFinishedAfterNil

`func (o *HistoricActivityInstanceQueryDto) SetFinishedAfterNil(b bool)`

 SetFinishedAfterNil sets the value for FinishedAfter to be an explicit nil

### UnsetFinishedAfter
`func (o *HistoricActivityInstanceQueryDto) UnsetFinishedAfter()`

UnsetFinishedAfter ensures that no value is present for FinishedAfter, not even an explicit nil
### GetTenantIdIn

`func (o *HistoricActivityInstanceQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *HistoricActivityInstanceQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *HistoricActivityInstanceQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *HistoricActivityInstanceQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *HistoricActivityInstanceQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *HistoricActivityInstanceQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *HistoricActivityInstanceQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *HistoricActivityInstanceQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *HistoricActivityInstanceQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *HistoricActivityInstanceQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *HistoricActivityInstanceQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *HistoricActivityInstanceQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetSorting

`func (o *HistoricActivityInstanceQueryDto) GetSorting() []HistoricActivityInstanceQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *HistoricActivityInstanceQueryDto) GetSortingOk() (*[]HistoricActivityInstanceQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *HistoricActivityInstanceQueryDto) SetSorting(v []HistoricActivityInstanceQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *HistoricActivityInstanceQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *HistoricActivityInstanceQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *HistoricActivityInstanceQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


