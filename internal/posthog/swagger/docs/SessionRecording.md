# SessionRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**DistinctId** | **NullableString** |  | [readonly] 
**Viewed** | **bool** |  | [readonly] 
**Viewers** | **[]string** |  | [readonly] 
**RecordingDuration** | **int32** |  | [readonly] 
**ActiveSeconds** | **NullableInt32** |  | [readonly] 
**InactiveSeconds** | **NullableInt32** |  | [readonly] 
**StartTime** | **NullableTime** |  | [readonly] 
**EndTime** | **NullableTime** |  | [readonly] 
**ClickCount** | **NullableInt32** |  | [readonly] 
**KeypressCount** | **NullableInt32** |  | [readonly] 
**MouseActivityCount** | **NullableInt32** |  | [readonly] 
**ConsoleLogCount** | **NullableInt32** |  | [readonly] 
**ConsoleWarnCount** | **NullableInt32** |  | [readonly] 
**ConsoleErrorCount** | **NullableInt32** |  | [readonly] 
**StartUrl** | **NullableString** |  | [readonly] 
**Person** | Pointer to [**MinimalPerson**](MinimalPerson.md) |  | [optional] 
**Storage** | **string** |  | [readonly] 
**RetentionPeriodDays** | **NullableInt32** |  | [readonly] 
**ExpiryTime** | **string** |  | [readonly] 
**RecordingTtl** | **string** |  | [readonly] 
**SnapshotSource** | **NullableString** |  | [readonly] 
**Ongoing** | **bool** |  | [readonly] 
**ActivityScore** | **NullableFloat64** |  | [readonly] 

## Methods

### NewSessionRecording

`func NewSessionRecording(id string, distinctId NullableString, viewed bool, viewers []string, recordingDuration int32, activeSeconds NullableInt32, inactiveSeconds NullableInt32, startTime NullableTime, endTime NullableTime, clickCount NullableInt32, keypressCount NullableInt32, mouseActivityCount NullableInt32, consoleLogCount NullableInt32, consoleWarnCount NullableInt32, consoleErrorCount NullableInt32, startUrl NullableString, storage string, retentionPeriodDays NullableInt32, expiryTime string, recordingTtl string, snapshotSource NullableString, ongoing bool, activityScore NullableFloat64, ) *SessionRecording`

NewSessionRecording instantiates a new SessionRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRecordingWithDefaults

`func NewSessionRecordingWithDefaults() *SessionRecording`

NewSessionRecordingWithDefaults instantiates a new SessionRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionRecording) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionRecording) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionRecording) SetId(v string)`

SetId sets Id field to given value.


### GetDistinctId

`func (o *SessionRecording) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *SessionRecording) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *SessionRecording) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.


### SetDistinctIdNil

`func (o *SessionRecording) SetDistinctIdNil(b bool)`

 SetDistinctIdNil sets the value for DistinctId to be an explicit nil

### UnsetDistinctId
`func (o *SessionRecording) UnsetDistinctId()`

UnsetDistinctId ensures that no value is present for DistinctId, not even an explicit nil
### GetViewed

`func (o *SessionRecording) GetViewed() bool`

GetViewed returns the Viewed field if non-nil, zero value otherwise.

### GetViewedOk

`func (o *SessionRecording) GetViewedOk() (*bool, bool)`

GetViewedOk returns a tuple with the Viewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewed

`func (o *SessionRecording) SetViewed(v bool)`

SetViewed sets Viewed field to given value.


### GetViewers

`func (o *SessionRecording) GetViewers() []string`

GetViewers returns the Viewers field if non-nil, zero value otherwise.

### GetViewersOk

`func (o *SessionRecording) GetViewersOk() (*[]string, bool)`

GetViewersOk returns a tuple with the Viewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewers

`func (o *SessionRecording) SetViewers(v []string)`

SetViewers sets Viewers field to given value.


### GetRecordingDuration

`func (o *SessionRecording) GetRecordingDuration() int32`

GetRecordingDuration returns the RecordingDuration field if non-nil, zero value otherwise.

### GetRecordingDurationOk

`func (o *SessionRecording) GetRecordingDurationOk() (*int32, bool)`

GetRecordingDurationOk returns a tuple with the RecordingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingDuration

`func (o *SessionRecording) SetRecordingDuration(v int32)`

SetRecordingDuration sets RecordingDuration field to given value.


### GetActiveSeconds

`func (o *SessionRecording) GetActiveSeconds() int32`

GetActiveSeconds returns the ActiveSeconds field if non-nil, zero value otherwise.

### GetActiveSecondsOk

`func (o *SessionRecording) GetActiveSecondsOk() (*int32, bool)`

GetActiveSecondsOk returns a tuple with the ActiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSeconds

`func (o *SessionRecording) SetActiveSeconds(v int32)`

SetActiveSeconds sets ActiveSeconds field to given value.


### SetActiveSecondsNil

`func (o *SessionRecording) SetActiveSecondsNil(b bool)`

 SetActiveSecondsNil sets the value for ActiveSeconds to be an explicit nil

### UnsetActiveSeconds
`func (o *SessionRecording) UnsetActiveSeconds()`

UnsetActiveSeconds ensures that no value is present for ActiveSeconds, not even an explicit nil
### GetInactiveSeconds

`func (o *SessionRecording) GetInactiveSeconds() int32`

GetInactiveSeconds returns the InactiveSeconds field if non-nil, zero value otherwise.

### GetInactiveSecondsOk

`func (o *SessionRecording) GetInactiveSecondsOk() (*int32, bool)`

GetInactiveSecondsOk returns a tuple with the InactiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveSeconds

`func (o *SessionRecording) SetInactiveSeconds(v int32)`

SetInactiveSeconds sets InactiveSeconds field to given value.


### SetInactiveSecondsNil

`func (o *SessionRecording) SetInactiveSecondsNil(b bool)`

 SetInactiveSecondsNil sets the value for InactiveSeconds to be an explicit nil

### UnsetInactiveSeconds
`func (o *SessionRecording) UnsetInactiveSeconds()`

UnsetInactiveSeconds ensures that no value is present for InactiveSeconds, not even an explicit nil
### GetStartTime

`func (o *SessionRecording) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SessionRecording) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SessionRecording) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### SetStartTimeNil

`func (o *SessionRecording) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *SessionRecording) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *SessionRecording) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SessionRecording) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SessionRecording) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### SetEndTimeNil

`func (o *SessionRecording) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *SessionRecording) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetClickCount

`func (o *SessionRecording) GetClickCount() int32`

GetClickCount returns the ClickCount field if non-nil, zero value otherwise.

### GetClickCountOk

`func (o *SessionRecording) GetClickCountOk() (*int32, bool)`

GetClickCountOk returns a tuple with the ClickCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickCount

`func (o *SessionRecording) SetClickCount(v int32)`

SetClickCount sets ClickCount field to given value.


### SetClickCountNil

`func (o *SessionRecording) SetClickCountNil(b bool)`

 SetClickCountNil sets the value for ClickCount to be an explicit nil

### UnsetClickCount
`func (o *SessionRecording) UnsetClickCount()`

UnsetClickCount ensures that no value is present for ClickCount, not even an explicit nil
### GetKeypressCount

`func (o *SessionRecording) GetKeypressCount() int32`

GetKeypressCount returns the KeypressCount field if non-nil, zero value otherwise.

### GetKeypressCountOk

`func (o *SessionRecording) GetKeypressCountOk() (*int32, bool)`

GetKeypressCountOk returns a tuple with the KeypressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypressCount

`func (o *SessionRecording) SetKeypressCount(v int32)`

SetKeypressCount sets KeypressCount field to given value.


### SetKeypressCountNil

`func (o *SessionRecording) SetKeypressCountNil(b bool)`

 SetKeypressCountNil sets the value for KeypressCount to be an explicit nil

### UnsetKeypressCount
`func (o *SessionRecording) UnsetKeypressCount()`

UnsetKeypressCount ensures that no value is present for KeypressCount, not even an explicit nil
### GetMouseActivityCount

`func (o *SessionRecording) GetMouseActivityCount() int32`

GetMouseActivityCount returns the MouseActivityCount field if non-nil, zero value otherwise.

### GetMouseActivityCountOk

`func (o *SessionRecording) GetMouseActivityCountOk() (*int32, bool)`

GetMouseActivityCountOk returns a tuple with the MouseActivityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMouseActivityCount

`func (o *SessionRecording) SetMouseActivityCount(v int32)`

SetMouseActivityCount sets MouseActivityCount field to given value.


### SetMouseActivityCountNil

`func (o *SessionRecording) SetMouseActivityCountNil(b bool)`

 SetMouseActivityCountNil sets the value for MouseActivityCount to be an explicit nil

### UnsetMouseActivityCount
`func (o *SessionRecording) UnsetMouseActivityCount()`

UnsetMouseActivityCount ensures that no value is present for MouseActivityCount, not even an explicit nil
### GetConsoleLogCount

`func (o *SessionRecording) GetConsoleLogCount() int32`

GetConsoleLogCount returns the ConsoleLogCount field if non-nil, zero value otherwise.

### GetConsoleLogCountOk

`func (o *SessionRecording) GetConsoleLogCountOk() (*int32, bool)`

GetConsoleLogCountOk returns a tuple with the ConsoleLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleLogCount

`func (o *SessionRecording) SetConsoleLogCount(v int32)`

SetConsoleLogCount sets ConsoleLogCount field to given value.


### SetConsoleLogCountNil

`func (o *SessionRecording) SetConsoleLogCountNil(b bool)`

 SetConsoleLogCountNil sets the value for ConsoleLogCount to be an explicit nil

### UnsetConsoleLogCount
`func (o *SessionRecording) UnsetConsoleLogCount()`

UnsetConsoleLogCount ensures that no value is present for ConsoleLogCount, not even an explicit nil
### GetConsoleWarnCount

`func (o *SessionRecording) GetConsoleWarnCount() int32`

GetConsoleWarnCount returns the ConsoleWarnCount field if non-nil, zero value otherwise.

### GetConsoleWarnCountOk

`func (o *SessionRecording) GetConsoleWarnCountOk() (*int32, bool)`

GetConsoleWarnCountOk returns a tuple with the ConsoleWarnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleWarnCount

`func (o *SessionRecording) SetConsoleWarnCount(v int32)`

SetConsoleWarnCount sets ConsoleWarnCount field to given value.


### SetConsoleWarnCountNil

`func (o *SessionRecording) SetConsoleWarnCountNil(b bool)`

 SetConsoleWarnCountNil sets the value for ConsoleWarnCount to be an explicit nil

### UnsetConsoleWarnCount
`func (o *SessionRecording) UnsetConsoleWarnCount()`

UnsetConsoleWarnCount ensures that no value is present for ConsoleWarnCount, not even an explicit nil
### GetConsoleErrorCount

`func (o *SessionRecording) GetConsoleErrorCount() int32`

GetConsoleErrorCount returns the ConsoleErrorCount field if non-nil, zero value otherwise.

### GetConsoleErrorCountOk

`func (o *SessionRecording) GetConsoleErrorCountOk() (*int32, bool)`

GetConsoleErrorCountOk returns a tuple with the ConsoleErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleErrorCount

`func (o *SessionRecording) SetConsoleErrorCount(v int32)`

SetConsoleErrorCount sets ConsoleErrorCount field to given value.


### SetConsoleErrorCountNil

`func (o *SessionRecording) SetConsoleErrorCountNil(b bool)`

 SetConsoleErrorCountNil sets the value for ConsoleErrorCount to be an explicit nil

### UnsetConsoleErrorCount
`func (o *SessionRecording) UnsetConsoleErrorCount()`

UnsetConsoleErrorCount ensures that no value is present for ConsoleErrorCount, not even an explicit nil
### GetStartUrl

`func (o *SessionRecording) GetStartUrl() string`

GetStartUrl returns the StartUrl field if non-nil, zero value otherwise.

### GetStartUrlOk

`func (o *SessionRecording) GetStartUrlOk() (*string, bool)`

GetStartUrlOk returns a tuple with the StartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUrl

`func (o *SessionRecording) SetStartUrl(v string)`

SetStartUrl sets StartUrl field to given value.


### SetStartUrlNil

`func (o *SessionRecording) SetStartUrlNil(b bool)`

 SetStartUrlNil sets the value for StartUrl to be an explicit nil

### UnsetStartUrl
`func (o *SessionRecording) UnsetStartUrl()`

UnsetStartUrl ensures that no value is present for StartUrl, not even an explicit nil
### GetPerson

`func (o *SessionRecording) GetPerson() MinimalPerson`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *SessionRecording) GetPersonOk() (*MinimalPerson, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *SessionRecording) SetPerson(v MinimalPerson)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *SessionRecording) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetStorage

`func (o *SessionRecording) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *SessionRecording) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *SessionRecording) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetRetentionPeriodDays

`func (o *SessionRecording) GetRetentionPeriodDays() int32`

GetRetentionPeriodDays returns the RetentionPeriodDays field if non-nil, zero value otherwise.

### GetRetentionPeriodDaysOk

`func (o *SessionRecording) GetRetentionPeriodDaysOk() (*int32, bool)`

GetRetentionPeriodDaysOk returns a tuple with the RetentionPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDays

`func (o *SessionRecording) SetRetentionPeriodDays(v int32)`

SetRetentionPeriodDays sets RetentionPeriodDays field to given value.


### SetRetentionPeriodDaysNil

`func (o *SessionRecording) SetRetentionPeriodDaysNil(b bool)`

 SetRetentionPeriodDaysNil sets the value for RetentionPeriodDays to be an explicit nil

### UnsetRetentionPeriodDays
`func (o *SessionRecording) UnsetRetentionPeriodDays()`

UnsetRetentionPeriodDays ensures that no value is present for RetentionPeriodDays, not even an explicit nil
### GetExpiryTime

`func (o *SessionRecording) GetExpiryTime() string`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *SessionRecording) GetExpiryTimeOk() (*string, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *SessionRecording) SetExpiryTime(v string)`

SetExpiryTime sets ExpiryTime field to given value.


### GetRecordingTtl

`func (o *SessionRecording) GetRecordingTtl() string`

GetRecordingTtl returns the RecordingTtl field if non-nil, zero value otherwise.

### GetRecordingTtlOk

`func (o *SessionRecording) GetRecordingTtlOk() (*string, bool)`

GetRecordingTtlOk returns a tuple with the RecordingTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingTtl

`func (o *SessionRecording) SetRecordingTtl(v string)`

SetRecordingTtl sets RecordingTtl field to given value.


### GetSnapshotSource

`func (o *SessionRecording) GetSnapshotSource() string`

GetSnapshotSource returns the SnapshotSource field if non-nil, zero value otherwise.

### GetSnapshotSourceOk

`func (o *SessionRecording) GetSnapshotSourceOk() (*string, bool)`

GetSnapshotSourceOk returns a tuple with the SnapshotSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSource

`func (o *SessionRecording) SetSnapshotSource(v string)`

SetSnapshotSource sets SnapshotSource field to given value.


### SetSnapshotSourceNil

`func (o *SessionRecording) SetSnapshotSourceNil(b bool)`

 SetSnapshotSourceNil sets the value for SnapshotSource to be an explicit nil

### UnsetSnapshotSource
`func (o *SessionRecording) UnsetSnapshotSource()`

UnsetSnapshotSource ensures that no value is present for SnapshotSource, not even an explicit nil
### GetOngoing

`func (o *SessionRecording) GetOngoing() bool`

GetOngoing returns the Ongoing field if non-nil, zero value otherwise.

### GetOngoingOk

`func (o *SessionRecording) GetOngoingOk() (*bool, bool)`

GetOngoingOk returns a tuple with the Ongoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoing

`func (o *SessionRecording) SetOngoing(v bool)`

SetOngoing sets Ongoing field to given value.


### GetActivityScore

`func (o *SessionRecording) GetActivityScore() float64`

GetActivityScore returns the ActivityScore field if non-nil, zero value otherwise.

### GetActivityScoreOk

`func (o *SessionRecording) GetActivityScoreOk() (*float64, bool)`

GetActivityScoreOk returns a tuple with the ActivityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityScore

`func (o *SessionRecording) SetActivityScore(v float64)`

SetActivityScore sets ActivityScore field to given value.


### SetActivityScoreNil

`func (o *SessionRecording) SetActivityScoreNil(b bool)`

 SetActivityScoreNil sets the value for ActivityScore to be an explicit nil

### UnsetActivityScore
`func (o *SessionRecording) UnsetActivityScore()`

UnsetActivityScore ensures that no value is present for ActivityScore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


