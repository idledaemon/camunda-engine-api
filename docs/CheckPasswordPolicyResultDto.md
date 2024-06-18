# CheckPasswordPolicyResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]PasswordPolicyRuleDto**](PasswordPolicyRuleDto.md) | An array of password policy rules. Each element of the array is representing one rule of the policy. | [optional] 
**Valid** | Pointer to **NullableBool** | &#x60;true&#x60; if the password is compliant with the policy, otherwise &#x60;false&#x60;. | [optional] 

## Methods

### NewCheckPasswordPolicyResultDto

`func NewCheckPasswordPolicyResultDto() *CheckPasswordPolicyResultDto`

NewCheckPasswordPolicyResultDto instantiates a new CheckPasswordPolicyResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckPasswordPolicyResultDtoWithDefaults

`func NewCheckPasswordPolicyResultDtoWithDefaults() *CheckPasswordPolicyResultDto`

NewCheckPasswordPolicyResultDtoWithDefaults instantiates a new CheckPasswordPolicyResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *CheckPasswordPolicyResultDto) GetRules() []PasswordPolicyRuleDto`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CheckPasswordPolicyResultDto) GetRulesOk() (*[]PasswordPolicyRuleDto, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CheckPasswordPolicyResultDto) SetRules(v []PasswordPolicyRuleDto)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CheckPasswordPolicyResultDto) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *CheckPasswordPolicyResultDto) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *CheckPasswordPolicyResultDto) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetValid

`func (o *CheckPasswordPolicyResultDto) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *CheckPasswordPolicyResultDto) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *CheckPasswordPolicyResultDto) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *CheckPasswordPolicyResultDto) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValidNil

`func (o *CheckPasswordPolicyResultDto) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *CheckPasswordPolicyResultDto) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


