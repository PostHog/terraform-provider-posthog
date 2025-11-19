# DashboardFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakdownFilter** | Pointer to [**BreakdownFilter**](BreakdownFilter.md) |  | [optional] 
**DateFrom** | Pointer to **NullableString** |  | [optional] 
**DateTo** | Pointer to **NullableString** |  | [optional] 
**Properties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) |  | [optional] 

## Methods

### NewDashboardFilter

`func NewDashboardFilter() *DashboardFilter`

NewDashboardFilter instantiates a new DashboardFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardFilterWithDefaults

`func NewDashboardFilterWithDefaults() *DashboardFilter`

NewDashboardFilterWithDefaults instantiates a new DashboardFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdownFilter

`func (o *DashboardFilter) GetBreakdownFilter() BreakdownFilter`

GetBreakdownFilter returns the BreakdownFilter field if non-nil, zero value otherwise.

### GetBreakdownFilterOk

`func (o *DashboardFilter) GetBreakdownFilterOk() (*BreakdownFilter, bool)`

GetBreakdownFilterOk returns a tuple with the BreakdownFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownFilter

`func (o *DashboardFilter) SetBreakdownFilter(v BreakdownFilter)`

SetBreakdownFilter sets BreakdownFilter field to given value.

### HasBreakdownFilter

`func (o *DashboardFilter) HasBreakdownFilter() bool`

HasBreakdownFilter returns a boolean if a field has been set.

### GetDateFrom

`func (o *DashboardFilter) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *DashboardFilter) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *DashboardFilter) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *DashboardFilter) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### SetDateFromNil

`func (o *DashboardFilter) SetDateFromNil(b bool)`

 SetDateFromNil sets the value for DateFrom to be an explicit nil

### UnsetDateFrom
`func (o *DashboardFilter) UnsetDateFrom()`

UnsetDateFrom ensures that no value is present for DateFrom, not even an explicit nil
### GetDateTo

`func (o *DashboardFilter) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *DashboardFilter) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *DashboardFilter) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *DashboardFilter) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### SetDateToNil

`func (o *DashboardFilter) SetDateToNil(b bool)`

 SetDateToNil sets the value for DateTo to be an explicit nil

### UnsetDateTo
`func (o *DashboardFilter) UnsetDateTo()`

UnsetDateTo ensures that no value is present for DateTo, not even an explicit nil
### GetProperties

`func (o *DashboardFilter) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DashboardFilter) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DashboardFilter) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DashboardFilter) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *DashboardFilter) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *DashboardFilter) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


