# DateRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateFrom** | Pointer to **NullableString** |  | [optional] 
**DateTo** | Pointer to **NullableString** |  | [optional] 
**ExplicitDate** | Pointer to **NullableBool** | Whether the date_from and date_to should be used verbatim. Disables rounding to the start and end of period. | [optional] [default to false]

## Methods

### NewDateRange

`func NewDateRange() *DateRange`

NewDateRange instantiates a new DateRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateRangeWithDefaults

`func NewDateRangeWithDefaults() *DateRange`

NewDateRangeWithDefaults instantiates a new DateRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateFrom

`func (o *DateRange) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *DateRange) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *DateRange) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *DateRange) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### SetDateFromNil

`func (o *DateRange) SetDateFromNil(b bool)`

 SetDateFromNil sets the value for DateFrom to be an explicit nil

### UnsetDateFrom
`func (o *DateRange) UnsetDateFrom()`

UnsetDateFrom ensures that no value is present for DateFrom, not even an explicit nil
### GetDateTo

`func (o *DateRange) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *DateRange) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *DateRange) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *DateRange) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### SetDateToNil

`func (o *DateRange) SetDateToNil(b bool)`

 SetDateToNil sets the value for DateTo to be an explicit nil

### UnsetDateTo
`func (o *DateRange) UnsetDateTo()`

UnsetDateTo ensures that no value is present for DateTo, not even an explicit nil
### GetExplicitDate

`func (o *DateRange) GetExplicitDate() bool`

GetExplicitDate returns the ExplicitDate field if non-nil, zero value otherwise.

### GetExplicitDateOk

`func (o *DateRange) GetExplicitDateOk() (*bool, bool)`

GetExplicitDateOk returns a tuple with the ExplicitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitDate

`func (o *DateRange) SetExplicitDate(v bool)`

SetExplicitDate sets ExplicitDate field to given value.

### HasExplicitDate

`func (o *DateRange) HasExplicitDate() bool`

HasExplicitDate returns a boolean if a field has been set.

### SetExplicitDateNil

`func (o *DateRange) SetExplicitDateNil(b bool)`

 SetExplicitDateNil sets the value for ExplicitDate to be an explicit nil

### UnsetExplicitDate
`func (o *DateRange) UnsetExplicitDate()`

UnsetExplicitDate ensures that no value is present for ExplicitDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


