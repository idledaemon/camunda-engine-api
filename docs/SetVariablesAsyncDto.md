# SetVariablesAsyncDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessInstanceIds** | Pointer to **[]string** | A list of process instance ids that define a group of process instances to which the operation will set variables.  Please note that if &#x60;processInstanceIds&#x60;, &#x60;processInstanceQuery&#x60; and &#x60;historicProcessInstanceQuery&#x60; are defined, the resulting operation will be performed on the union of these sets. | [optional] 
**ProcessInstanceQuery** | Pointer to [**ProcessInstanceQueryDto**](ProcessInstanceQueryDto.md) |  | [optional] 
**HistoricProcessInstanceQuery** | Pointer to [**HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) |  | [optional] 
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A variables the operation will set in the root scope of the process instances. | [optional] 

## Methods

### NewSetVariablesAsyncDto

`func NewSetVariablesAsyncDto() *SetVariablesAsyncDto`

NewSetVariablesAsyncDto instantiates a new SetVariablesAsyncDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetVariablesAsyncDtoWithDefaults

`func NewSetVariablesAsyncDtoWithDefaults() *SetVariablesAsyncDto`

NewSetVariablesAsyncDtoWithDefaults instantiates a new SetVariablesAsyncDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessInstanceIds

`func (o *SetVariablesAsyncDto) GetProcessInstanceIds() []string`

GetProcessInstanceIds returns the ProcessInstanceIds field if non-nil, zero value otherwise.

### GetProcessInstanceIdsOk

`func (o *SetVariablesAsyncDto) GetProcessInstanceIdsOk() (*[]string, bool)`

GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIds

`func (o *SetVariablesAsyncDto) SetProcessInstanceIds(v []string)`

SetProcessInstanceIds sets ProcessInstanceIds field to given value.

### HasProcessInstanceIds

`func (o *SetVariablesAsyncDto) HasProcessInstanceIds() bool`

HasProcessInstanceIds returns a boolean if a field has been set.

### SetProcessInstanceIdsNil

`func (o *SetVariablesAsyncDto) SetProcessInstanceIdsNil(b bool)`

 SetProcessInstanceIdsNil sets the value for ProcessInstanceIds to be an explicit nil

### UnsetProcessInstanceIds
`func (o *SetVariablesAsyncDto) UnsetProcessInstanceIds()`

UnsetProcessInstanceIds ensures that no value is present for ProcessInstanceIds, not even an explicit nil
### GetProcessInstanceQuery

`func (o *SetVariablesAsyncDto) GetProcessInstanceQuery() ProcessInstanceQueryDto`

GetProcessInstanceQuery returns the ProcessInstanceQuery field if non-nil, zero value otherwise.

### GetProcessInstanceQueryOk

`func (o *SetVariablesAsyncDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool)`

GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceQuery

`func (o *SetVariablesAsyncDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto)`

SetProcessInstanceQuery sets ProcessInstanceQuery field to given value.

### HasProcessInstanceQuery

`func (o *SetVariablesAsyncDto) HasProcessInstanceQuery() bool`

HasProcessInstanceQuery returns a boolean if a field has been set.

### GetHistoricProcessInstanceQuery

`func (o *SetVariablesAsyncDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto`

GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field if non-nil, zero value otherwise.

### GetHistoricProcessInstanceQueryOk

`func (o *SetVariablesAsyncDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool)`

GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricProcessInstanceQuery

`func (o *SetVariablesAsyncDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto)`

SetHistoricProcessInstanceQuery sets HistoricProcessInstanceQuery field to given value.

### HasHistoricProcessInstanceQuery

`func (o *SetVariablesAsyncDto) HasHistoricProcessInstanceQuery() bool`

HasHistoricProcessInstanceQuery returns a boolean if a field has been set.

### GetVariables

`func (o *SetVariablesAsyncDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SetVariablesAsyncDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *SetVariablesAsyncDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *SetVariablesAsyncDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *SetVariablesAsyncDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *SetVariablesAsyncDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


