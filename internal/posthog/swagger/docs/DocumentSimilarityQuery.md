# DocumentSimilarityQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRange** | [**DateRange**](DateRange.md) |  | 
**DistanceFunc** | [**DistanceFunc**](DistanceFunc.md) |  | 
**DocumentTypes** | **[]string** |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "DocumentSimilarityQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Model** | **string** |  | 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**OrderBy** | [**OrderBy**](OrderBy.md) |  | 
**OrderDirection** | [**OrderDirection**](OrderDirection.md) |  | 
**Origin** | [**EmbeddedDocument**](EmbeddedDocument.md) |  | 
**Products** | **[]string** |  | 
**Renderings** | **[]string** |  | 
**Response** | Pointer to [**DocumentSimilarityQueryResponse**](DocumentSimilarityQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Threshold** | Pointer to **NullableFloat32** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewDocumentSimilarityQuery

`func NewDocumentSimilarityQuery(dateRange DateRange, distanceFunc DistanceFunc, documentTypes []string, model string, orderBy OrderBy, orderDirection OrderDirection, origin EmbeddedDocument, products []string, renderings []string, ) *DocumentSimilarityQuery`

NewDocumentSimilarityQuery instantiates a new DocumentSimilarityQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentSimilarityQueryWithDefaults

`func NewDocumentSimilarityQueryWithDefaults() *DocumentSimilarityQuery`

NewDocumentSimilarityQueryWithDefaults instantiates a new DocumentSimilarityQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRange

`func (o *DocumentSimilarityQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *DocumentSimilarityQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *DocumentSimilarityQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetDistanceFunc

`func (o *DocumentSimilarityQuery) GetDistanceFunc() DistanceFunc`

GetDistanceFunc returns the DistanceFunc field if non-nil, zero value otherwise.

### GetDistanceFuncOk

`func (o *DocumentSimilarityQuery) GetDistanceFuncOk() (*DistanceFunc, bool)`

GetDistanceFuncOk returns a tuple with the DistanceFunc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceFunc

`func (o *DocumentSimilarityQuery) SetDistanceFunc(v DistanceFunc)`

SetDistanceFunc sets DistanceFunc field to given value.


### GetDocumentTypes

`func (o *DocumentSimilarityQuery) GetDocumentTypes() []string`

GetDocumentTypes returns the DocumentTypes field if non-nil, zero value otherwise.

### GetDocumentTypesOk

`func (o *DocumentSimilarityQuery) GetDocumentTypesOk() (*[]string, bool)`

GetDocumentTypesOk returns a tuple with the DocumentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypes

`func (o *DocumentSimilarityQuery) SetDocumentTypes(v []string)`

SetDocumentTypes sets DocumentTypes field to given value.


### GetKind

`func (o *DocumentSimilarityQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DocumentSimilarityQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DocumentSimilarityQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DocumentSimilarityQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *DocumentSimilarityQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DocumentSimilarityQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DocumentSimilarityQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DocumentSimilarityQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *DocumentSimilarityQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *DocumentSimilarityQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModel

`func (o *DocumentSimilarityQuery) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DocumentSimilarityQuery) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DocumentSimilarityQuery) SetModel(v string)`

SetModel sets Model field to given value.


### GetModifiers

`func (o *DocumentSimilarityQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *DocumentSimilarityQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *DocumentSimilarityQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *DocumentSimilarityQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *DocumentSimilarityQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DocumentSimilarityQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DocumentSimilarityQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DocumentSimilarityQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *DocumentSimilarityQuery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *DocumentSimilarityQuery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOrderBy

`func (o *DocumentSimilarityQuery) GetOrderBy() OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *DocumentSimilarityQuery) GetOrderByOk() (*OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *DocumentSimilarityQuery) SetOrderBy(v OrderBy)`

SetOrderBy sets OrderBy field to given value.


### GetOrderDirection

`func (o *DocumentSimilarityQuery) GetOrderDirection() OrderDirection`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *DocumentSimilarityQuery) GetOrderDirectionOk() (*OrderDirection, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *DocumentSimilarityQuery) SetOrderDirection(v OrderDirection)`

SetOrderDirection sets OrderDirection field to given value.


### GetOrigin

`func (o *DocumentSimilarityQuery) GetOrigin() EmbeddedDocument`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DocumentSimilarityQuery) GetOriginOk() (*EmbeddedDocument, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DocumentSimilarityQuery) SetOrigin(v EmbeddedDocument)`

SetOrigin sets Origin field to given value.


### GetProducts

`func (o *DocumentSimilarityQuery) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *DocumentSimilarityQuery) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *DocumentSimilarityQuery) SetProducts(v []string)`

SetProducts sets Products field to given value.


### GetRenderings

`func (o *DocumentSimilarityQuery) GetRenderings() []string`

GetRenderings returns the Renderings field if non-nil, zero value otherwise.

### GetRenderingsOk

`func (o *DocumentSimilarityQuery) GetRenderingsOk() (*[]string, bool)`

GetRenderingsOk returns a tuple with the Renderings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderings

`func (o *DocumentSimilarityQuery) SetRenderings(v []string)`

SetRenderings sets Renderings field to given value.


### GetResponse

`func (o *DocumentSimilarityQuery) GetResponse() DocumentSimilarityQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *DocumentSimilarityQuery) GetResponseOk() (*DocumentSimilarityQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *DocumentSimilarityQuery) SetResponse(v DocumentSimilarityQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *DocumentSimilarityQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *DocumentSimilarityQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DocumentSimilarityQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DocumentSimilarityQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DocumentSimilarityQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreshold

`func (o *DocumentSimilarityQuery) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *DocumentSimilarityQuery) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *DocumentSimilarityQuery) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *DocumentSimilarityQuery) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### SetThresholdNil

`func (o *DocumentSimilarityQuery) SetThresholdNil(b bool)`

 SetThresholdNil sets the value for Threshold to be an explicit nil

### UnsetThreshold
`func (o *DocumentSimilarityQuery) UnsetThreshold()`

UnsetThreshold ensures that no value is present for Threshold, not even an explicit nil
### GetVersion

`func (o *DocumentSimilarityQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DocumentSimilarityQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DocumentSimilarityQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DocumentSimilarityQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DocumentSimilarityQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DocumentSimilarityQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


