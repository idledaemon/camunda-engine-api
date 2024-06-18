# MissingAuthorizationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionName** | Pointer to **NullableString** | The permission name that the user is missing. | [optional] 
**ResourceName** | Pointer to **NullableString** | The name of the resource that the user is missing permission for. | [optional] 
**ResourceId** | Pointer to **NullableString** | The id of the resource that the user is missing permission for. | [optional] 

## Methods

### NewMissingAuthorizationDto

`func NewMissingAuthorizationDto() *MissingAuthorizationDto`

NewMissingAuthorizationDto instantiates a new MissingAuthorizationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissingAuthorizationDtoWithDefaults

`func NewMissingAuthorizationDtoWithDefaults() *MissingAuthorizationDto`

NewMissingAuthorizationDtoWithDefaults instantiates a new MissingAuthorizationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionName

`func (o *MissingAuthorizationDto) GetPermissionName() string`

GetPermissionName returns the PermissionName field if non-nil, zero value otherwise.

### GetPermissionNameOk

`func (o *MissingAuthorizationDto) GetPermissionNameOk() (*string, bool)`

GetPermissionNameOk returns a tuple with the PermissionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionName

`func (o *MissingAuthorizationDto) SetPermissionName(v string)`

SetPermissionName sets PermissionName field to given value.

### HasPermissionName

`func (o *MissingAuthorizationDto) HasPermissionName() bool`

HasPermissionName returns a boolean if a field has been set.

### SetPermissionNameNil

`func (o *MissingAuthorizationDto) SetPermissionNameNil(b bool)`

 SetPermissionNameNil sets the value for PermissionName to be an explicit nil

### UnsetPermissionName
`func (o *MissingAuthorizationDto) UnsetPermissionName()`

UnsetPermissionName ensures that no value is present for PermissionName, not even an explicit nil
### GetResourceName

`func (o *MissingAuthorizationDto) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *MissingAuthorizationDto) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *MissingAuthorizationDto) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *MissingAuthorizationDto) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *MissingAuthorizationDto) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *MissingAuthorizationDto) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceId

`func (o *MissingAuthorizationDto) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *MissingAuthorizationDto) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *MissingAuthorizationDto) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *MissingAuthorizationDto) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *MissingAuthorizationDto) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *MissingAuthorizationDto) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


