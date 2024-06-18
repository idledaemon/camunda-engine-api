# TaskCountByCandidateGroupResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **NullableString** | The name of the candidate group. If there are tasks without a group name, the value will be &#x60;null&#x60; | [optional] 
**TaskCount** | Pointer to **NullableInt32** | The number of tasks which have the group name as candidate group. | [optional] 

## Methods

### NewTaskCountByCandidateGroupResultDto

`func NewTaskCountByCandidateGroupResultDto() *TaskCountByCandidateGroupResultDto`

NewTaskCountByCandidateGroupResultDto instantiates a new TaskCountByCandidateGroupResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCountByCandidateGroupResultDtoWithDefaults

`func NewTaskCountByCandidateGroupResultDtoWithDefaults() *TaskCountByCandidateGroupResultDto`

NewTaskCountByCandidateGroupResultDtoWithDefaults instantiates a new TaskCountByCandidateGroupResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *TaskCountByCandidateGroupResultDto) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *TaskCountByCandidateGroupResultDto) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *TaskCountByCandidateGroupResultDto) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *TaskCountByCandidateGroupResultDto) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *TaskCountByCandidateGroupResultDto) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *TaskCountByCandidateGroupResultDto) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetTaskCount

`func (o *TaskCountByCandidateGroupResultDto) GetTaskCount() int32`

GetTaskCount returns the TaskCount field if non-nil, zero value otherwise.

### GetTaskCountOk

`func (o *TaskCountByCandidateGroupResultDto) GetTaskCountOk() (*int32, bool)`

GetTaskCountOk returns a tuple with the TaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCount

`func (o *TaskCountByCandidateGroupResultDto) SetTaskCount(v int32)`

SetTaskCount sets TaskCount field to given value.

### HasTaskCount

`func (o *TaskCountByCandidateGroupResultDto) HasTaskCount() bool`

HasTaskCount returns a boolean if a field has been set.

### SetTaskCountNil

`func (o *TaskCountByCandidateGroupResultDto) SetTaskCountNil(b bool)`

 SetTaskCountNil sets the value for TaskCount to be an explicit nil

### UnsetTaskCount
`func (o *TaskCountByCandidateGroupResultDto) UnsetTaskCount()`

UnsetTaskCount ensures that no value is present for TaskCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


