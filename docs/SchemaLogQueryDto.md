# SchemaLogQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **NullableString** | The version of the schema. | [optional] 
**Sorting** | Pointer to [**[]SchemaLogQueryDtoSortingInner**](SchemaLogQueryDtoSortingInner.md) | A JSON array of criteria to sort the result by. Each element of the array is                       a JSON object that specifies one ordering. The position in the array                       identifies the rank of an ordering, i.e., whether it is primary, secondary,                       etc.  | [optional] 

## Methods

### NewSchemaLogQueryDto

`func NewSchemaLogQueryDto() *SchemaLogQueryDto`

NewSchemaLogQueryDto instantiates a new SchemaLogQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaLogQueryDtoWithDefaults

`func NewSchemaLogQueryDtoWithDefaults() *SchemaLogQueryDto`

NewSchemaLogQueryDtoWithDefaults instantiates a new SchemaLogQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SchemaLogQueryDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaLogQueryDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaLogQueryDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SchemaLogQueryDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SchemaLogQueryDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SchemaLogQueryDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetSorting

`func (o *SchemaLogQueryDto) GetSorting() []SchemaLogQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *SchemaLogQueryDto) GetSortingOk() (*[]SchemaLogQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *SchemaLogQueryDto) SetSorting(v []SchemaLogQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *SchemaLogQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *SchemaLogQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *SchemaLogQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


