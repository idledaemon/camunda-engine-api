# JobDefinitionPriorityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **NullableInt64** | The new execution priority number for jobs of the given definition. The definition&#39;s priority can be reset by using the value &#x60;null&#x60;. In that case, the job definition&#39;s priority no longer applies but a new job&#39;s priority is determined as specified in the process model. | [optional] 
**IncludeJobs** | Pointer to **NullableBool** | A boolean value indicating whether existing jobs of the given definition should receive the priority as well. Default value is &#x60;false&#x60;. Can only be &#x60;true&#x60; when the __priority__ parameter is not &#x60;null&#x60;. | [optional] 

## Methods

### NewJobDefinitionPriorityDto

`func NewJobDefinitionPriorityDto() *JobDefinitionPriorityDto`

NewJobDefinitionPriorityDto instantiates a new JobDefinitionPriorityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDefinitionPriorityDtoWithDefaults

`func NewJobDefinitionPriorityDtoWithDefaults() *JobDefinitionPriorityDto`

NewJobDefinitionPriorityDtoWithDefaults instantiates a new JobDefinitionPriorityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *JobDefinitionPriorityDto) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *JobDefinitionPriorityDto) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *JobDefinitionPriorityDto) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *JobDefinitionPriorityDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *JobDefinitionPriorityDto) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *JobDefinitionPriorityDto) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetIncludeJobs

`func (o *JobDefinitionPriorityDto) GetIncludeJobs() bool`

GetIncludeJobs returns the IncludeJobs field if non-nil, zero value otherwise.

### GetIncludeJobsOk

`func (o *JobDefinitionPriorityDto) GetIncludeJobsOk() (*bool, bool)`

GetIncludeJobsOk returns a tuple with the IncludeJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeJobs

`func (o *JobDefinitionPriorityDto) SetIncludeJobs(v bool)`

SetIncludeJobs sets IncludeJobs field to given value.

### HasIncludeJobs

`func (o *JobDefinitionPriorityDto) HasIncludeJobs() bool`

HasIncludeJobs returns a boolean if a field has been set.

### SetIncludeJobsNil

`func (o *JobDefinitionPriorityDto) SetIncludeJobsNil(b bool)`

 SetIncludeJobsNil sets the value for IncludeJobs to be an explicit nil

### UnsetIncludeJobs
`func (o *JobDefinitionPriorityDto) UnsetIncludeJobs()`

UnsetIncludeJobs ensures that no value is present for IncludeJobs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


