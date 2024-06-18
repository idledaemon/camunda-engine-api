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

// checks if the MetricsIntervalResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsIntervalResultDto{}

// MetricsIntervalResultDto struct for MetricsIntervalResultDto
type MetricsIntervalResultDto struct {
	// The interval timestamp.
	Timestamp NullableTime `json:"timestamp,omitempty"`
	// The name of the metric.
	Name NullableString `json:"name,omitempty"`
	// The reporter of the metric. `null` if the metrics are aggregated by reporter.
	Reporter NullableString `json:"reporter,omitempty"`
	// The value of the metric aggregated by the interval.
	Value NullableInt64 `json:"value,omitempty"`
}

// NewMetricsIntervalResultDto instantiates a new MetricsIntervalResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsIntervalResultDto() *MetricsIntervalResultDto {
	this := MetricsIntervalResultDto{}
	return &this
}

// NewMetricsIntervalResultDtoWithDefaults instantiates a new MetricsIntervalResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsIntervalResultDtoWithDefaults() *MetricsIntervalResultDto {
	this := MetricsIntervalResultDto{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsIntervalResultDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsIntervalResultDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MetricsIntervalResultDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *MetricsIntervalResultDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *MetricsIntervalResultDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *MetricsIntervalResultDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsIntervalResultDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsIntervalResultDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MetricsIntervalResultDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MetricsIntervalResultDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MetricsIntervalResultDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MetricsIntervalResultDto) UnsetName() {
	o.Name.Unset()
}

// GetReporter returns the Reporter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsIntervalResultDto) GetReporter() string {
	if o == nil || IsNil(o.Reporter.Get()) {
		var ret string
		return ret
	}
	return *o.Reporter.Get()
}

// GetReporterOk returns a tuple with the Reporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsIntervalResultDto) GetReporterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reporter.Get(), o.Reporter.IsSet()
}

// HasReporter returns a boolean if a field has been set.
func (o *MetricsIntervalResultDto) HasReporter() bool {
	if o != nil && o.Reporter.IsSet() {
		return true
	}

	return false
}

// SetReporter gets a reference to the given NullableString and assigns it to the Reporter field.
func (o *MetricsIntervalResultDto) SetReporter(v string) {
	o.Reporter.Set(&v)
}
// SetReporterNil sets the value for Reporter to be an explicit nil
func (o *MetricsIntervalResultDto) SetReporterNil() {
	o.Reporter.Set(nil)
}

// UnsetReporter ensures that no value is present for Reporter, not even an explicit nil
func (o *MetricsIntervalResultDto) UnsetReporter() {
	o.Reporter.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsIntervalResultDto) GetValue() int64 {
	if o == nil || IsNil(o.Value.Get()) {
		var ret int64
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsIntervalResultDto) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *MetricsIntervalResultDto) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableInt64 and assigns it to the Value field.
func (o *MetricsIntervalResultDto) SetValue(v int64) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *MetricsIntervalResultDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *MetricsIntervalResultDto) UnsetValue() {
	o.Value.Unset()
}

func (o MetricsIntervalResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsIntervalResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Reporter.IsSet() {
		toSerialize["reporter"] = o.Reporter.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableMetricsIntervalResultDto struct {
	value *MetricsIntervalResultDto
	isSet bool
}

func (v NullableMetricsIntervalResultDto) Get() *MetricsIntervalResultDto {
	return v.value
}

func (v *NullableMetricsIntervalResultDto) Set(val *MetricsIntervalResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsIntervalResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsIntervalResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsIntervalResultDto(val *MetricsIntervalResultDto) *NullableMetricsIntervalResultDto {
	return &NullableMetricsIntervalResultDto{value: val, isSet: true}
}

func (v NullableMetricsIntervalResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsIntervalResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


