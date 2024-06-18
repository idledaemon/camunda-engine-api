# JobDuedateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duedate** | Pointer to **NullableTime** | The date to set when the job has the next execution. | [optional] 
**Cascade** | Pointer to **NullableBool** | A boolean value to indicate if modifications to the due date should cascade to subsequent jobs. (e.g. Modify the due date of a timer by +15 minutes. This flag indicates if a +15 minutes should be applied to all subsequent timers.) This flag only affects timer jobs and only works if due date is not null. Default: &#x60;false&#x60; | [optional] 

## Methods

### NewJobDuedateDto

`func NewJobDuedateDto() *JobDuedateDto`

NewJobDuedateDto instantiates a new JobDuedateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDuedateDtoWithDefaults

`func NewJobDuedateDtoWithDefaults() *JobDuedateDto`

NewJobDuedateDtoWithDefaults instantiates a new JobDuedateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuedate

`func (o *JobDuedateDto) GetDuedate() time.Time`

GetDuedate returns the Duedate field if non-nil, zero value otherwise.

### GetDuedateOk

`func (o *JobDuedateDto) GetDuedateOk() (*time.Time, bool)`

GetDuedateOk returns a tuple with the Duedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuedate

`func (o *JobDuedateDto) SetDuedate(v time.Time)`

SetDuedate sets Duedate field to given value.

### HasDuedate

`func (o *JobDuedateDto) HasDuedate() bool`

HasDuedate returns a boolean if a field has been set.

### SetDuedateNil

`func (o *JobDuedateDto) SetDuedateNil(b bool)`

 SetDuedateNil sets the value for Duedate to be an explicit nil

### UnsetDuedate
`func (o *JobDuedateDto) UnsetDuedate()`

UnsetDuedate ensures that no value is present for Duedate, not even an explicit nil
### GetCascade

`func (o *JobDuedateDto) GetCascade() bool`

GetCascade returns the Cascade field if non-nil, zero value otherwise.

### GetCascadeOk

`func (o *JobDuedateDto) GetCascadeOk() (*bool, bool)`

GetCascadeOk returns a tuple with the Cascade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascade

`func (o *JobDuedateDto) SetCascade(v bool)`

SetCascade sets Cascade field to given value.

### HasCascade

`func (o *JobDuedateDto) HasCascade() bool`

HasCascade returns a boolean if a field has been set.

### SetCascadeNil

`func (o *JobDuedateDto) SetCascadeNil(b bool)`

 SetCascadeNil sets the value for Cascade to be an explicit nil

### UnsetCascade
`func (o *JobDuedateDto) UnsetCascade()`

UnsetCascade ensures that no value is present for Cascade, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


