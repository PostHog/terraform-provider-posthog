# Annotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Content** | Pointer to **NullableString** |  | [optional] 
**DateMarker** | Pointer to **NullableTime** |  | [optional] 
**CreationType** | Pointer to [**CreationTypeEnum**](CreationTypeEnum.md) |  | [optional] 
**DashboardItem** | Pointer to **NullableInt32** |  | [optional] 
**DashboardId** | **NullableInt32** |  | [readonly] 
**DashboardName** | **NullableString** |  | [readonly] 
**InsightShortId** | **NullableString** |  | [readonly] 
**InsightName** | **NullableString** |  | [readonly] 
**InsightDerivedName** | **NullableString** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**CreatedAt** | **NullableTime** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Scope** | Pointer to [**AnnotationScopeEnum**](AnnotationScopeEnum.md) |  | [optional] 

## Methods

### NewAnnotation

`func NewAnnotation(id int32, dashboardId NullableInt32, dashboardName NullableString, insightShortId NullableString, insightName NullableString, insightDerivedName NullableString, createdBy UserBasic, createdAt NullableTime, updatedAt time.Time, ) *Annotation`

NewAnnotation instantiates a new Annotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationWithDefaults

`func NewAnnotationWithDefaults() *Annotation`

NewAnnotationWithDefaults instantiates a new Annotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Annotation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Annotation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Annotation) SetId(v int32)`

SetId sets Id field to given value.


### GetContent

`func (o *Annotation) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Annotation) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Annotation) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Annotation) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *Annotation) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *Annotation) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetDateMarker

`func (o *Annotation) GetDateMarker() time.Time`

GetDateMarker returns the DateMarker field if non-nil, zero value otherwise.

### GetDateMarkerOk

`func (o *Annotation) GetDateMarkerOk() (*time.Time, bool)`

GetDateMarkerOk returns a tuple with the DateMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateMarker

`func (o *Annotation) SetDateMarker(v time.Time)`

SetDateMarker sets DateMarker field to given value.

### HasDateMarker

`func (o *Annotation) HasDateMarker() bool`

HasDateMarker returns a boolean if a field has been set.

### SetDateMarkerNil

`func (o *Annotation) SetDateMarkerNil(b bool)`

 SetDateMarkerNil sets the value for DateMarker to be an explicit nil

### UnsetDateMarker
`func (o *Annotation) UnsetDateMarker()`

UnsetDateMarker ensures that no value is present for DateMarker, not even an explicit nil
### GetCreationType

`func (o *Annotation) GetCreationType() CreationTypeEnum`

GetCreationType returns the CreationType field if non-nil, zero value otherwise.

### GetCreationTypeOk

`func (o *Annotation) GetCreationTypeOk() (*CreationTypeEnum, bool)`

GetCreationTypeOk returns a tuple with the CreationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationType

`func (o *Annotation) SetCreationType(v CreationTypeEnum)`

SetCreationType sets CreationType field to given value.

### HasCreationType

`func (o *Annotation) HasCreationType() bool`

HasCreationType returns a boolean if a field has been set.

### GetDashboardItem

`func (o *Annotation) GetDashboardItem() int32`

GetDashboardItem returns the DashboardItem field if non-nil, zero value otherwise.

### GetDashboardItemOk

`func (o *Annotation) GetDashboardItemOk() (*int32, bool)`

GetDashboardItemOk returns a tuple with the DashboardItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardItem

`func (o *Annotation) SetDashboardItem(v int32)`

SetDashboardItem sets DashboardItem field to given value.

### HasDashboardItem

`func (o *Annotation) HasDashboardItem() bool`

HasDashboardItem returns a boolean if a field has been set.

### SetDashboardItemNil

`func (o *Annotation) SetDashboardItemNil(b bool)`

 SetDashboardItemNil sets the value for DashboardItem to be an explicit nil

### UnsetDashboardItem
`func (o *Annotation) UnsetDashboardItem()`

UnsetDashboardItem ensures that no value is present for DashboardItem, not even an explicit nil
### GetDashboardId

`func (o *Annotation) GetDashboardId() int32`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *Annotation) GetDashboardIdOk() (*int32, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *Annotation) SetDashboardId(v int32)`

SetDashboardId sets DashboardId field to given value.


### SetDashboardIdNil

`func (o *Annotation) SetDashboardIdNil(b bool)`

 SetDashboardIdNil sets the value for DashboardId to be an explicit nil

### UnsetDashboardId
`func (o *Annotation) UnsetDashboardId()`

UnsetDashboardId ensures that no value is present for DashboardId, not even an explicit nil
### GetDashboardName

`func (o *Annotation) GetDashboardName() string`

GetDashboardName returns the DashboardName field if non-nil, zero value otherwise.

### GetDashboardNameOk

`func (o *Annotation) GetDashboardNameOk() (*string, bool)`

GetDashboardNameOk returns a tuple with the DashboardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardName

`func (o *Annotation) SetDashboardName(v string)`

SetDashboardName sets DashboardName field to given value.


### SetDashboardNameNil

`func (o *Annotation) SetDashboardNameNil(b bool)`

 SetDashboardNameNil sets the value for DashboardName to be an explicit nil

### UnsetDashboardName
`func (o *Annotation) UnsetDashboardName()`

UnsetDashboardName ensures that no value is present for DashboardName, not even an explicit nil
### GetInsightShortId

`func (o *Annotation) GetInsightShortId() string`

GetInsightShortId returns the InsightShortId field if non-nil, zero value otherwise.

### GetInsightShortIdOk

`func (o *Annotation) GetInsightShortIdOk() (*string, bool)`

GetInsightShortIdOk returns a tuple with the InsightShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightShortId

`func (o *Annotation) SetInsightShortId(v string)`

SetInsightShortId sets InsightShortId field to given value.


### SetInsightShortIdNil

`func (o *Annotation) SetInsightShortIdNil(b bool)`

 SetInsightShortIdNil sets the value for InsightShortId to be an explicit nil

### UnsetInsightShortId
`func (o *Annotation) UnsetInsightShortId()`

UnsetInsightShortId ensures that no value is present for InsightShortId, not even an explicit nil
### GetInsightName

`func (o *Annotation) GetInsightName() string`

GetInsightName returns the InsightName field if non-nil, zero value otherwise.

### GetInsightNameOk

`func (o *Annotation) GetInsightNameOk() (*string, bool)`

GetInsightNameOk returns a tuple with the InsightName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightName

`func (o *Annotation) SetInsightName(v string)`

SetInsightName sets InsightName field to given value.


### SetInsightNameNil

`func (o *Annotation) SetInsightNameNil(b bool)`

 SetInsightNameNil sets the value for InsightName to be an explicit nil

### UnsetInsightName
`func (o *Annotation) UnsetInsightName()`

UnsetInsightName ensures that no value is present for InsightName, not even an explicit nil
### GetInsightDerivedName

`func (o *Annotation) GetInsightDerivedName() string`

GetInsightDerivedName returns the InsightDerivedName field if non-nil, zero value otherwise.

### GetInsightDerivedNameOk

`func (o *Annotation) GetInsightDerivedNameOk() (*string, bool)`

GetInsightDerivedNameOk returns a tuple with the InsightDerivedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightDerivedName

`func (o *Annotation) SetInsightDerivedName(v string)`

SetInsightDerivedName sets InsightDerivedName field to given value.


### SetInsightDerivedNameNil

`func (o *Annotation) SetInsightDerivedNameNil(b bool)`

 SetInsightDerivedNameNil sets the value for InsightDerivedName to be an explicit nil

### UnsetInsightDerivedName
`func (o *Annotation) UnsetInsightDerivedName()`

UnsetInsightDerivedName ensures that no value is present for InsightDerivedName, not even an explicit nil
### GetCreatedBy

`func (o *Annotation) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Annotation) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Annotation) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *Annotation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Annotation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Annotation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Annotation) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Annotation) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Annotation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Annotation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Annotation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeleted

`func (o *Annotation) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Annotation) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Annotation) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Annotation) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetScope

`func (o *Annotation) GetScope() AnnotationScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Annotation) GetScopeOk() (*AnnotationScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Annotation) SetScope(v AnnotationScopeEnum)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Annotation) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


