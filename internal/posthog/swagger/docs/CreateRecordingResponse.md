# CreateRecordingResponse

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
**UploadToken** | **string** | Recall.ai upload token for the desktop SDK | 

## Methods

### NewCreateRecordingResponse

`func NewCreateRecordingResponse(id string, team int32, createdBy NullableInt32, sdkUploadId string, platform Platform9aaEnum, transcriptText string, createdAt time.Time, updatedAt time.Time, uploadToken string, ) *CreateRecordingResponse`

NewCreateRecordingResponse instantiates a new CreateRecordingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecordingResponseWithDefaults

`func NewCreateRecordingResponseWithDefaults() *CreateRecordingResponse`

NewCreateRecordingResponseWithDefaults instantiates a new CreateRecordingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateRecordingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateRecordingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateRecordingResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTeam

`func (o *CreateRecordingResponse) GetTeam() int32`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *CreateRecordingResponse) GetTeamOk() (*int32, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *CreateRecordingResponse) SetTeam(v int32)`

SetTeam sets Team field to given value.


### GetCreatedBy

`func (o *CreateRecordingResponse) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreateRecordingResponse) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreateRecordingResponse) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *CreateRecordingResponse) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *CreateRecordingResponse) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetSdkUploadId

`func (o *CreateRecordingResponse) GetSdkUploadId() string`

GetSdkUploadId returns the SdkUploadId field if non-nil, zero value otherwise.

### GetSdkUploadIdOk

`func (o *CreateRecordingResponse) GetSdkUploadIdOk() (*string, bool)`

GetSdkUploadIdOk returns a tuple with the SdkUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkUploadId

`func (o *CreateRecordingResponse) SetSdkUploadId(v string)`

SetSdkUploadId sets SdkUploadId field to given value.


### GetRecallRecordingId

`func (o *CreateRecordingResponse) GetRecallRecordingId() string`

GetRecallRecordingId returns the RecallRecordingId field if non-nil, zero value otherwise.

### GetRecallRecordingIdOk

`func (o *CreateRecordingResponse) GetRecallRecordingIdOk() (*string, bool)`

GetRecallRecordingIdOk returns a tuple with the RecallRecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecallRecordingId

`func (o *CreateRecordingResponse) SetRecallRecordingId(v string)`

SetRecallRecordingId sets RecallRecordingId field to given value.

### HasRecallRecordingId

`func (o *CreateRecordingResponse) HasRecallRecordingId() bool`

HasRecallRecordingId returns a boolean if a field has been set.

### SetRecallRecordingIdNil

`func (o *CreateRecordingResponse) SetRecallRecordingIdNil(b bool)`

 SetRecallRecordingIdNil sets the value for RecallRecordingId to be an explicit nil

### UnsetRecallRecordingId
`func (o *CreateRecordingResponse) UnsetRecallRecordingId()`

UnsetRecallRecordingId ensures that no value is present for RecallRecordingId, not even an explicit nil
### GetPlatform

`func (o *CreateRecordingResponse) GetPlatform() Platform9aaEnum`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateRecordingResponse) GetPlatformOk() (*Platform9aaEnum, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateRecordingResponse) SetPlatform(v Platform9aaEnum)`

SetPlatform sets Platform field to given value.


### GetMeetingTitle

`func (o *CreateRecordingResponse) GetMeetingTitle() string`

GetMeetingTitle returns the MeetingTitle field if non-nil, zero value otherwise.

### GetMeetingTitleOk

`func (o *CreateRecordingResponse) GetMeetingTitleOk() (*string, bool)`

GetMeetingTitleOk returns a tuple with the MeetingTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingTitle

`func (o *CreateRecordingResponse) SetMeetingTitle(v string)`

SetMeetingTitle sets MeetingTitle field to given value.

### HasMeetingTitle

`func (o *CreateRecordingResponse) HasMeetingTitle() bool`

HasMeetingTitle returns a boolean if a field has been set.

### SetMeetingTitleNil

`func (o *CreateRecordingResponse) SetMeetingTitleNil(b bool)`

 SetMeetingTitleNil sets the value for MeetingTitle to be an explicit nil

### UnsetMeetingTitle
`func (o *CreateRecordingResponse) UnsetMeetingTitle()`

UnsetMeetingTitle ensures that no value is present for MeetingTitle, not even an explicit nil
### GetMeetingUrl

`func (o *CreateRecordingResponse) GetMeetingUrl() string`

GetMeetingUrl returns the MeetingUrl field if non-nil, zero value otherwise.

### GetMeetingUrlOk

`func (o *CreateRecordingResponse) GetMeetingUrlOk() (*string, bool)`

GetMeetingUrlOk returns a tuple with the MeetingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingUrl

`func (o *CreateRecordingResponse) SetMeetingUrl(v string)`

SetMeetingUrl sets MeetingUrl field to given value.

### HasMeetingUrl

`func (o *CreateRecordingResponse) HasMeetingUrl() bool`

HasMeetingUrl returns a boolean if a field has been set.

### SetMeetingUrlNil

`func (o *CreateRecordingResponse) SetMeetingUrlNil(b bool)`

 SetMeetingUrlNil sets the value for MeetingUrl to be an explicit nil

### UnsetMeetingUrl
`func (o *CreateRecordingResponse) UnsetMeetingUrl()`

UnsetMeetingUrl ensures that no value is present for MeetingUrl, not even an explicit nil
### GetDurationSeconds

`func (o *CreateRecordingResponse) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *CreateRecordingResponse) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *CreateRecordingResponse) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *CreateRecordingResponse) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *CreateRecordingResponse) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *CreateRecordingResponse) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetStatus

`func (o *CreateRecordingResponse) GetStatus() Status292Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateRecordingResponse) GetStatusOk() (*Status292Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateRecordingResponse) SetStatus(v Status292Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateRecordingResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNotes

`func (o *CreateRecordingResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateRecordingResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateRecordingResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateRecordingResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *CreateRecordingResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *CreateRecordingResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetErrorMessage

`func (o *CreateRecordingResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CreateRecordingResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CreateRecordingResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CreateRecordingResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *CreateRecordingResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *CreateRecordingResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetVideoUrl

`func (o *CreateRecordingResponse) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *CreateRecordingResponse) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *CreateRecordingResponse) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *CreateRecordingResponse) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### SetVideoUrlNil

`func (o *CreateRecordingResponse) SetVideoUrlNil(b bool)`

 SetVideoUrlNil sets the value for VideoUrl to be an explicit nil

### UnsetVideoUrl
`func (o *CreateRecordingResponse) UnsetVideoUrl()`

UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
### GetVideoSizeBytes

`func (o *CreateRecordingResponse) GetVideoSizeBytes() int64`

GetVideoSizeBytes returns the VideoSizeBytes field if non-nil, zero value otherwise.

### GetVideoSizeBytesOk

`func (o *CreateRecordingResponse) GetVideoSizeBytesOk() (*int64, bool)`

GetVideoSizeBytesOk returns a tuple with the VideoSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoSizeBytes

`func (o *CreateRecordingResponse) SetVideoSizeBytes(v int64)`

SetVideoSizeBytes sets VideoSizeBytes field to given value.

### HasVideoSizeBytes

`func (o *CreateRecordingResponse) HasVideoSizeBytes() bool`

HasVideoSizeBytes returns a boolean if a field has been set.

### SetVideoSizeBytesNil

`func (o *CreateRecordingResponse) SetVideoSizeBytesNil(b bool)`

 SetVideoSizeBytesNil sets the value for VideoSizeBytes to be an explicit nil

### UnsetVideoSizeBytes
`func (o *CreateRecordingResponse) UnsetVideoSizeBytes()`

UnsetVideoSizeBytes ensures that no value is present for VideoSizeBytes, not even an explicit nil
### GetParticipants

`func (o *CreateRecordingResponse) GetParticipants() []string`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *CreateRecordingResponse) GetParticipantsOk() (*[]string, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *CreateRecordingResponse) SetParticipants(v []string)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *CreateRecordingResponse) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetTranscriptText

`func (o *CreateRecordingResponse) GetTranscriptText() string`

GetTranscriptText returns the TranscriptText field if non-nil, zero value otherwise.

### GetTranscriptTextOk

`func (o *CreateRecordingResponse) GetTranscriptTextOk() (*string, bool)`

GetTranscriptTextOk returns a tuple with the TranscriptText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptText

`func (o *CreateRecordingResponse) SetTranscriptText(v string)`

SetTranscriptText sets TranscriptText field to given value.


### GetTranscriptSegments

`func (o *CreateRecordingResponse) GetTranscriptSegments() []TranscriptSegment`

GetTranscriptSegments returns the TranscriptSegments field if non-nil, zero value otherwise.

### GetTranscriptSegmentsOk

`func (o *CreateRecordingResponse) GetTranscriptSegmentsOk() (*[]TranscriptSegment, bool)`

GetTranscriptSegmentsOk returns a tuple with the TranscriptSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptSegments

`func (o *CreateRecordingResponse) SetTranscriptSegments(v []TranscriptSegment)`

SetTranscriptSegments sets TranscriptSegments field to given value.

### HasTranscriptSegments

`func (o *CreateRecordingResponse) HasTranscriptSegments() bool`

HasTranscriptSegments returns a boolean if a field has been set.

### GetSummary

`func (o *CreateRecordingResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CreateRecordingResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CreateRecordingResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CreateRecordingResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *CreateRecordingResponse) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *CreateRecordingResponse) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetExtractedTasks

`func (o *CreateRecordingResponse) GetExtractedTasks() []Task`

GetExtractedTasks returns the ExtractedTasks field if non-nil, zero value otherwise.

### GetExtractedTasksOk

`func (o *CreateRecordingResponse) GetExtractedTasksOk() (*[]Task, bool)`

GetExtractedTasksOk returns a tuple with the ExtractedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedTasks

`func (o *CreateRecordingResponse) SetExtractedTasks(v []Task)`

SetExtractedTasks sets ExtractedTasks field to given value.

### HasExtractedTasks

`func (o *CreateRecordingResponse) HasExtractedTasks() bool`

HasExtractedTasks returns a boolean if a field has been set.

### GetTasksGeneratedAt

`func (o *CreateRecordingResponse) GetTasksGeneratedAt() time.Time`

GetTasksGeneratedAt returns the TasksGeneratedAt field if non-nil, zero value otherwise.

### GetTasksGeneratedAtOk

`func (o *CreateRecordingResponse) GetTasksGeneratedAtOk() (*time.Time, bool)`

GetTasksGeneratedAtOk returns a tuple with the TasksGeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksGeneratedAt

`func (o *CreateRecordingResponse) SetTasksGeneratedAt(v time.Time)`

SetTasksGeneratedAt sets TasksGeneratedAt field to given value.

### HasTasksGeneratedAt

`func (o *CreateRecordingResponse) HasTasksGeneratedAt() bool`

HasTasksGeneratedAt returns a boolean if a field has been set.

### SetTasksGeneratedAtNil

`func (o *CreateRecordingResponse) SetTasksGeneratedAtNil(b bool)`

 SetTasksGeneratedAtNil sets the value for TasksGeneratedAt to be an explicit nil

### UnsetTasksGeneratedAt
`func (o *CreateRecordingResponse) UnsetTasksGeneratedAt()`

UnsetTasksGeneratedAt ensures that no value is present for TasksGeneratedAt, not even an explicit nil
### GetSummaryGeneratedAt

`func (o *CreateRecordingResponse) GetSummaryGeneratedAt() time.Time`

GetSummaryGeneratedAt returns the SummaryGeneratedAt field if non-nil, zero value otherwise.

### GetSummaryGeneratedAtOk

`func (o *CreateRecordingResponse) GetSummaryGeneratedAtOk() (*time.Time, bool)`

GetSummaryGeneratedAtOk returns a tuple with the SummaryGeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryGeneratedAt

`func (o *CreateRecordingResponse) SetSummaryGeneratedAt(v time.Time)`

SetSummaryGeneratedAt sets SummaryGeneratedAt field to given value.

### HasSummaryGeneratedAt

`func (o *CreateRecordingResponse) HasSummaryGeneratedAt() bool`

HasSummaryGeneratedAt returns a boolean if a field has been set.

### SetSummaryGeneratedAtNil

`func (o *CreateRecordingResponse) SetSummaryGeneratedAtNil(b bool)`

 SetSummaryGeneratedAtNil sets the value for SummaryGeneratedAt to be an explicit nil

### UnsetSummaryGeneratedAt
`func (o *CreateRecordingResponse) UnsetSummaryGeneratedAt()`

UnsetSummaryGeneratedAt ensures that no value is present for SummaryGeneratedAt, not even an explicit nil
### GetStartedAt

`func (o *CreateRecordingResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CreateRecordingResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CreateRecordingResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CreateRecordingResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *CreateRecordingResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *CreateRecordingResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *CreateRecordingResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *CreateRecordingResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *CreateRecordingResponse) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *CreateRecordingResponse) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCreatedAt

`func (o *CreateRecordingResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateRecordingResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateRecordingResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateRecordingResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateRecordingResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateRecordingResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUploadToken

`func (o *CreateRecordingResponse) GetUploadToken() string`

GetUploadToken returns the UploadToken field if non-nil, zero value otherwise.

### GetUploadTokenOk

`func (o *CreateRecordingResponse) GetUploadTokenOk() (*string, bool)`

GetUploadTokenOk returns a tuple with the UploadToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadToken

`func (o *CreateRecordingResponse) SetUploadToken(v string)`

SetUploadToken sets UploadToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


