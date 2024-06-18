# ProcessInstanceModificationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipCustomListeners** | Pointer to **NullableBool** | Skip execution listener invocation for activities that are started or ended as part of this request. | [optional] 
**SkipIoMappings** | Pointer to **NullableBool** | Skip execution of [input/output variable mappings](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#input-output-variable-mapping) for activities that are started or ended as part of this request. | [optional] 
**Instructions** | Pointer to [**[]ProcessInstanceModificationInstructionDto**](ProcessInstanceModificationInstructionDto.md) | JSON array of modification instructions. The instructions are executed in the order they are in. | [optional] 
**Annotation** | Pointer to **NullableString** | An arbitrary text annotation set by a user for auditing reasons. | [optional] 

## Methods

### NewProcessInstanceModificationDto

`func NewProcessInstanceModificationDto() *ProcessInstanceModificationDto`

NewProcessInstanceModificationDto instantiates a new ProcessInstanceModificationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessInstanceModificationDtoWithDefaults

`func NewProcessInstanceModificationDtoWithDefaults() *ProcessInstanceModificationDto`

NewProcessInstanceModificationDtoWithDefaults instantiates a new ProcessInstanceModificationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipCustomListeners

`func (o *ProcessInstanceModificationDto) GetSkipCustomListeners() bool`

GetSkipCustomListeners returns the SkipCustomListeners field if non-nil, zero value otherwise.

### GetSkipCustomListenersOk

`func (o *ProcessInstanceModificationDto) GetSkipCustomListenersOk() (*bool, bool)`

GetSkipCustomListenersOk returns a tuple with the SkipCustomListeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCustomListeners

`func (o *ProcessInstanceModificationDto) SetSkipCustomListeners(v bool)`

SetSkipCustomListeners sets SkipCustomListeners field to given value.

### HasSkipCustomListeners

`func (o *ProcessInstanceModificationDto) HasSkipCustomListeners() bool`

HasSkipCustomListeners returns a boolean if a field has been set.

### SetSkipCustomListenersNil

`func (o *ProcessInstanceModificationDto) SetSkipCustomListenersNil(b bool)`

 SetSkipCustomListenersNil sets the value for SkipCustomListeners to be an explicit nil

### UnsetSkipCustomListeners
`func (o *ProcessInstanceModificationDto) UnsetSkipCustomListeners()`

UnsetSkipCustomListeners ensures that no value is present for SkipCustomListeners, not even an explicit nil
### GetSkipIoMappings

`func (o *ProcessInstanceModificationDto) GetSkipIoMappings() bool`

GetSkipIoMappings returns the SkipIoMappings field if non-nil, zero value otherwise.

### GetSkipIoMappingsOk

`func (o *ProcessInstanceModificationDto) GetSkipIoMappingsOk() (*bool, bool)`

GetSkipIoMappingsOk returns a tuple with the SkipIoMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIoMappings

`func (o *ProcessInstanceModificationDto) SetSkipIoMappings(v bool)`

SetSkipIoMappings sets SkipIoMappings field to given value.

### HasSkipIoMappings

`func (o *ProcessInstanceModificationDto) HasSkipIoMappings() bool`

HasSkipIoMappings returns a boolean if a field has been set.

### SetSkipIoMappingsNil

`func (o *ProcessInstanceModificationDto) SetSkipIoMappingsNil(b bool)`

 SetSkipIoMappingsNil sets the value for SkipIoMappings to be an explicit nil

### UnsetSkipIoMappings
`func (o *ProcessInstanceModificationDto) UnsetSkipIoMappings()`

UnsetSkipIoMappings ensures that no value is present for SkipIoMappings, not even an explicit nil
### GetInstructions

`func (o *ProcessInstanceModificationDto) GetInstructions() []ProcessInstanceModificationInstructionDto`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ProcessInstanceModificationDto) GetInstructionsOk() (*[]ProcessInstanceModificationInstructionDto, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ProcessInstanceModificationDto) SetInstructions(v []ProcessInstanceModificationInstructionDto)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ProcessInstanceModificationDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ProcessInstanceModificationDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ProcessInstanceModificationDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetAnnotation

`func (o *ProcessInstanceModificationDto) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *ProcessInstanceModificationDto) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *ProcessInstanceModificationDto) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *ProcessInstanceModificationDto) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotationNil

`func (o *ProcessInstanceModificationDto) SetAnnotationNil(b bool)`

 SetAnnotationNil sets the value for Annotation to be an explicit nil

### UnsetAnnotation
`func (o *ProcessInstanceModificationDto) UnsetAnnotation()`

UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


