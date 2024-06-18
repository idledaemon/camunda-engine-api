# FormDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | The form key. | [optional] 
**CamundaFormRef** | Pointer to [**CamundaFormRef**](CamundaFormRef.md) |  | [optional] 
**ContextPath** | Pointer to **NullableString** | The context path of the process application. If the task (or the process definition) does not belong to a process application deployment or a process definition at all, this property is not set. | [optional] 

## Methods

### NewFormDto

`func NewFormDto() *FormDto`

NewFormDto instantiates a new FormDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormDtoWithDefaults

`func NewFormDtoWithDefaults() *FormDto`

NewFormDtoWithDefaults instantiates a new FormDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FormDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FormDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *FormDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *FormDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetCamundaFormRef

`func (o *FormDto) GetCamundaFormRef() CamundaFormRef`

GetCamundaFormRef returns the CamundaFormRef field if non-nil, zero value otherwise.

### GetCamundaFormRefOk

`func (o *FormDto) GetCamundaFormRefOk() (*CamundaFormRef, bool)`

GetCamundaFormRefOk returns a tuple with the CamundaFormRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCamundaFormRef

`func (o *FormDto) SetCamundaFormRef(v CamundaFormRef)`

SetCamundaFormRef sets CamundaFormRef field to given value.

### HasCamundaFormRef

`func (o *FormDto) HasCamundaFormRef() bool`

HasCamundaFormRef returns a boolean if a field has been set.

### GetContextPath

`func (o *FormDto) GetContextPath() string`

GetContextPath returns the ContextPath field if non-nil, zero value otherwise.

### GetContextPathOk

`func (o *FormDto) GetContextPathOk() (*string, bool)`

GetContextPathOk returns a tuple with the ContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextPath

`func (o *FormDto) SetContextPath(v string)`

SetContextPath sets ContextPath field to given value.

### HasContextPath

`func (o *FormDto) HasContextPath() bool`

HasContextPath returns a boolean if a field has been set.

### SetContextPathNil

`func (o *FormDto) SetContextPathNil(b bool)`

 SetContextPathNil sets the value for ContextPath to be an explicit nil

### UnsetContextPath
`func (o *FormDto) UnsetContextPath()`

UnsetContextPath ensures that no value is present for ContextPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


