# PasswordPolicyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]PasswordPolicyRuleDto**](PasswordPolicyRuleDto.md) | An array of password policy rules. Each element of the array is representing one rule of the policy. | [optional] 

## Methods

### NewPasswordPolicyDto

`func NewPasswordPolicyDto() *PasswordPolicyDto`

NewPasswordPolicyDto instantiates a new PasswordPolicyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyDtoWithDefaults

`func NewPasswordPolicyDtoWithDefaults() *PasswordPolicyDto`

NewPasswordPolicyDtoWithDefaults instantiates a new PasswordPolicyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *PasswordPolicyDto) GetRules() []PasswordPolicyRuleDto`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PasswordPolicyDto) GetRulesOk() (*[]PasswordPolicyRuleDto, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PasswordPolicyDto) SetRules(v []PasswordPolicyRuleDto)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PasswordPolicyDto) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *PasswordPolicyDto) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *PasswordPolicyDto) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


