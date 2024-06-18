# ExecutionTriggerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A JSON object containing variable key-value pairs. Each key is a variable name and each value a JSON variable value object. | [optional] 

## Methods

### NewExecutionTriggerDto

`func NewExecutionTriggerDto() *ExecutionTriggerDto`

NewExecutionTriggerDto instantiates a new ExecutionTriggerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionTriggerDtoWithDefaults

`func NewExecutionTriggerDtoWithDefaults() *ExecutionTriggerDto`

NewExecutionTriggerDtoWithDefaults instantiates a new ExecutionTriggerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariables

`func (o *ExecutionTriggerDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ExecutionTriggerDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ExecutionTriggerDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ExecutionTriggerDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *ExecutionTriggerDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *ExecutionTriggerDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


