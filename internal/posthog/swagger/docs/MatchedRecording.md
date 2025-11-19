# MatchedRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]MatchedRecordingEvent**](MatchedRecordingEvent.md) |  | 
**SessionId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMatchedRecording

`func NewMatchedRecording(events []MatchedRecordingEvent, ) *MatchedRecording`

NewMatchedRecording instantiates a new MatchedRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchedRecordingWithDefaults

`func NewMatchedRecordingWithDefaults() *MatchedRecording`

NewMatchedRecordingWithDefaults instantiates a new MatchedRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *MatchedRecording) GetEvents() []MatchedRecordingEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MatchedRecording) GetEventsOk() (*[]MatchedRecordingEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *MatchedRecording) SetEvents(v []MatchedRecordingEvent)`

SetEvents sets Events field to given value.


### GetSessionId

`func (o *MatchedRecording) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *MatchedRecording) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *MatchedRecording) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *MatchedRecording) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *MatchedRecording) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *MatchedRecording) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


