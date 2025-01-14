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

// checks if the ExtendLockOnExternalTaskDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtendLockOnExternalTaskDto{}

// ExtendLockOnExternalTaskDto struct for ExtendLockOnExternalTaskDto
type ExtendLockOnExternalTaskDto struct {
	// **Mandatory.** The ID of the worker who is performing the operation on the external task. If the task is already locked, must match the id of the worker who has most recently locked the task.
	WorkerId *string `json:"workerId,omitempty"`
	// An amount of time (in milliseconds). This is the new lock duration starting from the current moment.
	NewDuration NullableInt64 `json:"newDuration,omitempty"`
}

// NewExtendLockOnExternalTaskDto instantiates a new ExtendLockOnExternalTaskDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendLockOnExternalTaskDto() *ExtendLockOnExternalTaskDto {
	this := ExtendLockOnExternalTaskDto{}
	return &this
}

// NewExtendLockOnExternalTaskDtoWithDefaults instantiates a new ExtendLockOnExternalTaskDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendLockOnExternalTaskDtoWithDefaults() *ExtendLockOnExternalTaskDto {
	this := ExtendLockOnExternalTaskDto{}
	return &this
}

// GetWorkerId returns the WorkerId field value if set, zero value otherwise.
func (o *ExtendLockOnExternalTaskDto) GetWorkerId() string {
	if o == nil || IsNil(o.WorkerId) {
		var ret string
		return ret
	}
	return *o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendLockOnExternalTaskDto) GetWorkerIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkerId) {
		return nil, false
	}
	return o.WorkerId, true
}

// HasWorkerId returns a boolean if a field has been set.
func (o *ExtendLockOnExternalTaskDto) HasWorkerId() bool {
	if o != nil && !IsNil(o.WorkerId) {
		return true
	}

	return false
}

// SetWorkerId gets a reference to the given string and assigns it to the WorkerId field.
func (o *ExtendLockOnExternalTaskDto) SetWorkerId(v string) {
	o.WorkerId = &v
}

// GetNewDuration returns the NewDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtendLockOnExternalTaskDto) GetNewDuration() int64 {
	if o == nil || IsNil(o.NewDuration.Get()) {
		var ret int64
		return ret
	}
	return *o.NewDuration.Get()
}

// GetNewDurationOk returns a tuple with the NewDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtendLockOnExternalTaskDto) GetNewDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewDuration.Get(), o.NewDuration.IsSet()
}

// HasNewDuration returns a boolean if a field has been set.
func (o *ExtendLockOnExternalTaskDto) HasNewDuration() bool {
	if o != nil && o.NewDuration.IsSet() {
		return true
	}

	return false
}

// SetNewDuration gets a reference to the given NullableInt64 and assigns it to the NewDuration field.
func (o *ExtendLockOnExternalTaskDto) SetNewDuration(v int64) {
	o.NewDuration.Set(&v)
}
// SetNewDurationNil sets the value for NewDuration to be an explicit nil
func (o *ExtendLockOnExternalTaskDto) SetNewDurationNil() {
	o.NewDuration.Set(nil)
}

// UnsetNewDuration ensures that no value is present for NewDuration, not even an explicit nil
func (o *ExtendLockOnExternalTaskDto) UnsetNewDuration() {
	o.NewDuration.Unset()
}

func (o ExtendLockOnExternalTaskDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendLockOnExternalTaskDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WorkerId) {
		toSerialize["workerId"] = o.WorkerId
	}
	if o.NewDuration.IsSet() {
		toSerialize["newDuration"] = o.NewDuration.Get()
	}
	return toSerialize, nil
}

type NullableExtendLockOnExternalTaskDto struct {
	value *ExtendLockOnExternalTaskDto
	isSet bool
}

func (v NullableExtendLockOnExternalTaskDto) Get() *ExtendLockOnExternalTaskDto {
	return v.value
}

func (v *NullableExtendLockOnExternalTaskDto) Set(val *ExtendLockOnExternalTaskDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendLockOnExternalTaskDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendLockOnExternalTaskDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendLockOnExternalTaskDto(val *ExtendLockOnExternalTaskDto) *NullableExtendLockOnExternalTaskDto {
	return &NullableExtendLockOnExternalTaskDto{value: val, isSet: true}
}

func (v NullableExtendLockOnExternalTaskDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendLockOnExternalTaskDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


