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

// checks if the DeleteProcessInstancesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteProcessInstancesDto{}

// DeleteProcessInstancesDto struct for DeleteProcessInstancesDto
type DeleteProcessInstancesDto struct {
	// A list process instance ids to delete.
	ProcessInstanceIds []string `json:"processInstanceIds,omitempty"`
	// A string with delete reason.
	DeleteReason NullableString `json:"deleteReason,omitempty"`
	// Skip execution listener invocation for activities that are started or ended as part of this request.
	SkipCustomListeners NullableBool `json:"skipCustomListeners,omitempty"`
	// Skip deletion of the subprocesses related to deleted processes as part of this request.
	SkipSubprocesses NullableBool `json:"skipSubprocesses,omitempty"`
	// Skip execution of [input/output variable mappings](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#input-output-variable-mapping) for activities that are started or ended as part of this request.
	SkipIoMappings NullableBool `json:"skipIoMappings,omitempty"`
	ProcessInstanceQuery *ProcessInstanceQueryDto `json:"processInstanceQuery,omitempty"`
	HistoricProcessInstanceQuery *HistoricProcessInstanceQueryDto `json:"historicProcessInstanceQuery,omitempty"`
}

// NewDeleteProcessInstancesDto instantiates a new DeleteProcessInstancesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteProcessInstancesDto() *DeleteProcessInstancesDto {
	this := DeleteProcessInstancesDto{}
	return &this
}

// NewDeleteProcessInstancesDtoWithDefaults instantiates a new DeleteProcessInstancesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteProcessInstancesDtoWithDefaults() *DeleteProcessInstancesDto {
	this := DeleteProcessInstancesDto{}
	return &this
}

// GetProcessInstanceIds returns the ProcessInstanceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteProcessInstancesDto) GetProcessInstanceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProcessInstanceIds
}

// GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteProcessInstancesDto) GetProcessInstanceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessInstanceIds) {
		return nil, false
	}
	return o.ProcessInstanceIds, true
}

// HasProcessInstanceIds returns a boolean if a field has been set.
func (o *DeleteProcessInstancesDto) HasProcessInstanceIds() bool {
	if o != nil && !IsNil(o.ProcessInstanceIds) {
		return true
	}

	return false
}

// SetProcessInstanceIds gets a reference to the given []string and assigns it to the ProcessInstanceIds field.
func (o *DeleteProcessInstancesDto) SetProcessInstanceIds(v []string) {
	o.ProcessInstanceIds = v
}

// GetDeleteReason returns the DeleteReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteProcessInstancesDto) GetDeleteReason() string {
	if o == nil || IsNil(o.DeleteReason.Get()) {
		var ret string
		return ret
	}
	return *o.DeleteReason.Get()
}

// GetDeleteReasonOk returns a tuple with the DeleteReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteProcessInstancesDto) GetDeleteReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeleteReason.Get(), o.DeleteReason.IsSet()
}

// HasDeleteReason returns a boolean if a field has been set.
func (o *DeleteProcessInstancesDto) HasDeleteReason() bool {
	if o != nil && o.DeleteReason.IsSet() {
		return true
	}

	return false
}

// SetDeleteReason gets a reference to the given NullableString and assigns it to the DeleteReason field.
func (o *DeleteProcessInstancesDto) SetDeleteReason(v string) {
	o.DeleteReason.Set(&v)
}
// SetDeleteReasonNil sets the value for DeleteReason to be an explicit nil
func (o *DeleteProcessInstancesDto) SetDeleteReasonNil() {
	o.DeleteReason.Set(nil)
}

// UnsetDeleteReason ensures that no value is present for DeleteReason, not even an explicit nil
func (o *DeleteProcessInstancesDto) UnsetDeleteReason() {
	o.DeleteReason.Unset()
}

// GetSkipCustomListeners returns the SkipCustomListeners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteProcessInstancesDto) GetSkipCustomListeners() bool {
	if o == nil || IsNil(o.SkipCustomListeners.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipCustomListeners.Get()
}

// GetSkipCustomListenersOk returns a tuple with the SkipCustomListeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteProcessInstancesDto) GetSkipCustomListenersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipCustomListeners.Get(), o.SkipCustomListeners.IsSet()
}

// HasSkipCustomListeners returns a boolean if a field has been set.
func (o *DeleteProcessInstancesDto) HasSkipCustomListeners() bool {
	if o != nil && o.SkipCustomListeners.IsSet() {
		return true
	}

	return false
}

// SetSkipCustomListeners gets a reference to the given NullableBool and assigns it to the SkipCustomListeners field.
func (o *DeleteProcessInstancesDto) SetSkipCustomListeners(v bool) {
	o.SkipCustomListeners.Set(&v)
}
// SetSkipCustomListenersNil sets the value for SkipCustomListeners to be an explicit nil
func (o *DeleteProcessInstancesDto) SetSkipCustomListenersNil() {
	o.SkipCustomListeners.Set(nil)
}

// UnsetSkipCustomListeners ensures that no value is present for SkipCustomListeners, not even an explicit nil
func (o *DeleteProcessInstancesDto) UnsetSkipCustomListeners() {
	o.SkipCustomListeners.Unset()
}

// GetSkipSubprocesses returns the SkipSubprocesses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteProcessInstancesDto) GetSkipSubprocesses() bool {
	if o == nil || IsNil(o.SkipSubprocesses.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipSubprocesses.Get()
}

// GetSkipSubprocessesOk returns a tuple with the SkipSubprocesses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteProcessInstancesDto) GetSkipSubprocessesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipSubprocesses.Get(), o.SkipSubprocesses.IsSet()
}

// HasSkipSubprocesses returns a boolean if a field has been set.
func (o *DeleteProcessInstancesDto) HasSkipSubprocesses() bool {
	if o != nil && o.SkipSubprocesses.IsSet() {
		return true
	}

	return false
}

// SetSkipSubprocesses gets a reference to the given NullableBool and assigns it to the SkipSubprocesses field.
func (o *DeleteProcessInstancesDto) SetSkipSubprocesses(v bool) {
	o.SkipSubprocesses.Set(&v)
}
// SetSkipSubprocessesNil sets the value for SkipSubprocesses to be an explicit nil
func (o *DeleteProcessInstancesDto) SetSkipSubprocessesNil() {
	o.SkipSubprocesses.Set(nil)
}

// UnsetSkipSubprocesses ensures that no value is present for SkipSubprocesses, not even an explicit nil
func (o *DeleteProcessInstancesDto) UnsetSkipSubprocesses() {
	o.SkipSubprocesses.Unset()
}

// GetSkipIoMappings returns the SkipIoMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteProcessInstancesDto) GetSkipIoMappings() bool {
	if o == nil || IsNil(o.SkipIoMappings.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipIoMappings.Get()
}

// GetSkipIoMappingsOk returns a tuple with the SkipIoMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteProcessInstancesDto) GetSkipIoMappingsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipIoMappings.Get(), o.SkipIoMappings.IsSet()
}

// HasSkipIoMappings returns a boolean if a field has been set.
func (o *DeleteProcessInstancesDto) HasSkipIoMappings() bool {
	if o != nil && o.SkipIoMappings.IsSet() {
		return true
	}

	return false
}

// SetSkipIoMappings gets a reference to the given NullableBool and assigns it to the SkipIoMappings field.
func (o *DeleteProcessInstancesDto) SetSkipIoMappings(v bool) {
	o.SkipIoMappings.Set(&v)
}
// SetSkipIoMappingsNil sets the value for SkipIoMappings to be an explicit nil
func (o *DeleteProcessInstancesDto) SetSkipIoMappingsNil() {
	o.SkipIoMappings.Set(nil)
}

// UnsetSkipIoMappings ensures that no value is present for SkipIoMappings, not even an explicit nil
func (o *DeleteProcessInstancesDto) UnsetSkipIoMappings() {
	o.SkipIoMappings.Unset()
}

// GetProcessInstanceQuery returns the ProcessInstanceQuery field value if set, zero value otherwise.
func (o *DeleteProcessInstancesDto) GetProcessInstanceQuery() ProcessInstanceQueryDto {
	if o == nil || IsNil(o.ProcessInstanceQuery) {
		var ret ProcessInstanceQueryDto
		return ret
	}
	return *o.ProcessInstanceQuery
}

// GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteProcessInstancesDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool) {
	if o == nil || IsNil(o.ProcessInstanceQuery) {
		return nil, false
	}
	return o.ProcessInstanceQuery, true
}

// HasProcessInstanceQuery returns a boolean if a field has been set.
func (o *DeleteProcessInstancesDto) HasProcessInstanceQuery() bool {
	if o != nil && !IsNil(o.ProcessInstanceQuery) {
		return true
	}

	return false
}

// SetProcessInstanceQuery gets a reference to the given ProcessInstanceQueryDto and assigns it to the ProcessInstanceQuery field.
func (o *DeleteProcessInstancesDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto) {
	o.ProcessInstanceQuery = &v
}

// GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field value if set, zero value otherwise.
func (o *DeleteProcessInstancesDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto {
	if o == nil || IsNil(o.HistoricProcessInstanceQuery) {
		var ret HistoricProcessInstanceQueryDto
		return ret
	}
	return *o.HistoricProcessInstanceQuery
}

// GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteProcessInstancesDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool) {
	if o == nil || IsNil(o.HistoricProcessInstanceQuery) {
		return nil, false
	}
	return o.HistoricProcessInstanceQuery, true
}

// HasHistoricProcessInstanceQuery returns a boolean if a field has been set.
func (o *DeleteProcessInstancesDto) HasHistoricProcessInstanceQuery() bool {
	if o != nil && !IsNil(o.HistoricProcessInstanceQuery) {
		return true
	}

	return false
}

// SetHistoricProcessInstanceQuery gets a reference to the given HistoricProcessInstanceQueryDto and assigns it to the HistoricProcessInstanceQuery field.
func (o *DeleteProcessInstancesDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto) {
	o.HistoricProcessInstanceQuery = &v
}

func (o DeleteProcessInstancesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteProcessInstancesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProcessInstanceIds != nil {
		toSerialize["processInstanceIds"] = o.ProcessInstanceIds
	}
	if o.DeleteReason.IsSet() {
		toSerialize["deleteReason"] = o.DeleteReason.Get()
	}
	if o.SkipCustomListeners.IsSet() {
		toSerialize["skipCustomListeners"] = o.SkipCustomListeners.Get()
	}
	if o.SkipSubprocesses.IsSet() {
		toSerialize["skipSubprocesses"] = o.SkipSubprocesses.Get()
	}
	if o.SkipIoMappings.IsSet() {
		toSerialize["skipIoMappings"] = o.SkipIoMappings.Get()
	}
	if !IsNil(o.ProcessInstanceQuery) {
		toSerialize["processInstanceQuery"] = o.ProcessInstanceQuery
	}
	if !IsNil(o.HistoricProcessInstanceQuery) {
		toSerialize["historicProcessInstanceQuery"] = o.HistoricProcessInstanceQuery
	}
	return toSerialize, nil
}

type NullableDeleteProcessInstancesDto struct {
	value *DeleteProcessInstancesDto
	isSet bool
}

func (v NullableDeleteProcessInstancesDto) Get() *DeleteProcessInstancesDto {
	return v.value
}

func (v *NullableDeleteProcessInstancesDto) Set(val *DeleteProcessInstancesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteProcessInstancesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteProcessInstancesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteProcessInstancesDto(val *DeleteProcessInstancesDto) *NullableDeleteProcessInstancesDto {
	return &NullableDeleteProcessInstancesDto{value: val, isSet: true}
}

func (v NullableDeleteProcessInstancesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteProcessInstancesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


