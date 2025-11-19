# HogQLQueryResponse

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

### NewHogQLQueryResponse

`func NewHogQLQueryResponse(results []interface{}, ) *HogQLQueryResponse`

NewHogQLQueryResponse instantiates a new HogQLQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogQLQueryResponseWithDefaults

`func NewHogQLQueryResponseWithDefaults() *HogQLQueryResponse`

NewHogQLQueryResponseWithDefaults instantiates a new HogQLQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClickhouse

`func (o *HogQLQueryResponse) GetClickhouse() string`

GetClickhouse returns the Clickhouse field if non-nil, zero value otherwise.

### GetClickhouseOk

`func (o *HogQLQueryResponse) GetClickhouseOk() (*string, bool)`

GetClickhouseOk returns a tuple with the Clickhouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouse

`func (o *HogQLQueryResponse) SetClickhouse(v string)`

SetClickhouse sets Clickhouse field to given value.

### HasClickhouse

`func (o *HogQLQueryResponse) HasClickhouse() bool`

HasClickhouse returns a boolean if a field has been set.

### SetClickhouseNil

`func (o *HogQLQueryResponse) SetClickhouseNil(b bool)`

 SetClickhouseNil sets the value for Clickhouse to be an explicit nil

### UnsetClickhouse
`func (o *HogQLQueryResponse) UnsetClickhouse()`

UnsetClickhouse ensures that no value is present for Clickhouse, not even an explicit nil
### GetColumns

`func (o *HogQLQueryResponse) GetColumns() []interface{}`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *HogQLQueryResponse) GetColumnsOk() (*[]interface{}, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *HogQLQueryResponse) SetColumns(v []interface{})`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *HogQLQueryResponse) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *HogQLQueryResponse) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *HogQLQueryResponse) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetError

`func (o *HogQLQueryResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HogQLQueryResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HogQLQueryResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *HogQLQueryResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *HogQLQueryResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HogQLQueryResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetExplain

`func (o *HogQLQueryResponse) GetExplain() []string`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *HogQLQueryResponse) GetExplainOk() (*[]string, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *HogQLQueryResponse) SetExplain(v []string)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *HogQLQueryResponse) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *HogQLQueryResponse) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *HogQLQueryResponse) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetHasMore

`func (o *HogQLQueryResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *HogQLQueryResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *HogQLQueryResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *HogQLQueryResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMoreNil

`func (o *HogQLQueryResponse) SetHasMoreNil(b bool)`

 SetHasMoreNil sets the value for HasMore to be an explicit nil

### UnsetHasMore
`func (o *HogQLQueryResponse) UnsetHasMore()`

UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil
### GetHogql

`func (o *HogQLQueryResponse) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *HogQLQueryResponse) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *HogQLQueryResponse) SetHogql(v string)`

SetHogql sets Hogql field to given value.

### HasHogql

`func (o *HogQLQueryResponse) HasHogql() bool`

HasHogql returns a boolean if a field has been set.

### SetHogqlNil

`func (o *HogQLQueryResponse) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *HogQLQueryResponse) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetLimit

`func (o *HogQLQueryResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *HogQLQueryResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *HogQLQueryResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *HogQLQueryResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *HogQLQueryResponse) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *HogQLQueryResponse) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetMetadata

`func (o *HogQLQueryResponse) GetMetadata() HogQLMetadataResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HogQLQueryResponse) GetMetadataOk() (*HogQLMetadataResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HogQLQueryResponse) SetMetadata(v HogQLMetadataResponse)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HogQLQueryResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModifiers

`func (o *HogQLQueryResponse) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *HogQLQueryResponse) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *HogQLQueryResponse) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *HogQLQueryResponse) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *HogQLQueryResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *HogQLQueryResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *HogQLQueryResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *HogQLQueryResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *HogQLQueryResponse) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *HogQLQueryResponse) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetQuery

`func (o *HogQLQueryResponse) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *HogQLQueryResponse) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *HogQLQueryResponse) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *HogQLQueryResponse) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *HogQLQueryResponse) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *HogQLQueryResponse) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetQueryStatus

`func (o *HogQLQueryResponse) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *HogQLQueryResponse) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *HogQLQueryResponse) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *HogQLQueryResponse) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *HogQLQueryResponse) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *HogQLQueryResponse) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *HogQLQueryResponse) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *HogQLQueryResponse) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *HogQLQueryResponse) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *HogQLQueryResponse) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *HogQLQueryResponse) SetResults(v []interface{})`

SetResults sets Results field to given value.


### GetTimings

`func (o *HogQLQueryResponse) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *HogQLQueryResponse) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *HogQLQueryResponse) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *HogQLQueryResponse) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *HogQLQueryResponse) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *HogQLQueryResponse) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil
### GetTypes

`func (o *HogQLQueryResponse) GetTypes() []interface{}`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *HogQLQueryResponse) GetTypesOk() (*[]interface{}, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *HogQLQueryResponse) SetTypes(v []interface{})`

SetTypes sets Types field to given value.

### HasTypes

`func (o *HogQLQueryResponse) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *HogQLQueryResponse) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *HogQLQueryResponse) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


