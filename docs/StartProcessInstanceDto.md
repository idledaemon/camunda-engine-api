# StartProcessInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessKey** | Pointer to **NullableString** | The business key of the process instance. | [optional] 
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) |  | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | The case instance id the process instance is to be initialized with. | [optional] 
**StartInstructions** | Pointer to [**[]ProcessInstanceModificationInstructionDto**](ProcessInstanceModificationInstructionDto.md) | **Optional**. A JSON array of instructions that specify which activities to start the process instance at. If this property is omitted, the process instance starts at its default blank start event. | [optional] 
**SkipCustomListeners** | Pointer to **NullableBool** | Skip execution listener invocation for activities that are started or ended as part of this request. **Note**: This option is currently only respected when start instructions are submitted via the &#x60;startInstructions&#x60; property. | [optional] 
**SkipIoMappings** | Pointer to **NullableBool** | Skip execution of [input/output variable mappings](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#input-output-variable-mapping) for activities that are started or ended as part of this request. **Note**: This option is currently only respected when start instructions are submitted via the &#x60;startInstructions&#x60; property. | [optional] 
**WithVariablesInReturn** | Pointer to **NullableBool** | Indicates if the variables, which was used by the process instance during execution, should be returned. Default value: &#x60;false&#x60; | [optional] 

## Methods

### NewStartProcessInstanceDto

`func NewStartProcessInstanceDto() *StartProcessInstanceDto`

NewStartProcessInstanceDto instantiates a new StartProcessInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartProcessInstanceDtoWithDefaults

`func NewStartProcessInstanceDtoWithDefaults() *StartProcessInstanceDto`

NewStartProcessInstanceDtoWithDefaults instantiates a new StartProcessInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessKey

`func (o *StartProcessInstanceDto) GetBusinessKey() string`

GetBusinessKey returns the BusinessKey field if non-nil, zero value otherwise.

### GetBusinessKeyOk

`func (o *StartProcessInstanceDto) GetBusinessKeyOk() (*string, bool)`

GetBusinessKeyOk returns a tuple with the BusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKey

`func (o *StartProcessInstanceDto) SetBusinessKey(v string)`

SetBusinessKey sets BusinessKey field to given value.

### HasBusinessKey

`func (o *StartProcessInstanceDto) HasBusinessKey() bool`

HasBusinessKey returns a boolean if a field has been set.

### SetBusinessKeyNil

`func (o *StartProcessInstanceDto) SetBusinessKeyNil(b bool)`

 SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil

### UnsetBusinessKey
`func (o *StartProcessInstanceDto) UnsetBusinessKey()`

UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
### GetVariables

`func (o *StartProcessInstanceDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *StartProcessInstanceDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *StartProcessInstanceDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *StartProcessInstanceDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *StartProcessInstanceDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *StartProcessInstanceDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetCaseInstanceId

`func (o *StartProcessInstanceDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *StartProcessInstanceDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *StartProcessInstanceDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *StartProcessInstanceDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *StartProcessInstanceDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *StartProcessInstanceDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetStartInstructions

`func (o *StartProcessInstanceDto) GetStartInstructions() []ProcessInstanceModificationInstructionDto`

GetStartInstructions returns the StartInstructions field if non-nil, zero value otherwise.

### GetStartInstructionsOk

`func (o *StartProcessInstanceDto) GetStartInstructionsOk() (*[]ProcessInstanceModificationInstructionDto, bool)`

GetStartInstructionsOk returns a tuple with the StartInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInstructions

`func (o *StartProcessInstanceDto) SetStartInstructions(v []ProcessInstanceModificationInstructionDto)`

SetStartInstructions sets StartInstructions field to given value.

### HasStartInstructions

`func (o *StartProcessInstanceDto) HasStartInstructions() bool`

HasStartInstructions returns a boolean if a field has been set.

### SetStartInstructionsNil

`func (o *StartProcessInstanceDto) SetStartInstructionsNil(b bool)`

 SetStartInstructionsNil sets the value for StartInstructions to be an explicit nil

### UnsetStartInstructions
`func (o *StartProcessInstanceDto) UnsetStartInstructions()`

UnsetStartInstructions ensures that no value is present for StartInstructions, not even an explicit nil
### GetSkipCustomListeners

`func (o *StartProcessInstanceDto) GetSkipCustomListeners() bool`

GetSkipCustomListeners returns the SkipCustomListeners field if non-nil, zero value otherwise.

### GetSkipCustomListenersOk

`func (o *StartProcessInstanceDto) GetSkipCustomListenersOk() (*bool, bool)`

GetSkipCustomListenersOk returns a tuple with the SkipCustomListeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCustomListeners

`func (o *StartProcessInstanceDto) SetSkipCustomListeners(v bool)`

SetSkipCustomListeners sets SkipCustomListeners field to given value.

### HasSkipCustomListeners

`func (o *StartProcessInstanceDto) HasSkipCustomListeners() bool`

HasSkipCustomListeners returns a boolean if a field has been set.

### SetSkipCustomListenersNil

`func (o *StartProcessInstanceDto) SetSkipCustomListenersNil(b bool)`

 SetSkipCustomListenersNil sets the value for SkipCustomListeners to be an explicit nil

### UnsetSkipCustomListeners
`func (o *StartProcessInstanceDto) UnsetSkipCustomListeners()`

UnsetSkipCustomListeners ensures that no value is present for SkipCustomListeners, not even an explicit nil
### GetSkipIoMappings

`func (o *StartProcessInstanceDto) GetSkipIoMappings() bool`

GetSkipIoMappings returns the SkipIoMappings field if non-nil, zero value otherwise.

### GetSkipIoMappingsOk

`func (o *StartProcessInstanceDto) GetSkipIoMappingsOk() (*bool, bool)`

GetSkipIoMappingsOk returns a tuple with the SkipIoMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIoMappings

`func (o *StartProcessInstanceDto) SetSkipIoMappings(v bool)`

SetSkipIoMappings sets SkipIoMappings field to given value.

### HasSkipIoMappings

`func (o *StartProcessInstanceDto) HasSkipIoMappings() bool`

HasSkipIoMappings returns a boolean if a field has been set.

### SetSkipIoMappingsNil

`func (o *StartProcessInstanceDto) SetSkipIoMappingsNil(b bool)`

 SetSkipIoMappingsNil sets the value for SkipIoMappings to be an explicit nil

### UnsetSkipIoMappings
`func (o *StartProcessInstanceDto) UnsetSkipIoMappings()`

UnsetSkipIoMappings ensures that no value is present for SkipIoMappings, not even an explicit nil
### GetWithVariablesInReturn

`func (o *StartProcessInstanceDto) GetWithVariablesInReturn() bool`

GetWithVariablesInReturn returns the WithVariablesInReturn field if non-nil, zero value otherwise.

### GetWithVariablesInReturnOk

`func (o *StartProcessInstanceDto) GetWithVariablesInReturnOk() (*bool, bool)`

GetWithVariablesInReturnOk returns a tuple with the WithVariablesInReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithVariablesInReturn

`func (o *StartProcessInstanceDto) SetWithVariablesInReturn(v bool)`

SetWithVariablesInReturn sets WithVariablesInReturn field to given value.

### HasWithVariablesInReturn

`func (o *StartProcessInstanceDto) HasWithVariablesInReturn() bool`

HasWithVariablesInReturn returns a boolean if a field has been set.

### SetWithVariablesInReturnNil

`func (o *StartProcessInstanceDto) SetWithVariablesInReturnNil(b bool)`

 SetWithVariablesInReturnNil sets the value for WithVariablesInReturn to be an explicit nil

### UnsetWithVariablesInReturn
`func (o *StartProcessInstanceDto) UnsetWithVariablesInReturn()`

UnsetWithVariablesInReturn ensures that no value is present for WithVariablesInReturn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


