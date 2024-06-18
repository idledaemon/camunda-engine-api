# MigrationExecutionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MigrationPlan** | Pointer to [**MigrationPlanDto**](MigrationPlanDto.md) |  | [optional] 
**ProcessInstanceIds** | Pointer to **[]string** | A list of process instance ids to migrate. | [optional] 
**ProcessInstanceQuery** | Pointer to [**ProcessInstanceQueryDto**](ProcessInstanceQueryDto.md) |  | [optional] 
**SkipCustomListeners** | Pointer to **NullableBool** | A boolean value to control whether execution listeners should be invoked during migration. | [optional] 
**SkipIoMappings** | Pointer to **NullableBool** | A boolean value to control whether input/output mappings should be executed during migration. | [optional] 

## Methods

### NewMigrationExecutionDto

`func NewMigrationExecutionDto() *MigrationExecutionDto`

NewMigrationExecutionDto instantiates a new MigrationExecutionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationExecutionDtoWithDefaults

`func NewMigrationExecutionDtoWithDefaults() *MigrationExecutionDto`

NewMigrationExecutionDtoWithDefaults instantiates a new MigrationExecutionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrationPlan

`func (o *MigrationExecutionDto) GetMigrationPlan() MigrationPlanDto`

GetMigrationPlan returns the MigrationPlan field if non-nil, zero value otherwise.

### GetMigrationPlanOk

`func (o *MigrationExecutionDto) GetMigrationPlanOk() (*MigrationPlanDto, bool)`

GetMigrationPlanOk returns a tuple with the MigrationPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationPlan

`func (o *MigrationExecutionDto) SetMigrationPlan(v MigrationPlanDto)`

SetMigrationPlan sets MigrationPlan field to given value.

### HasMigrationPlan

`func (o *MigrationExecutionDto) HasMigrationPlan() bool`

HasMigrationPlan returns a boolean if a field has been set.

### GetProcessInstanceIds

`func (o *MigrationExecutionDto) GetProcessInstanceIds() []string`

GetProcessInstanceIds returns the ProcessInstanceIds field if non-nil, zero value otherwise.

### GetProcessInstanceIdsOk

`func (o *MigrationExecutionDto) GetProcessInstanceIdsOk() (*[]string, bool)`

GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIds

`func (o *MigrationExecutionDto) SetProcessInstanceIds(v []string)`

SetProcessInstanceIds sets ProcessInstanceIds field to given value.

### HasProcessInstanceIds

`func (o *MigrationExecutionDto) HasProcessInstanceIds() bool`

HasProcessInstanceIds returns a boolean if a field has been set.

### SetProcessInstanceIdsNil

`func (o *MigrationExecutionDto) SetProcessInstanceIdsNil(b bool)`

 SetProcessInstanceIdsNil sets the value for ProcessInstanceIds to be an explicit nil

### UnsetProcessInstanceIds
`func (o *MigrationExecutionDto) UnsetProcessInstanceIds()`

UnsetProcessInstanceIds ensures that no value is present for ProcessInstanceIds, not even an explicit nil
### GetProcessInstanceQuery

`func (o *MigrationExecutionDto) GetProcessInstanceQuery() ProcessInstanceQueryDto`

GetProcessInstanceQuery returns the ProcessInstanceQuery field if non-nil, zero value otherwise.

### GetProcessInstanceQueryOk

`func (o *MigrationExecutionDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool)`

GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceQuery

`func (o *MigrationExecutionDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto)`

SetProcessInstanceQuery sets ProcessInstanceQuery field to given value.

### HasProcessInstanceQuery

`func (o *MigrationExecutionDto) HasProcessInstanceQuery() bool`

HasProcessInstanceQuery returns a boolean if a field has been set.

### GetSkipCustomListeners

`func (o *MigrationExecutionDto) GetSkipCustomListeners() bool`

GetSkipCustomListeners returns the SkipCustomListeners field if non-nil, zero value otherwise.

### GetSkipCustomListenersOk

`func (o *MigrationExecutionDto) GetSkipCustomListenersOk() (*bool, bool)`

GetSkipCustomListenersOk returns a tuple with the SkipCustomListeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCustomListeners

`func (o *MigrationExecutionDto) SetSkipCustomListeners(v bool)`

SetSkipCustomListeners sets SkipCustomListeners field to given value.

### HasSkipCustomListeners

`func (o *MigrationExecutionDto) HasSkipCustomListeners() bool`

HasSkipCustomListeners returns a boolean if a field has been set.

### SetSkipCustomListenersNil

`func (o *MigrationExecutionDto) SetSkipCustomListenersNil(b bool)`

 SetSkipCustomListenersNil sets the value for SkipCustomListeners to be an explicit nil

### UnsetSkipCustomListeners
`func (o *MigrationExecutionDto) UnsetSkipCustomListeners()`

UnsetSkipCustomListeners ensures that no value is present for SkipCustomListeners, not even an explicit nil
### GetSkipIoMappings

`func (o *MigrationExecutionDto) GetSkipIoMappings() bool`

GetSkipIoMappings returns the SkipIoMappings field if non-nil, zero value otherwise.

### GetSkipIoMappingsOk

`func (o *MigrationExecutionDto) GetSkipIoMappingsOk() (*bool, bool)`

GetSkipIoMappingsOk returns a tuple with the SkipIoMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIoMappings

`func (o *MigrationExecutionDto) SetSkipIoMappings(v bool)`

SetSkipIoMappings sets SkipIoMappings field to given value.

### HasSkipIoMappings

`func (o *MigrationExecutionDto) HasSkipIoMappings() bool`

HasSkipIoMappings returns a boolean if a field has been set.

### SetSkipIoMappingsNil

`func (o *MigrationExecutionDto) SetSkipIoMappingsNil(b bool)`

 SetSkipIoMappingsNil sets the value for SkipIoMappings to be an explicit nil

### UnsetSkipIoMappings
`func (o *MigrationExecutionDto) UnsetSkipIoMappings()`

UnsetSkipIoMappings ensures that no value is present for SkipIoMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


