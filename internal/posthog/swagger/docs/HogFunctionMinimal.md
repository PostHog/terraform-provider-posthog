# HogFunctionMinimal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Type** | **NullableString** |  | [readonly] 
**Name** | **NullableString** |  | [readonly] 
**Description** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Enabled** | **bool** |  | [readonly] 
**Hog** | **string** |  | [readonly] 
**Filters** | **interface{}** |  | [readonly] 
**IconUrl** | **NullableString** |  | [readonly] 
**Template** | [**HogFunctionTemplate**](HogFunctionTemplate.md) |  | [readonly] 
**Status** | [**NullableHogFunctionStatus**](HogFunctionStatus.md) |  | [readonly] 
**ExecutionOrder** | **NullableInt32** |  | [readonly] 

## Methods

### NewHogFunctionMinimal

`func NewHogFunctionMinimal(id string, type_ NullableString, name NullableString, description string, createdAt time.Time, createdBy UserBasic, updatedAt time.Time, enabled bool, hog string, filters interface{}, iconUrl NullableString, template HogFunctionTemplate, status NullableHogFunctionStatus, executionOrder NullableInt32, ) *HogFunctionMinimal`

NewHogFunctionMinimal instantiates a new HogFunctionMinimal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogFunctionMinimalWithDefaults

`func NewHogFunctionMinimalWithDefaults() *HogFunctionMinimal`

NewHogFunctionMinimalWithDefaults instantiates a new HogFunctionMinimal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HogFunctionMinimal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HogFunctionMinimal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HogFunctionMinimal) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *HogFunctionMinimal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HogFunctionMinimal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HogFunctionMinimal) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *HogFunctionMinimal) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HogFunctionMinimal) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *HogFunctionMinimal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HogFunctionMinimal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HogFunctionMinimal) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *HogFunctionMinimal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HogFunctionMinimal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *HogFunctionMinimal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HogFunctionMinimal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HogFunctionMinimal) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *HogFunctionMinimal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HogFunctionMinimal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HogFunctionMinimal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *HogFunctionMinimal) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *HogFunctionMinimal) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *HogFunctionMinimal) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedAt

`func (o *HogFunctionMinimal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HogFunctionMinimal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HogFunctionMinimal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEnabled

`func (o *HogFunctionMinimal) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HogFunctionMinimal) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HogFunctionMinimal) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHog

`func (o *HogFunctionMinimal) GetHog() string`

GetHog returns the Hog field if non-nil, zero value otherwise.

### GetHogOk

`func (o *HogFunctionMinimal) GetHogOk() (*string, bool)`

GetHogOk returns a tuple with the Hog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHog

`func (o *HogFunctionMinimal) SetHog(v string)`

SetHog sets Hog field to given value.


### GetFilters

`func (o *HogFunctionMinimal) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *HogFunctionMinimal) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *HogFunctionMinimal) SetFilters(v interface{})`

SetFilters sets Filters field to given value.


### SetFiltersNil

`func (o *HogFunctionMinimal) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *HogFunctionMinimal) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetIconUrl

`func (o *HogFunctionMinimal) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *HogFunctionMinimal) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *HogFunctionMinimal) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### SetIconUrlNil

`func (o *HogFunctionMinimal) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *HogFunctionMinimal) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetTemplate

`func (o *HogFunctionMinimal) GetTemplate() HogFunctionTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *HogFunctionMinimal) GetTemplateOk() (*HogFunctionTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *HogFunctionMinimal) SetTemplate(v HogFunctionTemplate)`

SetTemplate sets Template field to given value.


### GetStatus

`func (o *HogFunctionMinimal) GetStatus() HogFunctionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HogFunctionMinimal) GetStatusOk() (*HogFunctionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HogFunctionMinimal) SetStatus(v HogFunctionStatus)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *HogFunctionMinimal) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *HogFunctionMinimal) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetExecutionOrder

`func (o *HogFunctionMinimal) GetExecutionOrder() int32`

GetExecutionOrder returns the ExecutionOrder field if non-nil, zero value otherwise.

### GetExecutionOrderOk

`func (o *HogFunctionMinimal) GetExecutionOrderOk() (*int32, bool)`

GetExecutionOrderOk returns a tuple with the ExecutionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOrder

`func (o *HogFunctionMinimal) SetExecutionOrder(v int32)`

SetExecutionOrder sets ExecutionOrder field to given value.


### SetExecutionOrderNil

`func (o *HogFunctionMinimal) SetExecutionOrderNil(b bool)`

 SetExecutionOrderNil sets the value for ExecutionOrder to be an explicit nil

### UnsetExecutionOrder
`func (o *HogFunctionMinimal) UnsetExecutionOrder()`

UnsetExecutionOrder ensures that no value is present for ExecutionOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


