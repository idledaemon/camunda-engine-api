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

// checks if the EvaluationConditionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvaluationConditionDto{}

// EvaluationConditionDto struct for EvaluationConditionDto
type EvaluationConditionDto struct {
	// A map of variables which are used for evaluation of the conditions and are injected into the process instances which have been triggered. Each key is a variable name and each value a JSON variable value object with the following properties.
	Variables map[string]VariableValueDto `json:"variables,omitempty"`
	// Used for the process instances that have been triggered after the evaluation.
	BusinessKey NullableString `json:"businessKey,omitempty"`
	// Used to evaluate a condition for a tenant with the given id. Will only evaluate conditions of process definitions which belong to the tenant.
	TenantId NullableString `json:"tenantId,omitempty"`
	// A Boolean value that indicates whether the conditions should only be evaluated of process definitions which belong to no tenant or not. Value may only be true, as false is the default behavior.
	WithoutTenantId NullableBool `json:"withoutTenantId,omitempty"`
	// Used to evaluate conditions of the process definition with the given id.
	ProcessDefinitionId NullableString `json:"processDefinitionId,omitempty"`
}

// NewEvaluationConditionDto instantiates a new EvaluationConditionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvaluationConditionDto() *EvaluationConditionDto {
	this := EvaluationConditionDto{}
	return &this
}

// NewEvaluationConditionDtoWithDefaults instantiates a new EvaluationConditionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvaluationConditionDtoWithDefaults() *EvaluationConditionDto {
	this := EvaluationConditionDto{}
	return &this
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EvaluationConditionDto) GetVariables() map[string]VariableValueDto {
	if o == nil {
		var ret map[string]VariableValueDto
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EvaluationConditionDto) GetVariablesOk() (*map[string]VariableValueDto, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *EvaluationConditionDto) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]VariableValueDto and assigns it to the Variables field.
func (o *EvaluationConditionDto) SetVariables(v map[string]VariableValueDto) {
	o.Variables = v
}

// GetBusinessKey returns the BusinessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EvaluationConditionDto) GetBusinessKey() string {
	if o == nil || IsNil(o.BusinessKey.Get()) {
		var ret string
		return ret
	}
	return *o.BusinessKey.Get()
}

// GetBusinessKeyOk returns a tuple with the BusinessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EvaluationConditionDto) GetBusinessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessKey.Get(), o.BusinessKey.IsSet()
}

// HasBusinessKey returns a boolean if a field has been set.
func (o *EvaluationConditionDto) HasBusinessKey() bool {
	if o != nil && o.BusinessKey.IsSet() {
		return true
	}

	return false
}

// SetBusinessKey gets a reference to the given NullableString and assigns it to the BusinessKey field.
func (o *EvaluationConditionDto) SetBusinessKey(v string) {
	o.BusinessKey.Set(&v)
}
// SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil
func (o *EvaluationConditionDto) SetBusinessKeyNil() {
	o.BusinessKey.Set(nil)
}

// UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
func (o *EvaluationConditionDto) UnsetBusinessKey() {
	o.BusinessKey.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EvaluationConditionDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EvaluationConditionDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *EvaluationConditionDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *EvaluationConditionDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *EvaluationConditionDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *EvaluationConditionDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetWithoutTenantId returns the WithoutTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EvaluationConditionDto) GetWithoutTenantId() bool {
	if o == nil || IsNil(o.WithoutTenantId.Get()) {
		var ret bool
		return ret
	}
	return *o.WithoutTenantId.Get()
}

// GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EvaluationConditionDto) GetWithoutTenantIdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithoutTenantId.Get(), o.WithoutTenantId.IsSet()
}

// HasWithoutTenantId returns a boolean if a field has been set.
func (o *EvaluationConditionDto) HasWithoutTenantId() bool {
	if o != nil && o.WithoutTenantId.IsSet() {
		return true
	}

	return false
}

// SetWithoutTenantId gets a reference to the given NullableBool and assigns it to the WithoutTenantId field.
func (o *EvaluationConditionDto) SetWithoutTenantId(v bool) {
	o.WithoutTenantId.Set(&v)
}
// SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil
func (o *EvaluationConditionDto) SetWithoutTenantIdNil() {
	o.WithoutTenantId.Set(nil)
}

// UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
func (o *EvaluationConditionDto) UnsetWithoutTenantId() {
	o.WithoutTenantId.Unset()
}

// GetProcessDefinitionId returns the ProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EvaluationConditionDto) GetProcessDefinitionId() string {
	if o == nil || IsNil(o.ProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionId.Get()
}

// GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EvaluationConditionDto) GetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionId.Get(), o.ProcessDefinitionId.IsSet()
}

// HasProcessDefinitionId returns a boolean if a field has been set.
func (o *EvaluationConditionDto) HasProcessDefinitionId() bool {
	if o != nil && o.ProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionId gets a reference to the given NullableString and assigns it to the ProcessDefinitionId field.
func (o *EvaluationConditionDto) SetProcessDefinitionId(v string) {
	o.ProcessDefinitionId.Set(&v)
}
// SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil
func (o *EvaluationConditionDto) SetProcessDefinitionIdNil() {
	o.ProcessDefinitionId.Set(nil)
}

// UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
func (o *EvaluationConditionDto) UnsetProcessDefinitionId() {
	o.ProcessDefinitionId.Unset()
}

func (o EvaluationConditionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvaluationConditionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	if o.BusinessKey.IsSet() {
		toSerialize["businessKey"] = o.BusinessKey.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.WithoutTenantId.IsSet() {
		toSerialize["withoutTenantId"] = o.WithoutTenantId.Get()
	}
	if o.ProcessDefinitionId.IsSet() {
		toSerialize["processDefinitionId"] = o.ProcessDefinitionId.Get()
	}
	return toSerialize, nil
}

type NullableEvaluationConditionDto struct {
	value *EvaluationConditionDto
	isSet bool
}

func (v NullableEvaluationConditionDto) Get() *EvaluationConditionDto {
	return v.value
}

func (v *NullableEvaluationConditionDto) Set(val *EvaluationConditionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEvaluationConditionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEvaluationConditionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvaluationConditionDto(val *EvaluationConditionDto) *NullableEvaluationConditionDto {
	return &NullableEvaluationConditionDto{value: val, isSet: true}
}

func (v NullableEvaluationConditionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvaluationConditionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


