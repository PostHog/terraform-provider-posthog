# EmbeddingRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** |  | 
**DocumentType** | **string** |  | 
**ModelName** | [**EmbeddingModelName**](EmbeddingModelName.md) |  | 
**Product** | **string** |  | 
**Rendering** | **string** |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewEmbeddingRecord

`func NewEmbeddingRecord(documentId string, documentType string, modelName EmbeddingModelName, product string, rendering string, timestamp time.Time, ) *EmbeddingRecord`

NewEmbeddingRecord instantiates a new EmbeddingRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingRecordWithDefaults

`func NewEmbeddingRecordWithDefaults() *EmbeddingRecord`

NewEmbeddingRecordWithDefaults instantiates a new EmbeddingRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *EmbeddingRecord) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *EmbeddingRecord) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *EmbeddingRecord) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetDocumentType

`func (o *EmbeddingRecord) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *EmbeddingRecord) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *EmbeddingRecord) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### GetModelName

`func (o *EmbeddingRecord) GetModelName() EmbeddingModelName`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *EmbeddingRecord) GetModelNameOk() (*EmbeddingModelName, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *EmbeddingRecord) SetModelName(v EmbeddingModelName)`

SetModelName sets ModelName field to given value.


### GetProduct

`func (o *EmbeddingRecord) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EmbeddingRecord) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EmbeddingRecord) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetRendering

`func (o *EmbeddingRecord) GetRendering() string`

GetRendering returns the Rendering field if non-nil, zero value otherwise.

### GetRenderingOk

`func (o *EmbeddingRecord) GetRenderingOk() (*string, bool)`

GetRenderingOk returns a tuple with the Rendering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendering

`func (o *EmbeddingRecord) SetRendering(v string)`

SetRendering sets Rendering field to given value.


### GetTimestamp

`func (o *EmbeddingRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmbeddingRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmbeddingRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


