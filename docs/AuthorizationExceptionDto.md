# AuthorizationExceptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | An exception class indicating the occurred error. | [optional] 
**Message** | Pointer to **NullableString** | A detailed message of the error. | [optional] 
**Code** | Pointer to **float32** | The code allows your client application to identify the error in an automated fashion. You can look up the meaning of all built-in codes and learn how to add custom codes in the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/error-handling/#exception-codes). | [optional] 
**UserId** | Pointer to **NullableString** | The id of the user that does not have expected permissions | [optional] 
**MissingAuthorizations** | Pointer to [**[]MissingAuthorizationDto**](MissingAuthorizationDto.md) |  | [optional] 

## Methods

### NewAuthorizationExceptionDto

`func NewAuthorizationExceptionDto() *AuthorizationExceptionDto`

NewAuthorizationExceptionDto instantiates a new AuthorizationExceptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationExceptionDtoWithDefaults

`func NewAuthorizationExceptionDtoWithDefaults() *AuthorizationExceptionDto`

NewAuthorizationExceptionDtoWithDefaults instantiates a new AuthorizationExceptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizationExceptionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizationExceptionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizationExceptionDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthorizationExceptionDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AuthorizationExceptionDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AuthorizationExceptionDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMessage

`func (o *AuthorizationExceptionDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthorizationExceptionDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthorizationExceptionDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthorizationExceptionDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AuthorizationExceptionDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AuthorizationExceptionDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCode

`func (o *AuthorizationExceptionDto) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthorizationExceptionDto) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthorizationExceptionDto) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthorizationExceptionDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetUserId

`func (o *AuthorizationExceptionDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthorizationExceptionDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthorizationExceptionDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthorizationExceptionDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *AuthorizationExceptionDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AuthorizationExceptionDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetMissingAuthorizations

`func (o *AuthorizationExceptionDto) GetMissingAuthorizations() []MissingAuthorizationDto`

GetMissingAuthorizations returns the MissingAuthorizations field if non-nil, zero value otherwise.

### GetMissingAuthorizationsOk

`func (o *AuthorizationExceptionDto) GetMissingAuthorizationsOk() (*[]MissingAuthorizationDto, bool)`

GetMissingAuthorizationsOk returns a tuple with the MissingAuthorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingAuthorizations

`func (o *AuthorizationExceptionDto) SetMissingAuthorizations(v []MissingAuthorizationDto)`

SetMissingAuthorizations sets MissingAuthorizations field to given value.

### HasMissingAuthorizations

`func (o *AuthorizationExceptionDto) HasMissingAuthorizations() bool`

HasMissingAuthorizations returns a boolean if a field has been set.

### SetMissingAuthorizationsNil

`func (o *AuthorizationExceptionDto) SetMissingAuthorizationsNil(b bool)`

 SetMissingAuthorizationsNil sets the value for MissingAuthorizations to be an explicit nil

### UnsetMissingAuthorizations
`func (o *AuthorizationExceptionDto) UnsetMissingAuthorizations()`

UnsetMissingAuthorizations ensures that no value is present for MissingAuthorizations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


