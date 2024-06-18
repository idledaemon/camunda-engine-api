# HandleExternalTaskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerId** | Pointer to **string** | **Mandatory.** The ID of the worker who is performing the operation on the external task. If the task is already locked, must match the id of the worker who has most recently locked the task. | [optional] 

## Methods

### NewHandleExternalTaskDto

`func NewHandleExternalTaskDto() *HandleExternalTaskDto`

NewHandleExternalTaskDto instantiates a new HandleExternalTaskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandleExternalTaskDtoWithDefaults

`func NewHandleExternalTaskDtoWithDefaults() *HandleExternalTaskDto`

NewHandleExternalTaskDtoWithDefaults instantiates a new HandleExternalTaskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerId

`func (o *HandleExternalTaskDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *HandleExternalTaskDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *HandleExternalTaskDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *HandleExternalTaskDto) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


