# PatchVariablesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modifications** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A JSON object containing variable key-value pairs. | [optional] 
**Deletions** | Pointer to **[]string** | An array of String keys of variables to be deleted. | [optional] 

## Methods

### NewPatchVariablesDto

`func NewPatchVariablesDto() *PatchVariablesDto`

NewPatchVariablesDto instantiates a new PatchVariablesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchVariablesDtoWithDefaults

`func NewPatchVariablesDtoWithDefaults() *PatchVariablesDto`

NewPatchVariablesDtoWithDefaults instantiates a new PatchVariablesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModifications

`func (o *PatchVariablesDto) GetModifications() map[string]VariableValueDto`

GetModifications returns the Modifications field if non-nil, zero value otherwise.

### GetModificationsOk

`func (o *PatchVariablesDto) GetModificationsOk() (*map[string]VariableValueDto, bool)`

GetModificationsOk returns a tuple with the Modifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifications

`func (o *PatchVariablesDto) SetModifications(v map[string]VariableValueDto)`

SetModifications sets Modifications field to given value.

### HasModifications

`func (o *PatchVariablesDto) HasModifications() bool`

HasModifications returns a boolean if a field has been set.

### SetModificationsNil

`func (o *PatchVariablesDto) SetModificationsNil(b bool)`

 SetModificationsNil sets the value for Modifications to be an explicit nil

### UnsetModifications
`func (o *PatchVariablesDto) UnsetModifications()`

UnsetModifications ensures that no value is present for Modifications, not even an explicit nil
### GetDeletions

`func (o *PatchVariablesDto) GetDeletions() []string`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *PatchVariablesDto) GetDeletionsOk() (*[]string, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *PatchVariablesDto) SetDeletions(v []string)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *PatchVariablesDto) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### SetDeletionsNil

`func (o *PatchVariablesDto) SetDeletionsNil(b bool)`

 SetDeletionsNil sets the value for Deletions to be an explicit nil

### UnsetDeletions
`func (o *PatchVariablesDto) UnsetDeletions()`

UnsetDeletions ensures that no value is present for Deletions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


