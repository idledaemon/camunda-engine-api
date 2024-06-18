# HistoricDecisionOutputInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the decision output value. | [optional] 
**DecisionInstanceId** | Pointer to **NullableString** | The id of the decision instance the output value belongs to. | [optional] 
**ClauseId** | Pointer to **NullableString** | The id of the clause the output value belongs to. | [optional] 
**ClauseName** | Pointer to **NullableString** | The name of the clause the output value belongs to. | [optional] 
**RuleId** | Pointer to **NullableString** | The id of the rule the output value belongs to. | [optional] 
**RuleOrder** | Pointer to **NullableInt32** | The order of the rule the output value belongs to. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | An error message in case a Java Serialized Object could not be de-serialized. | [optional] 
**VariableName** | Pointer to **NullableString** | The name of the output variable. | [optional] 
**Type** | Pointer to **NullableString** | The value type of the variable. | [optional] 
**CreateTime** | Pointer to **NullableTime** | The time the variable was inserted.  [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the entry should be removed by the History Cleanup job. [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process containing this entry. | [optional] 
**Value** | Pointer to **map[string]interface{}** | The variable&#39;s value. Value differs depending on the variable&#39;s type and on the &#x60;disableCustomObjectDeserialization&#x60; parameter. | [optional] 
**ValueInfo** | Pointer to **map[string]interface{}** | A JSON object containing additional, value-type-dependent properties.  For variables of type &#x60;Object&#x60;, the following properties are returned:  * &#x60;objectTypeName&#x60;: A string representation of the object&#39;s type name.  * &#x60;serializationDataFormat&#x60;: The serialization format used to store the variable. | [optional] 

## Methods

### NewHistoricDecisionOutputInstanceDto

`func NewHistoricDecisionOutputInstanceDto() *HistoricDecisionOutputInstanceDto`

NewHistoricDecisionOutputInstanceDto instantiates a new HistoricDecisionOutputInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricDecisionOutputInstanceDtoWithDefaults

`func NewHistoricDecisionOutputInstanceDtoWithDefaults() *HistoricDecisionOutputInstanceDto`

NewHistoricDecisionOutputInstanceDtoWithDefaults instantiates a new HistoricDecisionOutputInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricDecisionOutputInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricDecisionOutputInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricDecisionOutputInstanceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricDecisionOutputInstanceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricDecisionOutputInstanceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricDecisionOutputInstanceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDecisionInstanceId

`func (o *HistoricDecisionOutputInstanceDto) GetDecisionInstanceId() string`

GetDecisionInstanceId returns the DecisionInstanceId field if non-nil, zero value otherwise.

### GetDecisionInstanceIdOk

`func (o *HistoricDecisionOutputInstanceDto) GetDecisionInstanceIdOk() (*string, bool)`

GetDecisionInstanceIdOk returns a tuple with the DecisionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionInstanceId

`func (o *HistoricDecisionOutputInstanceDto) SetDecisionInstanceId(v string)`

SetDecisionInstanceId sets DecisionInstanceId field to given value.

### HasDecisionInstanceId

`func (o *HistoricDecisionOutputInstanceDto) HasDecisionInstanceId() bool`

HasDecisionInstanceId returns a boolean if a field has been set.

### SetDecisionInstanceIdNil

`func (o *HistoricDecisionOutputInstanceDto) SetDecisionInstanceIdNil(b bool)`

 SetDecisionInstanceIdNil sets the value for DecisionInstanceId to be an explicit nil

### UnsetDecisionInstanceId
`func (o *HistoricDecisionOutputInstanceDto) UnsetDecisionInstanceId()`

UnsetDecisionInstanceId ensures that no value is present for DecisionInstanceId, not even an explicit nil
### GetClauseId

`func (o *HistoricDecisionOutputInstanceDto) GetClauseId() string`

GetClauseId returns the ClauseId field if non-nil, zero value otherwise.

### GetClauseIdOk

`func (o *HistoricDecisionOutputInstanceDto) GetClauseIdOk() (*string, bool)`

GetClauseIdOk returns a tuple with the ClauseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauseId

`func (o *HistoricDecisionOutputInstanceDto) SetClauseId(v string)`

SetClauseId sets ClauseId field to given value.

### HasClauseId

`func (o *HistoricDecisionOutputInstanceDto) HasClauseId() bool`

HasClauseId returns a boolean if a field has been set.

### SetClauseIdNil

`func (o *HistoricDecisionOutputInstanceDto) SetClauseIdNil(b bool)`

 SetClauseIdNil sets the value for ClauseId to be an explicit nil

### UnsetClauseId
`func (o *HistoricDecisionOutputInstanceDto) UnsetClauseId()`

UnsetClauseId ensures that no value is present for ClauseId, not even an explicit nil
### GetClauseName

`func (o *HistoricDecisionOutputInstanceDto) GetClauseName() string`

GetClauseName returns the ClauseName field if non-nil, zero value otherwise.

### GetClauseNameOk

`func (o *HistoricDecisionOutputInstanceDto) GetClauseNameOk() (*string, bool)`

GetClauseNameOk returns a tuple with the ClauseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauseName

`func (o *HistoricDecisionOutputInstanceDto) SetClauseName(v string)`

SetClauseName sets ClauseName field to given value.

### HasClauseName

`func (o *HistoricDecisionOutputInstanceDto) HasClauseName() bool`

HasClauseName returns a boolean if a field has been set.

### SetClauseNameNil

`func (o *HistoricDecisionOutputInstanceDto) SetClauseNameNil(b bool)`

 SetClauseNameNil sets the value for ClauseName to be an explicit nil

### UnsetClauseName
`func (o *HistoricDecisionOutputInstanceDto) UnsetClauseName()`

UnsetClauseName ensures that no value is present for ClauseName, not even an explicit nil
### GetRuleId

`func (o *HistoricDecisionOutputInstanceDto) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *HistoricDecisionOutputInstanceDto) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *HistoricDecisionOutputInstanceDto) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *HistoricDecisionOutputInstanceDto) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *HistoricDecisionOutputInstanceDto) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *HistoricDecisionOutputInstanceDto) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetRuleOrder

`func (o *HistoricDecisionOutputInstanceDto) GetRuleOrder() int32`

GetRuleOrder returns the RuleOrder field if non-nil, zero value otherwise.

### GetRuleOrderOk

`func (o *HistoricDecisionOutputInstanceDto) GetRuleOrderOk() (*int32, bool)`

GetRuleOrderOk returns a tuple with the RuleOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOrder

`func (o *HistoricDecisionOutputInstanceDto) SetRuleOrder(v int32)`

SetRuleOrder sets RuleOrder field to given value.

### HasRuleOrder

`func (o *HistoricDecisionOutputInstanceDto) HasRuleOrder() bool`

HasRuleOrder returns a boolean if a field has been set.

### SetRuleOrderNil

`func (o *HistoricDecisionOutputInstanceDto) SetRuleOrderNil(b bool)`

 SetRuleOrderNil sets the value for RuleOrder to be an explicit nil

### UnsetRuleOrder
`func (o *HistoricDecisionOutputInstanceDto) UnsetRuleOrder()`

UnsetRuleOrder ensures that no value is present for RuleOrder, not even an explicit nil
### GetErrorMessage

`func (o *HistoricDecisionOutputInstanceDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *HistoricDecisionOutputInstanceDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *HistoricDecisionOutputInstanceDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *HistoricDecisionOutputInstanceDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *HistoricDecisionOutputInstanceDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *HistoricDecisionOutputInstanceDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetVariableName

`func (o *HistoricDecisionOutputInstanceDto) GetVariableName() string`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *HistoricDecisionOutputInstanceDto) GetVariableNameOk() (*string, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *HistoricDecisionOutputInstanceDto) SetVariableName(v string)`

SetVariableName sets VariableName field to given value.

### HasVariableName

`func (o *HistoricDecisionOutputInstanceDto) HasVariableName() bool`

HasVariableName returns a boolean if a field has been set.

### SetVariableNameNil

`func (o *HistoricDecisionOutputInstanceDto) SetVariableNameNil(b bool)`

 SetVariableNameNil sets the value for VariableName to be an explicit nil

### UnsetVariableName
`func (o *HistoricDecisionOutputInstanceDto) UnsetVariableName()`

UnsetVariableName ensures that no value is present for VariableName, not even an explicit nil
### GetType

`func (o *HistoricDecisionOutputInstanceDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoricDecisionOutputInstanceDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoricDecisionOutputInstanceDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HistoricDecisionOutputInstanceDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HistoricDecisionOutputInstanceDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HistoricDecisionOutputInstanceDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCreateTime

`func (o *HistoricDecisionOutputInstanceDto) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HistoricDecisionOutputInstanceDto) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HistoricDecisionOutputInstanceDto) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HistoricDecisionOutputInstanceDto) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### SetCreateTimeNil

`func (o *HistoricDecisionOutputInstanceDto) SetCreateTimeNil(b bool)`

 SetCreateTimeNil sets the value for CreateTime to be an explicit nil

### UnsetCreateTime
`func (o *HistoricDecisionOutputInstanceDto) UnsetCreateTime()`

UnsetCreateTime ensures that no value is present for CreateTime, not even an explicit nil
### GetRemovalTime

`func (o *HistoricDecisionOutputInstanceDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricDecisionOutputInstanceDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricDecisionOutputInstanceDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricDecisionOutputInstanceDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricDecisionOutputInstanceDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricDecisionOutputInstanceDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricDecisionOutputInstanceDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricDecisionOutputInstanceDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricDecisionOutputInstanceDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricDecisionOutputInstanceDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricDecisionOutputInstanceDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricDecisionOutputInstanceDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil
### GetValue

`func (o *HistoricDecisionOutputInstanceDto) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HistoricDecisionOutputInstanceDto) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HistoricDecisionOutputInstanceDto) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *HistoricDecisionOutputInstanceDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueInfo

`func (o *HistoricDecisionOutputInstanceDto) GetValueInfo() map[string]interface{}`

GetValueInfo returns the ValueInfo field if non-nil, zero value otherwise.

### GetValueInfoOk

`func (o *HistoricDecisionOutputInstanceDto) GetValueInfoOk() (*map[string]interface{}, bool)`

GetValueInfoOk returns a tuple with the ValueInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInfo

`func (o *HistoricDecisionOutputInstanceDto) SetValueInfo(v map[string]interface{})`

SetValueInfo sets ValueInfo field to given value.

### HasValueInfo

`func (o *HistoricDecisionOutputInstanceDto) HasValueInfo() bool`

HasValueInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


