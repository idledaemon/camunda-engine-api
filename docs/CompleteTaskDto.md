# CompleteTaskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A JSON object containing variable key-value pairs. | [optional] 
**WithVariablesInReturn** | Pointer to **NullableBool** | Indicates whether the response should contain the process variables or not. The default is &#x60;false&#x60; with a response code of &#x60;204&#x60;. If set to &#x60;true&#x60; the response contains the process variables and has a response code of &#x60;200&#x60;. If the task is not associated with a process instance (e.g. if it&#39;s part of a case instance) no variables will be returned. | [optional] [default to false]

## Methods

### NewCompleteTaskDto

`func NewCompleteTaskDto() *CompleteTaskDto`

NewCompleteTaskDto instantiates a new CompleteTaskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteTaskDtoWithDefaults

`func NewCompleteTaskDtoWithDefaults() *CompleteTaskDto`

NewCompleteTaskDtoWithDefaults instantiates a new CompleteTaskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariables

`func (o *CompleteTaskDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *CompleteTaskDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *CompleteTaskDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *CompleteTaskDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *CompleteTaskDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *CompleteTaskDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetWithVariablesInReturn

`func (o *CompleteTaskDto) GetWithVariablesInReturn() bool`

GetWithVariablesInReturn returns the WithVariablesInReturn field if non-nil, zero value otherwise.

### GetWithVariablesInReturnOk

`func (o *CompleteTaskDto) GetWithVariablesInReturnOk() (*bool, bool)`

GetWithVariablesInReturnOk returns a tuple with the WithVariablesInReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithVariablesInReturn

`func (o *CompleteTaskDto) SetWithVariablesInReturn(v bool)`

SetWithVariablesInReturn sets WithVariablesInReturn field to given value.

### HasWithVariablesInReturn

`func (o *CompleteTaskDto) HasWithVariablesInReturn() bool`

HasWithVariablesInReturn returns a boolean if a field has been set.

### SetWithVariablesInReturnNil

`func (o *CompleteTaskDto) SetWithVariablesInReturnNil(b bool)`

 SetWithVariablesInReturnNil sets the value for WithVariablesInReturn to be an explicit nil

### UnsetWithVariablesInReturn
`func (o *CompleteTaskDto) UnsetWithVariablesInReturn()`

UnsetWithVariablesInReturn ensures that no value is present for WithVariablesInReturn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


