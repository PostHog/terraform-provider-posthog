# QueryResponseAlternative21

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRange** | [**DateRange**](DateRange.md) |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentExposureQuery"]
**Timeseries** | [**[]ExperimentExposureTimeSeries**](ExperimentExposureTimeSeries.md) |  | 
**TotalExposures** | **map[string]float32** |  | 

## Methods

### NewQueryResponseAlternative21

`func NewQueryResponseAlternative21(dateRange DateRange, timeseries []ExperimentExposureTimeSeries, totalExposures map[string]float32, ) *QueryResponseAlternative21`

NewQueryResponseAlternative21 instantiates a new QueryResponseAlternative21 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternative21WithDefaults

`func NewQueryResponseAlternative21WithDefaults() *QueryResponseAlternative21`

NewQueryResponseAlternative21WithDefaults instantiates a new QueryResponseAlternative21 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRange

`func (o *QueryResponseAlternative21) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *QueryResponseAlternative21) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *QueryResponseAlternative21) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetKind

`func (o *QueryResponseAlternative21) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QueryResponseAlternative21) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QueryResponseAlternative21) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *QueryResponseAlternative21) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTimeseries

`func (o *QueryResponseAlternative21) GetTimeseries() []ExperimentExposureTimeSeries`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *QueryResponseAlternative21) GetTimeseriesOk() (*[]ExperimentExposureTimeSeries, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *QueryResponseAlternative21) SetTimeseries(v []ExperimentExposureTimeSeries)`

SetTimeseries sets Timeseries field to given value.


### GetTotalExposures

`func (o *QueryResponseAlternative21) GetTotalExposures() map[string]float32`

GetTotalExposures returns the TotalExposures field if non-nil, zero value otherwise.

### GetTotalExposuresOk

`func (o *QueryResponseAlternative21) GetTotalExposuresOk() (*map[string]float32, bool)`

GetTotalExposuresOk returns a tuple with the TotalExposures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExposures

`func (o *QueryResponseAlternative21) SetTotalExposures(v map[string]float32)`

SetTotalExposures sets TotalExposures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


