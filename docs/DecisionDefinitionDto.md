# DecisionDefinitionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the decision definition | [optional] 
**Key** | Pointer to **NullableString** | The key of the decision definition, i.e., the id of the DMN 1.0 XML decision definition. | [optional] 
**Category** | Pointer to **NullableString** | The category of the decision definition. | [optional] 
**Name** | Pointer to **NullableString** | The name of the decision definition. | [optional] 
**Version** | Pointer to **NullableInt32** | The version of the decision definition that the engine assigned to it. | [optional] 
**Resource** | Pointer to **NullableString** | The file name of the decision definition. | [optional] 
**DeploymentId** | Pointer to **NullableString** | The deployment id of the decision definition. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the decision definition. | [optional] 
**DecisionRequirementsDefinitionId** | Pointer to **NullableString** | The id of the decision requirements definition this decision definition belongs to. | [optional] 
**DecisionRequirementsDefinitionKey** | Pointer to **NullableString** | The key of the decision requirements definition this decision definition belongs to. | [optional] 
**HistoryTimeToLive** | Pointer to **NullableInt32** | History time to live value of the decision definition. Is used within [History cleanup](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#history-cleanup). | [optional] 
**VersionTag** | Pointer to **NullableString** | The version tag of the decision definition. | [optional] 

## Methods

### NewDecisionDefinitionDto

`func NewDecisionDefinitionDto() *DecisionDefinitionDto`

NewDecisionDefinitionDto instantiates a new DecisionDefinitionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionDefinitionDtoWithDefaults

`func NewDecisionDefinitionDtoWithDefaults() *DecisionDefinitionDto`

NewDecisionDefinitionDtoWithDefaults instantiates a new DecisionDefinitionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DecisionDefinitionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DecisionDefinitionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DecisionDefinitionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DecisionDefinitionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DecisionDefinitionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DecisionDefinitionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKey

`func (o *DecisionDefinitionDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DecisionDefinitionDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DecisionDefinitionDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DecisionDefinitionDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *DecisionDefinitionDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *DecisionDefinitionDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetCategory

`func (o *DecisionDefinitionDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DecisionDefinitionDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DecisionDefinitionDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DecisionDefinitionDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *DecisionDefinitionDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *DecisionDefinitionDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetName

`func (o *DecisionDefinitionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DecisionDefinitionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DecisionDefinitionDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DecisionDefinitionDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DecisionDefinitionDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DecisionDefinitionDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVersion

`func (o *DecisionDefinitionDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DecisionDefinitionDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DecisionDefinitionDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DecisionDefinitionDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DecisionDefinitionDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DecisionDefinitionDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetResource

`func (o *DecisionDefinitionDto) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DecisionDefinitionDto) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DecisionDefinitionDto) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *DecisionDefinitionDto) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *DecisionDefinitionDto) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *DecisionDefinitionDto) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetDeploymentId

`func (o *DecisionDefinitionDto) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DecisionDefinitionDto) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DecisionDefinitionDto) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *DecisionDefinitionDto) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *DecisionDefinitionDto) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *DecisionDefinitionDto) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetTenantId

`func (o *DecisionDefinitionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DecisionDefinitionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DecisionDefinitionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DecisionDefinitionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DecisionDefinitionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DecisionDefinitionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDecisionRequirementsDefinitionId

`func (o *DecisionDefinitionDto) GetDecisionRequirementsDefinitionId() string`

GetDecisionRequirementsDefinitionId returns the DecisionRequirementsDefinitionId field if non-nil, zero value otherwise.

### GetDecisionRequirementsDefinitionIdOk

`func (o *DecisionDefinitionDto) GetDecisionRequirementsDefinitionIdOk() (*string, bool)`

GetDecisionRequirementsDefinitionIdOk returns a tuple with the DecisionRequirementsDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionRequirementsDefinitionId

`func (o *DecisionDefinitionDto) SetDecisionRequirementsDefinitionId(v string)`

SetDecisionRequirementsDefinitionId sets DecisionRequirementsDefinitionId field to given value.

### HasDecisionRequirementsDefinitionId

`func (o *DecisionDefinitionDto) HasDecisionRequirementsDefinitionId() bool`

HasDecisionRequirementsDefinitionId returns a boolean if a field has been set.

### SetDecisionRequirementsDefinitionIdNil

`func (o *DecisionDefinitionDto) SetDecisionRequirementsDefinitionIdNil(b bool)`

 SetDecisionRequirementsDefinitionIdNil sets the value for DecisionRequirementsDefinitionId to be an explicit nil

### UnsetDecisionRequirementsDefinitionId
`func (o *DecisionDefinitionDto) UnsetDecisionRequirementsDefinitionId()`

UnsetDecisionRequirementsDefinitionId ensures that no value is present for DecisionRequirementsDefinitionId, not even an explicit nil
### GetDecisionRequirementsDefinitionKey

`func (o *DecisionDefinitionDto) GetDecisionRequirementsDefinitionKey() string`

GetDecisionRequirementsDefinitionKey returns the DecisionRequirementsDefinitionKey field if non-nil, zero value otherwise.

### GetDecisionRequirementsDefinitionKeyOk

`func (o *DecisionDefinitionDto) GetDecisionRequirementsDefinitionKeyOk() (*string, bool)`

GetDecisionRequirementsDefinitionKeyOk returns a tuple with the DecisionRequirementsDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionRequirementsDefinitionKey

`func (o *DecisionDefinitionDto) SetDecisionRequirementsDefinitionKey(v string)`

SetDecisionRequirementsDefinitionKey sets DecisionRequirementsDefinitionKey field to given value.

### HasDecisionRequirementsDefinitionKey

`func (o *DecisionDefinitionDto) HasDecisionRequirementsDefinitionKey() bool`

HasDecisionRequirementsDefinitionKey returns a boolean if a field has been set.

### SetDecisionRequirementsDefinitionKeyNil

`func (o *DecisionDefinitionDto) SetDecisionRequirementsDefinitionKeyNil(b bool)`

 SetDecisionRequirementsDefinitionKeyNil sets the value for DecisionRequirementsDefinitionKey to be an explicit nil

### UnsetDecisionRequirementsDefinitionKey
`func (o *DecisionDefinitionDto) UnsetDecisionRequirementsDefinitionKey()`

UnsetDecisionRequirementsDefinitionKey ensures that no value is present for DecisionRequirementsDefinitionKey, not even an explicit nil
### GetHistoryTimeToLive

`func (o *DecisionDefinitionDto) GetHistoryTimeToLive() int32`

GetHistoryTimeToLive returns the HistoryTimeToLive field if non-nil, zero value otherwise.

### GetHistoryTimeToLiveOk

`func (o *DecisionDefinitionDto) GetHistoryTimeToLiveOk() (*int32, bool)`

GetHistoryTimeToLiveOk returns a tuple with the HistoryTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryTimeToLive

`func (o *DecisionDefinitionDto) SetHistoryTimeToLive(v int32)`

SetHistoryTimeToLive sets HistoryTimeToLive field to given value.

### HasHistoryTimeToLive

`func (o *DecisionDefinitionDto) HasHistoryTimeToLive() bool`

HasHistoryTimeToLive returns a boolean if a field has been set.

### SetHistoryTimeToLiveNil

`func (o *DecisionDefinitionDto) SetHistoryTimeToLiveNil(b bool)`

 SetHistoryTimeToLiveNil sets the value for HistoryTimeToLive to be an explicit nil

### UnsetHistoryTimeToLive
`func (o *DecisionDefinitionDto) UnsetHistoryTimeToLive()`

UnsetHistoryTimeToLive ensures that no value is present for HistoryTimeToLive, not even an explicit nil
### GetVersionTag

`func (o *DecisionDefinitionDto) GetVersionTag() string`

GetVersionTag returns the VersionTag field if non-nil, zero value otherwise.

### GetVersionTagOk

`func (o *DecisionDefinitionDto) GetVersionTagOk() (*string, bool)`

GetVersionTagOk returns a tuple with the VersionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionTag

`func (o *DecisionDefinitionDto) SetVersionTag(v string)`

SetVersionTag sets VersionTag field to given value.

### HasVersionTag

`func (o *DecisionDefinitionDto) HasVersionTag() bool`

HasVersionTag returns a boolean if a field has been set.

### SetVersionTagNil

`func (o *DecisionDefinitionDto) SetVersionTagNil(b bool)`

 SetVersionTagNil sets the value for VersionTag to be an explicit nil

### UnsetVersionTag
`func (o *DecisionDefinitionDto) UnsetVersionTag()`

UnsetVersionTag ensures that no value is present for VersionTag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


