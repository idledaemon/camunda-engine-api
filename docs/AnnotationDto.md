# AnnotationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to **NullableString** | The annotation value to put. | [optional] 

## Methods

### NewAnnotationDto

`func NewAnnotationDto() *AnnotationDto`

NewAnnotationDto instantiates a new AnnotationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationDtoWithDefaults

`func NewAnnotationDtoWithDefaults() *AnnotationDto`

NewAnnotationDtoWithDefaults instantiates a new AnnotationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *AnnotationDto) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *AnnotationDto) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *AnnotationDto) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *AnnotationDto) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotationNil

`func (o *AnnotationDto) SetAnnotationNil(b bool)`

 SetAnnotationNil sets the value for Annotation to be an explicit nil

### UnsetAnnotation
`func (o *AnnotationDto) UnsetAnnotation()`

UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


