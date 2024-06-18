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

// checks if the CleanableHistoricDecisionInstanceReportResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CleanableHistoricDecisionInstanceReportResultDto{}

// CleanableHistoricDecisionInstanceReportResultDto struct for CleanableHistoricDecisionInstanceReportResultDto
type CleanableHistoricDecisionInstanceReportResultDto struct {
	// The id of the decision definition.
	DecisionDefinitionId NullableString `json:"decisionDefinitionId,omitempty"`
	// The key of the decision definition.
	DecisionDefinitionKey NullableString `json:"decisionDefinitionKey,omitempty"`
	// The name of the decision definition.
	DecisionDefinitionName NullableString `json:"decisionDefinitionName,omitempty"`
	// The version of the decision definition.
	DecisionDefinitionVersion NullableInt32 `json:"decisionDefinitionVersion,omitempty"`
	// The history time to live of the decision definition.
	HistoryTimeToLive NullableInt32 `json:"historyTimeToLive,omitempty"`
	// The count of the finished historic decision instances.
	FinishedDecisionInstanceCount NullableInt64 `json:"finishedDecisionInstanceCount,omitempty"`
	// The count of the cleanable historic decision instances, referring to history time to live.
	CleanableDecisionInstanceCount NullableInt64 `json:"cleanableDecisionInstanceCount,omitempty"`
	// The tenant id of the decision definition.
	TenantId NullableString `json:"tenantId,omitempty"`
}

// NewCleanableHistoricDecisionInstanceReportResultDto instantiates a new CleanableHistoricDecisionInstanceReportResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCleanableHistoricDecisionInstanceReportResultDto() *CleanableHistoricDecisionInstanceReportResultDto {
	this := CleanableHistoricDecisionInstanceReportResultDto{}
	return &this
}

// NewCleanableHistoricDecisionInstanceReportResultDtoWithDefaults instantiates a new CleanableHistoricDecisionInstanceReportResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCleanableHistoricDecisionInstanceReportResultDtoWithDefaults() *CleanableHistoricDecisionInstanceReportResultDto {
	this := CleanableHistoricDecisionInstanceReportResultDto{}
	return &this
}

// GetDecisionDefinitionId returns the DecisionDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionId() string {
	if o == nil || IsNil(o.DecisionDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.DecisionDefinitionId.Get()
}

// GetDecisionDefinitionIdOk returns a tuple with the DecisionDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecisionDefinitionId.Get(), o.DecisionDefinitionId.IsSet()
}

// HasDecisionDefinitionId returns a boolean if a field has been set.
func (o *CleanableHistoricDecisionInstanceReportResultDto) HasDecisionDefinitionId() bool {
	if o != nil && o.DecisionDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetDecisionDefinitionId gets a reference to the given NullableString and assigns it to the DecisionDefinitionId field.
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionId(v string) {
	o.DecisionDefinitionId.Set(&v)
}
// SetDecisionDefinitionIdNil sets the value for DecisionDefinitionId to be an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionIdNil() {
	o.DecisionDefinitionId.Set(nil)
}

// UnsetDecisionDefinitionId ensures that no value is present for DecisionDefinitionId, not even an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetDecisionDefinitionId() {
	o.DecisionDefinitionId.Unset()
}

// GetDecisionDefinitionKey returns the DecisionDefinitionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionKey() string {
	if o == nil || IsNil(o.DecisionDefinitionKey.Get()) {
		var ret string
		return ret
	}
	return *o.DecisionDefinitionKey.Get()
}

// GetDecisionDefinitionKeyOk returns a tuple with the DecisionDefinitionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecisionDefinitionKey.Get(), o.DecisionDefinitionKey.IsSet()
}

// HasDecisionDefinitionKey returns a boolean if a field has been set.
func (o *CleanableHistoricDecisionInstanceReportResultDto) HasDecisionDefinitionKey() bool {
	if o != nil && o.DecisionDefinitionKey.IsSet() {
		return true
	}

	return false
}

// SetDecisionDefinitionKey gets a reference to the given NullableString and assigns it to the DecisionDefinitionKey field.
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionKey(v string) {
	o.DecisionDefinitionKey.Set(&v)
}
// SetDecisionDefinitionKeyNil sets the value for DecisionDefinitionKey to be an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionKeyNil() {
	o.DecisionDefinitionKey.Set(nil)
}

// UnsetDecisionDefinitionKey ensures that no value is present for DecisionDefinitionKey, not even an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetDecisionDefinitionKey() {
	o.DecisionDefinitionKey.Unset()
}

// GetDecisionDefinitionName returns the DecisionDefinitionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionName() string {
	if o == nil || IsNil(o.DecisionDefinitionName.Get()) {
		var ret string
		return ret
	}
	return *o.DecisionDefinitionName.Get()
}

// GetDecisionDefinitionNameOk returns a tuple with the DecisionDefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecisionDefinitionName.Get(), o.DecisionDefinitionName.IsSet()
}

// HasDecisionDefinitionName returns a boolean if a field has been set.
func (o *CleanableHistoricDecisionInstanceReportResultDto) HasDecisionDefinitionName() bool {
	if o != nil && o.DecisionDefinitionName.IsSet() {
		return true
	}

	return false
}

// SetDecisionDefinitionName gets a reference to the given NullableString and assigns it to the DecisionDefinitionName field.
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionName(v string) {
	o.DecisionDefinitionName.Set(&v)
}
// SetDecisionDefinitionNameNil sets the value for DecisionDefinitionName to be an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionNameNil() {
	o.DecisionDefinitionName.Set(nil)
}

// UnsetDecisionDefinitionName ensures that no value is present for DecisionDefinitionName, not even an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetDecisionDefinitionName() {
	o.DecisionDefinitionName.Unset()
}

// GetDecisionDefinitionVersion returns the DecisionDefinitionVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionVersion() int32 {
	if o == nil || IsNil(o.DecisionDefinitionVersion.Get()) {
		var ret int32
		return ret
	}
	return *o.DecisionDefinitionVersion.Get()
}

// GetDecisionDefinitionVersionOk returns a tuple with the DecisionDefinitionVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetDecisionDefinitionVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecisionDefinitionVersion.Get(), o.DecisionDefinitionVersion.IsSet()
}

// HasDecisionDefinitionVersion returns a boolean if a field has been set.
func (o *CleanableHistoricDecisionInstanceReportResultDto) HasDecisionDefinitionVersion() bool {
	if o != nil && o.DecisionDefinitionVersion.IsSet() {
		return true
	}

	return false
}

// SetDecisionDefinitionVersion gets a reference to the given NullableInt32 and assigns it to the DecisionDefinitionVersion field.
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionVersion(v int32) {
	o.DecisionDefinitionVersion.Set(&v)
}
// SetDecisionDefinitionVersionNil sets the value for DecisionDefinitionVersion to be an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetDecisionDefinitionVersionNil() {
	o.DecisionDefinitionVersion.Set(nil)
}

// UnsetDecisionDefinitionVersion ensures that no value is present for DecisionDefinitionVersion, not even an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetDecisionDefinitionVersion() {
	o.DecisionDefinitionVersion.Unset()
}

// GetHistoryTimeToLive returns the HistoryTimeToLive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetHistoryTimeToLive() int32 {
	if o == nil || IsNil(o.HistoryTimeToLive.Get()) {
		var ret int32
		return ret
	}
	return *o.HistoryTimeToLive.Get()
}

// GetHistoryTimeToLiveOk returns a tuple with the HistoryTimeToLive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetHistoryTimeToLiveOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.HistoryTimeToLive.Get(), o.HistoryTimeToLive.IsSet()
}

// HasHistoryTimeToLive returns a boolean if a field has been set.
func (o *CleanableHistoricDecisionInstanceReportResultDto) HasHistoryTimeToLive() bool {
	if o != nil && o.HistoryTimeToLive.IsSet() {
		return true
	}

	return false
}

// SetHistoryTimeToLive gets a reference to the given NullableInt32 and assigns it to the HistoryTimeToLive field.
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetHistoryTimeToLive(v int32) {
	o.HistoryTimeToLive.Set(&v)
}
// SetHistoryTimeToLiveNil sets the value for HistoryTimeToLive to be an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetHistoryTimeToLiveNil() {
	o.HistoryTimeToLive.Set(nil)
}

// UnsetHistoryTimeToLive ensures that no value is present for HistoryTimeToLive, not even an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetHistoryTimeToLive() {
	o.HistoryTimeToLive.Unset()
}

// GetFinishedDecisionInstanceCount returns the FinishedDecisionInstanceCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetFinishedDecisionInstanceCount() int64 {
	if o == nil || IsNil(o.FinishedDecisionInstanceCount.Get()) {
		var ret int64
		return ret
	}
	return *o.FinishedDecisionInstanceCount.Get()
}

// GetFinishedDecisionInstanceCountOk returns a tuple with the FinishedDecisionInstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetFinishedDecisionInstanceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishedDecisionInstanceCount.Get(), o.FinishedDecisionInstanceCount.IsSet()
}

// HasFinishedDecisionInstanceCount returns a boolean if a field has been set.
func (o *CleanableHistoricDecisionInstanceReportResultDto) HasFinishedDecisionInstanceCount() bool {
	if o != nil && o.FinishedDecisionInstanceCount.IsSet() {
		return true
	}

	return false
}

// SetFinishedDecisionInstanceCount gets a reference to the given NullableInt64 and assigns it to the FinishedDecisionInstanceCount field.
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetFinishedDecisionInstanceCount(v int64) {
	o.FinishedDecisionInstanceCount.Set(&v)
}
// SetFinishedDecisionInstanceCountNil sets the value for FinishedDecisionInstanceCount to be an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetFinishedDecisionInstanceCountNil() {
	o.FinishedDecisionInstanceCount.Set(nil)
}

// UnsetFinishedDecisionInstanceCount ensures that no value is present for FinishedDecisionInstanceCount, not even an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetFinishedDecisionInstanceCount() {
	o.FinishedDecisionInstanceCount.Unset()
}

// GetCleanableDecisionInstanceCount returns the CleanableDecisionInstanceCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetCleanableDecisionInstanceCount() int64 {
	if o == nil || IsNil(o.CleanableDecisionInstanceCount.Get()) {
		var ret int64
		return ret
	}
	return *o.CleanableDecisionInstanceCount.Get()
}

// GetCleanableDecisionInstanceCountOk returns a tuple with the CleanableDecisionInstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetCleanableDecisionInstanceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CleanableDecisionInstanceCount.Get(), o.CleanableDecisionInstanceCount.IsSet()
}

// HasCleanableDecisionInstanceCount returns a boolean if a field has been set.
func (o *CleanableHistoricDecisionInstanceReportResultDto) HasCleanableDecisionInstanceCount() bool {
	if o != nil && o.CleanableDecisionInstanceCount.IsSet() {
		return true
	}

	return false
}

// SetCleanableDecisionInstanceCount gets a reference to the given NullableInt64 and assigns it to the CleanableDecisionInstanceCount field.
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetCleanableDecisionInstanceCount(v int64) {
	o.CleanableDecisionInstanceCount.Set(&v)
}
// SetCleanableDecisionInstanceCountNil sets the value for CleanableDecisionInstanceCount to be an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetCleanableDecisionInstanceCountNil() {
	o.CleanableDecisionInstanceCount.Set(nil)
}

// UnsetCleanableDecisionInstanceCount ensures that no value is present for CleanableDecisionInstanceCount, not even an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetCleanableDecisionInstanceCount() {
	o.CleanableDecisionInstanceCount.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CleanableHistoricDecisionInstanceReportResultDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *CleanableHistoricDecisionInstanceReportResultDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *CleanableHistoricDecisionInstanceReportResultDto) UnsetTenantId() {
	o.TenantId.Unset()
}

func (o CleanableHistoricDecisionInstanceReportResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CleanableHistoricDecisionInstanceReportResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DecisionDefinitionId.IsSet() {
		toSerialize["decisionDefinitionId"] = o.DecisionDefinitionId.Get()
	}
	if o.DecisionDefinitionKey.IsSet() {
		toSerialize["decisionDefinitionKey"] = o.DecisionDefinitionKey.Get()
	}
	if o.DecisionDefinitionName.IsSet() {
		toSerialize["decisionDefinitionName"] = o.DecisionDefinitionName.Get()
	}
	if o.DecisionDefinitionVersion.IsSet() {
		toSerialize["decisionDefinitionVersion"] = o.DecisionDefinitionVersion.Get()
	}
	if o.HistoryTimeToLive.IsSet() {
		toSerialize["historyTimeToLive"] = o.HistoryTimeToLive.Get()
	}
	if o.FinishedDecisionInstanceCount.IsSet() {
		toSerialize["finishedDecisionInstanceCount"] = o.FinishedDecisionInstanceCount.Get()
	}
	if o.CleanableDecisionInstanceCount.IsSet() {
		toSerialize["cleanableDecisionInstanceCount"] = o.CleanableDecisionInstanceCount.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	return toSerialize, nil
}

type NullableCleanableHistoricDecisionInstanceReportResultDto struct {
	value *CleanableHistoricDecisionInstanceReportResultDto
	isSet bool
}

func (v NullableCleanableHistoricDecisionInstanceReportResultDto) Get() *CleanableHistoricDecisionInstanceReportResultDto {
	return v.value
}

func (v *NullableCleanableHistoricDecisionInstanceReportResultDto) Set(val *CleanableHistoricDecisionInstanceReportResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCleanableHistoricDecisionInstanceReportResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCleanableHistoricDecisionInstanceReportResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCleanableHistoricDecisionInstanceReportResultDto(val *CleanableHistoricDecisionInstanceReportResultDto) *NullableCleanableHistoricDecisionInstanceReportResultDto {
	return &NullableCleanableHistoricDecisionInstanceReportResultDto{value: val, isSet: true}
}

func (v NullableCleanableHistoricDecisionInstanceReportResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCleanableHistoricDecisionInstanceReportResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

