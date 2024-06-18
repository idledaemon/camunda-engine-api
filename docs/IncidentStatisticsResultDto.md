# IncidentStatisticsResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncidentType** | Pointer to **NullableString** | The type of the incident the number of incidents is aggregated for. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/incidents/#incident-types) for a list of incident types. | [optional] 
**IncidentCount** | Pointer to **NullableInt32** | The total number of incidents for the corresponding incident type. | [optional] 

## Methods

### NewIncidentStatisticsResultDto

`func NewIncidentStatisticsResultDto() *IncidentStatisticsResultDto`

NewIncidentStatisticsResultDto instantiates a new IncidentStatisticsResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentStatisticsResultDtoWithDefaults

`func NewIncidentStatisticsResultDtoWithDefaults() *IncidentStatisticsResultDto`

NewIncidentStatisticsResultDtoWithDefaults instantiates a new IncidentStatisticsResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidentType

`func (o *IncidentStatisticsResultDto) GetIncidentType() string`

GetIncidentType returns the IncidentType field if non-nil, zero value otherwise.

### GetIncidentTypeOk

`func (o *IncidentStatisticsResultDto) GetIncidentTypeOk() (*string, bool)`

GetIncidentTypeOk returns a tuple with the IncidentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentType

`func (o *IncidentStatisticsResultDto) SetIncidentType(v string)`

SetIncidentType sets IncidentType field to given value.

### HasIncidentType

`func (o *IncidentStatisticsResultDto) HasIncidentType() bool`

HasIncidentType returns a boolean if a field has been set.

### SetIncidentTypeNil

`func (o *IncidentStatisticsResultDto) SetIncidentTypeNil(b bool)`

 SetIncidentTypeNil sets the value for IncidentType to be an explicit nil

### UnsetIncidentType
`func (o *IncidentStatisticsResultDto) UnsetIncidentType()`

UnsetIncidentType ensures that no value is present for IncidentType, not even an explicit nil
### GetIncidentCount

`func (o *IncidentStatisticsResultDto) GetIncidentCount() int32`

GetIncidentCount returns the IncidentCount field if non-nil, zero value otherwise.

### GetIncidentCountOk

`func (o *IncidentStatisticsResultDto) GetIncidentCountOk() (*int32, bool)`

GetIncidentCountOk returns a tuple with the IncidentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentCount

`func (o *IncidentStatisticsResultDto) SetIncidentCount(v int32)`

SetIncidentCount sets IncidentCount field to given value.

### HasIncidentCount

`func (o *IncidentStatisticsResultDto) HasIncidentCount() bool`

HasIncidentCount returns a boolean if a field has been set.

### SetIncidentCountNil

`func (o *IncidentStatisticsResultDto) SetIncidentCountNil(b bool)`

 SetIncidentCountNil sets the value for IncidentCount to be an explicit nil

### UnsetIncidentCount
`func (o *IncidentStatisticsResultDto) UnsetIncidentCount()`

UnsetIncidentCount ensures that no value is present for IncidentCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


