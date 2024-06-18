# UserCredentialsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **NullableString** | The users new password. | [optional] 
**AuthenticatedUserPassword** | Pointer to **NullableString** | The password of the authenticated user who changes the password of the user (i.e., the user with passed id as path parameter). | [optional] 

## Methods

### NewUserCredentialsDto

`func NewUserCredentialsDto() *UserCredentialsDto`

NewUserCredentialsDto instantiates a new UserCredentialsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCredentialsDtoWithDefaults

`func NewUserCredentialsDtoWithDefaults() *UserCredentialsDto`

NewUserCredentialsDtoWithDefaults instantiates a new UserCredentialsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UserCredentialsDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCredentialsDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCredentialsDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCredentialsDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UserCredentialsDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UserCredentialsDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetAuthenticatedUserPassword

`func (o *UserCredentialsDto) GetAuthenticatedUserPassword() string`

GetAuthenticatedUserPassword returns the AuthenticatedUserPassword field if non-nil, zero value otherwise.

### GetAuthenticatedUserPasswordOk

`func (o *UserCredentialsDto) GetAuthenticatedUserPasswordOk() (*string, bool)`

GetAuthenticatedUserPasswordOk returns a tuple with the AuthenticatedUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatedUserPassword

`func (o *UserCredentialsDto) SetAuthenticatedUserPassword(v string)`

SetAuthenticatedUserPassword sets AuthenticatedUserPassword field to given value.

### HasAuthenticatedUserPassword

`func (o *UserCredentialsDto) HasAuthenticatedUserPassword() bool`

HasAuthenticatedUserPassword returns a boolean if a field has been set.

### SetAuthenticatedUserPasswordNil

`func (o *UserCredentialsDto) SetAuthenticatedUserPasswordNil(b bool)`

 SetAuthenticatedUserPasswordNil sets the value for AuthenticatedUserPassword to be an explicit nil

### UnsetAuthenticatedUserPassword
`func (o *UserCredentialsDto) UnsetAuthenticatedUserPassword()`

UnsetAuthenticatedUserPassword ensures that no value is present for AuthenticatedUserPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


