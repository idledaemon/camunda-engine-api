# TransitionInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the transition instance. | [optional] 
**ParentActivityInstanceId** | Pointer to **NullableString** | The id of the parent activity instance, for example a sub process instance. | [optional] 
**ActivityId** | Pointer to **NullableString** | The id of the activity that this instance enters (asyncBefore job) or leaves (asyncAfter job) | [optional] 
**ActivityName** | Pointer to **NullableString** | The name of the activity that this instance enters (asyncBefore job) or leaves (asyncAfter job) | [optional] 
**ActivityType** | Pointer to **NullableString** | The type of the activity that this instance enters (asyncBefore job) or leaves (asyncAfter job) | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance this instance is part of. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The execution id. | [optional] 
**IncidentIds** | Pointer to **[]string** | A list of incident ids. | [optional] 
**Incidents** | Pointer to [**[]ActivityInstanceIncidentDto**](ActivityInstanceIncidentDto.md) | A list of JSON objects containing incident specific properties: * &#x60;id&#x60;: the id of the incident * &#x60;activityId&#x60;: the activity id in which the incident occurred | [optional] 

## Methods

### NewTransitionInstanceDto

`func NewTransitionInstanceDto() *TransitionInstanceDto`

NewTransitionInstanceDto instantiates a new TransitionInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionInstanceDtoWithDefaults

`func NewTransitionInstanceDtoWithDefaults() *TransitionInstanceDto`

NewTransitionInstanceDtoWithDefaults instantiates a new TransitionInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransitionInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransitionInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransitionInstanceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransitionInstanceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TransitionInstanceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TransitionInstanceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetParentActivityInstanceId

`func (o *TransitionInstanceDto) GetParentActivityInstanceId() string`

GetParentActivityInstanceId returns the ParentActivityInstanceId field if non-nil, zero value otherwise.

### GetParentActivityInstanceIdOk

`func (o *TransitionInstanceDto) GetParentActivityInstanceIdOk() (*string, bool)`

GetParentActivityInstanceIdOk returns a tuple with the ParentActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentActivityInstanceId

`func (o *TransitionInstanceDto) SetParentActivityInstanceId(v string)`

SetParentActivityInstanceId sets ParentActivityInstanceId field to given value.

### HasParentActivityInstanceId

`func (o *TransitionInstanceDto) HasParentActivityInstanceId() bool`

HasParentActivityInstanceId returns a boolean if a field has been set.

### SetParentActivityInstanceIdNil

`func (o *TransitionInstanceDto) SetParentActivityInstanceIdNil(b bool)`

 SetParentActivityInstanceIdNil sets the value for ParentActivityInstanceId to be an explicit nil

### UnsetParentActivityInstanceId
`func (o *TransitionInstanceDto) UnsetParentActivityInstanceId()`

UnsetParentActivityInstanceId ensures that no value is present for ParentActivityInstanceId, not even an explicit nil
### GetActivityId

`func (o *TransitionInstanceDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *TransitionInstanceDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *TransitionInstanceDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *TransitionInstanceDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *TransitionInstanceDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *TransitionInstanceDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetActivityName

`func (o *TransitionInstanceDto) GetActivityName() string`

GetActivityName returns the ActivityName field if non-nil, zero value otherwise.

### GetActivityNameOk

`func (o *TransitionInstanceDto) GetActivityNameOk() (*string, bool)`

GetActivityNameOk returns a tuple with the ActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityName

`func (o *TransitionInstanceDto) SetActivityName(v string)`

SetActivityName sets ActivityName field to given value.

### HasActivityName

`func (o *TransitionInstanceDto) HasActivityName() bool`

HasActivityName returns a boolean if a field has been set.

### SetActivityNameNil

`func (o *TransitionInstanceDto) SetActivityNameNil(b bool)`

 SetActivityNameNil sets the value for ActivityName to be an explicit nil

### UnsetActivityName
`func (o *TransitionInstanceDto) UnsetActivityName()`

UnsetActivityName ensures that no value is present for ActivityName, not even an explicit nil
### GetActivityType

`func (o *TransitionInstanceDto) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *TransitionInstanceDto) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *TransitionInstanceDto) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *TransitionInstanceDto) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### SetActivityTypeNil

`func (o *TransitionInstanceDto) SetActivityTypeNil(b bool)`

 SetActivityTypeNil sets the value for ActivityType to be an explicit nil

### UnsetActivityType
`func (o *TransitionInstanceDto) UnsetActivityType()`

UnsetActivityType ensures that no value is present for ActivityType, not even an explicit nil
### GetProcessInstanceId

`func (o *TransitionInstanceDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *TransitionInstanceDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *TransitionInstanceDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *TransitionInstanceDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *TransitionInstanceDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *TransitionInstanceDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessDefinitionId

`func (o *TransitionInstanceDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *TransitionInstanceDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *TransitionInstanceDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *TransitionInstanceDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *TransitionInstanceDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *TransitionInstanceDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetExecutionId

`func (o *TransitionInstanceDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *TransitionInstanceDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *TransitionInstanceDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *TransitionInstanceDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *TransitionInstanceDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *TransitionInstanceDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetIncidentIds

`func (o *TransitionInstanceDto) GetIncidentIds() []string`

GetIncidentIds returns the IncidentIds field if non-nil, zero value otherwise.

### GetIncidentIdsOk

`func (o *TransitionInstanceDto) GetIncidentIdsOk() (*[]string, bool)`

GetIncidentIdsOk returns a tuple with the IncidentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentIds

`func (o *TransitionInstanceDto) SetIncidentIds(v []string)`

SetIncidentIds sets IncidentIds field to given value.

### HasIncidentIds

`func (o *TransitionInstanceDto) HasIncidentIds() bool`

HasIncidentIds returns a boolean if a field has been set.

### SetIncidentIdsNil

`func (o *TransitionInstanceDto) SetIncidentIdsNil(b bool)`

 SetIncidentIdsNil sets the value for IncidentIds to be an explicit nil

### UnsetIncidentIds
`func (o *TransitionInstanceDto) UnsetIncidentIds()`

UnsetIncidentIds ensures that no value is present for IncidentIds, not even an explicit nil
### GetIncidents

`func (o *TransitionInstanceDto) GetIncidents() []ActivityInstanceIncidentDto`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *TransitionInstanceDto) GetIncidentsOk() (*[]ActivityInstanceIncidentDto, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *TransitionInstanceDto) SetIncidents(v []ActivityInstanceIncidentDto)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *TransitionInstanceDto) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### SetIncidentsNil

`func (o *TransitionInstanceDto) SetIncidentsNil(b bool)`

 SetIncidentsNil sets the value for Incidents to be an explicit nil

### UnsetIncidents
`func (o *TransitionInstanceDto) UnsetIncidents()`

UnsetIncidents ensures that no value is present for Incidents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


