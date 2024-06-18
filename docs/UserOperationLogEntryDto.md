# UserOperationLogEntryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The unique identifier of this log entry. | [optional] 
**UserId** | Pointer to **NullableString** | The user who performed this operation. | [optional] 
**Timestamp** | Pointer to **NullableTime** | Timestamp of this operation. | [optional] 
**OperationId** | Pointer to **NullableString** | The unique identifier of this operation. A composite operation that changes multiple properties has a common &#x60;operationId&#x60;. | [optional] 
**OperationType** | Pointer to **NullableString** | The type of this operation, e.g., &#x60;Assign&#x60;, &#x60;Claim&#x60; and so on. | [optional] 
**EntityType** | Pointer to **NullableString** | The type of the entity on which this operation was executed, e.g., &#x60;Task&#x60; or &#x60;Attachment&#x60;. | [optional] 
**Category** | Pointer to **NullableString** | The name of the category this operation was associated with, e.g., &#x60;TaskWorker&#x60; or &#x60;Admin&#x60;. | [optional] 
**Annotation** | Pointer to **NullableString** | An arbitrary annotation set by a user for auditing reasons. | [optional] 
**Property** | Pointer to **NullableString** | The property changed by this operation. | [optional] 
**OrgValue** | Pointer to **NullableString** | The original value of the changed property. | [optional] 
**NewValue** | Pointer to **NullableString** | The new value of the changed property. | [optional] 
**DeploymentId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this deployment. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this process definition. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to process definitions with this key. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this process instance. | [optional] 
**ExecutionId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this execution. | [optional] 
**CaseDefinitionId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this case definition. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this case instance. | [optional] 
**CaseExecutionId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this case execution. | [optional] 
**TaskId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this task. | [optional] 
**ExternalTaskId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this external task. | [optional] 
**BatchId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this batch. | [optional] 
**JobId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this job. | [optional] 
**JobDefinitionId** | Pointer to **NullableString** | If not &#x60;null&#x60;, the operation is restricted to entities in relation to this job definition. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the entry should be removed by the History Cleanup job. [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process containing this entry. | [optional] 

## Methods

### NewUserOperationLogEntryDto

`func NewUserOperationLogEntryDto() *UserOperationLogEntryDto`

NewUserOperationLogEntryDto instantiates a new UserOperationLogEntryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOperationLogEntryDtoWithDefaults

`func NewUserOperationLogEntryDtoWithDefaults() *UserOperationLogEntryDto`

NewUserOperationLogEntryDtoWithDefaults instantiates a new UserOperationLogEntryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserOperationLogEntryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserOperationLogEntryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserOperationLogEntryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserOperationLogEntryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserOperationLogEntryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserOperationLogEntryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUserId

`func (o *UserOperationLogEntryDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserOperationLogEntryDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserOperationLogEntryDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserOperationLogEntryDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UserOperationLogEntryDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UserOperationLogEntryDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTimestamp

`func (o *UserOperationLogEntryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserOperationLogEntryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserOperationLogEntryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UserOperationLogEntryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *UserOperationLogEntryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *UserOperationLogEntryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetOperationId

`func (o *UserOperationLogEntryDto) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *UserOperationLogEntryDto) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *UserOperationLogEntryDto) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *UserOperationLogEntryDto) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### SetOperationIdNil

`func (o *UserOperationLogEntryDto) SetOperationIdNil(b bool)`

 SetOperationIdNil sets the value for OperationId to be an explicit nil

### UnsetOperationId
`func (o *UserOperationLogEntryDto) UnsetOperationId()`

UnsetOperationId ensures that no value is present for OperationId, not even an explicit nil
### GetOperationType

`func (o *UserOperationLogEntryDto) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *UserOperationLogEntryDto) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *UserOperationLogEntryDto) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *UserOperationLogEntryDto) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationTypeNil

`func (o *UserOperationLogEntryDto) SetOperationTypeNil(b bool)`

 SetOperationTypeNil sets the value for OperationType to be an explicit nil

### UnsetOperationType
`func (o *UserOperationLogEntryDto) UnsetOperationType()`

UnsetOperationType ensures that no value is present for OperationType, not even an explicit nil
### GetEntityType

`func (o *UserOperationLogEntryDto) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *UserOperationLogEntryDto) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *UserOperationLogEntryDto) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *UserOperationLogEntryDto) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *UserOperationLogEntryDto) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *UserOperationLogEntryDto) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetCategory

`func (o *UserOperationLogEntryDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UserOperationLogEntryDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UserOperationLogEntryDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UserOperationLogEntryDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UserOperationLogEntryDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UserOperationLogEntryDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetAnnotation

`func (o *UserOperationLogEntryDto) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *UserOperationLogEntryDto) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *UserOperationLogEntryDto) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *UserOperationLogEntryDto) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotationNil

`func (o *UserOperationLogEntryDto) SetAnnotationNil(b bool)`

 SetAnnotationNil sets the value for Annotation to be an explicit nil

### UnsetAnnotation
`func (o *UserOperationLogEntryDto) UnsetAnnotation()`

UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil
### GetProperty

`func (o *UserOperationLogEntryDto) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *UserOperationLogEntryDto) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *UserOperationLogEntryDto) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *UserOperationLogEntryDto) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### SetPropertyNil

`func (o *UserOperationLogEntryDto) SetPropertyNil(b bool)`

 SetPropertyNil sets the value for Property to be an explicit nil

### UnsetProperty
`func (o *UserOperationLogEntryDto) UnsetProperty()`

UnsetProperty ensures that no value is present for Property, not even an explicit nil
### GetOrgValue

`func (o *UserOperationLogEntryDto) GetOrgValue() string`

GetOrgValue returns the OrgValue field if non-nil, zero value otherwise.

### GetOrgValueOk

`func (o *UserOperationLogEntryDto) GetOrgValueOk() (*string, bool)`

GetOrgValueOk returns a tuple with the OrgValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgValue

`func (o *UserOperationLogEntryDto) SetOrgValue(v string)`

SetOrgValue sets OrgValue field to given value.

### HasOrgValue

`func (o *UserOperationLogEntryDto) HasOrgValue() bool`

HasOrgValue returns a boolean if a field has been set.

### SetOrgValueNil

`func (o *UserOperationLogEntryDto) SetOrgValueNil(b bool)`

 SetOrgValueNil sets the value for OrgValue to be an explicit nil

### UnsetOrgValue
`func (o *UserOperationLogEntryDto) UnsetOrgValue()`

UnsetOrgValue ensures that no value is present for OrgValue, not even an explicit nil
### GetNewValue

`func (o *UserOperationLogEntryDto) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *UserOperationLogEntryDto) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *UserOperationLogEntryDto) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *UserOperationLogEntryDto) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### SetNewValueNil

`func (o *UserOperationLogEntryDto) SetNewValueNil(b bool)`

 SetNewValueNil sets the value for NewValue to be an explicit nil

### UnsetNewValue
`func (o *UserOperationLogEntryDto) UnsetNewValue()`

UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
### GetDeploymentId

`func (o *UserOperationLogEntryDto) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *UserOperationLogEntryDto) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *UserOperationLogEntryDto) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *UserOperationLogEntryDto) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *UserOperationLogEntryDto) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *UserOperationLogEntryDto) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetProcessDefinitionId

`func (o *UserOperationLogEntryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *UserOperationLogEntryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *UserOperationLogEntryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *UserOperationLogEntryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *UserOperationLogEntryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *UserOperationLogEntryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *UserOperationLogEntryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *UserOperationLogEntryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *UserOperationLogEntryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *UserOperationLogEntryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *UserOperationLogEntryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *UserOperationLogEntryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessInstanceId

`func (o *UserOperationLogEntryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *UserOperationLogEntryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *UserOperationLogEntryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *UserOperationLogEntryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *UserOperationLogEntryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *UserOperationLogEntryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetExecutionId

`func (o *UserOperationLogEntryDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *UserOperationLogEntryDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *UserOperationLogEntryDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *UserOperationLogEntryDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *UserOperationLogEntryDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *UserOperationLogEntryDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetCaseDefinitionId

`func (o *UserOperationLogEntryDto) GetCaseDefinitionId() string`

GetCaseDefinitionId returns the CaseDefinitionId field if non-nil, zero value otherwise.

### GetCaseDefinitionIdOk

`func (o *UserOperationLogEntryDto) GetCaseDefinitionIdOk() (*string, bool)`

GetCaseDefinitionIdOk returns a tuple with the CaseDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseDefinitionId

`func (o *UserOperationLogEntryDto) SetCaseDefinitionId(v string)`

SetCaseDefinitionId sets CaseDefinitionId field to given value.

### HasCaseDefinitionId

`func (o *UserOperationLogEntryDto) HasCaseDefinitionId() bool`

HasCaseDefinitionId returns a boolean if a field has been set.

### SetCaseDefinitionIdNil

`func (o *UserOperationLogEntryDto) SetCaseDefinitionIdNil(b bool)`

 SetCaseDefinitionIdNil sets the value for CaseDefinitionId to be an explicit nil

### UnsetCaseDefinitionId
`func (o *UserOperationLogEntryDto) UnsetCaseDefinitionId()`

UnsetCaseDefinitionId ensures that no value is present for CaseDefinitionId, not even an explicit nil
### GetCaseInstanceId

`func (o *UserOperationLogEntryDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *UserOperationLogEntryDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *UserOperationLogEntryDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *UserOperationLogEntryDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *UserOperationLogEntryDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *UserOperationLogEntryDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetCaseExecutionId

`func (o *UserOperationLogEntryDto) GetCaseExecutionId() string`

GetCaseExecutionId returns the CaseExecutionId field if non-nil, zero value otherwise.

### GetCaseExecutionIdOk

`func (o *UserOperationLogEntryDto) GetCaseExecutionIdOk() (*string, bool)`

GetCaseExecutionIdOk returns a tuple with the CaseExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExecutionId

`func (o *UserOperationLogEntryDto) SetCaseExecutionId(v string)`

SetCaseExecutionId sets CaseExecutionId field to given value.

### HasCaseExecutionId

`func (o *UserOperationLogEntryDto) HasCaseExecutionId() bool`

HasCaseExecutionId returns a boolean if a field has been set.

### SetCaseExecutionIdNil

`func (o *UserOperationLogEntryDto) SetCaseExecutionIdNil(b bool)`

 SetCaseExecutionIdNil sets the value for CaseExecutionId to be an explicit nil

### UnsetCaseExecutionId
`func (o *UserOperationLogEntryDto) UnsetCaseExecutionId()`

UnsetCaseExecutionId ensures that no value is present for CaseExecutionId, not even an explicit nil
### GetTaskId

`func (o *UserOperationLogEntryDto) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *UserOperationLogEntryDto) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *UserOperationLogEntryDto) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *UserOperationLogEntryDto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *UserOperationLogEntryDto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *UserOperationLogEntryDto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetExternalTaskId

`func (o *UserOperationLogEntryDto) GetExternalTaskId() string`

GetExternalTaskId returns the ExternalTaskId field if non-nil, zero value otherwise.

### GetExternalTaskIdOk

`func (o *UserOperationLogEntryDto) GetExternalTaskIdOk() (*string, bool)`

GetExternalTaskIdOk returns a tuple with the ExternalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTaskId

`func (o *UserOperationLogEntryDto) SetExternalTaskId(v string)`

SetExternalTaskId sets ExternalTaskId field to given value.

### HasExternalTaskId

`func (o *UserOperationLogEntryDto) HasExternalTaskId() bool`

HasExternalTaskId returns a boolean if a field has been set.

### SetExternalTaskIdNil

`func (o *UserOperationLogEntryDto) SetExternalTaskIdNil(b bool)`

 SetExternalTaskIdNil sets the value for ExternalTaskId to be an explicit nil

### UnsetExternalTaskId
`func (o *UserOperationLogEntryDto) UnsetExternalTaskId()`

UnsetExternalTaskId ensures that no value is present for ExternalTaskId, not even an explicit nil
### GetBatchId

`func (o *UserOperationLogEntryDto) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *UserOperationLogEntryDto) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *UserOperationLogEntryDto) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *UserOperationLogEntryDto) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchIdNil

`func (o *UserOperationLogEntryDto) SetBatchIdNil(b bool)`

 SetBatchIdNil sets the value for BatchId to be an explicit nil

### UnsetBatchId
`func (o *UserOperationLogEntryDto) UnsetBatchId()`

UnsetBatchId ensures that no value is present for BatchId, not even an explicit nil
### GetJobId

`func (o *UserOperationLogEntryDto) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *UserOperationLogEntryDto) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *UserOperationLogEntryDto) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *UserOperationLogEntryDto) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *UserOperationLogEntryDto) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *UserOperationLogEntryDto) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobDefinitionId

`func (o *UserOperationLogEntryDto) GetJobDefinitionId() string`

GetJobDefinitionId returns the JobDefinitionId field if non-nil, zero value otherwise.

### GetJobDefinitionIdOk

`func (o *UserOperationLogEntryDto) GetJobDefinitionIdOk() (*string, bool)`

GetJobDefinitionIdOk returns a tuple with the JobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionId

`func (o *UserOperationLogEntryDto) SetJobDefinitionId(v string)`

SetJobDefinitionId sets JobDefinitionId field to given value.

### HasJobDefinitionId

`func (o *UserOperationLogEntryDto) HasJobDefinitionId() bool`

HasJobDefinitionId returns a boolean if a field has been set.

### SetJobDefinitionIdNil

`func (o *UserOperationLogEntryDto) SetJobDefinitionIdNil(b bool)`

 SetJobDefinitionIdNil sets the value for JobDefinitionId to be an explicit nil

### UnsetJobDefinitionId
`func (o *UserOperationLogEntryDto) UnsetJobDefinitionId()`

UnsetJobDefinitionId ensures that no value is present for JobDefinitionId, not even an explicit nil
### GetRemovalTime

`func (o *UserOperationLogEntryDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *UserOperationLogEntryDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *UserOperationLogEntryDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *UserOperationLogEntryDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *UserOperationLogEntryDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *UserOperationLogEntryDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetRootProcessInstanceId

`func (o *UserOperationLogEntryDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *UserOperationLogEntryDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *UserOperationLogEntryDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *UserOperationLogEntryDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *UserOperationLogEntryDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *UserOperationLogEntryDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


