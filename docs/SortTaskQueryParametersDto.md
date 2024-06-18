# SortTaskQueryParametersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variable** | Pointer to **NullableString** | The name of the variable to sort by. | [optional] 
**Type** | Pointer to **NullableString** | The name of the type of the variable value. | [optional] 

## Methods

### NewSortTaskQueryParametersDto

`func NewSortTaskQueryParametersDto() *SortTaskQueryParametersDto`

NewSortTaskQueryParametersDto instantiates a new SortTaskQueryParametersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortTaskQueryParametersDtoWithDefaults

`func NewSortTaskQueryParametersDtoWithDefaults() *SortTaskQueryParametersDto`

NewSortTaskQueryParametersDtoWithDefaults instantiates a new SortTaskQueryParametersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariable

`func (o *SortTaskQueryParametersDto) GetVariable() string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *SortTaskQueryParametersDto) GetVariableOk() (*string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *SortTaskQueryParametersDto) SetVariable(v string)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *SortTaskQueryParametersDto) HasVariable() bool`

HasVariable returns a boolean if a field has been set.

### SetVariableNil

`func (o *SortTaskQueryParametersDto) SetVariableNil(b bool)`

 SetVariableNil sets the value for Variable to be an explicit nil

### UnsetVariable
`func (o *SortTaskQueryParametersDto) UnsetVariable()`

UnsetVariable ensures that no value is present for Variable, not even an explicit nil
### GetType

`func (o *SortTaskQueryParametersDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SortTaskQueryParametersDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SortTaskQueryParametersDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SortTaskQueryParametersDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SortTaskQueryParametersDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SortTaskQueryParametersDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


