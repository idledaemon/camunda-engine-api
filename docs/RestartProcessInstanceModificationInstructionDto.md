# RestartProcessInstanceModificationInstructionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | **Mandatory**. One of the following values: &#x60;startBeforeActivity&#x60;, &#x60;startAfterActivity&#x60;, &#x60;startTransition&#x60;.  * A &#x60;startBeforeActivity&#x60; instruction requests to enter a given activity. * A &#x60;startAfterActivity&#x60; instruction requests to execute the single outgoing sequence flow of a given activity. * A &#x60;startTransition&#x60; instruction requests to execute a specific sequence flow. | 
**ActivityId** | Pointer to **NullableString** | **Can be used with instructions of types** &#x60;startBeforeActivity&#x60; and &#x60;startAfterActivity&#x60;. Specifies the sequence flow to start. | [optional] 
**TransitionId** | Pointer to **NullableString** | **Can be used with instructions of types** &#x60;startTransition&#x60;. Specifies the sequence flow to start. | [optional] 

## Methods

### NewRestartProcessInstanceModificationInstructionDto

`func NewRestartProcessInstanceModificationInstructionDto(type_ string, ) *RestartProcessInstanceModificationInstructionDto`

NewRestartProcessInstanceModificationInstructionDto instantiates a new RestartProcessInstanceModificationInstructionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartProcessInstanceModificationInstructionDtoWithDefaults

`func NewRestartProcessInstanceModificationInstructionDtoWithDefaults() *RestartProcessInstanceModificationInstructionDto`

NewRestartProcessInstanceModificationInstructionDtoWithDefaults instantiates a new RestartProcessInstanceModificationInstructionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RestartProcessInstanceModificationInstructionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestartProcessInstanceModificationInstructionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestartProcessInstanceModificationInstructionDto) SetType(v string)`

SetType sets Type field to given value.


### GetActivityId

`func (o *RestartProcessInstanceModificationInstructionDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *RestartProcessInstanceModificationInstructionDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *RestartProcessInstanceModificationInstructionDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *RestartProcessInstanceModificationInstructionDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *RestartProcessInstanceModificationInstructionDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *RestartProcessInstanceModificationInstructionDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetTransitionId

`func (o *RestartProcessInstanceModificationInstructionDto) GetTransitionId() string`

GetTransitionId returns the TransitionId field if non-nil, zero value otherwise.

### GetTransitionIdOk

`func (o *RestartProcessInstanceModificationInstructionDto) GetTransitionIdOk() (*string, bool)`

GetTransitionIdOk returns a tuple with the TransitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionId

`func (o *RestartProcessInstanceModificationInstructionDto) SetTransitionId(v string)`

SetTransitionId sets TransitionId field to given value.

### HasTransitionId

`func (o *RestartProcessInstanceModificationInstructionDto) HasTransitionId() bool`

HasTransitionId returns a boolean if a field has been set.

### SetTransitionIdNil

`func (o *RestartProcessInstanceModificationInstructionDto) SetTransitionIdNil(b bool)`

 SetTransitionIdNil sets the value for TransitionId to be an explicit nil

### UnsetTransitionId
`func (o *RestartProcessInstanceModificationInstructionDto) UnsetTransitionId()`

UnsetTransitionId ensures that no value is present for TransitionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


