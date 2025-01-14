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

// checks if the MigrationPlanDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrationPlanDto{}

// MigrationPlanDto struct for MigrationPlanDto
type MigrationPlanDto struct {
	// The id of the source process definition for the migration.
	SourceProcessDefinitionId NullableString `json:"sourceProcessDefinitionId,omitempty"`
	// The id of the target process definition for the migration.
	TargetProcessDefinitionId NullableString `json:"targetProcessDefinitionId,omitempty"`
	// A list of migration instructions which map equal activities. Each migration instruction is a JSON object with the following properties:
	Instructions []MigrationInstructionDto `json:"instructions,omitempty"`
	// A map of variables which will be set into the process instances' scope. Each key is a variable name and each value a JSON variable value object.
	Variables map[string]VariableValueDto `json:"variables,omitempty"`
}

// NewMigrationPlanDto instantiates a new MigrationPlanDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationPlanDto() *MigrationPlanDto {
	this := MigrationPlanDto{}
	return &this
}

// NewMigrationPlanDtoWithDefaults instantiates a new MigrationPlanDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationPlanDtoWithDefaults() *MigrationPlanDto {
	this := MigrationPlanDto{}
	return &this
}

// GetSourceProcessDefinitionId returns the SourceProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationPlanDto) GetSourceProcessDefinitionId() string {
	if o == nil || IsNil(o.SourceProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.SourceProcessDefinitionId.Get()
}

// GetSourceProcessDefinitionIdOk returns a tuple with the SourceProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationPlanDto) GetSourceProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceProcessDefinitionId.Get(), o.SourceProcessDefinitionId.IsSet()
}

// HasSourceProcessDefinitionId returns a boolean if a field has been set.
func (o *MigrationPlanDto) HasSourceProcessDefinitionId() bool {
	if o != nil && o.SourceProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetSourceProcessDefinitionId gets a reference to the given NullableString and assigns it to the SourceProcessDefinitionId field.
func (o *MigrationPlanDto) SetSourceProcessDefinitionId(v string) {
	o.SourceProcessDefinitionId.Set(&v)
}
// SetSourceProcessDefinitionIdNil sets the value for SourceProcessDefinitionId to be an explicit nil
func (o *MigrationPlanDto) SetSourceProcessDefinitionIdNil() {
	o.SourceProcessDefinitionId.Set(nil)
}

// UnsetSourceProcessDefinitionId ensures that no value is present for SourceProcessDefinitionId, not even an explicit nil
func (o *MigrationPlanDto) UnsetSourceProcessDefinitionId() {
	o.SourceProcessDefinitionId.Unset()
}

// GetTargetProcessDefinitionId returns the TargetProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationPlanDto) GetTargetProcessDefinitionId() string {
	if o == nil || IsNil(o.TargetProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.TargetProcessDefinitionId.Get()
}

// GetTargetProcessDefinitionIdOk returns a tuple with the TargetProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationPlanDto) GetTargetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetProcessDefinitionId.Get(), o.TargetProcessDefinitionId.IsSet()
}

// HasTargetProcessDefinitionId returns a boolean if a field has been set.
func (o *MigrationPlanDto) HasTargetProcessDefinitionId() bool {
	if o != nil && o.TargetProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetTargetProcessDefinitionId gets a reference to the given NullableString and assigns it to the TargetProcessDefinitionId field.
func (o *MigrationPlanDto) SetTargetProcessDefinitionId(v string) {
	o.TargetProcessDefinitionId.Set(&v)
}
// SetTargetProcessDefinitionIdNil sets the value for TargetProcessDefinitionId to be an explicit nil
func (o *MigrationPlanDto) SetTargetProcessDefinitionIdNil() {
	o.TargetProcessDefinitionId.Set(nil)
}

// UnsetTargetProcessDefinitionId ensures that no value is present for TargetProcessDefinitionId, not even an explicit nil
func (o *MigrationPlanDto) UnsetTargetProcessDefinitionId() {
	o.TargetProcessDefinitionId.Unset()
}

// GetInstructions returns the Instructions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationPlanDto) GetInstructions() []MigrationInstructionDto {
	if o == nil {
		var ret []MigrationInstructionDto
		return ret
	}
	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationPlanDto) GetInstructionsOk() ([]MigrationInstructionDto, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *MigrationPlanDto) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given []MigrationInstructionDto and assigns it to the Instructions field.
func (o *MigrationPlanDto) SetInstructions(v []MigrationInstructionDto) {
	o.Instructions = v
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationPlanDto) GetVariables() map[string]VariableValueDto {
	if o == nil {
		var ret map[string]VariableValueDto
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationPlanDto) GetVariablesOk() (*map[string]VariableValueDto, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *MigrationPlanDto) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]VariableValueDto and assigns it to the Variables field.
func (o *MigrationPlanDto) SetVariables(v map[string]VariableValueDto) {
	o.Variables = v
}

func (o MigrationPlanDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrationPlanDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceProcessDefinitionId.IsSet() {
		toSerialize["sourceProcessDefinitionId"] = o.SourceProcessDefinitionId.Get()
	}
	if o.TargetProcessDefinitionId.IsSet() {
		toSerialize["targetProcessDefinitionId"] = o.TargetProcessDefinitionId.Get()
	}
	if o.Instructions != nil {
		toSerialize["instructions"] = o.Instructions
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

type NullableMigrationPlanDto struct {
	value *MigrationPlanDto
	isSet bool
}

func (v NullableMigrationPlanDto) Get() *MigrationPlanDto {
	return v.value
}

func (v *NullableMigrationPlanDto) Set(val *MigrationPlanDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationPlanDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationPlanDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationPlanDto(val *MigrationPlanDto) *NullableMigrationPlanDto {
	return &NullableMigrationPlanDto{value: val, isSet: true}
}

func (v NullableMigrationPlanDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationPlanDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


