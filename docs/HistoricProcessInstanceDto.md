# HistoricProcessInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the process instance. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process. | [optional] 
**SuperProcessInstanceId** | Pointer to **NullableString** | The id of the parent process instance, if it exists. | [optional] 
**SuperCaseInstanceId** | Pointer to **NullableString** | The id of the parent case instance, if it exists. | [optional] 
**CaseInstanceId** | Pointer to **NullableString** | The id of the parent case instance, if it exists. | [optional] 
**ProcessDefinitionName** | Pointer to **NullableString** | The name of the process definition that this process instance belongs to. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition that this process instance belongs to. | [optional] 
**ProcessDefinitionVersion** | Pointer to **NullableInt32** | The version of the process definition that this process instance belongs to. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition that this process instance belongs to. | [optional] 
**BusinessKey** | Pointer to **NullableString** | The business key of the process instance. | [optional] 
**StartTime** | Pointer to **NullableTime** | The time the instance was started. Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**EndTime** | Pointer to **NullableTime** | The time the instance ended. Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the instance should be removed by the History Cleanup job. Default [format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**DurationInMillis** | Pointer to **NullableInt64** | The time the instance took to finish (in milliseconds). | [optional] 
**StartUserId** | Pointer to **NullableString** | The id of the user who started the process instance. | [optional] 
**StartActivityId** | Pointer to **NullableString** | The id of the initial activity that was executed (e.g., a start event). | [optional] 
**DeleteReason** | Pointer to **NullableString** | The provided delete reason in case the process instance was canceled during execution. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the process instance. | [optional] 
**State** | Pointer to **NullableString** | Last state of the process instance, possible values are:  &#x60;ACTIVE&#x60; - running process instance  &#x60;SUSPENDED&#x60; - suspended process instances  &#x60;COMPLETED&#x60; - completed through normal end event  &#x60;EXTERNALLY_TERMINATED&#x60; - terminated externally, for instance through REST API  &#x60;INTERNALLY_TERMINATED&#x60; - terminated internally, for instance by terminating boundary event | [optional] 

## Methods

### NewHistoricProcessInstanceDto

`func NewHistoricProcessInstanceDto() *HistoricProcessInstanceDto`

NewHistoricProcessInstanceDto instantiates a new HistoricProcessInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricProcessInstanceDtoWithDefaults

`func NewHistoricProcessInstanceDtoWithDefaults() *HistoricProcessInstanceDto`

NewHistoricProcessInstanceDtoWithDefaults instantiates a new HistoricProcessInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricProcessInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricProcessInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricProcessInstanceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricProcessInstanceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricProcessInstanceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricProcessInstanceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricProcessInstanceDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricProcessInstanceDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricProcessInstanceDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricProcessInstanceDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricProcessInstanceDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricProcessInstanceDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil
### GetSuperProcessInstanceId

`func (o *HistoricProcessInstanceDto) GetSuperProcessInstanceId() string`

GetSuperProcessInstanceId returns the SuperProcessInstanceId field if non-nil, zero value otherwise.

### GetSuperProcessInstanceIdOk

`func (o *HistoricProcessInstanceDto) GetSuperProcessInstanceIdOk() (*string, bool)`

GetSuperProcessInstanceIdOk returns a tuple with the SuperProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperProcessInstanceId

`func (o *HistoricProcessInstanceDto) SetSuperProcessInstanceId(v string)`

SetSuperProcessInstanceId sets SuperProcessInstanceId field to given value.

### HasSuperProcessInstanceId

`func (o *HistoricProcessInstanceDto) HasSuperProcessInstanceId() bool`

HasSuperProcessInstanceId returns a boolean if a field has been set.

### SetSuperProcessInstanceIdNil

`func (o *HistoricProcessInstanceDto) SetSuperProcessInstanceIdNil(b bool)`

 SetSuperProcessInstanceIdNil sets the value for SuperProcessInstanceId to be an explicit nil

### UnsetSuperProcessInstanceId
`func (o *HistoricProcessInstanceDto) UnsetSuperProcessInstanceId()`

UnsetSuperProcessInstanceId ensures that no value is present for SuperProcessInstanceId, not even an explicit nil
### GetSuperCaseInstanceId

`func (o *HistoricProcessInstanceDto) GetSuperCaseInstanceId() string`

GetSuperCaseInstanceId returns the SuperCaseInstanceId field if non-nil, zero value otherwise.

### GetSuperCaseInstanceIdOk

`func (o *HistoricProcessInstanceDto) GetSuperCaseInstanceIdOk() (*string, bool)`

GetSuperCaseInstanceIdOk returns a tuple with the SuperCaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperCaseInstanceId

`func (o *HistoricProcessInstanceDto) SetSuperCaseInstanceId(v string)`

SetSuperCaseInstanceId sets SuperCaseInstanceId field to given value.

### HasSuperCaseInstanceId

`func (o *HistoricProcessInstanceDto) HasSuperCaseInstanceId() bool`

HasSuperCaseInstanceId returns a boolean if a field has been set.

### SetSuperCaseInstanceIdNil

`func (o *HistoricProcessInstanceDto) SetSuperCaseInstanceIdNil(b bool)`

 SetSuperCaseInstanceIdNil sets the value for SuperCaseInstanceId to be an explicit nil

### UnsetSuperCaseInstanceId
`func (o *HistoricProcessInstanceDto) UnsetSuperCaseInstanceId()`

UnsetSuperCaseInstanceId ensures that no value is present for SuperCaseInstanceId, not even an explicit nil
### GetCaseInstanceId

`func (o *HistoricProcessInstanceDto) GetCaseInstanceId() string`

GetCaseInstanceId returns the CaseInstanceId field if non-nil, zero value otherwise.

### GetCaseInstanceIdOk

`func (o *HistoricProcessInstanceDto) GetCaseInstanceIdOk() (*string, bool)`

GetCaseInstanceIdOk returns a tuple with the CaseInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInstanceId

`func (o *HistoricProcessInstanceDto) SetCaseInstanceId(v string)`

SetCaseInstanceId sets CaseInstanceId field to given value.

### HasCaseInstanceId

`func (o *HistoricProcessInstanceDto) HasCaseInstanceId() bool`

HasCaseInstanceId returns a boolean if a field has been set.

### SetCaseInstanceIdNil

`func (o *HistoricProcessInstanceDto) SetCaseInstanceIdNil(b bool)`

 SetCaseInstanceIdNil sets the value for CaseInstanceId to be an explicit nil

### UnsetCaseInstanceId
`func (o *HistoricProcessInstanceDto) UnsetCaseInstanceId()`

UnsetCaseInstanceId ensures that no value is present for CaseInstanceId, not even an explicit nil
### GetProcessDefinitionName

`func (o *HistoricProcessInstanceDto) GetProcessDefinitionName() string`

GetProcessDefinitionName returns the ProcessDefinitionName field if non-nil, zero value otherwise.

### GetProcessDefinitionNameOk

`func (o *HistoricProcessInstanceDto) GetProcessDefinitionNameOk() (*string, bool)`

GetProcessDefinitionNameOk returns a tuple with the ProcessDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionName

`func (o *HistoricProcessInstanceDto) SetProcessDefinitionName(v string)`

SetProcessDefinitionName sets ProcessDefinitionName field to given value.

### HasProcessDefinitionName

`func (o *HistoricProcessInstanceDto) HasProcessDefinitionName() bool`

HasProcessDefinitionName returns a boolean if a field has been set.

### SetProcessDefinitionNameNil

`func (o *HistoricProcessInstanceDto) SetProcessDefinitionNameNil(b bool)`

 SetProcessDefinitionNameNil sets the value for ProcessDefinitionName to be an explicit nil

### UnsetProcessDefinitionName
`func (o *HistoricProcessInstanceDto) UnsetProcessDefinitionName()`

UnsetProcessDefinitionName ensures that no value is present for ProcessDefinitionName, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricProcessInstanceDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricProcessInstanceDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricProcessInstanceDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricProcessInstanceDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricProcessInstanceDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricProcessInstanceDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionVersion

`func (o *HistoricProcessInstanceDto) GetProcessDefinitionVersion() int32`

GetProcessDefinitionVersion returns the ProcessDefinitionVersion field if non-nil, zero value otherwise.

### GetProcessDefinitionVersionOk

`func (o *HistoricProcessInstanceDto) GetProcessDefinitionVersionOk() (*int32, bool)`

GetProcessDefinitionVersionOk returns a tuple with the ProcessDefinitionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionVersion

`func (o *HistoricProcessInstanceDto) SetProcessDefinitionVersion(v int32)`

SetProcessDefinitionVersion sets ProcessDefinitionVersion field to given value.

### HasProcessDefinitionVersion

`func (o *HistoricProcessInstanceDto) HasProcessDefinitionVersion() bool`

HasProcessDefinitionVersion returns a boolean if a field has been set.

### SetProcessDefinitionVersionNil

`func (o *HistoricProcessInstanceDto) SetProcessDefinitionVersionNil(b bool)`

 SetProcessDefinitionVersionNil sets the value for ProcessDefinitionVersion to be an explicit nil

### UnsetProcessDefinitionVersion
`func (o *HistoricProcessInstanceDto) UnsetProcessDefinitionVersion()`

UnsetProcessDefinitionVersion ensures that no value is present for ProcessDefinitionVersion, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricProcessInstanceDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricProcessInstanceDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricProcessInstanceDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricProcessInstanceDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricProcessInstanceDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricProcessInstanceDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetBusinessKey

`func (o *HistoricProcessInstanceDto) GetBusinessKey() string`

GetBusinessKey returns the BusinessKey field if non-nil, zero value otherwise.

### GetBusinessKeyOk

`func (o *HistoricProcessInstanceDto) GetBusinessKeyOk() (*string, bool)`

GetBusinessKeyOk returns a tuple with the BusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKey

`func (o *HistoricProcessInstanceDto) SetBusinessKey(v string)`

SetBusinessKey sets BusinessKey field to given value.

### HasBusinessKey

`func (o *HistoricProcessInstanceDto) HasBusinessKey() bool`

HasBusinessKey returns a boolean if a field has been set.

### SetBusinessKeyNil

`func (o *HistoricProcessInstanceDto) SetBusinessKeyNil(b bool)`

 SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil

### UnsetBusinessKey
`func (o *HistoricProcessInstanceDto) UnsetBusinessKey()`

UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
### GetStartTime

`func (o *HistoricProcessInstanceDto) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HistoricProcessInstanceDto) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HistoricProcessInstanceDto) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HistoricProcessInstanceDto) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *HistoricProcessInstanceDto) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *HistoricProcessInstanceDto) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *HistoricProcessInstanceDto) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HistoricProcessInstanceDto) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HistoricProcessInstanceDto) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *HistoricProcessInstanceDto) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *HistoricProcessInstanceDto) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *HistoricProcessInstanceDto) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetRemovalTime

`func (o *HistoricProcessInstanceDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricProcessInstanceDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricProcessInstanceDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricProcessInstanceDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricProcessInstanceDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricProcessInstanceDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetDurationInMillis

`func (o *HistoricProcessInstanceDto) GetDurationInMillis() int64`

GetDurationInMillis returns the DurationInMillis field if non-nil, zero value otherwise.

### GetDurationInMillisOk

`func (o *HistoricProcessInstanceDto) GetDurationInMillisOk() (*int64, bool)`

GetDurationInMillisOk returns a tuple with the DurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMillis

`func (o *HistoricProcessInstanceDto) SetDurationInMillis(v int64)`

SetDurationInMillis sets DurationInMillis field to given value.

### HasDurationInMillis

`func (o *HistoricProcessInstanceDto) HasDurationInMillis() bool`

HasDurationInMillis returns a boolean if a field has been set.

### SetDurationInMillisNil

`func (o *HistoricProcessInstanceDto) SetDurationInMillisNil(b bool)`

 SetDurationInMillisNil sets the value for DurationInMillis to be an explicit nil

### UnsetDurationInMillis
`func (o *HistoricProcessInstanceDto) UnsetDurationInMillis()`

UnsetDurationInMillis ensures that no value is present for DurationInMillis, not even an explicit nil
### GetStartUserId

`func (o *HistoricProcessInstanceDto) GetStartUserId() string`

GetStartUserId returns the StartUserId field if non-nil, zero value otherwise.

### GetStartUserIdOk

`func (o *HistoricProcessInstanceDto) GetStartUserIdOk() (*string, bool)`

GetStartUserIdOk returns a tuple with the StartUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUserId

`func (o *HistoricProcessInstanceDto) SetStartUserId(v string)`

SetStartUserId sets StartUserId field to given value.

### HasStartUserId

`func (o *HistoricProcessInstanceDto) HasStartUserId() bool`

HasStartUserId returns a boolean if a field has been set.

### SetStartUserIdNil

`func (o *HistoricProcessInstanceDto) SetStartUserIdNil(b bool)`

 SetStartUserIdNil sets the value for StartUserId to be an explicit nil

### UnsetStartUserId
`func (o *HistoricProcessInstanceDto) UnsetStartUserId()`

UnsetStartUserId ensures that no value is present for StartUserId, not even an explicit nil
### GetStartActivityId

`func (o *HistoricProcessInstanceDto) GetStartActivityId() string`

GetStartActivityId returns the StartActivityId field if non-nil, zero value otherwise.

### GetStartActivityIdOk

`func (o *HistoricProcessInstanceDto) GetStartActivityIdOk() (*string, bool)`

GetStartActivityIdOk returns a tuple with the StartActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActivityId

`func (o *HistoricProcessInstanceDto) SetStartActivityId(v string)`

SetStartActivityId sets StartActivityId field to given value.

### HasStartActivityId

`func (o *HistoricProcessInstanceDto) HasStartActivityId() bool`

HasStartActivityId returns a boolean if a field has been set.

### SetStartActivityIdNil

`func (o *HistoricProcessInstanceDto) SetStartActivityIdNil(b bool)`

 SetStartActivityIdNil sets the value for StartActivityId to be an explicit nil

### UnsetStartActivityId
`func (o *HistoricProcessInstanceDto) UnsetStartActivityId()`

UnsetStartActivityId ensures that no value is present for StartActivityId, not even an explicit nil
### GetDeleteReason

`func (o *HistoricProcessInstanceDto) GetDeleteReason() string`

GetDeleteReason returns the DeleteReason field if non-nil, zero value otherwise.

### GetDeleteReasonOk

`func (o *HistoricProcessInstanceDto) GetDeleteReasonOk() (*string, bool)`

GetDeleteReasonOk returns a tuple with the DeleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteReason

`func (o *HistoricProcessInstanceDto) SetDeleteReason(v string)`

SetDeleteReason sets DeleteReason field to given value.

### HasDeleteReason

`func (o *HistoricProcessInstanceDto) HasDeleteReason() bool`

HasDeleteReason returns a boolean if a field has been set.

### SetDeleteReasonNil

`func (o *HistoricProcessInstanceDto) SetDeleteReasonNil(b bool)`

 SetDeleteReasonNil sets the value for DeleteReason to be an explicit nil

### UnsetDeleteReason
`func (o *HistoricProcessInstanceDto) UnsetDeleteReason()`

UnsetDeleteReason ensures that no value is present for DeleteReason, not even an explicit nil
### GetTenantId

`func (o *HistoricProcessInstanceDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricProcessInstanceDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricProcessInstanceDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricProcessInstanceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricProcessInstanceDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricProcessInstanceDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetState

`func (o *HistoricProcessInstanceDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HistoricProcessInstanceDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HistoricProcessInstanceDto) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HistoricProcessInstanceDto) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *HistoricProcessInstanceDto) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *HistoricProcessInstanceDto) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


