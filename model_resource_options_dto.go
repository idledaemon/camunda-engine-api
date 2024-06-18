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

// checks if the ResourceOptionsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceOptionsDto{}

// ResourceOptionsDto struct for ResourceOptionsDto
type ResourceOptionsDto struct {
	// The links associated to this resource, with `method`, `href` and `rel`.
	Links []AtomLink `json:"links,omitempty"`
}

// NewResourceOptionsDto instantiates a new ResourceOptionsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOptionsDto() *ResourceOptionsDto {
	this := ResourceOptionsDto{}
	return &this
}

// NewResourceOptionsDtoWithDefaults instantiates a new ResourceOptionsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOptionsDtoWithDefaults() *ResourceOptionsDto {
	this := ResourceOptionsDto{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceOptionsDto) GetLinks() []AtomLink {
	if o == nil {
		var ret []AtomLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceOptionsDto) GetLinksOk() ([]AtomLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceOptionsDto) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []AtomLink and assigns it to the Links field.
func (o *ResourceOptionsDto) SetLinks(v []AtomLink) {
	o.Links = v
}

func (o ResourceOptionsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceOptionsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableResourceOptionsDto struct {
	value *ResourceOptionsDto
	isSet bool
}

func (v NullableResourceOptionsDto) Get() *ResourceOptionsDto {
	return v.value
}

func (v *NullableResourceOptionsDto) Set(val *ResourceOptionsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOptionsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOptionsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOptionsDto(val *ResourceOptionsDto) *NullableResourceOptionsDto {
	return &NullableResourceOptionsDto{value: val, isSet: true}
}

func (v NullableResourceOptionsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOptionsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


