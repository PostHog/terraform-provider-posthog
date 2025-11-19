# QueryResponseAlternative8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clickhouse** | Pointer to **NullableString** | Executed ClickHouse query | [optional] 
**Columns** | Pointer to **[]interface{}** | Returned columns | [optional] 
**Error** | Pointer to **NullableString** | Query error. Returned only if &#39;explain&#39; or &#x60;modifiers.debug&#x60; is true. Throws an error otherwise. | [optional] 
**Explain** | Pointer to **[]string** | Query explanation output | [optional] 
**HasMore** | Pointer to **NullableBool** |  | [optional] 
**Hogql** | Pointer to **NullableString** | Generated HogQL query. | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Metadata** | Pointer to [**HogQLMetadataResponse**](HogQLMetadataResponse.md) |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**Query** | Pointer to **NullableString** | Input query string | [optional] 
**QueryStatus** | Pointer to [**QueryStatus**](QueryStatus.md) |  | [optional] 
**ResolvedDateRange** | Pointer to [**ResolvedDateRangeResponse**](ResolvedDateRangeResponse.md) |  | [optional] 
**Results** | **[]interface{}** |  | 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 
**Types** | Pointer to **[]interface{}** | Types of returned columns | [optional] 

## Methods

### NewQueryResponseAlternative8

`func NewQueryResponseAlternative8(results []interface{}, ) *QueryResponseAlternative8`

NewQueryResponseAlternative8 instantiates a new QueryResponseAlternative8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternative8WithDefaults

`func NewQueryResponseAlternative8WithDefaults() *QueryResponseAlternative8`

NewQueryResponseAlternative8WithDefaults instantiates a new QueryResponseAlternative8 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClickhouse

`func (o *QueryResponseAlternative8) GetClickhouse() string`

GetClickhouse returns the Clickhouse field if non-nil, zero value otherwise.

### GetClickhouseOk

`func (o *QueryResponseAlternative8) GetClickhouseOk() (*string, bool)`

GetClickhouseOk returns a tuple with the Clickhouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouse

`func (o *QueryResponseAlternative8) SetClickhouse(v string)`

SetClickhouse sets Clickhouse field to given value.

### HasClickhouse

`func (o *QueryResponseAlternative8) HasClickhouse() bool`

HasClickhouse returns a boolean if a field has been set.

### SetClickhouseNil

`func (o *QueryResponseAlternative8) SetClickhouseNil(b bool)`

 SetClickhouseNil sets the value for Clickhouse to be an explicit nil

### UnsetClickhouse
`func (o *QueryResponseAlternative8) UnsetClickhouse()`

UnsetClickhouse ensures that no value is present for Clickhouse, not even an explicit nil
### GetColumns

`func (o *QueryResponseAlternative8) GetColumns() []interface{}`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *QueryResponseAlternative8) GetColumnsOk() (*[]interface{}, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *QueryResponseAlternative8) SetColumns(v []interface{})`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *QueryResponseAlternative8) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *QueryResponseAlternative8) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *QueryResponseAlternative8) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetError

`func (o *QueryResponseAlternative8) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QueryResponseAlternative8) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QueryResponseAlternative8) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *QueryResponseAlternative8) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *QueryResponseAlternative8) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *QueryResponseAlternative8) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetExplain

`func (o *QueryResponseAlternative8) GetExplain() []string`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *QueryResponseAlternative8) GetExplainOk() (*[]string, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *QueryResponseAlternative8) SetExplain(v []string)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *QueryResponseAlternative8) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *QueryResponseAlternative8) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *QueryResponseAlternative8) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetHasMore

`func (o *QueryResponseAlternative8) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *QueryResponseAlternative8) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *QueryResponseAlternative8) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *QueryResponseAlternative8) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMoreNil

`func (o *QueryResponseAlternative8) SetHasMoreNil(b bool)`

 SetHasMoreNil sets the value for HasMore to be an explicit nil

### UnsetHasMore
`func (o *QueryResponseAlternative8) UnsetHasMore()`

UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil
### GetHogql

`func (o *QueryResponseAlternative8) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *QueryResponseAlternative8) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *QueryResponseAlternative8) SetHogql(v string)`

SetHogql sets Hogql field to given value.

### HasHogql

`func (o *QueryResponseAlternative8) HasHogql() bool`

HasHogql returns a boolean if a field has been set.

### SetHogqlNil

`func (o *QueryResponseAlternative8) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *QueryResponseAlternative8) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetLimit

`func (o *QueryResponseAlternative8) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryResponseAlternative8) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryResponseAlternative8) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryResponseAlternative8) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *QueryResponseAlternative8) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *QueryResponseAlternative8) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetMetadata

`func (o *QueryResponseAlternative8) GetMetadata() HogQLMetadataResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QueryResponseAlternative8) GetMetadataOk() (*HogQLMetadataResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QueryResponseAlternative8) SetMetadata(v HogQLMetadataResponse)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QueryResponseAlternative8) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModifiers

`func (o *QueryResponseAlternative8) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *QueryResponseAlternative8) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *QueryResponseAlternative8) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *QueryResponseAlternative8) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *QueryResponseAlternative8) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryResponseAlternative8) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryResponseAlternative8) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryResponseAlternative8) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *QueryResponseAlternative8) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *QueryResponseAlternative8) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetQuery

`func (o *QueryResponseAlternative8) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryResponseAlternative8) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryResponseAlternative8) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *QueryResponseAlternative8) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *QueryResponseAlternative8) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *QueryResponseAlternative8) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetQueryStatus

`func (o *QueryResponseAlternative8) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *QueryResponseAlternative8) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *QueryResponseAlternative8) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *QueryResponseAlternative8) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *QueryResponseAlternative8) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *QueryResponseAlternative8) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *QueryResponseAlternative8) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *QueryResponseAlternative8) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *QueryResponseAlternative8) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QueryResponseAlternative8) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QueryResponseAlternative8) SetResults(v []interface{})`

SetResults sets Results field to given value.


### GetTimings

`func (o *QueryResponseAlternative8) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *QueryResponseAlternative8) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *QueryResponseAlternative8) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *QueryResponseAlternative8) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *QueryResponseAlternative8) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *QueryResponseAlternative8) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil
### GetTypes

`func (o *QueryResponseAlternative8) GetTypes() []interface{}`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *QueryResponseAlternative8) GetTypesOk() (*[]interface{}, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *QueryResponseAlternative8) SetTypes(v []interface{})`

SetTypes sets Types field to given value.

### HasTypes

`func (o *QueryResponseAlternative8) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *QueryResponseAlternative8) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *QueryResponseAlternative8) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


