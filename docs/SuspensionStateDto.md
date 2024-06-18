# SuspensionStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suspended** | Pointer to **NullableBool** | A Boolean value which indicates whether to activate or suspend a given instance  (e.g. process instance, job, job definition, or batch). When the value is set to true,  the given instance will be suspended and when the value is set to false,  the given instance will be activated. | [optional] 

## Methods

### NewSuspensionStateDto

`func NewSuspensionStateDto() *SuspensionStateDto`

NewSuspensionStateDto instantiates a new SuspensionStateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuspensionStateDtoWithDefaults

`func NewSuspensionStateDtoWithDefaults() *SuspensionStateDto`

NewSuspensionStateDtoWithDefaults instantiates a new SuspensionStateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuspended

`func (o *SuspensionStateDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *SuspensionStateDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *SuspensionStateDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *SuspensionStateDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *SuspensionStateDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *SuspensionStateDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


