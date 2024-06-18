# CamundaFormRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | The key of the Camunda Form. | [optional] 
**Binding** | Pointer to **NullableString** | The binding of the Camunda Form. Can be &#x60;latest&#x60;, &#x60;deployment&#x60; or &#x60;version&#x60;. | [optional] 
**Version** | Pointer to **NullableInt32** | The specific version of a Camunda Form. This property is only set if &#x60;binding&#x60; is &#x60;version&#x60;. | [optional] 

## Methods

### NewCamundaFormRef

`func NewCamundaFormRef() *CamundaFormRef`

NewCamundaFormRef instantiates a new CamundaFormRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCamundaFormRefWithDefaults

`func NewCamundaFormRefWithDefaults() *CamundaFormRef`

NewCamundaFormRefWithDefaults instantiates a new CamundaFormRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CamundaFormRef) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CamundaFormRef) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CamundaFormRef) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CamundaFormRef) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CamundaFormRef) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CamundaFormRef) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetBinding

`func (o *CamundaFormRef) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *CamundaFormRef) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *CamundaFormRef) SetBinding(v string)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *CamundaFormRef) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### SetBindingNil

`func (o *CamundaFormRef) SetBindingNil(b bool)`

 SetBindingNil sets the value for Binding to be an explicit nil

### UnsetBinding
`func (o *CamundaFormRef) UnsetBinding()`

UnsetBinding ensures that no value is present for Binding, not even an explicit nil
### GetVersion

`func (o *CamundaFormRef) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CamundaFormRef) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CamundaFormRef) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CamundaFormRef) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *CamundaFormRef) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *CamundaFormRef) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


