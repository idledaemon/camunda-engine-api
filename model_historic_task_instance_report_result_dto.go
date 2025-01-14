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

// checks if the HistoricTaskInstanceReportResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoricTaskInstanceReportResultDto{}

// HistoricTaskInstanceReportResultDto struct for HistoricTaskInstanceReportResultDto
type HistoricTaskInstanceReportResultDto struct {
	// The name of the task. It is only available when the `groupBy` parameter is set to `taskName`. Else the value is `null`.  **Note:** This property is only set for a historic task report object. In these cases, the value of the `reportType` query parameter is `count`.
	TaskName NullableString `json:"taskName,omitempty"`
	// The number of tasks which have the given definition.  **Note:** This property is only set for a historic task report object. In these cases, the value of the `reportType` query parameter is `count`.
	Count NullableInt64 `json:"count,omitempty"`
	// The key of the process definition.  **Note:** This property is only set for a historic task report object. In these cases, the value of the `reportType` query parameter is `count`.
	ProcessDefinitionKey NullableString `json:"processDefinitionKey,omitempty"`
	// The id of the process definition.  **Note:** This property is only set for a historic task report object. In these cases, the value of the `reportType` query parameter is `count`.
	ProcessDefinitionId NullableString `json:"processDefinitionId,omitempty"`
	// The name of the process definition.  **Note:** This property is only set for a historic task report object. In these cases, the value of the `reportType` query parameter is `count`.
	ProcessDefinitionName NullableString `json:"processDefinitionName,omitempty"`
	// Specifies a span of time within a year. **Note:** The period must be interpreted in conjunction with the returned `periodUnit`.  **Note:** This property is only set for a duration report object. In these cases, the value of the `reportType` query parameter is `duration`.
	Period NullableInt32 `json:"period,omitempty"`
	// The unit of the given period. Possible values are `MONTH` and `QUARTER`.  **Note:** This property is only set for a duration report object. In these cases, the value of the `reportType` query parameter is `duration`.
	PeriodUnit NullableString `json:"periodUnit,omitempty"`
	// The smallest duration in milliseconds of all completed process instances which were started in the given period.  **Note:** This property is only set for a duration report object. In these cases, the value of the `reportType` query parameter is `duration`.
	Minimum NullableInt64 `json:"minimum,omitempty"`
	// The greatest duration in milliseconds of all completed process instances which were started in the given period.  **Note:** This property is only set for a duration report object. In these cases, the value of the `reportType` query parameter is `duration`.
	Maximum NullableInt64 `json:"maximum,omitempty"`
	// The average duration in milliseconds of all completed process instances which were started in the given period.  **Note:** This property is only set for a duration report object. In these cases, the value of the `reportType` query parameter is `duration`.
	Average NullableInt64 `json:"average,omitempty"`
	// The id of the tenant.
	TenantId NullableString `json:"tenantId,omitempty"`
}

// NewHistoricTaskInstanceReportResultDto instantiates a new HistoricTaskInstanceReportResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricTaskInstanceReportResultDto() *HistoricTaskInstanceReportResultDto {
	this := HistoricTaskInstanceReportResultDto{}
	return &this
}

// NewHistoricTaskInstanceReportResultDtoWithDefaults instantiates a new HistoricTaskInstanceReportResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricTaskInstanceReportResultDtoWithDefaults() *HistoricTaskInstanceReportResultDto {
	this := HistoricTaskInstanceReportResultDto{}
	return &this
}

// GetTaskName returns the TaskName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricTaskInstanceReportResultDto) GetTaskName() string {
	if o == nil || IsNil(o.TaskName.Get()) {
		var ret string
		return ret
	}
	return *o.TaskName.Get()
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricTaskInstanceReportResultDto) GetTaskNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskName.Get(), o.TaskName.IsSet()
}

// HasTaskName returns a boolean if a field has been set.
func (o *HistoricTaskInstanceReportResultDto) HasTaskName() bool {
	if o != nil && o.TaskName.IsSet() {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given NullableString and assigns it to the TaskName field.
func (o *HistoricTaskInstanceReportResultDto) SetTaskName(v string) {
	o.TaskName.Set(&v)
}
// SetTaskNameNil sets the value for TaskName to be an explicit nil
func (o *HistoricTaskInstanceReportResultDto) SetTaskNameNil() {
	o.TaskName.Set(nil)
}

// UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil
func (o *HistoricTaskInstanceReportResultDto) UnsetTaskName() {
	o.TaskName.Unset()
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricTaskInstanceReportResultDto) GetCount() int64 {
	if o == nil || IsNil(o.Count.Get()) {
		var ret int64
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricTaskInstanceReportResultDto) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *HistoricTaskInstanceReportResultDto) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt64 and assigns it to the Count field.
func (o *HistoricTaskInstanceReportResultDto) SetCount(v int64) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *HistoricTaskInstanceReportResultDto) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *HistoricTaskInstanceReportResultDto) UnsetCount() {
	o.Count.Unset()
}

// GetProcessDefinitionKey returns the ProcessDefinitionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionKey() string {
	if o == nil || IsNil(o.ProcessDefinitionKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionKey.Get()
}

// GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionKey.Get(), o.ProcessDefinitionKey.IsSet()
}

// HasProcessDefinitionKey returns a boolean if a field has been set.
func (o *HistoricTaskInstanceReportResultDto) HasProcessDefinitionKey() bool {
	if o != nil && o.ProcessDefinitionKey.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionKey gets a reference to the given NullableString and assigns it to the ProcessDefinitionKey field.
func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionKey(v string) {
	o.ProcessDefinitionKey.Set(&v)
}
// SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil
func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionKeyNil() {
	o.ProcessDefinitionKey.Set(nil)
}

// UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
func (o *HistoricTaskInstanceReportResultDto) UnsetProcessDefinitionKey() {
	o.ProcessDefinitionKey.Unset()
}

// GetProcessDefinitionId returns the ProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionId() string {
	if o == nil || IsNil(o.ProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionId.Get()
}

// GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionId.Get(), o.ProcessDefinitionId.IsSet()
}

// HasProcessDefinitionId returns a boolean if a field has been set.
func (o *HistoricTaskInstanceReportResultDto) HasProcessDefinitionId() bool {
	if o != nil && o.ProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionId gets a reference to the given NullableString and assigns it to the ProcessDefinitionId field.
func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionId(v string) {
	o.ProcessDefinitionId.Set(&v)
}
// SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil
func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionIdNil() {
	o.ProcessDefinitionId.Set(nil)
}

// UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
func (o *HistoricTaskInstanceReportResultDto) UnsetProcessDefinitionId() {
	o.ProcessDefinitionId.Unset()
}

// GetProcessDefinitionName returns the ProcessDefinitionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionName() string {
	if o == nil || IsNil(o.ProcessDefinitionName.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionName.Get()
}

// GetProcessDefinitionNameOk returns a tuple with the ProcessDefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionName.Get(), o.ProcessDefinitionName.IsSet()
}

// HasProcessDefinitionName returns a boolean if a field has been set.
func (o *HistoricTaskInstanceReportResultDto) HasProcessDefinitionName() bool {
	if o != nil && o.ProcessDefinitionName.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionName gets a reference to the given NullableString and assigns it to the ProcessDefinitionName field.
func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionName(v string) {
	o.ProcessDefinitionName.Set(&v)
}
// SetProcessDefinitionNameNil sets the value for ProcessDefinitionName to be an explicit nil
func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionNameNil() {
	o.ProcessDefinitionName.Set(nil)
}

// UnsetProcessDefinitionName ensures that no value is present for ProcessDefinitionName, not even an explicit nil
func (o *HistoricTaskInstanceReportResultDto) UnsetProcessDefinitionName() {
	o.ProcessDefinitionName.Unset()
}

// GetPeriod returns the Period field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricTaskInstanceReportResultDto) GetPeriod() int32 {
	if o == nil || IsNil(o.Period.Get()) {
		var ret int32
		return ret
	}
	return *o.Period.Get()
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricTaskInstanceReportResultDto) GetPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Period.Get(), o.Period.IsSet()
}

// HasPeriod returns a boolean if a field has been set.
func (o *HistoricTaskInstanceReportResultDto) HasPeriod() bool {
	if o != nil && o.Period.IsSet() {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given NullableInt32 and assigns it to the Period field.
func (o *HistoricTaskInstanceReportResultDto) SetPeriod(v int32) {
	o.Period.Set(&v)
}
// SetPeriodNil sets the value for Period to be an explicit nil
func (o *HistoricTaskInstanceReportResultDto) SetPeriodNil() {
	o.Period.Set(nil)
}

// UnsetPeriod ensures that no value is present for Period, not even an explicit nil
func (o *HistoricTaskInstanceReportResultDto) UnsetPeriod() {
	o.Period.Unset()
}

// GetPeriodUnit returns the PeriodUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricTaskInstanceReportResultDto) GetPeriodUnit() string {
	if o == nil || IsNil(o.PeriodUnit.Get()) {
		var ret string
		return ret
	}
	return *o.PeriodUnit.Get()
}

// GetPeriodUnitOk returns a tuple with the PeriodUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricTaskInstanceReportResultDto) GetPeriodUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeriodUnit.Get(), o.PeriodUnit.IsSet()
}

// HasPeriodUnit returns a boolean if a field has been set.
func (o *HistoricTaskInstanceReportResultDto) HasPeriodUnit() bool {
	if o != nil && o.PeriodUnit.IsSet() {
		return true
	}

	return false
}

// SetPeriodUnit gets a reference to the given NullableString and assigns it to the PeriodUnit field.
func (o *HistoricTaskInstanceReportResultDto) SetPeriodUnit(v string) {
	o.PeriodUnit.Set(&v)
}
// SetPeriodUnitNil sets the value for PeriodUnit to be an explicit nil
func (o *HistoricTaskInstanceReportResultDto) SetPeriodUnitNil() {
	o.PeriodUnit.Set(nil)
}

// UnsetPeriodUnit ensures that no value is present for PeriodUnit, not even an explicit nil
func (o *HistoricTaskInstanceReportResultDto) UnsetPeriodUnit() {
	o.PeriodUnit.Unset()
}

// GetMinimum returns the Minimum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricTaskInstanceReportResultDto) GetMinimum() int64 {
	if o == nil || IsNil(o.Minimum.Get()) {
		var ret int64
		return ret
	}
	return *o.Minimum.Get()
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricTaskInstanceReportResultDto) GetMinimumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Minimum.Get(), o.Minimum.IsSet()
}

// HasMinimum returns a boolean if a field has been set.
func (o *HistoricTaskInstanceReportResultDto) HasMinimum() bool {
	if o != nil && o.Minimum.IsSet() {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given NullableInt64 and assigns it to the Minimum field.
func (o *HistoricTaskInstanceReportResultDto) SetMinimum(v int64) {
	o.Minimum.Set(&v)
}
// SetMinimumNil sets the value for Minimum to be an explicit nil
func (o *HistoricTaskInstanceReportResultDto) SetMinimumNil() {
	o.Minimum.Set(nil)
}

// UnsetMinimum ensures that no value is present for Minimum, not even an explicit nil
func (o *HistoricTaskInstanceReportResultDto) UnsetMinimum() {
	o.Minimum.Unset()
}

// GetMaximum returns the Maximum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricTaskInstanceReportResultDto) GetMaximum() int64 {
	if o == nil || IsNil(o.Maximum.Get()) {
		var ret int64
		return ret
	}
	return *o.Maximum.Get()
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricTaskInstanceReportResultDto) GetMaximumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Maximum.Get(), o.Maximum.IsSet()
}

// HasMaximum returns a boolean if a field has been set.
func (o *HistoricTaskInstanceReportResultDto) HasMaximum() bool {
	if o != nil && o.Maximum.IsSet() {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given NullableInt64 and assigns it to the Maximum field.
func (o *HistoricTaskInstanceReportResultDto) SetMaximum(v int64) {
	o.Maximum.Set(&v)
}
// SetMaximumNil sets the value for Maximum to be an explicit nil
func (o *HistoricTaskInstanceReportResultDto) SetMaximumNil() {
	o.Maximum.Set(nil)
}

// UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil
func (o *HistoricTaskInstanceReportResultDto) UnsetMaximum() {
	o.Maximum.Unset()
}

// GetAverage returns the Average field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricTaskInstanceReportResultDto) GetAverage() int64 {
	if o == nil || IsNil(o.Average.Get()) {
		var ret int64
		return ret
	}
	return *o.Average.Get()
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricTaskInstanceReportResultDto) GetAverageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Average.Get(), o.Average.IsSet()
}

// HasAverage returns a boolean if a field has been set.
func (o *HistoricTaskInstanceReportResultDto) HasAverage() bool {
	if o != nil && o.Average.IsSet() {
		return true
	}

	return false
}

// SetAverage gets a reference to the given NullableInt64 and assigns it to the Average field.
func (o *HistoricTaskInstanceReportResultDto) SetAverage(v int64) {
	o.Average.Set(&v)
}
// SetAverageNil sets the value for Average to be an explicit nil
func (o *HistoricTaskInstanceReportResultDto) SetAverageNil() {
	o.Average.Set(nil)
}

// UnsetAverage ensures that no value is present for Average, not even an explicit nil
func (o *HistoricTaskInstanceReportResultDto) UnsetAverage() {
	o.Average.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricTaskInstanceReportResultDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricTaskInstanceReportResultDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *HistoricTaskInstanceReportResultDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *HistoricTaskInstanceReportResultDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *HistoricTaskInstanceReportResultDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *HistoricTaskInstanceReportResultDto) UnsetTenantId() {
	o.TenantId.Unset()
}

func (o HistoricTaskInstanceReportResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricTaskInstanceReportResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TaskName.IsSet() {
		toSerialize["taskName"] = o.TaskName.Get()
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.ProcessDefinitionKey.IsSet() {
		toSerialize["processDefinitionKey"] = o.ProcessDefinitionKey.Get()
	}
	if o.ProcessDefinitionId.IsSet() {
		toSerialize["processDefinitionId"] = o.ProcessDefinitionId.Get()
	}
	if o.ProcessDefinitionName.IsSet() {
		toSerialize["processDefinitionName"] = o.ProcessDefinitionName.Get()
	}
	if o.Period.IsSet() {
		toSerialize["period"] = o.Period.Get()
	}
	if o.PeriodUnit.IsSet() {
		toSerialize["periodUnit"] = o.PeriodUnit.Get()
	}
	if o.Minimum.IsSet() {
		toSerialize["minimum"] = o.Minimum.Get()
	}
	if o.Maximum.IsSet() {
		toSerialize["maximum"] = o.Maximum.Get()
	}
	if o.Average.IsSet() {
		toSerialize["average"] = o.Average.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	return toSerialize, nil
}

type NullableHistoricTaskInstanceReportResultDto struct {
	value *HistoricTaskInstanceReportResultDto
	isSet bool
}

func (v NullableHistoricTaskInstanceReportResultDto) Get() *HistoricTaskInstanceReportResultDto {
	return v.value
}

func (v *NullableHistoricTaskInstanceReportResultDto) Set(val *HistoricTaskInstanceReportResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricTaskInstanceReportResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricTaskInstanceReportResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricTaskInstanceReportResultDto(val *HistoricTaskInstanceReportResultDto) *NullableHistoricTaskInstanceReportResultDto {
	return &NullableHistoricTaskInstanceReportResultDto{value: val, isSet: true}
}

func (v NullableHistoricTaskInstanceReportResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricTaskInstanceReportResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


