# MigrationPlanDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceProcessDefinitionId** | Pointer to **NullableString** | The id of the source process definition for the migration. | [optional] 
**TargetProcessDefinitionId** | Pointer to **NullableString** | The id of the target process definition for the migration. | [optional] 
**Instructions** | Pointer to [**[]MigrationInstructionDto**](MigrationInstructionDto.md) | A list of migration instructions which map equal activities. Each migration instruction is a JSON object with the following properties: | [optional] 
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A map of variables which will be set into the process instances&#39; scope. Each key is a variable name and each value a JSON variable value object. | [optional] 

## Methods

### NewMigrationPlanDto

`func NewMigrationPlanDto() *MigrationPlanDto`

NewMigrationPlanDto instantiates a new MigrationPlanDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationPlanDtoWithDefaults

`func NewMigrationPlanDtoWithDefaults() *MigrationPlanDto`

NewMigrationPlanDtoWithDefaults instantiates a new MigrationPlanDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceProcessDefinitionId

`func (o *MigrationPlanDto) GetSourceProcessDefinitionId() string`

GetSourceProcessDefinitionId returns the SourceProcessDefinitionId field if non-nil, zero value otherwise.

### GetSourceProcessDefinitionIdOk

`func (o *MigrationPlanDto) GetSourceProcessDefinitionIdOk() (*string, bool)`

GetSourceProcessDefinitionIdOk returns a tuple with the SourceProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProcessDefinitionId

`func (o *MigrationPlanDto) SetSourceProcessDefinitionId(v string)`

SetSourceProcessDefinitionId sets SourceProcessDefinitionId field to given value.

### HasSourceProcessDefinitionId

`func (o *MigrationPlanDto) HasSourceProcessDefinitionId() bool`

HasSourceProcessDefinitionId returns a boolean if a field has been set.

### SetSourceProcessDefinitionIdNil

`func (o *MigrationPlanDto) SetSourceProcessDefinitionIdNil(b bool)`

 SetSourceProcessDefinitionIdNil sets the value for SourceProcessDefinitionId to be an explicit nil

### UnsetSourceProcessDefinitionId
`func (o *MigrationPlanDto) UnsetSourceProcessDefinitionId()`

UnsetSourceProcessDefinitionId ensures that no value is present for SourceProcessDefinitionId, not even an explicit nil
### GetTargetProcessDefinitionId

`func (o *MigrationPlanDto) GetTargetProcessDefinitionId() string`

GetTargetProcessDefinitionId returns the TargetProcessDefinitionId field if non-nil, zero value otherwise.

### GetTargetProcessDefinitionIdOk

`func (o *MigrationPlanDto) GetTargetProcessDefinitionIdOk() (*string, bool)`

GetTargetProcessDefinitionIdOk returns a tuple with the TargetProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProcessDefinitionId

`func (o *MigrationPlanDto) SetTargetProcessDefinitionId(v string)`

SetTargetProcessDefinitionId sets TargetProcessDefinitionId field to given value.

### HasTargetProcessDefinitionId

`func (o *MigrationPlanDto) HasTargetProcessDefinitionId() bool`

HasTargetProcessDefinitionId returns a boolean if a field has been set.

### SetTargetProcessDefinitionIdNil

`func (o *MigrationPlanDto) SetTargetProcessDefinitionIdNil(b bool)`

 SetTargetProcessDefinitionIdNil sets the value for TargetProcessDefinitionId to be an explicit nil

### UnsetTargetProcessDefinitionId
`func (o *MigrationPlanDto) UnsetTargetProcessDefinitionId()`

UnsetTargetProcessDefinitionId ensures that no value is present for TargetProcessDefinitionId, not even an explicit nil
### GetInstructions

`func (o *MigrationPlanDto) GetInstructions() []MigrationInstructionDto`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *MigrationPlanDto) GetInstructionsOk() (*[]MigrationInstructionDto, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *MigrationPlanDto) SetInstructions(v []MigrationInstructionDto)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *MigrationPlanDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *MigrationPlanDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *MigrationPlanDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetVariables

`func (o *MigrationPlanDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *MigrationPlanDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *MigrationPlanDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *MigrationPlanDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *MigrationPlanDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *MigrationPlanDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


