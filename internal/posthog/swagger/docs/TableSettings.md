# TableSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**[]ChartAxis**](ChartAxis.md) |  | [optional] 
**ConditionalFormatting** | Pointer to [**[]ConditionalFormattingRule**](ConditionalFormattingRule.md) |  | [optional] 

## Methods

### NewTableSettings

`func NewTableSettings() *TableSettings`

NewTableSettings instantiates a new TableSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableSettingsWithDefaults

`func NewTableSettingsWithDefaults() *TableSettings`

NewTableSettingsWithDefaults instantiates a new TableSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *TableSettings) GetColumns() []ChartAxis`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TableSettings) GetColumnsOk() (*[]ChartAxis, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TableSettings) SetColumns(v []ChartAxis)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TableSettings) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *TableSettings) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *TableSettings) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetConditionalFormatting

`func (o *TableSettings) GetConditionalFormatting() []ConditionalFormattingRule`

GetConditionalFormatting returns the ConditionalFormatting field if non-nil, zero value otherwise.

### GetConditionalFormattingOk

`func (o *TableSettings) GetConditionalFormattingOk() (*[]ConditionalFormattingRule, bool)`

GetConditionalFormattingOk returns a tuple with the ConditionalFormatting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalFormatting

`func (o *TableSettings) SetConditionalFormatting(v []ConditionalFormattingRule)`

SetConditionalFormatting sets ConditionalFormatting field to given value.

### HasConditionalFormatting

`func (o *TableSettings) HasConditionalFormatting() bool`

HasConditionalFormatting returns a boolean if a field has been set.

### SetConditionalFormattingNil

`func (o *TableSettings) SetConditionalFormattingNil(b bool)`

 SetConditionalFormattingNil sets the value for ConditionalFormatting to be an explicit nil

### UnsetConditionalFormatting
`func (o *TableSettings) UnsetConditionalFormatting()`

UnsetConditionalFormatting ensures that no value is present for ConditionalFormatting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


