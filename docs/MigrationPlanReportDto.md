# MigrationPlanReportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstructionReports** | Pointer to [**[]MigrationInstructionValidationReportDto**](MigrationInstructionValidationReportDto.md) | The list of instruction validation reports. If no validation errors are detected it is an empty list. | [optional] 
**VariableReports** | Pointer to [**map[string]MigrationVariableValidationReportDto**](MigrationVariableValidationReportDto.md) | A map of variable reports. Each key is a variable name and each value a JSON object consisting of the variable&#39;s type, value, value info object and a list of failures. | [optional] 

## Methods

### NewMigrationPlanReportDto

`func NewMigrationPlanReportDto() *MigrationPlanReportDto`

NewMigrationPlanReportDto instantiates a new MigrationPlanReportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationPlanReportDtoWithDefaults

`func NewMigrationPlanReportDtoWithDefaults() *MigrationPlanReportDto`

NewMigrationPlanReportDtoWithDefaults instantiates a new MigrationPlanReportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructionReports

`func (o *MigrationPlanReportDto) GetInstructionReports() []MigrationInstructionValidationReportDto`

GetInstructionReports returns the InstructionReports field if non-nil, zero value otherwise.

### GetInstructionReportsOk

`func (o *MigrationPlanReportDto) GetInstructionReportsOk() (*[]MigrationInstructionValidationReportDto, bool)`

GetInstructionReportsOk returns a tuple with the InstructionReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructionReports

`func (o *MigrationPlanReportDto) SetInstructionReports(v []MigrationInstructionValidationReportDto)`

SetInstructionReports sets InstructionReports field to given value.

### HasInstructionReports

`func (o *MigrationPlanReportDto) HasInstructionReports() bool`

HasInstructionReports returns a boolean if a field has been set.

### SetInstructionReportsNil

`func (o *MigrationPlanReportDto) SetInstructionReportsNil(b bool)`

 SetInstructionReportsNil sets the value for InstructionReports to be an explicit nil

### UnsetInstructionReports
`func (o *MigrationPlanReportDto) UnsetInstructionReports()`

UnsetInstructionReports ensures that no value is present for InstructionReports, not even an explicit nil
### GetVariableReports

`func (o *MigrationPlanReportDto) GetVariableReports() map[string]MigrationVariableValidationReportDto`

GetVariableReports returns the VariableReports field if non-nil, zero value otherwise.

### GetVariableReportsOk

`func (o *MigrationPlanReportDto) GetVariableReportsOk() (*map[string]MigrationVariableValidationReportDto, bool)`

GetVariableReportsOk returns a tuple with the VariableReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableReports

`func (o *MigrationPlanReportDto) SetVariableReports(v map[string]MigrationVariableValidationReportDto)`

SetVariableReports sets VariableReports field to given value.

### HasVariableReports

`func (o *MigrationPlanReportDto) HasVariableReports() bool`

HasVariableReports returns a boolean if a field has been set.

### SetVariableReportsNil

`func (o *MigrationPlanReportDto) SetVariableReportsNil(b bool)`

 SetVariableReportsNil sets the value for VariableReports to be an explicit nil

### UnsetVariableReports
`func (o *MigrationPlanReportDto) UnsetVariableReports()`

UnsetVariableReports ensures that no value is present for VariableReports, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


