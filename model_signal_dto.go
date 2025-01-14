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

// checks if the SignalDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignalDto{}

// SignalDto struct for SignalDto
type SignalDto struct {
	// The name of the signal to deliver.  **Note**: This property is mandatory.
	Name *string `json:"name,omitempty"`
	// Optionally specifies a single execution which is notified by the signal.  **Note**: If no execution id is defined the signal is broadcasted to all subscribed handlers. 
	ExecutionId NullableString `json:"executionId,omitempty"`
	// A JSON object containing variable key-value pairs. Each key is a variable name and each value a JSON variable value object.
	Variables map[string]VariableValueDto `json:"variables,omitempty"`
	// Specifies a tenant to deliver the signal. The signal can only be received on executions or process definitions which belongs to the given tenant.  **Note**: Cannot be used in combination with executionId.
	TenantId NullableString `json:"tenantId,omitempty"`
	// If true the signal can only be received on executions or process definitions which belongs to no tenant. Value may not be false as this is the default behavior.  **Note**: Cannot be used in combination with `executionId`.
	WithoutTenantId NullableBool `json:"withoutTenantId,omitempty"`
}

// NewSignalDto instantiates a new SignalDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalDto() *SignalDto {
	this := SignalDto{}
	return &this
}

// NewSignalDtoWithDefaults instantiates a new SignalDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalDtoWithDefaults() *SignalDto {
	this := SignalDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SignalDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SignalDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SignalDto) SetName(v string) {
	o.Name = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalDto) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionId.Get()
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalDto) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionId.Get(), o.ExecutionId.IsSet()
}

// HasExecutionId returns a boolean if a field has been set.
func (o *SignalDto) HasExecutionId() bool {
	if o != nil && o.ExecutionId.IsSet() {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given NullableString and assigns it to the ExecutionId field.
func (o *SignalDto) SetExecutionId(v string) {
	o.ExecutionId.Set(&v)
}
// SetExecutionIdNil sets the value for ExecutionId to be an explicit nil
func (o *SignalDto) SetExecutionIdNil() {
	o.ExecutionId.Set(nil)
}

// UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
func (o *SignalDto) UnsetExecutionId() {
	o.ExecutionId.Unset()
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalDto) GetVariables() map[string]VariableValueDto {
	if o == nil {
		var ret map[string]VariableValueDto
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalDto) GetVariablesOk() (*map[string]VariableValueDto, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *SignalDto) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]VariableValueDto and assigns it to the Variables field.
func (o *SignalDto) SetVariables(v map[string]VariableValueDto) {
	o.Variables = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *SignalDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *SignalDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *SignalDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *SignalDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetWithoutTenantId returns the WithoutTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalDto) GetWithoutTenantId() bool {
	if o == nil || IsNil(o.WithoutTenantId.Get()) {
		var ret bool
		return ret
	}
	return *o.WithoutTenantId.Get()
}

// GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalDto) GetWithoutTenantIdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithoutTenantId.Get(), o.WithoutTenantId.IsSet()
}

// HasWithoutTenantId returns a boolean if a field has been set.
func (o *SignalDto) HasWithoutTenantId() bool {
	if o != nil && o.WithoutTenantId.IsSet() {
		return true
	}

	return false
}

// SetWithoutTenantId gets a reference to the given NullableBool and assigns it to the WithoutTenantId field.
func (o *SignalDto) SetWithoutTenantId(v bool) {
	o.WithoutTenantId.Set(&v)
}
// SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil
func (o *SignalDto) SetWithoutTenantIdNil() {
	o.WithoutTenantId.Set(nil)
}

// UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
func (o *SignalDto) UnsetWithoutTenantId() {
	o.WithoutTenantId.Unset()
}

func (o SignalDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignalDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.ExecutionId.IsSet() {
		toSerialize["executionId"] = o.ExecutionId.Get()
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.WithoutTenantId.IsSet() {
		toSerialize["withoutTenantId"] = o.WithoutTenantId.Get()
	}
	return toSerialize, nil
}

type NullableSignalDto struct {
	value *SignalDto
	isSet bool
}

func (v NullableSignalDto) Get() *SignalDto {
	return v.value
}

func (v *NullableSignalDto) Set(val *SignalDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalDto(val *SignalDto) *NullableSignalDto {
	return &NullableSignalDto{value: val, isSet: true}
}

func (v NullableSignalDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


