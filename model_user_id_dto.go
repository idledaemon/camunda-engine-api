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

// checks if the UserIdDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdDto{}

// UserIdDto struct for UserIdDto
type UserIdDto struct {
	// The id of the user that the current action refers to.
	UserId NullableString `json:"userId,omitempty"`
}

// NewUserIdDto instantiates a new UserIdDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdDto() *UserIdDto {
	this := UserIdDto{}
	return &this
}

// NewUserIdDtoWithDefaults instantiates a new UserIdDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdDtoWithDefaults() *UserIdDto {
	this := UserIdDto{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserIdDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserIdDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *UserIdDto) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *UserIdDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *UserIdDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *UserIdDto) UnsetUserId() {
	o.UserId.Unset()
}

func (o UserIdDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	return toSerialize, nil
}

type NullableUserIdDto struct {
	value *UserIdDto
	isSet bool
}

func (v NullableUserIdDto) Get() *UserIdDto {
	return v.value
}

func (v *NullableUserIdDto) Set(val *UserIdDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdDto(val *UserIdDto) *NullableUserIdDto {
	return &NullableUserIdDto{value: val, isSet: true}
}

func (v NullableUserIdDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


