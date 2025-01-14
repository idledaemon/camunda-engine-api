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

// checks if the ExecutionTriggerDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionTriggerDto{}

// ExecutionTriggerDto struct for ExecutionTriggerDto
type ExecutionTriggerDto struct {
	// A JSON object containing variable key-value pairs. Each key is a variable name and each value a JSON variable value object.
	Variables map[string]VariableValueDto `json:"variables,omitempty"`
}

// NewExecutionTriggerDto instantiates a new ExecutionTriggerDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionTriggerDto() *ExecutionTriggerDto {
	this := ExecutionTriggerDto{}
	return &this
}

// NewExecutionTriggerDtoWithDefaults instantiates a new ExecutionTriggerDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionTriggerDtoWithDefaults() *ExecutionTriggerDto {
	this := ExecutionTriggerDto{}
	return &this
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecutionTriggerDto) GetVariables() map[string]VariableValueDto {
	if o == nil {
		var ret map[string]VariableValueDto
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExecutionTriggerDto) GetVariablesOk() (*map[string]VariableValueDto, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *ExecutionTriggerDto) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]VariableValueDto and assigns it to the Variables field.
func (o *ExecutionTriggerDto) SetVariables(v map[string]VariableValueDto) {
	o.Variables = v
}

func (o ExecutionTriggerDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionTriggerDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

type NullableExecutionTriggerDto struct {
	value *ExecutionTriggerDto
	isSet bool
}

func (v NullableExecutionTriggerDto) Get() *ExecutionTriggerDto {
	return v.value
}

func (v *NullableExecutionTriggerDto) Set(val *ExecutionTriggerDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionTriggerDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionTriggerDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionTriggerDto(val *ExecutionTriggerDto) *NullableExecutionTriggerDto {
	return &NullableExecutionTriggerDto{value: val, isSet: true}
}

func (v NullableExecutionTriggerDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionTriggerDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


