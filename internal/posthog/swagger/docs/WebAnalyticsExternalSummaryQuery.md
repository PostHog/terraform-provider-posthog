# WebAnalyticsExternalSummaryQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRange** | [**DateRange**](DateRange.md) |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "WebAnalyticsExternalSummaryQuery"]
**Properties** | [**[]Properties2Inner1**](Properties2Inner1.md) |  | 
**Response** | Pointer to [**WebAnalyticsExternalSummaryQueryResponse**](WebAnalyticsExternalSummaryQueryResponse.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewWebAnalyticsExternalSummaryQuery

`func NewWebAnalyticsExternalSummaryQuery(dateRange DateRange, properties []Properties2Inner1, ) *WebAnalyticsExternalSummaryQuery`

NewWebAnalyticsExternalSummaryQuery instantiates a new WebAnalyticsExternalSummaryQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAnalyticsExternalSummaryQueryWithDefaults

`func NewWebAnalyticsExternalSummaryQueryWithDefaults() *WebAnalyticsExternalSummaryQuery`

NewWebAnalyticsExternalSummaryQueryWithDefaults instantiates a new WebAnalyticsExternalSummaryQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRange

`func (o *WebAnalyticsExternalSummaryQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *WebAnalyticsExternalSummaryQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *WebAnalyticsExternalSummaryQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetKind

`func (o *WebAnalyticsExternalSummaryQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WebAnalyticsExternalSummaryQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WebAnalyticsExternalSummaryQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WebAnalyticsExternalSummaryQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProperties

`func (o *WebAnalyticsExternalSummaryQuery) GetProperties() []Properties2Inner1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WebAnalyticsExternalSummaryQuery) GetPropertiesOk() (*[]Properties2Inner1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WebAnalyticsExternalSummaryQuery) SetProperties(v []Properties2Inner1)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *WebAnalyticsExternalSummaryQuery) GetResponse() WebAnalyticsExternalSummaryQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebAnalyticsExternalSummaryQuery) GetResponseOk() (*WebAnalyticsExternalSummaryQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WebAnalyticsExternalSummaryQuery) SetResponse(v WebAnalyticsExternalSummaryQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WebAnalyticsExternalSummaryQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetVersion

`func (o *WebAnalyticsExternalSummaryQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WebAnalyticsExternalSummaryQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WebAnalyticsExternalSummaryQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WebAnalyticsExternalSummaryQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *WebAnalyticsExternalSummaryQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *WebAnalyticsExternalSummaryQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


