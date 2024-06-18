# AuthenticationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatedUser** | Pointer to **NullableString** | An id of authenticated user. | [optional] 
**Authenticated** | Pointer to **NullableBool** | A flag indicating if user is authenticated. | [optional] 
**Tenants** | Pointer to **[]string** | Will be null. | [optional] 
**Groups** | Pointer to **[]string** | Will be null. | [optional] 

## Methods

### NewAuthenticationResult

`func NewAuthenticationResult() *AuthenticationResult`

NewAuthenticationResult instantiates a new AuthenticationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationResultWithDefaults

`func NewAuthenticationResultWithDefaults() *AuthenticationResult`

NewAuthenticationResultWithDefaults instantiates a new AuthenticationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatedUser

`func (o *AuthenticationResult) GetAuthenticatedUser() string`

GetAuthenticatedUser returns the AuthenticatedUser field if non-nil, zero value otherwise.

### GetAuthenticatedUserOk

`func (o *AuthenticationResult) GetAuthenticatedUserOk() (*string, bool)`

GetAuthenticatedUserOk returns a tuple with the AuthenticatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatedUser

`func (o *AuthenticationResult) SetAuthenticatedUser(v string)`

SetAuthenticatedUser sets AuthenticatedUser field to given value.

### HasAuthenticatedUser

`func (o *AuthenticationResult) HasAuthenticatedUser() bool`

HasAuthenticatedUser returns a boolean if a field has been set.

### SetAuthenticatedUserNil

`func (o *AuthenticationResult) SetAuthenticatedUserNil(b bool)`

 SetAuthenticatedUserNil sets the value for AuthenticatedUser to be an explicit nil

### UnsetAuthenticatedUser
`func (o *AuthenticationResult) UnsetAuthenticatedUser()`

UnsetAuthenticatedUser ensures that no value is present for AuthenticatedUser, not even an explicit nil
### GetAuthenticated

`func (o *AuthenticationResult) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *AuthenticationResult) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *AuthenticationResult) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *AuthenticationResult) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### SetAuthenticatedNil

`func (o *AuthenticationResult) SetAuthenticatedNil(b bool)`

 SetAuthenticatedNil sets the value for Authenticated to be an explicit nil

### UnsetAuthenticated
`func (o *AuthenticationResult) UnsetAuthenticated()`

UnsetAuthenticated ensures that no value is present for Authenticated, not even an explicit nil
### GetTenants

`func (o *AuthenticationResult) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *AuthenticationResult) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *AuthenticationResult) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *AuthenticationResult) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *AuthenticationResult) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *AuthenticationResult) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetGroups

`func (o *AuthenticationResult) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AuthenticationResult) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AuthenticationResult) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AuthenticationResult) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *AuthenticationResult) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *AuthenticationResult) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


