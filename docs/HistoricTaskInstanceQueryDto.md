# HistoricTaskInstanceQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **NullableString** | Filter by task id. | [optional] 
**TaskParentTaskId** | Pointer to **NullableString** | Filter by parent task id. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | Filter by process instance id. | [optional] 
**ProcessInstanceBusinessKey** | Pointer to **NullableString** | Filter by process instance business key. | [optional] 
**ProcessInstanceBusinessKeyIn** | Pointer to **[]string** | Filter by process instances with one of the give business keys. The keys need to be in a comma-separated list. | [optional] 
**ProcessInstanceBusinessKeyLike** | Pointer to **NullableString** | Filter by  process instance business key that has the parameter value as a substring. | [optional] 
**ExecutionId** | Pointer to **NullableString** | Filter by the id of the execution that executed the task. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter by process definition id. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Restrict to tasks that belong to a process definition with the given key. | [optional] 
**ProcessDefinitionName** | Pointer to **NullableString** | Restrict to tasks that belong to a process definition with the given name. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | Filter by case instance id. | [optional] 
**CaseExecutionId** | Pointer to **NullableString** | Filter by the id of the case execution that executed the task. | [optional] 
**CaseDefinitionId** | Pointer to **NullableString** | Filter by case definition id. | [optional] 
**CaseDefinitionKey** | Pointer to **NullableString** | Restrict to tasks that belong to a case definition with the given key. | [optional] 
**CaseDefinitionName** | Pointer to **NullableString** | Restrict to tasks that belong to a case definition with the given name. | [optional] 
**ActivityInstanceIdIn** | Pointer to **[]string** | Only include tasks which belong to one of the passed  activity instance ids. | [optional] 
**TaskName** | Pointer to **NullableString** | Restrict to tasks that have the given name. | [optional] 
**TaskNameLike** | Pointer to **NullableString** | Restrict to tasks that have a name with the given parameter value as substring. | [optional] 
**TaskDescription** | Pointer to **NullableString** | Restrict to tasks that have the given description. | [optional] 
**TaskDescriptionLike** | Pointer to **NullableString** | Restrict to tasks that have a description that has the parameter value as a substring. | [optional] 
**TaskDefinitionKey** | Pointer to **NullableString** | Restrict to tasks that have the given key. | [optional] 
**TaskDefinitionKeyIn** | Pointer to **[]string** | Restrict to tasks that have one of the passed  task definition keys. | [optional] 
**TaskDeleteReason** | Pointer to **NullableString** | Restrict to tasks that have the given delete reason. | [optional] 
**TaskDeleteReasonLike** | Pointer to **NullableString** | Restrict to tasks that have a delete reason that has the parameter value as a substring. | [optional] 
**TaskAssignee** | Pointer to **NullableString** | Restrict to tasks that the given user is assigned to. | [optional] 
**TaskAssigneeLike** | Pointer to **NullableString** | Restrict to tasks that are assigned to users with the parameter value as a substring. | [optional] 
**TaskOwner** | Pointer to **NullableString** | Restrict to tasks that the given user owns. | [optional] 
**TaskOwnerLike** | Pointer to **NullableString** | Restrict to tasks that are owned by users with the parameter value as a substring. | [optional] 
**TaskPriority** | Pointer to **NullableInt32** | Restrict to tasks that have the given priority. | [optional] 
**Assigned** | Pointer to **NullableBool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are assigned. | [optional] 
**Unassigned** | Pointer to **NullableBool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are unassigned. | [optional] 
**Finished** | Pointer to **NullableBool** | Only include finished tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Unfinished** | Pointer to **NullableBool** | Only include unfinished tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**ProcessFinished** | Pointer to **NullableBool** | Only include tasks of finished processes. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**ProcessUnfinished** | Pointer to **NullableBool** | Only include tasks of unfinished processes. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**TaskDueDate** | Pointer to **NullableTime** | Restrict to tasks that are due on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**TaskDueDateBefore** | Pointer to **NullableTime** | Restrict to tasks that are due before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**TaskDueDateAfter** | Pointer to **NullableTime** | Restrict to tasks that are due after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**WithoutTaskDueDate** | Pointer to **NullableBool** | Only include tasks which have no due date. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**TaskFollowUpDate** | Pointer to **NullableTime** | Restrict to tasks that have a followUp date on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**TaskFollowUpDateBefore** | Pointer to **NullableTime** | Restrict to tasks that have a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**TaskFollowUpDateAfter** | Pointer to **NullableTime** | Restrict to tasks that have a followUp date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**StartedBefore** | Pointer to **NullableTime** | Restrict to tasks that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**StartedAfter** | Pointer to **NullableTime** | Restrict to tasks that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**FinishedBefore** | Pointer to **NullableTime** | Restrict to tasks that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**FinishedAfter** | Pointer to **NullableTime** | Restrict to tasks that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Filter by a  list of tenant ids. A task instance must have one of the given tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include historic task instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**TaskVariables** | Pointer to [**[]VariableQueryParameterDto**](VariableQueryParameterDto.md) | Only include tasks that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. **Note:** Values are always treated as &#x60;String&#x60; objects on server side.   Valid operator values are: * &#x60;eq&#x60; - equal to; * &#x60;neq&#x60; - not equal to; * &#x60;gt&#x60; - greater than; * &#x60;gteq&#x60; - greater than or equal to; * &#x60;lt&#x60; - lower than; * &#x60;lteq&#x60; - lower than or equal to; * &#x60;like&#x60;.  &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | [optional] 
**ProcessVariables** | Pointer to [**[]VariableQueryParameterDto**](VariableQueryParameterDto.md) | Only include tasks that belong to process instances that have variables with certain values. Variable filtering expressions are comma-separated and are structured as follows:  A valid parameter value has the form &#x60;key_operator_value&#x60;. &#x60;key&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. **Note:** Values are always treated as &#x60;String&#x60; objects on server side.   Valid operator values are: * &#x60;eq&#x60; - equal to; * &#x60;neq&#x60; - not equal to; * &#x60;gt&#x60; - greater than; * &#x60;gteq&#x60; - greater than or equal to; * &#x60;lt&#x60; - lower than; * &#x60;lteq&#x60; - lower than or equal to; * &#x60;like&#x60;; * &#x60;notLike&#x60;.  &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | [optional] 
**VariableNamesIgnoreCase** | Pointer to **NullableBool** | Match the variable name provided in &#x60;taskVariables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | [optional] 
**VariableValuesIgnoreCase** | Pointer to **NullableBool** | Match the variable value provided in &#x60;taskVariables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | [optional] 
**TaskInvolvedUser** | Pointer to **NullableString** | Restrict to tasks with a historic identity link to the given user. | [optional] 
**TaskInvolvedGroup** | Pointer to **NullableString** | Restrict to tasks with a historic identity link to the given group. | [optional] 
**TaskHadCandidateUser** | Pointer to **NullableString** | Restrict to tasks with a historic identity link to the given candidate user. | [optional] 
**TaskHadCandidateGroup** | Pointer to **NullableString** | Restrict to tasks with a historic identity link to the given candidate group. | [optional] 
**WithCandidateGroups** | Pointer to **NullableBool** | Only include tasks which have a candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**WithoutCandidateGroups** | Pointer to **NullableBool** | Only include tasks which have no candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**OrQueries** | Pointer to [**[]HistoricTaskInstanceQueryDto**](HistoricTaskInstanceQueryDto.md) | A JSON array of nested historic task instance queries with OR semantics.  A task instance matches a nested query if it fulfills at least one of the query&#39;s predicates.  With multiple nested queries, a task instance must fulfill at least one predicate of each query ([Conjunctive Normal Form](https://en.wikipedia.org/wiki/Conjunctive_normal_form)).  All task instance query properties can be used except for: &#x60;sorting&#x60;, &#x60;withCandidateGroups&#x60;, &#x60; withoutCandidateGroups&#x60;.  See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/process-engine-api/#or-queries) for more information about OR queries. | [optional] 
**Sorting** | Pointer to [**[]HistoricTaskInstanceQueryDtoSortingInner**](HistoricTaskInstanceQueryDtoSortingInner.md) | An array of criteria to sort the result by. Each element of the array is                     an object that specifies one ordering. The position in the array                     identifies the rank of an ordering, i.e., whether it is primary, secondary,                     etc. Sorting has no effect for &#x60;count&#x60; endpoints | [optional] 

## Methods

### NewHistoricTaskInstanceQueryDto

`func NewHistoricTaskInstanceQueryDto() *HistoricTaskInstanceQueryDto`

NewHistoricTaskInstanceQueryDto instantiates a new HistoricTaskInstanceQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricTaskInstanceQueryDtoWithDefaults

`func NewHistoricTaskInstanceQueryDtoWithDefaults() *HistoricTaskInstanceQueryDto`

NewHistoricTaskInstanceQueryDtoWithDefaults instantiates a new HistoricTaskInstanceQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *HistoricTaskInstanceQueryDto) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *HistoricTaskInstanceQueryDto) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *HistoricTaskInstanceQueryDto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetTaskParentTaskId

`func (o *HistoricTaskInstanceQueryDto) GetTaskParentTaskId() string`

GetTaskParentTaskId returns the TaskParentTaskId field if non-nil, zero value otherwise.

### GetTaskParentTaskIdOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskParentTaskIdOk() (*string, bool)`

GetTaskParentTaskIdOk returns a tuple with the TaskParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskParentTaskId

`func (o *HistoricTaskInstanceQueryDto) SetTaskParentTaskId(v string)`

SetTaskParentTaskId sets TaskParentTaskId field to given value.

### HasTaskParentTaskId

`func (o *HistoricTaskInstanceQueryDto) HasTaskParentTaskId() bool`

HasTaskParentTaskId returns a boolean if a field has been set.

### SetTaskParentTaskIdNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskParentTaskIdNil(b bool)`

 SetTaskParentTaskIdNil sets the value for TaskParentTaskId to be an explicit nil

### UnsetTaskParentTaskId
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskParentTaskId()`

UnsetTaskParentTaskId ensures that no value is present for TaskParentTaskId, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricTaskInstanceQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricTaskInstanceQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricTaskInstanceQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricTaskInstanceQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricTaskInstanceQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricTaskInstanceQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessInstanceBusinessKey

`func (o *HistoricTaskInstanceQueryDto) GetProcessInstanceBusinessKey() string`

GetProcessInstanceBusinessKey returns the ProcessInstanceBusinessKey field if non-nil, zero value otherwise.

### GetProcessInstanceBusinessKeyOk

`func (o *HistoricTaskInstanceQueryDto) GetProcessInstanceBusinessKeyOk() (*string, bool)`

GetProcessInstanceBusinessKeyOk returns a tuple with the ProcessInstanceBusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceBusinessKey

`func (o *HistoricTaskInstanceQueryDto) SetProcessInstanceBusinessKey(v string)`

SetProcessInstanceBusinessKey sets ProcessInstanceBusinessKey field to given value.

### HasProcessInstanceBusinessKey

`func (o *HistoricTaskInstanceQueryDto) HasProcessInstanceBusinessKey() bool`

HasProcessInstanceBusinessKey returns a boolean if a field has been set.

### SetProcessInstanceBusinessKeyNil

`func (o *HistoricTaskInstanceQueryDto) SetProcessInstanceBusinessKeyNil(b bool)`

 SetProcessInstanceBusinessKeyNil sets the value for ProcessInstanceBusinessKey to be an explicit nil

### UnsetProcessInstanceBusinessKey
`func (o *HistoricTaskInstanceQueryDto) UnsetProcessInstanceBusinessKey()`

UnsetProcessInstanceBusinessKey ensures that no value is present for ProcessInstanceBusinessKey, not even an explicit nil
### GetProcessInstanceBusinessKeyIn

`func (o *HistoricTaskInstanceQueryDto) GetProcessInstanceBusinessKeyIn() []string`

GetProcessInstanceBusinessKeyIn returns the ProcessInstanceBusinessKeyIn field if non-nil, zero value otherwise.

### GetProcessInstanceBusinessKeyInOk

`func (o *HistoricTaskInstanceQueryDto) GetProcessInstanceBusinessKeyInOk() (*[]string, bool)`

GetProcessInstanceBusinessKeyInOk returns a tuple with the ProcessInstanceBusinessKeyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceBusinessKeyIn

`func (o *HistoricTaskInstanceQueryDto) SetProcessInstanceBusinessKeyIn(v []string)`

SetProcessInstanceBusinessKeyIn sets ProcessInstanceBusinessKeyIn field to given value.

### HasProcessInstanceBusinessKeyIn

`func (o *HistoricTaskInstanceQueryDto) HasProcessInstanceBusinessKeyIn() bool`

HasProcessInstanceBusinessKeyIn returns a boolean if a field has been set.

### SetProcessInstanceBusinessKeyInNil

`func (o *HistoricTaskInstanceQueryDto) SetProcessInstanceBusinessKeyInNil(b bool)`

 SetProcessInstanceBusinessKeyInNil sets the value for ProcessInstanceBusinessKeyIn to be an explicit nil

### UnsetProcessInstanceBusinessKeyIn
`func (o *HistoricTaskInstanceQueryDto) UnsetProcessInstanceBusinessKeyIn()`

UnsetProcessInstanceBusinessKeyIn ensures that no value is present for ProcessInstanceBusinessKeyIn, not even an explicit nil
### GetProcessInstanceBusinessKeyLike

`func (o *HistoricTaskInstanceQueryDto) GetProcessInstanceBusinessKeyLike() string`

GetProcessInstanceBusinessKeyLike returns the ProcessInstanceBusinessKeyLike field if non-nil, zero value otherwise.

### GetProcessInstanceBusinessKeyLikeOk

`func (o *HistoricTaskInstanceQueryDto) GetProcessInstanceBusinessKeyLikeOk() (*string, bool)`

GetProcessInstanceBusinessKeyLikeOk returns a tuple with the ProcessInstanceBusinessKeyLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceBusinessKeyLike

`func (o *HistoricTaskInstanceQueryDto) SetProcessInstanceBusinessKeyLike(v string)`

SetProcessInstanceBusinessKeyLike sets ProcessInstanceBusinessKeyLike field to given value.

### HasProcessInstanceBusinessKeyLike

`func (o *HistoricTaskInstanceQueryDto) HasProcessInstanceBusinessKeyLike() bool`

HasProcessInstanceBusinessKeyLike returns a boolean if a field has been set.

### SetProcessInstanceBusinessKeyLikeNil

`func (o *HistoricTaskInstanceQueryDto) SetProcessInstanceBusinessKeyLikeNil(b bool)`

 SetProcessInstanceBusinessKeyLikeNil sets the value for ProcessInstanceBusinessKeyLike to be an explicit nil

### UnsetProcessInstanceBusinessKeyLike
`func (o *HistoricTaskInstanceQueryDto) UnsetProcessInstanceBusinessKeyLike()`

UnsetProcessInstanceBusinessKeyLike ensures that no value is present for ProcessInstanceBusinessKeyLike, not even an explicit nil
### GetExecutionId

`func (o *HistoricTaskInstanceQueryDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HistoricTaskInstanceQueryDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HistoricTaskInstanceQueryDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HistoricTaskInstanceQueryDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HistoricTaskInstanceQueryDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HistoricTaskInstanceQueryDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricTaskInstanceQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricTaskInstanceQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricTaskInstanceQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricTaskInstanceQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricTaskInstanceQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricTaskInstanceQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricTaskInstanceQueryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricTaskInstanceQueryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricTaskInstanceQueryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricTaskInstanceQueryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricTaskInstanceQueryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricTaskInstanceQueryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionName

`func (o *HistoricTaskInstanceQueryDto) GetProcessDefinitionName() string`

GetProcessDefinitionName returns the ProcessDefinitionName field if non-nil, zero value otherwise.

### GetProcessDefinitionNameOk

`func (o *HistoricTaskInstanceQueryDto) GetProcessDefinitionNameOk() (*string, bool)`

GetProcessDefinitionNameOk returns a tuple with the ProcessDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionName

`func (o *HistoricTaskInstanceQueryDto) SetProcessDefinitionName(v string)`

SetProcessDefinitionName sets ProcessDefinitionName field to given value.

### HasProcessDefinitionName

`func (o *HistoricTaskInstanceQueryDto) HasProcessDefinitionName() bool`

HasProcessDefinitionName returns a boolean if a field has been set.

### SetProcessDefinitionNameNil

`func (o *HistoricTaskInstanceQueryDto) SetProcessDefinitionNameNil(b bool)`

 SetProcessDefinitionNameNil sets the value for ProcessDefinitionName to be an explicit nil

### UnsetProcessDefinitionName
`func (o *HistoricTaskInstanceQueryDto) UnsetProcessDefinitionName()`

UnsetProcessDefinitionName ensures that no value is present for ProcessDefinitionName, not even an explicit nil
### GetCaseInstanceId

`func (o *HistoricTaskInstanceQueryDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *HistoricTaskInstanceQueryDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *HistoricTaskInstanceQueryDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *HistoricTaskInstanceQueryDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *HistoricTaskInstanceQueryDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *HistoricTaskInstanceQueryDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetCaseExecutionId

`func (o *HistoricTaskInstanceQueryDto) GetCaseExecutionId() string`

GetCaseExecutionId returns the CaseExecutionId field if non-nil, zero value otherwise.

### GetCaseExecutionIdOk

`func (o *HistoricTaskInstanceQueryDto) GetCaseExecutionIdOk() (*string, bool)`

GetCaseExecutionIdOk returns a tuple with the CaseExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExecutionId

`func (o *HistoricTaskInstanceQueryDto) SetCaseExecutionId(v string)`

SetCaseExecutionId sets CaseExecutionId field to given value.

### HasCaseExecutionId

`func (o *HistoricTaskInstanceQueryDto) HasCaseExecutionId() bool`

HasCaseExecutionId returns a boolean if a field has been set.

### SetCaseExecutionIdNil

`func (o *HistoricTaskInstanceQueryDto) SetCaseExecutionIdNil(b bool)`

 SetCaseExecutionIdNil sets the value for CaseExecutionId to be an explicit nil

### UnsetCaseExecutionId
`func (o *HistoricTaskInstanceQueryDto) UnsetCaseExecutionId()`

UnsetCaseExecutionId ensures that no value is present for CaseExecutionId, not even an explicit nil
### GetCaseDefinitionId

`func (o *HistoricTaskInstanceQueryDto) GetCaseDefinitionId() string`

GetCaseDefinitionId returns the CaseDefinitionId field if non-nil, zero value otherwise.

### GetCaseDefinitionIdOk

`func (o *HistoricTaskInstanceQueryDto) GetCaseDefinitionIdOk() (*string, bool)`

GetCaseDefinitionIdOk returns a tuple with the CaseDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionId

`func (o *HistoricTaskInstanceQueryDto) SetCaseDefinitionId(v string)`

SetCaseDefinitionId sets CaseDefinitionId field to given value.

### HasCaseDefinitionId

`func (o *HistoricTaskInstanceQueryDto) HasCaseDefinitionId() bool`

HasCaseDefinitionId returns a boolean if a field has been set.

### SetCaseDefinitionIdNil

`func (o *HistoricTaskInstanceQueryDto) SetCaseDefinitionIdNil(b bool)`

 SetCaseDefinitionIdNil sets the value for CaseDefinitionId to be an explicit nil

### UnsetCaseDefinitionId
`func (o *HistoricTaskInstanceQueryDto) UnsetCaseDefinitionId()`

UnsetCaseDefinitionId ensures that no value is present for CaseDefinitionId, not even an explicit nil
### GetCaseDefinitionKey

`func (o *HistoricTaskInstanceQueryDto) GetCaseDefinitionKey() string`

GetCaseDefinitionKey returns the CaseDefinitionKey field if non-nil, zero value otherwise.

### GetCaseDefinitionKeyOk

`func (o *HistoricTaskInstanceQueryDto) GetCaseDefinitionKeyOk() (*string, bool)`

GetCaseDefinitionKeyOk returns a tuple with the CaseDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionKey

`func (o *HistoricTaskInstanceQueryDto) SetCaseDefinitionKey(v string)`

SetCaseDefinitionKey sets CaseDefinitionKey field to given value.

### HasCaseDefinitionKey

`func (o *HistoricTaskInstanceQueryDto) HasCaseDefinitionKey() bool`

HasCaseDefinitionKey returns a boolean if a field has been set.

### SetCaseDefinitionKeyNil

`func (o *HistoricTaskInstanceQueryDto) SetCaseDefinitionKeyNil(b bool)`

 SetCaseDefinitionKeyNil sets the value for CaseDefinitionKey to be an explicit nil

### UnsetCaseDefinitionKey
`func (o *HistoricTaskInstanceQueryDto) UnsetCaseDefinitionKey()`

UnsetCaseDefinitionKey ensures that no value is present for CaseDefinitionKey, not even an explicit nil
### GetCaseDefinitionName

`func (o *HistoricTaskInstanceQueryDto) GetCaseDefinitionName() string`

GetCaseDefinitionName returns the CaseDefinitionName field if non-nil, zero value otherwise.

### GetCaseDefinitionNameOk

`func (o *HistoricTaskInstanceQueryDto) GetCaseDefinitionNameOk() (*string, bool)`

GetCaseDefinitionNameOk returns a tuple with the CaseDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionName

`func (o *HistoricTaskInstanceQueryDto) SetCaseDefinitionName(v string)`

SetCaseDefinitionName sets CaseDefinitionName field to given value.

### HasCaseDefinitionName

`func (o *HistoricTaskInstanceQueryDto) HasCaseDefinitionName() bool`

HasCaseDefinitionName returns a boolean if a field has been set.

### SetCaseDefinitionNameNil

`func (o *HistoricTaskInstanceQueryDto) SetCaseDefinitionNameNil(b bool)`

 SetCaseDefinitionNameNil sets the value for CaseDefinitionName to be an explicit nil

### UnsetCaseDefinitionName
`func (o *HistoricTaskInstanceQueryDto) UnsetCaseDefinitionName()`

UnsetCaseDefinitionName ensures that no value is present for CaseDefinitionName, not even an explicit nil
### GetActivityInstanceIdIn

`func (o *HistoricTaskInstanceQueryDto) GetActivityInstanceIdIn() []string`

GetActivityInstanceIdIn returns the ActivityInstanceIdIn field if non-nil, zero value otherwise.

### GetActivityInstanceIdInOk

`func (o *HistoricTaskInstanceQueryDto) GetActivityInstanceIdInOk() (*[]string, bool)`

GetActivityInstanceIdInOk returns a tuple with the ActivityInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceIdIn

`func (o *HistoricTaskInstanceQueryDto) SetActivityInstanceIdIn(v []string)`

SetActivityInstanceIdIn sets ActivityInstanceIdIn field to given value.

### HasActivityInstanceIdIn

`func (o *HistoricTaskInstanceQueryDto) HasActivityInstanceIdIn() bool`

HasActivityInstanceIdIn returns a boolean if a field has been set.

### SetActivityInstanceIdInNil

`func (o *HistoricTaskInstanceQueryDto) SetActivityInstanceIdInNil(b bool)`

 SetActivityInstanceIdInNil sets the value for ActivityInstanceIdIn to be an explicit nil

### UnsetActivityInstanceIdIn
`func (o *HistoricTaskInstanceQueryDto) UnsetActivityInstanceIdIn()`

UnsetActivityInstanceIdIn ensures that no value is present for ActivityInstanceIdIn, not even an explicit nil
### GetTaskName

`func (o *HistoricTaskInstanceQueryDto) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *HistoricTaskInstanceQueryDto) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *HistoricTaskInstanceQueryDto) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### SetTaskNameNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskNameNil(b bool)`

 SetTaskNameNil sets the value for TaskName to be an explicit nil

### UnsetTaskName
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskName()`

UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil
### GetTaskNameLike

`func (o *HistoricTaskInstanceQueryDto) GetTaskNameLike() string`

GetTaskNameLike returns the TaskNameLike field if non-nil, zero value otherwise.

### GetTaskNameLikeOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskNameLikeOk() (*string, bool)`

GetTaskNameLikeOk returns a tuple with the TaskNameLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskNameLike

`func (o *HistoricTaskInstanceQueryDto) SetTaskNameLike(v string)`

SetTaskNameLike sets TaskNameLike field to given value.

### HasTaskNameLike

`func (o *HistoricTaskInstanceQueryDto) HasTaskNameLike() bool`

HasTaskNameLike returns a boolean if a field has been set.

### SetTaskNameLikeNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskNameLikeNil(b bool)`

 SetTaskNameLikeNil sets the value for TaskNameLike to be an explicit nil

### UnsetTaskNameLike
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskNameLike()`

UnsetTaskNameLike ensures that no value is present for TaskNameLike, not even an explicit nil
### GetTaskDescription

`func (o *HistoricTaskInstanceQueryDto) GetTaskDescription() string`

GetTaskDescription returns the TaskDescription field if non-nil, zero value otherwise.

### GetTaskDescriptionOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskDescriptionOk() (*string, bool)`

GetTaskDescriptionOk returns a tuple with the TaskDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDescription

`func (o *HistoricTaskInstanceQueryDto) SetTaskDescription(v string)`

SetTaskDescription sets TaskDescription field to given value.

### HasTaskDescription

`func (o *HistoricTaskInstanceQueryDto) HasTaskDescription() bool`

HasTaskDescription returns a boolean if a field has been set.

### SetTaskDescriptionNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskDescriptionNil(b bool)`

 SetTaskDescriptionNil sets the value for TaskDescription to be an explicit nil

### UnsetTaskDescription
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskDescription()`

UnsetTaskDescription ensures that no value is present for TaskDescription, not even an explicit nil
### GetTaskDescriptionLike

`func (o *HistoricTaskInstanceQueryDto) GetTaskDescriptionLike() string`

GetTaskDescriptionLike returns the TaskDescriptionLike field if non-nil, zero value otherwise.

### GetTaskDescriptionLikeOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskDescriptionLikeOk() (*string, bool)`

GetTaskDescriptionLikeOk returns a tuple with the TaskDescriptionLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDescriptionLike

`func (o *HistoricTaskInstanceQueryDto) SetTaskDescriptionLike(v string)`

SetTaskDescriptionLike sets TaskDescriptionLike field to given value.

### HasTaskDescriptionLike

`func (o *HistoricTaskInstanceQueryDto) HasTaskDescriptionLike() bool`

HasTaskDescriptionLike returns a boolean if a field has been set.

### SetTaskDescriptionLikeNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskDescriptionLikeNil(b bool)`

 SetTaskDescriptionLikeNil sets the value for TaskDescriptionLike to be an explicit nil

### UnsetTaskDescriptionLike
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskDescriptionLike()`

UnsetTaskDescriptionLike ensures that no value is present for TaskDescriptionLike, not even an explicit nil
### GetTaskDefinitionKey

`func (o *HistoricTaskInstanceQueryDto) GetTaskDefinitionKey() string`

GetTaskDefinitionKey returns the TaskDefinitionKey field if non-nil, zero value otherwise.

### GetTaskDefinitionKeyOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskDefinitionKeyOk() (*string, bool)`

GetTaskDefinitionKeyOk returns a tuple with the TaskDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionKey

`func (o *HistoricTaskInstanceQueryDto) SetTaskDefinitionKey(v string)`

SetTaskDefinitionKey sets TaskDefinitionKey field to given value.

### HasTaskDefinitionKey

`func (o *HistoricTaskInstanceQueryDto) HasTaskDefinitionKey() bool`

HasTaskDefinitionKey returns a boolean if a field has been set.

### SetTaskDefinitionKeyNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskDefinitionKeyNil(b bool)`

 SetTaskDefinitionKeyNil sets the value for TaskDefinitionKey to be an explicit nil

### UnsetTaskDefinitionKey
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskDefinitionKey()`

UnsetTaskDefinitionKey ensures that no value is present for TaskDefinitionKey, not even an explicit nil
### GetTaskDefinitionKeyIn

`func (o *HistoricTaskInstanceQueryDto) GetTaskDefinitionKeyIn() []string`

GetTaskDefinitionKeyIn returns the TaskDefinitionKeyIn field if non-nil, zero value otherwise.

### GetTaskDefinitionKeyInOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskDefinitionKeyInOk() (*[]string, bool)`

GetTaskDefinitionKeyInOk returns a tuple with the TaskDefinitionKeyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionKeyIn

`func (o *HistoricTaskInstanceQueryDto) SetTaskDefinitionKeyIn(v []string)`

SetTaskDefinitionKeyIn sets TaskDefinitionKeyIn field to given value.

### HasTaskDefinitionKeyIn

`func (o *HistoricTaskInstanceQueryDto) HasTaskDefinitionKeyIn() bool`

HasTaskDefinitionKeyIn returns a boolean if a field has been set.

### SetTaskDefinitionKeyInNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskDefinitionKeyInNil(b bool)`

 SetTaskDefinitionKeyInNil sets the value for TaskDefinitionKeyIn to be an explicit nil

### UnsetTaskDefinitionKeyIn
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskDefinitionKeyIn()`

UnsetTaskDefinitionKeyIn ensures that no value is present for TaskDefinitionKeyIn, not even an explicit nil
### GetTaskDeleteReason

`func (o *HistoricTaskInstanceQueryDto) GetTaskDeleteReason() string`

GetTaskDeleteReason returns the TaskDeleteReason field if non-nil, zero value otherwise.

### GetTaskDeleteReasonOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskDeleteReasonOk() (*string, bool)`

GetTaskDeleteReasonOk returns a tuple with the TaskDeleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDeleteReason

`func (o *HistoricTaskInstanceQueryDto) SetTaskDeleteReason(v string)`

SetTaskDeleteReason sets TaskDeleteReason field to given value.

### HasTaskDeleteReason

`func (o *HistoricTaskInstanceQueryDto) HasTaskDeleteReason() bool`

HasTaskDeleteReason returns a boolean if a field has been set.

### SetTaskDeleteReasonNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskDeleteReasonNil(b bool)`

 SetTaskDeleteReasonNil sets the value for TaskDeleteReason to be an explicit nil

### UnsetTaskDeleteReason
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskDeleteReason()`

UnsetTaskDeleteReason ensures that no value is present for TaskDeleteReason, not even an explicit nil
### GetTaskDeleteReasonLike

`func (o *HistoricTaskInstanceQueryDto) GetTaskDeleteReasonLike() string`

GetTaskDeleteReasonLike returns the TaskDeleteReasonLike field if non-nil, zero value otherwise.

### GetTaskDeleteReasonLikeOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskDeleteReasonLikeOk() (*string, bool)`

GetTaskDeleteReasonLikeOk returns a tuple with the TaskDeleteReasonLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDeleteReasonLike

`func (o *HistoricTaskInstanceQueryDto) SetTaskDeleteReasonLike(v string)`

SetTaskDeleteReasonLike sets TaskDeleteReasonLike field to given value.

### HasTaskDeleteReasonLike

`func (o *HistoricTaskInstanceQueryDto) HasTaskDeleteReasonLike() bool`

HasTaskDeleteReasonLike returns a boolean if a field has been set.

### SetTaskDeleteReasonLikeNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskDeleteReasonLikeNil(b bool)`

 SetTaskDeleteReasonLikeNil sets the value for TaskDeleteReasonLike to be an explicit nil

### UnsetTaskDeleteReasonLike
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskDeleteReasonLike()`

UnsetTaskDeleteReasonLike ensures that no value is present for TaskDeleteReasonLike, not even an explicit nil
### GetTaskAssignee

`func (o *HistoricTaskInstanceQueryDto) GetTaskAssignee() string`

GetTaskAssignee returns the TaskAssignee field if non-nil, zero value otherwise.

### GetTaskAssigneeOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskAssigneeOk() (*string, bool)`

GetTaskAssigneeOk returns a tuple with the TaskAssignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAssignee

`func (o *HistoricTaskInstanceQueryDto) SetTaskAssignee(v string)`

SetTaskAssignee sets TaskAssignee field to given value.

### HasTaskAssignee

`func (o *HistoricTaskInstanceQueryDto) HasTaskAssignee() bool`

HasTaskAssignee returns a boolean if a field has been set.

### SetTaskAssigneeNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskAssigneeNil(b bool)`

 SetTaskAssigneeNil sets the value for TaskAssignee to be an explicit nil

### UnsetTaskAssignee
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskAssignee()`

UnsetTaskAssignee ensures that no value is present for TaskAssignee, not even an explicit nil
### GetTaskAssigneeLike

`func (o *HistoricTaskInstanceQueryDto) GetTaskAssigneeLike() string`

GetTaskAssigneeLike returns the TaskAssigneeLike field if non-nil, zero value otherwise.

### GetTaskAssigneeLikeOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskAssigneeLikeOk() (*string, bool)`

GetTaskAssigneeLikeOk returns a tuple with the TaskAssigneeLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAssigneeLike

`func (o *HistoricTaskInstanceQueryDto) SetTaskAssigneeLike(v string)`

SetTaskAssigneeLike sets TaskAssigneeLike field to given value.

### HasTaskAssigneeLike

`func (o *HistoricTaskInstanceQueryDto) HasTaskAssigneeLike() bool`

HasTaskAssigneeLike returns a boolean if a field has been set.

### SetTaskAssigneeLikeNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskAssigneeLikeNil(b bool)`

 SetTaskAssigneeLikeNil sets the value for TaskAssigneeLike to be an explicit nil

### UnsetTaskAssigneeLike
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskAssigneeLike()`

UnsetTaskAssigneeLike ensures that no value is present for TaskAssigneeLike, not even an explicit nil
### GetTaskOwner

`func (o *HistoricTaskInstanceQueryDto) GetTaskOwner() string`

GetTaskOwner returns the TaskOwner field if non-nil, zero value otherwise.

### GetTaskOwnerOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskOwnerOk() (*string, bool)`

GetTaskOwnerOk returns a tuple with the TaskOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOwner

`func (o *HistoricTaskInstanceQueryDto) SetTaskOwner(v string)`

SetTaskOwner sets TaskOwner field to given value.

### HasTaskOwner

`func (o *HistoricTaskInstanceQueryDto) HasTaskOwner() bool`

HasTaskOwner returns a boolean if a field has been set.

### SetTaskOwnerNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskOwnerNil(b bool)`

 SetTaskOwnerNil sets the value for TaskOwner to be an explicit nil

### UnsetTaskOwner
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskOwner()`

UnsetTaskOwner ensures that no value is present for TaskOwner, not even an explicit nil
### GetTaskOwnerLike

`func (o *HistoricTaskInstanceQueryDto) GetTaskOwnerLike() string`

GetTaskOwnerLike returns the TaskOwnerLike field if non-nil, zero value otherwise.

### GetTaskOwnerLikeOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskOwnerLikeOk() (*string, bool)`

GetTaskOwnerLikeOk returns a tuple with the TaskOwnerLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOwnerLike

`func (o *HistoricTaskInstanceQueryDto) SetTaskOwnerLike(v string)`

SetTaskOwnerLike sets TaskOwnerLike field to given value.

### HasTaskOwnerLike

`func (o *HistoricTaskInstanceQueryDto) HasTaskOwnerLike() bool`

HasTaskOwnerLike returns a boolean if a field has been set.

### SetTaskOwnerLikeNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskOwnerLikeNil(b bool)`

 SetTaskOwnerLikeNil sets the value for TaskOwnerLike to be an explicit nil

### UnsetTaskOwnerLike
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskOwnerLike()`

UnsetTaskOwnerLike ensures that no value is present for TaskOwnerLike, not even an explicit nil
### GetTaskPriority

`func (o *HistoricTaskInstanceQueryDto) GetTaskPriority() int32`

GetTaskPriority returns the TaskPriority field if non-nil, zero value otherwise.

### GetTaskPriorityOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskPriorityOk() (*int32, bool)`

GetTaskPriorityOk returns a tuple with the TaskPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPriority

`func (o *HistoricTaskInstanceQueryDto) SetTaskPriority(v int32)`

SetTaskPriority sets TaskPriority field to given value.

### HasTaskPriority

`func (o *HistoricTaskInstanceQueryDto) HasTaskPriority() bool`

HasTaskPriority returns a boolean if a field has been set.

### SetTaskPriorityNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskPriorityNil(b bool)`

 SetTaskPriorityNil sets the value for TaskPriority to be an explicit nil

### UnsetTaskPriority
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskPriority()`

UnsetTaskPriority ensures that no value is present for TaskPriority, not even an explicit nil
### GetAssigned

`func (o *HistoricTaskInstanceQueryDto) GetAssigned() bool`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *HistoricTaskInstanceQueryDto) GetAssignedOk() (*bool, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *HistoricTaskInstanceQueryDto) SetAssigned(v bool)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *HistoricTaskInstanceQueryDto) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### SetAssignedNil

`func (o *HistoricTaskInstanceQueryDto) SetAssignedNil(b bool)`

 SetAssignedNil sets the value for Assigned to be an explicit nil

### UnsetAssigned
`func (o *HistoricTaskInstanceQueryDto) UnsetAssigned()`

UnsetAssigned ensures that no value is present for Assigned, not even an explicit nil
### GetUnassigned

`func (o *HistoricTaskInstanceQueryDto) GetUnassigned() bool`

GetUnassigned returns the Unassigned field if non-nil, zero value otherwise.

### GetUnassignedOk

`func (o *HistoricTaskInstanceQueryDto) GetUnassignedOk() (*bool, bool)`

GetUnassignedOk returns a tuple with the Unassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassigned

`func (o *HistoricTaskInstanceQueryDto) SetUnassigned(v bool)`

SetUnassigned sets Unassigned field to given value.

### HasUnassigned

`func (o *HistoricTaskInstanceQueryDto) HasUnassigned() bool`

HasUnassigned returns a boolean if a field has been set.

### SetUnassignedNil

`func (o *HistoricTaskInstanceQueryDto) SetUnassignedNil(b bool)`

 SetUnassignedNil sets the value for Unassigned to be an explicit nil

### UnsetUnassigned
`func (o *HistoricTaskInstanceQueryDto) UnsetUnassigned()`

UnsetUnassigned ensures that no value is present for Unassigned, not even an explicit nil
### GetFinished

`func (o *HistoricTaskInstanceQueryDto) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *HistoricTaskInstanceQueryDto) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *HistoricTaskInstanceQueryDto) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *HistoricTaskInstanceQueryDto) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *HistoricTaskInstanceQueryDto) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *HistoricTaskInstanceQueryDto) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetUnfinished

`func (o *HistoricTaskInstanceQueryDto) GetUnfinished() bool`

GetUnfinished returns the Unfinished field if non-nil, zero value otherwise.

### GetUnfinishedOk

`func (o *HistoricTaskInstanceQueryDto) GetUnfinishedOk() (*bool, bool)`

GetUnfinishedOk returns a tuple with the Unfinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinished

`func (o *HistoricTaskInstanceQueryDto) SetUnfinished(v bool)`

SetUnfinished sets Unfinished field to given value.

### HasUnfinished

`func (o *HistoricTaskInstanceQueryDto) HasUnfinished() bool`

HasUnfinished returns a boolean if a field has been set.

### SetUnfinishedNil

`func (o *HistoricTaskInstanceQueryDto) SetUnfinishedNil(b bool)`

 SetUnfinishedNil sets the value for Unfinished to be an explicit nil

### UnsetUnfinished
`func (o *HistoricTaskInstanceQueryDto) UnsetUnfinished()`

UnsetUnfinished ensures that no value is present for Unfinished, not even an explicit nil
### GetProcessFinished

`func (o *HistoricTaskInstanceQueryDto) GetProcessFinished() bool`

GetProcessFinished returns the ProcessFinished field if non-nil, zero value otherwise.

### GetProcessFinishedOk

`func (o *HistoricTaskInstanceQueryDto) GetProcessFinishedOk() (*bool, bool)`

GetProcessFinishedOk returns a tuple with the ProcessFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessFinished

`func (o *HistoricTaskInstanceQueryDto) SetProcessFinished(v bool)`

SetProcessFinished sets ProcessFinished field to given value.

### HasProcessFinished

`func (o *HistoricTaskInstanceQueryDto) HasProcessFinished() bool`

HasProcessFinished returns a boolean if a field has been set.

### SetProcessFinishedNil

`func (o *HistoricTaskInstanceQueryDto) SetProcessFinishedNil(b bool)`

 SetProcessFinishedNil sets the value for ProcessFinished to be an explicit nil

### UnsetProcessFinished
`func (o *HistoricTaskInstanceQueryDto) UnsetProcessFinished()`

UnsetProcessFinished ensures that no value is present for ProcessFinished, not even an explicit nil
### GetProcessUnfinished

`func (o *HistoricTaskInstanceQueryDto) GetProcessUnfinished() bool`

GetProcessUnfinished returns the ProcessUnfinished field if non-nil, zero value otherwise.

### GetProcessUnfinishedOk

`func (o *HistoricTaskInstanceQueryDto) GetProcessUnfinishedOk() (*bool, bool)`

GetProcessUnfinishedOk returns a tuple with the ProcessUnfinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessUnfinished

`func (o *HistoricTaskInstanceQueryDto) SetProcessUnfinished(v bool)`

SetProcessUnfinished sets ProcessUnfinished field to given value.

### HasProcessUnfinished

`func (o *HistoricTaskInstanceQueryDto) HasProcessUnfinished() bool`

HasProcessUnfinished returns a boolean if a field has been set.

### SetProcessUnfinishedNil

`func (o *HistoricTaskInstanceQueryDto) SetProcessUnfinishedNil(b bool)`

 SetProcessUnfinishedNil sets the value for ProcessUnfinished to be an explicit nil

### UnsetProcessUnfinished
`func (o *HistoricTaskInstanceQueryDto) UnsetProcessUnfinished()`

UnsetProcessUnfinished ensures that no value is present for ProcessUnfinished, not even an explicit nil
### GetTaskDueDate

`func (o *HistoricTaskInstanceQueryDto) GetTaskDueDate() time.Time`

GetTaskDueDate returns the TaskDueDate field if non-nil, zero value otherwise.

### GetTaskDueDateOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskDueDateOk() (*time.Time, bool)`

GetTaskDueDateOk returns a tuple with the TaskDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDueDate

`func (o *HistoricTaskInstanceQueryDto) SetTaskDueDate(v time.Time)`

SetTaskDueDate sets TaskDueDate field to given value.

### HasTaskDueDate

`func (o *HistoricTaskInstanceQueryDto) HasTaskDueDate() bool`

HasTaskDueDate returns a boolean if a field has been set.

### SetTaskDueDateNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskDueDateNil(b bool)`

 SetTaskDueDateNil sets the value for TaskDueDate to be an explicit nil

### UnsetTaskDueDate
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskDueDate()`

UnsetTaskDueDate ensures that no value is present for TaskDueDate, not even an explicit nil
### GetTaskDueDateBefore

`func (o *HistoricTaskInstanceQueryDto) GetTaskDueDateBefore() time.Time`

GetTaskDueDateBefore returns the TaskDueDateBefore field if non-nil, zero value otherwise.

### GetTaskDueDateBeforeOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskDueDateBeforeOk() (*time.Time, bool)`

GetTaskDueDateBeforeOk returns a tuple with the TaskDueDateBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDueDateBefore

`func (o *HistoricTaskInstanceQueryDto) SetTaskDueDateBefore(v time.Time)`

SetTaskDueDateBefore sets TaskDueDateBefore field to given value.

### HasTaskDueDateBefore

`func (o *HistoricTaskInstanceQueryDto) HasTaskDueDateBefore() bool`

HasTaskDueDateBefore returns a boolean if a field has been set.

### SetTaskDueDateBeforeNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskDueDateBeforeNil(b bool)`

 SetTaskDueDateBeforeNil sets the value for TaskDueDateBefore to be an explicit nil

### UnsetTaskDueDateBefore
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskDueDateBefore()`

UnsetTaskDueDateBefore ensures that no value is present for TaskDueDateBefore, not even an explicit nil
### GetTaskDueDateAfter

`func (o *HistoricTaskInstanceQueryDto) GetTaskDueDateAfter() time.Time`

GetTaskDueDateAfter returns the TaskDueDateAfter field if non-nil, zero value otherwise.

### GetTaskDueDateAfterOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskDueDateAfterOk() (*time.Time, bool)`

GetTaskDueDateAfterOk returns a tuple with the TaskDueDateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDueDateAfter

`func (o *HistoricTaskInstanceQueryDto) SetTaskDueDateAfter(v time.Time)`

SetTaskDueDateAfter sets TaskDueDateAfter field to given value.

### HasTaskDueDateAfter

`func (o *HistoricTaskInstanceQueryDto) HasTaskDueDateAfter() bool`

HasTaskDueDateAfter returns a boolean if a field has been set.

### SetTaskDueDateAfterNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskDueDateAfterNil(b bool)`

 SetTaskDueDateAfterNil sets the value for TaskDueDateAfter to be an explicit nil

### UnsetTaskDueDateAfter
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskDueDateAfter()`

UnsetTaskDueDateAfter ensures that no value is present for TaskDueDateAfter, not even an explicit nil
### GetWithoutTaskDueDate

`func (o *HistoricTaskInstanceQueryDto) GetWithoutTaskDueDate() bool`

GetWithoutTaskDueDate returns the WithoutTaskDueDate field if non-nil, zero value otherwise.

### GetWithoutTaskDueDateOk

`func (o *HistoricTaskInstanceQueryDto) GetWithoutTaskDueDateOk() (*bool, bool)`

GetWithoutTaskDueDateOk returns a tuple with the WithoutTaskDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTaskDueDate

`func (o *HistoricTaskInstanceQueryDto) SetWithoutTaskDueDate(v bool)`

SetWithoutTaskDueDate sets WithoutTaskDueDate field to given value.

### HasWithoutTaskDueDate

`func (o *HistoricTaskInstanceQueryDto) HasWithoutTaskDueDate() bool`

HasWithoutTaskDueDate returns a boolean if a field has been set.

### SetWithoutTaskDueDateNil

`func (o *HistoricTaskInstanceQueryDto) SetWithoutTaskDueDateNil(b bool)`

 SetWithoutTaskDueDateNil sets the value for WithoutTaskDueDate to be an explicit nil

### UnsetWithoutTaskDueDate
`func (o *HistoricTaskInstanceQueryDto) UnsetWithoutTaskDueDate()`

UnsetWithoutTaskDueDate ensures that no value is present for WithoutTaskDueDate, not even an explicit nil
### GetTaskFollowUpDate

`func (o *HistoricTaskInstanceQueryDto) GetTaskFollowUpDate() time.Time`

GetTaskFollowUpDate returns the TaskFollowUpDate field if non-nil, zero value otherwise.

### GetTaskFollowUpDateOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskFollowUpDateOk() (*time.Time, bool)`

GetTaskFollowUpDateOk returns a tuple with the TaskFollowUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskFollowUpDate

`func (o *HistoricTaskInstanceQueryDto) SetTaskFollowUpDate(v time.Time)`

SetTaskFollowUpDate sets TaskFollowUpDate field to given value.

### HasTaskFollowUpDate

`func (o *HistoricTaskInstanceQueryDto) HasTaskFollowUpDate() bool`

HasTaskFollowUpDate returns a boolean if a field has been set.

### SetTaskFollowUpDateNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskFollowUpDateNil(b bool)`

 SetTaskFollowUpDateNil sets the value for TaskFollowUpDate to be an explicit nil

### UnsetTaskFollowUpDate
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskFollowUpDate()`

UnsetTaskFollowUpDate ensures that no value is present for TaskFollowUpDate, not even an explicit nil
### GetTaskFollowUpDateBefore

`func (o *HistoricTaskInstanceQueryDto) GetTaskFollowUpDateBefore() time.Time`

GetTaskFollowUpDateBefore returns the TaskFollowUpDateBefore field if non-nil, zero value otherwise.

### GetTaskFollowUpDateBeforeOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskFollowUpDateBeforeOk() (*time.Time, bool)`

GetTaskFollowUpDateBeforeOk returns a tuple with the TaskFollowUpDateBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskFollowUpDateBefore

`func (o *HistoricTaskInstanceQueryDto) SetTaskFollowUpDateBefore(v time.Time)`

SetTaskFollowUpDateBefore sets TaskFollowUpDateBefore field to given value.

### HasTaskFollowUpDateBefore

`func (o *HistoricTaskInstanceQueryDto) HasTaskFollowUpDateBefore() bool`

HasTaskFollowUpDateBefore returns a boolean if a field has been set.

### SetTaskFollowUpDateBeforeNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskFollowUpDateBeforeNil(b bool)`

 SetTaskFollowUpDateBeforeNil sets the value for TaskFollowUpDateBefore to be an explicit nil

### UnsetTaskFollowUpDateBefore
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskFollowUpDateBefore()`

UnsetTaskFollowUpDateBefore ensures that no value is present for TaskFollowUpDateBefore, not even an explicit nil
### GetTaskFollowUpDateAfter

`func (o *HistoricTaskInstanceQueryDto) GetTaskFollowUpDateAfter() time.Time`

GetTaskFollowUpDateAfter returns the TaskFollowUpDateAfter field if non-nil, zero value otherwise.

### GetTaskFollowUpDateAfterOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskFollowUpDateAfterOk() (*time.Time, bool)`

GetTaskFollowUpDateAfterOk returns a tuple with the TaskFollowUpDateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskFollowUpDateAfter

`func (o *HistoricTaskInstanceQueryDto) SetTaskFollowUpDateAfter(v time.Time)`

SetTaskFollowUpDateAfter sets TaskFollowUpDateAfter field to given value.

### HasTaskFollowUpDateAfter

`func (o *HistoricTaskInstanceQueryDto) HasTaskFollowUpDateAfter() bool`

HasTaskFollowUpDateAfter returns a boolean if a field has been set.

### SetTaskFollowUpDateAfterNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskFollowUpDateAfterNil(b bool)`

 SetTaskFollowUpDateAfterNil sets the value for TaskFollowUpDateAfter to be an explicit nil

### UnsetTaskFollowUpDateAfter
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskFollowUpDateAfter()`

UnsetTaskFollowUpDateAfter ensures that no value is present for TaskFollowUpDateAfter, not even an explicit nil
### GetStartedBefore

`func (o *HistoricTaskInstanceQueryDto) GetStartedBefore() time.Time`

GetStartedBefore returns the StartedBefore field if non-nil, zero value otherwise.

### GetStartedBeforeOk

`func (o *HistoricTaskInstanceQueryDto) GetStartedBeforeOk() (*time.Time, bool)`

GetStartedBeforeOk returns a tuple with the StartedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBefore

`func (o *HistoricTaskInstanceQueryDto) SetStartedBefore(v time.Time)`

SetStartedBefore sets StartedBefore field to given value.

### HasStartedBefore

`func (o *HistoricTaskInstanceQueryDto) HasStartedBefore() bool`

HasStartedBefore returns a boolean if a field has been set.

### SetStartedBeforeNil

`func (o *HistoricTaskInstanceQueryDto) SetStartedBeforeNil(b bool)`

 SetStartedBeforeNil sets the value for StartedBefore to be an explicit nil

### UnsetStartedBefore
`func (o *HistoricTaskInstanceQueryDto) UnsetStartedBefore()`

UnsetStartedBefore ensures that no value is present for StartedBefore, not even an explicit nil
### GetStartedAfter

`func (o *HistoricTaskInstanceQueryDto) GetStartedAfter() time.Time`

GetStartedAfter returns the StartedAfter field if non-nil, zero value otherwise.

### GetStartedAfterOk

`func (o *HistoricTaskInstanceQueryDto) GetStartedAfterOk() (*time.Time, bool)`

GetStartedAfterOk returns a tuple with the StartedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAfter

`func (o *HistoricTaskInstanceQueryDto) SetStartedAfter(v time.Time)`

SetStartedAfter sets StartedAfter field to given value.

### HasStartedAfter

`func (o *HistoricTaskInstanceQueryDto) HasStartedAfter() bool`

HasStartedAfter returns a boolean if a field has been set.

### SetStartedAfterNil

`func (o *HistoricTaskInstanceQueryDto) SetStartedAfterNil(b bool)`

 SetStartedAfterNil sets the value for StartedAfter to be an explicit nil

### UnsetStartedAfter
`func (o *HistoricTaskInstanceQueryDto) UnsetStartedAfter()`

UnsetStartedAfter ensures that no value is present for StartedAfter, not even an explicit nil
### GetFinishedBefore

`func (o *HistoricTaskInstanceQueryDto) GetFinishedBefore() time.Time`

GetFinishedBefore returns the FinishedBefore field if non-nil, zero value otherwise.

### GetFinishedBeforeOk

`func (o *HistoricTaskInstanceQueryDto) GetFinishedBeforeOk() (*time.Time, bool)`

GetFinishedBeforeOk returns a tuple with the FinishedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedBefore

`func (o *HistoricTaskInstanceQueryDto) SetFinishedBefore(v time.Time)`

SetFinishedBefore sets FinishedBefore field to given value.

### HasFinishedBefore

`func (o *HistoricTaskInstanceQueryDto) HasFinishedBefore() bool`

HasFinishedBefore returns a boolean if a field has been set.

### SetFinishedBeforeNil

`func (o *HistoricTaskInstanceQueryDto) SetFinishedBeforeNil(b bool)`

 SetFinishedBeforeNil sets the value for FinishedBefore to be an explicit nil

### UnsetFinishedBefore
`func (o *HistoricTaskInstanceQueryDto) UnsetFinishedBefore()`

UnsetFinishedBefore ensures that no value is present for FinishedBefore, not even an explicit nil
### GetFinishedAfter

`func (o *HistoricTaskInstanceQueryDto) GetFinishedAfter() time.Time`

GetFinishedAfter returns the FinishedAfter field if non-nil, zero value otherwise.

### GetFinishedAfterOk

`func (o *HistoricTaskInstanceQueryDto) GetFinishedAfterOk() (*time.Time, bool)`

GetFinishedAfterOk returns a tuple with the FinishedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAfter

`func (o *HistoricTaskInstanceQueryDto) SetFinishedAfter(v time.Time)`

SetFinishedAfter sets FinishedAfter field to given value.

### HasFinishedAfter

`func (o *HistoricTaskInstanceQueryDto) HasFinishedAfter() bool`

HasFinishedAfter returns a boolean if a field has been set.

### SetFinishedAfterNil

`func (o *HistoricTaskInstanceQueryDto) SetFinishedAfterNil(b bool)`

 SetFinishedAfterNil sets the value for FinishedAfter to be an explicit nil

### UnsetFinishedAfter
`func (o *HistoricTaskInstanceQueryDto) UnsetFinishedAfter()`

UnsetFinishedAfter ensures that no value is present for FinishedAfter, not even an explicit nil
### GetTenantIdIn

`func (o *HistoricTaskInstanceQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *HistoricTaskInstanceQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *HistoricTaskInstanceQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *HistoricTaskInstanceQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *HistoricTaskInstanceQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *HistoricTaskInstanceQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *HistoricTaskInstanceQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *HistoricTaskInstanceQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *HistoricTaskInstanceQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *HistoricTaskInstanceQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *HistoricTaskInstanceQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *HistoricTaskInstanceQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetTaskVariables

`func (o *HistoricTaskInstanceQueryDto) GetTaskVariables() []VariableQueryParameterDto`

GetTaskVariables returns the TaskVariables field if non-nil, zero value otherwise.

### GetTaskVariablesOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskVariablesOk() (*[]VariableQueryParameterDto, bool)`

GetTaskVariablesOk returns a tuple with the TaskVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskVariables

`func (o *HistoricTaskInstanceQueryDto) SetTaskVariables(v []VariableQueryParameterDto)`

SetTaskVariables sets TaskVariables field to given value.

### HasTaskVariables

`func (o *HistoricTaskInstanceQueryDto) HasTaskVariables() bool`

HasTaskVariables returns a boolean if a field has been set.

### SetTaskVariablesNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskVariablesNil(b bool)`

 SetTaskVariablesNil sets the value for TaskVariables to be an explicit nil

### UnsetTaskVariables
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskVariables()`

UnsetTaskVariables ensures that no value is present for TaskVariables, not even an explicit nil
### GetProcessVariables

`func (o *HistoricTaskInstanceQueryDto) GetProcessVariables() []VariableQueryParameterDto`

GetProcessVariables returns the ProcessVariables field if non-nil, zero value otherwise.

### GetProcessVariablesOk

`func (o *HistoricTaskInstanceQueryDto) GetProcessVariablesOk() (*[]VariableQueryParameterDto, bool)`

GetProcessVariablesOk returns a tuple with the ProcessVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessVariables

`func (o *HistoricTaskInstanceQueryDto) SetProcessVariables(v []VariableQueryParameterDto)`

SetProcessVariables sets ProcessVariables field to given value.

### HasProcessVariables

`func (o *HistoricTaskInstanceQueryDto) HasProcessVariables() bool`

HasProcessVariables returns a boolean if a field has been set.

### SetProcessVariablesNil

`func (o *HistoricTaskInstanceQueryDto) SetProcessVariablesNil(b bool)`

 SetProcessVariablesNil sets the value for ProcessVariables to be an explicit nil

### UnsetProcessVariables
`func (o *HistoricTaskInstanceQueryDto) UnsetProcessVariables()`

UnsetProcessVariables ensures that no value is present for ProcessVariables, not even an explicit nil
### GetVariableNamesIgnoreCase

`func (o *HistoricTaskInstanceQueryDto) GetVariableNamesIgnoreCase() bool`

GetVariableNamesIgnoreCase returns the VariableNamesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableNamesIgnoreCaseOk

`func (o *HistoricTaskInstanceQueryDto) GetVariableNamesIgnoreCaseOk() (*bool, bool)`

GetVariableNamesIgnoreCaseOk returns a tuple with the VariableNamesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNamesIgnoreCase

`func (o *HistoricTaskInstanceQueryDto) SetVariableNamesIgnoreCase(v bool)`

SetVariableNamesIgnoreCase sets VariableNamesIgnoreCase field to given value.

### HasVariableNamesIgnoreCase

`func (o *HistoricTaskInstanceQueryDto) HasVariableNamesIgnoreCase() bool`

HasVariableNamesIgnoreCase returns a boolean if a field has been set.

### SetVariableNamesIgnoreCaseNil

`func (o *HistoricTaskInstanceQueryDto) SetVariableNamesIgnoreCaseNil(b bool)`

 SetVariableNamesIgnoreCaseNil sets the value for VariableNamesIgnoreCase to be an explicit nil

### UnsetVariableNamesIgnoreCase
`func (o *HistoricTaskInstanceQueryDto) UnsetVariableNamesIgnoreCase()`

UnsetVariableNamesIgnoreCase ensures that no value is present for VariableNamesIgnoreCase, not even an explicit nil
### GetVariableValuesIgnoreCase

`func (o *HistoricTaskInstanceQueryDto) GetVariableValuesIgnoreCase() bool`

GetVariableValuesIgnoreCase returns the VariableValuesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableValuesIgnoreCaseOk

`func (o *HistoricTaskInstanceQueryDto) GetVariableValuesIgnoreCaseOk() (*bool, bool)`

GetVariableValuesIgnoreCaseOk returns a tuple with the VariableValuesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValuesIgnoreCase

`func (o *HistoricTaskInstanceQueryDto) SetVariableValuesIgnoreCase(v bool)`

SetVariableValuesIgnoreCase sets VariableValuesIgnoreCase field to given value.

### HasVariableValuesIgnoreCase

`func (o *HistoricTaskInstanceQueryDto) HasVariableValuesIgnoreCase() bool`

HasVariableValuesIgnoreCase returns a boolean if a field has been set.

### SetVariableValuesIgnoreCaseNil

`func (o *HistoricTaskInstanceQueryDto) SetVariableValuesIgnoreCaseNil(b bool)`

 SetVariableValuesIgnoreCaseNil sets the value for VariableValuesIgnoreCase to be an explicit nil

### UnsetVariableValuesIgnoreCase
`func (o *HistoricTaskInstanceQueryDto) UnsetVariableValuesIgnoreCase()`

UnsetVariableValuesIgnoreCase ensures that no value is present for VariableValuesIgnoreCase, not even an explicit nil
### GetTaskInvolvedUser

`func (o *HistoricTaskInstanceQueryDto) GetTaskInvolvedUser() string`

GetTaskInvolvedUser returns the TaskInvolvedUser field if non-nil, zero value otherwise.

### GetTaskInvolvedUserOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskInvolvedUserOk() (*string, bool)`

GetTaskInvolvedUserOk returns a tuple with the TaskInvolvedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInvolvedUser

`func (o *HistoricTaskInstanceQueryDto) SetTaskInvolvedUser(v string)`

SetTaskInvolvedUser sets TaskInvolvedUser field to given value.

### HasTaskInvolvedUser

`func (o *HistoricTaskInstanceQueryDto) HasTaskInvolvedUser() bool`

HasTaskInvolvedUser returns a boolean if a field has been set.

### SetTaskInvolvedUserNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskInvolvedUserNil(b bool)`

 SetTaskInvolvedUserNil sets the value for TaskInvolvedUser to be an explicit nil

### UnsetTaskInvolvedUser
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskInvolvedUser()`

UnsetTaskInvolvedUser ensures that no value is present for TaskInvolvedUser, not even an explicit nil
### GetTaskInvolvedGroup

`func (o *HistoricTaskInstanceQueryDto) GetTaskInvolvedGroup() string`

GetTaskInvolvedGroup returns the TaskInvolvedGroup field if non-nil, zero value otherwise.

### GetTaskInvolvedGroupOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskInvolvedGroupOk() (*string, bool)`

GetTaskInvolvedGroupOk returns a tuple with the TaskInvolvedGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInvolvedGroup

`func (o *HistoricTaskInstanceQueryDto) SetTaskInvolvedGroup(v string)`

SetTaskInvolvedGroup sets TaskInvolvedGroup field to given value.

### HasTaskInvolvedGroup

`func (o *HistoricTaskInstanceQueryDto) HasTaskInvolvedGroup() bool`

HasTaskInvolvedGroup returns a boolean if a field has been set.

### SetTaskInvolvedGroupNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskInvolvedGroupNil(b bool)`

 SetTaskInvolvedGroupNil sets the value for TaskInvolvedGroup to be an explicit nil

### UnsetTaskInvolvedGroup
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskInvolvedGroup()`

UnsetTaskInvolvedGroup ensures that no value is present for TaskInvolvedGroup, not even an explicit nil
### GetTaskHadCandidateUser

`func (o *HistoricTaskInstanceQueryDto) GetTaskHadCandidateUser() string`

GetTaskHadCandidateUser returns the TaskHadCandidateUser field if non-nil, zero value otherwise.

### GetTaskHadCandidateUserOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskHadCandidateUserOk() (*string, bool)`

GetTaskHadCandidateUserOk returns a tuple with the TaskHadCandidateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskHadCandidateUser

`func (o *HistoricTaskInstanceQueryDto) SetTaskHadCandidateUser(v string)`

SetTaskHadCandidateUser sets TaskHadCandidateUser field to given value.

### HasTaskHadCandidateUser

`func (o *HistoricTaskInstanceQueryDto) HasTaskHadCandidateUser() bool`

HasTaskHadCandidateUser returns a boolean if a field has been set.

### SetTaskHadCandidateUserNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskHadCandidateUserNil(b bool)`

 SetTaskHadCandidateUserNil sets the value for TaskHadCandidateUser to be an explicit nil

### UnsetTaskHadCandidateUser
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskHadCandidateUser()`

UnsetTaskHadCandidateUser ensures that no value is present for TaskHadCandidateUser, not even an explicit nil
### GetTaskHadCandidateGroup

`func (o *HistoricTaskInstanceQueryDto) GetTaskHadCandidateGroup() string`

GetTaskHadCandidateGroup returns the TaskHadCandidateGroup field if non-nil, zero value otherwise.

### GetTaskHadCandidateGroupOk

`func (o *HistoricTaskInstanceQueryDto) GetTaskHadCandidateGroupOk() (*string, bool)`

GetTaskHadCandidateGroupOk returns a tuple with the TaskHadCandidateGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskHadCandidateGroup

`func (o *HistoricTaskInstanceQueryDto) SetTaskHadCandidateGroup(v string)`

SetTaskHadCandidateGroup sets TaskHadCandidateGroup field to given value.

### HasTaskHadCandidateGroup

`func (o *HistoricTaskInstanceQueryDto) HasTaskHadCandidateGroup() bool`

HasTaskHadCandidateGroup returns a boolean if a field has been set.

### SetTaskHadCandidateGroupNil

`func (o *HistoricTaskInstanceQueryDto) SetTaskHadCandidateGroupNil(b bool)`

 SetTaskHadCandidateGroupNil sets the value for TaskHadCandidateGroup to be an explicit nil

### UnsetTaskHadCandidateGroup
`func (o *HistoricTaskInstanceQueryDto) UnsetTaskHadCandidateGroup()`

UnsetTaskHadCandidateGroup ensures that no value is present for TaskHadCandidateGroup, not even an explicit nil
### GetWithCandidateGroups

`func (o *HistoricTaskInstanceQueryDto) GetWithCandidateGroups() bool`

GetWithCandidateGroups returns the WithCandidateGroups field if non-nil, zero value otherwise.

### GetWithCandidateGroupsOk

`func (o *HistoricTaskInstanceQueryDto) GetWithCandidateGroupsOk() (*bool, bool)`

GetWithCandidateGroupsOk returns a tuple with the WithCandidateGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCandidateGroups

`func (o *HistoricTaskInstanceQueryDto) SetWithCandidateGroups(v bool)`

SetWithCandidateGroups sets WithCandidateGroups field to given value.

### HasWithCandidateGroups

`func (o *HistoricTaskInstanceQueryDto) HasWithCandidateGroups() bool`

HasWithCandidateGroups returns a boolean if a field has been set.

### SetWithCandidateGroupsNil

`func (o *HistoricTaskInstanceQueryDto) SetWithCandidateGroupsNil(b bool)`

 SetWithCandidateGroupsNil sets the value for WithCandidateGroups to be an explicit nil

### UnsetWithCandidateGroups
`func (o *HistoricTaskInstanceQueryDto) UnsetWithCandidateGroups()`

UnsetWithCandidateGroups ensures that no value is present for WithCandidateGroups, not even an explicit nil
### GetWithoutCandidateGroups

`func (o *HistoricTaskInstanceQueryDto) GetWithoutCandidateGroups() bool`

GetWithoutCandidateGroups returns the WithoutCandidateGroups field if non-nil, zero value otherwise.

### GetWithoutCandidateGroupsOk

`func (o *HistoricTaskInstanceQueryDto) GetWithoutCandidateGroupsOk() (*bool, bool)`

GetWithoutCandidateGroupsOk returns a tuple with the WithoutCandidateGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutCandidateGroups

`func (o *HistoricTaskInstanceQueryDto) SetWithoutCandidateGroups(v bool)`

SetWithoutCandidateGroups sets WithoutCandidateGroups field to given value.

### HasWithoutCandidateGroups

`func (o *HistoricTaskInstanceQueryDto) HasWithoutCandidateGroups() bool`

HasWithoutCandidateGroups returns a boolean if a field has been set.

### SetWithoutCandidateGroupsNil

`func (o *HistoricTaskInstanceQueryDto) SetWithoutCandidateGroupsNil(b bool)`

 SetWithoutCandidateGroupsNil sets the value for WithoutCandidateGroups to be an explicit nil

### UnsetWithoutCandidateGroups
`func (o *HistoricTaskInstanceQueryDto) UnsetWithoutCandidateGroups()`

UnsetWithoutCandidateGroups ensures that no value is present for WithoutCandidateGroups, not even an explicit nil
### GetOrQueries

`func (o *HistoricTaskInstanceQueryDto) GetOrQueries() []HistoricTaskInstanceQueryDto`

GetOrQueries returns the OrQueries field if non-nil, zero value otherwise.

### GetOrQueriesOk

`func (o *HistoricTaskInstanceQueryDto) GetOrQueriesOk() (*[]HistoricTaskInstanceQueryDto, bool)`

GetOrQueriesOk returns a tuple with the OrQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrQueries

`func (o *HistoricTaskInstanceQueryDto) SetOrQueries(v []HistoricTaskInstanceQueryDto)`

SetOrQueries sets OrQueries field to given value.

### HasOrQueries

`func (o *HistoricTaskInstanceQueryDto) HasOrQueries() bool`

HasOrQueries returns a boolean if a field has been set.

### SetOrQueriesNil

`func (o *HistoricTaskInstanceQueryDto) SetOrQueriesNil(b bool)`

 SetOrQueriesNil sets the value for OrQueries to be an explicit nil

### UnsetOrQueries
`func (o *HistoricTaskInstanceQueryDto) UnsetOrQueries()`

UnsetOrQueries ensures that no value is present for OrQueries, not even an explicit nil
### GetSorting

`func (o *HistoricTaskInstanceQueryDto) GetSorting() []HistoricTaskInstanceQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *HistoricTaskInstanceQueryDto) GetSortingOk() (*[]HistoricTaskInstanceQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *HistoricTaskInstanceQueryDto) SetSorting(v []HistoricTaskInstanceQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *HistoricTaskInstanceQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *HistoricTaskInstanceQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *HistoricTaskInstanceQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


