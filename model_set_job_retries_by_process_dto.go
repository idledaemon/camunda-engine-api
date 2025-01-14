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

// checks if the SetJobRetriesByProcessDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetJobRetriesByProcessDto{}

// SetJobRetriesByProcessDto struct for SetJobRetriesByProcessDto
type SetJobRetriesByProcessDto struct {
	// A list of job ids to set retries for.
	JobIds []string `json:"jobIds,omitempty"`
	JobQuery *JobQueryDto `json:"jobQuery,omitempty"`
	// The due date to set for the job. A due date indicates when this job is ready for execution. Jobs with due dates in the past will be scheduled for execution.
	DueDate NullableTime `json:"dueDate,omitempty"`
	// The number of retries to set for the resource.  Must be >= 0. If this is 0, an incident is created and the task, or job, cannot be fetched, or acquired anymore unless the retries are increased again. Can not be null.
	Retries NullableInt32 `json:"retries,omitempty"`
	// A list of process instance ids to fetch jobs, for which retries will be set.
	ProcessInstances []string `json:"processInstances,omitempty"`
	ProcessInstanceQuery *ProcessInstanceQueryDto `json:"processInstanceQuery,omitempty"`
	HistoricProcessInstanceQuery *HistoricProcessInstanceQueryDto `json:"historicProcessInstanceQuery,omitempty"`
}

// NewSetJobRetriesByProcessDto instantiates a new SetJobRetriesByProcessDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetJobRetriesByProcessDto() *SetJobRetriesByProcessDto {
	this := SetJobRetriesByProcessDto{}
	return &this
}

// NewSetJobRetriesByProcessDtoWithDefaults instantiates a new SetJobRetriesByProcessDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetJobRetriesByProcessDtoWithDefaults() *SetJobRetriesByProcessDto {
	this := SetJobRetriesByProcessDto{}
	return &this
}

// GetJobIds returns the JobIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetJobRetriesByProcessDto) GetJobIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.JobIds
}

// GetJobIdsOk returns a tuple with the JobIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetJobRetriesByProcessDto) GetJobIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.JobIds) {
		return nil, false
	}
	return o.JobIds, true
}

// HasJobIds returns a boolean if a field has been set.
func (o *SetJobRetriesByProcessDto) HasJobIds() bool {
	if o != nil && !IsNil(o.JobIds) {
		return true
	}

	return false
}

// SetJobIds gets a reference to the given []string and assigns it to the JobIds field.
func (o *SetJobRetriesByProcessDto) SetJobIds(v []string) {
	o.JobIds = v
}

// GetJobQuery returns the JobQuery field value if set, zero value otherwise.
func (o *SetJobRetriesByProcessDto) GetJobQuery() JobQueryDto {
	if o == nil || IsNil(o.JobQuery) {
		var ret JobQueryDto
		return ret
	}
	return *o.JobQuery
}

// GetJobQueryOk returns a tuple with the JobQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetJobRetriesByProcessDto) GetJobQueryOk() (*JobQueryDto, bool) {
	if o == nil || IsNil(o.JobQuery) {
		return nil, false
	}
	return o.JobQuery, true
}

// HasJobQuery returns a boolean if a field has been set.
func (o *SetJobRetriesByProcessDto) HasJobQuery() bool {
	if o != nil && !IsNil(o.JobQuery) {
		return true
	}

	return false
}

// SetJobQuery gets a reference to the given JobQueryDto and assigns it to the JobQuery field.
func (o *SetJobRetriesByProcessDto) SetJobQuery(v JobQueryDto) {
	o.JobQuery = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetJobRetriesByProcessDto) GetDueDate() time.Time {
	if o == nil || IsNil(o.DueDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetJobRetriesByProcessDto) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *SetJobRetriesByProcessDto) HasDueDate() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableTime and assigns it to the DueDate field.
func (o *SetJobRetriesByProcessDto) SetDueDate(v time.Time) {
	o.DueDate.Set(&v)
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *SetJobRetriesByProcessDto) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *SetJobRetriesByProcessDto) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetRetries returns the Retries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetJobRetriesByProcessDto) GetRetries() int32 {
	if o == nil || IsNil(o.Retries.Get()) {
		var ret int32
		return ret
	}
	return *o.Retries.Get()
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetJobRetriesByProcessDto) GetRetriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Retries.Get(), o.Retries.IsSet()
}

// HasRetries returns a boolean if a field has been set.
func (o *SetJobRetriesByProcessDto) HasRetries() bool {
	if o != nil && o.Retries.IsSet() {
		return true
	}

	return false
}

// SetRetries gets a reference to the given NullableInt32 and assigns it to the Retries field.
func (o *SetJobRetriesByProcessDto) SetRetries(v int32) {
	o.Retries.Set(&v)
}
// SetRetriesNil sets the value for Retries to be an explicit nil
func (o *SetJobRetriesByProcessDto) SetRetriesNil() {
	o.Retries.Set(nil)
}

// UnsetRetries ensures that no value is present for Retries, not even an explicit nil
func (o *SetJobRetriesByProcessDto) UnsetRetries() {
	o.Retries.Unset()
}

// GetProcessInstances returns the ProcessInstances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetJobRetriesByProcessDto) GetProcessInstances() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProcessInstances
}

// GetProcessInstancesOk returns a tuple with the ProcessInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetJobRetriesByProcessDto) GetProcessInstancesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessInstances) {
		return nil, false
	}
	return o.ProcessInstances, true
}

// HasProcessInstances returns a boolean if a field has been set.
func (o *SetJobRetriesByProcessDto) HasProcessInstances() bool {
	if o != nil && !IsNil(o.ProcessInstances) {
		return true
	}

	return false
}

// SetProcessInstances gets a reference to the given []string and assigns it to the ProcessInstances field.
func (o *SetJobRetriesByProcessDto) SetProcessInstances(v []string) {
	o.ProcessInstances = v
}

// GetProcessInstanceQuery returns the ProcessInstanceQuery field value if set, zero value otherwise.
func (o *SetJobRetriesByProcessDto) GetProcessInstanceQuery() ProcessInstanceQueryDto {
	if o == nil || IsNil(o.ProcessInstanceQuery) {
		var ret ProcessInstanceQueryDto
		return ret
	}
	return *o.ProcessInstanceQuery
}

// GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetJobRetriesByProcessDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool) {
	if o == nil || IsNil(o.ProcessInstanceQuery) {
		return nil, false
	}
	return o.ProcessInstanceQuery, true
}

// HasProcessInstanceQuery returns a boolean if a field has been set.
func (o *SetJobRetriesByProcessDto) HasProcessInstanceQuery() bool {
	if o != nil && !IsNil(o.ProcessInstanceQuery) {
		return true
	}

	return false
}

// SetProcessInstanceQuery gets a reference to the given ProcessInstanceQueryDto and assigns it to the ProcessInstanceQuery field.
func (o *SetJobRetriesByProcessDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto) {
	o.ProcessInstanceQuery = &v
}

// GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field value if set, zero value otherwise.
func (o *SetJobRetriesByProcessDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto {
	if o == nil || IsNil(o.HistoricProcessInstanceQuery) {
		var ret HistoricProcessInstanceQueryDto
		return ret
	}
	return *o.HistoricProcessInstanceQuery
}

// GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetJobRetriesByProcessDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool) {
	if o == nil || IsNil(o.HistoricProcessInstanceQuery) {
		return nil, false
	}
	return o.HistoricProcessInstanceQuery, true
}

// HasHistoricProcessInstanceQuery returns a boolean if a field has been set.
func (o *SetJobRetriesByProcessDto) HasHistoricProcessInstanceQuery() bool {
	if o != nil && !IsNil(o.HistoricProcessInstanceQuery) {
		return true
	}

	return false
}

// SetHistoricProcessInstanceQuery gets a reference to the given HistoricProcessInstanceQueryDto and assigns it to the HistoricProcessInstanceQuery field.
func (o *SetJobRetriesByProcessDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto) {
	o.HistoricProcessInstanceQuery = &v
}

func (o SetJobRetriesByProcessDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetJobRetriesByProcessDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.JobIds != nil {
		toSerialize["jobIds"] = o.JobIds
	}
	if !IsNil(o.JobQuery) {
		toSerialize["jobQuery"] = o.JobQuery
	}
	if o.DueDate.IsSet() {
		toSerialize["dueDate"] = o.DueDate.Get()
	}
	if o.Retries.IsSet() {
		toSerialize["retries"] = o.Retries.Get()
	}
	if o.ProcessInstances != nil {
		toSerialize["processInstances"] = o.ProcessInstances
	}
	if !IsNil(o.ProcessInstanceQuery) {
		toSerialize["processInstanceQuery"] = o.ProcessInstanceQuery
	}
	if !IsNil(o.HistoricProcessInstanceQuery) {
		toSerialize["historicProcessInstanceQuery"] = o.HistoricProcessInstanceQuery
	}
	return toSerialize, nil
}

type NullableSetJobRetriesByProcessDto struct {
	value *SetJobRetriesByProcessDto
	isSet bool
}

func (v NullableSetJobRetriesByProcessDto) Get() *SetJobRetriesByProcessDto {
	return v.value
}

func (v *NullableSetJobRetriesByProcessDto) Set(val *SetJobRetriesByProcessDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetJobRetriesByProcessDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetJobRetriesByProcessDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetJobRetriesByProcessDto(val *SetJobRetriesByProcessDto) *NullableSetJobRetriesByProcessDto {
	return &NullableSetJobRetriesByProcessDto{value: val, isSet: true}
}

func (v NullableSetJobRetriesByProcessDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetJobRetriesByProcessDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


