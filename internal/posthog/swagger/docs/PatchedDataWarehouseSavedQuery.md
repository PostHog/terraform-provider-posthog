# PatchedDataWarehouseSavedQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **interface{}** | HogQL query | [optional] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**SyncFrequency** | Pointer to **string** |  | [optional] [readonly] 
**Columns** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **NullableString** | The status of when this SavedQuery last ran.  * &#x60;Cancelled&#x60; - Cancelled * &#x60;Modified&#x60; - Modified * &#x60;Completed&#x60; - Completed * &#x60;Failed&#x60; - Failed * &#x60;Running&#x60; - Running | [optional] [readonly] 
**LastRunAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ManagedViewsetKind** | Pointer to **string** |  | [optional] [readonly] 
**LatestError** | Pointer to **NullableString** |  | [optional] [readonly] 
**EditedHistoryId** | Pointer to **NullableString** |  | [optional] 
**LatestHistoryId** | Pointer to **string** |  | [optional] [readonly] 
**SoftUpdate** | Pointer to **NullableBool** |  | [optional] 
**IsMaterialized** | Pointer to **NullableBool** |  | [optional] [readonly] 

## Methods

### NewPatchedDataWarehouseSavedQuery

`func NewPatchedDataWarehouseSavedQuery() *PatchedDataWarehouseSavedQuery`

NewPatchedDataWarehouseSavedQuery instantiates a new PatchedDataWarehouseSavedQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDataWarehouseSavedQueryWithDefaults

`func NewPatchedDataWarehouseSavedQueryWithDefaults() *PatchedDataWarehouseSavedQuery`

NewPatchedDataWarehouseSavedQueryWithDefaults instantiates a new PatchedDataWarehouseSavedQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDataWarehouseSavedQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDataWarehouseSavedQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDataWarehouseSavedQuery) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDataWarehouseSavedQuery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedDataWarehouseSavedQuery) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedDataWarehouseSavedQuery) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedDataWarehouseSavedQuery) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedDataWarehouseSavedQuery) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *PatchedDataWarehouseSavedQuery) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *PatchedDataWarehouseSavedQuery) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetName

`func (o *PatchedDataWarehouseSavedQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDataWarehouseSavedQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDataWarehouseSavedQuery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDataWarehouseSavedQuery) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuery

`func (o *PatchedDataWarehouseSavedQuery) GetQuery() interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PatchedDataWarehouseSavedQuery) GetQueryOk() (*interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PatchedDataWarehouseSavedQuery) SetQuery(v interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *PatchedDataWarehouseSavedQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *PatchedDataWarehouseSavedQuery) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *PatchedDataWarehouseSavedQuery) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetCreatedBy

`func (o *PatchedDataWarehouseSavedQuery) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedDataWarehouseSavedQuery) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedDataWarehouseSavedQuery) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedDataWarehouseSavedQuery) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedDataWarehouseSavedQuery) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedDataWarehouseSavedQuery) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedDataWarehouseSavedQuery) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedDataWarehouseSavedQuery) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSyncFrequency

`func (o *PatchedDataWarehouseSavedQuery) GetSyncFrequency() string`

GetSyncFrequency returns the SyncFrequency field if non-nil, zero value otherwise.

### GetSyncFrequencyOk

`func (o *PatchedDataWarehouseSavedQuery) GetSyncFrequencyOk() (*string, bool)`

GetSyncFrequencyOk returns a tuple with the SyncFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFrequency

`func (o *PatchedDataWarehouseSavedQuery) SetSyncFrequency(v string)`

SetSyncFrequency sets SyncFrequency field to given value.

### HasSyncFrequency

`func (o *PatchedDataWarehouseSavedQuery) HasSyncFrequency() bool`

HasSyncFrequency returns a boolean if a field has been set.

### GetColumns

`func (o *PatchedDataWarehouseSavedQuery) GetColumns() string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *PatchedDataWarehouseSavedQuery) GetColumnsOk() (*string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *PatchedDataWarehouseSavedQuery) SetColumns(v string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *PatchedDataWarehouseSavedQuery) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedDataWarehouseSavedQuery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedDataWarehouseSavedQuery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedDataWarehouseSavedQuery) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedDataWarehouseSavedQuery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PatchedDataWarehouseSavedQuery) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PatchedDataWarehouseSavedQuery) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastRunAt

`func (o *PatchedDataWarehouseSavedQuery) GetLastRunAt() time.Time`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *PatchedDataWarehouseSavedQuery) GetLastRunAtOk() (*time.Time, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *PatchedDataWarehouseSavedQuery) SetLastRunAt(v time.Time)`

SetLastRunAt sets LastRunAt field to given value.

### HasLastRunAt

`func (o *PatchedDataWarehouseSavedQuery) HasLastRunAt() bool`

HasLastRunAt returns a boolean if a field has been set.

### SetLastRunAtNil

`func (o *PatchedDataWarehouseSavedQuery) SetLastRunAtNil(b bool)`

 SetLastRunAtNil sets the value for LastRunAt to be an explicit nil

### UnsetLastRunAt
`func (o *PatchedDataWarehouseSavedQuery) UnsetLastRunAt()`

UnsetLastRunAt ensures that no value is present for LastRunAt, not even an explicit nil
### GetManagedViewsetKind

`func (o *PatchedDataWarehouseSavedQuery) GetManagedViewsetKind() string`

GetManagedViewsetKind returns the ManagedViewsetKind field if non-nil, zero value otherwise.

### GetManagedViewsetKindOk

`func (o *PatchedDataWarehouseSavedQuery) GetManagedViewsetKindOk() (*string, bool)`

GetManagedViewsetKindOk returns a tuple with the ManagedViewsetKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedViewsetKind

`func (o *PatchedDataWarehouseSavedQuery) SetManagedViewsetKind(v string)`

SetManagedViewsetKind sets ManagedViewsetKind field to given value.

### HasManagedViewsetKind

`func (o *PatchedDataWarehouseSavedQuery) HasManagedViewsetKind() bool`

HasManagedViewsetKind returns a boolean if a field has been set.

### GetLatestError

`func (o *PatchedDataWarehouseSavedQuery) GetLatestError() string`

GetLatestError returns the LatestError field if non-nil, zero value otherwise.

### GetLatestErrorOk

`func (o *PatchedDataWarehouseSavedQuery) GetLatestErrorOk() (*string, bool)`

GetLatestErrorOk returns a tuple with the LatestError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestError

`func (o *PatchedDataWarehouseSavedQuery) SetLatestError(v string)`

SetLatestError sets LatestError field to given value.

### HasLatestError

`func (o *PatchedDataWarehouseSavedQuery) HasLatestError() bool`

HasLatestError returns a boolean if a field has been set.

### SetLatestErrorNil

`func (o *PatchedDataWarehouseSavedQuery) SetLatestErrorNil(b bool)`

 SetLatestErrorNil sets the value for LatestError to be an explicit nil

### UnsetLatestError
`func (o *PatchedDataWarehouseSavedQuery) UnsetLatestError()`

UnsetLatestError ensures that no value is present for LatestError, not even an explicit nil
### GetEditedHistoryId

`func (o *PatchedDataWarehouseSavedQuery) GetEditedHistoryId() string`

GetEditedHistoryId returns the EditedHistoryId field if non-nil, zero value otherwise.

### GetEditedHistoryIdOk

`func (o *PatchedDataWarehouseSavedQuery) GetEditedHistoryIdOk() (*string, bool)`

GetEditedHistoryIdOk returns a tuple with the EditedHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedHistoryId

`func (o *PatchedDataWarehouseSavedQuery) SetEditedHistoryId(v string)`

SetEditedHistoryId sets EditedHistoryId field to given value.

### HasEditedHistoryId

`func (o *PatchedDataWarehouseSavedQuery) HasEditedHistoryId() bool`

HasEditedHistoryId returns a boolean if a field has been set.

### SetEditedHistoryIdNil

`func (o *PatchedDataWarehouseSavedQuery) SetEditedHistoryIdNil(b bool)`

 SetEditedHistoryIdNil sets the value for EditedHistoryId to be an explicit nil

### UnsetEditedHistoryId
`func (o *PatchedDataWarehouseSavedQuery) UnsetEditedHistoryId()`

UnsetEditedHistoryId ensures that no value is present for EditedHistoryId, not even an explicit nil
### GetLatestHistoryId

`func (o *PatchedDataWarehouseSavedQuery) GetLatestHistoryId() string`

GetLatestHistoryId returns the LatestHistoryId field if non-nil, zero value otherwise.

### GetLatestHistoryIdOk

`func (o *PatchedDataWarehouseSavedQuery) GetLatestHistoryIdOk() (*string, bool)`

GetLatestHistoryIdOk returns a tuple with the LatestHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestHistoryId

`func (o *PatchedDataWarehouseSavedQuery) SetLatestHistoryId(v string)`

SetLatestHistoryId sets LatestHistoryId field to given value.

### HasLatestHistoryId

`func (o *PatchedDataWarehouseSavedQuery) HasLatestHistoryId() bool`

HasLatestHistoryId returns a boolean if a field has been set.

### GetSoftUpdate

`func (o *PatchedDataWarehouseSavedQuery) GetSoftUpdate() bool`

GetSoftUpdate returns the SoftUpdate field if non-nil, zero value otherwise.

### GetSoftUpdateOk

`func (o *PatchedDataWarehouseSavedQuery) GetSoftUpdateOk() (*bool, bool)`

GetSoftUpdateOk returns a tuple with the SoftUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftUpdate

`func (o *PatchedDataWarehouseSavedQuery) SetSoftUpdate(v bool)`

SetSoftUpdate sets SoftUpdate field to given value.

### HasSoftUpdate

`func (o *PatchedDataWarehouseSavedQuery) HasSoftUpdate() bool`

HasSoftUpdate returns a boolean if a field has been set.

### SetSoftUpdateNil

`func (o *PatchedDataWarehouseSavedQuery) SetSoftUpdateNil(b bool)`

 SetSoftUpdateNil sets the value for SoftUpdate to be an explicit nil

### UnsetSoftUpdate
`func (o *PatchedDataWarehouseSavedQuery) UnsetSoftUpdate()`

UnsetSoftUpdate ensures that no value is present for SoftUpdate, not even an explicit nil
### GetIsMaterialized

`func (o *PatchedDataWarehouseSavedQuery) GetIsMaterialized() bool`

GetIsMaterialized returns the IsMaterialized field if non-nil, zero value otherwise.

### GetIsMaterializedOk

`func (o *PatchedDataWarehouseSavedQuery) GetIsMaterializedOk() (*bool, bool)`

GetIsMaterializedOk returns a tuple with the IsMaterialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaterialized

`func (o *PatchedDataWarehouseSavedQuery) SetIsMaterialized(v bool)`

SetIsMaterialized sets IsMaterialized field to given value.

### HasIsMaterialized

`func (o *PatchedDataWarehouseSavedQuery) HasIsMaterialized() bool`

HasIsMaterialized returns a boolean if a field has been set.

### SetIsMaterializedNil

`func (o *PatchedDataWarehouseSavedQuery) SetIsMaterializedNil(b bool)`

 SetIsMaterializedNil sets the value for IsMaterialized to be an explicit nil

### UnsetIsMaterialized
`func (o *PatchedDataWarehouseSavedQuery) UnsetIsMaterialized()`

UnsetIsMaterialized ensures that no value is present for IsMaterialized, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


