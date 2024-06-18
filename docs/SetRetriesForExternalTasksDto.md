# SetRetriesForExternalTasksDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retries** | Pointer to **NullableInt32** | The number of retries to set for the external task.  Must be &gt;&#x3D; 0. If this is 0, an incident is created and the task cannot be fetched anymore unless the retries are increased again. Can not be null. | [optional] 
**ExternalTaskIds** | Pointer to **[]string** | The ids of the external tasks to set the number of retries for. | [optional] 
**ProcessInstanceIds** | Pointer to **[]string** | The ids of process instances containing the tasks to set the number of retries for. | [optional] 
**ExternalTaskQuery** | Pointer to [**ExternalTaskQueryDto**](ExternalTaskQueryDto.md) |  | [optional] 
**ProcessInstanceQuery** | Pointer to [**ProcessInstanceQueryDto**](ProcessInstanceQueryDto.md) |  | [optional] 
**HistoricProcessInstanceQuery** | Pointer to [**HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) |  | [optional] 

## Methods

### NewSetRetriesForExternalTasksDto

`func NewSetRetriesForExternalTasksDto() *SetRetriesForExternalTasksDto`

NewSetRetriesForExternalTasksDto instantiates a new SetRetriesForExternalTasksDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRetriesForExternalTasksDtoWithDefaults

`func NewSetRetriesForExternalTasksDtoWithDefaults() *SetRetriesForExternalTasksDto`

NewSetRetriesForExternalTasksDtoWithDefaults instantiates a new SetRetriesForExternalTasksDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetries

`func (o *SetRetriesForExternalTasksDto) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *SetRetriesForExternalTasksDto) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *SetRetriesForExternalTasksDto) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *SetRetriesForExternalTasksDto) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *SetRetriesForExternalTasksDto) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *SetRetriesForExternalTasksDto) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetExternalTaskIds

`func (o *SetRetriesForExternalTasksDto) GetExternalTaskIds() []string`

GetExternalTaskIds returns the ExternalTaskIds field if non-nil, zero value otherwise.

### GetExternalTaskIdsOk

`func (o *SetRetriesForExternalTasksDto) GetExternalTaskIdsOk() (*[]string, bool)`

GetExternalTaskIdsOk returns a tuple with the ExternalTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTaskIds

`func (o *SetRetriesForExternalTasksDto) SetExternalTaskIds(v []string)`

SetExternalTaskIds sets ExternalTaskIds field to given value.

### HasExternalTaskIds

`func (o *SetRetriesForExternalTasksDto) HasExternalTaskIds() bool`

HasExternalTaskIds returns a boolean if a field has been set.

### SetExternalTaskIdsNil

`func (o *SetRetriesForExternalTasksDto) SetExternalTaskIdsNil(b bool)`

 SetExternalTaskIdsNil sets the value for ExternalTaskIds to be an explicit nil

### UnsetExternalTaskIds
`func (o *SetRetriesForExternalTasksDto) UnsetExternalTaskIds()`

UnsetExternalTaskIds ensures that no value is present for ExternalTaskIds, not even an explicit nil
### GetProcessInstanceIds

`func (o *SetRetriesForExternalTasksDto) GetProcessInstanceIds() []string`

GetProcessInstanceIds returns the ProcessInstanceIds field if non-nil, zero value otherwise.

### GetProcessInstanceIdsOk

`func (o *SetRetriesForExternalTasksDto) GetProcessInstanceIdsOk() (*[]string, bool)`

GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIds

`func (o *SetRetriesForExternalTasksDto) SetProcessInstanceIds(v []string)`

SetProcessInstanceIds sets ProcessInstanceIds field to given value.

### HasProcessInstanceIds

`func (o *SetRetriesForExternalTasksDto) HasProcessInstanceIds() bool`

HasProcessInstanceIds returns a boolean if a field has been set.

### SetProcessInstanceIdsNil

`func (o *SetRetriesForExternalTasksDto) SetProcessInstanceIdsNil(b bool)`

 SetProcessInstanceIdsNil sets the value for ProcessInstanceIds to be an explicit nil

### UnsetProcessInstanceIds
`func (o *SetRetriesForExternalTasksDto) UnsetProcessInstanceIds()`

UnsetProcessInstanceIds ensures that no value is present for ProcessInstanceIds, not even an explicit nil
### GetExternalTaskQuery

`func (o *SetRetriesForExternalTasksDto) GetExternalTaskQuery() ExternalTaskQueryDto`

GetExternalTaskQuery returns the ExternalTaskQuery field if non-nil, zero value otherwise.

### GetExternalTaskQueryOk

`func (o *SetRetriesForExternalTasksDto) GetExternalTaskQueryOk() (*ExternalTaskQueryDto, bool)`

GetExternalTaskQueryOk returns a tuple with the ExternalTaskQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTaskQuery

`func (o *SetRetriesForExternalTasksDto) SetExternalTaskQuery(v ExternalTaskQueryDto)`

SetExternalTaskQuery sets ExternalTaskQuery field to given value.

### HasExternalTaskQuery

`func (o *SetRetriesForExternalTasksDto) HasExternalTaskQuery() bool`

HasExternalTaskQuery returns a boolean if a field has been set.

### GetProcessInstanceQuery

`func (o *SetRetriesForExternalTasksDto) GetProcessInstanceQuery() ProcessInstanceQueryDto`

GetProcessInstanceQuery returns the ProcessInstanceQuery field if non-nil, zero value otherwise.

### GetProcessInstanceQueryOk

`func (o *SetRetriesForExternalTasksDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool)`

GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceQuery

`func (o *SetRetriesForExternalTasksDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto)`

SetProcessInstanceQuery sets ProcessInstanceQuery field to given value.

### HasProcessInstanceQuery

`func (o *SetRetriesForExternalTasksDto) HasProcessInstanceQuery() bool`

HasProcessInstanceQuery returns a boolean if a field has been set.

### GetHistoricProcessInstanceQuery

`func (o *SetRetriesForExternalTasksDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto`

GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field if non-nil, zero value otherwise.

### GetHistoricProcessInstanceQueryOk

`func (o *SetRetriesForExternalTasksDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool)`

GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricProcessInstanceQuery

`func (o *SetRetriesForExternalTasksDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto)`

SetHistoricProcessInstanceQuery sets HistoricProcessInstanceQuery field to given value.

### HasHistoricProcessInstanceQuery

`func (o *SetRetriesForExternalTasksDto) HasHistoricProcessInstanceQuery() bool`

HasHistoricProcessInstanceQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


