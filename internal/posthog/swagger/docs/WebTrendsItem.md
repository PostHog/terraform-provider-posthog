# WebTrendsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**Metrics** | [**Metrics**](Metrics.md) |  | 

## Methods

### NewWebTrendsItem

`func NewWebTrendsItem(bucket string, metrics Metrics, ) *WebTrendsItem`

NewWebTrendsItem instantiates a new WebTrendsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebTrendsItemWithDefaults

`func NewWebTrendsItemWithDefaults() *WebTrendsItem`

NewWebTrendsItemWithDefaults instantiates a new WebTrendsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *WebTrendsItem) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *WebTrendsItem) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *WebTrendsItem) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetMetrics

`func (o *WebTrendsItem) GetMetrics() Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *WebTrendsItem) GetMetricsOk() (*Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *WebTrendsItem) SetMetrics(v Metrics)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


