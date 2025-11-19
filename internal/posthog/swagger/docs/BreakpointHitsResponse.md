# BreakpointHitsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]BreakpointHit**](BreakpointHit.md) | List of breakpoint hit events | 
**Count** | **int32** | Number of results returned | 
**HasMore** | **bool** | Whether there are more results available | 

## Methods

### NewBreakpointHitsResponse

`func NewBreakpointHitsResponse(results []BreakpointHit, count int32, hasMore bool, ) *BreakpointHitsResponse`

NewBreakpointHitsResponse instantiates a new BreakpointHitsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreakpointHitsResponseWithDefaults

`func NewBreakpointHitsResponseWithDefaults() *BreakpointHitsResponse`

NewBreakpointHitsResponseWithDefaults instantiates a new BreakpointHitsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BreakpointHitsResponse) GetResults() []BreakpointHit`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BreakpointHitsResponse) GetResultsOk() (*[]BreakpointHit, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BreakpointHitsResponse) SetResults(v []BreakpointHit)`

SetResults sets Results field to given value.


### GetCount

`func (o *BreakpointHitsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BreakpointHitsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BreakpointHitsResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetHasMore

`func (o *BreakpointHitsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *BreakpointHitsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *BreakpointHitsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


