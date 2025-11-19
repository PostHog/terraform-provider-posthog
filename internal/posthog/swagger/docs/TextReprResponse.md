# TextReprResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Generated text representation of the event | 
**Metadata** | [**TextReprMetadata**](TextReprMetadata.md) | Metadata about the text representation | 

## Methods

### NewTextReprResponse

`func NewTextReprResponse(text string, metadata TextReprMetadata, ) *TextReprResponse`

NewTextReprResponse instantiates a new TextReprResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextReprResponseWithDefaults

`func NewTextReprResponseWithDefaults() *TextReprResponse`

NewTextReprResponseWithDefaults instantiates a new TextReprResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *TextReprResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TextReprResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TextReprResponse) SetText(v string)`

SetText sets Text field to given value.


### GetMetadata

`func (o *TextReprResponse) GetMetadata() TextReprMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TextReprResponse) GetMetadataOk() (*TextReprMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TextReprResponse) SetMetadata(v TextReprMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


