# GroupQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Filter by the id of the group. | [optional] 
**IdIn** | Pointer to **[]string** | Filter by a JSON string array of group ids. | [optional] 
**Name** | Pointer to **NullableString** | Filter by the name of the group. | [optional] 
**NameLike** | Pointer to **NullableString** | Filter by the name that the parameter is a substring of. | [optional] 
**Type** | Pointer to **NullableString** | Filter by the type of the group. | [optional] 
**Member** | Pointer to **NullableString** | Only retrieve groups where the given user id is a member of. | [optional] 
**MemberOfTenant** | Pointer to **NullableString** | Only retrieve groups which are members of the given tenant. | [optional] 
**Sorting** | Pointer to [**[]GroupQueryDtoSortingInner**](GroupQueryDtoSortingInner.md) | Apply sorting of the result | [optional] 

## Methods

### NewGroupQueryDto

`func NewGroupQueryDto() *GroupQueryDto`

NewGroupQueryDto instantiates a new GroupQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupQueryDtoWithDefaults

`func NewGroupQueryDtoWithDefaults() *GroupQueryDto`

NewGroupQueryDtoWithDefaults instantiates a new GroupQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupQueryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupQueryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupQueryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupQueryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GroupQueryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GroupQueryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdIn

`func (o *GroupQueryDto) GetIdIn() []string`

GetIdIn returns the IdIn field if non-nil, zero value otherwise.

### GetIdInOk

`func (o *GroupQueryDto) GetIdInOk() (*[]string, bool)`

GetIdInOk returns a tuple with the IdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdIn

`func (o *GroupQueryDto) SetIdIn(v []string)`

SetIdIn sets IdIn field to given value.

### HasIdIn

`func (o *GroupQueryDto) HasIdIn() bool`

HasIdIn returns a boolean if a field has been set.

### SetIdInNil

`func (o *GroupQueryDto) SetIdInNil(b bool)`

 SetIdInNil sets the value for IdIn to be an explicit nil

### UnsetIdIn
`func (o *GroupQueryDto) UnsetIdIn()`

UnsetIdIn ensures that no value is present for IdIn, not even an explicit nil
### GetName

`func (o *GroupQueryDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupQueryDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupQueryDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupQueryDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GroupQueryDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GroupQueryDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNameLike

`func (o *GroupQueryDto) GetNameLike() string`

GetNameLike returns the NameLike field if non-nil, zero value otherwise.

### GetNameLikeOk

`func (o *GroupQueryDto) GetNameLikeOk() (*string, bool)`

GetNameLikeOk returns a tuple with the NameLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLike

`func (o *GroupQueryDto) SetNameLike(v string)`

SetNameLike sets NameLike field to given value.

### HasNameLike

`func (o *GroupQueryDto) HasNameLike() bool`

HasNameLike returns a boolean if a field has been set.

### SetNameLikeNil

`func (o *GroupQueryDto) SetNameLikeNil(b bool)`

 SetNameLikeNil sets the value for NameLike to be an explicit nil

### UnsetNameLike
`func (o *GroupQueryDto) UnsetNameLike()`

UnsetNameLike ensures that no value is present for NameLike, not even an explicit nil
### GetType

`func (o *GroupQueryDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupQueryDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupQueryDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupQueryDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GroupQueryDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GroupQueryDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMember

`func (o *GroupQueryDto) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GroupQueryDto) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GroupQueryDto) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GroupQueryDto) HasMember() bool`

HasMember returns a boolean if a field has been set.

### SetMemberNil

`func (o *GroupQueryDto) SetMemberNil(b bool)`

 SetMemberNil sets the value for Member to be an explicit nil

### UnsetMember
`func (o *GroupQueryDto) UnsetMember()`

UnsetMember ensures that no value is present for Member, not even an explicit nil
### GetMemberOfTenant

`func (o *GroupQueryDto) GetMemberOfTenant() string`

GetMemberOfTenant returns the MemberOfTenant field if non-nil, zero value otherwise.

### GetMemberOfTenantOk

`func (o *GroupQueryDto) GetMemberOfTenantOk() (*string, bool)`

GetMemberOfTenantOk returns a tuple with the MemberOfTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfTenant

`func (o *GroupQueryDto) SetMemberOfTenant(v string)`

SetMemberOfTenant sets MemberOfTenant field to given value.

### HasMemberOfTenant

`func (o *GroupQueryDto) HasMemberOfTenant() bool`

HasMemberOfTenant returns a boolean if a field has been set.

### SetMemberOfTenantNil

`func (o *GroupQueryDto) SetMemberOfTenantNil(b bool)`

 SetMemberOfTenantNil sets the value for MemberOfTenant to be an explicit nil

### UnsetMemberOfTenant
`func (o *GroupQueryDto) UnsetMemberOfTenant()`

UnsetMemberOfTenant ensures that no value is present for MemberOfTenant, not even an explicit nil
### GetSorting

`func (o *GroupQueryDto) GetSorting() []GroupQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *GroupQueryDto) GetSortingOk() (*[]GroupQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *GroupQueryDto) SetSorting(v []GroupQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *GroupQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *GroupQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *GroupQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


