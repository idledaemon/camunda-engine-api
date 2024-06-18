# ActivityStatisticsResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the activity the results are aggregated for. | [optional] 
**Instances** | Pointer to **int32** | The total number of running process instances of this activity. | [optional] 
**FailedJobs** | Pointer to **int32** | The total number of failed jobs for the running instances. **Note**: Will be &#x60;0&#x60; (not &#x60;null&#x60;), if failed jobs were excluded. | [optional] 
**Incidents** | Pointer to [**[]IncidentStatisticsResultDto**](IncidentStatisticsResultDto.md) | Each item in the resulting array is an object which contains &#x60;incidentType&#x60; and &#x60;incidentCount&#x60;. **Note**: Will be an empty array, if &#x60;incidents&#x60; or &#x60;incidentsForType&#x60; were excluded. Furthermore, the array will be also empty if no incidents were found. | [optional] 
**Class** | Pointer to **NullableString** | The fully qualified class name of the data transfer object class. The class name might change in future releases. | [optional] 

## Methods

### NewActivityStatisticsResultDto

`func NewActivityStatisticsResultDto() *ActivityStatisticsResultDto`

NewActivityStatisticsResultDto instantiates a new ActivityStatisticsResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityStatisticsResultDtoWithDefaults

`func NewActivityStatisticsResultDtoWithDefaults() *ActivityStatisticsResultDto`

NewActivityStatisticsResultDtoWithDefaults instantiates a new ActivityStatisticsResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityStatisticsResultDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityStatisticsResultDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityStatisticsResultDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityStatisticsResultDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ActivityStatisticsResultDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ActivityStatisticsResultDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInstances

`func (o *ActivityStatisticsResultDto) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ActivityStatisticsResultDto) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ActivityStatisticsResultDto) SetInstances(v int32)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ActivityStatisticsResultDto) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetFailedJobs

`func (o *ActivityStatisticsResultDto) GetFailedJobs() int32`

GetFailedJobs returns the FailedJobs field if non-nil, zero value otherwise.

### GetFailedJobsOk

`func (o *ActivityStatisticsResultDto) GetFailedJobsOk() (*int32, bool)`

GetFailedJobsOk returns a tuple with the FailedJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedJobs

`func (o *ActivityStatisticsResultDto) SetFailedJobs(v int32)`

SetFailedJobs sets FailedJobs field to given value.

### HasFailedJobs

`func (o *ActivityStatisticsResultDto) HasFailedJobs() bool`

HasFailedJobs returns a boolean if a field has been set.

### GetIncidents

`func (o *ActivityStatisticsResultDto) GetIncidents() []IncidentStatisticsResultDto`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *ActivityStatisticsResultDto) GetIncidentsOk() (*[]IncidentStatisticsResultDto, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *ActivityStatisticsResultDto) SetIncidents(v []IncidentStatisticsResultDto)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *ActivityStatisticsResultDto) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### SetIncidentsNil

`func (o *ActivityStatisticsResultDto) SetIncidentsNil(b bool)`

 SetIncidentsNil sets the value for Incidents to be an explicit nil

### UnsetIncidents
`func (o *ActivityStatisticsResultDto) UnsetIncidents()`

UnsetIncidents ensures that no value is present for Incidents, not even an explicit nil
### GetClass

`func (o *ActivityStatisticsResultDto) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *ActivityStatisticsResultDto) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *ActivityStatisticsResultDto) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *ActivityStatisticsResultDto) HasClass() bool`

HasClass returns a boolean if a field has been set.

### SetClassNil

`func (o *ActivityStatisticsResultDto) SetClassNil(b bool)`

 SetClassNil sets the value for Class to be an explicit nil

### UnsetClass
`func (o *ActivityStatisticsResultDto) UnsetClass()`

UnsetClass ensures that no value is present for Class, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


