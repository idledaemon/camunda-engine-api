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

// checks if the HistoricVariableInstanceQueryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoricVariableInstanceQueryDto{}

// HistoricVariableInstanceQueryDto A Historic Variable Instance instance query which defines a list of Historic Variable Instance instances
type HistoricVariableInstanceQueryDto struct {
	// Filter by variable name.
	VariableName NullableString `json:"variableName,omitempty"`
	// Restrict to variables with a name like the parameter.
	VariableNameLike NullableString `json:"variableNameLike,omitempty"`
	// Filter by variable value. May be `String`, `Number` or `Boolean`.
	VariableValue map[string]interface{} `json:"variableValue,omitempty"`
	// Match the variable name provided in `variableName` and `variableNameLike` case- insensitively. If set to `true` **variableName** and **variablename** are treated as equal.
	VariableNamesIgnoreCase NullableBool `json:"variableNamesIgnoreCase,omitempty"`
	// Match the variable value provided in `variableValue` case-insensitively. If set to `true` **variableValue** and **variablevalue** are treated as equal.
	VariableValuesIgnoreCase NullableBool `json:"variableValuesIgnoreCase,omitempty"`
	// Only include historic variable instances which belong to one of the passed and comma- separated variable types. A list of all supported variable types can be found [here](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables/#supported-variable-values). **Note:** All non-primitive variables are associated with the type 'serializable'.
	VariableTypeIn []string `json:"variableTypeIn,omitempty"`
	// Include variables that has already been deleted during the execution.
	IncludeDeleted NullableBool `json:"includeDeleted,omitempty"`
	// Filter by the process instance the variable belongs to.
	ProcessInstanceId NullableString `json:"processInstanceId,omitempty"`
	// Only include historic variable instances which belong to one of the passed  process instance ids.
	ProcessInstanceIdIn []string `json:"processInstanceIdIn,omitempty"`
	// Filter by the process definition the variable belongs to.
	ProcessDefinitionId NullableString `json:"processDefinitionId,omitempty"`
	// Filter by a key of the process definition the variable belongs to.
	ProcessDefinitionKey NullableString `json:"processDefinitionKey,omitempty"`
	// Only include historic variable instances which belong to one of the passed and  execution ids.
	ExecutionIdIn []string `json:"executionIdIn,omitempty"`
	// Filter by the case instance the variable belongs to.
	CaseInstanceId NullableString `json:"caseInstanceId,omitempty"`
	// Only include historic variable instances which belong to one of the passed and  case execution ids.
	CaseExecutionIdIn []string `json:"caseExecutionIdIn,omitempty"`
	// Only include historic variable instances which belong to one of the passed and  case activity ids.
	CaseActivityIdIn []string `json:"caseActivityIdIn,omitempty"`
	// Only include historic variable instances which belong to one of the passed and  task ids.
	TaskIdIn []string `json:"taskIdIn,omitempty"`
	// Only include historic variable instances which belong to one of the passed and  activity instance ids.
	ActivityInstanceIdIn []string `json:"activityInstanceIdIn,omitempty"`
	// Only include historic variable instances which belong to one of the passed and comma- separated tenant ids.
	TenantIdIn []string `json:"tenantIdIn,omitempty"`
	// Only include historic variable instances that belong to no tenant. Value may only be `true`, as `false` is the default behavior.
	WithoutTenantId NullableBool `json:"withoutTenantId,omitempty"`
	// Only include historic variable instances which belong to one of the passed  variable names.
	VariableNameIn []string `json:"variableNameIn,omitempty"`
	// An array of criteria to sort the result by. Each element of the array is                      an object that specifies one ordering. The position in the array                      identifies the rank of an ordering, i.e., whether it is primary, secondary,                      etc. Sorting has no effect for `count` endpoints
	Sorting []HistoricVariableInstanceQueryDtoSortingInner `json:"sorting,omitempty"`
}

// NewHistoricVariableInstanceQueryDto instantiates a new HistoricVariableInstanceQueryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricVariableInstanceQueryDto() *HistoricVariableInstanceQueryDto {
	this := HistoricVariableInstanceQueryDto{}
	return &this
}

// NewHistoricVariableInstanceQueryDtoWithDefaults instantiates a new HistoricVariableInstanceQueryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricVariableInstanceQueryDtoWithDefaults() *HistoricVariableInstanceQueryDto {
	this := HistoricVariableInstanceQueryDto{}
	return &this
}

// GetVariableName returns the VariableName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetVariableName() string {
	if o == nil || IsNil(o.VariableName.Get()) {
		var ret string
		return ret
	}
	return *o.VariableName.Get()
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetVariableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariableName.Get(), o.VariableName.IsSet()
}

// HasVariableName returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasVariableName() bool {
	if o != nil && o.VariableName.IsSet() {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given NullableString and assigns it to the VariableName field.
func (o *HistoricVariableInstanceQueryDto) SetVariableName(v string) {
	o.VariableName.Set(&v)
}
// SetVariableNameNil sets the value for VariableName to be an explicit nil
func (o *HistoricVariableInstanceQueryDto) SetVariableNameNil() {
	o.VariableName.Set(nil)
}

// UnsetVariableName ensures that no value is present for VariableName, not even an explicit nil
func (o *HistoricVariableInstanceQueryDto) UnsetVariableName() {
	o.VariableName.Unset()
}

// GetVariableNameLike returns the VariableNameLike field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetVariableNameLike() string {
	if o == nil || IsNil(o.VariableNameLike.Get()) {
		var ret string
		return ret
	}
	return *o.VariableNameLike.Get()
}

// GetVariableNameLikeOk returns a tuple with the VariableNameLike field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetVariableNameLikeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariableNameLike.Get(), o.VariableNameLike.IsSet()
}

// HasVariableNameLike returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasVariableNameLike() bool {
	if o != nil && o.VariableNameLike.IsSet() {
		return true
	}

	return false
}

// SetVariableNameLike gets a reference to the given NullableString and assigns it to the VariableNameLike field.
func (o *HistoricVariableInstanceQueryDto) SetVariableNameLike(v string) {
	o.VariableNameLike.Set(&v)
}
// SetVariableNameLikeNil sets the value for VariableNameLike to be an explicit nil
func (o *HistoricVariableInstanceQueryDto) SetVariableNameLikeNil() {
	o.VariableNameLike.Set(nil)
}

// UnsetVariableNameLike ensures that no value is present for VariableNameLike, not even an explicit nil
func (o *HistoricVariableInstanceQueryDto) UnsetVariableNameLike() {
	o.VariableNameLike.Unset()
}

// GetVariableValue returns the VariableValue field value if set, zero value otherwise.
func (o *HistoricVariableInstanceQueryDto) GetVariableValue() map[string]interface{} {
	if o == nil || IsNil(o.VariableValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.VariableValue
}

// GetVariableValueOk returns a tuple with the VariableValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricVariableInstanceQueryDto) GetVariableValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.VariableValue) {
		return map[string]interface{}{}, false
	}
	return o.VariableValue, true
}

// HasVariableValue returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasVariableValue() bool {
	if o != nil && !IsNil(o.VariableValue) {
		return true
	}

	return false
}

// SetVariableValue gets a reference to the given map[string]interface{} and assigns it to the VariableValue field.
func (o *HistoricVariableInstanceQueryDto) SetVariableValue(v map[string]interface{}) {
	o.VariableValue = v
}

// GetVariableNamesIgnoreCase returns the VariableNamesIgnoreCase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetVariableNamesIgnoreCase() bool {
	if o == nil || IsNil(o.VariableNamesIgnoreCase.Get()) {
		var ret bool
		return ret
	}
	return *o.VariableNamesIgnoreCase.Get()
}

// GetVariableNamesIgnoreCaseOk returns a tuple with the VariableNamesIgnoreCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetVariableNamesIgnoreCaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariableNamesIgnoreCase.Get(), o.VariableNamesIgnoreCase.IsSet()
}

// HasVariableNamesIgnoreCase returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasVariableNamesIgnoreCase() bool {
	if o != nil && o.VariableNamesIgnoreCase.IsSet() {
		return true
	}

	return false
}

// SetVariableNamesIgnoreCase gets a reference to the given NullableBool and assigns it to the VariableNamesIgnoreCase field.
func (o *HistoricVariableInstanceQueryDto) SetVariableNamesIgnoreCase(v bool) {
	o.VariableNamesIgnoreCase.Set(&v)
}
// SetVariableNamesIgnoreCaseNil sets the value for VariableNamesIgnoreCase to be an explicit nil
func (o *HistoricVariableInstanceQueryDto) SetVariableNamesIgnoreCaseNil() {
	o.VariableNamesIgnoreCase.Set(nil)
}

// UnsetVariableNamesIgnoreCase ensures that no value is present for VariableNamesIgnoreCase, not even an explicit nil
func (o *HistoricVariableInstanceQueryDto) UnsetVariableNamesIgnoreCase() {
	o.VariableNamesIgnoreCase.Unset()
}

// GetVariableValuesIgnoreCase returns the VariableValuesIgnoreCase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetVariableValuesIgnoreCase() bool {
	if o == nil || IsNil(o.VariableValuesIgnoreCase.Get()) {
		var ret bool
		return ret
	}
	return *o.VariableValuesIgnoreCase.Get()
}

// GetVariableValuesIgnoreCaseOk returns a tuple with the VariableValuesIgnoreCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetVariableValuesIgnoreCaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariableValuesIgnoreCase.Get(), o.VariableValuesIgnoreCase.IsSet()
}

// HasVariableValuesIgnoreCase returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasVariableValuesIgnoreCase() bool {
	if o != nil && o.VariableValuesIgnoreCase.IsSet() {
		return true
	}

	return false
}

// SetVariableValuesIgnoreCase gets a reference to the given NullableBool and assigns it to the VariableValuesIgnoreCase field.
func (o *HistoricVariableInstanceQueryDto) SetVariableValuesIgnoreCase(v bool) {
	o.VariableValuesIgnoreCase.Set(&v)
}
// SetVariableValuesIgnoreCaseNil sets the value for VariableValuesIgnoreCase to be an explicit nil
func (o *HistoricVariableInstanceQueryDto) SetVariableValuesIgnoreCaseNil() {
	o.VariableValuesIgnoreCase.Set(nil)
}

// UnsetVariableValuesIgnoreCase ensures that no value is present for VariableValuesIgnoreCase, not even an explicit nil
func (o *HistoricVariableInstanceQueryDto) UnsetVariableValuesIgnoreCase() {
	o.VariableValuesIgnoreCase.Unset()
}

// GetVariableTypeIn returns the VariableTypeIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetVariableTypeIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.VariableTypeIn
}

// GetVariableTypeInOk returns a tuple with the VariableTypeIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetVariableTypeInOk() ([]string, bool) {
	if o == nil || IsNil(o.VariableTypeIn) {
		return nil, false
	}
	return o.VariableTypeIn, true
}

// HasVariableTypeIn returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasVariableTypeIn() bool {
	if o != nil && !IsNil(o.VariableTypeIn) {
		return true
	}

	return false
}

// SetVariableTypeIn gets a reference to the given []string and assigns it to the VariableTypeIn field.
func (o *HistoricVariableInstanceQueryDto) SetVariableTypeIn(v []string) {
	o.VariableTypeIn = v
}

// GetIncludeDeleted returns the IncludeDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetIncludeDeleted() bool {
	if o == nil || IsNil(o.IncludeDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeDeleted.Get()
}

// GetIncludeDeletedOk returns a tuple with the IncludeDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetIncludeDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeDeleted.Get(), o.IncludeDeleted.IsSet()
}

// HasIncludeDeleted returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasIncludeDeleted() bool {
	if o != nil && o.IncludeDeleted.IsSet() {
		return true
	}

	return false
}

// SetIncludeDeleted gets a reference to the given NullableBool and assigns it to the IncludeDeleted field.
func (o *HistoricVariableInstanceQueryDto) SetIncludeDeleted(v bool) {
	o.IncludeDeleted.Set(&v)
}
// SetIncludeDeletedNil sets the value for IncludeDeleted to be an explicit nil
func (o *HistoricVariableInstanceQueryDto) SetIncludeDeletedNil() {
	o.IncludeDeleted.Set(nil)
}

// UnsetIncludeDeleted ensures that no value is present for IncludeDeleted, not even an explicit nil
func (o *HistoricVariableInstanceQueryDto) UnsetIncludeDeleted() {
	o.IncludeDeleted.Unset()
}

// GetProcessInstanceId returns the ProcessInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetProcessInstanceId() string {
	if o == nil || IsNil(o.ProcessInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessInstanceId.Get()
}

// GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetProcessInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessInstanceId.Get(), o.ProcessInstanceId.IsSet()
}

// HasProcessInstanceId returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasProcessInstanceId() bool {
	if o != nil && o.ProcessInstanceId.IsSet() {
		return true
	}

	return false
}

// SetProcessInstanceId gets a reference to the given NullableString and assigns it to the ProcessInstanceId field.
func (o *HistoricVariableInstanceQueryDto) SetProcessInstanceId(v string) {
	o.ProcessInstanceId.Set(&v)
}
// SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil
func (o *HistoricVariableInstanceQueryDto) SetProcessInstanceIdNil() {
	o.ProcessInstanceId.Set(nil)
}

// UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
func (o *HistoricVariableInstanceQueryDto) UnsetProcessInstanceId() {
	o.ProcessInstanceId.Unset()
}

// GetProcessInstanceIdIn returns the ProcessInstanceIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetProcessInstanceIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProcessInstanceIdIn
}

// GetProcessInstanceIdInOk returns a tuple with the ProcessInstanceIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetProcessInstanceIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessInstanceIdIn) {
		return nil, false
	}
	return o.ProcessInstanceIdIn, true
}

// HasProcessInstanceIdIn returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasProcessInstanceIdIn() bool {
	if o != nil && !IsNil(o.ProcessInstanceIdIn) {
		return true
	}

	return false
}

// SetProcessInstanceIdIn gets a reference to the given []string and assigns it to the ProcessInstanceIdIn field.
func (o *HistoricVariableInstanceQueryDto) SetProcessInstanceIdIn(v []string) {
	o.ProcessInstanceIdIn = v
}

// GetProcessDefinitionId returns the ProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetProcessDefinitionId() string {
	if o == nil || IsNil(o.ProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionId.Get()
}

// GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionId.Get(), o.ProcessDefinitionId.IsSet()
}

// HasProcessDefinitionId returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasProcessDefinitionId() bool {
	if o != nil && o.ProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionId gets a reference to the given NullableString and assigns it to the ProcessDefinitionId field.
func (o *HistoricVariableInstanceQueryDto) SetProcessDefinitionId(v string) {
	o.ProcessDefinitionId.Set(&v)
}
// SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil
func (o *HistoricVariableInstanceQueryDto) SetProcessDefinitionIdNil() {
	o.ProcessDefinitionId.Set(nil)
}

// UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
func (o *HistoricVariableInstanceQueryDto) UnsetProcessDefinitionId() {
	o.ProcessDefinitionId.Unset()
}

// GetProcessDefinitionKey returns the ProcessDefinitionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetProcessDefinitionKey() string {
	if o == nil || IsNil(o.ProcessDefinitionKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionKey.Get()
}

// GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetProcessDefinitionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionKey.Get(), o.ProcessDefinitionKey.IsSet()
}

// HasProcessDefinitionKey returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasProcessDefinitionKey() bool {
	if o != nil && o.ProcessDefinitionKey.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionKey gets a reference to the given NullableString and assigns it to the ProcessDefinitionKey field.
func (o *HistoricVariableInstanceQueryDto) SetProcessDefinitionKey(v string) {
	o.ProcessDefinitionKey.Set(&v)
}
// SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil
func (o *HistoricVariableInstanceQueryDto) SetProcessDefinitionKeyNil() {
	o.ProcessDefinitionKey.Set(nil)
}

// UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
func (o *HistoricVariableInstanceQueryDto) UnsetProcessDefinitionKey() {
	o.ProcessDefinitionKey.Unset()
}

// GetExecutionIdIn returns the ExecutionIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetExecutionIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExecutionIdIn
}

// GetExecutionIdInOk returns a tuple with the ExecutionIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetExecutionIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.ExecutionIdIn) {
		return nil, false
	}
	return o.ExecutionIdIn, true
}

// HasExecutionIdIn returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasExecutionIdIn() bool {
	if o != nil && !IsNil(o.ExecutionIdIn) {
		return true
	}

	return false
}

// SetExecutionIdIn gets a reference to the given []string and assigns it to the ExecutionIdIn field.
func (o *HistoricVariableInstanceQueryDto) SetExecutionIdIn(v []string) {
	o.ExecutionIdIn = v
}

// GetCaseInstanceId returns the CaseInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetCaseInstanceId() string {
	if o == nil || IsNil(o.CaseInstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.CaseInstanceId.Get()
}

// GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetCaseInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaseInstanceId.Get(), o.CaseInstanceId.IsSet()
}

// HasCaseInstanceId returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasCaseInstanceId() bool {
	if o != nil && o.CaseInstanceId.IsSet() {
		return true
	}

	return false
}

// SetCaseInstanceId gets a reference to the given NullableString and assigns it to the CaseInstanceId field.
func (o *HistoricVariableInstanceQueryDto) SetCaseInstanceId(v string) {
	o.CaseInstanceId.Set(&v)
}
// SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil
func (o *HistoricVariableInstanceQueryDto) SetCaseInstanceIdNil() {
	o.CaseInstanceId.Set(nil)
}

// UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
func (o *HistoricVariableInstanceQueryDto) UnsetCaseInstanceId() {
	o.CaseInstanceId.Unset()
}

// GetCaseExecutionIdIn returns the CaseExecutionIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetCaseExecutionIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CaseExecutionIdIn
}

// GetCaseExecutionIdInOk returns a tuple with the CaseExecutionIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetCaseExecutionIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.CaseExecutionIdIn) {
		return nil, false
	}
	return o.CaseExecutionIdIn, true
}

// HasCaseExecutionIdIn returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasCaseExecutionIdIn() bool {
	if o != nil && !IsNil(o.CaseExecutionIdIn) {
		return true
	}

	return false
}

// SetCaseExecutionIdIn gets a reference to the given []string and assigns it to the CaseExecutionIdIn field.
func (o *HistoricVariableInstanceQueryDto) SetCaseExecutionIdIn(v []string) {
	o.CaseExecutionIdIn = v
}

// GetCaseActivityIdIn returns the CaseActivityIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetCaseActivityIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CaseActivityIdIn
}

// GetCaseActivityIdInOk returns a tuple with the CaseActivityIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetCaseActivityIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.CaseActivityIdIn) {
		return nil, false
	}
	return o.CaseActivityIdIn, true
}

// HasCaseActivityIdIn returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasCaseActivityIdIn() bool {
	if o != nil && !IsNil(o.CaseActivityIdIn) {
		return true
	}

	return false
}

// SetCaseActivityIdIn gets a reference to the given []string and assigns it to the CaseActivityIdIn field.
func (o *HistoricVariableInstanceQueryDto) SetCaseActivityIdIn(v []string) {
	o.CaseActivityIdIn = v
}

// GetTaskIdIn returns the TaskIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetTaskIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TaskIdIn
}

// GetTaskIdInOk returns a tuple with the TaskIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetTaskIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.TaskIdIn) {
		return nil, false
	}
	return o.TaskIdIn, true
}

// HasTaskIdIn returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasTaskIdIn() bool {
	if o != nil && !IsNil(o.TaskIdIn) {
		return true
	}

	return false
}

// SetTaskIdIn gets a reference to the given []string and assigns it to the TaskIdIn field.
func (o *HistoricVariableInstanceQueryDto) SetTaskIdIn(v []string) {
	o.TaskIdIn = v
}

// GetActivityInstanceIdIn returns the ActivityInstanceIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetActivityInstanceIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ActivityInstanceIdIn
}

// GetActivityInstanceIdInOk returns a tuple with the ActivityInstanceIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetActivityInstanceIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.ActivityInstanceIdIn) {
		return nil, false
	}
	return o.ActivityInstanceIdIn, true
}

// HasActivityInstanceIdIn returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasActivityInstanceIdIn() bool {
	if o != nil && !IsNil(o.ActivityInstanceIdIn) {
		return true
	}

	return false
}

// SetActivityInstanceIdIn gets a reference to the given []string and assigns it to the ActivityInstanceIdIn field.
func (o *HistoricVariableInstanceQueryDto) SetActivityInstanceIdIn(v []string) {
	o.ActivityInstanceIdIn = v
}

// GetTenantIdIn returns the TenantIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetTenantIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TenantIdIn
}

// GetTenantIdInOk returns a tuple with the TenantIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetTenantIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.TenantIdIn) {
		return nil, false
	}
	return o.TenantIdIn, true
}

// HasTenantIdIn returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasTenantIdIn() bool {
	if o != nil && !IsNil(o.TenantIdIn) {
		return true
	}

	return false
}

// SetTenantIdIn gets a reference to the given []string and assigns it to the TenantIdIn field.
func (o *HistoricVariableInstanceQueryDto) SetTenantIdIn(v []string) {
	o.TenantIdIn = v
}

// GetWithoutTenantId returns the WithoutTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetWithoutTenantId() bool {
	if o == nil || IsNil(o.WithoutTenantId.Get()) {
		var ret bool
		return ret
	}
	return *o.WithoutTenantId.Get()
}

// GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetWithoutTenantIdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithoutTenantId.Get(), o.WithoutTenantId.IsSet()
}

// HasWithoutTenantId returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasWithoutTenantId() bool {
	if o != nil && o.WithoutTenantId.IsSet() {
		return true
	}

	return false
}

// SetWithoutTenantId gets a reference to the given NullableBool and assigns it to the WithoutTenantId field.
func (o *HistoricVariableInstanceQueryDto) SetWithoutTenantId(v bool) {
	o.WithoutTenantId.Set(&v)
}
// SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil
func (o *HistoricVariableInstanceQueryDto) SetWithoutTenantIdNil() {
	o.WithoutTenantId.Set(nil)
}

// UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
func (o *HistoricVariableInstanceQueryDto) UnsetWithoutTenantId() {
	o.WithoutTenantId.Unset()
}

// GetVariableNameIn returns the VariableNameIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetVariableNameIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.VariableNameIn
}

// GetVariableNameInOk returns a tuple with the VariableNameIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetVariableNameInOk() ([]string, bool) {
	if o == nil || IsNil(o.VariableNameIn) {
		return nil, false
	}
	return o.VariableNameIn, true
}

// HasVariableNameIn returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasVariableNameIn() bool {
	if o != nil && !IsNil(o.VariableNameIn) {
		return true
	}

	return false
}

// SetVariableNameIn gets a reference to the given []string and assigns it to the VariableNameIn field.
func (o *HistoricVariableInstanceQueryDto) SetVariableNameIn(v []string) {
	o.VariableNameIn = v
}

// GetSorting returns the Sorting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricVariableInstanceQueryDto) GetSorting() []HistoricVariableInstanceQueryDtoSortingInner {
	if o == nil {
		var ret []HistoricVariableInstanceQueryDtoSortingInner
		return ret
	}
	return o.Sorting
}

// GetSortingOk returns a tuple with the Sorting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricVariableInstanceQueryDto) GetSortingOk() ([]HistoricVariableInstanceQueryDtoSortingInner, bool) {
	if o == nil || IsNil(o.Sorting) {
		return nil, false
	}
	return o.Sorting, true
}

// HasSorting returns a boolean if a field has been set.
func (o *HistoricVariableInstanceQueryDto) HasSorting() bool {
	if o != nil && !IsNil(o.Sorting) {
		return true
	}

	return false
}

// SetSorting gets a reference to the given []HistoricVariableInstanceQueryDtoSortingInner and assigns it to the Sorting field.
func (o *HistoricVariableInstanceQueryDto) SetSorting(v []HistoricVariableInstanceQueryDtoSortingInner) {
	o.Sorting = v
}

func (o HistoricVariableInstanceQueryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricVariableInstanceQueryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.VariableName.IsSet() {
		toSerialize["variableName"] = o.VariableName.Get()
	}
	if o.VariableNameLike.IsSet() {
		toSerialize["variableNameLike"] = o.VariableNameLike.Get()
	}
	if !IsNil(o.VariableValue) {
		toSerialize["variableValue"] = o.VariableValue
	}
	if o.VariableNamesIgnoreCase.IsSet() {
		toSerialize["variableNamesIgnoreCase"] = o.VariableNamesIgnoreCase.Get()
	}
	if o.VariableValuesIgnoreCase.IsSet() {
		toSerialize["variableValuesIgnoreCase"] = o.VariableValuesIgnoreCase.Get()
	}
	if o.VariableTypeIn != nil {
		toSerialize["variableTypeIn"] = o.VariableTypeIn
	}
	if o.IncludeDeleted.IsSet() {
		toSerialize["includeDeleted"] = o.IncludeDeleted.Get()
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
	if o.ProcessDefinitionKey.IsSet() {
		toSerialize["processDefinitionKey"] = o.ProcessDefinitionKey.Get()
	}
	if o.ExecutionIdIn != nil {
		toSerialize["executionIdIn"] = o.ExecutionIdIn
	}
	if o.CaseInstanceId.IsSet() {
		toSerialize["caseInstanceId"] = o.CaseInstanceId.Get()
	}
	if o.CaseExecutionIdIn != nil {
		toSerialize["caseExecutionIdIn"] = o.CaseExecutionIdIn
	}
	if o.CaseActivityIdIn != nil {
		toSerialize["caseActivityIdIn"] = o.CaseActivityIdIn
	}
	if o.TaskIdIn != nil {
		toSerialize["taskIdIn"] = o.TaskIdIn
	}
	if o.ActivityInstanceIdIn != nil {
		toSerialize["activityInstanceIdIn"] = o.ActivityInstanceIdIn
	}
	if o.TenantIdIn != nil {
		toSerialize["tenantIdIn"] = o.TenantIdIn
	}
	if o.WithoutTenantId.IsSet() {
		toSerialize["withoutTenantId"] = o.WithoutTenantId.Get()
	}
	if o.VariableNameIn != nil {
		toSerialize["variableNameIn"] = o.VariableNameIn
	}
	if o.Sorting != nil {
		toSerialize["sorting"] = o.Sorting
	}
	return toSerialize, nil
}

type NullableHistoricVariableInstanceQueryDto struct {
	value *HistoricVariableInstanceQueryDto
	isSet bool
}

func (v NullableHistoricVariableInstanceQueryDto) Get() *HistoricVariableInstanceQueryDto {
	return v.value
}

func (v *NullableHistoricVariableInstanceQueryDto) Set(val *HistoricVariableInstanceQueryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricVariableInstanceQueryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricVariableInstanceQueryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricVariableInstanceQueryDto(val *HistoricVariableInstanceQueryDto) *NullableHistoricVariableInstanceQueryDto {
	return &NullableHistoricVariableInstanceQueryDto{value: val, isSet: true}
}

func (v NullableHistoricVariableInstanceQueryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricVariableInstanceQueryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


