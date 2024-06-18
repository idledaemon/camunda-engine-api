# JobRetriesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retries** | Pointer to **NullableInt32** | The number of retries to set for the resource.  Must be &gt;&#x3D; 0. If this is 0, an incident is created and the task, or job, cannot be fetched, or acquired anymore unless the retries are increased again. Can not be null. | [optional] 
**DueDate** | Pointer to **NullableTime** | The due date to set for the job. A due date indicates when this job is ready for execution. Jobs with due dates in the past will be scheduled for execution. | [optional] 

## Methods

### NewJobRetriesDto

`func NewJobRetriesDto() *JobRetriesDto`

NewJobRetriesDto instantiates a new JobRetriesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRetriesDtoWithDefaults

`func NewJobRetriesDtoWithDefaults() *JobRetriesDto`

NewJobRetriesDtoWithDefaults instantiates a new JobRetriesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetries

`func (o *JobRetriesDto) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *JobRetriesDto) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *JobRetriesDto) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *JobRetriesDto) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *JobRetriesDto) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *JobRetriesDto) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetDueDate

`func (o *JobRetriesDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *JobRetriesDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *JobRetriesDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *JobRetriesDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *JobRetriesDto) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *JobRetriesDto) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


