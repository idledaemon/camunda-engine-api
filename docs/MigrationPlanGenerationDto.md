# MigrationPlanGenerationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceProcessDefinitionId** | Pointer to **NullableString** | The id of the source process definition for the migration. | [optional] 
**TargetProcessDefinitionId** | Pointer to **NullableString** | The id of the target process definition for the migration. | [optional] 
**UpdateEventTriggers** | Pointer to **NullableBool** | A boolean flag indicating whether instructions between events should be configured to update the event triggers. | [optional] 
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A map of variables which will be set into the process instances&#39; scope. Each key is a variable name and each value a JSON variable value object. | [optional] 

## Methods

### NewMigrationPlanGenerationDto

`func NewMigrationPlanGenerationDto() *MigrationPlanGenerationDto`

NewMigrationPlanGenerationDto instantiates a new MigrationPlanGenerationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationPlanGenerationDtoWithDefaults

`func NewMigrationPlanGenerationDtoWithDefaults() *MigrationPlanGenerationDto`

NewMigrationPlanGenerationDtoWithDefaults instantiates a new MigrationPlanGenerationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceProcessDefinitionId

`func (o *MigrationPlanGenerationDto) GetSourceProcessDefinitionId() string`

GetSourceProcessDefinitionId returns the SourceProcessDefinitionId field if non-nil, zero value otherwise.

### GetSourceProcessDefinitionIdOk

`func (o *MigrationPlanGenerationDto) GetSourceProcessDefinitionIdOk() (*string, bool)`

GetSourceProcessDefinitionIdOk returns a tuple with the SourceProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProcessDefinitionId

`func (o *MigrationPlanGenerationDto) SetSourceProcessDefinitionId(v string)`

SetSourceProcessDefinitionId sets SourceProcessDefinitionId field to given value.

### HasSourceProcessDefinitionId

`func (o *MigrationPlanGenerationDto) HasSourceProcessDefinitionId() bool`

HasSourceProcessDefinitionId returns a boolean if a field has been set.

### SetSourceProcessDefinitionIdNil

`func (o *MigrationPlanGenerationDto) SetSourceProcessDefinitionIdNil(b bool)`

 SetSourceProcessDefinitionIdNil sets the value for SourceProcessDefinitionId to be an explicit nil

### UnsetSourceProcessDefinitionId
`func (o *MigrationPlanGenerationDto) UnsetSourceProcessDefinitionId()`

UnsetSourceProcessDefinitionId ensures that no value is present for SourceProcessDefinitionId, not even an explicit nil
### GetTargetProcessDefinitionId

`func (o *MigrationPlanGenerationDto) GetTargetProcessDefinitionId() string`

GetTargetProcessDefinitionId returns the TargetProcessDefinitionId field if non-nil, zero value otherwise.

### GetTargetProcessDefinitionIdOk

`func (o *MigrationPlanGenerationDto) GetTargetProcessDefinitionIdOk() (*string, bool)`

GetTargetProcessDefinitionIdOk returns a tuple with the TargetProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProcessDefinitionId

`func (o *MigrationPlanGenerationDto) SetTargetProcessDefinitionId(v string)`

SetTargetProcessDefinitionId sets TargetProcessDefinitionId field to given value.

### HasTargetProcessDefinitionId

`func (o *MigrationPlanGenerationDto) HasTargetProcessDefinitionId() bool`

HasTargetProcessDefinitionId returns a boolean if a field has been set.

### SetTargetProcessDefinitionIdNil

`func (o *MigrationPlanGenerationDto) SetTargetProcessDefinitionIdNil(b bool)`

 SetTargetProcessDefinitionIdNil sets the value for TargetProcessDefinitionId to be an explicit nil

### UnsetTargetProcessDefinitionId
`func (o *MigrationPlanGenerationDto) UnsetTargetProcessDefinitionId()`

UnsetTargetProcessDefinitionId ensures that no value is present for TargetProcessDefinitionId, not even an explicit nil
### GetUpdateEventTriggers

`func (o *MigrationPlanGenerationDto) GetUpdateEventTriggers() bool`

GetUpdateEventTriggers returns the UpdateEventTriggers field if non-nil, zero value otherwise.

### GetUpdateEventTriggersOk

`func (o *MigrationPlanGenerationDto) GetUpdateEventTriggersOk() (*bool, bool)`

GetUpdateEventTriggersOk returns a tuple with the UpdateEventTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateEventTriggers

`func (o *MigrationPlanGenerationDto) SetUpdateEventTriggers(v bool)`

SetUpdateEventTriggers sets UpdateEventTriggers field to given value.

### HasUpdateEventTriggers

`func (o *MigrationPlanGenerationDto) HasUpdateEventTriggers() bool`

HasUpdateEventTriggers returns a boolean if a field has been set.

### SetUpdateEventTriggersNil

`func (o *MigrationPlanGenerationDto) SetUpdateEventTriggersNil(b bool)`

 SetUpdateEventTriggersNil sets the value for UpdateEventTriggers to be an explicit nil

### UnsetUpdateEventTriggers
`func (o *MigrationPlanGenerationDto) UnsetUpdateEventTriggers()`

UnsetUpdateEventTriggers ensures that no value is present for UpdateEventTriggers, not even an explicit nil
### GetVariables

`func (o *MigrationPlanGenerationDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *MigrationPlanGenerationDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *MigrationPlanGenerationDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *MigrationPlanGenerationDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *MigrationPlanGenerationDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *MigrationPlanGenerationDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


