# WebAnalyticsOverviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visitors** | **int32** | Unique visitors | 
**Views** | **int32** | Total page views | 
**Sessions** | **int32** | Total sessions | 
**BounceRate** | **float64** | Bounce rate | 
**SessionDuration** | **float64** | Average session duration in seconds | 

## Methods

### NewWebAnalyticsOverviewResponse

`func NewWebAnalyticsOverviewResponse(visitors int32, views int32, sessions int32, bounceRate float64, sessionDuration float64, ) *WebAnalyticsOverviewResponse`

NewWebAnalyticsOverviewResponse instantiates a new WebAnalyticsOverviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAnalyticsOverviewResponseWithDefaults

`func NewWebAnalyticsOverviewResponseWithDefaults() *WebAnalyticsOverviewResponse`

NewWebAnalyticsOverviewResponseWithDefaults instantiates a new WebAnalyticsOverviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisitors

`func (o *WebAnalyticsOverviewResponse) GetVisitors() int32`

GetVisitors returns the Visitors field if non-nil, zero value otherwise.

### GetVisitorsOk

`func (o *WebAnalyticsOverviewResponse) GetVisitorsOk() (*int32, bool)`

GetVisitorsOk returns a tuple with the Visitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitors

`func (o *WebAnalyticsOverviewResponse) SetVisitors(v int32)`

SetVisitors sets Visitors field to given value.


### GetViews

`func (o *WebAnalyticsOverviewResponse) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *WebAnalyticsOverviewResponse) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *WebAnalyticsOverviewResponse) SetViews(v int32)`

SetViews sets Views field to given value.


### GetSessions

`func (o *WebAnalyticsOverviewResponse) GetSessions() int32`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *WebAnalyticsOverviewResponse) GetSessionsOk() (*int32, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *WebAnalyticsOverviewResponse) SetSessions(v int32)`

SetSessions sets Sessions field to given value.


### GetBounceRate

`func (o *WebAnalyticsOverviewResponse) GetBounceRate() float64`

GetBounceRate returns the BounceRate field if non-nil, zero value otherwise.

### GetBounceRateOk

`func (o *WebAnalyticsOverviewResponse) GetBounceRateOk() (*float64, bool)`

GetBounceRateOk returns a tuple with the BounceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceRate

`func (o *WebAnalyticsOverviewResponse) SetBounceRate(v float64)`

SetBounceRate sets BounceRate field to given value.


### GetSessionDuration

`func (o *WebAnalyticsOverviewResponse) GetSessionDuration() float64`

GetSessionDuration returns the SessionDuration field if non-nil, zero value otherwise.

### GetSessionDurationOk

`func (o *WebAnalyticsOverviewResponse) GetSessionDurationOk() (*float64, bool)`

GetSessionDurationOk returns a tuple with the SessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDuration

`func (o *WebAnalyticsOverviewResponse) SetSessionDuration(v float64)`

SetSessionDuration sets SessionDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


