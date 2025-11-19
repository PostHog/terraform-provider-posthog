# PatchedEarlyAccessFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**FeatureFlag** | Pointer to [**MinimalFeatureFlag**](MinimalFeatureFlag.md) |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to [**StageEnum**](StageEnum.md) |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedEarlyAccessFeature

`func NewPatchedEarlyAccessFeature() *PatchedEarlyAccessFeature`

NewPatchedEarlyAccessFeature instantiates a new PatchedEarlyAccessFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEarlyAccessFeatureWithDefaults

`func NewPatchedEarlyAccessFeatureWithDefaults() *PatchedEarlyAccessFeature`

NewPatchedEarlyAccessFeatureWithDefaults instantiates a new PatchedEarlyAccessFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedEarlyAccessFeature) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedEarlyAccessFeature) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedEarlyAccessFeature) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedEarlyAccessFeature) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFeatureFlag

`func (o *PatchedEarlyAccessFeature) GetFeatureFlag() MinimalFeatureFlag`

GetFeatureFlag returns the FeatureFlag field if non-nil, zero value otherwise.

### GetFeatureFlagOk

`func (o *PatchedEarlyAccessFeature) GetFeatureFlagOk() (*MinimalFeatureFlag, bool)`

GetFeatureFlagOk returns a tuple with the FeatureFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlag

`func (o *PatchedEarlyAccessFeature) SetFeatureFlag(v MinimalFeatureFlag)`

SetFeatureFlag sets FeatureFlag field to given value.

### HasFeatureFlag

`func (o *PatchedEarlyAccessFeature) HasFeatureFlag() bool`

HasFeatureFlag returns a boolean if a field has been set.

### GetName

`func (o *PatchedEarlyAccessFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEarlyAccessFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEarlyAccessFeature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEarlyAccessFeature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedEarlyAccessFeature) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedEarlyAccessFeature) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedEarlyAccessFeature) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedEarlyAccessFeature) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStage

`func (o *PatchedEarlyAccessFeature) GetStage() StageEnum`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PatchedEarlyAccessFeature) GetStageOk() (*StageEnum, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PatchedEarlyAccessFeature) SetStage(v StageEnum)`

SetStage sets Stage field to given value.

### HasStage

`func (o *PatchedEarlyAccessFeature) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *PatchedEarlyAccessFeature) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *PatchedEarlyAccessFeature) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *PatchedEarlyAccessFeature) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *PatchedEarlyAccessFeature) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedEarlyAccessFeature) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedEarlyAccessFeature) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedEarlyAccessFeature) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedEarlyAccessFeature) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


