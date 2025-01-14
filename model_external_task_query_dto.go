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

// checks if the ExternalTaskQueryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalTaskQueryDto{}

// ExternalTaskQueryDto A JSON object with the following properties:
type ExternalTaskQueryDto struct {
	// Filter by an external task's id.
	ExternalTaskId NullableString `json:"externalTaskId,omitempty"`
	// Filter by the comma-separated list of external task ids.
	ExternalTaskIdIn []string `json:"externalTaskIdIn,omitempty"`
	// Filter by an external task topic.
	TopicName NullableString `json:"topicName,omitempty"`
	// Filter by the id of the worker that the task was most recently locked by.
	WorkerId NullableString `json:"workerId,omitempty"`
	// Only include external tasks that are currently locked (i.e., they have a lock time and it has not expired). Value may only be `true`, as `false` matches any external task.
	Locked NullableBool `json:"locked,omitempty"`
	// Only include external tasks that are currently not locked (i.e., they have no lock or it has expired). Value may only be `true`, as `false` matches any external task.
	NotLocked NullableBool `json:"notLocked,omitempty"`
	// Only include external tasks that have a positive (&gt; 0) number of retries (or `null`). Value may only be `true`, as `false` matches any external task.
	WithRetriesLeft NullableBool `json:"withRetriesLeft,omitempty"`
	// Only include external tasks that have 0 retries. Value may only be `true`, as `false` matches any external task.
	NoRetriesLeft NullableBool `json:"noRetriesLeft,omitempty"`
	// Restrict to external tasks that have a lock that expires after a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`.
	LockExpirationAfter NullableTime `json:"lockExpirationAfter,omitempty"`
	// Restrict to external tasks that have a lock that expires before a given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`.
	LockExpirationBefore NullableTime `json:"lockExpirationBefore,omitempty"`
	// Filter by the id of the activity that an external task is created for.
	ActivityId NullableString `json:"activityId,omitempty"`
	// Filter by the comma-separated list of ids of the activities that an external task is created for.
	ActivityIdIn []string `json:"activityIdIn,omitempty"`
	// Filter by the id of the execution that an external task belongs to.
	ExecutionId NullableString `json:"executionId,omitempty"`
	// Filter by the id of the process instance that an external task belongs to.
	ProcessInstanceId NullableString `json:"processInstanceId,omitempty"`
	// Filter by a comma-separated list of process instance ids that an external task may belong to.
	ProcessInstanceIdIn []string `json:"processInstanceIdIn,omitempty"`
	// Filter by the id of the process definition that an external task belongs to.
	ProcessDefinitionId NullableString `json:"processDefinitionId,omitempty"`
	// Filter by a comma-separated list of tenant ids. An external task must have one of the given tenant ids.
	TenantIdIn []string `json:"tenantIdIn,omitempty"`
	// Only include active tasks. Value may only be `true`, as `false` matches any external task.
	Active NullableBool `json:"active,omitempty"`
	// Only include suspended tasks. Value may only be `true`, as `false` matches any external task.
	Suspended NullableBool `json:"suspended,omitempty"`
	// Only include jobs with a priority higher than or equal to the given value. Value must be a valid `long` value.
	PriorityHigherThanOrEquals NullableInt64 `json:"priorityHigherThanOrEquals,omitempty"`
	// Only include jobs with a priority lower than or equal to the given value. Value must be a valid `long` value.
	PriorityLowerThanOrEquals NullableInt64 `json:"priorityLowerThanOrEquals,omitempty"`
	// A JSON array of criteria to sort the result by. Each element of the array is a JSON object that                     specifies one ordering. The position in the array identifies the rank of an ordering, i.e., whether                     it is primary, secondary, etc. The ordering objects have the following properties:                      **Note:** The `sorting` properties will not be applied to the External Task count query.
	Sorting []ExternalTaskQueryDtoSortingInner `json:"sorting,omitempty"`
}

// NewExternalTaskQueryDto instantiates a new ExternalTaskQueryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTaskQueryDto() *ExternalTaskQueryDto {
	this := ExternalTaskQueryDto{}
	return &this
}

// NewExternalTaskQueryDtoWithDefaults instantiates a new ExternalTaskQueryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTaskQueryDtoWithDefaults() *ExternalTaskQueryDto {
	this := ExternalTaskQueryDto{}
	return &this
}

// GetExternalTaskId returns the ExternalTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetExternalTaskId() string {
	if o == nil || IsNil(o.ExternalTaskId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalTaskId.Get()
}

// GetExternalTaskIdOk returns a tuple with the ExternalTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetExternalTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalTaskId.Get(), o.ExternalTaskId.IsSet()
}

// HasExternalTaskId returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasExternalTaskId() bool {
	if o != nil && o.ExternalTaskId.IsSet() {
		return true
	}

	return false
}

// SetExternalTaskId gets a reference to the given NullableString and assigns it to the ExternalTaskId field.
func (o *ExternalTaskQueryDto) SetExternalTaskId(v string) {
	o.ExternalTaskId.Set(&v)
}
// SetExternalTaskIdNil sets the value for ExternalTaskId to be an explicit nil
func (o *ExternalTaskQueryDto) SetExternalTaskIdNil() {
	o.ExternalTaskId.Set(nil)
}

// UnsetExternalTaskId ensures that no value is present for ExternalTaskId, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetExternalTaskId() {
	o.ExternalTaskId.Unset()
}

// GetExternalTaskIdIn returns the ExternalTaskIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetExternalTaskIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExternalTaskIdIn
}

// GetExternalTaskIdInOk returns a tuple with the ExternalTaskIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetExternalTaskIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.ExternalTaskIdIn) {
		return nil, false
	}
	return o.ExternalTaskIdIn, true
}

// HasExternalTaskIdIn returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasExternalTaskIdIn() bool {
	if o != nil && !IsNil(o.ExternalTaskIdIn) {
		return true
	}

	return false
}

// SetExternalTaskIdIn gets a reference to the given []string and assigns it to the ExternalTaskIdIn field.
func (o *ExternalTaskQueryDto) SetExternalTaskIdIn(v []string) {
	o.ExternalTaskIdIn = v
}

// GetTopicName returns the TopicName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetTopicName() string {
	if o == nil || IsNil(o.TopicName.Get()) {
		var ret string
		return ret
	}
	return *o.TopicName.Get()
}

// GetTopicNameOk returns a tuple with the TopicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopicName.Get(), o.TopicName.IsSet()
}

// HasTopicName returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasTopicName() bool {
	if o != nil && o.TopicName.IsSet() {
		return true
	}

	return false
}

// SetTopicName gets a reference to the given NullableString and assigns it to the TopicName field.
func (o *ExternalTaskQueryDto) SetTopicName(v string) {
	o.TopicName.Set(&v)
}
// SetTopicNameNil sets the value for TopicName to be an explicit nil
func (o *ExternalTaskQueryDto) SetTopicNameNil() {
	o.TopicName.Set(nil)
}

// UnsetTopicName ensures that no value is present for TopicName, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetTopicName() {
	o.TopicName.Unset()
}

// GetWorkerId returns the WorkerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetWorkerId() string {
	if o == nil || IsNil(o.WorkerId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkerId.Get()
}

// GetWorkerIdOk returns a tuple with the WorkerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetWorkerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkerId.Get(), o.WorkerId.IsSet()
}

// HasWorkerId returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasWorkerId() bool {
	if o != nil && o.WorkerId.IsSet() {
		return true
	}

	return false
}

// SetWorkerId gets a reference to the given NullableString and assigns it to the WorkerId field.
func (o *ExternalTaskQueryDto) SetWorkerId(v string) {
	o.WorkerId.Set(&v)
}
// SetWorkerIdNil sets the value for WorkerId to be an explicit nil
func (o *ExternalTaskQueryDto) SetWorkerIdNil() {
	o.WorkerId.Set(nil)
}

// UnsetWorkerId ensures that no value is present for WorkerId, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetWorkerId() {
	o.WorkerId.Unset()
}

// GetLocked returns the Locked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetLocked() bool {
	if o == nil || IsNil(o.Locked.Get()) {
		var ret bool
		return ret
	}
	return *o.Locked.Get()
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locked.Get(), o.Locked.IsSet()
}

// HasLocked returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasLocked() bool {
	if o != nil && o.Locked.IsSet() {
		return true
	}

	return false
}

// SetLocked gets a reference to the given NullableBool and assigns it to the Locked field.
func (o *ExternalTaskQueryDto) SetLocked(v bool) {
	o.Locked.Set(&v)
}
// SetLockedNil sets the value for Locked to be an explicit nil
func (o *ExternalTaskQueryDto) SetLockedNil() {
	o.Locked.Set(nil)
}

// UnsetLocked ensures that no value is present for Locked, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetLocked() {
	o.Locked.Unset()
}

// GetNotLocked returns the NotLocked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetNotLocked() bool {
	if o == nil || IsNil(o.NotLocked.Get()) {
		var ret bool
		return ret
	}
	return *o.NotLocked.Get()
}

// GetNotLockedOk returns a tuple with the NotLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetNotLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotLocked.Get(), o.NotLocked.IsSet()
}

// HasNotLocked returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasNotLocked() bool {
	if o != nil && o.NotLocked.IsSet() {
		return true
	}

	return false
}

// SetNotLocked gets a reference to the given NullableBool and assigns it to the NotLocked field.
func (o *ExternalTaskQueryDto) SetNotLocked(v bool) {
	o.NotLocked.Set(&v)
}
// SetNotLockedNil sets the value for NotLocked to be an explicit nil
func (o *ExternalTaskQueryDto) SetNotLockedNil() {
	o.NotLocked.Set(nil)
}

// UnsetNotLocked ensures that no value is present for NotLocked, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetNotLocked() {
	o.NotLocked.Unset()
}

// GetWithRetriesLeft returns the WithRetriesLeft field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetWithRetriesLeft() bool {
	if o == nil || IsNil(o.WithRetriesLeft.Get()) {
		var ret bool
		return ret
	}
	return *o.WithRetriesLeft.Get()
}

// GetWithRetriesLeftOk returns a tuple with the WithRetriesLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetWithRetriesLeftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithRetriesLeft.Get(), o.WithRetriesLeft.IsSet()
}

// HasWithRetriesLeft returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasWithRetriesLeft() bool {
	if o != nil && o.WithRetriesLeft.IsSet() {
		return true
	}

	return false
}

// SetWithRetriesLeft gets a reference to the given NullableBool and assigns it to the WithRetriesLeft field.
func (o *ExternalTaskQueryDto) SetWithRetriesLeft(v bool) {
	o.WithRetriesLeft.Set(&v)
}
// SetWithRetriesLeftNil sets the value for WithRetriesLeft to be an explicit nil
func (o *ExternalTaskQueryDto) SetWithRetriesLeftNil() {
	o.WithRetriesLeft.Set(nil)
}

// UnsetWithRetriesLeft ensures that no value is present for WithRetriesLeft, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetWithRetriesLeft() {
	o.WithRetriesLeft.Unset()
}

// GetNoRetriesLeft returns the NoRetriesLeft field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetNoRetriesLeft() bool {
	if o == nil || IsNil(o.NoRetriesLeft.Get()) {
		var ret bool
		return ret
	}
	return *o.NoRetriesLeft.Get()
}

// GetNoRetriesLeftOk returns a tuple with the NoRetriesLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetNoRetriesLeftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoRetriesLeft.Get(), o.NoRetriesLeft.IsSet()
}

// HasNoRetriesLeft returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasNoRetriesLeft() bool {
	if o != nil && o.NoRetriesLeft.IsSet() {
		return true
	}

	return false
}

// SetNoRetriesLeft gets a reference to the given NullableBool and assigns it to the NoRetriesLeft field.
func (o *ExternalTaskQueryDto) SetNoRetriesLeft(v bool) {
	o.NoRetriesLeft.Set(&v)
}
// SetNoRetriesLeftNil sets the value for NoRetriesLeft to be an explicit nil
func (o *ExternalTaskQueryDto) SetNoRetriesLeftNil() {
	o.NoRetriesLeft.Set(nil)
}

// UnsetNoRetriesLeft ensures that no value is present for NoRetriesLeft, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetNoRetriesLeft() {
	o.NoRetriesLeft.Unset()
}

// GetLockExpirationAfter returns the LockExpirationAfter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetLockExpirationAfter() time.Time {
	if o == nil || IsNil(o.LockExpirationAfter.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LockExpirationAfter.Get()
}

// GetLockExpirationAfterOk returns a tuple with the LockExpirationAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetLockExpirationAfterOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockExpirationAfter.Get(), o.LockExpirationAfter.IsSet()
}

// HasLockExpirationAfter returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasLockExpirationAfter() bool {
	if o != nil && o.LockExpirationAfter.IsSet() {
		return true
	}

	return false
}

// SetLockExpirationAfter gets a reference to the given NullableTime and assigns it to the LockExpirationAfter field.
func (o *ExternalTaskQueryDto) SetLockExpirationAfter(v time.Time) {
	o.LockExpirationAfter.Set(&v)
}
// SetLockExpirationAfterNil sets the value for LockExpirationAfter to be an explicit nil
func (o *ExternalTaskQueryDto) SetLockExpirationAfterNil() {
	o.LockExpirationAfter.Set(nil)
}

// UnsetLockExpirationAfter ensures that no value is present for LockExpirationAfter, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetLockExpirationAfter() {
	o.LockExpirationAfter.Unset()
}

// GetLockExpirationBefore returns the LockExpirationBefore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetLockExpirationBefore() time.Time {
	if o == nil || IsNil(o.LockExpirationBefore.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LockExpirationBefore.Get()
}

// GetLockExpirationBeforeOk returns a tuple with the LockExpirationBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetLockExpirationBeforeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockExpirationBefore.Get(), o.LockExpirationBefore.IsSet()
}

// HasLockExpirationBefore returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasLockExpirationBefore() bool {
	if o != nil && o.LockExpirationBefore.IsSet() {
		return true
	}

	return false
}

// SetLockExpirationBefore gets a reference to the given NullableTime and assigns it to the LockExpirationBefore field.
func (o *ExternalTaskQueryDto) SetLockExpirationBefore(v time.Time) {
	o.LockExpirationBefore.Set(&v)
}
// SetLockExpirationBeforeNil sets the value for LockExpirationBefore to be an explicit nil
func (o *ExternalTaskQueryDto) SetLockExpirationBeforeNil() {
	o.LockExpirationBefore.Set(nil)
}

// UnsetLockExpirationBefore ensures that no value is present for LockExpirationBefore, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetLockExpirationBefore() {
	o.LockExpirationBefore.Unset()
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId.Get()) {
		var ret string
		return ret
	}
	return *o.ActivityId.Get()
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetActivityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityId.Get(), o.ActivityId.IsSet()
}

// HasActivityId returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasActivityId() bool {
	if o != nil && o.ActivityId.IsSet() {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given NullableString and assigns it to the ActivityId field.
func (o *ExternalTaskQueryDto) SetActivityId(v string) {
	o.ActivityId.Set(&v)
}
// SetActivityIdNil sets the value for ActivityId to be an explicit nil
func (o *ExternalTaskQueryDto) SetActivityIdNil() {
	o.ActivityId.Set(nil)
}

// UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetActivityId() {
	o.ActivityId.Unset()
}

// GetActivityIdIn returns the ActivityIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetActivityIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ActivityIdIn
}

// GetActivityIdInOk returns a tuple with the ActivityIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetActivityIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.ActivityIdIn) {
		return nil, false
	}
	return o.ActivityIdIn, true
}

// HasActivityIdIn returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasActivityIdIn() bool {
	if o != nil && !IsNil(o.ActivityIdIn) {
		return true
	}

	return false
}

// SetActivityIdIn gets a reference to the given []string and assigns it to the ActivityIdIn field.
func (o *ExternalTaskQueryDto) SetActivityIdIn(v []string) {
	o.ActivityIdIn = v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionId.Get()
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionId.Get(), o.ExecutionId.IsSet()
}

// HasExecutionId returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasExecutionId() bool {
	if o != nil && o.ExecutionId.IsSet() {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given NullableString and assigns it to the ExecutionId field.
func (o *ExternalTaskQueryDto) SetExecutionId(v string) {
	o.ExecutionId.Set(&v)
}
// SetExecutionIdNil sets the value for ExecutionId to be an explicit nil
func (o *ExternalTaskQueryDto) SetExecutionIdNil() {
	o.ExecutionId.Set(nil)
}

// UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetExecutionId() {
	o.ExecutionId.Unset()
}

// GetProcessInstanceId returns the ProcessInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetProcessInstanceId() string {
	if o == nil || IsNil(o.ProcessInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessInstanceId.Get()
}

// GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetProcessInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessInstanceId.Get(), o.ProcessInstanceId.IsSet()
}

// HasProcessInstanceId returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasProcessInstanceId() bool {
	if o != nil && o.ProcessInstanceId.IsSet() {
		return true
	}

	return false
}

// SetProcessInstanceId gets a reference to the given NullableString and assigns it to the ProcessInstanceId field.
func (o *ExternalTaskQueryDto) SetProcessInstanceId(v string) {
	o.ProcessInstanceId.Set(&v)
}
// SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil
func (o *ExternalTaskQueryDto) SetProcessInstanceIdNil() {
	o.ProcessInstanceId.Set(nil)
}

// UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetProcessInstanceId() {
	o.ProcessInstanceId.Unset()
}

// GetProcessInstanceIdIn returns the ProcessInstanceIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetProcessInstanceIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProcessInstanceIdIn
}

// GetProcessInstanceIdInOk returns a tuple with the ProcessInstanceIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetProcessInstanceIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessInstanceIdIn) {
		return nil, false
	}
	return o.ProcessInstanceIdIn, true
}

// HasProcessInstanceIdIn returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasProcessInstanceIdIn() bool {
	if o != nil && !IsNil(o.ProcessInstanceIdIn) {
		return true
	}

	return false
}

// SetProcessInstanceIdIn gets a reference to the given []string and assigns it to the ProcessInstanceIdIn field.
func (o *ExternalTaskQueryDto) SetProcessInstanceIdIn(v []string) {
	o.ProcessInstanceIdIn = v
}

// GetProcessDefinitionId returns the ProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetProcessDefinitionId() string {
	if o == nil || IsNil(o.ProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionId.Get()
}

// GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionId.Get(), o.ProcessDefinitionId.IsSet()
}

// HasProcessDefinitionId returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasProcessDefinitionId() bool {
	if o != nil && o.ProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionId gets a reference to the given NullableString and assigns it to the ProcessDefinitionId field.
func (o *ExternalTaskQueryDto) SetProcessDefinitionId(v string) {
	o.ProcessDefinitionId.Set(&v)
}
// SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil
func (o *ExternalTaskQueryDto) SetProcessDefinitionIdNil() {
	o.ProcessDefinitionId.Set(nil)
}

// UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetProcessDefinitionId() {
	o.ProcessDefinitionId.Unset()
}

// GetTenantIdIn returns the TenantIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetTenantIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TenantIdIn
}

// GetTenantIdInOk returns a tuple with the TenantIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetTenantIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.TenantIdIn) {
		return nil, false
	}
	return o.TenantIdIn, true
}

// HasTenantIdIn returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasTenantIdIn() bool {
	if o != nil && !IsNil(o.TenantIdIn) {
		return true
	}

	return false
}

// SetTenantIdIn gets a reference to the given []string and assigns it to the TenantIdIn field.
func (o *ExternalTaskQueryDto) SetTenantIdIn(v []string) {
	o.TenantIdIn = v
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetActive() bool {
	if o == nil || IsNil(o.Active.Get()) {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *ExternalTaskQueryDto) SetActive(v bool) {
	o.Active.Set(&v)
}
// SetActiveNil sets the value for Active to be an explicit nil
func (o *ExternalTaskQueryDto) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetActive() {
	o.Active.Unset()
}

// GetSuspended returns the Suspended field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetSuspended() bool {
	if o == nil || IsNil(o.Suspended.Get()) {
		var ret bool
		return ret
	}
	return *o.Suspended.Get()
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetSuspendedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suspended.Get(), o.Suspended.IsSet()
}

// HasSuspended returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasSuspended() bool {
	if o != nil && o.Suspended.IsSet() {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given NullableBool and assigns it to the Suspended field.
func (o *ExternalTaskQueryDto) SetSuspended(v bool) {
	o.Suspended.Set(&v)
}
// SetSuspendedNil sets the value for Suspended to be an explicit nil
func (o *ExternalTaskQueryDto) SetSuspendedNil() {
	o.Suspended.Set(nil)
}

// UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetSuspended() {
	o.Suspended.Unset()
}

// GetPriorityHigherThanOrEquals returns the PriorityHigherThanOrEquals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetPriorityHigherThanOrEquals() int64 {
	if o == nil || IsNil(o.PriorityHigherThanOrEquals.Get()) {
		var ret int64
		return ret
	}
	return *o.PriorityHigherThanOrEquals.Get()
}

// GetPriorityHigherThanOrEqualsOk returns a tuple with the PriorityHigherThanOrEquals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetPriorityHigherThanOrEqualsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriorityHigherThanOrEquals.Get(), o.PriorityHigherThanOrEquals.IsSet()
}

// HasPriorityHigherThanOrEquals returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasPriorityHigherThanOrEquals() bool {
	if o != nil && o.PriorityHigherThanOrEquals.IsSet() {
		return true
	}

	return false
}

// SetPriorityHigherThanOrEquals gets a reference to the given NullableInt64 and assigns it to the PriorityHigherThanOrEquals field.
func (o *ExternalTaskQueryDto) SetPriorityHigherThanOrEquals(v int64) {
	o.PriorityHigherThanOrEquals.Set(&v)
}
// SetPriorityHigherThanOrEqualsNil sets the value for PriorityHigherThanOrEquals to be an explicit nil
func (o *ExternalTaskQueryDto) SetPriorityHigherThanOrEqualsNil() {
	o.PriorityHigherThanOrEquals.Set(nil)
}

// UnsetPriorityHigherThanOrEquals ensures that no value is present for PriorityHigherThanOrEquals, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetPriorityHigherThanOrEquals() {
	o.PriorityHigherThanOrEquals.Unset()
}

// GetPriorityLowerThanOrEquals returns the PriorityLowerThanOrEquals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetPriorityLowerThanOrEquals() int64 {
	if o == nil || IsNil(o.PriorityLowerThanOrEquals.Get()) {
		var ret int64
		return ret
	}
	return *o.PriorityLowerThanOrEquals.Get()
}

// GetPriorityLowerThanOrEqualsOk returns a tuple with the PriorityLowerThanOrEquals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetPriorityLowerThanOrEqualsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriorityLowerThanOrEquals.Get(), o.PriorityLowerThanOrEquals.IsSet()
}

// HasPriorityLowerThanOrEquals returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasPriorityLowerThanOrEquals() bool {
	if o != nil && o.PriorityLowerThanOrEquals.IsSet() {
		return true
	}

	return false
}

// SetPriorityLowerThanOrEquals gets a reference to the given NullableInt64 and assigns it to the PriorityLowerThanOrEquals field.
func (o *ExternalTaskQueryDto) SetPriorityLowerThanOrEquals(v int64) {
	o.PriorityLowerThanOrEquals.Set(&v)
}
// SetPriorityLowerThanOrEqualsNil sets the value for PriorityLowerThanOrEquals to be an explicit nil
func (o *ExternalTaskQueryDto) SetPriorityLowerThanOrEqualsNil() {
	o.PriorityLowerThanOrEquals.Set(nil)
}

// UnsetPriorityLowerThanOrEquals ensures that no value is present for PriorityLowerThanOrEquals, not even an explicit nil
func (o *ExternalTaskQueryDto) UnsetPriorityLowerThanOrEquals() {
	o.PriorityLowerThanOrEquals.Unset()
}

// GetSorting returns the Sorting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalTaskQueryDto) GetSorting() []ExternalTaskQueryDtoSortingInner {
	if o == nil {
		var ret []ExternalTaskQueryDtoSortingInner
		return ret
	}
	return o.Sorting
}

// GetSortingOk returns a tuple with the Sorting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalTaskQueryDto) GetSortingOk() ([]ExternalTaskQueryDtoSortingInner, bool) {
	if o == nil || IsNil(o.Sorting) {
		return nil, false
	}
	return o.Sorting, true
}

// HasSorting returns a boolean if a field has been set.
func (o *ExternalTaskQueryDto) HasSorting() bool {
	if o != nil && !IsNil(o.Sorting) {
		return true
	}

	return false
}

// SetSorting gets a reference to the given []ExternalTaskQueryDtoSortingInner and assigns it to the Sorting field.
func (o *ExternalTaskQueryDto) SetSorting(v []ExternalTaskQueryDtoSortingInner) {
	o.Sorting = v
}

func (o ExternalTaskQueryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalTaskQueryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalTaskId.IsSet() {
		toSerialize["externalTaskId"] = o.ExternalTaskId.Get()
	}
	if o.ExternalTaskIdIn != nil {
		toSerialize["externalTaskIdIn"] = o.ExternalTaskIdIn
	}
	if o.TopicName.IsSet() {
		toSerialize["topicName"] = o.TopicName.Get()
	}
	if o.WorkerId.IsSet() {
		toSerialize["workerId"] = o.WorkerId.Get()
	}
	if o.Locked.IsSet() {
		toSerialize["locked"] = o.Locked.Get()
	}
	if o.NotLocked.IsSet() {
		toSerialize["notLocked"] = o.NotLocked.Get()
	}
	if o.WithRetriesLeft.IsSet() {
		toSerialize["withRetriesLeft"] = o.WithRetriesLeft.Get()
	}
	if o.NoRetriesLeft.IsSet() {
		toSerialize["noRetriesLeft"] = o.NoRetriesLeft.Get()
	}
	if o.LockExpirationAfter.IsSet() {
		toSerialize["lockExpirationAfter"] = o.LockExpirationAfter.Get()
	}
	if o.LockExpirationBefore.IsSet() {
		toSerialize["lockExpirationBefore"] = o.LockExpirationBefore.Get()
	}
	if o.ActivityId.IsSet() {
		toSerialize["activityId"] = o.ActivityId.Get()
	}
	if o.ActivityIdIn != nil {
		toSerialize["activityIdIn"] = o.ActivityIdIn
	}
	if o.ExecutionId.IsSet() {
		toSerialize["executionId"] = o.ExecutionId.Get()
	}
	if o.ProcessInstanceId.IsSet() {
		toSerialize["processInstanceId"] = o.ProcessInstanceId.Get()
	}
	if o.ProcessInstanceIdIn != nil {
		toSerialize["processInstanceIdIn"] = o.ProcessInstanceIdIn
	}
	if o.ProcessDefinitionId.IsSet() {
		toSerialize["processDefinitionId"] = o.ProcessDefinitionId.Get()
	}
	if o.TenantIdIn != nil {
		toSerialize["tenantIdIn"] = o.TenantIdIn
	}
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
	if o.Suspended.IsSet() {
		toSerialize["suspended"] = o.Suspended.Get()
	}
	if o.PriorityHigherThanOrEquals.IsSet() {
		toSerialize["priorityHigherThanOrEquals"] = o.PriorityHigherThanOrEquals.Get()
	}
	if o.PriorityLowerThanOrEquals.IsSet() {
		toSerialize["priorityLowerThanOrEquals"] = o.PriorityLowerThanOrEquals.Get()
	}
	if o.Sorting != nil {
		toSerialize["sorting"] = o.Sorting
	}
	return toSerialize, nil
}

type NullableExternalTaskQueryDto struct {
	value *ExternalTaskQueryDto
	isSet bool
}

func (v NullableExternalTaskQueryDto) Get() *ExternalTaskQueryDto {
	return v.value
}

func (v *NullableExternalTaskQueryDto) Set(val *ExternalTaskQueryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTaskQueryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTaskQueryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTaskQueryDto(val *ExternalTaskQueryDto) *NullableExternalTaskQueryDto {
	return &NullableExternalTaskQueryDto{value: val, isSet: true}
}

func (v NullableExternalTaskQueryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTaskQueryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


