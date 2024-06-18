# DeleteHistoricProcessInstancesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoricProcessInstanceIds** | Pointer to **[]string** | A list historic process instance ids to delete. | [optional] 
**HistoricProcessInstanceQuery** | Pointer to [**HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) |  | [optional] 
**DeleteReason** | Pointer to **NullableString** | A string with delete reason. | [optional] 
**FailIfNotExists** | Pointer to **NullableBool** | If set to &#x60;false&#x60;, the request will still be successful if one ore more of the process ids are not found. | [optional] 

## Methods

### NewDeleteHistoricProcessInstancesDto

`func NewDeleteHistoricProcessInstancesDto() *DeleteHistoricProcessInstancesDto`

NewDeleteHistoricProcessInstancesDto instantiates a new DeleteHistoricProcessInstancesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteHistoricProcessInstancesDtoWithDefaults

`func NewDeleteHistoricProcessInstancesDtoWithDefaults() *DeleteHistoricProcessInstancesDto`

NewDeleteHistoricProcessInstancesDtoWithDefaults instantiates a new DeleteHistoricProcessInstancesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoricProcessInstanceIds

`func (o *DeleteHistoricProcessInstancesDto) GetHistoricProcessInstanceIds() []string`

GetHistoricProcessInstanceIds returns the HistoricProcessInstanceIds field if non-nil, zero value otherwise.

### GetHistoricProcessInstanceIdsOk

`func (o *DeleteHistoricProcessInstancesDto) GetHistoricProcessInstanceIdsOk() (*[]string, bool)`

GetHistoricProcessInstanceIdsOk returns a tuple with the HistoricProcessInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricProcessInstanceIds

`func (o *DeleteHistoricProcessInstancesDto) SetHistoricProcessInstanceIds(v []string)`

SetHistoricProcessInstanceIds sets HistoricProcessInstanceIds field to given value.

### HasHistoricProcessInstanceIds

`func (o *DeleteHistoricProcessInstancesDto) HasHistoricProcessInstanceIds() bool`

HasHistoricProcessInstanceIds returns a boolean if a field has been set.

### SetHistoricProcessInstanceIdsNil

`func (o *DeleteHistoricProcessInstancesDto) SetHistoricProcessInstanceIdsNil(b bool)`

 SetHistoricProcessInstanceIdsNil sets the value for HistoricProcessInstanceIds to be an explicit nil

### UnsetHistoricProcessInstanceIds
`func (o *DeleteHistoricProcessInstancesDto) UnsetHistoricProcessInstanceIds()`

UnsetHistoricProcessInstanceIds ensures that no value is present for HistoricProcessInstanceIds, not even an explicit nil
### GetHistoricProcessInstanceQuery

`func (o *DeleteHistoricProcessInstancesDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto`

GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field if non-nil, zero value otherwise.

### GetHistoricProcessInstanceQueryOk

`func (o *DeleteHistoricProcessInstancesDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool)`

GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricProcessInstanceQuery

`func (o *DeleteHistoricProcessInstancesDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto)`

SetHistoricProcessInstanceQuery sets HistoricProcessInstanceQuery field to given value.

### HasHistoricProcessInstanceQuery

`func (o *DeleteHistoricProcessInstancesDto) HasHistoricProcessInstanceQuery() bool`

HasHistoricProcessInstanceQuery returns a boolean if a field has been set.

### GetDeleteReason

`func (o *DeleteHistoricProcessInstancesDto) GetDeleteReason() string`

GetDeleteReason returns the DeleteReason field if non-nil, zero value otherwise.

### GetDeleteReasonOk

`func (o *DeleteHistoricProcessInstancesDto) GetDeleteReasonOk() (*string, bool)`

GetDeleteReasonOk returns a tuple with the DeleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteReason

`func (o *DeleteHistoricProcessInstancesDto) SetDeleteReason(v string)`

SetDeleteReason sets DeleteReason field to given value.

### HasDeleteReason

`func (o *DeleteHistoricProcessInstancesDto) HasDeleteReason() bool`

HasDeleteReason returns a boolean if a field has been set.

### SetDeleteReasonNil

`func (o *DeleteHistoricProcessInstancesDto) SetDeleteReasonNil(b bool)`

 SetDeleteReasonNil sets the value for DeleteReason to be an explicit nil

### UnsetDeleteReason
`func (o *DeleteHistoricProcessInstancesDto) UnsetDeleteReason()`

UnsetDeleteReason ensures that no value is present for DeleteReason, not even an explicit nil
### GetFailIfNotExists

`func (o *DeleteHistoricProcessInstancesDto) GetFailIfNotExists() bool`

GetFailIfNotExists returns the FailIfNotExists field if non-nil, zero value otherwise.

### GetFailIfNotExistsOk

`func (o *DeleteHistoricProcessInstancesDto) GetFailIfNotExistsOk() (*bool, bool)`

GetFailIfNotExistsOk returns a tuple with the FailIfNotExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailIfNotExists

`func (o *DeleteHistoricProcessInstancesDto) SetFailIfNotExists(v bool)`

SetFailIfNotExists sets FailIfNotExists field to given value.

### HasFailIfNotExists

`func (o *DeleteHistoricProcessInstancesDto) HasFailIfNotExists() bool`

HasFailIfNotExists returns a boolean if a field has been set.

### SetFailIfNotExistsNil

`func (o *DeleteHistoricProcessInstancesDto) SetFailIfNotExistsNil(b bool)`

 SetFailIfNotExistsNil sets the value for FailIfNotExists to be an explicit nil

### UnsetFailIfNotExists
`func (o *DeleteHistoricProcessInstancesDto) UnsetFailIfNotExists()`

UnsetFailIfNotExists ensures that no value is present for FailIfNotExists, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


