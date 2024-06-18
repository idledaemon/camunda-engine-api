# ExecutionQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessKey** | Pointer to **NullableString** | Filter by the business key of the process instances the executions belong to. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter by the process definition the executions run on. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Filter by the key of the process definition the executions run on. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | Filter by the id of the process instance the execution belongs to. | [optional] 
**ActivityId** | Pointer to **NullableString** | Filter by the id of the activity the execution currently executes. | [optional] 
**SignalEventSubscriptionName** | Pointer to **NullableString** | Select only those executions that expect a signal of the given name. | [optional] 
**MessageEventSubscriptionName** | Pointer to **NullableString** | Select only those executions that expect a message of the given name. | [optional] 
**Active** | Pointer to **NullableBool** | Only include active executions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Suspended** | Pointer to **NullableBool** | Only include suspended executions. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**IncidentId** | Pointer to **NullableString** | Filter by the incident id. | [optional] 
**IncidentType** | Pointer to **NullableString** | Filter by the incident type. See the [User Guide](/manual/develop/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | [optional] 
**IncidentMessage** | Pointer to **NullableString** | Filter by the incident message. Exact match. | [optional] 
**IncidentMessageLike** | Pointer to **NullableString** | Filter by the incident message that the parameter is a substring of. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Filter by a  list of tenant ids. An execution must have one of the given tenant ids. | [optional] 
**Variables** | Pointer to [**[]VariableQueryParameterDto**](VariableQueryParameterDto.md) | An array to only include executions that have variables with certain values.  The array consists of objects with the three properties &#x60;name&#x60;, &#x60;operator&#x60; and &#x60;value&#x60;. &#x60;name (String)&#x60; is the variable name, &#x60;operator (String)&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. &#x60;value&#x60; may be &#x60;String&#x60;, &#x60;Number&#x60; or &#x60;Boolean&#x60;.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to; &#x60;gt&#x60; - greater than; &#x60;gteq&#x60; - greater than or equal to; &#x60;lt&#x60; - lower than; &#x60;lteq&#x60; - lower than or equal to; &#x60;like&#x60;. | [optional] 
**ProcessVariables** | Pointer to [**[]VariableQueryParameterDto**](VariableQueryParameterDto.md) | An array to only include executions that belong to a process instance with variables with certain values.  The array consists of objects with the three properties &#x60;name&#x60;, &#x60;operator&#x60; and &#x60;value&#x60;. &#x60;name (String)&#x60; is the variable name, &#x60;operator (String)&#x60; is the comparison operator to be used and &#x60;value&#x60; the variable value. &#x60;value&#x60; may be &#x60;String&#x60;, &#x60;Number&#x60; or &#x60;Boolean&#x60;.  Valid operator values are: &#x60;eq&#x60; - equal to; &#x60;neq&#x60; - not equal to. | [optional] 
**VariableNamesIgnoreCase** | Pointer to **NullableBool** | Match all variable names provided in &#x60;variables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableName** and **variablename** are treated as equal. | [optional] 
**VariableValuesIgnoreCase** | Pointer to **NullableBool** | Match all variable values provided in &#x60;variables&#x60; and &#x60;processVariables&#x60; case- insensitively. If set to &#x60;true&#x60; **variableValue** and **variablevalue** are treated as equal. | [optional] 
**Sorting** | Pointer to [**[]ExecutionQueryDtoSortingInner**](ExecutionQueryDtoSortingInner.md) | An array of criteria to sort the result by. Each element of the array is                        an object that specifies one ordering. The position in the array                        identifies the rank of an ordering, i.e., whether it is primary, secondary,                        etc. Has no effect for the &#x60;/count&#x60; endpoint | [optional] 

## Methods

### NewExecutionQueryDto

`func NewExecutionQueryDto() *ExecutionQueryDto`

NewExecutionQueryDto instantiates a new ExecutionQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionQueryDtoWithDefaults

`func NewExecutionQueryDtoWithDefaults() *ExecutionQueryDto`

NewExecutionQueryDtoWithDefaults instantiates a new ExecutionQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessKey

`func (o *ExecutionQueryDto) GetBusinessKey() string`

GetBusinessKey returns the BusinessKey field if non-nil, zero value otherwise.

### GetBusinessKeyOk

`func (o *ExecutionQueryDto) GetBusinessKeyOk() (*string, bool)`

GetBusinessKeyOk returns a tuple with the BusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKey

`func (o *ExecutionQueryDto) SetBusinessKey(v string)`

SetBusinessKey sets BusinessKey field to given value.

### HasBusinessKey

`func (o *ExecutionQueryDto) HasBusinessKey() bool`

HasBusinessKey returns a boolean if a field has been set.

### SetBusinessKeyNil

`func (o *ExecutionQueryDto) SetBusinessKeyNil(b bool)`

 SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil

### UnsetBusinessKey
`func (o *ExecutionQueryDto) UnsetBusinessKey()`

UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
### GetProcessDefinitionId

`func (o *ExecutionQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *ExecutionQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *ExecutionQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *ExecutionQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *ExecutionQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *ExecutionQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *ExecutionQueryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *ExecutionQueryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *ExecutionQueryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *ExecutionQueryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *ExecutionQueryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *ExecutionQueryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessInstanceId

`func (o *ExecutionQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *ExecutionQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *ExecutionQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *ExecutionQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *ExecutionQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *ExecutionQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetActivityId

`func (o *ExecutionQueryDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *ExecutionQueryDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *ExecutionQueryDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *ExecutionQueryDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *ExecutionQueryDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *ExecutionQueryDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetSignalEventSubscriptionName

`func (o *ExecutionQueryDto) GetSignalEventSubscriptionName() string`

GetSignalEventSubscriptionName returns the SignalEventSubscriptionName field if non-nil, zero value otherwise.

### GetSignalEventSubscriptionNameOk

`func (o *ExecutionQueryDto) GetSignalEventSubscriptionNameOk() (*string, bool)`

GetSignalEventSubscriptionNameOk returns a tuple with the SignalEventSubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalEventSubscriptionName

`func (o *ExecutionQueryDto) SetSignalEventSubscriptionName(v string)`

SetSignalEventSubscriptionName sets SignalEventSubscriptionName field to given value.

### HasSignalEventSubscriptionName

`func (o *ExecutionQueryDto) HasSignalEventSubscriptionName() bool`

HasSignalEventSubscriptionName returns a boolean if a field has been set.

### SetSignalEventSubscriptionNameNil

`func (o *ExecutionQueryDto) SetSignalEventSubscriptionNameNil(b bool)`

 SetSignalEventSubscriptionNameNil sets the value for SignalEventSubscriptionName to be an explicit nil

### UnsetSignalEventSubscriptionName
`func (o *ExecutionQueryDto) UnsetSignalEventSubscriptionName()`

UnsetSignalEventSubscriptionName ensures that no value is present for SignalEventSubscriptionName, not even an explicit nil
### GetMessageEventSubscriptionName

`func (o *ExecutionQueryDto) GetMessageEventSubscriptionName() string`

GetMessageEventSubscriptionName returns the MessageEventSubscriptionName field if non-nil, zero value otherwise.

### GetMessageEventSubscriptionNameOk

`func (o *ExecutionQueryDto) GetMessageEventSubscriptionNameOk() (*string, bool)`

GetMessageEventSubscriptionNameOk returns a tuple with the MessageEventSubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEventSubscriptionName

`func (o *ExecutionQueryDto) SetMessageEventSubscriptionName(v string)`

SetMessageEventSubscriptionName sets MessageEventSubscriptionName field to given value.

### HasMessageEventSubscriptionName

`func (o *ExecutionQueryDto) HasMessageEventSubscriptionName() bool`

HasMessageEventSubscriptionName returns a boolean if a field has been set.

### SetMessageEventSubscriptionNameNil

`func (o *ExecutionQueryDto) SetMessageEventSubscriptionNameNil(b bool)`

 SetMessageEventSubscriptionNameNil sets the value for MessageEventSubscriptionName to be an explicit nil

### UnsetMessageEventSubscriptionName
`func (o *ExecutionQueryDto) UnsetMessageEventSubscriptionName()`

UnsetMessageEventSubscriptionName ensures that no value is present for MessageEventSubscriptionName, not even an explicit nil
### GetActive

`func (o *ExecutionQueryDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ExecutionQueryDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ExecutionQueryDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ExecutionQueryDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *ExecutionQueryDto) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *ExecutionQueryDto) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetSuspended

`func (o *ExecutionQueryDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ExecutionQueryDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ExecutionQueryDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ExecutionQueryDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *ExecutionQueryDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *ExecutionQueryDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetIncidentId

`func (o *ExecutionQueryDto) GetIncidentId() string`

GetIncidentId returns the IncidentId field if non-nil, zero value otherwise.

### GetIncidentIdOk

`func (o *ExecutionQueryDto) GetIncidentIdOk() (*string, bool)`

GetIncidentIdOk returns a tuple with the IncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentId

`func (o *ExecutionQueryDto) SetIncidentId(v string)`

SetIncidentId sets IncidentId field to given value.

### HasIncidentId

`func (o *ExecutionQueryDto) HasIncidentId() bool`

HasIncidentId returns a boolean if a field has been set.

### SetIncidentIdNil

`func (o *ExecutionQueryDto) SetIncidentIdNil(b bool)`

 SetIncidentIdNil sets the value for IncidentId to be an explicit nil

### UnsetIncidentId
`func (o *ExecutionQueryDto) UnsetIncidentId()`

UnsetIncidentId ensures that no value is present for IncidentId, not even an explicit nil
### GetIncidentType

`func (o *ExecutionQueryDto) GetIncidentType() string`

GetIncidentType returns the IncidentType field if non-nil, zero value otherwise.

### GetIncidentTypeOk

`func (o *ExecutionQueryDto) GetIncidentTypeOk() (*string, bool)`

GetIncidentTypeOk returns a tuple with the IncidentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentType

`func (o *ExecutionQueryDto) SetIncidentType(v string)`

SetIncidentType sets IncidentType field to given value.

### HasIncidentType

`func (o *ExecutionQueryDto) HasIncidentType() bool`

HasIncidentType returns a boolean if a field has been set.

### SetIncidentTypeNil

`func (o *ExecutionQueryDto) SetIncidentTypeNil(b bool)`

 SetIncidentTypeNil sets the value for IncidentType to be an explicit nil

### UnsetIncidentType
`func (o *ExecutionQueryDto) UnsetIncidentType()`

UnsetIncidentType ensures that no value is present for IncidentType, not even an explicit nil
### GetIncidentMessage

`func (o *ExecutionQueryDto) GetIncidentMessage() string`

GetIncidentMessage returns the IncidentMessage field if non-nil, zero value otherwise.

### GetIncidentMessageOk

`func (o *ExecutionQueryDto) GetIncidentMessageOk() (*string, bool)`

GetIncidentMessageOk returns a tuple with the IncidentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentMessage

`func (o *ExecutionQueryDto) SetIncidentMessage(v string)`

SetIncidentMessage sets IncidentMessage field to given value.

### HasIncidentMessage

`func (o *ExecutionQueryDto) HasIncidentMessage() bool`

HasIncidentMessage returns a boolean if a field has been set.

### SetIncidentMessageNil

`func (o *ExecutionQueryDto) SetIncidentMessageNil(b bool)`

 SetIncidentMessageNil sets the value for IncidentMessage to be an explicit nil

### UnsetIncidentMessage
`func (o *ExecutionQueryDto) UnsetIncidentMessage()`

UnsetIncidentMessage ensures that no value is present for IncidentMessage, not even an explicit nil
### GetIncidentMessageLike

`func (o *ExecutionQueryDto) GetIncidentMessageLike() string`

GetIncidentMessageLike returns the IncidentMessageLike field if non-nil, zero value otherwise.

### GetIncidentMessageLikeOk

`func (o *ExecutionQueryDto) GetIncidentMessageLikeOk() (*string, bool)`

GetIncidentMessageLikeOk returns a tuple with the IncidentMessageLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentMessageLike

`func (o *ExecutionQueryDto) SetIncidentMessageLike(v string)`

SetIncidentMessageLike sets IncidentMessageLike field to given value.

### HasIncidentMessageLike

`func (o *ExecutionQueryDto) HasIncidentMessageLike() bool`

HasIncidentMessageLike returns a boolean if a field has been set.

### SetIncidentMessageLikeNil

`func (o *ExecutionQueryDto) SetIncidentMessageLikeNil(b bool)`

 SetIncidentMessageLikeNil sets the value for IncidentMessageLike to be an explicit nil

### UnsetIncidentMessageLike
`func (o *ExecutionQueryDto) UnsetIncidentMessageLike()`

UnsetIncidentMessageLike ensures that no value is present for IncidentMessageLike, not even an explicit nil
### GetTenantIdIn

`func (o *ExecutionQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *ExecutionQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *ExecutionQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *ExecutionQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *ExecutionQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *ExecutionQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetVariables

`func (o *ExecutionQueryDto) GetVariables() []VariableQueryParameterDto`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ExecutionQueryDto) GetVariablesOk() (*[]VariableQueryParameterDto, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ExecutionQueryDto) SetVariables(v []VariableQueryParameterDto)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ExecutionQueryDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *ExecutionQueryDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *ExecutionQueryDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetProcessVariables

`func (o *ExecutionQueryDto) GetProcessVariables() []VariableQueryParameterDto`

GetProcessVariables returns the ProcessVariables field if non-nil, zero value otherwise.

### GetProcessVariablesOk

`func (o *ExecutionQueryDto) GetProcessVariablesOk() (*[]VariableQueryParameterDto, bool)`

GetProcessVariablesOk returns a tuple with the ProcessVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessVariables

`func (o *ExecutionQueryDto) SetProcessVariables(v []VariableQueryParameterDto)`

SetProcessVariables sets ProcessVariables field to given value.

### HasProcessVariables

`func (o *ExecutionQueryDto) HasProcessVariables() bool`

HasProcessVariables returns a boolean if a field has been set.

### SetProcessVariablesNil

`func (o *ExecutionQueryDto) SetProcessVariablesNil(b bool)`

 SetProcessVariablesNil sets the value for ProcessVariables to be an explicit nil

### UnsetProcessVariables
`func (o *ExecutionQueryDto) UnsetProcessVariables()`

UnsetProcessVariables ensures that no value is present for ProcessVariables, not even an explicit nil
### GetVariableNamesIgnoreCase

`func (o *ExecutionQueryDto) GetVariableNamesIgnoreCase() bool`

GetVariableNamesIgnoreCase returns the VariableNamesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableNamesIgnoreCaseOk

`func (o *ExecutionQueryDto) GetVariableNamesIgnoreCaseOk() (*bool, bool)`

GetVariableNamesIgnoreCaseOk returns a tuple with the VariableNamesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableNamesIgnoreCase

`func (o *ExecutionQueryDto) SetVariableNamesIgnoreCase(v bool)`

SetVariableNamesIgnoreCase sets VariableNamesIgnoreCase field to given value.

### HasVariableNamesIgnoreCase

`func (o *ExecutionQueryDto) HasVariableNamesIgnoreCase() bool`

HasVariableNamesIgnoreCase returns a boolean if a field has been set.

### SetVariableNamesIgnoreCaseNil

`func (o *ExecutionQueryDto) SetVariableNamesIgnoreCaseNil(b bool)`

 SetVariableNamesIgnoreCaseNil sets the value for VariableNamesIgnoreCase to be an explicit nil

### UnsetVariableNamesIgnoreCase
`func (o *ExecutionQueryDto) UnsetVariableNamesIgnoreCase()`

UnsetVariableNamesIgnoreCase ensures that no value is present for VariableNamesIgnoreCase, not even an explicit nil
### GetVariableValuesIgnoreCase

`func (o *ExecutionQueryDto) GetVariableValuesIgnoreCase() bool`

GetVariableValuesIgnoreCase returns the VariableValuesIgnoreCase field if non-nil, zero value otherwise.

### GetVariableValuesIgnoreCaseOk

`func (o *ExecutionQueryDto) GetVariableValuesIgnoreCaseOk() (*bool, bool)`

GetVariableValuesIgnoreCaseOk returns a tuple with the VariableValuesIgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValuesIgnoreCase

`func (o *ExecutionQueryDto) SetVariableValuesIgnoreCase(v bool)`

SetVariableValuesIgnoreCase sets VariableValuesIgnoreCase field to given value.

### HasVariableValuesIgnoreCase

`func (o *ExecutionQueryDto) HasVariableValuesIgnoreCase() bool`

HasVariableValuesIgnoreCase returns a boolean if a field has been set.

### SetVariableValuesIgnoreCaseNil

`func (o *ExecutionQueryDto) SetVariableValuesIgnoreCaseNil(b bool)`

 SetVariableValuesIgnoreCaseNil sets the value for VariableValuesIgnoreCase to be an explicit nil

### UnsetVariableValuesIgnoreCase
`func (o *ExecutionQueryDto) UnsetVariableValuesIgnoreCase()`

UnsetVariableValuesIgnoreCase ensures that no value is present for VariableValuesIgnoreCase, not even an explicit nil
### GetSorting

`func (o *ExecutionQueryDto) GetSorting() []ExecutionQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *ExecutionQueryDto) GetSortingOk() (*[]ExecutionQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *ExecutionQueryDto) SetSorting(v []ExecutionQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *ExecutionQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *ExecutionQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *ExecutionQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


