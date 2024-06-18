# HistoricDecisionInstanceStatisticsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecisionDefinitionKey** | Pointer to **NullableString** | A key of decision definition. | [optional] 
**Evaluations** | Pointer to **NullableInt32** | A number of evaluation for decision definition. | [optional] 

## Methods

### NewHistoricDecisionInstanceStatisticsDto

`func NewHistoricDecisionInstanceStatisticsDto() *HistoricDecisionInstanceStatisticsDto`

NewHistoricDecisionInstanceStatisticsDto instantiates a new HistoricDecisionInstanceStatisticsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricDecisionInstanceStatisticsDtoWithDefaults

`func NewHistoricDecisionInstanceStatisticsDtoWithDefaults() *HistoricDecisionInstanceStatisticsDto`

NewHistoricDecisionInstanceStatisticsDtoWithDefaults instantiates a new HistoricDecisionInstanceStatisticsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecisionDefinitionKey

`func (o *HistoricDecisionInstanceStatisticsDto) GetDecisionDefinitionKey() string`

GetDecisionDefinitionKey returns the DecisionDefinitionKey field if non-nil, zero value otherwise.

### GetDecisionDefinitionKeyOk

`func (o *HistoricDecisionInstanceStatisticsDto) GetDecisionDefinitionKeyOk() (*string, bool)`

GetDecisionDefinitionKeyOk returns a tuple with the DecisionDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDefinitionKey

`func (o *HistoricDecisionInstanceStatisticsDto) SetDecisionDefinitionKey(v string)`

SetDecisionDefinitionKey sets DecisionDefinitionKey field to given value.

### HasDecisionDefinitionKey

`func (o *HistoricDecisionInstanceStatisticsDto) HasDecisionDefinitionKey() bool`

HasDecisionDefinitionKey returns a boolean if a field has been set.

### SetDecisionDefinitionKeyNil

`func (o *HistoricDecisionInstanceStatisticsDto) SetDecisionDefinitionKeyNil(b bool)`

 SetDecisionDefinitionKeyNil sets the value for DecisionDefinitionKey to be an explicit nil

### UnsetDecisionDefinitionKey
`func (o *HistoricDecisionInstanceStatisticsDto) UnsetDecisionDefinitionKey()`

UnsetDecisionDefinitionKey ensures that no value is present for DecisionDefinitionKey, not even an explicit nil
### GetEvaluations

`func (o *HistoricDecisionInstanceStatisticsDto) GetEvaluations() int32`

GetEvaluations returns the Evaluations field if non-nil, zero value otherwise.

### GetEvaluationsOk

`func (o *HistoricDecisionInstanceStatisticsDto) GetEvaluationsOk() (*int32, bool)`

GetEvaluationsOk returns a tuple with the Evaluations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluations

`func (o *HistoricDecisionInstanceStatisticsDto) SetEvaluations(v int32)`

SetEvaluations sets Evaluations field to given value.

### HasEvaluations

`func (o *HistoricDecisionInstanceStatisticsDto) HasEvaluations() bool`

HasEvaluations returns a boolean if a field has been set.

### SetEvaluationsNil

`func (o *HistoricDecisionInstanceStatisticsDto) SetEvaluationsNil(b bool)`

 SetEvaluationsNil sets the value for Evaluations to be an explicit nil

### UnsetEvaluations
`func (o *HistoricDecisionInstanceStatisticsDto) UnsetEvaluations()`

UnsetEvaluations ensures that no value is present for Evaluations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


