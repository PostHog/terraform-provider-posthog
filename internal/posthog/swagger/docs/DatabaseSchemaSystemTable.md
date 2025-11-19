# DatabaseSchemaSystemTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**map[string]DatabaseSchemaField**](DatabaseSchemaField.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**RowCount** | Pointer to **NullableFloat32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "system"]

## Methods

### NewDatabaseSchemaSystemTable

`func NewDatabaseSchemaSystemTable(fields map[string]DatabaseSchemaField, id string, name string, ) *DatabaseSchemaSystemTable`

NewDatabaseSchemaSystemTable instantiates a new DatabaseSchemaSystemTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaSystemTableWithDefaults

`func NewDatabaseSchemaSystemTableWithDefaults() *DatabaseSchemaSystemTable`

NewDatabaseSchemaSystemTableWithDefaults instantiates a new DatabaseSchemaSystemTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *DatabaseSchemaSystemTable) GetFields() map[string]DatabaseSchemaField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DatabaseSchemaSystemTable) GetFieldsOk() (*map[string]DatabaseSchemaField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DatabaseSchemaSystemTable) SetFields(v map[string]DatabaseSchemaField)`

SetFields sets Fields field to given value.


### GetId

`func (o *DatabaseSchemaSystemTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseSchemaSystemTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseSchemaSystemTable) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DatabaseSchemaSystemTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseSchemaSystemTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseSchemaSystemTable) SetName(v string)`

SetName sets Name field to given value.


### GetRowCount

`func (o *DatabaseSchemaSystemTable) GetRowCount() float32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *DatabaseSchemaSystemTable) GetRowCountOk() (*float32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *DatabaseSchemaSystemTable) SetRowCount(v float32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *DatabaseSchemaSystemTable) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCountNil

`func (o *DatabaseSchemaSystemTable) SetRowCountNil(b bool)`

 SetRowCountNil sets the value for RowCount to be an explicit nil

### UnsetRowCount
`func (o *DatabaseSchemaSystemTable) UnsetRowCount()`

UnsetRowCount ensures that no value is present for RowCount, not even an explicit nil
### GetType

`func (o *DatabaseSchemaSystemTable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseSchemaSystemTable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseSchemaSystemTable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DatabaseSchemaSystemTable) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


