# JobDefinitionQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobDefinitionId** | Pointer to **NullableString** | Filter by job definition id. | [optional] 
**ActivityIdIn** | Pointer to **[]string** | Only include job definitions which belong to one of the passed activity ids. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Only include job definitions which exist for the given process definition id. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Only include job definitions which exist for the given process definition key. | [optional] 
**JobType** | Pointer to **NullableString** | Only include job definitions which exist for the given job type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job types. | [optional] 
**JobConfiguration** | Pointer to **NullableString** | Only include job definitions which exist for the given job configuration. For example: for timer jobs it is the timer configuration. | [optional] 
**Active** | Pointer to **NullableBool** | Only include active job definitions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Suspended** | Pointer to **NullableBool** | Only include suspended job definitions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**WithOverridingJobPriority** | Pointer to **NullableBool** | Only include job definitions that have an overriding job priority defined. The only effective value is &#x60;true&#x60;. If set to &#x60;false&#x60;, this filter is not applied. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Only include job definitions which belong to one of the passed tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include job definitions which belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**IncludeJobDefinitionsWithoutTenantId** | Pointer to **NullableBool** | Include job definitions which belong to no tenant. Can be used in combination with &#x60;tenantIdIn&#x60;. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Sorting** | Pointer to [**[]JobDefinitionQueryDtoSortingInner**](JobDefinitionQueryDtoSortingInner.md) | An array of criteria to sort the result by. Each element of the array is                        an object that specifies one ordering. The position in the array                        identifies the rank of an ordering, i.e., whether it is primary, secondary,                        etc. Sorting has no effect for &#x60;count&#x60; endpoints. | [optional] 

## Methods

### NewJobDefinitionQueryDto

`func NewJobDefinitionQueryDto() *JobDefinitionQueryDto`

NewJobDefinitionQueryDto instantiates a new JobDefinitionQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDefinitionQueryDtoWithDefaults

`func NewJobDefinitionQueryDtoWithDefaults() *JobDefinitionQueryDto`

NewJobDefinitionQueryDtoWithDefaults instantiates a new JobDefinitionQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobDefinitionId

`func (o *JobDefinitionQueryDto) GetJobDefinitionId() string`

GetJobDefinitionId returns the JobDefinitionId field if non-nil, zero value otherwise.

### GetJobDefinitionIdOk

`func (o *JobDefinitionQueryDto) GetJobDefinitionIdOk() (*string, bool)`

GetJobDefinitionIdOk returns a tuple with the JobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionId

`func (o *JobDefinitionQueryDto) SetJobDefinitionId(v string)`

SetJobDefinitionId sets JobDefinitionId field to given value.

### HasJobDefinitionId

`func (o *JobDefinitionQueryDto) HasJobDefinitionId() bool`

HasJobDefinitionId returns a boolean if a field has been set.

### SetJobDefinitionIdNil

`func (o *JobDefinitionQueryDto) SetJobDefinitionIdNil(b bool)`

 SetJobDefinitionIdNil sets the value for JobDefinitionId to be an explicit nil

### UnsetJobDefinitionId
`func (o *JobDefinitionQueryDto) UnsetJobDefinitionId()`

UnsetJobDefinitionId ensures that no value is present for JobDefinitionId, not even an explicit nil
### GetActivityIdIn

`func (o *JobDefinitionQueryDto) GetActivityIdIn() []string`

GetActivityIdIn returns the ActivityIdIn field if non-nil, zero value otherwise.

### GetActivityIdInOk

`func (o *JobDefinitionQueryDto) GetActivityIdInOk() (*[]string, bool)`

GetActivityIdInOk returns a tuple with the ActivityIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIdIn

`func (o *JobDefinitionQueryDto) SetActivityIdIn(v []string)`

SetActivityIdIn sets ActivityIdIn field to given value.

### HasActivityIdIn

`func (o *JobDefinitionQueryDto) HasActivityIdIn() bool`

HasActivityIdIn returns a boolean if a field has been set.

### SetActivityIdInNil

`func (o *JobDefinitionQueryDto) SetActivityIdInNil(b bool)`

 SetActivityIdInNil sets the value for ActivityIdIn to be an explicit nil

### UnsetActivityIdIn
`func (o *JobDefinitionQueryDto) UnsetActivityIdIn()`

UnsetActivityIdIn ensures that no value is present for ActivityIdIn, not even an explicit nil
### GetProcessDefinitionId

`func (o *JobDefinitionQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *JobDefinitionQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *JobDefinitionQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *JobDefinitionQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *JobDefinitionQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *JobDefinitionQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *JobDefinitionQueryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *JobDefinitionQueryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *JobDefinitionQueryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *JobDefinitionQueryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *JobDefinitionQueryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *JobDefinitionQueryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetJobType

`func (o *JobDefinitionQueryDto) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobDefinitionQueryDto) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobDefinitionQueryDto) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *JobDefinitionQueryDto) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *JobDefinitionQueryDto) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *JobDefinitionQueryDto) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetJobConfiguration

`func (o *JobDefinitionQueryDto) GetJobConfiguration() string`

GetJobConfiguration returns the JobConfiguration field if non-nil, zero value otherwise.

### GetJobConfigurationOk

`func (o *JobDefinitionQueryDto) GetJobConfigurationOk() (*string, bool)`

GetJobConfigurationOk returns a tuple with the JobConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobConfiguration

`func (o *JobDefinitionQueryDto) SetJobConfiguration(v string)`

SetJobConfiguration sets JobConfiguration field to given value.

### HasJobConfiguration

`func (o *JobDefinitionQueryDto) HasJobConfiguration() bool`

HasJobConfiguration returns a boolean if a field has been set.

### SetJobConfigurationNil

`func (o *JobDefinitionQueryDto) SetJobConfigurationNil(b bool)`

 SetJobConfigurationNil sets the value for JobConfiguration to be an explicit nil

### UnsetJobConfiguration
`func (o *JobDefinitionQueryDto) UnsetJobConfiguration()`

UnsetJobConfiguration ensures that no value is present for JobConfiguration, not even an explicit nil
### GetActive

`func (o *JobDefinitionQueryDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *JobDefinitionQueryDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *JobDefinitionQueryDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *JobDefinitionQueryDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *JobDefinitionQueryDto) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *JobDefinitionQueryDto) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetSuspended

`func (o *JobDefinitionQueryDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *JobDefinitionQueryDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *JobDefinitionQueryDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *JobDefinitionQueryDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *JobDefinitionQueryDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *JobDefinitionQueryDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetWithOverridingJobPriority

`func (o *JobDefinitionQueryDto) GetWithOverridingJobPriority() bool`

GetWithOverridingJobPriority returns the WithOverridingJobPriority field if non-nil, zero value otherwise.

### GetWithOverridingJobPriorityOk

`func (o *JobDefinitionQueryDto) GetWithOverridingJobPriorityOk() (*bool, bool)`

GetWithOverridingJobPriorityOk returns a tuple with the WithOverridingJobPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithOverridingJobPriority

`func (o *JobDefinitionQueryDto) SetWithOverridingJobPriority(v bool)`

SetWithOverridingJobPriority sets WithOverridingJobPriority field to given value.

### HasWithOverridingJobPriority

`func (o *JobDefinitionQueryDto) HasWithOverridingJobPriority() bool`

HasWithOverridingJobPriority returns a boolean if a field has been set.

### SetWithOverridingJobPriorityNil

`func (o *JobDefinitionQueryDto) SetWithOverridingJobPriorityNil(b bool)`

 SetWithOverridingJobPriorityNil sets the value for WithOverridingJobPriority to be an explicit nil

### UnsetWithOverridingJobPriority
`func (o *JobDefinitionQueryDto) UnsetWithOverridingJobPriority()`

UnsetWithOverridingJobPriority ensures that no value is present for WithOverridingJobPriority, not even an explicit nil
### GetTenantIdIn

`func (o *JobDefinitionQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *JobDefinitionQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *JobDefinitionQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *JobDefinitionQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *JobDefinitionQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *JobDefinitionQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *JobDefinitionQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *JobDefinitionQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *JobDefinitionQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *JobDefinitionQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *JobDefinitionQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *JobDefinitionQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetIncludeJobDefinitionsWithoutTenantId

`func (o *JobDefinitionQueryDto) GetIncludeJobDefinitionsWithoutTenantId() bool`

GetIncludeJobDefinitionsWithoutTenantId returns the IncludeJobDefinitionsWithoutTenantId field if non-nil, zero value otherwise.

### GetIncludeJobDefinitionsWithoutTenantIdOk

`func (o *JobDefinitionQueryDto) GetIncludeJobDefinitionsWithoutTenantIdOk() (*bool, bool)`

GetIncludeJobDefinitionsWithoutTenantIdOk returns a tuple with the IncludeJobDefinitionsWithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeJobDefinitionsWithoutTenantId

`func (o *JobDefinitionQueryDto) SetIncludeJobDefinitionsWithoutTenantId(v bool)`

SetIncludeJobDefinitionsWithoutTenantId sets IncludeJobDefinitionsWithoutTenantId field to given value.

### HasIncludeJobDefinitionsWithoutTenantId

`func (o *JobDefinitionQueryDto) HasIncludeJobDefinitionsWithoutTenantId() bool`

HasIncludeJobDefinitionsWithoutTenantId returns a boolean if a field has been set.

### SetIncludeJobDefinitionsWithoutTenantIdNil

`func (o *JobDefinitionQueryDto) SetIncludeJobDefinitionsWithoutTenantIdNil(b bool)`

 SetIncludeJobDefinitionsWithoutTenantIdNil sets the value for IncludeJobDefinitionsWithoutTenantId to be an explicit nil

### UnsetIncludeJobDefinitionsWithoutTenantId
`func (o *JobDefinitionQueryDto) UnsetIncludeJobDefinitionsWithoutTenantId()`

UnsetIncludeJobDefinitionsWithoutTenantId ensures that no value is present for IncludeJobDefinitionsWithoutTenantId, not even an explicit nil
### GetSorting

`func (o *JobDefinitionQueryDto) GetSorting() []JobDefinitionQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *JobDefinitionQueryDto) GetSortingOk() (*[]JobDefinitionQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *JobDefinitionQueryDto) SetSorting(v []JobDefinitionQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *JobDefinitionQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *JobDefinitionQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *JobDefinitionQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


