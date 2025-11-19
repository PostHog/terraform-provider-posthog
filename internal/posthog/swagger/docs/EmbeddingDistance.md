# EmbeddingDistance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distance** | **float32** |  | 
**Origin** | Pointer to [**EmbeddingRecord**](EmbeddingRecord.md) |  | [optional] 
**Result** | [**EmbeddingRecord**](EmbeddingRecord.md) |  | 

## Methods

### NewEmbeddingDistance

`func NewEmbeddingDistance(distance float32, result EmbeddingRecord, ) *EmbeddingDistance`

NewEmbeddingDistance instantiates a new EmbeddingDistance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingDistanceWithDefaults

`func NewEmbeddingDistanceWithDefaults() *EmbeddingDistance`

NewEmbeddingDistanceWithDefaults instantiates a new EmbeddingDistance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistance

`func (o *EmbeddingDistance) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *EmbeddingDistance) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *EmbeddingDistance) SetDistance(v float32)`

SetDistance sets Distance field to given value.


### GetOrigin

`func (o *EmbeddingDistance) GetOrigin() EmbeddingRecord`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *EmbeddingDistance) GetOriginOk() (*EmbeddingRecord, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *EmbeddingDistance) SetOrigin(v EmbeddingRecord)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *EmbeddingDistance) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetResult

`func (o *EmbeddingDistance) GetResult() EmbeddingRecord`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EmbeddingDistance) GetResultOk() (*EmbeddingRecord, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EmbeddingDistance) SetResult(v EmbeddingRecord)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


