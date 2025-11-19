# PatchedDatasetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Dataset** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 
**Output** | Pointer to **interface{}** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**RefTraceId** | Pointer to **NullableString** |  | [optional] 
**RefTimestamp** | Pointer to **NullableTime** |  | [optional] 
**RefSourceId** | Pointer to **NullableString** |  | [optional] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**Team** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewPatchedDatasetItem

`func NewPatchedDatasetItem() *PatchedDatasetItem`

NewPatchedDatasetItem instantiates a new PatchedDatasetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDatasetItemWithDefaults

`func NewPatchedDatasetItemWithDefaults() *PatchedDatasetItem`

NewPatchedDatasetItemWithDefaults instantiates a new PatchedDatasetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDatasetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDatasetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDatasetItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDatasetItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDataset

`func (o *PatchedDatasetItem) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *PatchedDatasetItem) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *PatchedDatasetItem) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *PatchedDatasetItem) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetInput

`func (o *PatchedDatasetItem) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *PatchedDatasetItem) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *PatchedDatasetItem) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *PatchedDatasetItem) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *PatchedDatasetItem) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *PatchedDatasetItem) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetOutput

`func (o *PatchedDatasetItem) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *PatchedDatasetItem) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *PatchedDatasetItem) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *PatchedDatasetItem) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *PatchedDatasetItem) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *PatchedDatasetItem) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetMetadata

`func (o *PatchedDatasetItem) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchedDatasetItem) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchedDatasetItem) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchedDatasetItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PatchedDatasetItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PatchedDatasetItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRefTraceId

`func (o *PatchedDatasetItem) GetRefTraceId() string`

GetRefTraceId returns the RefTraceId field if non-nil, zero value otherwise.

### GetRefTraceIdOk

`func (o *PatchedDatasetItem) GetRefTraceIdOk() (*string, bool)`

GetRefTraceIdOk returns a tuple with the RefTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTraceId

`func (o *PatchedDatasetItem) SetRefTraceId(v string)`

SetRefTraceId sets RefTraceId field to given value.

### HasRefTraceId

`func (o *PatchedDatasetItem) HasRefTraceId() bool`

HasRefTraceId returns a boolean if a field has been set.

### SetRefTraceIdNil

`func (o *PatchedDatasetItem) SetRefTraceIdNil(b bool)`

 SetRefTraceIdNil sets the value for RefTraceId to be an explicit nil

### UnsetRefTraceId
`func (o *PatchedDatasetItem) UnsetRefTraceId()`

UnsetRefTraceId ensures that no value is present for RefTraceId, not even an explicit nil
### GetRefTimestamp

`func (o *PatchedDatasetItem) GetRefTimestamp() time.Time`

GetRefTimestamp returns the RefTimestamp field if non-nil, zero value otherwise.

### GetRefTimestampOk

`func (o *PatchedDatasetItem) GetRefTimestampOk() (*time.Time, bool)`

GetRefTimestampOk returns a tuple with the RefTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTimestamp

`func (o *PatchedDatasetItem) SetRefTimestamp(v time.Time)`

SetRefTimestamp sets RefTimestamp field to given value.

### HasRefTimestamp

`func (o *PatchedDatasetItem) HasRefTimestamp() bool`

HasRefTimestamp returns a boolean if a field has been set.

### SetRefTimestampNil

`func (o *PatchedDatasetItem) SetRefTimestampNil(b bool)`

 SetRefTimestampNil sets the value for RefTimestamp to be an explicit nil

### UnsetRefTimestamp
`func (o *PatchedDatasetItem) UnsetRefTimestamp()`

UnsetRefTimestamp ensures that no value is present for RefTimestamp, not even an explicit nil
### GetRefSourceId

`func (o *PatchedDatasetItem) GetRefSourceId() string`

GetRefSourceId returns the RefSourceId field if non-nil, zero value otherwise.

### GetRefSourceIdOk

`func (o *PatchedDatasetItem) GetRefSourceIdOk() (*string, bool)`

GetRefSourceIdOk returns a tuple with the RefSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefSourceId

`func (o *PatchedDatasetItem) SetRefSourceId(v string)`

SetRefSourceId sets RefSourceId field to given value.

### HasRefSourceId

`func (o *PatchedDatasetItem) HasRefSourceId() bool`

HasRefSourceId returns a boolean if a field has been set.

### SetRefSourceIdNil

`func (o *PatchedDatasetItem) SetRefSourceIdNil(b bool)`

 SetRefSourceIdNil sets the value for RefSourceId to be an explicit nil

### UnsetRefSourceId
`func (o *PatchedDatasetItem) UnsetRefSourceId()`

UnsetRefSourceId ensures that no value is present for RefSourceId, not even an explicit nil
### GetDeleted

`func (o *PatchedDatasetItem) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedDatasetItem) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedDatasetItem) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedDatasetItem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *PatchedDatasetItem) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *PatchedDatasetItem) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetCreatedAt

`func (o *PatchedDatasetItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedDatasetItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedDatasetItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedDatasetItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedDatasetItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedDatasetItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedDatasetItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedDatasetItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *PatchedDatasetItem) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *PatchedDatasetItem) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCreatedBy

`func (o *PatchedDatasetItem) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedDatasetItem) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedDatasetItem) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedDatasetItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTeam

`func (o *PatchedDatasetItem) GetTeam() int32`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *PatchedDatasetItem) GetTeamOk() (*int32, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *PatchedDatasetItem) SetTeam(v int32)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *PatchedDatasetItem) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


