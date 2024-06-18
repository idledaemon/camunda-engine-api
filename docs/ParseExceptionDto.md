# ParseExceptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | An exception class indicating the occurred error. | [optional] 
**Message** | Pointer to **NullableString** | A detailed message of the error. | [optional] 
**Code** | Pointer to **float32** | The code allows your client application to identify the error in an automated fashion. You can look up the meaning of all built-in codes and learn how to add custom codes in the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/error-handling/#exception-codes). | [optional] 
**Details** | Pointer to [**map[string]ResourceReportDto**](ResourceReportDto.md) | A JSON Object containing list of errors and warnings occurred during deployment. | [optional] 

## Methods

### NewParseExceptionDto

`func NewParseExceptionDto() *ParseExceptionDto`

NewParseExceptionDto instantiates a new ParseExceptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseExceptionDtoWithDefaults

`func NewParseExceptionDtoWithDefaults() *ParseExceptionDto`

NewParseExceptionDtoWithDefaults instantiates a new ParseExceptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParseExceptionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParseExceptionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParseExceptionDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ParseExceptionDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ParseExceptionDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ParseExceptionDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMessage

`func (o *ParseExceptionDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ParseExceptionDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ParseExceptionDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ParseExceptionDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ParseExceptionDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ParseExceptionDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCode

`func (o *ParseExceptionDto) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ParseExceptionDto) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ParseExceptionDto) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ParseExceptionDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *ParseExceptionDto) GetDetails() map[string]ResourceReportDto`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ParseExceptionDto) GetDetailsOk() (*map[string]ResourceReportDto, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ParseExceptionDto) SetDetails(v map[string]ResourceReportDto)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ParseExceptionDto) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ParseExceptionDto) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ParseExceptionDto) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


