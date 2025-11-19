# PatchedGroupUsageMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Format** | Pointer to [**GroupUsageMetricFormatEnum**](GroupUsageMetricFormatEnum.md) |  | [optional] 
**Interval** | Pointer to **int32** | In days | [optional] 
**Display** | Pointer to [**DisplayEnum**](DisplayEnum.md) |  | [optional] 
**Filters** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchedGroupUsageMetric

`func NewPatchedGroupUsageMetric() *PatchedGroupUsageMetric`

NewPatchedGroupUsageMetric instantiates a new PatchedGroupUsageMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGroupUsageMetricWithDefaults

`func NewPatchedGroupUsageMetricWithDefaults() *PatchedGroupUsageMetric`

NewPatchedGroupUsageMetricWithDefaults instantiates a new PatchedGroupUsageMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedGroupUsageMetric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedGroupUsageMetric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedGroupUsageMetric) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedGroupUsageMetric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedGroupUsageMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedGroupUsageMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedGroupUsageMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedGroupUsageMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormat

`func (o *PatchedGroupUsageMetric) GetFormat() GroupUsageMetricFormatEnum`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PatchedGroupUsageMetric) GetFormatOk() (*GroupUsageMetricFormatEnum, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PatchedGroupUsageMetric) SetFormat(v GroupUsageMetricFormatEnum)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PatchedGroupUsageMetric) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetInterval

`func (o *PatchedGroupUsageMetric) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PatchedGroupUsageMetric) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PatchedGroupUsageMetric) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PatchedGroupUsageMetric) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedGroupUsageMetric) GetDisplay() DisplayEnum`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedGroupUsageMetric) GetDisplayOk() (*DisplayEnum, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedGroupUsageMetric) SetDisplay(v DisplayEnum)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedGroupUsageMetric) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetFilters

`func (o *PatchedGroupUsageMetric) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PatchedGroupUsageMetric) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PatchedGroupUsageMetric) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PatchedGroupUsageMetric) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *PatchedGroupUsageMetric) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *PatchedGroupUsageMetric) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


