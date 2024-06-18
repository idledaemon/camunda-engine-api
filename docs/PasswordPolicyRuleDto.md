# PasswordPolicyRuleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Placeholder** | Pointer to **NullableString** | A placeholder string that contains the name of a password policy rule. | [optional] 
**Parameter** | Pointer to **map[string]string** | A map that describes the characteristics of a password policy rule, such as the minimum number of digits. | [optional] 

## Methods

### NewPasswordPolicyRuleDto

`func NewPasswordPolicyRuleDto() *PasswordPolicyRuleDto`

NewPasswordPolicyRuleDto instantiates a new PasswordPolicyRuleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyRuleDtoWithDefaults

`func NewPasswordPolicyRuleDtoWithDefaults() *PasswordPolicyRuleDto`

NewPasswordPolicyRuleDtoWithDefaults instantiates a new PasswordPolicyRuleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaceholder

`func (o *PasswordPolicyRuleDto) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *PasswordPolicyRuleDto) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *PasswordPolicyRuleDto) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *PasswordPolicyRuleDto) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *PasswordPolicyRuleDto) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *PasswordPolicyRuleDto) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetParameter

`func (o *PasswordPolicyRuleDto) GetParameter() map[string]string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *PasswordPolicyRuleDto) GetParameterOk() (*map[string]string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *PasswordPolicyRuleDto) SetParameter(v map[string]string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *PasswordPolicyRuleDto) HasParameter() bool`

HasParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


