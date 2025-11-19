# DatabaseSchemaSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Incremental** | **bool** |  | 
**LastSyncedAt** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**ShouldSync** | **bool** |  | 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDatabaseSchemaSchema

`func NewDatabaseSchemaSchema(id string, incremental bool, name string, shouldSync bool, ) *DatabaseSchemaSchema`

NewDatabaseSchemaSchema instantiates a new DatabaseSchemaSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaSchemaWithDefaults

`func NewDatabaseSchemaSchemaWithDefaults() *DatabaseSchemaSchema`

NewDatabaseSchemaSchemaWithDefaults instantiates a new DatabaseSchemaSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseSchemaSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseSchemaSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseSchemaSchema) SetId(v string)`

SetId sets Id field to given value.


### GetIncremental

`func (o *DatabaseSchemaSchema) GetIncremental() bool`

GetIncremental returns the Incremental field if non-nil, zero value otherwise.

### GetIncrementalOk

`func (o *DatabaseSchemaSchema) GetIncrementalOk() (*bool, bool)`

GetIncrementalOk returns a tuple with the Incremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncremental

`func (o *DatabaseSchemaSchema) SetIncremental(v bool)`

SetIncremental sets Incremental field to given value.


### GetLastSyncedAt

`func (o *DatabaseSchemaSchema) GetLastSyncedAt() string`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *DatabaseSchemaSchema) GetLastSyncedAtOk() (*string, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *DatabaseSchemaSchema) SetLastSyncedAt(v string)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *DatabaseSchemaSchema) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### SetLastSyncedAtNil

`func (o *DatabaseSchemaSchema) SetLastSyncedAtNil(b bool)`

 SetLastSyncedAtNil sets the value for LastSyncedAt to be an explicit nil

### UnsetLastSyncedAt
`func (o *DatabaseSchemaSchema) UnsetLastSyncedAt()`

UnsetLastSyncedAt ensures that no value is present for LastSyncedAt, not even an explicit nil
### GetName

`func (o *DatabaseSchemaSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseSchemaSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseSchemaSchema) SetName(v string)`

SetName sets Name field to given value.


### GetShouldSync

`func (o *DatabaseSchemaSchema) GetShouldSync() bool`

GetShouldSync returns the ShouldSync field if non-nil, zero value otherwise.

### GetShouldSyncOk

`func (o *DatabaseSchemaSchema) GetShouldSyncOk() (*bool, bool)`

GetShouldSyncOk returns a tuple with the ShouldSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSync

`func (o *DatabaseSchemaSchema) SetShouldSync(v bool)`

SetShouldSync sets ShouldSync field to given value.


### GetStatus

`func (o *DatabaseSchemaSchema) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseSchemaSchema) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseSchemaSchema) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatabaseSchemaSchema) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DatabaseSchemaSchema) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DatabaseSchemaSchema) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


