/*
Camunda Platform REST API

OpenApi Spec for Camunda Platform REST API.

API version: 7.21.2-ee
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package camundarestgo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RestartProcessInstanceModificationInstructionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestartProcessInstanceModificationInstructionDto{}

// RestartProcessInstanceModificationInstructionDto struct for RestartProcessInstanceModificationInstructionDto
type RestartProcessInstanceModificationInstructionDto struct {
	// **Mandatory**. One of the following values: `startBeforeActivity`, `startAfterActivity`, `startTransition`.  * A `startBeforeActivity` instruction requests to enter a given activity. * A `startAfterActivity` instruction requests to execute the single outgoing sequence flow of a given activity. * A `startTransition` instruction requests to execute a specific sequence flow.
	Type string `json:"type"`
	// **Can be used with instructions of types** `startBeforeActivity` and `startAfterActivity`. Specifies the sequence flow to start.
	ActivityId NullableString `json:"activityId,omitempty"`
	// **Can be used with instructions of types** `startTransition`. Specifies the sequence flow to start.
	TransitionId NullableString `json:"transitionId,omitempty"`
}

type _RestartProcessInstanceModificationInstructionDto RestartProcessInstanceModificationInstructionDto

// NewRestartProcessInstanceModificationInstructionDto instantiates a new RestartProcessInstanceModificationInstructionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestartProcessInstanceModificationInstructionDto(type_ string) *RestartProcessInstanceModificationInstructionDto {
	this := RestartProcessInstanceModificationInstructionDto{}
	this.Type = type_
	return &this
}

// NewRestartProcessInstanceModificationInstructionDtoWithDefaults instantiates a new RestartProcessInstanceModificationInstructionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestartProcessInstanceModificationInstructionDtoWithDefaults() *RestartProcessInstanceModificationInstructionDto {
	this := RestartProcessInstanceModificationInstructionDto{}
	return &this
}

// GetType returns the Type field value
func (o *RestartProcessInstanceModificationInstructionDto) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RestartProcessInstanceModificationInstructionDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RestartProcessInstanceModificationInstructionDto) SetType(v string) {
	o.Type = v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestartProcessInstanceModificationInstructionDto) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId.Get()) {
		var ret string
		return ret
	}
	return *o.ActivityId.Get()
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestartProcessInstanceModificationInstructionDto) GetActivityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityId.Get(), o.ActivityId.IsSet()
}

// HasActivityId returns a boolean if a field has been set.
func (o *RestartProcessInstanceModificationInstructionDto) HasActivityId() bool {
	if o != nil && o.ActivityId.IsSet() {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given NullableString and assigns it to the ActivityId field.
func (o *RestartProcessInstanceModificationInstructionDto) SetActivityId(v string) {
	o.ActivityId.Set(&v)
}
// SetActivityIdNil sets the value for ActivityId to be an explicit nil
func (o *RestartProcessInstanceModificationInstructionDto) SetActivityIdNil() {
	o.ActivityId.Set(nil)
}

// UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
func (o *RestartProcessInstanceModificationInstructionDto) UnsetActivityId() {
	o.ActivityId.Unset()
}

// GetTransitionId returns the TransitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestartProcessInstanceModificationInstructionDto) GetTransitionId() string {
	if o == nil || IsNil(o.TransitionId.Get()) {
		var ret string
		return ret
	}
	return *o.TransitionId.Get()
}

// GetTransitionIdOk returns a tuple with the TransitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestartProcessInstanceModificationInstructionDto) GetTransitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransitionId.Get(), o.TransitionId.IsSet()
}

// HasTransitionId returns a boolean if a field has been set.
func (o *RestartProcessInstanceModificationInstructionDto) HasTransitionId() bool {
	if o != nil && o.TransitionId.IsSet() {
		return true
	}

	return false
}

// SetTransitionId gets a reference to the given NullableString and assigns it to the TransitionId field.
func (o *RestartProcessInstanceModificationInstructionDto) SetTransitionId(v string) {
	o.TransitionId.Set(&v)
}
// SetTransitionIdNil sets the value for TransitionId to be an explicit nil
func (o *RestartProcessInstanceModificationInstructionDto) SetTransitionIdNil() {
	o.TransitionId.Set(nil)
}

// UnsetTransitionId ensures that no value is present for TransitionId, not even an explicit nil
func (o *RestartProcessInstanceModificationInstructionDto) UnsetTransitionId() {
	o.TransitionId.Unset()
}

func (o RestartProcessInstanceModificationInstructionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestartProcessInstanceModificationInstructionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.ActivityId.IsSet() {
		toSerialize["activityId"] = o.ActivityId.Get()
	}
	if o.TransitionId.IsSet() {
		toSerialize["transitionId"] = o.TransitionId.Get()
	}
	return toSerialize, nil
}

func (o *RestartProcessInstanceModificationInstructionDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRestartProcessInstanceModificationInstructionDto := _RestartProcessInstanceModificationInstructionDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRestartProcessInstanceModificationInstructionDto)

	if err != nil {
		return err
	}

	*o = RestartProcessInstanceModificationInstructionDto(varRestartProcessInstanceModificationInstructionDto)

	return err
}

type NullableRestartProcessInstanceModificationInstructionDto struct {
	value *RestartProcessInstanceModificationInstructionDto
	isSet bool
}

func (v NullableRestartProcessInstanceModificationInstructionDto) Get() *RestartProcessInstanceModificationInstructionDto {
	return v.value
}

func (v *NullableRestartProcessInstanceModificationInstructionDto) Set(val *RestartProcessInstanceModificationInstructionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRestartProcessInstanceModificationInstructionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRestartProcessInstanceModificationInstructionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestartProcessInstanceModificationInstructionDto(val *RestartProcessInstanceModificationInstructionDto) *NullableRestartProcessInstanceModificationInstructionDto {
	return &NullableRestartProcessInstanceModificationInstructionDto{value: val, isSet: true}
}

func (v NullableRestartProcessInstanceModificationInstructionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestartProcessInstanceModificationInstructionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

