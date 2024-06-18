# GroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the group. | [optional] 
**Name** | Pointer to **NullableString** | The name of the group. | [optional] 
**Type** | Pointer to **NullableString** | The type of the group. | [optional] 

## Methods

### NewGroupDto

`func NewGroupDto() *GroupDto`

NewGroupDto instantiates a new GroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupDtoWithDefaults

`func NewGroupDtoWithDefaults() *GroupDto`

NewGroupDtoWithDefaults instantiates a new GroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GroupDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GroupDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *GroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GroupDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GroupDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *GroupDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GroupDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GroupDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


