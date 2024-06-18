# TriggerVariableValueDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **interface{}** | Can be any value - string, number, boolean, array or object.  **Note**: Not every endpoint supports every type. | [optional] 
**Type** | Pointer to **NullableString** | The value type of the variable. | [optional] 
**ValueInfo** | Pointer to **map[string]interface{}** | A JSON object containing additional, value-type-dependent properties. For serialized variables of type Object, the following properties can be provided:  * &#x60;objectTypeName&#x60;: A string representation of the object&#39;s type name. * &#x60;serializationDataFormat&#x60;: The serialization format used to store the variable.  For serialized variables of type File, the following properties can be provided:  * &#x60;filename&#x60;: The name of the file. This is not the variable name but the name that will be used when downloading the file again. * &#x60;mimetype&#x60;: The MIME type of the file that is being uploaded. * &#x60;encoding&#x60;: The encoding of the file that is being uploaded.  The following property can be provided for all value types:  * &#x60;transient&#x60;: Indicates whether the variable should be transient or not. See [documentation](https://docs.camunda.org/manual/7.21/user-guide/process-engine/variables#transient-variables) for more informations. (Not applicable for &#x60;decision-definition&#x60;, &#x60; /process-instance/variables-async&#x60;, and &#x60;/migration/executeAsync&#x60; endpoints) | [optional] 
**Local** | Pointer to **NullableBool** | Indicates whether the variable should be a local variable or not. If set to true, the variable becomes a local variable of the execution entering the target activity. | [optional] 

## Methods

### NewTriggerVariableValueDto

`func NewTriggerVariableValueDto() *TriggerVariableValueDto`

NewTriggerVariableValueDto instantiates a new TriggerVariableValueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerVariableValueDtoWithDefaults

`func NewTriggerVariableValueDtoWithDefaults() *TriggerVariableValueDto`

NewTriggerVariableValueDtoWithDefaults instantiates a new TriggerVariableValueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *TriggerVariableValueDto) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TriggerVariableValueDto) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TriggerVariableValueDto) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *TriggerVariableValueDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TriggerVariableValueDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TriggerVariableValueDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetType

`func (o *TriggerVariableValueDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerVariableValueDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerVariableValueDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TriggerVariableValueDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TriggerVariableValueDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TriggerVariableValueDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValueInfo

`func (o *TriggerVariableValueDto) GetValueInfo() map[string]interface{}`

GetValueInfo returns the ValueInfo field if non-nil, zero value otherwise.

### GetValueInfoOk

`func (o *TriggerVariableValueDto) GetValueInfoOk() (*map[string]interface{}, bool)`

GetValueInfoOk returns a tuple with the ValueInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInfo

`func (o *TriggerVariableValueDto) SetValueInfo(v map[string]interface{})`

SetValueInfo sets ValueInfo field to given value.

### HasValueInfo

`func (o *TriggerVariableValueDto) HasValueInfo() bool`

HasValueInfo returns a boolean if a field has been set.

### GetLocal

`func (o *TriggerVariableValueDto) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *TriggerVariableValueDto) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *TriggerVariableValueDto) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *TriggerVariableValueDto) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### SetLocalNil

`func (o *TriggerVariableValueDto) SetLocalNil(b bool)`

 SetLocalNil sets the value for Local to be an explicit nil

### UnsetLocal
`func (o *TriggerVariableValueDto) UnsetLocal()`

UnsetLocal ensures that no value is present for Local, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


