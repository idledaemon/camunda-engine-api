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

// checks if the AnnotationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnnotationDto{}

// AnnotationDto struct for AnnotationDto
type AnnotationDto struct {
	// The annotation value to put.
	Annotation NullableString `json:"annotation,omitempty"`
}

// NewAnnotationDto instantiates a new AnnotationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotationDto() *AnnotationDto {
	this := AnnotationDto{}
	return &this
}

// NewAnnotationDtoWithDefaults instantiates a new AnnotationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationDtoWithDefaults() *AnnotationDto {
	this := AnnotationDto{}
	return &this
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnnotationDto) GetAnnotation() string {
	if o == nil || IsNil(o.Annotation.Get()) {
		var ret string
		return ret
	}
	return *o.Annotation.Get()
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnnotationDto) GetAnnotationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Annotation.Get(), o.Annotation.IsSet()
}

// HasAnnotation returns a boolean if a field has been set.
func (o *AnnotationDto) HasAnnotation() bool {
	if o != nil && o.Annotation.IsSet() {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given NullableString and assigns it to the Annotation field.
func (o *AnnotationDto) SetAnnotation(v string) {
	o.Annotation.Set(&v)
}
// SetAnnotationNil sets the value for Annotation to be an explicit nil
func (o *AnnotationDto) SetAnnotationNil() {
	o.Annotation.Set(nil)
}

// UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil
func (o *AnnotationDto) UnsetAnnotation() {
	o.Annotation.Unset()
}

func (o AnnotationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnnotationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Annotation.IsSet() {
		toSerialize["annotation"] = o.Annotation.Get()
	}
	return toSerialize, nil
}

type NullableAnnotationDto struct {
	value *AnnotationDto
	isSet bool
}

func (v NullableAnnotationDto) Get() *AnnotationDto {
	return v.value
}

func (v *NullableAnnotationDto) Set(val *AnnotationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationDto(val *AnnotationDto) *NullableAnnotationDto {
	return &NullableAnnotationDto{value: val, isSet: true}
}

func (v NullableAnnotationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

