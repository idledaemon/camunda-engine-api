# CompleteExternalTaskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerId** | Pointer to **string** | **Mandatory.** The ID of the worker who is performing the operation on the external task. If the task is already locked, must match the id of the worker who has most recently locked the task. | [optional] 
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A JSON object containing variable key-value pairs. Each key is a variable name and each value a JSON variable value object with the following properties: | [optional] 
**LocalVariables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A JSON object containing local variable key-value pairs. Local variables are set only in the scope of external task. Each key is a variable name and each value a JSON variable value object with the following properties: | [optional] 

## Methods

### NewCompleteExternalTaskDto

`func NewCompleteExternalTaskDto() *CompleteExternalTaskDto`

NewCompleteExternalTaskDto instantiates a new CompleteExternalTaskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteExternalTaskDtoWithDefaults

`func NewCompleteExternalTaskDtoWithDefaults() *CompleteExternalTaskDto`

NewCompleteExternalTaskDtoWithDefaults instantiates a new CompleteExternalTaskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerId

`func (o *CompleteExternalTaskDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *CompleteExternalTaskDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *CompleteExternalTaskDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *CompleteExternalTaskDto) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetVariables

`func (o *CompleteExternalTaskDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *CompleteExternalTaskDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *CompleteExternalTaskDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *CompleteExternalTaskDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *CompleteExternalTaskDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *CompleteExternalTaskDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetLocalVariables

`func (o *CompleteExternalTaskDto) GetLocalVariables() map[string]VariableValueDto`

GetLocalVariables returns the LocalVariables field if non-nil, zero value otherwise.

### GetLocalVariablesOk

`func (o *CompleteExternalTaskDto) GetLocalVariablesOk() (*map[string]VariableValueDto, bool)`

GetLocalVariablesOk returns a tuple with the LocalVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVariables

`func (o *CompleteExternalTaskDto) SetLocalVariables(v map[string]VariableValueDto)`

SetLocalVariables sets LocalVariables field to given value.

### HasLocalVariables

`func (o *CompleteExternalTaskDto) HasLocalVariables() bool`

HasLocalVariables returns a boolean if a field has been set.

### SetLocalVariablesNil

`func (o *CompleteExternalTaskDto) SetLocalVariablesNil(b bool)`

 SetLocalVariablesNil sets the value for LocalVariables to be an explicit nil

### UnsetLocalVariables
`func (o *CompleteExternalTaskDto) UnsetLocalVariables()`

UnsetLocalVariables ensures that no value is present for LocalVariables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


