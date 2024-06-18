# HistoryTimeToLiveDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryTimeToLive** | Pointer to **NullableInt32** | New value for historyTimeToLive field of the definition. Can be &#x60;null&#x60; if &#x60;enforceHistoryTimeToLive&#x60; is configured to &#x60;false&#x60;. Cannot be negative. | [optional] 

## Methods

### NewHistoryTimeToLiveDto

`func NewHistoryTimeToLiveDto() *HistoryTimeToLiveDto`

NewHistoryTimeToLiveDto instantiates a new HistoryTimeToLiveDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryTimeToLiveDtoWithDefaults

`func NewHistoryTimeToLiveDtoWithDefaults() *HistoryTimeToLiveDto`

NewHistoryTimeToLiveDtoWithDefaults instantiates a new HistoryTimeToLiveDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryTimeToLive

`func (o *HistoryTimeToLiveDto) GetHistoryTimeToLive() int32`

GetHistoryTimeToLive returns the HistoryTimeToLive field if non-nil, zero value otherwise.

### GetHistoryTimeToLiveOk

`func (o *HistoryTimeToLiveDto) GetHistoryTimeToLiveOk() (*int32, bool)`

GetHistoryTimeToLiveOk returns a tuple with the HistoryTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryTimeToLive

`func (o *HistoryTimeToLiveDto) SetHistoryTimeToLive(v int32)`

SetHistoryTimeToLive sets HistoryTimeToLive field to given value.

### HasHistoryTimeToLive

`func (o *HistoryTimeToLiveDto) HasHistoryTimeToLive() bool`

HasHistoryTimeToLive returns a boolean if a field has been set.

### SetHistoryTimeToLiveNil

`func (o *HistoryTimeToLiveDto) SetHistoryTimeToLiveNil(b bool)`

 SetHistoryTimeToLiveNil sets the value for HistoryTimeToLive to be an explicit nil

### UnsetHistoryTimeToLive
`func (o *HistoryTimeToLiveDto) UnsetHistoryTimeToLive()`

UnsetHistoryTimeToLive ensures that no value is present for HistoryTimeToLive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


