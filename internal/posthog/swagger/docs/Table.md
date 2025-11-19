# Table

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**Name** | **string** |  | 
**Format** | [**TableFormatEnum**](TableFormatEnum.md) |  | 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UrlPattern** | **string** |  | 
**Credential** | [**Credential**](Credential.md) |  | 
**Columns** | **string** |  | [readonly] 
**ExternalDataSource** | [**SimpleExternalDataSourceSerializers**](SimpleExternalDataSourceSerializers.md) |  | [readonly] 
**ExternalSchema** | **string** |  | [readonly] 

## Methods

### NewTable

`func NewTable(id string, name string, format TableFormatEnum, createdBy UserBasic, createdAt time.Time, urlPattern string, credential Credential, columns string, externalDataSource SimpleExternalDataSourceSerializers, externalSchema string, ) *Table`

NewTable instantiates a new Table object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableWithDefaults

`func NewTableWithDefaults() *Table`

NewTableWithDefaults instantiates a new Table object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Table) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Table) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Table) SetId(v string)`

SetId sets Id field to given value.


### GetDeleted

`func (o *Table) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Table) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Table) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Table) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *Table) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *Table) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetName

`func (o *Table) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Table) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Table) SetName(v string)`

SetName sets Name field to given value.


### GetFormat

`func (o *Table) GetFormat() TableFormatEnum`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Table) GetFormatOk() (*TableFormatEnum, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Table) SetFormat(v TableFormatEnum)`

SetFormat sets Format field to given value.


### GetCreatedBy

`func (o *Table) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Table) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Table) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *Table) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Table) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Table) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUrlPattern

`func (o *Table) GetUrlPattern() string`

GetUrlPattern returns the UrlPattern field if non-nil, zero value otherwise.

### GetUrlPatternOk

`func (o *Table) GetUrlPatternOk() (*string, bool)`

GetUrlPatternOk returns a tuple with the UrlPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPattern

`func (o *Table) SetUrlPattern(v string)`

SetUrlPattern sets UrlPattern field to given value.


### GetCredential

`func (o *Table) GetCredential() Credential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *Table) GetCredentialOk() (*Credential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *Table) SetCredential(v Credential)`

SetCredential sets Credential field to given value.


### GetColumns

`func (o *Table) GetColumns() string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Table) GetColumnsOk() (*string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Table) SetColumns(v string)`

SetColumns sets Columns field to given value.


### GetExternalDataSource

`func (o *Table) GetExternalDataSource() SimpleExternalDataSourceSerializers`

GetExternalDataSource returns the ExternalDataSource field if non-nil, zero value otherwise.

### GetExternalDataSourceOk

`func (o *Table) GetExternalDataSourceOk() (*SimpleExternalDataSourceSerializers, bool)`

GetExternalDataSourceOk returns a tuple with the ExternalDataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDataSource

`func (o *Table) SetExternalDataSource(v SimpleExternalDataSourceSerializers)`

SetExternalDataSource sets ExternalDataSource field to given value.


### GetExternalSchema

`func (o *Table) GetExternalSchema() string`

GetExternalSchema returns the ExternalSchema field if non-nil, zero value otherwise.

### GetExternalSchemaOk

`func (o *Table) GetExternalSchemaOk() (*string, bool)`

GetExternalSchemaOk returns a tuple with the ExternalSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSchema

`func (o *Table) SetExternalSchema(v string)`

SetExternalSchema sets ExternalSchema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


