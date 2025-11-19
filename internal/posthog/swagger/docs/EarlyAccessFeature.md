# EarlyAccessFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**FeatureFlag** | [**MinimalFeatureFlag**](MinimalFeatureFlag.md) |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Stage** | [**StageEnum**](StageEnum.md) |  | 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewEarlyAccessFeature

`func NewEarlyAccessFeature(id string, featureFlag MinimalFeatureFlag, name string, stage StageEnum, createdAt time.Time, ) *EarlyAccessFeature`

NewEarlyAccessFeature instantiates a new EarlyAccessFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarlyAccessFeatureWithDefaults

`func NewEarlyAccessFeatureWithDefaults() *EarlyAccessFeature`

NewEarlyAccessFeatureWithDefaults instantiates a new EarlyAccessFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EarlyAccessFeature) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EarlyAccessFeature) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EarlyAccessFeature) SetId(v string)`

SetId sets Id field to given value.


### GetFeatureFlag

`func (o *EarlyAccessFeature) GetFeatureFlag() MinimalFeatureFlag`

GetFeatureFlag returns the FeatureFlag field if non-nil, zero value otherwise.

### GetFeatureFlagOk

`func (o *EarlyAccessFeature) GetFeatureFlagOk() (*MinimalFeatureFlag, bool)`

GetFeatureFlagOk returns a tuple with the FeatureFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlag

`func (o *EarlyAccessFeature) SetFeatureFlag(v MinimalFeatureFlag)`

SetFeatureFlag sets FeatureFlag field to given value.


### GetName

`func (o *EarlyAccessFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EarlyAccessFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EarlyAccessFeature) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EarlyAccessFeature) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EarlyAccessFeature) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EarlyAccessFeature) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EarlyAccessFeature) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStage

`func (o *EarlyAccessFeature) GetStage() StageEnum`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *EarlyAccessFeature) GetStageOk() (*StageEnum, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *EarlyAccessFeature) SetStage(v StageEnum)`

SetStage sets Stage field to given value.


### GetDocumentationUrl

`func (o *EarlyAccessFeature) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *EarlyAccessFeature) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *EarlyAccessFeature) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *EarlyAccessFeature) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EarlyAccessFeature) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EarlyAccessFeature) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EarlyAccessFeature) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


