# HogQLMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChTableNames** | Pointer to **[]string** |  | [optional] 
**Errors** | [**[]HogQLNotice**](HogQLNotice.md) |  | 
**IsUsingIndices** | Pointer to [**QueryIndexUsage**](QueryIndexUsage.md) |  | [optional] 
**IsValid** | Pointer to **NullableBool** |  | [optional] 
**Notices** | [**[]HogQLNotice**](HogQLNotice.md) |  | 
**Query** | Pointer to **NullableString** |  | [optional] 
**TableNames** | Pointer to **[]string** |  | [optional] 
**Warnings** | [**[]HogQLNotice**](HogQLNotice.md) |  | 

## Methods

### NewHogQLMetadataResponse

`func NewHogQLMetadataResponse(errors []HogQLNotice, notices []HogQLNotice, warnings []HogQLNotice, ) *HogQLMetadataResponse`

NewHogQLMetadataResponse instantiates a new HogQLMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogQLMetadataResponseWithDefaults

`func NewHogQLMetadataResponseWithDefaults() *HogQLMetadataResponse`

NewHogQLMetadataResponseWithDefaults instantiates a new HogQLMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChTableNames

`func (o *HogQLMetadataResponse) GetChTableNames() []string`

GetChTableNames returns the ChTableNames field if non-nil, zero value otherwise.

### GetChTableNamesOk

`func (o *HogQLMetadataResponse) GetChTableNamesOk() (*[]string, bool)`

GetChTableNamesOk returns a tuple with the ChTableNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChTableNames

`func (o *HogQLMetadataResponse) SetChTableNames(v []string)`

SetChTableNames sets ChTableNames field to given value.

### HasChTableNames

`func (o *HogQLMetadataResponse) HasChTableNames() bool`

HasChTableNames returns a boolean if a field has been set.

### SetChTableNamesNil

`func (o *HogQLMetadataResponse) SetChTableNamesNil(b bool)`

 SetChTableNamesNil sets the value for ChTableNames to be an explicit nil

### UnsetChTableNames
`func (o *HogQLMetadataResponse) UnsetChTableNames()`

UnsetChTableNames ensures that no value is present for ChTableNames, not even an explicit nil
### GetErrors

`func (o *HogQLMetadataResponse) GetErrors() []HogQLNotice`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HogQLMetadataResponse) GetErrorsOk() (*[]HogQLNotice, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HogQLMetadataResponse) SetErrors(v []HogQLNotice)`

SetErrors sets Errors field to given value.


### GetIsUsingIndices

`func (o *HogQLMetadataResponse) GetIsUsingIndices() QueryIndexUsage`

GetIsUsingIndices returns the IsUsingIndices field if non-nil, zero value otherwise.

### GetIsUsingIndicesOk

`func (o *HogQLMetadataResponse) GetIsUsingIndicesOk() (*QueryIndexUsage, bool)`

GetIsUsingIndicesOk returns a tuple with the IsUsingIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingIndices

`func (o *HogQLMetadataResponse) SetIsUsingIndices(v QueryIndexUsage)`

SetIsUsingIndices sets IsUsingIndices field to given value.

### HasIsUsingIndices

`func (o *HogQLMetadataResponse) HasIsUsingIndices() bool`

HasIsUsingIndices returns a boolean if a field has been set.

### GetIsValid

`func (o *HogQLMetadataResponse) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *HogQLMetadataResponse) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *HogQLMetadataResponse) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *HogQLMetadataResponse) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### SetIsValidNil

`func (o *HogQLMetadataResponse) SetIsValidNil(b bool)`

 SetIsValidNil sets the value for IsValid to be an explicit nil

### UnsetIsValid
`func (o *HogQLMetadataResponse) UnsetIsValid()`

UnsetIsValid ensures that no value is present for IsValid, not even an explicit nil
### GetNotices

`func (o *HogQLMetadataResponse) GetNotices() []HogQLNotice`

GetNotices returns the Notices field if non-nil, zero value otherwise.

### GetNoticesOk

`func (o *HogQLMetadataResponse) GetNoticesOk() (*[]HogQLNotice, bool)`

GetNoticesOk returns a tuple with the Notices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotices

`func (o *HogQLMetadataResponse) SetNotices(v []HogQLNotice)`

SetNotices sets Notices field to given value.


### GetQuery

`func (o *HogQLMetadataResponse) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *HogQLMetadataResponse) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *HogQLMetadataResponse) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *HogQLMetadataResponse) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *HogQLMetadataResponse) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *HogQLMetadataResponse) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetTableNames

`func (o *HogQLMetadataResponse) GetTableNames() []string`

GetTableNames returns the TableNames field if non-nil, zero value otherwise.

### GetTableNamesOk

`func (o *HogQLMetadataResponse) GetTableNamesOk() (*[]string, bool)`

GetTableNamesOk returns a tuple with the TableNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNames

`func (o *HogQLMetadataResponse) SetTableNames(v []string)`

SetTableNames sets TableNames field to given value.

### HasTableNames

`func (o *HogQLMetadataResponse) HasTableNames() bool`

HasTableNames returns a boolean if a field has been set.

### SetTableNamesNil

`func (o *HogQLMetadataResponse) SetTableNamesNil(b bool)`

 SetTableNamesNil sets the value for TableNames to be an explicit nil

### UnsetTableNames
`func (o *HogQLMetadataResponse) UnsetTableNames()`

UnsetTableNames ensures that no value is present for TableNames, not even an explicit nil
### GetWarnings

`func (o *HogQLMetadataResponse) GetWarnings() []HogQLNotice`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *HogQLMetadataResponse) GetWarningsOk() (*[]HogQLNotice, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *HogQLMetadataResponse) SetWarnings(v []HogQLNotice)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


