# HistoricDecisionInstanceQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecisionInstanceId** | Pointer to **NullableString** | Filter by decision instance id. | [optional] 
**DecisionInstanceIdIn** | Pointer to **[]string** | Filter by decision instance ids. Must be a comma-separated list of decision instance ids. | [optional] 
**DecisionDefinitionId** | Pointer to **NullableString** | Filter by the decision definition the instances belongs to. | [optional] 
**DecisionDefinitionIdIn** | Pointer to **[]string** | Filter by the decision definitions the instances belongs to. Must be a comma-separated list of decision definition ids. | [optional] 
**DecisionDefinitionKey** | Pointer to **NullableString** | Filter by the key of the decision definition the instances belongs to. | [optional] 
**DecisionDefinitionKeyIn** | Pointer to **[]string** | Filter by the keys of the decision definition the instances belongs to. Must be a comma- separated list of decision definition keys. | [optional] 
**DecisionDefinitionName** | Pointer to **NullableString** | Filter by the name of the decision definition the instances belongs to. | [optional] 
**DecisionDefinitionNameLike** | Pointer to **NullableString** | Filter by the name of the decision definition the instances belongs to, that the parameter is a substring of. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter by the process definition the instances belongs to. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Filter by the key of the process definition the instances belongs to. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | Filter by the process instance the instances belongs to. | [optional] 
**CaseDefinitionId** | Pointer to **NullableString** | Filter by the case definition the instances belongs to. | [optional] 
**CaseDefinitionKey** | Pointer to **NullableString** | Filter by the key of the case definition the instances belongs to. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | Filter by the case instance the instances belongs to. | [optional] 
**ActivityIdIn** | Pointer to **[]string** | Filter by the activity ids the instances belongs to. Must be a comma-separated list of acitvity ids. | [optional] 
**ActivityInstanceIdIn** | Pointer to **[]string** | Filter by the activity instance ids the instances belongs to. Must be a comma-separated list of acitvity instance ids. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Filter by a comma-separated list of tenant ids. A historic decision instance must have one of the given tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include historic decision instances that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**EvaluatedBefore** | Pointer to **NullableTime** | Restrict to instances that were evaluated before the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM- dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**EvaluatedAfter** | Pointer to **NullableTime** | Restrict to instances that were evaluated after the given date. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM- dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**UserId** | Pointer to **NullableString** | Restrict to instances that were evaluated by the given user. | [optional] 
**RootDecisionInstanceId** | Pointer to **NullableString** | Restrict to instances that have a given root decision instance id. This also includes the decision instance with the given id. | [optional] 
**RootDecisionInstancesOnly** | Pointer to **NullableBool** | Restrict to instances those are the root decision instance of an evaluation. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**DecisionRequirementsDefinitionId** | Pointer to **NullableString** | Filter by the decision requirements definition the instances belongs to. | [optional] 
**DecisionRequirementsDefinitionKey** | Pointer to **NullableString** | Filter by the key of the decision requirements definition the instances belongs to. | [optional] 
**IncludeInputs** | Pointer to **NullableBool** | Include input values in the result. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**IncludeOutputs** | Pointer to **NullableBool** | Include output values in the result. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**DisableBinaryFetching** | Pointer to **NullableBool** | Disables fetching of byte array input and output values. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**DisableCustomObjectDeserialization** | Pointer to **NullableBool** | Disables deserialization of input and output values that are custom objects. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 

## Methods

### NewHistoricDecisionInstanceQueryDto

`func NewHistoricDecisionInstanceQueryDto() *HistoricDecisionInstanceQueryDto`

NewHistoricDecisionInstanceQueryDto instantiates a new HistoricDecisionInstanceQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricDecisionInstanceQueryDtoWithDefaults

`func NewHistoricDecisionInstanceQueryDtoWithDefaults() *HistoricDecisionInstanceQueryDto`

NewHistoricDecisionInstanceQueryDtoWithDefaults instantiates a new HistoricDecisionInstanceQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecisionInstanceId

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionInstanceId() string`

GetDecisionInstanceId returns the DecisionInstanceId field if non-nil, zero value otherwise.

### GetDecisionInstanceIdOk

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionInstanceIdOk() (*string, bool)`

GetDecisionInstanceIdOk returns a tuple with the DecisionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionInstanceId

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionInstanceId(v string)`

SetDecisionInstanceId sets DecisionInstanceId field to given value.

### HasDecisionInstanceId

`func (o *HistoricDecisionInstanceQueryDto) HasDecisionInstanceId() bool`

HasDecisionInstanceId returns a boolean if a field has been set.

### SetDecisionInstanceIdNil

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionInstanceIdNil(b bool)`

 SetDecisionInstanceIdNil sets the value for DecisionInstanceId to be an explicit nil

### UnsetDecisionInstanceId
`func (o *HistoricDecisionInstanceQueryDto) UnsetDecisionInstanceId()`

UnsetDecisionInstanceId ensures that no value is present for DecisionInstanceId, not even an explicit nil
### GetDecisionInstanceIdIn

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionInstanceIdIn() []string`

GetDecisionInstanceIdIn returns the DecisionInstanceIdIn field if non-nil, zero value otherwise.

### GetDecisionInstanceIdInOk

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionInstanceIdInOk() (*[]string, bool)`

GetDecisionInstanceIdInOk returns a tuple with the DecisionInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionInstanceIdIn

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionInstanceIdIn(v []string)`

SetDecisionInstanceIdIn sets DecisionInstanceIdIn field to given value.

### HasDecisionInstanceIdIn

`func (o *HistoricDecisionInstanceQueryDto) HasDecisionInstanceIdIn() bool`

HasDecisionInstanceIdIn returns a boolean if a field has been set.

### SetDecisionInstanceIdInNil

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionInstanceIdInNil(b bool)`

 SetDecisionInstanceIdInNil sets the value for DecisionInstanceIdIn to be an explicit nil

### UnsetDecisionInstanceIdIn
`func (o *HistoricDecisionInstanceQueryDto) UnsetDecisionInstanceIdIn()`

UnsetDecisionInstanceIdIn ensures that no value is present for DecisionInstanceIdIn, not even an explicit nil
### GetDecisionDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionId() string`

GetDecisionDefinitionId returns the DecisionDefinitionId field if non-nil, zero value otherwise.

### GetDecisionDefinitionIdOk

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionIdOk() (*string, bool)`

GetDecisionDefinitionIdOk returns a tuple with the DecisionDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionId(v string)`

SetDecisionDefinitionId sets DecisionDefinitionId field to given value.

### HasDecisionDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) HasDecisionDefinitionId() bool`

HasDecisionDefinitionId returns a boolean if a field has been set.

### SetDecisionDefinitionIdNil

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionIdNil(b bool)`

 SetDecisionDefinitionIdNil sets the value for DecisionDefinitionId to be an explicit nil

### UnsetDecisionDefinitionId
`func (o *HistoricDecisionInstanceQueryDto) UnsetDecisionDefinitionId()`

UnsetDecisionDefinitionId ensures that no value is present for DecisionDefinitionId, not even an explicit nil
### GetDecisionDefinitionIdIn

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionIdIn() []string`

GetDecisionDefinitionIdIn returns the DecisionDefinitionIdIn field if non-nil, zero value otherwise.

### GetDecisionDefinitionIdInOk

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionIdInOk() (*[]string, bool)`

GetDecisionDefinitionIdInOk returns a tuple with the DecisionDefinitionIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionIdIn

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionIdIn(v []string)`

SetDecisionDefinitionIdIn sets DecisionDefinitionIdIn field to given value.

### HasDecisionDefinitionIdIn

`func (o *HistoricDecisionInstanceQueryDto) HasDecisionDefinitionIdIn() bool`

HasDecisionDefinitionIdIn returns a boolean if a field has been set.

### SetDecisionDefinitionIdInNil

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionIdInNil(b bool)`

 SetDecisionDefinitionIdInNil sets the value for DecisionDefinitionIdIn to be an explicit nil

### UnsetDecisionDefinitionIdIn
`func (o *HistoricDecisionInstanceQueryDto) UnsetDecisionDefinitionIdIn()`

UnsetDecisionDefinitionIdIn ensures that no value is present for DecisionDefinitionIdIn, not even an explicit nil
### GetDecisionDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionKey() string`

GetDecisionDefinitionKey returns the DecisionDefinitionKey field if non-nil, zero value otherwise.

### GetDecisionDefinitionKeyOk

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionKeyOk() (*string, bool)`

GetDecisionDefinitionKeyOk returns a tuple with the DecisionDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionKey(v string)`

SetDecisionDefinitionKey sets DecisionDefinitionKey field to given value.

### HasDecisionDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) HasDecisionDefinitionKey() bool`

HasDecisionDefinitionKey returns a boolean if a field has been set.

### SetDecisionDefinitionKeyNil

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionKeyNil(b bool)`

 SetDecisionDefinitionKeyNil sets the value for DecisionDefinitionKey to be an explicit nil

### UnsetDecisionDefinitionKey
`func (o *HistoricDecisionInstanceQueryDto) UnsetDecisionDefinitionKey()`

UnsetDecisionDefinitionKey ensures that no value is present for DecisionDefinitionKey, not even an explicit nil
### GetDecisionDefinitionKeyIn

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionKeyIn() []string`

GetDecisionDefinitionKeyIn returns the DecisionDefinitionKeyIn field if non-nil, zero value otherwise.

### GetDecisionDefinitionKeyInOk

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionKeyInOk() (*[]string, bool)`

GetDecisionDefinitionKeyInOk returns a tuple with the DecisionDefinitionKeyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionKeyIn

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionKeyIn(v []string)`

SetDecisionDefinitionKeyIn sets DecisionDefinitionKeyIn field to given value.

### HasDecisionDefinitionKeyIn

`func (o *HistoricDecisionInstanceQueryDto) HasDecisionDefinitionKeyIn() bool`

HasDecisionDefinitionKeyIn returns a boolean if a field has been set.

### SetDecisionDefinitionKeyInNil

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionKeyInNil(b bool)`

 SetDecisionDefinitionKeyInNil sets the value for DecisionDefinitionKeyIn to be an explicit nil

### UnsetDecisionDefinitionKeyIn
`func (o *HistoricDecisionInstanceQueryDto) UnsetDecisionDefinitionKeyIn()`

UnsetDecisionDefinitionKeyIn ensures that no value is present for DecisionDefinitionKeyIn, not even an explicit nil
### GetDecisionDefinitionName

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionName() string`

GetDecisionDefinitionName returns the DecisionDefinitionName field if non-nil, zero value otherwise.

### GetDecisionDefinitionNameOk

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionNameOk() (*string, bool)`

GetDecisionDefinitionNameOk returns a tuple with the DecisionDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionName

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionName(v string)`

SetDecisionDefinitionName sets DecisionDefinitionName field to given value.

### HasDecisionDefinitionName

`func (o *HistoricDecisionInstanceQueryDto) HasDecisionDefinitionName() bool`

HasDecisionDefinitionName returns a boolean if a field has been set.

### SetDecisionDefinitionNameNil

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionNameNil(b bool)`

 SetDecisionDefinitionNameNil sets the value for DecisionDefinitionName to be an explicit nil

### UnsetDecisionDefinitionName
`func (o *HistoricDecisionInstanceQueryDto) UnsetDecisionDefinitionName()`

UnsetDecisionDefinitionName ensures that no value is present for DecisionDefinitionName, not even an explicit nil
### GetDecisionDefinitionNameLike

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionNameLike() string`

GetDecisionDefinitionNameLike returns the DecisionDefinitionNameLike field if non-nil, zero value otherwise.

### GetDecisionDefinitionNameLikeOk

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionDefinitionNameLikeOk() (*string, bool)`

GetDecisionDefinitionNameLikeOk returns a tuple with the DecisionDefinitionNameLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionNameLike

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionNameLike(v string)`

SetDecisionDefinitionNameLike sets DecisionDefinitionNameLike field to given value.

### HasDecisionDefinitionNameLike

`func (o *HistoricDecisionInstanceQueryDto) HasDecisionDefinitionNameLike() bool`

HasDecisionDefinitionNameLike returns a boolean if a field has been set.

### SetDecisionDefinitionNameLikeNil

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionDefinitionNameLikeNil(b bool)`

 SetDecisionDefinitionNameLikeNil sets the value for DecisionDefinitionNameLike to be an explicit nil

### UnsetDecisionDefinitionNameLike
`func (o *HistoricDecisionInstanceQueryDto) UnsetDecisionDefinitionNameLike()`

UnsetDecisionDefinitionNameLike ensures that no value is present for DecisionDefinitionNameLike, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricDecisionInstanceQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricDecisionInstanceQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricDecisionInstanceQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricDecisionInstanceQueryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricDecisionInstanceQueryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricDecisionInstanceQueryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricDecisionInstanceQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricDecisionInstanceQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricDecisionInstanceQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricDecisionInstanceQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricDecisionInstanceQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricDecisionInstanceQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetCaseDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) GetCaseDefinitionId() string`

GetCaseDefinitionId returns the CaseDefinitionId field if non-nil, zero value otherwise.

### GetCaseDefinitionIdOk

`func (o *HistoricDecisionInstanceQueryDto) GetCaseDefinitionIdOk() (*string, bool)`

GetCaseDefinitionIdOk returns a tuple with the CaseDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) SetCaseDefinitionId(v string)`

SetCaseDefinitionId sets CaseDefinitionId field to given value.

### HasCaseDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) HasCaseDefinitionId() bool`

HasCaseDefinitionId returns a boolean if a field has been set.

### SetCaseDefinitionIdNil

`func (o *HistoricDecisionInstanceQueryDto) SetCaseDefinitionIdNil(b bool)`

 SetCaseDefinitionIdNil sets the value for CaseDefinitionId to be an explicit nil

### UnsetCaseDefinitionId
`func (o *HistoricDecisionInstanceQueryDto) UnsetCaseDefinitionId()`

UnsetCaseDefinitionId ensures that no value is present for CaseDefinitionId, not even an explicit nil
### GetCaseDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) GetCaseDefinitionKey() string`

GetCaseDefinitionKey returns the CaseDefinitionKey field if non-nil, zero value otherwise.

### GetCaseDefinitionKeyOk

`func (o *HistoricDecisionInstanceQueryDto) GetCaseDefinitionKeyOk() (*string, bool)`

GetCaseDefinitionKeyOk returns a tuple with the CaseDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) SetCaseDefinitionKey(v string)`

SetCaseDefinitionKey sets CaseDefinitionKey field to given value.

### HasCaseDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) HasCaseDefinitionKey() bool`

HasCaseDefinitionKey returns a boolean if a field has been set.

### SetCaseDefinitionKeyNil

`func (o *HistoricDecisionInstanceQueryDto) SetCaseDefinitionKeyNil(b bool)`

 SetCaseDefinitionKeyNil sets the value for CaseDefinitionKey to be an explicit nil

### UnsetCaseDefinitionKey
`func (o *HistoricDecisionInstanceQueryDto) UnsetCaseDefinitionKey()`

UnsetCaseDefinitionKey ensures that no value is present for CaseDefinitionKey, not even an explicit nil
### GetCaseInstanceId

`func (o *HistoricDecisionInstanceQueryDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *HistoricDecisionInstanceQueryDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *HistoricDecisionInstanceQueryDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *HistoricDecisionInstanceQueryDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *HistoricDecisionInstanceQueryDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *HistoricDecisionInstanceQueryDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetActivityIdIn

`func (o *HistoricDecisionInstanceQueryDto) GetActivityIdIn() []string`

GetActivityIdIn returns the ActivityIdIn field if non-nil, zero value otherwise.

### GetActivityIdInOk

`func (o *HistoricDecisionInstanceQueryDto) GetActivityIdInOk() (*[]string, bool)`

GetActivityIdInOk returns a tuple with the ActivityIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIdIn

`func (o *HistoricDecisionInstanceQueryDto) SetActivityIdIn(v []string)`

SetActivityIdIn sets ActivityIdIn field to given value.

### HasActivityIdIn

`func (o *HistoricDecisionInstanceQueryDto) HasActivityIdIn() bool`

HasActivityIdIn returns a boolean if a field has been set.

### SetActivityIdInNil

`func (o *HistoricDecisionInstanceQueryDto) SetActivityIdInNil(b bool)`

 SetActivityIdInNil sets the value for ActivityIdIn to be an explicit nil

### UnsetActivityIdIn
`func (o *HistoricDecisionInstanceQueryDto) UnsetActivityIdIn()`

UnsetActivityIdIn ensures that no value is present for ActivityIdIn, not even an explicit nil
### GetActivityInstanceIdIn

`func (o *HistoricDecisionInstanceQueryDto) GetActivityInstanceIdIn() []string`

GetActivityInstanceIdIn returns the ActivityInstanceIdIn field if non-nil, zero value otherwise.

### GetActivityInstanceIdInOk

`func (o *HistoricDecisionInstanceQueryDto) GetActivityInstanceIdInOk() (*[]string, bool)`

GetActivityInstanceIdInOk returns a tuple with the ActivityInstanceIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceIdIn

`func (o *HistoricDecisionInstanceQueryDto) SetActivityInstanceIdIn(v []string)`

SetActivityInstanceIdIn sets ActivityInstanceIdIn field to given value.

### HasActivityInstanceIdIn

`func (o *HistoricDecisionInstanceQueryDto) HasActivityInstanceIdIn() bool`

HasActivityInstanceIdIn returns a boolean if a field has been set.

### SetActivityInstanceIdInNil

`func (o *HistoricDecisionInstanceQueryDto) SetActivityInstanceIdInNil(b bool)`

 SetActivityInstanceIdInNil sets the value for ActivityInstanceIdIn to be an explicit nil

### UnsetActivityInstanceIdIn
`func (o *HistoricDecisionInstanceQueryDto) UnsetActivityInstanceIdIn()`

UnsetActivityInstanceIdIn ensures that no value is present for ActivityInstanceIdIn, not even an explicit nil
### GetTenantIdIn

`func (o *HistoricDecisionInstanceQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *HistoricDecisionInstanceQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *HistoricDecisionInstanceQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *HistoricDecisionInstanceQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *HistoricDecisionInstanceQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *HistoricDecisionInstanceQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *HistoricDecisionInstanceQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *HistoricDecisionInstanceQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *HistoricDecisionInstanceQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *HistoricDecisionInstanceQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *HistoricDecisionInstanceQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *HistoricDecisionInstanceQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetEvaluatedBefore

`func (o *HistoricDecisionInstanceQueryDto) GetEvaluatedBefore() time.Time`

GetEvaluatedBefore returns the EvaluatedBefore field if non-nil, zero value otherwise.

### GetEvaluatedBeforeOk

`func (o *HistoricDecisionInstanceQueryDto) GetEvaluatedBeforeOk() (*time.Time, bool)`

GetEvaluatedBeforeOk returns a tuple with the EvaluatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedBefore

`func (o *HistoricDecisionInstanceQueryDto) SetEvaluatedBefore(v time.Time)`

SetEvaluatedBefore sets EvaluatedBefore field to given value.

### HasEvaluatedBefore

`func (o *HistoricDecisionInstanceQueryDto) HasEvaluatedBefore() bool`

HasEvaluatedBefore returns a boolean if a field has been set.

### SetEvaluatedBeforeNil

`func (o *HistoricDecisionInstanceQueryDto) SetEvaluatedBeforeNil(b bool)`

 SetEvaluatedBeforeNil sets the value for EvaluatedBefore to be an explicit nil

### UnsetEvaluatedBefore
`func (o *HistoricDecisionInstanceQueryDto) UnsetEvaluatedBefore()`

UnsetEvaluatedBefore ensures that no value is present for EvaluatedBefore, not even an explicit nil
### GetEvaluatedAfter

`func (o *HistoricDecisionInstanceQueryDto) GetEvaluatedAfter() time.Time`

GetEvaluatedAfter returns the EvaluatedAfter field if non-nil, zero value otherwise.

### GetEvaluatedAfterOk

`func (o *HistoricDecisionInstanceQueryDto) GetEvaluatedAfterOk() (*time.Time, bool)`

GetEvaluatedAfterOk returns a tuple with the EvaluatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedAfter

`func (o *HistoricDecisionInstanceQueryDto) SetEvaluatedAfter(v time.Time)`

SetEvaluatedAfter sets EvaluatedAfter field to given value.

### HasEvaluatedAfter

`func (o *HistoricDecisionInstanceQueryDto) HasEvaluatedAfter() bool`

HasEvaluatedAfter returns a boolean if a field has been set.

### SetEvaluatedAfterNil

`func (o *HistoricDecisionInstanceQueryDto) SetEvaluatedAfterNil(b bool)`

 SetEvaluatedAfterNil sets the value for EvaluatedAfter to be an explicit nil

### UnsetEvaluatedAfter
`func (o *HistoricDecisionInstanceQueryDto) UnsetEvaluatedAfter()`

UnsetEvaluatedAfter ensures that no value is present for EvaluatedAfter, not even an explicit nil
### GetUserId

`func (o *HistoricDecisionInstanceQueryDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HistoricDecisionInstanceQueryDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HistoricDecisionInstanceQueryDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *HistoricDecisionInstanceQueryDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *HistoricDecisionInstanceQueryDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *HistoricDecisionInstanceQueryDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetRootDecisionInstanceId

`func (o *HistoricDecisionInstanceQueryDto) GetRootDecisionInstanceId() string`

GetRootDecisionInstanceId returns the RootDecisionInstanceId field if non-nil, zero value otherwise.

### GetRootDecisionInstanceIdOk

`func (o *HistoricDecisionInstanceQueryDto) GetRootDecisionInstanceIdOk() (*string, bool)`

GetRootDecisionInstanceIdOk returns a tuple with the RootDecisionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDecisionInstanceId

`func (o *HistoricDecisionInstanceQueryDto) SetRootDecisionInstanceId(v string)`

SetRootDecisionInstanceId sets RootDecisionInstanceId field to given value.

### HasRootDecisionInstanceId

`func (o *HistoricDecisionInstanceQueryDto) HasRootDecisionInstanceId() bool`

HasRootDecisionInstanceId returns a boolean if a field has been set.

### SetRootDecisionInstanceIdNil

`func (o *HistoricDecisionInstanceQueryDto) SetRootDecisionInstanceIdNil(b bool)`

 SetRootDecisionInstanceIdNil sets the value for RootDecisionInstanceId to be an explicit nil

### UnsetRootDecisionInstanceId
`func (o *HistoricDecisionInstanceQueryDto) UnsetRootDecisionInstanceId()`

UnsetRootDecisionInstanceId ensures that no value is present for RootDecisionInstanceId, not even an explicit nil
### GetRootDecisionInstancesOnly

`func (o *HistoricDecisionInstanceQueryDto) GetRootDecisionInstancesOnly() bool`

GetRootDecisionInstancesOnly returns the RootDecisionInstancesOnly field if non-nil, zero value otherwise.

### GetRootDecisionInstancesOnlyOk

`func (o *HistoricDecisionInstanceQueryDto) GetRootDecisionInstancesOnlyOk() (*bool, bool)`

GetRootDecisionInstancesOnlyOk returns a tuple with the RootDecisionInstancesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDecisionInstancesOnly

`func (o *HistoricDecisionInstanceQueryDto) SetRootDecisionInstancesOnly(v bool)`

SetRootDecisionInstancesOnly sets RootDecisionInstancesOnly field to given value.

### HasRootDecisionInstancesOnly

`func (o *HistoricDecisionInstanceQueryDto) HasRootDecisionInstancesOnly() bool`

HasRootDecisionInstancesOnly returns a boolean if a field has been set.

### SetRootDecisionInstancesOnlyNil

`func (o *HistoricDecisionInstanceQueryDto) SetRootDecisionInstancesOnlyNil(b bool)`

 SetRootDecisionInstancesOnlyNil sets the value for RootDecisionInstancesOnly to be an explicit nil

### UnsetRootDecisionInstancesOnly
`func (o *HistoricDecisionInstanceQueryDto) UnsetRootDecisionInstancesOnly()`

UnsetRootDecisionInstancesOnly ensures that no value is present for RootDecisionInstancesOnly, not even an explicit nil
### GetDecisionRequirementsDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionRequirementsDefinitionId() string`

GetDecisionRequirementsDefinitionId returns the DecisionRequirementsDefinitionId field if non-nil, zero value otherwise.

### GetDecisionRequirementsDefinitionIdOk

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionRequirementsDefinitionIdOk() (*string, bool)`

GetDecisionRequirementsDefinitionIdOk returns a tuple with the DecisionRequirementsDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionRequirementsDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionRequirementsDefinitionId(v string)`

SetDecisionRequirementsDefinitionId sets DecisionRequirementsDefinitionId field to given value.

### HasDecisionRequirementsDefinitionId

`func (o *HistoricDecisionInstanceQueryDto) HasDecisionRequirementsDefinitionId() bool`

HasDecisionRequirementsDefinitionId returns a boolean if a field has been set.

### SetDecisionRequirementsDefinitionIdNil

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionRequirementsDefinitionIdNil(b bool)`

 SetDecisionRequirementsDefinitionIdNil sets the value for DecisionRequirementsDefinitionId to be an explicit nil

### UnsetDecisionRequirementsDefinitionId
`func (o *HistoricDecisionInstanceQueryDto) UnsetDecisionRequirementsDefinitionId()`

UnsetDecisionRequirementsDefinitionId ensures that no value is present for DecisionRequirementsDefinitionId, not even an explicit nil
### GetDecisionRequirementsDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionRequirementsDefinitionKey() string`

GetDecisionRequirementsDefinitionKey returns the DecisionRequirementsDefinitionKey field if non-nil, zero value otherwise.

### GetDecisionRequirementsDefinitionKeyOk

`func (o *HistoricDecisionInstanceQueryDto) GetDecisionRequirementsDefinitionKeyOk() (*string, bool)`

GetDecisionRequirementsDefinitionKeyOk returns a tuple with the DecisionRequirementsDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionRequirementsDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionRequirementsDefinitionKey(v string)`

SetDecisionRequirementsDefinitionKey sets DecisionRequirementsDefinitionKey field to given value.

### HasDecisionRequirementsDefinitionKey

`func (o *HistoricDecisionInstanceQueryDto) HasDecisionRequirementsDefinitionKey() bool`

HasDecisionRequirementsDefinitionKey returns a boolean if a field has been set.

### SetDecisionRequirementsDefinitionKeyNil

`func (o *HistoricDecisionInstanceQueryDto) SetDecisionRequirementsDefinitionKeyNil(b bool)`

 SetDecisionRequirementsDefinitionKeyNil sets the value for DecisionRequirementsDefinitionKey to be an explicit nil

### UnsetDecisionRequirementsDefinitionKey
`func (o *HistoricDecisionInstanceQueryDto) UnsetDecisionRequirementsDefinitionKey()`

UnsetDecisionRequirementsDefinitionKey ensures that no value is present for DecisionRequirementsDefinitionKey, not even an explicit nil
### GetIncludeInputs

`func (o *HistoricDecisionInstanceQueryDto) GetIncludeInputs() bool`

GetIncludeInputs returns the IncludeInputs field if non-nil, zero value otherwise.

### GetIncludeInputsOk

`func (o *HistoricDecisionInstanceQueryDto) GetIncludeInputsOk() (*bool, bool)`

GetIncludeInputsOk returns a tuple with the IncludeInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInputs

`func (o *HistoricDecisionInstanceQueryDto) SetIncludeInputs(v bool)`

SetIncludeInputs sets IncludeInputs field to given value.

### HasIncludeInputs

`func (o *HistoricDecisionInstanceQueryDto) HasIncludeInputs() bool`

HasIncludeInputs returns a boolean if a field has been set.

### SetIncludeInputsNil

`func (o *HistoricDecisionInstanceQueryDto) SetIncludeInputsNil(b bool)`

 SetIncludeInputsNil sets the value for IncludeInputs to be an explicit nil

### UnsetIncludeInputs
`func (o *HistoricDecisionInstanceQueryDto) UnsetIncludeInputs()`

UnsetIncludeInputs ensures that no value is present for IncludeInputs, not even an explicit nil
### GetIncludeOutputs

`func (o *HistoricDecisionInstanceQueryDto) GetIncludeOutputs() bool`

GetIncludeOutputs returns the IncludeOutputs field if non-nil, zero value otherwise.

### GetIncludeOutputsOk

`func (o *HistoricDecisionInstanceQueryDto) GetIncludeOutputsOk() (*bool, bool)`

GetIncludeOutputsOk returns a tuple with the IncludeOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOutputs

`func (o *HistoricDecisionInstanceQueryDto) SetIncludeOutputs(v bool)`

SetIncludeOutputs sets IncludeOutputs field to given value.

### HasIncludeOutputs

`func (o *HistoricDecisionInstanceQueryDto) HasIncludeOutputs() bool`

HasIncludeOutputs returns a boolean if a field has been set.

### SetIncludeOutputsNil

`func (o *HistoricDecisionInstanceQueryDto) SetIncludeOutputsNil(b bool)`

 SetIncludeOutputsNil sets the value for IncludeOutputs to be an explicit nil

### UnsetIncludeOutputs
`func (o *HistoricDecisionInstanceQueryDto) UnsetIncludeOutputs()`

UnsetIncludeOutputs ensures that no value is present for IncludeOutputs, not even an explicit nil
### GetDisableBinaryFetching

`func (o *HistoricDecisionInstanceQueryDto) GetDisableBinaryFetching() bool`

GetDisableBinaryFetching returns the DisableBinaryFetching field if non-nil, zero value otherwise.

### GetDisableBinaryFetchingOk

`func (o *HistoricDecisionInstanceQueryDto) GetDisableBinaryFetchingOk() (*bool, bool)`

GetDisableBinaryFetchingOk returns a tuple with the DisableBinaryFetching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBinaryFetching

`func (o *HistoricDecisionInstanceQueryDto) SetDisableBinaryFetching(v bool)`

SetDisableBinaryFetching sets DisableBinaryFetching field to given value.

### HasDisableBinaryFetching

`func (o *HistoricDecisionInstanceQueryDto) HasDisableBinaryFetching() bool`

HasDisableBinaryFetching returns a boolean if a field has been set.

### SetDisableBinaryFetchingNil

`func (o *HistoricDecisionInstanceQueryDto) SetDisableBinaryFetchingNil(b bool)`

 SetDisableBinaryFetchingNil sets the value for DisableBinaryFetching to be an explicit nil

### UnsetDisableBinaryFetching
`func (o *HistoricDecisionInstanceQueryDto) UnsetDisableBinaryFetching()`

UnsetDisableBinaryFetching ensures that no value is present for DisableBinaryFetching, not even an explicit nil
### GetDisableCustomObjectDeserialization

`func (o *HistoricDecisionInstanceQueryDto) GetDisableCustomObjectDeserialization() bool`

GetDisableCustomObjectDeserialization returns the DisableCustomObjectDeserialization field if non-nil, zero value otherwise.

### GetDisableCustomObjectDeserializationOk

`func (o *HistoricDecisionInstanceQueryDto) GetDisableCustomObjectDeserializationOk() (*bool, bool)`

GetDisableCustomObjectDeserializationOk returns a tuple with the DisableCustomObjectDeserialization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCustomObjectDeserialization

`func (o *HistoricDecisionInstanceQueryDto) SetDisableCustomObjectDeserialization(v bool)`

SetDisableCustomObjectDeserialization sets DisableCustomObjectDeserialization field to given value.

### HasDisableCustomObjectDeserialization

`func (o *HistoricDecisionInstanceQueryDto) HasDisableCustomObjectDeserialization() bool`

HasDisableCustomObjectDeserialization returns a boolean if a field has been set.

### SetDisableCustomObjectDeserializationNil

`func (o *HistoricDecisionInstanceQueryDto) SetDisableCustomObjectDeserializationNil(b bool)`

 SetDisableCustomObjectDeserializationNil sets the value for DisableCustomObjectDeserialization to be an explicit nil

### UnsetDisableCustomObjectDeserialization
`func (o *HistoricDecisionInstanceQueryDto) UnsetDisableCustomObjectDeserialization()`

UnsetDisableCustomObjectDeserialization ensures that no value is present for DisableCustomObjectDeserialization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


