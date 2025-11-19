# SamplingRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denominator** | Pointer to **NullableFloat32** |  | [optional] 
**Numerator** | **float32** |  | 

## Methods

### NewSamplingRate

`func NewSamplingRate(numerator float32, ) *SamplingRate`

NewSamplingRate instantiates a new SamplingRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamplingRateWithDefaults

`func NewSamplingRateWithDefaults() *SamplingRate`

NewSamplingRateWithDefaults instantiates a new SamplingRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenominator

`func (o *SamplingRate) GetDenominator() float32`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *SamplingRate) GetDenominatorOk() (*float32, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *SamplingRate) SetDenominator(v float32)`

SetDenominator sets Denominator field to given value.

### HasDenominator

`func (o *SamplingRate) HasDenominator() bool`

HasDenominator returns a boolean if a field has been set.

### SetDenominatorNil

`func (o *SamplingRate) SetDenominatorNil(b bool)`

 SetDenominatorNil sets the value for Denominator to be an explicit nil

### UnsetDenominator
`func (o *SamplingRate) UnsetDenominator()`

UnsetDenominator ensures that no value is present for Denominator, not even an explicit nil
### GetNumerator

`func (o *SamplingRate) GetNumerator() float32`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *SamplingRate) GetNumeratorOk() (*float32, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *SamplingRate) SetNumerator(v float32)`

SetNumerator sets Numerator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


