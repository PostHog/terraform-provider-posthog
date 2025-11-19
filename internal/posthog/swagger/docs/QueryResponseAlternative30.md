# QueryResponseAlternative30

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** |  | [optional] 
**Error** | Pointer to **NullableString** | Query error. Returned only if &#39;explain&#39; or &#x60;modifiers.debug&#x60; is true. Throws an error otherwise. | [optional] 
**Hogql** | Pointer to **NullableString** | Generated HogQL query. | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**QueryStatus** | Pointer to [**QueryStatus**](QueryStatus.md) |  | [optional] 
**ResolvedDateRange** | Pointer to [**ResolvedDateRangeResponse**](ResolvedDateRangeResponse.md) |  | [optional] 
**Results** | **[]interface{}** |  | 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 

## Methods

### NewQueryResponseAlternative30

`func NewQueryResponseAlternative30(results []interface{}, ) *QueryResponseAlternative30`

NewQueryResponseAlternative30 instantiates a new QueryResponseAlternative30 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternative30WithDefaults

`func NewQueryResponseAlternative30WithDefaults() *QueryResponseAlternative30`

NewQueryResponseAlternative30WithDefaults instantiates a new QueryResponseAlternative30 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *QueryResponseAlternative30) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *QueryResponseAlternative30) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *QueryResponseAlternative30) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *QueryResponseAlternative30) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *QueryResponseAlternative30) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *QueryResponseAlternative30) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetError

`func (o *QueryResponseAlternative30) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QueryResponseAlternative30) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QueryResponseAlternative30) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *QueryResponseAlternative30) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *QueryResponseAlternative30) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *QueryResponseAlternative30) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHogql

`func (o *QueryResponseAlternative30) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *QueryResponseAlternative30) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *QueryResponseAlternative30) SetHogql(v string)`

SetHogql sets Hogql field to given value.

### HasHogql

`func (o *QueryResponseAlternative30) HasHogql() bool`

HasHogql returns a boolean if a field has been set.

### SetHogqlNil

`func (o *QueryResponseAlternative30) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *QueryResponseAlternative30) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetModifiers

`func (o *QueryResponseAlternative30) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *QueryResponseAlternative30) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *QueryResponseAlternative30) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *QueryResponseAlternative30) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetQueryStatus

`func (o *QueryResponseAlternative30) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *QueryResponseAlternative30) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *QueryResponseAlternative30) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *QueryResponseAlternative30) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *QueryResponseAlternative30) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *QueryResponseAlternative30) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *QueryResponseAlternative30) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *QueryResponseAlternative30) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *QueryResponseAlternative30) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QueryResponseAlternative30) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QueryResponseAlternative30) SetResults(v []interface{})`

SetResults sets Results field to given value.


### GetTimings

`func (o *QueryResponseAlternative30) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *QueryResponseAlternative30) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *QueryResponseAlternative30) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *QueryResponseAlternative30) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *QueryResponseAlternative30) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *QueryResponseAlternative30) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


