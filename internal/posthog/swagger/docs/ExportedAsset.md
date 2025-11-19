# ExportedAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Dashboard** | Pointer to **NullableInt32** |  | [optional] 
**Insight** | Pointer to **NullableInt32** |  | [optional] 
**ExportFormat** | [**ExportFormatEnum**](ExportFormatEnum.md) |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**HasContent** | **string** |  | [readonly] 
**ExportContext** | Pointer to **interface{}** |  | [optional] 
**Filename** | **string** |  | [readonly] 
**ExpiresAfter** | **NullableTime** |  | [readonly] 
**Exception** | **NullableString** |  | [readonly] 

## Methods

### NewExportedAsset

`func NewExportedAsset(id int32, exportFormat ExportFormatEnum, createdAt time.Time, hasContent string, filename string, expiresAfter NullableTime, exception NullableString, ) *ExportedAsset`

NewExportedAsset instantiates a new ExportedAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportedAssetWithDefaults

`func NewExportedAssetWithDefaults() *ExportedAsset`

NewExportedAssetWithDefaults instantiates a new ExportedAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExportedAsset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExportedAsset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExportedAsset) SetId(v int32)`

SetId sets Id field to given value.


### GetDashboard

`func (o *ExportedAsset) GetDashboard() int32`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *ExportedAsset) GetDashboardOk() (*int32, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *ExportedAsset) SetDashboard(v int32)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *ExportedAsset) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### SetDashboardNil

`func (o *ExportedAsset) SetDashboardNil(b bool)`

 SetDashboardNil sets the value for Dashboard to be an explicit nil

### UnsetDashboard
`func (o *ExportedAsset) UnsetDashboard()`

UnsetDashboard ensures that no value is present for Dashboard, not even an explicit nil
### GetInsight

`func (o *ExportedAsset) GetInsight() int32`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *ExportedAsset) GetInsightOk() (*int32, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *ExportedAsset) SetInsight(v int32)`

SetInsight sets Insight field to given value.

### HasInsight

`func (o *ExportedAsset) HasInsight() bool`

HasInsight returns a boolean if a field has been set.

### SetInsightNil

`func (o *ExportedAsset) SetInsightNil(b bool)`

 SetInsightNil sets the value for Insight to be an explicit nil

### UnsetInsight
`func (o *ExportedAsset) UnsetInsight()`

UnsetInsight ensures that no value is present for Insight, not even an explicit nil
### GetExportFormat

`func (o *ExportedAsset) GetExportFormat() ExportFormatEnum`

GetExportFormat returns the ExportFormat field if non-nil, zero value otherwise.

### GetExportFormatOk

`func (o *ExportedAsset) GetExportFormatOk() (*ExportFormatEnum, bool)`

GetExportFormatOk returns a tuple with the ExportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportFormat

`func (o *ExportedAsset) SetExportFormat(v ExportFormatEnum)`

SetExportFormat sets ExportFormat field to given value.


### GetCreatedAt

`func (o *ExportedAsset) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExportedAsset) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExportedAsset) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetHasContent

`func (o *ExportedAsset) GetHasContent() string`

GetHasContent returns the HasContent field if non-nil, zero value otherwise.

### GetHasContentOk

`func (o *ExportedAsset) GetHasContentOk() (*string, bool)`

GetHasContentOk returns a tuple with the HasContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasContent

`func (o *ExportedAsset) SetHasContent(v string)`

SetHasContent sets HasContent field to given value.


### GetExportContext

`func (o *ExportedAsset) GetExportContext() interface{}`

GetExportContext returns the ExportContext field if non-nil, zero value otherwise.

### GetExportContextOk

`func (o *ExportedAsset) GetExportContextOk() (*interface{}, bool)`

GetExportContextOk returns a tuple with the ExportContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportContext

`func (o *ExportedAsset) SetExportContext(v interface{})`

SetExportContext sets ExportContext field to given value.

### HasExportContext

`func (o *ExportedAsset) HasExportContext() bool`

HasExportContext returns a boolean if a field has been set.

### SetExportContextNil

`func (o *ExportedAsset) SetExportContextNil(b bool)`

 SetExportContextNil sets the value for ExportContext to be an explicit nil

### UnsetExportContext
`func (o *ExportedAsset) UnsetExportContext()`

UnsetExportContext ensures that no value is present for ExportContext, not even an explicit nil
### GetFilename

`func (o *ExportedAsset) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ExportedAsset) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ExportedAsset) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetExpiresAfter

`func (o *ExportedAsset) GetExpiresAfter() time.Time`

GetExpiresAfter returns the ExpiresAfter field if non-nil, zero value otherwise.

### GetExpiresAfterOk

`func (o *ExportedAsset) GetExpiresAfterOk() (*time.Time, bool)`

GetExpiresAfterOk returns a tuple with the ExpiresAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAfter

`func (o *ExportedAsset) SetExpiresAfter(v time.Time)`

SetExpiresAfter sets ExpiresAfter field to given value.


### SetExpiresAfterNil

`func (o *ExportedAsset) SetExpiresAfterNil(b bool)`

 SetExpiresAfterNil sets the value for ExpiresAfter to be an explicit nil

### UnsetExpiresAfter
`func (o *ExportedAsset) UnsetExpiresAfter()`

UnsetExpiresAfter ensures that no value is present for ExpiresAfter, not even an explicit nil
### GetException

`func (o *ExportedAsset) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *ExportedAsset) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *ExportedAsset) SetException(v string)`

SetException sets Exception field to given value.


### SetExceptionNil

`func (o *ExportedAsset) SetExceptionNil(b bool)`

 SetExceptionNil sets the value for Exception to be an explicit nil

### UnsetException
`func (o *ExportedAsset) UnsetException()`

UnsetException ensures that no value is present for Exception, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


