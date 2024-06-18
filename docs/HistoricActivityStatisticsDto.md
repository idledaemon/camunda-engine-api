# HistoricActivityStatisticsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the activity the results are aggregated for. | [optional] 
**Instances** | Pointer to **NullableInt64** | The total number of all running instances of the activity. | [optional] 
**Canceled** | Pointer to **NullableInt64** | The total number of all canceled instances of the activity. **Note:** Will be &#x60;0&#x60; (not &#x60;null&#x60;), if canceled activity instances were excluded. | [optional] 
**Finished** | Pointer to **NullableInt64** | The total number of all finished instances of the activity. **Note:** Will be &#x60;0&#x60; (not &#x60;null&#x60;), if finished activity instances were excluded. | [optional] 
**CompleteScope** | Pointer to **NullableInt64** | The total number of all instances which completed a scope of the activity. **Note:** Will be &#x60;0&#x60; (not &#x60;null&#x60;), if activity instances which completed a scope were excluded. | [optional] 
**OpenIncidents** | Pointer to **NullableInt64** | The total number of open incidents for the activity. **Note:** Will be &#x60;0&#x60; (not &#x60;null&#x60;), if &#x60;incidents&#x60; is set to &#x60;false&#x60;. | [optional] 
**ResolvedIncidents** | Pointer to **NullableInt64** | The total number of resolved incidents for the activity. **Note:** Will be &#x60;0&#x60; (not &#x60;null&#x60;), if &#x60;incidents&#x60; is set to &#x60;false&#x60;. | [optional] 
**DeletedIncidents** | Pointer to **NullableInt64** | The total number of deleted incidents for the activity. **Note:** Will be &#x60;0&#x60; (not &#x60;null&#x60;), if &#x60;incidents&#x60; is set to &#x60;false&#x60;. | [optional] 

## Methods

### NewHistoricActivityStatisticsDto

`func NewHistoricActivityStatisticsDto() *HistoricActivityStatisticsDto`

NewHistoricActivityStatisticsDto instantiates a new HistoricActivityStatisticsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricActivityStatisticsDtoWithDefaults

`func NewHistoricActivityStatisticsDtoWithDefaults() *HistoricActivityStatisticsDto`

NewHistoricActivityStatisticsDtoWithDefaults instantiates a new HistoricActivityStatisticsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricActivityStatisticsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricActivityStatisticsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricActivityStatisticsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricActivityStatisticsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricActivityStatisticsDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricActivityStatisticsDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInstances

`func (o *HistoricActivityStatisticsDto) GetInstances() int64`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *HistoricActivityStatisticsDto) GetInstancesOk() (*int64, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *HistoricActivityStatisticsDto) SetInstances(v int64)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *HistoricActivityStatisticsDto) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstancesNil

`func (o *HistoricActivityStatisticsDto) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *HistoricActivityStatisticsDto) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetCanceled

`func (o *HistoricActivityStatisticsDto) GetCanceled() int64`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *HistoricActivityStatisticsDto) GetCanceledOk() (*int64, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *HistoricActivityStatisticsDto) SetCanceled(v int64)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *HistoricActivityStatisticsDto) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### SetCanceledNil

`func (o *HistoricActivityStatisticsDto) SetCanceledNil(b bool)`

 SetCanceledNil sets the value for Canceled to be an explicit nil

### UnsetCanceled
`func (o *HistoricActivityStatisticsDto) UnsetCanceled()`

UnsetCanceled ensures that no value is present for Canceled, not even an explicit nil
### GetFinished

`func (o *HistoricActivityStatisticsDto) GetFinished() int64`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *HistoricActivityStatisticsDto) GetFinishedOk() (*int64, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *HistoricActivityStatisticsDto) SetFinished(v int64)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *HistoricActivityStatisticsDto) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *HistoricActivityStatisticsDto) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *HistoricActivityStatisticsDto) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetCompleteScope

`func (o *HistoricActivityStatisticsDto) GetCompleteScope() int64`

GetCompleteScope returns the CompleteScope field if non-nil, zero value otherwise.

### GetCompleteScopeOk

`func (o *HistoricActivityStatisticsDto) GetCompleteScopeOk() (*int64, bool)`

GetCompleteScopeOk returns a tuple with the CompleteScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteScope

`func (o *HistoricActivityStatisticsDto) SetCompleteScope(v int64)`

SetCompleteScope sets CompleteScope field to given value.

### HasCompleteScope

`func (o *HistoricActivityStatisticsDto) HasCompleteScope() bool`

HasCompleteScope returns a boolean if a field has been set.

### SetCompleteScopeNil

`func (o *HistoricActivityStatisticsDto) SetCompleteScopeNil(b bool)`

 SetCompleteScopeNil sets the value for CompleteScope to be an explicit nil

### UnsetCompleteScope
`func (o *HistoricActivityStatisticsDto) UnsetCompleteScope()`

UnsetCompleteScope ensures that no value is present for CompleteScope, not even an explicit nil
### GetOpenIncidents

`func (o *HistoricActivityStatisticsDto) GetOpenIncidents() int64`

GetOpenIncidents returns the OpenIncidents field if non-nil, zero value otherwise.

### GetOpenIncidentsOk

`func (o *HistoricActivityStatisticsDto) GetOpenIncidentsOk() (*int64, bool)`

GetOpenIncidentsOk returns a tuple with the OpenIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIncidents

`func (o *HistoricActivityStatisticsDto) SetOpenIncidents(v int64)`

SetOpenIncidents sets OpenIncidents field to given value.

### HasOpenIncidents

`func (o *HistoricActivityStatisticsDto) HasOpenIncidents() bool`

HasOpenIncidents returns a boolean if a field has been set.

### SetOpenIncidentsNil

`func (o *HistoricActivityStatisticsDto) SetOpenIncidentsNil(b bool)`

 SetOpenIncidentsNil sets the value for OpenIncidents to be an explicit nil

### UnsetOpenIncidents
`func (o *HistoricActivityStatisticsDto) UnsetOpenIncidents()`

UnsetOpenIncidents ensures that no value is present for OpenIncidents, not even an explicit nil
### GetResolvedIncidents

`func (o *HistoricActivityStatisticsDto) GetResolvedIncidents() int64`

GetResolvedIncidents returns the ResolvedIncidents field if non-nil, zero value otherwise.

### GetResolvedIncidentsOk

`func (o *HistoricActivityStatisticsDto) GetResolvedIncidentsOk() (*int64, bool)`

GetResolvedIncidentsOk returns a tuple with the ResolvedIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedIncidents

`func (o *HistoricActivityStatisticsDto) SetResolvedIncidents(v int64)`

SetResolvedIncidents sets ResolvedIncidents field to given value.

### HasResolvedIncidents

`func (o *HistoricActivityStatisticsDto) HasResolvedIncidents() bool`

HasResolvedIncidents returns a boolean if a field has been set.

### SetResolvedIncidentsNil

`func (o *HistoricActivityStatisticsDto) SetResolvedIncidentsNil(b bool)`

 SetResolvedIncidentsNil sets the value for ResolvedIncidents to be an explicit nil

### UnsetResolvedIncidents
`func (o *HistoricActivityStatisticsDto) UnsetResolvedIncidents()`

UnsetResolvedIncidents ensures that no value is present for ResolvedIncidents, not even an explicit nil
### GetDeletedIncidents

`func (o *HistoricActivityStatisticsDto) GetDeletedIncidents() int64`

GetDeletedIncidents returns the DeletedIncidents field if non-nil, zero value otherwise.

### GetDeletedIncidentsOk

`func (o *HistoricActivityStatisticsDto) GetDeletedIncidentsOk() (*int64, bool)`

GetDeletedIncidentsOk returns a tuple with the DeletedIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedIncidents

`func (o *HistoricActivityStatisticsDto) SetDeletedIncidents(v int64)`

SetDeletedIncidents sets DeletedIncidents field to given value.

### HasDeletedIncidents

`func (o *HistoricActivityStatisticsDto) HasDeletedIncidents() bool`

HasDeletedIncidents returns a boolean if a field has been set.

### SetDeletedIncidentsNil

`func (o *HistoricActivityStatisticsDto) SetDeletedIncidentsNil(b bool)`

 SetDeletedIncidentsNil sets the value for DeletedIncidents to be an explicit nil

### UnsetDeletedIncidents
`func (o *HistoricActivityStatisticsDto) UnsetDeletedIncidents()`

UnsetDeletedIncidents ensures that no value is present for DeletedIncidents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


