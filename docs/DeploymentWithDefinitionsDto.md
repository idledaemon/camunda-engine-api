# DeploymentWithDefinitionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the deployment. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id of the deployment. | [optional] 
**DeploymentTime** | Pointer to **NullableTime** | The time when the deployment was created. | [optional] 
**Source** | Pointer to **NullableString** | The source of the deployment. | [optional] 
**Name** | Pointer to **NullableString** | The name of the deployment. | [optional] 
**Links** | Pointer to [**[]AtomLink**](AtomLink.md) | The links associated to this resource, with &#x60;method&#x60;, &#x60;href&#x60; and &#x60;rel&#x60;. | [optional] 
**DeployedProcessDefinitions** | Pointer to [**map[string]ProcessDefinitionDto**](ProcessDefinitionDto.md) | A JSON Object containing a property for each of the process definitions, which are successfully deployed with that deployment. The key is the process definition id, the value is a JSON Object corresponding to the process definition. | [optional] 
**DeployedDecisionDefinitions** | Pointer to [**map[string]DecisionDefinitionDto**](DecisionDefinitionDto.md) | A JSON Object containing a property for each of the decision definitions, which are successfully deployed with that deployment. The key is the decision definition id, the value is a JSON Object corresponding to the decision definition. | [optional] 
**DeployedDecisionRequirementsDefinitions** | Pointer to [**map[string]DecisionRequirementsDefinitionDto**](DecisionRequirementsDefinitionDto.md) | A JSON Object containing a property for each of the decision requirements definitions, which are successfully deployed with that deployment. The key is the decision requirements definition id, the value is a JSON Object corresponding to the decision requirements definition. | [optional] 
**DeployedCaseDefinitions** | Pointer to [**map[string]CaseDefinitionDto**](CaseDefinitionDto.md) | A JSON Object containing a property for each of the case definitions, which are successfully deployed with that deployment. The key is the case definition id, the value is a JSON Object corresponding to the case definition. | [optional] 

## Methods

### NewDeploymentWithDefinitionsDto

`func NewDeploymentWithDefinitionsDto() *DeploymentWithDefinitionsDto`

NewDeploymentWithDefinitionsDto instantiates a new DeploymentWithDefinitionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefinitionsDtoWithDefaults

`func NewDeploymentWithDefinitionsDtoWithDefaults() *DeploymentWithDefinitionsDto`

NewDeploymentWithDefinitionsDtoWithDefaults instantiates a new DeploymentWithDefinitionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentWithDefinitionsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentWithDefinitionsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentWithDefinitionsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentWithDefinitionsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeploymentWithDefinitionsDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeploymentWithDefinitionsDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTenantId

`func (o *DeploymentWithDefinitionsDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DeploymentWithDefinitionsDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DeploymentWithDefinitionsDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DeploymentWithDefinitionsDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DeploymentWithDefinitionsDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DeploymentWithDefinitionsDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDeploymentTime

`func (o *DeploymentWithDefinitionsDto) GetDeploymentTime() time.Time`

GetDeploymentTime returns the DeploymentTime field if non-nil, zero value otherwise.

### GetDeploymentTimeOk

`func (o *DeploymentWithDefinitionsDto) GetDeploymentTimeOk() (*time.Time, bool)`

GetDeploymentTimeOk returns a tuple with the DeploymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTime

`func (o *DeploymentWithDefinitionsDto) SetDeploymentTime(v time.Time)`

SetDeploymentTime sets DeploymentTime field to given value.

### HasDeploymentTime

`func (o *DeploymentWithDefinitionsDto) HasDeploymentTime() bool`

HasDeploymentTime returns a boolean if a field has been set.

### SetDeploymentTimeNil

`func (o *DeploymentWithDefinitionsDto) SetDeploymentTimeNil(b bool)`

 SetDeploymentTimeNil sets the value for DeploymentTime to be an explicit nil

### UnsetDeploymentTime
`func (o *DeploymentWithDefinitionsDto) UnsetDeploymentTime()`

UnsetDeploymentTime ensures that no value is present for DeploymentTime, not even an explicit nil
### GetSource

`func (o *DeploymentWithDefinitionsDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeploymentWithDefinitionsDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeploymentWithDefinitionsDto) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DeploymentWithDefinitionsDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *DeploymentWithDefinitionsDto) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *DeploymentWithDefinitionsDto) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetName

`func (o *DeploymentWithDefinitionsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentWithDefinitionsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentWithDefinitionsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentWithDefinitionsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeploymentWithDefinitionsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeploymentWithDefinitionsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLinks

`func (o *DeploymentWithDefinitionsDto) GetLinks() []AtomLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeploymentWithDefinitionsDto) GetLinksOk() (*[]AtomLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeploymentWithDefinitionsDto) SetLinks(v []AtomLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeploymentWithDefinitionsDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *DeploymentWithDefinitionsDto) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *DeploymentWithDefinitionsDto) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetDeployedProcessDefinitions

`func (o *DeploymentWithDefinitionsDto) GetDeployedProcessDefinitions() map[string]ProcessDefinitionDto`

GetDeployedProcessDefinitions returns the DeployedProcessDefinitions field if non-nil, zero value otherwise.

### GetDeployedProcessDefinitionsOk

`func (o *DeploymentWithDefinitionsDto) GetDeployedProcessDefinitionsOk() (*map[string]ProcessDefinitionDto, bool)`

GetDeployedProcessDefinitionsOk returns a tuple with the DeployedProcessDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedProcessDefinitions

`func (o *DeploymentWithDefinitionsDto) SetDeployedProcessDefinitions(v map[string]ProcessDefinitionDto)`

SetDeployedProcessDefinitions sets DeployedProcessDefinitions field to given value.

### HasDeployedProcessDefinitions

`func (o *DeploymentWithDefinitionsDto) HasDeployedProcessDefinitions() bool`

HasDeployedProcessDefinitions returns a boolean if a field has been set.

### SetDeployedProcessDefinitionsNil

`func (o *DeploymentWithDefinitionsDto) SetDeployedProcessDefinitionsNil(b bool)`

 SetDeployedProcessDefinitionsNil sets the value for DeployedProcessDefinitions to be an explicit nil

### UnsetDeployedProcessDefinitions
`func (o *DeploymentWithDefinitionsDto) UnsetDeployedProcessDefinitions()`

UnsetDeployedProcessDefinitions ensures that no value is present for DeployedProcessDefinitions, not even an explicit nil
### GetDeployedDecisionDefinitions

`func (o *DeploymentWithDefinitionsDto) GetDeployedDecisionDefinitions() map[string]DecisionDefinitionDto`

GetDeployedDecisionDefinitions returns the DeployedDecisionDefinitions field if non-nil, zero value otherwise.

### GetDeployedDecisionDefinitionsOk

`func (o *DeploymentWithDefinitionsDto) GetDeployedDecisionDefinitionsOk() (*map[string]DecisionDefinitionDto, bool)`

GetDeployedDecisionDefinitionsOk returns a tuple with the DeployedDecisionDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedDecisionDefinitions

`func (o *DeploymentWithDefinitionsDto) SetDeployedDecisionDefinitions(v map[string]DecisionDefinitionDto)`

SetDeployedDecisionDefinitions sets DeployedDecisionDefinitions field to given value.

### HasDeployedDecisionDefinitions

`func (o *DeploymentWithDefinitionsDto) HasDeployedDecisionDefinitions() bool`

HasDeployedDecisionDefinitions returns a boolean if a field has been set.

### SetDeployedDecisionDefinitionsNil

`func (o *DeploymentWithDefinitionsDto) SetDeployedDecisionDefinitionsNil(b bool)`

 SetDeployedDecisionDefinitionsNil sets the value for DeployedDecisionDefinitions to be an explicit nil

### UnsetDeployedDecisionDefinitions
`func (o *DeploymentWithDefinitionsDto) UnsetDeployedDecisionDefinitions()`

UnsetDeployedDecisionDefinitions ensures that no value is present for DeployedDecisionDefinitions, not even an explicit nil
### GetDeployedDecisionRequirementsDefinitions

`func (o *DeploymentWithDefinitionsDto) GetDeployedDecisionRequirementsDefinitions() map[string]DecisionRequirementsDefinitionDto`

GetDeployedDecisionRequirementsDefinitions returns the DeployedDecisionRequirementsDefinitions field if non-nil, zero value otherwise.

### GetDeployedDecisionRequirementsDefinitionsOk

`func (o *DeploymentWithDefinitionsDto) GetDeployedDecisionRequirementsDefinitionsOk() (*map[string]DecisionRequirementsDefinitionDto, bool)`

GetDeployedDecisionRequirementsDefinitionsOk returns a tuple with the DeployedDecisionRequirementsDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedDecisionRequirementsDefinitions

`func (o *DeploymentWithDefinitionsDto) SetDeployedDecisionRequirementsDefinitions(v map[string]DecisionRequirementsDefinitionDto)`

SetDeployedDecisionRequirementsDefinitions sets DeployedDecisionRequirementsDefinitions field to given value.

### HasDeployedDecisionRequirementsDefinitions

`func (o *DeploymentWithDefinitionsDto) HasDeployedDecisionRequirementsDefinitions() bool`

HasDeployedDecisionRequirementsDefinitions returns a boolean if a field has been set.

### SetDeployedDecisionRequirementsDefinitionsNil

`func (o *DeploymentWithDefinitionsDto) SetDeployedDecisionRequirementsDefinitionsNil(b bool)`

 SetDeployedDecisionRequirementsDefinitionsNil sets the value for DeployedDecisionRequirementsDefinitions to be an explicit nil

### UnsetDeployedDecisionRequirementsDefinitions
`func (o *DeploymentWithDefinitionsDto) UnsetDeployedDecisionRequirementsDefinitions()`

UnsetDeployedDecisionRequirementsDefinitions ensures that no value is present for DeployedDecisionRequirementsDefinitions, not even an explicit nil
### GetDeployedCaseDefinitions

`func (o *DeploymentWithDefinitionsDto) GetDeployedCaseDefinitions() map[string]CaseDefinitionDto`

GetDeployedCaseDefinitions returns the DeployedCaseDefinitions field if non-nil, zero value otherwise.

### GetDeployedCaseDefinitionsOk

`func (o *DeploymentWithDefinitionsDto) GetDeployedCaseDefinitionsOk() (*map[string]CaseDefinitionDto, bool)`

GetDeployedCaseDefinitionsOk returns a tuple with the DeployedCaseDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCaseDefinitions

`func (o *DeploymentWithDefinitionsDto) SetDeployedCaseDefinitions(v map[string]CaseDefinitionDto)`

SetDeployedCaseDefinitions sets DeployedCaseDefinitions field to given value.

### HasDeployedCaseDefinitions

`func (o *DeploymentWithDefinitionsDto) HasDeployedCaseDefinitions() bool`

HasDeployedCaseDefinitions returns a boolean if a field has been set.

### SetDeployedCaseDefinitionsNil

`func (o *DeploymentWithDefinitionsDto) SetDeployedCaseDefinitionsNil(b bool)`

 SetDeployedCaseDefinitionsNil sets the value for DeployedCaseDefinitions to be an explicit nil

### UnsetDeployedCaseDefinitions
`func (o *DeploymentWithDefinitionsDto) UnsetDeployedCaseDefinitions()`

UnsetDeployedCaseDefinitions ensures that no value is present for DeployedCaseDefinitions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


