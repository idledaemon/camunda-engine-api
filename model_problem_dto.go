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

// checks if the ProblemDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProblemDto{}

// ProblemDto struct for ProblemDto
type ProblemDto struct {
	// The message of the problem.
	Message NullableString `json:"message,omitempty"`
	// The line where the problem occurred.
	Line NullableInt32 `json:"line,omitempty"`
	// The column where the problem occurred.
	Column NullableInt32 `json:"column,omitempty"`
	// The main element id where the problem occurred.
	MainElementId NullableString `json:"mainElementId,omitempty"`
	// A list of element id affected by the problem.
	ElementIds []string `json:"elementIds,omitempty"`
}

// NewProblemDto instantiates a new ProblemDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemDto() *ProblemDto {
	this := ProblemDto{}
	return &this
}

// NewProblemDtoWithDefaults instantiates a new ProblemDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemDtoWithDefaults() *ProblemDto {
	this := ProblemDto{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProblemDto) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProblemDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ProblemDto) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ProblemDto) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *ProblemDto) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ProblemDto) UnsetMessage() {
	o.Message.Unset()
}

// GetLine returns the Line field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProblemDto) GetLine() int32 {
	if o == nil || IsNil(o.Line.Get()) {
		var ret int32
		return ret
	}
	return *o.Line.Get()
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProblemDto) GetLineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line.Get(), o.Line.IsSet()
}

// HasLine returns a boolean if a field has been set.
func (o *ProblemDto) HasLine() bool {
	if o != nil && o.Line.IsSet() {
		return true
	}

	return false
}

// SetLine gets a reference to the given NullableInt32 and assigns it to the Line field.
func (o *ProblemDto) SetLine(v int32) {
	o.Line.Set(&v)
}
// SetLineNil sets the value for Line to be an explicit nil
func (o *ProblemDto) SetLineNil() {
	o.Line.Set(nil)
}

// UnsetLine ensures that no value is present for Line, not even an explicit nil
func (o *ProblemDto) UnsetLine() {
	o.Line.Unset()
}

// GetColumn returns the Column field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProblemDto) GetColumn() int32 {
	if o == nil || IsNil(o.Column.Get()) {
		var ret int32
		return ret
	}
	return *o.Column.Get()
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProblemDto) GetColumnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Column.Get(), o.Column.IsSet()
}

// HasColumn returns a boolean if a field has been set.
func (o *ProblemDto) HasColumn() bool {
	if o != nil && o.Column.IsSet() {
		return true
	}

	return false
}

// SetColumn gets a reference to the given NullableInt32 and assigns it to the Column field.
func (o *ProblemDto) SetColumn(v int32) {
	o.Column.Set(&v)
}
// SetColumnNil sets the value for Column to be an explicit nil
func (o *ProblemDto) SetColumnNil() {
	o.Column.Set(nil)
}

// UnsetColumn ensures that no value is present for Column, not even an explicit nil
func (o *ProblemDto) UnsetColumn() {
	o.Column.Unset()
}

// GetMainElementId returns the MainElementId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProblemDto) GetMainElementId() string {
	if o == nil || IsNil(o.MainElementId.Get()) {
		var ret string
		return ret
	}
	return *o.MainElementId.Get()
}

// GetMainElementIdOk returns a tuple with the MainElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProblemDto) GetMainElementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MainElementId.Get(), o.MainElementId.IsSet()
}

// HasMainElementId returns a boolean if a field has been set.
func (o *ProblemDto) HasMainElementId() bool {
	if o != nil && o.MainElementId.IsSet() {
		return true
	}

	return false
}

// SetMainElementId gets a reference to the given NullableString and assigns it to the MainElementId field.
func (o *ProblemDto) SetMainElementId(v string) {
	o.MainElementId.Set(&v)
}
// SetMainElementIdNil sets the value for MainElementId to be an explicit nil
func (o *ProblemDto) SetMainElementIdNil() {
	o.MainElementId.Set(nil)
}

// UnsetMainElementId ensures that no value is present for MainElementId, not even an explicit nil
func (o *ProblemDto) UnsetMainElementId() {
	o.MainElementId.Unset()
}

// GetElementIds returns the ElementIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProblemDto) GetElementIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ElementIds
}

// GetElementIdsOk returns a tuple with the ElementIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProblemDto) GetElementIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ElementIds) {
		return nil, false
	}
	return o.ElementIds, true
}

// HasElementIds returns a boolean if a field has been set.
func (o *ProblemDto) HasElementIds() bool {
	if o != nil && !IsNil(o.ElementIds) {
		return true
	}

	return false
}

// SetElementIds gets a reference to the given []string and assigns it to the ElementIds field.
func (o *ProblemDto) SetElementIds(v []string) {
	o.ElementIds = v
}

func (o ProblemDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProblemDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Line.IsSet() {
		toSerialize["line"] = o.Line.Get()
	}
	if o.Column.IsSet() {
		toSerialize["column"] = o.Column.Get()
	}
	if o.MainElementId.IsSet() {
		toSerialize["mainElementId"] = o.MainElementId.Get()
	}
	if o.ElementIds != nil {
		toSerialize["elementIds"] = o.ElementIds
	}
	return toSerialize, nil
}

type NullableProblemDto struct {
	value *ProblemDto
	isSet bool
}

func (v NullableProblemDto) Get() *ProblemDto {
	return v.value
}

func (v *NullableProblemDto) Set(val *ProblemDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemDto(val *ProblemDto) *NullableProblemDto {
	return &NullableProblemDto{value: val, isSet: true}
}

func (v NullableProblemDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


