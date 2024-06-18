# ProcessInstanceModificationInstructionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | **Mandatory**. One of the following values: &#x60;cancel&#x60;, &#x60;startBeforeActivity&#x60;, &#x60;startAfterActivity&#x60;, &#x60;startTransition&#x60;.  * A cancel instruction requests cancellation of a single activity instance or all instances of one activity. * A startBeforeActivity instruction requests to enter a given activity. * A startAfterActivity instruction requests to execute the single outgoing sequence flow of a given activity. * A startTransition instruction requests to execute a specific sequence flow. | 
**Variables** | Pointer to [**TriggerVariableValueDto**](TriggerVariableValueDto.md) |  | [optional] 
**ActivityId** | Pointer to **NullableString** | Can be used with instructions of types &#x60;startTransition&#x60;. Specifies the sequence flow to start. | [optional] 
**TransitionId** | Pointer to **NullableString** | Can be used with instructions of types &#x60;startTransition&#x60;. Specifies the sequence flow to start. | [optional] 
**ActivityInstanceId** | Pointer to **NullableString** | Can be used with instructions of type &#x60;cancel&#x60;. Specifies the activity instance to cancel. Valid values are the activity instance IDs supplied by the [Get Activity Instance request](https://docs.camunda.org/manual/7.21/reference/rest/process-instance/get-activity-instances/). | [optional] 
**TransitionInstanceId** | Pointer to **NullableString** | Can be used with instructions of type &#x60;cancel&#x60;. Specifies the transition instance to cancel. Valid values are the transition instance IDs supplied by the [Get Activity Instance request](https://docs.camunda.org/manual/7.21/reference/rest/process-instance/get-activity-instances/). | [optional] 
**AncestorActivityInstanceId** | Pointer to **NullableString** | Can be used with instructions of type &#x60;startBeforeActivity&#x60;, &#x60;startAfterActivity&#x60;, and &#x60;startTransition&#x60;. Valid values are the activity instance IDs supplied by the Get Activity Instance request. If there are multiple parent activity instances of the targeted activity, this specifies the ancestor scope in which hierarchy the activity/transition is to be instantiated.  Example: When there are two instances of a subprocess and an activity contained in the subprocess is to be started, this parameter allows to specifiy under which subprocess instance the activity should be started. | [optional] 
**CancelCurrentActiveActivityInstances** | Pointer to **NullableBool** | Can be used with instructions of type cancel. Prevents the deletion of new created activity instances. | [optional] 

## Methods

### NewProcessInstanceModificationInstructionDto

`func NewProcessInstanceModificationInstructionDto(type_ string, ) *ProcessInstanceModificationInstructionDto`

NewProcessInstanceModificationInstructionDto instantiates a new ProcessInstanceModificationInstructionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessInstanceModificationInstructionDtoWithDefaults

`func NewProcessInstanceModificationInstructionDtoWithDefaults() *ProcessInstanceModificationInstructionDto`

NewProcessInstanceModificationInstructionDtoWithDefaults instantiates a new ProcessInstanceModificationInstructionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProcessInstanceModificationInstructionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProcessInstanceModificationInstructionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProcessInstanceModificationInstructionDto) SetType(v string)`

SetType sets Type field to given value.


### GetVariables

`func (o *ProcessInstanceModificationInstructionDto) GetVariables() TriggerVariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ProcessInstanceModificationInstructionDto) GetVariablesOk() (*TriggerVariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ProcessInstanceModificationInstructionDto) SetVariables(v TriggerVariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ProcessInstanceModificationInstructionDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetActivityId

`func (o *ProcessInstanceModificationInstructionDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *ProcessInstanceModificationInstructionDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *ProcessInstanceModificationInstructionDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *ProcessInstanceModificationInstructionDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *ProcessInstanceModificationInstructionDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *ProcessInstanceModificationInstructionDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetTransitionId

`func (o *ProcessInstanceModificationInstructionDto) GetTransitionId() string`

GetTransitionId returns the TransitionId field if non-nil, zero value otherwise.

### GetTransitionIdOk

`func (o *ProcessInstanceModificationInstructionDto) GetTransitionIdOk() (*string, bool)`

GetTransitionIdOk returns a tuple with the TransitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionId

`func (o *ProcessInstanceModificationInstructionDto) SetTransitionId(v string)`

SetTransitionId sets TransitionId field to given value.

### HasTransitionId

`func (o *ProcessInstanceModificationInstructionDto) HasTransitionId() bool`

HasTransitionId returns a boolean if a field has been set.

### SetTransitionIdNil

`func (o *ProcessInstanceModificationInstructionDto) SetTransitionIdNil(b bool)`

 SetTransitionIdNil sets the value for TransitionId to be an explicit nil

### UnsetTransitionId
`func (o *ProcessInstanceModificationInstructionDto) UnsetTransitionId()`

UnsetTransitionId ensures that no value is present for TransitionId, not even an explicit nil
### GetActivityInstanceId

`func (o *ProcessInstanceModificationInstructionDto) GetActivityInstanceId() string`

GetActivityInstanceId returns the ActivityInstanceId field if non-nil, zero value otherwise.

### GetActivityInstanceIdOk

`func (o *ProcessInstanceModificationInstructionDto) GetActivityInstanceIdOk() (*string, bool)`

GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceId

`func (o *ProcessInstanceModificationInstructionDto) SetActivityInstanceId(v string)`

SetActivityInstanceId sets ActivityInstanceId field to given value.

### HasActivityInstanceId

`func (o *ProcessInstanceModificationInstructionDto) HasActivityInstanceId() bool`

HasActivityInstanceId returns a boolean if a field has been set.

### SetActivityInstanceIdNil

`func (o *ProcessInstanceModificationInstructionDto) SetActivityInstanceIdNil(b bool)`

 SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil

### UnsetActivityInstanceId
`func (o *ProcessInstanceModificationInstructionDto) UnsetActivityInstanceId()`

UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
### GetTransitionInstanceId

`func (o *ProcessInstanceModificationInstructionDto) GetTransitionInstanceId() string`

GetTransitionInstanceId returns the TransitionInstanceId field if non-nil, zero value otherwise.

### GetTransitionInstanceIdOk

`func (o *ProcessInstanceModificationInstructionDto) GetTransitionInstanceIdOk() (*string, bool)`

GetTransitionInstanceIdOk returns a tuple with the TransitionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionInstanceId

`func (o *ProcessInstanceModificationInstructionDto) SetTransitionInstanceId(v string)`

SetTransitionInstanceId sets TransitionInstanceId field to given value.

### HasTransitionInstanceId

`func (o *ProcessInstanceModificationInstructionDto) HasTransitionInstanceId() bool`

HasTransitionInstanceId returns a boolean if a field has been set.

### SetTransitionInstanceIdNil

`func (o *ProcessInstanceModificationInstructionDto) SetTransitionInstanceIdNil(b bool)`

 SetTransitionInstanceIdNil sets the value for TransitionInstanceId to be an explicit nil

### UnsetTransitionInstanceId
`func (o *ProcessInstanceModificationInstructionDto) UnsetTransitionInstanceId()`

UnsetTransitionInstanceId ensures that no value is present for TransitionInstanceId, not even an explicit nil
### GetAncestorActivityInstanceId

`func (o *ProcessInstanceModificationInstructionDto) GetAncestorActivityInstanceId() string`

GetAncestorActivityInstanceId returns the AncestorActivityInstanceId field if non-nil, zero value otherwise.

### GetAncestorActivityInstanceIdOk

`func (o *ProcessInstanceModificationInstructionDto) GetAncestorActivityInstanceIdOk() (*string, bool)`

GetAncestorActivityInstanceIdOk returns a tuple with the AncestorActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorActivityInstanceId

`func (o *ProcessInstanceModificationInstructionDto) SetAncestorActivityInstanceId(v string)`

SetAncestorActivityInstanceId sets AncestorActivityInstanceId field to given value.

### HasAncestorActivityInstanceId

`func (o *ProcessInstanceModificationInstructionDto) HasAncestorActivityInstanceId() bool`

HasAncestorActivityInstanceId returns a boolean if a field has been set.

### SetAncestorActivityInstanceIdNil

`func (o *ProcessInstanceModificationInstructionDto) SetAncestorActivityInstanceIdNil(b bool)`

 SetAncestorActivityInstanceIdNil sets the value for AncestorActivityInstanceId to be an explicit nil

### UnsetAncestorActivityInstanceId
`func (o *ProcessInstanceModificationInstructionDto) UnsetAncestorActivityInstanceId()`

UnsetAncestorActivityInstanceId ensures that no value is present for AncestorActivityInstanceId, not even an explicit nil
### GetCancelCurrentActiveActivityInstances

`func (o *ProcessInstanceModificationInstructionDto) GetCancelCurrentActiveActivityInstances() bool`

GetCancelCurrentActiveActivityInstances returns the CancelCurrentActiveActivityInstances field if non-nil, zero value otherwise.

### GetCancelCurrentActiveActivityInstancesOk

`func (o *ProcessInstanceModificationInstructionDto) GetCancelCurrentActiveActivityInstancesOk() (*bool, bool)`

GetCancelCurrentActiveActivityInstancesOk returns a tuple with the CancelCurrentActiveActivityInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelCurrentActiveActivityInstances

`func (o *ProcessInstanceModificationInstructionDto) SetCancelCurrentActiveActivityInstances(v bool)`

SetCancelCurrentActiveActivityInstances sets CancelCurrentActiveActivityInstances field to given value.

### HasCancelCurrentActiveActivityInstances

`func (o *ProcessInstanceModificationInstructionDto) HasCancelCurrentActiveActivityInstances() bool`

HasCancelCurrentActiveActivityInstances returns a boolean if a field has been set.

### SetCancelCurrentActiveActivityInstancesNil

`func (o *ProcessInstanceModificationInstructionDto) SetCancelCurrentActiveActivityInstancesNil(b bool)`

 SetCancelCurrentActiveActivityInstancesNil sets the value for CancelCurrentActiveActivityInstances to be an explicit nil

### UnsetCancelCurrentActiveActivityInstances
`func (o *ProcessInstanceModificationInstructionDto) UnsetCancelCurrentActiveActivityInstances()`

UnsetCancelCurrentActiveActivityInstances ensures that no value is present for CancelCurrentActiveActivityInstances, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


