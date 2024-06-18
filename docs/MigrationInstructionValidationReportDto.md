# MigrationInstructionValidationReportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instruction** | Pointer to [**MigrationInstructionDto**](MigrationInstructionDto.md) |  | [optional] 
**Failures** | Pointer to **[]string** | A list of instruction validation report messages. | [optional] 

## Methods

### NewMigrationInstructionValidationReportDto

`func NewMigrationInstructionValidationReportDto() *MigrationInstructionValidationReportDto`

NewMigrationInstructionValidationReportDto instantiates a new MigrationInstructionValidationReportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationInstructionValidationReportDtoWithDefaults

`func NewMigrationInstructionValidationReportDtoWithDefaults() *MigrationInstructionValidationReportDto`

NewMigrationInstructionValidationReportDtoWithDefaults instantiates a new MigrationInstructionValidationReportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstruction

`func (o *MigrationInstructionValidationReportDto) GetInstruction() MigrationInstructionDto`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *MigrationInstructionValidationReportDto) GetInstructionOk() (*MigrationInstructionDto, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *MigrationInstructionValidationReportDto) SetInstruction(v MigrationInstructionDto)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *MigrationInstructionValidationReportDto) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### GetFailures

`func (o *MigrationInstructionValidationReportDto) GetFailures() []string`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *MigrationInstructionValidationReportDto) GetFailuresOk() (*[]string, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *MigrationInstructionValidationReportDto) SetFailures(v []string)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *MigrationInstructionValidationReportDto) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### SetFailuresNil

`func (o *MigrationInstructionValidationReportDto) SetFailuresNil(b bool)`

 SetFailuresNil sets the value for Failures to be an explicit nil

### UnsetFailures
`func (o *MigrationInstructionValidationReportDto) UnsetFailures()`

UnsetFailures ensures that no value is present for Failures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


