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

// checks if the HistoricIdentityLinkLogDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoricIdentityLinkLogDto{}

// HistoricIdentityLinkLogDto struct for HistoricIdentityLinkLogDto
type HistoricIdentityLinkLogDto struct {
	// Id of the Historic identity link entry.
	Id NullableString `json:"id,omitempty"`
	// The time when the identity link is logged.  [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	Time NullableTime `json:"time,omitempty"`
	// The type of identity link (candidate/assignee/owner).
	Type NullableString `json:"type,omitempty"`
	// The id of the user/assignee.
	UserId NullableString `json:"userId,omitempty"`
	// The id of the group.
	GroupId NullableString `json:"groupId,omitempty"`
	// The id of the task.
	TaskId NullableString `json:"taskId,omitempty"`
	// The id of the process definition.
	ProcessDefinitionId NullableString `json:"processDefinitionId,omitempty"`
	// The key of the process definition.
	ProcessDefinitionKey NullableString `json:"processDefinitionKey,omitempty"`
	// Type of operation (add/delete).
	OperationType NullableString `json:"operationType,omitempty"`
	// The id of the assigner.
	AssignerId NullableString `json:"assignerId,omitempty"`
	// The id of the tenant.
	TenantId NullableString `json:"tenantId,omitempty"`
	// The time after which the identity link should be removed by the History Cleanup job.  [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	RemovalTime NullableTime `json:"removalTime,omitempty"`
	// The process instance id of the root process instance that initiated the process containing this identity link.
	RootProcessInstanceId NullableString `json:"rootProcessInstanceId,omitempty"`
}

// NewHistoricIdentityLinkLogDto instantiates a new HistoricIdentityLinkLogDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricIdentityLinkLogDto() *HistoricIdentityLinkLogDto {
	this := HistoricIdentityLinkLogDto{}
	return &this
}

// NewHistoricIdentityLinkLogDtoWithDefaults instantiates a new HistoricIdentityLinkLogDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricIdentityLinkLogDtoWithDefaults() *HistoricIdentityLinkLogDto {
	this := HistoricIdentityLinkLogDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HistoricIdentityLinkLogDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetId() {
	o.Id.Unset()
}

// GetTime returns the Time field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetTime() time.Time {
	if o == nil || IsNil(o.Time.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Time.Get()
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Time.Get(), o.Time.IsSet()
}

// HasTime returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasTime() bool {
	if o != nil && o.Time.IsSet() {
		return true
	}

	return false
}

// SetTime gets a reference to the given NullableTime and assigns it to the Time field.
func (o *HistoricIdentityLinkLogDto) SetTime(v time.Time) {
	o.Time.Set(&v)
}
// SetTimeNil sets the value for Time to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetTimeNil() {
	o.Time.Set(nil)
}

// UnsetTime ensures that no value is present for Time, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetTime() {
	o.Time.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *HistoricIdentityLinkLogDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetType() {
	o.Type.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *HistoricIdentityLinkLogDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetGroupId() string {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret string
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableString and assigns it to the GroupId field.
func (o *HistoricIdentityLinkLogDto) SetGroupId(v string) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetTaskId returns the TaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetTaskId() string {
	if o == nil || IsNil(o.TaskId.Get()) {
		var ret string
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// HasTaskId returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasTaskId() bool {
	if o != nil && o.TaskId.IsSet() {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given NullableString and assigns it to the TaskId field.
func (o *HistoricIdentityLinkLogDto) SetTaskId(v string) {
	o.TaskId.Set(&v)
}
// SetTaskIdNil sets the value for TaskId to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetTaskIdNil() {
	o.TaskId.Set(nil)
}

// UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetTaskId() {
	o.TaskId.Unset()
}

// GetProcessDefinitionId returns the ProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetProcessDefinitionId() string {
	if o == nil || IsNil(o.ProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionId.Get()
}

// GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionId.Get(), o.ProcessDefinitionId.IsSet()
}

// HasProcessDefinitionId returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasProcessDefinitionId() bool {
	if o != nil && o.ProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionId gets a reference to the given NullableString and assigns it to the ProcessDefinitionId field.
func (o *HistoricIdentityLinkLogDto) SetProcessDefinitionId(v string) {
	o.ProcessDefinitionId.Set(&v)
}
// SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetProcessDefinitionIdNil() {
	o.ProcessDefinitionId.Set(nil)
}

// UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetProcessDefinitionId() {
	o.ProcessDefinitionId.Unset()
}

// GetProcessDefinitionKey returns the ProcessDefinitionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetProcessDefinitionKey() string {
	if o == nil || IsNil(o.ProcessDefinitionKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionKey.Get()
}

// GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetProcessDefinitionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionKey.Get(), o.ProcessDefinitionKey.IsSet()
}

// HasProcessDefinitionKey returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasProcessDefinitionKey() bool {
	if o != nil && o.ProcessDefinitionKey.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionKey gets a reference to the given NullableString and assigns it to the ProcessDefinitionKey field.
func (o *HistoricIdentityLinkLogDto) SetProcessDefinitionKey(v string) {
	o.ProcessDefinitionKey.Set(&v)
}
// SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetProcessDefinitionKeyNil() {
	o.ProcessDefinitionKey.Set(nil)
}

// UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetProcessDefinitionKey() {
	o.ProcessDefinitionKey.Unset()
}

// GetOperationType returns the OperationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetOperationType() string {
	if o == nil || IsNil(o.OperationType.Get()) {
		var ret string
		return ret
	}
	return *o.OperationType.Get()
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetOperationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationType.Get(), o.OperationType.IsSet()
}

// HasOperationType returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasOperationType() bool {
	if o != nil && o.OperationType.IsSet() {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given NullableString and assigns it to the OperationType field.
func (o *HistoricIdentityLinkLogDto) SetOperationType(v string) {
	o.OperationType.Set(&v)
}
// SetOperationTypeNil sets the value for OperationType to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetOperationTypeNil() {
	o.OperationType.Set(nil)
}

// UnsetOperationType ensures that no value is present for OperationType, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetOperationType() {
	o.OperationType.Unset()
}

// GetAssignerId returns the AssignerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetAssignerId() string {
	if o == nil || IsNil(o.AssignerId.Get()) {
		var ret string
		return ret
	}
	return *o.AssignerId.Get()
}

// GetAssignerIdOk returns a tuple with the AssignerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetAssignerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignerId.Get(), o.AssignerId.IsSet()
}

// HasAssignerId returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasAssignerId() bool {
	if o != nil && o.AssignerId.IsSet() {
		return true
	}

	return false
}

// SetAssignerId gets a reference to the given NullableString and assigns it to the AssignerId field.
func (o *HistoricIdentityLinkLogDto) SetAssignerId(v string) {
	o.AssignerId.Set(&v)
}
// SetAssignerIdNil sets the value for AssignerId to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetAssignerIdNil() {
	o.AssignerId.Set(nil)
}

// UnsetAssignerId ensures that no value is present for AssignerId, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetAssignerId() {
	o.AssignerId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *HistoricIdentityLinkLogDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetRemovalTime returns the RemovalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetRemovalTime() time.Time {
	if o == nil || IsNil(o.RemovalTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RemovalTime.Get()
}

// GetRemovalTimeOk returns a tuple with the RemovalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetRemovalTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovalTime.Get(), o.RemovalTime.IsSet()
}

// HasRemovalTime returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasRemovalTime() bool {
	if o != nil && o.RemovalTime.IsSet() {
		return true
	}

	return false
}

// SetRemovalTime gets a reference to the given NullableTime and assigns it to the RemovalTime field.
func (o *HistoricIdentityLinkLogDto) SetRemovalTime(v time.Time) {
	o.RemovalTime.Set(&v)
}
// SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetRemovalTimeNil() {
	o.RemovalTime.Set(nil)
}

// UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetRemovalTime() {
	o.RemovalTime.Unset()
}

// GetRootProcessInstanceId returns the RootProcessInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricIdentityLinkLogDto) GetRootProcessInstanceId() string {
	if o == nil || IsNil(o.RootProcessInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.RootProcessInstanceId.Get()
}

// GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricIdentityLinkLogDto) GetRootProcessInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootProcessInstanceId.Get(), o.RootProcessInstanceId.IsSet()
}

// HasRootProcessInstanceId returns a boolean if a field has been set.
func (o *HistoricIdentityLinkLogDto) HasRootProcessInstanceId() bool {
	if o != nil && o.RootProcessInstanceId.IsSet() {
		return true
	}

	return false
}

// SetRootProcessInstanceId gets a reference to the given NullableString and assigns it to the RootProcessInstanceId field.
func (o *HistoricIdentityLinkLogDto) SetRootProcessInstanceId(v string) {
	o.RootProcessInstanceId.Set(&v)
}
// SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil
func (o *HistoricIdentityLinkLogDto) SetRootProcessInstanceIdNil() {
	o.RootProcessInstanceId.Set(nil)
}

// UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil
func (o *HistoricIdentityLinkLogDto) UnsetRootProcessInstanceId() {
	o.RootProcessInstanceId.Unset()
}

func (o HistoricIdentityLinkLogDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricIdentityLinkLogDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Time.IsSet() {
		toSerialize["time"] = o.Time.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.GroupId.IsSet() {
		toSerialize["groupId"] = o.GroupId.Get()
	}
	if o.TaskId.IsSet() {
		toSerialize["taskId"] = o.TaskId.Get()
	}
	if o.ProcessDefinitionId.IsSet() {
		toSerialize["processDefinitionId"] = o.ProcessDefinitionId.Get()
	}
	if o.ProcessDefinitionKey.IsSet() {
		toSerialize["processDefinitionKey"] = o.ProcessDefinitionKey.Get()
	}
	if o.OperationType.IsSet() {
		toSerialize["operationType"] = o.OperationType.Get()
	}
	if o.AssignerId.IsSet() {
		toSerialize["assignerId"] = o.AssignerId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.RemovalTime.IsSet() {
		toSerialize["removalTime"] = o.RemovalTime.Get()
	}
	if o.RootProcessInstanceId.IsSet() {
		toSerialize["rootProcessInstanceId"] = o.RootProcessInstanceId.Get()
	}
	return toSerialize, nil
}

type NullableHistoricIdentityLinkLogDto struct {
	value *HistoricIdentityLinkLogDto
	isSet bool
}

func (v NullableHistoricIdentityLinkLogDto) Get() *HistoricIdentityLinkLogDto {
	return v.value
}

func (v *NullableHistoricIdentityLinkLogDto) Set(val *HistoricIdentityLinkLogDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricIdentityLinkLogDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricIdentityLinkLogDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricIdentityLinkLogDto(val *HistoricIdentityLinkLogDto) *NullableHistoricIdentityLinkLogDto {
	return &NullableHistoricIdentityLinkLogDto{value: val, isSet: true}
}

func (v NullableHistoricIdentityLinkLogDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricIdentityLinkLogDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


