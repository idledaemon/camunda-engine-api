# BatchDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the batch. | [optional] 
**Type** | Pointer to **NullableString** | The type of the batch. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. | [optional] 
**TotalJobs** | Pointer to **NullableInt32** | The total jobs of a batch is the number of batch execution jobs required to complete the batch. | [optional] 
**JobsCreated** | Pointer to **NullableInt32** | The number of batch execution jobs already created by the seed job. | [optional] 
**BatchJobsPerSeed** | Pointer to **NullableInt32** | The number of batch execution jobs created per seed job invocation. The batch seed job is invoked until it has created all batch execution jobs required by the batch (see &#x60;totalJobs&#x60; property). | [optional] 
**InvocationsPerBatchJob** | Pointer to **NullableInt32** | Every batch execution job invokes the command executed by the batch &#x60;invocationsPerBatchJob&#x60; times. E.g., for a process instance migration batch this specifies the number of process instances which are migrated per batch execution job. | [optional] 
**SeedJobDefinitionId** | Pointer to **NullableString** | The job definition id for the seed jobs of this batch. | [optional] 
**MonitorJobDefinitionId** | Pointer to **NullableString** | The job definition id for the monitor jobs of this batch. | [optional] 
**BatchJobDefinitionId** | Pointer to **NullableString** | The job definition id for the batch execution jobs of this batch. | [optional] 
**Suspended** | Pointer to **NullableBool** | Indicates whether this batch is suspended or not. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the batch. | [optional] 
**CreateUserId** | Pointer to **NullableString** | The id of the user that created the batch. | [optional] 
**StartTime** | Pointer to **NullableTime** | The time the batch was started. Default format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. For further information, please see the [documentation] (https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) | [optional] 
**ExecutionStartTime** | Pointer to **NullableTime** | The time the batch execution was started, i.e., at least one batch job has been executed. Default format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. For further information, please see the [documentation] (https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) | [optional] 

## Methods

### NewBatchDto

`func NewBatchDto() *BatchDto`

NewBatchDto instantiates a new BatchDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchDtoWithDefaults

`func NewBatchDtoWithDefaults() *BatchDto`

NewBatchDtoWithDefaults instantiates a new BatchDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BatchDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BatchDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BatchDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *BatchDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BatchDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BatchDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BatchDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BatchDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BatchDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTotalJobs

`func (o *BatchDto) GetTotalJobs() int32`

GetTotalJobs returns the TotalJobs field if non-nil, zero value otherwise.

### GetTotalJobsOk

`func (o *BatchDto) GetTotalJobsOk() (*int32, bool)`

GetTotalJobsOk returns a tuple with the TotalJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalJobs

`func (o *BatchDto) SetTotalJobs(v int32)`

SetTotalJobs sets TotalJobs field to given value.

### HasTotalJobs

`func (o *BatchDto) HasTotalJobs() bool`

HasTotalJobs returns a boolean if a field has been set.

### SetTotalJobsNil

`func (o *BatchDto) SetTotalJobsNil(b bool)`

 SetTotalJobsNil sets the value for TotalJobs to be an explicit nil

### UnsetTotalJobs
`func (o *BatchDto) UnsetTotalJobs()`

UnsetTotalJobs ensures that no value is present for TotalJobs, not even an explicit nil
### GetJobsCreated

`func (o *BatchDto) GetJobsCreated() int32`

GetJobsCreated returns the JobsCreated field if non-nil, zero value otherwise.

### GetJobsCreatedOk

`func (o *BatchDto) GetJobsCreatedOk() (*int32, bool)`

GetJobsCreatedOk returns a tuple with the JobsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCreated

`func (o *BatchDto) SetJobsCreated(v int32)`

SetJobsCreated sets JobsCreated field to given value.

### HasJobsCreated

`func (o *BatchDto) HasJobsCreated() bool`

HasJobsCreated returns a boolean if a field has been set.

### SetJobsCreatedNil

`func (o *BatchDto) SetJobsCreatedNil(b bool)`

 SetJobsCreatedNil sets the value for JobsCreated to be an explicit nil

### UnsetJobsCreated
`func (o *BatchDto) UnsetJobsCreated()`

UnsetJobsCreated ensures that no value is present for JobsCreated, not even an explicit nil
### GetBatchJobsPerSeed

`func (o *BatchDto) GetBatchJobsPerSeed() int32`

GetBatchJobsPerSeed returns the BatchJobsPerSeed field if non-nil, zero value otherwise.

### GetBatchJobsPerSeedOk

`func (o *BatchDto) GetBatchJobsPerSeedOk() (*int32, bool)`

GetBatchJobsPerSeedOk returns a tuple with the BatchJobsPerSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchJobsPerSeed

`func (o *BatchDto) SetBatchJobsPerSeed(v int32)`

SetBatchJobsPerSeed sets BatchJobsPerSeed field to given value.

### HasBatchJobsPerSeed

`func (o *BatchDto) HasBatchJobsPerSeed() bool`

HasBatchJobsPerSeed returns a boolean if a field has been set.

### SetBatchJobsPerSeedNil

`func (o *BatchDto) SetBatchJobsPerSeedNil(b bool)`

 SetBatchJobsPerSeedNil sets the value for BatchJobsPerSeed to be an explicit nil

### UnsetBatchJobsPerSeed
`func (o *BatchDto) UnsetBatchJobsPerSeed()`

UnsetBatchJobsPerSeed ensures that no value is present for BatchJobsPerSeed, not even an explicit nil
### GetInvocationsPerBatchJob

`func (o *BatchDto) GetInvocationsPerBatchJob() int32`

GetInvocationsPerBatchJob returns the InvocationsPerBatchJob field if non-nil, zero value otherwise.

### GetInvocationsPerBatchJobOk

`func (o *BatchDto) GetInvocationsPerBatchJobOk() (*int32, bool)`

GetInvocationsPerBatchJobOk returns a tuple with the InvocationsPerBatchJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationsPerBatchJob

`func (o *BatchDto) SetInvocationsPerBatchJob(v int32)`

SetInvocationsPerBatchJob sets InvocationsPerBatchJob field to given value.

### HasInvocationsPerBatchJob

`func (o *BatchDto) HasInvocationsPerBatchJob() bool`

HasInvocationsPerBatchJob returns a boolean if a field has been set.

### SetInvocationsPerBatchJobNil

`func (o *BatchDto) SetInvocationsPerBatchJobNil(b bool)`

 SetInvocationsPerBatchJobNil sets the value for InvocationsPerBatchJob to be an explicit nil

### UnsetInvocationsPerBatchJob
`func (o *BatchDto) UnsetInvocationsPerBatchJob()`

UnsetInvocationsPerBatchJob ensures that no value is present for InvocationsPerBatchJob, not even an explicit nil
### GetSeedJobDefinitionId

`func (o *BatchDto) GetSeedJobDefinitionId() string`

GetSeedJobDefinitionId returns the SeedJobDefinitionId field if non-nil, zero value otherwise.

### GetSeedJobDefinitionIdOk

`func (o *BatchDto) GetSeedJobDefinitionIdOk() (*string, bool)`

GetSeedJobDefinitionIdOk returns a tuple with the SeedJobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedJobDefinitionId

`func (o *BatchDto) SetSeedJobDefinitionId(v string)`

SetSeedJobDefinitionId sets SeedJobDefinitionId field to given value.

### HasSeedJobDefinitionId

`func (o *BatchDto) HasSeedJobDefinitionId() bool`

HasSeedJobDefinitionId returns a boolean if a field has been set.

### SetSeedJobDefinitionIdNil

`func (o *BatchDto) SetSeedJobDefinitionIdNil(b bool)`

 SetSeedJobDefinitionIdNil sets the value for SeedJobDefinitionId to be an explicit nil

### UnsetSeedJobDefinitionId
`func (o *BatchDto) UnsetSeedJobDefinitionId()`

UnsetSeedJobDefinitionId ensures that no value is present for SeedJobDefinitionId, not even an explicit nil
### GetMonitorJobDefinitionId

`func (o *BatchDto) GetMonitorJobDefinitionId() string`

GetMonitorJobDefinitionId returns the MonitorJobDefinitionId field if non-nil, zero value otherwise.

### GetMonitorJobDefinitionIdOk

`func (o *BatchDto) GetMonitorJobDefinitionIdOk() (*string, bool)`

GetMonitorJobDefinitionIdOk returns a tuple with the MonitorJobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorJobDefinitionId

`func (o *BatchDto) SetMonitorJobDefinitionId(v string)`

SetMonitorJobDefinitionId sets MonitorJobDefinitionId field to given value.

### HasMonitorJobDefinitionId

`func (o *BatchDto) HasMonitorJobDefinitionId() bool`

HasMonitorJobDefinitionId returns a boolean if a field has been set.

### SetMonitorJobDefinitionIdNil

`func (o *BatchDto) SetMonitorJobDefinitionIdNil(b bool)`

 SetMonitorJobDefinitionIdNil sets the value for MonitorJobDefinitionId to be an explicit nil

### UnsetMonitorJobDefinitionId
`func (o *BatchDto) UnsetMonitorJobDefinitionId()`

UnsetMonitorJobDefinitionId ensures that no value is present for MonitorJobDefinitionId, not even an explicit nil
### GetBatchJobDefinitionId

`func (o *BatchDto) GetBatchJobDefinitionId() string`

GetBatchJobDefinitionId returns the BatchJobDefinitionId field if non-nil, zero value otherwise.

### GetBatchJobDefinitionIdOk

`func (o *BatchDto) GetBatchJobDefinitionIdOk() (*string, bool)`

GetBatchJobDefinitionIdOk returns a tuple with the BatchJobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchJobDefinitionId

`func (o *BatchDto) SetBatchJobDefinitionId(v string)`

SetBatchJobDefinitionId sets BatchJobDefinitionId field to given value.

### HasBatchJobDefinitionId

`func (o *BatchDto) HasBatchJobDefinitionId() bool`

HasBatchJobDefinitionId returns a boolean if a field has been set.

### SetBatchJobDefinitionIdNil

`func (o *BatchDto) SetBatchJobDefinitionIdNil(b bool)`

 SetBatchJobDefinitionIdNil sets the value for BatchJobDefinitionId to be an explicit nil

### UnsetBatchJobDefinitionId
`func (o *BatchDto) UnsetBatchJobDefinitionId()`

UnsetBatchJobDefinitionId ensures that no value is present for BatchJobDefinitionId, not even an explicit nil
### GetSuspended

`func (o *BatchDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *BatchDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *BatchDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *BatchDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *BatchDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *BatchDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetTenantId

`func (o *BatchDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BatchDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BatchDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BatchDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BatchDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BatchDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCreateUserId

`func (o *BatchDto) GetCreateUserId() string`

GetCreateUserId returns the CreateUserId field if non-nil, zero value otherwise.

### GetCreateUserIdOk

`func (o *BatchDto) GetCreateUserIdOk() (*string, bool)`

GetCreateUserIdOk returns a tuple with the CreateUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUserId

`func (o *BatchDto) SetCreateUserId(v string)`

SetCreateUserId sets CreateUserId field to given value.

### HasCreateUserId

`func (o *BatchDto) HasCreateUserId() bool`

HasCreateUserId returns a boolean if a field has been set.

### SetCreateUserIdNil

`func (o *BatchDto) SetCreateUserIdNil(b bool)`

 SetCreateUserIdNil sets the value for CreateUserId to be an explicit nil

### UnsetCreateUserId
`func (o *BatchDto) UnsetCreateUserId()`

UnsetCreateUserId ensures that no value is present for CreateUserId, not even an explicit nil
### GetStartTime

`func (o *BatchDto) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BatchDto) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BatchDto) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BatchDto) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *BatchDto) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *BatchDto) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetExecutionStartTime

`func (o *BatchDto) GetExecutionStartTime() time.Time`

GetExecutionStartTime returns the ExecutionStartTime field if non-nil, zero value otherwise.

### GetExecutionStartTimeOk

`func (o *BatchDto) GetExecutionStartTimeOk() (*time.Time, bool)`

GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStartTime

`func (o *BatchDto) SetExecutionStartTime(v time.Time)`

SetExecutionStartTime sets ExecutionStartTime field to given value.

### HasExecutionStartTime

`func (o *BatchDto) HasExecutionStartTime() bool`

HasExecutionStartTime returns a boolean if a field has been set.

### SetExecutionStartTimeNil

`func (o *BatchDto) SetExecutionStartTimeNil(b bool)`

 SetExecutionStartTimeNil sets the value for ExecutionStartTime to be an explicit nil

### UnsetExecutionStartTime
`func (o *BatchDto) UnsetExecutionStartTime()`

UnsetExecutionStartTime ensures that no value is present for ExecutionStartTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


