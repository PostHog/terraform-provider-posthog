# WebOverviewItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeFromPreviousPct** | Pointer to **NullableFloat32** |  | [optional] 
**IsIncreaseBad** | Pointer to **NullableBool** |  | [optional] 
**Key** | **string** |  | 
**Kind** | [**WebAnalyticsItemKind**](WebAnalyticsItemKind.md) |  | 
**Previous** | Pointer to **NullableFloat32** |  | [optional] 
**UsedPreAggregatedTables** | Pointer to **NullableBool** |  | [optional] 
**Value** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewWebOverviewItem

`func NewWebOverviewItem(key string, kind WebAnalyticsItemKind, ) *WebOverviewItem`

NewWebOverviewItem instantiates a new WebOverviewItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebOverviewItemWithDefaults

`func NewWebOverviewItemWithDefaults() *WebOverviewItem`

NewWebOverviewItemWithDefaults instantiates a new WebOverviewItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeFromPreviousPct

`func (o *WebOverviewItem) GetChangeFromPreviousPct() float32`

GetChangeFromPreviousPct returns the ChangeFromPreviousPct field if non-nil, zero value otherwise.

### GetChangeFromPreviousPctOk

`func (o *WebOverviewItem) GetChangeFromPreviousPctOk() (*float32, bool)`

GetChangeFromPreviousPctOk returns a tuple with the ChangeFromPreviousPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFromPreviousPct

`func (o *WebOverviewItem) SetChangeFromPreviousPct(v float32)`

SetChangeFromPreviousPct sets ChangeFromPreviousPct field to given value.

### HasChangeFromPreviousPct

`func (o *WebOverviewItem) HasChangeFromPreviousPct() bool`

HasChangeFromPreviousPct returns a boolean if a field has been set.

### SetChangeFromPreviousPctNil

`func (o *WebOverviewItem) SetChangeFromPreviousPctNil(b bool)`

 SetChangeFromPreviousPctNil sets the value for ChangeFromPreviousPct to be an explicit nil

### UnsetChangeFromPreviousPct
`func (o *WebOverviewItem) UnsetChangeFromPreviousPct()`

UnsetChangeFromPreviousPct ensures that no value is present for ChangeFromPreviousPct, not even an explicit nil
### GetIsIncreaseBad

`func (o *WebOverviewItem) GetIsIncreaseBad() bool`

GetIsIncreaseBad returns the IsIncreaseBad field if non-nil, zero value otherwise.

### GetIsIncreaseBadOk

`func (o *WebOverviewItem) GetIsIncreaseBadOk() (*bool, bool)`

GetIsIncreaseBadOk returns a tuple with the IsIncreaseBad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncreaseBad

`func (o *WebOverviewItem) SetIsIncreaseBad(v bool)`

SetIsIncreaseBad sets IsIncreaseBad field to given value.

### HasIsIncreaseBad

`func (o *WebOverviewItem) HasIsIncreaseBad() bool`

HasIsIncreaseBad returns a boolean if a field has been set.

### SetIsIncreaseBadNil

`func (o *WebOverviewItem) SetIsIncreaseBadNil(b bool)`

 SetIsIncreaseBadNil sets the value for IsIncreaseBad to be an explicit nil

### UnsetIsIncreaseBad
`func (o *WebOverviewItem) UnsetIsIncreaseBad()`

UnsetIsIncreaseBad ensures that no value is present for IsIncreaseBad, not even an explicit nil
### GetKey

`func (o *WebOverviewItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WebOverviewItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WebOverviewItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetKind

`func (o *WebOverviewItem) GetKind() WebAnalyticsItemKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WebOverviewItem) GetKindOk() (*WebAnalyticsItemKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WebOverviewItem) SetKind(v WebAnalyticsItemKind)`

SetKind sets Kind field to given value.


### GetPrevious

`func (o *WebOverviewItem) GetPrevious() float32`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *WebOverviewItem) GetPreviousOk() (*float32, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *WebOverviewItem) SetPrevious(v float32)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *WebOverviewItem) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *WebOverviewItem) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *WebOverviewItem) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetUsedPreAggregatedTables

`func (o *WebOverviewItem) GetUsedPreAggregatedTables() bool`

GetUsedPreAggregatedTables returns the UsedPreAggregatedTables field if non-nil, zero value otherwise.

### GetUsedPreAggregatedTablesOk

`func (o *WebOverviewItem) GetUsedPreAggregatedTablesOk() (*bool, bool)`

GetUsedPreAggregatedTablesOk returns a tuple with the UsedPreAggregatedTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPreAggregatedTables

`func (o *WebOverviewItem) SetUsedPreAggregatedTables(v bool)`

SetUsedPreAggregatedTables sets UsedPreAggregatedTables field to given value.

### HasUsedPreAggregatedTables

`func (o *WebOverviewItem) HasUsedPreAggregatedTables() bool`

HasUsedPreAggregatedTables returns a boolean if a field has been set.

### SetUsedPreAggregatedTablesNil

`func (o *WebOverviewItem) SetUsedPreAggregatedTablesNil(b bool)`

 SetUsedPreAggregatedTablesNil sets the value for UsedPreAggregatedTables to be an explicit nil

### UnsetUsedPreAggregatedTables
`func (o *WebOverviewItem) UnsetUsedPreAggregatedTables()`

UnsetUsedPreAggregatedTables ensures that no value is present for UsedPreAggregatedTables, not even an explicit nil
### GetValue

`func (o *WebOverviewItem) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebOverviewItem) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebOverviewItem) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *WebOverviewItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *WebOverviewItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WebOverviewItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


