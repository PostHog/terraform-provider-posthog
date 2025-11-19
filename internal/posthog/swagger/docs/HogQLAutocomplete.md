# HogQLAutocomplete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndPosition** | **int32** | End position of the editor word | 
**Filters** | Pointer to [**HogQLFilters**](HogQLFilters.md) |  | [optional] 
**Globals** | Pointer to **map[string]interface{}** | Global values in scope | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "HogQLAutocomplete"]
**Language** | [**HogLanguage**](HogLanguage.md) |  | 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Query** | **string** | Query to validate | 
**Response** | Pointer to [**HogQLAutocompleteResponse**](HogQLAutocompleteResponse.md) |  | [optional] 
**SourceQuery** | Pointer to [**NullableSourcequery**](Sourcequery.md) |  | [optional] [default to null]
**StartPosition** | **int32** | Start position of the editor word | 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewHogQLAutocomplete

`func NewHogQLAutocomplete(endPosition int32, language HogLanguage, query string, startPosition int32, ) *HogQLAutocomplete`

NewHogQLAutocomplete instantiates a new HogQLAutocomplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogQLAutocompleteWithDefaults

`func NewHogQLAutocompleteWithDefaults() *HogQLAutocomplete`

NewHogQLAutocompleteWithDefaults instantiates a new HogQLAutocomplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndPosition

`func (o *HogQLAutocomplete) GetEndPosition() int32`

GetEndPosition returns the EndPosition field if non-nil, zero value otherwise.

### GetEndPositionOk

`func (o *HogQLAutocomplete) GetEndPositionOk() (*int32, bool)`

GetEndPositionOk returns a tuple with the EndPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPosition

`func (o *HogQLAutocomplete) SetEndPosition(v int32)`

SetEndPosition sets EndPosition field to given value.


### GetFilters

`func (o *HogQLAutocomplete) GetFilters() HogQLFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *HogQLAutocomplete) GetFiltersOk() (*HogQLFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *HogQLAutocomplete) SetFilters(v HogQLFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *HogQLAutocomplete) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGlobals

`func (o *HogQLAutocomplete) GetGlobals() map[string]interface{}`

GetGlobals returns the Globals field if non-nil, zero value otherwise.

### GetGlobalsOk

`func (o *HogQLAutocomplete) GetGlobalsOk() (*map[string]interface{}, bool)`

GetGlobalsOk returns a tuple with the Globals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobals

`func (o *HogQLAutocomplete) SetGlobals(v map[string]interface{})`

SetGlobals sets Globals field to given value.

### HasGlobals

`func (o *HogQLAutocomplete) HasGlobals() bool`

HasGlobals returns a boolean if a field has been set.

### SetGlobalsNil

`func (o *HogQLAutocomplete) SetGlobalsNil(b bool)`

 SetGlobalsNil sets the value for Globals to be an explicit nil

### UnsetGlobals
`func (o *HogQLAutocomplete) UnsetGlobals()`

UnsetGlobals ensures that no value is present for Globals, not even an explicit nil
### GetKind

`func (o *HogQLAutocomplete) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HogQLAutocomplete) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HogQLAutocomplete) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HogQLAutocomplete) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLanguage

`func (o *HogQLAutocomplete) GetLanguage() HogLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *HogQLAutocomplete) GetLanguageOk() (*HogLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *HogQLAutocomplete) SetLanguage(v HogLanguage)`

SetLanguage sets Language field to given value.


### GetModifiers

`func (o *HogQLAutocomplete) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *HogQLAutocomplete) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *HogQLAutocomplete) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *HogQLAutocomplete) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetQuery

`func (o *HogQLAutocomplete) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *HogQLAutocomplete) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *HogQLAutocomplete) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetResponse

`func (o *HogQLAutocomplete) GetResponse() HogQLAutocompleteResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *HogQLAutocomplete) GetResponseOk() (*HogQLAutocompleteResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *HogQLAutocomplete) SetResponse(v HogQLAutocompleteResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *HogQLAutocomplete) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSourceQuery

`func (o *HogQLAutocomplete) GetSourceQuery() Sourcequery`

GetSourceQuery returns the SourceQuery field if non-nil, zero value otherwise.

### GetSourceQueryOk

`func (o *HogQLAutocomplete) GetSourceQueryOk() (*Sourcequery, bool)`

GetSourceQueryOk returns a tuple with the SourceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceQuery

`func (o *HogQLAutocomplete) SetSourceQuery(v Sourcequery)`

SetSourceQuery sets SourceQuery field to given value.

### HasSourceQuery

`func (o *HogQLAutocomplete) HasSourceQuery() bool`

HasSourceQuery returns a boolean if a field has been set.

### SetSourceQueryNil

`func (o *HogQLAutocomplete) SetSourceQueryNil(b bool)`

 SetSourceQueryNil sets the value for SourceQuery to be an explicit nil

### UnsetSourceQuery
`func (o *HogQLAutocomplete) UnsetSourceQuery()`

UnsetSourceQuery ensures that no value is present for SourceQuery, not even an explicit nil
### GetStartPosition

`func (o *HogQLAutocomplete) GetStartPosition() int32`

GetStartPosition returns the StartPosition field if non-nil, zero value otherwise.

### GetStartPositionOk

`func (o *HogQLAutocomplete) GetStartPositionOk() (*int32, bool)`

GetStartPositionOk returns a tuple with the StartPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPosition

`func (o *HogQLAutocomplete) SetStartPosition(v int32)`

SetStartPosition sets StartPosition field to given value.


### GetTags

`func (o *HogQLAutocomplete) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HogQLAutocomplete) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HogQLAutocomplete) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HogQLAutocomplete) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *HogQLAutocomplete) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HogQLAutocomplete) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HogQLAutocomplete) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HogQLAutocomplete) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *HogQLAutocomplete) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *HogQLAutocomplete) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


