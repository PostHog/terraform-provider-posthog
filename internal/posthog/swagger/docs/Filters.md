# Filters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**Properties** | Pointer to [**[]SessionPropertyFilter**](SessionPropertyFilter.md) |  | [optional] 

## Methods

### NewFilters

`func NewFilters() *Filters`

NewFilters instantiates a new Filters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersWithDefaults

`func NewFiltersWithDefaults() *Filters`

NewFiltersWithDefaults instantiates a new Filters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRange

`func (o *Filters) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *Filters) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *Filters) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *Filters) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetProperties

`func (o *Filters) GetProperties() []SessionPropertyFilter`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Filters) GetPropertiesOk() (*[]SessionPropertyFilter, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Filters) SetProperties(v []SessionPropertyFilter)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Filters) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *Filters) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *Filters) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


