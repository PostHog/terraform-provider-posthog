# UsageMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeFromPreviousPct** | Pointer to **NullableFloat32** |  | [optional] 
**Display** | [**UsageMetricDisplay**](UsageMetricDisplay.md) |  | 
**Format** | [**UsageMetricFormat**](UsageMetricFormat.md) |  | 
**Id** | **string** |  | 
**Interval** | **int32** |  | 
**Name** | **string** |  | 
**Previous** | **float32** |  | 
**Value** | **float32** |  | 

## Methods

### NewUsageMetric

`func NewUsageMetric(display UsageMetricDisplay, format UsageMetricFormat, id string, interval int32, name string, previous float32, value float32, ) *UsageMetric`

NewUsageMetric instantiates a new UsageMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMetricWithDefaults

`func NewUsageMetricWithDefaults() *UsageMetric`

NewUsageMetricWithDefaults instantiates a new UsageMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeFromPreviousPct

`func (o *UsageMetric) GetChangeFromPreviousPct() float32`

GetChangeFromPreviousPct returns the ChangeFromPreviousPct field if non-nil, zero value otherwise.

### GetChangeFromPreviousPctOk

`func (o *UsageMetric) GetChangeFromPreviousPctOk() (*float32, bool)`

GetChangeFromPreviousPctOk returns a tuple with the ChangeFromPreviousPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFromPreviousPct

`func (o *UsageMetric) SetChangeFromPreviousPct(v float32)`

SetChangeFromPreviousPct sets ChangeFromPreviousPct field to given value.

### HasChangeFromPreviousPct

`func (o *UsageMetric) HasChangeFromPreviousPct() bool`

HasChangeFromPreviousPct returns a boolean if a field has been set.

### SetChangeFromPreviousPctNil

`func (o *UsageMetric) SetChangeFromPreviousPctNil(b bool)`

 SetChangeFromPreviousPctNil sets the value for ChangeFromPreviousPct to be an explicit nil

### UnsetChangeFromPreviousPct
`func (o *UsageMetric) UnsetChangeFromPreviousPct()`

UnsetChangeFromPreviousPct ensures that no value is present for ChangeFromPreviousPct, not even an explicit nil
### GetDisplay

`func (o *UsageMetric) GetDisplay() UsageMetricDisplay`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *UsageMetric) GetDisplayOk() (*UsageMetricDisplay, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *UsageMetric) SetDisplay(v UsageMetricDisplay)`

SetDisplay sets Display field to given value.


### GetFormat

`func (o *UsageMetric) GetFormat() UsageMetricFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UsageMetric) GetFormatOk() (*UsageMetricFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UsageMetric) SetFormat(v UsageMetricFormat)`

SetFormat sets Format field to given value.


### GetId

`func (o *UsageMetric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageMetric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageMetric) SetId(v string)`

SetId sets Id field to given value.


### GetInterval

`func (o *UsageMetric) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UsageMetric) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UsageMetric) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetName

`func (o *UsageMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsageMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsageMetric) SetName(v string)`

SetName sets Name field to given value.


### GetPrevious

`func (o *UsageMetric) GetPrevious() float32`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *UsageMetric) GetPreviousOk() (*float32, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *UsageMetric) SetPrevious(v float32)`

SetPrevious sets Previous field to given value.


### GetValue

`func (o *UsageMetric) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UsageMetric) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UsageMetric) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


