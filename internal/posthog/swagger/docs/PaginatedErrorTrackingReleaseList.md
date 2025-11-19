# PaginatedErrorTrackingReleaseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]ErrorTrackingRelease**](ErrorTrackingRelease.md) |  | 

## Methods

### NewPaginatedErrorTrackingReleaseList

`func NewPaginatedErrorTrackingReleaseList(count int32, results []ErrorTrackingRelease, ) *PaginatedErrorTrackingReleaseList`

NewPaginatedErrorTrackingReleaseList instantiates a new PaginatedErrorTrackingReleaseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedErrorTrackingReleaseListWithDefaults

`func NewPaginatedErrorTrackingReleaseListWithDefaults() *PaginatedErrorTrackingReleaseList`

NewPaginatedErrorTrackingReleaseListWithDefaults instantiates a new PaginatedErrorTrackingReleaseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedErrorTrackingReleaseList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedErrorTrackingReleaseList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedErrorTrackingReleaseList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedErrorTrackingReleaseList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedErrorTrackingReleaseList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedErrorTrackingReleaseList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedErrorTrackingReleaseList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedErrorTrackingReleaseList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedErrorTrackingReleaseList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedErrorTrackingReleaseList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedErrorTrackingReleaseList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedErrorTrackingReleaseList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedErrorTrackingReleaseList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedErrorTrackingReleaseList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedErrorTrackingReleaseList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedErrorTrackingReleaseList) GetResults() []ErrorTrackingRelease`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedErrorTrackingReleaseList) GetResultsOk() (*[]ErrorTrackingRelease, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedErrorTrackingReleaseList) SetResults(v []ErrorTrackingRelease)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


