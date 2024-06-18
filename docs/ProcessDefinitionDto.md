# ProcessDefinitionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the process definition | [optional] 
**Key** | Pointer to **NullableString** | The key of the process definition, i.e., the id of the BPMN 2.0 XML process definition. | [optional] 
**Category** | Pointer to **NullableString** | The category of the process definition. | [optional] 
**Description** | Pointer to **NullableString** | The description of the process definition. | [optional] 
**Name** | Pointer to **NullableString** | The name of the process definition. | [optional] 
**Version** | Pointer to **NullableInt32** | The version of the process definition that the engine assigned to it. | [optional] 
**Resource** | Pointer to **NullableString** | The file name of the process definition. | [optional] 
**DeploymentId** | Pointer to **NullableString** | The deployment id of the process definition. | [optional] 
**Diagram** | Pointer to **NullableString** | The file name of the process definition diagram, if it exists. | [optional] 
**Suspended** | Pointer to **NullableBool** | A flag indicating whether the definition is suspended or not. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the process definition. | [optional] 
**VersionTag** | Pointer to **NullableString** | The version tag of the process definition. | [optional] 
**HistoryTimeToLive** | Pointer to **NullableInt32** | History time to live value of the process definition. Is used within [History cleanup](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#history-cleanup). | [optional] 
**StartableInTasklist** | Pointer to **NullableBool** | A flag indicating whether the process definition is startable in Tasklist or not. | [optional] 

## Methods

### NewProcessDefinitionDto

`func NewProcessDefinitionDto() *ProcessDefinitionDto`

NewProcessDefinitionDto instantiates a new ProcessDefinitionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessDefinitionDtoWithDefaults

`func NewProcessDefinitionDtoWithDefaults() *ProcessDefinitionDto`

NewProcessDefinitionDtoWithDefaults instantiates a new ProcessDefinitionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessDefinitionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessDefinitionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessDefinitionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessDefinitionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProcessDefinitionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProcessDefinitionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKey

`func (o *ProcessDefinitionDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProcessDefinitionDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProcessDefinitionDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProcessDefinitionDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ProcessDefinitionDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ProcessDefinitionDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetCategory

`func (o *ProcessDefinitionDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ProcessDefinitionDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ProcessDefinitionDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ProcessDefinitionDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ProcessDefinitionDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ProcessDefinitionDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *ProcessDefinitionDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProcessDefinitionDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProcessDefinitionDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProcessDefinitionDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProcessDefinitionDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProcessDefinitionDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *ProcessDefinitionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessDefinitionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessDefinitionDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProcessDefinitionDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProcessDefinitionDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProcessDefinitionDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVersion

`func (o *ProcessDefinitionDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProcessDefinitionDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProcessDefinitionDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProcessDefinitionDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ProcessDefinitionDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ProcessDefinitionDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetResource

`func (o *ProcessDefinitionDto) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ProcessDefinitionDto) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ProcessDefinitionDto) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ProcessDefinitionDto) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *ProcessDefinitionDto) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *ProcessDefinitionDto) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetDeploymentId

`func (o *ProcessDefinitionDto) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ProcessDefinitionDto) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ProcessDefinitionDto) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *ProcessDefinitionDto) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *ProcessDefinitionDto) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *ProcessDefinitionDto) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetDiagram

`func (o *ProcessDefinitionDto) GetDiagram() string`

GetDiagram returns the Diagram field if non-nil, zero value otherwise.

### GetDiagramOk

`func (o *ProcessDefinitionDto) GetDiagramOk() (*string, bool)`

GetDiagramOk returns a tuple with the Diagram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagram

`func (o *ProcessDefinitionDto) SetDiagram(v string)`

SetDiagram sets Diagram field to given value.

### HasDiagram

`func (o *ProcessDefinitionDto) HasDiagram() bool`

HasDiagram returns a boolean if a field has been set.

### SetDiagramNil

`func (o *ProcessDefinitionDto) SetDiagramNil(b bool)`

 SetDiagramNil sets the value for Diagram to be an explicit nil

### UnsetDiagram
`func (o *ProcessDefinitionDto) UnsetDiagram()`

UnsetDiagram ensures that no value is present for Diagram, not even an explicit nil
### GetSuspended

`func (o *ProcessDefinitionDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ProcessDefinitionDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ProcessDefinitionDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ProcessDefinitionDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *ProcessDefinitionDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *ProcessDefinitionDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetTenantId

`func (o *ProcessDefinitionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ProcessDefinitionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ProcessDefinitionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ProcessDefinitionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ProcessDefinitionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ProcessDefinitionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetVersionTag

`func (o *ProcessDefinitionDto) GetVersionTag() string`

GetVersionTag returns the VersionTag field if non-nil, zero value otherwise.

### GetVersionTagOk

`func (o *ProcessDefinitionDto) GetVersionTagOk() (*string, bool)`

GetVersionTagOk returns a tuple with the VersionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionTag

`func (o *ProcessDefinitionDto) SetVersionTag(v string)`

SetVersionTag sets VersionTag field to given value.

### HasVersionTag

`func (o *ProcessDefinitionDto) HasVersionTag() bool`

HasVersionTag returns a boolean if a field has been set.

### SetVersionTagNil

`func (o *ProcessDefinitionDto) SetVersionTagNil(b bool)`

 SetVersionTagNil sets the value for VersionTag to be an explicit nil

### UnsetVersionTag
`func (o *ProcessDefinitionDto) UnsetVersionTag()`

UnsetVersionTag ensures that no value is present for VersionTag, not even an explicit nil
### GetHistoryTimeToLive

`func (o *ProcessDefinitionDto) GetHistoryTimeToLive() int32`

GetHistoryTimeToLive returns the HistoryTimeToLive field if non-nil, zero value otherwise.

### GetHistoryTimeToLiveOk

`func (o *ProcessDefinitionDto) GetHistoryTimeToLiveOk() (*int32, bool)`

GetHistoryTimeToLiveOk returns a tuple with the HistoryTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryTimeToLive

`func (o *ProcessDefinitionDto) SetHistoryTimeToLive(v int32)`

SetHistoryTimeToLive sets HistoryTimeToLive field to given value.

### HasHistoryTimeToLive

`func (o *ProcessDefinitionDto) HasHistoryTimeToLive() bool`

HasHistoryTimeToLive returns a boolean if a field has been set.

### SetHistoryTimeToLiveNil

`func (o *ProcessDefinitionDto) SetHistoryTimeToLiveNil(b bool)`

 SetHistoryTimeToLiveNil sets the value for HistoryTimeToLive to be an explicit nil

### UnsetHistoryTimeToLive
`func (o *ProcessDefinitionDto) UnsetHistoryTimeToLive()`

UnsetHistoryTimeToLive ensures that no value is present for HistoryTimeToLive, not even an explicit nil
### GetStartableInTasklist

`func (o *ProcessDefinitionDto) GetStartableInTasklist() bool`

GetStartableInTasklist returns the StartableInTasklist field if non-nil, zero value otherwise.

### GetStartableInTasklistOk

`func (o *ProcessDefinitionDto) GetStartableInTasklistOk() (*bool, bool)`

GetStartableInTasklistOk returns a tuple with the StartableInTasklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartableInTasklist

`func (o *ProcessDefinitionDto) SetStartableInTasklist(v bool)`

SetStartableInTasklist sets StartableInTasklist field to given value.

### HasStartableInTasklist

`func (o *ProcessDefinitionDto) HasStartableInTasklist() bool`

HasStartableInTasklist returns a boolean if a field has been set.

### SetStartableInTasklistNil

`func (o *ProcessDefinitionDto) SetStartableInTasklistNil(b bool)`

 SetStartableInTasklistNil sets the value for StartableInTasklist to be an explicit nil

### UnsetStartableInTasklist
`func (o *ProcessDefinitionDto) UnsetStartableInTasklist()`

UnsetStartableInTasklist ensures that no value is present for StartableInTasklist, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


