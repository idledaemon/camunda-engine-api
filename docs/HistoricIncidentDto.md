# HistoricIncidentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the incident. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition this incident is associated with. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition this incident is associated with. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The key of the process definition this incident is associated with. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The id of the execution this incident is associated with. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process containing this incident. | [optional] 
**CreateTime** | Pointer to **NullableTime** | The time this incident happened.  [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**EndTime** | Pointer to **NullableTime** | The time this incident has been deleted or resolved.  [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the incident should be removed by the History Cleanup job. [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**IncidentType** | Pointer to **NullableString** | The type of incident, for example: &#x60;failedJobs&#x60; will be returned in case of an incident which identified a failed job during the execution of a process instance. See the [User Guide](/manual/develop/user- guide/process-engine/incidents/#incident-types) for a list of incident types. | [optional] 
**ActivityId** | Pointer to **NullableString** | The id of the activity this incident is associated with. | [optional] 
**FailedActivityId** | Pointer to **NullableString** | The id of the activity on which the last exception occurred. | [optional] 
**CauseIncidentId** | Pointer to **NullableString** | The id of the associated cause incident which has been triggered. | [optional] 
**RootCauseIncidentId** | Pointer to **NullableString** | The id of the associated root cause incident which has been triggered. | [optional] 
**Configuration** | Pointer to **NullableString** | The payload of this incident. | [optional] 
**HistoryConfiguration** | Pointer to **NullableString** | The payload of this incident at the time when it occurred. | [optional] 
**IncidentMessage** | Pointer to **NullableString** | The message of this incident. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant this incident is associated with. | [optional] 
**JobDefinitionId** | Pointer to **NullableString** | The job definition id the incident is associated with. | [optional] 
**Open** | Pointer to **NullableBool** | If true, this incident is open. | [optional] 
**Deleted** | Pointer to **NullableBool** | If true, this incident has been deleted. | [optional] 
**Resolved** | Pointer to **NullableBool** | If true, this incident has been resolved. | [optional] 
**Annotation** | Pointer to **NullableString** | The annotation set to the incident. | [optional] 

## Methods

### NewHistoricIncidentDto

`func NewHistoricIncidentDto() *HistoricIncidentDto`

NewHistoricIncidentDto instantiates a new HistoricIncidentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricIncidentDtoWithDefaults

`func NewHistoricIncidentDtoWithDefaults() *HistoricIncidentDto`

NewHistoricIncidentDtoWithDefaults instantiates a new HistoricIncidentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricIncidentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricIncidentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricIncidentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricIncidentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricIncidentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricIncidentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricIncidentDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricIncidentDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricIncidentDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricIncidentDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricIncidentDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricIncidentDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricIncidentDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricIncidentDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricIncidentDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricIncidentDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricIncidentDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricIncidentDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricIncidentDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricIncidentDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricIncidentDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricIncidentDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricIncidentDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricIncidentDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetExecutionId

`func (o *HistoricIncidentDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HistoricIncidentDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HistoricIncidentDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HistoricIncidentDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HistoricIncidentDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HistoricIncidentDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricIncidentDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricIncidentDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricIncidentDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricIncidentDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricIncidentDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricIncidentDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil
### GetCreateTime

`func (o *HistoricIncidentDto) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HistoricIncidentDto) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HistoricIncidentDto) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HistoricIncidentDto) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### SetCreateTimeNil

`func (o *HistoricIncidentDto) SetCreateTimeNil(b bool)`

 SetCreateTimeNil sets the value for CreateTime to be an explicit nil

### UnsetCreateTime
`func (o *HistoricIncidentDto) UnsetCreateTime()`

UnsetCreateTime ensures that no value is present for CreateTime, not even an explicit nil
### GetEndTime

`func (o *HistoricIncidentDto) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HistoricIncidentDto) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HistoricIncidentDto) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *HistoricIncidentDto) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *HistoricIncidentDto) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *HistoricIncidentDto) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetRemovalTime

`func (o *HistoricIncidentDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricIncidentDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricIncidentDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricIncidentDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricIncidentDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricIncidentDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetIncidentType

`func (o *HistoricIncidentDto) GetIncidentType() string`

GetIncidentType returns the IncidentType field if non-nil, zero value otherwise.

### GetIncidentTypeOk

`func (o *HistoricIncidentDto) GetIncidentTypeOk() (*string, bool)`

GetIncidentTypeOk returns a tuple with the IncidentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentType

`func (o *HistoricIncidentDto) SetIncidentType(v string)`

SetIncidentType sets IncidentType field to given value.

### HasIncidentType

`func (o *HistoricIncidentDto) HasIncidentType() bool`

HasIncidentType returns a boolean if a field has been set.

### SetIncidentTypeNil

`func (o *HistoricIncidentDto) SetIncidentTypeNil(b bool)`

 SetIncidentTypeNil sets the value for IncidentType to be an explicit nil

### UnsetIncidentType
`func (o *HistoricIncidentDto) UnsetIncidentType()`

UnsetIncidentType ensures that no value is present for IncidentType, not even an explicit nil
### GetActivityId

`func (o *HistoricIncidentDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *HistoricIncidentDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *HistoricIncidentDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *HistoricIncidentDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *HistoricIncidentDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *HistoricIncidentDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetFailedActivityId

`func (o *HistoricIncidentDto) GetFailedActivityId() string`

GetFailedActivityId returns the FailedActivityId field if non-nil, zero value otherwise.

### GetFailedActivityIdOk

`func (o *HistoricIncidentDto) GetFailedActivityIdOk() (*string, bool)`

GetFailedActivityIdOk returns a tuple with the FailedActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedActivityId

`func (o *HistoricIncidentDto) SetFailedActivityId(v string)`

SetFailedActivityId sets FailedActivityId field to given value.

### HasFailedActivityId

`func (o *HistoricIncidentDto) HasFailedActivityId() bool`

HasFailedActivityId returns a boolean if a field has been set.

### SetFailedActivityIdNil

`func (o *HistoricIncidentDto) SetFailedActivityIdNil(b bool)`

 SetFailedActivityIdNil sets the value for FailedActivityId to be an explicit nil

### UnsetFailedActivityId
`func (o *HistoricIncidentDto) UnsetFailedActivityId()`

UnsetFailedActivityId ensures that no value is present for FailedActivityId, not even an explicit nil
### GetCauseIncidentId

`func (o *HistoricIncidentDto) GetCauseIncidentId() string`

GetCauseIncidentId returns the CauseIncidentId field if non-nil, zero value otherwise.

### GetCauseIncidentIdOk

`func (o *HistoricIncidentDto) GetCauseIncidentIdOk() (*string, bool)`

GetCauseIncidentIdOk returns a tuple with the CauseIncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseIncidentId

`func (o *HistoricIncidentDto) SetCauseIncidentId(v string)`

SetCauseIncidentId sets CauseIncidentId field to given value.

### HasCauseIncidentId

`func (o *HistoricIncidentDto) HasCauseIncidentId() bool`

HasCauseIncidentId returns a boolean if a field has been set.

### SetCauseIncidentIdNil

`func (o *HistoricIncidentDto) SetCauseIncidentIdNil(b bool)`

 SetCauseIncidentIdNil sets the value for CauseIncidentId to be an explicit nil

### UnsetCauseIncidentId
`func (o *HistoricIncidentDto) UnsetCauseIncidentId()`

UnsetCauseIncidentId ensures that no value is present for CauseIncidentId, not even an explicit nil
### GetRootCauseIncidentId

`func (o *HistoricIncidentDto) GetRootCauseIncidentId() string`

GetRootCauseIncidentId returns the RootCauseIncidentId field if non-nil, zero value otherwise.

### GetRootCauseIncidentIdOk

`func (o *HistoricIncidentDto) GetRootCauseIncidentIdOk() (*string, bool)`

GetRootCauseIncidentIdOk returns a tuple with the RootCauseIncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIncidentId

`func (o *HistoricIncidentDto) SetRootCauseIncidentId(v string)`

SetRootCauseIncidentId sets RootCauseIncidentId field to given value.

### HasRootCauseIncidentId

`func (o *HistoricIncidentDto) HasRootCauseIncidentId() bool`

HasRootCauseIncidentId returns a boolean if a field has been set.

### SetRootCauseIncidentIdNil

`func (o *HistoricIncidentDto) SetRootCauseIncidentIdNil(b bool)`

 SetRootCauseIncidentIdNil sets the value for RootCauseIncidentId to be an explicit nil

### UnsetRootCauseIncidentId
`func (o *HistoricIncidentDto) UnsetRootCauseIncidentId()`

UnsetRootCauseIncidentId ensures that no value is present for RootCauseIncidentId, not even an explicit nil
### GetConfiguration

`func (o *HistoricIncidentDto) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *HistoricIncidentDto) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *HistoricIncidentDto) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *HistoricIncidentDto) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *HistoricIncidentDto) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *HistoricIncidentDto) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetHistoryConfiguration

`func (o *HistoricIncidentDto) GetHistoryConfiguration() string`

GetHistoryConfiguration returns the HistoryConfiguration field if non-nil, zero value otherwise.

### GetHistoryConfigurationOk

`func (o *HistoricIncidentDto) GetHistoryConfigurationOk() (*string, bool)`

GetHistoryConfigurationOk returns a tuple with the HistoryConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryConfiguration

`func (o *HistoricIncidentDto) SetHistoryConfiguration(v string)`

SetHistoryConfiguration sets HistoryConfiguration field to given value.

### HasHistoryConfiguration

`func (o *HistoricIncidentDto) HasHistoryConfiguration() bool`

HasHistoryConfiguration returns a boolean if a field has been set.

### SetHistoryConfigurationNil

`func (o *HistoricIncidentDto) SetHistoryConfigurationNil(b bool)`

 SetHistoryConfigurationNil sets the value for HistoryConfiguration to be an explicit nil

### UnsetHistoryConfiguration
`func (o *HistoricIncidentDto) UnsetHistoryConfiguration()`

UnsetHistoryConfiguration ensures that no value is present for HistoryConfiguration, not even an explicit nil
### GetIncidentMessage

`func (o *HistoricIncidentDto) GetIncidentMessage() string`

GetIncidentMessage returns the IncidentMessage field if non-nil, zero value otherwise.

### GetIncidentMessageOk

`func (o *HistoricIncidentDto) GetIncidentMessageOk() (*string, bool)`

GetIncidentMessageOk returns a tuple with the IncidentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentMessage

`func (o *HistoricIncidentDto) SetIncidentMessage(v string)`

SetIncidentMessage sets IncidentMessage field to given value.

### HasIncidentMessage

`func (o *HistoricIncidentDto) HasIncidentMessage() bool`

HasIncidentMessage returns a boolean if a field has been set.

### SetIncidentMessageNil

`func (o *HistoricIncidentDto) SetIncidentMessageNil(b bool)`

 SetIncidentMessageNil sets the value for IncidentMessage to be an explicit nil

### UnsetIncidentMessage
`func (o *HistoricIncidentDto) UnsetIncidentMessage()`

UnsetIncidentMessage ensures that no value is present for IncidentMessage, not even an explicit nil
### GetTenantId

`func (o *HistoricIncidentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricIncidentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricIncidentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricIncidentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricIncidentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricIncidentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetJobDefinitionId

`func (o *HistoricIncidentDto) GetJobDefinitionId() string`

GetJobDefinitionId returns the JobDefinitionId field if non-nil, zero value otherwise.

### GetJobDefinitionIdOk

`func (o *HistoricIncidentDto) GetJobDefinitionIdOk() (*string, bool)`

GetJobDefinitionIdOk returns a tuple with the JobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionId

`func (o *HistoricIncidentDto) SetJobDefinitionId(v string)`

SetJobDefinitionId sets JobDefinitionId field to given value.

### HasJobDefinitionId

`func (o *HistoricIncidentDto) HasJobDefinitionId() bool`

HasJobDefinitionId returns a boolean if a field has been set.

### SetJobDefinitionIdNil

`func (o *HistoricIncidentDto) SetJobDefinitionIdNil(b bool)`

 SetJobDefinitionIdNil sets the value for JobDefinitionId to be an explicit nil

### UnsetJobDefinitionId
`func (o *HistoricIncidentDto) UnsetJobDefinitionId()`

UnsetJobDefinitionId ensures that no value is present for JobDefinitionId, not even an explicit nil
### GetOpen

`func (o *HistoricIncidentDto) GetOpen() bool`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *HistoricIncidentDto) GetOpenOk() (*bool, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *HistoricIncidentDto) SetOpen(v bool)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *HistoricIncidentDto) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### SetOpenNil

`func (o *HistoricIncidentDto) SetOpenNil(b bool)`

 SetOpenNil sets the value for Open to be an explicit nil

### UnsetOpen
`func (o *HistoricIncidentDto) UnsetOpen()`

UnsetOpen ensures that no value is present for Open, not even an explicit nil
### GetDeleted

`func (o *HistoricIncidentDto) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *HistoricIncidentDto) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *HistoricIncidentDto) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *HistoricIncidentDto) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *HistoricIncidentDto) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *HistoricIncidentDto) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetResolved

`func (o *HistoricIncidentDto) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *HistoricIncidentDto) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *HistoricIncidentDto) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *HistoricIncidentDto) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *HistoricIncidentDto) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *HistoricIncidentDto) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetAnnotation

`func (o *HistoricIncidentDto) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *HistoricIncidentDto) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *HistoricIncidentDto) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *HistoricIncidentDto) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotationNil

`func (o *HistoricIncidentDto) SetAnnotationNil(b bool)`

 SetAnnotationNil sets the value for Annotation to be an explicit nil

### UnsetAnnotation
`func (o *HistoricIncidentDto) UnsetAnnotation()`

UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


