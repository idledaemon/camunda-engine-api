# AuthorizationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableInt32** | The type of the authorization (0&#x3D;global, 1&#x3D;grant, 2&#x3D;revoke). See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service.md#authorization-type) for more information about authorization types. | [optional] 
**Permissions** | Pointer to **[]string** | An array of Strings holding the permissions provided by this authorization. | [optional] 
**UserId** | Pointer to **NullableString** | The id of the user this authorization has been created for. The value &#x60;*&#x60; represents a global authorization ranging over all users. | [optional] 
**GroupId** | Pointer to **NullableString** | The id of the group this authorization has been created for. | [optional] 
**ResourceType** | Pointer to **NullableInt32** | An integer representing the resource type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/authorization-service/#resources) for a list of integer representations of resource types. | [optional] 
**ResourceId** | Pointer to **NullableString** | The resource Id. The value &#x60;*&#x60; represents an authorization ranging over all instances of a resource. | [optional] 

## Methods

### NewAuthorizationCreateDto

`func NewAuthorizationCreateDto() *AuthorizationCreateDto`

NewAuthorizationCreateDto instantiates a new AuthorizationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationCreateDtoWithDefaults

`func NewAuthorizationCreateDtoWithDefaults() *AuthorizationCreateDto`

NewAuthorizationCreateDtoWithDefaults instantiates a new AuthorizationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizationCreateDto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizationCreateDto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizationCreateDto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *AuthorizationCreateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AuthorizationCreateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AuthorizationCreateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPermissions

`func (o *AuthorizationCreateDto) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AuthorizationCreateDto) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AuthorizationCreateDto) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AuthorizationCreateDto) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *AuthorizationCreateDto) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *AuthorizationCreateDto) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetUserId

`func (o *AuthorizationCreateDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthorizationCreateDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthorizationCreateDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthorizationCreateDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *AuthorizationCreateDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AuthorizationCreateDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetGroupId

`func (o *AuthorizationCreateDto) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AuthorizationCreateDto) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AuthorizationCreateDto) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AuthorizationCreateDto) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *AuthorizationCreateDto) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *AuthorizationCreateDto) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetResourceType

`func (o *AuthorizationCreateDto) GetResourceType() int32`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AuthorizationCreateDto) GetResourceTypeOk() (*int32, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AuthorizationCreateDto) SetResourceType(v int32)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AuthorizationCreateDto) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *AuthorizationCreateDto) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *AuthorizationCreateDto) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetResourceId

`func (o *AuthorizationCreateDto) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AuthorizationCreateDto) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AuthorizationCreateDto) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AuthorizationCreateDto) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *AuthorizationCreateDto) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AuthorizationCreateDto) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


