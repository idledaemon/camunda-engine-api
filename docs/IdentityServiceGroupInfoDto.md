# IdentityServiceGroupInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]IdentityServiceGroupDto**](IdentityServiceGroupDto.md) | An array of group objects. | [optional] 
**GroupUsers** | Pointer to [**[]IdentityServiceUserDto**](IdentityServiceUserDto.md) | An array that contains all users that are member in one of the groups. | [optional] 

## Methods

### NewIdentityServiceGroupInfoDto

`func NewIdentityServiceGroupInfoDto() *IdentityServiceGroupInfoDto`

NewIdentityServiceGroupInfoDto instantiates a new IdentityServiceGroupInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityServiceGroupInfoDtoWithDefaults

`func NewIdentityServiceGroupInfoDtoWithDefaults() *IdentityServiceGroupInfoDto`

NewIdentityServiceGroupInfoDtoWithDefaults instantiates a new IdentityServiceGroupInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *IdentityServiceGroupInfoDto) GetGroups() []IdentityServiceGroupDto`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IdentityServiceGroupInfoDto) GetGroupsOk() (*[]IdentityServiceGroupDto, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IdentityServiceGroupInfoDto) SetGroups(v []IdentityServiceGroupDto)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IdentityServiceGroupInfoDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *IdentityServiceGroupInfoDto) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *IdentityServiceGroupInfoDto) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetGroupUsers

`func (o *IdentityServiceGroupInfoDto) GetGroupUsers() []IdentityServiceUserDto`

GetGroupUsers returns the GroupUsers field if non-nil, zero value otherwise.

### GetGroupUsersOk

`func (o *IdentityServiceGroupInfoDto) GetGroupUsersOk() (*[]IdentityServiceUserDto, bool)`

GetGroupUsersOk returns a tuple with the GroupUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUsers

`func (o *IdentityServiceGroupInfoDto) SetGroupUsers(v []IdentityServiceUserDto)`

SetGroupUsers sets GroupUsers field to given value.

### HasGroupUsers

`func (o *IdentityServiceGroupInfoDto) HasGroupUsers() bool`

HasGroupUsers returns a boolean if a field has been set.

### SetGroupUsersNil

`func (o *IdentityServiceGroupInfoDto) SetGroupUsersNil(b bool)`

 SetGroupUsersNil sets the value for GroupUsers to be an explicit nil

### UnsetGroupUsers
`func (o *IdentityServiceGroupInfoDto) UnsetGroupUsers()`

UnsetGroupUsers ensures that no value is present for GroupUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


