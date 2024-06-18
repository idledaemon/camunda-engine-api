# ExternalTaskBpmnError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **NullableString** | An error code that indicates the predefined error. It is used to identify the BPMN error handler. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | An error message that describes the error. | [optional] 
**Variables** | Pointer to [**map[string]VariableValueDto**](VariableValueDto.md) | A JSON object containing variable key-value pairs. | [optional] 
**WorkerId** | Pointer to **NullableString** | The id of the worker that reports the failure. Must match the id of the worker who has most recently locked the task. | [optional] 

## Methods

### NewExternalTaskBpmnError

`func NewExternalTaskBpmnError() *ExternalTaskBpmnError`

NewExternalTaskBpmnError instantiates a new ExternalTaskBpmnError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalTaskBpmnErrorWithDefaults

`func NewExternalTaskBpmnErrorWithDefaults() *ExternalTaskBpmnError`

NewExternalTaskBpmnErrorWithDefaults instantiates a new ExternalTaskBpmnError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ExternalTaskBpmnError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ExternalTaskBpmnError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ExternalTaskBpmnError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ExternalTaskBpmnError) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ExternalTaskBpmnError) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ExternalTaskBpmnError) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *ExternalTaskBpmnError) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ExternalTaskBpmnError) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ExternalTaskBpmnError) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ExternalTaskBpmnError) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ExternalTaskBpmnError) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ExternalTaskBpmnError) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetVariables

`func (o *ExternalTaskBpmnError) GetVariables() map[string]VariableValueDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ExternalTaskBpmnError) GetVariablesOk() (*map[string]VariableValueDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ExternalTaskBpmnError) SetVariables(v map[string]VariableValueDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ExternalTaskBpmnError) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *ExternalTaskBpmnError) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *ExternalTaskBpmnError) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetWorkerId

`func (o *ExternalTaskBpmnError) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ExternalTaskBpmnError) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ExternalTaskBpmnError) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *ExternalTaskBpmnError) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### SetWorkerIdNil

`func (o *ExternalTaskBpmnError) SetWorkerIdNil(b bool)`

 SetWorkerIdNil sets the value for WorkerId to be an explicit nil

### UnsetWorkerId
`func (o *ExternalTaskBpmnError) UnsetWorkerId()`

UnsetWorkerId ensures that no value is present for WorkerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


