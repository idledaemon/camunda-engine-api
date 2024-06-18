# FetchExternalTaskTopicDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicName** | **string** | **Mandatory.** The topic&#39;s name. | 
**LockDuration** | **NullableInt64** | **Mandatory.** The duration to lock the external tasks for in milliseconds. | 
**Variables** | Pointer to **[]string** | A JSON array of &#x60;String&#x60; values that represent variable names. For each result task belonging to this topic, the given variables are returned as well if they are accessible from the external task&#39;s execution. If not provided - all variables will be fetched. | [optional] 
**LocalVariables** | Pointer to **NullableBool** | If &#x60;true&#x60; only local variables will be fetched. | [optional] [default to false]
**BusinessKey** | Pointer to **NullableString** | A &#x60;String&#x60; value which enables the filtering of tasks based on process instance business key. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter tasks based on process definition id. | [optional] 
**ProcessDefinitionIdIn** | Pointer to **[]string** | Filter tasks based on process definition ids. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Filter tasks based on process definition key. | [optional] 
**ProcessDefinitionKeyIn** | Pointer to **[]string** | Filter tasks based on process definition keys. | [optional] 
**ProcessDefinitionVersionTag** | Pointer to **NullableString** | Filter tasks based on process definition version tag. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Filter tasks without tenant id. | [optional] [default to false]
**TenantIdIn** | Pointer to **[]string** | Filter tasks based on tenant ids. | [optional] 
**ProcessVariables** | Pointer to **map[string]interface{}** | A &#x60;JSON&#x60; object used for filtering tasks based on process instance variable values. A property name of the object represents a process variable name, while the property value represents the process variable value to filter tasks by. | [optional] 
**DeserializeValues** | Pointer to **NullableBool** | Determines whether serializable variable values (typically variables that store custom Java objects) should be deserialized on server side (default &#x60;false&#x60;).  If set to &#x60;true&#x60;, a serializable variable will be deserialized on server side and transformed to JSON using [Jackson&#39;s](https://github.com/FasterXML/jackson) POJO/bean property introspection feature. Note that this requires the Java classes of the variable value to be on the REST API&#39;s classpath.  If set to &#x60;false&#x60;, a serializable variable will be returned in its serialized format. For example, a variable that is serialized as XML will be returned as a JSON string containing XML. | [optional] [default to false]
**IncludeExtensionProperties** | Pointer to **NullableBool** | Determines whether custom extension properties defined in the BPMN activity of the external task (e.g. via the Extensions tab in the Camunda modeler) should be included in the response. Default: false | [optional] [default to false]

## Methods

### NewFetchExternalTaskTopicDto

`func NewFetchExternalTaskTopicDto(topicName string, lockDuration NullableInt64, ) *FetchExternalTaskTopicDto`

NewFetchExternalTaskTopicDto instantiates a new FetchExternalTaskTopicDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchExternalTaskTopicDtoWithDefaults

`func NewFetchExternalTaskTopicDtoWithDefaults() *FetchExternalTaskTopicDto`

NewFetchExternalTaskTopicDtoWithDefaults instantiates a new FetchExternalTaskTopicDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicName

`func (o *FetchExternalTaskTopicDto) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *FetchExternalTaskTopicDto) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *FetchExternalTaskTopicDto) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetLockDuration

`func (o *FetchExternalTaskTopicDto) GetLockDuration() int64`

GetLockDuration returns the LockDuration field if non-nil, zero value otherwise.

### GetLockDurationOk

`func (o *FetchExternalTaskTopicDto) GetLockDurationOk() (*int64, bool)`

GetLockDurationOk returns a tuple with the LockDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockDuration

`func (o *FetchExternalTaskTopicDto) SetLockDuration(v int64)`

SetLockDuration sets LockDuration field to given value.


### SetLockDurationNil

`func (o *FetchExternalTaskTopicDto) SetLockDurationNil(b bool)`

 SetLockDurationNil sets the value for LockDuration to be an explicit nil

### UnsetLockDuration
`func (o *FetchExternalTaskTopicDto) UnsetLockDuration()`

UnsetLockDuration ensures that no value is present for LockDuration, not even an explicit nil
### GetVariables

`func (o *FetchExternalTaskTopicDto) GetVariables() []string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *FetchExternalTaskTopicDto) GetVariablesOk() (*[]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *FetchExternalTaskTopicDto) SetVariables(v []string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *FetchExternalTaskTopicDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *FetchExternalTaskTopicDto) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *FetchExternalTaskTopicDto) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetLocalVariables

`func (o *FetchExternalTaskTopicDto) GetLocalVariables() bool`

GetLocalVariables returns the LocalVariables field if non-nil, zero value otherwise.

### GetLocalVariablesOk

`func (o *FetchExternalTaskTopicDto) GetLocalVariablesOk() (*bool, bool)`

GetLocalVariablesOk returns a tuple with the LocalVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVariables

`func (o *FetchExternalTaskTopicDto) SetLocalVariables(v bool)`

SetLocalVariables sets LocalVariables field to given value.

### HasLocalVariables

`func (o *FetchExternalTaskTopicDto) HasLocalVariables() bool`

HasLocalVariables returns a boolean if a field has been set.

### SetLocalVariablesNil

`func (o *FetchExternalTaskTopicDto) SetLocalVariablesNil(b bool)`

 SetLocalVariablesNil sets the value for LocalVariables to be an explicit nil

### UnsetLocalVariables
`func (o *FetchExternalTaskTopicDto) UnsetLocalVariables()`

UnsetLocalVariables ensures that no value is present for LocalVariables, not even an explicit nil
### GetBusinessKey

`func (o *FetchExternalTaskTopicDto) GetBusinessKey() string`

GetBusinessKey returns the BusinessKey field if non-nil, zero value otherwise.

### GetBusinessKeyOk

`func (o *FetchExternalTaskTopicDto) GetBusinessKeyOk() (*string, bool)`

GetBusinessKeyOk returns a tuple with the BusinessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessKey

`func (o *FetchExternalTaskTopicDto) SetBusinessKey(v string)`

SetBusinessKey sets BusinessKey field to given value.

### HasBusinessKey

`func (o *FetchExternalTaskTopicDto) HasBusinessKey() bool`

HasBusinessKey returns a boolean if a field has been set.

### SetBusinessKeyNil

`func (o *FetchExternalTaskTopicDto) SetBusinessKeyNil(b bool)`

 SetBusinessKeyNil sets the value for BusinessKey to be an explicit nil

### UnsetBusinessKey
`func (o *FetchExternalTaskTopicDto) UnsetBusinessKey()`

UnsetBusinessKey ensures that no value is present for BusinessKey, not even an explicit nil
### GetProcessDefinitionId

`func (o *FetchExternalTaskTopicDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *FetchExternalTaskTopicDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *FetchExternalTaskTopicDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *FetchExternalTaskTopicDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *FetchExternalTaskTopicDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *FetchExternalTaskTopicDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionIdIn

`func (o *FetchExternalTaskTopicDto) GetProcessDefinitionIdIn() []string`

GetProcessDefinitionIdIn returns the ProcessDefinitionIdIn field if non-nil, zero value otherwise.

### GetProcessDefinitionIdInOk

`func (o *FetchExternalTaskTopicDto) GetProcessDefinitionIdInOk() (*[]string, bool)`

GetProcessDefinitionIdInOk returns a tuple with the ProcessDefinitionIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionIdIn

`func (o *FetchExternalTaskTopicDto) SetProcessDefinitionIdIn(v []string)`

SetProcessDefinitionIdIn sets ProcessDefinitionIdIn field to given value.

### HasProcessDefinitionIdIn

`func (o *FetchExternalTaskTopicDto) HasProcessDefinitionIdIn() bool`

HasProcessDefinitionIdIn returns a boolean if a field has been set.

### SetProcessDefinitionIdInNil

`func (o *FetchExternalTaskTopicDto) SetProcessDefinitionIdInNil(b bool)`

 SetProcessDefinitionIdInNil sets the value for ProcessDefinitionIdIn to be an explicit nil

### UnsetProcessDefinitionIdIn
`func (o *FetchExternalTaskTopicDto) UnsetProcessDefinitionIdIn()`

UnsetProcessDefinitionIdIn ensures that no value is present for ProcessDefinitionIdIn, not even an explicit nil
### GetProcessDefinitionKey

`func (o *FetchExternalTaskTopicDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *FetchExternalTaskTopicDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *FetchExternalTaskTopicDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *FetchExternalTaskTopicDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *FetchExternalTaskTopicDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *FetchExternalTaskTopicDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionKeyIn

`func (o *FetchExternalTaskTopicDto) GetProcessDefinitionKeyIn() []string`

GetProcessDefinitionKeyIn returns the ProcessDefinitionKeyIn field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyInOk

`func (o *FetchExternalTaskTopicDto) GetProcessDefinitionKeyInOk() (*[]string, bool)`

GetProcessDefinitionKeyInOk returns a tuple with the ProcessDefinitionKeyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKeyIn

`func (o *FetchExternalTaskTopicDto) SetProcessDefinitionKeyIn(v []string)`

SetProcessDefinitionKeyIn sets ProcessDefinitionKeyIn field to given value.

### HasProcessDefinitionKeyIn

`func (o *FetchExternalTaskTopicDto) HasProcessDefinitionKeyIn() bool`

HasProcessDefinitionKeyIn returns a boolean if a field has been set.

### SetProcessDefinitionKeyInNil

`func (o *FetchExternalTaskTopicDto) SetProcessDefinitionKeyInNil(b bool)`

 SetProcessDefinitionKeyInNil sets the value for ProcessDefinitionKeyIn to be an explicit nil

### UnsetProcessDefinitionKeyIn
`func (o *FetchExternalTaskTopicDto) UnsetProcessDefinitionKeyIn()`

UnsetProcessDefinitionKeyIn ensures that no value is present for ProcessDefinitionKeyIn, not even an explicit nil
### GetProcessDefinitionVersionTag

`func (o *FetchExternalTaskTopicDto) GetProcessDefinitionVersionTag() string`

GetProcessDefinitionVersionTag returns the ProcessDefinitionVersionTag field if non-nil, zero value otherwise.

### GetProcessDefinitionVersionTagOk

`func (o *FetchExternalTaskTopicDto) GetProcessDefinitionVersionTagOk() (*string, bool)`

GetProcessDefinitionVersionTagOk returns a tuple with the ProcessDefinitionVersionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionVersionTag

`func (o *FetchExternalTaskTopicDto) SetProcessDefinitionVersionTag(v string)`

SetProcessDefinitionVersionTag sets ProcessDefinitionVersionTag field to given value.

### HasProcessDefinitionVersionTag

`func (o *FetchExternalTaskTopicDto) HasProcessDefinitionVersionTag() bool`

HasProcessDefinitionVersionTag returns a boolean if a field has been set.

### SetProcessDefinitionVersionTagNil

`func (o *FetchExternalTaskTopicDto) SetProcessDefinitionVersionTagNil(b bool)`

 SetProcessDefinitionVersionTagNil sets the value for ProcessDefinitionVersionTag to be an explicit nil

### UnsetProcessDefinitionVersionTag
`func (o *FetchExternalTaskTopicDto) UnsetProcessDefinitionVersionTag()`

UnsetProcessDefinitionVersionTag ensures that no value is present for ProcessDefinitionVersionTag, not even an explicit nil
### GetWithoutTenantId

`func (o *FetchExternalTaskTopicDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *FetchExternalTaskTopicDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *FetchExternalTaskTopicDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *FetchExternalTaskTopicDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *FetchExternalTaskTopicDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *FetchExternalTaskTopicDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetTenantIdIn

`func (o *FetchExternalTaskTopicDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *FetchExternalTaskTopicDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *FetchExternalTaskTopicDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *FetchExternalTaskTopicDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *FetchExternalTaskTopicDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *FetchExternalTaskTopicDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetProcessVariables

`func (o *FetchExternalTaskTopicDto) GetProcessVariables() map[string]interface{}`

GetProcessVariables returns the ProcessVariables field if non-nil, zero value otherwise.

### GetProcessVariablesOk

`func (o *FetchExternalTaskTopicDto) GetProcessVariablesOk() (*map[string]interface{}, bool)`

GetProcessVariablesOk returns a tuple with the ProcessVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessVariables

`func (o *FetchExternalTaskTopicDto) SetProcessVariables(v map[string]interface{})`

SetProcessVariables sets ProcessVariables field to given value.

### HasProcessVariables

`func (o *FetchExternalTaskTopicDto) HasProcessVariables() bool`

HasProcessVariables returns a boolean if a field has been set.

### GetDeserializeValues

`func (o *FetchExternalTaskTopicDto) GetDeserializeValues() bool`

GetDeserializeValues returns the DeserializeValues field if non-nil, zero value otherwise.

### GetDeserializeValuesOk

`func (o *FetchExternalTaskTopicDto) GetDeserializeValuesOk() (*bool, bool)`

GetDeserializeValuesOk returns a tuple with the DeserializeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeserializeValues

`func (o *FetchExternalTaskTopicDto) SetDeserializeValues(v bool)`

SetDeserializeValues sets DeserializeValues field to given value.

### HasDeserializeValues

`func (o *FetchExternalTaskTopicDto) HasDeserializeValues() bool`

HasDeserializeValues returns a boolean if a field has been set.

### SetDeserializeValuesNil

`func (o *FetchExternalTaskTopicDto) SetDeserializeValuesNil(b bool)`

 SetDeserializeValuesNil sets the value for DeserializeValues to be an explicit nil

### UnsetDeserializeValues
`func (o *FetchExternalTaskTopicDto) UnsetDeserializeValues()`

UnsetDeserializeValues ensures that no value is present for DeserializeValues, not even an explicit nil
### GetIncludeExtensionProperties

`func (o *FetchExternalTaskTopicDto) GetIncludeExtensionProperties() bool`

GetIncludeExtensionProperties returns the IncludeExtensionProperties field if non-nil, zero value otherwise.

### GetIncludeExtensionPropertiesOk

`func (o *FetchExternalTaskTopicDto) GetIncludeExtensionPropertiesOk() (*bool, bool)`

GetIncludeExtensionPropertiesOk returns a tuple with the IncludeExtensionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExtensionProperties

`func (o *FetchExternalTaskTopicDto) SetIncludeExtensionProperties(v bool)`

SetIncludeExtensionProperties sets IncludeExtensionProperties field to given value.

### HasIncludeExtensionProperties

`func (o *FetchExternalTaskTopicDto) HasIncludeExtensionProperties() bool`

HasIncludeExtensionProperties returns a boolean if a field has been set.

### SetIncludeExtensionPropertiesNil

`func (o *FetchExternalTaskTopicDto) SetIncludeExtensionPropertiesNil(b bool)`

 SetIncludeExtensionPropertiesNil sets the value for IncludeExtensionProperties to be an explicit nil

### UnsetIncludeExtensionProperties
`func (o *FetchExternalTaskTopicDto) UnsetIncludeExtensionProperties()`

UnsetIncludeExtensionProperties ensures that no value is present for IncludeExtensionProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


