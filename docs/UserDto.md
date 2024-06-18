# UserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**UserProfileDto**](UserProfileDto.md) |  | [optional] 
**Credentials** | Pointer to [**UserCredentialsDto**](UserCredentialsDto.md) |  | [optional] 

## Methods

### NewUserDto

`func NewUserDto() *UserDto`

NewUserDto instantiates a new UserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDtoWithDefaults

`func NewUserDtoWithDefaults() *UserDto`

NewUserDtoWithDefaults instantiates a new UserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *UserDto) GetProfile() UserProfileDto`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserDto) GetProfileOk() (*UserProfileDto, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserDto) SetProfile(v UserProfileDto)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserDto) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetCredentials

`func (o *UserDto) GetCredentials() UserCredentialsDto`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UserDto) GetCredentialsOk() (*UserCredentialsDto, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UserDto) SetCredentials(v UserCredentialsDto)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UserDto) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


