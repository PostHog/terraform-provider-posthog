# PaginatedBatchExportBackfillList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]BatchExportBackfill**](BatchExportBackfill.md) |  | 

## Methods

### NewPaginatedBatchExportBackfillList

`func NewPaginatedBatchExportBackfillList(results []BatchExportBackfill, ) *PaginatedBatchExportBackfillList`

NewPaginatedBatchExportBackfillList instantiates a new PaginatedBatchExportBackfillList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedBatchExportBackfillListWithDefaults

`func NewPaginatedBatchExportBackfillListWithDefaults() *PaginatedBatchExportBackfillList`

NewPaginatedBatchExportBackfillListWithDefaults instantiates a new PaginatedBatchExportBackfillList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedBatchExportBackfillList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedBatchExportBackfillList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedBatchExportBackfillList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedBatchExportBackfillList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedBatchExportBackfillList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedBatchExportBackfillList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedBatchExportBackfillList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedBatchExportBackfillList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedBatchExportBackfillList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedBatchExportBackfillList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedBatchExportBackfillList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedBatchExportBackfillList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedBatchExportBackfillList) GetResults() []BatchExportBackfill`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedBatchExportBackfillList) GetResultsOk() (*[]BatchExportBackfill, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedBatchExportBackfillList) SetResults(v []BatchExportBackfill)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


