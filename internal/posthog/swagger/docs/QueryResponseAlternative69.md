# QueryResponseAlternative69

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Joins** | [**[]DataWarehouseViewLink**](DataWarehouseViewLink.md) |  | 
**Tables** | [**map[string]TablesValue**](TablesValue.md) |  | 

## Methods

### NewQueryResponseAlternative69

`func NewQueryResponseAlternative69(joins []DataWarehouseViewLink, tables map[string]TablesValue, ) *QueryResponseAlternative69`

NewQueryResponseAlternative69 instantiates a new QueryResponseAlternative69 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternative69WithDefaults

`func NewQueryResponseAlternative69WithDefaults() *QueryResponseAlternative69`

NewQueryResponseAlternative69WithDefaults instantiates a new QueryResponseAlternative69 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoins

`func (o *QueryResponseAlternative69) GetJoins() []DataWarehouseViewLink`

GetJoins returns the Joins field if non-nil, zero value otherwise.

### GetJoinsOk

`func (o *QueryResponseAlternative69) GetJoinsOk() (*[]DataWarehouseViewLink, bool)`

GetJoinsOk returns a tuple with the Joins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoins

`func (o *QueryResponseAlternative69) SetJoins(v []DataWarehouseViewLink)`

SetJoins sets Joins field to given value.


### GetTables

`func (o *QueryResponseAlternative69) GetTables() map[string]TablesValue`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *QueryResponseAlternative69) GetTablesOk() (*map[string]TablesValue, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *QueryResponseAlternative69) SetTables(v map[string]TablesValue)`

SetTables sets Tables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


