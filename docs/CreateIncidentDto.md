# CreateIncidentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncidentType** | Pointer to **NullableString** | A type of the new incident. | [optional] 
**Configuration** | Pointer to **NullableString** | A configuration for the new incident. | [optional] 
**Message** | Pointer to **NullableString** | A message for the new incident. | [optional] 

## Methods

### NewCreateIncidentDto

`func NewCreateIncidentDto() *CreateIncidentDto`

NewCreateIncidentDto instantiates a new CreateIncidentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIncidentDtoWithDefaults

`func NewCreateIncidentDtoWithDefaults() *CreateIncidentDto`

NewCreateIncidentDtoWithDefaults instantiates a new CreateIncidentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidentType

`func (o *CreateIncidentDto) GetIncidentType() string`

GetIncidentType returns the IncidentType field if non-nil, zero value otherwise.

### GetIncidentTypeOk

`func (o *CreateIncidentDto) GetIncidentTypeOk() (*string, bool)`

GetIncidentTypeOk returns a tuple with the IncidentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentType

`func (o *CreateIncidentDto) SetIncidentType(v string)`

SetIncidentType sets IncidentType field to given value.

### HasIncidentType

`func (o *CreateIncidentDto) HasIncidentType() bool`

HasIncidentType returns a boolean if a field has been set.

### SetIncidentTypeNil

`func (o *CreateIncidentDto) SetIncidentTypeNil(b bool)`

 SetIncidentTypeNil sets the value for IncidentType to be an explicit nil

### UnsetIncidentType
`func (o *CreateIncidentDto) UnsetIncidentType()`

UnsetIncidentType ensures that no value is present for IncidentType, not even an explicit nil
### GetConfiguration

`func (o *CreateIncidentDto) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateIncidentDto) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateIncidentDto) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateIncidentDto) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *CreateIncidentDto) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *CreateIncidentDto) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetMessage

`func (o *CreateIncidentDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateIncidentDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateIncidentDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateIncidentDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CreateIncidentDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CreateIncidentDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


