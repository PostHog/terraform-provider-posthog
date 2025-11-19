# PaginatedClickhouseEventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]ClickhouseEvent**](ClickhouseEvent.md) |  | [optional] 

## Methods

### NewPaginatedClickhouseEventList

`func NewPaginatedClickhouseEventList() *PaginatedClickhouseEventList`

NewPaginatedClickhouseEventList instantiates a new PaginatedClickhouseEventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedClickhouseEventListWithDefaults

`func NewPaginatedClickhouseEventListWithDefaults() *PaginatedClickhouseEventList`

NewPaginatedClickhouseEventListWithDefaults instantiates a new PaginatedClickhouseEventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedClickhouseEventList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedClickhouseEventList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedClickhouseEventList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedClickhouseEventList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedClickhouseEventList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedClickhouseEventList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetResults

`func (o *PaginatedClickhouseEventList) GetResults() []ClickhouseEvent`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedClickhouseEventList) GetResultsOk() (*[]ClickhouseEvent, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedClickhouseEventList) SetResults(v []ClickhouseEvent)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedClickhouseEventList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


