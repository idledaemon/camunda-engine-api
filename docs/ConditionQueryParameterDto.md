# ConditionQueryParameterDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **NullableString** | Comparison operator to be used. &#x60;notLike&#x60; is not supported by all endpoints. | [optional] 
**Value** | Pointer to **interface{}** | Can be any value - string, number, boolean, array or object.  **Note**: Not every endpoint supports every type. | [optional] 

## Methods

### NewConditionQueryParameterDto

`func NewConditionQueryParameterDto() *ConditionQueryParameterDto`

NewConditionQueryParameterDto instantiates a new ConditionQueryParameterDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionQueryParameterDtoWithDefaults

`func NewConditionQueryParameterDtoWithDefaults() *ConditionQueryParameterDto`

NewConditionQueryParameterDtoWithDefaults instantiates a new ConditionQueryParameterDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ConditionQueryParameterDto) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ConditionQueryParameterDto) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ConditionQueryParameterDto) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ConditionQueryParameterDto) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *ConditionQueryParameterDto) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *ConditionQueryParameterDto) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetValue

`func (o *ConditionQueryParameterDto) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConditionQueryParameterDto) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConditionQueryParameterDto) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ConditionQueryParameterDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ConditionQueryParameterDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ConditionQueryParameterDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


