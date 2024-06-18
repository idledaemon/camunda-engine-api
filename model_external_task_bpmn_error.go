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

// checks if the ExternalTaskBpmnError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalTaskBpmnError{}

// ExternalTaskBpmnError struct for ExternalTaskBpmnError
type ExternalTaskBpmnError struct {
	// An error code that indicates the predefined error. It is used to identify the BPMN error handler.
	ErrorCode NullableString `json:"errorCode,omitempty"`
	// An error message that describes the error.
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	// A JSON object containing variable key-value pairs.
	Variables map[string]VariableValueDto `json:"variables,omitempty"`
	// The id of the worker that reports the failure. Must match the id of the worker who has most recently locked the task.
	WorkerId NullableString `json:"workerId,omitempty"`
}

// NewExternalTaskBpmnError instantiates a new ExternalTaskBpmnError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTaskBpmnError() *ExternalTaskBpmnError {
	this := ExternalTaskBpmnError{}
	return &this
}

// NewExternalTaskBpmnErrorWithDefaults instantiates a new ExternalTaskBpmnError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTaskBpmnErrorWithDefaults() *ExternalTaskBpmnError {
	this := ExternalTaskBpmnError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskBpmnError) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskBpmnError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ExternalTaskBpmnError) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *ExternalTaskBpmnError) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *ExternalTaskBpmnError) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *ExternalTaskBpmnError) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskBpmnError) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskBpmnError) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ExternalTaskBpmnError) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *ExternalTaskBpmnError) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *ExternalTaskBpmnError) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *ExternalTaskBpmnError) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskBpmnError) GetVariables() map[string]VariableValueDto {
	if o == nil {
		var ret map[string]VariableValueDto
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskBpmnError) GetVariablesOk() (*map[string]VariableValueDto, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *ExternalTaskBpmnError) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]VariableValueDto and assigns it to the Variables field.
func (o *ExternalTaskBpmnError) SetVariables(v map[string]VariableValueDto) {
	o.Variables = v
}

// GetWorkerId returns the WorkerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskBpmnError) GetWorkerId() string {
	if o == nil || IsNil(o.WorkerId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkerId.Get()
}

// GetWorkerIdOk returns a tuple with the WorkerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskBpmnError) GetWorkerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkerId.Get(), o.WorkerId.IsSet()
}

// HasWorkerId returns a boolean if a field has been set.
func (o *ExternalTaskBpmnError) HasWorkerId() bool {
	if o != nil && o.WorkerId.IsSet() {
		return true
	}

	return false
}

// SetWorkerId gets a reference to the given NullableString and assigns it to the WorkerId field.
func (o *ExternalTaskBpmnError) SetWorkerId(v string) {
	o.WorkerId.Set(&v)
}
// SetWorkerIdNil sets the value for WorkerId to be an explicit nil
func (o *ExternalTaskBpmnError) SetWorkerIdNil() {
	o.WorkerId.Set(nil)
}

// UnsetWorkerId ensures that no value is present for WorkerId, not even an explicit nil
func (o *ExternalTaskBpmnError) UnsetWorkerId() {
	o.WorkerId.Unset()
}

func (o ExternalTaskBpmnError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalTaskBpmnError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode.IsSet() {
		toSerialize["errorCode"] = o.ErrorCode.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	if o.WorkerId.IsSet() {
		toSerialize["workerId"] = o.WorkerId.Get()
	}
	return toSerialize, nil
}

type NullableExternalTaskBpmnError struct {
	value *ExternalTaskBpmnError
	isSet bool
}

func (v NullableExternalTaskBpmnError) Get() *ExternalTaskBpmnError {
	return v.value
}

func (v *NullableExternalTaskBpmnError) Set(val *ExternalTaskBpmnError) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTaskBpmnError) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTaskBpmnError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTaskBpmnError(val *ExternalTaskBpmnError) *NullableExternalTaskBpmnError {
	return &NullableExternalTaskBpmnError{value: val, isSet: true}
}

func (v NullableExternalTaskBpmnError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTaskBpmnError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

