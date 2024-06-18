# LinkableDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]AtomLink**](AtomLink.md) | The links associated to this resource, with &#x60;method&#x60;, &#x60;href&#x60; and &#x60;rel&#x60;. | [optional] 

## Methods

### NewLinkableDto

`func NewLinkableDto() *LinkableDto`

NewLinkableDto instantiates a new LinkableDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkableDtoWithDefaults

`func NewLinkableDtoWithDefaults() *LinkableDto`

NewLinkableDtoWithDefaults instantiates a new LinkableDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *LinkableDto) GetLinks() []AtomLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LinkableDto) GetLinksOk() (*[]AtomLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LinkableDto) SetLinks(v []AtomLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LinkableDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *LinkableDto) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *LinkableDto) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


