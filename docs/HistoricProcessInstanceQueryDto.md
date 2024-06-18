# HistoricProcessInstanceQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessInstanceId** | Pointer to **NullableString** | Filter by process instance id. | [optional] 
**ProcessInstanceIds** | Pointer to **[]string** | Filter by process instance ids. Must be a JSON array of &#x60;Strings&#x60;. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter by the process definition the instances run on. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Filter by the key of the process definition the instances run on. | [optional] 
**ProcessDefinitionKeyIn** | Pointer to **[]string** | Filter by a list of process definition keys. A process instance must have one of the given process definition keys. Must be a JSON array of &#x60;Strings&#x60;. | [optional] 
**ProcessDefinitionName** | Pointer to **NullableString** | Filter by the name of the process definition the instances run on. | [optional] 
**ProcessDefinitionNameLike** | Pointer to **NullableString** | Filter by process definition names that the parameter is a substring of. | [optional] 
**ProcessDefinitionKeyNotIn** | Pointer to **[]string** | Exclude instances that belong to a set of process definitions. Must be a JSON array of &#x60;Strings&#x60;. | [optional] 
**ProcessInstanceBusinessKey** | Pointer to **NullableString** | Filter by process instance business key. | [optional] 
**ProcessInstanceBusinessKeyIn** | Pointer to **[]string** | Filter by a list of business keys. A process instance must have one of the given business keys. Must be a JSON array of &#x60;Strings&#x60; | [optional] 
**ProcessInstanceBusinessKeyLike** | Pointer to **NullableString** | Filter by process instance business key that the parameter is a substring of. | [optional] 
**RootProcessInstances** | Pointer to **NullableBool** | Restrict the query to all process instances that are top level process instances. | [optional] 
**Finished** | Pointer to **NullableBool** | Only include finished process instances. This flag includes all process instances that are completed or terminated. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Unfinished** | Pointer to **NullableBool** | Only include unfinished process instances. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**WithIncidents** | Pointer to **NullableBool** | Only include process instances which have an incident. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**WithRootIncidents** | Pointer to **NullableBool** | Only include process instances which have a root incident. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**IncidentType** | Pointer to **NullableString** | Filter by the incident type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | [optional] 
**IncidentStatus** | Pointer to **NullableString** | Only include process instances which have an incident in status either open or resolved. To get all process instances, use the query parameter withIncidents. | [optional] 
**IncidentMessage** | Pointer to **NullableString** | Filter by the incident message. Exact match. | [optional] 
**IncidentMessageLike** | Pointer to **NullableString** | Filter by the incident message that the parameter is a substring of. | [optional] 
**StartedBefore** | Pointer to **NullableTime** | Restrict to instances that were started before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**StartedAfter** | Pointer to **NullableTime** | Restrict to instances that were started after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**FinishedBefore** | Pointer to **NullableTime** | Restrict to instances that were finished before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**FinishedAfter** | Pointer to **NullableTime** | Restrict to instances that were finished after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**ExecutedActivityAfter** | Pointer to **NullableTime** | Restrict to instances that executed an activity after the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**ExecutedActivityBefore** | Pointer to **NullableTime** | Restrict to instances that executed an activity before the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**ExecutedJobAfter** | Pointer to **NullableTime** | Restrict to instances that executed an job after the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**ExecutedJobBefore** | Pointer to **NullableTime** | Restrict to instances that executed an job before the given date (inclusive). By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**StartedBy** | Pointer to **NullableString** | Only include process instances that were started by the given user. | [optional] 
**SuperProcessInstanceId** | Pointer to **NullableString** | Restrict query to all process instances that are sub process instances of the given process instance. Takes a process instance id. | [optional] 
**SubProcessInstanceId** | Pointer to **NullableString** | Restrict query to one process instance that has a sub process instance with the given id. | [optional] 
**SuperCaseInstanceId** | Pointer to **NullableString** | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. | [optional] 
**SubCaseInstanceId** | Pointer to **NullableString** | Restrict query to one process instance that has a sub case instance with the given id. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Filter by a list of tenant ids. A process instance must have one of the given tenant ids. Must be a JSON array of &#x60;Strings&#x60; | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include historic process instances which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**ExecutedActivityIdIn** | Pointer to **[]string** | Restrict to instances that executed an activity with one of given ids. Must be a JSON array of &#x60;Strings&#x60; | [optional] 
**ActiveActivityIdIn** | Pointer to **[]string** | Restrict to instances that have an active activity with one of given ids. Must be a JSON array of &#x60;Strings&#x60; | [optional] 
**Active** | Pointer to **NullableBool** | Restrict to instances that are active. | [optional] 
**Suspended** | Pointer to **NullableBool** | Restrict to instances that are suspended. | [optional] 
**Completed** | Pointer to **NullableBool** | Restrict to instances that are completed. | [optional] 
**ExternallyTerminated** | Pointer to **NullableBool** | Restrict to instances that are externallyTerminated. | [optional] 
**InternallyTerminated** | Pointer to **NullableBool** | Restrict to instances that are internallyTerminated. | [optional] 
**Variables** | Pointer to [**[]VariableQueryParameterDto**](VariableQueryParameterDto.md) | A JSON array to only include process instances that have/had variables with certain values. The array consists of objects with the three properties &#x60;name&#x60;, &#x60;operator&#x60; and &#x60;value&#x60;. &#x60;name&#x60; (&#x60;String&#x60;) is the variable name, &#x60;operator&#x60; (&#x60;String&#x60;) is the comparison operator to be used and &#x60;value&#x60; the variable value.  Value may be &#x60;String&#x60;, &#x60;Number&#x60; or &#x60;Boolean&#x60;.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;.  | [optional] 
**VariableNamesIgnoreCase** | Pointer to **NullableBool** | Match all variable names provided in variables case-insensitively. If set to &#x60;true&#x60; variableName and variablename are treated as equal. | [optional] 
**VariableValuesIgnoreCase** | Pointer to **NullableBool** | Match all variable values provided in variables case-insensitively. If set to &#x60;true&#x60; variableValue and variablevalue are treated as equal. | [optional] 
**OrQueries** | Pointer to [**[]HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) | A JSON array of nested historic process instance queries with OR semantics.  A process instance matches a nested query if it fulfills at least one of the query&#39;s predicates.  With multiple nested queries, a process instance must fulfill at least one predicate of each query ([Conjunctive Normal Form](https://en.wikipedia.org/wiki/Conjunctive_normal_form)).  All process instance query properties can be used except for: &#x60;sorting&#x60;  See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/process-engine-api/#or-queries) for more information about OR queries. | [optional] 
**Sorting** | Pointer to [**[]HistoricProcessInstanceQueryDtoSortingInner**](HistoricProcessInstanceQueryDtoSortingInner.md) | Apply sorting of the result | [optional] 

## Methods

### NewHistoricProcessInstanceQueryDto

`func NewHistoricProcessInstanceQueryDto() *HistoricProcessInstanceQueryDto`

NewHistoricProcessInstanceQueryDto instantiates a new HistoricProcessInstanceQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricProcessInstanceQueryDtoWithDefaults

`func NewHistoricProcessInstanceQueryDtoWithDefaults() *HistoricProcessInstanceQueryDto`

NewHistoricProcessInstanceQueryDtoWithDefaults instantiates a new HistoricProcessInstanceQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessInstanceId

`func (o *HistoricProcessInstanceQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricProcessInstanceQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricProcessInstanceQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricProcessInstanceQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricProcessInstanceQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricProcessInstanceQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessInstanceIds

`func (o *HistoricProcessInstanceQueryDto) GetProcessInstanceIds() []string`

GetProcessInstanceIds returns the ProcessInstanceIds field if non-nil, zero value otherwise.

### GetProcessInstanceIdsOk

`func (o *HistoricProcessInstanceQueryDto) GetProcessInstanceIdsOk() (*[]string, bool)`

GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIds

`func (o *HistoricProcessInstanceQueryDto) SetProcessInstanceIds(v []string)`

SetProcessInstanceIds sets ProcessInstanceIds field to given value.

### HasProcessInstanceIds

`func (o *HistoricProcessInstanceQueryDto) HasProcessInstanceIds() bool`

HasProcessInstanceIds returns a boolean if a field has been set.

### SetProcessInstanceIdsNil

`func (o *HistoricProcessInstanceQueryDto) SetProcessInstanceIdsNil(b bool)`

 SetProcessInstanceIdsNil sets the value for ProcessInstanceIds to be an explicit nil

### UnsetProcessInstanceIds
`func (o *HistoricProcessInstanceQueryDto) UnsetProcessInstanceIds()`

UnsetProcessInstanceIds ensures that no value is present for ProcessInstanceIds, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricProcessInstanceQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricProcessInstanceQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricProcessInstanceQueryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricProcessInstanceQueryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionKeyIn

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionKeyIn() []string`

GetProcessDefinitionKeyIn returns the ProcessDefinitionKeyIn field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyInOk

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionKeyInOk() (*[]string, bool)`

GetProcessDefinitionKeyInOk returns a tuple with the ProcessDefinitionKeyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKeyIn

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionKeyIn(v []string)`

SetProcessDefinitionKeyIn sets ProcessDefinitionKeyIn field to given value.

### HasProcessDefinitionKeyIn

`func (o *HistoricProcessInstanceQueryDto) HasProcessDefinitionKeyIn() bool`

HasProcessDefinitionKeyIn returns a boolean if a field has been set.

### SetProcessDefinitionKeyInNil

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionKeyInNil(b bool)`

 SetProcessDefinitionKeyInNil sets the value for ProcessDefinitionKeyIn to be an explicit nil

### UnsetProcessDefinitionKeyIn
`func (o *HistoricProcessInstanceQueryDto) UnsetProcessDefinitionKeyIn()`

UnsetProcessDefinitionKeyIn ensures that no value is present for ProcessDefinitionKeyIn, not even an explicit nil
### GetProcessDefinitionName

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionName() string`

GetProcessDefinitionName returns the ProcessDefinitionName field if non-nil, zero value otherwise.

### GetProcessDefinitionNameOk

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionNameOk() (*string, bool)`

GetProcessDefinitionNameOk returns a tuple with the ProcessDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionName

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionName(v string)`

SetProcessDefinitionName sets ProcessDefinitionName field to given value.

### HasProcessDefinitionName

`func (o *HistoricProcessInstanceQueryDto) HasProcessDefinitionName() bool`

HasProcessDefinitionName returns a boolean if a field has been set.

### SetProcessDefinitionNameNil

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionNameNil(b bool)`

 SetProcessDefinitionNameNil sets the value for ProcessDefinitionName to be an explicit nil

### UnsetProcessDefinitionName
`func (o *HistoricProcessInstanceQueryDto) UnsetProcessDefinitionName()`

UnsetProcessDefinitionName ensures that no value is present for ProcessDefinitionName, not even an explicit nil
### GetProcessDefinitionNameLike

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionNameLike() string`

GetProcessDefinitionNameLike returns the ProcessDefinitionNameLike field if non-nil, zero value otherwise.

### GetProcessDefinitionNameLikeOk

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionNameLikeOk() (*string, bool)`

GetProcessDefinitionNameLikeOk returns a tuple with the ProcessDefinitionNameLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionNameLike

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionNameLike(v string)`

SetProcessDefinitionNameLike sets ProcessDefinitionNameLike field to given value.

### HasProcessDefinitionNameLike

`func (o *HistoricProcessInstanceQueryDto) HasProcessDefinitionNameLike() bool`

HasProcessDefinitionNameLike returns a boolean if a field has been set.

### SetProcessDefinitionNameLikeNil

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionNameLikeNil(b bool)`

 SetProcessDefinitionNameLikeNil sets the value for ProcessDefinitionNameLike to be an explicit nil

### UnsetProcessDefinitionNameLike
`func (o *HistoricProcessInstanceQueryDto) UnsetProcessDefinitionNameLike()`

UnsetProcessDefinitionNameLike ensures that no value is present for ProcessDefinitionNameLike, not even an explicit nil
### GetProcessDefinitionKeyNotIn

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionKeyNotIn() []string`

GetProcessDefinitionKeyNotIn returns the ProcessDefinitionKeyNotIn field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyNotInOk

`func (o *HistoricProcessInstanceQueryDto) GetProcessDefinitionKeyNotInOk() (*[]string, bool)`

GetProcessDefinitionKeyNotInOk returns a tuple with the ProcessDefinitionKeyNotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKeyNotIn

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionKeyNotIn(v []string)`

SetProcessDefinitionKeyNotIn sets ProcessDefinitionKeyNotIn field to given value.

### HasProcessDefinitionKeyNotIn

`func (o *HistoricProcessInstanceQueryDto) HasProcessDefinitionKeyNotIn() bool`

HasProcessDefinitionKeyNotIn returns a boolean if a field has been set.

### SetProcessDefinitionKeyNotInNil

`func (o *HistoricProcessInstanceQueryDto) SetProcessDefinitionKeyNotInNil(b bool)`

 SetProcessDefinitionKeyNotInNil sets the value for ProcessDefinitionKeyNotIn to be an explicit nil

### UnsetProcessDefinitionKeyNotIn
`func (o *HistoricProcessInstanceQueryDto) UnsetProcessDefinitionKeyNotIn()`

UnsetProcessDefinitionKeyNotIn ensures that no value is present for ProcessDefinitionKeyNotIn, not even an explicit nil
### GetProcessInstanceBusinessKey

`func (o *HistoricProcessInstanceQueryDto) GetProcessInstanceBusinessKey() string`

GetProcessInstanceBusinessKey returns the ProcessInstanceBusinessKey field if non-nil, zero value otherwise.

### GetProcessInstanceBusinessKeyOk

`func (o *HistoricProcessInstanceQueryDto) GetProcessInstanceBusinessKeyOk() (*string, bool)`

GetProcessInstanceBusinessKeyOk returns a tuple with the ProcessInstanceBusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceBusinessKey

`func (o *HistoricProcessInstanceQueryDto) SetProcessInstanceBusinessKey(v string)`

SetProcessInstanceBusinessKey sets ProcessInstanceBusinessKey field to given value.

### HasProcessInstanceBusinessKey

`func (o *HistoricProcessInstanceQueryDto) HasProcessInstanceBusinessKey() bool`

HasProcessInstanceBusinessKey returns a boolean if a field has been set.

### SetProcessInstanceBusinessKeyNil

`func (o *HistoricProcessInstanceQueryDto) SetProcessInstanceBusinessKeyNil(b bool)`

 SetProcessInstanceBusinessKeyNil sets the value for ProcessInstanceBusinessKey to be an explicit nil

### UnsetProcessInstanceBusinessKey
`func (o *HistoricProcessInstanceQueryDto) UnsetProcessInstanceBusinessKey()`

UnsetProcessInstanceBusinessKey ensures that no value is present for ProcessInstanceBusinessKey, not even an explicit nil
### GetProcessInstanceBusinessKeyIn

`func (o *HistoricProcessInstanceQueryDto) GetProcessInstanceBusinessKeyIn() []string`

GetProcessInstanceBusinessKeyIn returns the ProcessInstanceBusinessKeyIn field if non-nil, zero value otherwise.

### GetProcessInstanceBusinessKeyInOk

`func (o *HistoricProcessInstanceQueryDto) GetProcessInstanceBusinessKeyInOk() (*[]string, bool)`

GetProcessInstanceBusinessKeyInOk returns a tuple with the ProcessInstanceBusinessKeyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceBusinessKeyIn

`func (o *HistoricProcessInstanceQueryDto) SetProcessInstanceBusinessKeyIn(v []string)`

SetProcessInstanceBusinessKeyIn sets ProcessInstanceBusinessKeyIn field to given value.

### HasProcessInstanceBusinessKeyIn

`func (o *HistoricProcessInstanceQueryDto) HasProcessInstanceBusinessKeyIn() bool`

HasProcessInstanceBusinessKeyIn returns a boolean if a field has been set.

### SetProcessInstanceBusinessKeyInNil

`func (o *HistoricProcessInstanceQueryDto) SetProcessInstanceBusinessKeyInNil(b bool)`

 SetProcessInstanceBusinessKeyInNil sets the value for ProcessInstanceBusinessKeyIn to be an explicit nil

### UnsetProcessInstanceBusinessKeyIn
`func (o *HistoricProcessInstanceQueryDto) UnsetProcessInstanceBusinessKeyIn()`

UnsetProcessInstanceBusinessKeyIn ensures that no value is present for ProcessInstanceBusinessKeyIn, not even an explicit nil
### GetProcessInstanceBusinessKeyLike

`func (o *HistoricProcessInstanceQueryDto) GetProcessInstanceBusinessKeyLike() string`

GetProcessInstanceBusinessKeyLike returns the ProcessInstanceBusinessKeyLike field if non-nil, zero value otherwise.

### GetProcessInstanceBusinessKeyLikeOk

`func (o *HistoricProcessInstanceQueryDto) GetProcessInstanceBusinessKeyLikeOk() (*string, bool)`

GetProcessInstanceBusinessKeyLikeOk returns a tuple with the ProcessInstanceBusinessKeyLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceBusinessKeyLike

`func (o *HistoricProcessInstanceQueryDto) SetProcessInstanceBusinessKeyLike(v string)`

SetProcessInstanceBusinessKeyLike sets ProcessInstanceBusinessKeyLike field to given value.

### HasProcessInstanceBusinessKeyLike

`func (o *HistoricProcessInstanceQueryDto) HasProcessInstanceBusinessKeyLike() bool`

HasProcessInstanceBusinessKeyLike returns a boolean if a field has been set.

### SetProcessInstanceBusinessKeyLikeNil

`func (o *HistoricProcessInstanceQueryDto) SetProcessInstanceBusinessKeyLikeNil(b bool)`

 SetProcessInstanceBusinessKeyLikeNil sets the value for ProcessInstanceBusinessKeyLike to be an explicit nil

### UnsetProcessInstanceBusinessKeyLike
`func (o *HistoricProcessInstanceQueryDto) UnsetProcessInstanceBusinessKeyLike()`

UnsetProcessInstanceBusinessKeyLike ensures that no value is present for ProcessInstanceBusinessKeyLike, not even an explicit nil
### GetRootProcessInstances

`func (o *HistoricProcessInstanceQueryDto) GetRootProcessInstances() bool`

GetRootProcessInstances returns the RootProcessInstances field if non-nil, zero value otherwise.

### GetRootProcessInstancesOk

`func (o *HistoricProcessInstanceQueryDto) GetRootProcessInstancesOk() (*bool, bool)`

GetRootProcessInstancesOk returns a tuple with the RootProcessInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstances

`func (o *HistoricProcessInstanceQueryDto) SetRootProcessInstances(v bool)`

SetRootProcessInstances sets RootProcessInstances field to given value.

### HasRootProcessInstances

`func (o *HistoricProcessInstanceQueryDto) HasRootProcessInstances() bool`

HasRootProcessInstances returns a boolean if a field has been set.

### SetRootProcessInstancesNil

`func (o *HistoricProcessInstanceQueryDto) SetRootProcessInstancesNil(b bool)`

 SetRootProcessInstancesNil sets the value for RootProcessInstances to be an explicit nil

### UnsetRootProcessInstances
`func (o *HistoricProcessInstanceQueryDto) UnsetRootProcessInstances()`

UnsetRootProcessInstances ensures that no value is present for RootProcessInstances, not even an explicit nil
### GetFinished

`func (o *HistoricProcessInstanceQueryDto) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *HistoricProcessInstanceQueryDto) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *HistoricProcessInstanceQueryDto) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *HistoricProcessInstanceQueryDto) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *HistoricProcessInstanceQueryDto) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *HistoricProcessInstanceQueryDto) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetUnfinished

`func (o *HistoricProcessInstanceQueryDto) GetUnfinished() bool`

GetUnfinished returns the Unfinished field if non-nil, zero value otherwise.

### GetUnfinishedOk

`func (o *HistoricProcessInstanceQueryDto) GetUnfinishedOk() (*bool, bool)`

GetUnfinishedOk returns a tuple with the Unfinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinished

`func (o *HistoricProcessInstanceQueryDto) SetUnfinished(v bool)`

SetUnfinished sets Unfinished field to given value.

### HasUnfinished

`func (o *HistoricProcessInstanceQueryDto) HasUnfinished() bool`

HasUnfinished returns a boolean if a field has been set.

### SetUnfinishedNil

`func (o *HistoricProcessInstanceQueryDto) SetUnfinishedNil(b bool)`

 SetUnfinishedNil sets the value for Unfinished to be an explicit nil

### UnsetUnfinished
`func (o *HistoricProcessInstanceQueryDto) UnsetUnfinished()`

UnsetUnfinished ensures that no value is present for Unfinished, not even an explicit nil
### GetWithIncidents

`func (o *HistoricProcessInstanceQueryDto) GetWithIncidents() bool`

GetWithIncidents returns the WithIncidents field if non-nil, zero value otherwise.

### GetWithIncidentsOk

`func (o *HistoricProcessInstanceQueryDto) GetWithIncidentsOk() (*bool, bool)`

GetWithIncidentsOk returns a tuple with the WithIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithIncidents

`func (o *HistoricProcessInstanceQueryDto) SetWithIncidents(v bool)`

SetWithIncidents sets WithIncidents field to given value.

### HasWithIncidents

`func (o *HistoricProcessInstanceQueryDto) HasWithIncidents() bool`

HasWithIncidents returns a boolean if a field has been set.

### SetWithIncidentsNil

`func (o *HistoricProcessInstanceQueryDto) SetWithIncidentsNil(b bool)`

 SetWithIncidentsNil sets the value for WithIncidents to be an explicit nil

### UnsetWithIncidents
`func (o *HistoricProcessInstanceQueryDto) UnsetWithIncidents()`

UnsetWithIncidents ensures that no value is present for WithIncidents, not even an explicit nil
### GetWithRootIncidents

`func (o *HistoricProcessInstanceQueryDto) GetWithRootIncidents() bool`

GetWithRootIncidents returns the WithRootIncidents field if non-nil, zero value otherwise.

### GetWithRootIncidentsOk

`func (o *HistoricProcessInstanceQueryDto) GetWithRootIncidentsOk() (*bool, bool)`

GetWithRootIncidentsOk returns a tuple with the WithRootIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithRootIncidents

`func (o *HistoricProcessInstanceQueryDto) SetWithRootIncidents(v bool)`

SetWithRootIncidents sets WithRootIncidents field to given value.

### HasWithRootIncidents

`func (o *HistoricProcessInstanceQueryDto) HasWithRootIncidents() bool`

HasWithRootIncidents returns a boolean if a field has been set.

### SetWithRootIncidentsNil

`func (o *HistoricProcessInstanceQueryDto) SetWithRootIncidentsNil(b bool)`

 SetWithRootIncidentsNil sets the value for WithRootIncidents to be an explicit nil

### UnsetWithRootIncidents
`func (o *HistoricProcessInstanceQueryDto) UnsetWithRootIncidents()`

UnsetWithRootIncidents ensures that no value is present for WithRootIncidents, not even an explicit nil
### GetIncidentType

`func (o *HistoricProcessInstanceQueryDto) GetIncidentType() string`

GetIncidentType returns the IncidentType field if non-nil, zero value otherwise.

### GetIncidentTypeOk

`func (o *HistoricProcessInstanceQueryDto) GetIncidentTypeOk() (*string, bool)`

GetIncidentTypeOk returns a tuple with the IncidentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentType

`func (o *HistoricProcessInstanceQueryDto) SetIncidentType(v string)`

SetIncidentType sets IncidentType field to given value.

### HasIncidentType

`func (o *HistoricProcessInstanceQueryDto) HasIncidentType() bool`

HasIncidentType returns a boolean if a field has been set.

### SetIncidentTypeNil

`func (o *HistoricProcessInstanceQueryDto) SetIncidentTypeNil(b bool)`

 SetIncidentTypeNil sets the value for IncidentType to be an explicit nil

### UnsetIncidentType
`func (o *HistoricProcessInstanceQueryDto) UnsetIncidentType()`

UnsetIncidentType ensures that no value is present for IncidentType, not even an explicit nil
### GetIncidentStatus

`func (o *HistoricProcessInstanceQueryDto) GetIncidentStatus() string`

GetIncidentStatus returns the IncidentStatus field if non-nil, zero value otherwise.

### GetIncidentStatusOk

`func (o *HistoricProcessInstanceQueryDto) GetIncidentStatusOk() (*string, bool)`

GetIncidentStatusOk returns a tuple with the IncidentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentStatus

`func (o *HistoricProcessInstanceQueryDto) SetIncidentStatus(v string)`

SetIncidentStatus sets IncidentStatus field to given value.

### HasIncidentStatus

`func (o *HistoricProcessInstanceQueryDto) HasIncidentStatus() bool`

HasIncidentStatus returns a boolean if a field has been set.

### SetIncidentStatusNil

`func (o *HistoricProcessInstanceQueryDto) SetIncidentStatusNil(b bool)`

 SetIncidentStatusNil sets the value for IncidentStatus to be an explicit nil

### UnsetIncidentStatus
`func (o *HistoricProcessInstanceQueryDto) UnsetIncidentStatus()`

UnsetIncidentStatus ensures that no value is present for IncidentStatus, not even an explicit nil
### GetIncidentMessage

`func (o *HistoricProcessInstanceQueryDto) GetIncidentMessage() string`

GetIncidentMessage returns the IncidentMessage field if non-nil, zero value otherwise.

### GetIncidentMessageOk

`func (o *HistoricProcessInstanceQueryDto) GetIncidentMessageOk() (*string, bool)`

GetIncidentMessageOk returns a tuple with the IncidentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentMessage

`func (o *HistoricProcessInstanceQueryDto) SetIncidentMessage(v string)`

SetIncidentMessage sets IncidentMessage field to given value.

### HasIncidentMessage

`func (o *HistoricProcessInstanceQueryDto) HasIncidentMessage() bool`

HasIncidentMessage returns a boolean if a field has been set.

### SetIncidentMessageNil

`func (o *HistoricProcessInstanceQueryDto) SetIncidentMessageNil(b bool)`

 SetIncidentMessageNil sets the value for IncidentMessage to be an explicit nil

### UnsetIncidentMessage
`func (o *HistoricProcessInstanceQueryDto) UnsetIncidentMessage()`

UnsetIncidentMessage ensures that no value is present for IncidentMessage, not even an explicit nil
### GetIncidentMessageLike

`func (o *HistoricProcessInstanceQueryDto) GetIncidentMessageLike() string`

GetIncidentMessageLike returns the IncidentMessageLike field if non-nil, zero value otherwise.

### GetIncidentMessageLikeOk

`func (o *HistoricProcessInstanceQueryDto) GetIncidentMessageLikeOk() (*string, bool)`

GetIncidentMessageLikeOk returns a tuple with the IncidentMessageLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentMessageLike

`func (o *HistoricProcessInstanceQueryDto) SetIncidentMessageLike(v string)`

SetIncidentMessageLike sets IncidentMessageLike field to given value.

### HasIncidentMessageLike

`func (o *HistoricProcessInstanceQueryDto) HasIncidentMessageLike() bool`

HasIncidentMessageLike returns a boolean if a field has been set.

### SetIncidentMessageLikeNil

`func (o *HistoricProcessInstanceQueryDto) SetIncidentMessageLikeNil(b bool)`

 SetIncidentMessageLikeNil sets the value for IncidentMessageLike to be an explicit nil

### UnsetIncidentMessageLike
`func (o *HistoricProcessInstanceQueryDto) UnsetIncidentMessageLike()`

UnsetIncidentMessageLike ensures that no value is present for IncidentMessageLike, not even an explicit nil
### GetStartedBefore

`func (o *HistoricProcessInstanceQueryDto) GetStartedBefore() time.Time`

GetStartedBefore returns the StartedBefore field if non-nil, zero value otherwise.

### GetStartedBeforeOk

`func (o *HistoricProcessInstanceQueryDto) GetStartedBeforeOk() (*time.Time, bool)`

GetStartedBeforeOk returns a tuple with the StartedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBefore

`func (o *HistoricProcessInstanceQueryDto) SetStartedBefore(v time.Time)`

SetStartedBefore sets StartedBefore field to given value.

### HasStartedBefore

`func (o *HistoricProcessInstanceQueryDto) HasStartedBefore() bool`

HasStartedBefore returns a boolean if a field has been set.

### SetStartedBeforeNil

`func (o *HistoricProcessInstanceQueryDto) SetStartedBeforeNil(b bool)`

 SetStartedBeforeNil sets the value for StartedBefore to be an explicit nil

### UnsetStartedBefore
`func (o *HistoricProcessInstanceQueryDto) UnsetStartedBefore()`

UnsetStartedBefore ensures that no value is present for StartedBefore, not even an explicit nil
### GetStartedAfter

`func (o *HistoricProcessInstanceQueryDto) GetStartedAfter() time.Time`

GetStartedAfter returns the StartedAfter field if non-nil, zero value otherwise.

### GetStartedAfterOk

`func (o *HistoricProcessInstanceQueryDto) GetStartedAfterOk() (*time.Time, bool)`

GetStartedAfterOk returns a tuple with the StartedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAfter

`func (o *HistoricProcessInstanceQueryDto) SetStartedAfter(v time.Time)`

SetStartedAfter sets StartedAfter field to given value.

### HasStartedAfter

`func (o *HistoricProcessInstanceQueryDto) HasStartedAfter() bool`

HasStartedAfter returns a boolean if a field has been set.

### SetStartedAfterNil

`func (o *HistoricProcessInstanceQueryDto) SetStartedAfterNil(b bool)`

 SetStartedAfterNil sets the value for StartedAfter to be an explicit nil

### UnsetStartedAfter
`func (o *HistoricProcessInstanceQueryDto) UnsetStartedAfter()`

UnsetStartedAfter ensures that no value is present for StartedAfter, not even an explicit nil
### GetFinishedBefore

`func (o *HistoricProcessInstanceQueryDto) GetFinishedBefore() time.Time`

GetFinishedBefore returns the FinishedBefore field if non-nil, zero value otherwise.

### GetFinishedBeforeOk

`func (o *HistoricProcessInstanceQueryDto) GetFinishedBeforeOk() (*time.Time, bool)`

GetFinishedBeforeOk returns a tuple with the FinishedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedBefore

`func (o *HistoricProcessInstanceQueryDto) SetFinishedBefore(v time.Time)`

SetFinishedBefore sets FinishedBefore field to given value.

### HasFinishedBefore

`func (o *HistoricProcessInstanceQueryDto) HasFinishedBefore() bool`

HasFinishedBefore returns a boolean if a field has been set.

### SetFinishedBeforeNil

`func (o *HistoricProcessInstanceQueryDto) SetFinishedBeforeNil(b bool)`

 SetFinishedBeforeNil sets the value for FinishedBefore to be an explicit nil

### UnsetFinishedBefore
`func (o *HistoricProcessInstanceQueryDto) UnsetFinishedBefore()`

UnsetFinishedBefore ensures that no value is present for FinishedBefore, not even an explicit nil
### GetFinishedAfter

`func (o *HistoricProcessInstanceQueryDto) GetFinishedAfter() time.Time`

GetFinishedAfter returns the FinishedAfter field if non-nil, zero value otherwise.

### GetFinishedAfterOk

`func (o *HistoricProcessInstanceQueryDto) GetFinishedAfterOk() (*time.Time, bool)`

GetFinishedAfterOk returns a tuple with the FinishedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAfter

`func (o *HistoricProcessInstanceQueryDto) SetFinishedAfter(v time.Time)`

SetFinishedAfter sets FinishedAfter field to given value.

### HasFinishedAfter

`func (o *HistoricProcessInstanceQueryDto) HasFinishedAfter() bool`

HasFinishedAfter returns a boolean if a field has been set.

### SetFinishedAfterNil

`func (o *HistoricProcessInstanceQueryDto) SetFinishedAfterNil(b bool)`

 SetFinishedAfterNil sets the value for FinishedAfter to be an explicit nil

### UnsetFinishedAfter
`func (o *HistoricProcessInstanceQueryDto) UnsetFinishedAfter()`

UnsetFinishedAfter ensures that no value is present for FinishedAfter, not even an explicit nil
### GetExecutedActivityAfter

`func (o *HistoricProcessInstanceQueryDto) GetExecutedActivityAfter() time.Time`

GetExecutedActivityAfter returns the ExecutedActivityAfter field if non-nil, zero value otherwise.

### GetExecutedActivityAfterOk

`func (o *HistoricProcessInstanceQueryDto) GetExecutedActivityAfterOk() (*time.Time, bool)`

GetExecutedActivityAfterOk returns a tuple with the ExecutedActivityAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedActivityAfter

`func (o *HistoricProcessInstanceQueryDto) SetExecutedActivityAfter(v time.Time)`

SetExecutedActivityAfter sets ExecutedActivityAfter field to given value.

### HasExecutedActivityAfter

`func (o *HistoricProcessInstanceQueryDto) HasExecutedActivityAfter() bool`

HasExecutedActivityAfter returns a boolean if a field has been set.

### SetExecutedActivityAfterNil

`func (o *HistoricProcessInstanceQueryDto) SetExecutedActivityAfterNil(b bool)`

 SetExecutedActivityAfterNil sets the value for ExecutedActivityAfter to be an explicit nil

### UnsetExecutedActivityAfter
`func (o *HistoricProcessInstanceQueryDto) UnsetExecutedActivityAfter()`

UnsetExecutedActivityAfter ensures that no value is present for ExecutedActivityAfter, not even an explicit nil
### GetExecutedActivityBefore

`func (o *HistoricProcessInstanceQueryDto) GetExecutedActivityBefore() time.Time`

GetExecutedActivityBefore returns the ExecutedActivityBefore field if non-nil, zero value otherwise.

### GetExecutedActivityBeforeOk

`func (o *HistoricProcessInstanceQueryDto) GetExecutedActivityBeforeOk() (*time.Time, bool)`

GetExecutedActivityBeforeOk returns a tuple with the ExecutedActivityBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedActivityBefore

`func (o *HistoricProcessInstanceQueryDto) SetExecutedActivityBefore(v time.Time)`

SetExecutedActivityBefore sets ExecutedActivityBefore field to given value.

### HasExecutedActivityBefore

`func (o *HistoricProcessInstanceQueryDto) HasExecutedActivityBefore() bool`

HasExecutedActivityBefore returns a boolean if a field has been set.

### SetExecutedActivityBeforeNil

`func (o *HistoricProcessInstanceQueryDto) SetExecutedActivityBeforeNil(b bool)`

 SetExecutedActivityBeforeNil sets the value for ExecutedActivityBefore to be an explicit nil

### UnsetExecutedActivityBefore
`func (o *HistoricProcessInstanceQueryDto) UnsetExecutedActivityBefore()`

UnsetExecutedActivityBefore ensures that no value is present for ExecutedActivityBefore, not even an explicit nil
### GetExecutedJobAfter

`func (o *HistoricProcessInstanceQueryDto) GetExecutedJobAfter() time.Time`

GetExecutedJobAfter returns the ExecutedJobAfter field if non-nil, zero value otherwise.

### GetExecutedJobAfterOk

`func (o *HistoricProcessInstanceQueryDto) GetExecutedJobAfterOk() (*time.Time, bool)`

GetExecutedJobAfterOk returns a tuple with the ExecutedJobAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedJobAfter

`func (o *HistoricProcessInstanceQueryDto) SetExecutedJobAfter(v time.Time)`

SetExecutedJobAfter sets ExecutedJobAfter field to given value.

### HasExecutedJobAfter

`func (o *HistoricProcessInstanceQueryDto) HasExecutedJobAfter() bool`

HasExecutedJobAfter returns a boolean if a field has been set.

### SetExecutedJobAfterNil

`func (o *HistoricProcessInstanceQueryDto) SetExecutedJobAfterNil(b bool)`

 SetExecutedJobAfterNil sets the value for ExecutedJobAfter to be an explicit nil

### UnsetExecutedJobAfter
`func (o *HistoricProcessInstanceQueryDto) UnsetExecutedJobAfter()`

UnsetExecutedJobAfter ensures that no value is present for ExecutedJobAfter, not even an explicit nil
### GetExecutedJobBefore

`func (o *HistoricProcessInstanceQueryDto) GetExecutedJobBefore() time.Time`

GetExecutedJobBefore returns the ExecutedJobBefore field if non-nil, zero value otherwise.

### GetExecutedJobBeforeOk

`func (o *HistoricProcessInstanceQueryDto) GetExecutedJobBeforeOk() (*time.Time, bool)`

GetExecutedJobBeforeOk returns a tuple with the ExecutedJobBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedJobBefore

`func (o *HistoricProcessInstanceQueryDto) SetExecutedJobBefore(v time.Time)`

SetExecutedJobBefore sets ExecutedJobBefore field to given value.

### HasExecutedJobBefore

`func (o *HistoricProcessInstanceQueryDto) HasExecutedJobBefore() bool`

HasExecutedJobBefore returns a boolean if a field has been set.

### SetExecutedJobBeforeNil

`func (o *HistoricProcessInstanceQueryDto) SetExecutedJobBeforeNil(b bool)`

 SetExecutedJobBeforeNil sets the value for ExecutedJobBefore to be an explicit nil

### UnsetExecutedJobBefore
`func (o *HistoricProcessInstanceQueryDto) UnsetExecutedJobBefore()`

UnsetExecutedJobBefore ensures that no value is present for ExecutedJobBefore, not even an explicit nil
### GetStartedBy

`func (o *HistoricProcessInstanceQueryDto) GetStartedBy() string`

GetStartedBy returns the StartedBy field if non-nil, zero value otherwise.

### GetStartedByOk

`func (o *HistoricProcessInstanceQueryDto) GetStartedByOk() (*string, bool)`

GetStartedByOk returns a tuple with the StartedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBy

`func (o *HistoricProcessInstanceQueryDto) SetStartedBy(v string)`

SetStartedBy sets StartedBy field to given value.

### HasStartedBy

`func (o *HistoricProcessInstanceQueryDto) HasStartedBy() bool`

HasStartedBy returns a boolean if a field has been set.

### SetStartedByNil

`func (o *HistoricProcessInstanceQueryDto) SetStartedByNil(b bool)`

 SetStartedByNil sets the value for StartedBy to be an explicit nil

### UnsetStartedBy
`func (o *HistoricProcessInstanceQueryDto) UnsetStartedBy()`

UnsetStartedBy ensures that no value is present for StartedBy, not even an explicit nil
### GetSuperProcessInstanceId

`func (o *HistoricProcessInstanceQueryDto) GetSuperProcessInstanceId() string`

GetSuperProcessInstanceId returns the SuperProcessInstanceId field if non-nil, zero value otherwise.

### GetSuperProcessInstanceIdOk

`func (o *HistoricProcessInstanceQueryDto) GetSuperProcessInstanceIdOk() (*string, bool)`

GetSuperProcessInstanceIdOk returns a tuple with the SuperProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperProcessInstanceId

`func (o *HistoricProcessInstanceQueryDto) SetSuperProcessInstanceId(v string)`

SetSuperProcessInstanceId sets SuperProcessInstanceId field to given value.

### HasSuperProcessInstanceId

`func (o *HistoricProcessInstanceQueryDto) HasSuperProcessInstanceId() bool`

HasSuperProcessInstanceId returns a boolean if a field has been set.

### SetSuperProcessInstanceIdNil

`func (o *HistoricProcessInstanceQueryDto) SetSuperProcessInstanceIdNil(b bool)`

 SetSuperProcessInstanceIdNil sets the value for SuperProcessInstanceId to be an explicit nil

### UnsetSuperProcessInstanceId
`func (o *HistoricProcessInstanceQueryDto) UnsetSuperProcessInstanceId()`

UnsetSuperProcessInstanceId ensures that no value is present for SuperProcessInstanceId, not even an explicit nil
### GetSubProcessInstanceId

`func (o *HistoricProcessInstanceQueryDto) GetSubProcessInstanceId() string`

GetSubProcessInstanceId returns the SubProcessInstanceId field if non-nil, zero value otherwise.

### GetSubProcessInstanceIdOk

`func (o *HistoricProcessInstanceQueryDto) GetSubProcessInstanceIdOk() (*string, bool)`

GetSubProcessInstanceIdOk returns a tuple with the SubProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProcessInstanceId

`func (o *HistoricProcessInstanceQueryDto) SetSubProcessInstanceId(v string)`

SetSubProcessInstanceId sets SubProcessInstanceId field to given value.

### HasSubProcessInstanceId

`func (o *HistoricProcessInstanceQueryDto) HasSubProcessInstanceId() bool`

HasSubProcessInstanceId returns a boolean if a field has been set.

### SetSubProcessInstanceIdNil

`func (o *HistoricProcessInstanceQueryDto) SetSubProcessInstanceIdNil(b bool)`

 SetSubProcessInstanceIdNil sets the value for SubProcessInstanceId to be an explicit nil

### UnsetSubProcessInstanceId
`func (o *HistoricProcessInstanceQueryDto) UnsetSubProcessInstanceId()`

UnsetSubProcessInstanceId ensures that no value is present for SubProcessInstanceId, not even an explicit nil
### GetSuperCaseInstanceId

`func (o *HistoricProcessInstanceQueryDto) GetSuperCaseInstanceId() string`

GetSuperCaseInstanceId returns the SuperCaseInstanceId field if non-nil, zero value otherwise.

### GetSuperCaseInstanceIdOk

`func (o *HistoricProcessInstanceQueryDto) GetSuperCaseInstanceIdOk() (*string, bool)`

GetSuperCaseInstanceIdOk returns a tuple with the SuperCaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperCaseInstanceId

`func (o *HistoricProcessInstanceQueryDto) SetSuperCaseInstanceId(v string)`

SetSuperCaseInstanceId sets SuperCaseInstanceId field to given value.

### HasSuperCaseInstanceId

`func (o *HistoricProcessInstanceQueryDto) HasSuperCaseInstanceId() bool`

HasSuperCaseInstanceId returns a boolean if a field has been set.

### SetSuperCaseInstanceIdNil

`func (o *HistoricProcessInstanceQueryDto) SetSuperCaseInstanceIdNil(b bool)`

 SetSuperCaseInstanceIdNil sets the value for SuperCaseInstanceId to be an explicit nil

### UnsetSuperCaseInstanceId
`func (o *HistoricProcessInstanceQueryDto) UnsetSuperCaseInstanceId()`

UnsetSuperCaseInstanceId ensures that no value is present for SuperCaseInstanceId, not even an explicit nil
### GetSubCaseInstanceId

`func (o *HistoricProcessInstanceQueryDto) GetSubCaseInstanceId() string`

GetSubCaseInstanceId returns the SubCaseInstanceId field if non-nil, zero value otherwise.

### GetSubCaseInstanceIdOk

`func (o *HistoricProcessInstanceQueryDto) GetSubCaseInstanceIdOk() (*string, bool)`

GetSubCaseInstanceIdOk returns a tuple with the SubCaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCaseInstanceId

`func (o *HistoricProcessInstanceQueryDto) SetSubCaseInstanceId(v string)`

SetSubCaseInstanceId sets SubCaseInstanceId field to given value.

### HasSubCaseInstanceId

`func (o *HistoricProcessInstanceQueryDto) HasSubCaseInstanceId() bool`

HasSubCaseInstanceId returns a boolean if a field has been set.

### SetSubCaseInstanceIdNil

`func (o *HistoricProcessInstanceQueryDto) SetSubCaseInstanceIdNil(b bool)`

 SetSubCaseInstanceIdNil sets the value for SubCaseInstanceId to be an explicit nil

### UnsetSubCaseInstanceId
`func (o *HistoricProcessInstanceQueryDto) UnsetSubCaseInstanceId()`

UnsetSubCaseInstanceId ensures that no value is present for SubCaseInstanceId, not even an explicit nil
### GetCaseInstanceId

`func (o *HistoricProcessInstanceQueryDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *HistoricProcessInstanceQueryDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *HistoricProcessInstanceQueryDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *HistoricProcessInstanceQueryDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *HistoricProcessInstanceQueryDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *HistoricProcessInstanceQueryDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetTenantIdIn

`func (o *HistoricProcessInstanceQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *HistoricProcessInstanceQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *HistoricProcessInstanceQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *HistoricProcessInstanceQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *HistoricProcessInstanceQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *HistoricProcessInstanceQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *HistoricProcessInstanceQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *HistoricProcessInstanceQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *HistoricProcessInstanceQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *HistoricProcessInstanceQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *HistoricProcessInstanceQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *HistoricProcessInstanceQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetExecutedActivityIdIn

`func (o *HistoricProcessInstanceQueryDto) GetExecutedActivityIdIn() []string`

GetExecutedActivityIdIn returns the ExecutedActivityIdIn field if non-nil, zero value otherwise.

### GetExecutedActivityIdInOk

`func (o *HistoricProcessInstanceQueryDto) GetExecutedActivityIdInOk() (*[]string, bool)`

GetExecutedActivityIdInOk returns a tuple with the ExecutedActivityIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedActivityIdIn

`func (o *HistoricProcessInstanceQueryDto) SetExecutedActivityIdIn(v []string)`

SetExecutedActivityIdIn sets ExecutedActivityIdIn field to given value.

### HasExecutedActivityIdIn

`func (o *HistoricProcessInstanceQueryDto) HasExecutedActivityIdIn() bool`

HasExecutedActivityIdIn returns a boolean if a field has been set.

### SetExecutedActivityIdInNil

`func (o *HistoricProcessInstanceQueryDto) SetExecutedActivityIdInNil(b bool)`

 SetExecutedActivityIdInNil sets the value for ExecutedActivityIdIn to be an explicit nil

### UnsetExecutedActivityIdIn
`func (o *HistoricProcessInstanceQueryDto) UnsetExecutedActivityIdIn()`

UnsetExecutedActivityIdIn ensures that no value is present for ExecutedActivityIdIn, not even an explicit nil
### GetActiveActivityIdIn

`func (o *HistoricProcessInstanceQueryDto) GetActiveActivityIdIn() []string`

GetActiveActivityIdIn returns the ActiveActivityIdIn field if non-nil, zero value otherwise.

### GetActiveActivityIdInOk

`func (o *HistoricProcessInstanceQueryDto) GetActiveActivityIdInOk() (*[]string, bool)`

GetActiveActivityIdInOk returns a tuple with the ActiveActivityIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveActivityIdIn

`func (o *HistoricProcessInstanceQueryDto) SetActiveActivityIdIn(v []string)`

SetActiveActivityIdIn sets ActiveActivityIdIn field to given value.

### HasActiveActivityIdIn

`func (o *HistoricProcessInstanceQueryDto) HasActiveActivityIdIn() bool`

HasActiveActivityIdIn returns a boolean if a field has been set.

### SetActiveActivityIdInNil

`func (o *HistoricProcessInstanceQueryDto) SetActiveActivityIdInNil(b bool)`

 SetActiveActivityIdInNil sets the value for ActiveActivityIdIn to be an explicit nil

### UnsetActiveActivityIdIn
`func (o *HistoricProcessInstanceQueryDto) UnsetActiveActivityIdIn()`

UnsetActiveActivityIdIn ensures that no value is present for ActiveActivityIdIn, not even an explicit nil
### GetActive

`func (o *HistoricProcessInstanceQueryDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *HistoricProcessInstanceQueryDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *HistoricProcessInstanceQueryDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *HistoricProcessInstanceQueryDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *HistoricProcessInstanceQueryDto) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *HistoricProcessInstanceQueryDto) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetSuspended

`func (o *HistoricProcessInstanceQueryDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *HistoricProcessInstanceQueryDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *HistoricProcessInstanceQueryDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *HistoricProcessInstanceQueryDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *HistoricProcessInstanceQueryDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *HistoricProcessInstanceQueryDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetCompleted

`func (o *HistoricProcessInstanceQueryDto) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *HistoricProcessInstanceQueryDto) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *HistoricProcessInstanceQueryDto) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *HistoricProcessInstanceQueryDto) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### SetCompletedNil

`func (o *HistoricProcessInstanceQueryDto) SetCompletedNil(b bool)`

 SetCompletedNil sets the value for Completed to be an explicit nil

### UnsetCompleted
`func (o *HistoricProcessInstanceQueryDto) UnsetCompleted()`

UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
### GetExternallyTerminated

`func (o *HistoricProcessInstanceQueryDto) GetExternallyTerminated() bool`

GetExternallyTerminated returns the ExternallyTerminated field if non-nil, zero value otherwise.

### GetExternallyTerminatedOk

`func (o *HistoricProcessInstanceQueryDto) GetExternallyTerminatedOk() (*bool, bool)`

GetExternallyTerminatedOk returns a tuple with the ExternallyTerminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyTerminated

`func (o *HistoricProcessInstanceQueryDto) SetExternallyTerminated(v bool)`

SetExternallyTerminated sets ExternallyTerminated field to given value.

### HasExternallyTerminated

`func (o *HistoricProcessInstanceQueryDto) HasExternallyTerminated() bool`

HasExternallyTerminated returns a boolean if a field has been set.

### SetExternallyTerminatedNil

`func (o *HistoricProcessInstanceQueryDto) SetExternallyTerminatedNil(b bool)`

 SetExternallyTerminatedNil sets the value for ExternallyTerminated to be an explicit nil

### UnsetExternallyTerminated
`func (o *HistoricProcessInstanceQueryDto) UnsetExternallyTerminated()`

UnsetExternallyTerminated ensures that no value is present for ExternallyTerminated, not even an explicit nil
### GetInternallyTerminated

`func (o *HistoricProcessInstanceQueryDto) GetInternallyTerminated() bool`

GetInternallyTerminated returns the InternallyTerminated field if non-nil, zero value otherwise.

### GetInternallyTerminatedOk

`func (o *HistoricProcessInstanceQueryDto) GetInternallyTerminatedOk() (*bool, bool)`

GetInternallyTerminatedOk returns a tuple with the InternallyTerminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternallyTerminated

`func (o *HistoricProcessInstanceQueryDto) SetInternallyTerminated(v bool)`

SetInternallyTerminated sets InternallyTerminated field to given value.

### HasInternallyTerminated

`func (o *HistoricProcessInstanceQueryDto) HasInternallyTerminated() bool`

HasInternallyTerminated returns a boolean if a field has been set.

### SetInternallyTerminatedNil

`func (o *HistoricProcessInstanceQueryDto) SetInternallyTerminatedNil(b bool)`

 SetInternallyTerminatedNil sets the value for InternallyTerminated to be an explicit nil

### UnsetInternallyTerminated
`func (o *HistoricProcessInstanceQueryDto) UnsetInternallyTerminated()`

UnsetInternallyTerminated ensures that no value is present for InternallyTerminated, not even an explicit nil
### GetVariables

`func (o *HistoricProcessInstanceQueryDto) GetVariables() []VariableQueryParameterDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *HistoricProcessInstanceQueryDto) GetVariablesOk() (*[]VariableQueryParameterDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *HistoricProcessInstanceQueryDto) SetVariables(v []VariableQueryParameterDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *HistoricProcessInstanceQueryDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *HistoricProcessInstanceQueryDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *HistoricProcessInstanceQueryDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetVariableNamesIgnoreCase

`func (o *HistoricProcessInstanceQueryDto) GetVariableNamesIgnoreCase() bool`

GetVariableNamesIgnoreCase returns the VariableNamesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableNamesIgnoreCaseOk

`func (o *HistoricProcessInstanceQueryDto) GetVariableNamesIgnoreCaseOk() (*bool, bool)`

GetVariableNamesIgnoreCaseOk returns a tuple with the VariableNamesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNamesIgnoreCase

`func (o *HistoricProcessInstanceQueryDto) SetVariableNamesIgnoreCase(v bool)`

SetVariableNamesIgnoreCase sets VariableNamesIgnoreCase field to given value.

### HasVariableNamesIgnoreCase

`func (o *HistoricProcessInstanceQueryDto) HasVariableNamesIgnoreCase() bool`

HasVariableNamesIgnoreCase returns a boolean if a field has been set.

### SetVariableNamesIgnoreCaseNil

`func (o *HistoricProcessInstanceQueryDto) SetVariableNamesIgnoreCaseNil(b bool)`

 SetVariableNamesIgnoreCaseNil sets the value for VariableNamesIgnoreCase to be an explicit nil

### UnsetVariableNamesIgnoreCase
`func (o *HistoricProcessInstanceQueryDto) UnsetVariableNamesIgnoreCase()`

UnsetVariableNamesIgnoreCase ensures that no value is present for VariableNamesIgnoreCase, not even an explicit nil
### GetVariableValuesIgnoreCase

`func (o *HistoricProcessInstanceQueryDto) GetVariableValuesIgnoreCase() bool`

GetVariableValuesIgnoreCase returns the VariableValuesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableValuesIgnoreCaseOk

`func (o *HistoricProcessInstanceQueryDto) GetVariableValuesIgnoreCaseOk() (*bool, bool)`

GetVariableValuesIgnoreCaseOk returns a tuple with the VariableValuesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValuesIgnoreCase

`func (o *HistoricProcessInstanceQueryDto) SetVariableValuesIgnoreCase(v bool)`

SetVariableValuesIgnoreCase sets VariableValuesIgnoreCase field to given value.

### HasVariableValuesIgnoreCase

`func (o *HistoricProcessInstanceQueryDto) HasVariableValuesIgnoreCase() bool`

HasVariableValuesIgnoreCase returns a boolean if a field has been set.

### SetVariableValuesIgnoreCaseNil

`func (o *HistoricProcessInstanceQueryDto) SetVariableValuesIgnoreCaseNil(b bool)`

 SetVariableValuesIgnoreCaseNil sets the value for VariableValuesIgnoreCase to be an explicit nil

### UnsetVariableValuesIgnoreCase
`func (o *HistoricProcessInstanceQueryDto) UnsetVariableValuesIgnoreCase()`

UnsetVariableValuesIgnoreCase ensures that no value is present for VariableValuesIgnoreCase, not even an explicit nil
### GetOrQueries

`func (o *HistoricProcessInstanceQueryDto) GetOrQueries() []HistoricProcessInstanceQueryDto`

GetOrQueries returns the OrQueries field if non-nil, zero value otherwise.

### GetOrQueriesOk

`func (o *HistoricProcessInstanceQueryDto) GetOrQueriesOk() (*[]HistoricProcessInstanceQueryDto, bool)`

GetOrQueriesOk returns a tuple with the OrQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrQueries

`func (o *HistoricProcessInstanceQueryDto) SetOrQueries(v []HistoricProcessInstanceQueryDto)`

SetOrQueries sets OrQueries field to given value.

### HasOrQueries

`func (o *HistoricProcessInstanceQueryDto) HasOrQueries() bool`

HasOrQueries returns a boolean if a field has been set.

### SetOrQueriesNil

`func (o *HistoricProcessInstanceQueryDto) SetOrQueriesNil(b bool)`

 SetOrQueriesNil sets the value for OrQueries to be an explicit nil

### UnsetOrQueries
`func (o *HistoricProcessInstanceQueryDto) UnsetOrQueries()`

UnsetOrQueries ensures that no value is present for OrQueries, not even an explicit nil
### GetSorting

`func (o *HistoricProcessInstanceQueryDto) GetSorting() []HistoricProcessInstanceQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *HistoricProcessInstanceQueryDto) GetSortingOk() (*[]HistoricProcessInstanceQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *HistoricProcessInstanceQueryDto) SetSorting(v []HistoricProcessInstanceQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *HistoricProcessInstanceQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *HistoricProcessInstanceQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *HistoricProcessInstanceQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


