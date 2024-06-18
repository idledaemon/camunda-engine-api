# LockExternalTaskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerId** | Pointer to **string** | **Mandatory.** The ID of the worker who is performing the operation on the external task. If the task is already locked, must match the id of the worker who has most recently locked the task. | [optional] 
**LockDuration** | Pointer to **int64** | The duration to lock the external task for in milliseconds. **Note:** Attempting to lock an already locked external task with the same &#x60;workerId&#x60; will succeed and a new lock duration will be set, starting from the current moment. | [optional] 

## Methods

### NewLockExternalTaskDto

`func NewLockExternalTaskDto() *LockExternalTaskDto`

NewLockExternalTaskDto instantiates a new LockExternalTaskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockExternalTaskDtoWithDefaults

`func NewLockExternalTaskDtoWithDefaults() *LockExternalTaskDto`

NewLockExternalTaskDtoWithDefaults instantiates a new LockExternalTaskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerId

`func (o *LockExternalTaskDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *LockExternalTaskDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *LockExternalTaskDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *LockExternalTaskDto) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetLockDuration

`func (o *LockExternalTaskDto) GetLockDuration() int64`

GetLockDuration returns the LockDuration field if non-nil, zero value otherwise.

### GetLockDurationOk

`func (o *LockExternalTaskDto) GetLockDurationOk() (*int64, bool)`

GetLockDurationOk returns a tuple with the LockDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockDuration

`func (o *LockExternalTaskDto) SetLockDuration(v int64)`

SetLockDuration sets LockDuration field to given value.

### HasLockDuration

`func (o *LockExternalTaskDto) HasLockDuration() bool`

HasLockDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


