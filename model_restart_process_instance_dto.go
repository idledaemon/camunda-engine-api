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

// checks if the RestartProcessInstanceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestartProcessInstanceDto{}

// RestartProcessInstanceDto struct for RestartProcessInstanceDto
type RestartProcessInstanceDto struct {
	// A list of process instance ids to restart.
	ProcessInstanceIds []string `json:"processInstanceIds,omitempty"`
	HistoricProcessInstanceQuery *HistoricProcessInstanceQueryDto `json:"historicProcessInstanceQuery,omitempty"`
	// Skip execution listener invocation for activities that are started as part of this request.
	SkipCustomListeners NullableBool `json:"skipCustomListeners,omitempty"`
	// Skip execution of [input/output variable mappings](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#input-output-variable-mapping) for activities that are started as part of this request.
	SkipIoMappings NullableBool `json:"skipIoMappings,omitempty"`
	// Set the initial set of variables during restart. By default, the last set of variables is used.
	InitialVariables NullableBool `json:"initialVariables,omitempty"`
	// Do not take over the business key of the historic process instance.
	WithoutBusinessKey NullableBool `json:"withoutBusinessKey,omitempty"`
	// **Optional**. A JSON array of instructions that specify which activities to start the process instance at. If this property is omitted, the process instance starts at its default blank start event.
	Instructions []RestartProcessInstanceModificationInstructionDto `json:"instructions,omitempty"`
}

// NewRestartProcessInstanceDto instantiates a new RestartProcessInstanceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestartProcessInstanceDto() *RestartProcessInstanceDto {
	this := RestartProcessInstanceDto{}
	return &this
}

// NewRestartProcessInstanceDtoWithDefaults instantiates a new RestartProcessInstanceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestartProcessInstanceDtoWithDefaults() *RestartProcessInstanceDto {
	this := RestartProcessInstanceDto{}
	return &this
}

// GetProcessInstanceIds returns the ProcessInstanceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestartProcessInstanceDto) GetProcessInstanceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProcessInstanceIds
}

// GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestartProcessInstanceDto) GetProcessInstanceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessInstanceIds) {
		return nil, false
	}
	return o.ProcessInstanceIds, true
}

// HasProcessInstanceIds returns a boolean if a field has been set.
func (o *RestartProcessInstanceDto) HasProcessInstanceIds() bool {
	if o != nil && !IsNil(o.ProcessInstanceIds) {
		return true
	}

	return false
}

// SetProcessInstanceIds gets a reference to the given []string and assigns it to the ProcessInstanceIds field.
func (o *RestartProcessInstanceDto) SetProcessInstanceIds(v []string) {
	o.ProcessInstanceIds = v
}

// GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field value if set, zero value otherwise.
func (o *RestartProcessInstanceDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto {
	if o == nil || IsNil(o.HistoricProcessInstanceQuery) {
		var ret HistoricProcessInstanceQueryDto
		return ret
	}
	return *o.HistoricProcessInstanceQuery
}

// GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestartProcessInstanceDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool) {
	if o == nil || IsNil(o.HistoricProcessInstanceQuery) {
		return nil, false
	}
	return o.HistoricProcessInstanceQuery, true
}

// HasHistoricProcessInstanceQuery returns a boolean if a field has been set.
func (o *RestartProcessInstanceDto) HasHistoricProcessInstanceQuery() bool {
	if o != nil && !IsNil(o.HistoricProcessInstanceQuery) {
		return true
	}

	return false
}

// SetHistoricProcessInstanceQuery gets a reference to the given HistoricProcessInstanceQueryDto and assigns it to the HistoricProcessInstanceQuery field.
func (o *RestartProcessInstanceDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto) {
	o.HistoricProcessInstanceQuery = &v
}

// GetSkipCustomListeners returns the SkipCustomListeners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestartProcessInstanceDto) GetSkipCustomListeners() bool {
	if o == nil || IsNil(o.SkipCustomListeners.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipCustomListeners.Get()
}

// GetSkipCustomListenersOk returns a tuple with the SkipCustomListeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestartProcessInstanceDto) GetSkipCustomListenersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipCustomListeners.Get(), o.SkipCustomListeners.IsSet()
}

// HasSkipCustomListeners returns a boolean if a field has been set.
func (o *RestartProcessInstanceDto) HasSkipCustomListeners() bool {
	if o != nil && o.SkipCustomListeners.IsSet() {
		return true
	}

	return false
}

// SetSkipCustomListeners gets a reference to the given NullableBool and assigns it to the SkipCustomListeners field.
func (o *RestartProcessInstanceDto) SetSkipCustomListeners(v bool) {
	o.SkipCustomListeners.Set(&v)
}
// SetSkipCustomListenersNil sets the value for SkipCustomListeners to be an explicit nil
func (o *RestartProcessInstanceDto) SetSkipCustomListenersNil() {
	o.SkipCustomListeners.Set(nil)
}

// UnsetSkipCustomListeners ensures that no value is present for SkipCustomListeners, not even an explicit nil
func (o *RestartProcessInstanceDto) UnsetSkipCustomListeners() {
	o.SkipCustomListeners.Unset()
}

// GetSkipIoMappings returns the SkipIoMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestartProcessInstanceDto) GetSkipIoMappings() bool {
	if o == nil || IsNil(o.SkipIoMappings.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipIoMappings.Get()
}

// GetSkipIoMappingsOk returns a tuple with the SkipIoMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestartProcessInstanceDto) GetSkipIoMappingsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipIoMappings.Get(), o.SkipIoMappings.IsSet()
}

// HasSkipIoMappings returns a boolean if a field has been set.
func (o *RestartProcessInstanceDto) HasSkipIoMappings() bool {
	if o != nil && o.SkipIoMappings.IsSet() {
		return true
	}

	return false
}

// SetSkipIoMappings gets a reference to the given NullableBool and assigns it to the SkipIoMappings field.
func (o *RestartProcessInstanceDto) SetSkipIoMappings(v bool) {
	o.SkipIoMappings.Set(&v)
}
// SetSkipIoMappingsNil sets the value for SkipIoMappings to be an explicit nil
func (o *RestartProcessInstanceDto) SetSkipIoMappingsNil() {
	o.SkipIoMappings.Set(nil)
}

// UnsetSkipIoMappings ensures that no value is present for SkipIoMappings, not even an explicit nil
func (o *RestartProcessInstanceDto) UnsetSkipIoMappings() {
	o.SkipIoMappings.Unset()
}

// GetInitialVariables returns the InitialVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestartProcessInstanceDto) GetInitialVariables() bool {
	if o == nil || IsNil(o.InitialVariables.Get()) {
		var ret bool
		return ret
	}
	return *o.InitialVariables.Get()
}

// GetInitialVariablesOk returns a tuple with the InitialVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestartProcessInstanceDto) GetInitialVariablesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitialVariables.Get(), o.InitialVariables.IsSet()
}

// HasInitialVariables returns a boolean if a field has been set.
func (o *RestartProcessInstanceDto) HasInitialVariables() bool {
	if o != nil && o.InitialVariables.IsSet() {
		return true
	}

	return false
}

// SetInitialVariables gets a reference to the given NullableBool and assigns it to the InitialVariables field.
func (o *RestartProcessInstanceDto) SetInitialVariables(v bool) {
	o.InitialVariables.Set(&v)
}
// SetInitialVariablesNil sets the value for InitialVariables to be an explicit nil
func (o *RestartProcessInstanceDto) SetInitialVariablesNil() {
	o.InitialVariables.Set(nil)
}

// UnsetInitialVariables ensures that no value is present for InitialVariables, not even an explicit nil
func (o *RestartProcessInstanceDto) UnsetInitialVariables() {
	o.InitialVariables.Unset()
}

// GetWithoutBusinessKey returns the WithoutBusinessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestartProcessInstanceDto) GetWithoutBusinessKey() bool {
	if o == nil || IsNil(o.WithoutBusinessKey.Get()) {
		var ret bool
		return ret
	}
	return *o.WithoutBusinessKey.Get()
}

// GetWithoutBusinessKeyOk returns a tuple with the WithoutBusinessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestartProcessInstanceDto) GetWithoutBusinessKeyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithoutBusinessKey.Get(), o.WithoutBusinessKey.IsSet()
}

// HasWithoutBusinessKey returns a boolean if a field has been set.
func (o *RestartProcessInstanceDto) HasWithoutBusinessKey() bool {
	if o != nil && o.WithoutBusinessKey.IsSet() {
		return true
	}

	return false
}

// SetWithoutBusinessKey gets a reference to the given NullableBool and assigns it to the WithoutBusinessKey field.
func (o *RestartProcessInstanceDto) SetWithoutBusinessKey(v bool) {
	o.WithoutBusinessKey.Set(&v)
}
// SetWithoutBusinessKeyNil sets the value for WithoutBusinessKey to be an explicit nil
func (o *RestartProcessInstanceDto) SetWithoutBusinessKeyNil() {
	o.WithoutBusinessKey.Set(nil)
}

// UnsetWithoutBusinessKey ensures that no value is present for WithoutBusinessKey, not even an explicit nil
func (o *RestartProcessInstanceDto) UnsetWithoutBusinessKey() {
	o.WithoutBusinessKey.Unset()
}

// GetInstructions returns the Instructions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestartProcessInstanceDto) GetInstructions() []RestartProcessInstanceModificationInstructionDto {
	if o == nil {
		var ret []RestartProcessInstanceModificationInstructionDto
		return ret
	}
	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestartProcessInstanceDto) GetInstructionsOk() ([]RestartProcessInstanceModificationInstructionDto, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *RestartProcessInstanceDto) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given []RestartProcessInstanceModificationInstructionDto and assigns it to the Instructions field.
func (o *RestartProcessInstanceDto) SetInstructions(v []RestartProcessInstanceModificationInstructionDto) {
	o.Instructions = v
}

func (o RestartProcessInstanceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestartProcessInstanceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProcessInstanceIds != nil {
		toSerialize["processInstanceIds"] = o.ProcessInstanceIds
	}
	if !IsNil(o.HistoricProcessInstanceQuery) {
		toSerialize["historicProcessInstanceQuery"] = o.HistoricProcessInstanceQuery
	}
	if o.SkipCustomListeners.IsSet() {
		toSerialize["skipCustomListeners"] = o.SkipCustomListeners.Get()
	}
	if o.SkipIoMappings.IsSet() {
		toSerialize["skipIoMappings"] = o.SkipIoMappings.Get()
	}
	if o.InitialVariables.IsSet() {
		toSerialize["initialVariables"] = o.InitialVariables.Get()
	}
	if o.WithoutBusinessKey.IsSet() {
		toSerialize["withoutBusinessKey"] = o.WithoutBusinessKey.Get()
	}
	if o.Instructions != nil {
		toSerialize["instructions"] = o.Instructions
	}
	return toSerialize, nil
}

type NullableRestartProcessInstanceDto struct {
	value *RestartProcessInstanceDto
	isSet bool
}

func (v NullableRestartProcessInstanceDto) Get() *RestartProcessInstanceDto {
	return v.value
}

func (v *NullableRestartProcessInstanceDto) Set(val *RestartProcessInstanceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRestartProcessInstanceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRestartProcessInstanceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestartProcessInstanceDto(val *RestartProcessInstanceDto) *NullableRestartProcessInstanceDto {
	return &NullableRestartProcessInstanceDto{value: val, isSet: true}
}

func (v NullableRestartProcessInstanceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestartProcessInstanceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

