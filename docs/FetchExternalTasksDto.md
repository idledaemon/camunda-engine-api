# FetchExternalTasksDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerId** | **string** | **Mandatory.** The id of the worker on which behalf tasks are fetched. The returned tasks are locked for that worker and can only be completed when providing the same worker id. | 
**MaxTasks** | **NullableInt32** | **Mandatory.** The maximum number of tasks to return. | 
**UsePriority** | Pointer to **NullableBool** | A &#x60;boolean&#x60; value, which indicates whether the task should be fetched based on its priority or arbitrarily. | [optional] 
**AsyncResponseTimeout** | Pointer to **NullableInt64** | The [Long Polling](https://docs.camunda.org/manual/7.21/user-guide/process-engine/external-tasks/#long-polling-to-fetch-and-lock-external-tasks) timeout in milliseconds.  **Note:** The value cannot be set larger than 1.800.000 milliseconds (corresponds to 30 minutes). | [optional] 
**Topics** | Pointer to [**[]FetchExternalTaskTopicDto**](FetchExternalTaskTopicDto.md) | A JSON array of topic objects for which external tasks should be fetched. The returned tasks may be arbitrarily distributed among these topics. Each topic object has the following properties: | [optional] 
**Sorting** | Pointer to [**[]FetchExternalTasksDtoSortingInner**](FetchExternalTasksDtoSortingInner.md) | Apply sorting of the result | [optional] 

## Methods

### NewFetchExternalTasksDto

`func NewFetchExternalTasksDto(workerId string, maxTasks NullableInt32, ) *FetchExternalTasksDto`

NewFetchExternalTasksDto instantiates a new FetchExternalTasksDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchExternalTasksDtoWithDefaults

`func NewFetchExternalTasksDtoWithDefaults() *FetchExternalTasksDto`

NewFetchExternalTasksDtoWithDefaults instantiates a new FetchExternalTasksDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerId

`func (o *FetchExternalTasksDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *FetchExternalTasksDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *FetchExternalTasksDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.


### GetMaxTasks

`func (o *FetchExternalTasksDto) GetMaxTasks() int32`

GetMaxTasks returns the MaxTasks field if non-nil, zero value otherwise.

### GetMaxTasksOk

`func (o *FetchExternalTasksDto) GetMaxTasksOk() (*int32, bool)`

GetMaxTasksOk returns a tuple with the MaxTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTasks

`func (o *FetchExternalTasksDto) SetMaxTasks(v int32)`

SetMaxTasks sets MaxTasks field to given value.


### SetMaxTasksNil

`func (o *FetchExternalTasksDto) SetMaxTasksNil(b bool)`

 SetMaxTasksNil sets the value for MaxTasks to be an explicit nil

### UnsetMaxTasks
`func (o *FetchExternalTasksDto) UnsetMaxTasks()`

UnsetMaxTasks ensures that no value is present for MaxTasks, not even an explicit nil
### GetUsePriority

`func (o *FetchExternalTasksDto) GetUsePriority() bool`

GetUsePriority returns the UsePriority field if non-nil, zero value otherwise.

### GetUsePriorityOk

`func (o *FetchExternalTasksDto) GetUsePriorityOk() (*bool, bool)`

GetUsePriorityOk returns a tuple with the UsePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePriority

`func (o *FetchExternalTasksDto) SetUsePriority(v bool)`

SetUsePriority sets UsePriority field to given value.

### HasUsePriority

`func (o *FetchExternalTasksDto) HasUsePriority() bool`

HasUsePriority returns a boolean if a field has been set.

### SetUsePriorityNil

`func (o *FetchExternalTasksDto) SetUsePriorityNil(b bool)`

 SetUsePriorityNil sets the value for UsePriority to be an explicit nil

### UnsetUsePriority
`func (o *FetchExternalTasksDto) UnsetUsePriority()`

UnsetUsePriority ensures that no value is present for UsePriority, not even an explicit nil
### GetAsyncResponseTimeout

`func (o *FetchExternalTasksDto) GetAsyncResponseTimeout() int64`

GetAsyncResponseTimeout returns the AsyncResponseTimeout field if non-nil, zero value otherwise.

### GetAsyncResponseTimeoutOk

`func (o *FetchExternalTasksDto) GetAsyncResponseTimeoutOk() (*int64, bool)`

GetAsyncResponseTimeoutOk returns a tuple with the AsyncResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncResponseTimeout

`func (o *FetchExternalTasksDto) SetAsyncResponseTimeout(v int64)`

SetAsyncResponseTimeout sets AsyncResponseTimeout field to given value.

### HasAsyncResponseTimeout

`func (o *FetchExternalTasksDto) HasAsyncResponseTimeout() bool`

HasAsyncResponseTimeout returns a boolean if a field has been set.

### SetAsyncResponseTimeoutNil

`func (o *FetchExternalTasksDto) SetAsyncResponseTimeoutNil(b bool)`

 SetAsyncResponseTimeoutNil sets the value for AsyncResponseTimeout to be an explicit nil

### UnsetAsyncResponseTimeout
`func (o *FetchExternalTasksDto) UnsetAsyncResponseTimeout()`

UnsetAsyncResponseTimeout ensures that no value is present for AsyncResponseTimeout, not even an explicit nil
### GetTopics

`func (o *FetchExternalTasksDto) GetTopics() []FetchExternalTaskTopicDto`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *FetchExternalTasksDto) GetTopicsOk() (*[]FetchExternalTaskTopicDto, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *FetchExternalTasksDto) SetTopics(v []FetchExternalTaskTopicDto)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *FetchExternalTasksDto) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### SetTopicsNil

`func (o *FetchExternalTasksDto) SetTopicsNil(b bool)`

 SetTopicsNil sets the value for Topics to be an explicit nil

### UnsetTopics
`func (o *FetchExternalTasksDto) UnsetTopics()`

UnsetTopics ensures that no value is present for Topics, not even an explicit nil
### GetSorting

`func (o *FetchExternalTasksDto) GetSorting() []FetchExternalTasksDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *FetchExternalTasksDto) GetSortingOk() (*[]FetchExternalTasksDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *FetchExternalTasksDto) SetSorting(v []FetchExternalTasksDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *FetchExternalTasksDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *FetchExternalTasksDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *FetchExternalTasksDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


