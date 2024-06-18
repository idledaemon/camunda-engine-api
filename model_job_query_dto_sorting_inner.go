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

// checks if the JobQueryDtoSortingInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobQueryDtoSortingInner{}

// JobQueryDtoSortingInner struct for JobQueryDtoSortingInner
type JobQueryDtoSortingInner struct {
	// Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter.
	SortBy NullableString `json:"sortBy,omitempty"`
	// Sort the results in a given order. Values may be `asc` for ascending order or `desc` for descending order. Must be used in conjunction with the sortBy parameter.
	SortOrder NullableString `json:"sortOrder,omitempty"`
}

// NewJobQueryDtoSortingInner instantiates a new JobQueryDtoSortingInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobQueryDtoSortingInner() *JobQueryDtoSortingInner {
	this := JobQueryDtoSortingInner{}
	return &this
}

// NewJobQueryDtoSortingInnerWithDefaults instantiates a new JobQueryDtoSortingInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobQueryDtoSortingInnerWithDefaults() *JobQueryDtoSortingInner {
	this := JobQueryDtoSortingInner{}
	return &this
}

// GetSortBy returns the SortBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobQueryDtoSortingInner) GetSortBy() string {
	if o == nil || IsNil(o.SortBy.Get()) {
		var ret string
		return ret
	}
	return *o.SortBy.Get()
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobQueryDtoSortingInner) GetSortByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortBy.Get(), o.SortBy.IsSet()
}

// HasSortBy returns a boolean if a field has been set.
func (o *JobQueryDtoSortingInner) HasSortBy() bool {
	if o != nil && o.SortBy.IsSet() {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given NullableString and assigns it to the SortBy field.
func (o *JobQueryDtoSortingInner) SetSortBy(v string) {
	o.SortBy.Set(&v)
}
// SetSortByNil sets the value for SortBy to be an explicit nil
func (o *JobQueryDtoSortingInner) SetSortByNil() {
	o.SortBy.Set(nil)
}

// UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
func (o *JobQueryDtoSortingInner) UnsetSortBy() {
	o.SortBy.Unset()
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobQueryDtoSortingInner) GetSortOrder() string {
	if o == nil || IsNil(o.SortOrder.Get()) {
		var ret string
		return ret
	}
	return *o.SortOrder.Get()
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobQueryDtoSortingInner) GetSortOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortOrder.Get(), o.SortOrder.IsSet()
}

// HasSortOrder returns a boolean if a field has been set.
func (o *JobQueryDtoSortingInner) HasSortOrder() bool {
	if o != nil && o.SortOrder.IsSet() {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given NullableString and assigns it to the SortOrder field.
func (o *JobQueryDtoSortingInner) SetSortOrder(v string) {
	o.SortOrder.Set(&v)
}
// SetSortOrderNil sets the value for SortOrder to be an explicit nil
func (o *JobQueryDtoSortingInner) SetSortOrderNil() {
	o.SortOrder.Set(nil)
}

// UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
func (o *JobQueryDtoSortingInner) UnsetSortOrder() {
	o.SortOrder.Unset()
}

func (o JobQueryDtoSortingInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobQueryDtoSortingInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SortBy.IsSet() {
		toSerialize["sortBy"] = o.SortBy.Get()
	}
	if o.SortOrder.IsSet() {
		toSerialize["sortOrder"] = o.SortOrder.Get()
	}
	return toSerialize, nil
}

type NullableJobQueryDtoSortingInner struct {
	value *JobQueryDtoSortingInner
	isSet bool
}

func (v NullableJobQueryDtoSortingInner) Get() *JobQueryDtoSortingInner {
	return v.value
}

func (v *NullableJobQueryDtoSortingInner) Set(val *JobQueryDtoSortingInner) {
	v.value = val
	v.isSet = true
}

func (v NullableJobQueryDtoSortingInner) IsSet() bool {
	return v.isSet
}

func (v *NullableJobQueryDtoSortingInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobQueryDtoSortingInner(val *JobQueryDtoSortingInner) *NullableJobQueryDtoSortingInner {
	return &NullableJobQueryDtoSortingInner{value: val, isSet: true}
}

func (v NullableJobQueryDtoSortingInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobQueryDtoSortingInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


