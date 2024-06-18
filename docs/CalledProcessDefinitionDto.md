# CalledProcessDefinitionDto

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
**CalledFromActivityIds** | Pointer to **[]string** | Ids of the CallActivities which call this process. | [optional] 
**CallingProcessDefinitionId** | Pointer to **NullableString** | The id of the calling process definition | [optional] 

## Methods

### NewCalledProcessDefinitionDto

`func NewCalledProcessDefinitionDto() *CalledProcessDefinitionDto`

NewCalledProcessDefinitionDto instantiates a new CalledProcessDefinitionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalledProcessDefinitionDtoWithDefaults

`func NewCalledProcessDefinitionDtoWithDefaults() *CalledProcessDefinitionDto`

NewCalledProcessDefinitionDtoWithDefaults instantiates a new CalledProcessDefinitionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CalledProcessDefinitionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CalledProcessDefinitionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CalledProcessDefinitionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CalledProcessDefinitionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CalledProcessDefinitionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CalledProcessDefinitionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKey

`func (o *CalledProcessDefinitionDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CalledProcessDefinitionDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CalledProcessDefinitionDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CalledProcessDefinitionDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CalledProcessDefinitionDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CalledProcessDefinitionDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetCategory

`func (o *CalledProcessDefinitionDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CalledProcessDefinitionDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CalledProcessDefinitionDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CalledProcessDefinitionDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *CalledProcessDefinitionDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CalledProcessDefinitionDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *CalledProcessDefinitionDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CalledProcessDefinitionDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CalledProcessDefinitionDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CalledProcessDefinitionDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CalledProcessDefinitionDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CalledProcessDefinitionDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *CalledProcessDefinitionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CalledProcessDefinitionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CalledProcessDefinitionDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CalledProcessDefinitionDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CalledProcessDefinitionDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CalledProcessDefinitionDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVersion

`func (o *CalledProcessDefinitionDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CalledProcessDefinitionDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CalledProcessDefinitionDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CalledProcessDefinitionDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *CalledProcessDefinitionDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *CalledProcessDefinitionDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetResource

`func (o *CalledProcessDefinitionDto) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CalledProcessDefinitionDto) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CalledProcessDefinitionDto) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *CalledProcessDefinitionDto) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *CalledProcessDefinitionDto) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *CalledProcessDefinitionDto) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetDeploymentId

`func (o *CalledProcessDefinitionDto) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *CalledProcessDefinitionDto) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *CalledProcessDefinitionDto) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *CalledProcessDefinitionDto) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *CalledProcessDefinitionDto) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *CalledProcessDefinitionDto) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetDiagram

`func (o *CalledProcessDefinitionDto) GetDiagram() string`

GetDiagram returns the Diagram field if non-nil, zero value otherwise.

### GetDiagramOk

`func (o *CalledProcessDefinitionDto) GetDiagramOk() (*string, bool)`

GetDiagramOk returns a tuple with the Diagram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagram

`func (o *CalledProcessDefinitionDto) SetDiagram(v string)`

SetDiagram sets Diagram field to given value.

### HasDiagram

`func (o *CalledProcessDefinitionDto) HasDiagram() bool`

HasDiagram returns a boolean if a field has been set.

### SetDiagramNil

`func (o *CalledProcessDefinitionDto) SetDiagramNil(b bool)`

 SetDiagramNil sets the value for Diagram to be an explicit nil

### UnsetDiagram
`func (o *CalledProcessDefinitionDto) UnsetDiagram()`

UnsetDiagram ensures that no value is present for Diagram, not even an explicit nil
### GetSuspended

`func (o *CalledProcessDefinitionDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *CalledProcessDefinitionDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *CalledProcessDefinitionDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *CalledProcessDefinitionDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *CalledProcessDefinitionDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *CalledProcessDefinitionDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetTenantId

`func (o *CalledProcessDefinitionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CalledProcessDefinitionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CalledProcessDefinitionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CalledProcessDefinitionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CalledProcessDefinitionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CalledProcessDefinitionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetVersionTag

`func (o *CalledProcessDefinitionDto) GetVersionTag() string`

GetVersionTag returns the VersionTag field if non-nil, zero value otherwise.

### GetVersionTagOk

`func (o *CalledProcessDefinitionDto) GetVersionTagOk() (*string, bool)`

GetVersionTagOk returns a tuple with the VersionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionTag

`func (o *CalledProcessDefinitionDto) SetVersionTag(v string)`

SetVersionTag sets VersionTag field to given value.

### HasVersionTag

`func (o *CalledProcessDefinitionDto) HasVersionTag() bool`

HasVersionTag returns a boolean if a field has been set.

### SetVersionTagNil

`func (o *CalledProcessDefinitionDto) SetVersionTagNil(b bool)`

 SetVersionTagNil sets the value for VersionTag to be an explicit nil

### UnsetVersionTag
`func (o *CalledProcessDefinitionDto) UnsetVersionTag()`

UnsetVersionTag ensures that no value is present for VersionTag, not even an explicit nil
### GetHistoryTimeToLive

`func (o *CalledProcessDefinitionDto) GetHistoryTimeToLive() int32`

GetHistoryTimeToLive returns the HistoryTimeToLive field if non-nil, zero value otherwise.

### GetHistoryTimeToLiveOk

`func (o *CalledProcessDefinitionDto) GetHistoryTimeToLiveOk() (*int32, bool)`

GetHistoryTimeToLiveOk returns a tuple with the HistoryTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryTimeToLive

`func (o *CalledProcessDefinitionDto) SetHistoryTimeToLive(v int32)`

SetHistoryTimeToLive sets HistoryTimeToLive field to given value.

### HasHistoryTimeToLive

`func (o *CalledProcessDefinitionDto) HasHistoryTimeToLive() bool`

HasHistoryTimeToLive returns a boolean if a field has been set.

### SetHistoryTimeToLiveNil

`func (o *CalledProcessDefinitionDto) SetHistoryTimeToLiveNil(b bool)`

 SetHistoryTimeToLiveNil sets the value for HistoryTimeToLive to be an explicit nil

### UnsetHistoryTimeToLive
`func (o *CalledProcessDefinitionDto) UnsetHistoryTimeToLive()`

UnsetHistoryTimeToLive ensures that no value is present for HistoryTimeToLive, not even an explicit nil
### GetStartableInTasklist

`func (o *CalledProcessDefinitionDto) GetStartableInTasklist() bool`

GetStartableInTasklist returns the StartableInTasklist field if non-nil, zero value otherwise.

### GetStartableInTasklistOk

`func (o *CalledProcessDefinitionDto) GetStartableInTasklistOk() (*bool, bool)`

GetStartableInTasklistOk returns a tuple with the StartableInTasklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartableInTasklist

`func (o *CalledProcessDefinitionDto) SetStartableInTasklist(v bool)`

SetStartableInTasklist sets StartableInTasklist field to given value.

### HasStartableInTasklist

`func (o *CalledProcessDefinitionDto) HasStartableInTasklist() bool`

HasStartableInTasklist returns a boolean if a field has been set.

### SetStartableInTasklistNil

`func (o *CalledProcessDefinitionDto) SetStartableInTasklistNil(b bool)`

 SetStartableInTasklistNil sets the value for StartableInTasklist to be an explicit nil

### UnsetStartableInTasklist
`func (o *CalledProcessDefinitionDto) UnsetStartableInTasklist()`

UnsetStartableInTasklist ensures that no value is present for StartableInTasklist, not even an explicit nil
### GetCalledFromActivityIds

`func (o *CalledProcessDefinitionDto) GetCalledFromActivityIds() []string`

GetCalledFromActivityIds returns the CalledFromActivityIds field if non-nil, zero value otherwise.

### GetCalledFromActivityIdsOk

`func (o *CalledProcessDefinitionDto) GetCalledFromActivityIdsOk() (*[]string, bool)`

GetCalledFromActivityIdsOk returns a tuple with the CalledFromActivityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalledFromActivityIds

`func (o *CalledProcessDefinitionDto) SetCalledFromActivityIds(v []string)`

SetCalledFromActivityIds sets CalledFromActivityIds field to given value.

### HasCalledFromActivityIds

`func (o *CalledProcessDefinitionDto) HasCalledFromActivityIds() bool`

HasCalledFromActivityIds returns a boolean if a field has been set.

### SetCalledFromActivityIdsNil

`func (o *CalledProcessDefinitionDto) SetCalledFromActivityIdsNil(b bool)`

 SetCalledFromActivityIdsNil sets the value for CalledFromActivityIds to be an explicit nil

### UnsetCalledFromActivityIds
`func (o *CalledProcessDefinitionDto) UnsetCalledFromActivityIds()`

UnsetCalledFromActivityIds ensures that no value is present for CalledFromActivityIds, not even an explicit nil
### GetCallingProcessDefinitionId

`func (o *CalledProcessDefinitionDto) GetCallingProcessDefinitionId() string`

GetCallingProcessDefinitionId returns the CallingProcessDefinitionId field if non-nil, zero value otherwise.

### GetCallingProcessDefinitionIdOk

`func (o *CalledProcessDefinitionDto) GetCallingProcessDefinitionIdOk() (*string, bool)`

GetCallingProcessDefinitionIdOk returns a tuple with the CallingProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingProcessDefinitionId

`func (o *CalledProcessDefinitionDto) SetCallingProcessDefinitionId(v string)`

SetCallingProcessDefinitionId sets CallingProcessDefinitionId field to given value.

### HasCallingProcessDefinitionId

`func (o *CalledProcessDefinitionDto) HasCallingProcessDefinitionId() bool`

HasCallingProcessDefinitionId returns a boolean if a field has been set.

### SetCallingProcessDefinitionIdNil

`func (o *CalledProcessDefinitionDto) SetCallingProcessDefinitionIdNil(b bool)`

 SetCallingProcessDefinitionIdNil sets the value for CallingProcessDefinitionId to be an explicit nil

### UnsetCallingProcessDefinitionId
`func (o *CalledProcessDefinitionDto) UnsetCallingProcessDefinitionId()`

UnsetCallingProcessDefinitionId ensures that no value is present for CallingProcessDefinitionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


