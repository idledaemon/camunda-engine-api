# DeploymentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]AtomLink**](AtomLink.md) | The links associated to this resource, with &#x60;method&#x60;, &#x60;href&#x60; and &#x60;rel&#x60;. | [optional] 
**Id** | Pointer to **NullableString** | The id of the deployment. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the deployment. | [optional] 
**DeploymentTime** | Pointer to **NullableTime** | The time when the deployment was created. | [optional] 
**Source** | Pointer to **NullableString** | The source of the deployment. | [optional] 
**Name** | Pointer to **NullableString** | The name of the deployment. | [optional] 

## Methods

### NewDeploymentDto

`func NewDeploymentDto() *DeploymentDto`

NewDeploymentDto instantiates a new DeploymentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDtoWithDefaults

`func NewDeploymentDtoWithDefaults() *DeploymentDto`

NewDeploymentDtoWithDefaults instantiates a new DeploymentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DeploymentDto) GetLinks() []AtomLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeploymentDto) GetLinksOk() (*[]AtomLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeploymentDto) SetLinks(v []AtomLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeploymentDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *DeploymentDto) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *DeploymentDto) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetId

`func (o *DeploymentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeploymentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeploymentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTenantId

`func (o *DeploymentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DeploymentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DeploymentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DeploymentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DeploymentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DeploymentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDeploymentTime

`func (o *DeploymentDto) GetDeploymentTime() time.Time`

GetDeploymentTime returns the DeploymentTime field if non-nil, zero value otherwise.

### GetDeploymentTimeOk

`func (o *DeploymentDto) GetDeploymentTimeOk() (*time.Time, bool)`

GetDeploymentTimeOk returns a tuple with the DeploymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTime

`func (o *DeploymentDto) SetDeploymentTime(v time.Time)`

SetDeploymentTime sets DeploymentTime field to given value.

### HasDeploymentTime

`func (o *DeploymentDto) HasDeploymentTime() bool`

HasDeploymentTime returns a boolean if a field has been set.

### SetDeploymentTimeNil

`func (o *DeploymentDto) SetDeploymentTimeNil(b bool)`

 SetDeploymentTimeNil sets the value for DeploymentTime to be an explicit nil

### UnsetDeploymentTime
`func (o *DeploymentDto) UnsetDeploymentTime()`

UnsetDeploymentTime ensures that no value is present for DeploymentTime, not even an explicit nil
### GetSource

`func (o *DeploymentDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeploymentDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeploymentDto) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DeploymentDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *DeploymentDto) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *DeploymentDto) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetName

`func (o *DeploymentDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeploymentDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeploymentDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


