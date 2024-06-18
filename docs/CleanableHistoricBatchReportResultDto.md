# CleanableHistoricBatchReportResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchType** | Pointer to **NullableString** | The type of the batch operation. | [optional] 
**HistoryTimeToLive** | Pointer to **NullableInt32** | The history time to live of the batch operation. | [optional] 
**FinishedBatchesCount** | Pointer to **NullableInt64** | The count of the finished batch operations. | [optional] 
**CleanableBatchesCount** | Pointer to **NullableInt64** | The count of the cleanable historic batch operations, referring to history time to live. | [optional] 

## Methods

### NewCleanableHistoricBatchReportResultDto

`func NewCleanableHistoricBatchReportResultDto() *CleanableHistoricBatchReportResultDto`

NewCleanableHistoricBatchReportResultDto instantiates a new CleanableHistoricBatchReportResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanableHistoricBatchReportResultDtoWithDefaults

`func NewCleanableHistoricBatchReportResultDtoWithDefaults() *CleanableHistoricBatchReportResultDto`

NewCleanableHistoricBatchReportResultDtoWithDefaults instantiates a new CleanableHistoricBatchReportResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchType

`func (o *CleanableHistoricBatchReportResultDto) GetBatchType() string`

GetBatchType returns the BatchType field if non-nil, zero value otherwise.

### GetBatchTypeOk

`func (o *CleanableHistoricBatchReportResultDto) GetBatchTypeOk() (*string, bool)`

GetBatchTypeOk returns a tuple with the BatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchType

`func (o *CleanableHistoricBatchReportResultDto) SetBatchType(v string)`

SetBatchType sets BatchType field to given value.

### HasBatchType

`func (o *CleanableHistoricBatchReportResultDto) HasBatchType() bool`

HasBatchType returns a boolean if a field has been set.

### SetBatchTypeNil

`func (o *CleanableHistoricBatchReportResultDto) SetBatchTypeNil(b bool)`

 SetBatchTypeNil sets the value for BatchType to be an explicit nil

### UnsetBatchType
`func (o *CleanableHistoricBatchReportResultDto) UnsetBatchType()`

UnsetBatchType ensures that no value is present for BatchType, not even an explicit nil
### GetHistoryTimeToLive

`func (o *CleanableHistoricBatchReportResultDto) GetHistoryTimeToLive() int32`

GetHistoryTimeToLive returns the HistoryTimeToLive field if non-nil, zero value otherwise.

### GetHistoryTimeToLiveOk

`func (o *CleanableHistoricBatchReportResultDto) GetHistoryTimeToLiveOk() (*int32, bool)`

GetHistoryTimeToLiveOk returns a tuple with the HistoryTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryTimeToLive

`func (o *CleanableHistoricBatchReportResultDto) SetHistoryTimeToLive(v int32)`

SetHistoryTimeToLive sets HistoryTimeToLive field to given value.

### HasHistoryTimeToLive

`func (o *CleanableHistoricBatchReportResultDto) HasHistoryTimeToLive() bool`

HasHistoryTimeToLive returns a boolean if a field has been set.

### SetHistoryTimeToLiveNil

`func (o *CleanableHistoricBatchReportResultDto) SetHistoryTimeToLiveNil(b bool)`

 SetHistoryTimeToLiveNil sets the value for HistoryTimeToLive to be an explicit nil

### UnsetHistoryTimeToLive
`func (o *CleanableHistoricBatchReportResultDto) UnsetHistoryTimeToLive()`

UnsetHistoryTimeToLive ensures that no value is present for HistoryTimeToLive, not even an explicit nil
### GetFinishedBatchesCount

`func (o *CleanableHistoricBatchReportResultDto) GetFinishedBatchesCount() int64`

GetFinishedBatchesCount returns the FinishedBatchesCount field if non-nil, zero value otherwise.

### GetFinishedBatchesCountOk

`func (o *CleanableHistoricBatchReportResultDto) GetFinishedBatchesCountOk() (*int64, bool)`

GetFinishedBatchesCountOk returns a tuple with the FinishedBatchesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedBatchesCount

`func (o *CleanableHistoricBatchReportResultDto) SetFinishedBatchesCount(v int64)`

SetFinishedBatchesCount sets FinishedBatchesCount field to given value.

### HasFinishedBatchesCount

`func (o *CleanableHistoricBatchReportResultDto) HasFinishedBatchesCount() bool`

HasFinishedBatchesCount returns a boolean if a field has been set.

### SetFinishedBatchesCountNil

`func (o *CleanableHistoricBatchReportResultDto) SetFinishedBatchesCountNil(b bool)`

 SetFinishedBatchesCountNil sets the value for FinishedBatchesCount to be an explicit nil

### UnsetFinishedBatchesCount
`func (o *CleanableHistoricBatchReportResultDto) UnsetFinishedBatchesCount()`

UnsetFinishedBatchesCount ensures that no value is present for FinishedBatchesCount, not even an explicit nil
### GetCleanableBatchesCount

`func (o *CleanableHistoricBatchReportResultDto) GetCleanableBatchesCount() int64`

GetCleanableBatchesCount returns the CleanableBatchesCount field if non-nil, zero value otherwise.

### GetCleanableBatchesCountOk

`func (o *CleanableHistoricBatchReportResultDto) GetCleanableBatchesCountOk() (*int64, bool)`

GetCleanableBatchesCountOk returns a tuple with the CleanableBatchesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanableBatchesCount

`func (o *CleanableHistoricBatchReportResultDto) SetCleanableBatchesCount(v int64)`

SetCleanableBatchesCount sets CleanableBatchesCount field to given value.

### HasCleanableBatchesCount

`func (o *CleanableHistoricBatchReportResultDto) HasCleanableBatchesCount() bool`

HasCleanableBatchesCount returns a boolean if a field has been set.

### SetCleanableBatchesCountNil

`func (o *CleanableHistoricBatchReportResultDto) SetCleanableBatchesCountNil(b bool)`

 SetCleanableBatchesCountNil sets the value for CleanableBatchesCount to be an explicit nil

### UnsetCleanableBatchesCount
`func (o *CleanableHistoricBatchReportResultDto) UnsetCleanableBatchesCount()`

UnsetCleanableBatchesCount ensures that no value is present for CleanableBatchesCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


