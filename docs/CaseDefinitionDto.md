# CaseDefinitionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the case definition | [optional] 
**Key** | Pointer to **NullableString** | The key of the case definition, i.e., the id of the CMMN 2.0 XML case definition. | [optional] 
**Category** | Pointer to **NullableString** | The category of the case definition. | [optional] 
**Name** | Pointer to **NullableString** | The name of the case definition. | [optional] 
**Version** | Pointer to **NullableInt32** | The version of the case definition that the engine assigned to it. | [optional] 
**Resource** | Pointer to **NullableString** | The file name of the case definition. | [optional] 
**DeploymentId** | Pointer to **NullableString** | The deployment id of the case definition. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the case definition. | [optional] 
**HistoryTimeToLive** | Pointer to **NullableInt32** | History time to live value of the case definition. Is used within [History cleanup](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#history-cleanup). | [optional] 

## Methods

### NewCaseDefinitionDto

`func NewCaseDefinitionDto() *CaseDefinitionDto`

NewCaseDefinitionDto instantiates a new CaseDefinitionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseDefinitionDtoWithDefaults

`func NewCaseDefinitionDtoWithDefaults() *CaseDefinitionDto`

NewCaseDefinitionDtoWithDefaults instantiates a new CaseDefinitionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CaseDefinitionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaseDefinitionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaseDefinitionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaseDefinitionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CaseDefinitionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CaseDefinitionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKey

`func (o *CaseDefinitionDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CaseDefinitionDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CaseDefinitionDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CaseDefinitionDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CaseDefinitionDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CaseDefinitionDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetCategory

`func (o *CaseDefinitionDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CaseDefinitionDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CaseDefinitionDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CaseDefinitionDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *CaseDefinitionDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CaseDefinitionDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetName

`func (o *CaseDefinitionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaseDefinitionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaseDefinitionDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CaseDefinitionDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CaseDefinitionDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CaseDefinitionDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVersion

`func (o *CaseDefinitionDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CaseDefinitionDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CaseDefinitionDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CaseDefinitionDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *CaseDefinitionDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *CaseDefinitionDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetResource

`func (o *CaseDefinitionDto) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CaseDefinitionDto) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CaseDefinitionDto) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *CaseDefinitionDto) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *CaseDefinitionDto) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *CaseDefinitionDto) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetDeploymentId

`func (o *CaseDefinitionDto) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *CaseDefinitionDto) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *CaseDefinitionDto) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *CaseDefinitionDto) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *CaseDefinitionDto) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *CaseDefinitionDto) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetTenantId

`func (o *CaseDefinitionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CaseDefinitionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CaseDefinitionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CaseDefinitionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CaseDefinitionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CaseDefinitionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetHistoryTimeToLive

`func (o *CaseDefinitionDto) GetHistoryTimeToLive() int32`

GetHistoryTimeToLive returns the HistoryTimeToLive field if non-nil, zero value otherwise.

### GetHistoryTimeToLiveOk

`func (o *CaseDefinitionDto) GetHistoryTimeToLiveOk() (*int32, bool)`

GetHistoryTimeToLiveOk returns a tuple with the HistoryTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryTimeToLive

`func (o *CaseDefinitionDto) SetHistoryTimeToLive(v int32)`

SetHistoryTimeToLive sets HistoryTimeToLive field to given value.

### HasHistoryTimeToLive

`func (o *CaseDefinitionDto) HasHistoryTimeToLive() bool`

HasHistoryTimeToLive returns a boolean if a field has been set.

### SetHistoryTimeToLiveNil

`func (o *CaseDefinitionDto) SetHistoryTimeToLiveNil(b bool)`

 SetHistoryTimeToLiveNil sets the value for HistoryTimeToLive to be an explicit nil

### UnsetHistoryTimeToLive
`func (o *CaseDefinitionDto) UnsetHistoryTimeToLive()`

UnsetHistoryTimeToLive ensures that no value is present for HistoryTimeToLive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


