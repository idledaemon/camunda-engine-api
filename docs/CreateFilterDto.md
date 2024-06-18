# CreateFilterDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **NullableString** | The resource type of the filter. | [optional] 
**Name** | Pointer to **NullableString** | The name of the filter. | [optional] 
**Owner** | Pointer to **NullableString** | The user id of the owner of the filter. | [optional] 
**Query** | Pointer to **map[string]interface{}** | The query of the filter as a JSON object. | [optional] 
**Properties** | Pointer to **map[string]interface{}** | The properties of a filter as a JSON object. | [optional] 

## Methods

### NewCreateFilterDto

`func NewCreateFilterDto() *CreateFilterDto`

NewCreateFilterDto instantiates a new CreateFilterDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFilterDtoWithDefaults

`func NewCreateFilterDtoWithDefaults() *CreateFilterDto`

NewCreateFilterDtoWithDefaults instantiates a new CreateFilterDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *CreateFilterDto) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateFilterDto) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateFilterDto) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *CreateFilterDto) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *CreateFilterDto) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *CreateFilterDto) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetName

`func (o *CreateFilterDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFilterDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFilterDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateFilterDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateFilterDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateFilterDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwner

`func (o *CreateFilterDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateFilterDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateFilterDto) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateFilterDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *CreateFilterDto) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *CreateFilterDto) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetQuery

`func (o *CreateFilterDto) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateFilterDto) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateFilterDto) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CreateFilterDto) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetProperties

`func (o *CreateFilterDto) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateFilterDto) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateFilterDto) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateFilterDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


