# CleanableHistoricProcessInstanceReportResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition. | [optional] 
**ProcessDefinitionName** | Pointer to **NullableString** | The name of the process definition. | [optional] 
**ProcessDefinitionVersion** | Pointer to **NullableInt32** | The version of the process definition. | [optional] 
**HistoryTimeToLive** | Pointer to **NullableInt32** | The history time to live of the process definition. | [optional] 
**FinishedProcessInstanceCount** | Pointer to **NullableInt64** | The count of the finished historic process instances. | [optional] 
**CleanableProcessInstanceCount** | Pointer to **NullableInt64** | The count of the cleanable historic process instances, referring to history time to live. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the process definition. | [optional] 

## Methods

### NewCleanableHistoricProcessInstanceReportResultDto

`func NewCleanableHistoricProcessInstanceReportResultDto() *CleanableHistoricProcessInstanceReportResultDto`

NewCleanableHistoricProcessInstanceReportResultDto instantiates a new CleanableHistoricProcessInstanceReportResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanableHistoricProcessInstanceReportResultDtoWithDefaults

`func NewCleanableHistoricProcessInstanceReportResultDtoWithDefaults() *CleanableHistoricProcessInstanceReportResultDto`

NewCleanableHistoricProcessInstanceReportResultDtoWithDefaults instantiates a new CleanableHistoricProcessInstanceReportResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessDefinitionId

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *CleanableHistoricProcessInstanceReportResultDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *CleanableHistoricProcessInstanceReportResultDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *CleanableHistoricProcessInstanceReportResultDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *CleanableHistoricProcessInstanceReportResultDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionName

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetProcessDefinitionName() string`

GetProcessDefinitionName returns the ProcessDefinitionName field if non-nil, zero value otherwise.

### GetProcessDefinitionNameOk

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetProcessDefinitionNameOk() (*string, bool)`

GetProcessDefinitionNameOk returns a tuple with the ProcessDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionName

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetProcessDefinitionName(v string)`

SetProcessDefinitionName sets ProcessDefinitionName field to given value.

### HasProcessDefinitionName

`func (o *CleanableHistoricProcessInstanceReportResultDto) HasProcessDefinitionName() bool`

HasProcessDefinitionName returns a boolean if a field has been set.

### SetProcessDefinitionNameNil

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetProcessDefinitionNameNil(b bool)`

 SetProcessDefinitionNameNil sets the value for ProcessDefinitionName to be an explicit nil

### UnsetProcessDefinitionName
`func (o *CleanableHistoricProcessInstanceReportResultDto) UnsetProcessDefinitionName()`

UnsetProcessDefinitionName ensures that no value is present for ProcessDefinitionName, not even an explicit nil
### GetProcessDefinitionVersion

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetProcessDefinitionVersion() int32`

GetProcessDefinitionVersion returns the ProcessDefinitionVersion field if non-nil, zero value otherwise.

### GetProcessDefinitionVersionOk

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetProcessDefinitionVersionOk() (*int32, bool)`

GetProcessDefinitionVersionOk returns a tuple with the ProcessDefinitionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionVersion

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetProcessDefinitionVersion(v int32)`

SetProcessDefinitionVersion sets ProcessDefinitionVersion field to given value.

### HasProcessDefinitionVersion

`func (o *CleanableHistoricProcessInstanceReportResultDto) HasProcessDefinitionVersion() bool`

HasProcessDefinitionVersion returns a boolean if a field has been set.

### SetProcessDefinitionVersionNil

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetProcessDefinitionVersionNil(b bool)`

 SetProcessDefinitionVersionNil sets the value for ProcessDefinitionVersion to be an explicit nil

### UnsetProcessDefinitionVersion
`func (o *CleanableHistoricProcessInstanceReportResultDto) UnsetProcessDefinitionVersion()`

UnsetProcessDefinitionVersion ensures that no value is present for ProcessDefinitionVersion, not even an explicit nil
### GetHistoryTimeToLive

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetHistoryTimeToLive() int32`

GetHistoryTimeToLive returns the HistoryTimeToLive field if non-nil, zero value otherwise.

### GetHistoryTimeToLiveOk

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetHistoryTimeToLiveOk() (*int32, bool)`

GetHistoryTimeToLiveOk returns a tuple with the HistoryTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryTimeToLive

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetHistoryTimeToLive(v int32)`

SetHistoryTimeToLive sets HistoryTimeToLive field to given value.

### HasHistoryTimeToLive

`func (o *CleanableHistoricProcessInstanceReportResultDto) HasHistoryTimeToLive() bool`

HasHistoryTimeToLive returns a boolean if a field has been set.

### SetHistoryTimeToLiveNil

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetHistoryTimeToLiveNil(b bool)`

 SetHistoryTimeToLiveNil sets the value for HistoryTimeToLive to be an explicit nil

### UnsetHistoryTimeToLive
`func (o *CleanableHistoricProcessInstanceReportResultDto) UnsetHistoryTimeToLive()`

UnsetHistoryTimeToLive ensures that no value is present for HistoryTimeToLive, not even an explicit nil
### GetFinishedProcessInstanceCount

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetFinishedProcessInstanceCount() int64`

GetFinishedProcessInstanceCount returns the FinishedProcessInstanceCount field if non-nil, zero value otherwise.

### GetFinishedProcessInstanceCountOk

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetFinishedProcessInstanceCountOk() (*int64, bool)`

GetFinishedProcessInstanceCountOk returns a tuple with the FinishedProcessInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedProcessInstanceCount

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetFinishedProcessInstanceCount(v int64)`

SetFinishedProcessInstanceCount sets FinishedProcessInstanceCount field to given value.

### HasFinishedProcessInstanceCount

`func (o *CleanableHistoricProcessInstanceReportResultDto) HasFinishedProcessInstanceCount() bool`

HasFinishedProcessInstanceCount returns a boolean if a field has been set.

### SetFinishedProcessInstanceCountNil

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetFinishedProcessInstanceCountNil(b bool)`

 SetFinishedProcessInstanceCountNil sets the value for FinishedProcessInstanceCount to be an explicit nil

### UnsetFinishedProcessInstanceCount
`func (o *CleanableHistoricProcessInstanceReportResultDto) UnsetFinishedProcessInstanceCount()`

UnsetFinishedProcessInstanceCount ensures that no value is present for FinishedProcessInstanceCount, not even an explicit nil
### GetCleanableProcessInstanceCount

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetCleanableProcessInstanceCount() int64`

GetCleanableProcessInstanceCount returns the CleanableProcessInstanceCount field if non-nil, zero value otherwise.

### GetCleanableProcessInstanceCountOk

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetCleanableProcessInstanceCountOk() (*int64, bool)`

GetCleanableProcessInstanceCountOk returns a tuple with the CleanableProcessInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanableProcessInstanceCount

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetCleanableProcessInstanceCount(v int64)`

SetCleanableProcessInstanceCount sets CleanableProcessInstanceCount field to given value.

### HasCleanableProcessInstanceCount

`func (o *CleanableHistoricProcessInstanceReportResultDto) HasCleanableProcessInstanceCount() bool`

HasCleanableProcessInstanceCount returns a boolean if a field has been set.

### SetCleanableProcessInstanceCountNil

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetCleanableProcessInstanceCountNil(b bool)`

 SetCleanableProcessInstanceCountNil sets the value for CleanableProcessInstanceCount to be an explicit nil

### UnsetCleanableProcessInstanceCount
`func (o *CleanableHistoricProcessInstanceReportResultDto) UnsetCleanableProcessInstanceCount()`

UnsetCleanableProcessInstanceCount ensures that no value is present for CleanableProcessInstanceCount, not even an explicit nil
### GetTenantId

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CleanableHistoricProcessInstanceReportResultDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CleanableHistoricProcessInstanceReportResultDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CleanableHistoricProcessInstanceReportResultDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CleanableHistoricProcessInstanceReportResultDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


