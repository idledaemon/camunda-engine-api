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

// checks if the CreateFilterDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFilterDto{}

// CreateFilterDto struct for CreateFilterDto
type CreateFilterDto struct {
	// The resource type of the filter.
	ResourceType NullableString `json:"resourceType,omitempty"`
	// The name of the filter.
	Name NullableString `json:"name,omitempty"`
	// The user id of the owner of the filter.
	Owner NullableString `json:"owner,omitempty"`
	// The query of the filter as a JSON object.
	Query map[string]interface{} `json:"query,omitempty"`
	// The properties of a filter as a JSON object.
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// NewCreateFilterDto instantiates a new CreateFilterDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFilterDto() *CreateFilterDto {
	this := CreateFilterDto{}
	return &this
}

// NewCreateFilterDtoWithDefaults instantiates a new CreateFilterDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFilterDtoWithDefaults() *CreateFilterDto {
	this := CreateFilterDto{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFilterDto) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceType.Get()
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFilterDto) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType.Get(), o.ResourceType.IsSet()
}

// HasResourceType returns a boolean if a field has been set.
func (o *CreateFilterDto) HasResourceType() bool {
	if o != nil && o.ResourceType.IsSet() {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given NullableString and assigns it to the ResourceType field.
func (o *CreateFilterDto) SetResourceType(v string) {
	o.ResourceType.Set(&v)
}
// SetResourceTypeNil sets the value for ResourceType to be an explicit nil
func (o *CreateFilterDto) SetResourceTypeNil() {
	o.ResourceType.Set(nil)
}

// UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
func (o *CreateFilterDto) UnsetResourceType() {
	o.ResourceType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFilterDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFilterDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateFilterDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateFilterDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateFilterDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateFilterDto) UnsetName() {
	o.Name.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFilterDto) GetOwner() string {
	if o == nil || IsNil(o.Owner.Get()) {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFilterDto) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *CreateFilterDto) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *CreateFilterDto) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *CreateFilterDto) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *CreateFilterDto) UnsetOwner() {
	o.Owner.Unset()
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *CreateFilterDto) GetQuery() map[string]interface{} {
	if o == nil || IsNil(o.Query) {
		var ret map[string]interface{}
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFilterDto) GetQueryOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Query) {
		return map[string]interface{}{}, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *CreateFilterDto) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given map[string]interface{} and assigns it to the Query field.
func (o *CreateFilterDto) SetQuery(v map[string]interface{}) {
	o.Query = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CreateFilterDto) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFilterDto) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CreateFilterDto) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *CreateFilterDto) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o CreateFilterDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFilterDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceType.IsSet() {
		toSerialize["resourceType"] = o.ResourceType.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableCreateFilterDto struct {
	value *CreateFilterDto
	isSet bool
}

func (v NullableCreateFilterDto) Get() *CreateFilterDto {
	return v.value
}

func (v *NullableCreateFilterDto) Set(val *CreateFilterDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFilterDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFilterDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFilterDto(val *CreateFilterDto) *NullableCreateFilterDto {
	return &NullableCreateFilterDto{value: val, isSet: true}
}

func (v NullableCreateFilterDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFilterDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


