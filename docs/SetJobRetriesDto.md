# SetJobRetriesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DueDate** | Pointer to **NullableTime** | The due date to set for the job. A due date indicates when this job is ready for execution. Jobs with due dates in the past will be scheduled for execution. | [optional] 
**Retries** | Pointer to **NullableInt32** | The number of retries to set for the resource.  Must be &gt;&#x3D; 0. If this is 0, an incident is created and the task, or job, cannot be fetched, or acquired anymore unless the retries are increased again. Can not be null. | [optional] 
**JobIds** | Pointer to **[]string** | A list of job ids to set retries for. | [optional] 
**JobQuery** | Pointer to [**JobQueryDto**](JobQueryDto.md) |  | [optional] 

## Methods

### NewSetJobRetriesDto

`func NewSetJobRetriesDto() *SetJobRetriesDto`

NewSetJobRetriesDto instantiates a new SetJobRetriesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetJobRetriesDtoWithDefaults

`func NewSetJobRetriesDtoWithDefaults() *SetJobRetriesDto`

NewSetJobRetriesDtoWithDefaults instantiates a new SetJobRetriesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDueDate

`func (o *SetJobRetriesDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *SetJobRetriesDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *SetJobRetriesDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *SetJobRetriesDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *SetJobRetriesDto) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *SetJobRetriesDto) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetRetries

`func (o *SetJobRetriesDto) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *SetJobRetriesDto) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *SetJobRetriesDto) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *SetJobRetriesDto) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *SetJobRetriesDto) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *SetJobRetriesDto) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetJobIds

`func (o *SetJobRetriesDto) GetJobIds() []string`

GetJobIds returns the JobIds field if non-nil, zero value otherwise.

### GetJobIdsOk

`func (o *SetJobRetriesDto) GetJobIdsOk() (*[]string, bool)`

GetJobIdsOk returns a tuple with the JobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIds

`func (o *SetJobRetriesDto) SetJobIds(v []string)`

SetJobIds sets JobIds field to given value.

### HasJobIds

`func (o *SetJobRetriesDto) HasJobIds() bool`

HasJobIds returns a boolean if a field has been set.

### SetJobIdsNil

`func (o *SetJobRetriesDto) SetJobIdsNil(b bool)`

 SetJobIdsNil sets the value for JobIds to be an explicit nil

### UnsetJobIds
`func (o *SetJobRetriesDto) UnsetJobIds()`

UnsetJobIds ensures that no value is present for JobIds, not even an explicit nil
### GetJobQuery

`func (o *SetJobRetriesDto) GetJobQuery() JobQueryDto`

GetJobQuery returns the JobQuery field if non-nil, zero value otherwise.

### GetJobQueryOk

`func (o *SetJobRetriesDto) GetJobQueryOk() (*JobQueryDto, bool)`

GetJobQueryOk returns a tuple with the JobQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQuery

`func (o *SetJobRetriesDto) SetJobQuery(v JobQueryDto)`

SetJobQuery sets JobQuery field to given value.

### HasJobQuery

`func (o *SetJobRetriesDto) HasJobQuery() bool`

HasJobQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


