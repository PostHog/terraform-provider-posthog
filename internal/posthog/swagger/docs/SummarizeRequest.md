# SummarizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummarizeType** | [**SummarizeTypeEnum**](SummarizeTypeEnum.md) | Type of entity to summarize  * &#x60;trace&#x60; - trace * &#x60;event&#x60; - event | 
**Mode** | Pointer to [**ModeEnum**](ModeEnum.md) | Summary detail level: &#39;minimal&#39; for 3-5 points, &#39;detailed&#39; for 5-10 points  * &#x60;minimal&#x60; - minimal * &#x60;detailed&#x60; - detailed | [optional] [default to MINIMAL]
**Data** | **interface{}** | Data to summarize. For traces: {trace, hierarchy}. For events: {event}. | 
**ForceRefresh** | Pointer to **bool** | Force regenerate summary, bypassing cache | [optional] [default to false]

## Methods

### NewSummarizeRequest

`func NewSummarizeRequest(summarizeType SummarizeTypeEnum, data interface{}, ) *SummarizeRequest`

NewSummarizeRequest instantiates a new SummarizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummarizeRequestWithDefaults

`func NewSummarizeRequestWithDefaults() *SummarizeRequest`

NewSummarizeRequestWithDefaults instantiates a new SummarizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummarizeType

`func (o *SummarizeRequest) GetSummarizeType() SummarizeTypeEnum`

GetSummarizeType returns the SummarizeType field if non-nil, zero value otherwise.

### GetSummarizeTypeOk

`func (o *SummarizeRequest) GetSummarizeTypeOk() (*SummarizeTypeEnum, bool)`

GetSummarizeTypeOk returns a tuple with the SummarizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummarizeType

`func (o *SummarizeRequest) SetSummarizeType(v SummarizeTypeEnum)`

SetSummarizeType sets SummarizeType field to given value.


### GetMode

`func (o *SummarizeRequest) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SummarizeRequest) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SummarizeRequest) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SummarizeRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetData

`func (o *SummarizeRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SummarizeRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SummarizeRequest) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *SummarizeRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SummarizeRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetForceRefresh

`func (o *SummarizeRequest) GetForceRefresh() bool`

GetForceRefresh returns the ForceRefresh field if non-nil, zero value otherwise.

### GetForceRefreshOk

`func (o *SummarizeRequest) GetForceRefreshOk() (*bool, bool)`

GetForceRefreshOk returns a tuple with the ForceRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRefresh

`func (o *SummarizeRequest) SetForceRefresh(v bool)`

SetForceRefresh sets ForceRefresh field to given value.

### HasForceRefresh

`func (o *SummarizeRequest) HasForceRefresh() bool`

HasForceRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


