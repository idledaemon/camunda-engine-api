# ActivityInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the activity instance. | [optional] 
**ParentActivityInstanceId** | Pointer to **NullableString** | The id of the parent activity instance, for example a sub process instance. | [optional] 
**ActivityId** | Pointer to **NullableString** | The id of the activity. | [optional] 
**ActivityName** | Pointer to **NullableString** | The name of the activity | [optional] 
**Name** | Pointer to **NullableString** | The name of the activity. This property is deprecated. Please use &#39;activityName&#39;. | [optional] 
**ActivityType** | Pointer to **NullableString** | The type of activity (corresponds to the XML element name in the BPMN 2.0, e.g., &#39;userTask&#39;) | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance this activity instance is part of. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition. | [optional] 
**ChildActivityInstances** | Pointer to [**[]ActivityInstanceDto**](ActivityInstanceDto.md) | A list of child activity instances. | [optional] 
**ChildTransitionInstances** | Pointer to [**[]TransitionInstanceDto**](TransitionInstanceDto.md) | A list of child transition instances. A transition instance represents an execution waiting in an asynchronous continuation. | [optional] 
**ExecutionIds** | Pointer to **[]string** | A list of execution ids. | [optional] 
**IncidentIds** | Pointer to **[]string** | A list of incident ids. | [optional] 
**Incidents** | Pointer to [**[]ActivityInstanceIncidentDto**](ActivityInstanceIncidentDto.md) | A list of JSON objects containing incident specific properties: * &#x60;id&#x60;: the id of the incident * &#x60;activityId&#x60;: the activity id in which the incident occurred | [optional] 

## Methods

### NewActivityInstanceDto

`func NewActivityInstanceDto() *ActivityInstanceDto`

NewActivityInstanceDto instantiates a new ActivityInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityInstanceDtoWithDefaults

`func NewActivityInstanceDtoWithDefaults() *ActivityInstanceDto`

NewActivityInstanceDtoWithDefaults instantiates a new ActivityInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityInstanceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityInstanceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ActivityInstanceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ActivityInstanceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetParentActivityInstanceId

`func (o *ActivityInstanceDto) GetParentActivityInstanceId() string`

GetParentActivityInstanceId returns the ParentActivityInstanceId field if non-nil, zero value otherwise.

### GetParentActivityInstanceIdOk

`func (o *ActivityInstanceDto) GetParentActivityInstanceIdOk() (*string, bool)`

GetParentActivityInstanceIdOk returns a tuple with the ParentActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentActivityInstanceId

`func (o *ActivityInstanceDto) SetParentActivityInstanceId(v string)`

SetParentActivityInstanceId sets ParentActivityInstanceId field to given value.

### HasParentActivityInstanceId

`func (o *ActivityInstanceDto) HasParentActivityInstanceId() bool`

HasParentActivityInstanceId returns a boolean if a field has been set.

### SetParentActivityInstanceIdNil

`func (o *ActivityInstanceDto) SetParentActivityInstanceIdNil(b bool)`

 SetParentActivityInstanceIdNil sets the value for ParentActivityInstanceId to be an explicit nil

### UnsetParentActivityInstanceId
`func (o *ActivityInstanceDto) UnsetParentActivityInstanceId()`

UnsetParentActivityInstanceId ensures that no value is present for ParentActivityInstanceId, not even an explicit nil
### GetActivityId

`func (o *ActivityInstanceDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *ActivityInstanceDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *ActivityInstanceDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *ActivityInstanceDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *ActivityInstanceDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *ActivityInstanceDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetActivityName

`func (o *ActivityInstanceDto) GetActivityName() string`

GetActivityName returns the ActivityName field if non-nil, zero value otherwise.

### GetActivityNameOk

`func (o *ActivityInstanceDto) GetActivityNameOk() (*string, bool)`

GetActivityNameOk returns a tuple with the ActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityName

`func (o *ActivityInstanceDto) SetActivityName(v string)`

SetActivityName sets ActivityName field to given value.

### HasActivityName

`func (o *ActivityInstanceDto) HasActivityName() bool`

HasActivityName returns a boolean if a field has been set.

### SetActivityNameNil

`func (o *ActivityInstanceDto) SetActivityNameNil(b bool)`

 SetActivityNameNil sets the value for ActivityName to be an explicit nil

### UnsetActivityName
`func (o *ActivityInstanceDto) UnsetActivityName()`

UnsetActivityName ensures that no value is present for ActivityName, not even an explicit nil
### GetName

`func (o *ActivityInstanceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityInstanceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityInstanceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActivityInstanceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ActivityInstanceDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ActivityInstanceDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetActivityType

`func (o *ActivityInstanceDto) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *ActivityInstanceDto) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *ActivityInstanceDto) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *ActivityInstanceDto) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### SetActivityTypeNil

`func (o *ActivityInstanceDto) SetActivityTypeNil(b bool)`

 SetActivityTypeNil sets the value for ActivityType to be an explicit nil

### UnsetActivityType
`func (o *ActivityInstanceDto) UnsetActivityType()`

UnsetActivityType ensures that no value is present for ActivityType, not even an explicit nil
### GetProcessInstanceId

`func (o *ActivityInstanceDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *ActivityInstanceDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *ActivityInstanceDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *ActivityInstanceDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *ActivityInstanceDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *ActivityInstanceDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessDefinitionId

`func (o *ActivityInstanceDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *ActivityInstanceDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *ActivityInstanceDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *ActivityInstanceDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *ActivityInstanceDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *ActivityInstanceDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetChildActivityInstances

`func (o *ActivityInstanceDto) GetChildActivityInstances() []ActivityInstanceDto`

GetChildActivityInstances returns the ChildActivityInstances field if non-nil, zero value otherwise.

### GetChildActivityInstancesOk

`func (o *ActivityInstanceDto) GetChildActivityInstancesOk() (*[]ActivityInstanceDto, bool)`

GetChildActivityInstancesOk returns a tuple with the ChildActivityInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildActivityInstances

`func (o *ActivityInstanceDto) SetChildActivityInstances(v []ActivityInstanceDto)`

SetChildActivityInstances sets ChildActivityInstances field to given value.

### HasChildActivityInstances

`func (o *ActivityInstanceDto) HasChildActivityInstances() bool`

HasChildActivityInstances returns a boolean if a field has been set.

### SetChildActivityInstancesNil

`func (o *ActivityInstanceDto) SetChildActivityInstancesNil(b bool)`

 SetChildActivityInstancesNil sets the value for ChildActivityInstances to be an explicit nil

### UnsetChildActivityInstances
`func (o *ActivityInstanceDto) UnsetChildActivityInstances()`

UnsetChildActivityInstances ensures that no value is present for ChildActivityInstances, not even an explicit nil
### GetChildTransitionInstances

`func (o *ActivityInstanceDto) GetChildTransitionInstances() []TransitionInstanceDto`

GetChildTransitionInstances returns the ChildTransitionInstances field if non-nil, zero value otherwise.

### GetChildTransitionInstancesOk

`func (o *ActivityInstanceDto) GetChildTransitionInstancesOk() (*[]TransitionInstanceDto, bool)`

GetChildTransitionInstancesOk returns a tuple with the ChildTransitionInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTransitionInstances

`func (o *ActivityInstanceDto) SetChildTransitionInstances(v []TransitionInstanceDto)`

SetChildTransitionInstances sets ChildTransitionInstances field to given value.

### HasChildTransitionInstances

`func (o *ActivityInstanceDto) HasChildTransitionInstances() bool`

HasChildTransitionInstances returns a boolean if a field has been set.

### SetChildTransitionInstancesNil

`func (o *ActivityInstanceDto) SetChildTransitionInstancesNil(b bool)`

 SetChildTransitionInstancesNil sets the value for ChildTransitionInstances to be an explicit nil

### UnsetChildTransitionInstances
`func (o *ActivityInstanceDto) UnsetChildTransitionInstances()`

UnsetChildTransitionInstances ensures that no value is present for ChildTransitionInstances, not even an explicit nil
### GetExecutionIds

`func (o *ActivityInstanceDto) GetExecutionIds() []string`

GetExecutionIds returns the ExecutionIds field if non-nil, zero value otherwise.

### GetExecutionIdsOk

`func (o *ActivityInstanceDto) GetExecutionIdsOk() (*[]string, bool)`

GetExecutionIdsOk returns a tuple with the ExecutionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionIds

`func (o *ActivityInstanceDto) SetExecutionIds(v []string)`

SetExecutionIds sets ExecutionIds field to given value.

### HasExecutionIds

`func (o *ActivityInstanceDto) HasExecutionIds() bool`

HasExecutionIds returns a boolean if a field has been set.

### SetExecutionIdsNil

`func (o *ActivityInstanceDto) SetExecutionIdsNil(b bool)`

 SetExecutionIdsNil sets the value for ExecutionIds to be an explicit nil

### UnsetExecutionIds
`func (o *ActivityInstanceDto) UnsetExecutionIds()`

UnsetExecutionIds ensures that no value is present for ExecutionIds, not even an explicit nil
### GetIncidentIds

`func (o *ActivityInstanceDto) GetIncidentIds() []string`

GetIncidentIds returns the IncidentIds field if non-nil, zero value otherwise.

### GetIncidentIdsOk

`func (o *ActivityInstanceDto) GetIncidentIdsOk() (*[]string, bool)`

GetIncidentIdsOk returns a tuple with the IncidentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentIds

`func (o *ActivityInstanceDto) SetIncidentIds(v []string)`

SetIncidentIds sets IncidentIds field to given value.

### HasIncidentIds

`func (o *ActivityInstanceDto) HasIncidentIds() bool`

HasIncidentIds returns a boolean if a field has been set.

### SetIncidentIdsNil

`func (o *ActivityInstanceDto) SetIncidentIdsNil(b bool)`

 SetIncidentIdsNil sets the value for IncidentIds to be an explicit nil

### UnsetIncidentIds
`func (o *ActivityInstanceDto) UnsetIncidentIds()`

UnsetIncidentIds ensures that no value is present for IncidentIds, not even an explicit nil
### GetIncidents

`func (o *ActivityInstanceDto) GetIncidents() []ActivityInstanceIncidentDto`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *ActivityInstanceDto) GetIncidentsOk() (*[]ActivityInstanceIncidentDto, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *ActivityInstanceDto) SetIncidents(v []ActivityInstanceIncidentDto)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *ActivityInstanceDto) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### SetIncidentsNil

`func (o *ActivityInstanceDto) SetIncidentsNil(b bool)`

 SetIncidentsNil sets the value for Incidents to be an explicit nil

### UnsetIncidents
`func (o *ActivityInstanceDto) UnsetIncidents()`

UnsetIncidents ensures that no value is present for Incidents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


