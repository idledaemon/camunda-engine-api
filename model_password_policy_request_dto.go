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

// checks if the PasswordPolicyRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyRequestDto{}

// PasswordPolicyRequestDto struct for PasswordPolicyRequestDto
type PasswordPolicyRequestDto struct {
	// The candidate password to be check against the password policy.
	Password NullableString `json:"password,omitempty"`
	Profile *UserProfileDto `json:"profile,omitempty"`
}

// NewPasswordPolicyRequestDto instantiates a new PasswordPolicyRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRequestDto() *PasswordPolicyRequestDto {
	this := PasswordPolicyRequestDto{}
	return &this
}

// NewPasswordPolicyRequestDtoWithDefaults instantiates a new PasswordPolicyRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRequestDtoWithDefaults() *PasswordPolicyRequestDto {
	this := PasswordPolicyRequestDto{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicyRequestDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicyRequestDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *PasswordPolicyRequestDto) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *PasswordPolicyRequestDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *PasswordPolicyRequestDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *PasswordPolicyRequestDto) UnsetPassword() {
	o.Password.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *PasswordPolicyRequestDto) GetProfile() UserProfileDto {
	if o == nil || IsNil(o.Profile) {
		var ret UserProfileDto
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequestDto) GetProfileOk() (*UserProfileDto, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *PasswordPolicyRequestDto) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserProfileDto and assigns it to the Profile field.
func (o *PasswordPolicyRequestDto) SetProfile(v UserProfileDto) {
	o.Profile = &v
}

func (o PasswordPolicyRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	return toSerialize, nil
}

type NullablePasswordPolicyRequestDto struct {
	value *PasswordPolicyRequestDto
	isSet bool
}

func (v NullablePasswordPolicyRequestDto) Get() *PasswordPolicyRequestDto {
	return v.value
}

func (v *NullablePasswordPolicyRequestDto) Set(val *PasswordPolicyRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRequestDto(val *PasswordPolicyRequestDto) *NullablePasswordPolicyRequestDto {
	return &NullablePasswordPolicyRequestDto{value: val, isSet: true}
}

func (v NullablePasswordPolicyRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


