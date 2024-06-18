# JobDefinitionSuspensionStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suspended** | Pointer to **NullableBool** | A Boolean value which indicates whether to activate or suspend a given instance  (e.g. process instance, job, job definition, or batch). When the value is set to true,  the given instance will be suspended and when the value is set to false,  the given instance will be activated. | [optional] 
**IncludeJobs** | Pointer to **NullableBool** | A &#x60;Boolean&#x60; value which indicates whether to activate or suspend also all jobs of the referenced job definitions. When the value is set to &#x60;true&#x60;, all jobs of the provided job definitions will be activated or suspended and when the value is set to &#x60;false&#x60;, the suspension state of all jobs of the provided job definitions will not be updated. | [optional] 
**ExecutionDate** | Pointer to **NullableString** | The date on which the referenced job definitions will be activated or suspended. If null, the suspension state of the given job definitions is updated immediately. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM- dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 

## Methods

### NewJobDefinitionSuspensionStateDto

`func NewJobDefinitionSuspensionStateDto() *JobDefinitionSuspensionStateDto`

NewJobDefinitionSuspensionStateDto instantiates a new JobDefinitionSuspensionStateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDefinitionSuspensionStateDtoWithDefaults

`func NewJobDefinitionSuspensionStateDtoWithDefaults() *JobDefinitionSuspensionStateDto`

NewJobDefinitionSuspensionStateDtoWithDefaults instantiates a new JobDefinitionSuspensionStateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuspended

`func (o *JobDefinitionSuspensionStateDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *JobDefinitionSuspensionStateDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *JobDefinitionSuspensionStateDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *JobDefinitionSuspensionStateDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *JobDefinitionSuspensionStateDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *JobDefinitionSuspensionStateDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetIncludeJobs

`func (o *JobDefinitionSuspensionStateDto) GetIncludeJobs() bool`

GetIncludeJobs returns the IncludeJobs field if non-nil, zero value otherwise.

### GetIncludeJobsOk

`func (o *JobDefinitionSuspensionStateDto) GetIncludeJobsOk() (*bool, bool)`

GetIncludeJobsOk returns a tuple with the IncludeJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeJobs

`func (o *JobDefinitionSuspensionStateDto) SetIncludeJobs(v bool)`

SetIncludeJobs sets IncludeJobs field to given value.

### HasIncludeJobs

`func (o *JobDefinitionSuspensionStateDto) HasIncludeJobs() bool`

HasIncludeJobs returns a boolean if a field has been set.

### SetIncludeJobsNil

`func (o *JobDefinitionSuspensionStateDto) SetIncludeJobsNil(b bool)`

 SetIncludeJobsNil sets the value for IncludeJobs to be an explicit nil

### UnsetIncludeJobs
`func (o *JobDefinitionSuspensionStateDto) UnsetIncludeJobs()`

UnsetIncludeJobs ensures that no value is present for IncludeJobs, not even an explicit nil
### GetExecutionDate

`func (o *JobDefinitionSuspensionStateDto) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *JobDefinitionSuspensionStateDto) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *JobDefinitionSuspensionStateDto) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *JobDefinitionSuspensionStateDto) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### SetExecutionDateNil

`func (o *JobDefinitionSuspensionStateDto) SetExecutionDateNil(b bool)`

 SetExecutionDateNil sets the value for ExecutionDate to be an explicit nil

### UnsetExecutionDate
`func (o *JobDefinitionSuspensionStateDto) UnsetExecutionDate()`

UnsetExecutionDate ensures that no value is present for ExecutionDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


