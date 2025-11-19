# HogQLASTQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Explain** | Pointer to **NullableBool** |  | [optional] 
**Filters** | Pointer to [**HogQLFilters**](HogQLFilters.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "HogQLASTQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Client provided name of the query | [optional] 
**Query** | **map[string]interface{}** |  | 
**Response** | Pointer to [**HogQLQueryResponse**](HogQLQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Values** | Pointer to **map[string]interface{}** | Constant values that can be referenced with the {placeholder} syntax in the query | [optional] 
**Variables** | Pointer to [**map[string]HogQLVariable**](HogQLVariable.md) | Variables to be substituted into the query | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewHogQLASTQuery

`func NewHogQLASTQuery(query map[string]interface{}, ) *HogQLASTQuery`

NewHogQLASTQuery instantiates a new HogQLASTQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogQLASTQueryWithDefaults

`func NewHogQLASTQueryWithDefaults() *HogQLASTQuery`

NewHogQLASTQueryWithDefaults instantiates a new HogQLASTQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExplain

`func (o *HogQLASTQuery) GetExplain() bool`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *HogQLASTQuery) GetExplainOk() (*bool, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *HogQLASTQuery) SetExplain(v bool)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *HogQLASTQuery) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *HogQLASTQuery) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *HogQLASTQuery) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetFilters

`func (o *HogQLASTQuery) GetFilters() HogQLFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *HogQLASTQuery) GetFiltersOk() (*HogQLFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *HogQLASTQuery) SetFilters(v HogQLFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *HogQLASTQuery) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetKind

`func (o *HogQLASTQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HogQLASTQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HogQLASTQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HogQLASTQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *HogQLASTQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *HogQLASTQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *HogQLASTQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *HogQLASTQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetName

`func (o *HogQLASTQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HogQLASTQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HogQLASTQuery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HogQLASTQuery) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HogQLASTQuery) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HogQLASTQuery) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuery

`func (o *HogQLASTQuery) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *HogQLASTQuery) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *HogQLASTQuery) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.


### GetResponse

`func (o *HogQLASTQuery) GetResponse() HogQLQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *HogQLASTQuery) GetResponseOk() (*HogQLQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *HogQLASTQuery) SetResponse(v HogQLQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *HogQLASTQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *HogQLASTQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HogQLASTQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HogQLASTQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HogQLASTQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetValues

`func (o *HogQLASTQuery) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HogQLASTQuery) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HogQLASTQuery) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *HogQLASTQuery) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *HogQLASTQuery) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *HogQLASTQuery) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetVariables

`func (o *HogQLASTQuery) GetVariables() map[string]HogQLVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *HogQLASTQuery) GetVariablesOk() (*map[string]HogQLVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *HogQLASTQuery) SetVariables(v map[string]HogQLVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *HogQLASTQuery) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *HogQLASTQuery) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *HogQLASTQuery) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetVersion

`func (o *HogQLASTQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HogQLASTQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HogQLASTQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HogQLASTQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *HogQLASTQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *HogQLASTQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


