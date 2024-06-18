# JobSuspensionStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suspended** | Pointer to **NullableBool** | A Boolean value which indicates whether to activate or suspend a given instance  (e.g. process instance, job, job definition, or batch). When the value is set to true,  the given instance will be suspended and when the value is set to false,  the given instance will be activated. | [optional] 
**JobDefinitionId** | Pointer to **NullableString** | The job definition id of the jobs to activate or suspend. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The process definition id of the jobs to activate or suspend. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The process instance id of the jobs to activate or suspend. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The process definition key of the jobs to activate or suspend. | [optional] 
**ProcessDefinitionTenantId** | Pointer to **NullableString** | Only activate or suspend jobs of a process definition which belongs to a tenant with the given id. Works only when selecting with &#x60;processDefinitionKey&#x60;. | [optional] 
**ProcessDefinitionWithoutTenantId** | Pointer to **NullableBool** | Only activate or suspend jobs of a process definition which belongs to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. Works only when selecting with &#x60;processDefinitionKey&#x60;. | [optional] 

## Methods

### NewJobSuspensionStateDto

`func NewJobSuspensionStateDto() *JobSuspensionStateDto`

NewJobSuspensionStateDto instantiates a new JobSuspensionStateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobSuspensionStateDtoWithDefaults

`func NewJobSuspensionStateDtoWithDefaults() *JobSuspensionStateDto`

NewJobSuspensionStateDtoWithDefaults instantiates a new JobSuspensionStateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuspended

`func (o *JobSuspensionStateDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *JobSuspensionStateDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *JobSuspensionStateDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *JobSuspensionStateDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *JobSuspensionStateDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *JobSuspensionStateDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetJobDefinitionId

`func (o *JobSuspensionStateDto) GetJobDefinitionId() string`

GetJobDefinitionId returns the JobDefinitionId field if non-nil, zero value otherwise.

### GetJobDefinitionIdOk

`func (o *JobSuspensionStateDto) GetJobDefinitionIdOk() (*string, bool)`

GetJobDefinitionIdOk returns a tuple with the JobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionId

`func (o *JobSuspensionStateDto) SetJobDefinitionId(v string)`

SetJobDefinitionId sets JobDefinitionId field to given value.

### HasJobDefinitionId

`func (o *JobSuspensionStateDto) HasJobDefinitionId() bool`

HasJobDefinitionId returns a boolean if a field has been set.

### SetJobDefinitionIdNil

`func (o *JobSuspensionStateDto) SetJobDefinitionIdNil(b bool)`

 SetJobDefinitionIdNil sets the value for JobDefinitionId to be an explicit nil

### UnsetJobDefinitionId
`func (o *JobSuspensionStateDto) UnsetJobDefinitionId()`

UnsetJobDefinitionId ensures that no value is present for JobDefinitionId, not even an explicit nil
### GetProcessDefinitionId

`func (o *JobSuspensionStateDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *JobSuspensionStateDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *JobSuspensionStateDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *JobSuspensionStateDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *JobSuspensionStateDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *JobSuspensionStateDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessInstanceId

`func (o *JobSuspensionStateDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *JobSuspensionStateDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *JobSuspensionStateDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *JobSuspensionStateDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *JobSuspensionStateDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *JobSuspensionStateDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *JobSuspensionStateDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *JobSuspensionStateDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *JobSuspensionStateDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *JobSuspensionStateDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *JobSuspensionStateDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *JobSuspensionStateDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionTenantId

`func (o *JobSuspensionStateDto) GetProcessDefinitionTenantId() string`

GetProcessDefinitionTenantId returns the ProcessDefinitionTenantId field if non-nil, zero value otherwise.

### GetProcessDefinitionTenantIdOk

`func (o *JobSuspensionStateDto) GetProcessDefinitionTenantIdOk() (*string, bool)`

GetProcessDefinitionTenantIdOk returns a tuple with the ProcessDefinitionTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionTenantId

`func (o *JobSuspensionStateDto) SetProcessDefinitionTenantId(v string)`

SetProcessDefinitionTenantId sets ProcessDefinitionTenantId field to given value.

### HasProcessDefinitionTenantId

`func (o *JobSuspensionStateDto) HasProcessDefinitionTenantId() bool`

HasProcessDefinitionTenantId returns a boolean if a field has been set.

### SetProcessDefinitionTenantIdNil

`func (o *JobSuspensionStateDto) SetProcessDefinitionTenantIdNil(b bool)`

 SetProcessDefinitionTenantIdNil sets the value for ProcessDefinitionTenantId to be an explicit nil

### UnsetProcessDefinitionTenantId
`func (o *JobSuspensionStateDto) UnsetProcessDefinitionTenantId()`

UnsetProcessDefinitionTenantId ensures that no value is present for ProcessDefinitionTenantId, not even an explicit nil
### GetProcessDefinitionWithoutTenantId

`func (o *JobSuspensionStateDto) GetProcessDefinitionWithoutTenantId() bool`

GetProcessDefinitionWithoutTenantId returns the ProcessDefinitionWithoutTenantId field if non-nil, zero value otherwise.

### GetProcessDefinitionWithoutTenantIdOk

`func (o *JobSuspensionStateDto) GetProcessDefinitionWithoutTenantIdOk() (*bool, bool)`

GetProcessDefinitionWithoutTenantIdOk returns a tuple with the ProcessDefinitionWithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionWithoutTenantId

`func (o *JobSuspensionStateDto) SetProcessDefinitionWithoutTenantId(v bool)`

SetProcessDefinitionWithoutTenantId sets ProcessDefinitionWithoutTenantId field to given value.

### HasProcessDefinitionWithoutTenantId

`func (o *JobSuspensionStateDto) HasProcessDefinitionWithoutTenantId() bool`

HasProcessDefinitionWithoutTenantId returns a boolean if a field has been set.

### SetProcessDefinitionWithoutTenantIdNil

`func (o *JobSuspensionStateDto) SetProcessDefinitionWithoutTenantIdNil(b bool)`

 SetProcessDefinitionWithoutTenantIdNil sets the value for ProcessDefinitionWithoutTenantId to be an explicit nil

### UnsetProcessDefinitionWithoutTenantId
`func (o *JobSuspensionStateDto) UnsetProcessDefinitionWithoutTenantId()`

UnsetProcessDefinitionWithoutTenantId ensures that no value is present for ProcessDefinitionWithoutTenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


