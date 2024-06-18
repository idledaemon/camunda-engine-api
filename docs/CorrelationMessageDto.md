# CorrelationMessageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageName** | Pointer to **NullableString** | The name of the message to deliver. | [optional] 
**BusinessKey** | Pointer to **NullableString** | Used for correlation of process instances that wait for incoming messages. Will only correlate to executions that belong to a process instance with the provided business key. | [optional] 
**TenantId** | Pointer to **NullableString** | Used to correlate the message for a tenant with the given id. Will only correlate to executions and process definitions which belong to the tenant. Must not be supplied in conjunction with a &#x60;withoutTenantId&#x60;. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | A Boolean value that indicates whether the message should only be correlated to executions and process definitions which belong to no tenant or not. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. Must not be supplied in conjunction with a &#x60;tenantId&#x60;. | [optional] [default to false]
**ProcessInstanceId** | Pointer to **NullableString** | Used to correlate the message to the process instance with the given id. | [optional] 
**CorrelationKeys** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | Used for correlation of process instances that wait for incoming messages. Has to be a JSON object containing key-value pairs that are matched against process instance variables during correlation. Each key is a variable name and each value a JSON variable value object with the following properties. | [optional] 
**LocalCorrelationKeys** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | Local variables used for correlation of executions (process instances) that wait for incoming messages. Has to be a JSON object containing key-value pairs that are matched against local variables during correlation. Each key is a variable name and each value a JSON variable value object with the following properties. | [optional] 
**ProcessVariables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A map of variables that is injected into the triggered execution or process instance after the message has been delivered. Each key is a variable name and each value a JSON variable value object with the following properties. | [optional] 
**ProcessVariablesLocal** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A map of local variables that is injected into the execution waiting on the message. Each key is a variable name and each value a JSON variable value object with the following properties. | [optional] 
**ProcessVariablesToTriggeredScope** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A map of variables that is injected into the new scope triggered by message correlation. Each key is a variable name and each value a JSON variable value object with the following properties. | [optional] 
**All** | Pointer to **NullableBool** | A Boolean value that indicates whether the message should be correlated to exactly one entity or multiple entities. If the value is set to &#x60;false&#x60;, the message will be correlated to exactly one entity (execution or process definition). If the value is set to &#x60;true&#x60;, the message will be correlated to multiple executions and a process definition that can be instantiated by this message in one go. | [optional] [default to false]
**ResultEnabled** | Pointer to **NullableBool** | A Boolean value that indicates whether the result of the correlation should be returned or not. If this property is set to &#x60;true&#x60;, there will be returned a list of message correlation result objects. Depending on the all property, there will be either one ore more returned results in the list.  The default value is &#x60;false&#x60;, which means no result will be returned. | [optional] [default to false]
**VariablesInResultEnabled** | Pointer to **NullableBool** | A Boolean value that indicates whether the result of the correlation should contain process variables or not. The parameter resultEnabled should be set to &#x60;true&#x60; in order to use this it.  The default value is &#x60;false&#x60;, which means the variables will not be returned. | [optional] [default to false]

## Methods

### NewCorrelationMessageDto

`func NewCorrelationMessageDto() *CorrelationMessageDto`

NewCorrelationMessageDto instantiates a new CorrelationMessageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrelationMessageDtoWithDefaults

`func NewCorrelationMessageDtoWithDefaults() *CorrelationMessageDto`

NewCorrelationMessageDtoWithDefaults instantiates a new CorrelationMessageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageName

`func (o *CorrelationMessageDto) GetMessageName() string`

GetMessageName returns the MessageName field if non-nil, zero value otherwise.

### GetMessageNameOk

`func (o *CorrelationMessageDto) GetMessageNameOk() (*string, bool)`

GetMessageNameOk returns a tuple with the MessageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageName

`func (o *CorrelationMessageDto) SetMessageName(v string)`

SetMessageName sets MessageName field to given value.

### HasMessageName

`func (o *CorrelationMessageDto) HasMessageName() bool`

HasMessageName returns a boolean if a field has been set.

### SetMessageNameNil

`func (o *CorrelationMessageDto) SetMessageNameNil(b bool)`

 SetMessageNameNil sets the value for MessageName to be an explicit nil

### UnsetMessageName
`func (o *CorrelationMessageDto) UnsetMessageName()`

UnsetMessageName ensures that no value is present for MessageName, not even an explicit nil
### GetBusinessKey

`func (o *CorrelationMessageDto) GetBusinessKey() string`

GetBusinessKey returns the BusinessKey field if non-nil, zero value otherwise.

### GetBusinessKeyOk

`func (o *CorrelationMessageDto) GetBusinessKeyOk() (*string, bool)`

GetBusinessKeyOk returns a tuple with the BusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKey

`func (o *CorrelationMessageDto) SetBusinessKey(v string)`

SetBusinessKey sets BusinessKey field to given value.

### HasBusinessKey

`func (o *CorrelationMessageDto) HasBusinessKey() bool`

HasBusinessKey returns a boolean if a field has been set.

### SetBusinessKeyNil

`func (o *CorrelationMessageDto) SetBusinessKeyNil(b bool)`

 SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil

### UnsetBusinessKey
`func (o *CorrelationMessageDto) UnsetBusinessKey()`

UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
### GetTenantId

`func (o *CorrelationMessageDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CorrelationMessageDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CorrelationMessageDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CorrelationMessageDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CorrelationMessageDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CorrelationMessageDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetWithoutTenantId

`func (o *CorrelationMessageDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *CorrelationMessageDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *CorrelationMessageDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *CorrelationMessageDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *CorrelationMessageDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *CorrelationMessageDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetProcessInstanceId

`func (o *CorrelationMessageDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *CorrelationMessageDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *CorrelationMessageDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *CorrelationMessageDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *CorrelationMessageDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *CorrelationMessageDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetCorrelationKeys

`func (o *CorrelationMessageDto) GetCorrelationKeys() map[string]VariableValueDto`

GetCorrelationKeys returns the CorrelationKeys field if non-nil, zero value otherwise.

### GetCorrelationKeysOk

`func (o *CorrelationMessageDto) GetCorrelationKeysOk() (*map[string]VariableValueDto, bool)`

GetCorrelationKeysOk returns a tuple with the CorrelationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationKeys

`func (o *CorrelationMessageDto) SetCorrelationKeys(v map[string]VariableValueDto)`

SetCorrelationKeys sets CorrelationKeys field to given value.

### HasCorrelationKeys

`func (o *CorrelationMessageDto) HasCorrelationKeys() bool`

HasCorrelationKeys returns a boolean if a field has been set.

### SetCorrelationKeysNil

`func (o *CorrelationMessageDto) SetCorrelationKeysNil(b bool)`

 SetCorrelationKeysNil sets the value for CorrelationKeys to be an explicit nil

### UnsetCorrelationKeys
`func (o *CorrelationMessageDto) UnsetCorrelationKeys()`

UnsetCorrelationKeys ensures that no value is present for CorrelationKeys, not even an explicit nil
### GetLocalCorrelationKeys

`func (o *CorrelationMessageDto) GetLocalCorrelationKeys() map[string]VariableValueDto`

GetLocalCorrelationKeys returns the LocalCorrelationKeys field if non-nil, zero value otherwise.

### GetLocalCorrelationKeysOk

`func (o *CorrelationMessageDto) GetLocalCorrelationKeysOk() (*map[string]VariableValueDto, bool)`

GetLocalCorrelationKeysOk returns a tuple with the LocalCorrelationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCorrelationKeys

`func (o *CorrelationMessageDto) SetLocalCorrelationKeys(v map[string]VariableValueDto)`

SetLocalCorrelationKeys sets LocalCorrelationKeys field to given value.

### HasLocalCorrelationKeys

`func (o *CorrelationMessageDto) HasLocalCorrelationKeys() bool`

HasLocalCorrelationKeys returns a boolean if a field has been set.

### SetLocalCorrelationKeysNil

`func (o *CorrelationMessageDto) SetLocalCorrelationKeysNil(b bool)`

 SetLocalCorrelationKeysNil sets the value for LocalCorrelationKeys to be an explicit nil

### UnsetLocalCorrelationKeys
`func (o *CorrelationMessageDto) UnsetLocalCorrelationKeys()`

UnsetLocalCorrelationKeys ensures that no value is present for LocalCorrelationKeys, not even an explicit nil
### GetProcessVariables

`func (o *CorrelationMessageDto) GetProcessVariables() map[string]VariableValueDto`

GetProcessVariables returns the ProcessVariables field if non-nil, zero value otherwise.

### GetProcessVariablesOk

`func (o *CorrelationMessageDto) GetProcessVariablesOk() (*map[string]VariableValueDto, bool)`

GetProcessVariablesOk returns a tuple with the ProcessVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessVariables

`func (o *CorrelationMessageDto) SetProcessVariables(v map[string]VariableValueDto)`

SetProcessVariables sets ProcessVariables field to given value.

### HasProcessVariables

`func (o *CorrelationMessageDto) HasProcessVariables() bool`

HasProcessVariables returns a boolean if a field has been set.

### SetProcessVariablesNil

`func (o *CorrelationMessageDto) SetProcessVariablesNil(b bool)`

 SetProcessVariablesNil sets the value for ProcessVariables to be an explicit nil

### UnsetProcessVariables
`func (o *CorrelationMessageDto) UnsetProcessVariables()`

UnsetProcessVariables ensures that no value is present for ProcessVariables, not even an explicit nil
### GetProcessVariablesLocal

`func (o *CorrelationMessageDto) GetProcessVariablesLocal() map[string]VariableValueDto`

GetProcessVariablesLocal returns the ProcessVariablesLocal field if non-nil, zero value otherwise.

### GetProcessVariablesLocalOk

`func (o *CorrelationMessageDto) GetProcessVariablesLocalOk() (*map[string]VariableValueDto, bool)`

GetProcessVariablesLocalOk returns a tuple with the ProcessVariablesLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessVariablesLocal

`func (o *CorrelationMessageDto) SetProcessVariablesLocal(v map[string]VariableValueDto)`

SetProcessVariablesLocal sets ProcessVariablesLocal field to given value.

### HasProcessVariablesLocal

`func (o *CorrelationMessageDto) HasProcessVariablesLocal() bool`

HasProcessVariablesLocal returns a boolean if a field has been set.

### SetProcessVariablesLocalNil

`func (o *CorrelationMessageDto) SetProcessVariablesLocalNil(b bool)`

 SetProcessVariablesLocalNil sets the value for ProcessVariablesLocal to be an explicit nil

### UnsetProcessVariablesLocal
`func (o *CorrelationMessageDto) UnsetProcessVariablesLocal()`

UnsetProcessVariablesLocal ensures that no value is present for ProcessVariablesLocal, not even an explicit nil
### GetProcessVariablesToTriggeredScope

`func (o *CorrelationMessageDto) GetProcessVariablesToTriggeredScope() map[string]VariableValueDto`

GetProcessVariablesToTriggeredScope returns the ProcessVariablesToTriggeredScope field if non-nil, zero value otherwise.

### GetProcessVariablesToTriggeredScopeOk

`func (o *CorrelationMessageDto) GetProcessVariablesToTriggeredScopeOk() (*map[string]VariableValueDto, bool)`

GetProcessVariablesToTriggeredScopeOk returns a tuple with the ProcessVariablesToTriggeredScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessVariablesToTriggeredScope

`func (o *CorrelationMessageDto) SetProcessVariablesToTriggeredScope(v map[string]VariableValueDto)`

SetProcessVariablesToTriggeredScope sets ProcessVariablesToTriggeredScope field to given value.

### HasProcessVariablesToTriggeredScope

`func (o *CorrelationMessageDto) HasProcessVariablesToTriggeredScope() bool`

HasProcessVariablesToTriggeredScope returns a boolean if a field has been set.

### SetProcessVariablesToTriggeredScopeNil

`func (o *CorrelationMessageDto) SetProcessVariablesToTriggeredScopeNil(b bool)`

 SetProcessVariablesToTriggeredScopeNil sets the value for ProcessVariablesToTriggeredScope to be an explicit nil

### UnsetProcessVariablesToTriggeredScope
`func (o *CorrelationMessageDto) UnsetProcessVariablesToTriggeredScope()`

UnsetProcessVariablesToTriggeredScope ensures that no value is present for ProcessVariablesToTriggeredScope, not even an explicit nil
### GetAll

`func (o *CorrelationMessageDto) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *CorrelationMessageDto) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *CorrelationMessageDto) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *CorrelationMessageDto) HasAll() bool`

HasAll returns a boolean if a field has been set.

### SetAllNil

`func (o *CorrelationMessageDto) SetAllNil(b bool)`

 SetAllNil sets the value for All to be an explicit nil

### UnsetAll
`func (o *CorrelationMessageDto) UnsetAll()`

UnsetAll ensures that no value is present for All, not even an explicit nil
### GetResultEnabled

`func (o *CorrelationMessageDto) GetResultEnabled() bool`

GetResultEnabled returns the ResultEnabled field if non-nil, zero value otherwise.

### GetResultEnabledOk

`func (o *CorrelationMessageDto) GetResultEnabledOk() (*bool, bool)`

GetResultEnabledOk returns a tuple with the ResultEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEnabled

`func (o *CorrelationMessageDto) SetResultEnabled(v bool)`

SetResultEnabled sets ResultEnabled field to given value.

### HasResultEnabled

`func (o *CorrelationMessageDto) HasResultEnabled() bool`

HasResultEnabled returns a boolean if a field has been set.

### SetResultEnabledNil

`func (o *CorrelationMessageDto) SetResultEnabledNil(b bool)`

 SetResultEnabledNil sets the value for ResultEnabled to be an explicit nil

### UnsetResultEnabled
`func (o *CorrelationMessageDto) UnsetResultEnabled()`

UnsetResultEnabled ensures that no value is present for ResultEnabled, not even an explicit nil
### GetVariablesInResultEnabled

`func (o *CorrelationMessageDto) GetVariablesInResultEnabled() bool`

GetVariablesInResultEnabled returns the VariablesInResultEnabled field if non-nil, zero value otherwise.

### GetVariablesInResultEnabledOk

`func (o *CorrelationMessageDto) GetVariablesInResultEnabledOk() (*bool, bool)`

GetVariablesInResultEnabledOk returns a tuple with the VariablesInResultEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesInResultEnabled

`func (o *CorrelationMessageDto) SetVariablesInResultEnabled(v bool)`

SetVariablesInResultEnabled sets VariablesInResultEnabled field to given value.

### HasVariablesInResultEnabled

`func (o *CorrelationMessageDto) HasVariablesInResultEnabled() bool`

HasVariablesInResultEnabled returns a boolean if a field has been set.

### SetVariablesInResultEnabledNil

`func (o *CorrelationMessageDto) SetVariablesInResultEnabledNil(b bool)`

 SetVariablesInResultEnabledNil sets the value for VariablesInResultEnabled to be an explicit nil

### UnsetVariablesInResultEnabled
`func (o *CorrelationMessageDto) UnsetVariablesInResultEnabled()`

UnsetVariablesInResultEnabled ensures that no value is present for VariablesInResultEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


