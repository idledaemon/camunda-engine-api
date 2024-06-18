# SetRemovalTimeToHistoricProcessInstancesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteRemovalTime** | Pointer to **NullableTime** | The date for which the instances shall be removed. Value may not be &#x60;null&#x60;.  **Note:** Cannot be set in conjunction with &#x60;clearedRemovalTime&#x60; or &#x60;calculatedRemovalTime&#x60;. | [optional] 
**ClearedRemovalTime** | Pointer to **NullableBool** | Sets the removal time to &#x60;null&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior.  **Note:** Cannot be set in conjunction with &#x60;absoluteRemovalTime&#x60; or &#x60;calculatedRemovalTime&#x60;. | [optional] 
**CalculatedRemovalTime** | Pointer to **NullableBool** | The removal time is calculated based on the engine&#39;s configuration settings. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior.  **Note:** Cannot be set in conjunction with &#x60;absoluteRemovalTime&#x60; or &#x60;clearedRemovalTime&#x60;. | [optional] 
**HistoricProcessInstanceIds** | Pointer to **[]string** | The id of the process instance. | [optional] 
**HistoricProcessInstanceQuery** | Pointer to [**HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) |  | [optional] 
**Hierarchical** | Pointer to **NullableBool** | Sets the removal time to all historic process instances in the hierarchy. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**UpdateInChunks** | Pointer to **NullableBool** | Handles removal time updates in chunks, taking into account the defined size in &#x60;removalTimeUpdateChunkSize&#x60; in the process engine configuration. The size of the  chunks can also be overridden per call with the &#x60;updateChunkSize&#x60; parameter. Enabling this option can lead to multiple executions of the resulting jobs, preventing the database transaction from timing out by limiting the number of rows to update. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**UpdateChunkSize** | Pointer to **NullableInt32** | Defines the size of the chunks in which removal time updates are processed. The value must be a positive integer between &#x60;1&#x60; and &#x60;500&#x60;. This only has an  effect if &#x60;updateInChunks&#x60; is set to &#x60;true&#x60;. If undefined, the operation uses the  &#x60;removalTimeUpdateChunkSize&#x60; defined in the process engine configuration. | [optional] 

## Methods

### NewSetRemovalTimeToHistoricProcessInstancesDto

`func NewSetRemovalTimeToHistoricProcessInstancesDto() *SetRemovalTimeToHistoricProcessInstancesDto`

NewSetRemovalTimeToHistoricProcessInstancesDto instantiates a new SetRemovalTimeToHistoricProcessInstancesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRemovalTimeToHistoricProcessInstancesDtoWithDefaults

`func NewSetRemovalTimeToHistoricProcessInstancesDtoWithDefaults() *SetRemovalTimeToHistoricProcessInstancesDto`

NewSetRemovalTimeToHistoricProcessInstancesDtoWithDefaults instantiates a new SetRemovalTimeToHistoricProcessInstancesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteRemovalTime

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetAbsoluteRemovalTime() time.Time`

GetAbsoluteRemovalTime returns the AbsoluteRemovalTime field if non-nil, zero value otherwise.

### GetAbsoluteRemovalTimeOk

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetAbsoluteRemovalTimeOk() (*time.Time, bool)`

GetAbsoluteRemovalTimeOk returns a tuple with the AbsoluteRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteRemovalTime

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetAbsoluteRemovalTime(v time.Time)`

SetAbsoluteRemovalTime sets AbsoluteRemovalTime field to given value.

### HasAbsoluteRemovalTime

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) HasAbsoluteRemovalTime() bool`

HasAbsoluteRemovalTime returns a boolean if a field has been set.

### SetAbsoluteRemovalTimeNil

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetAbsoluteRemovalTimeNil(b bool)`

 SetAbsoluteRemovalTimeNil sets the value for AbsoluteRemovalTime to be an explicit nil

### UnsetAbsoluteRemovalTime
`func (o *SetRemovalTimeToHistoricProcessInstancesDto) UnsetAbsoluteRemovalTime()`

UnsetAbsoluteRemovalTime ensures that no value is present for AbsoluteRemovalTime, not even an explicit nil
### GetClearedRemovalTime

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetClearedRemovalTime() bool`

GetClearedRemovalTime returns the ClearedRemovalTime field if non-nil, zero value otherwise.

### GetClearedRemovalTimeOk

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetClearedRemovalTimeOk() (*bool, bool)`

GetClearedRemovalTimeOk returns a tuple with the ClearedRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearedRemovalTime

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetClearedRemovalTime(v bool)`

SetClearedRemovalTime sets ClearedRemovalTime field to given value.

### HasClearedRemovalTime

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) HasClearedRemovalTime() bool`

HasClearedRemovalTime returns a boolean if a field has been set.

### SetClearedRemovalTimeNil

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetClearedRemovalTimeNil(b bool)`

 SetClearedRemovalTimeNil sets the value for ClearedRemovalTime to be an explicit nil

### UnsetClearedRemovalTime
`func (o *SetRemovalTimeToHistoricProcessInstancesDto) UnsetClearedRemovalTime()`

UnsetClearedRemovalTime ensures that no value is present for ClearedRemovalTime, not even an explicit nil
### GetCalculatedRemovalTime

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetCalculatedRemovalTime() bool`

GetCalculatedRemovalTime returns the CalculatedRemovalTime field if non-nil, zero value otherwise.

### GetCalculatedRemovalTimeOk

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetCalculatedRemovalTimeOk() (*bool, bool)`

GetCalculatedRemovalTimeOk returns a tuple with the CalculatedRemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedRemovalTime

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetCalculatedRemovalTime(v bool)`

SetCalculatedRemovalTime sets CalculatedRemovalTime field to given value.

### HasCalculatedRemovalTime

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) HasCalculatedRemovalTime() bool`

HasCalculatedRemovalTime returns a boolean if a field has been set.

### SetCalculatedRemovalTimeNil

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetCalculatedRemovalTimeNil(b bool)`

 SetCalculatedRemovalTimeNil sets the value for CalculatedRemovalTime to be an explicit nil

### UnsetCalculatedRemovalTime
`func (o *SetRemovalTimeToHistoricProcessInstancesDto) UnsetCalculatedRemovalTime()`

UnsetCalculatedRemovalTime ensures that no value is present for CalculatedRemovalTime, not even an explicit nil
### GetHistoricProcessInstanceIds

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetHistoricProcessInstanceIds() []string`

GetHistoricProcessInstanceIds returns the HistoricProcessInstanceIds field if non-nil, zero value otherwise.

### GetHistoricProcessInstanceIdsOk

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetHistoricProcessInstanceIdsOk() (*[]string, bool)`

GetHistoricProcessInstanceIdsOk returns a tuple with the HistoricProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricProcessInstanceIds

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetHistoricProcessInstanceIds(v []string)`

SetHistoricProcessInstanceIds sets HistoricProcessInstanceIds field to given value.

### HasHistoricProcessInstanceIds

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) HasHistoricProcessInstanceIds() bool`

HasHistoricProcessInstanceIds returns a boolean if a field has been set.

### SetHistoricProcessInstanceIdsNil

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetHistoricProcessInstanceIdsNil(b bool)`

 SetHistoricProcessInstanceIdsNil sets the value for HistoricProcessInstanceIds to be an explicit nil

### UnsetHistoricProcessInstanceIds
`func (o *SetRemovalTimeToHistoricProcessInstancesDto) UnsetHistoricProcessInstanceIds()`

UnsetHistoricProcessInstanceIds ensures that no value is present for HistoricProcessInstanceIds, not even an explicit nil
### GetHistoricProcessInstanceQuery

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto`

GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field if non-nil, zero value otherwise.

### GetHistoricProcessInstanceQueryOk

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool)`

GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricProcessInstanceQuery

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto)`

SetHistoricProcessInstanceQuery sets HistoricProcessInstanceQuery field to given value.

### HasHistoricProcessInstanceQuery

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) HasHistoricProcessInstanceQuery() bool`

HasHistoricProcessInstanceQuery returns a boolean if a field has been set.

### GetHierarchical

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetHierarchical() bool`

GetHierarchical returns the Hierarchical field if non-nil, zero value otherwise.

### GetHierarchicalOk

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetHierarchicalOk() (*bool, bool)`

GetHierarchicalOk returns a tuple with the Hierarchical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchical

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetHierarchical(v bool)`

SetHierarchical sets Hierarchical field to given value.

### HasHierarchical

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) HasHierarchical() bool`

HasHierarchical returns a boolean if a field has been set.

### SetHierarchicalNil

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetHierarchicalNil(b bool)`

 SetHierarchicalNil sets the value for Hierarchical to be an explicit nil

### UnsetHierarchical
`func (o *SetRemovalTimeToHistoricProcessInstancesDto) UnsetHierarchical()`

UnsetHierarchical ensures that no value is present for Hierarchical, not even an explicit nil
### GetUpdateInChunks

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetUpdateInChunks() bool`

GetUpdateInChunks returns the UpdateInChunks field if non-nil, zero value otherwise.

### GetUpdateInChunksOk

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetUpdateInChunksOk() (*bool, bool)`

GetUpdateInChunksOk returns a tuple with the UpdateInChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInChunks

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetUpdateInChunks(v bool)`

SetUpdateInChunks sets UpdateInChunks field to given value.

### HasUpdateInChunks

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) HasUpdateInChunks() bool`

HasUpdateInChunks returns a boolean if a field has been set.

### SetUpdateInChunksNil

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetUpdateInChunksNil(b bool)`

 SetUpdateInChunksNil sets the value for UpdateInChunks to be an explicit nil

### UnsetUpdateInChunks
`func (o *SetRemovalTimeToHistoricProcessInstancesDto) UnsetUpdateInChunks()`

UnsetUpdateInChunks ensures that no value is present for UpdateInChunks, not even an explicit nil
### GetUpdateChunkSize

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetUpdateChunkSize() int32`

GetUpdateChunkSize returns the UpdateChunkSize field if non-nil, zero value otherwise.

### GetUpdateChunkSizeOk

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) GetUpdateChunkSizeOk() (*int32, bool)`

GetUpdateChunkSizeOk returns a tuple with the UpdateChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateChunkSize

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetUpdateChunkSize(v int32)`

SetUpdateChunkSize sets UpdateChunkSize field to given value.

### HasUpdateChunkSize

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) HasUpdateChunkSize() bool`

HasUpdateChunkSize returns a boolean if a field has been set.

### SetUpdateChunkSizeNil

`func (o *SetRemovalTimeToHistoricProcessInstancesDto) SetUpdateChunkSizeNil(b bool)`

 SetUpdateChunkSizeNil sets the value for UpdateChunkSize to be an explicit nil

### UnsetUpdateChunkSize
`func (o *SetRemovalTimeToHistoricProcessInstancesDto) UnsetUpdateChunkSize()`

UnsetUpdateChunkSize ensures that no value is present for UpdateChunkSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


