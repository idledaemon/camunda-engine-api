# EventSubscriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the event subscription. | [optional] 
**EventType** | Pointer to **NullableString** | The type of the event subscription. | [optional] 
**EventName** | Pointer to **NullableString** | The name of the event this subscription belongs to as defined in the process model. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The execution that is subscribed on the referenced event. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The process instance this subscription belongs to. | [optional] 
**ActivityId** | Pointer to **NullableString** | The identifier of the activity that this event subscription belongs to. This could for example be the id of a receive task. | [optional] 
**CreatedDate** | Pointer to **NullableTime** | The time this event subscription was created. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant this event subscription belongs to. Can be &#x60;null&#x60; if the subscription belongs to no single tenant. | [optional] 

## Methods

### NewEventSubscriptionDto

`func NewEventSubscriptionDto() *EventSubscriptionDto`

NewEventSubscriptionDto instantiates a new EventSubscriptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSubscriptionDtoWithDefaults

`func NewEventSubscriptionDtoWithDefaults() *EventSubscriptionDto`

NewEventSubscriptionDtoWithDefaults instantiates a new EventSubscriptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventSubscriptionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventSubscriptionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventSubscriptionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventSubscriptionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EventSubscriptionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EventSubscriptionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetEventType

`func (o *EventSubscriptionDto) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventSubscriptionDto) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventSubscriptionDto) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventSubscriptionDto) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *EventSubscriptionDto) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *EventSubscriptionDto) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetEventName

`func (o *EventSubscriptionDto) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *EventSubscriptionDto) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *EventSubscriptionDto) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *EventSubscriptionDto) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### SetEventNameNil

`func (o *EventSubscriptionDto) SetEventNameNil(b bool)`

 SetEventNameNil sets the value for EventName to be an explicit nil

### UnsetEventName
`func (o *EventSubscriptionDto) UnsetEventName()`

UnsetEventName ensures that no value is present for EventName, not even an explicit nil
### GetExecutionId

`func (o *EventSubscriptionDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *EventSubscriptionDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *EventSubscriptionDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *EventSubscriptionDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *EventSubscriptionDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *EventSubscriptionDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetProcessInstanceId

`func (o *EventSubscriptionDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *EventSubscriptionDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *EventSubscriptionDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *EventSubscriptionDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *EventSubscriptionDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *EventSubscriptionDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetActivityId

`func (o *EventSubscriptionDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *EventSubscriptionDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *EventSubscriptionDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *EventSubscriptionDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *EventSubscriptionDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *EventSubscriptionDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetCreatedDate

`func (o *EventSubscriptionDto) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *EventSubscriptionDto) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *EventSubscriptionDto) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *EventSubscriptionDto) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *EventSubscriptionDto) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *EventSubscriptionDto) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetTenantId

`func (o *EventSubscriptionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EventSubscriptionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EventSubscriptionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *EventSubscriptionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *EventSubscriptionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *EventSubscriptionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


