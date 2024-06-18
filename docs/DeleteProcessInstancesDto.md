# DeleteProcessInstancesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessInstanceIds** | Pointer to **[]string** | A list process instance ids to delete. | [optional] 
**DeleteReason** | Pointer to **NullableString** | A string with delete reason. | [optional] 
**SkipCustomListeners** | Pointer to **NullableBool** | Skip execution listener invocation for activities that are started or ended as part of this request. | [optional] 
**SkipSubprocesses** | Pointer to **NullableBool** | Skip deletion of the subprocesses related to deleted processes as part of this request. | [optional] 
**SkipIoMappings** | Pointer to **NullableBool** | Skip execution of [input/output variable mappings](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#input-output-variable-mapping) for activities that are started or ended as part of this request. | [optional] 
**ProcessInstanceQuery** | Pointer to [**ProcessInstanceQueryDto**](ProcessInstanceQueryDto.md) |  | [optional] 
**HistoricProcessInstanceQuery** | Pointer to [**HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) |  | [optional] 

## Methods

### NewDeleteProcessInstancesDto

`func NewDeleteProcessInstancesDto() *DeleteProcessInstancesDto`

NewDeleteProcessInstancesDto instantiates a new DeleteProcessInstancesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteProcessInstancesDtoWithDefaults

`func NewDeleteProcessInstancesDtoWithDefaults() *DeleteProcessInstancesDto`

NewDeleteProcessInstancesDtoWithDefaults instantiates a new DeleteProcessInstancesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessInstanceIds

`func (o *DeleteProcessInstancesDto) GetProcessInstanceIds() []string`

GetProcessInstanceIds returns the ProcessInstanceIds field if non-nil, zero value otherwise.

### GetProcessInstanceIdsOk

`func (o *DeleteProcessInstancesDto) GetProcessInstanceIdsOk() (*[]string, bool)`

GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIds

`func (o *DeleteProcessInstancesDto) SetProcessInstanceIds(v []string)`

SetProcessInstanceIds sets ProcessInstanceIds field to given value.

### HasProcessInstanceIds

`func (o *DeleteProcessInstancesDto) HasProcessInstanceIds() bool`

HasProcessInstanceIds returns a boolean if a field has been set.

### SetProcessInstanceIdsNil

`func (o *DeleteProcessInstancesDto) SetProcessInstanceIdsNil(b bool)`

 SetProcessInstanceIdsNil sets the value for ProcessInstanceIds to be an explicit nil

### UnsetProcessInstanceIds
`func (o *DeleteProcessInstancesDto) UnsetProcessInstanceIds()`

UnsetProcessInstanceIds ensures that no value is present for ProcessInstanceIds, not even an explicit nil
### GetDeleteReason

`func (o *DeleteProcessInstancesDto) GetDeleteReason() string`

GetDeleteReason returns the DeleteReason field if non-nil, zero value otherwise.

### GetDeleteReasonOk

`func (o *DeleteProcessInstancesDto) GetDeleteReasonOk() (*string, bool)`

GetDeleteReasonOk returns a tuple with the DeleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteReason

`func (o *DeleteProcessInstancesDto) SetDeleteReason(v string)`

SetDeleteReason sets DeleteReason field to given value.

### HasDeleteReason

`func (o *DeleteProcessInstancesDto) HasDeleteReason() bool`

HasDeleteReason returns a boolean if a field has been set.

### SetDeleteReasonNil

`func (o *DeleteProcessInstancesDto) SetDeleteReasonNil(b bool)`

 SetDeleteReasonNil sets the value for DeleteReason to be an explicit nil

### UnsetDeleteReason
`func (o *DeleteProcessInstancesDto) UnsetDeleteReason()`

UnsetDeleteReason ensures that no value is present for DeleteReason, not even an explicit nil
### GetSkipCustomListeners

`func (o *DeleteProcessInstancesDto) GetSkipCustomListeners() bool`

GetSkipCustomListeners returns the SkipCustomListeners field if non-nil, zero value otherwise.

### GetSkipCustomListenersOk

`func (o *DeleteProcessInstancesDto) GetSkipCustomListenersOk() (*bool, bool)`

GetSkipCustomListenersOk returns a tuple with the SkipCustomListeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCustomListeners

`func (o *DeleteProcessInstancesDto) SetSkipCustomListeners(v bool)`

SetSkipCustomListeners sets SkipCustomListeners field to given value.

### HasSkipCustomListeners

`func (o *DeleteProcessInstancesDto) HasSkipCustomListeners() bool`

HasSkipCustomListeners returns a boolean if a field has been set.

### SetSkipCustomListenersNil

`func (o *DeleteProcessInstancesDto) SetSkipCustomListenersNil(b bool)`

 SetSkipCustomListenersNil sets the value for SkipCustomListeners to be an explicit nil

### UnsetSkipCustomListeners
`func (o *DeleteProcessInstancesDto) UnsetSkipCustomListeners()`

UnsetSkipCustomListeners ensures that no value is present for SkipCustomListeners, not even an explicit nil
### GetSkipSubprocesses

`func (o *DeleteProcessInstancesDto) GetSkipSubprocesses() bool`

GetSkipSubprocesses returns the SkipSubprocesses field if non-nil, zero value otherwise.

### GetSkipSubprocessesOk

`func (o *DeleteProcessInstancesDto) GetSkipSubprocessesOk() (*bool, bool)`

GetSkipSubprocessesOk returns a tuple with the SkipSubprocesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubprocesses

`func (o *DeleteProcessInstancesDto) SetSkipSubprocesses(v bool)`

SetSkipSubprocesses sets SkipSubprocesses field to given value.

### HasSkipSubprocesses

`func (o *DeleteProcessInstancesDto) HasSkipSubprocesses() bool`

HasSkipSubprocesses returns a boolean if a field has been set.

### SetSkipSubprocessesNil

`func (o *DeleteProcessInstancesDto) SetSkipSubprocessesNil(b bool)`

 SetSkipSubprocessesNil sets the value for SkipSubprocesses to be an explicit nil

### UnsetSkipSubprocesses
`func (o *DeleteProcessInstancesDto) UnsetSkipSubprocesses()`

UnsetSkipSubprocesses ensures that no value is present for SkipSubprocesses, not even an explicit nil
### GetSkipIoMappings

`func (o *DeleteProcessInstancesDto) GetSkipIoMappings() bool`

GetSkipIoMappings returns the SkipIoMappings field if non-nil, zero value otherwise.

### GetSkipIoMappingsOk

`func (o *DeleteProcessInstancesDto) GetSkipIoMappingsOk() (*bool, bool)`

GetSkipIoMappingsOk returns a tuple with the SkipIoMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIoMappings

`func (o *DeleteProcessInstancesDto) SetSkipIoMappings(v bool)`

SetSkipIoMappings sets SkipIoMappings field to given value.

### HasSkipIoMappings

`func (o *DeleteProcessInstancesDto) HasSkipIoMappings() bool`

HasSkipIoMappings returns a boolean if a field has been set.

### SetSkipIoMappingsNil

`func (o *DeleteProcessInstancesDto) SetSkipIoMappingsNil(b bool)`

 SetSkipIoMappingsNil sets the value for SkipIoMappings to be an explicit nil

### UnsetSkipIoMappings
`func (o *DeleteProcessInstancesDto) UnsetSkipIoMappings()`

UnsetSkipIoMappings ensures that no value is present for SkipIoMappings, not even an explicit nil
### GetProcessInstanceQuery

`func (o *DeleteProcessInstancesDto) GetProcessInstanceQuery() ProcessInstanceQueryDto`

GetProcessInstanceQuery returns the ProcessInstanceQuery field if non-nil, zero value otherwise.

### GetProcessInstanceQueryOk

`func (o *DeleteProcessInstancesDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool)`

GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceQuery

`func (o *DeleteProcessInstancesDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto)`

SetProcessInstanceQuery sets ProcessInstanceQuery field to given value.

### HasProcessInstanceQuery

`func (o *DeleteProcessInstancesDto) HasProcessInstanceQuery() bool`

HasProcessInstanceQuery returns a boolean if a field has been set.

### GetHistoricProcessInstanceQuery

`func (o *DeleteProcessInstancesDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto`

GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field if non-nil, zero value otherwise.

### GetHistoricProcessInstanceQueryOk

`func (o *DeleteProcessInstancesDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool)`

GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricProcessInstanceQuery

`func (o *DeleteProcessInstancesDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto)`

SetHistoricProcessInstanceQuery sets HistoricProcessInstanceQuery field to given value.

### HasHistoricProcessInstanceQuery

`func (o *DeleteProcessInstancesDto) HasHistoricProcessInstanceQuery() bool`

HasHistoricProcessInstanceQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


