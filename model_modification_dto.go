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

// checks if the ModificationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModificationDto{}

// ModificationDto struct for ModificationDto
type ModificationDto struct {
	// The id of the process definition for the modification
	ProcessDefinitionId NullableString `json:"processDefinitionId,omitempty"`
	// Skip execution listener invocation for activities that are started or ended as part of this request.
	SkipCustomListeners NullableBool `json:"skipCustomListeners,omitempty"`
	// Skip execution of [input/output variable mappings](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#input-output-variable-mapping) for activities that are started or ended as part of this request.
	SkipIoMappings NullableBool `json:"skipIoMappings,omitempty"`
	// A list of process instance ids to modify.
	ProcessInstanceIds []string `json:"processInstanceIds,omitempty"`
	ProcessInstanceQuery *ProcessInstanceQueryDto `json:"processInstanceQuery,omitempty"`
	// An array of modification instructions. The instructions are executed in the order they are in. 
	Instructions []MultipleProcessInstanceModificationInstructionDto `json:"instructions,omitempty"`
	// An arbitrary text annotation set by a user for auditing reasons.
	Annotation NullableString `json:"annotation,omitempty"`
}

// NewModificationDto instantiates a new ModificationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModificationDto() *ModificationDto {
	this := ModificationDto{}
	return &this
}

// NewModificationDtoWithDefaults instantiates a new ModificationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModificationDtoWithDefaults() *ModificationDto {
	this := ModificationDto{}
	return &this
}

// GetProcessDefinitionId returns the ProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModificationDto) GetProcessDefinitionId() string {
	if o == nil || IsNil(o.ProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionId.Get()
}

// GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModificationDto) GetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionId.Get(), o.ProcessDefinitionId.IsSet()
}

// HasProcessDefinitionId returns a boolean if a field has been set.
func (o *ModificationDto) HasProcessDefinitionId() bool {
	if o != nil && o.ProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionId gets a reference to the given NullableString and assigns it to the ProcessDefinitionId field.
func (o *ModificationDto) SetProcessDefinitionId(v string) {
	o.ProcessDefinitionId.Set(&v)
}
// SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil
func (o *ModificationDto) SetProcessDefinitionIdNil() {
	o.ProcessDefinitionId.Set(nil)
}

// UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
func (o *ModificationDto) UnsetProcessDefinitionId() {
	o.ProcessDefinitionId.Unset()
}

// GetSkipCustomListeners returns the SkipCustomListeners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModificationDto) GetSkipCustomListeners() bool {
	if o == nil || IsNil(o.SkipCustomListeners.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipCustomListeners.Get()
}

// GetSkipCustomListenersOk returns a tuple with the SkipCustomListeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModificationDto) GetSkipCustomListenersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipCustomListeners.Get(), o.SkipCustomListeners.IsSet()
}

// HasSkipCustomListeners returns a boolean if a field has been set.
func (o *ModificationDto) HasSkipCustomListeners() bool {
	if o != nil && o.SkipCustomListeners.IsSet() {
		return true
	}

	return false
}

// SetSkipCustomListeners gets a reference to the given NullableBool and assigns it to the SkipCustomListeners field.
func (o *ModificationDto) SetSkipCustomListeners(v bool) {
	o.SkipCustomListeners.Set(&v)
}
// SetSkipCustomListenersNil sets the value for SkipCustomListeners to be an explicit nil
func (o *ModificationDto) SetSkipCustomListenersNil() {
	o.SkipCustomListeners.Set(nil)
}

// UnsetSkipCustomListeners ensures that no value is present for SkipCustomListeners, not even an explicit nil
func (o *ModificationDto) UnsetSkipCustomListeners() {
	o.SkipCustomListeners.Unset()
}

// GetSkipIoMappings returns the SkipIoMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModificationDto) GetSkipIoMappings() bool {
	if o == nil || IsNil(o.SkipIoMappings.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipIoMappings.Get()
}

// GetSkipIoMappingsOk returns a tuple with the SkipIoMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModificationDto) GetSkipIoMappingsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipIoMappings.Get(), o.SkipIoMappings.IsSet()
}

// HasSkipIoMappings returns a boolean if a field has been set.
func (o *ModificationDto) HasSkipIoMappings() bool {
	if o != nil && o.SkipIoMappings.IsSet() {
		return true
	}

	return false
}

// SetSkipIoMappings gets a reference to the given NullableBool and assigns it to the SkipIoMappings field.
func (o *ModificationDto) SetSkipIoMappings(v bool) {
	o.SkipIoMappings.Set(&v)
}
// SetSkipIoMappingsNil sets the value for SkipIoMappings to be an explicit nil
func (o *ModificationDto) SetSkipIoMappingsNil() {
	o.SkipIoMappings.Set(nil)
}

// UnsetSkipIoMappings ensures that no value is present for SkipIoMappings, not even an explicit nil
func (o *ModificationDto) UnsetSkipIoMappings() {
	o.SkipIoMappings.Unset()
}

// GetProcessInstanceIds returns the ProcessInstanceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModificationDto) GetProcessInstanceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProcessInstanceIds
}

// GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModificationDto) GetProcessInstanceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessInstanceIds) {
		return nil, false
	}
	return o.ProcessInstanceIds, true
}

// HasProcessInstanceIds returns a boolean if a field has been set.
func (o *ModificationDto) HasProcessInstanceIds() bool {
	if o != nil && !IsNil(o.ProcessInstanceIds) {
		return true
	}

	return false
}

// SetProcessInstanceIds gets a reference to the given []string and assigns it to the ProcessInstanceIds field.
func (o *ModificationDto) SetProcessInstanceIds(v []string) {
	o.ProcessInstanceIds = v
}

// GetProcessInstanceQuery returns the ProcessInstanceQuery field value if set, zero value otherwise.
func (o *ModificationDto) GetProcessInstanceQuery() ProcessInstanceQueryDto {
	if o == nil || IsNil(o.ProcessInstanceQuery) {
		var ret ProcessInstanceQueryDto
		return ret
	}
	return *o.ProcessInstanceQuery
}

// GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModificationDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool) {
	if o == nil || IsNil(o.ProcessInstanceQuery) {
		return nil, false
	}
	return o.ProcessInstanceQuery, true
}

// HasProcessInstanceQuery returns a boolean if a field has been set.
func (o *ModificationDto) HasProcessInstanceQuery() bool {
	if o != nil && !IsNil(o.ProcessInstanceQuery) {
		return true
	}

	return false
}

// SetProcessInstanceQuery gets a reference to the given ProcessInstanceQueryDto and assigns it to the ProcessInstanceQuery field.
func (o *ModificationDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto) {
	o.ProcessInstanceQuery = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModificationDto) GetInstructions() []MultipleProcessInstanceModificationInstructionDto {
	if o == nil {
		var ret []MultipleProcessInstanceModificationInstructionDto
		return ret
	}
	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModificationDto) GetInstructionsOk() ([]MultipleProcessInstanceModificationInstructionDto, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *ModificationDto) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given []MultipleProcessInstanceModificationInstructionDto and assigns it to the Instructions field.
func (o *ModificationDto) SetInstructions(v []MultipleProcessInstanceModificationInstructionDto) {
	o.Instructions = v
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModificationDto) GetAnnotation() string {
	if o == nil || IsNil(o.Annotation.Get()) {
		var ret string
		return ret
	}
	return *o.Annotation.Get()
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModificationDto) GetAnnotationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Annotation.Get(), o.Annotation.IsSet()
}

// HasAnnotation returns a boolean if a field has been set.
func (o *ModificationDto) HasAnnotation() bool {
	if o != nil && o.Annotation.IsSet() {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given NullableString and assigns it to the Annotation field.
func (o *ModificationDto) SetAnnotation(v string) {
	o.Annotation.Set(&v)
}
// SetAnnotationNil sets the value for Annotation to be an explicit nil
func (o *ModificationDto) SetAnnotationNil() {
	o.Annotation.Set(nil)
}

// UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil
func (o *ModificationDto) UnsetAnnotation() {
	o.Annotation.Unset()
}

func (o ModificationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModificationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProcessDefinitionId.IsSet() {
		toSerialize["processDefinitionId"] = o.ProcessDefinitionId.Get()
	}
	if o.SkipCustomListeners.IsSet() {
		toSerialize["skipCustomListeners"] = o.SkipCustomListeners.Get()
	}
	if o.SkipIoMappings.IsSet() {
		toSerialize["skipIoMappings"] = o.SkipIoMappings.Get()
	}
	if o.ProcessInstanceIds != nil {
		toSerialize["processInstanceIds"] = o.ProcessInstanceIds
	}
	if !IsNil(o.ProcessInstanceQuery) {
		toSerialize["processInstanceQuery"] = o.ProcessInstanceQuery
	}
	if o.Instructions != nil {
		toSerialize["instructions"] = o.Instructions
	}
	if o.Annotation.IsSet() {
		toSerialize["annotation"] = o.Annotation.Get()
	}
	return toSerialize, nil
}

type NullableModificationDto struct {
	value *ModificationDto
	isSet bool
}

func (v NullableModificationDto) Get() *ModificationDto {
	return v.value
}

func (v *NullableModificationDto) Set(val *ModificationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableModificationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableModificationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModificationDto(val *ModificationDto) *NullableModificationDto {
	return &NullableModificationDto{value: val, isSet: true}
}

func (v NullableModificationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModificationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


