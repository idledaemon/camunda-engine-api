# SchemaLogEntryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the schema log entry. | [optional] 
**Timestamp** | Pointer to **NullableTime** | The date and time of the schema update. | [optional] 
**Version** | Pointer to **NullableString** | The version of the schema. | [optional] 

## Methods

### NewSchemaLogEntryDto

`func NewSchemaLogEntryDto() *SchemaLogEntryDto`

NewSchemaLogEntryDto instantiates a new SchemaLogEntryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaLogEntryDtoWithDefaults

`func NewSchemaLogEntryDtoWithDefaults() *SchemaLogEntryDto`

NewSchemaLogEntryDtoWithDefaults instantiates a new SchemaLogEntryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchemaLogEntryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemaLogEntryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemaLogEntryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemaLogEntryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SchemaLogEntryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SchemaLogEntryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SchemaLogEntryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SchemaLogEntryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SchemaLogEntryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SchemaLogEntryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SchemaLogEntryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SchemaLogEntryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetVersion

`func (o *SchemaLogEntryDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaLogEntryDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaLogEntryDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SchemaLogEntryDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SchemaLogEntryDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SchemaLogEntryDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


