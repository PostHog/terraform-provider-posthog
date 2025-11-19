# DashboardTileBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**DashboardId** | **int32** |  | [readonly] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewDashboardTileBasic

`func NewDashboardTileBasic(id int32, dashboardId int32, ) *DashboardTileBasic`

NewDashboardTileBasic instantiates a new DashboardTileBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardTileBasicWithDefaults

`func NewDashboardTileBasicWithDefaults() *DashboardTileBasic`

NewDashboardTileBasicWithDefaults instantiates a new DashboardTileBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardTileBasic) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardTileBasic) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardTileBasic) SetId(v int32)`

SetId sets Id field to given value.


### GetDashboardId

`func (o *DashboardTileBasic) GetDashboardId() int32`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *DashboardTileBasic) GetDashboardIdOk() (*int32, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *DashboardTileBasic) SetDashboardId(v int32)`

SetDashboardId sets DashboardId field to given value.


### GetDeleted

`func (o *DashboardTileBasic) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DashboardTileBasic) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DashboardTileBasic) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DashboardTileBasic) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *DashboardTileBasic) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *DashboardTileBasic) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


