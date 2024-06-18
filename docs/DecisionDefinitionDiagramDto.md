# DecisionDefinitionDiagramDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the decision definition. | [optional] 
**DmnXml** | Pointer to **NullableString** | An escaped XML string containing the XML that this decision definition was deployed with. Carriage returns, line feeds and quotation marks are escaped. | [optional] 

## Methods

### NewDecisionDefinitionDiagramDto

`func NewDecisionDefinitionDiagramDto() *DecisionDefinitionDiagramDto`

NewDecisionDefinitionDiagramDto instantiates a new DecisionDefinitionDiagramDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionDefinitionDiagramDtoWithDefaults

`func NewDecisionDefinitionDiagramDtoWithDefaults() *DecisionDefinitionDiagramDto`

NewDecisionDefinitionDiagramDtoWithDefaults instantiates a new DecisionDefinitionDiagramDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DecisionDefinitionDiagramDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DecisionDefinitionDiagramDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DecisionDefinitionDiagramDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DecisionDefinitionDiagramDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DecisionDefinitionDiagramDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DecisionDefinitionDiagramDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDmnXml

`func (o *DecisionDefinitionDiagramDto) GetDmnXml() string`

GetDmnXml returns the DmnXml field if non-nil, zero value otherwise.

### GetDmnXmlOk

`func (o *DecisionDefinitionDiagramDto) GetDmnXmlOk() (*string, bool)`

GetDmnXmlOk returns a tuple with the DmnXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmnXml

`func (o *DecisionDefinitionDiagramDto) SetDmnXml(v string)`

SetDmnXml sets DmnXml field to given value.

### HasDmnXml

`func (o *DecisionDefinitionDiagramDto) HasDmnXml() bool`

HasDmnXml returns a boolean if a field has been set.

### SetDmnXmlNil

`func (o *DecisionDefinitionDiagramDto) SetDmnXmlNil(b bool)`

 SetDmnXmlNil sets the value for DmnXml to be an explicit nil

### UnsetDmnXml
`func (o *DecisionDefinitionDiagramDto) UnsetDmnXml()`

UnsetDmnXml ensures that no value is present for DmnXml, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


