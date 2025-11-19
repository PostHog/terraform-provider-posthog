# ExperimentToSavedMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Experiment** | **int32** |  | 
**SavedMetric** | **int32** |  | 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**Query** | **interface{}** |  | [readonly] 
**Name** | **string** |  | [readonly] 

## Methods

### NewExperimentToSavedMetric

`func NewExperimentToSavedMetric(id int32, experiment int32, savedMetric int32, createdAt time.Time, query interface{}, name string, ) *ExperimentToSavedMetric`

NewExperimentToSavedMetric instantiates a new ExperimentToSavedMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentToSavedMetricWithDefaults

`func NewExperimentToSavedMetricWithDefaults() *ExperimentToSavedMetric`

NewExperimentToSavedMetricWithDefaults instantiates a new ExperimentToSavedMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExperimentToSavedMetric) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperimentToSavedMetric) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperimentToSavedMetric) SetId(v int32)`

SetId sets Id field to given value.


### GetExperiment

`func (o *ExperimentToSavedMetric) GetExperiment() int32`

GetExperiment returns the Experiment field if non-nil, zero value otherwise.

### GetExperimentOk

`func (o *ExperimentToSavedMetric) GetExperimentOk() (*int32, bool)`

GetExperimentOk returns a tuple with the Experiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiment

`func (o *ExperimentToSavedMetric) SetExperiment(v int32)`

SetExperiment sets Experiment field to given value.


### GetSavedMetric

`func (o *ExperimentToSavedMetric) GetSavedMetric() int32`

GetSavedMetric returns the SavedMetric field if non-nil, zero value otherwise.

### GetSavedMetricOk

`func (o *ExperimentToSavedMetric) GetSavedMetricOk() (*int32, bool)`

GetSavedMetricOk returns a tuple with the SavedMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedMetric

`func (o *ExperimentToSavedMetric) SetSavedMetric(v int32)`

SetSavedMetric sets SavedMetric field to given value.


### GetMetadata

`func (o *ExperimentToSavedMetric) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExperimentToSavedMetric) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExperimentToSavedMetric) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExperimentToSavedMetric) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ExperimentToSavedMetric) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ExperimentToSavedMetric) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAt

`func (o *ExperimentToSavedMetric) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExperimentToSavedMetric) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExperimentToSavedMetric) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetQuery

`func (o *ExperimentToSavedMetric) GetQuery() interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ExperimentToSavedMetric) GetQueryOk() (*interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ExperimentToSavedMetric) SetQuery(v interface{})`

SetQuery sets Query field to given value.


### SetQueryNil

`func (o *ExperimentToSavedMetric) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *ExperimentToSavedMetric) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetName

`func (o *ExperimentToSavedMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentToSavedMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentToSavedMetric) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


