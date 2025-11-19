# ErrorTrackingSimilarIssuesQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**IssueId** | **string** |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "ErrorTrackingSimilarIssuesQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**MaxDistance** | Pointer to **NullableFloat32** |  | [optional] 
**ModelName** | Pointer to [**EmbeddingModelName**](EmbeddingModelName.md) |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**Rendering** | Pointer to **NullableString** |  | [optional] 
**Response** | Pointer to [**ErrorTrackingSimilarIssuesQueryResponse**](ErrorTrackingSimilarIssuesQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewErrorTrackingSimilarIssuesQuery

`func NewErrorTrackingSimilarIssuesQuery(issueId string, ) *ErrorTrackingSimilarIssuesQuery`

NewErrorTrackingSimilarIssuesQuery instantiates a new ErrorTrackingSimilarIssuesQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingSimilarIssuesQueryWithDefaults

`func NewErrorTrackingSimilarIssuesQueryWithDefaults() *ErrorTrackingSimilarIssuesQuery`

NewErrorTrackingSimilarIssuesQueryWithDefaults instantiates a new ErrorTrackingSimilarIssuesQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRange

`func (o *ErrorTrackingSimilarIssuesQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *ErrorTrackingSimilarIssuesQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *ErrorTrackingSimilarIssuesQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetIssueId

`func (o *ErrorTrackingSimilarIssuesQuery) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *ErrorTrackingSimilarIssuesQuery) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.


### GetKind

`func (o *ErrorTrackingSimilarIssuesQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ErrorTrackingSimilarIssuesQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ErrorTrackingSimilarIssuesQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *ErrorTrackingSimilarIssuesQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ErrorTrackingSimilarIssuesQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ErrorTrackingSimilarIssuesQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ErrorTrackingSimilarIssuesQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ErrorTrackingSimilarIssuesQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetMaxDistance

`func (o *ErrorTrackingSimilarIssuesQuery) GetMaxDistance() float32`

GetMaxDistance returns the MaxDistance field if non-nil, zero value otherwise.

### GetMaxDistanceOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetMaxDistanceOk() (*float32, bool)`

GetMaxDistanceOk returns a tuple with the MaxDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistance

`func (o *ErrorTrackingSimilarIssuesQuery) SetMaxDistance(v float32)`

SetMaxDistance sets MaxDistance field to given value.

### HasMaxDistance

`func (o *ErrorTrackingSimilarIssuesQuery) HasMaxDistance() bool`

HasMaxDistance returns a boolean if a field has been set.

### SetMaxDistanceNil

`func (o *ErrorTrackingSimilarIssuesQuery) SetMaxDistanceNil(b bool)`

 SetMaxDistanceNil sets the value for MaxDistance to be an explicit nil

### UnsetMaxDistance
`func (o *ErrorTrackingSimilarIssuesQuery) UnsetMaxDistance()`

UnsetMaxDistance ensures that no value is present for MaxDistance, not even an explicit nil
### GetModelName

`func (o *ErrorTrackingSimilarIssuesQuery) GetModelName() EmbeddingModelName`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetModelNameOk() (*EmbeddingModelName, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ErrorTrackingSimilarIssuesQuery) SetModelName(v EmbeddingModelName)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *ErrorTrackingSimilarIssuesQuery) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModifiers

`func (o *ErrorTrackingSimilarIssuesQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *ErrorTrackingSimilarIssuesQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *ErrorTrackingSimilarIssuesQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *ErrorTrackingSimilarIssuesQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ErrorTrackingSimilarIssuesQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ErrorTrackingSimilarIssuesQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *ErrorTrackingSimilarIssuesQuery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *ErrorTrackingSimilarIssuesQuery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetRendering

`func (o *ErrorTrackingSimilarIssuesQuery) GetRendering() string`

GetRendering returns the Rendering field if non-nil, zero value otherwise.

### GetRenderingOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetRenderingOk() (*string, bool)`

GetRenderingOk returns a tuple with the Rendering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendering

`func (o *ErrorTrackingSimilarIssuesQuery) SetRendering(v string)`

SetRendering sets Rendering field to given value.

### HasRendering

`func (o *ErrorTrackingSimilarIssuesQuery) HasRendering() bool`

HasRendering returns a boolean if a field has been set.

### SetRenderingNil

`func (o *ErrorTrackingSimilarIssuesQuery) SetRenderingNil(b bool)`

 SetRenderingNil sets the value for Rendering to be an explicit nil

### UnsetRendering
`func (o *ErrorTrackingSimilarIssuesQuery) UnsetRendering()`

UnsetRendering ensures that no value is present for Rendering, not even an explicit nil
### GetResponse

`func (o *ErrorTrackingSimilarIssuesQuery) GetResponse() ErrorTrackingSimilarIssuesQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetResponseOk() (*ErrorTrackingSimilarIssuesQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ErrorTrackingSimilarIssuesQuery) SetResponse(v ErrorTrackingSimilarIssuesQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ErrorTrackingSimilarIssuesQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *ErrorTrackingSimilarIssuesQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ErrorTrackingSimilarIssuesQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ErrorTrackingSimilarIssuesQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *ErrorTrackingSimilarIssuesQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ErrorTrackingSimilarIssuesQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ErrorTrackingSimilarIssuesQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ErrorTrackingSimilarIssuesQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ErrorTrackingSimilarIssuesQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ErrorTrackingSimilarIssuesQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


