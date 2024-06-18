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

// checks if the TransitionInstanceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransitionInstanceDto{}

// TransitionInstanceDto A JSON object corresponding to the Activity Instance tree of the given process instance.
type TransitionInstanceDto struct {
	// The id of the transition instance.
	Id NullableString `json:"id,omitempty"`
	// The id of the parent activity instance, for example a sub process instance.
	ParentActivityInstanceId NullableString `json:"parentActivityInstanceId,omitempty"`
	// The id of the activity that this instance enters (asyncBefore job) or leaves (asyncAfter job)
	ActivityId NullableString `json:"activityId,omitempty"`
	// The name of the activity that this instance enters (asyncBefore job) or leaves (asyncAfter job)
	ActivityName NullableString `json:"activityName,omitempty"`
	// The type of the activity that this instance enters (asyncBefore job) or leaves (asyncAfter job)
	ActivityType NullableString `json:"activityType,omitempty"`
	// The id of the process instance this instance is part of.
	ProcessInstanceId NullableString `json:"processInstanceId,omitempty"`
	// The id of the process definition.
	ProcessDefinitionId NullableString `json:"processDefinitionId,omitempty"`
	// The execution id.
	ExecutionId NullableString `json:"executionId,omitempty"`
	// A list of incident ids.
	IncidentIds []string `json:"incidentIds,omitempty"`
	// A list of JSON objects containing incident specific properties: * `id`: the id of the incident * `activityId`: the activity id in which the incident occurred
	Incidents []ActivityInstanceIncidentDto `json:"incidents,omitempty"`
}

// NewTransitionInstanceDto instantiates a new TransitionInstanceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransitionInstanceDto() *TransitionInstanceDto {
	this := TransitionInstanceDto{}
	return &this
}

// NewTransitionInstanceDtoWithDefaults instantiates a new TransitionInstanceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitionInstanceDtoWithDefaults() *TransitionInstanceDto {
	this := TransitionInstanceDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransitionInstanceDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransitionInstanceDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TransitionInstanceDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *TransitionInstanceDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TransitionInstanceDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TransitionInstanceDto) UnsetId() {
	o.Id.Unset()
}

// GetParentActivityInstanceId returns the ParentActivityInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransitionInstanceDto) GetParentActivityInstanceId() string {
	if o == nil || IsNil(o.ParentActivityInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentActivityInstanceId.Get()
}

// GetParentActivityInstanceIdOk returns a tuple with the ParentActivityInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransitionInstanceDto) GetParentActivityInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentActivityInstanceId.Get(), o.ParentActivityInstanceId.IsSet()
}

// HasParentActivityInstanceId returns a boolean if a field has been set.
func (o *TransitionInstanceDto) HasParentActivityInstanceId() bool {
	if o != nil && o.ParentActivityInstanceId.IsSet() {
		return true
	}

	return false
}

// SetParentActivityInstanceId gets a reference to the given NullableString and assigns it to the ParentActivityInstanceId field.
func (o *TransitionInstanceDto) SetParentActivityInstanceId(v string) {
	o.ParentActivityInstanceId.Set(&v)
}
// SetParentActivityInstanceIdNil sets the value for ParentActivityInstanceId to be an explicit nil
func (o *TransitionInstanceDto) SetParentActivityInstanceIdNil() {
	o.ParentActivityInstanceId.Set(nil)
}

// UnsetParentActivityInstanceId ensures that no value is present for ParentActivityInstanceId, not even an explicit nil
func (o *TransitionInstanceDto) UnsetParentActivityInstanceId() {
	o.ParentActivityInstanceId.Unset()
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransitionInstanceDto) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId.Get()) {
		var ret string
		return ret
	}
	return *o.ActivityId.Get()
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransitionInstanceDto) GetActivityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityId.Get(), o.ActivityId.IsSet()
}

// HasActivityId returns a boolean if a field has been set.
func (o *TransitionInstanceDto) HasActivityId() bool {
	if o != nil && o.ActivityId.IsSet() {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given NullableString and assigns it to the ActivityId field.
func (o *TransitionInstanceDto) SetActivityId(v string) {
	o.ActivityId.Set(&v)
}
// SetActivityIdNil sets the value for ActivityId to be an explicit nil
func (o *TransitionInstanceDto) SetActivityIdNil() {
	o.ActivityId.Set(nil)
}

// UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
func (o *TransitionInstanceDto) UnsetActivityId() {
	o.ActivityId.Unset()
}

// GetActivityName returns the ActivityName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransitionInstanceDto) GetActivityName() string {
	if o == nil || IsNil(o.ActivityName.Get()) {
		var ret string
		return ret
	}
	return *o.ActivityName.Get()
}

// GetActivityNameOk returns a tuple with the ActivityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransitionInstanceDto) GetActivityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityName.Get(), o.ActivityName.IsSet()
}

// HasActivityName returns a boolean if a field has been set.
func (o *TransitionInstanceDto) HasActivityName() bool {
	if o != nil && o.ActivityName.IsSet() {
		return true
	}

	return false
}

// SetActivityName gets a reference to the given NullableString and assigns it to the ActivityName field.
func (o *TransitionInstanceDto) SetActivityName(v string) {
	o.ActivityName.Set(&v)
}
// SetActivityNameNil sets the value for ActivityName to be an explicit nil
func (o *TransitionInstanceDto) SetActivityNameNil() {
	o.ActivityName.Set(nil)
}

// UnsetActivityName ensures that no value is present for ActivityName, not even an explicit nil
func (o *TransitionInstanceDto) UnsetActivityName() {
	o.ActivityName.Unset()
}

// GetActivityType returns the ActivityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransitionInstanceDto) GetActivityType() string {
	if o == nil || IsNil(o.ActivityType.Get()) {
		var ret string
		return ret
	}
	return *o.ActivityType.Get()
}

// GetActivityTypeOk returns a tuple with the ActivityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransitionInstanceDto) GetActivityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityType.Get(), o.ActivityType.IsSet()
}

// HasActivityType returns a boolean if a field has been set.
func (o *TransitionInstanceDto) HasActivityType() bool {
	if o != nil && o.ActivityType.IsSet() {
		return true
	}

	return false
}

// SetActivityType gets a reference to the given NullableString and assigns it to the ActivityType field.
func (o *TransitionInstanceDto) SetActivityType(v string) {
	o.ActivityType.Set(&v)
}
// SetActivityTypeNil sets the value for ActivityType to be an explicit nil
func (o *TransitionInstanceDto) SetActivityTypeNil() {
	o.ActivityType.Set(nil)
}

// UnsetActivityType ensures that no value is present for ActivityType, not even an explicit nil
func (o *TransitionInstanceDto) UnsetActivityType() {
	o.ActivityType.Unset()
}

// GetProcessInstanceId returns the ProcessInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransitionInstanceDto) GetProcessInstanceId() string {
	if o == nil || IsNil(o.ProcessInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessInstanceId.Get()
}

// GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransitionInstanceDto) GetProcessInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessInstanceId.Get(), o.ProcessInstanceId.IsSet()
}

// HasProcessInstanceId returns a boolean if a field has been set.
func (o *TransitionInstanceDto) HasProcessInstanceId() bool {
	if o != nil && o.ProcessInstanceId.IsSet() {
		return true
	}

	return false
}

// SetProcessInstanceId gets a reference to the given NullableString and assigns it to the ProcessInstanceId field.
func (o *TransitionInstanceDto) SetProcessInstanceId(v string) {
	o.ProcessInstanceId.Set(&v)
}
// SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil
func (o *TransitionInstanceDto) SetProcessInstanceIdNil() {
	o.ProcessInstanceId.Set(nil)
}

// UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
func (o *TransitionInstanceDto) UnsetProcessInstanceId() {
	o.ProcessInstanceId.Unset()
}

// GetProcessDefinitionId returns the ProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransitionInstanceDto) GetProcessDefinitionId() string {
	if o == nil || IsNil(o.ProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionId.Get()
}

// GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransitionInstanceDto) GetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionId.Get(), o.ProcessDefinitionId.IsSet()
}

// HasProcessDefinitionId returns a boolean if a field has been set.
func (o *TransitionInstanceDto) HasProcessDefinitionId() bool {
	if o != nil && o.ProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionId gets a reference to the given NullableString and assigns it to the ProcessDefinitionId field.
func (o *TransitionInstanceDto) SetProcessDefinitionId(v string) {
	o.ProcessDefinitionId.Set(&v)
}
// SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil
func (o *TransitionInstanceDto) SetProcessDefinitionIdNil() {
	o.ProcessDefinitionId.Set(nil)
}

// UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
func (o *TransitionInstanceDto) UnsetProcessDefinitionId() {
	o.ProcessDefinitionId.Unset()
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransitionInstanceDto) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionId.Get()
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransitionInstanceDto) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionId.Get(), o.ExecutionId.IsSet()
}

// HasExecutionId returns a boolean if a field has been set.
func (o *TransitionInstanceDto) HasExecutionId() bool {
	if o != nil && o.ExecutionId.IsSet() {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given NullableString and assigns it to the ExecutionId field.
func (o *TransitionInstanceDto) SetExecutionId(v string) {
	o.ExecutionId.Set(&v)
}
// SetExecutionIdNil sets the value for ExecutionId to be an explicit nil
func (o *TransitionInstanceDto) SetExecutionIdNil() {
	o.ExecutionId.Set(nil)
}

// UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
func (o *TransitionInstanceDto) UnsetExecutionId() {
	o.ExecutionId.Unset()
}

// GetIncidentIds returns the IncidentIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransitionInstanceDto) GetIncidentIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IncidentIds
}

// GetIncidentIdsOk returns a tuple with the IncidentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransitionInstanceDto) GetIncidentIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncidentIds) {
		return nil, false
	}
	return o.IncidentIds, true
}

// HasIncidentIds returns a boolean if a field has been set.
func (o *TransitionInstanceDto) HasIncidentIds() bool {
	if o != nil && !IsNil(o.IncidentIds) {
		return true
	}

	return false
}

// SetIncidentIds gets a reference to the given []string and assigns it to the IncidentIds field.
func (o *TransitionInstanceDto) SetIncidentIds(v []string) {
	o.IncidentIds = v
}

// GetIncidents returns the Incidents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransitionInstanceDto) GetIncidents() []ActivityInstanceIncidentDto {
	if o == nil {
		var ret []ActivityInstanceIncidentDto
		return ret
	}
	return o.Incidents
}

// GetIncidentsOk returns a tuple with the Incidents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransitionInstanceDto) GetIncidentsOk() ([]ActivityInstanceIncidentDto, bool) {
	if o == nil || IsNil(o.Incidents) {
		return nil, false
	}
	return o.Incidents, true
}

// HasIncidents returns a boolean if a field has been set.
func (o *TransitionInstanceDto) HasIncidents() bool {
	if o != nil && !IsNil(o.Incidents) {
		return true
	}

	return false
}

// SetIncidents gets a reference to the given []ActivityInstanceIncidentDto and assigns it to the Incidents field.
func (o *TransitionInstanceDto) SetIncidents(v []ActivityInstanceIncidentDto) {
	o.Incidents = v
}

func (o TransitionInstanceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransitionInstanceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ParentActivityInstanceId.IsSet() {
		toSerialize["parentActivityInstanceId"] = o.ParentActivityInstanceId.Get()
	}
	if o.ActivityId.IsSet() {
		toSerialize["activityId"] = o.ActivityId.Get()
	}
	if o.ActivityName.IsSet() {
		toSerialize["activityName"] = o.ActivityName.Get()
	}
	if o.ActivityType.IsSet() {
		toSerialize["activityType"] = o.ActivityType.Get()
	}
	if o.ProcessInstanceId.IsSet() {
		toSerialize["processInstanceId"] = o.ProcessInstanceId.Get()
	}
	if o.ProcessDefinitionId.IsSet() {
		toSerialize["processDefinitionId"] = o.ProcessDefinitionId.Get()
	}
	if o.ExecutionId.IsSet() {
		toSerialize["executionId"] = o.ExecutionId.Get()
	}
	if o.IncidentIds != nil {
		toSerialize["incidentIds"] = o.IncidentIds
	}
	if o.Incidents != nil {
		toSerialize["incidents"] = o.Incidents
	}
	return toSerialize, nil
}

type NullableTransitionInstanceDto struct {
	value *TransitionInstanceDto
	isSet bool
}

func (v NullableTransitionInstanceDto) Get() *TransitionInstanceDto {
	return v.value
}

func (v *NullableTransitionInstanceDto) Set(val *TransitionInstanceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTransitionInstanceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTransitionInstanceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransitionInstanceDto(val *TransitionInstanceDto) *NullableTransitionInstanceDto {
	return &NullableTransitionInstanceDto{value: val, isSet: true}
}

func (v NullableTransitionInstanceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransitionInstanceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


