# HistoricDecisionInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the decision instance. | [optional] 
**DecisionDefinitionId** | Pointer to **NullableString** | The id of the decision definition that this decision instance belongs to. | [optional] 
**DecisionDefinitionKey** | Pointer to **NullableString** | The key of the decision definition that this decision instance belongs to. | [optional] 
**DecisionDefinitionName** | Pointer to **NullableString** | The name of the decision definition that this decision instance belongs to. | [optional] 
**EvaluationTime** | Pointer to **NullableTime** | The time the instance was evaluated.  [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the instance should be removed by the History Cleanup job. [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition that this decision instance belongs to. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition that this decision instance belongs to. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance that this decision instance belongs to. | [optional] 
**CaseDefinitionId** | Pointer to **NullableString** | The id of the case definition that this decision instance belongs to. | [optional] 
**CaseDefinitionKey** | Pointer to **NullableString** | The key of the case definition that this decision instance belongs to. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | The id of the case instance that this decision instance belongs to. | [optional] 
**ActivityId** | Pointer to **NullableString** | The id of the activity that this decision instance belongs to. | [optional] 
**ActivityInstanceId** | Pointer to **NullableString** | The id of the activity instance that this decision instance belongs to. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the historic decision instance. | [optional] 
**UserId** | Pointer to **NullableString** | The id of the authenticated user that has evaluated this decision instance without a process or case instance. | [optional] 
**Inputs** | Pointer to [**[]HistoricDecisionInputInstanceDto**](HistoricDecisionInputInstanceDto.md) | The list of decision input values. **Only exists** if &#x60;includeInputs&#x60; was set to &#x60;true&#x60; in the query. | [optional] 
**Outputs** | Pointer to [**[]HistoricDecisionOutputInstanceDto**](HistoricDecisionOutputInstanceDto.md) | The list of decision output values. **Only exists** if &#x60;includeOutputs&#x60; was set to &#x60;true&#x60; in the query. | [optional] 
**CollectResultValue** | Pointer to **NullableFloat64** | The result of the collect aggregation of the decision result if used. &#x60;null&#x60; if no aggregation was used. | [optional] 
**RootDecisionInstanceId** | Pointer to **NullableString** | The decision instance id of the evaluated root decision. Can be &#x60;null&#x60; if this instance is the root decision instance of the evaluation. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the evaluation of this decision. Can be &#x60;null&#x60; if this decision instance is not evaluated as part of a BPMN process. | [optional] 
**DecisionRequirementsDefinitionId** | Pointer to **NullableString** | The id of the decision requirements definition that this decision instance belongs to. | [optional] 
**DecisionRequirementsDefinitionKey** | Pointer to **NullableString** | The key of the decision requirements definition that this decision instance belongs to. | [optional] 

## Methods

### NewHistoricDecisionInstanceDto

`func NewHistoricDecisionInstanceDto() *HistoricDecisionInstanceDto`

NewHistoricDecisionInstanceDto instantiates a new HistoricDecisionInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricDecisionInstanceDtoWithDefaults

`func NewHistoricDecisionInstanceDtoWithDefaults() *HistoricDecisionInstanceDto`

NewHistoricDecisionInstanceDtoWithDefaults instantiates a new HistoricDecisionInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricDecisionInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricDecisionInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricDecisionInstanceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricDecisionInstanceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricDecisionInstanceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricDecisionInstanceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDecisionDefinitionId

`func (o *HistoricDecisionInstanceDto) GetDecisionDefinitionId() string`

GetDecisionDefinitionId returns the DecisionDefinitionId field if non-nil, zero value otherwise.

### GetDecisionDefinitionIdOk

`func (o *HistoricDecisionInstanceDto) GetDecisionDefinitionIdOk() (*string, bool)`

GetDecisionDefinitionIdOk returns a tuple with the DecisionDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionId

`func (o *HistoricDecisionInstanceDto) SetDecisionDefinitionId(v string)`

SetDecisionDefinitionId sets DecisionDefinitionId field to given value.

### HasDecisionDefinitionId

`func (o *HistoricDecisionInstanceDto) HasDecisionDefinitionId() bool`

HasDecisionDefinitionId returns a boolean if a field has been set.

### SetDecisionDefinitionIdNil

`func (o *HistoricDecisionInstanceDto) SetDecisionDefinitionIdNil(b bool)`

 SetDecisionDefinitionIdNil sets the value for DecisionDefinitionId to be an explicit nil

### UnsetDecisionDefinitionId
`func (o *HistoricDecisionInstanceDto) UnsetDecisionDefinitionId()`

UnsetDecisionDefinitionId ensures that no value is present for DecisionDefinitionId, not even an explicit nil
### GetDecisionDefinitionKey

`func (o *HistoricDecisionInstanceDto) GetDecisionDefinitionKey() string`

GetDecisionDefinitionKey returns the DecisionDefinitionKey field if non-nil, zero value otherwise.

### GetDecisionDefinitionKeyOk

`func (o *HistoricDecisionInstanceDto) GetDecisionDefinitionKeyOk() (*string, bool)`

GetDecisionDefinitionKeyOk returns a tuple with the DecisionDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionKey

`func (o *HistoricDecisionInstanceDto) SetDecisionDefinitionKey(v string)`

SetDecisionDefinitionKey sets DecisionDefinitionKey field to given value.

### HasDecisionDefinitionKey

`func (o *HistoricDecisionInstanceDto) HasDecisionDefinitionKey() bool`

HasDecisionDefinitionKey returns a boolean if a field has been set.

### SetDecisionDefinitionKeyNil

`func (o *HistoricDecisionInstanceDto) SetDecisionDefinitionKeyNil(b bool)`

 SetDecisionDefinitionKeyNil sets the value for DecisionDefinitionKey to be an explicit nil

### UnsetDecisionDefinitionKey
`func (o *HistoricDecisionInstanceDto) UnsetDecisionDefinitionKey()`

UnsetDecisionDefinitionKey ensures that no value is present for DecisionDefinitionKey, not even an explicit nil
### GetDecisionDefinitionName

`func (o *HistoricDecisionInstanceDto) GetDecisionDefinitionName() string`

GetDecisionDefinitionName returns the DecisionDefinitionName field if non-nil, zero value otherwise.

### GetDecisionDefinitionNameOk

`func (o *HistoricDecisionInstanceDto) GetDecisionDefinitionNameOk() (*string, bool)`

GetDecisionDefinitionNameOk returns a tuple with the DecisionDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionName

`func (o *HistoricDecisionInstanceDto) SetDecisionDefinitionName(v string)`

SetDecisionDefinitionName sets DecisionDefinitionName field to given value.

### HasDecisionDefinitionName

`func (o *HistoricDecisionInstanceDto) HasDecisionDefinitionName() bool`

HasDecisionDefinitionName returns a boolean if a field has been set.

### SetDecisionDefinitionNameNil

`func (o *HistoricDecisionInstanceDto) SetDecisionDefinitionNameNil(b bool)`

 SetDecisionDefinitionNameNil sets the value for DecisionDefinitionName to be an explicit nil

### UnsetDecisionDefinitionName
`func (o *HistoricDecisionInstanceDto) UnsetDecisionDefinitionName()`

UnsetDecisionDefinitionName ensures that no value is present for DecisionDefinitionName, not even an explicit nil
### GetEvaluationTime

`func (o *HistoricDecisionInstanceDto) GetEvaluationTime() time.Time`

GetEvaluationTime returns the EvaluationTime field if non-nil, zero value otherwise.

### GetEvaluationTimeOk

`func (o *HistoricDecisionInstanceDto) GetEvaluationTimeOk() (*time.Time, bool)`

GetEvaluationTimeOk returns a tuple with the EvaluationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTime

`func (o *HistoricDecisionInstanceDto) SetEvaluationTime(v time.Time)`

SetEvaluationTime sets EvaluationTime field to given value.

### HasEvaluationTime

`func (o *HistoricDecisionInstanceDto) HasEvaluationTime() bool`

HasEvaluationTime returns a boolean if a field has been set.

### SetEvaluationTimeNil

`func (o *HistoricDecisionInstanceDto) SetEvaluationTimeNil(b bool)`

 SetEvaluationTimeNil sets the value for EvaluationTime to be an explicit nil

### UnsetEvaluationTime
`func (o *HistoricDecisionInstanceDto) UnsetEvaluationTime()`

UnsetEvaluationTime ensures that no value is present for EvaluationTime, not even an explicit nil
### GetRemovalTime

`func (o *HistoricDecisionInstanceDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricDecisionInstanceDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricDecisionInstanceDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricDecisionInstanceDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricDecisionInstanceDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricDecisionInstanceDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricDecisionInstanceDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricDecisionInstanceDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricDecisionInstanceDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricDecisionInstanceDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricDecisionInstanceDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricDecisionInstanceDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricDecisionInstanceDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricDecisionInstanceDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricDecisionInstanceDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricDecisionInstanceDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricDecisionInstanceDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricDecisionInstanceDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricDecisionInstanceDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricDecisionInstanceDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricDecisionInstanceDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricDecisionInstanceDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricDecisionInstanceDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricDecisionInstanceDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetCaseDefinitionId

`func (o *HistoricDecisionInstanceDto) GetCaseDefinitionId() string`

GetCaseDefinitionId returns the CaseDefinitionId field if non-nil, zero value otherwise.

### GetCaseDefinitionIdOk

`func (o *HistoricDecisionInstanceDto) GetCaseDefinitionIdOk() (*string, bool)`

GetCaseDefinitionIdOk returns a tuple with the CaseDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionId

`func (o *HistoricDecisionInstanceDto) SetCaseDefinitionId(v string)`

SetCaseDefinitionId sets CaseDefinitionId field to given value.

### HasCaseDefinitionId

`func (o *HistoricDecisionInstanceDto) HasCaseDefinitionId() bool`

HasCaseDefinitionId returns a boolean if a field has been set.

### SetCaseDefinitionIdNil

`func (o *HistoricDecisionInstanceDto) SetCaseDefinitionIdNil(b bool)`

 SetCaseDefinitionIdNil sets the value for CaseDefinitionId to be an explicit nil

### UnsetCaseDefinitionId
`func (o *HistoricDecisionInstanceDto) UnsetCaseDefinitionId()`

UnsetCaseDefinitionId ensures that no value is present for CaseDefinitionId, not even an explicit nil
### GetCaseDefinitionKey

`func (o *HistoricDecisionInstanceDto) GetCaseDefinitionKey() string`

GetCaseDefinitionKey returns the CaseDefinitionKey field if non-nil, zero value otherwise.

### GetCaseDefinitionKeyOk

`func (o *HistoricDecisionInstanceDto) GetCaseDefinitionKeyOk() (*string, bool)`

GetCaseDefinitionKeyOk returns a tuple with the CaseDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionKey

`func (o *HistoricDecisionInstanceDto) SetCaseDefinitionKey(v string)`

SetCaseDefinitionKey sets CaseDefinitionKey field to given value.

### HasCaseDefinitionKey

`func (o *HistoricDecisionInstanceDto) HasCaseDefinitionKey() bool`

HasCaseDefinitionKey returns a boolean if a field has been set.

### SetCaseDefinitionKeyNil

`func (o *HistoricDecisionInstanceDto) SetCaseDefinitionKeyNil(b bool)`

 SetCaseDefinitionKeyNil sets the value for CaseDefinitionKey to be an explicit nil

### UnsetCaseDefinitionKey
`func (o *HistoricDecisionInstanceDto) UnsetCaseDefinitionKey()`

UnsetCaseDefinitionKey ensures that no value is present for CaseDefinitionKey, not even an explicit nil
### GetCaseInstanceId

`func (o *HistoricDecisionInstanceDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *HistoricDecisionInstanceDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *HistoricDecisionInstanceDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *HistoricDecisionInstanceDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *HistoricDecisionInstanceDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *HistoricDecisionInstanceDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetActivityId

`func (o *HistoricDecisionInstanceDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *HistoricDecisionInstanceDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *HistoricDecisionInstanceDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *HistoricDecisionInstanceDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *HistoricDecisionInstanceDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *HistoricDecisionInstanceDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetActivityInstanceId

`func (o *HistoricDecisionInstanceDto) GetActivityInstanceId() string`

GetActivityInstanceId returns the ActivityInstanceId field if non-nil, zero value otherwise.

### GetActivityInstanceIdOk

`func (o *HistoricDecisionInstanceDto) GetActivityInstanceIdOk() (*string, bool)`

GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceId

`func (o *HistoricDecisionInstanceDto) SetActivityInstanceId(v string)`

SetActivityInstanceId sets ActivityInstanceId field to given value.

### HasActivityInstanceId

`func (o *HistoricDecisionInstanceDto) HasActivityInstanceId() bool`

HasActivityInstanceId returns a boolean if a field has been set.

### SetActivityInstanceIdNil

`func (o *HistoricDecisionInstanceDto) SetActivityInstanceIdNil(b bool)`

 SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil

### UnsetActivityInstanceId
`func (o *HistoricDecisionInstanceDto) UnsetActivityInstanceId()`

UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
### GetTenantId

`func (o *HistoricDecisionInstanceDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricDecisionInstanceDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricDecisionInstanceDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricDecisionInstanceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricDecisionInstanceDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricDecisionInstanceDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUserId

`func (o *HistoricDecisionInstanceDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HistoricDecisionInstanceDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HistoricDecisionInstanceDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *HistoricDecisionInstanceDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *HistoricDecisionInstanceDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *HistoricDecisionInstanceDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetInputs

`func (o *HistoricDecisionInstanceDto) GetInputs() []HistoricDecisionInputInstanceDto`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *HistoricDecisionInstanceDto) GetInputsOk() (*[]HistoricDecisionInputInstanceDto, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *HistoricDecisionInstanceDto) SetInputs(v []HistoricDecisionInputInstanceDto)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *HistoricDecisionInstanceDto) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### SetInputsNil

`func (o *HistoricDecisionInstanceDto) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *HistoricDecisionInstanceDto) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil
### GetOutputs

`func (o *HistoricDecisionInstanceDto) GetOutputs() []HistoricDecisionOutputInstanceDto`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *HistoricDecisionInstanceDto) GetOutputsOk() (*[]HistoricDecisionOutputInstanceDto, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *HistoricDecisionInstanceDto) SetOutputs(v []HistoricDecisionOutputInstanceDto)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *HistoricDecisionInstanceDto) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *HistoricDecisionInstanceDto) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *HistoricDecisionInstanceDto) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetCollectResultValue

`func (o *HistoricDecisionInstanceDto) GetCollectResultValue() float64`

GetCollectResultValue returns the CollectResultValue field if non-nil, zero value otherwise.

### GetCollectResultValueOk

`func (o *HistoricDecisionInstanceDto) GetCollectResultValueOk() (*float64, bool)`

GetCollectResultValueOk returns a tuple with the CollectResultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectResultValue

`func (o *HistoricDecisionInstanceDto) SetCollectResultValue(v float64)`

SetCollectResultValue sets CollectResultValue field to given value.

### HasCollectResultValue

`func (o *HistoricDecisionInstanceDto) HasCollectResultValue() bool`

HasCollectResultValue returns a boolean if a field has been set.

### SetCollectResultValueNil

`func (o *HistoricDecisionInstanceDto) SetCollectResultValueNil(b bool)`

 SetCollectResultValueNil sets the value for CollectResultValue to be an explicit nil

### UnsetCollectResultValue
`func (o *HistoricDecisionInstanceDto) UnsetCollectResultValue()`

UnsetCollectResultValue ensures that no value is present for CollectResultValue, not even an explicit nil
### GetRootDecisionInstanceId

`func (o *HistoricDecisionInstanceDto) GetRootDecisionInstanceId() string`

GetRootDecisionInstanceId returns the RootDecisionInstanceId field if non-nil, zero value otherwise.

### GetRootDecisionInstanceIdOk

`func (o *HistoricDecisionInstanceDto) GetRootDecisionInstanceIdOk() (*string, bool)`

GetRootDecisionInstanceIdOk returns a tuple with the RootDecisionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDecisionInstanceId

`func (o *HistoricDecisionInstanceDto) SetRootDecisionInstanceId(v string)`

SetRootDecisionInstanceId sets RootDecisionInstanceId field to given value.

### HasRootDecisionInstanceId

`func (o *HistoricDecisionInstanceDto) HasRootDecisionInstanceId() bool`

HasRootDecisionInstanceId returns a boolean if a field has been set.

### SetRootDecisionInstanceIdNil

`func (o *HistoricDecisionInstanceDto) SetRootDecisionInstanceIdNil(b bool)`

 SetRootDecisionInstanceIdNil sets the value for RootDecisionInstanceId to be an explicit nil

### UnsetRootDecisionInstanceId
`func (o *HistoricDecisionInstanceDto) UnsetRootDecisionInstanceId()`

UnsetRootDecisionInstanceId ensures that no value is present for RootDecisionInstanceId, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricDecisionInstanceDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricDecisionInstanceDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricDecisionInstanceDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricDecisionInstanceDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricDecisionInstanceDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricDecisionInstanceDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil
### GetDecisionRequirementsDefinitionId

`func (o *HistoricDecisionInstanceDto) GetDecisionRequirementsDefinitionId() string`

GetDecisionRequirementsDefinitionId returns the DecisionRequirementsDefinitionId field if non-nil, zero value otherwise.

### GetDecisionRequirementsDefinitionIdOk

`func (o *HistoricDecisionInstanceDto) GetDecisionRequirementsDefinitionIdOk() (*string, bool)`

GetDecisionRequirementsDefinitionIdOk returns a tuple with the DecisionRequirementsDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionRequirementsDefinitionId

`func (o *HistoricDecisionInstanceDto) SetDecisionRequirementsDefinitionId(v string)`

SetDecisionRequirementsDefinitionId sets DecisionRequirementsDefinitionId field to given value.

### HasDecisionRequirementsDefinitionId

`func (o *HistoricDecisionInstanceDto) HasDecisionRequirementsDefinitionId() bool`

HasDecisionRequirementsDefinitionId returns a boolean if a field has been set.

### SetDecisionRequirementsDefinitionIdNil

`func (o *HistoricDecisionInstanceDto) SetDecisionRequirementsDefinitionIdNil(b bool)`

 SetDecisionRequirementsDefinitionIdNil sets the value for DecisionRequirementsDefinitionId to be an explicit nil

### UnsetDecisionRequirementsDefinitionId
`func (o *HistoricDecisionInstanceDto) UnsetDecisionRequirementsDefinitionId()`

UnsetDecisionRequirementsDefinitionId ensures that no value is present for DecisionRequirementsDefinitionId, not even an explicit nil
### GetDecisionRequirementsDefinitionKey

`func (o *HistoricDecisionInstanceDto) GetDecisionRequirementsDefinitionKey() string`

GetDecisionRequirementsDefinitionKey returns the DecisionRequirementsDefinitionKey field if non-nil, zero value otherwise.

### GetDecisionRequirementsDefinitionKeyOk

`func (o *HistoricDecisionInstanceDto) GetDecisionRequirementsDefinitionKeyOk() (*string, bool)`

GetDecisionRequirementsDefinitionKeyOk returns a tuple with the DecisionRequirementsDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionRequirementsDefinitionKey

`func (o *HistoricDecisionInstanceDto) SetDecisionRequirementsDefinitionKey(v string)`

SetDecisionRequirementsDefinitionKey sets DecisionRequirementsDefinitionKey field to given value.

### HasDecisionRequirementsDefinitionKey

`func (o *HistoricDecisionInstanceDto) HasDecisionRequirementsDefinitionKey() bool`

HasDecisionRequirementsDefinitionKey returns a boolean if a field has been set.

### SetDecisionRequirementsDefinitionKeyNil

`func (o *HistoricDecisionInstanceDto) SetDecisionRequirementsDefinitionKeyNil(b bool)`

 SetDecisionRequirementsDefinitionKeyNil sets the value for DecisionRequirementsDefinitionKey to be an explicit nil

### UnsetDecisionRequirementsDefinitionKey
`func (o *HistoricDecisionInstanceDto) UnsetDecisionRequirementsDefinitionKey()`

UnsetDecisionRequirementsDefinitionKey ensures that no value is present for DecisionRequirementsDefinitionKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


