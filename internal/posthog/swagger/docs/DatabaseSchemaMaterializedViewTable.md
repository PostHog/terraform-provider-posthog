# DatabaseSchemaMaterializedViewTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**map[string]DatabaseSchemaField**](DatabaseSchemaField.md) |  | 
**Id** | **string** |  | 
**LastRunAt** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Query** | [**HogQLQuery**](HogQLQuery.md) |  | 
**RowCount** | Pointer to **NullableFloat32** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "materialized_view"]

## Methods

### NewDatabaseSchemaMaterializedViewTable

`func NewDatabaseSchemaMaterializedViewTable(fields map[string]DatabaseSchemaField, id string, name string, query HogQLQuery, ) *DatabaseSchemaMaterializedViewTable`

NewDatabaseSchemaMaterializedViewTable instantiates a new DatabaseSchemaMaterializedViewTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaMaterializedViewTableWithDefaults

`func NewDatabaseSchemaMaterializedViewTableWithDefaults() *DatabaseSchemaMaterializedViewTable`

NewDatabaseSchemaMaterializedViewTableWithDefaults instantiates a new DatabaseSchemaMaterializedViewTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *DatabaseSchemaMaterializedViewTable) GetFields() map[string]DatabaseSchemaField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DatabaseSchemaMaterializedViewTable) GetFieldsOk() (*map[string]DatabaseSchemaField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DatabaseSchemaMaterializedViewTable) SetFields(v map[string]DatabaseSchemaField)`

SetFields sets Fields field to given value.


### GetId

`func (o *DatabaseSchemaMaterializedViewTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseSchemaMaterializedViewTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseSchemaMaterializedViewTable) SetId(v string)`

SetId sets Id field to given value.


### GetLastRunAt

`func (o *DatabaseSchemaMaterializedViewTable) GetLastRunAt() string`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *DatabaseSchemaMaterializedViewTable) GetLastRunAtOk() (*string, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *DatabaseSchemaMaterializedViewTable) SetLastRunAt(v string)`

SetLastRunAt sets LastRunAt field to given value.

### HasLastRunAt

`func (o *DatabaseSchemaMaterializedViewTable) HasLastRunAt() bool`

HasLastRunAt returns a boolean if a field has been set.

### SetLastRunAtNil

`func (o *DatabaseSchemaMaterializedViewTable) SetLastRunAtNil(b bool)`

 SetLastRunAtNil sets the value for LastRunAt to be an explicit nil

### UnsetLastRunAt
`func (o *DatabaseSchemaMaterializedViewTable) UnsetLastRunAt()`

UnsetLastRunAt ensures that no value is present for LastRunAt, not even an explicit nil
### GetName

`func (o *DatabaseSchemaMaterializedViewTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseSchemaMaterializedViewTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseSchemaMaterializedViewTable) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *DatabaseSchemaMaterializedViewTable) GetQuery() HogQLQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DatabaseSchemaMaterializedViewTable) GetQueryOk() (*HogQLQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DatabaseSchemaMaterializedViewTable) SetQuery(v HogQLQuery)`

SetQuery sets Query field to given value.


### GetRowCount

`func (o *DatabaseSchemaMaterializedViewTable) GetRowCount() float32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *DatabaseSchemaMaterializedViewTable) GetRowCountOk() (*float32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *DatabaseSchemaMaterializedViewTable) SetRowCount(v float32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *DatabaseSchemaMaterializedViewTable) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCountNil

`func (o *DatabaseSchemaMaterializedViewTable) SetRowCountNil(b bool)`

 SetRowCountNil sets the value for RowCount to be an explicit nil

### UnsetRowCount
`func (o *DatabaseSchemaMaterializedViewTable) UnsetRowCount()`

UnsetRowCount ensures that no value is present for RowCount, not even an explicit nil
### GetStatus

`func (o *DatabaseSchemaMaterializedViewTable) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseSchemaMaterializedViewTable) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseSchemaMaterializedViewTable) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatabaseSchemaMaterializedViewTable) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DatabaseSchemaMaterializedViewTable) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DatabaseSchemaMaterializedViewTable) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *DatabaseSchemaMaterializedViewTable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseSchemaMaterializedViewTable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseSchemaMaterializedViewTable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DatabaseSchemaMaterializedViewTable) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


