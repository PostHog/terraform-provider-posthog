# ExperimentExposureQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRange** | [**DateRange**](DateRange.md) |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentExposureQuery"]
**Timeseries** | [**[]ExperimentExposureTimeSeries**](ExperimentExposureTimeSeries.md) |  | 
**TotalExposures** | **map[string]float32** |  | 

## Methods

### NewExperimentExposureQueryResponse

`func NewExperimentExposureQueryResponse(dateRange DateRange, timeseries []ExperimentExposureTimeSeries, totalExposures map[string]float32, ) *ExperimentExposureQueryResponse`

NewExperimentExposureQueryResponse instantiates a new ExperimentExposureQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentExposureQueryResponseWithDefaults

`func NewExperimentExposureQueryResponseWithDefaults() *ExperimentExposureQueryResponse`

NewExperimentExposureQueryResponseWithDefaults instantiates a new ExperimentExposureQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRange

`func (o *ExperimentExposureQueryResponse) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *ExperimentExposureQueryResponse) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *ExperimentExposureQueryResponse) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetKind

`func (o *ExperimentExposureQueryResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentExposureQueryResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentExposureQueryResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentExposureQueryResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTimeseries

`func (o *ExperimentExposureQueryResponse) GetTimeseries() []ExperimentExposureTimeSeries`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *ExperimentExposureQueryResponse) GetTimeseriesOk() (*[]ExperimentExposureTimeSeries, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *ExperimentExposureQueryResponse) SetTimeseries(v []ExperimentExposureTimeSeries)`

SetTimeseries sets Timeseries field to given value.


### GetTotalExposures

`func (o *ExperimentExposureQueryResponse) GetTotalExposures() map[string]float32`

GetTotalExposures returns the TotalExposures field if non-nil, zero value otherwise.

### GetTotalExposuresOk

`func (o *ExperimentExposureQueryResponse) GetTotalExposuresOk() (*map[string]float32, bool)`

GetTotalExposuresOk returns a tuple with the TotalExposures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExposures

`func (o *ExperimentExposureQueryResponse) SetTotalExposures(v map[string]float32)`

SetTotalExposures sets TotalExposures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


