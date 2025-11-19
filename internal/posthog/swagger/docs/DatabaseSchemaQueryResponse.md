# DatabaseSchemaQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Joins** | [**[]DataWarehouseViewLink**](DataWarehouseViewLink.md) |  | 
**Tables** | [**map[string]TablesValue**](TablesValue.md) |  | 

## Methods

### NewDatabaseSchemaQueryResponse

`func NewDatabaseSchemaQueryResponse(joins []DataWarehouseViewLink, tables map[string]TablesValue, ) *DatabaseSchemaQueryResponse`

NewDatabaseSchemaQueryResponse instantiates a new DatabaseSchemaQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaQueryResponseWithDefaults

`func NewDatabaseSchemaQueryResponseWithDefaults() *DatabaseSchemaQueryResponse`

NewDatabaseSchemaQueryResponseWithDefaults instantiates a new DatabaseSchemaQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoins

`func (o *DatabaseSchemaQueryResponse) GetJoins() []DataWarehouseViewLink`

GetJoins returns the Joins field if non-nil, zero value otherwise.

### GetJoinsOk

`func (o *DatabaseSchemaQueryResponse) GetJoinsOk() (*[]DataWarehouseViewLink, bool)`

GetJoinsOk returns a tuple with the Joins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoins

`func (o *DatabaseSchemaQueryResponse) SetJoins(v []DataWarehouseViewLink)`

SetJoins sets Joins field to given value.


### GetTables

`func (o *DatabaseSchemaQueryResponse) GetTables() map[string]TablesValue`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *DatabaseSchemaQueryResponse) GetTablesOk() (*map[string]TablesValue, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *DatabaseSchemaQueryResponse) SetTables(v map[string]TablesValue)`

SetTables sets Tables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


