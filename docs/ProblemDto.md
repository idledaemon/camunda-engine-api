# ProblemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | The message of the problem. | [optional] 
**Line** | Pointer to **NullableInt32** | The line where the problem occurred. | [optional] 
**Column** | Pointer to **NullableInt32** | The column where the problem occurred. | [optional] 
**MainElementId** | Pointer to **NullableString** | The main element id where the problem occurred. | [optional] 
**ElementIds** | Pointer to **[]string** | A list of element id affected by the problem. | [optional] 

## Methods

### NewProblemDto

`func NewProblemDto() *ProblemDto`

NewProblemDto instantiates a new ProblemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemDtoWithDefaults

`func NewProblemDtoWithDefaults() *ProblemDto`

NewProblemDtoWithDefaults instantiates a new ProblemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ProblemDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProblemDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProblemDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProblemDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ProblemDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ProblemDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetLine

`func (o *ProblemDto) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ProblemDto) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ProblemDto) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *ProblemDto) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLineNil

`func (o *ProblemDto) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *ProblemDto) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil
### GetColumn

`func (o *ProblemDto) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *ProblemDto) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *ProblemDto) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *ProblemDto) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### SetColumnNil

`func (o *ProblemDto) SetColumnNil(b bool)`

 SetColumnNil sets the value for Column to be an explicit nil

### UnsetColumn
`func (o *ProblemDto) UnsetColumn()`

UnsetColumn ensures that no value is present for Column, not even an explicit nil
### GetMainElementId

`func (o *ProblemDto) GetMainElementId() string`

GetMainElementId returns the MainElementId field if non-nil, zero value otherwise.

### GetMainElementIdOk

`func (o *ProblemDto) GetMainElementIdOk() (*string, bool)`

GetMainElementIdOk returns a tuple with the MainElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainElementId

`func (o *ProblemDto) SetMainElementId(v string)`

SetMainElementId sets MainElementId field to given value.

### HasMainElementId

`func (o *ProblemDto) HasMainElementId() bool`

HasMainElementId returns a boolean if a field has been set.

### SetMainElementIdNil

`func (o *ProblemDto) SetMainElementIdNil(b bool)`

 SetMainElementIdNil sets the value for MainElementId to be an explicit nil

### UnsetMainElementId
`func (o *ProblemDto) UnsetMainElementId()`

UnsetMainElementId ensures that no value is present for MainElementId, not even an explicit nil
### GetElementIds

`func (o *ProblemDto) GetElementIds() []string`

GetElementIds returns the ElementIds field if non-nil, zero value otherwise.

### GetElementIdsOk

`func (o *ProblemDto) GetElementIdsOk() (*[]string, bool)`

GetElementIdsOk returns a tuple with the ElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIds

`func (o *ProblemDto) SetElementIds(v []string)`

SetElementIds sets ElementIds field to given value.

### HasElementIds

`func (o *ProblemDto) HasElementIds() bool`

HasElementIds returns a boolean if a field has been set.

### SetElementIdsNil

`func (o *ProblemDto) SetElementIdsNil(b bool)`

 SetElementIdsNil sets the value for ElementIds to be an explicit nil

### UnsetElementIds
`func (o *ProblemDto) UnsetElementIds()`

UnsetElementIds ensures that no value is present for ElementIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


