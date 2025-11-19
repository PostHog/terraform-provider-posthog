# DatabaseSchemaSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**LastSyncedAt** | Pointer to **NullableString** |  | [optional] 
**Prefix** | **string** |  | 
**SourceType** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewDatabaseSchemaSource

`func NewDatabaseSchemaSource(id string, prefix string, sourceType string, status string, ) *DatabaseSchemaSource`

NewDatabaseSchemaSource instantiates a new DatabaseSchemaSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaSourceWithDefaults

`func NewDatabaseSchemaSourceWithDefaults() *DatabaseSchemaSource`

NewDatabaseSchemaSourceWithDefaults instantiates a new DatabaseSchemaSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseSchemaSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseSchemaSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseSchemaSource) SetId(v string)`

SetId sets Id field to given value.


### GetLastSyncedAt

`func (o *DatabaseSchemaSource) GetLastSyncedAt() string`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *DatabaseSchemaSource) GetLastSyncedAtOk() (*string, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *DatabaseSchemaSource) SetLastSyncedAt(v string)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *DatabaseSchemaSource) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### SetLastSyncedAtNil

`func (o *DatabaseSchemaSource) SetLastSyncedAtNil(b bool)`

 SetLastSyncedAtNil sets the value for LastSyncedAt to be an explicit nil

### UnsetLastSyncedAt
`func (o *DatabaseSchemaSource) UnsetLastSyncedAt()`

UnsetLastSyncedAt ensures that no value is present for LastSyncedAt, not even an explicit nil
### GetPrefix

`func (o *DatabaseSchemaSource) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DatabaseSchemaSource) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DatabaseSchemaSource) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetSourceType

`func (o *DatabaseSchemaSource) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *DatabaseSchemaSource) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *DatabaseSchemaSource) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetStatus

`func (o *DatabaseSchemaSource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseSchemaSource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseSchemaSource) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


