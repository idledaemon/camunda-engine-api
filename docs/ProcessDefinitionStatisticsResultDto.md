# ProcessDefinitionStatisticsResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the process definition the results are aggregated for. | [optional] 
**Instances** | Pointer to **int32** | The total number of running process instances of this process definition. | [optional] 
**FailedJobs** | Pointer to **int32** | The total number of failed jobs for the running instances. **Note**: Will be &#x60;0&#x60; (not &#x60;null&#x60;), if failed jobs were excluded. | [optional] 
**Incidents** | Pointer to [**[]IncidentStatisticsResultDto**](IncidentStatisticsResultDto.md) | Each item in the resulting array is an object which contains &#x60;incidentType&#x60; and &#x60;incidentCount&#x60;. **Note**: Will be an empty array, if &#x60;incidents&#x60; or &#x60;incidentsForType&#x60; were excluded. Furthermore, the array will be also empty if no incidents were found. | [optional] 
**Class** | Pointer to **NullableString** | The fully qualified class name of the data transfer object class. The class name might change in future releases. | [optional] 
**Definition** | Pointer to [**ProcessDefinitionDto**](ProcessDefinitionDto.md) |  | [optional] 

## Methods

### NewProcessDefinitionStatisticsResultDto

`func NewProcessDefinitionStatisticsResultDto() *ProcessDefinitionStatisticsResultDto`

NewProcessDefinitionStatisticsResultDto instantiates a new ProcessDefinitionStatisticsResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessDefinitionStatisticsResultDtoWithDefaults

`func NewProcessDefinitionStatisticsResultDtoWithDefaults() *ProcessDefinitionStatisticsResultDto`

NewProcessDefinitionStatisticsResultDtoWithDefaults instantiates a new ProcessDefinitionStatisticsResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessDefinitionStatisticsResultDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessDefinitionStatisticsResultDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessDefinitionStatisticsResultDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessDefinitionStatisticsResultDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProcessDefinitionStatisticsResultDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProcessDefinitionStatisticsResultDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInstances

`func (o *ProcessDefinitionStatisticsResultDto) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ProcessDefinitionStatisticsResultDto) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ProcessDefinitionStatisticsResultDto) SetInstances(v int32)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ProcessDefinitionStatisticsResultDto) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetFailedJobs

`func (o *ProcessDefinitionStatisticsResultDto) GetFailedJobs() int32`

GetFailedJobs returns the FailedJobs field if non-nil, zero value otherwise.

### GetFailedJobsOk

`func (o *ProcessDefinitionStatisticsResultDto) GetFailedJobsOk() (*int32, bool)`

GetFailedJobsOk returns a tuple with the FailedJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedJobs

`func (o *ProcessDefinitionStatisticsResultDto) SetFailedJobs(v int32)`

SetFailedJobs sets FailedJobs field to given value.

### HasFailedJobs

`func (o *ProcessDefinitionStatisticsResultDto) HasFailedJobs() bool`

HasFailedJobs returns a boolean if a field has been set.

### GetIncidents

`func (o *ProcessDefinitionStatisticsResultDto) GetIncidents() []IncidentStatisticsResultDto`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *ProcessDefinitionStatisticsResultDto) GetIncidentsOk() (*[]IncidentStatisticsResultDto, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *ProcessDefinitionStatisticsResultDto) SetIncidents(v []IncidentStatisticsResultDto)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *ProcessDefinitionStatisticsResultDto) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### SetIncidentsNil

`func (o *ProcessDefinitionStatisticsResultDto) SetIncidentsNil(b bool)`

 SetIncidentsNil sets the value for Incidents to be an explicit nil

### UnsetIncidents
`func (o *ProcessDefinitionStatisticsResultDto) UnsetIncidents()`

UnsetIncidents ensures that no value is present for Incidents, not even an explicit nil
### GetClass

`func (o *ProcessDefinitionStatisticsResultDto) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *ProcessDefinitionStatisticsResultDto) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *ProcessDefinitionStatisticsResultDto) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *ProcessDefinitionStatisticsResultDto) HasClass() bool`

HasClass returns a boolean if a field has been set.

### SetClassNil

`func (o *ProcessDefinitionStatisticsResultDto) SetClassNil(b bool)`

 SetClassNil sets the value for Class to be an explicit nil

### UnsetClass
`func (o *ProcessDefinitionStatisticsResultDto) UnsetClass()`

UnsetClass ensures that no value is present for Class, not even an explicit nil
### GetDefinition

`func (o *ProcessDefinitionStatisticsResultDto) GetDefinition() ProcessDefinitionDto`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ProcessDefinitionStatisticsResultDto) GetDefinitionOk() (*ProcessDefinitionDto, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ProcessDefinitionStatisticsResultDto) SetDefinition(v ProcessDefinitionDto)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *ProcessDefinitionStatisticsResultDto) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


