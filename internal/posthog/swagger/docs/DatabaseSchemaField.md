# DatabaseSchemaField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to [**[]Breakdown1AnyOfInner**](Breakdown1AnyOfInner.md) |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**HogqlValue** | **string** |  | 
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**SchemaValid** | **bool** |  | 
**Table** | Pointer to **NullableString** |  | [optional] 
**Type** | [**DatabaseSerializedFieldType**](DatabaseSerializedFieldType.md) |  | 

## Methods

### NewDatabaseSchemaField

`func NewDatabaseSchemaField(hogqlValue string, name string, schemaValid bool, type_ DatabaseSerializedFieldType, ) *DatabaseSchemaField`

NewDatabaseSchemaField instantiates a new DatabaseSchemaField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaFieldWithDefaults

`func NewDatabaseSchemaFieldWithDefaults() *DatabaseSchemaField`

NewDatabaseSchemaFieldWithDefaults instantiates a new DatabaseSchemaField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *DatabaseSchemaField) GetChain() []Breakdown1AnyOfInner`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *DatabaseSchemaField) GetChainOk() (*[]Breakdown1AnyOfInner, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *DatabaseSchemaField) SetChain(v []Breakdown1AnyOfInner)`

SetChain sets Chain field to given value.

### HasChain

`func (o *DatabaseSchemaField) HasChain() bool`

HasChain returns a boolean if a field has been set.

### SetChainNil

`func (o *DatabaseSchemaField) SetChainNil(b bool)`

 SetChainNil sets the value for Chain to be an explicit nil

### UnsetChain
`func (o *DatabaseSchemaField) UnsetChain()`

UnsetChain ensures that no value is present for Chain, not even an explicit nil
### GetFields

`func (o *DatabaseSchemaField) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DatabaseSchemaField) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DatabaseSchemaField) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DatabaseSchemaField) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *DatabaseSchemaField) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *DatabaseSchemaField) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetHogqlValue

`func (o *DatabaseSchemaField) GetHogqlValue() string`

GetHogqlValue returns the HogqlValue field if non-nil, zero value otherwise.

### GetHogqlValueOk

`func (o *DatabaseSchemaField) GetHogqlValueOk() (*string, bool)`

GetHogqlValueOk returns a tuple with the HogqlValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogqlValue

`func (o *DatabaseSchemaField) SetHogqlValue(v string)`

SetHogqlValue sets HogqlValue field to given value.


### GetId

`func (o *DatabaseSchemaField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseSchemaField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseSchemaField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatabaseSchemaField) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DatabaseSchemaField) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DatabaseSchemaField) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *DatabaseSchemaField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseSchemaField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseSchemaField) SetName(v string)`

SetName sets Name field to given value.


### GetSchemaValid

`func (o *DatabaseSchemaField) GetSchemaValid() bool`

GetSchemaValid returns the SchemaValid field if non-nil, zero value otherwise.

### GetSchemaValidOk

`func (o *DatabaseSchemaField) GetSchemaValidOk() (*bool, bool)`

GetSchemaValidOk returns a tuple with the SchemaValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaValid

`func (o *DatabaseSchemaField) SetSchemaValid(v bool)`

SetSchemaValid sets SchemaValid field to given value.


### GetTable

`func (o *DatabaseSchemaField) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *DatabaseSchemaField) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *DatabaseSchemaField) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *DatabaseSchemaField) HasTable() bool`

HasTable returns a boolean if a field has been set.

### SetTableNil

`func (o *DatabaseSchemaField) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *DatabaseSchemaField) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil
### GetType

`func (o *DatabaseSchemaField) GetType() DatabaseSerializedFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseSchemaField) GetTypeOk() (*DatabaseSerializedFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseSchemaField) SetType(v DatabaseSerializedFieldType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


