# CompareFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compare** | Pointer to **NullableBool** | Whether to compare the current date range to a previous date range. | [optional] [default to false]
**CompareTo** | Pointer to **NullableString** | The date range to compare to. The value is a relative date. Examples of relative dates are: &#x60;-1y&#x60; for 1 year ago, &#x60;-14m&#x60; for 14 months ago, &#x60;-100w&#x60; for 100 weeks ago, &#x60;-14d&#x60; for 14 days ago, &#x60;-30h&#x60; for 30 hours ago. | [optional] 

## Methods

### NewCompareFilter

`func NewCompareFilter() *CompareFilter`

NewCompareFilter instantiates a new CompareFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompareFilterWithDefaults

`func NewCompareFilterWithDefaults() *CompareFilter`

NewCompareFilterWithDefaults instantiates a new CompareFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompare

`func (o *CompareFilter) GetCompare() bool`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *CompareFilter) GetCompareOk() (*bool, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *CompareFilter) SetCompare(v bool)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *CompareFilter) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### SetCompareNil

`func (o *CompareFilter) SetCompareNil(b bool)`

 SetCompareNil sets the value for Compare to be an explicit nil

### UnsetCompare
`func (o *CompareFilter) UnsetCompare()`

UnsetCompare ensures that no value is present for Compare, not even an explicit nil
### GetCompareTo

`func (o *CompareFilter) GetCompareTo() string`

GetCompareTo returns the CompareTo field if non-nil, zero value otherwise.

### GetCompareToOk

`func (o *CompareFilter) GetCompareToOk() (*string, bool)`

GetCompareToOk returns a tuple with the CompareTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareTo

`func (o *CompareFilter) SetCompareTo(v string)`

SetCompareTo sets CompareTo field to given value.

### HasCompareTo

`func (o *CompareFilter) HasCompareTo() bool`

HasCompareTo returns a boolean if a field has been set.

### SetCompareToNil

`func (o *CompareFilter) SetCompareToNil(b bool)`

 SetCompareToNil sets the value for CompareTo to be an explicit nil

### UnsetCompareTo
`func (o *CompareFilter) UnsetCompareTo()`

UnsetCompareTo ensures that no value is present for CompareTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


