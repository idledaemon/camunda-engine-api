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

// checks if the MigrationInstructionValidationReportDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrationInstructionValidationReportDto{}

// MigrationInstructionValidationReportDto struct for MigrationInstructionValidationReportDto
type MigrationInstructionValidationReportDto struct {
	Instruction *MigrationInstructionDto `json:"instruction,omitempty"`
	// A list of instruction validation report messages.
	Failures []string `json:"failures,omitempty"`
}

// NewMigrationInstructionValidationReportDto instantiates a new MigrationInstructionValidationReportDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationInstructionValidationReportDto() *MigrationInstructionValidationReportDto {
	this := MigrationInstructionValidationReportDto{}
	return &this
}

// NewMigrationInstructionValidationReportDtoWithDefaults instantiates a new MigrationInstructionValidationReportDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationInstructionValidationReportDtoWithDefaults() *MigrationInstructionValidationReportDto {
	this := MigrationInstructionValidationReportDto{}
	return &this
}

// GetInstruction returns the Instruction field value if set, zero value otherwise.
func (o *MigrationInstructionValidationReportDto) GetInstruction() MigrationInstructionDto {
	if o == nil || IsNil(o.Instruction) {
		var ret MigrationInstructionDto
		return ret
	}
	return *o.Instruction
}

// GetInstructionOk returns a tuple with the Instruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationInstructionValidationReportDto) GetInstructionOk() (*MigrationInstructionDto, bool) {
	if o == nil || IsNil(o.Instruction) {
		return nil, false
	}
	return o.Instruction, true
}

// HasInstruction returns a boolean if a field has been set.
func (o *MigrationInstructionValidationReportDto) HasInstruction() bool {
	if o != nil && !IsNil(o.Instruction) {
		return true
	}

	return false
}

// SetInstruction gets a reference to the given MigrationInstructionDto and assigns it to the Instruction field.
func (o *MigrationInstructionValidationReportDto) SetInstruction(v MigrationInstructionDto) {
	o.Instruction = &v
}

// GetFailures returns the Failures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationInstructionValidationReportDto) GetFailures() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationInstructionValidationReportDto) GetFailuresOk() ([]string, bool) {
	if o == nil || IsNil(o.Failures) {
		return nil, false
	}
	return o.Failures, true
}

// HasFailures returns a boolean if a field has been set.
func (o *MigrationInstructionValidationReportDto) HasFailures() bool {
	if o != nil && !IsNil(o.Failures) {
		return true
	}

	return false
}

// SetFailures gets a reference to the given []string and assigns it to the Failures field.
func (o *MigrationInstructionValidationReportDto) SetFailures(v []string) {
	o.Failures = v
}

func (o MigrationInstructionValidationReportDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrationInstructionValidationReportDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Instruction) {
		toSerialize["instruction"] = o.Instruction
	}
	if o.Failures != nil {
		toSerialize["failures"] = o.Failures
	}
	return toSerialize, nil
}

type NullableMigrationInstructionValidationReportDto struct {
	value *MigrationInstructionValidationReportDto
	isSet bool
}

func (v NullableMigrationInstructionValidationReportDto) Get() *MigrationInstructionValidationReportDto {
	return v.value
}

func (v *NullableMigrationInstructionValidationReportDto) Set(val *MigrationInstructionValidationReportDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationInstructionValidationReportDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationInstructionValidationReportDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationInstructionValidationReportDto(val *MigrationInstructionValidationReportDto) *NullableMigrationInstructionValidationReportDto {
	return &NullableMigrationInstructionValidationReportDto{value: val, isSet: true}
}

func (v NullableMigrationInstructionValidationReportDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationInstructionValidationReportDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


