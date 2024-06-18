# ExceptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | An exception class indicating the occurred error. | [optional] 
**Message** | Pointer to **NullableString** | A detailed message of the error. | [optional] 
**Code** | Pointer to **float32** | The code allows your client application to identify the error in an automated fashion. You can look up the meaning of all built-in codes and learn how to add custom codes in the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/error-handling/#exception-codes). | [optional] 

## Methods

### NewExceptionDto

`func NewExceptionDto() *ExceptionDto`

NewExceptionDto instantiates a new ExceptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionDtoWithDefaults

`func NewExceptionDtoWithDefaults() *ExceptionDto`

NewExceptionDtoWithDefaults instantiates a new ExceptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExceptionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExceptionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExceptionDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExceptionDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ExceptionDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExceptionDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMessage

`func (o *ExceptionDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExceptionDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExceptionDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ExceptionDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ExceptionDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ExceptionDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCode

`func (o *ExceptionDto) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExceptionDto) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExceptionDto) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ExceptionDto) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


