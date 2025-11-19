# DataWarehouseSavedQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**Name** | **string** |  | 
**Query** | Pointer to **interface{}** | HogQL query | [optional] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**SyncFrequency** | **string** |  | [readonly] 
**Columns** | **string** |  | [readonly] 
**Status** | **NullableString** | The status of when this SavedQuery last ran.  * &#x60;Cancelled&#x60; - Cancelled * &#x60;Modified&#x60; - Modified * &#x60;Completed&#x60; - Completed * &#x60;Failed&#x60; - Failed * &#x60;Running&#x60; - Running | [readonly] 
**LastRunAt** | **NullableTime** |  | [readonly] 
**ManagedViewsetKind** | **string** |  | [readonly] 
**LatestError** | **NullableString** |  | [readonly] 
**EditedHistoryId** | Pointer to **NullableString** |  | [optional] 
**LatestHistoryId** | **string** |  | [readonly] 
**SoftUpdate** | Pointer to **NullableBool** |  | [optional] 
**IsMaterialized** | **NullableBool** |  | [readonly] 

## Methods

### NewDataWarehouseSavedQuery

`func NewDataWarehouseSavedQuery(id string, name string, createdBy UserBasic, createdAt time.Time, syncFrequency string, columns string, status NullableString, lastRunAt NullableTime, managedViewsetKind string, latestError NullableString, latestHistoryId string, isMaterialized NullableBool, ) *DataWarehouseSavedQuery`

NewDataWarehouseSavedQuery instantiates a new DataWarehouseSavedQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWarehouseSavedQueryWithDefaults

`func NewDataWarehouseSavedQueryWithDefaults() *DataWarehouseSavedQuery`

NewDataWarehouseSavedQueryWithDefaults instantiates a new DataWarehouseSavedQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataWarehouseSavedQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataWarehouseSavedQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataWarehouseSavedQuery) SetId(v string)`

SetId sets Id field to given value.


### GetDeleted

`func (o *DataWarehouseSavedQuery) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DataWarehouseSavedQuery) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DataWarehouseSavedQuery) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DataWarehouseSavedQuery) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *DataWarehouseSavedQuery) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *DataWarehouseSavedQuery) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetName

`func (o *DataWarehouseSavedQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataWarehouseSavedQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataWarehouseSavedQuery) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *DataWarehouseSavedQuery) GetQuery() interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DataWarehouseSavedQuery) GetQueryOk() (*interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DataWarehouseSavedQuery) SetQuery(v interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DataWarehouseSavedQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *DataWarehouseSavedQuery) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *DataWarehouseSavedQuery) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetCreatedBy

`func (o *DataWarehouseSavedQuery) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DataWarehouseSavedQuery) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DataWarehouseSavedQuery) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *DataWarehouseSavedQuery) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataWarehouseSavedQuery) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataWarehouseSavedQuery) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSyncFrequency

`func (o *DataWarehouseSavedQuery) GetSyncFrequency() string`

GetSyncFrequency returns the SyncFrequency field if non-nil, zero value otherwise.

### GetSyncFrequencyOk

`func (o *DataWarehouseSavedQuery) GetSyncFrequencyOk() (*string, bool)`

GetSyncFrequencyOk returns a tuple with the SyncFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFrequency

`func (o *DataWarehouseSavedQuery) SetSyncFrequency(v string)`

SetSyncFrequency sets SyncFrequency field to given value.


### GetColumns

`func (o *DataWarehouseSavedQuery) GetColumns() string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *DataWarehouseSavedQuery) GetColumnsOk() (*string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *DataWarehouseSavedQuery) SetColumns(v string)`

SetColumns sets Columns field to given value.


### GetStatus

`func (o *DataWarehouseSavedQuery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataWarehouseSavedQuery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataWarehouseSavedQuery) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *DataWarehouseSavedQuery) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DataWarehouseSavedQuery) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastRunAt

`func (o *DataWarehouseSavedQuery) GetLastRunAt() time.Time`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *DataWarehouseSavedQuery) GetLastRunAtOk() (*time.Time, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *DataWarehouseSavedQuery) SetLastRunAt(v time.Time)`

SetLastRunAt sets LastRunAt field to given value.


### SetLastRunAtNil

`func (o *DataWarehouseSavedQuery) SetLastRunAtNil(b bool)`

 SetLastRunAtNil sets the value for LastRunAt to be an explicit nil

### UnsetLastRunAt
`func (o *DataWarehouseSavedQuery) UnsetLastRunAt()`

UnsetLastRunAt ensures that no value is present for LastRunAt, not even an explicit nil
### GetManagedViewsetKind

`func (o *DataWarehouseSavedQuery) GetManagedViewsetKind() string`

GetManagedViewsetKind returns the ManagedViewsetKind field if non-nil, zero value otherwise.

### GetManagedViewsetKindOk

`func (o *DataWarehouseSavedQuery) GetManagedViewsetKindOk() (*string, bool)`

GetManagedViewsetKindOk returns a tuple with the ManagedViewsetKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedViewsetKind

`func (o *DataWarehouseSavedQuery) SetManagedViewsetKind(v string)`

SetManagedViewsetKind sets ManagedViewsetKind field to given value.


### GetLatestError

`func (o *DataWarehouseSavedQuery) GetLatestError() string`

GetLatestError returns the LatestError field if non-nil, zero value otherwise.

### GetLatestErrorOk

`func (o *DataWarehouseSavedQuery) GetLatestErrorOk() (*string, bool)`

GetLatestErrorOk returns a tuple with the LatestError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestError

`func (o *DataWarehouseSavedQuery) SetLatestError(v string)`

SetLatestError sets LatestError field to given value.


### SetLatestErrorNil

`func (o *DataWarehouseSavedQuery) SetLatestErrorNil(b bool)`

 SetLatestErrorNil sets the value for LatestError to be an explicit nil

### UnsetLatestError
`func (o *DataWarehouseSavedQuery) UnsetLatestError()`

UnsetLatestError ensures that no value is present for LatestError, not even an explicit nil
### GetEditedHistoryId

`func (o *DataWarehouseSavedQuery) GetEditedHistoryId() string`

GetEditedHistoryId returns the EditedHistoryId field if non-nil, zero value otherwise.

### GetEditedHistoryIdOk

`func (o *DataWarehouseSavedQuery) GetEditedHistoryIdOk() (*string, bool)`

GetEditedHistoryIdOk returns a tuple with the EditedHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedHistoryId

`func (o *DataWarehouseSavedQuery) SetEditedHistoryId(v string)`

SetEditedHistoryId sets EditedHistoryId field to given value.

### HasEditedHistoryId

`func (o *DataWarehouseSavedQuery) HasEditedHistoryId() bool`

HasEditedHistoryId returns a boolean if a field has been set.

### SetEditedHistoryIdNil

`func (o *DataWarehouseSavedQuery) SetEditedHistoryIdNil(b bool)`

 SetEditedHistoryIdNil sets the value for EditedHistoryId to be an explicit nil

### UnsetEditedHistoryId
`func (o *DataWarehouseSavedQuery) UnsetEditedHistoryId()`

UnsetEditedHistoryId ensures that no value is present for EditedHistoryId, not even an explicit nil
### GetLatestHistoryId

`func (o *DataWarehouseSavedQuery) GetLatestHistoryId() string`

GetLatestHistoryId returns the LatestHistoryId field if non-nil, zero value otherwise.

### GetLatestHistoryIdOk

`func (o *DataWarehouseSavedQuery) GetLatestHistoryIdOk() (*string, bool)`

GetLatestHistoryIdOk returns a tuple with the LatestHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestHistoryId

`func (o *DataWarehouseSavedQuery) SetLatestHistoryId(v string)`

SetLatestHistoryId sets LatestHistoryId field to given value.


### GetSoftUpdate

`func (o *DataWarehouseSavedQuery) GetSoftUpdate() bool`

GetSoftUpdate returns the SoftUpdate field if non-nil, zero value otherwise.

### GetSoftUpdateOk

`func (o *DataWarehouseSavedQuery) GetSoftUpdateOk() (*bool, bool)`

GetSoftUpdateOk returns a tuple with the SoftUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftUpdate

`func (o *DataWarehouseSavedQuery) SetSoftUpdate(v bool)`

SetSoftUpdate sets SoftUpdate field to given value.

### HasSoftUpdate

`func (o *DataWarehouseSavedQuery) HasSoftUpdate() bool`

HasSoftUpdate returns a boolean if a field has been set.

### SetSoftUpdateNil

`func (o *DataWarehouseSavedQuery) SetSoftUpdateNil(b bool)`

 SetSoftUpdateNil sets the value for SoftUpdate to be an explicit nil

### UnsetSoftUpdate
`func (o *DataWarehouseSavedQuery) UnsetSoftUpdate()`

UnsetSoftUpdate ensures that no value is present for SoftUpdate, not even an explicit nil
### GetIsMaterialized

`func (o *DataWarehouseSavedQuery) GetIsMaterialized() bool`

GetIsMaterialized returns the IsMaterialized field if non-nil, zero value otherwise.

### GetIsMaterializedOk

`func (o *DataWarehouseSavedQuery) GetIsMaterializedOk() (*bool, bool)`

GetIsMaterializedOk returns a tuple with the IsMaterialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaterialized

`func (o *DataWarehouseSavedQuery) SetIsMaterialized(v bool)`

SetIsMaterialized sets IsMaterialized field to given value.


### SetIsMaterializedNil

`func (o *DataWarehouseSavedQuery) SetIsMaterializedNil(b bool)`

 SetIsMaterializedNil sets the value for IsMaterialized to be an explicit nil

### UnsetIsMaterialized
`func (o *DataWarehouseSavedQuery) UnsetIsMaterialized()`

UnsetIsMaterialized ensures that no value is present for IsMaterialized, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


