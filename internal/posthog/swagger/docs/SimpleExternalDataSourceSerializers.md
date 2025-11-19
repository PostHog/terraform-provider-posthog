# SimpleExternalDataSourceSerializers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**CreatedBy** | **NullableInt32** |  | [readonly] 
**Status** | **string** |  | [readonly] 
**SourceType** | [**SourceTypeEnum**](SourceTypeEnum.md) |  | [readonly] 

## Methods

### NewSimpleExternalDataSourceSerializers

`func NewSimpleExternalDataSourceSerializers(id string, createdAt time.Time, createdBy NullableInt32, status string, sourceType SourceTypeEnum, ) *SimpleExternalDataSourceSerializers`

NewSimpleExternalDataSourceSerializers instantiates a new SimpleExternalDataSourceSerializers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleExternalDataSourceSerializersWithDefaults

`func NewSimpleExternalDataSourceSerializersWithDefaults() *SimpleExternalDataSourceSerializers`

NewSimpleExternalDataSourceSerializersWithDefaults instantiates a new SimpleExternalDataSourceSerializers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleExternalDataSourceSerializers) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleExternalDataSourceSerializers) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleExternalDataSourceSerializers) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *SimpleExternalDataSourceSerializers) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimpleExternalDataSourceSerializers) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimpleExternalDataSourceSerializers) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SimpleExternalDataSourceSerializers) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SimpleExternalDataSourceSerializers) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SimpleExternalDataSourceSerializers) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *SimpleExternalDataSourceSerializers) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *SimpleExternalDataSourceSerializers) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetStatus

`func (o *SimpleExternalDataSourceSerializers) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimpleExternalDataSourceSerializers) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimpleExternalDataSourceSerializers) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSourceType

`func (o *SimpleExternalDataSourceSerializers) GetSourceType() SourceTypeEnum`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *SimpleExternalDataSourceSerializers) GetSourceTypeOk() (*SourceTypeEnum, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *SimpleExternalDataSourceSerializers) SetSourceType(v SourceTypeEnum)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


