# SummarizeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | [**StructuredSummary**](StructuredSummary.md) | Structured AI-generated summary with flow, bullets, and optional notes | 
**TextRepr** | **string** | Line-numbered text representation that the summary references | 
**Metadata** | Pointer to **interface{}** | Metadata about the summarization | [optional] 

## Methods

### NewSummarizeResponse

`func NewSummarizeResponse(summary StructuredSummary, textRepr string, ) *SummarizeResponse`

NewSummarizeResponse instantiates a new SummarizeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummarizeResponseWithDefaults

`func NewSummarizeResponseWithDefaults() *SummarizeResponse`

NewSummarizeResponseWithDefaults instantiates a new SummarizeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *SummarizeResponse) GetSummary() StructuredSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SummarizeResponse) GetSummaryOk() (*StructuredSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SummarizeResponse) SetSummary(v StructuredSummary)`

SetSummary sets Summary field to given value.


### GetTextRepr

`func (o *SummarizeResponse) GetTextRepr() string`

GetTextRepr returns the TextRepr field if non-nil, zero value otherwise.

### GetTextReprOk

`func (o *SummarizeResponse) GetTextReprOk() (*string, bool)`

GetTextReprOk returns a tuple with the TextRepr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextRepr

`func (o *SummarizeResponse) SetTextRepr(v string)`

SetTextRepr sets TextRepr field to given value.


### GetMetadata

`func (o *SummarizeResponse) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SummarizeResponse) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SummarizeResponse) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SummarizeResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *SummarizeResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SummarizeResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


