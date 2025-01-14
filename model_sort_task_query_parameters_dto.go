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

// checks if the SortTaskQueryParametersDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SortTaskQueryParametersDto{}

// SortTaskQueryParametersDto Mandatory when `sortBy` is one of the following values: `processVariable`, `executionVariable`, `taskVariable`, `caseExecutionVariable` or `caseInstanceVariable`. Must be a JSON object with the properties `variable` and `type` where `variable` is a variable name and `type` is the name of a variable value type.
type SortTaskQueryParametersDto struct {
	// The name of the variable to sort by.
	Variable NullableString `json:"variable,omitempty"`
	// The name of the type of the variable value.
	Type NullableString `json:"type,omitempty"`
}

// NewSortTaskQueryParametersDto instantiates a new SortTaskQueryParametersDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortTaskQueryParametersDto() *SortTaskQueryParametersDto {
	this := SortTaskQueryParametersDto{}
	return &this
}

// NewSortTaskQueryParametersDtoWithDefaults instantiates a new SortTaskQueryParametersDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortTaskQueryParametersDtoWithDefaults() *SortTaskQueryParametersDto {
	this := SortTaskQueryParametersDto{}
	return &this
}

// GetVariable returns the Variable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SortTaskQueryParametersDto) GetVariable() string {
	if o == nil || IsNil(o.Variable.Get()) {
		var ret string
		return ret
	}
	return *o.Variable.Get()
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SortTaskQueryParametersDto) GetVariableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variable.Get(), o.Variable.IsSet()
}

// HasVariable returns a boolean if a field has been set.
func (o *SortTaskQueryParametersDto) HasVariable() bool {
	if o != nil && o.Variable.IsSet() {
		return true
	}

	return false
}

// SetVariable gets a reference to the given NullableString and assigns it to the Variable field.
func (o *SortTaskQueryParametersDto) SetVariable(v string) {
	o.Variable.Set(&v)
}
// SetVariableNil sets the value for Variable to be an explicit nil
func (o *SortTaskQueryParametersDto) SetVariableNil() {
	o.Variable.Set(nil)
}

// UnsetVariable ensures that no value is present for Variable, not even an explicit nil
func (o *SortTaskQueryParametersDto) UnsetVariable() {
	o.Variable.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SortTaskQueryParametersDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SortTaskQueryParametersDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *SortTaskQueryParametersDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *SortTaskQueryParametersDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *SortTaskQueryParametersDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *SortTaskQueryParametersDto) UnsetType() {
	o.Type.Unset()
}

func (o SortTaskQueryParametersDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SortTaskQueryParametersDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Variable.IsSet() {
		toSerialize["variable"] = o.Variable.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullableSortTaskQueryParametersDto struct {
	value *SortTaskQueryParametersDto
	isSet bool
}

func (v NullableSortTaskQueryParametersDto) Get() *SortTaskQueryParametersDto {
	return v.value
}

func (v *NullableSortTaskQueryParametersDto) Set(val *SortTaskQueryParametersDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSortTaskQueryParametersDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSortTaskQueryParametersDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortTaskQueryParametersDto(val *SortTaskQueryParametersDto) *NullableSortTaskQueryParametersDto {
	return &NullableSortTaskQueryParametersDto{value: val, isSet: true}
}

func (v NullableSortTaskQueryParametersDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortTaskQueryParametersDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


