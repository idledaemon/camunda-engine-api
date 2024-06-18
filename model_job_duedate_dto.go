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

// checks if the JobDuedateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobDuedateDto{}

// JobDuedateDto struct for JobDuedateDto
type JobDuedateDto struct {
	// The date to set when the job has the next execution.
	Duedate NullableTime `json:"duedate,omitempty"`
	// A boolean value to indicate if modifications to the due date should cascade to subsequent jobs. (e.g. Modify the due date of a timer by +15 minutes. This flag indicates if a +15 minutes should be applied to all subsequent timers.) This flag only affects timer jobs and only works if due date is not null. Default: `false`
	Cascade NullableBool `json:"cascade,omitempty"`
}

// NewJobDuedateDto instantiates a new JobDuedateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDuedateDto() *JobDuedateDto {
	this := JobDuedateDto{}
	return &this
}

// NewJobDuedateDtoWithDefaults instantiates a new JobDuedateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDuedateDtoWithDefaults() *JobDuedateDto {
	this := JobDuedateDto{}
	return &this
}

// GetDuedate returns the Duedate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDuedateDto) GetDuedate() time.Time {
	if o == nil || IsNil(o.Duedate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Duedate.Get()
}

// GetDuedateOk returns a tuple with the Duedate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDuedateDto) GetDuedateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duedate.Get(), o.Duedate.IsSet()
}

// HasDuedate returns a boolean if a field has been set.
func (o *JobDuedateDto) HasDuedate() bool {
	if o != nil && o.Duedate.IsSet() {
		return true
	}

	return false
}

// SetDuedate gets a reference to the given NullableTime and assigns it to the Duedate field.
func (o *JobDuedateDto) SetDuedate(v time.Time) {
	o.Duedate.Set(&v)
}
// SetDuedateNil sets the value for Duedate to be an explicit nil
func (o *JobDuedateDto) SetDuedateNil() {
	o.Duedate.Set(nil)
}

// UnsetDuedate ensures that no value is present for Duedate, not even an explicit nil
func (o *JobDuedateDto) UnsetDuedate() {
	o.Duedate.Unset()
}

// GetCascade returns the Cascade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDuedateDto) GetCascade() bool {
	if o == nil || IsNil(o.Cascade.Get()) {
		var ret bool
		return ret
	}
	return *o.Cascade.Get()
}

// GetCascadeOk returns a tuple with the Cascade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDuedateDto) GetCascadeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cascade.Get(), o.Cascade.IsSet()
}

// HasCascade returns a boolean if a field has been set.
func (o *JobDuedateDto) HasCascade() bool {
	if o != nil && o.Cascade.IsSet() {
		return true
	}

	return false
}

// SetCascade gets a reference to the given NullableBool and assigns it to the Cascade field.
func (o *JobDuedateDto) SetCascade(v bool) {
	o.Cascade.Set(&v)
}
// SetCascadeNil sets the value for Cascade to be an explicit nil
func (o *JobDuedateDto) SetCascadeNil() {
	o.Cascade.Set(nil)
}

// UnsetCascade ensures that no value is present for Cascade, not even an explicit nil
func (o *JobDuedateDto) UnsetCascade() {
	o.Cascade.Unset()
}

func (o JobDuedateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobDuedateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Duedate.IsSet() {
		toSerialize["duedate"] = o.Duedate.Get()
	}
	if o.Cascade.IsSet() {
		toSerialize["cascade"] = o.Cascade.Get()
	}
	return toSerialize, nil
}

type NullableJobDuedateDto struct {
	value *JobDuedateDto
	isSet bool
}

func (v NullableJobDuedateDto) Get() *JobDuedateDto {
	return v.value
}

func (v *NullableJobDuedateDto) Set(val *JobDuedateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDuedateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDuedateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDuedateDto(val *JobDuedateDto) *NullableJobDuedateDto {
	return &NullableJobDuedateDto{value: val, isSet: true}
}

func (v NullableJobDuedateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDuedateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

