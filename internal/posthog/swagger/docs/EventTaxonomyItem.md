# EventTaxonomyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | **string** |  | 
**SampleCount** | **int32** |  | 
**SampleValues** | **[]string** |  | 

## Methods

### NewEventTaxonomyItem

`func NewEventTaxonomyItem(property string, sampleCount int32, sampleValues []string, ) *EventTaxonomyItem`

NewEventTaxonomyItem instantiates a new EventTaxonomyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTaxonomyItemWithDefaults

`func NewEventTaxonomyItemWithDefaults() *EventTaxonomyItem`

NewEventTaxonomyItemWithDefaults instantiates a new EventTaxonomyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *EventTaxonomyItem) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *EventTaxonomyItem) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *EventTaxonomyItem) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetSampleCount

`func (o *EventTaxonomyItem) GetSampleCount() int32`

GetSampleCount returns the SampleCount field if non-nil, zero value otherwise.

### GetSampleCountOk

`func (o *EventTaxonomyItem) GetSampleCountOk() (*int32, bool)`

GetSampleCountOk returns a tuple with the SampleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleCount

`func (o *EventTaxonomyItem) SetSampleCount(v int32)`

SetSampleCount sets SampleCount field to given value.


### GetSampleValues

`func (o *EventTaxonomyItem) GetSampleValues() []string`

GetSampleValues returns the SampleValues field if non-nil, zero value otherwise.

### GetSampleValuesOk

`func (o *EventTaxonomyItem) GetSampleValuesOk() (*[]string, bool)`

GetSampleValuesOk returns a tuple with the SampleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleValues

`func (o *EventTaxonomyItem) SetSampleValues(v []string)`

SetSampleValues sets SampleValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


