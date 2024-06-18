# EventSubscriptionQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSubscriptionId** | Pointer to **NullableString** | The id of the event subscription. | [optional] 
**EventName** | Pointer to **NullableString** | The name of the event this subscription belongs to as defined in the process model. | [optional] 
**EventType** | Pointer to **NullableString** | The type of the event subscription. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The execution that is subscribed on the referenced event. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The process instance this subscription belongs to. | [optional] 
**ActivityId** | Pointer to **NullableString** | The identifier of the activity that this event subscription belongs to. This could for example be the id of a receive task. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Filter by a comma-separated list of tenant ids. Only select subscriptions that belong to one of the given tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only select subscriptions which have no tenant id. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**IncludeEventSubscriptionsWithoutTenantId** | Pointer to **NullableBool** | Select event subscriptions which have no tenant id. Can be used in combination with tenantIdIn parameter. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Sorting** | Pointer to [**[]EventSubscriptionQueryDtoSortingInner**](EventSubscriptionQueryDtoSortingInner.md) | Apply sorting of the result | [optional] 

## Methods

### NewEventSubscriptionQueryDto

`func NewEventSubscriptionQueryDto() *EventSubscriptionQueryDto`

NewEventSubscriptionQueryDto instantiates a new EventSubscriptionQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSubscriptionQueryDtoWithDefaults

`func NewEventSubscriptionQueryDtoWithDefaults() *EventSubscriptionQueryDto`

NewEventSubscriptionQueryDtoWithDefaults instantiates a new EventSubscriptionQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSubscriptionId

`func (o *EventSubscriptionQueryDto) GetEventSubscriptionId() string`

GetEventSubscriptionId returns the EventSubscriptionId field if non-nil, zero value otherwise.

### GetEventSubscriptionIdOk

`func (o *EventSubscriptionQueryDto) GetEventSubscriptionIdOk() (*string, bool)`

GetEventSubscriptionIdOk returns a tuple with the EventSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubscriptionId

`func (o *EventSubscriptionQueryDto) SetEventSubscriptionId(v string)`

SetEventSubscriptionId sets EventSubscriptionId field to given value.

### HasEventSubscriptionId

`func (o *EventSubscriptionQueryDto) HasEventSubscriptionId() bool`

HasEventSubscriptionId returns a boolean if a field has been set.

### SetEventSubscriptionIdNil

`func (o *EventSubscriptionQueryDto) SetEventSubscriptionIdNil(b bool)`

 SetEventSubscriptionIdNil sets the value for EventSubscriptionId to be an explicit nil

### UnsetEventSubscriptionId
`func (o *EventSubscriptionQueryDto) UnsetEventSubscriptionId()`

UnsetEventSubscriptionId ensures that no value is present for EventSubscriptionId, not even an explicit nil
### GetEventName

`func (o *EventSubscriptionQueryDto) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *EventSubscriptionQueryDto) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *EventSubscriptionQueryDto) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *EventSubscriptionQueryDto) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### SetEventNameNil

`func (o *EventSubscriptionQueryDto) SetEventNameNil(b bool)`

 SetEventNameNil sets the value for EventName to be an explicit nil

### UnsetEventName
`func (o *EventSubscriptionQueryDto) UnsetEventName()`

UnsetEventName ensures that no value is present for EventName, not even an explicit nil
### GetEventType

`func (o *EventSubscriptionQueryDto) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventSubscriptionQueryDto) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventSubscriptionQueryDto) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventSubscriptionQueryDto) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *EventSubscriptionQueryDto) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *EventSubscriptionQueryDto) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetExecutionId

`func (o *EventSubscriptionQueryDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *EventSubscriptionQueryDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *EventSubscriptionQueryDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *EventSubscriptionQueryDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *EventSubscriptionQueryDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *EventSubscriptionQueryDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetProcessInstanceId

`func (o *EventSubscriptionQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *EventSubscriptionQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *EventSubscriptionQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *EventSubscriptionQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *EventSubscriptionQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *EventSubscriptionQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetActivityId

`func (o *EventSubscriptionQueryDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *EventSubscriptionQueryDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *EventSubscriptionQueryDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *EventSubscriptionQueryDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *EventSubscriptionQueryDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *EventSubscriptionQueryDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetTenantIdIn

`func (o *EventSubscriptionQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *EventSubscriptionQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *EventSubscriptionQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *EventSubscriptionQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *EventSubscriptionQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *EventSubscriptionQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *EventSubscriptionQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *EventSubscriptionQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *EventSubscriptionQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *EventSubscriptionQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *EventSubscriptionQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *EventSubscriptionQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetIncludeEventSubscriptionsWithoutTenantId

`func (o *EventSubscriptionQueryDto) GetIncludeEventSubscriptionsWithoutTenantId() bool`

GetIncludeEventSubscriptionsWithoutTenantId returns the IncludeEventSubscriptionsWithoutTenantId field if non-nil, zero value otherwise.

### GetIncludeEventSubscriptionsWithoutTenantIdOk

`func (o *EventSubscriptionQueryDto) GetIncludeEventSubscriptionsWithoutTenantIdOk() (*bool, bool)`

GetIncludeEventSubscriptionsWithoutTenantIdOk returns a tuple with the IncludeEventSubscriptionsWithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEventSubscriptionsWithoutTenantId

`func (o *EventSubscriptionQueryDto) SetIncludeEventSubscriptionsWithoutTenantId(v bool)`

SetIncludeEventSubscriptionsWithoutTenantId sets IncludeEventSubscriptionsWithoutTenantId field to given value.

### HasIncludeEventSubscriptionsWithoutTenantId

`func (o *EventSubscriptionQueryDto) HasIncludeEventSubscriptionsWithoutTenantId() bool`

HasIncludeEventSubscriptionsWithoutTenantId returns a boolean if a field has been set.

### SetIncludeEventSubscriptionsWithoutTenantIdNil

`func (o *EventSubscriptionQueryDto) SetIncludeEventSubscriptionsWithoutTenantIdNil(b bool)`

 SetIncludeEventSubscriptionsWithoutTenantIdNil sets the value for IncludeEventSubscriptionsWithoutTenantId to be an explicit nil

### UnsetIncludeEventSubscriptionsWithoutTenantId
`func (o *EventSubscriptionQueryDto) UnsetIncludeEventSubscriptionsWithoutTenantId()`

UnsetIncludeEventSubscriptionsWithoutTenantId ensures that no value is present for IncludeEventSubscriptionsWithoutTenantId, not even an explicit nil
### GetSorting

`func (o *EventSubscriptionQueryDto) GetSorting() []EventSubscriptionQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *EventSubscriptionQueryDto) GetSortingOk() (*[]EventSubscriptionQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *EventSubscriptionQueryDto) SetSorting(v []EventSubscriptionQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *EventSubscriptionQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *EventSubscriptionQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *EventSubscriptionQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


