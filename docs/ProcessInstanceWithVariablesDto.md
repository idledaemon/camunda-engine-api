# ProcessInstanceWithVariablesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the process instance. | [optional] 
**DefinitionId** | Pointer to **NullableString** | The id of the process definition that this process instance belongs to. | [optional] 
**BusinessKey** | Pointer to **NullableString** | The business key of the process instance. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | The id of the case instance associated with the process instance. | [optional] 
**Ended** | Pointer to **NullableBool** | A flag indicating whether the process instance has ended or not. Deprecated: will always be false! | [optional] 
**Suspended** | Pointer to **NullableBool** | A flag indicating whether the process instance is suspended or not. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the process instance. | [optional] 
**Links** | Pointer to [**[]AtomLink**](AtomLink.md) | The links associated to this resource, with &#x60;method&#x60;, &#x60;href&#x60; and &#x60;rel&#x60;. | [optional] 
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | The id of the process instance. | [optional] 

## Methods

### NewProcessInstanceWithVariablesDto

`func NewProcessInstanceWithVariablesDto() *ProcessInstanceWithVariablesDto`

NewProcessInstanceWithVariablesDto instantiates a new ProcessInstanceWithVariablesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessInstanceWithVariablesDtoWithDefaults

`func NewProcessInstanceWithVariablesDtoWithDefaults() *ProcessInstanceWithVariablesDto`

NewProcessInstanceWithVariablesDtoWithDefaults instantiates a new ProcessInstanceWithVariablesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessInstanceWithVariablesDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessInstanceWithVariablesDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessInstanceWithVariablesDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessInstanceWithVariablesDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProcessInstanceWithVariablesDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProcessInstanceWithVariablesDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDefinitionId

`func (o *ProcessInstanceWithVariablesDto) GetDefinitionId() string`

GetDefinitionId returns the DefinitionId field if non-nil, zero value otherwise.

### GetDefinitionIdOk

`func (o *ProcessInstanceWithVariablesDto) GetDefinitionIdOk() (*string, bool)`

GetDefinitionIdOk returns a tuple with the DefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionId

`func (o *ProcessInstanceWithVariablesDto) SetDefinitionId(v string)`

SetDefinitionId sets DefinitionId field to given value.

### HasDefinitionId

`func (o *ProcessInstanceWithVariablesDto) HasDefinitionId() bool`

HasDefinitionId returns a boolean if a field has been set.

### SetDefinitionIdNil

`func (o *ProcessInstanceWithVariablesDto) SetDefinitionIdNil(b bool)`

 SetDefinitionIdNil sets the value for DefinitionId to be an explicit nil

### UnsetDefinitionId
`func (o *ProcessInstanceWithVariablesDto) UnsetDefinitionId()`

UnsetDefinitionId ensures that no value is present for DefinitionId, not even an explicit nil
### GetBusinessKey

`func (o *ProcessInstanceWithVariablesDto) GetBusinessKey() string`

GetBusinessKey returns the BusinessKey field if non-nil, zero value otherwise.

### GetBusinessKeyOk

`func (o *ProcessInstanceWithVariablesDto) GetBusinessKeyOk() (*string, bool)`

GetBusinessKeyOk returns a tuple with the BusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKey

`func (o *ProcessInstanceWithVariablesDto) SetBusinessKey(v string)`

SetBusinessKey sets BusinessKey field to given value.

### HasBusinessKey

`func (o *ProcessInstanceWithVariablesDto) HasBusinessKey() bool`

HasBusinessKey returns a boolean if a field has been set.

### SetBusinessKeyNil

`func (o *ProcessInstanceWithVariablesDto) SetBusinessKeyNil(b bool)`

 SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil

### UnsetBusinessKey
`func (o *ProcessInstanceWithVariablesDto) UnsetBusinessKey()`

UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
### GetCaseInstanceId

`func (o *ProcessInstanceWithVariablesDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *ProcessInstanceWithVariablesDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *ProcessInstanceWithVariablesDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *ProcessInstanceWithVariablesDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *ProcessInstanceWithVariablesDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *ProcessInstanceWithVariablesDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetEnded

`func (o *ProcessInstanceWithVariablesDto) GetEnded() bool`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *ProcessInstanceWithVariablesDto) GetEndedOk() (*bool, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *ProcessInstanceWithVariablesDto) SetEnded(v bool)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *ProcessInstanceWithVariablesDto) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### SetEndedNil

`func (o *ProcessInstanceWithVariablesDto) SetEndedNil(b bool)`

 SetEndedNil sets the value for Ended to be an explicit nil

### UnsetEnded
`func (o *ProcessInstanceWithVariablesDto) UnsetEnded()`

UnsetEnded ensures that no value is present for Ended, not even an explicit nil
### GetSuspended

`func (o *ProcessInstanceWithVariablesDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ProcessInstanceWithVariablesDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ProcessInstanceWithVariablesDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ProcessInstanceWithVariablesDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *ProcessInstanceWithVariablesDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *ProcessInstanceWithVariablesDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetTenantId

`func (o *ProcessInstanceWithVariablesDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ProcessInstanceWithVariablesDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ProcessInstanceWithVariablesDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ProcessInstanceWithVariablesDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ProcessInstanceWithVariablesDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ProcessInstanceWithVariablesDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetLinks

`func (o *ProcessInstanceWithVariablesDto) GetLinks() []AtomLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProcessInstanceWithVariablesDto) GetLinksOk() (*[]AtomLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProcessInstanceWithVariablesDto) SetLinks(v []AtomLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProcessInstanceWithVariablesDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ProcessInstanceWithVariablesDto) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ProcessInstanceWithVariablesDto) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetVariables

`func (o *ProcessInstanceWithVariablesDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ProcessInstanceWithVariablesDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ProcessInstanceWithVariablesDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ProcessInstanceWithVariablesDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *ProcessInstanceWithVariablesDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *ProcessInstanceWithVariablesDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


