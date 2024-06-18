# AuthorizationCheckResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionName** | Pointer to **NullableString** | Name of the permission which was checked. | [optional] 
**ResourceName** | Pointer to **NullableString** | The name of the resource for which the permission check was performed. | [optional] 
**ResourceId** | Pointer to **NullableString** | The id of the resource for which the permission check was performed. | [optional] 
**Authorized** | Pointer to **NullableBool** | Returns true or false depending on whether the user is authorized or not. | [optional] 

## Methods

### NewAuthorizationCheckResultDto

`func NewAuthorizationCheckResultDto() *AuthorizationCheckResultDto`

NewAuthorizationCheckResultDto instantiates a new AuthorizationCheckResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationCheckResultDtoWithDefaults

`func NewAuthorizationCheckResultDtoWithDefaults() *AuthorizationCheckResultDto`

NewAuthorizationCheckResultDtoWithDefaults instantiates a new AuthorizationCheckResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionName

`func (o *AuthorizationCheckResultDto) GetPermissionName() string`

GetPermissionName returns the PermissionName field if non-nil, zero value otherwise.

### GetPermissionNameOk

`func (o *AuthorizationCheckResultDto) GetPermissionNameOk() (*string, bool)`

GetPermissionNameOk returns a tuple with the PermissionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionName

`func (o *AuthorizationCheckResultDto) SetPermissionName(v string)`

SetPermissionName sets PermissionName field to given value.

### HasPermissionName

`func (o *AuthorizationCheckResultDto) HasPermissionName() bool`

HasPermissionName returns a boolean if a field has been set.

### SetPermissionNameNil

`func (o *AuthorizationCheckResultDto) SetPermissionNameNil(b bool)`

 SetPermissionNameNil sets the value for PermissionName to be an explicit nil

### UnsetPermissionName
`func (o *AuthorizationCheckResultDto) UnsetPermissionName()`

UnsetPermissionName ensures that no value is present for PermissionName, not even an explicit nil
### GetResourceName

`func (o *AuthorizationCheckResultDto) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AuthorizationCheckResultDto) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AuthorizationCheckResultDto) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AuthorizationCheckResultDto) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *AuthorizationCheckResultDto) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *AuthorizationCheckResultDto) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceId

`func (o *AuthorizationCheckResultDto) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AuthorizationCheckResultDto) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AuthorizationCheckResultDto) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AuthorizationCheckResultDto) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *AuthorizationCheckResultDto) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AuthorizationCheckResultDto) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetAuthorized

`func (o *AuthorizationCheckResultDto) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *AuthorizationCheckResultDto) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *AuthorizationCheckResultDto) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *AuthorizationCheckResultDto) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### SetAuthorizedNil

`func (o *AuthorizationCheckResultDto) SetAuthorizedNil(b bool)`

 SetAuthorizedNil sets the value for Authorized to be an explicit nil

### UnsetAuthorized
`func (o *AuthorizationCheckResultDto) UnsetAuthorized()`

UnsetAuthorized ensures that no value is present for Authorized, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


