# SetRemovalTimeToHistoricDecisionInstancesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteRemovalTime** | Pointer to **NullableTime** | The date for which the instances shall be removed. Value may not be &#x60;null&#x60;.  **Note:** Cannot be set in conjunction with &#x60;clearedRemovalTime&#x60; or &#x60;calculatedRemovalTime&#x60;. | [optional] 
**ClearedRemovalTime** | Pointer to **NullableBool** | Sets the removal time to &#x60;null&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior.  **Note:** Cannot be set in conjunction with &#x60;absoluteRemovalTime&#x60; or &#x60;calculatedRemovalTime&#x60;. | [optional] 
**CalculatedRemovalTime** | Pointer to **NullableBool** | The removal time is calculated based on the engine&#39;s configuration settings. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior.  **Note:** Cannot be set in conjunction with &#x60;absoluteRemovalTime&#x60; or &#x60;clearedRemovalTime&#x60;. | [optional] 
**Hierarchical** | Pointer to **NullableBool** | Sets the removal time to all historic decision instances in the hierarchy. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**HistoricDecisionInstanceQuery** | Pointer to [**HistoricDecisionInstanceQueryDto**](HistoricDecisionInstanceQueryDto.md) |  | [optional] 
**HistoricDecisionInstanceIds** | Pointer to **[]string** | The ids of the historic decision instances to set the removal time for. | [optional] 

## Methods

### NewSetRemovalTimeToHistoricDecisionInstancesDto

`func NewSetRemovalTimeToHistoricDecisionInstancesDto() *SetRemovalTimeToHistoricDecisionInstancesDto`

NewSetRemovalTimeToHistoricDecisionInstancesDto instantiates a new SetRemovalTimeToHistoricDecisionInstancesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRemovalTimeToHistoricDecisionInstancesDtoWithDefaults

`func NewSetRemovalTimeToHistoricDecisionInstancesDtoWithDefaults() *SetRemovalTimeToHistoricDecisionInstancesDto`

NewSetRemovalTimeToHistoricDecisionInstancesDtoWithDefaults instantiates a new SetRemovalTimeToHistoricDecisionInstancesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteRemovalTime

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetAbsoluteRemovalTime() time.Time`

GetAbsoluteRemovalTime returns the AbsoluteRemovalTime field if non-nil, zero value otherwise.

### GetAbsoluteRemovalTimeOk

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetAbsoluteRemovalTimeOk() (*time.Time, bool)`

GetAbsoluteRemovalTimeOk returns a tuple with the AbsoluteRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteRemovalTime

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) SetAbsoluteRemovalTime(v time.Time)`

SetAbsoluteRemovalTime sets AbsoluteRemovalTime field to given value.

### HasAbsoluteRemovalTime

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) HasAbsoluteRemovalTime() bool`

HasAbsoluteRemovalTime returns a boolean if a field has been set.

### SetAbsoluteRemovalTimeNil

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) SetAbsoluteRemovalTimeNil(b bool)`

 SetAbsoluteRemovalTimeNil sets the value for AbsoluteRemovalTime to be an explicit nil

### UnsetAbsoluteRemovalTime
`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) UnsetAbsoluteRemovalTime()`

UnsetAbsoluteRemovalTime ensures that no value is present for AbsoluteRemovalTime, not even an explicit nil
### GetClearedRemovalTime

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetClearedRemovalTime() bool`

GetClearedRemovalTime returns the ClearedRemovalTime field if non-nil, zero value otherwise.

### GetClearedRemovalTimeOk

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetClearedRemovalTimeOk() (*bool, bool)`

GetClearedRemovalTimeOk returns a tuple with the ClearedRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearedRemovalTime

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) SetClearedRemovalTime(v bool)`

SetClearedRemovalTime sets ClearedRemovalTime field to given value.

### HasClearedRemovalTime

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) HasClearedRemovalTime() bool`

HasClearedRemovalTime returns a boolean if a field has been set.

### SetClearedRemovalTimeNil

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) SetClearedRemovalTimeNil(b bool)`

 SetClearedRemovalTimeNil sets the value for ClearedRemovalTime to be an explicit nil

### UnsetClearedRemovalTime
`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) UnsetClearedRemovalTime()`

UnsetClearedRemovalTime ensures that no value is present for ClearedRemovalTime, not even an explicit nil
### GetCalculatedRemovalTime

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetCalculatedRemovalTime() bool`

GetCalculatedRemovalTime returns the CalculatedRemovalTime field if non-nil, zero value otherwise.

### GetCalculatedRemovalTimeOk

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetCalculatedRemovalTimeOk() (*bool, bool)`

GetCalculatedRemovalTimeOk returns a tuple with the CalculatedRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedRemovalTime

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) SetCalculatedRemovalTime(v bool)`

SetCalculatedRemovalTime sets CalculatedRemovalTime field to given value.

### HasCalculatedRemovalTime

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) HasCalculatedRemovalTime() bool`

HasCalculatedRemovalTime returns a boolean if a field has been set.

### SetCalculatedRemovalTimeNil

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) SetCalculatedRemovalTimeNil(b bool)`

 SetCalculatedRemovalTimeNil sets the value for CalculatedRemovalTime to be an explicit nil

### UnsetCalculatedRemovalTime
`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) UnsetCalculatedRemovalTime()`

UnsetCalculatedRemovalTime ensures that no value is present for CalculatedRemovalTime, not even an explicit nil
### GetHierarchical

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetHierarchical() bool`

GetHierarchical returns the Hierarchical field if non-nil, zero value otherwise.

### GetHierarchicalOk

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetHierarchicalOk() (*bool, bool)`

GetHierarchicalOk returns a tuple with the Hierarchical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchical

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) SetHierarchical(v bool)`

SetHierarchical sets Hierarchical field to given value.

### HasHierarchical

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) HasHierarchical() bool`

HasHierarchical returns a boolean if a field has been set.

### SetHierarchicalNil

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) SetHierarchicalNil(b bool)`

 SetHierarchicalNil sets the value for Hierarchical to be an explicit nil

### UnsetHierarchical
`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) UnsetHierarchical()`

UnsetHierarchical ensures that no value is present for Hierarchical, not even an explicit nil
### GetHistoricDecisionInstanceQuery

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetHistoricDecisionInstanceQuery() HistoricDecisionInstanceQueryDto`

GetHistoricDecisionInstanceQuery returns the HistoricDecisionInstanceQuery field if non-nil, zero value otherwise.

### GetHistoricDecisionInstanceQueryOk

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetHistoricDecisionInstanceQueryOk() (*HistoricDecisionInstanceQueryDto, bool)`

GetHistoricDecisionInstanceQueryOk returns a tuple with the HistoricDecisionInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricDecisionInstanceQuery

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) SetHistoricDecisionInstanceQuery(v HistoricDecisionInstanceQueryDto)`

SetHistoricDecisionInstanceQuery sets HistoricDecisionInstanceQuery field to given value.

### HasHistoricDecisionInstanceQuery

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) HasHistoricDecisionInstanceQuery() bool`

HasHistoricDecisionInstanceQuery returns a boolean if a field has been set.

### GetHistoricDecisionInstanceIds

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetHistoricDecisionInstanceIds() []string`

GetHistoricDecisionInstanceIds returns the HistoricDecisionInstanceIds field if non-nil, zero value otherwise.

### GetHistoricDecisionInstanceIdsOk

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) GetHistoricDecisionInstanceIdsOk() (*[]string, bool)`

GetHistoricDecisionInstanceIdsOk returns a tuple with the HistoricDecisionInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricDecisionInstanceIds

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) SetHistoricDecisionInstanceIds(v []string)`

SetHistoricDecisionInstanceIds sets HistoricDecisionInstanceIds field to given value.

### HasHistoricDecisionInstanceIds

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) HasHistoricDecisionInstanceIds() bool`

HasHistoricDecisionInstanceIds returns a boolean if a field has been set.

### SetHistoricDecisionInstanceIdsNil

`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) SetHistoricDecisionInstanceIdsNil(b bool)`

 SetHistoricDecisionInstanceIdsNil sets the value for HistoricDecisionInstanceIds to be an explicit nil

### UnsetHistoricDecisionInstanceIds
`func (o *SetRemovalTimeToHistoricDecisionInstancesDto) UnsetHistoricDecisionInstanceIds()`

UnsetHistoricDecisionInstanceIds ensures that no value is present for HistoricDecisionInstanceIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


