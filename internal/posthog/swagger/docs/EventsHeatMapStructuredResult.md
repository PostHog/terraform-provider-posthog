# EventsHeatMapStructuredResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllAggregations** | **int32** |  | 
**ColumnAggregations** | [**[]EventsHeatMapColumnAggregationResult**](EventsHeatMapColumnAggregationResult.md) |  | 
**Data** | [**[]EventsHeatMapDataResult**](EventsHeatMapDataResult.md) |  | 
**RowAggregations** | [**[]EventsHeatMapRowAggregationResult**](EventsHeatMapRowAggregationResult.md) |  | 

## Methods

### NewEventsHeatMapStructuredResult

`func NewEventsHeatMapStructuredResult(allAggregations int32, columnAggregations []EventsHeatMapColumnAggregationResult, data []EventsHeatMapDataResult, rowAggregations []EventsHeatMapRowAggregationResult, ) *EventsHeatMapStructuredResult`

NewEventsHeatMapStructuredResult instantiates a new EventsHeatMapStructuredResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsHeatMapStructuredResultWithDefaults

`func NewEventsHeatMapStructuredResultWithDefaults() *EventsHeatMapStructuredResult`

NewEventsHeatMapStructuredResultWithDefaults instantiates a new EventsHeatMapStructuredResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllAggregations

`func (o *EventsHeatMapStructuredResult) GetAllAggregations() int32`

GetAllAggregations returns the AllAggregations field if non-nil, zero value otherwise.

### GetAllAggregationsOk

`func (o *EventsHeatMapStructuredResult) GetAllAggregationsOk() (*int32, bool)`

GetAllAggregationsOk returns a tuple with the AllAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAggregations

`func (o *EventsHeatMapStructuredResult) SetAllAggregations(v int32)`

SetAllAggregations sets AllAggregations field to given value.


### GetColumnAggregations

`func (o *EventsHeatMapStructuredResult) GetColumnAggregations() []EventsHeatMapColumnAggregationResult`

GetColumnAggregations returns the ColumnAggregations field if non-nil, zero value otherwise.

### GetColumnAggregationsOk

`func (o *EventsHeatMapStructuredResult) GetColumnAggregationsOk() (*[]EventsHeatMapColumnAggregationResult, bool)`

GetColumnAggregationsOk returns a tuple with the ColumnAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnAggregations

`func (o *EventsHeatMapStructuredResult) SetColumnAggregations(v []EventsHeatMapColumnAggregationResult)`

SetColumnAggregations sets ColumnAggregations field to given value.


### GetData

`func (o *EventsHeatMapStructuredResult) GetData() []EventsHeatMapDataResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventsHeatMapStructuredResult) GetDataOk() (*[]EventsHeatMapDataResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventsHeatMapStructuredResult) SetData(v []EventsHeatMapDataResult)`

SetData sets Data field to given value.


### GetRowAggregations

`func (o *EventsHeatMapStructuredResult) GetRowAggregations() []EventsHeatMapRowAggregationResult`

GetRowAggregations returns the RowAggregations field if non-nil, zero value otherwise.

### GetRowAggregationsOk

`func (o *EventsHeatMapStructuredResult) GetRowAggregationsOk() (*[]EventsHeatMapRowAggregationResult, bool)`

GetRowAggregationsOk returns a tuple with the RowAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowAggregations

`func (o *EventsHeatMapStructuredResult) SetRowAggregations(v []EventsHeatMapRowAggregationResult)`

SetRowAggregations sets RowAggregations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


