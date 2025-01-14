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

// checks if the VariableInstanceQueryDtoSortingInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableInstanceQueryDtoSortingInner{}

// VariableInstanceQueryDtoSortingInner struct for VariableInstanceQueryDtoSortingInner
type VariableInstanceQueryDtoSortingInner struct {
	// Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter.
	SortBy NullableString `json:"sortBy,omitempty"`
	// Sort the results in a given order. Values may be `asc` for ascending order or `desc` for descending order. Must be used in conjunction with the sortBy parameter.
	SortOrder NullableString `json:"sortOrder,omitempty"`
}

// NewVariableInstanceQueryDtoSortingInner instantiates a new VariableInstanceQueryDtoSortingInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableInstanceQueryDtoSortingInner() *VariableInstanceQueryDtoSortingInner {
	this := VariableInstanceQueryDtoSortingInner{}
	return &this
}

// NewVariableInstanceQueryDtoSortingInnerWithDefaults instantiates a new VariableInstanceQueryDtoSortingInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableInstanceQueryDtoSortingInnerWithDefaults() *VariableInstanceQueryDtoSortingInner {
	this := VariableInstanceQueryDtoSortingInner{}
	return &this
}

// GetSortBy returns the SortBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VariableInstanceQueryDtoSortingInner) GetSortBy() string {
	if o == nil || IsNil(o.SortBy.Get()) {
		var ret string
		return ret
	}
	return *o.SortBy.Get()
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariableInstanceQueryDtoSortingInner) GetSortByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortBy.Get(), o.SortBy.IsSet()
}

// HasSortBy returns a boolean if a field has been set.
func (o *VariableInstanceQueryDtoSortingInner) HasSortBy() bool {
	if o != nil && o.SortBy.IsSet() {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given NullableString and assigns it to the SortBy field.
func (o *VariableInstanceQueryDtoSortingInner) SetSortBy(v string) {
	o.SortBy.Set(&v)
}
// SetSortByNil sets the value for SortBy to be an explicit nil
func (o *VariableInstanceQueryDtoSortingInner) SetSortByNil() {
	o.SortBy.Set(nil)
}

// UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
func (o *VariableInstanceQueryDtoSortingInner) UnsetSortBy() {
	o.SortBy.Unset()
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VariableInstanceQueryDtoSortingInner) GetSortOrder() string {
	if o == nil || IsNil(o.SortOrder.Get()) {
		var ret string
		return ret
	}
	return *o.SortOrder.Get()
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariableInstanceQueryDtoSortingInner) GetSortOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortOrder.Get(), o.SortOrder.IsSet()
}

// HasSortOrder returns a boolean if a field has been set.
func (o *VariableInstanceQueryDtoSortingInner) HasSortOrder() bool {
	if o != nil && o.SortOrder.IsSet() {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given NullableString and assigns it to the SortOrder field.
func (o *VariableInstanceQueryDtoSortingInner) SetSortOrder(v string) {
	o.SortOrder.Set(&v)
}
// SetSortOrderNil sets the value for SortOrder to be an explicit nil
func (o *VariableInstanceQueryDtoSortingInner) SetSortOrderNil() {
	o.SortOrder.Set(nil)
}

// UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
func (o *VariableInstanceQueryDtoSortingInner) UnsetSortOrder() {
	o.SortOrder.Unset()
}

func (o VariableInstanceQueryDtoSortingInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableInstanceQueryDtoSortingInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SortBy.IsSet() {
		toSerialize["sortBy"] = o.SortBy.Get()
	}
	if o.SortOrder.IsSet() {
		toSerialize["sortOrder"] = o.SortOrder.Get()
	}
	return toSerialize, nil
}

type NullableVariableInstanceQueryDtoSortingInner struct {
	value *VariableInstanceQueryDtoSortingInner
	isSet bool
}

func (v NullableVariableInstanceQueryDtoSortingInner) Get() *VariableInstanceQueryDtoSortingInner {
	return v.value
}

func (v *NullableVariableInstanceQueryDtoSortingInner) Set(val *VariableInstanceQueryDtoSortingInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableInstanceQueryDtoSortingInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableInstanceQueryDtoSortingInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableInstanceQueryDtoSortingInner(val *VariableInstanceQueryDtoSortingInner) *NullableVariableInstanceQueryDtoSortingInner {
	return &NullableVariableInstanceQueryDtoSortingInner{value: val, isSet: true}
}

func (v NullableVariableInstanceQueryDtoSortingInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableInstanceQueryDtoSortingInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


