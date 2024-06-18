# RestartProcessInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessInstanceIds** | Pointer to **[]string** | A list of process instance ids to restart. | [optional] 
**HistoricProcessInstanceQuery** | Pointer to [**HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) |  | [optional] 
**SkipCustomListeners** | Pointer to **NullableBool** | Skip execution listener invocation for activities that are started as part of this request. | [optional] 
**SkipIoMappings** | Pointer to **NullableBool** | Skip execution of [input/output variable mappings](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#input-output-variable-mapping) for activities that are started as part of this request. | [optional] 
**InitialVariables** | Pointer to **NullableBool** | Set the initial set of variables during restart. By default, the last set of variables is used. | [optional] 
**WithoutBusinessKey** | Pointer to **NullableBool** | Do not take over the business key of the historic process instance. | [optional] 
**Instructions** | Pointer to [**[]RestartProcessInstanceModificationInstructionDto**](RestartProcessInstanceModificationInstructionDto.md) | **Optional**. A JSON array of instructions that specify which activities to start the process instance at. If this property is omitted, the process instance starts at its default blank start event. | [optional] 

## Methods

### NewRestartProcessInstanceDto

`func NewRestartProcessInstanceDto() *RestartProcessInstanceDto`

NewRestartProcessInstanceDto instantiates a new RestartProcessInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartProcessInstanceDtoWithDefaults

`func NewRestartProcessInstanceDtoWithDefaults() *RestartProcessInstanceDto`

NewRestartProcessInstanceDtoWithDefaults instantiates a new RestartProcessInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessInstanceIds

`func (o *RestartProcessInstanceDto) GetProcessInstanceIds() []string`

GetProcessInstanceIds returns the ProcessInstanceIds field if non-nil, zero value otherwise.

### GetProcessInstanceIdsOk

`func (o *RestartProcessInstanceDto) GetProcessInstanceIdsOk() (*[]string, bool)`

GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIds

`func (o *RestartProcessInstanceDto) SetProcessInstanceIds(v []string)`

SetProcessInstanceIds sets ProcessInstanceIds field to given value.

### HasProcessInstanceIds

`func (o *RestartProcessInstanceDto) HasProcessInstanceIds() bool`

HasProcessInstanceIds returns a boolean if a field has been set.

### SetProcessInstanceIdsNil

`func (o *RestartProcessInstanceDto) SetProcessInstanceIdsNil(b bool)`

 SetProcessInstanceIdsNil sets the value for ProcessInstanceIds to be an explicit nil

### UnsetProcessInstanceIds
`func (o *RestartProcessInstanceDto) UnsetProcessInstanceIds()`

UnsetProcessInstanceIds ensures that no value is present for ProcessInstanceIds, not even an explicit nil
### GetHistoricProcessInstanceQuery

`func (o *RestartProcessInstanceDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto`

GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field if non-nil, zero value otherwise.

### GetHistoricProcessInstanceQueryOk

`func (o *RestartProcessInstanceDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool)`

GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricProcessInstanceQuery

`func (o *RestartProcessInstanceDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto)`

SetHistoricProcessInstanceQuery sets HistoricProcessInstanceQuery field to given value.

### HasHistoricProcessInstanceQuery

`func (o *RestartProcessInstanceDto) HasHistoricProcessInstanceQuery() bool`

HasHistoricProcessInstanceQuery returns a boolean if a field has been set.

### GetSkipCustomListeners

`func (o *RestartProcessInstanceDto) GetSkipCustomListeners() bool`

GetSkipCustomListeners returns the SkipCustomListeners field if non-nil, zero value otherwise.

### GetSkipCustomListenersOk

`func (o *RestartProcessInstanceDto) GetSkipCustomListenersOk() (*bool, bool)`

GetSkipCustomListenersOk returns a tuple with the SkipCustomListeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCustomListeners

`func (o *RestartProcessInstanceDto) SetSkipCustomListeners(v bool)`

SetSkipCustomListeners sets SkipCustomListeners field to given value.

### HasSkipCustomListeners

`func (o *RestartProcessInstanceDto) HasSkipCustomListeners() bool`

HasSkipCustomListeners returns a boolean if a field has been set.

### SetSkipCustomListenersNil

`func (o *RestartProcessInstanceDto) SetSkipCustomListenersNil(b bool)`

 SetSkipCustomListenersNil sets the value for SkipCustomListeners to be an explicit nil

### UnsetSkipCustomListeners
`func (o *RestartProcessInstanceDto) UnsetSkipCustomListeners()`

UnsetSkipCustomListeners ensures that no value is present for SkipCustomListeners, not even an explicit nil
### GetSkipIoMappings

`func (o *RestartProcessInstanceDto) GetSkipIoMappings() bool`

GetSkipIoMappings returns the SkipIoMappings field if non-nil, zero value otherwise.

### GetSkipIoMappingsOk

`func (o *RestartProcessInstanceDto) GetSkipIoMappingsOk() (*bool, bool)`

GetSkipIoMappingsOk returns a tuple with the SkipIoMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIoMappings

`func (o *RestartProcessInstanceDto) SetSkipIoMappings(v bool)`

SetSkipIoMappings sets SkipIoMappings field to given value.

### HasSkipIoMappings

`func (o *RestartProcessInstanceDto) HasSkipIoMappings() bool`

HasSkipIoMappings returns a boolean if a field has been set.

### SetSkipIoMappingsNil

`func (o *RestartProcessInstanceDto) SetSkipIoMappingsNil(b bool)`

 SetSkipIoMappingsNil sets the value for SkipIoMappings to be an explicit nil

### UnsetSkipIoMappings
`func (o *RestartProcessInstanceDto) UnsetSkipIoMappings()`

UnsetSkipIoMappings ensures that no value is present for SkipIoMappings, not even an explicit nil
### GetInitialVariables

`func (o *RestartProcessInstanceDto) GetInitialVariables() bool`

GetInitialVariables returns the InitialVariables field if non-nil, zero value otherwise.

### GetInitialVariablesOk

`func (o *RestartProcessInstanceDto) GetInitialVariablesOk() (*bool, bool)`

GetInitialVariablesOk returns a tuple with the InitialVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialVariables

`func (o *RestartProcessInstanceDto) SetInitialVariables(v bool)`

SetInitialVariables sets InitialVariables field to given value.

### HasInitialVariables

`func (o *RestartProcessInstanceDto) HasInitialVariables() bool`

HasInitialVariables returns a boolean if a field has been set.

### SetInitialVariablesNil

`func (o *RestartProcessInstanceDto) SetInitialVariablesNil(b bool)`

 SetInitialVariablesNil sets the value for InitialVariables to be an explicit nil

### UnsetInitialVariables
`func (o *RestartProcessInstanceDto) UnsetInitialVariables()`

UnsetInitialVariables ensures that no value is present for InitialVariables, not even an explicit nil
### GetWithoutBusinessKey

`func (o *RestartProcessInstanceDto) GetWithoutBusinessKey() bool`

GetWithoutBusinessKey returns the WithoutBusinessKey field if non-nil, zero value otherwise.

### GetWithoutBusinessKeyOk

`func (o *RestartProcessInstanceDto) GetWithoutBusinessKeyOk() (*bool, bool)`

GetWithoutBusinessKeyOk returns a tuple with the WithoutBusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutBusinessKey

`func (o *RestartProcessInstanceDto) SetWithoutBusinessKey(v bool)`

SetWithoutBusinessKey sets WithoutBusinessKey field to given value.

### HasWithoutBusinessKey

`func (o *RestartProcessInstanceDto) HasWithoutBusinessKey() bool`

HasWithoutBusinessKey returns a boolean if a field has been set.

### SetWithoutBusinessKeyNil

`func (o *RestartProcessInstanceDto) SetWithoutBusinessKeyNil(b bool)`

 SetWithoutBusinessKeyNil sets the value for WithoutBusinessKey to be an explicit nil

### UnsetWithoutBusinessKey
`func (o *RestartProcessInstanceDto) UnsetWithoutBusinessKey()`

UnsetWithoutBusinessKey ensures that no value is present for WithoutBusinessKey, not even an explicit nil
### GetInstructions

`func (o *RestartProcessInstanceDto) GetInstructions() []RestartProcessInstanceModificationInstructionDto`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *RestartProcessInstanceDto) GetInstructionsOk() (*[]RestartProcessInstanceModificationInstructionDto, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *RestartProcessInstanceDto) SetInstructions(v []RestartProcessInstanceModificationInstructionDto)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *RestartProcessInstanceDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *RestartProcessInstanceDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *RestartProcessInstanceDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


