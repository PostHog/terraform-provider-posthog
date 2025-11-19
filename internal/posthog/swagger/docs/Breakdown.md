# Breakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupTypeIndex** | Pointer to **NullableInt32** |  | [optional] 
**HistogramBinCount** | Pointer to **NullableInt32** |  | [optional] 
**NormalizeUrl** | Pointer to **NullableBool** |  | [optional] 
**Property** | [**Property1**](Property1.md) |  | 
**Type** | Pointer to [**MultipleBreakdownType**](MultipleBreakdownType.md) |  | [optional] 

## Methods

### NewBreakdown

`func NewBreakdown(property Property1, ) *Breakdown`

NewBreakdown instantiates a new Breakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreakdownWithDefaults

`func NewBreakdownWithDefaults() *Breakdown`

NewBreakdownWithDefaults instantiates a new Breakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupTypeIndex

`func (o *Breakdown) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *Breakdown) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *Breakdown) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.

### HasGroupTypeIndex

`func (o *Breakdown) HasGroupTypeIndex() bool`

HasGroupTypeIndex returns a boolean if a field has been set.

### SetGroupTypeIndexNil

`func (o *Breakdown) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *Breakdown) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil
### GetHistogramBinCount

`func (o *Breakdown) GetHistogramBinCount() int32`

GetHistogramBinCount returns the HistogramBinCount field if non-nil, zero value otherwise.

### GetHistogramBinCountOk

`func (o *Breakdown) GetHistogramBinCountOk() (*int32, bool)`

GetHistogramBinCountOk returns a tuple with the HistogramBinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramBinCount

`func (o *Breakdown) SetHistogramBinCount(v int32)`

SetHistogramBinCount sets HistogramBinCount field to given value.

### HasHistogramBinCount

`func (o *Breakdown) HasHistogramBinCount() bool`

HasHistogramBinCount returns a boolean if a field has been set.

### SetHistogramBinCountNil

`func (o *Breakdown) SetHistogramBinCountNil(b bool)`

 SetHistogramBinCountNil sets the value for HistogramBinCount to be an explicit nil

### UnsetHistogramBinCount
`func (o *Breakdown) UnsetHistogramBinCount()`

UnsetHistogramBinCount ensures that no value is present for HistogramBinCount, not even an explicit nil
### GetNormalizeUrl

`func (o *Breakdown) GetNormalizeUrl() bool`

GetNormalizeUrl returns the NormalizeUrl field if non-nil, zero value otherwise.

### GetNormalizeUrlOk

`func (o *Breakdown) GetNormalizeUrlOk() (*bool, bool)`

GetNormalizeUrlOk returns a tuple with the NormalizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizeUrl

`func (o *Breakdown) SetNormalizeUrl(v bool)`

SetNormalizeUrl sets NormalizeUrl field to given value.

### HasNormalizeUrl

`func (o *Breakdown) HasNormalizeUrl() bool`

HasNormalizeUrl returns a boolean if a field has been set.

### SetNormalizeUrlNil

`func (o *Breakdown) SetNormalizeUrlNil(b bool)`

 SetNormalizeUrlNil sets the value for NormalizeUrl to be an explicit nil

### UnsetNormalizeUrl
`func (o *Breakdown) UnsetNormalizeUrl()`

UnsetNormalizeUrl ensures that no value is present for NormalizeUrl, not even an explicit nil
### GetProperty

`func (o *Breakdown) GetProperty() Property1`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *Breakdown) GetPropertyOk() (*Property1, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *Breakdown) SetProperty(v Property1)`

SetProperty sets Property field to given value.


### GetType

`func (o *Breakdown) GetType() MultipleBreakdownType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Breakdown) GetTypeOk() (*MultipleBreakdownType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Breakdown) SetType(v MultipleBreakdownType)`

SetType sets Type field to given value.

### HasType

`func (o *Breakdown) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


