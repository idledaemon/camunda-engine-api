# VariableQueryParameterDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **NullableString** | Comparison operator to be used. &#x60;notLike&#x60; is not supported by all endpoints. | [optional] 
**Value** | Pointer to **interface{}** | Can be any value - string, number, boolean, array or object.  **Note**: Not every endpoint supports every type. | [optional] 
**Name** | Pointer to **NullableString** | Variable name | [optional] 

## Methods

### NewVariableQueryParameterDto

`func NewVariableQueryParameterDto() *VariableQueryParameterDto`

NewVariableQueryParameterDto instantiates a new VariableQueryParameterDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableQueryParameterDtoWithDefaults

`func NewVariableQueryParameterDtoWithDefaults() *VariableQueryParameterDto`

NewVariableQueryParameterDtoWithDefaults instantiates a new VariableQueryParameterDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *VariableQueryParameterDto) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *VariableQueryParameterDto) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *VariableQueryParameterDto) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *VariableQueryParameterDto) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *VariableQueryParameterDto) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *VariableQueryParameterDto) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetValue

`func (o *VariableQueryParameterDto) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableQueryParameterDto) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableQueryParameterDto) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableQueryParameterDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *VariableQueryParameterDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *VariableQueryParameterDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetName

`func (o *VariableQueryParameterDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableQueryParameterDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableQueryParameterDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VariableQueryParameterDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VariableQueryParameterDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VariableQueryParameterDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


