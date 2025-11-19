# EmbeddedDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** |  | 
**DocumentType** | **string** |  | 
**Product** | **string** |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewEmbeddedDocument

`func NewEmbeddedDocument(documentId string, documentType string, product string, timestamp time.Time, ) *EmbeddedDocument`

NewEmbeddedDocument instantiates a new EmbeddedDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedDocumentWithDefaults

`func NewEmbeddedDocumentWithDefaults() *EmbeddedDocument`

NewEmbeddedDocumentWithDefaults instantiates a new EmbeddedDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *EmbeddedDocument) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *EmbeddedDocument) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *EmbeddedDocument) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetDocumentType

`func (o *EmbeddedDocument) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *EmbeddedDocument) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *EmbeddedDocument) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### GetProduct

`func (o *EmbeddedDocument) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EmbeddedDocument) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EmbeddedDocument) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetTimestamp

`func (o *EmbeddedDocument) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmbeddedDocument) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmbeddedDocument) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


