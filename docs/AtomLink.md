# AtomLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rel** | Pointer to **NullableString** | The relation of the link to the object that belongs to. | [optional] 
**Href** | Pointer to **NullableString** | The url of the link. | [optional] 
**Method** | Pointer to **NullableString** | The http method. | [optional] 

## Methods

### NewAtomLink

`func NewAtomLink() *AtomLink`

NewAtomLink instantiates a new AtomLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtomLinkWithDefaults

`func NewAtomLinkWithDefaults() *AtomLink`

NewAtomLinkWithDefaults instantiates a new AtomLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRel

`func (o *AtomLink) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *AtomLink) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *AtomLink) SetRel(v string)`

SetRel sets Rel field to given value.

### HasRel

`func (o *AtomLink) HasRel() bool`

HasRel returns a boolean if a field has been set.

### SetRelNil

`func (o *AtomLink) SetRelNil(b bool)`

 SetRelNil sets the value for Rel to be an explicit nil

### UnsetRel
`func (o *AtomLink) UnsetRel()`

UnsetRel ensures that no value is present for Rel, not even an explicit nil
### GetHref

`func (o *AtomLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AtomLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AtomLink) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AtomLink) HasHref() bool`

HasHref returns a boolean if a field has been set.

### SetHrefNil

`func (o *AtomLink) SetHrefNil(b bool)`

 SetHrefNil sets the value for Href to be an explicit nil

### UnsetHref
`func (o *AtomLink) UnsetHref()`

UnsetHref ensures that no value is present for Href, not even an explicit nil
### GetMethod

`func (o *AtomLink) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AtomLink) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AtomLink) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AtomLink) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *AtomLink) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *AtomLink) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


