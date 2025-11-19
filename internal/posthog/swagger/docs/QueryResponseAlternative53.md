# QueryResponseAlternative53

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]interface{}** |  | [optional] 
**Error** | Pointer to **NullableString** | Query error. Returned only if &#39;explain&#39; or &#x60;modifiers.debug&#x60; is true. Throws an error otherwise. | [optional] 
**HasMore** | Pointer to **NullableBool** |  | [optional] 
**Hogql** | Pointer to **NullableString** | Generated HogQL query. | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**QueryStatus** | Pointer to [**QueryStatus**](QueryStatus.md) |  | [optional] 
**ResolvedDateRange** | Pointer to [**ResolvedDateRangeResponse**](ResolvedDateRangeResponse.md) |  | [optional] 
**Results** | **interface{}** |  | 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 
**Types** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewQueryResponseAlternative53

`func NewQueryResponseAlternative53(results interface{}, ) *QueryResponseAlternative53`

NewQueryResponseAlternative53 instantiates a new QueryResponseAlternative53 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternative53WithDefaults

`func NewQueryResponseAlternative53WithDefaults() *QueryResponseAlternative53`

NewQueryResponseAlternative53WithDefaults instantiates a new QueryResponseAlternative53 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *QueryResponseAlternative53) GetColumns() []interface{}`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *QueryResponseAlternative53) GetColumnsOk() (*[]interface{}, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *QueryResponseAlternative53) SetColumns(v []interface{})`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *QueryResponseAlternative53) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *QueryResponseAlternative53) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *QueryResponseAlternative53) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetError

`func (o *QueryResponseAlternative53) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QueryResponseAlternative53) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QueryResponseAlternative53) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *QueryResponseAlternative53) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *QueryResponseAlternative53) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *QueryResponseAlternative53) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHasMore

`func (o *QueryResponseAlternative53) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *QueryResponseAlternative53) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *QueryResponseAlternative53) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *QueryResponseAlternative53) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMoreNil

`func (o *QueryResponseAlternative53) SetHasMoreNil(b bool)`

 SetHasMoreNil sets the value for HasMore to be an explicit nil

### UnsetHasMore
`func (o *QueryResponseAlternative53) UnsetHasMore()`

UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil
### GetHogql

`func (o *QueryResponseAlternative53) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *QueryResponseAlternative53) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *QueryResponseAlternative53) SetHogql(v string)`

SetHogql sets Hogql field to given value.

### HasHogql

`func (o *QueryResponseAlternative53) HasHogql() bool`

HasHogql returns a boolean if a field has been set.

### SetHogqlNil

`func (o *QueryResponseAlternative53) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *QueryResponseAlternative53) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetLimit

`func (o *QueryResponseAlternative53) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryResponseAlternative53) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryResponseAlternative53) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryResponseAlternative53) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *QueryResponseAlternative53) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *QueryResponseAlternative53) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *QueryResponseAlternative53) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *QueryResponseAlternative53) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *QueryResponseAlternative53) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *QueryResponseAlternative53) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *QueryResponseAlternative53) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryResponseAlternative53) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryResponseAlternative53) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryResponseAlternative53) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *QueryResponseAlternative53) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *QueryResponseAlternative53) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetQueryStatus

`func (o *QueryResponseAlternative53) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *QueryResponseAlternative53) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *QueryResponseAlternative53) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *QueryResponseAlternative53) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *QueryResponseAlternative53) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *QueryResponseAlternative53) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *QueryResponseAlternative53) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *QueryResponseAlternative53) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *QueryResponseAlternative53) GetResults() interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QueryResponseAlternative53) GetResultsOk() (*interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QueryResponseAlternative53) SetResults(v interface{})`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *QueryResponseAlternative53) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *QueryResponseAlternative53) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetTimings

`func (o *QueryResponseAlternative53) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *QueryResponseAlternative53) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *QueryResponseAlternative53) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *QueryResponseAlternative53) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *QueryResponseAlternative53) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *QueryResponseAlternative53) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil
### GetTypes

`func (o *QueryResponseAlternative53) GetTypes() []interface{}`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *QueryResponseAlternative53) GetTypesOk() (*[]interface{}, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *QueryResponseAlternative53) SetTypes(v []interface{})`

SetTypes sets Types field to given value.

### HasTypes

`func (o *QueryResponseAlternative53) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *QueryResponseAlternative53) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *QueryResponseAlternative53) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


