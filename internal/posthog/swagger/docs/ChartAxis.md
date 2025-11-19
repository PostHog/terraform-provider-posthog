# ChartAxis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | **string** |  | 
**Settings** | Pointer to [**Settings**](Settings.md) |  | [optional] 

## Methods

### NewChartAxis

`func NewChartAxis(column string, ) *ChartAxis`

NewChartAxis instantiates a new ChartAxis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartAxisWithDefaults

`func NewChartAxisWithDefaults() *ChartAxis`

NewChartAxisWithDefaults instantiates a new ChartAxis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *ChartAxis) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *ChartAxis) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *ChartAxis) SetColumn(v string)`

SetColumn sets Column field to given value.


### GetSettings

`func (o *ChartAxis) GetSettings() Settings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ChartAxis) GetSettingsOk() (*Settings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ChartAxis) SetSettings(v Settings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ChartAxis) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


