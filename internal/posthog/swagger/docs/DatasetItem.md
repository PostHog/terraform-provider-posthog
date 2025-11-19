# DatasetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Dataset** | **string** |  | 
**Input** | Pointer to **interface{}** |  | [optional] 
**Output** | Pointer to **interface{}** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**RefTraceId** | Pointer to **NullableString** |  | [optional] 
**RefTimestamp** | Pointer to **NullableTime** |  | [optional] 
**RefSourceId** | Pointer to **NullableString** |  | [optional] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **NullableTime** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**Team** | **int32** |  | [readonly] 

## Methods

### NewDatasetItem

`func NewDatasetItem(id string, dataset string, createdAt time.Time, updatedAt NullableTime, createdBy UserBasic, team int32, ) *DatasetItem`

NewDatasetItem instantiates a new DatasetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetItemWithDefaults

`func NewDatasetItemWithDefaults() *DatasetItem`

NewDatasetItemWithDefaults instantiates a new DatasetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatasetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatasetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatasetItem) SetId(v string)`

SetId sets Id field to given value.


### GetDataset

`func (o *DatasetItem) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *DatasetItem) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *DatasetItem) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetInput

`func (o *DatasetItem) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *DatasetItem) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *DatasetItem) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *DatasetItem) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *DatasetItem) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *DatasetItem) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetOutput

`func (o *DatasetItem) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *DatasetItem) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *DatasetItem) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *DatasetItem) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *DatasetItem) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *DatasetItem) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetMetadata

`func (o *DatasetItem) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DatasetItem) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DatasetItem) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DatasetItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DatasetItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DatasetItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRefTraceId

`func (o *DatasetItem) GetRefTraceId() string`

GetRefTraceId returns the RefTraceId field if non-nil, zero value otherwise.

### GetRefTraceIdOk

`func (o *DatasetItem) GetRefTraceIdOk() (*string, bool)`

GetRefTraceIdOk returns a tuple with the RefTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTraceId

`func (o *DatasetItem) SetRefTraceId(v string)`

SetRefTraceId sets RefTraceId field to given value.

### HasRefTraceId

`func (o *DatasetItem) HasRefTraceId() bool`

HasRefTraceId returns a boolean if a field has been set.

### SetRefTraceIdNil

`func (o *DatasetItem) SetRefTraceIdNil(b bool)`

 SetRefTraceIdNil sets the value for RefTraceId to be an explicit nil

### UnsetRefTraceId
`func (o *DatasetItem) UnsetRefTraceId()`

UnsetRefTraceId ensures that no value is present for RefTraceId, not even an explicit nil
### GetRefTimestamp

`func (o *DatasetItem) GetRefTimestamp() time.Time`

GetRefTimestamp returns the RefTimestamp field if non-nil, zero value otherwise.

### GetRefTimestampOk

`func (o *DatasetItem) GetRefTimestampOk() (*time.Time, bool)`

GetRefTimestampOk returns a tuple with the RefTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTimestamp

`func (o *DatasetItem) SetRefTimestamp(v time.Time)`

SetRefTimestamp sets RefTimestamp field to given value.

### HasRefTimestamp

`func (o *DatasetItem) HasRefTimestamp() bool`

HasRefTimestamp returns a boolean if a field has been set.

### SetRefTimestampNil

`func (o *DatasetItem) SetRefTimestampNil(b bool)`

 SetRefTimestampNil sets the value for RefTimestamp to be an explicit nil

### UnsetRefTimestamp
`func (o *DatasetItem) UnsetRefTimestamp()`

UnsetRefTimestamp ensures that no value is present for RefTimestamp, not even an explicit nil
### GetRefSourceId

`func (o *DatasetItem) GetRefSourceId() string`

GetRefSourceId returns the RefSourceId field if non-nil, zero value otherwise.

### GetRefSourceIdOk

`func (o *DatasetItem) GetRefSourceIdOk() (*string, bool)`

GetRefSourceIdOk returns a tuple with the RefSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefSourceId

`func (o *DatasetItem) SetRefSourceId(v string)`

SetRefSourceId sets RefSourceId field to given value.

### HasRefSourceId

`func (o *DatasetItem) HasRefSourceId() bool`

HasRefSourceId returns a boolean if a field has been set.

### SetRefSourceIdNil

`func (o *DatasetItem) SetRefSourceIdNil(b bool)`

 SetRefSourceIdNil sets the value for RefSourceId to be an explicit nil

### UnsetRefSourceId
`func (o *DatasetItem) UnsetRefSourceId()`

UnsetRefSourceId ensures that no value is present for RefSourceId, not even an explicit nil
### GetDeleted

`func (o *DatasetItem) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DatasetItem) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DatasetItem) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DatasetItem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *DatasetItem) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *DatasetItem) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetCreatedAt

`func (o *DatasetItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatasetItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatasetItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DatasetItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DatasetItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DatasetItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *DatasetItem) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *DatasetItem) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCreatedBy

`func (o *DatasetItem) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DatasetItem) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DatasetItem) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetTeam

`func (o *DatasetItem) GetTeam() int32`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *DatasetItem) GetTeamOk() (*int32, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *DatasetItem) SetTeam(v int32)`

SetTeam sets Team field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


