# DatabaseSchemaManagedViewTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**map[string]DatabaseSchemaField**](DatabaseSchemaField.md) |  | 
**Id** | **string** |  | 
**Kind** | [**DatabaseSchemaManagedViewTableKind**](DatabaseSchemaManagedViewTableKind.md) |  | 
**Name** | **string** |  | 
**Query** | [**HogQLQuery**](HogQLQuery.md) |  | 
**RowCount** | Pointer to **NullableFloat32** |  | [optional] 
**SourceId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "managed_view"]

## Methods

### NewDatabaseSchemaManagedViewTable

`func NewDatabaseSchemaManagedViewTable(fields map[string]DatabaseSchemaField, id string, kind DatabaseSchemaManagedViewTableKind, name string, query HogQLQuery, ) *DatabaseSchemaManagedViewTable`

NewDatabaseSchemaManagedViewTable instantiates a new DatabaseSchemaManagedViewTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaManagedViewTableWithDefaults

`func NewDatabaseSchemaManagedViewTableWithDefaults() *DatabaseSchemaManagedViewTable`

NewDatabaseSchemaManagedViewTableWithDefaults instantiates a new DatabaseSchemaManagedViewTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *DatabaseSchemaManagedViewTable) GetFields() map[string]DatabaseSchemaField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DatabaseSchemaManagedViewTable) GetFieldsOk() (*map[string]DatabaseSchemaField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DatabaseSchemaManagedViewTable) SetFields(v map[string]DatabaseSchemaField)`

SetFields sets Fields field to given value.


### GetId

`func (o *DatabaseSchemaManagedViewTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseSchemaManagedViewTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseSchemaManagedViewTable) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *DatabaseSchemaManagedViewTable) GetKind() DatabaseSchemaManagedViewTableKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DatabaseSchemaManagedViewTable) GetKindOk() (*DatabaseSchemaManagedViewTableKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DatabaseSchemaManagedViewTable) SetKind(v DatabaseSchemaManagedViewTableKind)`

SetKind sets Kind field to given value.


### GetName

`func (o *DatabaseSchemaManagedViewTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseSchemaManagedViewTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseSchemaManagedViewTable) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *DatabaseSchemaManagedViewTable) GetQuery() HogQLQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DatabaseSchemaManagedViewTable) GetQueryOk() (*HogQLQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DatabaseSchemaManagedViewTable) SetQuery(v HogQLQuery)`

SetQuery sets Query field to given value.


### GetRowCount

`func (o *DatabaseSchemaManagedViewTable) GetRowCount() float32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *DatabaseSchemaManagedViewTable) GetRowCountOk() (*float32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *DatabaseSchemaManagedViewTable) SetRowCount(v float32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *DatabaseSchemaManagedViewTable) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCountNil

`func (o *DatabaseSchemaManagedViewTable) SetRowCountNil(b bool)`

 SetRowCountNil sets the value for RowCount to be an explicit nil

### UnsetRowCount
`func (o *DatabaseSchemaManagedViewTable) UnsetRowCount()`

UnsetRowCount ensures that no value is present for RowCount, not even an explicit nil
### GetSourceId

`func (o *DatabaseSchemaManagedViewTable) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DatabaseSchemaManagedViewTable) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DatabaseSchemaManagedViewTable) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *DatabaseSchemaManagedViewTable) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *DatabaseSchemaManagedViewTable) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *DatabaseSchemaManagedViewTable) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetType

`func (o *DatabaseSchemaManagedViewTable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseSchemaManagedViewTable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseSchemaManagedViewTable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DatabaseSchemaManagedViewTable) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


