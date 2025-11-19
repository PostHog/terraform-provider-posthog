# EndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheAgeSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**IsMaterialized** | Pointer to **NullableBool** | Whether this endpoint&#39;s query results are materialized to S3 | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Query** | Pointer to [**NullableQuery**](Query.md) |  | [optional] [default to null]
**SyncFrequency** | Pointer to [**DataWarehouseSyncInterval**](DataWarehouseSyncInterval.md) |  | [optional] 

## Methods

### NewEndpointRequest

`func NewEndpointRequest() *EndpointRequest`

NewEndpointRequest instantiates a new EndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointRequestWithDefaults

`func NewEndpointRequestWithDefaults() *EndpointRequest`

NewEndpointRequestWithDefaults instantiates a new EndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheAgeSeconds

`func (o *EndpointRequest) GetCacheAgeSeconds() float32`

GetCacheAgeSeconds returns the CacheAgeSeconds field if non-nil, zero value otherwise.

### GetCacheAgeSecondsOk

`func (o *EndpointRequest) GetCacheAgeSecondsOk() (*float32, bool)`

GetCacheAgeSecondsOk returns a tuple with the CacheAgeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAgeSeconds

`func (o *EndpointRequest) SetCacheAgeSeconds(v float32)`

SetCacheAgeSeconds sets CacheAgeSeconds field to given value.

### HasCacheAgeSeconds

`func (o *EndpointRequest) HasCacheAgeSeconds() bool`

HasCacheAgeSeconds returns a boolean if a field has been set.

### SetCacheAgeSecondsNil

`func (o *EndpointRequest) SetCacheAgeSecondsNil(b bool)`

 SetCacheAgeSecondsNil sets the value for CacheAgeSeconds to be an explicit nil

### UnsetCacheAgeSeconds
`func (o *EndpointRequest) UnsetCacheAgeSeconds()`

UnsetCacheAgeSeconds ensures that no value is present for CacheAgeSeconds, not even an explicit nil
### GetDescription

`func (o *EndpointRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EndpointRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EndpointRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EndpointRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EndpointRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EndpointRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsActive

`func (o *EndpointRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *EndpointRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *EndpointRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *EndpointRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *EndpointRequest) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *EndpointRequest) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsMaterialized

`func (o *EndpointRequest) GetIsMaterialized() bool`

GetIsMaterialized returns the IsMaterialized field if non-nil, zero value otherwise.

### GetIsMaterializedOk

`func (o *EndpointRequest) GetIsMaterializedOk() (*bool, bool)`

GetIsMaterializedOk returns a tuple with the IsMaterialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaterialized

`func (o *EndpointRequest) SetIsMaterialized(v bool)`

SetIsMaterialized sets IsMaterialized field to given value.

### HasIsMaterialized

`func (o *EndpointRequest) HasIsMaterialized() bool`

HasIsMaterialized returns a boolean if a field has been set.

### SetIsMaterializedNil

`func (o *EndpointRequest) SetIsMaterializedNil(b bool)`

 SetIsMaterializedNil sets the value for IsMaterialized to be an explicit nil

### UnsetIsMaterialized
`func (o *EndpointRequest) UnsetIsMaterialized()`

UnsetIsMaterialized ensures that no value is present for IsMaterialized, not even an explicit nil
### GetName

`func (o *EndpointRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EndpointRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EndpointRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EndpointRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuery

`func (o *EndpointRequest) GetQuery() Query`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EndpointRequest) GetQueryOk() (*Query, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EndpointRequest) SetQuery(v Query)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *EndpointRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *EndpointRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *EndpointRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetSyncFrequency

`func (o *EndpointRequest) GetSyncFrequency() DataWarehouseSyncInterval`

GetSyncFrequency returns the SyncFrequency field if non-nil, zero value otherwise.

### GetSyncFrequencyOk

`func (o *EndpointRequest) GetSyncFrequencyOk() (*DataWarehouseSyncInterval, bool)`

GetSyncFrequencyOk returns a tuple with the SyncFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFrequency

`func (o *EndpointRequest) SetSyncFrequency(v DataWarehouseSyncInterval)`

SetSyncFrequency sets SyncFrequency field to given value.

### HasSyncFrequency

`func (o *EndpointRequest) HasSyncFrequency() bool`

HasSyncFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


