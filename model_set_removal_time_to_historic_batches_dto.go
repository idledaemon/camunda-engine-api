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

// checks if the SetRemovalTimeToHistoricBatchesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetRemovalTimeToHistoricBatchesDto{}

// SetRemovalTimeToHistoricBatchesDto struct for SetRemovalTimeToHistoricBatchesDto
type SetRemovalTimeToHistoricBatchesDto struct {
	// The date for which the instances shall be removed. Value may not be `null`.  **Note:** Cannot be set in conjunction with `clearedRemovalTime` or `calculatedRemovalTime`.
	AbsoluteRemovalTime NullableTime `json:"absoluteRemovalTime,omitempty"`
	// Sets the removal time to `null`. Value may only be `true`, as `false` is the default behavior.  **Note:** Cannot be set in conjunction with `absoluteRemovalTime` or `calculatedRemovalTime`.
	ClearedRemovalTime NullableBool `json:"clearedRemovalTime,omitempty"`
	// The removal time is calculated based on the engine's configuration settings. Value may only be `true`, as `false` is the default behavior.  **Note:** Cannot be set in conjunction with `absoluteRemovalTime` or `clearedRemovalTime`.
	CalculatedRemovalTime NullableBool `json:"calculatedRemovalTime,omitempty"`
	// Query for the historic batches to set the removal time for.
	HistoricBatchQuery map[string]interface{} `json:"historicBatchQuery,omitempty"`
	// The ids of the historic batches to set the removal time for.
	HistoricBatchIds []string `json:"historicBatchIds,omitempty"`
}

// NewSetRemovalTimeToHistoricBatchesDto instantiates a new SetRemovalTimeToHistoricBatchesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetRemovalTimeToHistoricBatchesDto() *SetRemovalTimeToHistoricBatchesDto {
	this := SetRemovalTimeToHistoricBatchesDto{}
	return &this
}

// NewSetRemovalTimeToHistoricBatchesDtoWithDefaults instantiates a new SetRemovalTimeToHistoricBatchesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetRemovalTimeToHistoricBatchesDtoWithDefaults() *SetRemovalTimeToHistoricBatchesDto {
	this := SetRemovalTimeToHistoricBatchesDto{}
	return &this
}

// GetAbsoluteRemovalTime returns the AbsoluteRemovalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetRemovalTimeToHistoricBatchesDto) GetAbsoluteRemovalTime() time.Time {
	if o == nil || IsNil(o.AbsoluteRemovalTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AbsoluteRemovalTime.Get()
}

// GetAbsoluteRemovalTimeOk returns a tuple with the AbsoluteRemovalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetRemovalTimeToHistoricBatchesDto) GetAbsoluteRemovalTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbsoluteRemovalTime.Get(), o.AbsoluteRemovalTime.IsSet()
}

// HasAbsoluteRemovalTime returns a boolean if a field has been set.
func (o *SetRemovalTimeToHistoricBatchesDto) HasAbsoluteRemovalTime() bool {
	if o != nil && o.AbsoluteRemovalTime.IsSet() {
		return true
	}

	return false
}

// SetAbsoluteRemovalTime gets a reference to the given NullableTime and assigns it to the AbsoluteRemovalTime field.
func (o *SetRemovalTimeToHistoricBatchesDto) SetAbsoluteRemovalTime(v time.Time) {
	o.AbsoluteRemovalTime.Set(&v)
}
// SetAbsoluteRemovalTimeNil sets the value for AbsoluteRemovalTime to be an explicit nil
func (o *SetRemovalTimeToHistoricBatchesDto) SetAbsoluteRemovalTimeNil() {
	o.AbsoluteRemovalTime.Set(nil)
}

// UnsetAbsoluteRemovalTime ensures that no value is present for AbsoluteRemovalTime, not even an explicit nil
func (o *SetRemovalTimeToHistoricBatchesDto) UnsetAbsoluteRemovalTime() {
	o.AbsoluteRemovalTime.Unset()
}

// GetClearedRemovalTime returns the ClearedRemovalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetRemovalTimeToHistoricBatchesDto) GetClearedRemovalTime() bool {
	if o == nil || IsNil(o.ClearedRemovalTime.Get()) {
		var ret bool
		return ret
	}
	return *o.ClearedRemovalTime.Get()
}

// GetClearedRemovalTimeOk returns a tuple with the ClearedRemovalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetRemovalTimeToHistoricBatchesDto) GetClearedRemovalTimeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClearedRemovalTime.Get(), o.ClearedRemovalTime.IsSet()
}

// HasClearedRemovalTime returns a boolean if a field has been set.
func (o *SetRemovalTimeToHistoricBatchesDto) HasClearedRemovalTime() bool {
	if o != nil && o.ClearedRemovalTime.IsSet() {
		return true
	}

	return false
}

// SetClearedRemovalTime gets a reference to the given NullableBool and assigns it to the ClearedRemovalTime field.
func (o *SetRemovalTimeToHistoricBatchesDto) SetClearedRemovalTime(v bool) {
	o.ClearedRemovalTime.Set(&v)
}
// SetClearedRemovalTimeNil sets the value for ClearedRemovalTime to be an explicit nil
func (o *SetRemovalTimeToHistoricBatchesDto) SetClearedRemovalTimeNil() {
	o.ClearedRemovalTime.Set(nil)
}

// UnsetClearedRemovalTime ensures that no value is present for ClearedRemovalTime, not even an explicit nil
func (o *SetRemovalTimeToHistoricBatchesDto) UnsetClearedRemovalTime() {
	o.ClearedRemovalTime.Unset()
}

// GetCalculatedRemovalTime returns the CalculatedRemovalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetRemovalTimeToHistoricBatchesDto) GetCalculatedRemovalTime() bool {
	if o == nil || IsNil(o.CalculatedRemovalTime.Get()) {
		var ret bool
		return ret
	}
	return *o.CalculatedRemovalTime.Get()
}

// GetCalculatedRemovalTimeOk returns a tuple with the CalculatedRemovalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetRemovalTimeToHistoricBatchesDto) GetCalculatedRemovalTimeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CalculatedRemovalTime.Get(), o.CalculatedRemovalTime.IsSet()
}

// HasCalculatedRemovalTime returns a boolean if a field has been set.
func (o *SetRemovalTimeToHistoricBatchesDto) HasCalculatedRemovalTime() bool {
	if o != nil && o.CalculatedRemovalTime.IsSet() {
		return true
	}

	return false
}

// SetCalculatedRemovalTime gets a reference to the given NullableBool and assigns it to the CalculatedRemovalTime field.
func (o *SetRemovalTimeToHistoricBatchesDto) SetCalculatedRemovalTime(v bool) {
	o.CalculatedRemovalTime.Set(&v)
}
// SetCalculatedRemovalTimeNil sets the value for CalculatedRemovalTime to be an explicit nil
func (o *SetRemovalTimeToHistoricBatchesDto) SetCalculatedRemovalTimeNil() {
	o.CalculatedRemovalTime.Set(nil)
}

// UnsetCalculatedRemovalTime ensures that no value is present for CalculatedRemovalTime, not even an explicit nil
func (o *SetRemovalTimeToHistoricBatchesDto) UnsetCalculatedRemovalTime() {
	o.CalculatedRemovalTime.Unset()
}

// GetHistoricBatchQuery returns the HistoricBatchQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetRemovalTimeToHistoricBatchesDto) GetHistoricBatchQuery() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.HistoricBatchQuery
}

// GetHistoricBatchQueryOk returns a tuple with the HistoricBatchQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetRemovalTimeToHistoricBatchesDto) GetHistoricBatchQueryOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.HistoricBatchQuery) {
		return map[string]interface{}{}, false
	}
	return o.HistoricBatchQuery, true
}

// HasHistoricBatchQuery returns a boolean if a field has been set.
func (o *SetRemovalTimeToHistoricBatchesDto) HasHistoricBatchQuery() bool {
	if o != nil && !IsNil(o.HistoricBatchQuery) {
		return true
	}

	return false
}

// SetHistoricBatchQuery gets a reference to the given map[string]interface{} and assigns it to the HistoricBatchQuery field.
func (o *SetRemovalTimeToHistoricBatchesDto) SetHistoricBatchQuery(v map[string]interface{}) {
	o.HistoricBatchQuery = v
}

// GetHistoricBatchIds returns the HistoricBatchIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetRemovalTimeToHistoricBatchesDto) GetHistoricBatchIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.HistoricBatchIds
}

// GetHistoricBatchIdsOk returns a tuple with the HistoricBatchIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetRemovalTimeToHistoricBatchesDto) GetHistoricBatchIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.HistoricBatchIds) {
		return nil, false
	}
	return o.HistoricBatchIds, true
}

// HasHistoricBatchIds returns a boolean if a field has been set.
func (o *SetRemovalTimeToHistoricBatchesDto) HasHistoricBatchIds() bool {
	if o != nil && !IsNil(o.HistoricBatchIds) {
		return true
	}

	return false
}

// SetHistoricBatchIds gets a reference to the given []string and assigns it to the HistoricBatchIds field.
func (o *SetRemovalTimeToHistoricBatchesDto) SetHistoricBatchIds(v []string) {
	o.HistoricBatchIds = v
}

func (o SetRemovalTimeToHistoricBatchesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetRemovalTimeToHistoricBatchesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AbsoluteRemovalTime.IsSet() {
		toSerialize["absoluteRemovalTime"] = o.AbsoluteRemovalTime.Get()
	}
	if o.ClearedRemovalTime.IsSet() {
		toSerialize["clearedRemovalTime"] = o.ClearedRemovalTime.Get()
	}
	if o.CalculatedRemovalTime.IsSet() {
		toSerialize["calculatedRemovalTime"] = o.CalculatedRemovalTime.Get()
	}
	if o.HistoricBatchQuery != nil {
		toSerialize["historicBatchQuery"] = o.HistoricBatchQuery
	}
	if o.HistoricBatchIds != nil {
		toSerialize["historicBatchIds"] = o.HistoricBatchIds
	}
	return toSerialize, nil
}

type NullableSetRemovalTimeToHistoricBatchesDto struct {
	value *SetRemovalTimeToHistoricBatchesDto
	isSet bool
}

func (v NullableSetRemovalTimeToHistoricBatchesDto) Get() *SetRemovalTimeToHistoricBatchesDto {
	return v.value
}

func (v *NullableSetRemovalTimeToHistoricBatchesDto) Set(val *SetRemovalTimeToHistoricBatchesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetRemovalTimeToHistoricBatchesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetRemovalTimeToHistoricBatchesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetRemovalTimeToHistoricBatchesDto(val *SetRemovalTimeToHistoricBatchesDto) *NullableSetRemovalTimeToHistoricBatchesDto {
	return &NullableSetRemovalTimeToHistoricBatchesDto{value: val, isSet: true}
}

func (v NullableSetRemovalTimeToHistoricBatchesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetRemovalTimeToHistoricBatchesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


