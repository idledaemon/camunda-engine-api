# JobConditionQueryParameterDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **NullableString** | Comparison operator to be used. | [optional] 
**Value** | Pointer to **NullableTime** | Date value to compare with. | [optional] 

## Methods

### NewJobConditionQueryParameterDto

`func NewJobConditionQueryParameterDto() *JobConditionQueryParameterDto`

NewJobConditionQueryParameterDto instantiates a new JobConditionQueryParameterDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobConditionQueryParameterDtoWithDefaults

`func NewJobConditionQueryParameterDtoWithDefaults() *JobConditionQueryParameterDto`

NewJobConditionQueryParameterDtoWithDefaults instantiates a new JobConditionQueryParameterDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *JobConditionQueryParameterDto) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *JobConditionQueryParameterDto) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *JobConditionQueryParameterDto) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *JobConditionQueryParameterDto) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *JobConditionQueryParameterDto) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *JobConditionQueryParameterDto) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetValue

`func (o *JobConditionQueryParameterDto) GetValue() time.Time`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JobConditionQueryParameterDto) GetValueOk() (*time.Time, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JobConditionQueryParameterDto) SetValue(v time.Time)`

SetValue sets Value field to given value.

### HasValue

`func (o *JobConditionQueryParameterDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *JobConditionQueryParameterDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *JobConditionQueryParameterDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


