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

// checks if the TriggerVariableValueDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerVariableValueDto{}

// TriggerVariableValueDto struct for TriggerVariableValueDto
type TriggerVariableValueDto struct {
	// Can be any value - string, number, boolean, array or object.  **Note**: Not every endpoint supports every type.
	Value interface{} `json:"value,omitempty"`
	// The value type of the variable.
	Type NullableString `json:"type,omitempty"`
	// A JSON object containing additional, value-type-dependent properties. For serialized variables of type Object, the following properties can be provided:  * `objectTypeName`: A string representation of the object's type name. * `serializationDataFormat`: The serialization format used to store the variable.  For serialized variables of type File, the following properties can be provided:  * `filename`: The name of the file. This is not the variable name but the name that will be used when downloading the file again. * `mimetype`: The MIME type of the file that is being uploaded. * `encoding`: The encoding of the file that is being uploaded.  The following property can be provided for all value types:  * `transient`: Indicates whether the variable should be transient or not. See [documentation](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables#transient-variables) for more informations. (Not applicable for `decision-definition`, ` /process-instance/variables-async`, and `/migration/executeAsync` endpoints)
	ValueInfo map[string]interface{} `json:"valueInfo,omitempty"`
	// Indicates whether the variable should be a local variable or not. If set to true, the variable becomes a local variable of the execution entering the target activity.
	Local NullableBool `json:"local,omitempty"`
}

// NewTriggerVariableValueDto instantiates a new TriggerVariableValueDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerVariableValueDto() *TriggerVariableValueDto {
	this := TriggerVariableValueDto{}
	return &this
}

// NewTriggerVariableValueDtoWithDefaults instantiates a new TriggerVariableValueDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerVariableValueDtoWithDefaults() *TriggerVariableValueDto {
	this := TriggerVariableValueDto{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerVariableValueDto) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerVariableValueDto) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TriggerVariableValueDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *TriggerVariableValueDto) SetValue(v interface{}) {
	o.Value = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerVariableValueDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerVariableValueDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *TriggerVariableValueDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *TriggerVariableValueDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *TriggerVariableValueDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *TriggerVariableValueDto) UnsetType() {
	o.Type.Unset()
}

// GetValueInfo returns the ValueInfo field value if set, zero value otherwise.
func (o *TriggerVariableValueDto) GetValueInfo() map[string]interface{} {
	if o == nil || IsNil(o.ValueInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.ValueInfo
}

// GetValueInfoOk returns a tuple with the ValueInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerVariableValueDto) GetValueInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ValueInfo) {
		return map[string]interface{}{}, false
	}
	return o.ValueInfo, true
}

// HasValueInfo returns a boolean if a field has been set.
func (o *TriggerVariableValueDto) HasValueInfo() bool {
	if o != nil && !IsNil(o.ValueInfo) {
		return true
	}

	return false
}

// SetValueInfo gets a reference to the given map[string]interface{} and assigns it to the ValueInfo field.
func (o *TriggerVariableValueDto) SetValueInfo(v map[string]interface{}) {
	o.ValueInfo = v
}

// GetLocal returns the Local field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerVariableValueDto) GetLocal() bool {
	if o == nil || IsNil(o.Local.Get()) {
		var ret bool
		return ret
	}
	return *o.Local.Get()
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerVariableValueDto) GetLocalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Local.Get(), o.Local.IsSet()
}

// HasLocal returns a boolean if a field has been set.
func (o *TriggerVariableValueDto) HasLocal() bool {
	if o != nil && o.Local.IsSet() {
		return true
	}

	return false
}

// SetLocal gets a reference to the given NullableBool and assigns it to the Local field.
func (o *TriggerVariableValueDto) SetLocal(v bool) {
	o.Local.Set(&v)
}
// SetLocalNil sets the value for Local to be an explicit nil
func (o *TriggerVariableValueDto) SetLocalNil() {
	o.Local.Set(nil)
}

// UnsetLocal ensures that no value is present for Local, not even an explicit nil
func (o *TriggerVariableValueDto) UnsetLocal() {
	o.Local.Unset()
}

func (o TriggerVariableValueDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerVariableValueDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.ValueInfo) {
		toSerialize["valueInfo"] = o.ValueInfo
	}
	if o.Local.IsSet() {
		toSerialize["local"] = o.Local.Get()
	}
	return toSerialize, nil
}

type NullableTriggerVariableValueDto struct {
	value *TriggerVariableValueDto
	isSet bool
}

func (v NullableTriggerVariableValueDto) Get() *TriggerVariableValueDto {
	return v.value
}

func (v *NullableTriggerVariableValueDto) Set(val *TriggerVariableValueDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerVariableValueDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerVariableValueDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerVariableValueDto(val *TriggerVariableValueDto) *NullableTriggerVariableValueDto {
	return &NullableTriggerVariableValueDto{value: val, isSet: true}
}

func (v NullableTriggerVariableValueDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerVariableValueDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

