# DatabaseSchemaDataWarehouseTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**map[string]DatabaseSchemaField**](DatabaseSchemaField.md) |  | 
**Format** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**RowCount** | Pointer to **NullableFloat32** |  | [optional] 
**Schema** | Pointer to [**DatabaseSchemaSchema**](DatabaseSchemaSchema.md) |  | [optional] 
**Source** | Pointer to [**DatabaseSchemaSource**](DatabaseSchemaSource.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "data_warehouse"]
**UrlPattern** | **string** |  | 

## Methods

### NewDatabaseSchemaDataWarehouseTable

`func NewDatabaseSchemaDataWarehouseTable(fields map[string]DatabaseSchemaField, format string, id string, name string, urlPattern string, ) *DatabaseSchemaDataWarehouseTable`

NewDatabaseSchemaDataWarehouseTable instantiates a new DatabaseSchemaDataWarehouseTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaDataWarehouseTableWithDefaults

`func NewDatabaseSchemaDataWarehouseTableWithDefaults() *DatabaseSchemaDataWarehouseTable`

NewDatabaseSchemaDataWarehouseTableWithDefaults instantiates a new DatabaseSchemaDataWarehouseTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *DatabaseSchemaDataWarehouseTable) GetFields() map[string]DatabaseSchemaField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DatabaseSchemaDataWarehouseTable) GetFieldsOk() (*map[string]DatabaseSchemaField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DatabaseSchemaDataWarehouseTable) SetFields(v map[string]DatabaseSchemaField)`

SetFields sets Fields field to given value.


### GetFormat

`func (o *DatabaseSchemaDataWarehouseTable) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DatabaseSchemaDataWarehouseTable) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DatabaseSchemaDataWarehouseTable) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetId

`func (o *DatabaseSchemaDataWarehouseTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseSchemaDataWarehouseTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseSchemaDataWarehouseTable) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DatabaseSchemaDataWarehouseTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseSchemaDataWarehouseTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseSchemaDataWarehouseTable) SetName(v string)`

SetName sets Name field to given value.


### GetRowCount

`func (o *DatabaseSchemaDataWarehouseTable) GetRowCount() float32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *DatabaseSchemaDataWarehouseTable) GetRowCountOk() (*float32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *DatabaseSchemaDataWarehouseTable) SetRowCount(v float32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *DatabaseSchemaDataWarehouseTable) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCountNil

`func (o *DatabaseSchemaDataWarehouseTable) SetRowCountNil(b bool)`

 SetRowCountNil sets the value for RowCount to be an explicit nil

### UnsetRowCount
`func (o *DatabaseSchemaDataWarehouseTable) UnsetRowCount()`

UnsetRowCount ensures that no value is present for RowCount, not even an explicit nil
### GetSchema

`func (o *DatabaseSchemaDataWarehouseTable) GetSchema() DatabaseSchemaSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *DatabaseSchemaDataWarehouseTable) GetSchemaOk() (*DatabaseSchemaSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *DatabaseSchemaDataWarehouseTable) SetSchema(v DatabaseSchemaSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *DatabaseSchemaDataWarehouseTable) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSource

`func (o *DatabaseSchemaDataWarehouseTable) GetSource() DatabaseSchemaSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DatabaseSchemaDataWarehouseTable) GetSourceOk() (*DatabaseSchemaSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DatabaseSchemaDataWarehouseTable) SetSource(v DatabaseSchemaSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *DatabaseSchemaDataWarehouseTable) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *DatabaseSchemaDataWarehouseTable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseSchemaDataWarehouseTable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseSchemaDataWarehouseTable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DatabaseSchemaDataWarehouseTable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrlPattern

`func (o *DatabaseSchemaDataWarehouseTable) GetUrlPattern() string`

GetUrlPattern returns the UrlPattern field if non-nil, zero value otherwise.

### GetUrlPatternOk

`func (o *DatabaseSchemaDataWarehouseTable) GetUrlPatternOk() (*string, bool)`

GetUrlPatternOk returns a tuple with the UrlPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPattern

`func (o *DatabaseSchemaDataWarehouseTable) SetUrlPattern(v string)`

SetUrlPattern sets UrlPattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


