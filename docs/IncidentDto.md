# IncidentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the incident. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition this incident is associated with. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance this incident is associated with. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The id of the execution this incident is associated with. | [optional] 
**IncidentTimestamp** | Pointer to **NullableTime** | The time this incident happened. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**IncidentType** | Pointer to **NullableString** | The type of incident, for example: &#x60;failedJobs&#x60; will be returned in case of an incident which identified a failed job during the execution of a process instance. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | [optional] 
**ActivityId** | Pointer to **NullableString** | The id of the activity this incident is associated with. | [optional] 
**FailedActivityId** | Pointer to **NullableString** | The id of the activity on which the last exception occurred. | [optional] 
**CauseIncidentId** | Pointer to **NullableString** | The id of the associated cause incident which has been triggered. | [optional] 
**RootCauseIncidentId** | Pointer to **NullableString** | The id of the associated root cause incident which has been triggered. | [optional] 
**Configuration** | Pointer to **NullableString** | The payload of this incident. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant this incident is associated with. | [optional] 
**IncidentMessage** | Pointer to **NullableString** | The message of this incident. | [optional] 
**JobDefinitionId** | Pointer to **NullableString** | The job definition id the incident is associated with. | [optional] 
**Annotation** | Pointer to **NullableString** | The annotation set to the incident. | [optional] 

## Methods

### NewIncidentDto

`func NewIncidentDto() *IncidentDto`

NewIncidentDto instantiates a new IncidentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentDtoWithDefaults

`func NewIncidentDtoWithDefaults() *IncidentDto`

NewIncidentDtoWithDefaults instantiates a new IncidentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IncidentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IncidentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IncidentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProcessDefinitionId

`func (o *IncidentDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *IncidentDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *IncidentDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *IncidentDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *IncidentDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *IncidentDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessInstanceId

`func (o *IncidentDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *IncidentDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *IncidentDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *IncidentDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *IncidentDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *IncidentDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetExecutionId

`func (o *IncidentDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *IncidentDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *IncidentDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *IncidentDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *IncidentDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *IncidentDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetIncidentTimestamp

`func (o *IncidentDto) GetIncidentTimestamp() time.Time`

GetIncidentTimestamp returns the IncidentTimestamp field if non-nil, zero value otherwise.

### GetIncidentTimestampOk

`func (o *IncidentDto) GetIncidentTimestampOk() (*time.Time, bool)`

GetIncidentTimestampOk returns a tuple with the IncidentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentTimestamp

`func (o *IncidentDto) SetIncidentTimestamp(v time.Time)`

SetIncidentTimestamp sets IncidentTimestamp field to given value.

### HasIncidentTimestamp

`func (o *IncidentDto) HasIncidentTimestamp() bool`

HasIncidentTimestamp returns a boolean if a field has been set.

### SetIncidentTimestampNil

`func (o *IncidentDto) SetIncidentTimestampNil(b bool)`

 SetIncidentTimestampNil sets the value for IncidentTimestamp to be an explicit nil

### UnsetIncidentTimestamp
`func (o *IncidentDto) UnsetIncidentTimestamp()`

UnsetIncidentTimestamp ensures that no value is present for IncidentTimestamp, not even an explicit nil
### GetIncidentType

`func (o *IncidentDto) GetIncidentType() string`

GetIncidentType returns the IncidentType field if non-nil, zero value otherwise.

### GetIncidentTypeOk

`func (o *IncidentDto) GetIncidentTypeOk() (*string, bool)`

GetIncidentTypeOk returns a tuple with the IncidentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentType

`func (o *IncidentDto) SetIncidentType(v string)`

SetIncidentType sets IncidentType field to given value.

### HasIncidentType

`func (o *IncidentDto) HasIncidentType() bool`

HasIncidentType returns a boolean if a field has been set.

### SetIncidentTypeNil

`func (o *IncidentDto) SetIncidentTypeNil(b bool)`

 SetIncidentTypeNil sets the value for IncidentType to be an explicit nil

### UnsetIncidentType
`func (o *IncidentDto) UnsetIncidentType()`

UnsetIncidentType ensures that no value is present for IncidentType, not even an explicit nil
### GetActivityId

`func (o *IncidentDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *IncidentDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *IncidentDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *IncidentDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *IncidentDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *IncidentDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetFailedActivityId

`func (o *IncidentDto) GetFailedActivityId() string`

GetFailedActivityId returns the FailedActivityId field if non-nil, zero value otherwise.

### GetFailedActivityIdOk

`func (o *IncidentDto) GetFailedActivityIdOk() (*string, bool)`

GetFailedActivityIdOk returns a tuple with the FailedActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedActivityId

`func (o *IncidentDto) SetFailedActivityId(v string)`

SetFailedActivityId sets FailedActivityId field to given value.

### HasFailedActivityId

`func (o *IncidentDto) HasFailedActivityId() bool`

HasFailedActivityId returns a boolean if a field has been set.

### SetFailedActivityIdNil

`func (o *IncidentDto) SetFailedActivityIdNil(b bool)`

 SetFailedActivityIdNil sets the value for FailedActivityId to be an explicit nil

### UnsetFailedActivityId
`func (o *IncidentDto) UnsetFailedActivityId()`

UnsetFailedActivityId ensures that no value is present for FailedActivityId, not even an explicit nil
### GetCauseIncidentId

`func (o *IncidentDto) GetCauseIncidentId() string`

GetCauseIncidentId returns the CauseIncidentId field if non-nil, zero value otherwise.

### GetCauseIncidentIdOk

`func (o *IncidentDto) GetCauseIncidentIdOk() (*string, bool)`

GetCauseIncidentIdOk returns a tuple with the CauseIncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseIncidentId

`func (o *IncidentDto) SetCauseIncidentId(v string)`

SetCauseIncidentId sets CauseIncidentId field to given value.

### HasCauseIncidentId

`func (o *IncidentDto) HasCauseIncidentId() bool`

HasCauseIncidentId returns a boolean if a field has been set.

### SetCauseIncidentIdNil

`func (o *IncidentDto) SetCauseIncidentIdNil(b bool)`

 SetCauseIncidentIdNil sets the value for CauseIncidentId to be an explicit nil

### UnsetCauseIncidentId
`func (o *IncidentDto) UnsetCauseIncidentId()`

UnsetCauseIncidentId ensures that no value is present for CauseIncidentId, not even an explicit nil
### GetRootCauseIncidentId

`func (o *IncidentDto) GetRootCauseIncidentId() string`

GetRootCauseIncidentId returns the RootCauseIncidentId field if non-nil, zero value otherwise.

### GetRootCauseIncidentIdOk

`func (o *IncidentDto) GetRootCauseIncidentIdOk() (*string, bool)`

GetRootCauseIncidentIdOk returns a tuple with the RootCauseIncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIncidentId

`func (o *IncidentDto) SetRootCauseIncidentId(v string)`

SetRootCauseIncidentId sets RootCauseIncidentId field to given value.

### HasRootCauseIncidentId

`func (o *IncidentDto) HasRootCauseIncidentId() bool`

HasRootCauseIncidentId returns a boolean if a field has been set.

### SetRootCauseIncidentIdNil

`func (o *IncidentDto) SetRootCauseIncidentIdNil(b bool)`

 SetRootCauseIncidentIdNil sets the value for RootCauseIncidentId to be an explicit nil

### UnsetRootCauseIncidentId
`func (o *IncidentDto) UnsetRootCauseIncidentId()`

UnsetRootCauseIncidentId ensures that no value is present for RootCauseIncidentId, not even an explicit nil
### GetConfiguration

`func (o *IncidentDto) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IncidentDto) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IncidentDto) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *IncidentDto) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *IncidentDto) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *IncidentDto) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetTenantId

`func (o *IncidentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IncidentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IncidentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IncidentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *IncidentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *IncidentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetIncidentMessage

`func (o *IncidentDto) GetIncidentMessage() string`

GetIncidentMessage returns the IncidentMessage field if non-nil, zero value otherwise.

### GetIncidentMessageOk

`func (o *IncidentDto) GetIncidentMessageOk() (*string, bool)`

GetIncidentMessageOk returns a tuple with the IncidentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentMessage

`func (o *IncidentDto) SetIncidentMessage(v string)`

SetIncidentMessage sets IncidentMessage field to given value.

### HasIncidentMessage

`func (o *IncidentDto) HasIncidentMessage() bool`

HasIncidentMessage returns a boolean if a field has been set.

### SetIncidentMessageNil

`func (o *IncidentDto) SetIncidentMessageNil(b bool)`

 SetIncidentMessageNil sets the value for IncidentMessage to be an explicit nil

### UnsetIncidentMessage
`func (o *IncidentDto) UnsetIncidentMessage()`

UnsetIncidentMessage ensures that no value is present for IncidentMessage, not even an explicit nil
### GetJobDefinitionId

`func (o *IncidentDto) GetJobDefinitionId() string`

GetJobDefinitionId returns the JobDefinitionId field if non-nil, zero value otherwise.

### GetJobDefinitionIdOk

`func (o *IncidentDto) GetJobDefinitionIdOk() (*string, bool)`

GetJobDefinitionIdOk returns a tuple with the JobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionId

`func (o *IncidentDto) SetJobDefinitionId(v string)`

SetJobDefinitionId sets JobDefinitionId field to given value.

### HasJobDefinitionId

`func (o *IncidentDto) HasJobDefinitionId() bool`

HasJobDefinitionId returns a boolean if a field has been set.

### SetJobDefinitionIdNil

`func (o *IncidentDto) SetJobDefinitionIdNil(b bool)`

 SetJobDefinitionIdNil sets the value for JobDefinitionId to be an explicit nil

### UnsetJobDefinitionId
`func (o *IncidentDto) UnsetJobDefinitionId()`

UnsetJobDefinitionId ensures that no value is present for JobDefinitionId, not even an explicit nil
### GetAnnotation

`func (o *IncidentDto) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *IncidentDto) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *IncidentDto) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *IncidentDto) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotationNil

`func (o *IncidentDto) SetAnnotationNil(b bool)`

 SetAnnotationNil sets the value for Annotation to be an explicit nil

### UnsetAnnotation
`func (o *IncidentDto) UnsetAnnotation()`

UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


