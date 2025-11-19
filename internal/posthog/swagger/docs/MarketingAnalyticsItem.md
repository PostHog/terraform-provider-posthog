# MarketingAnalyticsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeFromPreviousPct** | Pointer to **NullableFloat32** |  | [optional] 
**HasComparison** | Pointer to **NullableBool** |  | [optional] 
**IsIncreaseBad** | Pointer to **NullableBool** |  | [optional] 
**Key** | **string** |  | 
**Kind** | [**WebAnalyticsItemKind**](WebAnalyticsItemKind.md) |  | 
**Previous** | Pointer to [**NullablePrevious**](Previous.md) |  | [optional] [default to null]
**Value** | Pointer to [**NullableValue5**](Value5.md) |  | [optional] [default to null]

## Methods

### NewMarketingAnalyticsItem

`func NewMarketingAnalyticsItem(key string, kind WebAnalyticsItemKind, ) *MarketingAnalyticsItem`

NewMarketingAnalyticsItem instantiates a new MarketingAnalyticsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingAnalyticsItemWithDefaults

`func NewMarketingAnalyticsItemWithDefaults() *MarketingAnalyticsItem`

NewMarketingAnalyticsItemWithDefaults instantiates a new MarketingAnalyticsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeFromPreviousPct

`func (o *MarketingAnalyticsItem) GetChangeFromPreviousPct() float32`

GetChangeFromPreviousPct returns the ChangeFromPreviousPct field if non-nil, zero value otherwise.

### GetChangeFromPreviousPctOk

`func (o *MarketingAnalyticsItem) GetChangeFromPreviousPctOk() (*float32, bool)`

GetChangeFromPreviousPctOk returns a tuple with the ChangeFromPreviousPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFromPreviousPct

`func (o *MarketingAnalyticsItem) SetChangeFromPreviousPct(v float32)`

SetChangeFromPreviousPct sets ChangeFromPreviousPct field to given value.

### HasChangeFromPreviousPct

`func (o *MarketingAnalyticsItem) HasChangeFromPreviousPct() bool`

HasChangeFromPreviousPct returns a boolean if a field has been set.

### SetChangeFromPreviousPctNil

`func (o *MarketingAnalyticsItem) SetChangeFromPreviousPctNil(b bool)`

 SetChangeFromPreviousPctNil sets the value for ChangeFromPreviousPct to be an explicit nil

### UnsetChangeFromPreviousPct
`func (o *MarketingAnalyticsItem) UnsetChangeFromPreviousPct()`

UnsetChangeFromPreviousPct ensures that no value is present for ChangeFromPreviousPct, not even an explicit nil
### GetHasComparison

`func (o *MarketingAnalyticsItem) GetHasComparison() bool`

GetHasComparison returns the HasComparison field if non-nil, zero value otherwise.

### GetHasComparisonOk

`func (o *MarketingAnalyticsItem) GetHasComparisonOk() (*bool, bool)`

GetHasComparisonOk returns a tuple with the HasComparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasComparison

`func (o *MarketingAnalyticsItem) SetHasComparison(v bool)`

SetHasComparison sets HasComparison field to given value.

### HasHasComparison

`func (o *MarketingAnalyticsItem) HasHasComparison() bool`

HasHasComparison returns a boolean if a field has been set.

### SetHasComparisonNil

`func (o *MarketingAnalyticsItem) SetHasComparisonNil(b bool)`

 SetHasComparisonNil sets the value for HasComparison to be an explicit nil

### UnsetHasComparison
`func (o *MarketingAnalyticsItem) UnsetHasComparison()`

UnsetHasComparison ensures that no value is present for HasComparison, not even an explicit nil
### GetIsIncreaseBad

`func (o *MarketingAnalyticsItem) GetIsIncreaseBad() bool`

GetIsIncreaseBad returns the IsIncreaseBad field if non-nil, zero value otherwise.

### GetIsIncreaseBadOk

`func (o *MarketingAnalyticsItem) GetIsIncreaseBadOk() (*bool, bool)`

GetIsIncreaseBadOk returns a tuple with the IsIncreaseBad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncreaseBad

`func (o *MarketingAnalyticsItem) SetIsIncreaseBad(v bool)`

SetIsIncreaseBad sets IsIncreaseBad field to given value.

### HasIsIncreaseBad

`func (o *MarketingAnalyticsItem) HasIsIncreaseBad() bool`

HasIsIncreaseBad returns a boolean if a field has been set.

### SetIsIncreaseBadNil

`func (o *MarketingAnalyticsItem) SetIsIncreaseBadNil(b bool)`

 SetIsIncreaseBadNil sets the value for IsIncreaseBad to be an explicit nil

### UnsetIsIncreaseBad
`func (o *MarketingAnalyticsItem) UnsetIsIncreaseBad()`

UnsetIsIncreaseBad ensures that no value is present for IsIncreaseBad, not even an explicit nil
### GetKey

`func (o *MarketingAnalyticsItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MarketingAnalyticsItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MarketingAnalyticsItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetKind

`func (o *MarketingAnalyticsItem) GetKind() WebAnalyticsItemKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MarketingAnalyticsItem) GetKindOk() (*WebAnalyticsItemKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MarketingAnalyticsItem) SetKind(v WebAnalyticsItemKind)`

SetKind sets Kind field to given value.


### GetPrevious

`func (o *MarketingAnalyticsItem) GetPrevious() Previous`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *MarketingAnalyticsItem) GetPreviousOk() (*Previous, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *MarketingAnalyticsItem) SetPrevious(v Previous)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *MarketingAnalyticsItem) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *MarketingAnalyticsItem) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *MarketingAnalyticsItem) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetValue

`func (o *MarketingAnalyticsItem) GetValue() Value5`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MarketingAnalyticsItem) GetValueOk() (*Value5, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MarketingAnalyticsItem) SetValue(v Value5)`

SetValue sets Value field to given value.

### HasValue

`func (o *MarketingAnalyticsItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MarketingAnalyticsItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MarketingAnalyticsItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


