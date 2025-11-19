# EarlyAccessFeatureSerializerCreateOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Stage** | [**StageEnum**](StageEnum.md) |  | 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**FeatureFlagId** | Pointer to **int32** |  | [optional] 
**FeatureFlag** | [**MinimalFeatureFlag**](MinimalFeatureFlag.md) |  | [readonly] 
**CreateInFolder** | Pointer to **string** |  | [optional] 

## Methods

### NewEarlyAccessFeatureSerializerCreateOnly

`func NewEarlyAccessFeatureSerializerCreateOnly(id string, name string, stage StageEnum, createdAt time.Time, featureFlag MinimalFeatureFlag, ) *EarlyAccessFeatureSerializerCreateOnly`

NewEarlyAccessFeatureSerializerCreateOnly instantiates a new EarlyAccessFeatureSerializerCreateOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarlyAccessFeatureSerializerCreateOnlyWithDefaults

`func NewEarlyAccessFeatureSerializerCreateOnlyWithDefaults() *EarlyAccessFeatureSerializerCreateOnly`

NewEarlyAccessFeatureSerializerCreateOnlyWithDefaults instantiates a new EarlyAccessFeatureSerializerCreateOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EarlyAccessFeatureSerializerCreateOnly) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EarlyAccessFeatureSerializerCreateOnly) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EarlyAccessFeatureSerializerCreateOnly) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EarlyAccessFeatureSerializerCreateOnly) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStage

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetStage() StageEnum`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetStageOk() (*StageEnum, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *EarlyAccessFeatureSerializerCreateOnly) SetStage(v StageEnum)`

SetStage sets Stage field to given value.


### GetDocumentationUrl

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *EarlyAccessFeatureSerializerCreateOnly) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *EarlyAccessFeatureSerializerCreateOnly) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EarlyAccessFeatureSerializerCreateOnly) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFeatureFlagId

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetFeatureFlagId() int32`

GetFeatureFlagId returns the FeatureFlagId field if non-nil, zero value otherwise.

### GetFeatureFlagIdOk

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetFeatureFlagIdOk() (*int32, bool)`

GetFeatureFlagIdOk returns a tuple with the FeatureFlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagId

`func (o *EarlyAccessFeatureSerializerCreateOnly) SetFeatureFlagId(v int32)`

SetFeatureFlagId sets FeatureFlagId field to given value.

### HasFeatureFlagId

`func (o *EarlyAccessFeatureSerializerCreateOnly) HasFeatureFlagId() bool`

HasFeatureFlagId returns a boolean if a field has been set.

### GetFeatureFlag

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetFeatureFlag() MinimalFeatureFlag`

GetFeatureFlag returns the FeatureFlag field if non-nil, zero value otherwise.

### GetFeatureFlagOk

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetFeatureFlagOk() (*MinimalFeatureFlag, bool)`

GetFeatureFlagOk returns a tuple with the FeatureFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlag

`func (o *EarlyAccessFeatureSerializerCreateOnly) SetFeatureFlag(v MinimalFeatureFlag)`

SetFeatureFlag sets FeatureFlag field to given value.


### GetCreateInFolder

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *EarlyAccessFeatureSerializerCreateOnly) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *EarlyAccessFeatureSerializerCreateOnly) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *EarlyAccessFeatureSerializerCreateOnly) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


