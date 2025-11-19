# TimelineEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]EventType**](EventType.md) |  | 
**RecordingDurationS** | Pointer to **NullableFloat32** | Duration of the recording in seconds. | [optional] 
**SessionId** | Pointer to **NullableString** | Session ID. None means out-of-session events | [optional] 

## Methods

### NewTimelineEntry

`func NewTimelineEntry(events []EventType, ) *TimelineEntry`

NewTimelineEntry instantiates a new TimelineEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEntryWithDefaults

`func NewTimelineEntryWithDefaults() *TimelineEntry`

NewTimelineEntryWithDefaults instantiates a new TimelineEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *TimelineEntry) GetEvents() []EventType`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TimelineEntry) GetEventsOk() (*[]EventType, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TimelineEntry) SetEvents(v []EventType)`

SetEvents sets Events field to given value.


### GetRecordingDurationS

`func (o *TimelineEntry) GetRecordingDurationS() float32`

GetRecordingDurationS returns the RecordingDurationS field if non-nil, zero value otherwise.

### GetRecordingDurationSOk

`func (o *TimelineEntry) GetRecordingDurationSOk() (*float32, bool)`

GetRecordingDurationSOk returns a tuple with the RecordingDurationS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingDurationS

`func (o *TimelineEntry) SetRecordingDurationS(v float32)`

SetRecordingDurationS sets RecordingDurationS field to given value.

### HasRecordingDurationS

`func (o *TimelineEntry) HasRecordingDurationS() bool`

HasRecordingDurationS returns a boolean if a field has been set.

### SetRecordingDurationSNil

`func (o *TimelineEntry) SetRecordingDurationSNil(b bool)`

 SetRecordingDurationSNil sets the value for RecordingDurationS to be an explicit nil

### UnsetRecordingDurationS
`func (o *TimelineEntry) UnsetRecordingDurationS()`

UnsetRecordingDurationS ensures that no value is present for RecordingDurationS, not even an explicit nil
### GetSessionId

`func (o *TimelineEntry) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *TimelineEntry) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *TimelineEntry) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *TimelineEntry) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *TimelineEntry) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *TimelineEntry) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


