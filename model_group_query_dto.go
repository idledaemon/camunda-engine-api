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

// checks if the GroupQueryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupQueryDto{}

// GroupQueryDto A group instance query which defines a list of group instances
type GroupQueryDto struct {
	// Filter by the id of the group.
	Id NullableString `json:"id,omitempty"`
	// Filter by a JSON string array of group ids.
	IdIn []string `json:"idIn,omitempty"`
	// Filter by the name of the group.
	Name NullableString `json:"name,omitempty"`
	// Filter by the name that the parameter is a substring of.
	NameLike NullableString `json:"nameLike,omitempty"`
	// Filter by the type of the group.
	Type NullableString `json:"type,omitempty"`
	// Only retrieve groups where the given user id is a member of.
	Member NullableString `json:"member,omitempty"`
	// Only retrieve groups which are members of the given tenant.
	MemberOfTenant NullableString `json:"memberOfTenant,omitempty"`
	// Apply sorting of the result
	Sorting []GroupQueryDtoSortingInner `json:"sorting,omitempty"`
}

// NewGroupQueryDto instantiates a new GroupQueryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupQueryDto() *GroupQueryDto {
	this := GroupQueryDto{}
	return &this
}

// NewGroupQueryDtoWithDefaults instantiates a new GroupQueryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupQueryDtoWithDefaults() *GroupQueryDto {
	this := GroupQueryDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupQueryDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupQueryDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GroupQueryDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *GroupQueryDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GroupQueryDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GroupQueryDto) UnsetId() {
	o.Id.Unset()
}

// GetIdIn returns the IdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupQueryDto) GetIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IdIn
}

// GetIdInOk returns a tuple with the IdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupQueryDto) GetIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.IdIn) {
		return nil, false
	}
	return o.IdIn, true
}

// HasIdIn returns a boolean if a field has been set.
func (o *GroupQueryDto) HasIdIn() bool {
	if o != nil && !IsNil(o.IdIn) {
		return true
	}

	return false
}

// SetIdIn gets a reference to the given []string and assigns it to the IdIn field.
func (o *GroupQueryDto) SetIdIn(v []string) {
	o.IdIn = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupQueryDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupQueryDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *GroupQueryDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *GroupQueryDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *GroupQueryDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GroupQueryDto) UnsetName() {
	o.Name.Unset()
}

// GetNameLike returns the NameLike field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupQueryDto) GetNameLike() string {
	if o == nil || IsNil(o.NameLike.Get()) {
		var ret string
		return ret
	}
	return *o.NameLike.Get()
}

// GetNameLikeOk returns a tuple with the NameLike field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupQueryDto) GetNameLikeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameLike.Get(), o.NameLike.IsSet()
}

// HasNameLike returns a boolean if a field has been set.
func (o *GroupQueryDto) HasNameLike() bool {
	if o != nil && o.NameLike.IsSet() {
		return true
	}

	return false
}

// SetNameLike gets a reference to the given NullableString and assigns it to the NameLike field.
func (o *GroupQueryDto) SetNameLike(v string) {
	o.NameLike.Set(&v)
}
// SetNameLikeNil sets the value for NameLike to be an explicit nil
func (o *GroupQueryDto) SetNameLikeNil() {
	o.NameLike.Set(nil)
}

// UnsetNameLike ensures that no value is present for NameLike, not even an explicit nil
func (o *GroupQueryDto) UnsetNameLike() {
	o.NameLike.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupQueryDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupQueryDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *GroupQueryDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *GroupQueryDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *GroupQueryDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *GroupQueryDto) UnsetType() {
	o.Type.Unset()
}

// GetMember returns the Member field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupQueryDto) GetMember() string {
	if o == nil || IsNil(o.Member.Get()) {
		var ret string
		return ret
	}
	return *o.Member.Get()
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupQueryDto) GetMemberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Member.Get(), o.Member.IsSet()
}

// HasMember returns a boolean if a field has been set.
func (o *GroupQueryDto) HasMember() bool {
	if o != nil && o.Member.IsSet() {
		return true
	}

	return false
}

// SetMember gets a reference to the given NullableString and assigns it to the Member field.
func (o *GroupQueryDto) SetMember(v string) {
	o.Member.Set(&v)
}
// SetMemberNil sets the value for Member to be an explicit nil
func (o *GroupQueryDto) SetMemberNil() {
	o.Member.Set(nil)
}

// UnsetMember ensures that no value is present for Member, not even an explicit nil
func (o *GroupQueryDto) UnsetMember() {
	o.Member.Unset()
}

// GetMemberOfTenant returns the MemberOfTenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupQueryDto) GetMemberOfTenant() string {
	if o == nil || IsNil(o.MemberOfTenant.Get()) {
		var ret string
		return ret
	}
	return *o.MemberOfTenant.Get()
}

// GetMemberOfTenantOk returns a tuple with the MemberOfTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupQueryDto) GetMemberOfTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemberOfTenant.Get(), o.MemberOfTenant.IsSet()
}

// HasMemberOfTenant returns a boolean if a field has been set.
func (o *GroupQueryDto) HasMemberOfTenant() bool {
	if o != nil && o.MemberOfTenant.IsSet() {
		return true
	}

	return false
}

// SetMemberOfTenant gets a reference to the given NullableString and assigns it to the MemberOfTenant field.
func (o *GroupQueryDto) SetMemberOfTenant(v string) {
	o.MemberOfTenant.Set(&v)
}
// SetMemberOfTenantNil sets the value for MemberOfTenant to be an explicit nil
func (o *GroupQueryDto) SetMemberOfTenantNil() {
	o.MemberOfTenant.Set(nil)
}

// UnsetMemberOfTenant ensures that no value is present for MemberOfTenant, not even an explicit nil
func (o *GroupQueryDto) UnsetMemberOfTenant() {
	o.MemberOfTenant.Unset()
}

// GetSorting returns the Sorting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupQueryDto) GetSorting() []GroupQueryDtoSortingInner {
	if o == nil {
		var ret []GroupQueryDtoSortingInner
		return ret
	}
	return o.Sorting
}

// GetSortingOk returns a tuple with the Sorting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupQueryDto) GetSortingOk() ([]GroupQueryDtoSortingInner, bool) {
	if o == nil || IsNil(o.Sorting) {
		return nil, false
	}
	return o.Sorting, true
}

// HasSorting returns a boolean if a field has been set.
func (o *GroupQueryDto) HasSorting() bool {
	if o != nil && !IsNil(o.Sorting) {
		return true
	}

	return false
}

// SetSorting gets a reference to the given []GroupQueryDtoSortingInner and assigns it to the Sorting field.
func (o *GroupQueryDto) SetSorting(v []GroupQueryDtoSortingInner) {
	o.Sorting = v
}

func (o GroupQueryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupQueryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IdIn != nil {
		toSerialize["idIn"] = o.IdIn
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.NameLike.IsSet() {
		toSerialize["nameLike"] = o.NameLike.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Member.IsSet() {
		toSerialize["member"] = o.Member.Get()
	}
	if o.MemberOfTenant.IsSet() {
		toSerialize["memberOfTenant"] = o.MemberOfTenant.Get()
	}
	if o.Sorting != nil {
		toSerialize["sorting"] = o.Sorting
	}
	return toSerialize, nil
}

type NullableGroupQueryDto struct {
	value *GroupQueryDto
	isSet bool
}

func (v NullableGroupQueryDto) Get() *GroupQueryDto {
	return v.value
}

func (v *NullableGroupQueryDto) Set(val *GroupQueryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupQueryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupQueryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupQueryDto(val *GroupQueryDto) *NullableGroupQueryDto {
	return &NullableGroupQueryDto{value: val, isSet: true}
}

func (v NullableGroupQueryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupQueryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


