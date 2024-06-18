# ExecutionQueryDtoSortingInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SortBy** | Pointer to **NullableString** | Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter. | [optional] 
**SortOrder** | Pointer to **NullableString** | Sort the results in a given order. Values may be &#x60;asc&#x60; for ascending order or &#x60;desc&#x60; for descending order. Must be used in conjunction with the sortBy parameter. | [optional] 

## Methods

### NewExecutionQueryDtoSortingInner

`func NewExecutionQueryDtoSortingInner() *ExecutionQueryDtoSortingInner`

NewExecutionQueryDtoSortingInner instantiates a new ExecutionQueryDtoSortingInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionQueryDtoSortingInnerWithDefaults

`func NewExecutionQueryDtoSortingInnerWithDefaults() *ExecutionQueryDtoSortingInner`

NewExecutionQueryDtoSortingInnerWithDefaults instantiates a new ExecutionQueryDtoSortingInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSortBy

`func (o *ExecutionQueryDtoSortingInner) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ExecutionQueryDtoSortingInner) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ExecutionQueryDtoSortingInner) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ExecutionQueryDtoSortingInner) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *ExecutionQueryDtoSortingInner) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *ExecutionQueryDtoSortingInner) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortOrder

`func (o *ExecutionQueryDtoSortingInner) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ExecutionQueryDtoSortingInner) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ExecutionQueryDtoSortingInner) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ExecutionQueryDtoSortingInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *ExecutionQueryDtoSortingInner) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *ExecutionQueryDtoSortingInner) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


