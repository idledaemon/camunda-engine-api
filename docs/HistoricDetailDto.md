# HistoricDetailDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the historic detail. | [optional] 
**Type** | Pointer to **NullableString** | The type of the historic detail. Either &#x60;formField&#x60; for a submitted form field value or &#x60;variableUpdate&#x60; for variable updates. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition that this historic detail belongs to. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition that this historic detail belongs to. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance the historic detail belongs to. | [optional] 
**ActivityInstanceId** | Pointer to **NullableString** | The id of the activity instance the historic detail belongs to. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The id of the execution the historic detail belongs to. | [optional] 
**CaseDefinitionKey** | Pointer to **NullableString** | The key of the case definition that this historic detail belongs to. | [optional] 
**CaseDefinitionId** | Pointer to **NullableString** | The id of the case definition that this historic detail belongs to. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | The id of the case instance the historic detail belongs to. | [optional] 
**CaseExecutionId** | Pointer to **NullableString** | The id of the case execution the historic detail belongs to. | [optional] 
**TaskId** | Pointer to **NullableString** | The id of the task the historic detail belongs to. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant that this historic detail belongs to. | [optional] 
**UserOperationId** | Pointer to **NullableString** | The id of user operation which links historic detail with [user operation log](https://docs.camunda.org/manual/7.21/reference/rest/history/user-operation-log/) entries. | [optional] 
**Time** | Pointer to **NullableTime** | The time when this historic detail occurred. Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the historic detail should be removed by the History Cleanup job. Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process containing this historic detail. | [optional] 
**FieldId** | Pointer to **NullableString** | The id of the form field.  **Note:** This property is only set for a &#x60;HistoricVariableUpdate&#x60; historic details. In these cases, the value of the &#x60;type&#x60; property is &#x60;formField&#x60;. | [optional] 
**FieldValue** | Pointer to **map[string]interface{}** | The submitted form field value. The value differs depending on the form field&#39;s type and on the &#x60;deserializeValue&#x60; parameter.  **Note:** This property is only set for a &#x60;HistoricVariableUpdate&#x60; historic details. In these cases, the value of the &#x60;type&#x60; property is &#x60;formField&#x60;. | [optional] 
**VariableName** | Pointer to **NullableString** | The name of the variable which has been updated.  **Note:** This property is only set for a &#x60;HistoricVariableUpdate&#x60; historic details. In these cases, the value of the &#x60;type&#x60; property is &#x60;variableUpdate&#x60;. | [optional] 
**VariableInstanceId** | Pointer to **NullableString** | The id of the associated variable instance.  **Note:** This property is only set for a &#x60;HistoricVariableUpdate&#x60; historic details. In these cases, the value of the &#x60;type&#x60; property is &#x60;variableUpdate&#x60;. | [optional] 
**VariableType** | Pointer to **NullableString** | The value type of the variable.  **Note:** This property is only set for a &#x60;HistoricVariableUpdate&#x60; historic details. In these cases, the value of the &#x60;type&#x60; property is &#x60;variableUpdate&#x60;. | [optional] 
**Value** | Pointer to **map[string]interface{}** | The variable&#39;s value. Value differs depending on the variable&#39;s type and on the deserializeValues parameter.  **Note:** This property is only set for a &#x60;HistoricVariableUpdate&#x60; historic details. In these cases, the value of the &#x60;type&#x60; property is &#x60;variableUpdate&#x60;. | [optional] 
**ValueInfo** | Pointer to **map[string]interface{}** | A JSON object containing additional, value-type-dependent properties. For variables of type &#x60;Object&#x60;, the following properties are returned:  * &#x60;objectTypeName&#x60;: A string representation of the object&#39;s type name. * &#x60;serializationDataFormat&#x60;: The serialization format used to store the variable.  **Note:** This property is only set for a &#x60;HistoricVariableUpdate&#x60; historic details. In these cases, the value of the &#x60;type&#x60; property is &#x60;variableUpdate&#x60;. | [optional] 
**Initial** | Pointer to **NullableBool** | Returns &#x60;true&#x60; for variable updates that contains the initial values of the variables.  **Note:** This property is only set for a &#x60;HistoricVariableUpdate&#x60; historic details. In these cases, the value of the &#x60;type&#x60; property is &#x60;variableUpdate&#x60;. | [optional] 
**Revision** | Pointer to **NullableInt32** | The revision of the historic variable update.  **Note:** This property is only set for a &#x60;HistoricVariableUpdate&#x60; historic details. In these cases, the value of the &#x60;type&#x60; property is &#x60;variableUpdate&#x60;. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | An error message in case a Java Serialized Object could not be de-serialized.  **Note:** This property is only set for a &#x60;HistoricVariableUpdate&#x60; historic details. In these cases, the value of the &#x60;type&#x60; property is &#x60;variableUpdate&#x60;. | [optional] 

## Methods

### NewHistoricDetailDto

`func NewHistoricDetailDto() *HistoricDetailDto`

NewHistoricDetailDto instantiates a new HistoricDetailDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricDetailDtoWithDefaults

`func NewHistoricDetailDtoWithDefaults() *HistoricDetailDto`

NewHistoricDetailDtoWithDefaults instantiates a new HistoricDetailDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricDetailDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricDetailDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricDetailDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricDetailDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricDetailDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricDetailDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *HistoricDetailDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoricDetailDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoricDetailDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HistoricDetailDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HistoricDetailDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HistoricDetailDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricDetailDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricDetailDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricDetailDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricDetailDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricDetailDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricDetailDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricDetailDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricDetailDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricDetailDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricDetailDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricDetailDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricDetailDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricDetailDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricDetailDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricDetailDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricDetailDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricDetailDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricDetailDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetActivityInstanceId

`func (o *HistoricDetailDto) GetActivityInstanceId() string`

GetActivityInstanceId returns the ActivityInstanceId field if non-nil, zero value otherwise.

### GetActivityInstanceIdOk

`func (o *HistoricDetailDto) GetActivityInstanceIdOk() (*string, bool)`

GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceId

`func (o *HistoricDetailDto) SetActivityInstanceId(v string)`

SetActivityInstanceId sets ActivityInstanceId field to given value.

### HasActivityInstanceId

`func (o *HistoricDetailDto) HasActivityInstanceId() bool`

HasActivityInstanceId returns a boolean if a field has been set.

### SetActivityInstanceIdNil

`func (o *HistoricDetailDto) SetActivityInstanceIdNil(b bool)`

 SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil

### UnsetActivityInstanceId
`func (o *HistoricDetailDto) UnsetActivityInstanceId()`

UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
### GetExecutionId

`func (o *HistoricDetailDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HistoricDetailDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HistoricDetailDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HistoricDetailDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HistoricDetailDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HistoricDetailDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetCaseDefinitionKey

`func (o *HistoricDetailDto) GetCaseDefinitionKey() string`

GetCaseDefinitionKey returns the CaseDefinitionKey field if non-nil, zero value otherwise.

### GetCaseDefinitionKeyOk

`func (o *HistoricDetailDto) GetCaseDefinitionKeyOk() (*string, bool)`

GetCaseDefinitionKeyOk returns a tuple with the CaseDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionKey

`func (o *HistoricDetailDto) SetCaseDefinitionKey(v string)`

SetCaseDefinitionKey sets CaseDefinitionKey field to given value.

### HasCaseDefinitionKey

`func (o *HistoricDetailDto) HasCaseDefinitionKey() bool`

HasCaseDefinitionKey returns a boolean if a field has been set.

### SetCaseDefinitionKeyNil

`func (o *HistoricDetailDto) SetCaseDefinitionKeyNil(b bool)`

 SetCaseDefinitionKeyNil sets the value for CaseDefinitionKey to be an explicit nil

### UnsetCaseDefinitionKey
`func (o *HistoricDetailDto) UnsetCaseDefinitionKey()`

UnsetCaseDefinitionKey ensures that no value is present for CaseDefinitionKey, not even an explicit nil
### GetCaseDefinitionId

`func (o *HistoricDetailDto) GetCaseDefinitionId() string`

GetCaseDefinitionId returns the CaseDefinitionId field if non-nil, zero value otherwise.

### GetCaseDefinitionIdOk

`func (o *HistoricDetailDto) GetCaseDefinitionIdOk() (*string, bool)`

GetCaseDefinitionIdOk returns a tuple with the CaseDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionId

`func (o *HistoricDetailDto) SetCaseDefinitionId(v string)`

SetCaseDefinitionId sets CaseDefinitionId field to given value.

### HasCaseDefinitionId

`func (o *HistoricDetailDto) HasCaseDefinitionId() bool`

HasCaseDefinitionId returns a boolean if a field has been set.

### SetCaseDefinitionIdNil

`func (o *HistoricDetailDto) SetCaseDefinitionIdNil(b bool)`

 SetCaseDefinitionIdNil sets the value for CaseDefinitionId to be an explicit nil

### UnsetCaseDefinitionId
`func (o *HistoricDetailDto) UnsetCaseDefinitionId()`

UnsetCaseDefinitionId ensures that no value is present for CaseDefinitionId, not even an explicit nil
### GetCaseInstanceId

`func (o *HistoricDetailDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *HistoricDetailDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *HistoricDetailDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *HistoricDetailDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *HistoricDetailDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *HistoricDetailDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetCaseExecutionId

`func (o *HistoricDetailDto) GetCaseExecutionId() string`

GetCaseExecutionId returns the CaseExecutionId field if non-nil, zero value otherwise.

### GetCaseExecutionIdOk

`func (o *HistoricDetailDto) GetCaseExecutionIdOk() (*string, bool)`

GetCaseExecutionIdOk returns a tuple with the CaseExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExecutionId

`func (o *HistoricDetailDto) SetCaseExecutionId(v string)`

SetCaseExecutionId sets CaseExecutionId field to given value.

### HasCaseExecutionId

`func (o *HistoricDetailDto) HasCaseExecutionId() bool`

HasCaseExecutionId returns a boolean if a field has been set.

### SetCaseExecutionIdNil

`func (o *HistoricDetailDto) SetCaseExecutionIdNil(b bool)`

 SetCaseExecutionIdNil sets the value for CaseExecutionId to be an explicit nil

### UnsetCaseExecutionId
`func (o *HistoricDetailDto) UnsetCaseExecutionId()`

UnsetCaseExecutionId ensures that no value is present for CaseExecutionId, not even an explicit nil
### GetTaskId

`func (o *HistoricDetailDto) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *HistoricDetailDto) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *HistoricDetailDto) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *HistoricDetailDto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *HistoricDetailDto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *HistoricDetailDto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetTenantId

`func (o *HistoricDetailDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricDetailDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricDetailDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricDetailDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricDetailDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricDetailDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUserOperationId

`func (o *HistoricDetailDto) GetUserOperationId() string`

GetUserOperationId returns the UserOperationId field if non-nil, zero value otherwise.

### GetUserOperationIdOk

`func (o *HistoricDetailDto) GetUserOperationIdOk() (*string, bool)`

GetUserOperationIdOk returns a tuple with the UserOperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOperationId

`func (o *HistoricDetailDto) SetUserOperationId(v string)`

SetUserOperationId sets UserOperationId field to given value.

### HasUserOperationId

`func (o *HistoricDetailDto) HasUserOperationId() bool`

HasUserOperationId returns a boolean if a field has been set.

### SetUserOperationIdNil

`func (o *HistoricDetailDto) SetUserOperationIdNil(b bool)`

 SetUserOperationIdNil sets the value for UserOperationId to be an explicit nil

### UnsetUserOperationId
`func (o *HistoricDetailDto) UnsetUserOperationId()`

UnsetUserOperationId ensures that no value is present for UserOperationId, not even an explicit nil
### GetTime

`func (o *HistoricDetailDto) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *HistoricDetailDto) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *HistoricDetailDto) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *HistoricDetailDto) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *HistoricDetailDto) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *HistoricDetailDto) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil
### GetRemovalTime

`func (o *HistoricDetailDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricDetailDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricDetailDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricDetailDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricDetailDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricDetailDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricDetailDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricDetailDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricDetailDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricDetailDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricDetailDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricDetailDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil
### GetFieldId

`func (o *HistoricDetailDto) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *HistoricDetailDto) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *HistoricDetailDto) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *HistoricDetailDto) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### SetFieldIdNil

`func (o *HistoricDetailDto) SetFieldIdNil(b bool)`

 SetFieldIdNil sets the value for FieldId to be an explicit nil

### UnsetFieldId
`func (o *HistoricDetailDto) UnsetFieldId()`

UnsetFieldId ensures that no value is present for FieldId, not even an explicit nil
### GetFieldValue

`func (o *HistoricDetailDto) GetFieldValue() map[string]interface{}`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *HistoricDetailDto) GetFieldValueOk() (*map[string]interface{}, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *HistoricDetailDto) SetFieldValue(v map[string]interface{})`

SetFieldValue sets FieldValue field to given value.

### HasFieldValue

`func (o *HistoricDetailDto) HasFieldValue() bool`

HasFieldValue returns a boolean if a field has been set.

### GetVariableName

`func (o *HistoricDetailDto) GetVariableName() string`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *HistoricDetailDto) GetVariableNameOk() (*string, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *HistoricDetailDto) SetVariableName(v string)`

SetVariableName sets VariableName field to given value.

### HasVariableName

`func (o *HistoricDetailDto) HasVariableName() bool`

HasVariableName returns a boolean if a field has been set.

### SetVariableNameNil

`func (o *HistoricDetailDto) SetVariableNameNil(b bool)`

 SetVariableNameNil sets the value for VariableName to be an explicit nil

### UnsetVariableName
`func (o *HistoricDetailDto) UnsetVariableName()`

UnsetVariableName ensures that no value is present for VariableName, not even an explicit nil
### GetVariableInstanceId

`func (o *HistoricDetailDto) GetVariableInstanceId() string`

GetVariableInstanceId returns the VariableInstanceId field if non-nil, zero value otherwise.

### GetVariableInstanceIdOk

`func (o *HistoricDetailDto) GetVariableInstanceIdOk() (*string, bool)`

GetVariableInstanceIdOk returns a tuple with the VariableInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableInstanceId

`func (o *HistoricDetailDto) SetVariableInstanceId(v string)`

SetVariableInstanceId sets VariableInstanceId field to given value.

### HasVariableInstanceId

`func (o *HistoricDetailDto) HasVariableInstanceId() bool`

HasVariableInstanceId returns a boolean if a field has been set.

### SetVariableInstanceIdNil

`func (o *HistoricDetailDto) SetVariableInstanceIdNil(b bool)`

 SetVariableInstanceIdNil sets the value for VariableInstanceId to be an explicit nil

### UnsetVariableInstanceId
`func (o *HistoricDetailDto) UnsetVariableInstanceId()`

UnsetVariableInstanceId ensures that no value is present for VariableInstanceId, not even an explicit nil
### GetVariableType

`func (o *HistoricDetailDto) GetVariableType() string`

GetVariableType returns the VariableType field if non-nil, zero value otherwise.

### GetVariableTypeOk

`func (o *HistoricDetailDto) GetVariableTypeOk() (*string, bool)`

GetVariableTypeOk returns a tuple with the VariableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableType

`func (o *HistoricDetailDto) SetVariableType(v string)`

SetVariableType sets VariableType field to given value.

### HasVariableType

`func (o *HistoricDetailDto) HasVariableType() bool`

HasVariableType returns a boolean if a field has been set.

### SetVariableTypeNil

`func (o *HistoricDetailDto) SetVariableTypeNil(b bool)`

 SetVariableTypeNil sets the value for VariableType to be an explicit nil

### UnsetVariableType
`func (o *HistoricDetailDto) UnsetVariableType()`

UnsetVariableType ensures that no value is present for VariableType, not even an explicit nil
### GetValue

`func (o *HistoricDetailDto) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HistoricDetailDto) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HistoricDetailDto) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *HistoricDetailDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueInfo

`func (o *HistoricDetailDto) GetValueInfo() map[string]interface{}`

GetValueInfo returns the ValueInfo field if non-nil, zero value otherwise.

### GetValueInfoOk

`func (o *HistoricDetailDto) GetValueInfoOk() (*map[string]interface{}, bool)`

GetValueInfoOk returns a tuple with the ValueInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInfo

`func (o *HistoricDetailDto) SetValueInfo(v map[string]interface{})`

SetValueInfo sets ValueInfo field to given value.

### HasValueInfo

`func (o *HistoricDetailDto) HasValueInfo() bool`

HasValueInfo returns a boolean if a field has been set.

### GetInitial

`func (o *HistoricDetailDto) GetInitial() bool`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *HistoricDetailDto) GetInitialOk() (*bool, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *HistoricDetailDto) SetInitial(v bool)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *HistoricDetailDto) HasInitial() bool`

HasInitial returns a boolean if a field has been set.

### SetInitialNil

`func (o *HistoricDetailDto) SetInitialNil(b bool)`

 SetInitialNil sets the value for Initial to be an explicit nil

### UnsetInitial
`func (o *HistoricDetailDto) UnsetInitial()`

UnsetInitial ensures that no value is present for Initial, not even an explicit nil
### GetRevision

`func (o *HistoricDetailDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *HistoricDetailDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *HistoricDetailDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *HistoricDetailDto) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *HistoricDetailDto) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *HistoricDetailDto) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetErrorMessage

`func (o *HistoricDetailDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *HistoricDetailDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *HistoricDetailDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *HistoricDetailDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *HistoricDetailDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *HistoricDetailDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


