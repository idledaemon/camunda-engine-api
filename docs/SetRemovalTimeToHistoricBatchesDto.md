# SetRemovalTimeToHistoricBatchesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteRemovalTime** | Pointer to **NullableTime** | The date for which the instances shall be removed. Value may not be &#x60;null&#x60;.  **Note:** Cannot be set in conjunction with &#x60;clearedRemovalTime&#x60; or &#x60;calculatedRemovalTime&#x60;. | [optional] 
**ClearedRemovalTime** | Pointer to **NullableBool** | Sets the removal time to &#x60;null&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior.  **Note:** Cannot be set in conjunction with &#x60;absoluteRemovalTime&#x60; or &#x60;calculatedRemovalTime&#x60;. | [optional] 
**CalculatedRemovalTime** | Pointer to **NullableBool** | The removal time is calculated based on the engine&#39;s configuration settings. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior.  **Note:** Cannot be set in conjunction with &#x60;absoluteRemovalTime&#x60; or &#x60;clearedRemovalTime&#x60;. | [optional] 
**HistoricBatchQuery** | Pointer to **map[string]interface{}** | Query for the historic batches to set the removal time for. | [optional] 
**HistoricBatchIds** | Pointer to **[]string** | The ids of the historic batches to set the removal time for. | [optional] 

## Methods

### NewSetRemovalTimeToHistoricBatchesDto

`func NewSetRemovalTimeToHistoricBatchesDto() *SetRemovalTimeToHistoricBatchesDto`

NewSetRemovalTimeToHistoricBatchesDto instantiates a new SetRemovalTimeToHistoricBatchesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRemovalTimeToHistoricBatchesDtoWithDefaults

`func NewSetRemovalTimeToHistoricBatchesDtoWithDefaults() *SetRemovalTimeToHistoricBatchesDto`

NewSetRemovalTimeToHistoricBatchesDtoWithDefaults instantiates a new SetRemovalTimeToHistoricBatchesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteRemovalTime

`func (o *SetRemovalTimeToHistoricBatchesDto) GetAbsoluteRemovalTime() time.Time`

GetAbsoluteRemovalTime returns the AbsoluteRemovalTime field if non-nil, zero value otherwise.

### GetAbsoluteRemovalTimeOk

`func (o *SetRemovalTimeToHistoricBatchesDto) GetAbsoluteRemovalTimeOk() (*time.Time, bool)`

GetAbsoluteRemovalTimeOk returns a tuple with the AbsoluteRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteRemovalTime

`func (o *SetRemovalTimeToHistoricBatchesDto) SetAbsoluteRemovalTime(v time.Time)`

SetAbsoluteRemovalTime sets AbsoluteRemovalTime field to given value.

### HasAbsoluteRemovalTime

`func (o *SetRemovalTimeToHistoricBatchesDto) HasAbsoluteRemovalTime() bool`

HasAbsoluteRemovalTime returns a boolean if a field has been set.

### SetAbsoluteRemovalTimeNil

`func (o *SetRemovalTimeToHistoricBatchesDto) SetAbsoluteRemovalTimeNil(b bool)`

 SetAbsoluteRemovalTimeNil sets the value for AbsoluteRemovalTime to be an explicit nil

### UnsetAbsoluteRemovalTime
`func (o *SetRemovalTimeToHistoricBatchesDto) UnsetAbsoluteRemovalTime()`

UnsetAbsoluteRemovalTime ensures that no value is present for AbsoluteRemovalTime, not even an explicit nil
### GetClearedRemovalTime

`func (o *SetRemovalTimeToHistoricBatchesDto) GetClearedRemovalTime() bool`

GetClearedRemovalTime returns the ClearedRemovalTime field if non-nil, zero value otherwise.

### GetClearedRemovalTimeOk

`func (o *SetRemovalTimeToHistoricBatchesDto) GetClearedRemovalTimeOk() (*bool, bool)`

GetClearedRemovalTimeOk returns a tuple with the ClearedRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearedRemovalTime

`func (o *SetRemovalTimeToHistoricBatchesDto) SetClearedRemovalTime(v bool)`

SetClearedRemovalTime sets ClearedRemovalTime field to given value.

### HasClearedRemovalTime

`func (o *SetRemovalTimeToHistoricBatchesDto) HasClearedRemovalTime() bool`

HasClearedRemovalTime returns a boolean if a field has been set.

### SetClearedRemovalTimeNil

`func (o *SetRemovalTimeToHistoricBatchesDto) SetClearedRemovalTimeNil(b bool)`

 SetClearedRemovalTimeNil sets the value for ClearedRemovalTime to be an explicit nil

### UnsetClearedRemovalTime
`func (o *SetRemovalTimeToHistoricBatchesDto) UnsetClearedRemovalTime()`

UnsetClearedRemovalTime ensures that no value is present for ClearedRemovalTime, not even an explicit nil
### GetCalculatedRemovalTime

`func (o *SetRemovalTimeToHistoricBatchesDto) GetCalculatedRemovalTime() bool`

GetCalculatedRemovalTime returns the CalculatedRemovalTime field if non-nil, zero value otherwise.

### GetCalculatedRemovalTimeOk

`func (o *SetRemovalTimeToHistoricBatchesDto) GetCalculatedRemovalTimeOk() (*bool, bool)`

GetCalculatedRemovalTimeOk returns a tuple with the CalculatedRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedRemovalTime

`func (o *SetRemovalTimeToHistoricBatchesDto) SetCalculatedRemovalTime(v bool)`

SetCalculatedRemovalTime sets CalculatedRemovalTime field to given value.

### HasCalculatedRemovalTime

`func (o *SetRemovalTimeToHistoricBatchesDto) HasCalculatedRemovalTime() bool`

HasCalculatedRemovalTime returns a boolean if a field has been set.

### SetCalculatedRemovalTimeNil

`func (o *SetRemovalTimeToHistoricBatchesDto) SetCalculatedRemovalTimeNil(b bool)`

 SetCalculatedRemovalTimeNil sets the value for CalculatedRemovalTime to be an explicit nil

### UnsetCalculatedRemovalTime
`func (o *SetRemovalTimeToHistoricBatchesDto) UnsetCalculatedRemovalTime()`

UnsetCalculatedRemovalTime ensures that no value is present for CalculatedRemovalTime, not even an explicit nil
### GetHistoricBatchQuery

`func (o *SetRemovalTimeToHistoricBatchesDto) GetHistoricBatchQuery() map[string]interface{}`

GetHistoricBatchQuery returns the HistoricBatchQuery field if non-nil, zero value otherwise.

### GetHistoricBatchQueryOk

`func (o *SetRemovalTimeToHistoricBatchesDto) GetHistoricBatchQueryOk() (*map[string]interface{}, bool)`

GetHistoricBatchQueryOk returns a tuple with the HistoricBatchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricBatchQuery

`func (o *SetRemovalTimeToHistoricBatchesDto) SetHistoricBatchQuery(v map[string]interface{})`

SetHistoricBatchQuery sets HistoricBatchQuery field to given value.

### HasHistoricBatchQuery

`func (o *SetRemovalTimeToHistoricBatchesDto) HasHistoricBatchQuery() bool`

HasHistoricBatchQuery returns a boolean if a field has been set.

### SetHistoricBatchQueryNil

`func (o *SetRemovalTimeToHistoricBatchesDto) SetHistoricBatchQueryNil(b bool)`

 SetHistoricBatchQueryNil sets the value for HistoricBatchQuery to be an explicit nil

### UnsetHistoricBatchQuery
`func (o *SetRemovalTimeToHistoricBatchesDto) UnsetHistoricBatchQuery()`

UnsetHistoricBatchQuery ensures that no value is present for HistoricBatchQuery, not even an explicit nil
### GetHistoricBatchIds

`func (o *SetRemovalTimeToHistoricBatchesDto) GetHistoricBatchIds() []string`

GetHistoricBatchIds returns the HistoricBatchIds field if non-nil, zero value otherwise.

### GetHistoricBatchIdsOk

`func (o *SetRemovalTimeToHistoricBatchesDto) GetHistoricBatchIdsOk() (*[]string, bool)`

GetHistoricBatchIdsOk returns a tuple with the HistoricBatchIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricBatchIds

`func (o *SetRemovalTimeToHistoricBatchesDto) SetHistoricBatchIds(v []string)`

SetHistoricBatchIds sets HistoricBatchIds field to given value.

### HasHistoricBatchIds

`func (o *SetRemovalTimeToHistoricBatchesDto) HasHistoricBatchIds() bool`

HasHistoricBatchIds returns a boolean if a field has been set.

### SetHistoricBatchIdsNil

`func (o *SetRemovalTimeToHistoricBatchesDto) SetHistoricBatchIdsNil(b bool)`

 SetHistoricBatchIdsNil sets the value for HistoricBatchIds to be an explicit nil

### UnsetHistoricBatchIds
`func (o *SetRemovalTimeToHistoricBatchesDto) UnsetHistoricBatchIds()`

UnsetHistoricBatchIds ensures that no value is present for HistoricBatchIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


