# WebAnalyticsBreakdownResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** | URL for next page of results | [optional] 
**Results** | **[]interface{}** | Array of breakdown items | 

## Methods

### NewWebAnalyticsBreakdownResponse

`func NewWebAnalyticsBreakdownResponse(results []interface{}, ) *WebAnalyticsBreakdownResponse`

NewWebAnalyticsBreakdownResponse instantiates a new WebAnalyticsBreakdownResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAnalyticsBreakdownResponseWithDefaults

`func NewWebAnalyticsBreakdownResponseWithDefaults() *WebAnalyticsBreakdownResponse`

NewWebAnalyticsBreakdownResponseWithDefaults instantiates a new WebAnalyticsBreakdownResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *WebAnalyticsBreakdownResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *WebAnalyticsBreakdownResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *WebAnalyticsBreakdownResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *WebAnalyticsBreakdownResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *WebAnalyticsBreakdownResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *WebAnalyticsBreakdownResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetResults

`func (o *WebAnalyticsBreakdownResponse) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WebAnalyticsBreakdownResponse) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WebAnalyticsBreakdownResponse) SetResults(v []interface{})`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


