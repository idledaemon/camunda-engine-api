# RedeploymentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceIds** | Pointer to **[]string** | A list of deployment resource ids to re-deploy. | [optional] 
**ResourceNames** | Pointer to **[]string** | A list of deployment resource names to re-deploy. | [optional] 
**Source** | Pointer to **NullableString** | Sets the source of the deployment. | [optional] 

## Methods

### NewRedeploymentDto

`func NewRedeploymentDto() *RedeploymentDto`

NewRedeploymentDto instantiates a new RedeploymentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeploymentDtoWithDefaults

`func NewRedeploymentDtoWithDefaults() *RedeploymentDto`

NewRedeploymentDtoWithDefaults instantiates a new RedeploymentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceIds

`func (o *RedeploymentDto) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *RedeploymentDto) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *RedeploymentDto) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.

### HasResourceIds

`func (o *RedeploymentDto) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.

### SetResourceIdsNil

`func (o *RedeploymentDto) SetResourceIdsNil(b bool)`

 SetResourceIdsNil sets the value for ResourceIds to be an explicit nil

### UnsetResourceIds
`func (o *RedeploymentDto) UnsetResourceIds()`

UnsetResourceIds ensures that no value is present for ResourceIds, not even an explicit nil
### GetResourceNames

`func (o *RedeploymentDto) GetResourceNames() []string`

GetResourceNames returns the ResourceNames field if non-nil, zero value otherwise.

### GetResourceNamesOk

`func (o *RedeploymentDto) GetResourceNamesOk() (*[]string, bool)`

GetResourceNamesOk returns a tuple with the ResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceNames

`func (o *RedeploymentDto) SetResourceNames(v []string)`

SetResourceNames sets ResourceNames field to given value.

### HasResourceNames

`func (o *RedeploymentDto) HasResourceNames() bool`

HasResourceNames returns a boolean if a field has been set.

### SetResourceNamesNil

`func (o *RedeploymentDto) SetResourceNamesNil(b bool)`

 SetResourceNamesNil sets the value for ResourceNames to be an explicit nil

### UnsetResourceNames
`func (o *RedeploymentDto) UnsetResourceNames()`

UnsetResourceNames ensures that no value is present for ResourceNames, not even an explicit nil
### GetSource

`func (o *RedeploymentDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RedeploymentDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RedeploymentDto) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RedeploymentDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *RedeploymentDto) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RedeploymentDto) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


