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

// checks if the DeleteHistoricDecisionInstancesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteHistoricDecisionInstancesDto{}

// DeleteHistoricDecisionInstancesDto struct for DeleteHistoricDecisionInstancesDto
type DeleteHistoricDecisionInstancesDto struct {
	// A list of historic decision instance ids to delete.
	HistoricDecisionInstanceIds []string `json:"historicDecisionInstanceIds,omitempty"`
	HistoricDecisionInstanceQuery *HistoricDecisionInstanceQueryDto `json:"historicDecisionInstanceQuery,omitempty"`
	// A string with delete reason.
	DeleteReason NullableString `json:"deleteReason,omitempty"`
}

// NewDeleteHistoricDecisionInstancesDto instantiates a new DeleteHistoricDecisionInstancesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteHistoricDecisionInstancesDto() *DeleteHistoricDecisionInstancesDto {
	this := DeleteHistoricDecisionInstancesDto{}
	return &this
}

// NewDeleteHistoricDecisionInstancesDtoWithDefaults instantiates a new DeleteHistoricDecisionInstancesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteHistoricDecisionInstancesDtoWithDefaults() *DeleteHistoricDecisionInstancesDto {
	this := DeleteHistoricDecisionInstancesDto{}
	return &this
}

// GetHistoricDecisionInstanceIds returns the HistoricDecisionInstanceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteHistoricDecisionInstancesDto) GetHistoricDecisionInstanceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.HistoricDecisionInstanceIds
}

// GetHistoricDecisionInstanceIdsOk returns a tuple with the HistoricDecisionInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteHistoricDecisionInstancesDto) GetHistoricDecisionInstanceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.HistoricDecisionInstanceIds) {
		return nil, false
	}
	return o.HistoricDecisionInstanceIds, true
}

// HasHistoricDecisionInstanceIds returns a boolean if a field has been set.
func (o *DeleteHistoricDecisionInstancesDto) HasHistoricDecisionInstanceIds() bool {
	if o != nil && !IsNil(o.HistoricDecisionInstanceIds) {
		return true
	}

	return false
}

// SetHistoricDecisionInstanceIds gets a reference to the given []string and assigns it to the HistoricDecisionInstanceIds field.
func (o *DeleteHistoricDecisionInstancesDto) SetHistoricDecisionInstanceIds(v []string) {
	o.HistoricDecisionInstanceIds = v
}

// GetHistoricDecisionInstanceQuery returns the HistoricDecisionInstanceQuery field value if set, zero value otherwise.
func (o *DeleteHistoricDecisionInstancesDto) GetHistoricDecisionInstanceQuery() HistoricDecisionInstanceQueryDto {
	if o == nil || IsNil(o.HistoricDecisionInstanceQuery) {
		var ret HistoricDecisionInstanceQueryDto
		return ret
	}
	return *o.HistoricDecisionInstanceQuery
}

// GetHistoricDecisionInstanceQueryOk returns a tuple with the HistoricDecisionInstanceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteHistoricDecisionInstancesDto) GetHistoricDecisionInstanceQueryOk() (*HistoricDecisionInstanceQueryDto, bool) {
	if o == nil || IsNil(o.HistoricDecisionInstanceQuery) {
		return nil, false
	}
	return o.HistoricDecisionInstanceQuery, true
}

// HasHistoricDecisionInstanceQuery returns a boolean if a field has been set.
func (o *DeleteHistoricDecisionInstancesDto) HasHistoricDecisionInstanceQuery() bool {
	if o != nil && !IsNil(o.HistoricDecisionInstanceQuery) {
		return true
	}

	return false
}

// SetHistoricDecisionInstanceQuery gets a reference to the given HistoricDecisionInstanceQueryDto and assigns it to the HistoricDecisionInstanceQuery field.
func (o *DeleteHistoricDecisionInstancesDto) SetHistoricDecisionInstanceQuery(v HistoricDecisionInstanceQueryDto) {
	o.HistoricDecisionInstanceQuery = &v
}

// GetDeleteReason returns the DeleteReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteHistoricDecisionInstancesDto) GetDeleteReason() string {
	if o == nil || IsNil(o.DeleteReason.Get()) {
		var ret string
		return ret
	}
	return *o.DeleteReason.Get()
}

// GetDeleteReasonOk returns a tuple with the DeleteReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteHistoricDecisionInstancesDto) GetDeleteReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeleteReason.Get(), o.DeleteReason.IsSet()
}

// HasDeleteReason returns a boolean if a field has been set.
func (o *DeleteHistoricDecisionInstancesDto) HasDeleteReason() bool {
	if o != nil && o.DeleteReason.IsSet() {
		return true
	}

	return false
}

// SetDeleteReason gets a reference to the given NullableString and assigns it to the DeleteReason field.
func (o *DeleteHistoricDecisionInstancesDto) SetDeleteReason(v string) {
	o.DeleteReason.Set(&v)
}
// SetDeleteReasonNil sets the value for DeleteReason to be an explicit nil
func (o *DeleteHistoricDecisionInstancesDto) SetDeleteReasonNil() {
	o.DeleteReason.Set(nil)
}

// UnsetDeleteReason ensures that no value is present for DeleteReason, not even an explicit nil
func (o *DeleteHistoricDecisionInstancesDto) UnsetDeleteReason() {
	o.DeleteReason.Unset()
}

func (o DeleteHistoricDecisionInstancesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteHistoricDecisionInstancesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.HistoricDecisionInstanceIds != nil {
		toSerialize["historicDecisionInstanceIds"] = o.HistoricDecisionInstanceIds
	}
	if !IsNil(o.HistoricDecisionInstanceQuery) {
		toSerialize["historicDecisionInstanceQuery"] = o.HistoricDecisionInstanceQuery
	}
	if o.DeleteReason.IsSet() {
		toSerialize["deleteReason"] = o.DeleteReason.Get()
	}
	return toSerialize, nil
}

type NullableDeleteHistoricDecisionInstancesDto struct {
	value *DeleteHistoricDecisionInstancesDto
	isSet bool
}

func (v NullableDeleteHistoricDecisionInstancesDto) Get() *DeleteHistoricDecisionInstancesDto {
	return v.value
}

func (v *NullableDeleteHistoricDecisionInstancesDto) Set(val *DeleteHistoricDecisionInstancesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteHistoricDecisionInstancesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteHistoricDecisionInstancesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteHistoricDecisionInstancesDto(val *DeleteHistoricDecisionInstancesDto) *NullableDeleteHistoricDecisionInstancesDto {
	return &NullableDeleteHistoricDecisionInstancesDto{value: val, isSet: true}
}

func (v NullableDeleteHistoricDecisionInstancesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteHistoricDecisionInstancesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

