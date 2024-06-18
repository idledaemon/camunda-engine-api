# HistoricIdentityLinkLogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the Historic identity link entry. | [optional] 
**Time** | Pointer to **NullableTime** | The time when the identity link is logged.  [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**Type** | Pointer to **NullableString** | The type of identity link (candidate/assignee/owner). | [optional] 
**UserId** | Pointer to **NullableString** | The id of the user/assignee. | [optional] 
**GroupId** | Pointer to **NullableString** | The id of the group. | [optional] 
**TaskId** | Pointer to **NullableString** | The id of the task. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition. | [optional] 
**OperationType** | Pointer to **NullableString** | Type of operation (add/delete). | [optional] 
**AssignerId** | Pointer to **NullableString** | The id of the assigner. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the identity link should be removed by the History Cleanup job.  [Default format](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process containing this identity link. | [optional] 

## Methods

### NewHistoricIdentityLinkLogDto

`func NewHistoricIdentityLinkLogDto() *HistoricIdentityLinkLogDto`

NewHistoricIdentityLinkLogDto instantiates a new HistoricIdentityLinkLogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricIdentityLinkLogDtoWithDefaults

`func NewHistoricIdentityLinkLogDtoWithDefaults() *HistoricIdentityLinkLogDto`

NewHistoricIdentityLinkLogDtoWithDefaults instantiates a new HistoricIdentityLinkLogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricIdentityLinkLogDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricIdentityLinkLogDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricIdentityLinkLogDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricIdentityLinkLogDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricIdentityLinkLogDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricIdentityLinkLogDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTime

`func (o *HistoricIdentityLinkLogDto) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *HistoricIdentityLinkLogDto) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *HistoricIdentityLinkLogDto) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *HistoricIdentityLinkLogDto) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *HistoricIdentityLinkLogDto) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *HistoricIdentityLinkLogDto) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil
### GetType

`func (o *HistoricIdentityLinkLogDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoricIdentityLinkLogDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoricIdentityLinkLogDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HistoricIdentityLinkLogDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HistoricIdentityLinkLogDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HistoricIdentityLinkLogDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUserId

`func (o *HistoricIdentityLinkLogDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HistoricIdentityLinkLogDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HistoricIdentityLinkLogDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *HistoricIdentityLinkLogDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *HistoricIdentityLinkLogDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *HistoricIdentityLinkLogDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetGroupId

`func (o *HistoricIdentityLinkLogDto) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *HistoricIdentityLinkLogDto) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *HistoricIdentityLinkLogDto) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *HistoricIdentityLinkLogDto) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *HistoricIdentityLinkLogDto) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *HistoricIdentityLinkLogDto) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetTaskId

`func (o *HistoricIdentityLinkLogDto) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *HistoricIdentityLinkLogDto) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *HistoricIdentityLinkLogDto) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *HistoricIdentityLinkLogDto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *HistoricIdentityLinkLogDto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *HistoricIdentityLinkLogDto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricIdentityLinkLogDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricIdentityLinkLogDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricIdentityLinkLogDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricIdentityLinkLogDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricIdentityLinkLogDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricIdentityLinkLogDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricIdentityLinkLogDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricIdentityLinkLogDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricIdentityLinkLogDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricIdentityLinkLogDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricIdentityLinkLogDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricIdentityLinkLogDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetOperationType

`func (o *HistoricIdentityLinkLogDto) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *HistoricIdentityLinkLogDto) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *HistoricIdentityLinkLogDto) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *HistoricIdentityLinkLogDto) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationTypeNil

`func (o *HistoricIdentityLinkLogDto) SetOperationTypeNil(b bool)`

 SetOperationTypeNil sets the value for OperationType to be an explicit nil

### UnsetOperationType
`func (o *HistoricIdentityLinkLogDto) UnsetOperationType()`

UnsetOperationType ensures that no value is present for OperationType, not even an explicit nil
### GetAssignerId

`func (o *HistoricIdentityLinkLogDto) GetAssignerId() string`

GetAssignerId returns the AssignerId field if non-nil, zero value otherwise.

### GetAssignerIdOk

`func (o *HistoricIdentityLinkLogDto) GetAssignerIdOk() (*string, bool)`

GetAssignerIdOk returns a tuple with the AssignerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignerId

`func (o *HistoricIdentityLinkLogDto) SetAssignerId(v string)`

SetAssignerId sets AssignerId field to given value.

### HasAssignerId

`func (o *HistoricIdentityLinkLogDto) HasAssignerId() bool`

HasAssignerId returns a boolean if a field has been set.

### SetAssignerIdNil

`func (o *HistoricIdentityLinkLogDto) SetAssignerIdNil(b bool)`

 SetAssignerIdNil sets the value for AssignerId to be an explicit nil

### UnsetAssignerId
`func (o *HistoricIdentityLinkLogDto) UnsetAssignerId()`

UnsetAssignerId ensures that no value is present for AssignerId, not even an explicit nil
### GetTenantId

`func (o *HistoricIdentityLinkLogDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricIdentityLinkLogDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricIdentityLinkLogDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricIdentityLinkLogDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricIdentityLinkLogDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricIdentityLinkLogDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetRemovalTime

`func (o *HistoricIdentityLinkLogDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricIdentityLinkLogDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricIdentityLinkLogDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricIdentityLinkLogDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricIdentityLinkLogDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricIdentityLinkLogDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricIdentityLinkLogDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricIdentityLinkLogDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricIdentityLinkLogDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricIdentityLinkLogDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricIdentityLinkLogDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricIdentityLinkLogDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


