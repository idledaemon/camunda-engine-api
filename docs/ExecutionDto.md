# ExecutionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the Execution. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the root of the execution tree representing the process instance. | [optional] 
**Ended** | Pointer to **NullableBool** | Indicates if the execution is ended. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant this execution belongs to. Can be &#x60;null&#x60; if the execution belongs to no single tenant. | [optional] 

## Methods

### NewExecutionDto

`func NewExecutionDto() *ExecutionDto`

NewExecutionDto instantiates a new ExecutionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionDtoWithDefaults

`func NewExecutionDtoWithDefaults() *ExecutionDto`

NewExecutionDtoWithDefaults instantiates a new ExecutionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExecutionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExecutionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProcessInstanceId

`func (o *ExecutionDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *ExecutionDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *ExecutionDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *ExecutionDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *ExecutionDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *ExecutionDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetEnded

`func (o *ExecutionDto) GetEnded() bool`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *ExecutionDto) GetEndedOk() (*bool, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *ExecutionDto) SetEnded(v bool)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *ExecutionDto) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### SetEndedNil

`func (o *ExecutionDto) SetEndedNil(b bool)`

 SetEndedNil sets the value for Ended to be an explicit nil

### UnsetEnded
`func (o *ExecutionDto) UnsetEnded()`

UnsetEnded ensures that no value is present for Ended, not even an explicit nil
### GetTenantId

`func (o *ExecutionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExecutionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExecutionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ExecutionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ExecutionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ExecutionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


