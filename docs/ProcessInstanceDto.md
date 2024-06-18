# ProcessInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]AtomLink**](AtomLink.md) | The links associated to this resource, with &#x60;method&#x60;, &#x60;href&#x60; and &#x60;rel&#x60;. | [optional] 
**Id** | Pointer to **NullableString** | The id of the process instance. | [optional] 
**DefinitionId** | Pointer to **NullableString** | The id of the process definition that this process instance belongs to. | [optional] 
**BusinessKey** | Pointer to **NullableString** | The business key of the process instance. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | The id of the case instance associated with the process instance. | [optional] 
**Ended** | Pointer to **NullableBool** | A flag indicating whether the process instance has ended or not. Deprecated: will always be false! | [optional] 
**Suspended** | Pointer to **NullableBool** | A flag indicating whether the process instance is suspended or not. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the process instance. | [optional] 

## Methods

### NewProcessInstanceDto

`func NewProcessInstanceDto() *ProcessInstanceDto`

NewProcessInstanceDto instantiates a new ProcessInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessInstanceDtoWithDefaults

`func NewProcessInstanceDtoWithDefaults() *ProcessInstanceDto`

NewProcessInstanceDtoWithDefaults instantiates a new ProcessInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProcessInstanceDto) GetLinks() []AtomLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProcessInstanceDto) GetLinksOk() (*[]AtomLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProcessInstanceDto) SetLinks(v []AtomLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProcessInstanceDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ProcessInstanceDto) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ProcessInstanceDto) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetId

`func (o *ProcessInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessInstanceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessInstanceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProcessInstanceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProcessInstanceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDefinitionId

`func (o *ProcessInstanceDto) GetDefinitionId() string`

GetDefinitionId returns the DefinitionId field if non-nil, zero value otherwise.

### GetDefinitionIdOk

`func (o *ProcessInstanceDto) GetDefinitionIdOk() (*string, bool)`

GetDefinitionIdOk returns a tuple with the DefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionId

`func (o *ProcessInstanceDto) SetDefinitionId(v string)`

SetDefinitionId sets DefinitionId field to given value.

### HasDefinitionId

`func (o *ProcessInstanceDto) HasDefinitionId() bool`

HasDefinitionId returns a boolean if a field has been set.

### SetDefinitionIdNil

`func (o *ProcessInstanceDto) SetDefinitionIdNil(b bool)`

 SetDefinitionIdNil sets the value for DefinitionId to be an explicit nil

### UnsetDefinitionId
`func (o *ProcessInstanceDto) UnsetDefinitionId()`

UnsetDefinitionId ensures that no value is present for DefinitionId, not even an explicit nil
### GetBusinessKey

`func (o *ProcessInstanceDto) GetBusinessKey() string`

GetBusinessKey returns the BusinessKey field if non-nil, zero value otherwise.

### GetBusinessKeyOk

`func (o *ProcessInstanceDto) GetBusinessKeyOk() (*string, bool)`

GetBusinessKeyOk returns a tuple with the BusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKey

`func (o *ProcessInstanceDto) SetBusinessKey(v string)`

SetBusinessKey sets BusinessKey field to given value.

### HasBusinessKey

`func (o *ProcessInstanceDto) HasBusinessKey() bool`

HasBusinessKey returns a boolean if a field has been set.

### SetBusinessKeyNil

`func (o *ProcessInstanceDto) SetBusinessKeyNil(b bool)`

 SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil

### UnsetBusinessKey
`func (o *ProcessInstanceDto) UnsetBusinessKey()`

UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
### GetCaseInstanceId

`func (o *ProcessInstanceDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *ProcessInstanceDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *ProcessInstanceDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *ProcessInstanceDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *ProcessInstanceDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *ProcessInstanceDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetEnded

`func (o *ProcessInstanceDto) GetEnded() bool`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *ProcessInstanceDto) GetEndedOk() (*bool, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *ProcessInstanceDto) SetEnded(v bool)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *ProcessInstanceDto) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### SetEndedNil

`func (o *ProcessInstanceDto) SetEndedNil(b bool)`

 SetEndedNil sets the value for Ended to be an explicit nil

### UnsetEnded
`func (o *ProcessInstanceDto) UnsetEnded()`

UnsetEnded ensures that no value is present for Ended, not even an explicit nil
### GetSuspended

`func (o *ProcessInstanceDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ProcessInstanceDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ProcessInstanceDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ProcessInstanceDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *ProcessInstanceDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *ProcessInstanceDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetTenantId

`func (o *ProcessInstanceDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ProcessInstanceDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ProcessInstanceDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ProcessInstanceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ProcessInstanceDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ProcessInstanceDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


