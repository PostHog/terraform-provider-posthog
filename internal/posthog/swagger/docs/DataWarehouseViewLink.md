# DataWarehouseViewLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**DataWarehouseViewLinkConfiguration**](DataWarehouseViewLinkConfiguration.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**UserBasicType**](UserBasicType.md) |  | [optional] 
**FieldName** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**JoiningTableKey** | Pointer to **NullableString** |  | [optional] 
**JoiningTableName** | Pointer to **NullableString** |  | [optional] 
**SourceTableKey** | Pointer to **NullableString** |  | [optional] 
**SourceTableName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDataWarehouseViewLink

`func NewDataWarehouseViewLink(id string, ) *DataWarehouseViewLink`

NewDataWarehouseViewLink instantiates a new DataWarehouseViewLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWarehouseViewLinkWithDefaults

`func NewDataWarehouseViewLinkWithDefaults() *DataWarehouseViewLink`

NewDataWarehouseViewLinkWithDefaults instantiates a new DataWarehouseViewLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *DataWarehouseViewLink) GetConfiguration() DataWarehouseViewLinkConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DataWarehouseViewLink) GetConfigurationOk() (*DataWarehouseViewLinkConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DataWarehouseViewLink) SetConfiguration(v DataWarehouseViewLinkConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DataWarehouseViewLink) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DataWarehouseViewLink) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataWarehouseViewLink) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataWarehouseViewLink) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataWarehouseViewLink) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *DataWarehouseViewLink) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DataWarehouseViewLink) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *DataWarehouseViewLink) GetCreatedBy() UserBasicType`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DataWarehouseViewLink) GetCreatedByOk() (*UserBasicType, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DataWarehouseViewLink) SetCreatedBy(v UserBasicType)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DataWarehouseViewLink) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetFieldName

`func (o *DataWarehouseViewLink) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *DataWarehouseViewLink) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *DataWarehouseViewLink) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *DataWarehouseViewLink) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### SetFieldNameNil

`func (o *DataWarehouseViewLink) SetFieldNameNil(b bool)`

 SetFieldNameNil sets the value for FieldName to be an explicit nil

### UnsetFieldName
`func (o *DataWarehouseViewLink) UnsetFieldName()`

UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
### GetId

`func (o *DataWarehouseViewLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataWarehouseViewLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataWarehouseViewLink) SetId(v string)`

SetId sets Id field to given value.


### GetJoiningTableKey

`func (o *DataWarehouseViewLink) GetJoiningTableKey() string`

GetJoiningTableKey returns the JoiningTableKey field if non-nil, zero value otherwise.

### GetJoiningTableKeyOk

`func (o *DataWarehouseViewLink) GetJoiningTableKeyOk() (*string, bool)`

GetJoiningTableKeyOk returns a tuple with the JoiningTableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoiningTableKey

`func (o *DataWarehouseViewLink) SetJoiningTableKey(v string)`

SetJoiningTableKey sets JoiningTableKey field to given value.

### HasJoiningTableKey

`func (o *DataWarehouseViewLink) HasJoiningTableKey() bool`

HasJoiningTableKey returns a boolean if a field has been set.

### SetJoiningTableKeyNil

`func (o *DataWarehouseViewLink) SetJoiningTableKeyNil(b bool)`

 SetJoiningTableKeyNil sets the value for JoiningTableKey to be an explicit nil

### UnsetJoiningTableKey
`func (o *DataWarehouseViewLink) UnsetJoiningTableKey()`

UnsetJoiningTableKey ensures that no value is present for JoiningTableKey, not even an explicit nil
### GetJoiningTableName

`func (o *DataWarehouseViewLink) GetJoiningTableName() string`

GetJoiningTableName returns the JoiningTableName field if non-nil, zero value otherwise.

### GetJoiningTableNameOk

`func (o *DataWarehouseViewLink) GetJoiningTableNameOk() (*string, bool)`

GetJoiningTableNameOk returns a tuple with the JoiningTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoiningTableName

`func (o *DataWarehouseViewLink) SetJoiningTableName(v string)`

SetJoiningTableName sets JoiningTableName field to given value.

### HasJoiningTableName

`func (o *DataWarehouseViewLink) HasJoiningTableName() bool`

HasJoiningTableName returns a boolean if a field has been set.

### SetJoiningTableNameNil

`func (o *DataWarehouseViewLink) SetJoiningTableNameNil(b bool)`

 SetJoiningTableNameNil sets the value for JoiningTableName to be an explicit nil

### UnsetJoiningTableName
`func (o *DataWarehouseViewLink) UnsetJoiningTableName()`

UnsetJoiningTableName ensures that no value is present for JoiningTableName, not even an explicit nil
### GetSourceTableKey

`func (o *DataWarehouseViewLink) GetSourceTableKey() string`

GetSourceTableKey returns the SourceTableKey field if non-nil, zero value otherwise.

### GetSourceTableKeyOk

`func (o *DataWarehouseViewLink) GetSourceTableKeyOk() (*string, bool)`

GetSourceTableKeyOk returns a tuple with the SourceTableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTableKey

`func (o *DataWarehouseViewLink) SetSourceTableKey(v string)`

SetSourceTableKey sets SourceTableKey field to given value.

### HasSourceTableKey

`func (o *DataWarehouseViewLink) HasSourceTableKey() bool`

HasSourceTableKey returns a boolean if a field has been set.

### SetSourceTableKeyNil

`func (o *DataWarehouseViewLink) SetSourceTableKeyNil(b bool)`

 SetSourceTableKeyNil sets the value for SourceTableKey to be an explicit nil

### UnsetSourceTableKey
`func (o *DataWarehouseViewLink) UnsetSourceTableKey()`

UnsetSourceTableKey ensures that no value is present for SourceTableKey, not even an explicit nil
### GetSourceTableName

`func (o *DataWarehouseViewLink) GetSourceTableName() string`

GetSourceTableName returns the SourceTableName field if non-nil, zero value otherwise.

### GetSourceTableNameOk

`func (o *DataWarehouseViewLink) GetSourceTableNameOk() (*string, bool)`

GetSourceTableNameOk returns a tuple with the SourceTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTableName

`func (o *DataWarehouseViewLink) SetSourceTableName(v string)`

SetSourceTableName sets SourceTableName field to given value.

### HasSourceTableName

`func (o *DataWarehouseViewLink) HasSourceTableName() bool`

HasSourceTableName returns a boolean if a field has been set.

### SetSourceTableNameNil

`func (o *DataWarehouseViewLink) SetSourceTableNameNil(b bool)`

 SetSourceTableNameNil sets the value for SourceTableName to be an explicit nil

### UnsetSourceTableName
`func (o *DataWarehouseViewLink) UnsetSourceTableName()`

UnsetSourceTableName ensures that no value is present for SourceTableName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


