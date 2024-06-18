# CleanableHistoricDecisionInstanceReportResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecisionDefinitionId** | Pointer to **NullableString** | The id of the decision definition. | [optional] 
**DecisionDefinitionKey** | Pointer to **NullableString** | The key of the decision definition. | [optional] 
**DecisionDefinitionName** | Pointer to **NullableString** | The name of the decision definition. | [optional] 
**DecisionDefinitionVersion** | Pointer to **NullableInt32** | The version of the decision definition. | [optional] 
**HistoryTimeToLive** | Pointer to **NullableInt32** | The history time to live of the decision definition. | [optional] 
**FinishedDecisionInstanceCount** | Pointer to **NullableInt64** | The count of the finished historic decision instances. | [optional] 
**CleanableDecisionInstanceCount** | Pointer to **NullableInt64** | The count of the cleanable historic decision instances, referring to history time to live. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the decision definition. | [optional] 

## Methods

### NewCleanableHistoricDecisionInstanceReportResultDto

`func NewCleanableHistoricDecisionInstanceReportResultDto() *CleanableHistoricDecisionInstanceReportResultDto`

NewCleanableHistoricDecisionInstanceReportResultDto instantiates a new CleanableHistoricDecisionInstanceReportResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanableHistoricDecisionInstanceReportResultDtoWithDefaults

`func NewCleanableHistoricDecisionInstanceReportResultDtoWithDefaults() *CleanableHistoricDecisionInstanceReportResultDto`

NewCleanableHistoricDecisionInstanceReportResultDtoWithDefaults instantiates a new CleanableHistoricDecisionInstanceReportResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecisionDefinitionId

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionId() string`

GetDecisionDefinitionId returns the DecisionDefinitionId field if non-nil, zero value otherwise.

### GetDecisionDefinitionIdOk

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionIdOk() (*string, bool)`

GetDecisionDefinitionIdOk returns a tuple with the DecisionDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionId

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionId(v string)`

SetDecisionDefinitionId sets DecisionDefinitionId field to given value.

### HasDecisionDefinitionId

`func (o *CleanableHistoricDecisionInstanceReportResultDto) HasDecisionDefinitionId() bool`

HasDecisionDefinitionId returns a boolean if a field has been set.

### SetDecisionDefinitionIdNil

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionIdNil(b bool)`

 SetDecisionDefinitionIdNil sets the value for DecisionDefinitionId to be an explicit nil

### UnsetDecisionDefinitionId
`func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetDecisionDefinitionId()`

UnsetDecisionDefinitionId ensures that no value is present for DecisionDefinitionId, not even an explicit nil
### GetDecisionDefinitionKey

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionKey() string`

GetDecisionDefinitionKey returns the DecisionDefinitionKey field if non-nil, zero value otherwise.

### GetDecisionDefinitionKeyOk

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionKeyOk() (*string, bool)`

GetDecisionDefinitionKeyOk returns a tuple with the DecisionDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionKey

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionKey(v string)`

SetDecisionDefinitionKey sets DecisionDefinitionKey field to given value.

### HasDecisionDefinitionKey

`func (o *CleanableHistoricDecisionInstanceReportResultDto) HasDecisionDefinitionKey() bool`

HasDecisionDefinitionKey returns a boolean if a field has been set.

### SetDecisionDefinitionKeyNil

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionKeyNil(b bool)`

 SetDecisionDefinitionKeyNil sets the value for DecisionDefinitionKey to be an explicit nil

### UnsetDecisionDefinitionKey
`func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetDecisionDefinitionKey()`

UnsetDecisionDefinitionKey ensures that no value is present for DecisionDefinitionKey, not even an explicit nil
### GetDecisionDefinitionName

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionName() string`

GetDecisionDefinitionName returns the DecisionDefinitionName field if non-nil, zero value otherwise.

### GetDecisionDefinitionNameOk

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionNameOk() (*string, bool)`

GetDecisionDefinitionNameOk returns a tuple with the DecisionDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionName

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionName(v string)`

SetDecisionDefinitionName sets DecisionDefinitionName field to given value.

### HasDecisionDefinitionName

`func (o *CleanableHistoricDecisionInstanceReportResultDto) HasDecisionDefinitionName() bool`

HasDecisionDefinitionName returns a boolean if a field has been set.

### SetDecisionDefinitionNameNil

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionNameNil(b bool)`

 SetDecisionDefinitionNameNil sets the value for DecisionDefinitionName to be an explicit nil

### UnsetDecisionDefinitionName
`func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetDecisionDefinitionName()`

UnsetDecisionDefinitionName ensures that no value is present for DecisionDefinitionName, not even an explicit nil
### GetDecisionDefinitionVersion

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionVersion() int32`

GetDecisionDefinitionVersion returns the DecisionDefinitionVersion field if non-nil, zero value otherwise.

### GetDecisionDefinitionVersionOk

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionVersionOk() (*int32, bool)`

GetDecisionDefinitionVersionOk returns a tuple with the DecisionDefinitionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionVersion

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionVersion(v int32)`

SetDecisionDefinitionVersion sets DecisionDefinitionVersion field to given value.

### HasDecisionDefinitionVersion

`func (o *CleanableHistoricDecisionInstanceReportResultDto) HasDecisionDefinitionVersion() bool`

HasDecisionDefinitionVersion returns a boolean if a field has been set.

### SetDecisionDefinitionVersionNil

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionVersionNil(b bool)`

 SetDecisionDefinitionVersionNil sets the value for DecisionDefinitionVersion to be an explicit nil

### UnsetDecisionDefinitionVersion
`func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetDecisionDefinitionVersion()`

UnsetDecisionDefinitionVersion ensures that no value is present for DecisionDefinitionVersion, not even an explicit nil
### GetHistoryTimeToLive

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetHistoryTimeToLive() int32`

GetHistoryTimeToLive returns the HistoryTimeToLive field if non-nil, zero value otherwise.

### GetHistoryTimeToLiveOk

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetHistoryTimeToLiveOk() (*int32, bool)`

GetHistoryTimeToLiveOk returns a tuple with the HistoryTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryTimeToLive

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetHistoryTimeToLive(v int32)`

SetHistoryTimeToLive sets HistoryTimeToLive field to given value.

### HasHistoryTimeToLive

`func (o *CleanableHistoricDecisionInstanceReportResultDto) HasHistoryTimeToLive() bool`

HasHistoryTimeToLive returns a boolean if a field has been set.

### SetHistoryTimeToLiveNil

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetHistoryTimeToLiveNil(b bool)`

 SetHistoryTimeToLiveNil sets the value for HistoryTimeToLive to be an explicit nil

### UnsetHistoryTimeToLive
`func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetHistoryTimeToLive()`

UnsetHistoryTimeToLive ensures that no value is present for HistoryTimeToLive, not even an explicit nil
### GetFinishedDecisionInstanceCount

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetFinishedDecisionInstanceCount() int64`

GetFinishedDecisionInstanceCount returns the FinishedDecisionInstanceCount field if non-nil, zero value otherwise.

### GetFinishedDecisionInstanceCountOk

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetFinishedDecisionInstanceCountOk() (*int64, bool)`

GetFinishedDecisionInstanceCountOk returns a tuple with the FinishedDecisionInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedDecisionInstanceCount

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetFinishedDecisionInstanceCount(v int64)`

SetFinishedDecisionInstanceCount sets FinishedDecisionInstanceCount field to given value.

### HasFinishedDecisionInstanceCount

`func (o *CleanableHistoricDecisionInstanceReportResultDto) HasFinishedDecisionInstanceCount() bool`

HasFinishedDecisionInstanceCount returns a boolean if a field has been set.

### SetFinishedDecisionInstanceCountNil

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetFinishedDecisionInstanceCountNil(b bool)`

 SetFinishedDecisionInstanceCountNil sets the value for FinishedDecisionInstanceCount to be an explicit nil

### UnsetFinishedDecisionInstanceCount
`func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetFinishedDecisionInstanceCount()`

UnsetFinishedDecisionInstanceCount ensures that no value is present for FinishedDecisionInstanceCount, not even an explicit nil
### GetCleanableDecisionInstanceCount

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetCleanableDecisionInstanceCount() int64`

GetCleanableDecisionInstanceCount returns the CleanableDecisionInstanceCount field if non-nil, zero value otherwise.

### GetCleanableDecisionInstanceCountOk

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetCleanableDecisionInstanceCountOk() (*int64, bool)`

GetCleanableDecisionInstanceCountOk returns a tuple with the CleanableDecisionInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanableDecisionInstanceCount

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetCleanableDecisionInstanceCount(v int64)`

SetCleanableDecisionInstanceCount sets CleanableDecisionInstanceCount field to given value.

### HasCleanableDecisionInstanceCount

`func (o *CleanableHistoricDecisionInstanceReportResultDto) HasCleanableDecisionInstanceCount() bool`

HasCleanableDecisionInstanceCount returns a boolean if a field has been set.

### SetCleanableDecisionInstanceCountNil

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetCleanableDecisionInstanceCountNil(b bool)`

 SetCleanableDecisionInstanceCountNil sets the value for CleanableDecisionInstanceCount to be an explicit nil

### UnsetCleanableDecisionInstanceCount
`func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetCleanableDecisionInstanceCount()`

UnsetCleanableDecisionInstanceCount ensures that no value is present for CleanableDecisionInstanceCount, not even an explicit nil
### GetTenantId

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CleanableHistoricDecisionInstanceReportResultDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CleanableHistoricDecisionInstanceReportResultDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CleanableHistoricDecisionInstanceReportResultDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


