# DatabaseSchemaViewTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**map[string]DatabaseSchemaField**](DatabaseSchemaField.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Query** | [**HogQLQuery**](HogQLQuery.md) |  | 
**RowCount** | Pointer to **NullableFloat32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "view"]

## Methods

### NewDatabaseSchemaViewTable

`func NewDatabaseSchemaViewTable(fields map[string]DatabaseSchemaField, id string, name string, query HogQLQuery, ) *DatabaseSchemaViewTable`

NewDatabaseSchemaViewTable instantiates a new DatabaseSchemaViewTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaViewTableWithDefaults

`func NewDatabaseSchemaViewTableWithDefaults() *DatabaseSchemaViewTable`

NewDatabaseSchemaViewTableWithDefaults instantiates a new DatabaseSchemaViewTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *DatabaseSchemaViewTable) GetFields() map[string]DatabaseSchemaField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DatabaseSchemaViewTable) GetFieldsOk() (*map[string]DatabaseSchemaField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DatabaseSchemaViewTable) SetFields(v map[string]DatabaseSchemaField)`

SetFields sets Fields field to given value.


### GetId

`func (o *DatabaseSchemaViewTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseSchemaViewTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseSchemaViewTable) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DatabaseSchemaViewTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseSchemaViewTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseSchemaViewTable) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *DatabaseSchemaViewTable) GetQuery() HogQLQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DatabaseSchemaViewTable) GetQueryOk() (*HogQLQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DatabaseSchemaViewTable) SetQuery(v HogQLQuery)`

SetQuery sets Query field to given value.


### GetRowCount

`func (o *DatabaseSchemaViewTable) GetRowCount() float32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *DatabaseSchemaViewTable) GetRowCountOk() (*float32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *DatabaseSchemaViewTable) SetRowCount(v float32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *DatabaseSchemaViewTable) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCountNil

`func (o *DatabaseSchemaViewTable) SetRowCountNil(b bool)`

 SetRowCountNil sets the value for RowCount to be an explicit nil

### UnsetRowCount
`func (o *DatabaseSchemaViewTable) UnsetRowCount()`

UnsetRowCount ensures that no value is present for RowCount, not even an explicit nil
### GetType

`func (o *DatabaseSchemaViewTable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseSchemaViewTable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseSchemaViewTable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DatabaseSchemaViewTable) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


