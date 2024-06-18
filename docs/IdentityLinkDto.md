# IdentityLinkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** | The id of the user participating in this link. Either &#x60;userId&#x60; or &#x60;groupId&#x60; is set. | [optional] 
**GroupId** | Pointer to **NullableString** | The id of the group participating in this link. Either &#x60;groupId&#x60; or &#x60;userId&#x60; is set. | [optional] 
**Type** | **NullableString** | The type of the identity link. The value of the this property can be user-defined. The Process Engine provides three pre-defined Identity Link &#x60;type&#x60;s:  * &#x60;candidate&#x60; * &#x60;assignee&#x60; - reserved for the task assignee * &#x60;owner&#x60; - reserved for the task owner  **Note**: When adding or removing an Identity Link, the &#x60;type&#x60; property must be defined. | 

## Methods

### NewIdentityLinkDto

`func NewIdentityLinkDto(type_ NullableString, ) *IdentityLinkDto`

NewIdentityLinkDto instantiates a new IdentityLinkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityLinkDtoWithDefaults

`func NewIdentityLinkDtoWithDefaults() *IdentityLinkDto`

NewIdentityLinkDtoWithDefaults instantiates a new IdentityLinkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *IdentityLinkDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *IdentityLinkDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *IdentityLinkDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *IdentityLinkDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *IdentityLinkDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *IdentityLinkDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetGroupId

`func (o *IdentityLinkDto) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *IdentityLinkDto) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *IdentityLinkDto) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *IdentityLinkDto) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *IdentityLinkDto) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *IdentityLinkDto) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetType

`func (o *IdentityLinkDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityLinkDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityLinkDto) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *IdentityLinkDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *IdentityLinkDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


