# PatchedAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Content** | Pointer to **NullableString** |  | [optional] 
**DateMarker** | Pointer to **NullableTime** |  | [optional] 
**CreationType** | Pointer to [**CreationTypeEnum**](CreationTypeEnum.md) |  | [optional] 
**DashboardItem** | Pointer to **NullableInt32** |  | [optional] 
**DashboardId** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**DashboardName** | Pointer to **NullableString** |  | [optional] [readonly] 
**InsightShortId** | Pointer to **NullableString** |  | [optional] [readonly] 
**InsightName** | Pointer to **NullableString** |  | [optional] [readonly] 
**InsightDerivedName** | Pointer to **NullableString** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Scope** | Pointer to [**AnnotationScopeEnum**](AnnotationScopeEnum.md) |  | [optional] 

## Methods

### NewPatchedAnnotation

`func NewPatchedAnnotation() *PatchedAnnotation`

NewPatchedAnnotation instantiates a new PatchedAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAnnotationWithDefaults

`func NewPatchedAnnotationWithDefaults() *PatchedAnnotation`

NewPatchedAnnotationWithDefaults instantiates a new PatchedAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedAnnotation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedAnnotation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedAnnotation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedAnnotation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContent

`func (o *PatchedAnnotation) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PatchedAnnotation) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PatchedAnnotation) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *PatchedAnnotation) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *PatchedAnnotation) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *PatchedAnnotation) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetDateMarker

`func (o *PatchedAnnotation) GetDateMarker() time.Time`

GetDateMarker returns the DateMarker field if non-nil, zero value otherwise.

### GetDateMarkerOk

`func (o *PatchedAnnotation) GetDateMarkerOk() (*time.Time, bool)`

GetDateMarkerOk returns a tuple with the DateMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateMarker

`func (o *PatchedAnnotation) SetDateMarker(v time.Time)`

SetDateMarker sets DateMarker field to given value.

### HasDateMarker

`func (o *PatchedAnnotation) HasDateMarker() bool`

HasDateMarker returns a boolean if a field has been set.

### SetDateMarkerNil

`func (o *PatchedAnnotation) SetDateMarkerNil(b bool)`

 SetDateMarkerNil sets the value for DateMarker to be an explicit nil

### UnsetDateMarker
`func (o *PatchedAnnotation) UnsetDateMarker()`

UnsetDateMarker ensures that no value is present for DateMarker, not even an explicit nil
### GetCreationType

`func (o *PatchedAnnotation) GetCreationType() CreationTypeEnum`

GetCreationType returns the CreationType field if non-nil, zero value otherwise.

### GetCreationTypeOk

`func (o *PatchedAnnotation) GetCreationTypeOk() (*CreationTypeEnum, bool)`

GetCreationTypeOk returns a tuple with the CreationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationType

`func (o *PatchedAnnotation) SetCreationType(v CreationTypeEnum)`

SetCreationType sets CreationType field to given value.

### HasCreationType

`func (o *PatchedAnnotation) HasCreationType() bool`

HasCreationType returns a boolean if a field has been set.

### GetDashboardItem

`func (o *PatchedAnnotation) GetDashboardItem() int32`

GetDashboardItem returns the DashboardItem field if non-nil, zero value otherwise.

### GetDashboardItemOk

`func (o *PatchedAnnotation) GetDashboardItemOk() (*int32, bool)`

GetDashboardItemOk returns a tuple with the DashboardItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardItem

`func (o *PatchedAnnotation) SetDashboardItem(v int32)`

SetDashboardItem sets DashboardItem field to given value.

### HasDashboardItem

`func (o *PatchedAnnotation) HasDashboardItem() bool`

HasDashboardItem returns a boolean if a field has been set.

### SetDashboardItemNil

`func (o *PatchedAnnotation) SetDashboardItemNil(b bool)`

 SetDashboardItemNil sets the value for DashboardItem to be an explicit nil

### UnsetDashboardItem
`func (o *PatchedAnnotation) UnsetDashboardItem()`

UnsetDashboardItem ensures that no value is present for DashboardItem, not even an explicit nil
### GetDashboardId

`func (o *PatchedAnnotation) GetDashboardId() int32`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *PatchedAnnotation) GetDashboardIdOk() (*int32, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *PatchedAnnotation) SetDashboardId(v int32)`

SetDashboardId sets DashboardId field to given value.

### HasDashboardId

`func (o *PatchedAnnotation) HasDashboardId() bool`

HasDashboardId returns a boolean if a field has been set.

### SetDashboardIdNil

`func (o *PatchedAnnotation) SetDashboardIdNil(b bool)`

 SetDashboardIdNil sets the value for DashboardId to be an explicit nil

### UnsetDashboardId
`func (o *PatchedAnnotation) UnsetDashboardId()`

UnsetDashboardId ensures that no value is present for DashboardId, not even an explicit nil
### GetDashboardName

`func (o *PatchedAnnotation) GetDashboardName() string`

GetDashboardName returns the DashboardName field if non-nil, zero value otherwise.

### GetDashboardNameOk

`func (o *PatchedAnnotation) GetDashboardNameOk() (*string, bool)`

GetDashboardNameOk returns a tuple with the DashboardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardName

`func (o *PatchedAnnotation) SetDashboardName(v string)`

SetDashboardName sets DashboardName field to given value.

### HasDashboardName

`func (o *PatchedAnnotation) HasDashboardName() bool`

HasDashboardName returns a boolean if a field has been set.

### SetDashboardNameNil

`func (o *PatchedAnnotation) SetDashboardNameNil(b bool)`

 SetDashboardNameNil sets the value for DashboardName to be an explicit nil

### UnsetDashboardName
`func (o *PatchedAnnotation) UnsetDashboardName()`

UnsetDashboardName ensures that no value is present for DashboardName, not even an explicit nil
### GetInsightShortId

`func (o *PatchedAnnotation) GetInsightShortId() string`

GetInsightShortId returns the InsightShortId field if non-nil, zero value otherwise.

### GetInsightShortIdOk

`func (o *PatchedAnnotation) GetInsightShortIdOk() (*string, bool)`

GetInsightShortIdOk returns a tuple with the InsightShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightShortId

`func (o *PatchedAnnotation) SetInsightShortId(v string)`

SetInsightShortId sets InsightShortId field to given value.

### HasInsightShortId

`func (o *PatchedAnnotation) HasInsightShortId() bool`

HasInsightShortId returns a boolean if a field has been set.

### SetInsightShortIdNil

`func (o *PatchedAnnotation) SetInsightShortIdNil(b bool)`

 SetInsightShortIdNil sets the value for InsightShortId to be an explicit nil

### UnsetInsightShortId
`func (o *PatchedAnnotation) UnsetInsightShortId()`

UnsetInsightShortId ensures that no value is present for InsightShortId, not even an explicit nil
### GetInsightName

`func (o *PatchedAnnotation) GetInsightName() string`

GetInsightName returns the InsightName field if non-nil, zero value otherwise.

### GetInsightNameOk

`func (o *PatchedAnnotation) GetInsightNameOk() (*string, bool)`

GetInsightNameOk returns a tuple with the InsightName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightName

`func (o *PatchedAnnotation) SetInsightName(v string)`

SetInsightName sets InsightName field to given value.

### HasInsightName

`func (o *PatchedAnnotation) HasInsightName() bool`

HasInsightName returns a boolean if a field has been set.

### SetInsightNameNil

`func (o *PatchedAnnotation) SetInsightNameNil(b bool)`

 SetInsightNameNil sets the value for InsightName to be an explicit nil

### UnsetInsightName
`func (o *PatchedAnnotation) UnsetInsightName()`

UnsetInsightName ensures that no value is present for InsightName, not even an explicit nil
### GetInsightDerivedName

`func (o *PatchedAnnotation) GetInsightDerivedName() string`

GetInsightDerivedName returns the InsightDerivedName field if non-nil, zero value otherwise.

### GetInsightDerivedNameOk

`func (o *PatchedAnnotation) GetInsightDerivedNameOk() (*string, bool)`

GetInsightDerivedNameOk returns a tuple with the InsightDerivedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightDerivedName

`func (o *PatchedAnnotation) SetInsightDerivedName(v string)`

SetInsightDerivedName sets InsightDerivedName field to given value.

### HasInsightDerivedName

`func (o *PatchedAnnotation) HasInsightDerivedName() bool`

HasInsightDerivedName returns a boolean if a field has been set.

### SetInsightDerivedNameNil

`func (o *PatchedAnnotation) SetInsightDerivedNameNil(b bool)`

 SetInsightDerivedNameNil sets the value for InsightDerivedName to be an explicit nil

### UnsetInsightDerivedName
`func (o *PatchedAnnotation) UnsetInsightDerivedName()`

UnsetInsightDerivedName ensures that no value is present for InsightDerivedName, not even an explicit nil
### GetCreatedBy

`func (o *PatchedAnnotation) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedAnnotation) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedAnnotation) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedAnnotation) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedAnnotation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedAnnotation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedAnnotation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedAnnotation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PatchedAnnotation) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PatchedAnnotation) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *PatchedAnnotation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedAnnotation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedAnnotation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedAnnotation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedAnnotation) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedAnnotation) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedAnnotation) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedAnnotation) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetScope

`func (o *PatchedAnnotation) GetScope() AnnotationScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PatchedAnnotation) GetScopeOk() (*AnnotationScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PatchedAnnotation) SetScope(v AnnotationScopeEnum)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PatchedAnnotation) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


