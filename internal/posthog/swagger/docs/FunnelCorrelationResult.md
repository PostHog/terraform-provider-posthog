# FunnelCorrelationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]EventOddsRatioSerialized**](EventOddsRatioSerialized.md) |  | 
**Skewed** | **bool** |  | 

## Methods

### NewFunnelCorrelationResult

`func NewFunnelCorrelationResult(events []EventOddsRatioSerialized, skewed bool, ) *FunnelCorrelationResult`

NewFunnelCorrelationResult instantiates a new FunnelCorrelationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelCorrelationResultWithDefaults

`func NewFunnelCorrelationResultWithDefaults() *FunnelCorrelationResult`

NewFunnelCorrelationResultWithDefaults instantiates a new FunnelCorrelationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *FunnelCorrelationResult) GetEvents() []EventOddsRatioSerialized`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FunnelCorrelationResult) GetEventsOk() (*[]EventOddsRatioSerialized, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FunnelCorrelationResult) SetEvents(v []EventOddsRatioSerialized)`

SetEvents sets Events field to given value.


### GetSkewed

`func (o *FunnelCorrelationResult) GetSkewed() bool`

GetSkewed returns the Skewed field if non-nil, zero value otherwise.

### GetSkewedOk

`func (o *FunnelCorrelationResult) GetSkewedOk() (*bool, bool)`

GetSkewedOk returns a tuple with the Skewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkewed

`func (o *FunnelCorrelationResult) SetSkewed(v bool)`

SetSkewed sets Skewed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


