# PatchedTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Format** | Pointer to [**TableFormatEnum**](TableFormatEnum.md) |  | [optional] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UrlPattern** | Pointer to **string** |  | [optional] 
**Credential** | Pointer to [**Credential**](Credential.md) |  | [optional] 
**Columns** | Pointer to **string** |  | [optional] [readonly] 
**ExternalDataSource** | Pointer to [**SimpleExternalDataSourceSerializers**](SimpleExternalDataSourceSerializers.md) |  | [optional] [readonly] 
**ExternalSchema** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedTable

`func NewPatchedTable() *PatchedTable`

NewPatchedTable instantiates a new PatchedTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTableWithDefaults

`func NewPatchedTableWithDefaults() *PatchedTable`

NewPatchedTableWithDefaults instantiates a new PatchedTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedTable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedTable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedTable) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedTable) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedTable) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedTable) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *PatchedTable) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *PatchedTable) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetName

`func (o *PatchedTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormat

`func (o *PatchedTable) GetFormat() TableFormatEnum`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PatchedTable) GetFormatOk() (*TableFormatEnum, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PatchedTable) SetFormat(v TableFormatEnum)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PatchedTable) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedTable) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedTable) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedTable) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedTable) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedTable) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedTable) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedTable) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedTable) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUrlPattern

`func (o *PatchedTable) GetUrlPattern() string`

GetUrlPattern returns the UrlPattern field if non-nil, zero value otherwise.

### GetUrlPatternOk

`func (o *PatchedTable) GetUrlPatternOk() (*string, bool)`

GetUrlPatternOk returns a tuple with the UrlPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPattern

`func (o *PatchedTable) SetUrlPattern(v string)`

SetUrlPattern sets UrlPattern field to given value.

### HasUrlPattern

`func (o *PatchedTable) HasUrlPattern() bool`

HasUrlPattern returns a boolean if a field has been set.

### GetCredential

`func (o *PatchedTable) GetCredential() Credential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *PatchedTable) GetCredentialOk() (*Credential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *PatchedTable) SetCredential(v Credential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *PatchedTable) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetColumns

`func (o *PatchedTable) GetColumns() string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *PatchedTable) GetColumnsOk() (*string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *PatchedTable) SetColumns(v string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *PatchedTable) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetExternalDataSource

`func (o *PatchedTable) GetExternalDataSource() SimpleExternalDataSourceSerializers`

GetExternalDataSource returns the ExternalDataSource field if non-nil, zero value otherwise.

### GetExternalDataSourceOk

`func (o *PatchedTable) GetExternalDataSourceOk() (*SimpleExternalDataSourceSerializers, bool)`

GetExternalDataSourceOk returns a tuple with the ExternalDataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDataSource

`func (o *PatchedTable) SetExternalDataSource(v SimpleExternalDataSourceSerializers)`

SetExternalDataSource sets ExternalDataSource field to given value.

### HasExternalDataSource

`func (o *PatchedTable) HasExternalDataSource() bool`

HasExternalDataSource returns a boolean if a field has been set.

### GetExternalSchema

`func (o *PatchedTable) GetExternalSchema() string`

GetExternalSchema returns the ExternalSchema field if non-nil, zero value otherwise.

### GetExternalSchemaOk

`func (o *PatchedTable) GetExternalSchemaOk() (*string, bool)`

GetExternalSchemaOk returns a tuple with the ExternalSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSchema

`func (o *PatchedTable) SetExternalSchema(v string)`

SetExternalSchema sets ExternalSchema field to given value.

### HasExternalSchema

`func (o *PatchedTable) HasExternalSchema() bool`

HasExternalSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


