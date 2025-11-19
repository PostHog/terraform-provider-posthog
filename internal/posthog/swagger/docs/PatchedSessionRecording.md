# PatchedSessionRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**DistinctId** | Pointer to **NullableString** |  | [optional] [readonly] 
**Viewed** | Pointer to **bool** |  | [optional] [readonly] 
**Viewers** | Pointer to **[]string** |  | [optional] [readonly] 
**RecordingDuration** | Pointer to **int32** |  | [optional] [readonly] 
**ActiveSeconds** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**InactiveSeconds** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**StartTime** | Pointer to **NullableTime** |  | [optional] [readonly] 
**EndTime** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ClickCount** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**KeypressCount** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**MouseActivityCount** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**ConsoleLogCount** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**ConsoleWarnCount** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**ConsoleErrorCount** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**StartUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**Person** | Pointer to [**MinimalPerson**](MinimalPerson.md) |  | [optional] 
**Storage** | Pointer to **string** |  | [optional] [readonly] 
**RetentionPeriodDays** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**ExpiryTime** | Pointer to **string** |  | [optional] [readonly] 
**RecordingTtl** | Pointer to **string** |  | [optional] [readonly] 
**SnapshotSource** | Pointer to **NullableString** |  | [optional] [readonly] 
**Ongoing** | Pointer to **bool** |  | [optional] [readonly] 
**ActivityScore** | Pointer to **NullableFloat64** |  | [optional] [readonly] 

## Methods

### NewPatchedSessionRecording

`func NewPatchedSessionRecording() *PatchedSessionRecording`

NewPatchedSessionRecording instantiates a new PatchedSessionRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSessionRecordingWithDefaults

`func NewPatchedSessionRecordingWithDefaults() *PatchedSessionRecording`

NewPatchedSessionRecordingWithDefaults instantiates a new PatchedSessionRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedSessionRecording) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedSessionRecording) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedSessionRecording) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedSessionRecording) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDistinctId

`func (o *PatchedSessionRecording) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *PatchedSessionRecording) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *PatchedSessionRecording) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *PatchedSessionRecording) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### SetDistinctIdNil

`func (o *PatchedSessionRecording) SetDistinctIdNil(b bool)`

 SetDistinctIdNil sets the value for DistinctId to be an explicit nil

### UnsetDistinctId
`func (o *PatchedSessionRecording) UnsetDistinctId()`

UnsetDistinctId ensures that no value is present for DistinctId, not even an explicit nil
### GetViewed

`func (o *PatchedSessionRecording) GetViewed() bool`

GetViewed returns the Viewed field if non-nil, zero value otherwise.

### GetViewedOk

`func (o *PatchedSessionRecording) GetViewedOk() (*bool, bool)`

GetViewedOk returns a tuple with the Viewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewed

`func (o *PatchedSessionRecording) SetViewed(v bool)`

SetViewed sets Viewed field to given value.

### HasViewed

`func (o *PatchedSessionRecording) HasViewed() bool`

HasViewed returns a boolean if a field has been set.

### GetViewers

`func (o *PatchedSessionRecording) GetViewers() []string`

GetViewers returns the Viewers field if non-nil, zero value otherwise.

### GetViewersOk

`func (o *PatchedSessionRecording) GetViewersOk() (*[]string, bool)`

GetViewersOk returns a tuple with the Viewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewers

`func (o *PatchedSessionRecording) SetViewers(v []string)`

SetViewers sets Viewers field to given value.

### HasViewers

`func (o *PatchedSessionRecording) HasViewers() bool`

HasViewers returns a boolean if a field has been set.

### GetRecordingDuration

`func (o *PatchedSessionRecording) GetRecordingDuration() int32`

GetRecordingDuration returns the RecordingDuration field if non-nil, zero value otherwise.

### GetRecordingDurationOk

`func (o *PatchedSessionRecording) GetRecordingDurationOk() (*int32, bool)`

GetRecordingDurationOk returns a tuple with the RecordingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingDuration

`func (o *PatchedSessionRecording) SetRecordingDuration(v int32)`

SetRecordingDuration sets RecordingDuration field to given value.

### HasRecordingDuration

`func (o *PatchedSessionRecording) HasRecordingDuration() bool`

HasRecordingDuration returns a boolean if a field has been set.

### GetActiveSeconds

`func (o *PatchedSessionRecording) GetActiveSeconds() int32`

GetActiveSeconds returns the ActiveSeconds field if non-nil, zero value otherwise.

### GetActiveSecondsOk

`func (o *PatchedSessionRecording) GetActiveSecondsOk() (*int32, bool)`

GetActiveSecondsOk returns a tuple with the ActiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSeconds

`func (o *PatchedSessionRecording) SetActiveSeconds(v int32)`

SetActiveSeconds sets ActiveSeconds field to given value.

### HasActiveSeconds

`func (o *PatchedSessionRecording) HasActiveSeconds() bool`

HasActiveSeconds returns a boolean if a field has been set.

### SetActiveSecondsNil

`func (o *PatchedSessionRecording) SetActiveSecondsNil(b bool)`

 SetActiveSecondsNil sets the value for ActiveSeconds to be an explicit nil

### UnsetActiveSeconds
`func (o *PatchedSessionRecording) UnsetActiveSeconds()`

UnsetActiveSeconds ensures that no value is present for ActiveSeconds, not even an explicit nil
### GetInactiveSeconds

`func (o *PatchedSessionRecording) GetInactiveSeconds() int32`

GetInactiveSeconds returns the InactiveSeconds field if non-nil, zero value otherwise.

### GetInactiveSecondsOk

`func (o *PatchedSessionRecording) GetInactiveSecondsOk() (*int32, bool)`

GetInactiveSecondsOk returns a tuple with the InactiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveSeconds

`func (o *PatchedSessionRecording) SetInactiveSeconds(v int32)`

SetInactiveSeconds sets InactiveSeconds field to given value.

### HasInactiveSeconds

`func (o *PatchedSessionRecording) HasInactiveSeconds() bool`

HasInactiveSeconds returns a boolean if a field has been set.

### SetInactiveSecondsNil

`func (o *PatchedSessionRecording) SetInactiveSecondsNil(b bool)`

 SetInactiveSecondsNil sets the value for InactiveSeconds to be an explicit nil

### UnsetInactiveSeconds
`func (o *PatchedSessionRecording) UnsetInactiveSeconds()`

UnsetInactiveSeconds ensures that no value is present for InactiveSeconds, not even an explicit nil
### GetStartTime

`func (o *PatchedSessionRecording) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PatchedSessionRecording) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PatchedSessionRecording) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *PatchedSessionRecording) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *PatchedSessionRecording) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *PatchedSessionRecording) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *PatchedSessionRecording) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *PatchedSessionRecording) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *PatchedSessionRecording) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *PatchedSessionRecording) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *PatchedSessionRecording) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *PatchedSessionRecording) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetClickCount

`func (o *PatchedSessionRecording) GetClickCount() int32`

GetClickCount returns the ClickCount field if non-nil, zero value otherwise.

### GetClickCountOk

`func (o *PatchedSessionRecording) GetClickCountOk() (*int32, bool)`

GetClickCountOk returns a tuple with the ClickCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickCount

`func (o *PatchedSessionRecording) SetClickCount(v int32)`

SetClickCount sets ClickCount field to given value.

### HasClickCount

`func (o *PatchedSessionRecording) HasClickCount() bool`

HasClickCount returns a boolean if a field has been set.

### SetClickCountNil

`func (o *PatchedSessionRecording) SetClickCountNil(b bool)`

 SetClickCountNil sets the value for ClickCount to be an explicit nil

### UnsetClickCount
`func (o *PatchedSessionRecording) UnsetClickCount()`

UnsetClickCount ensures that no value is present for ClickCount, not even an explicit nil
### GetKeypressCount

`func (o *PatchedSessionRecording) GetKeypressCount() int32`

GetKeypressCount returns the KeypressCount field if non-nil, zero value otherwise.

### GetKeypressCountOk

`func (o *PatchedSessionRecording) GetKeypressCountOk() (*int32, bool)`

GetKeypressCountOk returns a tuple with the KeypressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypressCount

`func (o *PatchedSessionRecording) SetKeypressCount(v int32)`

SetKeypressCount sets KeypressCount field to given value.

### HasKeypressCount

`func (o *PatchedSessionRecording) HasKeypressCount() bool`

HasKeypressCount returns a boolean if a field has been set.

### SetKeypressCountNil

`func (o *PatchedSessionRecording) SetKeypressCountNil(b bool)`

 SetKeypressCountNil sets the value for KeypressCount to be an explicit nil

### UnsetKeypressCount
`func (o *PatchedSessionRecording) UnsetKeypressCount()`

UnsetKeypressCount ensures that no value is present for KeypressCount, not even an explicit nil
### GetMouseActivityCount

`func (o *PatchedSessionRecording) GetMouseActivityCount() int32`

GetMouseActivityCount returns the MouseActivityCount field if non-nil, zero value otherwise.

### GetMouseActivityCountOk

`func (o *PatchedSessionRecording) GetMouseActivityCountOk() (*int32, bool)`

GetMouseActivityCountOk returns a tuple with the MouseActivityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMouseActivityCount

`func (o *PatchedSessionRecording) SetMouseActivityCount(v int32)`

SetMouseActivityCount sets MouseActivityCount field to given value.

### HasMouseActivityCount

`func (o *PatchedSessionRecording) HasMouseActivityCount() bool`

HasMouseActivityCount returns a boolean if a field has been set.

### SetMouseActivityCountNil

`func (o *PatchedSessionRecording) SetMouseActivityCountNil(b bool)`

 SetMouseActivityCountNil sets the value for MouseActivityCount to be an explicit nil

### UnsetMouseActivityCount
`func (o *PatchedSessionRecording) UnsetMouseActivityCount()`

UnsetMouseActivityCount ensures that no value is present for MouseActivityCount, not even an explicit nil
### GetConsoleLogCount

`func (o *PatchedSessionRecording) GetConsoleLogCount() int32`

GetConsoleLogCount returns the ConsoleLogCount field if non-nil, zero value otherwise.

### GetConsoleLogCountOk

`func (o *PatchedSessionRecording) GetConsoleLogCountOk() (*int32, bool)`

GetConsoleLogCountOk returns a tuple with the ConsoleLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleLogCount

`func (o *PatchedSessionRecording) SetConsoleLogCount(v int32)`

SetConsoleLogCount sets ConsoleLogCount field to given value.

### HasConsoleLogCount

`func (o *PatchedSessionRecording) HasConsoleLogCount() bool`

HasConsoleLogCount returns a boolean if a field has been set.

### SetConsoleLogCountNil

`func (o *PatchedSessionRecording) SetConsoleLogCountNil(b bool)`

 SetConsoleLogCountNil sets the value for ConsoleLogCount to be an explicit nil

### UnsetConsoleLogCount
`func (o *PatchedSessionRecording) UnsetConsoleLogCount()`

UnsetConsoleLogCount ensures that no value is present for ConsoleLogCount, not even an explicit nil
### GetConsoleWarnCount

`func (o *PatchedSessionRecording) GetConsoleWarnCount() int32`

GetConsoleWarnCount returns the ConsoleWarnCount field if non-nil, zero value otherwise.

### GetConsoleWarnCountOk

`func (o *PatchedSessionRecording) GetConsoleWarnCountOk() (*int32, bool)`

GetConsoleWarnCountOk returns a tuple with the ConsoleWarnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleWarnCount

`func (o *PatchedSessionRecording) SetConsoleWarnCount(v int32)`

SetConsoleWarnCount sets ConsoleWarnCount field to given value.

### HasConsoleWarnCount

`func (o *PatchedSessionRecording) HasConsoleWarnCount() bool`

HasConsoleWarnCount returns a boolean if a field has been set.

### SetConsoleWarnCountNil

`func (o *PatchedSessionRecording) SetConsoleWarnCountNil(b bool)`

 SetConsoleWarnCountNil sets the value for ConsoleWarnCount to be an explicit nil

### UnsetConsoleWarnCount
`func (o *PatchedSessionRecording) UnsetConsoleWarnCount()`

UnsetConsoleWarnCount ensures that no value is present for ConsoleWarnCount, not even an explicit nil
### GetConsoleErrorCount

`func (o *PatchedSessionRecording) GetConsoleErrorCount() int32`

GetConsoleErrorCount returns the ConsoleErrorCount field if non-nil, zero value otherwise.

### GetConsoleErrorCountOk

`func (o *PatchedSessionRecording) GetConsoleErrorCountOk() (*int32, bool)`

GetConsoleErrorCountOk returns a tuple with the ConsoleErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleErrorCount

`func (o *PatchedSessionRecording) SetConsoleErrorCount(v int32)`

SetConsoleErrorCount sets ConsoleErrorCount field to given value.

### HasConsoleErrorCount

`func (o *PatchedSessionRecording) HasConsoleErrorCount() bool`

HasConsoleErrorCount returns a boolean if a field has been set.

### SetConsoleErrorCountNil

`func (o *PatchedSessionRecording) SetConsoleErrorCountNil(b bool)`

 SetConsoleErrorCountNil sets the value for ConsoleErrorCount to be an explicit nil

### UnsetConsoleErrorCount
`func (o *PatchedSessionRecording) UnsetConsoleErrorCount()`

UnsetConsoleErrorCount ensures that no value is present for ConsoleErrorCount, not even an explicit nil
### GetStartUrl

`func (o *PatchedSessionRecording) GetStartUrl() string`

GetStartUrl returns the StartUrl field if non-nil, zero value otherwise.

### GetStartUrlOk

`func (o *PatchedSessionRecording) GetStartUrlOk() (*string, bool)`

GetStartUrlOk returns a tuple with the StartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUrl

`func (o *PatchedSessionRecording) SetStartUrl(v string)`

SetStartUrl sets StartUrl field to given value.

### HasStartUrl

`func (o *PatchedSessionRecording) HasStartUrl() bool`

HasStartUrl returns a boolean if a field has been set.

### SetStartUrlNil

`func (o *PatchedSessionRecording) SetStartUrlNil(b bool)`

 SetStartUrlNil sets the value for StartUrl to be an explicit nil

### UnsetStartUrl
`func (o *PatchedSessionRecording) UnsetStartUrl()`

UnsetStartUrl ensures that no value is present for StartUrl, not even an explicit nil
### GetPerson

`func (o *PatchedSessionRecording) GetPerson() MinimalPerson`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *PatchedSessionRecording) GetPersonOk() (*MinimalPerson, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *PatchedSessionRecording) SetPerson(v MinimalPerson)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *PatchedSessionRecording) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetStorage

`func (o *PatchedSessionRecording) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *PatchedSessionRecording) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *PatchedSessionRecording) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *PatchedSessionRecording) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetRetentionPeriodDays

`func (o *PatchedSessionRecording) GetRetentionPeriodDays() int32`

GetRetentionPeriodDays returns the RetentionPeriodDays field if non-nil, zero value otherwise.

### GetRetentionPeriodDaysOk

`func (o *PatchedSessionRecording) GetRetentionPeriodDaysOk() (*int32, bool)`

GetRetentionPeriodDaysOk returns a tuple with the RetentionPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDays

`func (o *PatchedSessionRecording) SetRetentionPeriodDays(v int32)`

SetRetentionPeriodDays sets RetentionPeriodDays field to given value.

### HasRetentionPeriodDays

`func (o *PatchedSessionRecording) HasRetentionPeriodDays() bool`

HasRetentionPeriodDays returns a boolean if a field has been set.

### SetRetentionPeriodDaysNil

`func (o *PatchedSessionRecording) SetRetentionPeriodDaysNil(b bool)`

 SetRetentionPeriodDaysNil sets the value for RetentionPeriodDays to be an explicit nil

### UnsetRetentionPeriodDays
`func (o *PatchedSessionRecording) UnsetRetentionPeriodDays()`

UnsetRetentionPeriodDays ensures that no value is present for RetentionPeriodDays, not even an explicit nil
### GetExpiryTime

`func (o *PatchedSessionRecording) GetExpiryTime() string`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *PatchedSessionRecording) GetExpiryTimeOk() (*string, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *PatchedSessionRecording) SetExpiryTime(v string)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *PatchedSessionRecording) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetRecordingTtl

`func (o *PatchedSessionRecording) GetRecordingTtl() string`

GetRecordingTtl returns the RecordingTtl field if non-nil, zero value otherwise.

### GetRecordingTtlOk

`func (o *PatchedSessionRecording) GetRecordingTtlOk() (*string, bool)`

GetRecordingTtlOk returns a tuple with the RecordingTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingTtl

`func (o *PatchedSessionRecording) SetRecordingTtl(v string)`

SetRecordingTtl sets RecordingTtl field to given value.

### HasRecordingTtl

`func (o *PatchedSessionRecording) HasRecordingTtl() bool`

HasRecordingTtl returns a boolean if a field has been set.

### GetSnapshotSource

`func (o *PatchedSessionRecording) GetSnapshotSource() string`

GetSnapshotSource returns the SnapshotSource field if non-nil, zero value otherwise.

### GetSnapshotSourceOk

`func (o *PatchedSessionRecording) GetSnapshotSourceOk() (*string, bool)`

GetSnapshotSourceOk returns a tuple with the SnapshotSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSource

`func (o *PatchedSessionRecording) SetSnapshotSource(v string)`

SetSnapshotSource sets SnapshotSource field to given value.

### HasSnapshotSource

`func (o *PatchedSessionRecording) HasSnapshotSource() bool`

HasSnapshotSource returns a boolean if a field has been set.

### SetSnapshotSourceNil

`func (o *PatchedSessionRecording) SetSnapshotSourceNil(b bool)`

 SetSnapshotSourceNil sets the value for SnapshotSource to be an explicit nil

### UnsetSnapshotSource
`func (o *PatchedSessionRecording) UnsetSnapshotSource()`

UnsetSnapshotSource ensures that no value is present for SnapshotSource, not even an explicit nil
### GetOngoing

`func (o *PatchedSessionRecording) GetOngoing() bool`

GetOngoing returns the Ongoing field if non-nil, zero value otherwise.

### GetOngoingOk

`func (o *PatchedSessionRecording) GetOngoingOk() (*bool, bool)`

GetOngoingOk returns a tuple with the Ongoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoing

`func (o *PatchedSessionRecording) SetOngoing(v bool)`

SetOngoing sets Ongoing field to given value.

### HasOngoing

`func (o *PatchedSessionRecording) HasOngoing() bool`

HasOngoing returns a boolean if a field has been set.

### GetActivityScore

`func (o *PatchedSessionRecording) GetActivityScore() float64`

GetActivityScore returns the ActivityScore field if non-nil, zero value otherwise.

### GetActivityScoreOk

`func (o *PatchedSessionRecording) GetActivityScoreOk() (*float64, bool)`

GetActivityScoreOk returns a tuple with the ActivityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityScore

`func (o *PatchedSessionRecording) SetActivityScore(v float64)`

SetActivityScore sets ActivityScore field to given value.

### HasActivityScore

`func (o *PatchedSessionRecording) HasActivityScore() bool`

HasActivityScore returns a boolean if a field has been set.

### SetActivityScoreNil

`func (o *PatchedSessionRecording) SetActivityScoreNil(b bool)`

 SetActivityScoreNil sets the value for ActivityScore to be an explicit nil

### UnsetActivityScore
`func (o *PatchedSessionRecording) UnsetActivityScore()`

UnsetActivityScore ensures that no value is present for ActivityScore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


