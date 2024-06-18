# HistoricBatchDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the batch. | [optional] 
**Type** | Pointer to **NullableString** | The type of the batch. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. | [optional] 
**TotalJobs** | Pointer to **NullableInt32** |  The total jobs of a batch is the number of batch execution jobs required to complete the batch.  | [optional] 
**BatchJobsPerSeed** | Pointer to **NullableInt32** |  The number of batch execution jobs created per seed job invocation. The batch seed job is invoked until it has created all batch execution jobs required by the batch (see &#x60;totalJobs&#x60; property).  | [optional] 
**InvocationsPerBatchJob** | Pointer to **NullableInt32** |  Every batch execution job invokes the command executed by the batch &#x60;invocationsPerBatchJob&#x60; times. E.g., for a process instance migration batch this specifies the number of process instances which are migrated per batch execution job.  | [optional] 
**SeedJobDefinitionId** | Pointer to **NullableString** | The job definition id for the seed jobs of this batch. | [optional] 
**MonitorJobDefinitionId** | Pointer to **NullableString** | The job definition id for the monitor jobs of this batch. | [optional] 
**BatchJobDefinitionId** | Pointer to **NullableString** | The job definition id for the batch execution jobs of this batch. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the batch. | [optional] 
**CreateUserId** | Pointer to **NullableString** | The batch creator&#39;s user id. | [optional] 
**StartTime** | Pointer to **NullableTime** | The time the batch was started. Default format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. For further information, please see the [documentation](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) | [optional] 
**ExecutionStartTime** | Pointer to **NullableTime** | The time the batch execution was started, i.e., at least one batch job has been executed. Default format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. For further information, please see the [documentation] (https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) | [optional] 
**EndTime** | Pointer to **NullableTime** | The time the batch ended. Default format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. For further information, please see the [documentation](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the historic batch should be removed by the History Cleanup job. Default format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. For further information, please see the [documentation](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) | [optional] 

## Methods

### NewHistoricBatchDto

`func NewHistoricBatchDto() *HistoricBatchDto`

NewHistoricBatchDto instantiates a new HistoricBatchDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricBatchDtoWithDefaults

`func NewHistoricBatchDtoWithDefaults() *HistoricBatchDto`

NewHistoricBatchDtoWithDefaults instantiates a new HistoricBatchDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricBatchDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricBatchDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricBatchDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricBatchDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricBatchDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricBatchDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *HistoricBatchDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoricBatchDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoricBatchDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HistoricBatchDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HistoricBatchDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HistoricBatchDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTotalJobs

`func (o *HistoricBatchDto) GetTotalJobs() int32`

GetTotalJobs returns the TotalJobs field if non-nil, zero value otherwise.

### GetTotalJobsOk

`func (o *HistoricBatchDto) GetTotalJobsOk() (*int32, bool)`

GetTotalJobsOk returns a tuple with the TotalJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalJobs

`func (o *HistoricBatchDto) SetTotalJobs(v int32)`

SetTotalJobs sets TotalJobs field to given value.

### HasTotalJobs

`func (o *HistoricBatchDto) HasTotalJobs() bool`

HasTotalJobs returns a boolean if a field has been set.

### SetTotalJobsNil

`func (o *HistoricBatchDto) SetTotalJobsNil(b bool)`

 SetTotalJobsNil sets the value for TotalJobs to be an explicit nil

### UnsetTotalJobs
`func (o *HistoricBatchDto) UnsetTotalJobs()`

UnsetTotalJobs ensures that no value is present for TotalJobs, not even an explicit nil
### GetBatchJobsPerSeed

`func (o *HistoricBatchDto) GetBatchJobsPerSeed() int32`

GetBatchJobsPerSeed returns the BatchJobsPerSeed field if non-nil, zero value otherwise.

### GetBatchJobsPerSeedOk

`func (o *HistoricBatchDto) GetBatchJobsPerSeedOk() (*int32, bool)`

GetBatchJobsPerSeedOk returns a tuple with the BatchJobsPerSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchJobsPerSeed

`func (o *HistoricBatchDto) SetBatchJobsPerSeed(v int32)`

SetBatchJobsPerSeed sets BatchJobsPerSeed field to given value.

### HasBatchJobsPerSeed

`func (o *HistoricBatchDto) HasBatchJobsPerSeed() bool`

HasBatchJobsPerSeed returns a boolean if a field has been set.

### SetBatchJobsPerSeedNil

`func (o *HistoricBatchDto) SetBatchJobsPerSeedNil(b bool)`

 SetBatchJobsPerSeedNil sets the value for BatchJobsPerSeed to be an explicit nil

### UnsetBatchJobsPerSeed
`func (o *HistoricBatchDto) UnsetBatchJobsPerSeed()`

UnsetBatchJobsPerSeed ensures that no value is present for BatchJobsPerSeed, not even an explicit nil
### GetInvocationsPerBatchJob

`func (o *HistoricBatchDto) GetInvocationsPerBatchJob() int32`

GetInvocationsPerBatchJob returns the InvocationsPerBatchJob field if non-nil, zero value otherwise.

### GetInvocationsPerBatchJobOk

`func (o *HistoricBatchDto) GetInvocationsPerBatchJobOk() (*int32, bool)`

GetInvocationsPerBatchJobOk returns a tuple with the InvocationsPerBatchJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationsPerBatchJob

`func (o *HistoricBatchDto) SetInvocationsPerBatchJob(v int32)`

SetInvocationsPerBatchJob sets InvocationsPerBatchJob field to given value.

### HasInvocationsPerBatchJob

`func (o *HistoricBatchDto) HasInvocationsPerBatchJob() bool`

HasInvocationsPerBatchJob returns a boolean if a field has been set.

### SetInvocationsPerBatchJobNil

`func (o *HistoricBatchDto) SetInvocationsPerBatchJobNil(b bool)`

 SetInvocationsPerBatchJobNil sets the value for InvocationsPerBatchJob to be an explicit nil

### UnsetInvocationsPerBatchJob
`func (o *HistoricBatchDto) UnsetInvocationsPerBatchJob()`

UnsetInvocationsPerBatchJob ensures that no value is present for InvocationsPerBatchJob, not even an explicit nil
### GetSeedJobDefinitionId

`func (o *HistoricBatchDto) GetSeedJobDefinitionId() string`

GetSeedJobDefinitionId returns the SeedJobDefinitionId field if non-nil, zero value otherwise.

### GetSeedJobDefinitionIdOk

`func (o *HistoricBatchDto) GetSeedJobDefinitionIdOk() (*string, bool)`

GetSeedJobDefinitionIdOk returns a tuple with the SeedJobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedJobDefinitionId

`func (o *HistoricBatchDto) SetSeedJobDefinitionId(v string)`

SetSeedJobDefinitionId sets SeedJobDefinitionId field to given value.

### HasSeedJobDefinitionId

`func (o *HistoricBatchDto) HasSeedJobDefinitionId() bool`

HasSeedJobDefinitionId returns a boolean if a field has been set.

### SetSeedJobDefinitionIdNil

`func (o *HistoricBatchDto) SetSeedJobDefinitionIdNil(b bool)`

 SetSeedJobDefinitionIdNil sets the value for SeedJobDefinitionId to be an explicit nil

### UnsetSeedJobDefinitionId
`func (o *HistoricBatchDto) UnsetSeedJobDefinitionId()`

UnsetSeedJobDefinitionId ensures that no value is present for SeedJobDefinitionId, not even an explicit nil
### GetMonitorJobDefinitionId

`func (o *HistoricBatchDto) GetMonitorJobDefinitionId() string`

GetMonitorJobDefinitionId returns the MonitorJobDefinitionId field if non-nil, zero value otherwise.

### GetMonitorJobDefinitionIdOk

`func (o *HistoricBatchDto) GetMonitorJobDefinitionIdOk() (*string, bool)`

GetMonitorJobDefinitionIdOk returns a tuple with the MonitorJobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorJobDefinitionId

`func (o *HistoricBatchDto) SetMonitorJobDefinitionId(v string)`

SetMonitorJobDefinitionId sets MonitorJobDefinitionId field to given value.

### HasMonitorJobDefinitionId

`func (o *HistoricBatchDto) HasMonitorJobDefinitionId() bool`

HasMonitorJobDefinitionId returns a boolean if a field has been set.

### SetMonitorJobDefinitionIdNil

`func (o *HistoricBatchDto) SetMonitorJobDefinitionIdNil(b bool)`

 SetMonitorJobDefinitionIdNil sets the value for MonitorJobDefinitionId to be an explicit nil

### UnsetMonitorJobDefinitionId
`func (o *HistoricBatchDto) UnsetMonitorJobDefinitionId()`

UnsetMonitorJobDefinitionId ensures that no value is present for MonitorJobDefinitionId, not even an explicit nil
### GetBatchJobDefinitionId

`func (o *HistoricBatchDto) GetBatchJobDefinitionId() string`

GetBatchJobDefinitionId returns the BatchJobDefinitionId field if non-nil, zero value otherwise.

### GetBatchJobDefinitionIdOk

`func (o *HistoricBatchDto) GetBatchJobDefinitionIdOk() (*string, bool)`

GetBatchJobDefinitionIdOk returns a tuple with the BatchJobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchJobDefinitionId

`func (o *HistoricBatchDto) SetBatchJobDefinitionId(v string)`

SetBatchJobDefinitionId sets BatchJobDefinitionId field to given value.

### HasBatchJobDefinitionId

`func (o *HistoricBatchDto) HasBatchJobDefinitionId() bool`

HasBatchJobDefinitionId returns a boolean if a field has been set.

### SetBatchJobDefinitionIdNil

`func (o *HistoricBatchDto) SetBatchJobDefinitionIdNil(b bool)`

 SetBatchJobDefinitionIdNil sets the value for BatchJobDefinitionId to be an explicit nil

### UnsetBatchJobDefinitionId
`func (o *HistoricBatchDto) UnsetBatchJobDefinitionId()`

UnsetBatchJobDefinitionId ensures that no value is present for BatchJobDefinitionId, not even an explicit nil
### GetTenantId

`func (o *HistoricBatchDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricBatchDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricBatchDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricBatchDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricBatchDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricBatchDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCreateUserId

`func (o *HistoricBatchDto) GetCreateUserId() string`

GetCreateUserId returns the CreateUserId field if non-nil, zero value otherwise.

### GetCreateUserIdOk

`func (o *HistoricBatchDto) GetCreateUserIdOk() (*string, bool)`

GetCreateUserIdOk returns a tuple with the CreateUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUserId

`func (o *HistoricBatchDto) SetCreateUserId(v string)`

SetCreateUserId sets CreateUserId field to given value.

### HasCreateUserId

`func (o *HistoricBatchDto) HasCreateUserId() bool`

HasCreateUserId returns a boolean if a field has been set.

### SetCreateUserIdNil

`func (o *HistoricBatchDto) SetCreateUserIdNil(b bool)`

 SetCreateUserIdNil sets the value for CreateUserId to be an explicit nil

### UnsetCreateUserId
`func (o *HistoricBatchDto) UnsetCreateUserId()`

UnsetCreateUserId ensures that no value is present for CreateUserId, not even an explicit nil
### GetStartTime

`func (o *HistoricBatchDto) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HistoricBatchDto) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HistoricBatchDto) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HistoricBatchDto) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *HistoricBatchDto) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *HistoricBatchDto) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetExecutionStartTime

`func (o *HistoricBatchDto) GetExecutionStartTime() time.Time`

GetExecutionStartTime returns the ExecutionStartTime field if non-nil, zero value otherwise.

### GetExecutionStartTimeOk

`func (o *HistoricBatchDto) GetExecutionStartTimeOk() (*time.Time, bool)`

GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStartTime

`func (o *HistoricBatchDto) SetExecutionStartTime(v time.Time)`

SetExecutionStartTime sets ExecutionStartTime field to given value.

### HasExecutionStartTime

`func (o *HistoricBatchDto) HasExecutionStartTime() bool`

HasExecutionStartTime returns a boolean if a field has been set.

### SetExecutionStartTimeNil

`func (o *HistoricBatchDto) SetExecutionStartTimeNil(b bool)`

 SetExecutionStartTimeNil sets the value for ExecutionStartTime to be an explicit nil

### UnsetExecutionStartTime
`func (o *HistoricBatchDto) UnsetExecutionStartTime()`

UnsetExecutionStartTime ensures that no value is present for ExecutionStartTime, not even an explicit nil
### GetEndTime

`func (o *HistoricBatchDto) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HistoricBatchDto) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HistoricBatchDto) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *HistoricBatchDto) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *HistoricBatchDto) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *HistoricBatchDto) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetRemovalTime

`func (o *HistoricBatchDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricBatchDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricBatchDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricBatchDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricBatchDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricBatchDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


