# HogQLMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Debug** | Pointer to **NullableBool** | Enable more verbose output, usually run from the /debug page | [optional] 
**Filters** | Pointer to [**HogQLFilters**](HogQLFilters.md) |  | [optional] 
**Globals** | Pointer to **map[string]interface{}** | Extra globals for the query | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "HogQLMetadata"]
**Language** | [**HogLanguage**](HogLanguage.md) |  | 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Query** | **string** | Query to validate | 
**Response** | Pointer to [**HogQLMetadataResponse**](HogQLMetadataResponse.md) |  | [optional] 
**SourceQuery** | Pointer to [**NullableSourcequery1**](Sourcequery1.md) |  | [optional] [default to null]
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Variables** | Pointer to [**map[string]HogQLVariable**](HogQLVariable.md) | Variables to be subsituted into the query | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewHogQLMetadata

`func NewHogQLMetadata(language HogLanguage, query string, ) *HogQLMetadata`

NewHogQLMetadata instantiates a new HogQLMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogQLMetadataWithDefaults

`func NewHogQLMetadataWithDefaults() *HogQLMetadata`

NewHogQLMetadataWithDefaults instantiates a new HogQLMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebug

`func (o *HogQLMetadata) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *HogQLMetadata) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *HogQLMetadata) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *HogQLMetadata) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### SetDebugNil

`func (o *HogQLMetadata) SetDebugNil(b bool)`

 SetDebugNil sets the value for Debug to be an explicit nil

### UnsetDebug
`func (o *HogQLMetadata) UnsetDebug()`

UnsetDebug ensures that no value is present for Debug, not even an explicit nil
### GetFilters

`func (o *HogQLMetadata) GetFilters() HogQLFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *HogQLMetadata) GetFiltersOk() (*HogQLFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *HogQLMetadata) SetFilters(v HogQLFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *HogQLMetadata) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGlobals

`func (o *HogQLMetadata) GetGlobals() map[string]interface{}`

GetGlobals returns the Globals field if non-nil, zero value otherwise.

### GetGlobalsOk

`func (o *HogQLMetadata) GetGlobalsOk() (*map[string]interface{}, bool)`

GetGlobalsOk returns a tuple with the Globals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobals

`func (o *HogQLMetadata) SetGlobals(v map[string]interface{})`

SetGlobals sets Globals field to given value.

### HasGlobals

`func (o *HogQLMetadata) HasGlobals() bool`

HasGlobals returns a boolean if a field has been set.

### SetGlobalsNil

`func (o *HogQLMetadata) SetGlobalsNil(b bool)`

 SetGlobalsNil sets the value for Globals to be an explicit nil

### UnsetGlobals
`func (o *HogQLMetadata) UnsetGlobals()`

UnsetGlobals ensures that no value is present for Globals, not even an explicit nil
### GetKind

`func (o *HogQLMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HogQLMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HogQLMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HogQLMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLanguage

`func (o *HogQLMetadata) GetLanguage() HogLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *HogQLMetadata) GetLanguageOk() (*HogLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *HogQLMetadata) SetLanguage(v HogLanguage)`

SetLanguage sets Language field to given value.


### GetModifiers

`func (o *HogQLMetadata) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *HogQLMetadata) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *HogQLMetadata) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *HogQLMetadata) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetQuery

`func (o *HogQLMetadata) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *HogQLMetadata) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *HogQLMetadata) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetResponse

`func (o *HogQLMetadata) GetResponse() HogQLMetadataResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *HogQLMetadata) GetResponseOk() (*HogQLMetadataResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *HogQLMetadata) SetResponse(v HogQLMetadataResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *HogQLMetadata) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSourceQuery

`func (o *HogQLMetadata) GetSourceQuery() Sourcequery1`

GetSourceQuery returns the SourceQuery field if non-nil, zero value otherwise.

### GetSourceQueryOk

`func (o *HogQLMetadata) GetSourceQueryOk() (*Sourcequery1, bool)`

GetSourceQueryOk returns a tuple with the SourceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceQuery

`func (o *HogQLMetadata) SetSourceQuery(v Sourcequery1)`

SetSourceQuery sets SourceQuery field to given value.

### HasSourceQuery

`func (o *HogQLMetadata) HasSourceQuery() bool`

HasSourceQuery returns a boolean if a field has been set.

### SetSourceQueryNil

`func (o *HogQLMetadata) SetSourceQueryNil(b bool)`

 SetSourceQueryNil sets the value for SourceQuery to be an explicit nil

### UnsetSourceQuery
`func (o *HogQLMetadata) UnsetSourceQuery()`

UnsetSourceQuery ensures that no value is present for SourceQuery, not even an explicit nil
### GetTags

`func (o *HogQLMetadata) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HogQLMetadata) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HogQLMetadata) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HogQLMetadata) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVariables

`func (o *HogQLMetadata) GetVariables() map[string]HogQLVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *HogQLMetadata) GetVariablesOk() (*map[string]HogQLVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *HogQLMetadata) SetVariables(v map[string]HogQLVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *HogQLMetadata) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *HogQLMetadata) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *HogQLMetadata) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetVersion

`func (o *HogQLMetadata) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HogQLMetadata) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HogQLMetadata) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HogQLMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *HogQLMetadata) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *HogQLMetadata) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


