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

// checks if the DecisionRequirementsDefinitionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DecisionRequirementsDefinitionDto{}

// DecisionRequirementsDefinitionDto struct for DecisionRequirementsDefinitionDto
type DecisionRequirementsDefinitionDto struct {
	// The id of the decision requirements definition.
	Id NullableString `json:"id,omitempty"`
	// The key of the decision requirements definition.
	Key NullableString `json:"key,omitempty"`
	// The category of the decision requirements definition.
	Category NullableString `json:"category,omitempty"`
	// The name of the decision requirements definition.
	Name NullableString `json:"name,omitempty"`
	// The version of the decision requirements definition that the engine assigned to it.
	Version NullableInt32 `json:"version,omitempty"`
	// The file name of the decision requirements definition.
	Resource NullableString `json:"resource,omitempty"`
	// The deployment id of the decision requirements definition.
	DeploymentId NullableString `json:"deploymentId,omitempty"`
	// The tenant id of the decision requirements definition.
	TenantId NullableString `json:"tenantId,omitempty"`
}

// NewDecisionRequirementsDefinitionDto instantiates a new DecisionRequirementsDefinitionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecisionRequirementsDefinitionDto() *DecisionRequirementsDefinitionDto {
	this := DecisionRequirementsDefinitionDto{}
	return &this
}

// NewDecisionRequirementsDefinitionDtoWithDefaults instantiates a new DecisionRequirementsDefinitionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecisionRequirementsDefinitionDtoWithDefaults() *DecisionRequirementsDefinitionDto {
	this := DecisionRequirementsDefinitionDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecisionRequirementsDefinitionDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecisionRequirementsDefinitionDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DecisionRequirementsDefinitionDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DecisionRequirementsDefinitionDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DecisionRequirementsDefinitionDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DecisionRequirementsDefinitionDto) UnsetId() {
	o.Id.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecisionRequirementsDefinitionDto) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecisionRequirementsDefinitionDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *DecisionRequirementsDefinitionDto) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *DecisionRequirementsDefinitionDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *DecisionRequirementsDefinitionDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *DecisionRequirementsDefinitionDto) UnsetKey() {
	o.Key.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecisionRequirementsDefinitionDto) GetCategory() string {
	if o == nil || IsNil(o.Category.Get()) {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecisionRequirementsDefinitionDto) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *DecisionRequirementsDefinitionDto) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *DecisionRequirementsDefinitionDto) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *DecisionRequirementsDefinitionDto) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *DecisionRequirementsDefinitionDto) UnsetCategory() {
	o.Category.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecisionRequirementsDefinitionDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecisionRequirementsDefinitionDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DecisionRequirementsDefinitionDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DecisionRequirementsDefinitionDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DecisionRequirementsDefinitionDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DecisionRequirementsDefinitionDto) UnsetName() {
	o.Name.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecisionRequirementsDefinitionDto) GetVersion() int32 {
	if o == nil || IsNil(o.Version.Get()) {
		var ret int32
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecisionRequirementsDefinitionDto) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *DecisionRequirementsDefinitionDto) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableInt32 and assigns it to the Version field.
func (o *DecisionRequirementsDefinitionDto) SetVersion(v int32) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *DecisionRequirementsDefinitionDto) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *DecisionRequirementsDefinitionDto) UnsetVersion() {
	o.Version.Unset()
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecisionRequirementsDefinitionDto) GetResource() string {
	if o == nil || IsNil(o.Resource.Get()) {
		var ret string
		return ret
	}
	return *o.Resource.Get()
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecisionRequirementsDefinitionDto) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resource.Get(), o.Resource.IsSet()
}

// HasResource returns a boolean if a field has been set.
func (o *DecisionRequirementsDefinitionDto) HasResource() bool {
	if o != nil && o.Resource.IsSet() {
		return true
	}

	return false
}

// SetResource gets a reference to the given NullableString and assigns it to the Resource field.
func (o *DecisionRequirementsDefinitionDto) SetResource(v string) {
	o.Resource.Set(&v)
}
// SetResourceNil sets the value for Resource to be an explicit nil
func (o *DecisionRequirementsDefinitionDto) SetResourceNil() {
	o.Resource.Set(nil)
}

// UnsetResource ensures that no value is present for Resource, not even an explicit nil
func (o *DecisionRequirementsDefinitionDto) UnsetResource() {
	o.Resource.Unset()
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecisionRequirementsDefinitionDto) GetDeploymentId() string {
	if o == nil || IsNil(o.DeploymentId.Get()) {
		var ret string
		return ret
	}
	return *o.DeploymentId.Get()
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecisionRequirementsDefinitionDto) GetDeploymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeploymentId.Get(), o.DeploymentId.IsSet()
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *DecisionRequirementsDefinitionDto) HasDeploymentId() bool {
	if o != nil && o.DeploymentId.IsSet() {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given NullableString and assigns it to the DeploymentId field.
func (o *DecisionRequirementsDefinitionDto) SetDeploymentId(v string) {
	o.DeploymentId.Set(&v)
}
// SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil
func (o *DecisionRequirementsDefinitionDto) SetDeploymentIdNil() {
	o.DeploymentId.Set(nil)
}

// UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
func (o *DecisionRequirementsDefinitionDto) UnsetDeploymentId() {
	o.DeploymentId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DecisionRequirementsDefinitionDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DecisionRequirementsDefinitionDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *DecisionRequirementsDefinitionDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *DecisionRequirementsDefinitionDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *DecisionRequirementsDefinitionDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *DecisionRequirementsDefinitionDto) UnsetTenantId() {
	o.TenantId.Unset()
}

func (o DecisionRequirementsDefinitionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DecisionRequirementsDefinitionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Resource.IsSet() {
		toSerialize["resource"] = o.Resource.Get()
	}
	if o.DeploymentId.IsSet() {
		toSerialize["deploymentId"] = o.DeploymentId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	return toSerialize, nil
}

type NullableDecisionRequirementsDefinitionDto struct {
	value *DecisionRequirementsDefinitionDto
	isSet bool
}

func (v NullableDecisionRequirementsDefinitionDto) Get() *DecisionRequirementsDefinitionDto {
	return v.value
}

func (v *NullableDecisionRequirementsDefinitionDto) Set(val *DecisionRequirementsDefinitionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDecisionRequirementsDefinitionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDecisionRequirementsDefinitionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecisionRequirementsDefinitionDto(val *DecisionRequirementsDefinitionDto) *NullableDecisionRequirementsDefinitionDto {
	return &NullableDecisionRequirementsDefinitionDto{value: val, isSet: true}
}

func (v NullableDecisionRequirementsDefinitionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecisionRequirementsDefinitionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


