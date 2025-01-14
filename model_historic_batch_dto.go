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

// checks if the HistoricBatchDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoricBatchDto{}

// HistoricBatchDto struct for HistoricBatchDto
type HistoricBatchDto struct {
	// The id of the batch.
	Id NullableString `json:"id,omitempty"`
	// The type of the batch. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types.
	Type NullableString `json:"type,omitempty"`
	//  The total jobs of a batch is the number of batch execution jobs required to complete the batch. 
	TotalJobs NullableInt32 `json:"totalJobs,omitempty"`
	//  The number of batch execution jobs created per seed job invocation. The batch seed job is invoked until it has created all batch execution jobs required by the batch (see `totalJobs` property). 
	BatchJobsPerSeed NullableInt32 `json:"batchJobsPerSeed,omitempty"`
	//  Every batch execution job invokes the command executed by the batch `invocationsPerBatchJob` times. E.g., for a process instance migration batch this specifies the number of process instances which are migrated per batch execution job. 
	InvocationsPerBatchJob NullableInt32 `json:"invocationsPerBatchJob,omitempty"`
	// The job definition id for the seed jobs of this batch.
	SeedJobDefinitionId NullableString `json:"seedJobDefinitionId,omitempty"`
	// The job definition id for the monitor jobs of this batch.
	MonitorJobDefinitionId NullableString `json:"monitorJobDefinitionId,omitempty"`
	// The job definition id for the batch execution jobs of this batch.
	BatchJobDefinitionId NullableString `json:"batchJobDefinitionId,omitempty"`
	// The tenant id of the batch.
	TenantId NullableString `json:"tenantId,omitempty"`
	// The batch creator's user id.
	CreateUserId NullableString `json:"createUserId,omitempty"`
	// The time the batch was started. Default format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`. For further information, please see the [documentation](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/)
	StartTime NullableTime `json:"startTime,omitempty"`
	// The time the batch execution was started, i.e., at least one batch job has been executed. Default format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`. For further information, please see the [documentation] (https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/)
	ExecutionStartTime NullableTime `json:"executionStartTime,omitempty"`
	// The time the batch ended. Default format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`. For further information, please see the [documentation](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/)
	EndTime NullableTime `json:"endTime,omitempty"`
	// The time after which the historic batch should be removed by the History Cleanup job. Default format `yyyy-MM-dd'T'HH:mm:ss.SSSZ`. For further information, please see the [documentation](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/)
	RemovalTime NullableTime `json:"removalTime,omitempty"`
}

// NewHistoricBatchDto instantiates a new HistoricBatchDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricBatchDto() *HistoricBatchDto {
	this := HistoricBatchDto{}
	return &this
}

// NewHistoricBatchDtoWithDefaults instantiates a new HistoricBatchDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricBatchDtoWithDefaults() *HistoricBatchDto {
	this := HistoricBatchDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HistoricBatchDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *HistoricBatchDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HistoricBatchDto) UnsetId() {
	o.Id.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *HistoricBatchDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *HistoricBatchDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *HistoricBatchDto) UnsetType() {
	o.Type.Unset()
}

// GetTotalJobs returns the TotalJobs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetTotalJobs() int32 {
	if o == nil || IsNil(o.TotalJobs.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalJobs.Get()
}

// GetTotalJobsOk returns a tuple with the TotalJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetTotalJobsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalJobs.Get(), o.TotalJobs.IsSet()
}

// HasTotalJobs returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasTotalJobs() bool {
	if o != nil && o.TotalJobs.IsSet() {
		return true
	}

	return false
}

// SetTotalJobs gets a reference to the given NullableInt32 and assigns it to the TotalJobs field.
func (o *HistoricBatchDto) SetTotalJobs(v int32) {
	o.TotalJobs.Set(&v)
}
// SetTotalJobsNil sets the value for TotalJobs to be an explicit nil
func (o *HistoricBatchDto) SetTotalJobsNil() {
	o.TotalJobs.Set(nil)
}

// UnsetTotalJobs ensures that no value is present for TotalJobs, not even an explicit nil
func (o *HistoricBatchDto) UnsetTotalJobs() {
	o.TotalJobs.Unset()
}

// GetBatchJobsPerSeed returns the BatchJobsPerSeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetBatchJobsPerSeed() int32 {
	if o == nil || IsNil(o.BatchJobsPerSeed.Get()) {
		var ret int32
		return ret
	}
	return *o.BatchJobsPerSeed.Get()
}

// GetBatchJobsPerSeedOk returns a tuple with the BatchJobsPerSeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetBatchJobsPerSeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BatchJobsPerSeed.Get(), o.BatchJobsPerSeed.IsSet()
}

// HasBatchJobsPerSeed returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasBatchJobsPerSeed() bool {
	if o != nil && o.BatchJobsPerSeed.IsSet() {
		return true
	}

	return false
}

// SetBatchJobsPerSeed gets a reference to the given NullableInt32 and assigns it to the BatchJobsPerSeed field.
func (o *HistoricBatchDto) SetBatchJobsPerSeed(v int32) {
	o.BatchJobsPerSeed.Set(&v)
}
// SetBatchJobsPerSeedNil sets the value for BatchJobsPerSeed to be an explicit nil
func (o *HistoricBatchDto) SetBatchJobsPerSeedNil() {
	o.BatchJobsPerSeed.Set(nil)
}

// UnsetBatchJobsPerSeed ensures that no value is present for BatchJobsPerSeed, not even an explicit nil
func (o *HistoricBatchDto) UnsetBatchJobsPerSeed() {
	o.BatchJobsPerSeed.Unset()
}

// GetInvocationsPerBatchJob returns the InvocationsPerBatchJob field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetInvocationsPerBatchJob() int32 {
	if o == nil || IsNil(o.InvocationsPerBatchJob.Get()) {
		var ret int32
		return ret
	}
	return *o.InvocationsPerBatchJob.Get()
}

// GetInvocationsPerBatchJobOk returns a tuple with the InvocationsPerBatchJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetInvocationsPerBatchJobOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvocationsPerBatchJob.Get(), o.InvocationsPerBatchJob.IsSet()
}

// HasInvocationsPerBatchJob returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasInvocationsPerBatchJob() bool {
	if o != nil && o.InvocationsPerBatchJob.IsSet() {
		return true
	}

	return false
}

// SetInvocationsPerBatchJob gets a reference to the given NullableInt32 and assigns it to the InvocationsPerBatchJob field.
func (o *HistoricBatchDto) SetInvocationsPerBatchJob(v int32) {
	o.InvocationsPerBatchJob.Set(&v)
}
// SetInvocationsPerBatchJobNil sets the value for InvocationsPerBatchJob to be an explicit nil
func (o *HistoricBatchDto) SetInvocationsPerBatchJobNil() {
	o.InvocationsPerBatchJob.Set(nil)
}

// UnsetInvocationsPerBatchJob ensures that no value is present for InvocationsPerBatchJob, not even an explicit nil
func (o *HistoricBatchDto) UnsetInvocationsPerBatchJob() {
	o.InvocationsPerBatchJob.Unset()
}

// GetSeedJobDefinitionId returns the SeedJobDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetSeedJobDefinitionId() string {
	if o == nil || IsNil(o.SeedJobDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.SeedJobDefinitionId.Get()
}

// GetSeedJobDefinitionIdOk returns a tuple with the SeedJobDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetSeedJobDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeedJobDefinitionId.Get(), o.SeedJobDefinitionId.IsSet()
}

// HasSeedJobDefinitionId returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasSeedJobDefinitionId() bool {
	if o != nil && o.SeedJobDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetSeedJobDefinitionId gets a reference to the given NullableString and assigns it to the SeedJobDefinitionId field.
func (o *HistoricBatchDto) SetSeedJobDefinitionId(v string) {
	o.SeedJobDefinitionId.Set(&v)
}
// SetSeedJobDefinitionIdNil sets the value for SeedJobDefinitionId to be an explicit nil
func (o *HistoricBatchDto) SetSeedJobDefinitionIdNil() {
	o.SeedJobDefinitionId.Set(nil)
}

// UnsetSeedJobDefinitionId ensures that no value is present for SeedJobDefinitionId, not even an explicit nil
func (o *HistoricBatchDto) UnsetSeedJobDefinitionId() {
	o.SeedJobDefinitionId.Unset()
}

// GetMonitorJobDefinitionId returns the MonitorJobDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetMonitorJobDefinitionId() string {
	if o == nil || IsNil(o.MonitorJobDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.MonitorJobDefinitionId.Get()
}

// GetMonitorJobDefinitionIdOk returns a tuple with the MonitorJobDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetMonitorJobDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitorJobDefinitionId.Get(), o.MonitorJobDefinitionId.IsSet()
}

// HasMonitorJobDefinitionId returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasMonitorJobDefinitionId() bool {
	if o != nil && o.MonitorJobDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetMonitorJobDefinitionId gets a reference to the given NullableString and assigns it to the MonitorJobDefinitionId field.
func (o *HistoricBatchDto) SetMonitorJobDefinitionId(v string) {
	o.MonitorJobDefinitionId.Set(&v)
}
// SetMonitorJobDefinitionIdNil sets the value for MonitorJobDefinitionId to be an explicit nil
func (o *HistoricBatchDto) SetMonitorJobDefinitionIdNil() {
	o.MonitorJobDefinitionId.Set(nil)
}

// UnsetMonitorJobDefinitionId ensures that no value is present for MonitorJobDefinitionId, not even an explicit nil
func (o *HistoricBatchDto) UnsetMonitorJobDefinitionId() {
	o.MonitorJobDefinitionId.Unset()
}

// GetBatchJobDefinitionId returns the BatchJobDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetBatchJobDefinitionId() string {
	if o == nil || IsNil(o.BatchJobDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.BatchJobDefinitionId.Get()
}

// GetBatchJobDefinitionIdOk returns a tuple with the BatchJobDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetBatchJobDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BatchJobDefinitionId.Get(), o.BatchJobDefinitionId.IsSet()
}

// HasBatchJobDefinitionId returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasBatchJobDefinitionId() bool {
	if o != nil && o.BatchJobDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetBatchJobDefinitionId gets a reference to the given NullableString and assigns it to the BatchJobDefinitionId field.
func (o *HistoricBatchDto) SetBatchJobDefinitionId(v string) {
	o.BatchJobDefinitionId.Set(&v)
}
// SetBatchJobDefinitionIdNil sets the value for BatchJobDefinitionId to be an explicit nil
func (o *HistoricBatchDto) SetBatchJobDefinitionIdNil() {
	o.BatchJobDefinitionId.Set(nil)
}

// UnsetBatchJobDefinitionId ensures that no value is present for BatchJobDefinitionId, not even an explicit nil
func (o *HistoricBatchDto) UnsetBatchJobDefinitionId() {
	o.BatchJobDefinitionId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *HistoricBatchDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *HistoricBatchDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *HistoricBatchDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetCreateUserId returns the CreateUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetCreateUserId() string {
	if o == nil || IsNil(o.CreateUserId.Get()) {
		var ret string
		return ret
	}
	return *o.CreateUserId.Get()
}

// GetCreateUserIdOk returns a tuple with the CreateUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetCreateUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateUserId.Get(), o.CreateUserId.IsSet()
}

// HasCreateUserId returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasCreateUserId() bool {
	if o != nil && o.CreateUserId.IsSet() {
		return true
	}

	return false
}

// SetCreateUserId gets a reference to the given NullableString and assigns it to the CreateUserId field.
func (o *HistoricBatchDto) SetCreateUserId(v string) {
	o.CreateUserId.Set(&v)
}
// SetCreateUserIdNil sets the value for CreateUserId to be an explicit nil
func (o *HistoricBatchDto) SetCreateUserIdNil() {
	o.CreateUserId.Set(nil)
}

// UnsetCreateUserId ensures that no value is present for CreateUserId, not even an explicit nil
func (o *HistoricBatchDto) UnsetCreateUserId() {
	o.CreateUserId.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.
func (o *HistoricBatchDto) SetStartTime(v time.Time) {
	o.StartTime.Set(&v)
}
// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *HistoricBatchDto) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *HistoricBatchDto) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetExecutionStartTime returns the ExecutionStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetExecutionStartTime() time.Time {
	if o == nil || IsNil(o.ExecutionStartTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExecutionStartTime.Get()
}

// GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetExecutionStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionStartTime.Get(), o.ExecutionStartTime.IsSet()
}

// HasExecutionStartTime returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasExecutionStartTime() bool {
	if o != nil && o.ExecutionStartTime.IsSet() {
		return true
	}

	return false
}

// SetExecutionStartTime gets a reference to the given NullableTime and assigns it to the ExecutionStartTime field.
func (o *HistoricBatchDto) SetExecutionStartTime(v time.Time) {
	o.ExecutionStartTime.Set(&v)
}
// SetExecutionStartTimeNil sets the value for ExecutionStartTime to be an explicit nil
func (o *HistoricBatchDto) SetExecutionStartTimeNil() {
	o.ExecutionStartTime.Set(nil)
}

// UnsetExecutionStartTime ensures that no value is present for ExecutionStartTime, not even an explicit nil
func (o *HistoricBatchDto) UnsetExecutionStartTime() {
	o.ExecutionStartTime.Unset()
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// HasEndTime returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasEndTime() bool {
	if o != nil && o.EndTime.IsSet() {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given NullableTime and assigns it to the EndTime field.
func (o *HistoricBatchDto) SetEndTime(v time.Time) {
	o.EndTime.Set(&v)
}
// SetEndTimeNil sets the value for EndTime to be an explicit nil
func (o *HistoricBatchDto) SetEndTimeNil() {
	o.EndTime.Set(nil)
}

// UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
func (o *HistoricBatchDto) UnsetEndTime() {
	o.EndTime.Unset()
}

// GetRemovalTime returns the RemovalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricBatchDto) GetRemovalTime() time.Time {
	if o == nil || IsNil(o.RemovalTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RemovalTime.Get()
}

// GetRemovalTimeOk returns a tuple with the RemovalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricBatchDto) GetRemovalTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovalTime.Get(), o.RemovalTime.IsSet()
}

// HasRemovalTime returns a boolean if a field has been set.
func (o *HistoricBatchDto) HasRemovalTime() bool {
	if o != nil && o.RemovalTime.IsSet() {
		return true
	}

	return false
}

// SetRemovalTime gets a reference to the given NullableTime and assigns it to the RemovalTime field.
func (o *HistoricBatchDto) SetRemovalTime(v time.Time) {
	o.RemovalTime.Set(&v)
}
// SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil
func (o *HistoricBatchDto) SetRemovalTimeNil() {
	o.RemovalTime.Set(nil)
}

// UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
func (o *HistoricBatchDto) UnsetRemovalTime() {
	o.RemovalTime.Unset()
}

func (o HistoricBatchDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricBatchDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.TotalJobs.IsSet() {
		toSerialize["totalJobs"] = o.TotalJobs.Get()
	}
	if o.BatchJobsPerSeed.IsSet() {
		toSerialize["batchJobsPerSeed"] = o.BatchJobsPerSeed.Get()
	}
	if o.InvocationsPerBatchJob.IsSet() {
		toSerialize["invocationsPerBatchJob"] = o.InvocationsPerBatchJob.Get()
	}
	if o.SeedJobDefinitionId.IsSet() {
		toSerialize["seedJobDefinitionId"] = o.SeedJobDefinitionId.Get()
	}
	if o.MonitorJobDefinitionId.IsSet() {
		toSerialize["monitorJobDefinitionId"] = o.MonitorJobDefinitionId.Get()
	}
	if o.BatchJobDefinitionId.IsSet() {
		toSerialize["batchJobDefinitionId"] = o.BatchJobDefinitionId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.CreateUserId.IsSet() {
		toSerialize["createUserId"] = o.CreateUserId.Get()
	}
	if o.StartTime.IsSet() {
		toSerialize["startTime"] = o.StartTime.Get()
	}
	if o.ExecutionStartTime.IsSet() {
		toSerialize["executionStartTime"] = o.ExecutionStartTime.Get()
	}
	if o.EndTime.IsSet() {
		toSerialize["endTime"] = o.EndTime.Get()
	}
	if o.RemovalTime.IsSet() {
		toSerialize["removalTime"] = o.RemovalTime.Get()
	}
	return toSerialize, nil
}

type NullableHistoricBatchDto struct {
	value *HistoricBatchDto
	isSet bool
}

func (v NullableHistoricBatchDto) Get() *HistoricBatchDto {
	return v.value
}

func (v *NullableHistoricBatchDto) Set(val *HistoricBatchDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricBatchDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricBatchDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricBatchDto(val *HistoricBatchDto) *NullableHistoricBatchDto {
	return &NullableHistoricBatchDto{value: val, isSet: true}
}

func (v NullableHistoricBatchDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricBatchDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


