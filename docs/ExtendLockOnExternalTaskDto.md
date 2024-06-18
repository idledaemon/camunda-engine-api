# ExtendLockOnExternalTaskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerId** | Pointer to **string** | **Mandatory.** The ID of the worker who is performing the operation on the external task. If the task is already locked, must match the id of the worker who has most recently locked the task. | [optional] 
**NewDuration** | Pointer to **NullableInt64** | An amount of time (in milliseconds). This is the new lock duration starting from the current moment. | [optional] 

## Methods

### NewExtendLockOnExternalTaskDto

`func NewExtendLockOnExternalTaskDto() *ExtendLockOnExternalTaskDto`

NewExtendLockOnExternalTaskDto instantiates a new ExtendLockOnExternalTaskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendLockOnExternalTaskDtoWithDefaults

`func NewExtendLockOnExternalTaskDtoWithDefaults() *ExtendLockOnExternalTaskDto`

NewExtendLockOnExternalTaskDtoWithDefaults instantiates a new ExtendLockOnExternalTaskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerId

`func (o *ExtendLockOnExternalTaskDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ExtendLockOnExternalTaskDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ExtendLockOnExternalTaskDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *ExtendLockOnExternalTaskDto) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetNewDuration

`func (o *ExtendLockOnExternalTaskDto) GetNewDuration() int64`

GetNewDuration returns the NewDuration field if non-nil, zero value otherwise.

### GetNewDurationOk

`func (o *ExtendLockOnExternalTaskDto) GetNewDurationOk() (*int64, bool)`

GetNewDurationOk returns a tuple with the NewDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDuration

`func (o *ExtendLockOnExternalTaskDto) SetNewDuration(v int64)`

SetNewDuration sets NewDuration field to given value.

### HasNewDuration

`func (o *ExtendLockOnExternalTaskDto) HasNewDuration() bool`

HasNewDuration returns a boolean if a field has been set.

### SetNewDurationNil

`func (o *ExtendLockOnExternalTaskDto) SetNewDurationNil(b bool)`

 SetNewDurationNil sets the value for NewDuration to be an explicit nil

### UnsetNewDuration
`func (o *ExtendLockOnExternalTaskDto) UnsetNewDuration()`

UnsetNewDuration ensures that no value is present for NewDuration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


