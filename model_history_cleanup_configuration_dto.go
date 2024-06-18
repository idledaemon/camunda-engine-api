/*
Camunda Platform REST API

OpenApi Spec for Camunda Platform REST API.

API version: 7.21.2-ee
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package camundarestgo

import (
	"encoding/json"
	"time"
)

// checks if the HistoryCleanupConfigurationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryCleanupConfigurationDto{}

// HistoryCleanupConfigurationDto struct for HistoryCleanupConfigurationDto
type HistoryCleanupConfigurationDto struct {
	// Start time of the current or next batch window. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`.
	BatchWindowStartTime NullableTime `json:"batchWindowStartTime,omitempty"`
	// End time of the current or next batch window. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`.
	BatchWindowEndTime NullableTime `json:"batchWindowEndTime,omitempty"`
	// Indicates whether the engine node participates in history cleanup or not. The default is `true`. Participation can be disabled via [Process Engine Configuration](https://docs.camunda.org/manual/7.21/reference/deployment-descriptors/tags/process-engine/#history-cleanup-enabled).  For more details, see [Cleanup Execution Participation per Node](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#cleanup-execution-participation-per-node).
	Enabled NullableBool `json:"enabled,omitempty"`
}

// NewHistoryCleanupConfigurationDto instantiates a new HistoryCleanupConfigurationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryCleanupConfigurationDto() *HistoryCleanupConfigurationDto {
	this := HistoryCleanupConfigurationDto{}
	return &this
}

// NewHistoryCleanupConfigurationDtoWithDefaults instantiates a new HistoryCleanupConfigurationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryCleanupConfigurationDtoWithDefaults() *HistoryCleanupConfigurationDto {
	this := HistoryCleanupConfigurationDto{}
	return &this
}

// GetBatchWindowStartTime returns the BatchWindowStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryCleanupConfigurationDto) GetBatchWindowStartTime() time.Time {
	if o == nil || IsNil(o.BatchWindowStartTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.BatchWindowStartTime.Get()
}

// GetBatchWindowStartTimeOk returns a tuple with the BatchWindowStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryCleanupConfigurationDto) GetBatchWindowStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.BatchWindowStartTime.Get(), o.BatchWindowStartTime.IsSet()
}

// HasBatchWindowStartTime returns a boolean if a field has been set.
func (o *HistoryCleanupConfigurationDto) HasBatchWindowStartTime() bool {
	if o != nil && o.BatchWindowStartTime.IsSet() {
		return true
	}

	return false
}

// SetBatchWindowStartTime gets a reference to the given NullableTime and assigns it to the BatchWindowStartTime field.
func (o *HistoryCleanupConfigurationDto) SetBatchWindowStartTime(v time.Time) {
	o.BatchWindowStartTime.Set(&v)
}
// SetBatchWindowStartTimeNil sets the value for BatchWindowStartTime to be an explicit nil
func (o *HistoryCleanupConfigurationDto) SetBatchWindowStartTimeNil() {
	o.BatchWindowStartTime.Set(nil)
}

// UnsetBatchWindowStartTime ensures that no value is present for BatchWindowStartTime, not even an explicit nil
func (o *HistoryCleanupConfigurationDto) UnsetBatchWindowStartTime() {
	o.BatchWindowStartTime.Unset()
}

// GetBatchWindowEndTime returns the BatchWindowEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryCleanupConfigurationDto) GetBatchWindowEndTime() time.Time {
	if o == nil || IsNil(o.BatchWindowEndTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.BatchWindowEndTime.Get()
}

// GetBatchWindowEndTimeOk returns a tuple with the BatchWindowEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryCleanupConfigurationDto) GetBatchWindowEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.BatchWindowEndTime.Get(), o.BatchWindowEndTime.IsSet()
}

// HasBatchWindowEndTime returns a boolean if a field has been set.
func (o *HistoryCleanupConfigurationDto) HasBatchWindowEndTime() bool {
	if o != nil && o.BatchWindowEndTime.IsSet() {
		return true
	}

	return false
}

// SetBatchWindowEndTime gets a reference to the given NullableTime and assigns it to the BatchWindowEndTime field.
func (o *HistoryCleanupConfigurationDto) SetBatchWindowEndTime(v time.Time) {
	o.BatchWindowEndTime.Set(&v)
}
// SetBatchWindowEndTimeNil sets the value for BatchWindowEndTime to be an explicit nil
func (o *HistoryCleanupConfigurationDto) SetBatchWindowEndTimeNil() {
	o.BatchWindowEndTime.Set(nil)
}

// UnsetBatchWindowEndTime ensures that no value is present for BatchWindowEndTime, not even an explicit nil
func (o *HistoryCleanupConfigurationDto) UnsetBatchWindowEndTime() {
	o.BatchWindowEndTime.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryCleanupConfigurationDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryCleanupConfigurationDto) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *HistoryCleanupConfigurationDto) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *HistoryCleanupConfigurationDto) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *HistoryCleanupConfigurationDto) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *HistoryCleanupConfigurationDto) UnsetEnabled() {
	o.Enabled.Unset()
}

func (o HistoryCleanupConfigurationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryCleanupConfigurationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BatchWindowStartTime.IsSet() {
		toSerialize["batchWindowStartTime"] = o.BatchWindowStartTime.Get()
	}
	if o.BatchWindowEndTime.IsSet() {
		toSerialize["batchWindowEndTime"] = o.BatchWindowEndTime.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	return toSerialize, nil
}

type NullableHistoryCleanupConfigurationDto struct {
	value *HistoryCleanupConfigurationDto
	isSet bool
}

func (v NullableHistoryCleanupConfigurationDto) Get() *HistoryCleanupConfigurationDto {
	return v.value
}

func (v *NullableHistoryCleanupConfigurationDto) Set(val *HistoryCleanupConfigurationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryCleanupConfigurationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryCleanupConfigurationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryCleanupConfigurationDto(val *HistoryCleanupConfigurationDto) *NullableHistoryCleanupConfigurationDto {
	return &NullableHistoryCleanupConfigurationDto{value: val, isSet: true}
}

func (v NullableHistoryCleanupConfigurationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryCleanupConfigurationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

