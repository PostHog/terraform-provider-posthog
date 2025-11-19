# TranscriptSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **NullableFloat64** | Milliseconds from recording start | [optional] 
**Speaker** | Pointer to **NullableString** |  | [optional] 
**Text** | **string** |  | 
**Confidence** | Pointer to **NullableFloat64** | Transcription confidence score | [optional] 
**IsFinal** | Pointer to **NullableBool** | Whether this is the final version | [optional] 

## Methods

### NewTranscriptSegment

`func NewTranscriptSegment(text string, ) *TranscriptSegment`

NewTranscriptSegment instantiates a new TranscriptSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptSegmentWithDefaults

`func NewTranscriptSegmentWithDefaults() *TranscriptSegment`

NewTranscriptSegmentWithDefaults instantiates a new TranscriptSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *TranscriptSegment) GetTimestamp() float64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TranscriptSegment) GetTimestampOk() (*float64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TranscriptSegment) SetTimestamp(v float64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TranscriptSegment) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TranscriptSegment) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TranscriptSegment) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetSpeaker

`func (o *TranscriptSegment) GetSpeaker() string`

GetSpeaker returns the Speaker field if non-nil, zero value otherwise.

### GetSpeakerOk

`func (o *TranscriptSegment) GetSpeakerOk() (*string, bool)`

GetSpeakerOk returns a tuple with the Speaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeaker

`func (o *TranscriptSegment) SetSpeaker(v string)`

SetSpeaker sets Speaker field to given value.

### HasSpeaker

`func (o *TranscriptSegment) HasSpeaker() bool`

HasSpeaker returns a boolean if a field has been set.

### SetSpeakerNil

`func (o *TranscriptSegment) SetSpeakerNil(b bool)`

 SetSpeakerNil sets the value for Speaker to be an explicit nil

### UnsetSpeaker
`func (o *TranscriptSegment) UnsetSpeaker()`

UnsetSpeaker ensures that no value is present for Speaker, not even an explicit nil
### GetText

`func (o *TranscriptSegment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TranscriptSegment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TranscriptSegment) SetText(v string)`

SetText sets Text field to given value.


### GetConfidence

`func (o *TranscriptSegment) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *TranscriptSegment) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *TranscriptSegment) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *TranscriptSegment) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *TranscriptSegment) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *TranscriptSegment) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetIsFinal

`func (o *TranscriptSegment) GetIsFinal() bool`

GetIsFinal returns the IsFinal field if non-nil, zero value otherwise.

### GetIsFinalOk

`func (o *TranscriptSegment) GetIsFinalOk() (*bool, bool)`

GetIsFinalOk returns a tuple with the IsFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinal

`func (o *TranscriptSegment) SetIsFinal(v bool)`

SetIsFinal sets IsFinal field to given value.

### HasIsFinal

`func (o *TranscriptSegment) HasIsFinal() bool`

HasIsFinal returns a boolean if a field has been set.

### SetIsFinalNil

`func (o *TranscriptSegment) SetIsFinalNil(b bool)`

 SetIsFinalNil sets the value for IsFinal to be an explicit nil

### UnsetIsFinal
`func (o *TranscriptSegment) UnsetIsFinal()`

UnsetIsFinal ensures that no value is present for IsFinal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


