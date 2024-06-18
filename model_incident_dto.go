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

// checks if the IncidentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncidentDto{}

// IncidentDto struct for IncidentDto
type IncidentDto struct {
	// The id of the incident.
	Id NullableString `json:"id,omitempty"`
	// The id of the process definition this incident is associated with.
	ProcessDefinitionId NullableString `json:"processDefinitionId,omitempty"`
	// The id of the process instance this incident is associated with.
	ProcessInstanceId NullableString `json:"processInstanceId,omitempty"`
	// The id of the execution this incident is associated with.
	ExecutionId NullableString `json:"executionId,omitempty"`
	// The time this incident happened. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`, e.g., `2013-01-23T14:42:45.000+0200`.
	IncidentTimestamp NullableTime `json:"incidentTimestamp,omitempty"`
	// The type of incident, for example: `failedJobs` will be returned in case of an incident which identified a failed job during the execution of a process instance. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types.
	IncidentType NullableString `json:"incidentType,omitempty"`
	// The id of the activity this incident is associated with.
	ActivityId NullableString `json:"activityId,omitempty"`
	// The id of the activity on which the last exception occurred.
	FailedActivityId NullableString `json:"failedActivityId,omitempty"`
	// The id of the associated cause incident which has been triggered.
	CauseIncidentId NullableString `json:"causeIncidentId,omitempty"`
	// The id of the associated root cause incident which has been triggered.
	RootCauseIncidentId NullableString `json:"rootCauseIncidentId,omitempty"`
	// The payload of this incident.
	Configuration NullableString `json:"configuration,omitempty"`
	// The id of the tenant this incident is associated with.
	TenantId NullableString `json:"tenantId,omitempty"`
	// The message of this incident.
	IncidentMessage NullableString `json:"incidentMessage,omitempty"`
	// The job definition id the incident is associated with.
	JobDefinitionId NullableString `json:"jobDefinitionId,omitempty"`
	// The annotation set to the incident.
	Annotation NullableString `json:"annotation,omitempty"`
}

// NewIncidentDto instantiates a new IncidentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentDto() *IncidentDto {
	this := IncidentDto{}
	return &this
}

// NewIncidentDtoWithDefaults instantiates a new IncidentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentDtoWithDefaults() *IncidentDto {
	this := IncidentDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *IncidentDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *IncidentDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *IncidentDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *IncidentDto) UnsetId() {
	o.Id.Unset()
}

// GetProcessDefinitionId returns the ProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetProcessDefinitionId() string {
	if o == nil || IsNil(o.ProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionId.Get()
}

// GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionId.Get(), o.ProcessDefinitionId.IsSet()
}

// HasProcessDefinitionId returns a boolean if a field has been set.
func (o *IncidentDto) HasProcessDefinitionId() bool {
	if o != nil && o.ProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionId gets a reference to the given NullableString and assigns it to the ProcessDefinitionId field.
func (o *IncidentDto) SetProcessDefinitionId(v string) {
	o.ProcessDefinitionId.Set(&v)
}
// SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil
func (o *IncidentDto) SetProcessDefinitionIdNil() {
	o.ProcessDefinitionId.Set(nil)
}

// UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
func (o *IncidentDto) UnsetProcessDefinitionId() {
	o.ProcessDefinitionId.Unset()
}

// GetProcessInstanceId returns the ProcessInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetProcessInstanceId() string {
	if o == nil || IsNil(o.ProcessInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessInstanceId.Get()
}

// GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetProcessInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessInstanceId.Get(), o.ProcessInstanceId.IsSet()
}

// HasProcessInstanceId returns a boolean if a field has been set.
func (o *IncidentDto) HasProcessInstanceId() bool {
	if o != nil && o.ProcessInstanceId.IsSet() {
		return true
	}

	return false
}

// SetProcessInstanceId gets a reference to the given NullableString and assigns it to the ProcessInstanceId field.
func (o *IncidentDto) SetProcessInstanceId(v string) {
	o.ProcessInstanceId.Set(&v)
}
// SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil
func (o *IncidentDto) SetProcessInstanceIdNil() {
	o.ProcessInstanceId.Set(nil)
}

// UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
func (o *IncidentDto) UnsetProcessInstanceId() {
	o.ProcessInstanceId.Unset()
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionId.Get()
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionId.Get(), o.ExecutionId.IsSet()
}

// HasExecutionId returns a boolean if a field has been set.
func (o *IncidentDto) HasExecutionId() bool {
	if o != nil && o.ExecutionId.IsSet() {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given NullableString and assigns it to the ExecutionId field.
func (o *IncidentDto) SetExecutionId(v string) {
	o.ExecutionId.Set(&v)
}
// SetExecutionIdNil sets the value for ExecutionId to be an explicit nil
func (o *IncidentDto) SetExecutionIdNil() {
	o.ExecutionId.Set(nil)
}

// UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
func (o *IncidentDto) UnsetExecutionId() {
	o.ExecutionId.Unset()
}

// GetIncidentTimestamp returns the IncidentTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetIncidentTimestamp() time.Time {
	if o == nil || IsNil(o.IncidentTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.IncidentTimestamp.Get()
}

// GetIncidentTimestampOk returns a tuple with the IncidentTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetIncidentTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncidentTimestamp.Get(), o.IncidentTimestamp.IsSet()
}

// HasIncidentTimestamp returns a boolean if a field has been set.
func (o *IncidentDto) HasIncidentTimestamp() bool {
	if o != nil && o.IncidentTimestamp.IsSet() {
		return true
	}

	return false
}

// SetIncidentTimestamp gets a reference to the given NullableTime and assigns it to the IncidentTimestamp field.
func (o *IncidentDto) SetIncidentTimestamp(v time.Time) {
	o.IncidentTimestamp.Set(&v)
}
// SetIncidentTimestampNil sets the value for IncidentTimestamp to be an explicit nil
func (o *IncidentDto) SetIncidentTimestampNil() {
	o.IncidentTimestamp.Set(nil)
}

// UnsetIncidentTimestamp ensures that no value is present for IncidentTimestamp, not even an explicit nil
func (o *IncidentDto) UnsetIncidentTimestamp() {
	o.IncidentTimestamp.Unset()
}

// GetIncidentType returns the IncidentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetIncidentType() string {
	if o == nil || IsNil(o.IncidentType.Get()) {
		var ret string
		return ret
	}
	return *o.IncidentType.Get()
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetIncidentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncidentType.Get(), o.IncidentType.IsSet()
}

// HasIncidentType returns a boolean if a field has been set.
func (o *IncidentDto) HasIncidentType() bool {
	if o != nil && o.IncidentType.IsSet() {
		return true
	}

	return false
}

// SetIncidentType gets a reference to the given NullableString and assigns it to the IncidentType field.
func (o *IncidentDto) SetIncidentType(v string) {
	o.IncidentType.Set(&v)
}
// SetIncidentTypeNil sets the value for IncidentType to be an explicit nil
func (o *IncidentDto) SetIncidentTypeNil() {
	o.IncidentType.Set(nil)
}

// UnsetIncidentType ensures that no value is present for IncidentType, not even an explicit nil
func (o *IncidentDto) UnsetIncidentType() {
	o.IncidentType.Unset()
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId.Get()) {
		var ret string
		return ret
	}
	return *o.ActivityId.Get()
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetActivityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivityId.Get(), o.ActivityId.IsSet()
}

// HasActivityId returns a boolean if a field has been set.
func (o *IncidentDto) HasActivityId() bool {
	if o != nil && o.ActivityId.IsSet() {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given NullableString and assigns it to the ActivityId field.
func (o *IncidentDto) SetActivityId(v string) {
	o.ActivityId.Set(&v)
}
// SetActivityIdNil sets the value for ActivityId to be an explicit nil
func (o *IncidentDto) SetActivityIdNil() {
	o.ActivityId.Set(nil)
}

// UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
func (o *IncidentDto) UnsetActivityId() {
	o.ActivityId.Unset()
}

// GetFailedActivityId returns the FailedActivityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetFailedActivityId() string {
	if o == nil || IsNil(o.FailedActivityId.Get()) {
		var ret string
		return ret
	}
	return *o.FailedActivityId.Get()
}

// GetFailedActivityIdOk returns a tuple with the FailedActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetFailedActivityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedActivityId.Get(), o.FailedActivityId.IsSet()
}

// HasFailedActivityId returns a boolean if a field has been set.
func (o *IncidentDto) HasFailedActivityId() bool {
	if o != nil && o.FailedActivityId.IsSet() {
		return true
	}

	return false
}

// SetFailedActivityId gets a reference to the given NullableString and assigns it to the FailedActivityId field.
func (o *IncidentDto) SetFailedActivityId(v string) {
	o.FailedActivityId.Set(&v)
}
// SetFailedActivityIdNil sets the value for FailedActivityId to be an explicit nil
func (o *IncidentDto) SetFailedActivityIdNil() {
	o.FailedActivityId.Set(nil)
}

// UnsetFailedActivityId ensures that no value is present for FailedActivityId, not even an explicit nil
func (o *IncidentDto) UnsetFailedActivityId() {
	o.FailedActivityId.Unset()
}

// GetCauseIncidentId returns the CauseIncidentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetCauseIncidentId() string {
	if o == nil || IsNil(o.CauseIncidentId.Get()) {
		var ret string
		return ret
	}
	return *o.CauseIncidentId.Get()
}

// GetCauseIncidentIdOk returns a tuple with the CauseIncidentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetCauseIncidentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CauseIncidentId.Get(), o.CauseIncidentId.IsSet()
}

// HasCauseIncidentId returns a boolean if a field has been set.
func (o *IncidentDto) HasCauseIncidentId() bool {
	if o != nil && o.CauseIncidentId.IsSet() {
		return true
	}

	return false
}

// SetCauseIncidentId gets a reference to the given NullableString and assigns it to the CauseIncidentId field.
func (o *IncidentDto) SetCauseIncidentId(v string) {
	o.CauseIncidentId.Set(&v)
}
// SetCauseIncidentIdNil sets the value for CauseIncidentId to be an explicit nil
func (o *IncidentDto) SetCauseIncidentIdNil() {
	o.CauseIncidentId.Set(nil)
}

// UnsetCauseIncidentId ensures that no value is present for CauseIncidentId, not even an explicit nil
func (o *IncidentDto) UnsetCauseIncidentId() {
	o.CauseIncidentId.Unset()
}

// GetRootCauseIncidentId returns the RootCauseIncidentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetRootCauseIncidentId() string {
	if o == nil || IsNil(o.RootCauseIncidentId.Get()) {
		var ret string
		return ret
	}
	return *o.RootCauseIncidentId.Get()
}

// GetRootCauseIncidentIdOk returns a tuple with the RootCauseIncidentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetRootCauseIncidentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootCauseIncidentId.Get(), o.RootCauseIncidentId.IsSet()
}

// HasRootCauseIncidentId returns a boolean if a field has been set.
func (o *IncidentDto) HasRootCauseIncidentId() bool {
	if o != nil && o.RootCauseIncidentId.IsSet() {
		return true
	}

	return false
}

// SetRootCauseIncidentId gets a reference to the given NullableString and assigns it to the RootCauseIncidentId field.
func (o *IncidentDto) SetRootCauseIncidentId(v string) {
	o.RootCauseIncidentId.Set(&v)
}
// SetRootCauseIncidentIdNil sets the value for RootCauseIncidentId to be an explicit nil
func (o *IncidentDto) SetRootCauseIncidentIdNil() {
	o.RootCauseIncidentId.Set(nil)
}

// UnsetRootCauseIncidentId ensures that no value is present for RootCauseIncidentId, not even an explicit nil
func (o *IncidentDto) UnsetRootCauseIncidentId() {
	o.RootCauseIncidentId.Unset()
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetConfiguration() string {
	if o == nil || IsNil(o.Configuration.Get()) {
		var ret string
		return ret
	}
	return *o.Configuration.Get()
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetConfigurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configuration.Get(), o.Configuration.IsSet()
}

// HasConfiguration returns a boolean if a field has been set.
func (o *IncidentDto) HasConfiguration() bool {
	if o != nil && o.Configuration.IsSet() {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given NullableString and assigns it to the Configuration field.
func (o *IncidentDto) SetConfiguration(v string) {
	o.Configuration.Set(&v)
}
// SetConfigurationNil sets the value for Configuration to be an explicit nil
func (o *IncidentDto) SetConfigurationNil() {
	o.Configuration.Set(nil)
}

// UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
func (o *IncidentDto) UnsetConfiguration() {
	o.Configuration.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *IncidentDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *IncidentDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *IncidentDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *IncidentDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetIncidentMessage returns the IncidentMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetIncidentMessage() string {
	if o == nil || IsNil(o.IncidentMessage.Get()) {
		var ret string
		return ret
	}
	return *o.IncidentMessage.Get()
}

// GetIncidentMessageOk returns a tuple with the IncidentMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetIncidentMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncidentMessage.Get(), o.IncidentMessage.IsSet()
}

// HasIncidentMessage returns a boolean if a field has been set.
func (o *IncidentDto) HasIncidentMessage() bool {
	if o != nil && o.IncidentMessage.IsSet() {
		return true
	}

	return false
}

// SetIncidentMessage gets a reference to the given NullableString and assigns it to the IncidentMessage field.
func (o *IncidentDto) SetIncidentMessage(v string) {
	o.IncidentMessage.Set(&v)
}
// SetIncidentMessageNil sets the value for IncidentMessage to be an explicit nil
func (o *IncidentDto) SetIncidentMessageNil() {
	o.IncidentMessage.Set(nil)
}

// UnsetIncidentMessage ensures that no value is present for IncidentMessage, not even an explicit nil
func (o *IncidentDto) UnsetIncidentMessage() {
	o.IncidentMessage.Unset()
}

// GetJobDefinitionId returns the JobDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetJobDefinitionId() string {
	if o == nil || IsNil(o.JobDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.JobDefinitionId.Get()
}

// GetJobDefinitionIdOk returns a tuple with the JobDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetJobDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobDefinitionId.Get(), o.JobDefinitionId.IsSet()
}

// HasJobDefinitionId returns a boolean if a field has been set.
func (o *IncidentDto) HasJobDefinitionId() bool {
	if o != nil && o.JobDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetJobDefinitionId gets a reference to the given NullableString and assigns it to the JobDefinitionId field.
func (o *IncidentDto) SetJobDefinitionId(v string) {
	o.JobDefinitionId.Set(&v)
}
// SetJobDefinitionIdNil sets the value for JobDefinitionId to be an explicit nil
func (o *IncidentDto) SetJobDefinitionIdNil() {
	o.JobDefinitionId.Set(nil)
}

// UnsetJobDefinitionId ensures that no value is present for JobDefinitionId, not even an explicit nil
func (o *IncidentDto) UnsetJobDefinitionId() {
	o.JobDefinitionId.Unset()
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentDto) GetAnnotation() string {
	if o == nil || IsNil(o.Annotation.Get()) {
		var ret string
		return ret
	}
	return *o.Annotation.Get()
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentDto) GetAnnotationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Annotation.Get(), o.Annotation.IsSet()
}

// HasAnnotation returns a boolean if a field has been set.
func (o *IncidentDto) HasAnnotation() bool {
	if o != nil && o.Annotation.IsSet() {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given NullableString and assigns it to the Annotation field.
func (o *IncidentDto) SetAnnotation(v string) {
	o.Annotation.Set(&v)
}
// SetAnnotationNil sets the value for Annotation to be an explicit nil
func (o *IncidentDto) SetAnnotationNil() {
	o.Annotation.Set(nil)
}

// UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil
func (o *IncidentDto) UnsetAnnotation() {
	o.Annotation.Unset()
}

func (o IncidentDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncidentDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ProcessDefinitionId.IsSet() {
		toSerialize["processDefinitionId"] = o.ProcessDefinitionId.Get()
	}
	if o.ProcessInstanceId.IsSet() {
		toSerialize["processInstanceId"] = o.ProcessInstanceId.Get()
	}
	if o.ExecutionId.IsSet() {
		toSerialize["executionId"] = o.ExecutionId.Get()
	}
	if o.IncidentTimestamp.IsSet() {
		toSerialize["incidentTimestamp"] = o.IncidentTimestamp.Get()
	}
	if o.IncidentType.IsSet() {
		toSerialize["incidentType"] = o.IncidentType.Get()
	}
	if o.ActivityId.IsSet() {
		toSerialize["activityId"] = o.ActivityId.Get()
	}
	if o.FailedActivityId.IsSet() {
		toSerialize["failedActivityId"] = o.FailedActivityId.Get()
	}
	if o.CauseIncidentId.IsSet() {
		toSerialize["causeIncidentId"] = o.CauseIncidentId.Get()
	}
	if o.RootCauseIncidentId.IsSet() {
		toSerialize["rootCauseIncidentId"] = o.RootCauseIncidentId.Get()
	}
	if o.Configuration.IsSet() {
		toSerialize["configuration"] = o.Configuration.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.IncidentMessage.IsSet() {
		toSerialize["incidentMessage"] = o.IncidentMessage.Get()
	}
	if o.JobDefinitionId.IsSet() {
		toSerialize["jobDefinitionId"] = o.JobDefinitionId.Get()
	}
	if o.Annotation.IsSet() {
		toSerialize["annotation"] = o.Annotation.Get()
	}
	return toSerialize, nil
}

type NullableIncidentDto struct {
	value *IncidentDto
	isSet bool
}

func (v NullableIncidentDto) Get() *IncidentDto {
	return v.value
}

func (v *NullableIncidentDto) Set(val *IncidentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentDto(val *IncidentDto) *NullableIncidentDto {
	return &NullableIncidentDto{value: val, isSet: true}
}

func (v NullableIncidentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

