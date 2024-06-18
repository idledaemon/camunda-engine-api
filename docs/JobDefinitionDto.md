# JobDefinitionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the job definition. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition this job definition is associated with. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition this job definition is associated with. | [optional] 
**ActivityId** | Pointer to **NullableString** | The id of the activity this job definition is associated with. | [optional] 
**JobType** | Pointer to **NullableString** | The type of the job which is running for this job definition. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job types. | [optional] 
**JobConfiguration** | Pointer to **NullableString** | The configuration of a job definition provides details about the jobs which will be created. For example: for timer jobs it is the timer configuration. | [optional] 
**OverridingJobPriority** | Pointer to **NullableInt64** | The execution priority defined for jobs that are created based on this definition. May be &#x60;null&#x60; when the priority has not been overridden on the job definition level. | [optional] 
**Suspended** | Pointer to **NullableBool** | Indicates whether this job definition is suspended or not. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant this job definition is associated with. | [optional] 
**DeploymentId** | Pointer to **NullableString** | The id of the deployment this job definition is related to. In a deployment-aware setup, this leads to all jobs of the same definition being executed on the same node. | [optional] 

## Methods

### NewJobDefinitionDto

`func NewJobDefinitionDto() *JobDefinitionDto`

NewJobDefinitionDto instantiates a new JobDefinitionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDefinitionDtoWithDefaults

`func NewJobDefinitionDtoWithDefaults() *JobDefinitionDto`

NewJobDefinitionDtoWithDefaults instantiates a new JobDefinitionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobDefinitionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobDefinitionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobDefinitionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobDefinitionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JobDefinitionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JobDefinitionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProcessDefinitionId

`func (o *JobDefinitionDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *JobDefinitionDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *JobDefinitionDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *JobDefinitionDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *JobDefinitionDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *JobDefinitionDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *JobDefinitionDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *JobDefinitionDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *JobDefinitionDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *JobDefinitionDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *JobDefinitionDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *JobDefinitionDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetActivityId

`func (o *JobDefinitionDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *JobDefinitionDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *JobDefinitionDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *JobDefinitionDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *JobDefinitionDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *JobDefinitionDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetJobType

`func (o *JobDefinitionDto) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobDefinitionDto) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobDefinitionDto) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *JobDefinitionDto) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *JobDefinitionDto) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *JobDefinitionDto) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetJobConfiguration

`func (o *JobDefinitionDto) GetJobConfiguration() string`

GetJobConfiguration returns the JobConfiguration field if non-nil, zero value otherwise.

### GetJobConfigurationOk

`func (o *JobDefinitionDto) GetJobConfigurationOk() (*string, bool)`

GetJobConfigurationOk returns a tuple with the JobConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobConfiguration

`func (o *JobDefinitionDto) SetJobConfiguration(v string)`

SetJobConfiguration sets JobConfiguration field to given value.

### HasJobConfiguration

`func (o *JobDefinitionDto) HasJobConfiguration() bool`

HasJobConfiguration returns a boolean if a field has been set.

### SetJobConfigurationNil

`func (o *JobDefinitionDto) SetJobConfigurationNil(b bool)`

 SetJobConfigurationNil sets the value for JobConfiguration to be an explicit nil

### UnsetJobConfiguration
`func (o *JobDefinitionDto) UnsetJobConfiguration()`

UnsetJobConfiguration ensures that no value is present for JobConfiguration, not even an explicit nil
### GetOverridingJobPriority

`func (o *JobDefinitionDto) GetOverridingJobPriority() int64`

GetOverridingJobPriority returns the OverridingJobPriority field if non-nil, zero value otherwise.

### GetOverridingJobPriorityOk

`func (o *JobDefinitionDto) GetOverridingJobPriorityOk() (*int64, bool)`

GetOverridingJobPriorityOk returns a tuple with the OverridingJobPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridingJobPriority

`func (o *JobDefinitionDto) SetOverridingJobPriority(v int64)`

SetOverridingJobPriority sets OverridingJobPriority field to given value.

### HasOverridingJobPriority

`func (o *JobDefinitionDto) HasOverridingJobPriority() bool`

HasOverridingJobPriority returns a boolean if a field has been set.

### SetOverridingJobPriorityNil

`func (o *JobDefinitionDto) SetOverridingJobPriorityNil(b bool)`

 SetOverridingJobPriorityNil sets the value for OverridingJobPriority to be an explicit nil

### UnsetOverridingJobPriority
`func (o *JobDefinitionDto) UnsetOverridingJobPriority()`

UnsetOverridingJobPriority ensures that no value is present for OverridingJobPriority, not even an explicit nil
### GetSuspended

`func (o *JobDefinitionDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *JobDefinitionDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *JobDefinitionDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *JobDefinitionDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *JobDefinitionDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *JobDefinitionDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetTenantId

`func (o *JobDefinitionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *JobDefinitionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *JobDefinitionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *JobDefinitionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *JobDefinitionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *JobDefinitionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDeploymentId

`func (o *JobDefinitionDto) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *JobDefinitionDto) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *JobDefinitionDto) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *JobDefinitionDto) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *JobDefinitionDto) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *JobDefinitionDto) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


