# EventOddsRatioSerialized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationType** | [**CorrelationType**](CorrelationType.md) |  | 
**Event** | [**EventDefinition**](EventDefinition.md) |  | 
**FailureCount** | **int32** |  | 
**OddsRatio** | **float32** |  | 
**SuccessCount** | **int32** |  | 

## Methods

### NewEventOddsRatioSerialized

`func NewEventOddsRatioSerialized(correlationType CorrelationType, event EventDefinition, failureCount int32, oddsRatio float32, successCount int32, ) *EventOddsRatioSerialized`

NewEventOddsRatioSerialized instantiates a new EventOddsRatioSerialized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventOddsRatioSerializedWithDefaults

`func NewEventOddsRatioSerializedWithDefaults() *EventOddsRatioSerialized`

NewEventOddsRatioSerializedWithDefaults instantiates a new EventOddsRatioSerialized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationType

`func (o *EventOddsRatioSerialized) GetCorrelationType() CorrelationType`

GetCorrelationType returns the CorrelationType field if non-nil, zero value otherwise.

### GetCorrelationTypeOk

`func (o *EventOddsRatioSerialized) GetCorrelationTypeOk() (*CorrelationType, bool)`

GetCorrelationTypeOk returns a tuple with the CorrelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationType

`func (o *EventOddsRatioSerialized) SetCorrelationType(v CorrelationType)`

SetCorrelationType sets CorrelationType field to given value.


### GetEvent

`func (o *EventOddsRatioSerialized) GetEvent() EventDefinition`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventOddsRatioSerialized) GetEventOk() (*EventDefinition, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventOddsRatioSerialized) SetEvent(v EventDefinition)`

SetEvent sets Event field to given value.


### GetFailureCount

`func (o *EventOddsRatioSerialized) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *EventOddsRatioSerialized) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *EventOddsRatioSerialized) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.


### GetOddsRatio

`func (o *EventOddsRatioSerialized) GetOddsRatio() float32`

GetOddsRatio returns the OddsRatio field if non-nil, zero value otherwise.

### GetOddsRatioOk

`func (o *EventOddsRatioSerialized) GetOddsRatioOk() (*float32, bool)`

GetOddsRatioOk returns a tuple with the OddsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOddsRatio

`func (o *EventOddsRatioSerialized) SetOddsRatio(v float32)`

SetOddsRatio sets OddsRatio field to given value.


### GetSuccessCount

`func (o *EventOddsRatioSerialized) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *EventOddsRatioSerialized) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *EventOddsRatioSerialized) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


