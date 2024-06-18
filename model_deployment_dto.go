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

// checks if the DeploymentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentDto{}

// DeploymentDto struct for DeploymentDto
type DeploymentDto struct {
	// The links associated to this resource, with `method`, `href` and `rel`.
	Links []AtomLink `json:"links,omitempty"`
	// The id of the deployment.
	Id NullableString `json:"id,omitempty"`
	// The tenant id of the deployment.
	TenantId NullableString `json:"tenantId,omitempty"`
	// The time when the deployment was created.
	DeploymentTime NullableTime `json:"deploymentTime,omitempty"`
	// The source of the deployment.
	Source NullableString `json:"source,omitempty"`
	// The name of the deployment.
	Name NullableString `json:"name,omitempty"`
}

// NewDeploymentDto instantiates a new DeploymentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentDto() *DeploymentDto {
	this := DeploymentDto{}
	return &this
}

// NewDeploymentDtoWithDefaults instantiates a new DeploymentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentDtoWithDefaults() *DeploymentDto {
	this := DeploymentDto{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentDto) GetLinks() []AtomLink {
	if o == nil {
		var ret []AtomLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentDto) GetLinksOk() ([]AtomLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DeploymentDto) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []AtomLink and assigns it to the Links field.
func (o *DeploymentDto) SetLinks(v []AtomLink) {
	o.Links = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DeploymentDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DeploymentDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DeploymentDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DeploymentDto) UnsetId() {
	o.Id.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *DeploymentDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *DeploymentDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *DeploymentDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *DeploymentDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetDeploymentTime returns the DeploymentTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentDto) GetDeploymentTime() time.Time {
	if o == nil || IsNil(o.DeploymentTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DeploymentTime.Get()
}

// GetDeploymentTimeOk returns a tuple with the DeploymentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentDto) GetDeploymentTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeploymentTime.Get(), o.DeploymentTime.IsSet()
}

// HasDeploymentTime returns a boolean if a field has been set.
func (o *DeploymentDto) HasDeploymentTime() bool {
	if o != nil && o.DeploymentTime.IsSet() {
		return true
	}

	return false
}

// SetDeploymentTime gets a reference to the given NullableTime and assigns it to the DeploymentTime field.
func (o *DeploymentDto) SetDeploymentTime(v time.Time) {
	o.DeploymentTime.Set(&v)
}
// SetDeploymentTimeNil sets the value for DeploymentTime to be an explicit nil
func (o *DeploymentDto) SetDeploymentTimeNil() {
	o.DeploymentTime.Set(nil)
}

// UnsetDeploymentTime ensures that no value is present for DeploymentTime, not even an explicit nil
func (o *DeploymentDto) UnsetDeploymentTime() {
	o.DeploymentTime.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentDto) GetSource() string {
	if o == nil || IsNil(o.Source.Get()) {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentDto) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *DeploymentDto) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableString and assigns it to the Source field.
func (o *DeploymentDto) SetSource(v string) {
	o.Source.Set(&v)
}
// SetSourceNil sets the value for Source to be an explicit nil
func (o *DeploymentDto) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *DeploymentDto) UnsetSource() {
	o.Source.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DeploymentDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DeploymentDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DeploymentDto) UnsetName() {
	o.Name.Unset()
}

func (o DeploymentDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.DeploymentTime.IsSet() {
		toSerialize["deploymentTime"] = o.DeploymentTime.Get()
	}
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableDeploymentDto struct {
	value *DeploymentDto
	isSet bool
}

func (v NullableDeploymentDto) Get() *DeploymentDto {
	return v.value
}

func (v *NullableDeploymentDto) Set(val *DeploymentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentDto(val *DeploymentDto) *NullableDeploymentDto {
	return &NullableDeploymentDto{value: val, isSet: true}
}

func (v NullableDeploymentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

