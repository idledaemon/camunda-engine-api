# ProcessDefinitionDiagramDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the process definition. | [optional] 
**Bpmn20Xml** | Pointer to **NullableString** | An escaped XML string containing the XML that this definition was deployed with. Carriage returns, line feeds and quotation marks are escaped. | [optional] 

## Methods

### NewProcessDefinitionDiagramDto

`func NewProcessDefinitionDiagramDto() *ProcessDefinitionDiagramDto`

NewProcessDefinitionDiagramDto instantiates a new ProcessDefinitionDiagramDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessDefinitionDiagramDtoWithDefaults

`func NewProcessDefinitionDiagramDtoWithDefaults() *ProcessDefinitionDiagramDto`

NewProcessDefinitionDiagramDtoWithDefaults instantiates a new ProcessDefinitionDiagramDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessDefinitionDiagramDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessDefinitionDiagramDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessDefinitionDiagramDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessDefinitionDiagramDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProcessDefinitionDiagramDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProcessDefinitionDiagramDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetBpmn20Xml

`func (o *ProcessDefinitionDiagramDto) GetBpmn20Xml() string`

GetBpmn20Xml returns the Bpmn20Xml field if non-nil, zero value otherwise.

### GetBpmn20XmlOk

`func (o *ProcessDefinitionDiagramDto) GetBpmn20XmlOk() (*string, bool)`

GetBpmn20XmlOk returns a tuple with the Bpmn20Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpmn20Xml

`func (o *ProcessDefinitionDiagramDto) SetBpmn20Xml(v string)`

SetBpmn20Xml sets Bpmn20Xml field to given value.

### HasBpmn20Xml

`func (o *ProcessDefinitionDiagramDto) HasBpmn20Xml() bool`

HasBpmn20Xml returns a boolean if a field has been set.

### SetBpmn20XmlNil

`func (o *ProcessDefinitionDiagramDto) SetBpmn20XmlNil(b bool)`

 SetBpmn20XmlNil sets the value for Bpmn20Xml to be an explicit nil

### UnsetBpmn20Xml
`func (o *ProcessDefinitionDiagramDto) UnsetBpmn20Xml()`

UnsetBpmn20Xml ensures that no value is present for Bpmn20Xml, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


