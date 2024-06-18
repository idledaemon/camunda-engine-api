# CheckPasswordPolicyRuleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Placeholder** | Pointer to **NullableString** | A placeholder string that contains the name of a password policy rule. | [optional] 
**Parameter** | Pointer to **map[string]string** | A map that describes the characteristics of a password policy rule, such as the minimum number of digits. | [optional] 
**Valid** | Pointer to **NullableBool** | &#x60;true&#x60; if the password is compliant with this rule, otherwise &#x60;false&#x60;. | [optional] 

## Methods

### NewCheckPasswordPolicyRuleDto

`func NewCheckPasswordPolicyRuleDto() *CheckPasswordPolicyRuleDto`

NewCheckPasswordPolicyRuleDto instantiates a new CheckPasswordPolicyRuleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckPasswordPolicyRuleDtoWithDefaults

`func NewCheckPasswordPolicyRuleDtoWithDefaults() *CheckPasswordPolicyRuleDto`

NewCheckPasswordPolicyRuleDtoWithDefaults instantiates a new CheckPasswordPolicyRuleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaceholder

`func (o *CheckPasswordPolicyRuleDto) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *CheckPasswordPolicyRuleDto) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *CheckPasswordPolicyRuleDto) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *CheckPasswordPolicyRuleDto) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *CheckPasswordPolicyRuleDto) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *CheckPasswordPolicyRuleDto) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetParameter

`func (o *CheckPasswordPolicyRuleDto) GetParameter() map[string]string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *CheckPasswordPolicyRuleDto) GetParameterOk() (*map[string]string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *CheckPasswordPolicyRuleDto) SetParameter(v map[string]string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *CheckPasswordPolicyRuleDto) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetValid

`func (o *CheckPasswordPolicyRuleDto) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *CheckPasswordPolicyRuleDto) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *CheckPasswordPolicyRuleDto) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *CheckPasswordPolicyRuleDto) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValidNil

`func (o *CheckPasswordPolicyRuleDto) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *CheckPasswordPolicyRuleDto) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


