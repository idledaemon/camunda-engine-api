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

// checks if the StartProcessInstanceFormDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartProcessInstanceFormDto{}

// StartProcessInstanceFormDto struct for StartProcessInstanceFormDto
type StartProcessInstanceFormDto struct {
	// 
	Variables map[string]VariableValueDto `json:"variables,omitempty"`
	// The business key the process instance is to be initialized with. The business key uniquely identifies the process instance in the context of the given process definition.
	BusinessKey NullableString `json:"businessKey,omitempty"`
}

// NewStartProcessInstanceFormDto instantiates a new StartProcessInstanceFormDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartProcessInstanceFormDto() *StartProcessInstanceFormDto {
	this := StartProcessInstanceFormDto{}
	return &this
}

// NewStartProcessInstanceFormDtoWithDefaults instantiates a new StartProcessInstanceFormDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartProcessInstanceFormDtoWithDefaults() *StartProcessInstanceFormDto {
	this := StartProcessInstanceFormDto{}
	return &this
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartProcessInstanceFormDto) GetVariables() map[string]VariableValueDto {
	if o == nil {
		var ret map[string]VariableValueDto
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartProcessInstanceFormDto) GetVariablesOk() (*map[string]VariableValueDto, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *StartProcessInstanceFormDto) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]VariableValueDto and assigns it to the Variables field.
func (o *StartProcessInstanceFormDto) SetVariables(v map[string]VariableValueDto) {
	o.Variables = v
}

// GetBusinessKey returns the BusinessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartProcessInstanceFormDto) GetBusinessKey() string {
	if o == nil || IsNil(o.BusinessKey.Get()) {
		var ret string
		return ret
	}
	return *o.BusinessKey.Get()
}

// GetBusinessKeyOk returns a tuple with the BusinessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartProcessInstanceFormDto) GetBusinessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessKey.Get(), o.BusinessKey.IsSet()
}

// HasBusinessKey returns a boolean if a field has been set.
func (o *StartProcessInstanceFormDto) HasBusinessKey() bool {
	if o != nil && o.BusinessKey.IsSet() {
		return true
	}

	return false
}

// SetBusinessKey gets a reference to the given NullableString and assigns it to the BusinessKey field.
func (o *StartProcessInstanceFormDto) SetBusinessKey(v string) {
	o.BusinessKey.Set(&v)
}
// SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil
func (o *StartProcessInstanceFormDto) SetBusinessKeyNil() {
	o.BusinessKey.Set(nil)
}

// UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
func (o *StartProcessInstanceFormDto) UnsetBusinessKey() {
	o.BusinessKey.Unset()
}

func (o StartProcessInstanceFormDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartProcessInstanceFormDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	if o.BusinessKey.IsSet() {
		toSerialize["businessKey"] = o.BusinessKey.Get()
	}
	return toSerialize, nil
}

type NullableStartProcessInstanceFormDto struct {
	value *StartProcessInstanceFormDto
	isSet bool
}

func (v NullableStartProcessInstanceFormDto) Get() *StartProcessInstanceFormDto {
	return v.value
}

func (v *NullableStartProcessInstanceFormDto) Set(val *StartProcessInstanceFormDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStartProcessInstanceFormDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStartProcessInstanceFormDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartProcessInstanceFormDto(val *StartProcessInstanceFormDto) *NullableStartProcessInstanceFormDto {
	return &NullableStartProcessInstanceFormDto{value: val, isSet: true}
}

func (v NullableStartProcessInstanceFormDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartProcessInstanceFormDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

