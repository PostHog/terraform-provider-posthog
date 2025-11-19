# QueryResponseAlternative9

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

### NewQueryResponseAlternative9

`func NewQueryResponseAlternative9(errors []HogQLNotice, notices []HogQLNotice, warnings []HogQLNotice, ) *QueryResponseAlternative9`

NewQueryResponseAlternative9 instantiates a new QueryResponseAlternative9 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternative9WithDefaults

`func NewQueryResponseAlternative9WithDefaults() *QueryResponseAlternative9`

NewQueryResponseAlternative9WithDefaults instantiates a new QueryResponseAlternative9 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChTableNames

`func (o *QueryResponseAlternative9) GetChTableNames() []string`

GetChTableNames returns the ChTableNames field if non-nil, zero value otherwise.

### GetChTableNamesOk

`func (o *QueryResponseAlternative9) GetChTableNamesOk() (*[]string, bool)`

GetChTableNamesOk returns a tuple with the ChTableNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChTableNames

`func (o *QueryResponseAlternative9) SetChTableNames(v []string)`

SetChTableNames sets ChTableNames field to given value.

### HasChTableNames

`func (o *QueryResponseAlternative9) HasChTableNames() bool`

HasChTableNames returns a boolean if a field has been set.

### SetChTableNamesNil

`func (o *QueryResponseAlternative9) SetChTableNamesNil(b bool)`

 SetChTableNamesNil sets the value for ChTableNames to be an explicit nil

### UnsetChTableNames
`func (o *QueryResponseAlternative9) UnsetChTableNames()`

UnsetChTableNames ensures that no value is present for ChTableNames, not even an explicit nil
### GetErrors

`func (o *QueryResponseAlternative9) GetErrors() []HogQLNotice`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *QueryResponseAlternative9) GetErrorsOk() (*[]HogQLNotice, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *QueryResponseAlternative9) SetErrors(v []HogQLNotice)`

SetErrors sets Errors field to given value.


### GetIsUsingIndices

`func (o *QueryResponseAlternative9) GetIsUsingIndices() QueryIndexUsage`

GetIsUsingIndices returns the IsUsingIndices field if non-nil, zero value otherwise.

### GetIsUsingIndicesOk

`func (o *QueryResponseAlternative9) GetIsUsingIndicesOk() (*QueryIndexUsage, bool)`

GetIsUsingIndicesOk returns a tuple with the IsUsingIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingIndices

`func (o *QueryResponseAlternative9) SetIsUsingIndices(v QueryIndexUsage)`

SetIsUsingIndices sets IsUsingIndices field to given value.

### HasIsUsingIndices

`func (o *QueryResponseAlternative9) HasIsUsingIndices() bool`

HasIsUsingIndices returns a boolean if a field has been set.

### GetIsValid

`func (o *QueryResponseAlternative9) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *QueryResponseAlternative9) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *QueryResponseAlternative9) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *QueryResponseAlternative9) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### SetIsValidNil

`func (o *QueryResponseAlternative9) SetIsValidNil(b bool)`

 SetIsValidNil sets the value for IsValid to be an explicit nil

### UnsetIsValid
`func (o *QueryResponseAlternative9) UnsetIsValid()`

UnsetIsValid ensures that no value is present for IsValid, not even an explicit nil
### GetNotices

`func (o *QueryResponseAlternative9) GetNotices() []HogQLNotice`

GetNotices returns the Notices field if non-nil, zero value otherwise.

### GetNoticesOk

`func (o *QueryResponseAlternative9) GetNoticesOk() (*[]HogQLNotice, bool)`

GetNoticesOk returns a tuple with the Notices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotices

`func (o *QueryResponseAlternative9) SetNotices(v []HogQLNotice)`

SetNotices sets Notices field to given value.


### GetQuery

`func (o *QueryResponseAlternative9) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryResponseAlternative9) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryResponseAlternative9) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *QueryResponseAlternative9) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *QueryResponseAlternative9) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *QueryResponseAlternative9) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetTableNames

`func (o *QueryResponseAlternative9) GetTableNames() []string`

GetTableNames returns the TableNames field if non-nil, zero value otherwise.

### GetTableNamesOk

`func (o *QueryResponseAlternative9) GetTableNamesOk() (*[]string, bool)`

GetTableNamesOk returns a tuple with the TableNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNames

`func (o *QueryResponseAlternative9) SetTableNames(v []string)`

SetTableNames sets TableNames field to given value.

### HasTableNames

`func (o *QueryResponseAlternative9) HasTableNames() bool`

HasTableNames returns a boolean if a field has been set.

### SetTableNamesNil

`func (o *QueryResponseAlternative9) SetTableNamesNil(b bool)`

 SetTableNamesNil sets the value for TableNames to be an explicit nil

### UnsetTableNames
`func (o *QueryResponseAlternative9) UnsetTableNames()`

UnsetTableNames ensures that no value is present for TableNames, not even an explicit nil
### GetWarnings

`func (o *QueryResponseAlternative9) GetWarnings() []HogQLNotice`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *QueryResponseAlternative9) GetWarningsOk() (*[]HogQLNotice, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *QueryResponseAlternative9) SetWarnings(v []HogQLNotice)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


