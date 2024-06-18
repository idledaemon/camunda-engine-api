# MigrationInstructionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceActivityIds** | Pointer to **[]string** | The activity ids from the source process definition being mapped. | [optional] 
**TargetActivityIds** | Pointer to **[]string** | The activity ids from the target process definition being mapped. | [optional] 
**UpdateEventTrigger** | Pointer to **NullableBool** | Configuration flag whether event triggers defined are going to be updated during migration. | [optional] 

## Methods

### NewMigrationInstructionDto

`func NewMigrationInstructionDto() *MigrationInstructionDto`

NewMigrationInstructionDto instantiates a new MigrationInstructionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationInstructionDtoWithDefaults

`func NewMigrationInstructionDtoWithDefaults() *MigrationInstructionDto`

NewMigrationInstructionDtoWithDefaults instantiates a new MigrationInstructionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceActivityIds

`func (o *MigrationInstructionDto) GetSourceActivityIds() []string`

GetSourceActivityIds returns the SourceActivityIds field if non-nil, zero value otherwise.

### GetSourceActivityIdsOk

`func (o *MigrationInstructionDto) GetSourceActivityIdsOk() (*[]string, bool)`

GetSourceActivityIdsOk returns a tuple with the SourceActivityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceActivityIds

`func (o *MigrationInstructionDto) SetSourceActivityIds(v []string)`

SetSourceActivityIds sets SourceActivityIds field to given value.

### HasSourceActivityIds

`func (o *MigrationInstructionDto) HasSourceActivityIds() bool`

HasSourceActivityIds returns a boolean if a field has been set.

### SetSourceActivityIdsNil

`func (o *MigrationInstructionDto) SetSourceActivityIdsNil(b bool)`

 SetSourceActivityIdsNil sets the value for SourceActivityIds to be an explicit nil

### UnsetSourceActivityIds
`func (o *MigrationInstructionDto) UnsetSourceActivityIds()`

UnsetSourceActivityIds ensures that no value is present for SourceActivityIds, not even an explicit nil
### GetTargetActivityIds

`func (o *MigrationInstructionDto) GetTargetActivityIds() []string`

GetTargetActivityIds returns the TargetActivityIds field if non-nil, zero value otherwise.

### GetTargetActivityIdsOk

`func (o *MigrationInstructionDto) GetTargetActivityIdsOk() (*[]string, bool)`

GetTargetActivityIdsOk returns a tuple with the TargetActivityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetActivityIds

`func (o *MigrationInstructionDto) SetTargetActivityIds(v []string)`

SetTargetActivityIds sets TargetActivityIds field to given value.

### HasTargetActivityIds

`func (o *MigrationInstructionDto) HasTargetActivityIds() bool`

HasTargetActivityIds returns a boolean if a field has been set.

### SetTargetActivityIdsNil

`func (o *MigrationInstructionDto) SetTargetActivityIdsNil(b bool)`

 SetTargetActivityIdsNil sets the value for TargetActivityIds to be an explicit nil

### UnsetTargetActivityIds
`func (o *MigrationInstructionDto) UnsetTargetActivityIds()`

UnsetTargetActivityIds ensures that no value is present for TargetActivityIds, not even an explicit nil
### GetUpdateEventTrigger

`func (o *MigrationInstructionDto) GetUpdateEventTrigger() bool`

GetUpdateEventTrigger returns the UpdateEventTrigger field if non-nil, zero value otherwise.

### GetUpdateEventTriggerOk

`func (o *MigrationInstructionDto) GetUpdateEventTriggerOk() (*bool, bool)`

GetUpdateEventTriggerOk returns a tuple with the UpdateEventTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateEventTrigger

`func (o *MigrationInstructionDto) SetUpdateEventTrigger(v bool)`

SetUpdateEventTrigger sets UpdateEventTrigger field to given value.

### HasUpdateEventTrigger

`func (o *MigrationInstructionDto) HasUpdateEventTrigger() bool`

HasUpdateEventTrigger returns a boolean if a field has been set.

### SetUpdateEventTriggerNil

`func (o *MigrationInstructionDto) SetUpdateEventTriggerNil(b bool)`

 SetUpdateEventTriggerNil sets the value for UpdateEventTrigger to be an explicit nil

### UnsetUpdateEventTrigger
`func (o *MigrationInstructionDto) UnsetUpdateEventTrigger()`

UnsetUpdateEventTrigger ensures that no value is present for UpdateEventTrigger, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


