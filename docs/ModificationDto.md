# ModificationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition for the modification | [optional] 
**SkipCustomListeners** | Pointer to **NullableBool** | Skip execution listener invocation for activities that are started or ended as part of this request. | [optional] 
**SkipIoMappings** | Pointer to **NullableBool** | Skip execution of [input/output variable mappings](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#input-output-variable-mapping) for activities that are started or ended as part of this request. | [optional] 
**ProcessInstanceIds** | Pointer to **[]string** | A list of process instance ids to modify. | [optional] 
**ProcessInstanceQuery** | Pointer to [**ProcessInstanceQueryDto**](ProcessInstanceQueryDto.md) |  | [optional] 
**Instructions** | Pointer to [**[]MultipleProcessInstanceModificationInstructionDto**](MultipleProcessInstanceModificationInstructionDto.md) | An array of modification instructions. The instructions are executed in the order they are in.  | [optional] 
**Annotation** | Pointer to **NullableString** | An arbitrary text annotation set by a user for auditing reasons. | [optional] 

## Methods

### NewModificationDto

`func NewModificationDto() *ModificationDto`

NewModificationDto instantiates a new ModificationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModificationDtoWithDefaults

`func NewModificationDtoWithDefaults() *ModificationDto`

NewModificationDtoWithDefaults instantiates a new ModificationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessDefinitionId

`func (o *ModificationDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *ModificationDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *ModificationDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *ModificationDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *ModificationDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *ModificationDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetSkipCustomListeners

`func (o *ModificationDto) GetSkipCustomListeners() bool`

GetSkipCustomListeners returns the SkipCustomListeners field if non-nil, zero value otherwise.

### GetSkipCustomListenersOk

`func (o *ModificationDto) GetSkipCustomListenersOk() (*bool, bool)`

GetSkipCustomListenersOk returns a tuple with the SkipCustomListeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCustomListeners

`func (o *ModificationDto) SetSkipCustomListeners(v bool)`

SetSkipCustomListeners sets SkipCustomListeners field to given value.

### HasSkipCustomListeners

`func (o *ModificationDto) HasSkipCustomListeners() bool`

HasSkipCustomListeners returns a boolean if a field has been set.

### SetSkipCustomListenersNil

`func (o *ModificationDto) SetSkipCustomListenersNil(b bool)`

 SetSkipCustomListenersNil sets the value for SkipCustomListeners to be an explicit nil

### UnsetSkipCustomListeners
`func (o *ModificationDto) UnsetSkipCustomListeners()`

UnsetSkipCustomListeners ensures that no value is present for SkipCustomListeners, not even an explicit nil
### GetSkipIoMappings

`func (o *ModificationDto) GetSkipIoMappings() bool`

GetSkipIoMappings returns the SkipIoMappings field if non-nil, zero value otherwise.

### GetSkipIoMappingsOk

`func (o *ModificationDto) GetSkipIoMappingsOk() (*bool, bool)`

GetSkipIoMappingsOk returns a tuple with the SkipIoMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIoMappings

`func (o *ModificationDto) SetSkipIoMappings(v bool)`

SetSkipIoMappings sets SkipIoMappings field to given value.

### HasSkipIoMappings

`func (o *ModificationDto) HasSkipIoMappings() bool`

HasSkipIoMappings returns a boolean if a field has been set.

### SetSkipIoMappingsNil

`func (o *ModificationDto) SetSkipIoMappingsNil(b bool)`

 SetSkipIoMappingsNil sets the value for SkipIoMappings to be an explicit nil

### UnsetSkipIoMappings
`func (o *ModificationDto) UnsetSkipIoMappings()`

UnsetSkipIoMappings ensures that no value is present for SkipIoMappings, not even an explicit nil
### GetProcessInstanceIds

`func (o *ModificationDto) GetProcessInstanceIds() []string`

GetProcessInstanceIds returns the ProcessInstanceIds field if non-nil, zero value otherwise.

### GetProcessInstanceIdsOk

`func (o *ModificationDto) GetProcessInstanceIdsOk() (*[]string, bool)`

GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIds

`func (o *ModificationDto) SetProcessInstanceIds(v []string)`

SetProcessInstanceIds sets ProcessInstanceIds field to given value.

### HasProcessInstanceIds

`func (o *ModificationDto) HasProcessInstanceIds() bool`

HasProcessInstanceIds returns a boolean if a field has been set.

### SetProcessInstanceIdsNil

`func (o *ModificationDto) SetProcessInstanceIdsNil(b bool)`

 SetProcessInstanceIdsNil sets the value for ProcessInstanceIds to be an explicit nil

### UnsetProcessInstanceIds
`func (o *ModificationDto) UnsetProcessInstanceIds()`

UnsetProcessInstanceIds ensures that no value is present for ProcessInstanceIds, not even an explicit nil
### GetProcessInstanceQuery

`func (o *ModificationDto) GetProcessInstanceQuery() ProcessInstanceQueryDto`

GetProcessInstanceQuery returns the ProcessInstanceQuery field if non-nil, zero value otherwise.

### GetProcessInstanceQueryOk

`func (o *ModificationDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool)`

GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceQuery

`func (o *ModificationDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto)`

SetProcessInstanceQuery sets ProcessInstanceQuery field to given value.

### HasProcessInstanceQuery

`func (o *ModificationDto) HasProcessInstanceQuery() bool`

HasProcessInstanceQuery returns a boolean if a field has been set.

### GetInstructions

`func (o *ModificationDto) GetInstructions() []MultipleProcessInstanceModificationInstructionDto`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ModificationDto) GetInstructionsOk() (*[]MultipleProcessInstanceModificationInstructionDto, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ModificationDto) SetInstructions(v []MultipleProcessInstanceModificationInstructionDto)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ModificationDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ModificationDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ModificationDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetAnnotation

`func (o *ModificationDto) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *ModificationDto) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *ModificationDto) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *ModificationDto) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotationNil

`func (o *ModificationDto) SetAnnotationNil(b bool)`

 SetAnnotationNil sets the value for Annotation to be an explicit nil

### UnsetAnnotation
`func (o *ModificationDto) UnsetAnnotation()`

UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


