# WebAnalyticsExternalSummaryQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]interface{}** |  | 
**Error** | Pointer to [**ExternalQueryError**](ExternalQueryError.md) |  | [optional] 
**Status** | [**ExternalQueryStatus**](ExternalQueryStatus.md) |  | 

## Methods

### NewWebAnalyticsExternalSummaryQueryResponse

`func NewWebAnalyticsExternalSummaryQueryResponse(data map[string]interface{}, status ExternalQueryStatus, ) *WebAnalyticsExternalSummaryQueryResponse`

NewWebAnalyticsExternalSummaryQueryResponse instantiates a new WebAnalyticsExternalSummaryQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAnalyticsExternalSummaryQueryResponseWithDefaults

`func NewWebAnalyticsExternalSummaryQueryResponseWithDefaults() *WebAnalyticsExternalSummaryQueryResponse`

NewWebAnalyticsExternalSummaryQueryResponseWithDefaults instantiates a new WebAnalyticsExternalSummaryQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WebAnalyticsExternalSummaryQueryResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebAnalyticsExternalSummaryQueryResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebAnalyticsExternalSummaryQueryResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetError

`func (o *WebAnalyticsExternalSummaryQueryResponse) GetError() ExternalQueryError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebAnalyticsExternalSummaryQueryResponse) GetErrorOk() (*ExternalQueryError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebAnalyticsExternalSummaryQueryResponse) SetError(v ExternalQueryError)`

SetError sets Error field to given value.

### HasError

`func (o *WebAnalyticsExternalSummaryQueryResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatus

`func (o *WebAnalyticsExternalSummaryQueryResponse) GetStatus() ExternalQueryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebAnalyticsExternalSummaryQueryResponse) GetStatusOk() (*ExternalQueryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebAnalyticsExternalSummaryQueryResponse) SetStatus(v ExternalQueryStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


