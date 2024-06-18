# HistoricDecisionInputInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the decision input value. | [optional] 
**DecisionInstanceId** | Pointer to **NullableString** | The id of the decision instance the input value belongs to. | [optional] 
**ClauseId** | Pointer to **NullableString** | The id of the clause the input value belongs to. | [optional] 
**ClauseName** | Pointer to **NullableString** | The name of the clause the input value belongs to. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | An error message in case a Java Serialized Object could not be de-serialized. | [optional] 
**Type** | Pointer to **NullableString** | The value type of the variable. | [optional] 
**CreateTime** | Pointer to **NullableTime** | The time the variable was inserted.  [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the entry should be removed by the History Cleanup job. [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process containing this entry. | [optional] 
**Value** | Pointer to **map[string]interface{}** | The variable&#39;s value. Value differs depending on the variable&#39;s type and on the &#x60;disableCustomObjectDeserialization&#x60; parameter. | [optional] 
**ValueInfo** | Pointer to **map[string]interface{}** | A JSON object containing additional, value-type-dependent properties.  For variables of type &#x60;Object&#x60;, the following properties are returned:  * &#x60;objectTypeName&#x60;: A string representation of the object&#39;s type name.  * &#x60;serializationDataFormat&#x60;: The serialization format used to store the variable. | [optional] 

## Methods

### NewHistoricDecisionInputInstanceDto

`func NewHistoricDecisionInputInstanceDto() *HistoricDecisionInputInstanceDto`

NewHistoricDecisionInputInstanceDto instantiates a new HistoricDecisionInputInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricDecisionInputInstanceDtoWithDefaults

`func NewHistoricDecisionInputInstanceDtoWithDefaults() *HistoricDecisionInputInstanceDto`

NewHistoricDecisionInputInstanceDtoWithDefaults instantiates a new HistoricDecisionInputInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricDecisionInputInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricDecisionInputInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricDecisionInputInstanceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricDecisionInputInstanceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricDecisionInputInstanceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricDecisionInputInstanceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDecisionInstanceId

`func (o *HistoricDecisionInputInstanceDto) GetDecisionInstanceId() string`

GetDecisionInstanceId returns the DecisionInstanceId field if non-nil, zero value otherwise.

### GetDecisionInstanceIdOk

`func (o *HistoricDecisionInputInstanceDto) GetDecisionInstanceIdOk() (*string, bool)`

GetDecisionInstanceIdOk returns a tuple with the DecisionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionInstanceId

`func (o *HistoricDecisionInputInstanceDto) SetDecisionInstanceId(v string)`

SetDecisionInstanceId sets DecisionInstanceId field to given value.

### HasDecisionInstanceId

`func (o *HistoricDecisionInputInstanceDto) HasDecisionInstanceId() bool`

HasDecisionInstanceId returns a boolean if a field has been set.

### SetDecisionInstanceIdNil

`func (o *HistoricDecisionInputInstanceDto) SetDecisionInstanceIdNil(b bool)`

 SetDecisionInstanceIdNil sets the value for DecisionInstanceId to be an explicit nil

### UnsetDecisionInstanceId
`func (o *HistoricDecisionInputInstanceDto) UnsetDecisionInstanceId()`

UnsetDecisionInstanceId ensures that no value is present for DecisionInstanceId, not even an explicit nil
### GetClauseId

`func (o *HistoricDecisionInputInstanceDto) GetClauseId() string`

GetClauseId returns the ClauseId field if non-nil, zero value otherwise.

### GetClauseIdOk

`func (o *HistoricDecisionInputInstanceDto) GetClauseIdOk() (*string, bool)`

GetClauseIdOk returns a tuple with the ClauseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauseId

`func (o *HistoricDecisionInputInstanceDto) SetClauseId(v string)`

SetClauseId sets ClauseId field to given value.

### HasClauseId

`func (o *HistoricDecisionInputInstanceDto) HasClauseId() bool`

HasClauseId returns a boolean if a field has been set.

### SetClauseIdNil

`func (o *HistoricDecisionInputInstanceDto) SetClauseIdNil(b bool)`

 SetClauseIdNil sets the value for ClauseId to be an explicit nil

### UnsetClauseId
`func (o *HistoricDecisionInputInstanceDto) UnsetClauseId()`

UnsetClauseId ensures that no value is present for ClauseId, not even an explicit nil
### GetClauseName

`func (o *HistoricDecisionInputInstanceDto) GetClauseName() string`

GetClauseName returns the ClauseName field if non-nil, zero value otherwise.

### GetClauseNameOk

`func (o *HistoricDecisionInputInstanceDto) GetClauseNameOk() (*string, bool)`

GetClauseNameOk returns a tuple with the ClauseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauseName

`func (o *HistoricDecisionInputInstanceDto) SetClauseName(v string)`

SetClauseName sets ClauseName field to given value.

### HasClauseName

`func (o *HistoricDecisionInputInstanceDto) HasClauseName() bool`

HasClauseName returns a boolean if a field has been set.

### SetClauseNameNil

`func (o *HistoricDecisionInputInstanceDto) SetClauseNameNil(b bool)`

 SetClauseNameNil sets the value for ClauseName to be an explicit nil

### UnsetClauseName
`func (o *HistoricDecisionInputInstanceDto) UnsetClauseName()`

UnsetClauseName ensures that no value is present for ClauseName, not even an explicit nil
### GetErrorMessage

`func (o *HistoricDecisionInputInstanceDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *HistoricDecisionInputInstanceDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *HistoricDecisionInputInstanceDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *HistoricDecisionInputInstanceDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *HistoricDecisionInputInstanceDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *HistoricDecisionInputInstanceDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetType

`func (o *HistoricDecisionInputInstanceDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoricDecisionInputInstanceDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoricDecisionInputInstanceDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HistoricDecisionInputInstanceDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HistoricDecisionInputInstanceDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HistoricDecisionInputInstanceDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCreateTime

`func (o *HistoricDecisionInputInstanceDto) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HistoricDecisionInputInstanceDto) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HistoricDecisionInputInstanceDto) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HistoricDecisionInputInstanceDto) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### SetCreateTimeNil

`func (o *HistoricDecisionInputInstanceDto) SetCreateTimeNil(b bool)`

 SetCreateTimeNil sets the value for CreateTime to be an explicit nil

### UnsetCreateTime
`func (o *HistoricDecisionInputInstanceDto) UnsetCreateTime()`

UnsetCreateTime ensures that no value is present for CreateTime, not even an explicit nil
### GetRemovalTime

`func (o *HistoricDecisionInputInstanceDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricDecisionInputInstanceDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricDecisionInputInstanceDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricDecisionInputInstanceDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricDecisionInputInstanceDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricDecisionInputInstanceDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricDecisionInputInstanceDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricDecisionInputInstanceDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricDecisionInputInstanceDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricDecisionInputInstanceDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricDecisionInputInstanceDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricDecisionInputInstanceDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil
### GetValue

`func (o *HistoricDecisionInputInstanceDto) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HistoricDecisionInputInstanceDto) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HistoricDecisionInputInstanceDto) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *HistoricDecisionInputInstanceDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueInfo

`func (o *HistoricDecisionInputInstanceDto) GetValueInfo() map[string]interface{}`

GetValueInfo returns the ValueInfo field if non-nil, zero value otherwise.

### GetValueInfoOk

`func (o *HistoricDecisionInputInstanceDto) GetValueInfoOk() (*map[string]interface{}, bool)`

GetValueInfoOk returns a tuple with the ValueInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInfo

`func (o *HistoricDecisionInputInstanceDto) SetValueInfo(v map[string]interface{})`

SetValueInfo sets ValueInfo field to given value.

### HasValueInfo

`func (o *HistoricDecisionInputInstanceDto) HasValueInfo() bool`

HasValueInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


