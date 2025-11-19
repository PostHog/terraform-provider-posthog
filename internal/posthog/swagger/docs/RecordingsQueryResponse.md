# RecordingsQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | **bool** |  | 
**NextCursor** | Pointer to **NullableString** | Cursor for the next page. Contains the ordering value and session_id from the last record. | [optional] 
**Results** | [**[]SessionRecordingType**](SessionRecordingType.md) |  | 

## Methods

### NewRecordingsQueryResponse

`func NewRecordingsQueryResponse(hasNext bool, results []SessionRecordingType, ) *RecordingsQueryResponse`

NewRecordingsQueryResponse instantiates a new RecordingsQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingsQueryResponseWithDefaults

`func NewRecordingsQueryResponseWithDefaults() *RecordingsQueryResponse`

NewRecordingsQueryResponseWithDefaults instantiates a new RecordingsQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *RecordingsQueryResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *RecordingsQueryResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *RecordingsQueryResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetNextCursor

`func (o *RecordingsQueryResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *RecordingsQueryResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *RecordingsQueryResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *RecordingsQueryResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *RecordingsQueryResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *RecordingsQueryResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetResults

`func (o *RecordingsQueryResponse) GetResults() []SessionRecordingType`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RecordingsQueryResponse) GetResultsOk() (*[]SessionRecordingType, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RecordingsQueryResponse) SetResults(v []SessionRecordingType)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


