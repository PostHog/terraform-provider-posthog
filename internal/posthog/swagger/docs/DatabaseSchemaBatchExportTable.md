# DatabaseSchemaBatchExportTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**map[string]DatabaseSchemaField**](DatabaseSchemaField.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**RowCount** | Pointer to **NullableFloat32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "batch_export"]

## Methods

### NewDatabaseSchemaBatchExportTable

`func NewDatabaseSchemaBatchExportTable(fields map[string]DatabaseSchemaField, id string, name string, ) *DatabaseSchemaBatchExportTable`

NewDatabaseSchemaBatchExportTable instantiates a new DatabaseSchemaBatchExportTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaBatchExportTableWithDefaults

`func NewDatabaseSchemaBatchExportTableWithDefaults() *DatabaseSchemaBatchExportTable`

NewDatabaseSchemaBatchExportTableWithDefaults instantiates a new DatabaseSchemaBatchExportTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *DatabaseSchemaBatchExportTable) GetFields() map[string]DatabaseSchemaField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DatabaseSchemaBatchExportTable) GetFieldsOk() (*map[string]DatabaseSchemaField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DatabaseSchemaBatchExportTable) SetFields(v map[string]DatabaseSchemaField)`

SetFields sets Fields field to given value.


### GetId

`func (o *DatabaseSchemaBatchExportTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseSchemaBatchExportTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseSchemaBatchExportTable) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DatabaseSchemaBatchExportTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseSchemaBatchExportTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseSchemaBatchExportTable) SetName(v string)`

SetName sets Name field to given value.


### GetRowCount

`func (o *DatabaseSchemaBatchExportTable) GetRowCount() float32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *DatabaseSchemaBatchExportTable) GetRowCountOk() (*float32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *DatabaseSchemaBatchExportTable) SetRowCount(v float32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *DatabaseSchemaBatchExportTable) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCountNil

`func (o *DatabaseSchemaBatchExportTable) SetRowCountNil(b bool)`

 SetRowCountNil sets the value for RowCount to be an explicit nil

### UnsetRowCount
`func (o *DatabaseSchemaBatchExportTable) UnsetRowCount()`

UnsetRowCount ensures that no value is present for RowCount, not even an explicit nil
### GetType

`func (o *DatabaseSchemaBatchExportTable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseSchemaBatchExportTable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseSchemaBatchExportTable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DatabaseSchemaBatchExportTable) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


