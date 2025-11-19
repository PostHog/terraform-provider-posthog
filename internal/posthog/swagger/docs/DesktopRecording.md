# DesktopRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Team** | **int32** |  | [readonly] 
**CreatedBy** | **NullableInt32** |  | [readonly] 
**SdkUploadId** | **string** |  | [readonly] 
**RecallRecordingId** | Pointer to **NullableString** |  | [optional] 
**Platform** | [**Platform9aaEnum**](Platform9aaEnum.md) |  | 
**MeetingTitle** | Pointer to **NullableString** |  | [optional] 
**MeetingUrl** | Pointer to **NullableString** |  | [optional] 
**DurationSeconds** | Pointer to **NullableInt32** |  | [optional] 
**Status** | Pointer to [**Status292Enum**](Status292Enum.md) |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**VideoUrl** | Pointer to **NullableString** |  | [optional] 
**VideoSizeBytes** | Pointer to **NullableInt64** |  | [optional] 
**Participants** | Pointer to **[]string** | List of participant names | [optional] 
**TranscriptText** | **string** |  | [readonly] 
**TranscriptSegments** | Pointer to [**[]TranscriptSegment**](TranscriptSegment.md) | Transcript segments with timestamps | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 
**ExtractedTasks** | Pointer to [**[]Task**](Task.md) | AI-extracted tasks from transcript | [optional] 
**TasksGeneratedAt** | Pointer to **NullableTime** |  | [optional] 
**SummaryGeneratedAt** | Pointer to **NullableTime** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewDesktopRecording

`func NewDesktopRecording(id string, team int32, createdBy NullableInt32, sdkUploadId string, platform Platform9aaEnum, transcriptText string, createdAt time.Time, updatedAt time.Time, ) *DesktopRecording`

NewDesktopRecording instantiates a new DesktopRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopRecordingWithDefaults

`func NewDesktopRecordingWithDefaults() *DesktopRecording`

NewDesktopRecordingWithDefaults instantiates a new DesktopRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DesktopRecording) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DesktopRecording) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DesktopRecording) SetId(v string)`

SetId sets Id field to given value.


### GetTeam

`func (o *DesktopRecording) GetTeam() int32`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *DesktopRecording) GetTeamOk() (*int32, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *DesktopRecording) SetTeam(v int32)`

SetTeam sets Team field to given value.


### GetCreatedBy

`func (o *DesktopRecording) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DesktopRecording) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DesktopRecording) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *DesktopRecording) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *DesktopRecording) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetSdkUploadId

`func (o *DesktopRecording) GetSdkUploadId() string`

GetSdkUploadId returns the SdkUploadId field if non-nil, zero value otherwise.

### GetSdkUploadIdOk

`func (o *DesktopRecording) GetSdkUploadIdOk() (*string, bool)`

GetSdkUploadIdOk returns a tuple with the SdkUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkUploadId

`func (o *DesktopRecording) SetSdkUploadId(v string)`

SetSdkUploadId sets SdkUploadId field to given value.


### GetRecallRecordingId

`func (o *DesktopRecording) GetRecallRecordingId() string`

GetRecallRecordingId returns the RecallRecordingId field if non-nil, zero value otherwise.

### GetRecallRecordingIdOk

`func (o *DesktopRecording) GetRecallRecordingIdOk() (*string, bool)`

GetRecallRecordingIdOk returns a tuple with the RecallRecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecallRecordingId

`func (o *DesktopRecording) SetRecallRecordingId(v string)`

SetRecallRecordingId sets RecallRecordingId field to given value.

### HasRecallRecordingId

`func (o *DesktopRecording) HasRecallRecordingId() bool`

HasRecallRecordingId returns a boolean if a field has been set.

### SetRecallRecordingIdNil

`func (o *DesktopRecording) SetRecallRecordingIdNil(b bool)`

 SetRecallRecordingIdNil sets the value for RecallRecordingId to be an explicit nil

### UnsetRecallRecordingId
`func (o *DesktopRecording) UnsetRecallRecordingId()`

UnsetRecallRecordingId ensures that no value is present for RecallRecordingId, not even an explicit nil
### GetPlatform

`func (o *DesktopRecording) GetPlatform() Platform9aaEnum`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DesktopRecording) GetPlatformOk() (*Platform9aaEnum, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DesktopRecording) SetPlatform(v Platform9aaEnum)`

SetPlatform sets Platform field to given value.


### GetMeetingTitle

`func (o *DesktopRecording) GetMeetingTitle() string`

GetMeetingTitle returns the MeetingTitle field if non-nil, zero value otherwise.

### GetMeetingTitleOk

`func (o *DesktopRecording) GetMeetingTitleOk() (*string, bool)`

GetMeetingTitleOk returns a tuple with the MeetingTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingTitle

`func (o *DesktopRecording) SetMeetingTitle(v string)`

SetMeetingTitle sets MeetingTitle field to given value.

### HasMeetingTitle

`func (o *DesktopRecording) HasMeetingTitle() bool`

HasMeetingTitle returns a boolean if a field has been set.

### SetMeetingTitleNil

`func (o *DesktopRecording) SetMeetingTitleNil(b bool)`

 SetMeetingTitleNil sets the value for MeetingTitle to be an explicit nil

### UnsetMeetingTitle
`func (o *DesktopRecording) UnsetMeetingTitle()`

UnsetMeetingTitle ensures that no value is present for MeetingTitle, not even an explicit nil
### GetMeetingUrl

`func (o *DesktopRecording) GetMeetingUrl() string`

GetMeetingUrl returns the MeetingUrl field if non-nil, zero value otherwise.

### GetMeetingUrlOk

`func (o *DesktopRecording) GetMeetingUrlOk() (*string, bool)`

GetMeetingUrlOk returns a tuple with the MeetingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingUrl

`func (o *DesktopRecording) SetMeetingUrl(v string)`

SetMeetingUrl sets MeetingUrl field to given value.

### HasMeetingUrl

`func (o *DesktopRecording) HasMeetingUrl() bool`

HasMeetingUrl returns a boolean if a field has been set.

### SetMeetingUrlNil

`func (o *DesktopRecording) SetMeetingUrlNil(b bool)`

 SetMeetingUrlNil sets the value for MeetingUrl to be an explicit nil

### UnsetMeetingUrl
`func (o *DesktopRecording) UnsetMeetingUrl()`

UnsetMeetingUrl ensures that no value is present for MeetingUrl, not even an explicit nil
### GetDurationSeconds

`func (o *DesktopRecording) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *DesktopRecording) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *DesktopRecording) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *DesktopRecording) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *DesktopRecording) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *DesktopRecording) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetStatus

`func (o *DesktopRecording) GetStatus() Status292Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DesktopRecording) GetStatusOk() (*Status292Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DesktopRecording) SetStatus(v Status292Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DesktopRecording) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNotes

`func (o *DesktopRecording) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DesktopRecording) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DesktopRecording) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DesktopRecording) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *DesktopRecording) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *DesktopRecording) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetErrorMessage

`func (o *DesktopRecording) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *DesktopRecording) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *DesktopRecording) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *DesktopRecording) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *DesktopRecording) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *DesktopRecording) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetVideoUrl

`func (o *DesktopRecording) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *DesktopRecording) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *DesktopRecording) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *DesktopRecording) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### SetVideoUrlNil

`func (o *DesktopRecording) SetVideoUrlNil(b bool)`

 SetVideoUrlNil sets the value for VideoUrl to be an explicit nil

### UnsetVideoUrl
`func (o *DesktopRecording) UnsetVideoUrl()`

UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
### GetVideoSizeBytes

`func (o *DesktopRecording) GetVideoSizeBytes() int64`

GetVideoSizeBytes returns the VideoSizeBytes field if non-nil, zero value otherwise.

### GetVideoSizeBytesOk

`func (o *DesktopRecording) GetVideoSizeBytesOk() (*int64, bool)`

GetVideoSizeBytesOk returns a tuple with the VideoSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoSizeBytes

`func (o *DesktopRecording) SetVideoSizeBytes(v int64)`

SetVideoSizeBytes sets VideoSizeBytes field to given value.

### HasVideoSizeBytes

`func (o *DesktopRecording) HasVideoSizeBytes() bool`

HasVideoSizeBytes returns a boolean if a field has been set.

### SetVideoSizeBytesNil

`func (o *DesktopRecording) SetVideoSizeBytesNil(b bool)`

 SetVideoSizeBytesNil sets the value for VideoSizeBytes to be an explicit nil

### UnsetVideoSizeBytes
`func (o *DesktopRecording) UnsetVideoSizeBytes()`

UnsetVideoSizeBytes ensures that no value is present for VideoSizeBytes, not even an explicit nil
### GetParticipants

`func (o *DesktopRecording) GetParticipants() []string`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *DesktopRecording) GetParticipantsOk() (*[]string, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *DesktopRecording) SetParticipants(v []string)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *DesktopRecording) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetTranscriptText

`func (o *DesktopRecording) GetTranscriptText() string`

GetTranscriptText returns the TranscriptText field if non-nil, zero value otherwise.

### GetTranscriptTextOk

`func (o *DesktopRecording) GetTranscriptTextOk() (*string, bool)`

GetTranscriptTextOk returns a tuple with the TranscriptText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptText

`func (o *DesktopRecording) SetTranscriptText(v string)`

SetTranscriptText sets TranscriptText field to given value.


### GetTranscriptSegments

`func (o *DesktopRecording) GetTranscriptSegments() []TranscriptSegment`

GetTranscriptSegments returns the TranscriptSegments field if non-nil, zero value otherwise.

### GetTranscriptSegmentsOk

`func (o *DesktopRecording) GetTranscriptSegmentsOk() (*[]TranscriptSegment, bool)`

GetTranscriptSegmentsOk returns a tuple with the TranscriptSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptSegments

`func (o *DesktopRecording) SetTranscriptSegments(v []TranscriptSegment)`

SetTranscriptSegments sets TranscriptSegments field to given value.

### HasTranscriptSegments

`func (o *DesktopRecording) HasTranscriptSegments() bool`

HasTranscriptSegments returns a boolean if a field has been set.

### GetSummary

`func (o *DesktopRecording) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DesktopRecording) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DesktopRecording) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *DesktopRecording) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *DesktopRecording) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *DesktopRecording) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetExtractedTasks

`func (o *DesktopRecording) GetExtractedTasks() []Task`

GetExtractedTasks returns the ExtractedTasks field if non-nil, zero value otherwise.

### GetExtractedTasksOk

`func (o *DesktopRecording) GetExtractedTasksOk() (*[]Task, bool)`

GetExtractedTasksOk returns a tuple with the ExtractedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedTasks

`func (o *DesktopRecording) SetExtractedTasks(v []Task)`

SetExtractedTasks sets ExtractedTasks field to given value.

### HasExtractedTasks

`func (o *DesktopRecording) HasExtractedTasks() bool`

HasExtractedTasks returns a boolean if a field has been set.

### GetTasksGeneratedAt

`func (o *DesktopRecording) GetTasksGeneratedAt() time.Time`

GetTasksGeneratedAt returns the TasksGeneratedAt field if non-nil, zero value otherwise.

### GetTasksGeneratedAtOk

`func (o *DesktopRecording) GetTasksGeneratedAtOk() (*time.Time, bool)`

GetTasksGeneratedAtOk returns a tuple with the TasksGeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksGeneratedAt

`func (o *DesktopRecording) SetTasksGeneratedAt(v time.Time)`

SetTasksGeneratedAt sets TasksGeneratedAt field to given value.

### HasTasksGeneratedAt

`func (o *DesktopRecording) HasTasksGeneratedAt() bool`

HasTasksGeneratedAt returns a boolean if a field has been set.

### SetTasksGeneratedAtNil

`func (o *DesktopRecording) SetTasksGeneratedAtNil(b bool)`

 SetTasksGeneratedAtNil sets the value for TasksGeneratedAt to be an explicit nil

### UnsetTasksGeneratedAt
`func (o *DesktopRecording) UnsetTasksGeneratedAt()`

UnsetTasksGeneratedAt ensures that no value is present for TasksGeneratedAt, not even an explicit nil
### GetSummaryGeneratedAt

`func (o *DesktopRecording) GetSummaryGeneratedAt() time.Time`

GetSummaryGeneratedAt returns the SummaryGeneratedAt field if non-nil, zero value otherwise.

### GetSummaryGeneratedAtOk

`func (o *DesktopRecording) GetSummaryGeneratedAtOk() (*time.Time, bool)`

GetSummaryGeneratedAtOk returns a tuple with the SummaryGeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryGeneratedAt

`func (o *DesktopRecording) SetSummaryGeneratedAt(v time.Time)`

SetSummaryGeneratedAt sets SummaryGeneratedAt field to given value.

### HasSummaryGeneratedAt

`func (o *DesktopRecording) HasSummaryGeneratedAt() bool`

HasSummaryGeneratedAt returns a boolean if a field has been set.

### SetSummaryGeneratedAtNil

`func (o *DesktopRecording) SetSummaryGeneratedAtNil(b bool)`

 SetSummaryGeneratedAtNil sets the value for SummaryGeneratedAt to be an explicit nil

### UnsetSummaryGeneratedAt
`func (o *DesktopRecording) UnsetSummaryGeneratedAt()`

UnsetSummaryGeneratedAt ensures that no value is present for SummaryGeneratedAt, not even an explicit nil
### GetStartedAt

`func (o *DesktopRecording) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DesktopRecording) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DesktopRecording) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DesktopRecording) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *DesktopRecording) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *DesktopRecording) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *DesktopRecording) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *DesktopRecording) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *DesktopRecording) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *DesktopRecording) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCreatedAt

`func (o *DesktopRecording) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DesktopRecording) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DesktopRecording) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DesktopRecording) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DesktopRecording) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DesktopRecording) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


