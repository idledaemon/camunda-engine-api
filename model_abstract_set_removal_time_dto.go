/*
Camunda Platform REST API

OpenApi Spec for Camunda Platform REST API.

API version: 7.21.2-ee
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package camundarestgo

import (
	"encoding/json"
	"time"
)

// checks if the AbstractSetRemovalTimeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AbstractSetRemovalTimeDto{}

// AbstractSetRemovalTimeDto struct for AbstractSetRemovalTimeDto
type AbstractSetRemovalTimeDto struct {
	// The date for which the instances shall be removed. Value may not be `null`.  **Note:** Cannot be set in conjunction with `clearedRemovalTime` or `calculatedRemovalTime`.
	AbsoluteRemovalTime NullableTime `json:"absoluteRemovalTime,omitempty"`
	// Sets the removal time to `null`. Value may only be `true`, as `false` is the default behavior.  **Note:** Cannot be set in conjunction with `absoluteRemovalTime` or `calculatedRemovalTime`.
	ClearedRemovalTime NullableBool `json:"clearedRemovalTime,omitempty"`
	// The removal time is calculated based on the engine's configuration settings. Value may only be `true`, as `false` is the default behavior.  **Note:** Cannot be set in conjunction with `absoluteRemovalTime` or `clearedRemovalTime`.
	CalculatedRemovalTime NullableBool `json:"calculatedRemovalTime,omitempty"`
}

// NewAbstractSetRemovalTimeDto instantiates a new AbstractSetRemovalTimeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbstractSetRemovalTimeDto() *AbstractSetRemovalTimeDto {
	this := AbstractSetRemovalTimeDto{}
	return &this
}

// NewAbstractSetRemovalTimeDtoWithDefaults instantiates a new AbstractSetRemovalTimeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbstractSetRemovalTimeDtoWithDefaults() *AbstractSetRemovalTimeDto {
	this := AbstractSetRemovalTimeDto{}
	return &this
}

// GetAbsoluteRemovalTime returns the AbsoluteRemovalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractSetRemovalTimeDto) GetAbsoluteRemovalTime() time.Time {
	if o == nil || IsNil(o.AbsoluteRemovalTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AbsoluteRemovalTime.Get()
}

// GetAbsoluteRemovalTimeOk returns a tuple with the AbsoluteRemovalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractSetRemovalTimeDto) GetAbsoluteRemovalTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbsoluteRemovalTime.Get(), o.AbsoluteRemovalTime.IsSet()
}

// HasAbsoluteRemovalTime returns a boolean if a field has been set.
func (o *AbstractSetRemovalTimeDto) HasAbsoluteRemovalTime() bool {
	if o != nil && o.AbsoluteRemovalTime.IsSet() {
		return true
	}

	return false
}

// SetAbsoluteRemovalTime gets a reference to the given NullableTime and assigns it to the AbsoluteRemovalTime field.
func (o *AbstractSetRemovalTimeDto) SetAbsoluteRemovalTime(v time.Time) {
	o.AbsoluteRemovalTime.Set(&v)
}
// SetAbsoluteRemovalTimeNil sets the value for AbsoluteRemovalTime to be an explicit nil
func (o *AbstractSetRemovalTimeDto) SetAbsoluteRemovalTimeNil() {
	o.AbsoluteRemovalTime.Set(nil)
}

// UnsetAbsoluteRemovalTime ensures that no value is present for AbsoluteRemovalTime, not even an explicit nil
func (o *AbstractSetRemovalTimeDto) UnsetAbsoluteRemovalTime() {
	o.AbsoluteRemovalTime.Unset()
}

// GetClearedRemovalTime returns the ClearedRemovalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractSetRemovalTimeDto) GetClearedRemovalTime() bool {
	if o == nil || IsNil(o.ClearedRemovalTime.Get()) {
		var ret bool
		return ret
	}
	return *o.ClearedRemovalTime.Get()
}

// GetClearedRemovalTimeOk returns a tuple with the ClearedRemovalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractSetRemovalTimeDto) GetClearedRemovalTimeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClearedRemovalTime.Get(), o.ClearedRemovalTime.IsSet()
}

// HasClearedRemovalTime returns a boolean if a field has been set.
func (o *AbstractSetRemovalTimeDto) HasClearedRemovalTime() bool {
	if o != nil && o.ClearedRemovalTime.IsSet() {
		return true
	}

	return false
}

// SetClearedRemovalTime gets a reference to the given NullableBool and assigns it to the ClearedRemovalTime field.
func (o *AbstractSetRemovalTimeDto) SetClearedRemovalTime(v bool) {
	o.ClearedRemovalTime.Set(&v)
}
// SetClearedRemovalTimeNil sets the value for ClearedRemovalTime to be an explicit nil
func (o *AbstractSetRemovalTimeDto) SetClearedRemovalTimeNil() {
	o.ClearedRemovalTime.Set(nil)
}

// UnsetClearedRemovalTime ensures that no value is present for ClearedRemovalTime, not even an explicit nil
func (o *AbstractSetRemovalTimeDto) UnsetClearedRemovalTime() {
	o.ClearedRemovalTime.Unset()
}

// GetCalculatedRemovalTime returns the CalculatedRemovalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractSetRemovalTimeDto) GetCalculatedRemovalTime() bool {
	if o == nil || IsNil(o.CalculatedRemovalTime.Get()) {
		var ret bool
		return ret
	}
	return *o.CalculatedRemovalTime.Get()
}

// GetCalculatedRemovalTimeOk returns a tuple with the CalculatedRemovalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractSetRemovalTimeDto) GetCalculatedRemovalTimeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CalculatedRemovalTime.Get(), o.CalculatedRemovalTime.IsSet()
}

// HasCalculatedRemovalTime returns a boolean if a field has been set.
func (o *AbstractSetRemovalTimeDto) HasCalculatedRemovalTime() bool {
	if o != nil && o.CalculatedRemovalTime.IsSet() {
		return true
	}

	return false
}

// SetCalculatedRemovalTime gets a reference to the given NullableBool and assigns it to the CalculatedRemovalTime field.
func (o *AbstractSetRemovalTimeDto) SetCalculatedRemovalTime(v bool) {
	o.CalculatedRemovalTime.Set(&v)
}
// SetCalculatedRemovalTimeNil sets the value for CalculatedRemovalTime to be an explicit nil
func (o *AbstractSetRemovalTimeDto) SetCalculatedRemovalTimeNil() {
	o.CalculatedRemovalTime.Set(nil)
}

// UnsetCalculatedRemovalTime ensures that no value is present for CalculatedRemovalTime, not even an explicit nil
func (o *AbstractSetRemovalTimeDto) UnsetCalculatedRemovalTime() {
	o.CalculatedRemovalTime.Unset()
}

func (o AbstractSetRemovalTimeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AbstractSetRemovalTimeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AbsoluteRemovalTime.IsSet() {
		toSerialize["absoluteRemovalTime"] = o.AbsoluteRemovalTime.Get()
	}
	if o.ClearedRemovalTime.IsSet() {
		toSerialize["clearedRemovalTime"] = o.ClearedRemovalTime.Get()
	}
	if o.CalculatedRemovalTime.IsSet() {
		toSerialize["calculatedRemovalTime"] = o.CalculatedRemovalTime.Get()
	}
	return toSerialize, nil
}

type NullableAbstractSetRemovalTimeDto struct {
	value *AbstractSetRemovalTimeDto
	isSet bool
}

func (v NullableAbstractSetRemovalTimeDto) Get() *AbstractSetRemovalTimeDto {
	return v.value
}

func (v *NullableAbstractSetRemovalTimeDto) Set(val *AbstractSetRemovalTimeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAbstractSetRemovalTimeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAbstractSetRemovalTimeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbstractSetRemovalTimeDto(val *AbstractSetRemovalTimeDto) *NullableAbstractSetRemovalTimeDto {
	return &NullableAbstractSetRemovalTimeDto{value: val, isSet: true}
}

func (v NullableAbstractSetRemovalTimeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbstractSetRemovalTimeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


