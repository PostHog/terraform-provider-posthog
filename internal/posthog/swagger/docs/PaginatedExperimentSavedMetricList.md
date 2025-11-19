# PaginatedExperimentSavedMetricList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]ExperimentSavedMetric**](ExperimentSavedMetric.md) |  | 

## Methods

### NewPaginatedExperimentSavedMetricList

`func NewPaginatedExperimentSavedMetricList(count int32, results []ExperimentSavedMetric, ) *PaginatedExperimentSavedMetricList`

NewPaginatedExperimentSavedMetricList instantiates a new PaginatedExperimentSavedMetricList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedExperimentSavedMetricListWithDefaults

`func NewPaginatedExperimentSavedMetricListWithDefaults() *PaginatedExperimentSavedMetricList`

NewPaginatedExperimentSavedMetricListWithDefaults instantiates a new PaginatedExperimentSavedMetricList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedExperimentSavedMetricList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedExperimentSavedMetricList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedExperimentSavedMetricList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedExperimentSavedMetricList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedExperimentSavedMetricList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedExperimentSavedMetricList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedExperimentSavedMetricList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedExperimentSavedMetricList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedExperimentSavedMetricList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedExperimentSavedMetricList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedExperimentSavedMetricList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedExperimentSavedMetricList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedExperimentSavedMetricList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedExperimentSavedMetricList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedExperimentSavedMetricList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedExperimentSavedMetricList) GetResults() []ExperimentSavedMetric`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedExperimentSavedMetricList) GetResultsOk() (*[]ExperimentSavedMetric, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedExperimentSavedMetricList) SetResults(v []ExperimentSavedMetric)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


