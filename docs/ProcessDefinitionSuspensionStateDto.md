# ProcessDefinitionSuspensionStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suspended** | Pointer to **NullableBool** | A &#x60;Boolean&#x60; value which indicates whether to activate or suspend all process definitions with the given key. When the value is set to &#x60;true&#x60;, all process definitions with the given key will be suspended and when the value is set to &#x60;false&#x60;, all process definitions with the given key will be activated. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definitions to activate or suspend. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definitions to activate or suspend. | [optional] 
**IncludeProcessInstances** | Pointer to **NullableBool** | A &#x60;Boolean&#x60; value which indicates whether to activate or suspend also all process instances of  the process definitions with the given key. When the value is set to &#x60;true&#x60;, all process instances of the process definitions with the given key will be activated or suspended and when the value is set to &#x60;false&#x60;, the suspension state of  all process instances of the process definitions with the given key will not be updated. | [optional] 
**ExecutionDate** | Pointer to **NullableTime** | The date on which all process definitions with the given key will be activated or suspended. If &#x60;null&#x60;, the suspension state of all process definitions with the given key is updated immediately. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 

## Methods

### NewProcessDefinitionSuspensionStateDto

`func NewProcessDefinitionSuspensionStateDto() *ProcessDefinitionSuspensionStateDto`

NewProcessDefinitionSuspensionStateDto instantiates a new ProcessDefinitionSuspensionStateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessDefinitionSuspensionStateDtoWithDefaults

`func NewProcessDefinitionSuspensionStateDtoWithDefaults() *ProcessDefinitionSuspensionStateDto`

NewProcessDefinitionSuspensionStateDtoWithDefaults instantiates a new ProcessDefinitionSuspensionStateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuspended

`func (o *ProcessDefinitionSuspensionStateDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ProcessDefinitionSuspensionStateDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ProcessDefinitionSuspensionStateDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ProcessDefinitionSuspensionStateDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *ProcessDefinitionSuspensionStateDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *ProcessDefinitionSuspensionStateDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetProcessDefinitionId

`func (o *ProcessDefinitionSuspensionStateDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *ProcessDefinitionSuspensionStateDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *ProcessDefinitionSuspensionStateDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *ProcessDefinitionSuspensionStateDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *ProcessDefinitionSuspensionStateDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *ProcessDefinitionSuspensionStateDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *ProcessDefinitionSuspensionStateDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *ProcessDefinitionSuspensionStateDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *ProcessDefinitionSuspensionStateDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *ProcessDefinitionSuspensionStateDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *ProcessDefinitionSuspensionStateDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *ProcessDefinitionSuspensionStateDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetIncludeProcessInstances

`func (o *ProcessDefinitionSuspensionStateDto) GetIncludeProcessInstances() bool`

GetIncludeProcessInstances returns the IncludeProcessInstances field if non-nil, zero value otherwise.

### GetIncludeProcessInstancesOk

`func (o *ProcessDefinitionSuspensionStateDto) GetIncludeProcessInstancesOk() (*bool, bool)`

GetIncludeProcessInstancesOk returns a tuple with the IncludeProcessInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProcessInstances

`func (o *ProcessDefinitionSuspensionStateDto) SetIncludeProcessInstances(v bool)`

SetIncludeProcessInstances sets IncludeProcessInstances field to given value.

### HasIncludeProcessInstances

`func (o *ProcessDefinitionSuspensionStateDto) HasIncludeProcessInstances() bool`

HasIncludeProcessInstances returns a boolean if a field has been set.

### SetIncludeProcessInstancesNil

`func (o *ProcessDefinitionSuspensionStateDto) SetIncludeProcessInstancesNil(b bool)`

 SetIncludeProcessInstancesNil sets the value for IncludeProcessInstances to be an explicit nil

### UnsetIncludeProcessInstances
`func (o *ProcessDefinitionSuspensionStateDto) UnsetIncludeProcessInstances()`

UnsetIncludeProcessInstances ensures that no value is present for IncludeProcessInstances, not even an explicit nil
### GetExecutionDate

`func (o *ProcessDefinitionSuspensionStateDto) GetExecutionDate() time.Time`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ProcessDefinitionSuspensionStateDto) GetExecutionDateOk() (*time.Time, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ProcessDefinitionSuspensionStateDto) SetExecutionDate(v time.Time)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *ProcessDefinitionSuspensionStateDto) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### SetExecutionDateNil

`func (o *ProcessDefinitionSuspensionStateDto) SetExecutionDateNil(b bool)`

 SetExecutionDateNil sets the value for ExecutionDate to be an explicit nil

### UnsetExecutionDate
`func (o *ProcessDefinitionSuspensionStateDto) UnsetExecutionDate()`

UnsetExecutionDate ensures that no value is present for ExecutionDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


