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

// checks if the ProcessInstanceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessInstanceDto{}

// ProcessInstanceDto struct for ProcessInstanceDto
type ProcessInstanceDto struct {
	// The links associated to this resource, with `method`, `href` and `rel`.
	Links []AtomLink `json:"links,omitempty"`
	// The id of the process instance.
	Id NullableString `json:"id,omitempty"`
	// The id of the process definition that this process instance belongs to.
	DefinitionId NullableString `json:"definitionId,omitempty"`
	// The business key of the process instance.
	BusinessKey NullableString `json:"businessKey,omitempty"`
	// The id of the case instance associated with the process instance.
	CaseInstanceId NullableString `json:"caseInstanceId,omitempty"`
	// A flag indicating whether the process instance has ended or not. Deprecated: will always be false!
	// Deprecated
	Ended NullableBool `json:"ended,omitempty"`
	// A flag indicating whether the process instance is suspended or not.
	Suspended NullableBool `json:"suspended,omitempty"`
	// The tenant id of the process instance.
	TenantId NullableString `json:"tenantId,omitempty"`
}

// NewProcessInstanceDto instantiates a new ProcessInstanceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessInstanceDto() *ProcessInstanceDto {
	this := ProcessInstanceDto{}
	return &this
}

// NewProcessInstanceDtoWithDefaults instantiates a new ProcessInstanceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessInstanceDtoWithDefaults() *ProcessInstanceDto {
	this := ProcessInstanceDto{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessInstanceDto) GetLinks() []AtomLink {
	if o == nil {
		var ret []AtomLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessInstanceDto) GetLinksOk() ([]AtomLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProcessInstanceDto) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []AtomLink and assigns it to the Links field.
func (o *ProcessInstanceDto) SetLinks(v []AtomLink) {
	o.Links = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessInstanceDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessInstanceDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ProcessInstanceDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ProcessInstanceDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ProcessInstanceDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ProcessInstanceDto) UnsetId() {
	o.Id.Unset()
}

// GetDefinitionId returns the DefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessInstanceDto) GetDefinitionId() string {
	if o == nil || IsNil(o.DefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.DefinitionId.Get()
}

// GetDefinitionIdOk returns a tuple with the DefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessInstanceDto) GetDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefinitionId.Get(), o.DefinitionId.IsSet()
}

// HasDefinitionId returns a boolean if a field has been set.
func (o *ProcessInstanceDto) HasDefinitionId() bool {
	if o != nil && o.DefinitionId.IsSet() {
		return true
	}

	return false
}

// SetDefinitionId gets a reference to the given NullableString and assigns it to the DefinitionId field.
func (o *ProcessInstanceDto) SetDefinitionId(v string) {
	o.DefinitionId.Set(&v)
}
// SetDefinitionIdNil sets the value for DefinitionId to be an explicit nil
func (o *ProcessInstanceDto) SetDefinitionIdNil() {
	o.DefinitionId.Set(nil)
}

// UnsetDefinitionId ensures that no value is present for DefinitionId, not even an explicit nil
func (o *ProcessInstanceDto) UnsetDefinitionId() {
	o.DefinitionId.Unset()
}

// GetBusinessKey returns the BusinessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessInstanceDto) GetBusinessKey() string {
	if o == nil || IsNil(o.BusinessKey.Get()) {
		var ret string
		return ret
	}
	return *o.BusinessKey.Get()
}

// GetBusinessKeyOk returns a tuple with the BusinessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessInstanceDto) GetBusinessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessKey.Get(), o.BusinessKey.IsSet()
}

// HasBusinessKey returns a boolean if a field has been set.
func (o *ProcessInstanceDto) HasBusinessKey() bool {
	if o != nil && o.BusinessKey.IsSet() {
		return true
	}

	return false
}

// SetBusinessKey gets a reference to the given NullableString and assigns it to the BusinessKey field.
func (o *ProcessInstanceDto) SetBusinessKey(v string) {
	o.BusinessKey.Set(&v)
}
// SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil
func (o *ProcessInstanceDto) SetBusinessKeyNil() {
	o.BusinessKey.Set(nil)
}

// UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
func (o *ProcessInstanceDto) UnsetBusinessKey() {
	o.BusinessKey.Unset()
}

// GetCaseInstanceId returns the CaseInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessInstanceDto) GetCaseInstanceId() string {
	if o == nil || IsNil(o.CaseInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.CaseInstanceId.Get()
}

// GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessInstanceDto) GetCaseInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaseInstanceId.Get(), o.CaseInstanceId.IsSet()
}

// HasCaseInstanceId returns a boolean if a field has been set.
func (o *ProcessInstanceDto) HasCaseInstanceId() bool {
	if o != nil && o.CaseInstanceId.IsSet() {
		return true
	}

	return false
}

// SetCaseInstanceId gets a reference to the given NullableString and assigns it to the CaseInstanceId field.
func (o *ProcessInstanceDto) SetCaseInstanceId(v string) {
	o.CaseInstanceId.Set(&v)
}
// SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil
func (o *ProcessInstanceDto) SetCaseInstanceIdNil() {
	o.CaseInstanceId.Set(nil)
}

// UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
func (o *ProcessInstanceDto) UnsetCaseInstanceId() {
	o.CaseInstanceId.Unset()
}

// GetEnded returns the Ended field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ProcessInstanceDto) GetEnded() bool {
	if o == nil || IsNil(o.Ended.Get()) {
		var ret bool
		return ret
	}
	return *o.Ended.Get()
}

// GetEndedOk returns a tuple with the Ended field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ProcessInstanceDto) GetEndedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ended.Get(), o.Ended.IsSet()
}

// HasEnded returns a boolean if a field has been set.
func (o *ProcessInstanceDto) HasEnded() bool {
	if o != nil && o.Ended.IsSet() {
		return true
	}

	return false
}

// SetEnded gets a reference to the given NullableBool and assigns it to the Ended field.
// Deprecated
func (o *ProcessInstanceDto) SetEnded(v bool) {
	o.Ended.Set(&v)
}
// SetEndedNil sets the value for Ended to be an explicit nil
func (o *ProcessInstanceDto) SetEndedNil() {
	o.Ended.Set(nil)
}

// UnsetEnded ensures that no value is present for Ended, not even an explicit nil
func (o *ProcessInstanceDto) UnsetEnded() {
	o.Ended.Unset()
}

// GetSuspended returns the Suspended field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessInstanceDto) GetSuspended() bool {
	if o == nil || IsNil(o.Suspended.Get()) {
		var ret bool
		return ret
	}
	return *o.Suspended.Get()
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessInstanceDto) GetSuspendedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suspended.Get(), o.Suspended.IsSet()
}

// HasSuspended returns a boolean if a field has been set.
func (o *ProcessInstanceDto) HasSuspended() bool {
	if o != nil && o.Suspended.IsSet() {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given NullableBool and assigns it to the Suspended field.
func (o *ProcessInstanceDto) SetSuspended(v bool) {
	o.Suspended.Set(&v)
}
// SetSuspendedNil sets the value for Suspended to be an explicit nil
func (o *ProcessInstanceDto) SetSuspendedNil() {
	o.Suspended.Set(nil)
}

// UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
func (o *ProcessInstanceDto) UnsetSuspended() {
	o.Suspended.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessInstanceDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessInstanceDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProcessInstanceDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *ProcessInstanceDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *ProcessInstanceDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *ProcessInstanceDto) UnsetTenantId() {
	o.TenantId.Unset()
}

func (o ProcessInstanceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessInstanceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.DefinitionId.IsSet() {
		toSerialize["definitionId"] = o.DefinitionId.Get()
	}
	if o.BusinessKey.IsSet() {
		toSerialize["businessKey"] = o.BusinessKey.Get()
	}
	if o.CaseInstanceId.IsSet() {
		toSerialize["caseInstanceId"] = o.CaseInstanceId.Get()
	}
	if o.Ended.IsSet() {
		toSerialize["ended"] = o.Ended.Get()
	}
	if o.Suspended.IsSet() {
		toSerialize["suspended"] = o.Suspended.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	return toSerialize, nil
}

type NullableProcessInstanceDto struct {
	value *ProcessInstanceDto
	isSet bool
}

func (v NullableProcessInstanceDto) Get() *ProcessInstanceDto {
	return v.value
}

func (v *NullableProcessInstanceDto) Set(val *ProcessInstanceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessInstanceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessInstanceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessInstanceDto(val *ProcessInstanceDto) *NullableProcessInstanceDto {
	return &NullableProcessInstanceDto{value: val, isSet: true}
}

func (v NullableProcessInstanceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessInstanceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


