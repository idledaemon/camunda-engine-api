# AbstractSetRemovalTimeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteRemovalTime** | Pointer to **NullableTime** | The date for which the instances shall be removed. Value may not be &#x60;null&#x60;.  **Note:** Cannot be set in conjunction with &#x60;clearedRemovalTime&#x60; or &#x60;calculatedRemovalTime&#x60;. | [optional] 
**ClearedRemovalTime** | Pointer to **NullableBool** | Sets the removal time to &#x60;null&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior.  **Note:** Cannot be set in conjunction with &#x60;absoluteRemovalTime&#x60; or &#x60;calculatedRemovalTime&#x60;. | [optional] 
**CalculatedRemovalTime** | Pointer to **NullableBool** | The removal time is calculated based on the engine&#39;s configuration settings. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior.  **Note:** Cannot be set in conjunction with &#x60;absoluteRemovalTime&#x60; or &#x60;clearedRemovalTime&#x60;. | [optional] 

## Methods

### NewAbstractSetRemovalTimeDto

`func NewAbstractSetRemovalTimeDto() *AbstractSetRemovalTimeDto`

NewAbstractSetRemovalTimeDto instantiates a new AbstractSetRemovalTimeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractSetRemovalTimeDtoWithDefaults

`func NewAbstractSetRemovalTimeDtoWithDefaults() *AbstractSetRemovalTimeDto`

NewAbstractSetRemovalTimeDtoWithDefaults instantiates a new AbstractSetRemovalTimeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteRemovalTime

`func (o *AbstractSetRemovalTimeDto) GetAbsoluteRemovalTime() time.Time`

GetAbsoluteRemovalTime returns the AbsoluteRemovalTime field if non-nil, zero value otherwise.

### GetAbsoluteRemovalTimeOk

`func (o *AbstractSetRemovalTimeDto) GetAbsoluteRemovalTimeOk() (*time.Time, bool)`

GetAbsoluteRemovalTimeOk returns a tuple with the AbsoluteRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteRemovalTime

`func (o *AbstractSetRemovalTimeDto) SetAbsoluteRemovalTime(v time.Time)`

SetAbsoluteRemovalTime sets AbsoluteRemovalTime field to given value.

### HasAbsoluteRemovalTime

`func (o *AbstractSetRemovalTimeDto) HasAbsoluteRemovalTime() bool`

HasAbsoluteRemovalTime returns a boolean if a field has been set.

### SetAbsoluteRemovalTimeNil

`func (o *AbstractSetRemovalTimeDto) SetAbsoluteRemovalTimeNil(b bool)`

 SetAbsoluteRemovalTimeNil sets the value for AbsoluteRemovalTime to be an explicit nil

### UnsetAbsoluteRemovalTime
`func (o *AbstractSetRemovalTimeDto) UnsetAbsoluteRemovalTime()`

UnsetAbsoluteRemovalTime ensures that no value is present for AbsoluteRemovalTime, not even an explicit nil
### GetClearedRemovalTime

`func (o *AbstractSetRemovalTimeDto) GetClearedRemovalTime() bool`

GetClearedRemovalTime returns the ClearedRemovalTime field if non-nil, zero value otherwise.

### GetClearedRemovalTimeOk

`func (o *AbstractSetRemovalTimeDto) GetClearedRemovalTimeOk() (*bool, bool)`

GetClearedRemovalTimeOk returns a tuple with the ClearedRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearedRemovalTime

`func (o *AbstractSetRemovalTimeDto) SetClearedRemovalTime(v bool)`

SetClearedRemovalTime sets ClearedRemovalTime field to given value.

### HasClearedRemovalTime

`func (o *AbstractSetRemovalTimeDto) HasClearedRemovalTime() bool`

HasClearedRemovalTime returns a boolean if a field has been set.

### SetClearedRemovalTimeNil

`func (o *AbstractSetRemovalTimeDto) SetClearedRemovalTimeNil(b bool)`

 SetClearedRemovalTimeNil sets the value for ClearedRemovalTime to be an explicit nil

### UnsetClearedRemovalTime
`func (o *AbstractSetRemovalTimeDto) UnsetClearedRemovalTime()`

UnsetClearedRemovalTime ensures that no value is present for ClearedRemovalTime, not even an explicit nil
### GetCalculatedRemovalTime

`func (o *AbstractSetRemovalTimeDto) GetCalculatedRemovalTime() bool`

GetCalculatedRemovalTime returns the CalculatedRemovalTime field if non-nil, zero value otherwise.

### GetCalculatedRemovalTimeOk

`func (o *AbstractSetRemovalTimeDto) GetCalculatedRemovalTimeOk() (*bool, bool)`

GetCalculatedRemovalTimeOk returns a tuple with the CalculatedRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedRemovalTime

`func (o *AbstractSetRemovalTimeDto) SetCalculatedRemovalTime(v bool)`

SetCalculatedRemovalTime sets CalculatedRemovalTime field to given value.

### HasCalculatedRemovalTime

`func (o *AbstractSetRemovalTimeDto) HasCalculatedRemovalTime() bool`

HasCalculatedRemovalTime returns a boolean if a field has been set.

### SetCalculatedRemovalTimeNil

`func (o *AbstractSetRemovalTimeDto) SetCalculatedRemovalTimeNil(b bool)`

 SetCalculatedRemovalTimeNil sets the value for CalculatedRemovalTime to be an explicit nil

### UnsetCalculatedRemovalTime
`func (o *AbstractSetRemovalTimeDto) UnsetCalculatedRemovalTime()`

UnsetCalculatedRemovalTime ensures that no value is present for CalculatedRemovalTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


