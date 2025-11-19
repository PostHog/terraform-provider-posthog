# SessionRecordingType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**ActivityScore** | Pointer to **NullableFloat32** | calculated on the backend so that we can sort by it, definition may change over time | [optional] 
**ClickCount** | Pointer to **NullableFloat32** |  | [optional] 
**ConsoleErrorCount** | Pointer to **NullableFloat32** |  | [optional] 
**ConsoleLogCount** | Pointer to **NullableFloat32** |  | [optional] 
**ConsoleWarnCount** | Pointer to **NullableFloat32** |  | [optional] 
**DistinctId** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**EndTime** | **string** | When the recording ends in ISO format. | 
**ExpiryTime** | Pointer to **NullableString** | When the recording expires, in ISO format. | [optional] 
**Id** | **string** |  | 
**InactiveSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**KeypressCount** | Pointer to **NullableFloat32** |  | [optional] 
**MatchingEvents** | Pointer to [**[]MatchedRecording**](MatchedRecording.md) | List of matching events. * | [optional] 
**MouseActivityCount** | Pointer to **NullableFloat32** | count of all mouse activity in the recording, not just clicks | [optional] 
**Ongoing** | Pointer to **NullableBool** | whether we have received data for this recording in the last 5 minutes (assumes the recording was loaded from ClickHouse) * | [optional] 
**Person** | Pointer to [**PersonType**](PersonType.md) |  | [optional] 
**RecordingDuration** | **float32** | Length of recording in seconds. | 
**RecordingTtl** | Pointer to **NullableFloat32** | Number of whole days left until the recording expires. | [optional] 
**RetentionPeriodDays** | Pointer to **NullableFloat32** | retention period for this recording | [optional] 
**SnapshotSource** | [**SnapshotSource**](SnapshotSource.md) |  | 
**StartTime** | **string** | When the recording starts in ISO format. | 
**StartUrl** | Pointer to **NullableString** |  | [optional] 
**Storage** | Pointer to [**Storage**](Storage.md) |  | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 
**Viewed** | **bool** | Whether this recording has been viewed by you already. | 
**Viewers** | **[]string** | user ids of other users who have viewed this recording | 

## Methods

### NewSessionRecordingType

`func NewSessionRecordingType(endTime string, id string, recordingDuration float32, snapshotSource SnapshotSource, startTime string, viewed bool, viewers []string, ) *SessionRecordingType`

NewSessionRecordingType instantiates a new SessionRecordingType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRecordingTypeWithDefaults

`func NewSessionRecordingTypeWithDefaults() *SessionRecordingType`

NewSessionRecordingTypeWithDefaults instantiates a new SessionRecordingType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSeconds

`func (o *SessionRecordingType) GetActiveSeconds() float32`

GetActiveSeconds returns the ActiveSeconds field if non-nil, zero value otherwise.

### GetActiveSecondsOk

`func (o *SessionRecordingType) GetActiveSecondsOk() (*float32, bool)`

GetActiveSecondsOk returns a tuple with the ActiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSeconds

`func (o *SessionRecordingType) SetActiveSeconds(v float32)`

SetActiveSeconds sets ActiveSeconds field to given value.

### HasActiveSeconds

`func (o *SessionRecordingType) HasActiveSeconds() bool`

HasActiveSeconds returns a boolean if a field has been set.

### SetActiveSecondsNil

`func (o *SessionRecordingType) SetActiveSecondsNil(b bool)`

 SetActiveSecondsNil sets the value for ActiveSeconds to be an explicit nil

### UnsetActiveSeconds
`func (o *SessionRecordingType) UnsetActiveSeconds()`

UnsetActiveSeconds ensures that no value is present for ActiveSeconds, not even an explicit nil
### GetActivityScore

`func (o *SessionRecordingType) GetActivityScore() float32`

GetActivityScore returns the ActivityScore field if non-nil, zero value otherwise.

### GetActivityScoreOk

`func (o *SessionRecordingType) GetActivityScoreOk() (*float32, bool)`

GetActivityScoreOk returns a tuple with the ActivityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityScore

`func (o *SessionRecordingType) SetActivityScore(v float32)`

SetActivityScore sets ActivityScore field to given value.

### HasActivityScore

`func (o *SessionRecordingType) HasActivityScore() bool`

HasActivityScore returns a boolean if a field has been set.

### SetActivityScoreNil

`func (o *SessionRecordingType) SetActivityScoreNil(b bool)`

 SetActivityScoreNil sets the value for ActivityScore to be an explicit nil

### UnsetActivityScore
`func (o *SessionRecordingType) UnsetActivityScore()`

UnsetActivityScore ensures that no value is present for ActivityScore, not even an explicit nil
### GetClickCount

`func (o *SessionRecordingType) GetClickCount() float32`

GetClickCount returns the ClickCount field if non-nil, zero value otherwise.

### GetClickCountOk

`func (o *SessionRecordingType) GetClickCountOk() (*float32, bool)`

GetClickCountOk returns a tuple with the ClickCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickCount

`func (o *SessionRecordingType) SetClickCount(v float32)`

SetClickCount sets ClickCount field to given value.

### HasClickCount

`func (o *SessionRecordingType) HasClickCount() bool`

HasClickCount returns a boolean if a field has been set.

### SetClickCountNil

`func (o *SessionRecordingType) SetClickCountNil(b bool)`

 SetClickCountNil sets the value for ClickCount to be an explicit nil

### UnsetClickCount
`func (o *SessionRecordingType) UnsetClickCount()`

UnsetClickCount ensures that no value is present for ClickCount, not even an explicit nil
### GetConsoleErrorCount

`func (o *SessionRecordingType) GetConsoleErrorCount() float32`

GetConsoleErrorCount returns the ConsoleErrorCount field if non-nil, zero value otherwise.

### GetConsoleErrorCountOk

`func (o *SessionRecordingType) GetConsoleErrorCountOk() (*float32, bool)`

GetConsoleErrorCountOk returns a tuple with the ConsoleErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleErrorCount

`func (o *SessionRecordingType) SetConsoleErrorCount(v float32)`

SetConsoleErrorCount sets ConsoleErrorCount field to given value.

### HasConsoleErrorCount

`func (o *SessionRecordingType) HasConsoleErrorCount() bool`

HasConsoleErrorCount returns a boolean if a field has been set.

### SetConsoleErrorCountNil

`func (o *SessionRecordingType) SetConsoleErrorCountNil(b bool)`

 SetConsoleErrorCountNil sets the value for ConsoleErrorCount to be an explicit nil

### UnsetConsoleErrorCount
`func (o *SessionRecordingType) UnsetConsoleErrorCount()`

UnsetConsoleErrorCount ensures that no value is present for ConsoleErrorCount, not even an explicit nil
### GetConsoleLogCount

`func (o *SessionRecordingType) GetConsoleLogCount() float32`

GetConsoleLogCount returns the ConsoleLogCount field if non-nil, zero value otherwise.

### GetConsoleLogCountOk

`func (o *SessionRecordingType) GetConsoleLogCountOk() (*float32, bool)`

GetConsoleLogCountOk returns a tuple with the ConsoleLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleLogCount

`func (o *SessionRecordingType) SetConsoleLogCount(v float32)`

SetConsoleLogCount sets ConsoleLogCount field to given value.

### HasConsoleLogCount

`func (o *SessionRecordingType) HasConsoleLogCount() bool`

HasConsoleLogCount returns a boolean if a field has been set.

### SetConsoleLogCountNil

`func (o *SessionRecordingType) SetConsoleLogCountNil(b bool)`

 SetConsoleLogCountNil sets the value for ConsoleLogCount to be an explicit nil

### UnsetConsoleLogCount
`func (o *SessionRecordingType) UnsetConsoleLogCount()`

UnsetConsoleLogCount ensures that no value is present for ConsoleLogCount, not even an explicit nil
### GetConsoleWarnCount

`func (o *SessionRecordingType) GetConsoleWarnCount() float32`

GetConsoleWarnCount returns the ConsoleWarnCount field if non-nil, zero value otherwise.

### GetConsoleWarnCountOk

`func (o *SessionRecordingType) GetConsoleWarnCountOk() (*float32, bool)`

GetConsoleWarnCountOk returns a tuple with the ConsoleWarnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleWarnCount

`func (o *SessionRecordingType) SetConsoleWarnCount(v float32)`

SetConsoleWarnCount sets ConsoleWarnCount field to given value.

### HasConsoleWarnCount

`func (o *SessionRecordingType) HasConsoleWarnCount() bool`

HasConsoleWarnCount returns a boolean if a field has been set.

### SetConsoleWarnCountNil

`func (o *SessionRecordingType) SetConsoleWarnCountNil(b bool)`

 SetConsoleWarnCountNil sets the value for ConsoleWarnCount to be an explicit nil

### UnsetConsoleWarnCount
`func (o *SessionRecordingType) UnsetConsoleWarnCount()`

UnsetConsoleWarnCount ensures that no value is present for ConsoleWarnCount, not even an explicit nil
### GetDistinctId

`func (o *SessionRecordingType) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *SessionRecordingType) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *SessionRecordingType) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *SessionRecordingType) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### SetDistinctIdNil

`func (o *SessionRecordingType) SetDistinctIdNil(b bool)`

 SetDistinctIdNil sets the value for DistinctId to be an explicit nil

### UnsetDistinctId
`func (o *SessionRecordingType) UnsetDistinctId()`

UnsetDistinctId ensures that no value is present for DistinctId, not even an explicit nil
### GetEmail

`func (o *SessionRecordingType) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SessionRecordingType) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SessionRecordingType) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SessionRecordingType) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SessionRecordingType) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SessionRecordingType) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEndTime

`func (o *SessionRecordingType) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SessionRecordingType) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SessionRecordingType) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetExpiryTime

`func (o *SessionRecordingType) GetExpiryTime() string`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *SessionRecordingType) GetExpiryTimeOk() (*string, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *SessionRecordingType) SetExpiryTime(v string)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *SessionRecordingType) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### SetExpiryTimeNil

`func (o *SessionRecordingType) SetExpiryTimeNil(b bool)`

 SetExpiryTimeNil sets the value for ExpiryTime to be an explicit nil

### UnsetExpiryTime
`func (o *SessionRecordingType) UnsetExpiryTime()`

UnsetExpiryTime ensures that no value is present for ExpiryTime, not even an explicit nil
### GetId

`func (o *SessionRecordingType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionRecordingType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionRecordingType) SetId(v string)`

SetId sets Id field to given value.


### GetInactiveSeconds

`func (o *SessionRecordingType) GetInactiveSeconds() float32`

GetInactiveSeconds returns the InactiveSeconds field if non-nil, zero value otherwise.

### GetInactiveSecondsOk

`func (o *SessionRecordingType) GetInactiveSecondsOk() (*float32, bool)`

GetInactiveSecondsOk returns a tuple with the InactiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveSeconds

`func (o *SessionRecordingType) SetInactiveSeconds(v float32)`

SetInactiveSeconds sets InactiveSeconds field to given value.

### HasInactiveSeconds

`func (o *SessionRecordingType) HasInactiveSeconds() bool`

HasInactiveSeconds returns a boolean if a field has been set.

### SetInactiveSecondsNil

`func (o *SessionRecordingType) SetInactiveSecondsNil(b bool)`

 SetInactiveSecondsNil sets the value for InactiveSeconds to be an explicit nil

### UnsetInactiveSeconds
`func (o *SessionRecordingType) UnsetInactiveSeconds()`

UnsetInactiveSeconds ensures that no value is present for InactiveSeconds, not even an explicit nil
### GetKeypressCount

`func (o *SessionRecordingType) GetKeypressCount() float32`

GetKeypressCount returns the KeypressCount field if non-nil, zero value otherwise.

### GetKeypressCountOk

`func (o *SessionRecordingType) GetKeypressCountOk() (*float32, bool)`

GetKeypressCountOk returns a tuple with the KeypressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypressCount

`func (o *SessionRecordingType) SetKeypressCount(v float32)`

SetKeypressCount sets KeypressCount field to given value.

### HasKeypressCount

`func (o *SessionRecordingType) HasKeypressCount() bool`

HasKeypressCount returns a boolean if a field has been set.

### SetKeypressCountNil

`func (o *SessionRecordingType) SetKeypressCountNil(b bool)`

 SetKeypressCountNil sets the value for KeypressCount to be an explicit nil

### UnsetKeypressCount
`func (o *SessionRecordingType) UnsetKeypressCount()`

UnsetKeypressCount ensures that no value is present for KeypressCount, not even an explicit nil
### GetMatchingEvents

`func (o *SessionRecordingType) GetMatchingEvents() []MatchedRecording`

GetMatchingEvents returns the MatchingEvents field if non-nil, zero value otherwise.

### GetMatchingEventsOk

`func (o *SessionRecordingType) GetMatchingEventsOk() (*[]MatchedRecording, bool)`

GetMatchingEventsOk returns a tuple with the MatchingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingEvents

`func (o *SessionRecordingType) SetMatchingEvents(v []MatchedRecording)`

SetMatchingEvents sets MatchingEvents field to given value.

### HasMatchingEvents

`func (o *SessionRecordingType) HasMatchingEvents() bool`

HasMatchingEvents returns a boolean if a field has been set.

### SetMatchingEventsNil

`func (o *SessionRecordingType) SetMatchingEventsNil(b bool)`

 SetMatchingEventsNil sets the value for MatchingEvents to be an explicit nil

### UnsetMatchingEvents
`func (o *SessionRecordingType) UnsetMatchingEvents()`

UnsetMatchingEvents ensures that no value is present for MatchingEvents, not even an explicit nil
### GetMouseActivityCount

`func (o *SessionRecordingType) GetMouseActivityCount() float32`

GetMouseActivityCount returns the MouseActivityCount field if non-nil, zero value otherwise.

### GetMouseActivityCountOk

`func (o *SessionRecordingType) GetMouseActivityCountOk() (*float32, bool)`

GetMouseActivityCountOk returns a tuple with the MouseActivityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMouseActivityCount

`func (o *SessionRecordingType) SetMouseActivityCount(v float32)`

SetMouseActivityCount sets MouseActivityCount field to given value.

### HasMouseActivityCount

`func (o *SessionRecordingType) HasMouseActivityCount() bool`

HasMouseActivityCount returns a boolean if a field has been set.

### SetMouseActivityCountNil

`func (o *SessionRecordingType) SetMouseActivityCountNil(b bool)`

 SetMouseActivityCountNil sets the value for MouseActivityCount to be an explicit nil

### UnsetMouseActivityCount
`func (o *SessionRecordingType) UnsetMouseActivityCount()`

UnsetMouseActivityCount ensures that no value is present for MouseActivityCount, not even an explicit nil
### GetOngoing

`func (o *SessionRecordingType) GetOngoing() bool`

GetOngoing returns the Ongoing field if non-nil, zero value otherwise.

### GetOngoingOk

`func (o *SessionRecordingType) GetOngoingOk() (*bool, bool)`

GetOngoingOk returns a tuple with the Ongoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoing

`func (o *SessionRecordingType) SetOngoing(v bool)`

SetOngoing sets Ongoing field to given value.

### HasOngoing

`func (o *SessionRecordingType) HasOngoing() bool`

HasOngoing returns a boolean if a field has been set.

### SetOngoingNil

`func (o *SessionRecordingType) SetOngoingNil(b bool)`

 SetOngoingNil sets the value for Ongoing to be an explicit nil

### UnsetOngoing
`func (o *SessionRecordingType) UnsetOngoing()`

UnsetOngoing ensures that no value is present for Ongoing, not even an explicit nil
### GetPerson

`func (o *SessionRecordingType) GetPerson() PersonType`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *SessionRecordingType) GetPersonOk() (*PersonType, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *SessionRecordingType) SetPerson(v PersonType)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *SessionRecordingType) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetRecordingDuration

`func (o *SessionRecordingType) GetRecordingDuration() float32`

GetRecordingDuration returns the RecordingDuration field if non-nil, zero value otherwise.

### GetRecordingDurationOk

`func (o *SessionRecordingType) GetRecordingDurationOk() (*float32, bool)`

GetRecordingDurationOk returns a tuple with the RecordingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingDuration

`func (o *SessionRecordingType) SetRecordingDuration(v float32)`

SetRecordingDuration sets RecordingDuration field to given value.


### GetRecordingTtl

`func (o *SessionRecordingType) GetRecordingTtl() float32`

GetRecordingTtl returns the RecordingTtl field if non-nil, zero value otherwise.

### GetRecordingTtlOk

`func (o *SessionRecordingType) GetRecordingTtlOk() (*float32, bool)`

GetRecordingTtlOk returns a tuple with the RecordingTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingTtl

`func (o *SessionRecordingType) SetRecordingTtl(v float32)`

SetRecordingTtl sets RecordingTtl field to given value.

### HasRecordingTtl

`func (o *SessionRecordingType) HasRecordingTtl() bool`

HasRecordingTtl returns a boolean if a field has been set.

### SetRecordingTtlNil

`func (o *SessionRecordingType) SetRecordingTtlNil(b bool)`

 SetRecordingTtlNil sets the value for RecordingTtl to be an explicit nil

### UnsetRecordingTtl
`func (o *SessionRecordingType) UnsetRecordingTtl()`

UnsetRecordingTtl ensures that no value is present for RecordingTtl, not even an explicit nil
### GetRetentionPeriodDays

`func (o *SessionRecordingType) GetRetentionPeriodDays() float32`

GetRetentionPeriodDays returns the RetentionPeriodDays field if non-nil, zero value otherwise.

### GetRetentionPeriodDaysOk

`func (o *SessionRecordingType) GetRetentionPeriodDaysOk() (*float32, bool)`

GetRetentionPeriodDaysOk returns a tuple with the RetentionPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDays

`func (o *SessionRecordingType) SetRetentionPeriodDays(v float32)`

SetRetentionPeriodDays sets RetentionPeriodDays field to given value.

### HasRetentionPeriodDays

`func (o *SessionRecordingType) HasRetentionPeriodDays() bool`

HasRetentionPeriodDays returns a boolean if a field has been set.

### SetRetentionPeriodDaysNil

`func (o *SessionRecordingType) SetRetentionPeriodDaysNil(b bool)`

 SetRetentionPeriodDaysNil sets the value for RetentionPeriodDays to be an explicit nil

### UnsetRetentionPeriodDays
`func (o *SessionRecordingType) UnsetRetentionPeriodDays()`

UnsetRetentionPeriodDays ensures that no value is present for RetentionPeriodDays, not even an explicit nil
### GetSnapshotSource

`func (o *SessionRecordingType) GetSnapshotSource() SnapshotSource`

GetSnapshotSource returns the SnapshotSource field if non-nil, zero value otherwise.

### GetSnapshotSourceOk

`func (o *SessionRecordingType) GetSnapshotSourceOk() (*SnapshotSource, bool)`

GetSnapshotSourceOk returns a tuple with the SnapshotSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSource

`func (o *SessionRecordingType) SetSnapshotSource(v SnapshotSource)`

SetSnapshotSource sets SnapshotSource field to given value.


### GetStartTime

`func (o *SessionRecordingType) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SessionRecordingType) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SessionRecordingType) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetStartUrl

`func (o *SessionRecordingType) GetStartUrl() string`

GetStartUrl returns the StartUrl field if non-nil, zero value otherwise.

### GetStartUrlOk

`func (o *SessionRecordingType) GetStartUrlOk() (*string, bool)`

GetStartUrlOk returns a tuple with the StartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUrl

`func (o *SessionRecordingType) SetStartUrl(v string)`

SetStartUrl sets StartUrl field to given value.

### HasStartUrl

`func (o *SessionRecordingType) HasStartUrl() bool`

HasStartUrl returns a boolean if a field has been set.

### SetStartUrlNil

`func (o *SessionRecordingType) SetStartUrlNil(b bool)`

 SetStartUrlNil sets the value for StartUrl to be an explicit nil

### UnsetStartUrl
`func (o *SessionRecordingType) UnsetStartUrl()`

UnsetStartUrl ensures that no value is present for StartUrl, not even an explicit nil
### GetStorage

`func (o *SessionRecordingType) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *SessionRecordingType) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *SessionRecordingType) SetStorage(v Storage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *SessionRecordingType) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetSummary

`func (o *SessionRecordingType) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SessionRecordingType) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SessionRecordingType) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *SessionRecordingType) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *SessionRecordingType) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *SessionRecordingType) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetViewed

`func (o *SessionRecordingType) GetViewed() bool`

GetViewed returns the Viewed field if non-nil, zero value otherwise.

### GetViewedOk

`func (o *SessionRecordingType) GetViewedOk() (*bool, bool)`

GetViewedOk returns a tuple with the Viewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewed

`func (o *SessionRecordingType) SetViewed(v bool)`

SetViewed sets Viewed field to given value.


### GetViewers

`func (o *SessionRecordingType) GetViewers() []string`

GetViewers returns the Viewers field if non-nil, zero value otherwise.

### GetViewersOk

`func (o *SessionRecordingType) GetViewersOk() (*[]string, bool)`

GetViewersOk returns a tuple with the Viewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewers

`func (o *SessionRecordingType) SetViewers(v []string)`

SetViewers sets Viewers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


