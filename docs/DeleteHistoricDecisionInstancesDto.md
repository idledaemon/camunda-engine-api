# DeleteHistoricDecisionInstancesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoricDecisionInstanceIds** | Pointer to **[]string** | A list of historic decision instance ids to delete. | [optional] 
**HistoricDecisionInstanceQuery** | Pointer to [**HistoricDecisionInstanceQueryDto**](HistoricDecisionInstanceQueryDto.md) |  | [optional] 
**DeleteReason** | Pointer to **NullableString** | A string with delete reason. | [optional] 

## Methods

### NewDeleteHistoricDecisionInstancesDto

`func NewDeleteHistoricDecisionInstancesDto() *DeleteHistoricDecisionInstancesDto`

NewDeleteHistoricDecisionInstancesDto instantiates a new DeleteHistoricDecisionInstancesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteHistoricDecisionInstancesDtoWithDefaults

`func NewDeleteHistoricDecisionInstancesDtoWithDefaults() *DeleteHistoricDecisionInstancesDto`

NewDeleteHistoricDecisionInstancesDtoWithDefaults instantiates a new DeleteHistoricDecisionInstancesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoricDecisionInstanceIds

`func (o *DeleteHistoricDecisionInstancesDto) GetHistoricDecisionInstanceIds() []string`

GetHistoricDecisionInstanceIds returns the HistoricDecisionInstanceIds field if non-nil, zero value otherwise.

### GetHistoricDecisionInstanceIdsOk

`func (o *DeleteHistoricDecisionInstancesDto) GetHistoricDecisionInstanceIdsOk() (*[]string, bool)`

GetHistoricDecisionInstanceIdsOk returns a tuple with the HistoricDecisionInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricDecisionInstanceIds

`func (o *DeleteHistoricDecisionInstancesDto) SetHistoricDecisionInstanceIds(v []string)`

SetHistoricDecisionInstanceIds sets HistoricDecisionInstanceIds field to given value.

### HasHistoricDecisionInstanceIds

`func (o *DeleteHistoricDecisionInstancesDto) HasHistoricDecisionInstanceIds() bool`

HasHistoricDecisionInstanceIds returns a boolean if a field has been set.

### SetHistoricDecisionInstanceIdsNil

`func (o *DeleteHistoricDecisionInstancesDto) SetHistoricDecisionInstanceIdsNil(b bool)`

 SetHistoricDecisionInstanceIdsNil sets the value for HistoricDecisionInstanceIds to be an explicit nil

### UnsetHistoricDecisionInstanceIds
`func (o *DeleteHistoricDecisionInstancesDto) UnsetHistoricDecisionInstanceIds()`

UnsetHistoricDecisionInstanceIds ensures that no value is present for HistoricDecisionInstanceIds, not even an explicit nil
### GetHistoricDecisionInstanceQuery

`func (o *DeleteHistoricDecisionInstancesDto) GetHistoricDecisionInstanceQuery() HistoricDecisionInstanceQueryDto`

GetHistoricDecisionInstanceQuery returns the HistoricDecisionInstanceQuery field if non-nil, zero value otherwise.

### GetHistoricDecisionInstanceQueryOk

`func (o *DeleteHistoricDecisionInstancesDto) GetHistoricDecisionInstanceQueryOk() (*HistoricDecisionInstanceQueryDto, bool)`

GetHistoricDecisionInstanceQueryOk returns a tuple with the HistoricDecisionInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricDecisionInstanceQuery

`func (o *DeleteHistoricDecisionInstancesDto) SetHistoricDecisionInstanceQuery(v HistoricDecisionInstanceQueryDto)`

SetHistoricDecisionInstanceQuery sets HistoricDecisionInstanceQuery field to given value.

### HasHistoricDecisionInstanceQuery

`func (o *DeleteHistoricDecisionInstancesDto) HasHistoricDecisionInstanceQuery() bool`

HasHistoricDecisionInstanceQuery returns a boolean if a field has been set.

### GetDeleteReason

`func (o *DeleteHistoricDecisionInstancesDto) GetDeleteReason() string`

GetDeleteReason returns the DeleteReason field if non-nil, zero value otherwise.

### GetDeleteReasonOk

`func (o *DeleteHistoricDecisionInstancesDto) GetDeleteReasonOk() (*string, bool)`

GetDeleteReasonOk returns a tuple with the DeleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteReason

`func (o *DeleteHistoricDecisionInstancesDto) SetDeleteReason(v string)`

SetDeleteReason sets DeleteReason field to given value.

### HasDeleteReason

`func (o *DeleteHistoricDecisionInstancesDto) HasDeleteReason() bool`

HasDeleteReason returns a boolean if a field has been set.

### SetDeleteReasonNil

`func (o *DeleteHistoricDecisionInstancesDto) SetDeleteReasonNil(b bool)`

 SetDeleteReasonNil sets the value for DeleteReason to be an explicit nil

### UnsetDeleteReason
`func (o *DeleteHistoricDecisionInstancesDto) UnsetDeleteReason()`

UnsetDeleteReason ensures that no value is present for DeleteReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


