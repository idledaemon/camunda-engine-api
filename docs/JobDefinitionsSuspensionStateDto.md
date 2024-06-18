# JobDefinitionsSuspensionStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeJobs** | Pointer to **NullableBool** | A &#x60;Boolean&#x60; value which indicates whether to activate or suspend also all jobs of the referenced job definitions. When the value is set to &#x60;true&#x60;, all jobs of the provided job definitions will be activated or suspended and when the value is set to &#x60;false&#x60;, the suspension state of all jobs of the provided job definitions will not be updated. | [optional] 
**ExecutionDate** | Pointer to **NullableString** | The date on which the referenced job definitions will be activated or suspended. If null, the suspension state of the given job definitions is updated immediately. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM- dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**Suspended** | Pointer to **NullableBool** | A Boolean value which indicates whether to activate or suspend a given instance  (e.g. process instance, job, job definition, or batch). When the value is set to true,  the given instance will be suspended and when the value is set to false,  the given instance will be activated. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The process definition id of the job definitions to activate or suspend. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The process definition key of the job definitions to activate or suspend. | [optional] 
**ProcessDefinitionTenantId** | Pointer to **NullableString** | Only activate or suspend job definitions of a process definition which belongs to a tenant with the given id.  Note that this parameter will only be considered  in combination with &#x60;processDefinitionKey&#x60;. | [optional] 
**ProcessDefinitionWithoutTenantId** | Pointer to **NullableBool** | Only activate or suspend job definitions of a process definition which belongs to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior.  Note that this parameter will only be considered  in combination with &#x60;processDefinitionKey&#x60;. | [optional] 

## Methods

### NewJobDefinitionsSuspensionStateDto

`func NewJobDefinitionsSuspensionStateDto() *JobDefinitionsSuspensionStateDto`

NewJobDefinitionsSuspensionStateDto instantiates a new JobDefinitionsSuspensionStateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDefinitionsSuspensionStateDtoWithDefaults

`func NewJobDefinitionsSuspensionStateDtoWithDefaults() *JobDefinitionsSuspensionStateDto`

NewJobDefinitionsSuspensionStateDtoWithDefaults instantiates a new JobDefinitionsSuspensionStateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeJobs

`func (o *JobDefinitionsSuspensionStateDto) GetIncludeJobs() bool`

GetIncludeJobs returns the IncludeJobs field if non-nil, zero value otherwise.

### GetIncludeJobsOk

`func (o *JobDefinitionsSuspensionStateDto) GetIncludeJobsOk() (*bool, bool)`

GetIncludeJobsOk returns a tuple with the IncludeJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeJobs

`func (o *JobDefinitionsSuspensionStateDto) SetIncludeJobs(v bool)`

SetIncludeJobs sets IncludeJobs field to given value.

### HasIncludeJobs

`func (o *JobDefinitionsSuspensionStateDto) HasIncludeJobs() bool`

HasIncludeJobs returns a boolean if a field has been set.

### SetIncludeJobsNil

`func (o *JobDefinitionsSuspensionStateDto) SetIncludeJobsNil(b bool)`

 SetIncludeJobsNil sets the value for IncludeJobs to be an explicit nil

### UnsetIncludeJobs
`func (o *JobDefinitionsSuspensionStateDto) UnsetIncludeJobs()`

UnsetIncludeJobs ensures that no value is present for IncludeJobs, not even an explicit nil
### GetExecutionDate

`func (o *JobDefinitionsSuspensionStateDto) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *JobDefinitionsSuspensionStateDto) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *JobDefinitionsSuspensionStateDto) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *JobDefinitionsSuspensionStateDto) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### SetExecutionDateNil

`func (o *JobDefinitionsSuspensionStateDto) SetExecutionDateNil(b bool)`

 SetExecutionDateNil sets the value for ExecutionDate to be an explicit nil

### UnsetExecutionDate
`func (o *JobDefinitionsSuspensionStateDto) UnsetExecutionDate()`

UnsetExecutionDate ensures that no value is present for ExecutionDate, not even an explicit nil
### GetSuspended

`func (o *JobDefinitionsSuspensionStateDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *JobDefinitionsSuspensionStateDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *JobDefinitionsSuspensionStateDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *JobDefinitionsSuspensionStateDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *JobDefinitionsSuspensionStateDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *JobDefinitionsSuspensionStateDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetProcessDefinitionId

`func (o *JobDefinitionsSuspensionStateDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *JobDefinitionsSuspensionStateDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *JobDefinitionsSuspensionStateDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *JobDefinitionsSuspensionStateDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *JobDefinitionsSuspensionStateDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *JobDefinitionsSuspensionStateDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *JobDefinitionsSuspensionStateDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *JobDefinitionsSuspensionStateDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *JobDefinitionsSuspensionStateDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *JobDefinitionsSuspensionStateDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *JobDefinitionsSuspensionStateDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *JobDefinitionsSuspensionStateDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionTenantId

`func (o *JobDefinitionsSuspensionStateDto) GetProcessDefinitionTenantId() string`

GetProcessDefinitionTenantId returns the ProcessDefinitionTenantId field if non-nil, zero value otherwise.

### GetProcessDefinitionTenantIdOk

`func (o *JobDefinitionsSuspensionStateDto) GetProcessDefinitionTenantIdOk() (*string, bool)`

GetProcessDefinitionTenantIdOk returns a tuple with the ProcessDefinitionTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionTenantId

`func (o *JobDefinitionsSuspensionStateDto) SetProcessDefinitionTenantId(v string)`

SetProcessDefinitionTenantId sets ProcessDefinitionTenantId field to given value.

### HasProcessDefinitionTenantId

`func (o *JobDefinitionsSuspensionStateDto) HasProcessDefinitionTenantId() bool`

HasProcessDefinitionTenantId returns a boolean if a field has been set.

### SetProcessDefinitionTenantIdNil

`func (o *JobDefinitionsSuspensionStateDto) SetProcessDefinitionTenantIdNil(b bool)`

 SetProcessDefinitionTenantIdNil sets the value for ProcessDefinitionTenantId to be an explicit nil

### UnsetProcessDefinitionTenantId
`func (o *JobDefinitionsSuspensionStateDto) UnsetProcessDefinitionTenantId()`

UnsetProcessDefinitionTenantId ensures that no value is present for ProcessDefinitionTenantId, not even an explicit nil
### GetProcessDefinitionWithoutTenantId

`func (o *JobDefinitionsSuspensionStateDto) GetProcessDefinitionWithoutTenantId() bool`

GetProcessDefinitionWithoutTenantId returns the ProcessDefinitionWithoutTenantId field if non-nil, zero value otherwise.

### GetProcessDefinitionWithoutTenantIdOk

`func (o *JobDefinitionsSuspensionStateDto) GetProcessDefinitionWithoutTenantIdOk() (*bool, bool)`

GetProcessDefinitionWithoutTenantIdOk returns a tuple with the ProcessDefinitionWithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionWithoutTenantId

`func (o *JobDefinitionsSuspensionStateDto) SetProcessDefinitionWithoutTenantId(v bool)`

SetProcessDefinitionWithoutTenantId sets ProcessDefinitionWithoutTenantId field to given value.

### HasProcessDefinitionWithoutTenantId

`func (o *JobDefinitionsSuspensionStateDto) HasProcessDefinitionWithoutTenantId() bool`

HasProcessDefinitionWithoutTenantId returns a boolean if a field has been set.

### SetProcessDefinitionWithoutTenantIdNil

`func (o *JobDefinitionsSuspensionStateDto) SetProcessDefinitionWithoutTenantIdNil(b bool)`

 SetProcessDefinitionWithoutTenantIdNil sets the value for ProcessDefinitionWithoutTenantId to be an explicit nil

### UnsetProcessDefinitionWithoutTenantId
`func (o *JobDefinitionsSuspensionStateDto) UnsetProcessDefinitionWithoutTenantId()`

UnsetProcessDefinitionWithoutTenantId ensures that no value is present for ProcessDefinitionWithoutTenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


