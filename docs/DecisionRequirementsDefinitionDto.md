# DecisionRequirementsDefinitionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the decision requirements definition. | [optional] 
**Key** | Pointer to **NullableString** | The key of the decision requirements definition. | [optional] 
**Category** | Pointer to **NullableString** | The category of the decision requirements definition. | [optional] 
**Name** | Pointer to **NullableString** | The name of the decision requirements definition. | [optional] 
**Version** | Pointer to **NullableInt32** | The version of the decision requirements definition that the engine assigned to it. | [optional] 
**Resource** | Pointer to **NullableString** | The file name of the decision requirements definition. | [optional] 
**DeploymentId** | Pointer to **NullableString** | The deployment id of the decision requirements definition. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the decision requirements definition. | [optional] 

## Methods

### NewDecisionRequirementsDefinitionDto

`func NewDecisionRequirementsDefinitionDto() *DecisionRequirementsDefinitionDto`

NewDecisionRequirementsDefinitionDto instantiates a new DecisionRequirementsDefinitionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionRequirementsDefinitionDtoWithDefaults

`func NewDecisionRequirementsDefinitionDtoWithDefaults() *DecisionRequirementsDefinitionDto`

NewDecisionRequirementsDefinitionDtoWithDefaults instantiates a new DecisionRequirementsDefinitionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DecisionRequirementsDefinitionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DecisionRequirementsDefinitionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DecisionRequirementsDefinitionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DecisionRequirementsDefinitionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DecisionRequirementsDefinitionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DecisionRequirementsDefinitionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKey

`func (o *DecisionRequirementsDefinitionDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DecisionRequirementsDefinitionDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DecisionRequirementsDefinitionDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DecisionRequirementsDefinitionDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *DecisionRequirementsDefinitionDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *DecisionRequirementsDefinitionDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetCategory

`func (o *DecisionRequirementsDefinitionDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DecisionRequirementsDefinitionDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DecisionRequirementsDefinitionDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DecisionRequirementsDefinitionDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *DecisionRequirementsDefinitionDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *DecisionRequirementsDefinitionDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetName

`func (o *DecisionRequirementsDefinitionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DecisionRequirementsDefinitionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DecisionRequirementsDefinitionDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DecisionRequirementsDefinitionDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DecisionRequirementsDefinitionDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DecisionRequirementsDefinitionDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVersion

`func (o *DecisionRequirementsDefinitionDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DecisionRequirementsDefinitionDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DecisionRequirementsDefinitionDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DecisionRequirementsDefinitionDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DecisionRequirementsDefinitionDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DecisionRequirementsDefinitionDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetResource

`func (o *DecisionRequirementsDefinitionDto) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DecisionRequirementsDefinitionDto) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DecisionRequirementsDefinitionDto) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *DecisionRequirementsDefinitionDto) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *DecisionRequirementsDefinitionDto) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *DecisionRequirementsDefinitionDto) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetDeploymentId

`func (o *DecisionRequirementsDefinitionDto) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DecisionRequirementsDefinitionDto) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DecisionRequirementsDefinitionDto) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *DecisionRequirementsDefinitionDto) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *DecisionRequirementsDefinitionDto) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *DecisionRequirementsDefinitionDto) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetTenantId

`func (o *DecisionRequirementsDefinitionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DecisionRequirementsDefinitionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DecisionRequirementsDefinitionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DecisionRequirementsDefinitionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DecisionRequirementsDefinitionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DecisionRequirementsDefinitionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


