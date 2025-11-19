# ClickhouseQueryProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveCpuTime** | **int32** |  | 
**BytesRead** | **int32** |  | 
**EstimatedRowsTotal** | **int32** |  | 
**RowsRead** | **int32** |  | 
**TimeElapsed** | **int32** |  | 

## Methods

### NewClickhouseQueryProgress

`func NewClickhouseQueryProgress(activeCpuTime int32, bytesRead int32, estimatedRowsTotal int32, rowsRead int32, timeElapsed int32, ) *ClickhouseQueryProgress`

NewClickhouseQueryProgress instantiates a new ClickhouseQueryProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickhouseQueryProgressWithDefaults

`func NewClickhouseQueryProgressWithDefaults() *ClickhouseQueryProgress`

NewClickhouseQueryProgressWithDefaults instantiates a new ClickhouseQueryProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveCpuTime

`func (o *ClickhouseQueryProgress) GetActiveCpuTime() int32`

GetActiveCpuTime returns the ActiveCpuTime field if non-nil, zero value otherwise.

### GetActiveCpuTimeOk

`func (o *ClickhouseQueryProgress) GetActiveCpuTimeOk() (*int32, bool)`

GetActiveCpuTimeOk returns a tuple with the ActiveCpuTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCpuTime

`func (o *ClickhouseQueryProgress) SetActiveCpuTime(v int32)`

SetActiveCpuTime sets ActiveCpuTime field to given value.


### GetBytesRead

`func (o *ClickhouseQueryProgress) GetBytesRead() int32`

GetBytesRead returns the BytesRead field if non-nil, zero value otherwise.

### GetBytesReadOk

`func (o *ClickhouseQueryProgress) GetBytesReadOk() (*int32, bool)`

GetBytesReadOk returns a tuple with the BytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesRead

`func (o *ClickhouseQueryProgress) SetBytesRead(v int32)`

SetBytesRead sets BytesRead field to given value.


### GetEstimatedRowsTotal

`func (o *ClickhouseQueryProgress) GetEstimatedRowsTotal() int32`

GetEstimatedRowsTotal returns the EstimatedRowsTotal field if non-nil, zero value otherwise.

### GetEstimatedRowsTotalOk

`func (o *ClickhouseQueryProgress) GetEstimatedRowsTotalOk() (*int32, bool)`

GetEstimatedRowsTotalOk returns a tuple with the EstimatedRowsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRowsTotal

`func (o *ClickhouseQueryProgress) SetEstimatedRowsTotal(v int32)`

SetEstimatedRowsTotal sets EstimatedRowsTotal field to given value.


### GetRowsRead

`func (o *ClickhouseQueryProgress) GetRowsRead() int32`

GetRowsRead returns the RowsRead field if non-nil, zero value otherwise.

### GetRowsReadOk

`func (o *ClickhouseQueryProgress) GetRowsReadOk() (*int32, bool)`

GetRowsReadOk returns a tuple with the RowsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsRead

`func (o *ClickhouseQueryProgress) SetRowsRead(v int32)`

SetRowsRead sets RowsRead field to given value.


### GetTimeElapsed

`func (o *ClickhouseQueryProgress) GetTimeElapsed() int32`

GetTimeElapsed returns the TimeElapsed field if non-nil, zero value otherwise.

### GetTimeElapsedOk

`func (o *ClickhouseQueryProgress) GetTimeElapsedOk() (*int32, bool)`

GetTimeElapsedOk returns a tuple with the TimeElapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeElapsed

`func (o *ClickhouseQueryProgress) SetTimeElapsed(v int32)`

SetTimeElapsed sets TimeElapsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


