# DataVisualizationNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartSettings** | Pointer to [**ChartSettings**](ChartSettings.md) |  | [optional] 
**Display** | Pointer to [**ChartDisplayType**](ChartDisplayType.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "DataVisualizationNode"]
**Source** | [**HogQLQuery**](HogQLQuery.md) |  | 
**TableSettings** | Pointer to [**TableSettings**](TableSettings.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewDataVisualizationNode

`func NewDataVisualizationNode(source HogQLQuery, ) *DataVisualizationNode`

NewDataVisualizationNode instantiates a new DataVisualizationNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataVisualizationNodeWithDefaults

`func NewDataVisualizationNodeWithDefaults() *DataVisualizationNode`

NewDataVisualizationNodeWithDefaults instantiates a new DataVisualizationNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartSettings

`func (o *DataVisualizationNode) GetChartSettings() ChartSettings`

GetChartSettings returns the ChartSettings field if non-nil, zero value otherwise.

### GetChartSettingsOk

`func (o *DataVisualizationNode) GetChartSettingsOk() (*ChartSettings, bool)`

GetChartSettingsOk returns a tuple with the ChartSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartSettings

`func (o *DataVisualizationNode) SetChartSettings(v ChartSettings)`

SetChartSettings sets ChartSettings field to given value.

### HasChartSettings

`func (o *DataVisualizationNode) HasChartSettings() bool`

HasChartSettings returns a boolean if a field has been set.

### GetDisplay

`func (o *DataVisualizationNode) GetDisplay() ChartDisplayType`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DataVisualizationNode) GetDisplayOk() (*ChartDisplayType, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DataVisualizationNode) SetDisplay(v ChartDisplayType)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *DataVisualizationNode) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetKind

`func (o *DataVisualizationNode) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DataVisualizationNode) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DataVisualizationNode) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DataVisualizationNode) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSource

`func (o *DataVisualizationNode) GetSource() HogQLQuery`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataVisualizationNode) GetSourceOk() (*HogQLQuery, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataVisualizationNode) SetSource(v HogQLQuery)`

SetSource sets Source field to given value.


### GetTableSettings

`func (o *DataVisualizationNode) GetTableSettings() TableSettings`

GetTableSettings returns the TableSettings field if non-nil, zero value otherwise.

### GetTableSettingsOk

`func (o *DataVisualizationNode) GetTableSettingsOk() (*TableSettings, bool)`

GetTableSettingsOk returns a tuple with the TableSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableSettings

`func (o *DataVisualizationNode) SetTableSettings(v TableSettings)`

SetTableSettings sets TableSettings field to given value.

### HasTableSettings

`func (o *DataVisualizationNode) HasTableSettings() bool`

HasTableSettings returns a boolean if a field has been set.

### GetVersion

`func (o *DataVisualizationNode) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DataVisualizationNode) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DataVisualizationNode) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DataVisualizationNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DataVisualizationNode) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DataVisualizationNode) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


