# ActivityInstanceIncidentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the incident. | [optional] 
**ActivityId** | Pointer to **NullableString** | The activity id in which the incident happened. | [optional] 

## Methods

### NewActivityInstanceIncidentDto

`func NewActivityInstanceIncidentDto() *ActivityInstanceIncidentDto`

NewActivityInstanceIncidentDto instantiates a new ActivityInstanceIncidentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityInstanceIncidentDtoWithDefaults

`func NewActivityInstanceIncidentDtoWithDefaults() *ActivityInstanceIncidentDto`

NewActivityInstanceIncidentDtoWithDefaults instantiates a new ActivityInstanceIncidentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityInstanceIncidentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityInstanceIncidentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityInstanceIncidentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityInstanceIncidentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ActivityInstanceIncidentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ActivityInstanceIncidentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetActivityId

`func (o *ActivityInstanceIncidentDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *ActivityInstanceIncidentDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *ActivityInstanceIncidentDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *ActivityInstanceIncidentDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *ActivityInstanceIncidentDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *ActivityInstanceIncidentDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


