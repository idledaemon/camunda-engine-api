# CorrelationMessageAsyncDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageName** | Pointer to **NullableString** | The name of the message to correlate. Corresponds to the &#39;name&#39; element of the message defined in BPMN 2.0 XML. Can be null to correlate by other criteria only. | [optional] 
**ProcessInstanceIds** | Pointer to **[]string** | A list of process instance ids that define a group of process instances to which the operation will correlate a message.  Please note that if &#x60;processInstanceIds&#x60;, &#x60;processInstanceQuery&#x60; and &#x60;historicProcessInstanceQuery&#x60; are defined, the resulting operation will be performed on the union of these sets. | [optional] 
**ProcessInstanceQuery** | Pointer to [**ProcessInstanceQueryDto**](ProcessInstanceQueryDto.md) |  | [optional] 
**HistoricProcessInstanceQuery** | Pointer to [**HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) |  | [optional] 
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | All variables the operation will set in the root scope of the process instances the message is correlated to. | [optional] 

## Methods

### NewCorrelationMessageAsyncDto

`func NewCorrelationMessageAsyncDto() *CorrelationMessageAsyncDto`

NewCorrelationMessageAsyncDto instantiates a new CorrelationMessageAsyncDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrelationMessageAsyncDtoWithDefaults

`func NewCorrelationMessageAsyncDtoWithDefaults() *CorrelationMessageAsyncDto`

NewCorrelationMessageAsyncDtoWithDefaults instantiates a new CorrelationMessageAsyncDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageName

`func (o *CorrelationMessageAsyncDto) GetMessageName() string`

GetMessageName returns the MessageName field if non-nil, zero value otherwise.

### GetMessageNameOk

`func (o *CorrelationMessageAsyncDto) GetMessageNameOk() (*string, bool)`

GetMessageNameOk returns a tuple with the MessageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageName

`func (o *CorrelationMessageAsyncDto) SetMessageName(v string)`

SetMessageName sets MessageName field to given value.

### HasMessageName

`func (o *CorrelationMessageAsyncDto) HasMessageName() bool`

HasMessageName returns a boolean if a field has been set.

### SetMessageNameNil

`func (o *CorrelationMessageAsyncDto) SetMessageNameNil(b bool)`

 SetMessageNameNil sets the value for MessageName to be an explicit nil

### UnsetMessageName
`func (o *CorrelationMessageAsyncDto) UnsetMessageName()`

UnsetMessageName ensures that no value is present for MessageName, not even an explicit nil
### GetProcessInstanceIds

`func (o *CorrelationMessageAsyncDto) GetProcessInstanceIds() []string`

GetProcessInstanceIds returns the ProcessInstanceIds field if non-nil, zero value otherwise.

### GetProcessInstanceIdsOk

`func (o *CorrelationMessageAsyncDto) GetProcessInstanceIdsOk() (*[]string, bool)`

GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIds

`func (o *CorrelationMessageAsyncDto) SetProcessInstanceIds(v []string)`

SetProcessInstanceIds sets ProcessInstanceIds field to given value.

### HasProcessInstanceIds

`func (o *CorrelationMessageAsyncDto) HasProcessInstanceIds() bool`

HasProcessInstanceIds returns a boolean if a field has been set.

### SetProcessInstanceIdsNil

`func (o *CorrelationMessageAsyncDto) SetProcessInstanceIdsNil(b bool)`

 SetProcessInstanceIdsNil sets the value for ProcessInstanceIds to be an explicit nil

### UnsetProcessInstanceIds
`func (o *CorrelationMessageAsyncDto) UnsetProcessInstanceIds()`

UnsetProcessInstanceIds ensures that no value is present for ProcessInstanceIds, not even an explicit nil
### GetProcessInstanceQuery

`func (o *CorrelationMessageAsyncDto) GetProcessInstanceQuery() ProcessInstanceQueryDto`

GetProcessInstanceQuery returns the ProcessInstanceQuery field if non-nil, zero value otherwise.

### GetProcessInstanceQueryOk

`func (o *CorrelationMessageAsyncDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool)`

GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceQuery

`func (o *CorrelationMessageAsyncDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto)`

SetProcessInstanceQuery sets ProcessInstanceQuery field to given value.

### HasProcessInstanceQuery

`func (o *CorrelationMessageAsyncDto) HasProcessInstanceQuery() bool`

HasProcessInstanceQuery returns a boolean if a field has been set.

### GetHistoricProcessInstanceQuery

`func (o *CorrelationMessageAsyncDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto`

GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field if non-nil, zero value otherwise.

### GetHistoricProcessInstanceQueryOk

`func (o *CorrelationMessageAsyncDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool)`

GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricProcessInstanceQuery

`func (o *CorrelationMessageAsyncDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto)`

SetHistoricProcessInstanceQuery sets HistoricProcessInstanceQuery field to given value.

### HasHistoricProcessInstanceQuery

`func (o *CorrelationMessageAsyncDto) HasHistoricProcessInstanceQuery() bool`

HasHistoricProcessInstanceQuery returns a boolean if a field has been set.

### GetVariables

`func (o *CorrelationMessageAsyncDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *CorrelationMessageAsyncDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *CorrelationMessageAsyncDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *CorrelationMessageAsyncDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *CorrelationMessageAsyncDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *CorrelationMessageAsyncDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


