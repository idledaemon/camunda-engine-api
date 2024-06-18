# AuthorizationUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to **[]string** | An array of Strings holding the permissions provided by this authorization. | [optional] 
**UserId** | Pointer to **NullableString** | The id of the user this authorization has been created for. The value &#x60;*&#x60; represents a global authorization ranging over all users. | [optional] 
**GroupId** | Pointer to **NullableString** | The id of the group this authorization has been created for. | [optional] 
**ResourceType** | Pointer to **NullableInt32** | An integer representing the resource type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#resources) for a list of integer representations of resource types. | [optional] 
**ResourceId** | Pointer to **NullableString** | The resource Id. The value &#x60;*&#x60; represents an authorization ranging over all instances of a resource. | [optional] 

## Methods

### NewAuthorizationUpdateDto

`func NewAuthorizationUpdateDto() *AuthorizationUpdateDto`

NewAuthorizationUpdateDto instantiates a new AuthorizationUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationUpdateDtoWithDefaults

`func NewAuthorizationUpdateDtoWithDefaults() *AuthorizationUpdateDto`

NewAuthorizationUpdateDtoWithDefaults instantiates a new AuthorizationUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *AuthorizationUpdateDto) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AuthorizationUpdateDto) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AuthorizationUpdateDto) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AuthorizationUpdateDto) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *AuthorizationUpdateDto) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *AuthorizationUpdateDto) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetUserId

`func (o *AuthorizationUpdateDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthorizationUpdateDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthorizationUpdateDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthorizationUpdateDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *AuthorizationUpdateDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AuthorizationUpdateDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetGroupId

`func (o *AuthorizationUpdateDto) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AuthorizationUpdateDto) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AuthorizationUpdateDto) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AuthorizationUpdateDto) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *AuthorizationUpdateDto) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *AuthorizationUpdateDto) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetResourceType

`func (o *AuthorizationUpdateDto) GetResourceType() int32`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AuthorizationUpdateDto) GetResourceTypeOk() (*int32, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AuthorizationUpdateDto) SetResourceType(v int32)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AuthorizationUpdateDto) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *AuthorizationUpdateDto) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *AuthorizationUpdateDto) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetResourceId

`func (o *AuthorizationUpdateDto) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AuthorizationUpdateDto) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AuthorizationUpdateDto) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AuthorizationUpdateDto) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *AuthorizationUpdateDto) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AuthorizationUpdateDto) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


