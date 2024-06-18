# TaskQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **NullableString** | Restrict to task with the given id. | [optional] 
**TaskIdIn** | Pointer to **[]string** | Restrict to tasks with any of the given ids. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | Restrict to tasks that belong to process instances with the given id. | [optional] 
**ProcessInstanceIdIn** | Pointer to **[]string** | Restrict to tasks that belong to process instances with the given ids. | [optional] 
**ProcessInstanceBusinessKey** | Pointer to **NullableString** | Restrict to tasks that belong to process instances with the given business key. | [optional] 
**ProcessInstanceBusinessKeyExpression** | Pointer to **NullableString** | Restrict to tasks that belong to process instances with the given business key which  is described by an expression. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. | [optional] 
**ProcessInstanceBusinessKeyIn** | Pointer to **[]string** | Restrict to tasks that belong to process instances with one of the give business keys.  The keys need to be in a comma-separated list. | [optional] 
**ProcessInstanceBusinessKeyLike** | Pointer to **NullableString** | Restrict to tasks that have a process instance business key that has the parameter  value as a substring. | [optional] 
**ProcessInstanceBusinessKeyLikeExpression** | Pointer to **NullableString** | Restrict to tasks that have a process instance business key that has the parameter  value as a substring and is described by an expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Restrict to tasks that belong to a process definition with the given id. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Restrict to tasks that belong to a process definition with the given key. | [optional] 
**ProcessDefinitionKeyIn** | Pointer to **[]string** | Restrict to tasks that belong to a process definition with one of the given keys. The  keys need to be in a comma-separated list. | [optional] 
**ProcessDefinitionName** | Pointer to **NullableString** | Restrict to tasks that belong to a process definition with the given name. | [optional] 
**ProcessDefinitionNameLike** | Pointer to **NullableString** | Restrict to tasks that have a process definition name that has the parameter value as  a substring. | [optional] 
**ExecutionId** | Pointer to **NullableString** | Restrict to tasks that belong to an execution with the given id. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | Restrict to tasks that belong to case instances with the given id. | [optional] 
**CaseInstanceBusinessKey** | Pointer to **NullableString** | Restrict to tasks that belong to case instances with the given business key. | [optional] 
**CaseInstanceBusinessKeyLike** | Pointer to **NullableString** | Restrict to tasks that have a case instance business key that has the parameter value  as a substring. | [optional] 
**CaseDefinitionId** | Pointer to **NullableString** | Restrict to tasks that belong to a case definition with the given id. | [optional] 
**CaseDefinitionKey** | Pointer to **NullableString** | Restrict to tasks that belong to a case definition with the given key. | [optional] 
**CaseDefinitionName** | Pointer to **NullableString** | Restrict to tasks that belong to a case definition with the given name. | [optional] 
**CaseDefinitionNameLike** | Pointer to **NullableString** | Restrict to tasks that have a case definition name that has the parameter value as a  substring. | [optional] 
**CaseExecutionId** | Pointer to **NullableString** | Restrict to tasks that belong to a case execution with the given id. | [optional] 
**ActivityInstanceIdIn** | Pointer to **[]string** | Only include tasks which belong to one of the passed and comma-separated activity  instance ids. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Only include tasks which belong to one of the passed and comma-separated  tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include tasks which belong to no tenant. Value may only be &#x60;true&#x60;,  as &#x60;false&#x60; is the default behavior. | [optional] [default to false]
**Assignee** | Pointer to **NullableString** | Restrict to tasks that the given user is assigned to. | [optional] 
**AssigneeExpression** | Pointer to **NullableString** | Restrict to tasks that the user described by the given expression is assigned to. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | [optional] 
**AssigneeLike** | Pointer to **NullableString** | Restrict to tasks that have an assignee that has the parameter  value as a substring. | [optional] 
**AssigneeLikeExpression** | Pointer to **NullableString** | Restrict to tasks that have an assignee that has the parameter value described by the  given expression as a substring. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | [optional] 
**AssigneeIn** | Pointer to **[]string** | Only include tasks which are assigned to one of the passed and comma-separated user ids. | [optional] 
**AssigneeNotIn** | Pointer to **[]string** | Only include tasks which are not assigned to one of the passed and comma-separated user ids. | [optional] 
**Owner** | Pointer to **NullableString** | Restrict to tasks that the given user owns. | [optional] 
**OwnerExpression** | Pointer to **NullableString** | Restrict to tasks that the user described by the given expression owns. See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | [optional] 
**CandidateGroup** | Pointer to **NullableString** | Only include tasks that are offered to the given group. | [optional] 
**CandidateGroupExpression** | Pointer to **NullableString** | Only include tasks that are offered to the group described by the given expression.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | [optional] 
**CandidateUser** | Pointer to **NullableString** | Only include tasks that are offered to the given user or to one of his groups. | [optional] 
**CandidateUserExpression** | Pointer to **NullableString** | Only include tasks that are offered to the user described by the given expression.  See the  [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions)  for more information on available functions. | [optional] 
**IncludeAssignedTasks** | Pointer to **NullableBool** | Also include tasks that are assigned to users in candidate queries. Default is to only  include tasks that are not assigned to any user if you query by candidate user or group(s). | [optional] [default to false]
**InvolvedUser** | Pointer to **NullableString** | Only include tasks that the given user is involved in. A user is involved in a task if  an identity link exists between task and user (e.g., the user is the assignee). | [optional] 
**InvolvedUserExpression** | Pointer to **NullableString** | Only include tasks that the user described by the given expression is involved in. A user is involved in a task if an identity link exists between task and user (e.g., the user is the assignee). See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. | [optional] 
**Assigned** | Pointer to **NullableBool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are assigned. | [optional] [default to false]
**Unassigned** | Pointer to **NullableBool** | If set to &#x60;true&#x60;, restricts the query to all tasks that are unassigned. | [optional] [default to false]
**TaskDefinitionKey** | Pointer to **NullableString** | Restrict to tasks that have the given key. | [optional] 
**TaskDefinitionKeyIn** | Pointer to **[]string** | Restrict to tasks that have one of the given keys. The keys need to be in a comma-separated list. | [optional] 
**TaskDefinitionKeyLike** | Pointer to **NullableString** | Restrict to tasks that have a key that has the parameter value as a substring. | [optional] 
**Name** | Pointer to **NullableString** | Restrict to tasks that have the given name. | [optional] 
**NameNotEqual** | Pointer to **NullableString** | Restrict to tasks that do not have the given name. | [optional] 
**NameLike** | Pointer to **NullableString** | Restrict to tasks that have a name with the given parameter value as substring. | [optional] 
**NameNotLike** | Pointer to **NullableString** | Restrict to tasks that do not have a name with the given parameter value as substring. | [optional] 
**Description** | Pointer to **NullableString** | Restrict to tasks that have the given description. | [optional] 
**DescriptionLike** | Pointer to **NullableString** | Restrict to tasks that have a description that has the parameter value as a substring. | [optional] 
**Priority** | Pointer to **NullableInt32** | Restrict to tasks that have the given priority. | [optional] 
**MaxPriority** | Pointer to **NullableInt32** | Restrict to tasks that have a lower or equal priority. | [optional] 
**MinPriority** | Pointer to **NullableInt32** | Restrict to tasks that have a higher or equal priority. | [optional] 
**DueDate** | Pointer to **NullableTime** | Restrict to tasks that are due on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.546+0200&#x60;. | [optional] 
**DueDateExpression** | Pointer to **NullableString** | Restrict to tasks that are due on the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | [optional] 
**DueAfter** | Pointer to **NullableTime** | Restrict to tasks that are due after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.435+0200&#x60;. | [optional] 
**DueAfterExpression** | Pointer to **NullableString** | Restrict to tasks that are due after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | [optional] 
**DueBefore** | Pointer to **NullableTime** | Restrict to tasks that are due before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.243+0200&#x60;. | [optional] 
**DueBeforeExpression** | Pointer to **NullableString** | Restrict to tasks that are due before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | [optional] 
**WithoutDueDate** | Pointer to **NullableBool** | Only include tasks which have no due date. Value may only be &#x60;true&#x60;,  as &#x60;false&#x60; is the default behavior. | [optional] [default to false]
**FollowUpDate** | Pointer to **NullableTime** | Restrict to tasks that have a followUp date on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.342+0200&#x60;. | [optional] 
**FollowUpDateExpression** | Pointer to **NullableString** | Restrict to tasks that have a followUp date on the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | [optional] 
**FollowUpAfter** | Pointer to **NullableTime** | Restrict to tasks that have a followUp date after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.542+0200&#x60;. | [optional] 
**FollowUpAfterExpression** | Pointer to **NullableString** | Restrict to tasks that have a followUp date after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | [optional] 
**FollowUpBefore** | Pointer to **NullableString** | Restrict to tasks that have a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.234+0200&#x60;. | [optional] 
**FollowUpBeforeExpression** | Pointer to **NullableString** | Restrict to tasks that have a followUp date before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | [optional] 
**FollowUpBeforeOrNotExistent** | Pointer to **NullableTime** | Restrict to tasks that have no followUp date or a followUp date before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.432+0200&#x60;. The typical use case is to query all &#x60;active&#x60; tasks for a user for a given date. | [optional] 
**FollowUpBeforeOrNotExistentExpression** | Pointer to **NullableString** | Restrict to tasks that have no followUp date or a followUp date before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | [optional] 
**CreatedOn** | Pointer to **NullableTime** | Restrict to tasks that were created on the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.324+0200&#x60;. | [optional] 
**CreatedOnExpression** | Pointer to **NullableString** | Restrict to tasks that were created on the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | [optional] 
**CreatedAfter** | Pointer to **NullableTime** | Restrict to tasks that were created after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.342+0200&#x60;. | [optional] 
**CreatedAfterExpression** | Pointer to **NullableString** | Restrict to tasks that were created after the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | [optional] 
**CreatedBefore** | Pointer to **NullableTime** | Restrict to tasks that were created before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.332+0200&#x60;. | [optional] 
**CreatedBeforeExpression** | Pointer to **NullableString** | Restrict to tasks that were created before the date described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | [optional] 
**UpdatedAfter** | Pointer to **NullableTime** | Restrict to tasks that were updated after the given date. Every action that fires  a [task update event](https://docs.camunda.org/manual/7.21/user-guide/process-engine/delegation-code/#task-listener-event-lifecycle) is considered as updating the task. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.332+0200&#x60;. | [optional] 
**UpdatedAfterExpression** | Pointer to **NullableString** | Restrict to tasks that were updated after the date described by the given expression. Every action that fires  a [task update event](https://docs.camunda.org/manual/7.21/user-guide/process-engine/delegation-code/#task-listener-event-lifecycle) is considered as updating the task. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to a &#x60;java.util.Date&#x60; or &#x60;org.joda.time.DateTime&#x60; object. | [optional] 
**DelegationState** | Pointer to **NullableString** | Restrict to tasks that are in the given delegation state. Valid values are &#x60;PENDING&#x60; and &#x60;RESOLVED&#x60;. | [optional] 
**CandidateGroups** | Pointer to **[]string** | Restrict to tasks that are offered to any of the given candidate groups. Takes a comma-separated list of group names, so for example &#x60;developers,support,sales&#x60;. | [optional] 
**CandidateGroupsExpression** | Pointer to **NullableString** | Restrict to tasks that are offered to any of the candidate groups described by the given expression. See the [user guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/expression-language/#internal-context-functions) for more information on available functions. The expression must evaluate to &#x60;java.util.List&#x60; of Strings. | [optional] 
**WithCandidateGroups** | Pointer to **NullableBool** | Only include tasks which have a candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] [default to false]
**WithoutCandidateGroups** | Pointer to **NullableBool** | Only include tasks which have no candidate group. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] [default to false]
**WithCandidateUsers** | Pointer to **NullableBool** | Only include tasks which have a candidate user. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] [default to false]
**WithoutCandidateUsers** | Pointer to **NullableBool** | Only include tasks which have no candidate users. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] [default to false]
**Active** | Pointer to **NullableBool** | Only include active tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] [default to false]
**Suspended** | Pointer to **NullableBool** | Only include suspended tasks. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] [default to false]
**TaskVariables** | Pointer to [**[]VariableQueryParameterDto**](VariableQueryParameterDto.md) | A JSON array to only include tasks that have variables with certain values. The array consists of JSON objects with three properties &#x60;name&#x60;, &#x60;operator&#x60; and &#x60;value&#x60;. &#x60;name&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. &#x60;value&#x60; may be of type &#x60;String&#x60;, &#x60;Number&#x60; or &#x60;Boolean&#x60;.  Valid &#x60;operator&#x60; values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | [optional] 
**ProcessVariables** | Pointer to [**[]VariableQueryParameterDto**](VariableQueryParameterDto.md) | A JSON array to only include tasks that belong to a process instance with variables with certain values. The array consists of JSON objects with three properties &#x60;name&#x60;, &#x60;operator&#x60; and &#x60;value&#x60;. &#x60;name&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. &#x60;value&#x60; may be of type &#x60;String&#x60;, &#x60;Number&#x60; or &#x60;Boolean&#x60;.  Valid &#x60;operator&#x60; values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;; &#x60;notLike&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | [optional] 
**CaseInstanceVariables** | Pointer to [**[]VariableQueryParameterDto**](VariableQueryParameterDto.md) | A JSON array to only include tasks that belong to a case instance with variables with certain values. The array consists of JSON objects with three properties &#x60;name&#x60;, &#x60;operator&#x60; and &#x60;value&#x60;. &#x60;name&#x60; is the variable name, &#x60;operator&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. &#x60;value&#x60; may be of type &#x60;String&#x60;, &#x60;Number&#x60; or &#x60;Boolean&#x60;.  Valid &#x60;operator&#x60; values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. &#x60;key&#x60; and &#x60;value&#x60; may not contain underscore or comma characters. | [optional] 
**VariableNamesIgnoreCase** | Pointer to **NullableBool** | Match all variable names in this query case-insensitively. If set &#x60;variableName&#x60; and &#x60;variablename&#x60; are treated as equal. | [optional] [default to false]
**VariableValuesIgnoreCase** | Pointer to **NullableBool** | Match all variable values in this query case-insensitively. If set &#x60;variableValue&#x60; and &#x60;variablevalue&#x60; are treated as equal. | [optional] [default to false]
**ParentTaskId** | Pointer to **NullableString** | Restrict query to all tasks that are sub tasks of the given task. Takes a task id. | [optional] 
**OrQueries** | Pointer to [**[]TaskQueryDto**](TaskQueryDto.md) | A JSON array of nested task queries with OR semantics. A task matches a nested query if it fulfills *at least one* of the query&#39;s predicates. With multiple nested queries, a task must fulfill at least one predicate of *each* query ([Conjunctive Normal Form](https://en.wikipedia.org/wiki/Conjunctive_normal_form)).  All task query properties can be used except for: &#x60;sorting&#x60;, &#x60;withCandidateGroups&#x60;, &#x60;withoutCandidateGroups&#x60;, &#x60;withCandidateUsers&#x60;, &#x60;withoutCandidateUsers&#x60;  See the [User guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/process-engine-api/#or-queries) for more information about OR queries. | [optional] 
**Sorting** | Pointer to [**[]TaskQueryDtoSortingInner**](TaskQueryDtoSortingInner.md) | Apply sorting of the result | [optional] 

## Methods

### NewTaskQueryDto

`func NewTaskQueryDto() *TaskQueryDto`

NewTaskQueryDto instantiates a new TaskQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskQueryDtoWithDefaults

`func NewTaskQueryDtoWithDefaults() *TaskQueryDto`

NewTaskQueryDtoWithDefaults instantiates a new TaskQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *TaskQueryDto) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskQueryDto) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskQueryDto) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TaskQueryDto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *TaskQueryDto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *TaskQueryDto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetTaskIdIn

`func (o *TaskQueryDto) GetTaskIdIn() []string`

GetTaskIdIn returns the TaskIdIn field if non-nil, zero value otherwise.

### GetTaskIdInOk

`func (o *TaskQueryDto) GetTaskIdInOk() (*[]string, bool)`

GetTaskIdInOk returns a tuple with the TaskIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIdIn

`func (o *TaskQueryDto) SetTaskIdIn(v []string)`

SetTaskIdIn sets TaskIdIn field to given value.

### HasTaskIdIn

`func (o *TaskQueryDto) HasTaskIdIn() bool`

HasTaskIdIn returns a boolean if a field has been set.

### SetTaskIdInNil

`func (o *TaskQueryDto) SetTaskIdInNil(b bool)`

 SetTaskIdInNil sets the value for TaskIdIn to be an explicit nil

### UnsetTaskIdIn
`func (o *TaskQueryDto) UnsetTaskIdIn()`

UnsetTaskIdIn ensures that no value is present for TaskIdIn, not even an explicit nil
### GetProcessInstanceId

`func (o *TaskQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *TaskQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *TaskQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *TaskQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *TaskQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *TaskQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessInstanceIdIn

`func (o *TaskQueryDto) GetProcessInstanceIdIn() []string`

GetProcessInstanceIdIn returns the ProcessInstanceIdIn field if non-nil, zero value otherwise.

### GetProcessInstanceIdInOk

`func (o *TaskQueryDto) GetProcessInstanceIdInOk() (*[]string, bool)`

GetProcessInstanceIdInOk returns a tuple with the ProcessInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIdIn

`func (o *TaskQueryDto) SetProcessInstanceIdIn(v []string)`

SetProcessInstanceIdIn sets ProcessInstanceIdIn field to given value.

### HasProcessInstanceIdIn

`func (o *TaskQueryDto) HasProcessInstanceIdIn() bool`

HasProcessInstanceIdIn returns a boolean if a field has been set.

### SetProcessInstanceIdInNil

`func (o *TaskQueryDto) SetProcessInstanceIdInNil(b bool)`

 SetProcessInstanceIdInNil sets the value for ProcessInstanceIdIn to be an explicit nil

### UnsetProcessInstanceIdIn
`func (o *TaskQueryDto) UnsetProcessInstanceIdIn()`

UnsetProcessInstanceIdIn ensures that no value is present for ProcessInstanceIdIn, not even an explicit nil
### GetProcessInstanceBusinessKey

`func (o *TaskQueryDto) GetProcessInstanceBusinessKey() string`

GetProcessInstanceBusinessKey returns the ProcessInstanceBusinessKey field if non-nil, zero value otherwise.

### GetProcessInstanceBusinessKeyOk

`func (o *TaskQueryDto) GetProcessInstanceBusinessKeyOk() (*string, bool)`

GetProcessInstanceBusinessKeyOk returns a tuple with the ProcessInstanceBusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceBusinessKey

`func (o *TaskQueryDto) SetProcessInstanceBusinessKey(v string)`

SetProcessInstanceBusinessKey sets ProcessInstanceBusinessKey field to given value.

### HasProcessInstanceBusinessKey

`func (o *TaskQueryDto) HasProcessInstanceBusinessKey() bool`

HasProcessInstanceBusinessKey returns a boolean if a field has been set.

### SetProcessInstanceBusinessKeyNil

`func (o *TaskQueryDto) SetProcessInstanceBusinessKeyNil(b bool)`

 SetProcessInstanceBusinessKeyNil sets the value for ProcessInstanceBusinessKey to be an explicit nil

### UnsetProcessInstanceBusinessKey
`func (o *TaskQueryDto) UnsetProcessInstanceBusinessKey()`

UnsetProcessInstanceBusinessKey ensures that no value is present for ProcessInstanceBusinessKey, not even an explicit nil
### GetProcessInstanceBusinessKeyExpression

`func (o *TaskQueryDto) GetProcessInstanceBusinessKeyExpression() string`

GetProcessInstanceBusinessKeyExpression returns the ProcessInstanceBusinessKeyExpression field if non-nil, zero value otherwise.

### GetProcessInstanceBusinessKeyExpressionOk

`func (o *TaskQueryDto) GetProcessInstanceBusinessKeyExpressionOk() (*string, bool)`

GetProcessInstanceBusinessKeyExpressionOk returns a tuple with the ProcessInstanceBusinessKeyExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceBusinessKeyExpression

`func (o *TaskQueryDto) SetProcessInstanceBusinessKeyExpression(v string)`

SetProcessInstanceBusinessKeyExpression sets ProcessInstanceBusinessKeyExpression field to given value.

### HasProcessInstanceBusinessKeyExpression

`func (o *TaskQueryDto) HasProcessInstanceBusinessKeyExpression() bool`

HasProcessInstanceBusinessKeyExpression returns a boolean if a field has been set.

### SetProcessInstanceBusinessKeyExpressionNil

`func (o *TaskQueryDto) SetProcessInstanceBusinessKeyExpressionNil(b bool)`

 SetProcessInstanceBusinessKeyExpressionNil sets the value for ProcessInstanceBusinessKeyExpression to be an explicit nil

### UnsetProcessInstanceBusinessKeyExpression
`func (o *TaskQueryDto) UnsetProcessInstanceBusinessKeyExpression()`

UnsetProcessInstanceBusinessKeyExpression ensures that no value is present for ProcessInstanceBusinessKeyExpression, not even an explicit nil
### GetProcessInstanceBusinessKeyIn

`func (o *TaskQueryDto) GetProcessInstanceBusinessKeyIn() []string`

GetProcessInstanceBusinessKeyIn returns the ProcessInstanceBusinessKeyIn field if non-nil, zero value otherwise.

### GetProcessInstanceBusinessKeyInOk

`func (o *TaskQueryDto) GetProcessInstanceBusinessKeyInOk() (*[]string, bool)`

GetProcessInstanceBusinessKeyInOk returns a tuple with the ProcessInstanceBusinessKeyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceBusinessKeyIn

`func (o *TaskQueryDto) SetProcessInstanceBusinessKeyIn(v []string)`

SetProcessInstanceBusinessKeyIn sets ProcessInstanceBusinessKeyIn field to given value.

### HasProcessInstanceBusinessKeyIn

`func (o *TaskQueryDto) HasProcessInstanceBusinessKeyIn() bool`

HasProcessInstanceBusinessKeyIn returns a boolean if a field has been set.

### SetProcessInstanceBusinessKeyInNil

`func (o *TaskQueryDto) SetProcessInstanceBusinessKeyInNil(b bool)`

 SetProcessInstanceBusinessKeyInNil sets the value for ProcessInstanceBusinessKeyIn to be an explicit nil

### UnsetProcessInstanceBusinessKeyIn
`func (o *TaskQueryDto) UnsetProcessInstanceBusinessKeyIn()`

UnsetProcessInstanceBusinessKeyIn ensures that no value is present for ProcessInstanceBusinessKeyIn, not even an explicit nil
### GetProcessInstanceBusinessKeyLike

`func (o *TaskQueryDto) GetProcessInstanceBusinessKeyLike() string`

GetProcessInstanceBusinessKeyLike returns the ProcessInstanceBusinessKeyLike field if non-nil, zero value otherwise.

### GetProcessInstanceBusinessKeyLikeOk

`func (o *TaskQueryDto) GetProcessInstanceBusinessKeyLikeOk() (*string, bool)`

GetProcessInstanceBusinessKeyLikeOk returns a tuple with the ProcessInstanceBusinessKeyLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceBusinessKeyLike

`func (o *TaskQueryDto) SetProcessInstanceBusinessKeyLike(v string)`

SetProcessInstanceBusinessKeyLike sets ProcessInstanceBusinessKeyLike field to given value.

### HasProcessInstanceBusinessKeyLike

`func (o *TaskQueryDto) HasProcessInstanceBusinessKeyLike() bool`

HasProcessInstanceBusinessKeyLike returns a boolean if a field has been set.

### SetProcessInstanceBusinessKeyLikeNil

`func (o *TaskQueryDto) SetProcessInstanceBusinessKeyLikeNil(b bool)`

 SetProcessInstanceBusinessKeyLikeNil sets the value for ProcessInstanceBusinessKeyLike to be an explicit nil

### UnsetProcessInstanceBusinessKeyLike
`func (o *TaskQueryDto) UnsetProcessInstanceBusinessKeyLike()`

UnsetProcessInstanceBusinessKeyLike ensures that no value is present for ProcessInstanceBusinessKeyLike, not even an explicit nil
### GetProcessInstanceBusinessKeyLikeExpression

`func (o *TaskQueryDto) GetProcessInstanceBusinessKeyLikeExpression() string`

GetProcessInstanceBusinessKeyLikeExpression returns the ProcessInstanceBusinessKeyLikeExpression field if non-nil, zero value otherwise.

### GetProcessInstanceBusinessKeyLikeExpressionOk

`func (o *TaskQueryDto) GetProcessInstanceBusinessKeyLikeExpressionOk() (*string, bool)`

GetProcessInstanceBusinessKeyLikeExpressionOk returns a tuple with the ProcessInstanceBusinessKeyLikeExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceBusinessKeyLikeExpression

`func (o *TaskQueryDto) SetProcessInstanceBusinessKeyLikeExpression(v string)`

SetProcessInstanceBusinessKeyLikeExpression sets ProcessInstanceBusinessKeyLikeExpression field to given value.

### HasProcessInstanceBusinessKeyLikeExpression

`func (o *TaskQueryDto) HasProcessInstanceBusinessKeyLikeExpression() bool`

HasProcessInstanceBusinessKeyLikeExpression returns a boolean if a field has been set.

### SetProcessInstanceBusinessKeyLikeExpressionNil

`func (o *TaskQueryDto) SetProcessInstanceBusinessKeyLikeExpressionNil(b bool)`

 SetProcessInstanceBusinessKeyLikeExpressionNil sets the value for ProcessInstanceBusinessKeyLikeExpression to be an explicit nil

### UnsetProcessInstanceBusinessKeyLikeExpression
`func (o *TaskQueryDto) UnsetProcessInstanceBusinessKeyLikeExpression()`

UnsetProcessInstanceBusinessKeyLikeExpression ensures that no value is present for ProcessInstanceBusinessKeyLikeExpression, not even an explicit nil
### GetProcessDefinitionId

`func (o *TaskQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *TaskQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *TaskQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *TaskQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *TaskQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *TaskQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *TaskQueryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *TaskQueryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *TaskQueryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *TaskQueryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *TaskQueryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *TaskQueryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionKeyIn

`func (o *TaskQueryDto) GetProcessDefinitionKeyIn() []string`

GetProcessDefinitionKeyIn returns the ProcessDefinitionKeyIn field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyInOk

`func (o *TaskQueryDto) GetProcessDefinitionKeyInOk() (*[]string, bool)`

GetProcessDefinitionKeyInOk returns a tuple with the ProcessDefinitionKeyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKeyIn

`func (o *TaskQueryDto) SetProcessDefinitionKeyIn(v []string)`

SetProcessDefinitionKeyIn sets ProcessDefinitionKeyIn field to given value.

### HasProcessDefinitionKeyIn

`func (o *TaskQueryDto) HasProcessDefinitionKeyIn() bool`

HasProcessDefinitionKeyIn returns a boolean if a field has been set.

### SetProcessDefinitionKeyInNil

`func (o *TaskQueryDto) SetProcessDefinitionKeyInNil(b bool)`

 SetProcessDefinitionKeyInNil sets the value for ProcessDefinitionKeyIn to be an explicit nil

### UnsetProcessDefinitionKeyIn
`func (o *TaskQueryDto) UnsetProcessDefinitionKeyIn()`

UnsetProcessDefinitionKeyIn ensures that no value is present for ProcessDefinitionKeyIn, not even an explicit nil
### GetProcessDefinitionName

`func (o *TaskQueryDto) GetProcessDefinitionName() string`

GetProcessDefinitionName returns the ProcessDefinitionName field if non-nil, zero value otherwise.

### GetProcessDefinitionNameOk

`func (o *TaskQueryDto) GetProcessDefinitionNameOk() (*string, bool)`

GetProcessDefinitionNameOk returns a tuple with the ProcessDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionName

`func (o *TaskQueryDto) SetProcessDefinitionName(v string)`

SetProcessDefinitionName sets ProcessDefinitionName field to given value.

### HasProcessDefinitionName

`func (o *TaskQueryDto) HasProcessDefinitionName() bool`

HasProcessDefinitionName returns a boolean if a field has been set.

### SetProcessDefinitionNameNil

`func (o *TaskQueryDto) SetProcessDefinitionNameNil(b bool)`

 SetProcessDefinitionNameNil sets the value for ProcessDefinitionName to be an explicit nil

### UnsetProcessDefinitionName
`func (o *TaskQueryDto) UnsetProcessDefinitionName()`

UnsetProcessDefinitionName ensures that no value is present for ProcessDefinitionName, not even an explicit nil
### GetProcessDefinitionNameLike

`func (o *TaskQueryDto) GetProcessDefinitionNameLike() string`

GetProcessDefinitionNameLike returns the ProcessDefinitionNameLike field if non-nil, zero value otherwise.

### GetProcessDefinitionNameLikeOk

`func (o *TaskQueryDto) GetProcessDefinitionNameLikeOk() (*string, bool)`

GetProcessDefinitionNameLikeOk returns a tuple with the ProcessDefinitionNameLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionNameLike

`func (o *TaskQueryDto) SetProcessDefinitionNameLike(v string)`

SetProcessDefinitionNameLike sets ProcessDefinitionNameLike field to given value.

### HasProcessDefinitionNameLike

`func (o *TaskQueryDto) HasProcessDefinitionNameLike() bool`

HasProcessDefinitionNameLike returns a boolean if a field has been set.

### SetProcessDefinitionNameLikeNil

`func (o *TaskQueryDto) SetProcessDefinitionNameLikeNil(b bool)`

 SetProcessDefinitionNameLikeNil sets the value for ProcessDefinitionNameLike to be an explicit nil

### UnsetProcessDefinitionNameLike
`func (o *TaskQueryDto) UnsetProcessDefinitionNameLike()`

UnsetProcessDefinitionNameLike ensures that no value is present for ProcessDefinitionNameLike, not even an explicit nil
### GetExecutionId

`func (o *TaskQueryDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *TaskQueryDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *TaskQueryDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *TaskQueryDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *TaskQueryDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *TaskQueryDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetCaseInstanceId

`func (o *TaskQueryDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *TaskQueryDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *TaskQueryDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *TaskQueryDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *TaskQueryDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *TaskQueryDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetCaseInstanceBusinessKey

`func (o *TaskQueryDto) GetCaseInstanceBusinessKey() string`

GetCaseInstanceBusinessKey returns the CaseInstanceBusinessKey field if non-nil, zero value otherwise.

### GetCaseInstanceBusinessKeyOk

`func (o *TaskQueryDto) GetCaseInstanceBusinessKeyOk() (*string, bool)`

GetCaseInstanceBusinessKeyOk returns a tuple with the CaseInstanceBusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceBusinessKey

`func (o *TaskQueryDto) SetCaseInstanceBusinessKey(v string)`

SetCaseInstanceBusinessKey sets CaseInstanceBusinessKey field to given value.

### HasCaseInstanceBusinessKey

`func (o *TaskQueryDto) HasCaseInstanceBusinessKey() bool`

HasCaseInstanceBusinessKey returns a boolean if a field has been set.

### SetCaseInstanceBusinessKeyNil

`func (o *TaskQueryDto) SetCaseInstanceBusinessKeyNil(b bool)`

 SetCaseInstanceBusinessKeyNil sets the value for CaseInstanceBusinessKey to be an explicit nil

### UnsetCaseInstanceBusinessKey
`func (o *TaskQueryDto) UnsetCaseInstanceBusinessKey()`

UnsetCaseInstanceBusinessKey ensures that no value is present for CaseInstanceBusinessKey, not even an explicit nil
### GetCaseInstanceBusinessKeyLike

`func (o *TaskQueryDto) GetCaseInstanceBusinessKeyLike() string`

GetCaseInstanceBusinessKeyLike returns the CaseInstanceBusinessKeyLike field if non-nil, zero value otherwise.

### GetCaseInstanceBusinessKeyLikeOk

`func (o *TaskQueryDto) GetCaseInstanceBusinessKeyLikeOk() (*string, bool)`

GetCaseInstanceBusinessKeyLikeOk returns a tuple with the CaseInstanceBusinessKeyLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceBusinessKeyLike

`func (o *TaskQueryDto) SetCaseInstanceBusinessKeyLike(v string)`

SetCaseInstanceBusinessKeyLike sets CaseInstanceBusinessKeyLike field to given value.

### HasCaseInstanceBusinessKeyLike

`func (o *TaskQueryDto) HasCaseInstanceBusinessKeyLike() bool`

HasCaseInstanceBusinessKeyLike returns a boolean if a field has been set.

### SetCaseInstanceBusinessKeyLikeNil

`func (o *TaskQueryDto) SetCaseInstanceBusinessKeyLikeNil(b bool)`

 SetCaseInstanceBusinessKeyLikeNil sets the value for CaseInstanceBusinessKeyLike to be an explicit nil

### UnsetCaseInstanceBusinessKeyLike
`func (o *TaskQueryDto) UnsetCaseInstanceBusinessKeyLike()`

UnsetCaseInstanceBusinessKeyLike ensures that no value is present for CaseInstanceBusinessKeyLike, not even an explicit nil
### GetCaseDefinitionId

`func (o *TaskQueryDto) GetCaseDefinitionId() string`

GetCaseDefinitionId returns the CaseDefinitionId field if non-nil, zero value otherwise.

### GetCaseDefinitionIdOk

`func (o *TaskQueryDto) GetCaseDefinitionIdOk() (*string, bool)`

GetCaseDefinitionIdOk returns a tuple with the CaseDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionId

`func (o *TaskQueryDto) SetCaseDefinitionId(v string)`

SetCaseDefinitionId sets CaseDefinitionId field to given value.

### HasCaseDefinitionId

`func (o *TaskQueryDto) HasCaseDefinitionId() bool`

HasCaseDefinitionId returns a boolean if a field has been set.

### SetCaseDefinitionIdNil

`func (o *TaskQueryDto) SetCaseDefinitionIdNil(b bool)`

 SetCaseDefinitionIdNil sets the value for CaseDefinitionId to be an explicit nil

### UnsetCaseDefinitionId
`func (o *TaskQueryDto) UnsetCaseDefinitionId()`

UnsetCaseDefinitionId ensures that no value is present for CaseDefinitionId, not even an explicit nil
### GetCaseDefinitionKey

`func (o *TaskQueryDto) GetCaseDefinitionKey() string`

GetCaseDefinitionKey returns the CaseDefinitionKey field if non-nil, zero value otherwise.

### GetCaseDefinitionKeyOk

`func (o *TaskQueryDto) GetCaseDefinitionKeyOk() (*string, bool)`

GetCaseDefinitionKeyOk returns a tuple with the CaseDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionKey

`func (o *TaskQueryDto) SetCaseDefinitionKey(v string)`

SetCaseDefinitionKey sets CaseDefinitionKey field to given value.

### HasCaseDefinitionKey

`func (o *TaskQueryDto) HasCaseDefinitionKey() bool`

HasCaseDefinitionKey returns a boolean if a field has been set.

### SetCaseDefinitionKeyNil

`func (o *TaskQueryDto) SetCaseDefinitionKeyNil(b bool)`

 SetCaseDefinitionKeyNil sets the value for CaseDefinitionKey to be an explicit nil

### UnsetCaseDefinitionKey
`func (o *TaskQueryDto) UnsetCaseDefinitionKey()`

UnsetCaseDefinitionKey ensures that no value is present for CaseDefinitionKey, not even an explicit nil
### GetCaseDefinitionName

`func (o *TaskQueryDto) GetCaseDefinitionName() string`

GetCaseDefinitionName returns the CaseDefinitionName field if non-nil, zero value otherwise.

### GetCaseDefinitionNameOk

`func (o *TaskQueryDto) GetCaseDefinitionNameOk() (*string, bool)`

GetCaseDefinitionNameOk returns a tuple with the CaseDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionName

`func (o *TaskQueryDto) SetCaseDefinitionName(v string)`

SetCaseDefinitionName sets CaseDefinitionName field to given value.

### HasCaseDefinitionName

`func (o *TaskQueryDto) HasCaseDefinitionName() bool`

HasCaseDefinitionName returns a boolean if a field has been set.

### SetCaseDefinitionNameNil

`func (o *TaskQueryDto) SetCaseDefinitionNameNil(b bool)`

 SetCaseDefinitionNameNil sets the value for CaseDefinitionName to be an explicit nil

### UnsetCaseDefinitionName
`func (o *TaskQueryDto) UnsetCaseDefinitionName()`

UnsetCaseDefinitionName ensures that no value is present for CaseDefinitionName, not even an explicit nil
### GetCaseDefinitionNameLike

`func (o *TaskQueryDto) GetCaseDefinitionNameLike() string`

GetCaseDefinitionNameLike returns the CaseDefinitionNameLike field if non-nil, zero value otherwise.

### GetCaseDefinitionNameLikeOk

`func (o *TaskQueryDto) GetCaseDefinitionNameLikeOk() (*string, bool)`

GetCaseDefinitionNameLikeOk returns a tuple with the CaseDefinitionNameLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionNameLike

`func (o *TaskQueryDto) SetCaseDefinitionNameLike(v string)`

SetCaseDefinitionNameLike sets CaseDefinitionNameLike field to given value.

### HasCaseDefinitionNameLike

`func (o *TaskQueryDto) HasCaseDefinitionNameLike() bool`

HasCaseDefinitionNameLike returns a boolean if a field has been set.

### SetCaseDefinitionNameLikeNil

`func (o *TaskQueryDto) SetCaseDefinitionNameLikeNil(b bool)`

 SetCaseDefinitionNameLikeNil sets the value for CaseDefinitionNameLike to be an explicit nil

### UnsetCaseDefinitionNameLike
`func (o *TaskQueryDto) UnsetCaseDefinitionNameLike()`

UnsetCaseDefinitionNameLike ensures that no value is present for CaseDefinitionNameLike, not even an explicit nil
### GetCaseExecutionId

`func (o *TaskQueryDto) GetCaseExecutionId() string`

GetCaseExecutionId returns the CaseExecutionId field if non-nil, zero value otherwise.

### GetCaseExecutionIdOk

`func (o *TaskQueryDto) GetCaseExecutionIdOk() (*string, bool)`

GetCaseExecutionIdOk returns a tuple with the CaseExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExecutionId

`func (o *TaskQueryDto) SetCaseExecutionId(v string)`

SetCaseExecutionId sets CaseExecutionId field to given value.

### HasCaseExecutionId

`func (o *TaskQueryDto) HasCaseExecutionId() bool`

HasCaseExecutionId returns a boolean if a field has been set.

### SetCaseExecutionIdNil

`func (o *TaskQueryDto) SetCaseExecutionIdNil(b bool)`

 SetCaseExecutionIdNil sets the value for CaseExecutionId to be an explicit nil

### UnsetCaseExecutionId
`func (o *TaskQueryDto) UnsetCaseExecutionId()`

UnsetCaseExecutionId ensures that no value is present for CaseExecutionId, not even an explicit nil
### GetActivityInstanceIdIn

`func (o *TaskQueryDto) GetActivityInstanceIdIn() []string`

GetActivityInstanceIdIn returns the ActivityInstanceIdIn field if non-nil, zero value otherwise.

### GetActivityInstanceIdInOk

`func (o *TaskQueryDto) GetActivityInstanceIdInOk() (*[]string, bool)`

GetActivityInstanceIdInOk returns a tuple with the ActivityInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceIdIn

`func (o *TaskQueryDto) SetActivityInstanceIdIn(v []string)`

SetActivityInstanceIdIn sets ActivityInstanceIdIn field to given value.

### HasActivityInstanceIdIn

`func (o *TaskQueryDto) HasActivityInstanceIdIn() bool`

HasActivityInstanceIdIn returns a boolean if a field has been set.

### SetActivityInstanceIdInNil

`func (o *TaskQueryDto) SetActivityInstanceIdInNil(b bool)`

 SetActivityInstanceIdInNil sets the value for ActivityInstanceIdIn to be an explicit nil

### UnsetActivityInstanceIdIn
`func (o *TaskQueryDto) UnsetActivityInstanceIdIn()`

UnsetActivityInstanceIdIn ensures that no value is present for ActivityInstanceIdIn, not even an explicit nil
### GetTenantIdIn

`func (o *TaskQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *TaskQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *TaskQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *TaskQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *TaskQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *TaskQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *TaskQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *TaskQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *TaskQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *TaskQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *TaskQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *TaskQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetAssignee

`func (o *TaskQueryDto) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *TaskQueryDto) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *TaskQueryDto) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *TaskQueryDto) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *TaskQueryDto) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *TaskQueryDto) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetAssigneeExpression

`func (o *TaskQueryDto) GetAssigneeExpression() string`

GetAssigneeExpression returns the AssigneeExpression field if non-nil, zero value otherwise.

### GetAssigneeExpressionOk

`func (o *TaskQueryDto) GetAssigneeExpressionOk() (*string, bool)`

GetAssigneeExpressionOk returns a tuple with the AssigneeExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeExpression

`func (o *TaskQueryDto) SetAssigneeExpression(v string)`

SetAssigneeExpression sets AssigneeExpression field to given value.

### HasAssigneeExpression

`func (o *TaskQueryDto) HasAssigneeExpression() bool`

HasAssigneeExpression returns a boolean if a field has been set.

### SetAssigneeExpressionNil

`func (o *TaskQueryDto) SetAssigneeExpressionNil(b bool)`

 SetAssigneeExpressionNil sets the value for AssigneeExpression to be an explicit nil

### UnsetAssigneeExpression
`func (o *TaskQueryDto) UnsetAssigneeExpression()`

UnsetAssigneeExpression ensures that no value is present for AssigneeExpression, not even an explicit nil
### GetAssigneeLike

`func (o *TaskQueryDto) GetAssigneeLike() string`

GetAssigneeLike returns the AssigneeLike field if non-nil, zero value otherwise.

### GetAssigneeLikeOk

`func (o *TaskQueryDto) GetAssigneeLikeOk() (*string, bool)`

GetAssigneeLikeOk returns a tuple with the AssigneeLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeLike

`func (o *TaskQueryDto) SetAssigneeLike(v string)`

SetAssigneeLike sets AssigneeLike field to given value.

### HasAssigneeLike

`func (o *TaskQueryDto) HasAssigneeLike() bool`

HasAssigneeLike returns a boolean if a field has been set.

### SetAssigneeLikeNil

`func (o *TaskQueryDto) SetAssigneeLikeNil(b bool)`

 SetAssigneeLikeNil sets the value for AssigneeLike to be an explicit nil

### UnsetAssigneeLike
`func (o *TaskQueryDto) UnsetAssigneeLike()`

UnsetAssigneeLike ensures that no value is present for AssigneeLike, not even an explicit nil
### GetAssigneeLikeExpression

`func (o *TaskQueryDto) GetAssigneeLikeExpression() string`

GetAssigneeLikeExpression returns the AssigneeLikeExpression field if non-nil, zero value otherwise.

### GetAssigneeLikeExpressionOk

`func (o *TaskQueryDto) GetAssigneeLikeExpressionOk() (*string, bool)`

GetAssigneeLikeExpressionOk returns a tuple with the AssigneeLikeExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeLikeExpression

`func (o *TaskQueryDto) SetAssigneeLikeExpression(v string)`

SetAssigneeLikeExpression sets AssigneeLikeExpression field to given value.

### HasAssigneeLikeExpression

`func (o *TaskQueryDto) HasAssigneeLikeExpression() bool`

HasAssigneeLikeExpression returns a boolean if a field has been set.

### SetAssigneeLikeExpressionNil

`func (o *TaskQueryDto) SetAssigneeLikeExpressionNil(b bool)`

 SetAssigneeLikeExpressionNil sets the value for AssigneeLikeExpression to be an explicit nil

### UnsetAssigneeLikeExpression
`func (o *TaskQueryDto) UnsetAssigneeLikeExpression()`

UnsetAssigneeLikeExpression ensures that no value is present for AssigneeLikeExpression, not even an explicit nil
### GetAssigneeIn

`func (o *TaskQueryDto) GetAssigneeIn() []string`

GetAssigneeIn returns the AssigneeIn field if non-nil, zero value otherwise.

### GetAssigneeInOk

`func (o *TaskQueryDto) GetAssigneeInOk() (*[]string, bool)`

GetAssigneeInOk returns a tuple with the AssigneeIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeIn

`func (o *TaskQueryDto) SetAssigneeIn(v []string)`

SetAssigneeIn sets AssigneeIn field to given value.

### HasAssigneeIn

`func (o *TaskQueryDto) HasAssigneeIn() bool`

HasAssigneeIn returns a boolean if a field has been set.

### SetAssigneeInNil

`func (o *TaskQueryDto) SetAssigneeInNil(b bool)`

 SetAssigneeInNil sets the value for AssigneeIn to be an explicit nil

### UnsetAssigneeIn
`func (o *TaskQueryDto) UnsetAssigneeIn()`

UnsetAssigneeIn ensures that no value is present for AssigneeIn, not even an explicit nil
### GetAssigneeNotIn

`func (o *TaskQueryDto) GetAssigneeNotIn() []string`

GetAssigneeNotIn returns the AssigneeNotIn field if non-nil, zero value otherwise.

### GetAssigneeNotInOk

`func (o *TaskQueryDto) GetAssigneeNotInOk() (*[]string, bool)`

GetAssigneeNotInOk returns a tuple with the AssigneeNotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeNotIn

`func (o *TaskQueryDto) SetAssigneeNotIn(v []string)`

SetAssigneeNotIn sets AssigneeNotIn field to given value.

### HasAssigneeNotIn

`func (o *TaskQueryDto) HasAssigneeNotIn() bool`

HasAssigneeNotIn returns a boolean if a field has been set.

### SetAssigneeNotInNil

`func (o *TaskQueryDto) SetAssigneeNotInNil(b bool)`

 SetAssigneeNotInNil sets the value for AssigneeNotIn to be an explicit nil

### UnsetAssigneeNotIn
`func (o *TaskQueryDto) UnsetAssigneeNotIn()`

UnsetAssigneeNotIn ensures that no value is present for AssigneeNotIn, not even an explicit nil
### GetOwner

`func (o *TaskQueryDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TaskQueryDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TaskQueryDto) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TaskQueryDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *TaskQueryDto) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *TaskQueryDto) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetOwnerExpression

`func (o *TaskQueryDto) GetOwnerExpression() string`

GetOwnerExpression returns the OwnerExpression field if non-nil, zero value otherwise.

### GetOwnerExpressionOk

`func (o *TaskQueryDto) GetOwnerExpressionOk() (*string, bool)`

GetOwnerExpressionOk returns a tuple with the OwnerExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerExpression

`func (o *TaskQueryDto) SetOwnerExpression(v string)`

SetOwnerExpression sets OwnerExpression field to given value.

### HasOwnerExpression

`func (o *TaskQueryDto) HasOwnerExpression() bool`

HasOwnerExpression returns a boolean if a field has been set.

### SetOwnerExpressionNil

`func (o *TaskQueryDto) SetOwnerExpressionNil(b bool)`

 SetOwnerExpressionNil sets the value for OwnerExpression to be an explicit nil

### UnsetOwnerExpression
`func (o *TaskQueryDto) UnsetOwnerExpression()`

UnsetOwnerExpression ensures that no value is present for OwnerExpression, not even an explicit nil
### GetCandidateGroup

`func (o *TaskQueryDto) GetCandidateGroup() string`

GetCandidateGroup returns the CandidateGroup field if non-nil, zero value otherwise.

### GetCandidateGroupOk

`func (o *TaskQueryDto) GetCandidateGroupOk() (*string, bool)`

GetCandidateGroupOk returns a tuple with the CandidateGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateGroup

`func (o *TaskQueryDto) SetCandidateGroup(v string)`

SetCandidateGroup sets CandidateGroup field to given value.

### HasCandidateGroup

`func (o *TaskQueryDto) HasCandidateGroup() bool`

HasCandidateGroup returns a boolean if a field has been set.

### SetCandidateGroupNil

`func (o *TaskQueryDto) SetCandidateGroupNil(b bool)`

 SetCandidateGroupNil sets the value for CandidateGroup to be an explicit nil

### UnsetCandidateGroup
`func (o *TaskQueryDto) UnsetCandidateGroup()`

UnsetCandidateGroup ensures that no value is present for CandidateGroup, not even an explicit nil
### GetCandidateGroupExpression

`func (o *TaskQueryDto) GetCandidateGroupExpression() string`

GetCandidateGroupExpression returns the CandidateGroupExpression field if non-nil, zero value otherwise.

### GetCandidateGroupExpressionOk

`func (o *TaskQueryDto) GetCandidateGroupExpressionOk() (*string, bool)`

GetCandidateGroupExpressionOk returns a tuple with the CandidateGroupExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateGroupExpression

`func (o *TaskQueryDto) SetCandidateGroupExpression(v string)`

SetCandidateGroupExpression sets CandidateGroupExpression field to given value.

### HasCandidateGroupExpression

`func (o *TaskQueryDto) HasCandidateGroupExpression() bool`

HasCandidateGroupExpression returns a boolean if a field has been set.

### SetCandidateGroupExpressionNil

`func (o *TaskQueryDto) SetCandidateGroupExpressionNil(b bool)`

 SetCandidateGroupExpressionNil sets the value for CandidateGroupExpression to be an explicit nil

### UnsetCandidateGroupExpression
`func (o *TaskQueryDto) UnsetCandidateGroupExpression()`

UnsetCandidateGroupExpression ensures that no value is present for CandidateGroupExpression, not even an explicit nil
### GetCandidateUser

`func (o *TaskQueryDto) GetCandidateUser() string`

GetCandidateUser returns the CandidateUser field if non-nil, zero value otherwise.

### GetCandidateUserOk

`func (o *TaskQueryDto) GetCandidateUserOk() (*string, bool)`

GetCandidateUserOk returns a tuple with the CandidateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateUser

`func (o *TaskQueryDto) SetCandidateUser(v string)`

SetCandidateUser sets CandidateUser field to given value.

### HasCandidateUser

`func (o *TaskQueryDto) HasCandidateUser() bool`

HasCandidateUser returns a boolean if a field has been set.

### SetCandidateUserNil

`func (o *TaskQueryDto) SetCandidateUserNil(b bool)`

 SetCandidateUserNil sets the value for CandidateUser to be an explicit nil

### UnsetCandidateUser
`func (o *TaskQueryDto) UnsetCandidateUser()`

UnsetCandidateUser ensures that no value is present for CandidateUser, not even an explicit nil
### GetCandidateUserExpression

`func (o *TaskQueryDto) GetCandidateUserExpression() string`

GetCandidateUserExpression returns the CandidateUserExpression field if non-nil, zero value otherwise.

### GetCandidateUserExpressionOk

`func (o *TaskQueryDto) GetCandidateUserExpressionOk() (*string, bool)`

GetCandidateUserExpressionOk returns a tuple with the CandidateUserExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateUserExpression

`func (o *TaskQueryDto) SetCandidateUserExpression(v string)`

SetCandidateUserExpression sets CandidateUserExpression field to given value.

### HasCandidateUserExpression

`func (o *TaskQueryDto) HasCandidateUserExpression() bool`

HasCandidateUserExpression returns a boolean if a field has been set.

### SetCandidateUserExpressionNil

`func (o *TaskQueryDto) SetCandidateUserExpressionNil(b bool)`

 SetCandidateUserExpressionNil sets the value for CandidateUserExpression to be an explicit nil

### UnsetCandidateUserExpression
`func (o *TaskQueryDto) UnsetCandidateUserExpression()`

UnsetCandidateUserExpression ensures that no value is present for CandidateUserExpression, not even an explicit nil
### GetIncludeAssignedTasks

`func (o *TaskQueryDto) GetIncludeAssignedTasks() bool`

GetIncludeAssignedTasks returns the IncludeAssignedTasks field if non-nil, zero value otherwise.

### GetIncludeAssignedTasksOk

`func (o *TaskQueryDto) GetIncludeAssignedTasksOk() (*bool, bool)`

GetIncludeAssignedTasksOk returns a tuple with the IncludeAssignedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAssignedTasks

`func (o *TaskQueryDto) SetIncludeAssignedTasks(v bool)`

SetIncludeAssignedTasks sets IncludeAssignedTasks field to given value.

### HasIncludeAssignedTasks

`func (o *TaskQueryDto) HasIncludeAssignedTasks() bool`

HasIncludeAssignedTasks returns a boolean if a field has been set.

### SetIncludeAssignedTasksNil

`func (o *TaskQueryDto) SetIncludeAssignedTasksNil(b bool)`

 SetIncludeAssignedTasksNil sets the value for IncludeAssignedTasks to be an explicit nil

### UnsetIncludeAssignedTasks
`func (o *TaskQueryDto) UnsetIncludeAssignedTasks()`

UnsetIncludeAssignedTasks ensures that no value is present for IncludeAssignedTasks, not even an explicit nil
### GetInvolvedUser

`func (o *TaskQueryDto) GetInvolvedUser() string`

GetInvolvedUser returns the InvolvedUser field if non-nil, zero value otherwise.

### GetInvolvedUserOk

`func (o *TaskQueryDto) GetInvolvedUserOk() (*string, bool)`

GetInvolvedUserOk returns a tuple with the InvolvedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedUser

`func (o *TaskQueryDto) SetInvolvedUser(v string)`

SetInvolvedUser sets InvolvedUser field to given value.

### HasInvolvedUser

`func (o *TaskQueryDto) HasInvolvedUser() bool`

HasInvolvedUser returns a boolean if a field has been set.

### SetInvolvedUserNil

`func (o *TaskQueryDto) SetInvolvedUserNil(b bool)`

 SetInvolvedUserNil sets the value for InvolvedUser to be an explicit nil

### UnsetInvolvedUser
`func (o *TaskQueryDto) UnsetInvolvedUser()`

UnsetInvolvedUser ensures that no value is present for InvolvedUser, not even an explicit nil
### GetInvolvedUserExpression

`func (o *TaskQueryDto) GetInvolvedUserExpression() string`

GetInvolvedUserExpression returns the InvolvedUserExpression field if non-nil, zero value otherwise.

### GetInvolvedUserExpressionOk

`func (o *TaskQueryDto) GetInvolvedUserExpressionOk() (*string, bool)`

GetInvolvedUserExpressionOk returns a tuple with the InvolvedUserExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedUserExpression

`func (o *TaskQueryDto) SetInvolvedUserExpression(v string)`

SetInvolvedUserExpression sets InvolvedUserExpression field to given value.

### HasInvolvedUserExpression

`func (o *TaskQueryDto) HasInvolvedUserExpression() bool`

HasInvolvedUserExpression returns a boolean if a field has been set.

### SetInvolvedUserExpressionNil

`func (o *TaskQueryDto) SetInvolvedUserExpressionNil(b bool)`

 SetInvolvedUserExpressionNil sets the value for InvolvedUserExpression to be an explicit nil

### UnsetInvolvedUserExpression
`func (o *TaskQueryDto) UnsetInvolvedUserExpression()`

UnsetInvolvedUserExpression ensures that no value is present for InvolvedUserExpression, not even an explicit nil
### GetAssigned

`func (o *TaskQueryDto) GetAssigned() bool`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *TaskQueryDto) GetAssignedOk() (*bool, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *TaskQueryDto) SetAssigned(v bool)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *TaskQueryDto) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### SetAssignedNil

`func (o *TaskQueryDto) SetAssignedNil(b bool)`

 SetAssignedNil sets the value for Assigned to be an explicit nil

### UnsetAssigned
`func (o *TaskQueryDto) UnsetAssigned()`

UnsetAssigned ensures that no value is present for Assigned, not even an explicit nil
### GetUnassigned

`func (o *TaskQueryDto) GetUnassigned() bool`

GetUnassigned returns the Unassigned field if non-nil, zero value otherwise.

### GetUnassignedOk

`func (o *TaskQueryDto) GetUnassignedOk() (*bool, bool)`

GetUnassignedOk returns a tuple with the Unassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassigned

`func (o *TaskQueryDto) SetUnassigned(v bool)`

SetUnassigned sets Unassigned field to given value.

### HasUnassigned

`func (o *TaskQueryDto) HasUnassigned() bool`

HasUnassigned returns a boolean if a field has been set.

### SetUnassignedNil

`func (o *TaskQueryDto) SetUnassignedNil(b bool)`

 SetUnassignedNil sets the value for Unassigned to be an explicit nil

### UnsetUnassigned
`func (o *TaskQueryDto) UnsetUnassigned()`

UnsetUnassigned ensures that no value is present for Unassigned, not even an explicit nil
### GetTaskDefinitionKey

`func (o *TaskQueryDto) GetTaskDefinitionKey() string`

GetTaskDefinitionKey returns the TaskDefinitionKey field if non-nil, zero value otherwise.

### GetTaskDefinitionKeyOk

`func (o *TaskQueryDto) GetTaskDefinitionKeyOk() (*string, bool)`

GetTaskDefinitionKeyOk returns a tuple with the TaskDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionKey

`func (o *TaskQueryDto) SetTaskDefinitionKey(v string)`

SetTaskDefinitionKey sets TaskDefinitionKey field to given value.

### HasTaskDefinitionKey

`func (o *TaskQueryDto) HasTaskDefinitionKey() bool`

HasTaskDefinitionKey returns a boolean if a field has been set.

### SetTaskDefinitionKeyNil

`func (o *TaskQueryDto) SetTaskDefinitionKeyNil(b bool)`

 SetTaskDefinitionKeyNil sets the value for TaskDefinitionKey to be an explicit nil

### UnsetTaskDefinitionKey
`func (o *TaskQueryDto) UnsetTaskDefinitionKey()`

UnsetTaskDefinitionKey ensures that no value is present for TaskDefinitionKey, not even an explicit nil
### GetTaskDefinitionKeyIn

`func (o *TaskQueryDto) GetTaskDefinitionKeyIn() []string`

GetTaskDefinitionKeyIn returns the TaskDefinitionKeyIn field if non-nil, zero value otherwise.

### GetTaskDefinitionKeyInOk

`func (o *TaskQueryDto) GetTaskDefinitionKeyInOk() (*[]string, bool)`

GetTaskDefinitionKeyInOk returns a tuple with the TaskDefinitionKeyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionKeyIn

`func (o *TaskQueryDto) SetTaskDefinitionKeyIn(v []string)`

SetTaskDefinitionKeyIn sets TaskDefinitionKeyIn field to given value.

### HasTaskDefinitionKeyIn

`func (o *TaskQueryDto) HasTaskDefinitionKeyIn() bool`

HasTaskDefinitionKeyIn returns a boolean if a field has been set.

### SetTaskDefinitionKeyInNil

`func (o *TaskQueryDto) SetTaskDefinitionKeyInNil(b bool)`

 SetTaskDefinitionKeyInNil sets the value for TaskDefinitionKeyIn to be an explicit nil

### UnsetTaskDefinitionKeyIn
`func (o *TaskQueryDto) UnsetTaskDefinitionKeyIn()`

UnsetTaskDefinitionKeyIn ensures that no value is present for TaskDefinitionKeyIn, not even an explicit nil
### GetTaskDefinitionKeyLike

`func (o *TaskQueryDto) GetTaskDefinitionKeyLike() string`

GetTaskDefinitionKeyLike returns the TaskDefinitionKeyLike field if non-nil, zero value otherwise.

### GetTaskDefinitionKeyLikeOk

`func (o *TaskQueryDto) GetTaskDefinitionKeyLikeOk() (*string, bool)`

GetTaskDefinitionKeyLikeOk returns a tuple with the TaskDefinitionKeyLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionKeyLike

`func (o *TaskQueryDto) SetTaskDefinitionKeyLike(v string)`

SetTaskDefinitionKeyLike sets TaskDefinitionKeyLike field to given value.

### HasTaskDefinitionKeyLike

`func (o *TaskQueryDto) HasTaskDefinitionKeyLike() bool`

HasTaskDefinitionKeyLike returns a boolean if a field has been set.

### SetTaskDefinitionKeyLikeNil

`func (o *TaskQueryDto) SetTaskDefinitionKeyLikeNil(b bool)`

 SetTaskDefinitionKeyLikeNil sets the value for TaskDefinitionKeyLike to be an explicit nil

### UnsetTaskDefinitionKeyLike
`func (o *TaskQueryDto) UnsetTaskDefinitionKeyLike()`

UnsetTaskDefinitionKeyLike ensures that no value is present for TaskDefinitionKeyLike, not even an explicit nil
### GetName

`func (o *TaskQueryDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskQueryDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskQueryDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskQueryDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TaskQueryDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TaskQueryDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNameNotEqual

`func (o *TaskQueryDto) GetNameNotEqual() string`

GetNameNotEqual returns the NameNotEqual field if non-nil, zero value otherwise.

### GetNameNotEqualOk

`func (o *TaskQueryDto) GetNameNotEqualOk() (*string, bool)`

GetNameNotEqualOk returns a tuple with the NameNotEqual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameNotEqual

`func (o *TaskQueryDto) SetNameNotEqual(v string)`

SetNameNotEqual sets NameNotEqual field to given value.

### HasNameNotEqual

`func (o *TaskQueryDto) HasNameNotEqual() bool`

HasNameNotEqual returns a boolean if a field has been set.

### SetNameNotEqualNil

`func (o *TaskQueryDto) SetNameNotEqualNil(b bool)`

 SetNameNotEqualNil sets the value for NameNotEqual to be an explicit nil

### UnsetNameNotEqual
`func (o *TaskQueryDto) UnsetNameNotEqual()`

UnsetNameNotEqual ensures that no value is present for NameNotEqual, not even an explicit nil
### GetNameLike

`func (o *TaskQueryDto) GetNameLike() string`

GetNameLike returns the NameLike field if non-nil, zero value otherwise.

### GetNameLikeOk

`func (o *TaskQueryDto) GetNameLikeOk() (*string, bool)`

GetNameLikeOk returns a tuple with the NameLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLike

`func (o *TaskQueryDto) SetNameLike(v string)`

SetNameLike sets NameLike field to given value.

### HasNameLike

`func (o *TaskQueryDto) HasNameLike() bool`

HasNameLike returns a boolean if a field has been set.

### SetNameLikeNil

`func (o *TaskQueryDto) SetNameLikeNil(b bool)`

 SetNameLikeNil sets the value for NameLike to be an explicit nil

### UnsetNameLike
`func (o *TaskQueryDto) UnsetNameLike()`

UnsetNameLike ensures that no value is present for NameLike, not even an explicit nil
### GetNameNotLike

`func (o *TaskQueryDto) GetNameNotLike() string`

GetNameNotLike returns the NameNotLike field if non-nil, zero value otherwise.

### GetNameNotLikeOk

`func (o *TaskQueryDto) GetNameNotLikeOk() (*string, bool)`

GetNameNotLikeOk returns a tuple with the NameNotLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameNotLike

`func (o *TaskQueryDto) SetNameNotLike(v string)`

SetNameNotLike sets NameNotLike field to given value.

### HasNameNotLike

`func (o *TaskQueryDto) HasNameNotLike() bool`

HasNameNotLike returns a boolean if a field has been set.

### SetNameNotLikeNil

`func (o *TaskQueryDto) SetNameNotLikeNil(b bool)`

 SetNameNotLikeNil sets the value for NameNotLike to be an explicit nil

### UnsetNameNotLike
`func (o *TaskQueryDto) UnsetNameNotLike()`

UnsetNameNotLike ensures that no value is present for NameNotLike, not even an explicit nil
### GetDescription

`func (o *TaskQueryDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskQueryDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskQueryDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskQueryDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TaskQueryDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TaskQueryDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDescriptionLike

`func (o *TaskQueryDto) GetDescriptionLike() string`

GetDescriptionLike returns the DescriptionLike field if non-nil, zero value otherwise.

### GetDescriptionLikeOk

`func (o *TaskQueryDto) GetDescriptionLikeOk() (*string, bool)`

GetDescriptionLikeOk returns a tuple with the DescriptionLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLike

`func (o *TaskQueryDto) SetDescriptionLike(v string)`

SetDescriptionLike sets DescriptionLike field to given value.

### HasDescriptionLike

`func (o *TaskQueryDto) HasDescriptionLike() bool`

HasDescriptionLike returns a boolean if a field has been set.

### SetDescriptionLikeNil

`func (o *TaskQueryDto) SetDescriptionLikeNil(b bool)`

 SetDescriptionLikeNil sets the value for DescriptionLike to be an explicit nil

### UnsetDescriptionLike
`func (o *TaskQueryDto) UnsetDescriptionLike()`

UnsetDescriptionLike ensures that no value is present for DescriptionLike, not even an explicit nil
### GetPriority

`func (o *TaskQueryDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskQueryDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskQueryDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaskQueryDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *TaskQueryDto) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *TaskQueryDto) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetMaxPriority

`func (o *TaskQueryDto) GetMaxPriority() int32`

GetMaxPriority returns the MaxPriority field if non-nil, zero value otherwise.

### GetMaxPriorityOk

`func (o *TaskQueryDto) GetMaxPriorityOk() (*int32, bool)`

GetMaxPriorityOk returns a tuple with the MaxPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriority

`func (o *TaskQueryDto) SetMaxPriority(v int32)`

SetMaxPriority sets MaxPriority field to given value.

### HasMaxPriority

`func (o *TaskQueryDto) HasMaxPriority() bool`

HasMaxPriority returns a boolean if a field has been set.

### SetMaxPriorityNil

`func (o *TaskQueryDto) SetMaxPriorityNil(b bool)`

 SetMaxPriorityNil sets the value for MaxPriority to be an explicit nil

### UnsetMaxPriority
`func (o *TaskQueryDto) UnsetMaxPriority()`

UnsetMaxPriority ensures that no value is present for MaxPriority, not even an explicit nil
### GetMinPriority

`func (o *TaskQueryDto) GetMinPriority() int32`

GetMinPriority returns the MinPriority field if non-nil, zero value otherwise.

### GetMinPriorityOk

`func (o *TaskQueryDto) GetMinPriorityOk() (*int32, bool)`

GetMinPriorityOk returns a tuple with the MinPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPriority

`func (o *TaskQueryDto) SetMinPriority(v int32)`

SetMinPriority sets MinPriority field to given value.

### HasMinPriority

`func (o *TaskQueryDto) HasMinPriority() bool`

HasMinPriority returns a boolean if a field has been set.

### SetMinPriorityNil

`func (o *TaskQueryDto) SetMinPriorityNil(b bool)`

 SetMinPriorityNil sets the value for MinPriority to be an explicit nil

### UnsetMinPriority
`func (o *TaskQueryDto) UnsetMinPriority()`

UnsetMinPriority ensures that no value is present for MinPriority, not even an explicit nil
### GetDueDate

`func (o *TaskQueryDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TaskQueryDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TaskQueryDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TaskQueryDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *TaskQueryDto) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *TaskQueryDto) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetDueDateExpression

`func (o *TaskQueryDto) GetDueDateExpression() string`

GetDueDateExpression returns the DueDateExpression field if non-nil, zero value otherwise.

### GetDueDateExpressionOk

`func (o *TaskQueryDto) GetDueDateExpressionOk() (*string, bool)`

GetDueDateExpressionOk returns a tuple with the DueDateExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateExpression

`func (o *TaskQueryDto) SetDueDateExpression(v string)`

SetDueDateExpression sets DueDateExpression field to given value.

### HasDueDateExpression

`func (o *TaskQueryDto) HasDueDateExpression() bool`

HasDueDateExpression returns a boolean if a field has been set.

### SetDueDateExpressionNil

`func (o *TaskQueryDto) SetDueDateExpressionNil(b bool)`

 SetDueDateExpressionNil sets the value for DueDateExpression to be an explicit nil

### UnsetDueDateExpression
`func (o *TaskQueryDto) UnsetDueDateExpression()`

UnsetDueDateExpression ensures that no value is present for DueDateExpression, not even an explicit nil
### GetDueAfter

`func (o *TaskQueryDto) GetDueAfter() time.Time`

GetDueAfter returns the DueAfter field if non-nil, zero value otherwise.

### GetDueAfterOk

`func (o *TaskQueryDto) GetDueAfterOk() (*time.Time, bool)`

GetDueAfterOk returns a tuple with the DueAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAfter

`func (o *TaskQueryDto) SetDueAfter(v time.Time)`

SetDueAfter sets DueAfter field to given value.

### HasDueAfter

`func (o *TaskQueryDto) HasDueAfter() bool`

HasDueAfter returns a boolean if a field has been set.

### SetDueAfterNil

`func (o *TaskQueryDto) SetDueAfterNil(b bool)`

 SetDueAfterNil sets the value for DueAfter to be an explicit nil

### UnsetDueAfter
`func (o *TaskQueryDto) UnsetDueAfter()`

UnsetDueAfter ensures that no value is present for DueAfter, not even an explicit nil
### GetDueAfterExpression

`func (o *TaskQueryDto) GetDueAfterExpression() string`

GetDueAfterExpression returns the DueAfterExpression field if non-nil, zero value otherwise.

### GetDueAfterExpressionOk

`func (o *TaskQueryDto) GetDueAfterExpressionOk() (*string, bool)`

GetDueAfterExpressionOk returns a tuple with the DueAfterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAfterExpression

`func (o *TaskQueryDto) SetDueAfterExpression(v string)`

SetDueAfterExpression sets DueAfterExpression field to given value.

### HasDueAfterExpression

`func (o *TaskQueryDto) HasDueAfterExpression() bool`

HasDueAfterExpression returns a boolean if a field has been set.

### SetDueAfterExpressionNil

`func (o *TaskQueryDto) SetDueAfterExpressionNil(b bool)`

 SetDueAfterExpressionNil sets the value for DueAfterExpression to be an explicit nil

### UnsetDueAfterExpression
`func (o *TaskQueryDto) UnsetDueAfterExpression()`

UnsetDueAfterExpression ensures that no value is present for DueAfterExpression, not even an explicit nil
### GetDueBefore

`func (o *TaskQueryDto) GetDueBefore() time.Time`

GetDueBefore returns the DueBefore field if non-nil, zero value otherwise.

### GetDueBeforeOk

`func (o *TaskQueryDto) GetDueBeforeOk() (*time.Time, bool)`

GetDueBeforeOk returns a tuple with the DueBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueBefore

`func (o *TaskQueryDto) SetDueBefore(v time.Time)`

SetDueBefore sets DueBefore field to given value.

### HasDueBefore

`func (o *TaskQueryDto) HasDueBefore() bool`

HasDueBefore returns a boolean if a field has been set.

### SetDueBeforeNil

`func (o *TaskQueryDto) SetDueBeforeNil(b bool)`

 SetDueBeforeNil sets the value for DueBefore to be an explicit nil

### UnsetDueBefore
`func (o *TaskQueryDto) UnsetDueBefore()`

UnsetDueBefore ensures that no value is present for DueBefore, not even an explicit nil
### GetDueBeforeExpression

`func (o *TaskQueryDto) GetDueBeforeExpression() string`

GetDueBeforeExpression returns the DueBeforeExpression field if non-nil, zero value otherwise.

### GetDueBeforeExpressionOk

`func (o *TaskQueryDto) GetDueBeforeExpressionOk() (*string, bool)`

GetDueBeforeExpressionOk returns a tuple with the DueBeforeExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueBeforeExpression

`func (o *TaskQueryDto) SetDueBeforeExpression(v string)`

SetDueBeforeExpression sets DueBeforeExpression field to given value.

### HasDueBeforeExpression

`func (o *TaskQueryDto) HasDueBeforeExpression() bool`

HasDueBeforeExpression returns a boolean if a field has been set.

### SetDueBeforeExpressionNil

`func (o *TaskQueryDto) SetDueBeforeExpressionNil(b bool)`

 SetDueBeforeExpressionNil sets the value for DueBeforeExpression to be an explicit nil

### UnsetDueBeforeExpression
`func (o *TaskQueryDto) UnsetDueBeforeExpression()`

UnsetDueBeforeExpression ensures that no value is present for DueBeforeExpression, not even an explicit nil
### GetWithoutDueDate

`func (o *TaskQueryDto) GetWithoutDueDate() bool`

GetWithoutDueDate returns the WithoutDueDate field if non-nil, zero value otherwise.

### GetWithoutDueDateOk

`func (o *TaskQueryDto) GetWithoutDueDateOk() (*bool, bool)`

GetWithoutDueDateOk returns a tuple with the WithoutDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutDueDate

`func (o *TaskQueryDto) SetWithoutDueDate(v bool)`

SetWithoutDueDate sets WithoutDueDate field to given value.

### HasWithoutDueDate

`func (o *TaskQueryDto) HasWithoutDueDate() bool`

HasWithoutDueDate returns a boolean if a field has been set.

### SetWithoutDueDateNil

`func (o *TaskQueryDto) SetWithoutDueDateNil(b bool)`

 SetWithoutDueDateNil sets the value for WithoutDueDate to be an explicit nil

### UnsetWithoutDueDate
`func (o *TaskQueryDto) UnsetWithoutDueDate()`

UnsetWithoutDueDate ensures that no value is present for WithoutDueDate, not even an explicit nil
### GetFollowUpDate

`func (o *TaskQueryDto) GetFollowUpDate() time.Time`

GetFollowUpDate returns the FollowUpDate field if non-nil, zero value otherwise.

### GetFollowUpDateOk

`func (o *TaskQueryDto) GetFollowUpDateOk() (*time.Time, bool)`

GetFollowUpDateOk returns a tuple with the FollowUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpDate

`func (o *TaskQueryDto) SetFollowUpDate(v time.Time)`

SetFollowUpDate sets FollowUpDate field to given value.

### HasFollowUpDate

`func (o *TaskQueryDto) HasFollowUpDate() bool`

HasFollowUpDate returns a boolean if a field has been set.

### SetFollowUpDateNil

`func (o *TaskQueryDto) SetFollowUpDateNil(b bool)`

 SetFollowUpDateNil sets the value for FollowUpDate to be an explicit nil

### UnsetFollowUpDate
`func (o *TaskQueryDto) UnsetFollowUpDate()`

UnsetFollowUpDate ensures that no value is present for FollowUpDate, not even an explicit nil
### GetFollowUpDateExpression

`func (o *TaskQueryDto) GetFollowUpDateExpression() string`

GetFollowUpDateExpression returns the FollowUpDateExpression field if non-nil, zero value otherwise.

### GetFollowUpDateExpressionOk

`func (o *TaskQueryDto) GetFollowUpDateExpressionOk() (*string, bool)`

GetFollowUpDateExpressionOk returns a tuple with the FollowUpDateExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpDateExpression

`func (o *TaskQueryDto) SetFollowUpDateExpression(v string)`

SetFollowUpDateExpression sets FollowUpDateExpression field to given value.

### HasFollowUpDateExpression

`func (o *TaskQueryDto) HasFollowUpDateExpression() bool`

HasFollowUpDateExpression returns a boolean if a field has been set.

### SetFollowUpDateExpressionNil

`func (o *TaskQueryDto) SetFollowUpDateExpressionNil(b bool)`

 SetFollowUpDateExpressionNil sets the value for FollowUpDateExpression to be an explicit nil

### UnsetFollowUpDateExpression
`func (o *TaskQueryDto) UnsetFollowUpDateExpression()`

UnsetFollowUpDateExpression ensures that no value is present for FollowUpDateExpression, not even an explicit nil
### GetFollowUpAfter

`func (o *TaskQueryDto) GetFollowUpAfter() time.Time`

GetFollowUpAfter returns the FollowUpAfter field if non-nil, zero value otherwise.

### GetFollowUpAfterOk

`func (o *TaskQueryDto) GetFollowUpAfterOk() (*time.Time, bool)`

GetFollowUpAfterOk returns a tuple with the FollowUpAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpAfter

`func (o *TaskQueryDto) SetFollowUpAfter(v time.Time)`

SetFollowUpAfter sets FollowUpAfter field to given value.

### HasFollowUpAfter

`func (o *TaskQueryDto) HasFollowUpAfter() bool`

HasFollowUpAfter returns a boolean if a field has been set.

### SetFollowUpAfterNil

`func (o *TaskQueryDto) SetFollowUpAfterNil(b bool)`

 SetFollowUpAfterNil sets the value for FollowUpAfter to be an explicit nil

### UnsetFollowUpAfter
`func (o *TaskQueryDto) UnsetFollowUpAfter()`

UnsetFollowUpAfter ensures that no value is present for FollowUpAfter, not even an explicit nil
### GetFollowUpAfterExpression

`func (o *TaskQueryDto) GetFollowUpAfterExpression() string`

GetFollowUpAfterExpression returns the FollowUpAfterExpression field if non-nil, zero value otherwise.

### GetFollowUpAfterExpressionOk

`func (o *TaskQueryDto) GetFollowUpAfterExpressionOk() (*string, bool)`

GetFollowUpAfterExpressionOk returns a tuple with the FollowUpAfterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpAfterExpression

`func (o *TaskQueryDto) SetFollowUpAfterExpression(v string)`

SetFollowUpAfterExpression sets FollowUpAfterExpression field to given value.

### HasFollowUpAfterExpression

`func (o *TaskQueryDto) HasFollowUpAfterExpression() bool`

HasFollowUpAfterExpression returns a boolean if a field has been set.

### SetFollowUpAfterExpressionNil

`func (o *TaskQueryDto) SetFollowUpAfterExpressionNil(b bool)`

 SetFollowUpAfterExpressionNil sets the value for FollowUpAfterExpression to be an explicit nil

### UnsetFollowUpAfterExpression
`func (o *TaskQueryDto) UnsetFollowUpAfterExpression()`

UnsetFollowUpAfterExpression ensures that no value is present for FollowUpAfterExpression, not even an explicit nil
### GetFollowUpBefore

`func (o *TaskQueryDto) GetFollowUpBefore() string`

GetFollowUpBefore returns the FollowUpBefore field if non-nil, zero value otherwise.

### GetFollowUpBeforeOk

`func (o *TaskQueryDto) GetFollowUpBeforeOk() (*string, bool)`

GetFollowUpBeforeOk returns a tuple with the FollowUpBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpBefore

`func (o *TaskQueryDto) SetFollowUpBefore(v string)`

SetFollowUpBefore sets FollowUpBefore field to given value.

### HasFollowUpBefore

`func (o *TaskQueryDto) HasFollowUpBefore() bool`

HasFollowUpBefore returns a boolean if a field has been set.

### SetFollowUpBeforeNil

`func (o *TaskQueryDto) SetFollowUpBeforeNil(b bool)`

 SetFollowUpBeforeNil sets the value for FollowUpBefore to be an explicit nil

### UnsetFollowUpBefore
`func (o *TaskQueryDto) UnsetFollowUpBefore()`

UnsetFollowUpBefore ensures that no value is present for FollowUpBefore, not even an explicit nil
### GetFollowUpBeforeExpression

`func (o *TaskQueryDto) GetFollowUpBeforeExpression() string`

GetFollowUpBeforeExpression returns the FollowUpBeforeExpression field if non-nil, zero value otherwise.

### GetFollowUpBeforeExpressionOk

`func (o *TaskQueryDto) GetFollowUpBeforeExpressionOk() (*string, bool)`

GetFollowUpBeforeExpressionOk returns a tuple with the FollowUpBeforeExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpBeforeExpression

`func (o *TaskQueryDto) SetFollowUpBeforeExpression(v string)`

SetFollowUpBeforeExpression sets FollowUpBeforeExpression field to given value.

### HasFollowUpBeforeExpression

`func (o *TaskQueryDto) HasFollowUpBeforeExpression() bool`

HasFollowUpBeforeExpression returns a boolean if a field has been set.

### SetFollowUpBeforeExpressionNil

`func (o *TaskQueryDto) SetFollowUpBeforeExpressionNil(b bool)`

 SetFollowUpBeforeExpressionNil sets the value for FollowUpBeforeExpression to be an explicit nil

### UnsetFollowUpBeforeExpression
`func (o *TaskQueryDto) UnsetFollowUpBeforeExpression()`

UnsetFollowUpBeforeExpression ensures that no value is present for FollowUpBeforeExpression, not even an explicit nil
### GetFollowUpBeforeOrNotExistent

`func (o *TaskQueryDto) GetFollowUpBeforeOrNotExistent() time.Time`

GetFollowUpBeforeOrNotExistent returns the FollowUpBeforeOrNotExistent field if non-nil, zero value otherwise.

### GetFollowUpBeforeOrNotExistentOk

`func (o *TaskQueryDto) GetFollowUpBeforeOrNotExistentOk() (*time.Time, bool)`

GetFollowUpBeforeOrNotExistentOk returns a tuple with the FollowUpBeforeOrNotExistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpBeforeOrNotExistent

`func (o *TaskQueryDto) SetFollowUpBeforeOrNotExistent(v time.Time)`

SetFollowUpBeforeOrNotExistent sets FollowUpBeforeOrNotExistent field to given value.

### HasFollowUpBeforeOrNotExistent

`func (o *TaskQueryDto) HasFollowUpBeforeOrNotExistent() bool`

HasFollowUpBeforeOrNotExistent returns a boolean if a field has been set.

### SetFollowUpBeforeOrNotExistentNil

`func (o *TaskQueryDto) SetFollowUpBeforeOrNotExistentNil(b bool)`

 SetFollowUpBeforeOrNotExistentNil sets the value for FollowUpBeforeOrNotExistent to be an explicit nil

### UnsetFollowUpBeforeOrNotExistent
`func (o *TaskQueryDto) UnsetFollowUpBeforeOrNotExistent()`

UnsetFollowUpBeforeOrNotExistent ensures that no value is present for FollowUpBeforeOrNotExistent, not even an explicit nil
### GetFollowUpBeforeOrNotExistentExpression

`func (o *TaskQueryDto) GetFollowUpBeforeOrNotExistentExpression() string`

GetFollowUpBeforeOrNotExistentExpression returns the FollowUpBeforeOrNotExistentExpression field if non-nil, zero value otherwise.

### GetFollowUpBeforeOrNotExistentExpressionOk

`func (o *TaskQueryDto) GetFollowUpBeforeOrNotExistentExpressionOk() (*string, bool)`

GetFollowUpBeforeOrNotExistentExpressionOk returns a tuple with the FollowUpBeforeOrNotExistentExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpBeforeOrNotExistentExpression

`func (o *TaskQueryDto) SetFollowUpBeforeOrNotExistentExpression(v string)`

SetFollowUpBeforeOrNotExistentExpression sets FollowUpBeforeOrNotExistentExpression field to given value.

### HasFollowUpBeforeOrNotExistentExpression

`func (o *TaskQueryDto) HasFollowUpBeforeOrNotExistentExpression() bool`

HasFollowUpBeforeOrNotExistentExpression returns a boolean if a field has been set.

### SetFollowUpBeforeOrNotExistentExpressionNil

`func (o *TaskQueryDto) SetFollowUpBeforeOrNotExistentExpressionNil(b bool)`

 SetFollowUpBeforeOrNotExistentExpressionNil sets the value for FollowUpBeforeOrNotExistentExpression to be an explicit nil

### UnsetFollowUpBeforeOrNotExistentExpression
`func (o *TaskQueryDto) UnsetFollowUpBeforeOrNotExistentExpression()`

UnsetFollowUpBeforeOrNotExistentExpression ensures that no value is present for FollowUpBeforeOrNotExistentExpression, not even an explicit nil
### GetCreatedOn

`func (o *TaskQueryDto) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *TaskQueryDto) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *TaskQueryDto) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *TaskQueryDto) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### SetCreatedOnNil

`func (o *TaskQueryDto) SetCreatedOnNil(b bool)`

 SetCreatedOnNil sets the value for CreatedOn to be an explicit nil

### UnsetCreatedOn
`func (o *TaskQueryDto) UnsetCreatedOn()`

UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
### GetCreatedOnExpression

`func (o *TaskQueryDto) GetCreatedOnExpression() string`

GetCreatedOnExpression returns the CreatedOnExpression field if non-nil, zero value otherwise.

### GetCreatedOnExpressionOk

`func (o *TaskQueryDto) GetCreatedOnExpressionOk() (*string, bool)`

GetCreatedOnExpressionOk returns a tuple with the CreatedOnExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOnExpression

`func (o *TaskQueryDto) SetCreatedOnExpression(v string)`

SetCreatedOnExpression sets CreatedOnExpression field to given value.

### HasCreatedOnExpression

`func (o *TaskQueryDto) HasCreatedOnExpression() bool`

HasCreatedOnExpression returns a boolean if a field has been set.

### SetCreatedOnExpressionNil

`func (o *TaskQueryDto) SetCreatedOnExpressionNil(b bool)`

 SetCreatedOnExpressionNil sets the value for CreatedOnExpression to be an explicit nil

### UnsetCreatedOnExpression
`func (o *TaskQueryDto) UnsetCreatedOnExpression()`

UnsetCreatedOnExpression ensures that no value is present for CreatedOnExpression, not even an explicit nil
### GetCreatedAfter

`func (o *TaskQueryDto) GetCreatedAfter() time.Time`

GetCreatedAfter returns the CreatedAfter field if non-nil, zero value otherwise.

### GetCreatedAfterOk

`func (o *TaskQueryDto) GetCreatedAfterOk() (*time.Time, bool)`

GetCreatedAfterOk returns a tuple with the CreatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfter

`func (o *TaskQueryDto) SetCreatedAfter(v time.Time)`

SetCreatedAfter sets CreatedAfter field to given value.

### HasCreatedAfter

`func (o *TaskQueryDto) HasCreatedAfter() bool`

HasCreatedAfter returns a boolean if a field has been set.

### SetCreatedAfterNil

`func (o *TaskQueryDto) SetCreatedAfterNil(b bool)`

 SetCreatedAfterNil sets the value for CreatedAfter to be an explicit nil

### UnsetCreatedAfter
`func (o *TaskQueryDto) UnsetCreatedAfter()`

UnsetCreatedAfter ensures that no value is present for CreatedAfter, not even an explicit nil
### GetCreatedAfterExpression

`func (o *TaskQueryDto) GetCreatedAfterExpression() string`

GetCreatedAfterExpression returns the CreatedAfterExpression field if non-nil, zero value otherwise.

### GetCreatedAfterExpressionOk

`func (o *TaskQueryDto) GetCreatedAfterExpressionOk() (*string, bool)`

GetCreatedAfterExpressionOk returns a tuple with the CreatedAfterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfterExpression

`func (o *TaskQueryDto) SetCreatedAfterExpression(v string)`

SetCreatedAfterExpression sets CreatedAfterExpression field to given value.

### HasCreatedAfterExpression

`func (o *TaskQueryDto) HasCreatedAfterExpression() bool`

HasCreatedAfterExpression returns a boolean if a field has been set.

### SetCreatedAfterExpressionNil

`func (o *TaskQueryDto) SetCreatedAfterExpressionNil(b bool)`

 SetCreatedAfterExpressionNil sets the value for CreatedAfterExpression to be an explicit nil

### UnsetCreatedAfterExpression
`func (o *TaskQueryDto) UnsetCreatedAfterExpression()`

UnsetCreatedAfterExpression ensures that no value is present for CreatedAfterExpression, not even an explicit nil
### GetCreatedBefore

`func (o *TaskQueryDto) GetCreatedBefore() time.Time`

GetCreatedBefore returns the CreatedBefore field if non-nil, zero value otherwise.

### GetCreatedBeforeOk

`func (o *TaskQueryDto) GetCreatedBeforeOk() (*time.Time, bool)`

GetCreatedBeforeOk returns a tuple with the CreatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBefore

`func (o *TaskQueryDto) SetCreatedBefore(v time.Time)`

SetCreatedBefore sets CreatedBefore field to given value.

### HasCreatedBefore

`func (o *TaskQueryDto) HasCreatedBefore() bool`

HasCreatedBefore returns a boolean if a field has been set.

### SetCreatedBeforeNil

`func (o *TaskQueryDto) SetCreatedBeforeNil(b bool)`

 SetCreatedBeforeNil sets the value for CreatedBefore to be an explicit nil

### UnsetCreatedBefore
`func (o *TaskQueryDto) UnsetCreatedBefore()`

UnsetCreatedBefore ensures that no value is present for CreatedBefore, not even an explicit nil
### GetCreatedBeforeExpression

`func (o *TaskQueryDto) GetCreatedBeforeExpression() string`

GetCreatedBeforeExpression returns the CreatedBeforeExpression field if non-nil, zero value otherwise.

### GetCreatedBeforeExpressionOk

`func (o *TaskQueryDto) GetCreatedBeforeExpressionOk() (*string, bool)`

GetCreatedBeforeExpressionOk returns a tuple with the CreatedBeforeExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBeforeExpression

`func (o *TaskQueryDto) SetCreatedBeforeExpression(v string)`

SetCreatedBeforeExpression sets CreatedBeforeExpression field to given value.

### HasCreatedBeforeExpression

`func (o *TaskQueryDto) HasCreatedBeforeExpression() bool`

HasCreatedBeforeExpression returns a boolean if a field has been set.

### SetCreatedBeforeExpressionNil

`func (o *TaskQueryDto) SetCreatedBeforeExpressionNil(b bool)`

 SetCreatedBeforeExpressionNil sets the value for CreatedBeforeExpression to be an explicit nil

### UnsetCreatedBeforeExpression
`func (o *TaskQueryDto) UnsetCreatedBeforeExpression()`

UnsetCreatedBeforeExpression ensures that no value is present for CreatedBeforeExpression, not even an explicit nil
### GetUpdatedAfter

`func (o *TaskQueryDto) GetUpdatedAfter() time.Time`

GetUpdatedAfter returns the UpdatedAfter field if non-nil, zero value otherwise.

### GetUpdatedAfterOk

`func (o *TaskQueryDto) GetUpdatedAfterOk() (*time.Time, bool)`

GetUpdatedAfterOk returns a tuple with the UpdatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAfter

`func (o *TaskQueryDto) SetUpdatedAfter(v time.Time)`

SetUpdatedAfter sets UpdatedAfter field to given value.

### HasUpdatedAfter

`func (o *TaskQueryDto) HasUpdatedAfter() bool`

HasUpdatedAfter returns a boolean if a field has been set.

### SetUpdatedAfterNil

`func (o *TaskQueryDto) SetUpdatedAfterNil(b bool)`

 SetUpdatedAfterNil sets the value for UpdatedAfter to be an explicit nil

### UnsetUpdatedAfter
`func (o *TaskQueryDto) UnsetUpdatedAfter()`

UnsetUpdatedAfter ensures that no value is present for UpdatedAfter, not even an explicit nil
### GetUpdatedAfterExpression

`func (o *TaskQueryDto) GetUpdatedAfterExpression() string`

GetUpdatedAfterExpression returns the UpdatedAfterExpression field if non-nil, zero value otherwise.

### GetUpdatedAfterExpressionOk

`func (o *TaskQueryDto) GetUpdatedAfterExpressionOk() (*string, bool)`

GetUpdatedAfterExpressionOk returns a tuple with the UpdatedAfterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAfterExpression

`func (o *TaskQueryDto) SetUpdatedAfterExpression(v string)`

SetUpdatedAfterExpression sets UpdatedAfterExpression field to given value.

### HasUpdatedAfterExpression

`func (o *TaskQueryDto) HasUpdatedAfterExpression() bool`

HasUpdatedAfterExpression returns a boolean if a field has been set.

### SetUpdatedAfterExpressionNil

`func (o *TaskQueryDto) SetUpdatedAfterExpressionNil(b bool)`

 SetUpdatedAfterExpressionNil sets the value for UpdatedAfterExpression to be an explicit nil

### UnsetUpdatedAfterExpression
`func (o *TaskQueryDto) UnsetUpdatedAfterExpression()`

UnsetUpdatedAfterExpression ensures that no value is present for UpdatedAfterExpression, not even an explicit nil
### GetDelegationState

`func (o *TaskQueryDto) GetDelegationState() string`

GetDelegationState returns the DelegationState field if non-nil, zero value otherwise.

### GetDelegationStateOk

`func (o *TaskQueryDto) GetDelegationStateOk() (*string, bool)`

GetDelegationStateOk returns a tuple with the DelegationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationState

`func (o *TaskQueryDto) SetDelegationState(v string)`

SetDelegationState sets DelegationState field to given value.

### HasDelegationState

`func (o *TaskQueryDto) HasDelegationState() bool`

HasDelegationState returns a boolean if a field has been set.

### SetDelegationStateNil

`func (o *TaskQueryDto) SetDelegationStateNil(b bool)`

 SetDelegationStateNil sets the value for DelegationState to be an explicit nil

### UnsetDelegationState
`func (o *TaskQueryDto) UnsetDelegationState()`

UnsetDelegationState ensures that no value is present for DelegationState, not even an explicit nil
### GetCandidateGroups

`func (o *TaskQueryDto) GetCandidateGroups() []string`

GetCandidateGroups returns the CandidateGroups field if non-nil, zero value otherwise.

### GetCandidateGroupsOk

`func (o *TaskQueryDto) GetCandidateGroupsOk() (*[]string, bool)`

GetCandidateGroupsOk returns a tuple with the CandidateGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateGroups

`func (o *TaskQueryDto) SetCandidateGroups(v []string)`

SetCandidateGroups sets CandidateGroups field to given value.

### HasCandidateGroups

`func (o *TaskQueryDto) HasCandidateGroups() bool`

HasCandidateGroups returns a boolean if a field has been set.

### SetCandidateGroupsNil

`func (o *TaskQueryDto) SetCandidateGroupsNil(b bool)`

 SetCandidateGroupsNil sets the value for CandidateGroups to be an explicit nil

### UnsetCandidateGroups
`func (o *TaskQueryDto) UnsetCandidateGroups()`

UnsetCandidateGroups ensures that no value is present for CandidateGroups, not even an explicit nil
### GetCandidateGroupsExpression

`func (o *TaskQueryDto) GetCandidateGroupsExpression() string`

GetCandidateGroupsExpression returns the CandidateGroupsExpression field if non-nil, zero value otherwise.

### GetCandidateGroupsExpressionOk

`func (o *TaskQueryDto) GetCandidateGroupsExpressionOk() (*string, bool)`

GetCandidateGroupsExpressionOk returns a tuple with the CandidateGroupsExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateGroupsExpression

`func (o *TaskQueryDto) SetCandidateGroupsExpression(v string)`

SetCandidateGroupsExpression sets CandidateGroupsExpression field to given value.

### HasCandidateGroupsExpression

`func (o *TaskQueryDto) HasCandidateGroupsExpression() bool`

HasCandidateGroupsExpression returns a boolean if a field has been set.

### SetCandidateGroupsExpressionNil

`func (o *TaskQueryDto) SetCandidateGroupsExpressionNil(b bool)`

 SetCandidateGroupsExpressionNil sets the value for CandidateGroupsExpression to be an explicit nil

### UnsetCandidateGroupsExpression
`func (o *TaskQueryDto) UnsetCandidateGroupsExpression()`

UnsetCandidateGroupsExpression ensures that no value is present for CandidateGroupsExpression, not even an explicit nil
### GetWithCandidateGroups

`func (o *TaskQueryDto) GetWithCandidateGroups() bool`

GetWithCandidateGroups returns the WithCandidateGroups field if non-nil, zero value otherwise.

### GetWithCandidateGroupsOk

`func (o *TaskQueryDto) GetWithCandidateGroupsOk() (*bool, bool)`

GetWithCandidateGroupsOk returns a tuple with the WithCandidateGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCandidateGroups

`func (o *TaskQueryDto) SetWithCandidateGroups(v bool)`

SetWithCandidateGroups sets WithCandidateGroups field to given value.

### HasWithCandidateGroups

`func (o *TaskQueryDto) HasWithCandidateGroups() bool`

HasWithCandidateGroups returns a boolean if a field has been set.

### SetWithCandidateGroupsNil

`func (o *TaskQueryDto) SetWithCandidateGroupsNil(b bool)`

 SetWithCandidateGroupsNil sets the value for WithCandidateGroups to be an explicit nil

### UnsetWithCandidateGroups
`func (o *TaskQueryDto) UnsetWithCandidateGroups()`

UnsetWithCandidateGroups ensures that no value is present for WithCandidateGroups, not even an explicit nil
### GetWithoutCandidateGroups

`func (o *TaskQueryDto) GetWithoutCandidateGroups() bool`

GetWithoutCandidateGroups returns the WithoutCandidateGroups field if non-nil, zero value otherwise.

### GetWithoutCandidateGroupsOk

`func (o *TaskQueryDto) GetWithoutCandidateGroupsOk() (*bool, bool)`

GetWithoutCandidateGroupsOk returns a tuple with the WithoutCandidateGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutCandidateGroups

`func (o *TaskQueryDto) SetWithoutCandidateGroups(v bool)`

SetWithoutCandidateGroups sets WithoutCandidateGroups field to given value.

### HasWithoutCandidateGroups

`func (o *TaskQueryDto) HasWithoutCandidateGroups() bool`

HasWithoutCandidateGroups returns a boolean if a field has been set.

### SetWithoutCandidateGroupsNil

`func (o *TaskQueryDto) SetWithoutCandidateGroupsNil(b bool)`

 SetWithoutCandidateGroupsNil sets the value for WithoutCandidateGroups to be an explicit nil

### UnsetWithoutCandidateGroups
`func (o *TaskQueryDto) UnsetWithoutCandidateGroups()`

UnsetWithoutCandidateGroups ensures that no value is present for WithoutCandidateGroups, not even an explicit nil
### GetWithCandidateUsers

`func (o *TaskQueryDto) GetWithCandidateUsers() bool`

GetWithCandidateUsers returns the WithCandidateUsers field if non-nil, zero value otherwise.

### GetWithCandidateUsersOk

`func (o *TaskQueryDto) GetWithCandidateUsersOk() (*bool, bool)`

GetWithCandidateUsersOk returns a tuple with the WithCandidateUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCandidateUsers

`func (o *TaskQueryDto) SetWithCandidateUsers(v bool)`

SetWithCandidateUsers sets WithCandidateUsers field to given value.

### HasWithCandidateUsers

`func (o *TaskQueryDto) HasWithCandidateUsers() bool`

HasWithCandidateUsers returns a boolean if a field has been set.

### SetWithCandidateUsersNil

`func (o *TaskQueryDto) SetWithCandidateUsersNil(b bool)`

 SetWithCandidateUsersNil sets the value for WithCandidateUsers to be an explicit nil

### UnsetWithCandidateUsers
`func (o *TaskQueryDto) UnsetWithCandidateUsers()`

UnsetWithCandidateUsers ensures that no value is present for WithCandidateUsers, not even an explicit nil
### GetWithoutCandidateUsers

`func (o *TaskQueryDto) GetWithoutCandidateUsers() bool`

GetWithoutCandidateUsers returns the WithoutCandidateUsers field if non-nil, zero value otherwise.

### GetWithoutCandidateUsersOk

`func (o *TaskQueryDto) GetWithoutCandidateUsersOk() (*bool, bool)`

GetWithoutCandidateUsersOk returns a tuple with the WithoutCandidateUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutCandidateUsers

`func (o *TaskQueryDto) SetWithoutCandidateUsers(v bool)`

SetWithoutCandidateUsers sets WithoutCandidateUsers field to given value.

### HasWithoutCandidateUsers

`func (o *TaskQueryDto) HasWithoutCandidateUsers() bool`

HasWithoutCandidateUsers returns a boolean if a field has been set.

### SetWithoutCandidateUsersNil

`func (o *TaskQueryDto) SetWithoutCandidateUsersNil(b bool)`

 SetWithoutCandidateUsersNil sets the value for WithoutCandidateUsers to be an explicit nil

### UnsetWithoutCandidateUsers
`func (o *TaskQueryDto) UnsetWithoutCandidateUsers()`

UnsetWithoutCandidateUsers ensures that no value is present for WithoutCandidateUsers, not even an explicit nil
### GetActive

`func (o *TaskQueryDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TaskQueryDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TaskQueryDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TaskQueryDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *TaskQueryDto) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *TaskQueryDto) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetSuspended

`func (o *TaskQueryDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *TaskQueryDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *TaskQueryDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *TaskQueryDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *TaskQueryDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *TaskQueryDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetTaskVariables

`func (o *TaskQueryDto) GetTaskVariables() []VariableQueryParameterDto`

GetTaskVariables returns the TaskVariables field if non-nil, zero value otherwise.

### GetTaskVariablesOk

`func (o *TaskQueryDto) GetTaskVariablesOk() (*[]VariableQueryParameterDto, bool)`

GetTaskVariablesOk returns a tuple with the TaskVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskVariables

`func (o *TaskQueryDto) SetTaskVariables(v []VariableQueryParameterDto)`

SetTaskVariables sets TaskVariables field to given value.

### HasTaskVariables

`func (o *TaskQueryDto) HasTaskVariables() bool`

HasTaskVariables returns a boolean if a field has been set.

### SetTaskVariablesNil

`func (o *TaskQueryDto) SetTaskVariablesNil(b bool)`

 SetTaskVariablesNil sets the value for TaskVariables to be an explicit nil

### UnsetTaskVariables
`func (o *TaskQueryDto) UnsetTaskVariables()`

UnsetTaskVariables ensures that no value is present for TaskVariables, not even an explicit nil
### GetProcessVariables

`func (o *TaskQueryDto) GetProcessVariables() []VariableQueryParameterDto`

GetProcessVariables returns the ProcessVariables field if non-nil, zero value otherwise.

### GetProcessVariablesOk

`func (o *TaskQueryDto) GetProcessVariablesOk() (*[]VariableQueryParameterDto, bool)`

GetProcessVariablesOk returns a tuple with the ProcessVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessVariables

`func (o *TaskQueryDto) SetProcessVariables(v []VariableQueryParameterDto)`

SetProcessVariables sets ProcessVariables field to given value.

### HasProcessVariables

`func (o *TaskQueryDto) HasProcessVariables() bool`

HasProcessVariables returns a boolean if a field has been set.

### SetProcessVariablesNil

`func (o *TaskQueryDto) SetProcessVariablesNil(b bool)`

 SetProcessVariablesNil sets the value for ProcessVariables to be an explicit nil

### UnsetProcessVariables
`func (o *TaskQueryDto) UnsetProcessVariables()`

UnsetProcessVariables ensures that no value is present for ProcessVariables, not even an explicit nil
### GetCaseInstanceVariables

`func (o *TaskQueryDto) GetCaseInstanceVariables() []VariableQueryParameterDto`

GetCaseInstanceVariables returns the CaseInstanceVariables field if non-nil, zero value otherwise.

### GetCaseInstanceVariablesOk

`func (o *TaskQueryDto) GetCaseInstanceVariablesOk() (*[]VariableQueryParameterDto, bool)`

GetCaseInstanceVariablesOk returns a tuple with the CaseInstanceVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceVariables

`func (o *TaskQueryDto) SetCaseInstanceVariables(v []VariableQueryParameterDto)`

SetCaseInstanceVariables sets CaseInstanceVariables field to given value.

### HasCaseInstanceVariables

`func (o *TaskQueryDto) HasCaseInstanceVariables() bool`

HasCaseInstanceVariables returns a boolean if a field has been set.

### SetCaseInstanceVariablesNil

`func (o *TaskQueryDto) SetCaseInstanceVariablesNil(b bool)`

 SetCaseInstanceVariablesNil sets the value for CaseInstanceVariables to be an explicit nil

### UnsetCaseInstanceVariables
`func (o *TaskQueryDto) UnsetCaseInstanceVariables()`

UnsetCaseInstanceVariables ensures that no value is present for CaseInstanceVariables, not even an explicit nil
### GetVariableNamesIgnoreCase

`func (o *TaskQueryDto) GetVariableNamesIgnoreCase() bool`

GetVariableNamesIgnoreCase returns the VariableNamesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableNamesIgnoreCaseOk

`func (o *TaskQueryDto) GetVariableNamesIgnoreCaseOk() (*bool, bool)`

GetVariableNamesIgnoreCaseOk returns a tuple with the VariableNamesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNamesIgnoreCase

`func (o *TaskQueryDto) SetVariableNamesIgnoreCase(v bool)`

SetVariableNamesIgnoreCase sets VariableNamesIgnoreCase field to given value.

### HasVariableNamesIgnoreCase

`func (o *TaskQueryDto) HasVariableNamesIgnoreCase() bool`

HasVariableNamesIgnoreCase returns a boolean if a field has been set.

### SetVariableNamesIgnoreCaseNil

`func (o *TaskQueryDto) SetVariableNamesIgnoreCaseNil(b bool)`

 SetVariableNamesIgnoreCaseNil sets the value for VariableNamesIgnoreCase to be an explicit nil

### UnsetVariableNamesIgnoreCase
`func (o *TaskQueryDto) UnsetVariableNamesIgnoreCase()`

UnsetVariableNamesIgnoreCase ensures that no value is present for VariableNamesIgnoreCase, not even an explicit nil
### GetVariableValuesIgnoreCase

`func (o *TaskQueryDto) GetVariableValuesIgnoreCase() bool`

GetVariableValuesIgnoreCase returns the VariableValuesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableValuesIgnoreCaseOk

`func (o *TaskQueryDto) GetVariableValuesIgnoreCaseOk() (*bool, bool)`

GetVariableValuesIgnoreCaseOk returns a tuple with the VariableValuesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValuesIgnoreCase

`func (o *TaskQueryDto) SetVariableValuesIgnoreCase(v bool)`

SetVariableValuesIgnoreCase sets VariableValuesIgnoreCase field to given value.

### HasVariableValuesIgnoreCase

`func (o *TaskQueryDto) HasVariableValuesIgnoreCase() bool`

HasVariableValuesIgnoreCase returns a boolean if a field has been set.

### SetVariableValuesIgnoreCaseNil

`func (o *TaskQueryDto) SetVariableValuesIgnoreCaseNil(b bool)`

 SetVariableValuesIgnoreCaseNil sets the value for VariableValuesIgnoreCase to be an explicit nil

### UnsetVariableValuesIgnoreCase
`func (o *TaskQueryDto) UnsetVariableValuesIgnoreCase()`

UnsetVariableValuesIgnoreCase ensures that no value is present for VariableValuesIgnoreCase, not even an explicit nil
### GetParentTaskId

`func (o *TaskQueryDto) GetParentTaskId() string`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *TaskQueryDto) GetParentTaskIdOk() (*string, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *TaskQueryDto) SetParentTaskId(v string)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *TaskQueryDto) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### SetParentTaskIdNil

`func (o *TaskQueryDto) SetParentTaskIdNil(b bool)`

 SetParentTaskIdNil sets the value for ParentTaskId to be an explicit nil

### UnsetParentTaskId
`func (o *TaskQueryDto) UnsetParentTaskId()`

UnsetParentTaskId ensures that no value is present for ParentTaskId, not even an explicit nil
### GetOrQueries

`func (o *TaskQueryDto) GetOrQueries() []TaskQueryDto`

GetOrQueries returns the OrQueries field if non-nil, zero value otherwise.

### GetOrQueriesOk

`func (o *TaskQueryDto) GetOrQueriesOk() (*[]TaskQueryDto, bool)`

GetOrQueriesOk returns a tuple with the OrQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrQueries

`func (o *TaskQueryDto) SetOrQueries(v []TaskQueryDto)`

SetOrQueries sets OrQueries field to given value.

### HasOrQueries

`func (o *TaskQueryDto) HasOrQueries() bool`

HasOrQueries returns a boolean if a field has been set.

### SetOrQueriesNil

`func (o *TaskQueryDto) SetOrQueriesNil(b bool)`

 SetOrQueriesNil sets the value for OrQueries to be an explicit nil

### UnsetOrQueries
`func (o *TaskQueryDto) UnsetOrQueries()`

UnsetOrQueries ensures that no value is present for OrQueries, not even an explicit nil
### GetSorting

`func (o *TaskQueryDto) GetSorting() []TaskQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *TaskQueryDto) GetSortingOk() (*[]TaskQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *TaskQueryDto) SetSorting(v []TaskQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *TaskQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *TaskQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *TaskQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


