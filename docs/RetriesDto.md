# RetriesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retries** | Pointer to **NullableInt32** | The number of retries to set for the resource.  Must be &gt;&#x3D; 0. If this is 0, an incident is created and the task, or job, cannot be fetched, or acquired anymore unless the retries are increased again. Can not be null. | [optional] 

## Methods

### NewRetriesDto

`func NewRetriesDto() *RetriesDto`

NewRetriesDto instantiates a new RetriesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetriesDtoWithDefaults

`func NewRetriesDtoWithDefaults() *RetriesDto`

NewRetriesDtoWithDefaults instantiates a new RetriesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetries

`func (o *RetriesDto) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *RetriesDto) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *RetriesDto) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *RetriesDto) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *RetriesDto) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *RetriesDto) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


