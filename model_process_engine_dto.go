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

// checks if the ProcessEngineDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessEngineDto{}

// ProcessEngineDto struct for ProcessEngineDto
type ProcessEngineDto struct {
	// The name of the process engine.
	Name NullableString `json:"name,omitempty"`
}

// NewProcessEngineDto instantiates a new ProcessEngineDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessEngineDto() *ProcessEngineDto {
	this := ProcessEngineDto{}
	return &this
}

// NewProcessEngineDtoWithDefaults instantiates a new ProcessEngineDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessEngineDtoWithDefaults() *ProcessEngineDto {
	this := ProcessEngineDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessEngineDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessEngineDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProcessEngineDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProcessEngineDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProcessEngineDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProcessEngineDto) UnsetName() {
	o.Name.Unset()
}

func (o ProcessEngineDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessEngineDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableProcessEngineDto struct {
	value *ProcessEngineDto
	isSet bool
}

func (v NullableProcessEngineDto) Get() *ProcessEngineDto {
	return v.value
}

func (v *NullableProcessEngineDto) Set(val *ProcessEngineDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessEngineDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessEngineDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessEngineDto(val *ProcessEngineDto) *NullableProcessEngineDto {
	return &NullableProcessEngineDto{value: val, isSet: true}
}

func (v NullableProcessEngineDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessEngineDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


