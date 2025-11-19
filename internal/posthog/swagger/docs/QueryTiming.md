# QueryTiming

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**K** | **string** | Key. Shortened to &#39;k&#39; to save on data. | 
**T** | **float32** | Time in seconds. Shortened to &#39;t&#39; to save on data. | 

## Methods

### NewQueryTiming

`func NewQueryTiming(k string, t float32, ) *QueryTiming`

NewQueryTiming instantiates a new QueryTiming object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryTimingWithDefaults

`func NewQueryTimingWithDefaults() *QueryTiming`

NewQueryTimingWithDefaults instantiates a new QueryTiming object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetK

`func (o *QueryTiming) GetK() string`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *QueryTiming) GetKOk() (*string, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *QueryTiming) SetK(v string)`

SetK sets K field to given value.


### GetT

`func (o *QueryTiming) GetT() float32`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *QueryTiming) GetTOk() (*float32, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *QueryTiming) SetT(v float32)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


