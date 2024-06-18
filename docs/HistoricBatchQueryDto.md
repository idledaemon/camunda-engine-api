# HistoricBatchQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **NullableString** | Filter by batch id. | [optional] 
**Type** | Pointer to **NullableString** | Filter by batch type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/batch/#creating-a-batch) for more information about batch types. | [optional] 
**Completed** | Pointer to **NullableBool** |  Filter completed or not completed batches. If the value is &#x60;true&#x60;, only completed batches, i.e., end time is set, are returned. Otherwise, if the value is &#x60;false&#x60;, only running batches, i.e., end time is null, are returned. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Filter by a comma-separated list of tenant ids. A batch matches if it has one of the given tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include batches which belong to no tenant. Value can effectively only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Sorting** | Pointer to [**[]HistoricBatchQueryDtoSortingInner**](HistoricBatchQueryDtoSortingInner.md) | An array of criteria to sort the result by. Each element of the array is                        an object that specifies one ordering. The position in the array                        identifies the rank of an ordering, i.e., whether it is primary, secondary,                        etc. Has no effect for the &#x60;/count&#x60; endpoint | [optional] 

## Methods

### NewHistoricBatchQueryDto

`func NewHistoricBatchQueryDto() *HistoricBatchQueryDto`

NewHistoricBatchQueryDto instantiates a new HistoricBatchQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricBatchQueryDtoWithDefaults

`func NewHistoricBatchQueryDtoWithDefaults() *HistoricBatchQueryDto`

NewHistoricBatchQueryDtoWithDefaults instantiates a new HistoricBatchQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *HistoricBatchQueryDto) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *HistoricBatchQueryDto) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *HistoricBatchQueryDto) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *HistoricBatchQueryDto) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchIdNil

`func (o *HistoricBatchQueryDto) SetBatchIdNil(b bool)`

 SetBatchIdNil sets the value for BatchId to be an explicit nil

### UnsetBatchId
`func (o *HistoricBatchQueryDto) UnsetBatchId()`

UnsetBatchId ensures that no value is present for BatchId, not even an explicit nil
### GetType

`func (o *HistoricBatchQueryDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoricBatchQueryDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoricBatchQueryDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HistoricBatchQueryDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HistoricBatchQueryDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HistoricBatchQueryDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCompleted

`func (o *HistoricBatchQueryDto) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *HistoricBatchQueryDto) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *HistoricBatchQueryDto) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *HistoricBatchQueryDto) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### SetCompletedNil

`func (o *HistoricBatchQueryDto) SetCompletedNil(b bool)`

 SetCompletedNil sets the value for Completed to be an explicit nil

### UnsetCompleted
`func (o *HistoricBatchQueryDto) UnsetCompleted()`

UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
### GetTenantIdIn

`func (o *HistoricBatchQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *HistoricBatchQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *HistoricBatchQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *HistoricBatchQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *HistoricBatchQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *HistoricBatchQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *HistoricBatchQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *HistoricBatchQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *HistoricBatchQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *HistoricBatchQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *HistoricBatchQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *HistoricBatchQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetSorting

`func (o *HistoricBatchQueryDto) GetSorting() []HistoricBatchQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *HistoricBatchQueryDto) GetSortingOk() (*[]HistoricBatchQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *HistoricBatchQueryDto) SetSorting(v []HistoricBatchQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *HistoricBatchQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *HistoricBatchQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *HistoricBatchQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


