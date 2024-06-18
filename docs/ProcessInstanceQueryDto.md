# ProcessInstanceQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **NullableString** | Filter by the deployment the id belongs to. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter by the process definition the instances run on. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Filter by the key of the process definition the instances run on. | [optional] 
**ProcessDefinitionKeyIn** | Pointer to **[]string** | Filter by a list of process definition keys. A process instance must have one of the given process definition keys. Must be a JSON array of Strings. | [optional] 
**ProcessDefinitionKeyNotIn** | Pointer to **[]string** | Exclude instances by a list of process definition keys. A process instance must not have one of the given process definition keys. Must be a JSON array of Strings. | [optional] 
**BusinessKey** | Pointer to **NullableString** | Filter by process instance business key. | [optional] 
**BusinessKeyLike** | Pointer to **NullableString** | Filter by process instance business key that the parameter is a substring of. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | Filter by case instance id. | [optional] 
**SuperProcessInstance** | Pointer to **NullableString** | Restrict query to all process instances that are sub process instances of the given process instance. Takes a process instance id. | [optional] 
**SubProcessInstance** | Pointer to **NullableString** | Restrict query to all process instances that have the given process instance as a sub process instance. Takes a process instance id. | [optional] 
**SuperCaseInstance** | Pointer to **NullableString** | Restrict query to all process instances that are sub process instances of the given case instance. Takes a case instance id. | [optional] 
**SubCaseInstance** | Pointer to **NullableString** | Restrict query to all process instances that have the given case instance as a sub case instance. Takes a case instance id. | [optional] 
**Active** | Pointer to **NullableBool** | Only include active process instances. Value may only be true, as false is the default behavior. | [optional] 
**Suspended** | Pointer to **NullableBool** | Only include suspended process instances. Value may only be true, as false is the default behavior. | [optional] 
**ProcessInstanceIds** | Pointer to **[]string** | Filter by a list of process instance ids. Must be a JSON array of Strings. | [optional] 
**WithIncident** | Pointer to **NullableBool** | Filter by presence of incidents. Selects only process instances that have an incident. | [optional] 
**IncidentId** | Pointer to **NullableString** | Filter by the incident id. | [optional] 
**IncidentType** | Pointer to **NullableString** | Filter by the incident type. See the User Guide for a list of incident types. | [optional] 
**IncidentMessage** | Pointer to **NullableString** | Filter by the incident message. Exact match. | [optional] 
**IncidentMessageLike** | Pointer to **NullableString** | Filter by the incident message that the parameter is a substring of. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Filter by a list of tenant ids. A process instance must have one of the given tenant ids. Must be a JSON array of Strings. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include process instances which belong to no tenant. Value may only be true, as false is the default behavior. | [optional] 
**ProcessDefinitionWithoutTenantId** | Pointer to **NullableBool** | Only include process instances which process definition has no tenant id. | [optional] 
**ActivityIdIn** | Pointer to **[]string** | Filter by a list of activity ids. A process instance must currently wait in a leaf activity with one of the given activity ids. | [optional] 
**RootProcessInstances** | Pointer to **NullableBool** | Restrict the query to all process instances that are top level process instances. | [optional] 
**LeafProcessInstances** | Pointer to **NullableBool** | Restrict the query to all process instances that are leaf instances. (i.e. don&#39;t have any sub instances) | [optional] 
**Variables** | Pointer to [**[]VariableQueryParameterDto**](VariableQueryParameterDto.md) | A JSON array to only include process instances that have variables with certain values. The array consists of objects with the three properties &#x60;name&#x60;, &#x60;operator&#x60; and &#x60;value&#x60;. &#x60;name&#x60; (String) is the variable name, &#x60;operator&#x60; (String) is the comparison operator to be used and &#x60;value&#x60; the variable value. The &#x60;value&#x60; may be String, Number or Boolean.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. | [optional] 
**VariableNamesIgnoreCase** | Pointer to **NullableBool** | Match all variable names in this query case-insensitively. If set to true variableName and variablename are treated as equal. | [optional] 
**VariableValuesIgnoreCase** | Pointer to **NullableBool** | Match all variable values in this query case-insensitively. If set to true variableValue and variablevalue are treated as equal. | [optional] 
**OrQueries** | Pointer to [**[]ProcessInstanceQueryDto**](ProcessInstanceQueryDto.md) | A JSON array of nested process instance queries with OR semantics. A process instance matches a nested query if it fulfills at least one of the query&#39;s predicates. With multiple nested queries, a process instance must fulfill at least one predicate of each query (Conjunctive Normal Form). All process instance query properties can be used except for: &#x60;sorting&#x60; See the [User guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/process-engine-api/#or-queries) for more information about OR queries. | [optional] 
**Sorting** | Pointer to [**[]ProcessInstanceQueryDtoSortingInner**](ProcessInstanceQueryDtoSortingInner.md) | Apply sorting of the result | [optional] 

## Methods

### NewProcessInstanceQueryDto

`func NewProcessInstanceQueryDto() *ProcessInstanceQueryDto`

NewProcessInstanceQueryDto instantiates a new ProcessInstanceQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessInstanceQueryDtoWithDefaults

`func NewProcessInstanceQueryDtoWithDefaults() *ProcessInstanceQueryDto`

NewProcessInstanceQueryDtoWithDefaults instantiates a new ProcessInstanceQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *ProcessInstanceQueryDto) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ProcessInstanceQueryDto) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ProcessInstanceQueryDto) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *ProcessInstanceQueryDto) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *ProcessInstanceQueryDto) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *ProcessInstanceQueryDto) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetProcessDefinitionId

`func (o *ProcessInstanceQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *ProcessInstanceQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *ProcessInstanceQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *ProcessInstanceQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *ProcessInstanceQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *ProcessInstanceQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *ProcessInstanceQueryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *ProcessInstanceQueryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *ProcessInstanceQueryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *ProcessInstanceQueryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *ProcessInstanceQueryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *ProcessInstanceQueryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionKeyIn

`func (o *ProcessInstanceQueryDto) GetProcessDefinitionKeyIn() []string`

GetProcessDefinitionKeyIn returns the ProcessDefinitionKeyIn field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyInOk

`func (o *ProcessInstanceQueryDto) GetProcessDefinitionKeyInOk() (*[]string, bool)`

GetProcessDefinitionKeyInOk returns a tuple with the ProcessDefinitionKeyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKeyIn

`func (o *ProcessInstanceQueryDto) SetProcessDefinitionKeyIn(v []string)`

SetProcessDefinitionKeyIn sets ProcessDefinitionKeyIn field to given value.

### HasProcessDefinitionKeyIn

`func (o *ProcessInstanceQueryDto) HasProcessDefinitionKeyIn() bool`

HasProcessDefinitionKeyIn returns a boolean if a field has been set.

### SetProcessDefinitionKeyInNil

`func (o *ProcessInstanceQueryDto) SetProcessDefinitionKeyInNil(b bool)`

 SetProcessDefinitionKeyInNil sets the value for ProcessDefinitionKeyIn to be an explicit nil

### UnsetProcessDefinitionKeyIn
`func (o *ProcessInstanceQueryDto) UnsetProcessDefinitionKeyIn()`

UnsetProcessDefinitionKeyIn ensures that no value is present for ProcessDefinitionKeyIn, not even an explicit nil
### GetProcessDefinitionKeyNotIn

`func (o *ProcessInstanceQueryDto) GetProcessDefinitionKeyNotIn() []string`

GetProcessDefinitionKeyNotIn returns the ProcessDefinitionKeyNotIn field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyNotInOk

`func (o *ProcessInstanceQueryDto) GetProcessDefinitionKeyNotInOk() (*[]string, bool)`

GetProcessDefinitionKeyNotInOk returns a tuple with the ProcessDefinitionKeyNotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKeyNotIn

`func (o *ProcessInstanceQueryDto) SetProcessDefinitionKeyNotIn(v []string)`

SetProcessDefinitionKeyNotIn sets ProcessDefinitionKeyNotIn field to given value.

### HasProcessDefinitionKeyNotIn

`func (o *ProcessInstanceQueryDto) HasProcessDefinitionKeyNotIn() bool`

HasProcessDefinitionKeyNotIn returns a boolean if a field has been set.

### SetProcessDefinitionKeyNotInNil

`func (o *ProcessInstanceQueryDto) SetProcessDefinitionKeyNotInNil(b bool)`

 SetProcessDefinitionKeyNotInNil sets the value for ProcessDefinitionKeyNotIn to be an explicit nil

### UnsetProcessDefinitionKeyNotIn
`func (o *ProcessInstanceQueryDto) UnsetProcessDefinitionKeyNotIn()`

UnsetProcessDefinitionKeyNotIn ensures that no value is present for ProcessDefinitionKeyNotIn, not even an explicit nil
### GetBusinessKey

`func (o *ProcessInstanceQueryDto) GetBusinessKey() string`

GetBusinessKey returns the BusinessKey field if non-nil, zero value otherwise.

### GetBusinessKeyOk

`func (o *ProcessInstanceQueryDto) GetBusinessKeyOk() (*string, bool)`

GetBusinessKeyOk returns a tuple with the BusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKey

`func (o *ProcessInstanceQueryDto) SetBusinessKey(v string)`

SetBusinessKey sets BusinessKey field to given value.

### HasBusinessKey

`func (o *ProcessInstanceQueryDto) HasBusinessKey() bool`

HasBusinessKey returns a boolean if a field has been set.

### SetBusinessKeyNil

`func (o *ProcessInstanceQueryDto) SetBusinessKeyNil(b bool)`

 SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil

### UnsetBusinessKey
`func (o *ProcessInstanceQueryDto) UnsetBusinessKey()`

UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
### GetBusinessKeyLike

`func (o *ProcessInstanceQueryDto) GetBusinessKeyLike() string`

GetBusinessKeyLike returns the BusinessKeyLike field if non-nil, zero value otherwise.

### GetBusinessKeyLikeOk

`func (o *ProcessInstanceQueryDto) GetBusinessKeyLikeOk() (*string, bool)`

GetBusinessKeyLikeOk returns a tuple with the BusinessKeyLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKeyLike

`func (o *ProcessInstanceQueryDto) SetBusinessKeyLike(v string)`

SetBusinessKeyLike sets BusinessKeyLike field to given value.

### HasBusinessKeyLike

`func (o *ProcessInstanceQueryDto) HasBusinessKeyLike() bool`

HasBusinessKeyLike returns a boolean if a field has been set.

### SetBusinessKeyLikeNil

`func (o *ProcessInstanceQueryDto) SetBusinessKeyLikeNil(b bool)`

 SetBusinessKeyLikeNil sets the value for BusinessKeyLike to be an explicit nil

### UnsetBusinessKeyLike
`func (o *ProcessInstanceQueryDto) UnsetBusinessKeyLike()`

UnsetBusinessKeyLike ensures that no value is present for BusinessKeyLike, not even an explicit nil
### GetCaseInstanceId

`func (o *ProcessInstanceQueryDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *ProcessInstanceQueryDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *ProcessInstanceQueryDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *ProcessInstanceQueryDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *ProcessInstanceQueryDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *ProcessInstanceQueryDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetSuperProcessInstance

`func (o *ProcessInstanceQueryDto) GetSuperProcessInstance() string`

GetSuperProcessInstance returns the SuperProcessInstance field if non-nil, zero value otherwise.

### GetSuperProcessInstanceOk

`func (o *ProcessInstanceQueryDto) GetSuperProcessInstanceOk() (*string, bool)`

GetSuperProcessInstanceOk returns a tuple with the SuperProcessInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperProcessInstance

`func (o *ProcessInstanceQueryDto) SetSuperProcessInstance(v string)`

SetSuperProcessInstance sets SuperProcessInstance field to given value.

### HasSuperProcessInstance

`func (o *ProcessInstanceQueryDto) HasSuperProcessInstance() bool`

HasSuperProcessInstance returns a boolean if a field has been set.

### SetSuperProcessInstanceNil

`func (o *ProcessInstanceQueryDto) SetSuperProcessInstanceNil(b bool)`

 SetSuperProcessInstanceNil sets the value for SuperProcessInstance to be an explicit nil

### UnsetSuperProcessInstance
`func (o *ProcessInstanceQueryDto) UnsetSuperProcessInstance()`

UnsetSuperProcessInstance ensures that no value is present for SuperProcessInstance, not even an explicit nil
### GetSubProcessInstance

`func (o *ProcessInstanceQueryDto) GetSubProcessInstance() string`

GetSubProcessInstance returns the SubProcessInstance field if non-nil, zero value otherwise.

### GetSubProcessInstanceOk

`func (o *ProcessInstanceQueryDto) GetSubProcessInstanceOk() (*string, bool)`

GetSubProcessInstanceOk returns a tuple with the SubProcessInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProcessInstance

`func (o *ProcessInstanceQueryDto) SetSubProcessInstance(v string)`

SetSubProcessInstance sets SubProcessInstance field to given value.

### HasSubProcessInstance

`func (o *ProcessInstanceQueryDto) HasSubProcessInstance() bool`

HasSubProcessInstance returns a boolean if a field has been set.

### SetSubProcessInstanceNil

`func (o *ProcessInstanceQueryDto) SetSubProcessInstanceNil(b bool)`

 SetSubProcessInstanceNil sets the value for SubProcessInstance to be an explicit nil

### UnsetSubProcessInstance
`func (o *ProcessInstanceQueryDto) UnsetSubProcessInstance()`

UnsetSubProcessInstance ensures that no value is present for SubProcessInstance, not even an explicit nil
### GetSuperCaseInstance

`func (o *ProcessInstanceQueryDto) GetSuperCaseInstance() string`

GetSuperCaseInstance returns the SuperCaseInstance field if non-nil, zero value otherwise.

### GetSuperCaseInstanceOk

`func (o *ProcessInstanceQueryDto) GetSuperCaseInstanceOk() (*string, bool)`

GetSuperCaseInstanceOk returns a tuple with the SuperCaseInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperCaseInstance

`func (o *ProcessInstanceQueryDto) SetSuperCaseInstance(v string)`

SetSuperCaseInstance sets SuperCaseInstance field to given value.

### HasSuperCaseInstance

`func (o *ProcessInstanceQueryDto) HasSuperCaseInstance() bool`

HasSuperCaseInstance returns a boolean if a field has been set.

### SetSuperCaseInstanceNil

`func (o *ProcessInstanceQueryDto) SetSuperCaseInstanceNil(b bool)`

 SetSuperCaseInstanceNil sets the value for SuperCaseInstance to be an explicit nil

### UnsetSuperCaseInstance
`func (o *ProcessInstanceQueryDto) UnsetSuperCaseInstance()`

UnsetSuperCaseInstance ensures that no value is present for SuperCaseInstance, not even an explicit nil
### GetSubCaseInstance

`func (o *ProcessInstanceQueryDto) GetSubCaseInstance() string`

GetSubCaseInstance returns the SubCaseInstance field if non-nil, zero value otherwise.

### GetSubCaseInstanceOk

`func (o *ProcessInstanceQueryDto) GetSubCaseInstanceOk() (*string, bool)`

GetSubCaseInstanceOk returns a tuple with the SubCaseInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCaseInstance

`func (o *ProcessInstanceQueryDto) SetSubCaseInstance(v string)`

SetSubCaseInstance sets SubCaseInstance field to given value.

### HasSubCaseInstance

`func (o *ProcessInstanceQueryDto) HasSubCaseInstance() bool`

HasSubCaseInstance returns a boolean if a field has been set.

### SetSubCaseInstanceNil

`func (o *ProcessInstanceQueryDto) SetSubCaseInstanceNil(b bool)`

 SetSubCaseInstanceNil sets the value for SubCaseInstance to be an explicit nil

### UnsetSubCaseInstance
`func (o *ProcessInstanceQueryDto) UnsetSubCaseInstance()`

UnsetSubCaseInstance ensures that no value is present for SubCaseInstance, not even an explicit nil
### GetActive

`func (o *ProcessInstanceQueryDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ProcessInstanceQueryDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ProcessInstanceQueryDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ProcessInstanceQueryDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *ProcessInstanceQueryDto) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *ProcessInstanceQueryDto) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetSuspended

`func (o *ProcessInstanceQueryDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ProcessInstanceQueryDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ProcessInstanceQueryDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ProcessInstanceQueryDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *ProcessInstanceQueryDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *ProcessInstanceQueryDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetProcessInstanceIds

`func (o *ProcessInstanceQueryDto) GetProcessInstanceIds() []string`

GetProcessInstanceIds returns the ProcessInstanceIds field if non-nil, zero value otherwise.

### GetProcessInstanceIdsOk

`func (o *ProcessInstanceQueryDto) GetProcessInstanceIdsOk() (*[]string, bool)`

GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIds

`func (o *ProcessInstanceQueryDto) SetProcessInstanceIds(v []string)`

SetProcessInstanceIds sets ProcessInstanceIds field to given value.

### HasProcessInstanceIds

`func (o *ProcessInstanceQueryDto) HasProcessInstanceIds() bool`

HasProcessInstanceIds returns a boolean if a field has been set.

### SetProcessInstanceIdsNil

`func (o *ProcessInstanceQueryDto) SetProcessInstanceIdsNil(b bool)`

 SetProcessInstanceIdsNil sets the value for ProcessInstanceIds to be an explicit nil

### UnsetProcessInstanceIds
`func (o *ProcessInstanceQueryDto) UnsetProcessInstanceIds()`

UnsetProcessInstanceIds ensures that no value is present for ProcessInstanceIds, not even an explicit nil
### GetWithIncident

`func (o *ProcessInstanceQueryDto) GetWithIncident() bool`

GetWithIncident returns the WithIncident field if non-nil, zero value otherwise.

### GetWithIncidentOk

`func (o *ProcessInstanceQueryDto) GetWithIncidentOk() (*bool, bool)`

GetWithIncidentOk returns a tuple with the WithIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithIncident

`func (o *ProcessInstanceQueryDto) SetWithIncident(v bool)`

SetWithIncident sets WithIncident field to given value.

### HasWithIncident

`func (o *ProcessInstanceQueryDto) HasWithIncident() bool`

HasWithIncident returns a boolean if a field has been set.

### SetWithIncidentNil

`func (o *ProcessInstanceQueryDto) SetWithIncidentNil(b bool)`

 SetWithIncidentNil sets the value for WithIncident to be an explicit nil

### UnsetWithIncident
`func (o *ProcessInstanceQueryDto) UnsetWithIncident()`

UnsetWithIncident ensures that no value is present for WithIncident, not even an explicit nil
### GetIncidentId

`func (o *ProcessInstanceQueryDto) GetIncidentId() string`

GetIncidentId returns the IncidentId field if non-nil, zero value otherwise.

### GetIncidentIdOk

`func (o *ProcessInstanceQueryDto) GetIncidentIdOk() (*string, bool)`

GetIncidentIdOk returns a tuple with the IncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentId

`func (o *ProcessInstanceQueryDto) SetIncidentId(v string)`

SetIncidentId sets IncidentId field to given value.

### HasIncidentId

`func (o *ProcessInstanceQueryDto) HasIncidentId() bool`

HasIncidentId returns a boolean if a field has been set.

### SetIncidentIdNil

`func (o *ProcessInstanceQueryDto) SetIncidentIdNil(b bool)`

 SetIncidentIdNil sets the value for IncidentId to be an explicit nil

### UnsetIncidentId
`func (o *ProcessInstanceQueryDto) UnsetIncidentId()`

UnsetIncidentId ensures that no value is present for IncidentId, not even an explicit nil
### GetIncidentType

`func (o *ProcessInstanceQueryDto) GetIncidentType() string`

GetIncidentType returns the IncidentType field if non-nil, zero value otherwise.

### GetIncidentTypeOk

`func (o *ProcessInstanceQueryDto) GetIncidentTypeOk() (*string, bool)`

GetIncidentTypeOk returns a tuple with the IncidentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentType

`func (o *ProcessInstanceQueryDto) SetIncidentType(v string)`

SetIncidentType sets IncidentType field to given value.

### HasIncidentType

`func (o *ProcessInstanceQueryDto) HasIncidentType() bool`

HasIncidentType returns a boolean if a field has been set.

### SetIncidentTypeNil

`func (o *ProcessInstanceQueryDto) SetIncidentTypeNil(b bool)`

 SetIncidentTypeNil sets the value for IncidentType to be an explicit nil

### UnsetIncidentType
`func (o *ProcessInstanceQueryDto) UnsetIncidentType()`

UnsetIncidentType ensures that no value is present for IncidentType, not even an explicit nil
### GetIncidentMessage

`func (o *ProcessInstanceQueryDto) GetIncidentMessage() string`

GetIncidentMessage returns the IncidentMessage field if non-nil, zero value otherwise.

### GetIncidentMessageOk

`func (o *ProcessInstanceQueryDto) GetIncidentMessageOk() (*string, bool)`

GetIncidentMessageOk returns a tuple with the IncidentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentMessage

`func (o *ProcessInstanceQueryDto) SetIncidentMessage(v string)`

SetIncidentMessage sets IncidentMessage field to given value.

### HasIncidentMessage

`func (o *ProcessInstanceQueryDto) HasIncidentMessage() bool`

HasIncidentMessage returns a boolean if a field has been set.

### SetIncidentMessageNil

`func (o *ProcessInstanceQueryDto) SetIncidentMessageNil(b bool)`

 SetIncidentMessageNil sets the value for IncidentMessage to be an explicit nil

### UnsetIncidentMessage
`func (o *ProcessInstanceQueryDto) UnsetIncidentMessage()`

UnsetIncidentMessage ensures that no value is present for IncidentMessage, not even an explicit nil
### GetIncidentMessageLike

`func (o *ProcessInstanceQueryDto) GetIncidentMessageLike() string`

GetIncidentMessageLike returns the IncidentMessageLike field if non-nil, zero value otherwise.

### GetIncidentMessageLikeOk

`func (o *ProcessInstanceQueryDto) GetIncidentMessageLikeOk() (*string, bool)`

GetIncidentMessageLikeOk returns a tuple with the IncidentMessageLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentMessageLike

`func (o *ProcessInstanceQueryDto) SetIncidentMessageLike(v string)`

SetIncidentMessageLike sets IncidentMessageLike field to given value.

### HasIncidentMessageLike

`func (o *ProcessInstanceQueryDto) HasIncidentMessageLike() bool`

HasIncidentMessageLike returns a boolean if a field has been set.

### SetIncidentMessageLikeNil

`func (o *ProcessInstanceQueryDto) SetIncidentMessageLikeNil(b bool)`

 SetIncidentMessageLikeNil sets the value for IncidentMessageLike to be an explicit nil

### UnsetIncidentMessageLike
`func (o *ProcessInstanceQueryDto) UnsetIncidentMessageLike()`

UnsetIncidentMessageLike ensures that no value is present for IncidentMessageLike, not even an explicit nil
### GetTenantIdIn

`func (o *ProcessInstanceQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *ProcessInstanceQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *ProcessInstanceQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *ProcessInstanceQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *ProcessInstanceQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *ProcessInstanceQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *ProcessInstanceQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *ProcessInstanceQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *ProcessInstanceQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *ProcessInstanceQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *ProcessInstanceQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *ProcessInstanceQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetProcessDefinitionWithoutTenantId

`func (o *ProcessInstanceQueryDto) GetProcessDefinitionWithoutTenantId() bool`

GetProcessDefinitionWithoutTenantId returns the ProcessDefinitionWithoutTenantId field if non-nil, zero value otherwise.

### GetProcessDefinitionWithoutTenantIdOk

`func (o *ProcessInstanceQueryDto) GetProcessDefinitionWithoutTenantIdOk() (*bool, bool)`

GetProcessDefinitionWithoutTenantIdOk returns a tuple with the ProcessDefinitionWithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionWithoutTenantId

`func (o *ProcessInstanceQueryDto) SetProcessDefinitionWithoutTenantId(v bool)`

SetProcessDefinitionWithoutTenantId sets ProcessDefinitionWithoutTenantId field to given value.

### HasProcessDefinitionWithoutTenantId

`func (o *ProcessInstanceQueryDto) HasProcessDefinitionWithoutTenantId() bool`

HasProcessDefinitionWithoutTenantId returns a boolean if a field has been set.

### SetProcessDefinitionWithoutTenantIdNil

`func (o *ProcessInstanceQueryDto) SetProcessDefinitionWithoutTenantIdNil(b bool)`

 SetProcessDefinitionWithoutTenantIdNil sets the value for ProcessDefinitionWithoutTenantId to be an explicit nil

### UnsetProcessDefinitionWithoutTenantId
`func (o *ProcessInstanceQueryDto) UnsetProcessDefinitionWithoutTenantId()`

UnsetProcessDefinitionWithoutTenantId ensures that no value is present for ProcessDefinitionWithoutTenantId, not even an explicit nil
### GetActivityIdIn

`func (o *ProcessInstanceQueryDto) GetActivityIdIn() []string`

GetActivityIdIn returns the ActivityIdIn field if non-nil, zero value otherwise.

### GetActivityIdInOk

`func (o *ProcessInstanceQueryDto) GetActivityIdInOk() (*[]string, bool)`

GetActivityIdInOk returns a tuple with the ActivityIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIdIn

`func (o *ProcessInstanceQueryDto) SetActivityIdIn(v []string)`

SetActivityIdIn sets ActivityIdIn field to given value.

### HasActivityIdIn

`func (o *ProcessInstanceQueryDto) HasActivityIdIn() bool`

HasActivityIdIn returns a boolean if a field has been set.

### SetActivityIdInNil

`func (o *ProcessInstanceQueryDto) SetActivityIdInNil(b bool)`

 SetActivityIdInNil sets the value for ActivityIdIn to be an explicit nil

### UnsetActivityIdIn
`func (o *ProcessInstanceQueryDto) UnsetActivityIdIn()`

UnsetActivityIdIn ensures that no value is present for ActivityIdIn, not even an explicit nil
### GetRootProcessInstances

`func (o *ProcessInstanceQueryDto) GetRootProcessInstances() bool`

GetRootProcessInstances returns the RootProcessInstances field if non-nil, zero value otherwise.

### GetRootProcessInstancesOk

`func (o *ProcessInstanceQueryDto) GetRootProcessInstancesOk() (*bool, bool)`

GetRootProcessInstancesOk returns a tuple with the RootProcessInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstances

`func (o *ProcessInstanceQueryDto) SetRootProcessInstances(v bool)`

SetRootProcessInstances sets RootProcessInstances field to given value.

### HasRootProcessInstances

`func (o *ProcessInstanceQueryDto) HasRootProcessInstances() bool`

HasRootProcessInstances returns a boolean if a field has been set.

### SetRootProcessInstancesNil

`func (o *ProcessInstanceQueryDto) SetRootProcessInstancesNil(b bool)`

 SetRootProcessInstancesNil sets the value for RootProcessInstances to be an explicit nil

### UnsetRootProcessInstances
`func (o *ProcessInstanceQueryDto) UnsetRootProcessInstances()`

UnsetRootProcessInstances ensures that no value is present for RootProcessInstances, not even an explicit nil
### GetLeafProcessInstances

`func (o *ProcessInstanceQueryDto) GetLeafProcessInstances() bool`

GetLeafProcessInstances returns the LeafProcessInstances field if non-nil, zero value otherwise.

### GetLeafProcessInstancesOk

`func (o *ProcessInstanceQueryDto) GetLeafProcessInstancesOk() (*bool, bool)`

GetLeafProcessInstancesOk returns a tuple with the LeafProcessInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafProcessInstances

`func (o *ProcessInstanceQueryDto) SetLeafProcessInstances(v bool)`

SetLeafProcessInstances sets LeafProcessInstances field to given value.

### HasLeafProcessInstances

`func (o *ProcessInstanceQueryDto) HasLeafProcessInstances() bool`

HasLeafProcessInstances returns a boolean if a field has been set.

### SetLeafProcessInstancesNil

`func (o *ProcessInstanceQueryDto) SetLeafProcessInstancesNil(b bool)`

 SetLeafProcessInstancesNil sets the value for LeafProcessInstances to be an explicit nil

### UnsetLeafProcessInstances
`func (o *ProcessInstanceQueryDto) UnsetLeafProcessInstances()`

UnsetLeafProcessInstances ensures that no value is present for LeafProcessInstances, not even an explicit nil
### GetVariables

`func (o *ProcessInstanceQueryDto) GetVariables() []VariableQueryParameterDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ProcessInstanceQueryDto) GetVariablesOk() (*[]VariableQueryParameterDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ProcessInstanceQueryDto) SetVariables(v []VariableQueryParameterDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ProcessInstanceQueryDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *ProcessInstanceQueryDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *ProcessInstanceQueryDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetVariableNamesIgnoreCase

`func (o *ProcessInstanceQueryDto) GetVariableNamesIgnoreCase() bool`

GetVariableNamesIgnoreCase returns the VariableNamesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableNamesIgnoreCaseOk

`func (o *ProcessInstanceQueryDto) GetVariableNamesIgnoreCaseOk() (*bool, bool)`

GetVariableNamesIgnoreCaseOk returns a tuple with the VariableNamesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNamesIgnoreCase

`func (o *ProcessInstanceQueryDto) SetVariableNamesIgnoreCase(v bool)`

SetVariableNamesIgnoreCase sets VariableNamesIgnoreCase field to given value.

### HasVariableNamesIgnoreCase

`func (o *ProcessInstanceQueryDto) HasVariableNamesIgnoreCase() bool`

HasVariableNamesIgnoreCase returns a boolean if a field has been set.

### SetVariableNamesIgnoreCaseNil

`func (o *ProcessInstanceQueryDto) SetVariableNamesIgnoreCaseNil(b bool)`

 SetVariableNamesIgnoreCaseNil sets the value for VariableNamesIgnoreCase to be an explicit nil

### UnsetVariableNamesIgnoreCase
`func (o *ProcessInstanceQueryDto) UnsetVariableNamesIgnoreCase()`

UnsetVariableNamesIgnoreCase ensures that no value is present for VariableNamesIgnoreCase, not even an explicit nil
### GetVariableValuesIgnoreCase

`func (o *ProcessInstanceQueryDto) GetVariableValuesIgnoreCase() bool`

GetVariableValuesIgnoreCase returns the VariableValuesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableValuesIgnoreCaseOk

`func (o *ProcessInstanceQueryDto) GetVariableValuesIgnoreCaseOk() (*bool, bool)`

GetVariableValuesIgnoreCaseOk returns a tuple with the VariableValuesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValuesIgnoreCase

`func (o *ProcessInstanceQueryDto) SetVariableValuesIgnoreCase(v bool)`

SetVariableValuesIgnoreCase sets VariableValuesIgnoreCase field to given value.

### HasVariableValuesIgnoreCase

`func (o *ProcessInstanceQueryDto) HasVariableValuesIgnoreCase() bool`

HasVariableValuesIgnoreCase returns a boolean if a field has been set.

### SetVariableValuesIgnoreCaseNil

`func (o *ProcessInstanceQueryDto) SetVariableValuesIgnoreCaseNil(b bool)`

 SetVariableValuesIgnoreCaseNil sets the value for VariableValuesIgnoreCase to be an explicit nil

### UnsetVariableValuesIgnoreCase
`func (o *ProcessInstanceQueryDto) UnsetVariableValuesIgnoreCase()`

UnsetVariableValuesIgnoreCase ensures that no value is present for VariableValuesIgnoreCase, not even an explicit nil
### GetOrQueries

`func (o *ProcessInstanceQueryDto) GetOrQueries() []ProcessInstanceQueryDto`

GetOrQueries returns the OrQueries field if non-nil, zero value otherwise.

### GetOrQueriesOk

`func (o *ProcessInstanceQueryDto) GetOrQueriesOk() (*[]ProcessInstanceQueryDto, bool)`

GetOrQueriesOk returns a tuple with the OrQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrQueries

`func (o *ProcessInstanceQueryDto) SetOrQueries(v []ProcessInstanceQueryDto)`

SetOrQueries sets OrQueries field to given value.

### HasOrQueries

`func (o *ProcessInstanceQueryDto) HasOrQueries() bool`

HasOrQueries returns a boolean if a field has been set.

### SetOrQueriesNil

`func (o *ProcessInstanceQueryDto) SetOrQueriesNil(b bool)`

 SetOrQueriesNil sets the value for OrQueries to be an explicit nil

### UnsetOrQueries
`func (o *ProcessInstanceQueryDto) UnsetOrQueries()`

UnsetOrQueries ensures that no value is present for OrQueries, not even an explicit nil
### GetSorting

`func (o *ProcessInstanceQueryDto) GetSorting() []ProcessInstanceQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *ProcessInstanceQueryDto) GetSortingOk() (*[]ProcessInstanceQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *ProcessInstanceQueryDto) SetSorting(v []ProcessInstanceQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *ProcessInstanceQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *ProcessInstanceQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *ProcessInstanceQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


