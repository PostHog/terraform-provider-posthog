# SessionSummaries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionIds** | **[]string** | List of session IDs to summarize (max 300) | 
**FocusArea** | Pointer to **string** | Optional focus area for the summarization | [optional] 

## Methods

### NewSessionSummaries

`func NewSessionSummaries(sessionIds []string, ) *SessionSummaries`

NewSessionSummaries instantiates a new SessionSummaries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionSummariesWithDefaults

`func NewSessionSummariesWithDefaults() *SessionSummaries`

NewSessionSummariesWithDefaults instantiates a new SessionSummaries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionIds

`func (o *SessionSummaries) GetSessionIds() []string`

GetSessionIds returns the SessionIds field if non-nil, zero value otherwise.

### GetSessionIdsOk

`func (o *SessionSummaries) GetSessionIdsOk() (*[]string, bool)`

GetSessionIdsOk returns a tuple with the SessionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIds

`func (o *SessionSummaries) SetSessionIds(v []string)`

SetSessionIds sets SessionIds field to given value.


### GetFocusArea

`func (o *SessionSummaries) GetFocusArea() string`

GetFocusArea returns the FocusArea field if non-nil, zero value otherwise.

### GetFocusAreaOk

`func (o *SessionSummaries) GetFocusAreaOk() (*string, bool)`

GetFocusAreaOk returns a tuple with the FocusArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocusArea

`func (o *SessionSummaries) SetFocusArea(v string)`

SetFocusArea sets FocusArea field to given value.

### HasFocusArea

`func (o *SessionSummaries) HasFocusArea() bool`

HasFocusArea returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


