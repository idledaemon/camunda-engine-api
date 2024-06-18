# BasicUserCredentialsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** | The username of a user. | [optional] 
**Password** | Pointer to **NullableString** | A password of a user. | [optional] 

## Methods

### NewBasicUserCredentialsDto

`func NewBasicUserCredentialsDto() *BasicUserCredentialsDto`

NewBasicUserCredentialsDto instantiates a new BasicUserCredentialsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicUserCredentialsDtoWithDefaults

`func NewBasicUserCredentialsDtoWithDefaults() *BasicUserCredentialsDto`

NewBasicUserCredentialsDtoWithDefaults instantiates a new BasicUserCredentialsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *BasicUserCredentialsDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BasicUserCredentialsDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BasicUserCredentialsDto) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *BasicUserCredentialsDto) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *BasicUserCredentialsDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *BasicUserCredentialsDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *BasicUserCredentialsDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BasicUserCredentialsDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BasicUserCredentialsDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BasicUserCredentialsDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *BasicUserCredentialsDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *BasicUserCredentialsDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


