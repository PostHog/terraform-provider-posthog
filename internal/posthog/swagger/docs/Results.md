# Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **float32** |  | 
**Values** | [**[]BreakdownValue**](BreakdownValue.md) |  | 

## Methods

### NewResults

`func NewResults(totalCount float32, values []BreakdownValue, ) *Results`

NewResults instantiates a new Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsWithDefaults

`func NewResultsWithDefaults() *Results`

NewResultsWithDefaults instantiates a new Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *Results) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Results) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Results) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.


### GetValues

`func (o *Results) GetValues() []BreakdownValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Results) GetValuesOk() (*[]BreakdownValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Results) SetValues(v []BreakdownValue)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


