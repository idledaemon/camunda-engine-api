# StartProcessInstanceFormDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) |  | [optional] 
**BusinessKey** | Pointer to **NullableString** | The business key the process instance is to be initialized with. The business key uniquely identifies the process instance in the context of the given process definition. | [optional] 

## Methods

### NewStartProcessInstanceFormDto

`func NewStartProcessInstanceFormDto() *StartProcessInstanceFormDto`

NewStartProcessInstanceFormDto instantiates a new StartProcessInstanceFormDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartProcessInstanceFormDtoWithDefaults

`func NewStartProcessInstanceFormDtoWithDefaults() *StartProcessInstanceFormDto`

NewStartProcessInstanceFormDtoWithDefaults instantiates a new StartProcessInstanceFormDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariables

`func (o *StartProcessInstanceFormDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *StartProcessInstanceFormDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *StartProcessInstanceFormDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *StartProcessInstanceFormDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *StartProcessInstanceFormDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *StartProcessInstanceFormDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetBusinessKey

`func (o *StartProcessInstanceFormDto) GetBusinessKey() string`

GetBusinessKey returns the BusinessKey field if non-nil, zero value otherwise.

### GetBusinessKeyOk

`func (o *StartProcessInstanceFormDto) GetBusinessKeyOk() (*string, bool)`

GetBusinessKeyOk returns a tuple with the BusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKey

`func (o *StartProcessInstanceFormDto) SetBusinessKey(v string)`

SetBusinessKey sets BusinessKey field to given value.

### HasBusinessKey

`func (o *StartProcessInstanceFormDto) HasBusinessKey() bool`

HasBusinessKey returns a boolean if a field has been set.

### SetBusinessKeyNil

`func (o *StartProcessInstanceFormDto) SetBusinessKeyNil(b bool)`

 SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil

### UnsetBusinessKey
`func (o *StartProcessInstanceFormDto) UnsetBusinessKey()`

UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


