# TextReprMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**EventId** | Pointer to **string** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 
**Rendering** | **string** |  | 
**CharCount** | **int32** |  | 
**Truncated** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewTextReprMetadata

`func NewTextReprMetadata(rendering string, charCount int32, truncated bool, ) *TextReprMetadata`

NewTextReprMetadata instantiates a new TextReprMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextReprMetadataWithDefaults

`func NewTextReprMetadataWithDefaults() *TextReprMetadata`

NewTextReprMetadataWithDefaults instantiates a new TextReprMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *TextReprMetadata) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TextReprMetadata) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TextReprMetadata) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TextReprMetadata) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventId

`func (o *TextReprMetadata) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *TextReprMetadata) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *TextReprMetadata) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *TextReprMetadata) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetTraceId

`func (o *TextReprMetadata) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *TextReprMetadata) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *TextReprMetadata) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *TextReprMetadata) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetRendering

`func (o *TextReprMetadata) GetRendering() string`

GetRendering returns the Rendering field if non-nil, zero value otherwise.

### GetRenderingOk

`func (o *TextReprMetadata) GetRenderingOk() (*string, bool)`

GetRenderingOk returns a tuple with the Rendering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendering

`func (o *TextReprMetadata) SetRendering(v string)`

SetRendering sets Rendering field to given value.


### GetCharCount

`func (o *TextReprMetadata) GetCharCount() int32`

GetCharCount returns the CharCount field if non-nil, zero value otherwise.

### GetCharCountOk

`func (o *TextReprMetadata) GetCharCountOk() (*int32, bool)`

GetCharCountOk returns a tuple with the CharCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharCount

`func (o *TextReprMetadata) SetCharCount(v int32)`

SetCharCount sets CharCount field to given value.


### GetTruncated

`func (o *TextReprMetadata) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *TextReprMetadata) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *TextReprMetadata) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.


### GetError

`func (o *TextReprMetadata) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TextReprMetadata) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TextReprMetadata) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TextReprMetadata) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


