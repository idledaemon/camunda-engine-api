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

// checks if the JobDefinitionQueryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobDefinitionQueryDto{}

// JobDefinitionQueryDto A Job definition query which defines a list of Job definitions
type JobDefinitionQueryDto struct {
	// Filter by job definition id.
	JobDefinitionId NullableString `json:"jobDefinitionId,omitempty"`
	// Only include job definitions which belong to one of the passed activity ids.
	ActivityIdIn []string `json:"activityIdIn,omitempty"`
	// Only include job definitions which exist for the given process definition id.
	ProcessDefinitionId NullableString `json:"processDefinitionId,omitempty"`
	// Only include job definitions which exist for the given process definition key.
	ProcessDefinitionKey NullableString `json:"processDefinitionKey,omitempty"`
	// Only include job definitions which exist for the given job type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job types.
	JobType NullableString `json:"jobType,omitempty"`
	// Only include job definitions which exist for the given job configuration. For example: for timer jobs it is the timer configuration.
	JobConfiguration NullableString `json:"jobConfiguration,omitempty"`
	// Only include active job definitions. Value may only be `true`, as `false` is the default behavior.
	Active NullableBool `json:"active,omitempty"`
	// Only include suspended job definitions. Value may only be `true`, as `false` is the default behavior.
	Suspended NullableBool `json:"suspended,omitempty"`
	// Only include job definitions that have an overriding job priority defined. The only effective value is `true`. If set to `false`, this filter is not applied.
	WithOverridingJobPriority NullableBool `json:"withOverridingJobPriority,omitempty"`
	// Only include job definitions which belong to one of the passed tenant ids.
	TenantIdIn []string `json:"tenantIdIn,omitempty"`
	// Only include job definitions which belong to no tenant. Value may only be `true`, as `false` is the default behavior.
	WithoutTenantId NullableBool `json:"withoutTenantId,omitempty"`
	// Include job definitions which belong to no tenant. Can be used in combination with `tenantIdIn`. Value may only be `true`, as `false` is the default behavior.
	IncludeJobDefinitionsWithoutTenantId NullableBool `json:"includeJobDefinitionsWithoutTenantId,omitempty"`
	// An array of criteria to sort the result by. Each element of the array is                        an object that specifies one ordering. The position in the array                        identifies the rank of an ordering, i.e., whether it is primary, secondary,                        etc. Sorting has no effect for `count` endpoints.
	Sorting []JobDefinitionQueryDtoSortingInner `json:"sorting,omitempty"`
}

// NewJobDefinitionQueryDto instantiates a new JobDefinitionQueryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDefinitionQueryDto() *JobDefinitionQueryDto {
	this := JobDefinitionQueryDto{}
	return &this
}

// NewJobDefinitionQueryDtoWithDefaults instantiates a new JobDefinitionQueryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDefinitionQueryDtoWithDefaults() *JobDefinitionQueryDto {
	this := JobDefinitionQueryDto{}
	return &this
}

// GetJobDefinitionId returns the JobDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetJobDefinitionId() string {
	if o == nil || IsNil(o.JobDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.JobDefinitionId.Get()
}

// GetJobDefinitionIdOk returns a tuple with the JobDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetJobDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobDefinitionId.Get(), o.JobDefinitionId.IsSet()
}

// HasJobDefinitionId returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasJobDefinitionId() bool {
	if o != nil && o.JobDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetJobDefinitionId gets a reference to the given NullableString and assigns it to the JobDefinitionId field.
func (o *JobDefinitionQueryDto) SetJobDefinitionId(v string) {
	o.JobDefinitionId.Set(&v)
}
// SetJobDefinitionIdNil sets the value for JobDefinitionId to be an explicit nil
func (o *JobDefinitionQueryDto) SetJobDefinitionIdNil() {
	o.JobDefinitionId.Set(nil)
}

// UnsetJobDefinitionId ensures that no value is present for JobDefinitionId, not even an explicit nil
func (o *JobDefinitionQueryDto) UnsetJobDefinitionId() {
	o.JobDefinitionId.Unset()
}

// GetActivityIdIn returns the ActivityIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetActivityIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ActivityIdIn
}

// GetActivityIdInOk returns a tuple with the ActivityIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetActivityIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.ActivityIdIn) {
		return nil, false
	}
	return o.ActivityIdIn, true
}

// HasActivityIdIn returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasActivityIdIn() bool {
	if o != nil && !IsNil(o.ActivityIdIn) {
		return true
	}

	return false
}

// SetActivityIdIn gets a reference to the given []string and assigns it to the ActivityIdIn field.
func (o *JobDefinitionQueryDto) SetActivityIdIn(v []string) {
	o.ActivityIdIn = v
}

// GetProcessDefinitionId returns the ProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetProcessDefinitionId() string {
	if o == nil || IsNil(o.ProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionId.Get()
}

// GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionId.Get(), o.ProcessDefinitionId.IsSet()
}

// HasProcessDefinitionId returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasProcessDefinitionId() bool {
	if o != nil && o.ProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionId gets a reference to the given NullableString and assigns it to the ProcessDefinitionId field.
func (o *JobDefinitionQueryDto) SetProcessDefinitionId(v string) {
	o.ProcessDefinitionId.Set(&v)
}
// SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil
func (o *JobDefinitionQueryDto) SetProcessDefinitionIdNil() {
	o.ProcessDefinitionId.Set(nil)
}

// UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
func (o *JobDefinitionQueryDto) UnsetProcessDefinitionId() {
	o.ProcessDefinitionId.Unset()
}

// GetProcessDefinitionKey returns the ProcessDefinitionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetProcessDefinitionKey() string {
	if o == nil || IsNil(o.ProcessDefinitionKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionKey.Get()
}

// GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetProcessDefinitionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionKey.Get(), o.ProcessDefinitionKey.IsSet()
}

// HasProcessDefinitionKey returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasProcessDefinitionKey() bool {
	if o != nil && o.ProcessDefinitionKey.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionKey gets a reference to the given NullableString and assigns it to the ProcessDefinitionKey field.
func (o *JobDefinitionQueryDto) SetProcessDefinitionKey(v string) {
	o.ProcessDefinitionKey.Set(&v)
}
// SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil
func (o *JobDefinitionQueryDto) SetProcessDefinitionKeyNil() {
	o.ProcessDefinitionKey.Set(nil)
}

// UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
func (o *JobDefinitionQueryDto) UnsetProcessDefinitionKey() {
	o.ProcessDefinitionKey.Unset()
}

// GetJobType returns the JobType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetJobType() string {
	if o == nil || IsNil(o.JobType.Get()) {
		var ret string
		return ret
	}
	return *o.JobType.Get()
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetJobTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobType.Get(), o.JobType.IsSet()
}

// HasJobType returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasJobType() bool {
	if o != nil && o.JobType.IsSet() {
		return true
	}

	return false
}

// SetJobType gets a reference to the given NullableString and assigns it to the JobType field.
func (o *JobDefinitionQueryDto) SetJobType(v string) {
	o.JobType.Set(&v)
}
// SetJobTypeNil sets the value for JobType to be an explicit nil
func (o *JobDefinitionQueryDto) SetJobTypeNil() {
	o.JobType.Set(nil)
}

// UnsetJobType ensures that no value is present for JobType, not even an explicit nil
func (o *JobDefinitionQueryDto) UnsetJobType() {
	o.JobType.Unset()
}

// GetJobConfiguration returns the JobConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetJobConfiguration() string {
	if o == nil || IsNil(o.JobConfiguration.Get()) {
		var ret string
		return ret
	}
	return *o.JobConfiguration.Get()
}

// GetJobConfigurationOk returns a tuple with the JobConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetJobConfigurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobConfiguration.Get(), o.JobConfiguration.IsSet()
}

// HasJobConfiguration returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasJobConfiguration() bool {
	if o != nil && o.JobConfiguration.IsSet() {
		return true
	}

	return false
}

// SetJobConfiguration gets a reference to the given NullableString and assigns it to the JobConfiguration field.
func (o *JobDefinitionQueryDto) SetJobConfiguration(v string) {
	o.JobConfiguration.Set(&v)
}
// SetJobConfigurationNil sets the value for JobConfiguration to be an explicit nil
func (o *JobDefinitionQueryDto) SetJobConfigurationNil() {
	o.JobConfiguration.Set(nil)
}

// UnsetJobConfiguration ensures that no value is present for JobConfiguration, not even an explicit nil
func (o *JobDefinitionQueryDto) UnsetJobConfiguration() {
	o.JobConfiguration.Unset()
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetActive() bool {
	if o == nil || IsNil(o.Active.Get()) {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *JobDefinitionQueryDto) SetActive(v bool) {
	o.Active.Set(&v)
}
// SetActiveNil sets the value for Active to be an explicit nil
func (o *JobDefinitionQueryDto) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *JobDefinitionQueryDto) UnsetActive() {
	o.Active.Unset()
}

// GetSuspended returns the Suspended field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetSuspended() bool {
	if o == nil || IsNil(o.Suspended.Get()) {
		var ret bool
		return ret
	}
	return *o.Suspended.Get()
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetSuspendedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suspended.Get(), o.Suspended.IsSet()
}

// HasSuspended returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasSuspended() bool {
	if o != nil && o.Suspended.IsSet() {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given NullableBool and assigns it to the Suspended field.
func (o *JobDefinitionQueryDto) SetSuspended(v bool) {
	o.Suspended.Set(&v)
}
// SetSuspendedNil sets the value for Suspended to be an explicit nil
func (o *JobDefinitionQueryDto) SetSuspendedNil() {
	o.Suspended.Set(nil)
}

// UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
func (o *JobDefinitionQueryDto) UnsetSuspended() {
	o.Suspended.Unset()
}

// GetWithOverridingJobPriority returns the WithOverridingJobPriority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetWithOverridingJobPriority() bool {
	if o == nil || IsNil(o.WithOverridingJobPriority.Get()) {
		var ret bool
		return ret
	}
	return *o.WithOverridingJobPriority.Get()
}

// GetWithOverridingJobPriorityOk returns a tuple with the WithOverridingJobPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetWithOverridingJobPriorityOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithOverridingJobPriority.Get(), o.WithOverridingJobPriority.IsSet()
}

// HasWithOverridingJobPriority returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasWithOverridingJobPriority() bool {
	if o != nil && o.WithOverridingJobPriority.IsSet() {
		return true
	}

	return false
}

// SetWithOverridingJobPriority gets a reference to the given NullableBool and assigns it to the WithOverridingJobPriority field.
func (o *JobDefinitionQueryDto) SetWithOverridingJobPriority(v bool) {
	o.WithOverridingJobPriority.Set(&v)
}
// SetWithOverridingJobPriorityNil sets the value for WithOverridingJobPriority to be an explicit nil
func (o *JobDefinitionQueryDto) SetWithOverridingJobPriorityNil() {
	o.WithOverridingJobPriority.Set(nil)
}

// UnsetWithOverridingJobPriority ensures that no value is present for WithOverridingJobPriority, not even an explicit nil
func (o *JobDefinitionQueryDto) UnsetWithOverridingJobPriority() {
	o.WithOverridingJobPriority.Unset()
}

// GetTenantIdIn returns the TenantIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetTenantIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TenantIdIn
}

// GetTenantIdInOk returns a tuple with the TenantIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetTenantIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.TenantIdIn) {
		return nil, false
	}
	return o.TenantIdIn, true
}

// HasTenantIdIn returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasTenantIdIn() bool {
	if o != nil && !IsNil(o.TenantIdIn) {
		return true
	}

	return false
}

// SetTenantIdIn gets a reference to the given []string and assigns it to the TenantIdIn field.
func (o *JobDefinitionQueryDto) SetTenantIdIn(v []string) {
	o.TenantIdIn = v
}

// GetWithoutTenantId returns the WithoutTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetWithoutTenantId() bool {
	if o == nil || IsNil(o.WithoutTenantId.Get()) {
		var ret bool
		return ret
	}
	return *o.WithoutTenantId.Get()
}

// GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetWithoutTenantIdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithoutTenantId.Get(), o.WithoutTenantId.IsSet()
}

// HasWithoutTenantId returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasWithoutTenantId() bool {
	if o != nil && o.WithoutTenantId.IsSet() {
		return true
	}

	return false
}

// SetWithoutTenantId gets a reference to the given NullableBool and assigns it to the WithoutTenantId field.
func (o *JobDefinitionQueryDto) SetWithoutTenantId(v bool) {
	o.WithoutTenantId.Set(&v)
}
// SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil
func (o *JobDefinitionQueryDto) SetWithoutTenantIdNil() {
	o.WithoutTenantId.Set(nil)
}

// UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
func (o *JobDefinitionQueryDto) UnsetWithoutTenantId() {
	o.WithoutTenantId.Unset()
}

// GetIncludeJobDefinitionsWithoutTenantId returns the IncludeJobDefinitionsWithoutTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetIncludeJobDefinitionsWithoutTenantId() bool {
	if o == nil || IsNil(o.IncludeJobDefinitionsWithoutTenantId.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeJobDefinitionsWithoutTenantId.Get()
}

// GetIncludeJobDefinitionsWithoutTenantIdOk returns a tuple with the IncludeJobDefinitionsWithoutTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetIncludeJobDefinitionsWithoutTenantIdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeJobDefinitionsWithoutTenantId.Get(), o.IncludeJobDefinitionsWithoutTenantId.IsSet()
}

// HasIncludeJobDefinitionsWithoutTenantId returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasIncludeJobDefinitionsWithoutTenantId() bool {
	if o != nil && o.IncludeJobDefinitionsWithoutTenantId.IsSet() {
		return true
	}

	return false
}

// SetIncludeJobDefinitionsWithoutTenantId gets a reference to the given NullableBool and assigns it to the IncludeJobDefinitionsWithoutTenantId field.
func (o *JobDefinitionQueryDto) SetIncludeJobDefinitionsWithoutTenantId(v bool) {
	o.IncludeJobDefinitionsWithoutTenantId.Set(&v)
}
// SetIncludeJobDefinitionsWithoutTenantIdNil sets the value for IncludeJobDefinitionsWithoutTenantId to be an explicit nil
func (o *JobDefinitionQueryDto) SetIncludeJobDefinitionsWithoutTenantIdNil() {
	o.IncludeJobDefinitionsWithoutTenantId.Set(nil)
}

// UnsetIncludeJobDefinitionsWithoutTenantId ensures that no value is present for IncludeJobDefinitionsWithoutTenantId, not even an explicit nil
func (o *JobDefinitionQueryDto) UnsetIncludeJobDefinitionsWithoutTenantId() {
	o.IncludeJobDefinitionsWithoutTenantId.Unset()
}

// GetSorting returns the Sorting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobDefinitionQueryDto) GetSorting() []JobDefinitionQueryDtoSortingInner {
	if o == nil {
		var ret []JobDefinitionQueryDtoSortingInner
		return ret
	}
	return o.Sorting
}

// GetSortingOk returns a tuple with the Sorting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDefinitionQueryDto) GetSortingOk() ([]JobDefinitionQueryDtoSortingInner, bool) {
	if o == nil || IsNil(o.Sorting) {
		return nil, false
	}
	return o.Sorting, true
}

// HasSorting returns a boolean if a field has been set.
func (o *JobDefinitionQueryDto) HasSorting() bool {
	if o != nil && !IsNil(o.Sorting) {
		return true
	}

	return false
}

// SetSorting gets a reference to the given []JobDefinitionQueryDtoSortingInner and assigns it to the Sorting field.
func (o *JobDefinitionQueryDto) SetSorting(v []JobDefinitionQueryDtoSortingInner) {
	o.Sorting = v
}

func (o JobDefinitionQueryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobDefinitionQueryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.JobDefinitionId.IsSet() {
		toSerialize["jobDefinitionId"] = o.JobDefinitionId.Get()
	}
	if o.ActivityIdIn != nil {
		toSerialize["activityIdIn"] = o.ActivityIdIn
	}
	if o.ProcessDefinitionId.IsSet() {
		toSerialize["processDefinitionId"] = o.ProcessDefinitionId.Get()
	}
	if o.ProcessDefinitionKey.IsSet() {
		toSerialize["processDefinitionKey"] = o.ProcessDefinitionKey.Get()
	}
	if o.JobType.IsSet() {
		toSerialize["jobType"] = o.JobType.Get()
	}
	if o.JobConfiguration.IsSet() {
		toSerialize["jobConfiguration"] = o.JobConfiguration.Get()
	}
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
	if o.Suspended.IsSet() {
		toSerialize["suspended"] = o.Suspended.Get()
	}
	if o.WithOverridingJobPriority.IsSet() {
		toSerialize["withOverridingJobPriority"] = o.WithOverridingJobPriority.Get()
	}
	if o.TenantIdIn != nil {
		toSerialize["tenantIdIn"] = o.TenantIdIn
	}
	if o.WithoutTenantId.IsSet() {
		toSerialize["withoutTenantId"] = o.WithoutTenantId.Get()
	}
	if o.IncludeJobDefinitionsWithoutTenantId.IsSet() {
		toSerialize["includeJobDefinitionsWithoutTenantId"] = o.IncludeJobDefinitionsWithoutTenantId.Get()
	}
	if o.Sorting != nil {
		toSerialize["sorting"] = o.Sorting
	}
	return toSerialize, nil
}

type NullableJobDefinitionQueryDto struct {
	value *JobDefinitionQueryDto
	isSet bool
}

func (v NullableJobDefinitionQueryDto) Get() *JobDefinitionQueryDto {
	return v.value
}

func (v *NullableJobDefinitionQueryDto) Set(val *JobDefinitionQueryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDefinitionQueryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDefinitionQueryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDefinitionQueryDto(val *JobDefinitionQueryDto) *NullableJobDefinitionQueryDto {
	return &NullableJobDefinitionQueryDto{value: val, isSet: true}
}

func (v NullableJobDefinitionQueryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDefinitionQueryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


