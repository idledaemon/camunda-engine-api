# IdentityServiceUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the user. | [optional] 
**FirstName** | Pointer to **NullableString** | The firstname of the user. | [optional] 
**LastName** | Pointer to **NullableString** | The lastname of the user. | [optional] 
**DisplayName** | Pointer to **NullableString** | The displayName is generated from the id or firstName and lastName if available. | [optional] 

## Methods

### NewIdentityServiceUserDto

`func NewIdentityServiceUserDto() *IdentityServiceUserDto`

NewIdentityServiceUserDto instantiates a new IdentityServiceUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityServiceUserDtoWithDefaults

`func NewIdentityServiceUserDtoWithDefaults() *IdentityServiceUserDto`

NewIdentityServiceUserDtoWithDefaults instantiates a new IdentityServiceUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityServiceUserDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityServiceUserDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityServiceUserDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityServiceUserDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IdentityServiceUserDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IdentityServiceUserDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetFirstName

`func (o *IdentityServiceUserDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IdentityServiceUserDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IdentityServiceUserDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IdentityServiceUserDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *IdentityServiceUserDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *IdentityServiceUserDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *IdentityServiceUserDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IdentityServiceUserDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IdentityServiceUserDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IdentityServiceUserDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *IdentityServiceUserDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *IdentityServiceUserDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDisplayName

`func (o *IdentityServiceUserDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityServiceUserDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityServiceUserDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityServiceUserDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *IdentityServiceUserDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *IdentityServiceUserDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


