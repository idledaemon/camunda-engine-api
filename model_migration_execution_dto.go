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

// checks if the MigrationExecutionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrationExecutionDto{}

// MigrationExecutionDto struct for MigrationExecutionDto
type MigrationExecutionDto struct {
	MigrationPlan *MigrationPlanDto `json:"migrationPlan,omitempty"`
	// A list of process instance ids to migrate.
	ProcessInstanceIds []string `json:"processInstanceIds,omitempty"`
	ProcessInstanceQuery *ProcessInstanceQueryDto `json:"processInstanceQuery,omitempty"`
	// A boolean value to control whether execution listeners should be invoked during migration.
	SkipCustomListeners NullableBool `json:"skipCustomListeners,omitempty"`
	// A boolean value to control whether input/output mappings should be executed during migration.
	SkipIoMappings NullableBool `json:"skipIoMappings,omitempty"`
}

// NewMigrationExecutionDto instantiates a new MigrationExecutionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationExecutionDto() *MigrationExecutionDto {
	this := MigrationExecutionDto{}
	return &this
}

// NewMigrationExecutionDtoWithDefaults instantiates a new MigrationExecutionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationExecutionDtoWithDefaults() *MigrationExecutionDto {
	this := MigrationExecutionDto{}
	return &this
}

// GetMigrationPlan returns the MigrationPlan field value if set, zero value otherwise.
func (o *MigrationExecutionDto) GetMigrationPlan() MigrationPlanDto {
	if o == nil || IsNil(o.MigrationPlan) {
		var ret MigrationPlanDto
		return ret
	}
	return *o.MigrationPlan
}

// GetMigrationPlanOk returns a tuple with the MigrationPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationExecutionDto) GetMigrationPlanOk() (*MigrationPlanDto, bool) {
	if o == nil || IsNil(o.MigrationPlan) {
		return nil, false
	}
	return o.MigrationPlan, true
}

// HasMigrationPlan returns a boolean if a field has been set.
func (o *MigrationExecutionDto) HasMigrationPlan() bool {
	if o != nil && !IsNil(o.MigrationPlan) {
		return true
	}

	return false
}

// SetMigrationPlan gets a reference to the given MigrationPlanDto and assigns it to the MigrationPlan field.
func (o *MigrationExecutionDto) SetMigrationPlan(v MigrationPlanDto) {
	o.MigrationPlan = &v
}

// GetProcessInstanceIds returns the ProcessInstanceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationExecutionDto) GetProcessInstanceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProcessInstanceIds
}

// GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationExecutionDto) GetProcessInstanceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessInstanceIds) {
		return nil, false
	}
	return o.ProcessInstanceIds, true
}

// HasProcessInstanceIds returns a boolean if a field has been set.
func (o *MigrationExecutionDto) HasProcessInstanceIds() bool {
	if o != nil && !IsNil(o.ProcessInstanceIds) {
		return true
	}

	return false
}

// SetProcessInstanceIds gets a reference to the given []string and assigns it to the ProcessInstanceIds field.
func (o *MigrationExecutionDto) SetProcessInstanceIds(v []string) {
	o.ProcessInstanceIds = v
}

// GetProcessInstanceQuery returns the ProcessInstanceQuery field value if set, zero value otherwise.
func (o *MigrationExecutionDto) GetProcessInstanceQuery() ProcessInstanceQueryDto {
	if o == nil || IsNil(o.ProcessInstanceQuery) {
		var ret ProcessInstanceQueryDto
		return ret
	}
	return *o.ProcessInstanceQuery
}

// GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationExecutionDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool) {
	if o == nil || IsNil(o.ProcessInstanceQuery) {
		return nil, false
	}
	return o.ProcessInstanceQuery, true
}

// HasProcessInstanceQuery returns a boolean if a field has been set.
func (o *MigrationExecutionDto) HasProcessInstanceQuery() bool {
	if o != nil && !IsNil(o.ProcessInstanceQuery) {
		return true
	}

	return false
}

// SetProcessInstanceQuery gets a reference to the given ProcessInstanceQueryDto and assigns it to the ProcessInstanceQuery field.
func (o *MigrationExecutionDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto) {
	o.ProcessInstanceQuery = &v
}

// GetSkipCustomListeners returns the SkipCustomListeners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationExecutionDto) GetSkipCustomListeners() bool {
	if o == nil || IsNil(o.SkipCustomListeners.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipCustomListeners.Get()
}

// GetSkipCustomListenersOk returns a tuple with the SkipCustomListeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationExecutionDto) GetSkipCustomListenersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipCustomListeners.Get(), o.SkipCustomListeners.IsSet()
}

// HasSkipCustomListeners returns a boolean if a field has been set.
func (o *MigrationExecutionDto) HasSkipCustomListeners() bool {
	if o != nil && o.SkipCustomListeners.IsSet() {
		return true
	}

	return false
}

// SetSkipCustomListeners gets a reference to the given NullableBool and assigns it to the SkipCustomListeners field.
func (o *MigrationExecutionDto) SetSkipCustomListeners(v bool) {
	o.SkipCustomListeners.Set(&v)
}
// SetSkipCustomListenersNil sets the value for SkipCustomListeners to be an explicit nil
func (o *MigrationExecutionDto) SetSkipCustomListenersNil() {
	o.SkipCustomListeners.Set(nil)
}

// UnsetSkipCustomListeners ensures that no value is present for SkipCustomListeners, not even an explicit nil
func (o *MigrationExecutionDto) UnsetSkipCustomListeners() {
	o.SkipCustomListeners.Unset()
}

// GetSkipIoMappings returns the SkipIoMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrationExecutionDto) GetSkipIoMappings() bool {
	if o == nil || IsNil(o.SkipIoMappings.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipIoMappings.Get()
}

// GetSkipIoMappingsOk returns a tuple with the SkipIoMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrationExecutionDto) GetSkipIoMappingsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipIoMappings.Get(), o.SkipIoMappings.IsSet()
}

// HasSkipIoMappings returns a boolean if a field has been set.
func (o *MigrationExecutionDto) HasSkipIoMappings() bool {
	if o != nil && o.SkipIoMappings.IsSet() {
		return true
	}

	return false
}

// SetSkipIoMappings gets a reference to the given NullableBool and assigns it to the SkipIoMappings field.
func (o *MigrationExecutionDto) SetSkipIoMappings(v bool) {
	o.SkipIoMappings.Set(&v)
}
// SetSkipIoMappingsNil sets the value for SkipIoMappings to be an explicit nil
func (o *MigrationExecutionDto) SetSkipIoMappingsNil() {
	o.SkipIoMappings.Set(nil)
}

// UnsetSkipIoMappings ensures that no value is present for SkipIoMappings, not even an explicit nil
func (o *MigrationExecutionDto) UnsetSkipIoMappings() {
	o.SkipIoMappings.Unset()
}

func (o MigrationExecutionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrationExecutionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MigrationPlan) {
		toSerialize["migrationPlan"] = o.MigrationPlan
	}
	if o.ProcessInstanceIds != nil {
		toSerialize["processInstanceIds"] = o.ProcessInstanceIds
	}
	if !IsNil(o.ProcessInstanceQuery) {
		toSerialize["processInstanceQuery"] = o.ProcessInstanceQuery
	}
	if o.SkipCustomListeners.IsSet() {
		toSerialize["skipCustomListeners"] = o.SkipCustomListeners.Get()
	}
	if o.SkipIoMappings.IsSet() {
		toSerialize["skipIoMappings"] = o.SkipIoMappings.Get()
	}
	return toSerialize, nil
}

type NullableMigrationExecutionDto struct {
	value *MigrationExecutionDto
	isSet bool
}

func (v NullableMigrationExecutionDto) Get() *MigrationExecutionDto {
	return v.value
}

func (v *NullableMigrationExecutionDto) Set(val *MigrationExecutionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationExecutionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationExecutionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationExecutionDto(val *MigrationExecutionDto) *NullableMigrationExecutionDto {
	return &NullableMigrationExecutionDto{value: val, isSet: true}
}

func (v NullableMigrationExecutionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationExecutionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


