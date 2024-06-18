# PasswordPolicyRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **NullableString** | The candidate password to be check against the password policy. | [optional] 
**Profile** | Pointer to [**UserProfileDto**](UserProfileDto.md) |  | [optional] 

## Methods

### NewPasswordPolicyRequestDto

`func NewPasswordPolicyRequestDto() *PasswordPolicyRequestDto`

NewPasswordPolicyRequestDto instantiates a new PasswordPolicyRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyRequestDtoWithDefaults

`func NewPasswordPolicyRequestDtoWithDefaults() *PasswordPolicyRequestDto`

NewPasswordPolicyRequestDtoWithDefaults instantiates a new PasswordPolicyRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PasswordPolicyRequestDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordPolicyRequestDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordPolicyRequestDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PasswordPolicyRequestDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PasswordPolicyRequestDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PasswordPolicyRequestDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetProfile

`func (o *PasswordPolicyRequestDto) GetProfile() UserProfileDto`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PasswordPolicyRequestDto) GetProfileOk() (*UserProfileDto, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PasswordPolicyRequestDto) SetProfile(v UserProfileDto)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PasswordPolicyRequestDto) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


