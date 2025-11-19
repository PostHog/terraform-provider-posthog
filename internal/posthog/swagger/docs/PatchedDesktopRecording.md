# PatchedDesktopRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Team** | Pointer to **int32** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**SdkUploadId** | Pointer to **string** |  | [optional] [readonly] 
**RecallRecordingId** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to [**Platform9aaEnum**](Platform9aaEnum.md) |  | [optional] 
**MeetingTitle** | Pointer to **NullableString** |  | [optional] 
**MeetingUrl** | Pointer to **NullableString** |  | [optional] 
**DurationSeconds** | Pointer to **NullableInt32** |  | [optional] 
**Status** | Pointer to [**Status292Enum**](Status292Enum.md) |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**VideoUrl** | Pointer to **NullableString** |  | [optional] 
**VideoSizeBytes** | Pointer to **NullableInt64** |  | [optional] 
**Participants** | Pointer to **[]string** | List of participant names | [optional] 
**TranscriptText** | Pointer to **string** |  | [optional] [readonly] 
**TranscriptSegments** | Pointer to [**[]TranscriptSegment**](TranscriptSegment.md) | Transcript segments with timestamps | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 
**ExtractedTasks** | Pointer to [**[]Task**](Task.md) | AI-extracted tasks from transcript | [optional] 
**TasksGeneratedAt** | Pointer to **NullableTime** |  | [optional] 
**SummaryGeneratedAt** | Pointer to **NullableTime** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedDesktopRecording

`func NewPatchedDesktopRecording() *PatchedDesktopRecording`

NewPatchedDesktopRecording instantiates a new PatchedDesktopRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDesktopRecordingWithDefaults

`func NewPatchedDesktopRecordingWithDefaults() *PatchedDesktopRecording`

NewPatchedDesktopRecordingWithDefaults instantiates a new PatchedDesktopRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDesktopRecording) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDesktopRecording) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDesktopRecording) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDesktopRecording) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeam

`func (o *PatchedDesktopRecording) GetTeam() int32`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *PatchedDesktopRecording) GetTeamOk() (*int32, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *PatchedDesktopRecording) SetTeam(v int32)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *PatchedDesktopRecording) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedDesktopRecording) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedDesktopRecording) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedDesktopRecording) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedDesktopRecording) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *PatchedDesktopRecording) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *PatchedDesktopRecording) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetSdkUploadId

`func (o *PatchedDesktopRecording) GetSdkUploadId() string`

GetSdkUploadId returns the SdkUploadId field if non-nil, zero value otherwise.

### GetSdkUploadIdOk

`func (o *PatchedDesktopRecording) GetSdkUploadIdOk() (*string, bool)`

GetSdkUploadIdOk returns a tuple with the SdkUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkUploadId

`func (o *PatchedDesktopRecording) SetSdkUploadId(v string)`

SetSdkUploadId sets SdkUploadId field to given value.

### HasSdkUploadId

`func (o *PatchedDesktopRecording) HasSdkUploadId() bool`

HasSdkUploadId returns a boolean if a field has been set.

### GetRecallRecordingId

`func (o *PatchedDesktopRecording) GetRecallRecordingId() string`

GetRecallRecordingId returns the RecallRecordingId field if non-nil, zero value otherwise.

### GetRecallRecordingIdOk

`func (o *PatchedDesktopRecording) GetRecallRecordingIdOk() (*string, bool)`

GetRecallRecordingIdOk returns a tuple with the RecallRecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecallRecordingId

`func (o *PatchedDesktopRecording) SetRecallRecordingId(v string)`

SetRecallRecordingId sets RecallRecordingId field to given value.

### HasRecallRecordingId

`func (o *PatchedDesktopRecording) HasRecallRecordingId() bool`

HasRecallRecordingId returns a boolean if a field has been set.

### SetRecallRecordingIdNil

`func (o *PatchedDesktopRecording) SetRecallRecordingIdNil(b bool)`

 SetRecallRecordingIdNil sets the value for RecallRecordingId to be an explicit nil

### UnsetRecallRecordingId
`func (o *PatchedDesktopRecording) UnsetRecallRecordingId()`

UnsetRecallRecordingId ensures that no value is present for RecallRecordingId, not even an explicit nil
### GetPlatform

`func (o *PatchedDesktopRecording) GetPlatform() Platform9aaEnum`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedDesktopRecording) GetPlatformOk() (*Platform9aaEnum, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedDesktopRecording) SetPlatform(v Platform9aaEnum)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedDesktopRecording) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetMeetingTitle

`func (o *PatchedDesktopRecording) GetMeetingTitle() string`

GetMeetingTitle returns the MeetingTitle field if non-nil, zero value otherwise.

### GetMeetingTitleOk

`func (o *PatchedDesktopRecording) GetMeetingTitleOk() (*string, bool)`

GetMeetingTitleOk returns a tuple with the MeetingTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingTitle

`func (o *PatchedDesktopRecording) SetMeetingTitle(v string)`

SetMeetingTitle sets MeetingTitle field to given value.

### HasMeetingTitle

`func (o *PatchedDesktopRecording) HasMeetingTitle() bool`

HasMeetingTitle returns a boolean if a field has been set.

### SetMeetingTitleNil

`func (o *PatchedDesktopRecording) SetMeetingTitleNil(b bool)`

 SetMeetingTitleNil sets the value for MeetingTitle to be an explicit nil

### UnsetMeetingTitle
`func (o *PatchedDesktopRecording) UnsetMeetingTitle()`

UnsetMeetingTitle ensures that no value is present for MeetingTitle, not even an explicit nil
### GetMeetingUrl

`func (o *PatchedDesktopRecording) GetMeetingUrl() string`

GetMeetingUrl returns the MeetingUrl field if non-nil, zero value otherwise.

### GetMeetingUrlOk

`func (o *PatchedDesktopRecording) GetMeetingUrlOk() (*string, bool)`

GetMeetingUrlOk returns a tuple with the MeetingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingUrl

`func (o *PatchedDesktopRecording) SetMeetingUrl(v string)`

SetMeetingUrl sets MeetingUrl field to given value.

### HasMeetingUrl

`func (o *PatchedDesktopRecording) HasMeetingUrl() bool`

HasMeetingUrl returns a boolean if a field has been set.

### SetMeetingUrlNil

`func (o *PatchedDesktopRecording) SetMeetingUrlNil(b bool)`

 SetMeetingUrlNil sets the value for MeetingUrl to be an explicit nil

### UnsetMeetingUrl
`func (o *PatchedDesktopRecording) UnsetMeetingUrl()`

UnsetMeetingUrl ensures that no value is present for MeetingUrl, not even an explicit nil
### GetDurationSeconds

`func (o *PatchedDesktopRecording) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *PatchedDesktopRecording) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *PatchedDesktopRecording) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *PatchedDesktopRecording) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *PatchedDesktopRecording) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *PatchedDesktopRecording) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetStatus

`func (o *PatchedDesktopRecording) GetStatus() Status292Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedDesktopRecording) GetStatusOk() (*Status292Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedDesktopRecording) SetStatus(v Status292Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedDesktopRecording) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNotes

`func (o *PatchedDesktopRecording) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PatchedDesktopRecording) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PatchedDesktopRecording) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PatchedDesktopRecording) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *PatchedDesktopRecording) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *PatchedDesktopRecording) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetErrorMessage

`func (o *PatchedDesktopRecording) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PatchedDesktopRecording) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PatchedDesktopRecording) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PatchedDesktopRecording) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *PatchedDesktopRecording) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *PatchedDesktopRecording) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetVideoUrl

`func (o *PatchedDesktopRecording) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *PatchedDesktopRecording) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *PatchedDesktopRecording) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *PatchedDesktopRecording) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### SetVideoUrlNil

`func (o *PatchedDesktopRecording) SetVideoUrlNil(b bool)`

 SetVideoUrlNil sets the value for VideoUrl to be an explicit nil

### UnsetVideoUrl
`func (o *PatchedDesktopRecording) UnsetVideoUrl()`

UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
### GetVideoSizeBytes

`func (o *PatchedDesktopRecording) GetVideoSizeBytes() int64`

GetVideoSizeBytes returns the VideoSizeBytes field if non-nil, zero value otherwise.

### GetVideoSizeBytesOk

`func (o *PatchedDesktopRecording) GetVideoSizeBytesOk() (*int64, bool)`

GetVideoSizeBytesOk returns a tuple with the VideoSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoSizeBytes

`func (o *PatchedDesktopRecording) SetVideoSizeBytes(v int64)`

SetVideoSizeBytes sets VideoSizeBytes field to given value.

### HasVideoSizeBytes

`func (o *PatchedDesktopRecording) HasVideoSizeBytes() bool`

HasVideoSizeBytes returns a boolean if a field has been set.

### SetVideoSizeBytesNil

`func (o *PatchedDesktopRecording) SetVideoSizeBytesNil(b bool)`

 SetVideoSizeBytesNil sets the value for VideoSizeBytes to be an explicit nil

### UnsetVideoSizeBytes
`func (o *PatchedDesktopRecording) UnsetVideoSizeBytes()`

UnsetVideoSizeBytes ensures that no value is present for VideoSizeBytes, not even an explicit nil
### GetParticipants

`func (o *PatchedDesktopRecording) GetParticipants() []string`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *PatchedDesktopRecording) GetParticipantsOk() (*[]string, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *PatchedDesktopRecording) SetParticipants(v []string)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *PatchedDesktopRecording) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetTranscriptText

`func (o *PatchedDesktopRecording) GetTranscriptText() string`

GetTranscriptText returns the TranscriptText field if non-nil, zero value otherwise.

### GetTranscriptTextOk

`func (o *PatchedDesktopRecording) GetTranscriptTextOk() (*string, bool)`

GetTranscriptTextOk returns a tuple with the TranscriptText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptText

`func (o *PatchedDesktopRecording) SetTranscriptText(v string)`

SetTranscriptText sets TranscriptText field to given value.

### HasTranscriptText

`func (o *PatchedDesktopRecording) HasTranscriptText() bool`

HasTranscriptText returns a boolean if a field has been set.

### GetTranscriptSegments

`func (o *PatchedDesktopRecording) GetTranscriptSegments() []TranscriptSegment`

GetTranscriptSegments returns the TranscriptSegments field if non-nil, zero value otherwise.

### GetTranscriptSegmentsOk

`func (o *PatchedDesktopRecording) GetTranscriptSegmentsOk() (*[]TranscriptSegment, bool)`

GetTranscriptSegmentsOk returns a tuple with the TranscriptSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptSegments

`func (o *PatchedDesktopRecording) SetTranscriptSegments(v []TranscriptSegment)`

SetTranscriptSegments sets TranscriptSegments field to given value.

### HasTranscriptSegments

`func (o *PatchedDesktopRecording) HasTranscriptSegments() bool`

HasTranscriptSegments returns a boolean if a field has been set.

### GetSummary

`func (o *PatchedDesktopRecording) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PatchedDesktopRecording) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PatchedDesktopRecording) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PatchedDesktopRecording) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *PatchedDesktopRecording) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *PatchedDesktopRecording) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetExtractedTasks

`func (o *PatchedDesktopRecording) GetExtractedTasks() []Task`

GetExtractedTasks returns the ExtractedTasks field if non-nil, zero value otherwise.

### GetExtractedTasksOk

`func (o *PatchedDesktopRecording) GetExtractedTasksOk() (*[]Task, bool)`

GetExtractedTasksOk returns a tuple with the ExtractedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedTasks

`func (o *PatchedDesktopRecording) SetExtractedTasks(v []Task)`

SetExtractedTasks sets ExtractedTasks field to given value.

### HasExtractedTasks

`func (o *PatchedDesktopRecording) HasExtractedTasks() bool`

HasExtractedTasks returns a boolean if a field has been set.

### GetTasksGeneratedAt

`func (o *PatchedDesktopRecording) GetTasksGeneratedAt() time.Time`

GetTasksGeneratedAt returns the TasksGeneratedAt field if non-nil, zero value otherwise.

### GetTasksGeneratedAtOk

`func (o *PatchedDesktopRecording) GetTasksGeneratedAtOk() (*time.Time, bool)`

GetTasksGeneratedAtOk returns a tuple with the TasksGeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksGeneratedAt

`func (o *PatchedDesktopRecording) SetTasksGeneratedAt(v time.Time)`

SetTasksGeneratedAt sets TasksGeneratedAt field to given value.

### HasTasksGeneratedAt

`func (o *PatchedDesktopRecording) HasTasksGeneratedAt() bool`

HasTasksGeneratedAt returns a boolean if a field has been set.

### SetTasksGeneratedAtNil

`func (o *PatchedDesktopRecording) SetTasksGeneratedAtNil(b bool)`

 SetTasksGeneratedAtNil sets the value for TasksGeneratedAt to be an explicit nil

### UnsetTasksGeneratedAt
`func (o *PatchedDesktopRecording) UnsetTasksGeneratedAt()`

UnsetTasksGeneratedAt ensures that no value is present for TasksGeneratedAt, not even an explicit nil
### GetSummaryGeneratedAt

`func (o *PatchedDesktopRecording) GetSummaryGeneratedAt() time.Time`

GetSummaryGeneratedAt returns the SummaryGeneratedAt field if non-nil, zero value otherwise.

### GetSummaryGeneratedAtOk

`func (o *PatchedDesktopRecording) GetSummaryGeneratedAtOk() (*time.Time, bool)`

GetSummaryGeneratedAtOk returns a tuple with the SummaryGeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryGeneratedAt

`func (o *PatchedDesktopRecording) SetSummaryGeneratedAt(v time.Time)`

SetSummaryGeneratedAt sets SummaryGeneratedAt field to given value.

### HasSummaryGeneratedAt

`func (o *PatchedDesktopRecording) HasSummaryGeneratedAt() bool`

HasSummaryGeneratedAt returns a boolean if a field has been set.

### SetSummaryGeneratedAtNil

`func (o *PatchedDesktopRecording) SetSummaryGeneratedAtNil(b bool)`

 SetSummaryGeneratedAtNil sets the value for SummaryGeneratedAt to be an explicit nil

### UnsetSummaryGeneratedAt
`func (o *PatchedDesktopRecording) UnsetSummaryGeneratedAt()`

UnsetSummaryGeneratedAt ensures that no value is present for SummaryGeneratedAt, not even an explicit nil
### GetStartedAt

`func (o *PatchedDesktopRecording) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *PatchedDesktopRecording) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *PatchedDesktopRecording) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *PatchedDesktopRecording) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *PatchedDesktopRecording) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *PatchedDesktopRecording) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *PatchedDesktopRecording) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *PatchedDesktopRecording) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *PatchedDesktopRecording) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *PatchedDesktopRecording) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCreatedAt

`func (o *PatchedDesktopRecording) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedDesktopRecording) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedDesktopRecording) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedDesktopRecording) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedDesktopRecording) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedDesktopRecording) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedDesktopRecording) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedDesktopRecording) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


