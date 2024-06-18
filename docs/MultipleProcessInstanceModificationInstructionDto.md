# MultipleProcessInstanceModificationInstructionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | **Mandatory**. One of the following values: &#x60;cancel&#x60;, &#x60;startBeforeActivity&#x60;, &#x60;startAfterActivity&#x60;, &#x60;startTransition&#x60;.  * A cancel instruction requests cancellation of a single activity instance or all instances of one activity. * A startBeforeActivity instruction requests to enter a given activity. * A startAfterActivity instruction requests to execute the single outgoing sequence flow of a given activity. * A startTransition instruction requests to execute a specific sequence flow. | 
**ActivityId** | Pointer to **NullableString** | Can be used with instructions of types &#x60;startTransition&#x60;. Specifies the sequence flow to start. | [optional] 
**TransitionId** | Pointer to **NullableString** | Can be used with instructions of types &#x60;startTransition&#x60;. Specifies the sequence flow to start. | [optional] 
**CancelCurrentActiveActivityInstances** | Pointer to **NullableBool** | Can be used with instructions of type cancel. Prevents the deletion of new created activity instances. | [optional] 

## Methods

### NewMultipleProcessInstanceModificationInstructionDto

`func NewMultipleProcessInstanceModificationInstructionDto(type_ string, ) *MultipleProcessInstanceModificationInstructionDto`

NewMultipleProcessInstanceModificationInstructionDto instantiates a new MultipleProcessInstanceModificationInstructionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleProcessInstanceModificationInstructionDtoWithDefaults

`func NewMultipleProcessInstanceModificationInstructionDtoWithDefaults() *MultipleProcessInstanceModificationInstructionDto`

NewMultipleProcessInstanceModificationInstructionDtoWithDefaults instantiates a new MultipleProcessInstanceModificationInstructionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MultipleProcessInstanceModificationInstructionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultipleProcessInstanceModificationInstructionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultipleProcessInstanceModificationInstructionDto) SetType(v string)`

SetType sets Type field to given value.


### GetActivityId

`func (o *MultipleProcessInstanceModificationInstructionDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *MultipleProcessInstanceModificationInstructionDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *MultipleProcessInstanceModificationInstructionDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *MultipleProcessInstanceModificationInstructionDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *MultipleProcessInstanceModificationInstructionDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *MultipleProcessInstanceModificationInstructionDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetTransitionId

`func (o *MultipleProcessInstanceModificationInstructionDto) GetTransitionId() string`

GetTransitionId returns the TransitionId field if non-nil, zero value otherwise.

### GetTransitionIdOk

`func (o *MultipleProcessInstanceModificationInstructionDto) GetTransitionIdOk() (*string, bool)`

GetTransitionIdOk returns a tuple with the TransitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionId

`func (o *MultipleProcessInstanceModificationInstructionDto) SetTransitionId(v string)`

SetTransitionId sets TransitionId field to given value.

### HasTransitionId

`func (o *MultipleProcessInstanceModificationInstructionDto) HasTransitionId() bool`

HasTransitionId returns a boolean if a field has been set.

### SetTransitionIdNil

`func (o *MultipleProcessInstanceModificationInstructionDto) SetTransitionIdNil(b bool)`

 SetTransitionIdNil sets the value for TransitionId to be an explicit nil

### UnsetTransitionId
`func (o *MultipleProcessInstanceModificationInstructionDto) UnsetTransitionId()`

UnsetTransitionId ensures that no value is present for TransitionId, not even an explicit nil
### GetCancelCurrentActiveActivityInstances

`func (o *MultipleProcessInstanceModificationInstructionDto) GetCancelCurrentActiveActivityInstances() bool`

GetCancelCurrentActiveActivityInstances returns the CancelCurrentActiveActivityInstances field if non-nil, zero value otherwise.

### GetCancelCurrentActiveActivityInstancesOk

`func (o *MultipleProcessInstanceModificationInstructionDto) GetCancelCurrentActiveActivityInstancesOk() (*bool, bool)`

GetCancelCurrentActiveActivityInstancesOk returns a tuple with the CancelCurrentActiveActivityInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelCurrentActiveActivityInstances

`func (o *MultipleProcessInstanceModificationInstructionDto) SetCancelCurrentActiveActivityInstances(v bool)`

SetCancelCurrentActiveActivityInstances sets CancelCurrentActiveActivityInstances field to given value.

### HasCancelCurrentActiveActivityInstances

`func (o *MultipleProcessInstanceModificationInstructionDto) HasCancelCurrentActiveActivityInstances() bool`

HasCancelCurrentActiveActivityInstances returns a boolean if a field has been set.

### SetCancelCurrentActiveActivityInstancesNil

`func (o *MultipleProcessInstanceModificationInstructionDto) SetCancelCurrentActiveActivityInstancesNil(b bool)`

 SetCancelCurrentActiveActivityInstancesNil sets the value for CancelCurrentActiveActivityInstances to be an explicit nil

### UnsetCancelCurrentActiveActivityInstances
`func (o *MultipleProcessInstanceModificationInstructionDto) UnsetCancelCurrentActiveActivityInstances()`

UnsetCancelCurrentActiveActivityInstances ensures that no value is present for CancelCurrentActiveActivityInstances, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


