# UserInterview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**IntervieweeEmails** | Pointer to **[]string** |  | [optional] 
**Transcript** | **string** |  | [readonly] 
**Summary** | Pointer to **string** |  | [optional] 
**Audio** | **string** |  | 

## Methods

### NewUserInterview

`func NewUserInterview(id string, createdBy UserBasic, createdAt time.Time, transcript string, audio string, ) *UserInterview`

NewUserInterview instantiates a new UserInterview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInterviewWithDefaults

`func NewUserInterviewWithDefaults() *UserInterview`

NewUserInterviewWithDefaults instantiates a new UserInterview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserInterview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInterview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInterview) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *UserInterview) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UserInterview) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UserInterview) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *UserInterview) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserInterview) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserInterview) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIntervieweeEmails

`func (o *UserInterview) GetIntervieweeEmails() []string`

GetIntervieweeEmails returns the IntervieweeEmails field if non-nil, zero value otherwise.

### GetIntervieweeEmailsOk

`func (o *UserInterview) GetIntervieweeEmailsOk() (*[]string, bool)`

GetIntervieweeEmailsOk returns a tuple with the IntervieweeEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervieweeEmails

`func (o *UserInterview) SetIntervieweeEmails(v []string)`

SetIntervieweeEmails sets IntervieweeEmails field to given value.

### HasIntervieweeEmails

`func (o *UserInterview) HasIntervieweeEmails() bool`

HasIntervieweeEmails returns a boolean if a field has been set.

### GetTranscript

`func (o *UserInterview) GetTranscript() string`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *UserInterview) GetTranscriptOk() (*string, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *UserInterview) SetTranscript(v string)`

SetTranscript sets Transcript field to given value.


### GetSummary

`func (o *UserInterview) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *UserInterview) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *UserInterview) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *UserInterview) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetAudio

`func (o *UserInterview) GetAudio() string`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *UserInterview) GetAudioOk() (*string, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *UserInterview) SetAudio(v string)`

SetAudio sets Audio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


