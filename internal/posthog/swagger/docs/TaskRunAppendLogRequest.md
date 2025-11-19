# TaskRunAppendLogRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | **[]map[string]interface{}** | Array of log entry dictionaries to append | 

## Methods

### NewTaskRunAppendLogRequest

`func NewTaskRunAppendLogRequest(entries []map[string]interface{}, ) *TaskRunAppendLogRequest`

NewTaskRunAppendLogRequest instantiates a new TaskRunAppendLogRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskRunAppendLogRequestWithDefaults

`func NewTaskRunAppendLogRequestWithDefaults() *TaskRunAppendLogRequest`

NewTaskRunAppendLogRequestWithDefaults instantiates a new TaskRunAppendLogRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *TaskRunAppendLogRequest) GetEntries() []map[string]interface{}`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *TaskRunAppendLogRequest) GetEntriesOk() (*[]map[string]interface{}, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *TaskRunAppendLogRequest) SetEntries(v []map[string]interface{})`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


