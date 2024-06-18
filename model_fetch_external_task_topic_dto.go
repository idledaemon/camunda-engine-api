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

// checks if the FetchExternalTaskTopicDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchExternalTaskTopicDto{}

// FetchExternalTaskTopicDto struct for FetchExternalTaskTopicDto
type FetchExternalTaskTopicDto struct {
	// **Mandatory.** The topic's name.
	TopicName string `json:"topicName"`
	// **Mandatory.** The duration to lock the external tasks for in milliseconds.
	LockDuration NullableInt64 `json:"lockDuration"`
	// A JSON array of `String` values that represent variable names. For each result task belonging to this topic, the given variables are returned as well if they are accessible from the external task's execution. If not provided - all variables will be fetched.
	Variables []string `json:"variables,omitempty"`
	// If `true` only local variables will be fetched.
	LocalVariables NullableBool `json:"localVariables,omitempty"`
	// A `String` value which enables the filtering of tasks based on process instance business key.
	BusinessKey NullableString `json:"businessKey,omitempty"`
	// Filter tasks based on process definition id.
	ProcessDefinitionId NullableString `json:"processDefinitionId,omitempty"`
	// Filter tasks based on process definition ids.
	ProcessDefinitionIdIn []string `json:"processDefinitionIdIn,omitempty"`
	// Filter tasks based on process definition key.
	ProcessDefinitionKey NullableString `json:"processDefinitionKey,omitempty"`
	// Filter tasks based on process definition keys.
	ProcessDefinitionKeyIn []string `json:"processDefinitionKeyIn,omitempty"`
	// Filter tasks based on process definition version tag.
	ProcessDefinitionVersionTag NullableString `json:"processDefinitionVersionTag,omitempty"`
	// Filter tasks without tenant id.
	WithoutTenantId NullableBool `json:"withoutTenantId,omitempty"`
	// Filter tasks based on tenant ids.
	TenantIdIn []string `json:"tenantIdIn,omitempty"`
	// A `JSON` object used for filtering tasks based on process instance variable values. A property name of the object represents a process variable name, while the property value represents the process variable value to filter tasks by.
	ProcessVariables map[string]interface{} `json:"processVariables,omitempty"`
	// Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default `false`).  If set to `true`, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson's](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API's classpath.  If set to `false`, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML.
	DeserializeValues NullableBool `json:"deserializeValues,omitempty"`
	// Determines whether custom extension properties defined in the BPMN activity of the external task (e.g. via the Extensions tab in the Camunda modeler) should be included in the response. Default: false
	IncludeExtensionProperties NullableBool `json:"includeExtensionProperties,omitempty"`
}

type _FetchExternalTaskTopicDto FetchExternalTaskTopicDto

// NewFetchExternalTaskTopicDto instantiates a new FetchExternalTaskTopicDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchExternalTaskTopicDto(topicName string, lockDuration NullableInt64) *FetchExternalTaskTopicDto {
	this := FetchExternalTaskTopicDto{}
	this.TopicName = topicName
	this.LockDuration = lockDuration
	var localVariables bool = false
	this.LocalVariables = *NewNullableBool(&localVariables)
	var withoutTenantId bool = false
	this.WithoutTenantId = *NewNullableBool(&withoutTenantId)
	var deserializeValues bool = false
	this.DeserializeValues = *NewNullableBool(&deserializeValues)
	var includeExtensionProperties bool = false
	this.IncludeExtensionProperties = *NewNullableBool(&includeExtensionProperties)
	return &this
}

// NewFetchExternalTaskTopicDtoWithDefaults instantiates a new FetchExternalTaskTopicDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchExternalTaskTopicDtoWithDefaults() *FetchExternalTaskTopicDto {
	this := FetchExternalTaskTopicDto{}
	var localVariables bool = false
	this.LocalVariables = *NewNullableBool(&localVariables)
	var withoutTenantId bool = false
	this.WithoutTenantId = *NewNullableBool(&withoutTenantId)
	var deserializeValues bool = false
	this.DeserializeValues = *NewNullableBool(&deserializeValues)
	var includeExtensionProperties bool = false
	this.IncludeExtensionProperties = *NewNullableBool(&includeExtensionProperties)
	return &this
}

// GetTopicName returns the TopicName field value
func (o *FetchExternalTaskTopicDto) GetTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value
// and a boolean to check if the value has been set.
func (o *FetchExternalTaskTopicDto) GetTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopicName, true
}

// SetTopicName sets field value
func (o *FetchExternalTaskTopicDto) SetTopicName(v string) {
	o.TopicName = v
}

// GetLockDuration returns the LockDuration field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *FetchExternalTaskTopicDto) GetLockDuration() int64 {
	if o == nil || o.LockDuration.Get() == nil {
		var ret int64
		return ret
	}

	return *o.LockDuration.Get()
}

// GetLockDurationOk returns a tuple with the LockDuration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetLockDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockDuration.Get(), o.LockDuration.IsSet()
}

// SetLockDuration sets field value
func (o *FetchExternalTaskTopicDto) SetLockDuration(v int64) {
	o.LockDuration.Set(&v)
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetVariables() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetVariablesOk() ([]string, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []string and assigns it to the Variables field.
func (o *FetchExternalTaskTopicDto) SetVariables(v []string) {
	o.Variables = v
}

// GetLocalVariables returns the LocalVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetLocalVariables() bool {
	if o == nil || IsNil(o.LocalVariables.Get()) {
		var ret bool
		return ret
	}
	return *o.LocalVariables.Get()
}

// GetLocalVariablesOk returns a tuple with the LocalVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetLocalVariablesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalVariables.Get(), o.LocalVariables.IsSet()
}

// HasLocalVariables returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasLocalVariables() bool {
	if o != nil && o.LocalVariables.IsSet() {
		return true
	}

	return false
}

// SetLocalVariables gets a reference to the given NullableBool and assigns it to the LocalVariables field.
func (o *FetchExternalTaskTopicDto) SetLocalVariables(v bool) {
	o.LocalVariables.Set(&v)
}
// SetLocalVariablesNil sets the value for LocalVariables to be an explicit nil
func (o *FetchExternalTaskTopicDto) SetLocalVariablesNil() {
	o.LocalVariables.Set(nil)
}

// UnsetLocalVariables ensures that no value is present for LocalVariables, not even an explicit nil
func (o *FetchExternalTaskTopicDto) UnsetLocalVariables() {
	o.LocalVariables.Unset()
}

// GetBusinessKey returns the BusinessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetBusinessKey() string {
	if o == nil || IsNil(o.BusinessKey.Get()) {
		var ret string
		return ret
	}
	return *o.BusinessKey.Get()
}

// GetBusinessKeyOk returns a tuple with the BusinessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetBusinessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessKey.Get(), o.BusinessKey.IsSet()
}

// HasBusinessKey returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasBusinessKey() bool {
	if o != nil && o.BusinessKey.IsSet() {
		return true
	}

	return false
}

// SetBusinessKey gets a reference to the given NullableString and assigns it to the BusinessKey field.
func (o *FetchExternalTaskTopicDto) SetBusinessKey(v string) {
	o.BusinessKey.Set(&v)
}
// SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil
func (o *FetchExternalTaskTopicDto) SetBusinessKeyNil() {
	o.BusinessKey.Set(nil)
}

// UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
func (o *FetchExternalTaskTopicDto) UnsetBusinessKey() {
	o.BusinessKey.Unset()
}

// GetProcessDefinitionId returns the ProcessDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetProcessDefinitionId() string {
	if o == nil || IsNil(o.ProcessDefinitionId.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionId.Get()
}

// GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetProcessDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionId.Get(), o.ProcessDefinitionId.IsSet()
}

// HasProcessDefinitionId returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasProcessDefinitionId() bool {
	if o != nil && o.ProcessDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionId gets a reference to the given NullableString and assigns it to the ProcessDefinitionId field.
func (o *FetchExternalTaskTopicDto) SetProcessDefinitionId(v string) {
	o.ProcessDefinitionId.Set(&v)
}
// SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil
func (o *FetchExternalTaskTopicDto) SetProcessDefinitionIdNil() {
	o.ProcessDefinitionId.Set(nil)
}

// UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
func (o *FetchExternalTaskTopicDto) UnsetProcessDefinitionId() {
	o.ProcessDefinitionId.Unset()
}

// GetProcessDefinitionIdIn returns the ProcessDefinitionIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetProcessDefinitionIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProcessDefinitionIdIn
}

// GetProcessDefinitionIdInOk returns a tuple with the ProcessDefinitionIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetProcessDefinitionIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessDefinitionIdIn) {
		return nil, false
	}
	return o.ProcessDefinitionIdIn, true
}

// HasProcessDefinitionIdIn returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasProcessDefinitionIdIn() bool {
	if o != nil && !IsNil(o.ProcessDefinitionIdIn) {
		return true
	}

	return false
}

// SetProcessDefinitionIdIn gets a reference to the given []string and assigns it to the ProcessDefinitionIdIn field.
func (o *FetchExternalTaskTopicDto) SetProcessDefinitionIdIn(v []string) {
	o.ProcessDefinitionIdIn = v
}

// GetProcessDefinitionKey returns the ProcessDefinitionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetProcessDefinitionKey() string {
	if o == nil || IsNil(o.ProcessDefinitionKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionKey.Get()
}

// GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetProcessDefinitionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionKey.Get(), o.ProcessDefinitionKey.IsSet()
}

// HasProcessDefinitionKey returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasProcessDefinitionKey() bool {
	if o != nil && o.ProcessDefinitionKey.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionKey gets a reference to the given NullableString and assigns it to the ProcessDefinitionKey field.
func (o *FetchExternalTaskTopicDto) SetProcessDefinitionKey(v string) {
	o.ProcessDefinitionKey.Set(&v)
}
// SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil
func (o *FetchExternalTaskTopicDto) SetProcessDefinitionKeyNil() {
	o.ProcessDefinitionKey.Set(nil)
}

// UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
func (o *FetchExternalTaskTopicDto) UnsetProcessDefinitionKey() {
	o.ProcessDefinitionKey.Unset()
}

// GetProcessDefinitionKeyIn returns the ProcessDefinitionKeyIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetProcessDefinitionKeyIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProcessDefinitionKeyIn
}

// GetProcessDefinitionKeyInOk returns a tuple with the ProcessDefinitionKeyIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetProcessDefinitionKeyInOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessDefinitionKeyIn) {
		return nil, false
	}
	return o.ProcessDefinitionKeyIn, true
}

// HasProcessDefinitionKeyIn returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasProcessDefinitionKeyIn() bool {
	if o != nil && !IsNil(o.ProcessDefinitionKeyIn) {
		return true
	}

	return false
}

// SetProcessDefinitionKeyIn gets a reference to the given []string and assigns it to the ProcessDefinitionKeyIn field.
func (o *FetchExternalTaskTopicDto) SetProcessDefinitionKeyIn(v []string) {
	o.ProcessDefinitionKeyIn = v
}

// GetProcessDefinitionVersionTag returns the ProcessDefinitionVersionTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetProcessDefinitionVersionTag() string {
	if o == nil || IsNil(o.ProcessDefinitionVersionTag.Get()) {
		var ret string
		return ret
	}
	return *o.ProcessDefinitionVersionTag.Get()
}

// GetProcessDefinitionVersionTagOk returns a tuple with the ProcessDefinitionVersionTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetProcessDefinitionVersionTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessDefinitionVersionTag.Get(), o.ProcessDefinitionVersionTag.IsSet()
}

// HasProcessDefinitionVersionTag returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasProcessDefinitionVersionTag() bool {
	if o != nil && o.ProcessDefinitionVersionTag.IsSet() {
		return true
	}

	return false
}

// SetProcessDefinitionVersionTag gets a reference to the given NullableString and assigns it to the ProcessDefinitionVersionTag field.
func (o *FetchExternalTaskTopicDto) SetProcessDefinitionVersionTag(v string) {
	o.ProcessDefinitionVersionTag.Set(&v)
}
// SetProcessDefinitionVersionTagNil sets the value for ProcessDefinitionVersionTag to be an explicit nil
func (o *FetchExternalTaskTopicDto) SetProcessDefinitionVersionTagNil() {
	o.ProcessDefinitionVersionTag.Set(nil)
}

// UnsetProcessDefinitionVersionTag ensures that no value is present for ProcessDefinitionVersionTag, not even an explicit nil
func (o *FetchExternalTaskTopicDto) UnsetProcessDefinitionVersionTag() {
	o.ProcessDefinitionVersionTag.Unset()
}

// GetWithoutTenantId returns the WithoutTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetWithoutTenantId() bool {
	if o == nil || IsNil(o.WithoutTenantId.Get()) {
		var ret bool
		return ret
	}
	return *o.WithoutTenantId.Get()
}

// GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetWithoutTenantIdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithoutTenantId.Get(), o.WithoutTenantId.IsSet()
}

// HasWithoutTenantId returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasWithoutTenantId() bool {
	if o != nil && o.WithoutTenantId.IsSet() {
		return true
	}

	return false
}

// SetWithoutTenantId gets a reference to the given NullableBool and assigns it to the WithoutTenantId field.
func (o *FetchExternalTaskTopicDto) SetWithoutTenantId(v bool) {
	o.WithoutTenantId.Set(&v)
}
// SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil
func (o *FetchExternalTaskTopicDto) SetWithoutTenantIdNil() {
	o.WithoutTenantId.Set(nil)
}

// UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
func (o *FetchExternalTaskTopicDto) UnsetWithoutTenantId() {
	o.WithoutTenantId.Unset()
}

// GetTenantIdIn returns the TenantIdIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetTenantIdIn() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TenantIdIn
}

// GetTenantIdInOk returns a tuple with the TenantIdIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetTenantIdInOk() ([]string, bool) {
	if o == nil || IsNil(o.TenantIdIn) {
		return nil, false
	}
	return o.TenantIdIn, true
}

// HasTenantIdIn returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasTenantIdIn() bool {
	if o != nil && !IsNil(o.TenantIdIn) {
		return true
	}

	return false
}

// SetTenantIdIn gets a reference to the given []string and assigns it to the TenantIdIn field.
func (o *FetchExternalTaskTopicDto) SetTenantIdIn(v []string) {
	o.TenantIdIn = v
}

// GetProcessVariables returns the ProcessVariables field value if set, zero value otherwise.
func (o *FetchExternalTaskTopicDto) GetProcessVariables() map[string]interface{} {
	if o == nil || IsNil(o.ProcessVariables) {
		var ret map[string]interface{}
		return ret
	}
	return o.ProcessVariables
}

// GetProcessVariablesOk returns a tuple with the ProcessVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchExternalTaskTopicDto) GetProcessVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ProcessVariables) {
		return map[string]interface{}{}, false
	}
	return o.ProcessVariables, true
}

// HasProcessVariables returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasProcessVariables() bool {
	if o != nil && !IsNil(o.ProcessVariables) {
		return true
	}

	return false
}

// SetProcessVariables gets a reference to the given map[string]interface{} and assigns it to the ProcessVariables field.
func (o *FetchExternalTaskTopicDto) SetProcessVariables(v map[string]interface{}) {
	o.ProcessVariables = v
}

// GetDeserializeValues returns the DeserializeValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetDeserializeValues() bool {
	if o == nil || IsNil(o.DeserializeValues.Get()) {
		var ret bool
		return ret
	}
	return *o.DeserializeValues.Get()
}

// GetDeserializeValuesOk returns a tuple with the DeserializeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetDeserializeValuesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeserializeValues.Get(), o.DeserializeValues.IsSet()
}

// HasDeserializeValues returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasDeserializeValues() bool {
	if o != nil && o.DeserializeValues.IsSet() {
		return true
	}

	return false
}

// SetDeserializeValues gets a reference to the given NullableBool and assigns it to the DeserializeValues field.
func (o *FetchExternalTaskTopicDto) SetDeserializeValues(v bool) {
	o.DeserializeValues.Set(&v)
}
// SetDeserializeValuesNil sets the value for DeserializeValues to be an explicit nil
func (o *FetchExternalTaskTopicDto) SetDeserializeValuesNil() {
	o.DeserializeValues.Set(nil)
}

// UnsetDeserializeValues ensures that no value is present for DeserializeValues, not even an explicit nil
func (o *FetchExternalTaskTopicDto) UnsetDeserializeValues() {
	o.DeserializeValues.Unset()
}

// GetIncludeExtensionProperties returns the IncludeExtensionProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchExternalTaskTopicDto) GetIncludeExtensionProperties() bool {
	if o == nil || IsNil(o.IncludeExtensionProperties.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeExtensionProperties.Get()
}

// GetIncludeExtensionPropertiesOk returns a tuple with the IncludeExtensionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchExternalTaskTopicDto) GetIncludeExtensionPropertiesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeExtensionProperties.Get(), o.IncludeExtensionProperties.IsSet()
}

// HasIncludeExtensionProperties returns a boolean if a field has been set.
func (o *FetchExternalTaskTopicDto) HasIncludeExtensionProperties() bool {
	if o != nil && o.IncludeExtensionProperties.IsSet() {
		return true
	}

	return false
}

// SetIncludeExtensionProperties gets a reference to the given NullableBool and assigns it to the IncludeExtensionProperties field.
func (o *FetchExternalTaskTopicDto) SetIncludeExtensionProperties(v bool) {
	o.IncludeExtensionProperties.Set(&v)
}
// SetIncludeExtensionPropertiesNil sets the value for IncludeExtensionProperties to be an explicit nil
func (o *FetchExternalTaskTopicDto) SetIncludeExtensionPropertiesNil() {
	o.IncludeExtensionProperties.Set(nil)
}

// UnsetIncludeExtensionProperties ensures that no value is present for IncludeExtensionProperties, not even an explicit nil
func (o *FetchExternalTaskTopicDto) UnsetIncludeExtensionProperties() {
	o.IncludeExtensionProperties.Unset()
}

func (o FetchExternalTaskTopicDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchExternalTaskTopicDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["topicName"] = o.TopicName
	toSerialize["lockDuration"] = o.LockDuration.Get()
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	if o.LocalVariables.IsSet() {
		toSerialize["localVariables"] = o.LocalVariables.Get()
	}
	if o.BusinessKey.IsSet() {
		toSerialize["businessKey"] = o.BusinessKey.Get()
	}
	if o.ProcessDefinitionId.IsSet() {
		toSerialize["processDefinitionId"] = o.ProcessDefinitionId.Get()
	}
	if o.ProcessDefinitionIdIn != nil {
		toSerialize["processDefinitionIdIn"] = o.ProcessDefinitionIdIn
	}
	if o.ProcessDefinitionKey.IsSet() {
		toSerialize["processDefinitionKey"] = o.ProcessDefinitionKey.Get()
	}
	if o.ProcessDefinitionKeyIn != nil {
		toSerialize["processDefinitionKeyIn"] = o.ProcessDefinitionKeyIn
	}
	if o.ProcessDefinitionVersionTag.IsSet() {
		toSerialize["processDefinitionVersionTag"] = o.ProcessDefinitionVersionTag.Get()
	}
	if o.WithoutTenantId.IsSet() {
		toSerialize["withoutTenantId"] = o.WithoutTenantId.Get()
	}
	if o.TenantIdIn != nil {
		toSerialize["tenantIdIn"] = o.TenantIdIn
	}
	if !IsNil(o.ProcessVariables) {
		toSerialize["processVariables"] = o.ProcessVariables
	}
	if o.DeserializeValues.IsSet() {
		toSerialize["deserializeValues"] = o.DeserializeValues.Get()
	}
	if o.IncludeExtensionProperties.IsSet() {
		toSerialize["includeExtensionProperties"] = o.IncludeExtensionProperties.Get()
	}
	return toSerialize, nil
}

func (o *FetchExternalTaskTopicDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"topicName",
		"lockDuration",
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

	varFetchExternalTaskTopicDto := _FetchExternalTaskTopicDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFetchExternalTaskTopicDto)

	if err != nil {
		return err
	}

	*o = FetchExternalTaskTopicDto(varFetchExternalTaskTopicDto)

	return err
}

type NullableFetchExternalTaskTopicDto struct {
	value *FetchExternalTaskTopicDto
	isSet bool
}

func (v NullableFetchExternalTaskTopicDto) Get() *FetchExternalTaskTopicDto {
	return v.value
}

func (v *NullableFetchExternalTaskTopicDto) Set(val *FetchExternalTaskTopicDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchExternalTaskTopicDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchExternalTaskTopicDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchExternalTaskTopicDto(val *FetchExternalTaskTopicDto) *NullableFetchExternalTaskTopicDto {
	return &NullableFetchExternalTaskTopicDto{value: val, isSet: true}
}

func (v NullableFetchExternalTaskTopicDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchExternalTaskTopicDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

