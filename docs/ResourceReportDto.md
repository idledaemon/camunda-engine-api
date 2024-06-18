# ResourceReportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ProblemDto**](ProblemDto.md) | A list of errors occurred during parsing. | [optional] 
**Warnings** | Pointer to [**[]ProblemDto**](ProblemDto.md) | A list of warnings occurred during parsing. | [optional] 

## Methods

### NewResourceReportDto

`func NewResourceReportDto() *ResourceReportDto`

NewResourceReportDto instantiates a new ResourceReportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceReportDtoWithDefaults

`func NewResourceReportDtoWithDefaults() *ResourceReportDto`

NewResourceReportDtoWithDefaults instantiates a new ResourceReportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ResourceReportDto) GetErrors() []ProblemDto`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ResourceReportDto) GetErrorsOk() (*[]ProblemDto, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ResourceReportDto) SetErrors(v []ProblemDto)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ResourceReportDto) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *ResourceReportDto) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *ResourceReportDto) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetWarnings

`func (o *ResourceReportDto) GetWarnings() []ProblemDto`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ResourceReportDto) GetWarningsOk() (*[]ProblemDto, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ResourceReportDto) SetWarnings(v []ProblemDto)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ResourceReportDto) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *ResourceReportDto) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *ResourceReportDto) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


