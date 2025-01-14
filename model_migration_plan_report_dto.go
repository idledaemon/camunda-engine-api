/*
Camunda Platform REST API

OpenApi Spec for Camunda Platform REST API.

API version: 7.21.2-ee
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package camundarestgo

import (
	"encoding/json"
)

// checks if the MigrationPlanReportDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrationPlanReportDto{}

// MigrationPlanReportDto struct for MigrationPlanReportDto
type MigrationPlanReportDto struct {
	// The list of instruction validation reports. If no validation errors are detected it is an empty list.
	InstructionReports []MigrationInstructionValidationReportDto `json:"instructionReports,omitempty"`
	// A map of variable reports. Each key is a variable name and each value a JSON object consisting of the variable's type, value, value info object and a list of failures.
	VariableReports map[string]MigrationVariableValidationReportDto `json:"variableReports,omitempty"`
}

// NewMigrationPlanReportDto instantiates a new MigrationPlanReportDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationPlanReportDto() *MigrationPlanReportDto {
	this := MigrationPlanReportDto{}
	return &this
}

// NewMigrationPlanReportDtoWithDefaults instantiates a new MigrationPlanReportDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationPlanReportDtoWithDefaults() *MigrationPlanReportDto {
	this := MigrationPlanReportDto{}
	return &this
}

// GetInstructionReports returns the InstructionReports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationPlanReportDto) GetInstructionReports() []MigrationInstructionValidationReportDto {
	if o == nil {
		var ret []MigrationInstructionValidationReportDto
		return ret
	}
	return o.InstructionReports
}

// GetInstructionReportsOk returns a tuple with the InstructionReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationPlanReportDto) GetInstructionReportsOk() ([]MigrationInstructionValidationReportDto, bool) {
	if o == nil || IsNil(o.InstructionReports) {
		return nil, false
	}
	return o.InstructionReports, true
}

// HasInstructionReports returns a boolean if a field has been set.
func (o *MigrationPlanReportDto) HasInstructionReports() bool {
	if o != nil && !IsNil(o.InstructionReports) {
		return true
	}

	return false
}

// SetInstructionReports gets a reference to the given []MigrationInstructionValidationReportDto and assigns it to the InstructionReports field.
func (o *MigrationPlanReportDto) SetInstructionReports(v []MigrationInstructionValidationReportDto) {
	o.InstructionReports = v
}

// GetVariableReports returns the VariableReports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationPlanReportDto) GetVariableReports() map[string]MigrationVariableValidationReportDto {
	if o == nil {
		var ret map[string]MigrationVariableValidationReportDto
		return ret
	}
	return o.VariableReports
}

// GetVariableReportsOk returns a tuple with the VariableReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationPlanReportDto) GetVariableReportsOk() (*map[string]MigrationVariableValidationReportDto, bool) {
	if o == nil || IsNil(o.VariableReports) {
		return nil, false
	}
	return &o.VariableReports, true
}

// HasVariableReports returns a boolean if a field has been set.
func (o *MigrationPlanReportDto) HasVariableReports() bool {
	if o != nil && !IsNil(o.VariableReports) {
		return true
	}

	return false
}

// SetVariableReports gets a reference to the given map[string]MigrationVariableValidationReportDto and assigns it to the VariableReports field.
func (o *MigrationPlanReportDto) SetVariableReports(v map[string]MigrationVariableValidationReportDto) {
	o.VariableReports = v
}

func (o MigrationPlanReportDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrationPlanReportDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InstructionReports != nil {
		toSerialize["instructionReports"] = o.InstructionReports
	}
	if o.VariableReports != nil {
		toSerialize["variableReports"] = o.VariableReports
	}
	return toSerialize, nil
}

type NullableMigrationPlanReportDto struct {
	value *MigrationPlanReportDto
	isSet bool
}

func (v NullableMigrationPlanReportDto) Get() *MigrationPlanReportDto {
	return v.value
}

func (v *NullableMigrationPlanReportDto) Set(val *MigrationPlanReportDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationPlanReportDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationPlanReportDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationPlanReportDto(val *MigrationPlanReportDto) *NullableMigrationPlanReportDto {
	return &NullableMigrationPlanReportDto{value: val, isSet: true}
}

func (v NullableMigrationPlanReportDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationPlanReportDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


