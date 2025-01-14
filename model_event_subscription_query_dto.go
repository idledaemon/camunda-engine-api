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

// checks if the EventSubscriptionQueryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventSubscriptionQueryDto{}

// EventSubscriptionQueryDto A event subscription query which retrieves a list of event subscriptions
type EventSubscriptionQueryDto struct {
	// The id of the event subscription.
	EventSubscriptionId NullableString `json:"eventSubscriptionId,omitempty"`
	// The name of the event this subscription belongs to as defined in the process model.
	EventName NullableString `json:"eventName,omitempty"`
	// The type of the event subscription.
	EventType NullableString `json:"eventType,omitempty"`
	// The execution that is subscribed on the referenced event.
	ExecutionId NullableString `json:"executionId,omitempty"`
	// The process instance this subscription belongs to.
	ProcessInstanceId NullableString `json:"processInstanceId,omitempty"`
	// The identifier of the activity that this event subscription belongs to. This could for example be the id of a receive task.
	ActivityId NullableString `json:"activityId,omitempty"`
	// Filter by a comma-separated list of tenant ids. Only select subscriptions that belong to one of the given tenant ids.
	TenantIdIn []string `json:"tenantIdIn,omitempty"`
	// Only select subscriptions which have no tenant id. Value may only be `true`, as `false` is the default behavior.
	WithoutTenantId NullableBool `json:"withoutTenantId,omitempty"`
	// Select event subscriptions which have no tenant id. Can be used in combination with tenantIdIn parameter. Value may only be `true`, as `false` is the default behavior.
	IncludeEventSubscriptionsWithoutTenantId NullableBool `json:"includeEventSubscriptionsWithoutTenantId,omitempty"`
	// Apply sorting of the result
	Sorting []EventSubscriptionQueryDtoSortingInner `json:"sorting,omitempty"`
}

// NewEventSubscriptionQueryDto instantiates a new EventSubscriptionQueryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSubscriptionQueryDto() *EventSubscriptionQueryDto {
	this := EventSubscriptionQueryDto{}
	return &this
}

// NewEventSubscriptionQueryDtoWithDefaults instantiates a new EventSubscriptionQueryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSubscriptionQueryDtoWithDefaults() *EventSubscriptionQueryDto {
	this := EventSubscriptionQueryDto{}
	return &this
}

// GetEventSubscriptionId returns the EventSubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSubscriptionQueryDto) GetEventSubscriptionId() string {
	if o == nil || IsNil(o.EventSubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.EventSubscriptionId.Get()
}

// GetEventSubscriptionIdOk returns a tuple with the EventSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSubscriptionQueryDto) GetEventSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventSubscriptionId.Get(), o.EventSubscriptionId.IsSet()
}

// HasEventSubscriptionId returns a boolean if a field has been set.
func (o *EventSubscriptionQueryDto) HasEventSubscriptionId() bool {
	if o != nil && o.EventSubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetEventSubscriptionId gets a reference to the given NullableString and assigns it to the EventSubscriptionId field.
func (o *EventSubscriptionQueryDto) SetEventSubscriptionId(v string) {
	o.EventSubscriptionId.Set(&v)
}
// SetEventSubscriptionIdNil sets the value for EventSubscriptionId to be an explicit nil
func (o *EventSubscriptionQueryDto) SetEventSubscriptionIdNil() {
	o.EventSubscriptionId.Set(nil)
}

// UnsetEventSubscriptionId ensures that no value is present for EventSubscriptionId, not even an explicit nil
func (o *EventSubscriptionQueryDto) UnsetEventSubscriptionId() {
	o.EventSubscriptionId.Unset()
}

// GetEventName returns the EventName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSubscriptionQueryDto) GetEventName() string {
	if o == nil || IsNil(o.EventName.Get()) {
		var ret string
		return ret
	}
	return *o.EventName.Get()
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSubscriptionQueryDto) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventName.Get(), o.EventName.IsSet()
}

// HasEventName returns a boolean if a field has been set.
func (o *EventSubscriptionQueryDto) HasEventName() bool {
	if o != nil && o.EventName.IsSet() {
		return true
	}

	return false
}

// SetEventName gets a reference to the given NullableString and assigns it to the EventName field.
func (o *EventSubscriptionQueryDto) SetEventName(v string) {
	o.EventName.Set(&v)
}
// SetEventNameNil sets the value for EventName to be an explicit nil
func (o *EventSubscriptionQueryDto) SetEventNameNil() {
	o.EventName.Set(nil)
}

// UnsetEventName ensures that no value is present for EventName, not even an explicit nil
func (o *EventSubscriptionQueryDto) UnsetEventName() {
	o.EventName.Unset()
}

// GetEventType returns the EventType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSubscriptionQueryDto) GetEventType() string {
	if o == nil || IsNil(o.EventType.Get()) {
		var ret string
		return ret
	}
	return *o.EventType.Get()
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSubscriptionQueryDto) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventType.Get(), o.EventType.IsSet()
}

// HasEventType returns a boolean if a field has been set.
func (o *EventSubscriptionQueryDto) HasEventType() bool {
	if o != nil && o.EventType.IsSet() {
		return true
	}

	return false
}

// SetEventType gets a reference to the given NullableString and assigns it to the EventType field.
func (o *EventSubscriptionQueryDto) SetEventType(v string) {
	o.EventType.Set(&v)
}
// SetEventTypeNil sets the value for EventType to be an explicit nil
func (o *EventSubscriptionQueryDto) SetEventTypeNil() {
	o.EventType.Set(nil)
}

// UnsetEventType ensures that no value is present for EventType, not even an explicit nil
func (o *EventSubscriptionQueryDto) UnsetEventType() {
	o.EventType.Unset()
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSubscriptionQueryDto) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionId.Get()
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSubscriptionQueryDto) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionId.Get(), o.ExecutionId.IsSet()
}

// HasExecutionId returns a boolean if a field has been set.
func (o *EventSubscriptionQueryDto) HasExecutionId() bool {
	if o != nil && o.ExecutionId.IsSet() {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given NullableString and assigns it to the ExecutionId field.
func (o *EventSubscriptionQueryDto) SetExecutionId(v string) {
	o.ExecutionId.Set(&v)
}
// SetExecutionIdNil sets the value for ExecutionId to be an explicit nil
func (o *EventSubscriptionQueryDto) SetExecutionIdNil() {
	o.ExecutionId.Set(nil)
}

// UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
func (o *EventSubscriptionQueryDto) UnsetExecutionId() {
	o.ExecutionId.Unset()
}

// GetProcessInstanceId returns the ProcessInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSubscriptionQueryDto) GetProcessInstanceId() string {
	if o == nil || IsNil(o.ProcessInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessInstanceId.Get()
}

// GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSubscriptionQueryDto) GetProcessInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessInstanceId.Get(), o.ProcessInstanceId.IsSet()
}

// HasProcessInstanceId returns a boolean if a field has been set.
func (o *EventSubscriptionQueryDto) HasProcessInstanceId() bool {
	if o != nil && o.ProcessInstanceId.IsSet() {
		return true
	}

	return false
}

// SetProcessInstanceId gets a reference to the given NullableString and assigns it to the ProcessInstanceId field.
func (o *EventSubscriptionQueryDto) SetProcessInstanceId(v string) {
	o.ProcessInstanceId.Set(&v)
}
// SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil
func (o *EventSubscriptionQueryDto) SetProcessInstanceIdNil() {
	o.ProcessInstanceId.Set(nil)
}

// UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
func (o *EventSubscriptionQueryDto) UnsetProcessInstanceId() {
	o.ProcessInstanceId.Unset()
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSubscriptionQueryDto) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId.Get()) {
		var ret string
		return ret
	}
	return *o.ActivityId.Get()
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSubscriptionQueryDto) GetActivityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityId.Get(), o.ActivityId.IsSet()
}

// HasActivityId returns a boolean if a field has been set.
func (o *EventSubscriptionQueryDto) HasActivityId() bool {
	if o != nil && o.ActivityId.IsSet() {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given NullableString and assigns it to the ActivityId field.
func (o *EventSubscriptionQueryDto) SetActivityId(v string) {
	o.ActivityId.Set(&v)
}
// SetActivityIdNil sets the value for ActivityId to be an explicit nil
func (o *EventSubscriptionQueryDto) SetActivityIdNil() {
	o.ActivityId.Set(nil)
}

// UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
func (o *EventSubscriptionQueryDto) UnsetActivityId() {
	o.ActivityId.Unset()
}

// GetTenantIdIn returns the TenantIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSubscriptionQueryDto) GetTenantIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TenantIdIn
}

// GetTenantIdInOk returns a tuple with the TenantIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSubscriptionQueryDto) GetTenantIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.TenantIdIn) {
		return nil, false
	}
	return o.TenantIdIn, true
}

// HasTenantIdIn returns a boolean if a field has been set.
func (o *EventSubscriptionQueryDto) HasTenantIdIn() bool {
	if o != nil && !IsNil(o.TenantIdIn) {
		return true
	}

	return false
}

// SetTenantIdIn gets a reference to the given []string and assigns it to the TenantIdIn field.
func (o *EventSubscriptionQueryDto) SetTenantIdIn(v []string) {
	o.TenantIdIn = v
}

// GetWithoutTenantId returns the WithoutTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSubscriptionQueryDto) GetWithoutTenantId() bool {
	if o == nil || IsNil(o.WithoutTenantId.Get()) {
		var ret bool
		return ret
	}
	return *o.WithoutTenantId.Get()
}

// GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSubscriptionQueryDto) GetWithoutTenantIdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithoutTenantId.Get(), o.WithoutTenantId.IsSet()
}

// HasWithoutTenantId returns a boolean if a field has been set.
func (o *EventSubscriptionQueryDto) HasWithoutTenantId() bool {
	if o != nil && o.WithoutTenantId.IsSet() {
		return true
	}

	return false
}

// SetWithoutTenantId gets a reference to the given NullableBool and assigns it to the WithoutTenantId field.
func (o *EventSubscriptionQueryDto) SetWithoutTenantId(v bool) {
	o.WithoutTenantId.Set(&v)
}
// SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil
func (o *EventSubscriptionQueryDto) SetWithoutTenantIdNil() {
	o.WithoutTenantId.Set(nil)
}

// UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
func (o *EventSubscriptionQueryDto) UnsetWithoutTenantId() {
	o.WithoutTenantId.Unset()
}

// GetIncludeEventSubscriptionsWithoutTenantId returns the IncludeEventSubscriptionsWithoutTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSubscriptionQueryDto) GetIncludeEventSubscriptionsWithoutTenantId() bool {
	if o == nil || IsNil(o.IncludeEventSubscriptionsWithoutTenantId.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeEventSubscriptionsWithoutTenantId.Get()
}

// GetIncludeEventSubscriptionsWithoutTenantIdOk returns a tuple with the IncludeEventSubscriptionsWithoutTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSubscriptionQueryDto) GetIncludeEventSubscriptionsWithoutTenantIdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeEventSubscriptionsWithoutTenantId.Get(), o.IncludeEventSubscriptionsWithoutTenantId.IsSet()
}

// HasIncludeEventSubscriptionsWithoutTenantId returns a boolean if a field has been set.
func (o *EventSubscriptionQueryDto) HasIncludeEventSubscriptionsWithoutTenantId() bool {
	if o != nil && o.IncludeEventSubscriptionsWithoutTenantId.IsSet() {
		return true
	}

	return false
}

// SetIncludeEventSubscriptionsWithoutTenantId gets a reference to the given NullableBool and assigns it to the IncludeEventSubscriptionsWithoutTenantId field.
func (o *EventSubscriptionQueryDto) SetIncludeEventSubscriptionsWithoutTenantId(v bool) {
	o.IncludeEventSubscriptionsWithoutTenantId.Set(&v)
}
// SetIncludeEventSubscriptionsWithoutTenantIdNil sets the value for IncludeEventSubscriptionsWithoutTenantId to be an explicit nil
func (o *EventSubscriptionQueryDto) SetIncludeEventSubscriptionsWithoutTenantIdNil() {
	o.IncludeEventSubscriptionsWithoutTenantId.Set(nil)
}

// UnsetIncludeEventSubscriptionsWithoutTenantId ensures that no value is present for IncludeEventSubscriptionsWithoutTenantId, not even an explicit nil
func (o *EventSubscriptionQueryDto) UnsetIncludeEventSubscriptionsWithoutTenantId() {
	o.IncludeEventSubscriptionsWithoutTenantId.Unset()
}

// GetSorting returns the Sorting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSubscriptionQueryDto) GetSorting() []EventSubscriptionQueryDtoSortingInner {
	if o == nil {
		var ret []EventSubscriptionQueryDtoSortingInner
		return ret
	}
	return o.Sorting
}

// GetSortingOk returns a tuple with the Sorting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSubscriptionQueryDto) GetSortingOk() ([]EventSubscriptionQueryDtoSortingInner, bool) {
	if o == nil || IsNil(o.Sorting) {
		return nil, false
	}
	return o.Sorting, true
}

// HasSorting returns a boolean if a field has been set.
func (o *EventSubscriptionQueryDto) HasSorting() bool {
	if o != nil && !IsNil(o.Sorting) {
		return true
	}

	return false
}

// SetSorting gets a reference to the given []EventSubscriptionQueryDtoSortingInner and assigns it to the Sorting field.
func (o *EventSubscriptionQueryDto) SetSorting(v []EventSubscriptionQueryDtoSortingInner) {
	o.Sorting = v
}

func (o EventSubscriptionQueryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSubscriptionQueryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EventSubscriptionId.IsSet() {
		toSerialize["eventSubscriptionId"] = o.EventSubscriptionId.Get()
	}
	if o.EventName.IsSet() {
		toSerialize["eventName"] = o.EventName.Get()
	}
	if o.EventType.IsSet() {
		toSerialize["eventType"] = o.EventType.Get()
	}
	if o.ExecutionId.IsSet() {
		toSerialize["executionId"] = o.ExecutionId.Get()
	}
	if o.ProcessInstanceId.IsSet() {
		toSerialize["processInstanceId"] = o.ProcessInstanceId.Get()
	}
	if o.ActivityId.IsSet() {
		toSerialize["activityId"] = o.ActivityId.Get()
	}
	if o.TenantIdIn != nil {
		toSerialize["tenantIdIn"] = o.TenantIdIn
	}
	if o.WithoutTenantId.IsSet() {
		toSerialize["withoutTenantId"] = o.WithoutTenantId.Get()
	}
	if o.IncludeEventSubscriptionsWithoutTenantId.IsSet() {
		toSerialize["includeEventSubscriptionsWithoutTenantId"] = o.IncludeEventSubscriptionsWithoutTenantId.Get()
	}
	if o.Sorting != nil {
		toSerialize["sorting"] = o.Sorting
	}
	return toSerialize, nil
}

type NullableEventSubscriptionQueryDto struct {
	value *EventSubscriptionQueryDto
	isSet bool
}

func (v NullableEventSubscriptionQueryDto) Get() *EventSubscriptionQueryDto {
	return v.value
}

func (v *NullableEventSubscriptionQueryDto) Set(val *EventSubscriptionQueryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSubscriptionQueryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSubscriptionQueryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSubscriptionQueryDto(val *EventSubscriptionQueryDto) *NullableEventSubscriptionQueryDto {
	return &NullableEventSubscriptionQueryDto{value: val, isSet: true}
}

func (v NullableEventSubscriptionQueryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSubscriptionQueryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


