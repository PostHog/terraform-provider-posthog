# PatchedUserInterview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**IntervieweeEmails** | Pointer to **[]string** |  | [optional] 
**Transcript** | Pointer to **string** |  | [optional] [readonly] 
**Summary** | Pointer to **string** |  | [optional] 
**Audio** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedUserInterview

`func NewPatchedUserInterview() *PatchedUserInterview`

NewPatchedUserInterview instantiates a new PatchedUserInterview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUserInterviewWithDefaults

`func NewPatchedUserInterviewWithDefaults() *PatchedUserInterview`

NewPatchedUserInterviewWithDefaults instantiates a new PatchedUserInterview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedUserInterview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedUserInterview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedUserInterview) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedUserInterview) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedUserInterview) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedUserInterview) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedUserInterview) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedUserInterview) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedUserInterview) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedUserInterview) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedUserInterview) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedUserInterview) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIntervieweeEmails

`func (o *PatchedUserInterview) GetIntervieweeEmails() []string`

GetIntervieweeEmails returns the IntervieweeEmails field if non-nil, zero value otherwise.

### GetIntervieweeEmailsOk

`func (o *PatchedUserInterview) GetIntervieweeEmailsOk() (*[]string, bool)`

GetIntervieweeEmailsOk returns a tuple with the IntervieweeEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervieweeEmails

`func (o *PatchedUserInterview) SetIntervieweeEmails(v []string)`

SetIntervieweeEmails sets IntervieweeEmails field to given value.

### HasIntervieweeEmails

`func (o *PatchedUserInterview) HasIntervieweeEmails() bool`

HasIntervieweeEmails returns a boolean if a field has been set.

### GetTranscript

`func (o *PatchedUserInterview) GetTranscript() string`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *PatchedUserInterview) GetTranscriptOk() (*string, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *PatchedUserInterview) SetTranscript(v string)`

SetTranscript sets Transcript field to given value.

### HasTranscript

`func (o *PatchedUserInterview) HasTranscript() bool`

HasTranscript returns a boolean if a field has been set.

### GetSummary

`func (o *PatchedUserInterview) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PatchedUserInterview) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PatchedUserInterview) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PatchedUserInterview) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetAudio

`func (o *PatchedUserInterview) GetAudio() string`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *PatchedUserInterview) GetAudioOk() (*string, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *PatchedUserInterview) SetAudio(v string)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *PatchedUserInterview) HasAudio() bool`

HasAudio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


