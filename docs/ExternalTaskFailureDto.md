# ExternalTaskFailureDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerId** | Pointer to **string** | **Mandatory.** The ID of the worker who is performing the operation on the external task. If the task is already locked, must match the id of the worker who has most recently locked the task. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | An message indicating the reason of the failure. | [optional] 
**ErrorDetails** | Pointer to **NullableString** | A detailed error description. | [optional] 
**Retries** | Pointer to **NullableInt32** | A number of how often the task should be retried. Must be &gt;&#x3D; 0. If this is 0, an incident is created and the task cannot be fetched anymore unless the retries are increased again. The incident&#39;s message is set to the &#x60;errorMessage&#x60; parameter. | [optional] 
**RetryTimeout** | Pointer to **NullableInt64** | A timeout in milliseconds before the external task becomes available again for fetching. Must be &gt;&#x3D; 0. | [optional] 
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A JSON object containing variable key-value pairs. Each key is a variable name and each value a JSON variable value object with the following properties: | [optional] 
**LocalVariables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A JSON object containing local variable key-value pairs. Local variables are set only in the scope of external task. Each key is a variable name and each value a JSON variable value object with the following properties: | [optional] 

## Methods

### NewExternalTaskFailureDto

`func NewExternalTaskFailureDto() *ExternalTaskFailureDto`

NewExternalTaskFailureDto instantiates a new ExternalTaskFailureDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalTaskFailureDtoWithDefaults

`func NewExternalTaskFailureDtoWithDefaults() *ExternalTaskFailureDto`

NewExternalTaskFailureDtoWithDefaults instantiates a new ExternalTaskFailureDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerId

`func (o *ExternalTaskFailureDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ExternalTaskFailureDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ExternalTaskFailureDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *ExternalTaskFailureDto) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ExternalTaskFailureDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ExternalTaskFailureDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ExternalTaskFailureDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ExternalTaskFailureDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ExternalTaskFailureDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ExternalTaskFailureDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetErrorDetails

`func (o *ExternalTaskFailureDto) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *ExternalTaskFailureDto) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *ExternalTaskFailureDto) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *ExternalTaskFailureDto) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### SetErrorDetailsNil

`func (o *ExternalTaskFailureDto) SetErrorDetailsNil(b bool)`

 SetErrorDetailsNil sets the value for ErrorDetails to be an explicit nil

### UnsetErrorDetails
`func (o *ExternalTaskFailureDto) UnsetErrorDetails()`

UnsetErrorDetails ensures that no value is present for ErrorDetails, not even an explicit nil
### GetRetries

`func (o *ExternalTaskFailureDto) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *ExternalTaskFailureDto) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *ExternalTaskFailureDto) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *ExternalTaskFailureDto) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *ExternalTaskFailureDto) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *ExternalTaskFailureDto) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetRetryTimeout

`func (o *ExternalTaskFailureDto) GetRetryTimeout() int64`

GetRetryTimeout returns the RetryTimeout field if non-nil, zero value otherwise.

### GetRetryTimeoutOk

`func (o *ExternalTaskFailureDto) GetRetryTimeoutOk() (*int64, bool)`

GetRetryTimeoutOk returns a tuple with the RetryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTimeout

`func (o *ExternalTaskFailureDto) SetRetryTimeout(v int64)`

SetRetryTimeout sets RetryTimeout field to given value.

### HasRetryTimeout

`func (o *ExternalTaskFailureDto) HasRetryTimeout() bool`

HasRetryTimeout returns a boolean if a field has been set.

### SetRetryTimeoutNil

`func (o *ExternalTaskFailureDto) SetRetryTimeoutNil(b bool)`

 SetRetryTimeoutNil sets the value for RetryTimeout to be an explicit nil

### UnsetRetryTimeout
`func (o *ExternalTaskFailureDto) UnsetRetryTimeout()`

UnsetRetryTimeout ensures that no value is present for RetryTimeout, not even an explicit nil
### GetVariables

`func (o *ExternalTaskFailureDto) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ExternalTaskFailureDto) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ExternalTaskFailureDto) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ExternalTaskFailureDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *ExternalTaskFailureDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *ExternalTaskFailureDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetLocalVariables

`func (o *ExternalTaskFailureDto) GetLocalVariables() map[string]VariableValueDto`

GetLocalVariables returns the LocalVariables field if non-nil, zero value otherwise.

### GetLocalVariablesOk

`func (o *ExternalTaskFailureDto) GetLocalVariablesOk() (*map[string]VariableValueDto, bool)`

GetLocalVariablesOk returns a tuple with the LocalVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVariables

`func (o *ExternalTaskFailureDto) SetLocalVariables(v map[string]VariableValueDto)`

SetLocalVariables sets LocalVariables field to given value.

### HasLocalVariables

`func (o *ExternalTaskFailureDto) HasLocalVariables() bool`

HasLocalVariables returns a boolean if a field has been set.

### SetLocalVariablesNil

`func (o *ExternalTaskFailureDto) SetLocalVariablesNil(b bool)`

 SetLocalVariablesNil sets the value for LocalVariables to be an explicit nil

### UnsetLocalVariables
`func (o *ExternalTaskFailureDto) UnsetLocalVariables()`

UnsetLocalVariables ensures that no value is present for LocalVariables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


