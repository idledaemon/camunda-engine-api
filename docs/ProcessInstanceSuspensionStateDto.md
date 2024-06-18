# ProcessInstanceSuspensionStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suspended** | Pointer to **NullableBool** | A Boolean value which indicates whether to activate or suspend a given instance  (e.g. process instance, job, job definition, or batch). When the value is set to true,  the given instance will be suspended and when the value is set to false,  the given instance will be activated. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The process definition id of the process instances to activate or suspend.  **Note**: This parameter can be used only with combination of &#x60;suspended&#x60;. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The process definition key of the process instances to activate or suspend.  **Note**: This parameter can be used only with combination of &#x60;suspended&#x60;, &#x60;processDefinitionTenantId&#x60;, and &#x60;processDefinitionWithoutTenantId&#x60;. | [optional] 
**ProcessDefinitionTenantId** | Pointer to **NullableString** | Only activate or suspend process instances of a process definition which belongs to a tenant with the given id.  **Note**: This parameter can be used only with combination of &#x60;suspended&#x60;, &#x60;processDefinitionKey&#x60;, and &#x60;processDefinitionWithoutTenantId&#x60;. | [optional] 
**ProcessDefinitionWithoutTenantId** | Pointer to **NullableBool** | Only activate or suspend process instances of a process definition which belongs to no tenant. Value may only be true, as false is the default behavior.  **Note**: This parameter can be used only with combination of &#x60;suspended&#x60;, &#x60;processDefinitionKey&#x60;, and &#x60;processDefinitionTenantId&#x60;. | [optional] 
**ProcessInstanceIds** | Pointer to **[]string** | A list of process instance ids which defines a group of process instances which will be activated or suspended by the operation.  **Note**: This parameter can be used only with combination of &#x60;suspended&#x60;, &#x60;processInstanceQuery&#x60;, and &#x60;historicProcessInstanceQuery&#x60;. | [optional] 
**ProcessInstanceQuery** | Pointer to [**ProcessInstanceQueryDto**](ProcessInstanceQueryDto.md) |  | [optional] 
**HistoricProcessInstanceQuery** | Pointer to [**HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) |  | [optional] 

## Methods

### NewProcessInstanceSuspensionStateDto

`func NewProcessInstanceSuspensionStateDto() *ProcessInstanceSuspensionStateDto`

NewProcessInstanceSuspensionStateDto instantiates a new ProcessInstanceSuspensionStateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessInstanceSuspensionStateDtoWithDefaults

`func NewProcessInstanceSuspensionStateDtoWithDefaults() *ProcessInstanceSuspensionStateDto`

NewProcessInstanceSuspensionStateDtoWithDefaults instantiates a new ProcessInstanceSuspensionStateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuspended

`func (o *ProcessInstanceSuspensionStateDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ProcessInstanceSuspensionStateDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ProcessInstanceSuspensionStateDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ProcessInstanceSuspensionStateDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *ProcessInstanceSuspensionStateDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *ProcessInstanceSuspensionStateDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetProcessDefinitionId

`func (o *ProcessInstanceSuspensionStateDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *ProcessInstanceSuspensionStateDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *ProcessInstanceSuspensionStateDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *ProcessInstanceSuspensionStateDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *ProcessInstanceSuspensionStateDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *ProcessInstanceSuspensionStateDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *ProcessInstanceSuspensionStateDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *ProcessInstanceSuspensionStateDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *ProcessInstanceSuspensionStateDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *ProcessInstanceSuspensionStateDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *ProcessInstanceSuspensionStateDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *ProcessInstanceSuspensionStateDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionTenantId

`func (o *ProcessInstanceSuspensionStateDto) GetProcessDefinitionTenantId() string`

GetProcessDefinitionTenantId returns the ProcessDefinitionTenantId field if non-nil, zero value otherwise.

### GetProcessDefinitionTenantIdOk

`func (o *ProcessInstanceSuspensionStateDto) GetProcessDefinitionTenantIdOk() (*string, bool)`

GetProcessDefinitionTenantIdOk returns a tuple with the ProcessDefinitionTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionTenantId

`func (o *ProcessInstanceSuspensionStateDto) SetProcessDefinitionTenantId(v string)`

SetProcessDefinitionTenantId sets ProcessDefinitionTenantId field to given value.

### HasProcessDefinitionTenantId

`func (o *ProcessInstanceSuspensionStateDto) HasProcessDefinitionTenantId() bool`

HasProcessDefinitionTenantId returns a boolean if a field has been set.

### SetProcessDefinitionTenantIdNil

`func (o *ProcessInstanceSuspensionStateDto) SetProcessDefinitionTenantIdNil(b bool)`

 SetProcessDefinitionTenantIdNil sets the value for ProcessDefinitionTenantId to be an explicit nil

### UnsetProcessDefinitionTenantId
`func (o *ProcessInstanceSuspensionStateDto) UnsetProcessDefinitionTenantId()`

UnsetProcessDefinitionTenantId ensures that no value is present for ProcessDefinitionTenantId, not even an explicit nil
### GetProcessDefinitionWithoutTenantId

`func (o *ProcessInstanceSuspensionStateDto) GetProcessDefinitionWithoutTenantId() bool`

GetProcessDefinitionWithoutTenantId returns the ProcessDefinitionWithoutTenantId field if non-nil, zero value otherwise.

### GetProcessDefinitionWithoutTenantIdOk

`func (o *ProcessInstanceSuspensionStateDto) GetProcessDefinitionWithoutTenantIdOk() (*bool, bool)`

GetProcessDefinitionWithoutTenantIdOk returns a tuple with the ProcessDefinitionWithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionWithoutTenantId

`func (o *ProcessInstanceSuspensionStateDto) SetProcessDefinitionWithoutTenantId(v bool)`

SetProcessDefinitionWithoutTenantId sets ProcessDefinitionWithoutTenantId field to given value.

### HasProcessDefinitionWithoutTenantId

`func (o *ProcessInstanceSuspensionStateDto) HasProcessDefinitionWithoutTenantId() bool`

HasProcessDefinitionWithoutTenantId returns a boolean if a field has been set.

### SetProcessDefinitionWithoutTenantIdNil

`func (o *ProcessInstanceSuspensionStateDto) SetProcessDefinitionWithoutTenantIdNil(b bool)`

 SetProcessDefinitionWithoutTenantIdNil sets the value for ProcessDefinitionWithoutTenantId to be an explicit nil

### UnsetProcessDefinitionWithoutTenantId
`func (o *ProcessInstanceSuspensionStateDto) UnsetProcessDefinitionWithoutTenantId()`

UnsetProcessDefinitionWithoutTenantId ensures that no value is present for ProcessDefinitionWithoutTenantId, not even an explicit nil
### GetProcessInstanceIds

`func (o *ProcessInstanceSuspensionStateDto) GetProcessInstanceIds() []string`

GetProcessInstanceIds returns the ProcessInstanceIds field if non-nil, zero value otherwise.

### GetProcessInstanceIdsOk

`func (o *ProcessInstanceSuspensionStateDto) GetProcessInstanceIdsOk() (*[]string, bool)`

GetProcessInstanceIdsOk returns a tuple with the ProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceIds

`func (o *ProcessInstanceSuspensionStateDto) SetProcessInstanceIds(v []string)`

SetProcessInstanceIds sets ProcessInstanceIds field to given value.

### HasProcessInstanceIds

`func (o *ProcessInstanceSuspensionStateDto) HasProcessInstanceIds() bool`

HasProcessInstanceIds returns a boolean if a field has been set.

### SetProcessInstanceIdsNil

`func (o *ProcessInstanceSuspensionStateDto) SetProcessInstanceIdsNil(b bool)`

 SetProcessInstanceIdsNil sets the value for ProcessInstanceIds to be an explicit nil

### UnsetProcessInstanceIds
`func (o *ProcessInstanceSuspensionStateDto) UnsetProcessInstanceIds()`

UnsetProcessInstanceIds ensures that no value is present for ProcessInstanceIds, not even an explicit nil
### GetProcessInstanceQuery

`func (o *ProcessInstanceSuspensionStateDto) GetProcessInstanceQuery() ProcessInstanceQueryDto`

GetProcessInstanceQuery returns the ProcessInstanceQuery field if non-nil, zero value otherwise.

### GetProcessInstanceQueryOk

`func (o *ProcessInstanceSuspensionStateDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool)`

GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceQuery

`func (o *ProcessInstanceSuspensionStateDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto)`

SetProcessInstanceQuery sets ProcessInstanceQuery field to given value.

### HasProcessInstanceQuery

`func (o *ProcessInstanceSuspensionStateDto) HasProcessInstanceQuery() bool`

HasProcessInstanceQuery returns a boolean if a field has been set.

### GetHistoricProcessInstanceQuery

`func (o *ProcessInstanceSuspensionStateDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto`

GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field if non-nil, zero value otherwise.

### GetHistoricProcessInstanceQueryOk

`func (o *ProcessInstanceSuspensionStateDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool)`

GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricProcessInstanceQuery

`func (o *ProcessInstanceSuspensionStateDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto)`

SetHistoricProcessInstanceQuery sets HistoricProcessInstanceQuery field to given value.

### HasHistoricProcessInstanceQuery

`func (o *ProcessInstanceSuspensionStateDto) HasHistoricProcessInstanceQuery() bool`

HasHistoricProcessInstanceQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


