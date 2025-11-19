# Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Display** | Pointer to [**ChartSettingsDisplay**](ChartSettingsDisplay.md) |  | [optional] 
**Formatting** | Pointer to [**ChartSettingsFormatting**](ChartSettingsFormatting.md) |  | [optional] 

## Methods

### NewSettings

`func NewSettings() *Settings`

NewSettings instantiates a new Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsWithDefaults

`func NewSettingsWithDefaults() *Settings`

NewSettingsWithDefaults instantiates a new Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplay

`func (o *Settings) GetDisplay() ChartSettingsDisplay`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Settings) GetDisplayOk() (*ChartSettingsDisplay, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Settings) SetDisplay(v ChartSettingsDisplay)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Settings) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetFormatting

`func (o *Settings) GetFormatting() ChartSettingsFormatting`

GetFormatting returns the Formatting field if non-nil, zero value otherwise.

### GetFormattingOk

`func (o *Settings) GetFormattingOk() (*ChartSettingsFormatting, bool)`

GetFormattingOk returns a tuple with the Formatting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatting

`func (o *Settings) SetFormatting(v ChartSettingsFormatting)`

SetFormatting sets Formatting field to given value.

### HasFormatting

`func (o *Settings) HasFormatting() bool`

HasFormatting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


