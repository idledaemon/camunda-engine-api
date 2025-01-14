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

// checks if the HistoricExternalTaskLogDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoricExternalTaskLogDto{}

// HistoricExternalTaskLogDto struct for HistoricExternalTaskLogDto
type HistoricExternalTaskLogDto struct {
	// The id of the log entry.
	Id NullableString `json:"id,omitempty"`
	// The id of the external task.
	ExternalTaskId NullableString `json:"externalTaskId,omitempty"`
	// The time when the log entry has been written.
	Timestamp NullableTime `json:"timestamp,omitempty"`
	// The topic name of the associated external task.
	TopicName NullableString `json:"topicName,omitempty"`
	// The id of the worker that posessed the most recent lock.
	WorkerId NullableString `json:"workerId,omitempty"`
	// The number of retries the associated external task has left.
	Retries NullableInt32 `json:"retries,omitempty"`
	// The execution priority the external task had when the log entry was created.
	Priority NullableInt64 `json:"priority,omitempty"`
	// The message of the error that occurred by executing the associated external task.
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	// The id of the activity on which the associated external task was created.
	ActivityId NullableString `json:"activityId,omitempty"`
	// The id of the activity instance on which the associated external task was created.
	ActivityInstanceId NullableString `json:"activityInstanceId,omitempty"`
	// The execution id on which the associated external task was created.
	ExecutionId NullableString `json:"executionId,omitempty"`
	// The id of the process instance on which the associated external task was created.
	ProcessInstanceId NullableString `json:"processInstanceId,omitempty"`
	// The id of the process definition which the associated external task belongs to.
	ProcessDefinitionId NullableString `json:"processDefinitionId,omitempty"`
	// The key of the process definition which the associated external task belongs to.
	ProcessDefinitionKey NullableString `json:"processDefinitionKey,omitempty"`
	// The id of the tenant that this historic external task log entry belongs to.
	TenantId NullableString `json:"tenantId,omitempty"`
	// A flag indicating whether this log represents the creation of the associated external task.
	CreationLog NullableBool `json:"creationLog,omitempty"`
	// A flag indicating whether this log represents the failed execution of the associated external task.
	FailureLog NullableBool `json:"failureLog,omitempty"`
	// A flag indicating whether this log represents the successful execution of the associated external task.
	SuccessLog NullableBool `json:"successLog,omitempty"`
	// A flag indicating whether this log represents the deletion of the associated external task.
	DeletionLog NullableBool `json:"deletionLog,omitempty"`
	// The time after which this log should be removed by the History Cleanup job. Default format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.  For further information, please see the [documentation](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/)
	RemovalTime NullableTime `json:"removalTime,omitempty"`
	// The process instance id of the root process instance that initiated the process containing this log.
	RootProcessInstanceId NullableString `json:"rootProcessInstanceId,omitempty"`
}

// NewHistoricExternalTaskLogDto instantiates a new HistoricExternalTaskLogDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricExternalTaskLogDto() *HistoricExternalTaskLogDto {
	this := HistoricExternalTaskLogDto{}
	return &this
}

// NewHistoricExternalTaskLogDtoWithDefaults instantiates a new HistoricExternalTaskLogDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricExternalTaskLogDtoWithDefaults() *HistoricExternalTaskLogDto {
	this := HistoricExternalTaskLogDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HistoricExternalTaskLogDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetId() {
	o.Id.Unset()
}

// GetExternalTaskId returns the ExternalTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetExternalTaskId() string {
	if o == nil || IsNil(o.ExternalTaskId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalTaskId.Get()
}

// GetExternalTaskIdOk returns a tuple with the ExternalTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetExternalTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalTaskId.Get(), o.ExternalTaskId.IsSet()
}

// HasExternalTaskId returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasExternalTaskId() bool {
	if o != nil && o.ExternalTaskId.IsSet() {
		return true
	}

	return false
}

// SetExternalTaskId gets a reference to the given NullableString and assigns it to the ExternalTaskId field.
func (o *HistoricExternalTaskLogDto) SetExternalTaskId(v string) {
	o.ExternalTaskId.Set(&v)
}
// SetExternalTaskIdNil sets the value for ExternalTaskId to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetExternalTaskIdNil() {
	o.ExternalTaskId.Set(nil)
}

// UnsetExternalTaskId ensures that no value is present for ExternalTaskId, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetExternalTaskId() {
	o.ExternalTaskId.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *HistoricExternalTaskLogDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetTopicName returns the TopicName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetTopicName() string {
	if o == nil || IsNil(o.TopicName.Get()) {
		var ret string
		return ret
	}
	return *o.TopicName.Get()
}

// GetTopicNameOk returns a tuple with the TopicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopicName.Get(), o.TopicName.IsSet()
}

// HasTopicName returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasTopicName() bool {
	if o != nil && o.TopicName.IsSet() {
		return true
	}

	return false
}

// SetTopicName gets a reference to the given NullableString and assigns it to the TopicName field.
func (o *HistoricExternalTaskLogDto) SetTopicName(v string) {
	o.TopicName.Set(&v)
}
// SetTopicNameNil sets the value for TopicName to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetTopicNameNil() {
	o.TopicName.Set(nil)
}

// UnsetTopicName ensures that no value is present for TopicName, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetTopicName() {
	o.TopicName.Unset()
}

// GetWorkerId returns the WorkerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetWorkerId() string {
	if o == nil || IsNil(o.WorkerId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkerId.Get()
}

// GetWorkerIdOk returns a tuple with the WorkerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetWorkerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkerId.Get(), o.WorkerId.IsSet()
}

// HasWorkerId returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasWorkerId() bool {
	if o != nil && o.WorkerId.IsSet() {
		return true
	}

	return false
}

// SetWorkerId gets a reference to the given NullableString and assigns it to the WorkerId field.
func (o *HistoricExternalTaskLogDto) SetWorkerId(v string) {
	o.WorkerId.Set(&v)
}
// SetWorkerIdNil sets the value for WorkerId to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetWorkerIdNil() {
	o.WorkerId.Set(nil)
}

// UnsetWorkerId ensures that no value is present for WorkerId, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetWorkerId() {
	o.WorkerId.Unset()
}

// GetRetries returns the Retries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetRetries() int32 {
	if o == nil || IsNil(o.Retries.Get()) {
		var ret int32
		return ret
	}
	return *o.Retries.Get()
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetRetriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Retries.Get(), o.Retries.IsSet()
}

// HasRetries returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasRetries() bool {
	if o != nil && o.Retries.IsSet() {
		return true
	}

	return false
}

// SetRetries gets a reference to the given NullableInt32 and assigns it to the Retries field.
func (o *HistoricExternalTaskLogDto) SetRetries(v int32) {
	o.Retries.Set(&v)
}
// SetRetriesNil sets the value for Retries to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetRetriesNil() {
	o.Retries.Set(nil)
}

// UnsetRetries ensures that no value is present for Retries, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetRetries() {
	o.Retries.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetPriority() int64 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int64
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt64 and assigns it to the Priority field.
func (o *HistoricExternalTaskLogDto) SetPriority(v int64) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetPriority() {
	o.Priority.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *HistoricExternalTaskLogDto) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId.Get()) {
		var ret string
		return ret
	}
	return *o.ActivityId.Get()
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetActivityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityId.Get(), o.ActivityId.IsSet()
}

// HasActivityId returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasActivityId() bool {
	if o != nil && o.ActivityId.IsSet() {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given NullableString and assigns it to the ActivityId field.
func (o *HistoricExternalTaskLogDto) SetActivityId(v string) {
	o.ActivityId.Set(&v)
}
// SetActivityIdNil sets the value for ActivityId to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetActivityIdNil() {
	o.ActivityId.Set(nil)
}

// UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetActivityId() {
	o.ActivityId.Unset()
}

// GetActivityInstanceId returns the ActivityInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetActivityInstanceId() string {
	if o == nil || IsNil(o.ActivityInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.ActivityInstanceId.Get()
}

// GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetActivityInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityInstanceId.Get(), o.ActivityInstanceId.IsSet()
}

// HasActivityInstanceId returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasActivityInstanceId() bool {
	if o != nil && o.ActivityInstanceId.IsSet() {
		return true
	}

	return false
}

// SetActivityInstanceId gets a reference to the given NullableString and assigns it to the ActivityInstanceId field.
func (o *HistoricExternalTaskLogDto) SetActivityInstanceId(v string) {
	o.ActivityInstanceId.Set(&v)
}
// SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetActivityInstanceIdNil() {
	o.ActivityInstanceId.Set(nil)
}

// UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetActivityInstanceId() {
	o.ActivityInstanceId.Unset()
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionId.Get()
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionId.Get(), o.ExecutionId.IsSet()
}

// HasExecutionId returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasExecutionId() bool {
	if o != nil && o.ExecutionId.IsSet() {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given NullableString and assigns it to the ExecutionId field.
func (o *HistoricExternalTaskLogDto) SetExecutionId(v string) {
	o.ExecutionId.Set(&v)
}
// SetExecutionIdNil sets the value for ExecutionId to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetExecutionIdNil() {
	o.ExecutionId.Set(nil)
}

// UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetExecutionId() {
	o.ExecutionId.Unset()
}

// GetProcessInstanceId returns the ProcessInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetProcessInstanceId() string {
	if o == nil || IsNil(o.ProcessInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessInstanceId.Get()
}

// GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetProcessInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessInstanceId.Get(), o.ProcessInstanceId.IsSet()
}

// HasProcessInstanceId returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasProcessInstanceId() bool {
	if o != nil && o.ProcessInstanceId.IsSet() {
		return true
	}

	return false
}

// SetProcessInstanceId gets a reference to the given NullableString and assigns it to the ProcessInstanceId field.
func (o *HistoricExternalTaskLogDto) SetProcessInstanceId(v string) {
	o.ProcessInstanceId.Set(&v)
}
// SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetProcessInstanceIdNil() {
	o.ProcessInstanceId.Set(nil)
}

// UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetProcessInstanceId() {
	o.ProcessInstanceId.Unset()
}

// GetProcessDefinitionId returns the ProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetProcessDefinitionId() string {
	if o == nil || IsNil(o.ProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionId.Get()
}

// GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionId.Get(), o.ProcessDefinitionId.IsSet()
}

// HasProcessDefinitionId returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasProcessDefinitionId() bool {
	if o != nil && o.ProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionId gets a reference to the given NullableString and assigns it to the ProcessDefinitionId field.
func (o *HistoricExternalTaskLogDto) SetProcessDefinitionId(v string) {
	o.ProcessDefinitionId.Set(&v)
}
// SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetProcessDefinitionIdNil() {
	o.ProcessDefinitionId.Set(nil)
}

// UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetProcessDefinitionId() {
	o.ProcessDefinitionId.Unset()
}

// GetProcessDefinitionKey returns the ProcessDefinitionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetProcessDefinitionKey() string {
	if o == nil || IsNil(o.ProcessDefinitionKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionKey.Get()
}

// GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetProcessDefinitionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionKey.Get(), o.ProcessDefinitionKey.IsSet()
}

// HasProcessDefinitionKey returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasProcessDefinitionKey() bool {
	if o != nil && o.ProcessDefinitionKey.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionKey gets a reference to the given NullableString and assigns it to the ProcessDefinitionKey field.
func (o *HistoricExternalTaskLogDto) SetProcessDefinitionKey(v string) {
	o.ProcessDefinitionKey.Set(&v)
}
// SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetProcessDefinitionKeyNil() {
	o.ProcessDefinitionKey.Set(nil)
}

// UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetProcessDefinitionKey() {
	o.ProcessDefinitionKey.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *HistoricExternalTaskLogDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetCreationLog returns the CreationLog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetCreationLog() bool {
	if o == nil || IsNil(o.CreationLog.Get()) {
		var ret bool
		return ret
	}
	return *o.CreationLog.Get()
}

// GetCreationLogOk returns a tuple with the CreationLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetCreationLogOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreationLog.Get(), o.CreationLog.IsSet()
}

// HasCreationLog returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasCreationLog() bool {
	if o != nil && o.CreationLog.IsSet() {
		return true
	}

	return false
}

// SetCreationLog gets a reference to the given NullableBool and assigns it to the CreationLog field.
func (o *HistoricExternalTaskLogDto) SetCreationLog(v bool) {
	o.CreationLog.Set(&v)
}
// SetCreationLogNil sets the value for CreationLog to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetCreationLogNil() {
	o.CreationLog.Set(nil)
}

// UnsetCreationLog ensures that no value is present for CreationLog, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetCreationLog() {
	o.CreationLog.Unset()
}

// GetFailureLog returns the FailureLog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetFailureLog() bool {
	if o == nil || IsNil(o.FailureLog.Get()) {
		var ret bool
		return ret
	}
	return *o.FailureLog.Get()
}

// GetFailureLogOk returns a tuple with the FailureLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetFailureLogOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureLog.Get(), o.FailureLog.IsSet()
}

// HasFailureLog returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasFailureLog() bool {
	if o != nil && o.FailureLog.IsSet() {
		return true
	}

	return false
}

// SetFailureLog gets a reference to the given NullableBool and assigns it to the FailureLog field.
func (o *HistoricExternalTaskLogDto) SetFailureLog(v bool) {
	o.FailureLog.Set(&v)
}
// SetFailureLogNil sets the value for FailureLog to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetFailureLogNil() {
	o.FailureLog.Set(nil)
}

// UnsetFailureLog ensures that no value is present for FailureLog, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetFailureLog() {
	o.FailureLog.Unset()
}

// GetSuccessLog returns the SuccessLog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetSuccessLog() bool {
	if o == nil || IsNil(o.SuccessLog.Get()) {
		var ret bool
		return ret
	}
	return *o.SuccessLog.Get()
}

// GetSuccessLogOk returns a tuple with the SuccessLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetSuccessLogOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuccessLog.Get(), o.SuccessLog.IsSet()
}

// HasSuccessLog returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasSuccessLog() bool {
	if o != nil && o.SuccessLog.IsSet() {
		return true
	}

	return false
}

// SetSuccessLog gets a reference to the given NullableBool and assigns it to the SuccessLog field.
func (o *HistoricExternalTaskLogDto) SetSuccessLog(v bool) {
	o.SuccessLog.Set(&v)
}
// SetSuccessLogNil sets the value for SuccessLog to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetSuccessLogNil() {
	o.SuccessLog.Set(nil)
}

// UnsetSuccessLog ensures that no value is present for SuccessLog, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetSuccessLog() {
	o.SuccessLog.Unset()
}

// GetDeletionLog returns the DeletionLog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetDeletionLog() bool {
	if o == nil || IsNil(o.DeletionLog.Get()) {
		var ret bool
		return ret
	}
	return *o.DeletionLog.Get()
}

// GetDeletionLogOk returns a tuple with the DeletionLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetDeletionLogOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletionLog.Get(), o.DeletionLog.IsSet()
}

// HasDeletionLog returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasDeletionLog() bool {
	if o != nil && o.DeletionLog.IsSet() {
		return true
	}

	return false
}

// SetDeletionLog gets a reference to the given NullableBool and assigns it to the DeletionLog field.
func (o *HistoricExternalTaskLogDto) SetDeletionLog(v bool) {
	o.DeletionLog.Set(&v)
}
// SetDeletionLogNil sets the value for DeletionLog to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetDeletionLogNil() {
	o.DeletionLog.Set(nil)
}

// UnsetDeletionLog ensures that no value is present for DeletionLog, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetDeletionLog() {
	o.DeletionLog.Unset()
}

// GetRemovalTime returns the RemovalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetRemovalTime() time.Time {
	if o == nil || IsNil(o.RemovalTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RemovalTime.Get()
}

// GetRemovalTimeOk returns a tuple with the RemovalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetRemovalTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovalTime.Get(), o.RemovalTime.IsSet()
}

// HasRemovalTime returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasRemovalTime() bool {
	if o != nil && o.RemovalTime.IsSet() {
		return true
	}

	return false
}

// SetRemovalTime gets a reference to the given NullableTime and assigns it to the RemovalTime field.
func (o *HistoricExternalTaskLogDto) SetRemovalTime(v time.Time) {
	o.RemovalTime.Set(&v)
}
// SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetRemovalTimeNil() {
	o.RemovalTime.Set(nil)
}

// UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetRemovalTime() {
	o.RemovalTime.Unset()
}

// GetRootProcessInstanceId returns the RootProcessInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricExternalTaskLogDto) GetRootProcessInstanceId() string {
	if o == nil || IsNil(o.RootProcessInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.RootProcessInstanceId.Get()
}

// GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricExternalTaskLogDto) GetRootProcessInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootProcessInstanceId.Get(), o.RootProcessInstanceId.IsSet()
}

// HasRootProcessInstanceId returns a boolean if a field has been set.
func (o *HistoricExternalTaskLogDto) HasRootProcessInstanceId() bool {
	if o != nil && o.RootProcessInstanceId.IsSet() {
		return true
	}

	return false
}

// SetRootProcessInstanceId gets a reference to the given NullableString and assigns it to the RootProcessInstanceId field.
func (o *HistoricExternalTaskLogDto) SetRootProcessInstanceId(v string) {
	o.RootProcessInstanceId.Set(&v)
}
// SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil
func (o *HistoricExternalTaskLogDto) SetRootProcessInstanceIdNil() {
	o.RootProcessInstanceId.Set(nil)
}

// UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil
func (o *HistoricExternalTaskLogDto) UnsetRootProcessInstanceId() {
	o.RootProcessInstanceId.Unset()
}

func (o HistoricExternalTaskLogDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricExternalTaskLogDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ExternalTaskId.IsSet() {
		toSerialize["externalTaskId"] = o.ExternalTaskId.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.TopicName.IsSet() {
		toSerialize["topicName"] = o.TopicName.Get()
	}
	if o.WorkerId.IsSet() {
		toSerialize["workerId"] = o.WorkerId.Get()
	}
	if o.Retries.IsSet() {
		toSerialize["retries"] = o.Retries.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.ActivityId.IsSet() {
		toSerialize["activityId"] = o.ActivityId.Get()
	}
	if o.ActivityInstanceId.IsSet() {
		toSerialize["activityInstanceId"] = o.ActivityInstanceId.Get()
	}
	if o.ExecutionId.IsSet() {
		toSerialize["executionId"] = o.ExecutionId.Get()
	}
	if o.ProcessInstanceId.IsSet() {
		toSerialize["processInstanceId"] = o.ProcessInstanceId.Get()
	}
	if o.ProcessDefinitionId.IsSet() {
		toSerialize["processDefinitionId"] = o.ProcessDefinitionId.Get()
	}
	if o.ProcessDefinitionKey.IsSet() {
		toSerialize["processDefinitionKey"] = o.ProcessDefinitionKey.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.CreationLog.IsSet() {
		toSerialize["creationLog"] = o.CreationLog.Get()
	}
	if o.FailureLog.IsSet() {
		toSerialize["failureLog"] = o.FailureLog.Get()
	}
	if o.SuccessLog.IsSet() {
		toSerialize["successLog"] = o.SuccessLog.Get()
	}
	if o.DeletionLog.IsSet() {
		toSerialize["deletionLog"] = o.DeletionLog.Get()
	}
	if o.RemovalTime.IsSet() {
		toSerialize["removalTime"] = o.RemovalTime.Get()
	}
	if o.RootProcessInstanceId.IsSet() {
		toSerialize["rootProcessInstanceId"] = o.RootProcessInstanceId.Get()
	}
	return toSerialize, nil
}

type NullableHistoricExternalTaskLogDto struct {
	value *HistoricExternalTaskLogDto
	isSet bool
}

func (v NullableHistoricExternalTaskLogDto) Get() *HistoricExternalTaskLogDto {
	return v.value
}

func (v *NullableHistoricExternalTaskLogDto) Set(val *HistoricExternalTaskLogDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricExternalTaskLogDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricExternalTaskLogDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricExternalTaskLogDto(val *HistoricExternalTaskLogDto) *NullableHistoricExternalTaskLogDto {
	return &NullableHistoricExternalTaskLogDto{value: val, isSet: true}
}

func (v NullableHistoricExternalTaskLogDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricExternalTaskLogDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


