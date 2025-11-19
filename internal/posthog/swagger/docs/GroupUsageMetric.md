# GroupUsageMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Format** | Pointer to [**GroupUsageMetricFormatEnum**](GroupUsageMetricFormatEnum.md) |  | [optional] 
**Interval** | Pointer to **int32** | In days | [optional] 
**Display** | Pointer to [**DisplayEnum**](DisplayEnum.md) |  | [optional] 
**Filters** | **interface{}** |  | 

## Methods

### NewGroupUsageMetric

`func NewGroupUsageMetric(id string, name string, filters interface{}, ) *GroupUsageMetric`

NewGroupUsageMetric instantiates a new GroupUsageMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUsageMetricWithDefaults

`func NewGroupUsageMetricWithDefaults() *GroupUsageMetric`

NewGroupUsageMetricWithDefaults instantiates a new GroupUsageMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupUsageMetric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupUsageMetric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupUsageMetric) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupUsageMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupUsageMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupUsageMetric) SetName(v string)`

SetName sets Name field to given value.


### GetFormat

`func (o *GroupUsageMetric) GetFormat() GroupUsageMetricFormatEnum`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GroupUsageMetric) GetFormatOk() (*GroupUsageMetricFormatEnum, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GroupUsageMetric) SetFormat(v GroupUsageMetricFormatEnum)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GroupUsageMetric) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetInterval

`func (o *GroupUsageMetric) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GroupUsageMetric) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GroupUsageMetric) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GroupUsageMetric) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetDisplay

`func (o *GroupUsageMetric) GetDisplay() DisplayEnum`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *GroupUsageMetric) GetDisplayOk() (*DisplayEnum, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *GroupUsageMetric) SetDisplay(v DisplayEnum)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *GroupUsageMetric) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetFilters

`func (o *GroupUsageMetric) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GroupUsageMetric) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GroupUsageMetric) SetFilters(v interface{})`

SetFilters sets Filters field to given value.


### SetFiltersNil

`func (o *GroupUsageMetric) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *GroupUsageMetric) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


