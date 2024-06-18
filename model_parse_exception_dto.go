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

// checks if the ParseExceptionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParseExceptionDto{}

// ParseExceptionDto struct for ParseExceptionDto
type ParseExceptionDto struct {
	// An exception class indicating the occurred error.
	Type NullableString `json:"type,omitempty"`
	// A detailed message of the error.
	Message NullableString `json:"message,omitempty"`
	// The code allows your client application to identify the error in an automated fashion. You can look up the meaning of all built-in codes and learn how to add custom codes in the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/error-handling/#exception-codes).
	Code *float32 `json:"code,omitempty"`
	// A JSON Object containing list of errors and warnings occurred during deployment.
	Details map[string]ResourceReportDto `json:"details,omitempty"`
}

// NewParseExceptionDto instantiates a new ParseExceptionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParseExceptionDto() *ParseExceptionDto {
	this := ParseExceptionDto{}
	return &this
}

// NewParseExceptionDtoWithDefaults instantiates a new ParseExceptionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParseExceptionDtoWithDefaults() *ParseExceptionDto {
	this := ParseExceptionDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParseExceptionDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseExceptionDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ParseExceptionDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ParseExceptionDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ParseExceptionDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ParseExceptionDto) UnsetType() {
	o.Type.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParseExceptionDto) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseExceptionDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ParseExceptionDto) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ParseExceptionDto) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *ParseExceptionDto) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ParseExceptionDto) UnsetMessage() {
	o.Message.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ParseExceptionDto) GetCode() float32 {
	if o == nil || IsNil(o.Code) {
		var ret float32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParseExceptionDto) GetCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ParseExceptionDto) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given float32 and assigns it to the Code field.
func (o *ParseExceptionDto) SetCode(v float32) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParseExceptionDto) GetDetails() map[string]ResourceReportDto {
	if o == nil {
		var ret map[string]ResourceReportDto
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseExceptionDto) GetDetailsOk() (*map[string]ResourceReportDto, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return &o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ParseExceptionDto) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]ResourceReportDto and assigns it to the Details field.
func (o *ParseExceptionDto) SetDetails(v map[string]ResourceReportDto) {
	o.Details = v
}

func (o ParseExceptionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParseExceptionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableParseExceptionDto struct {
	value *ParseExceptionDto
	isSet bool
}

func (v NullableParseExceptionDto) Get() *ParseExceptionDto {
	return v.value
}

func (v *NullableParseExceptionDto) Set(val *ParseExceptionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableParseExceptionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableParseExceptionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParseExceptionDto(val *ParseExceptionDto) *NullableParseExceptionDto {
	return &NullableParseExceptionDto{value: val, isSet: true}
}

func (v NullableParseExceptionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParseExceptionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

