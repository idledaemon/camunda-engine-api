# FilterDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the filter. | [optional] 
**ResourceType** | Pointer to **NullableString** | The resource type of the filter. | [optional] 
**Name** | Pointer to **NullableString** | The name of the filter. | [optional] 
**Owner** | Pointer to **NullableString** | The user id of the owner of the filter. | [optional] 
**Query** | Pointer to **map[string]interface{}** | The query of the filter as a JSON object. | [optional] 
**Properties** | Pointer to **map[string]interface{}** | The properties of a filter as a JSON object. | [optional] 
**ItemCount** | Pointer to **NullableInt64** |  The number of items matched by the filter itself. Note: Only exists if the query parameter &#x60;itemCount&#x60; was set to &#x60;true&#x60; | [optional] 

## Methods

### NewFilterDto

`func NewFilterDto() *FilterDto`

NewFilterDto instantiates a new FilterDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterDtoWithDefaults

`func NewFilterDtoWithDefaults() *FilterDto`

NewFilterDtoWithDefaults instantiates a new FilterDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilterDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilterDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilterDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FilterDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FilterDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FilterDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetResourceType

`func (o *FilterDto) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *FilterDto) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *FilterDto) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *FilterDto) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *FilterDto) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *FilterDto) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetName

`func (o *FilterDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilterDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilterDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilterDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FilterDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FilterDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwner

`func (o *FilterDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FilterDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FilterDto) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *FilterDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *FilterDto) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *FilterDto) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetQuery

`func (o *FilterDto) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *FilterDto) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *FilterDto) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *FilterDto) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetProperties

`func (o *FilterDto) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FilterDto) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FilterDto) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *FilterDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetItemCount

`func (o *FilterDto) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *FilterDto) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *FilterDto) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *FilterDto) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### SetItemCountNil

`func (o *FilterDto) SetItemCountNil(b bool)`

 SetItemCountNil sets the value for ItemCount to be an explicit nil

### UnsetItemCount
`func (o *FilterDto) UnsetItemCount()`

UnsetItemCount ensures that no value is present for ItemCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


