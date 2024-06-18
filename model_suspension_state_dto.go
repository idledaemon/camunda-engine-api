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

// checks if the SuspensionStateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuspensionStateDto{}

// SuspensionStateDto struct for SuspensionStateDto
type SuspensionStateDto struct {
	// A Boolean value which indicates whether to activate or suspend a given instance  (e.g. process instance, job, job definition, or batch). When the value is set to true,  the given instance will be suspended and when the value is set to false,  the given instance will be activated.
	Suspended NullableBool `json:"suspended,omitempty"`
}

// NewSuspensionStateDto instantiates a new SuspensionStateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuspensionStateDto() *SuspensionStateDto {
	this := SuspensionStateDto{}
	return &this
}

// NewSuspensionStateDtoWithDefaults instantiates a new SuspensionStateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuspensionStateDtoWithDefaults() *SuspensionStateDto {
	this := SuspensionStateDto{}
	return &this
}

// GetSuspended returns the Suspended field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SuspensionStateDto) GetSuspended() bool {
	if o == nil || IsNil(o.Suspended.Get()) {
		var ret bool
		return ret
	}
	return *o.Suspended.Get()
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SuspensionStateDto) GetSuspendedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suspended.Get(), o.Suspended.IsSet()
}

// HasSuspended returns a boolean if a field has been set.
func (o *SuspensionStateDto) HasSuspended() bool {
	if o != nil && o.Suspended.IsSet() {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given NullableBool and assigns it to the Suspended field.
func (o *SuspensionStateDto) SetSuspended(v bool) {
	o.Suspended.Set(&v)
}
// SetSuspendedNil sets the value for Suspended to be an explicit nil
func (o *SuspensionStateDto) SetSuspendedNil() {
	o.Suspended.Set(nil)
}

// UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
func (o *SuspensionStateDto) UnsetSuspended() {
	o.Suspended.Unset()
}

func (o SuspensionStateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuspensionStateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Suspended.IsSet() {
		toSerialize["suspended"] = o.Suspended.Get()
	}
	return toSerialize, nil
}

type NullableSuspensionStateDto struct {
	value *SuspensionStateDto
	isSet bool
}

func (v NullableSuspensionStateDto) Get() *SuspensionStateDto {
	return v.value
}

func (v *NullableSuspensionStateDto) Set(val *SuspensionStateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSuspensionStateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSuspensionStateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuspensionStateDto(val *SuspensionStateDto) *NullableSuspensionStateDto {
	return &NullableSuspensionStateDto{value: val, isSet: true}
}

func (v NullableSuspensionStateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuspensionStateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

