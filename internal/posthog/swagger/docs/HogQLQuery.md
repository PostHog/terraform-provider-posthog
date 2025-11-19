# HogQLQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Explain** | Pointer to **NullableBool** |  | [optional] 
**Filters** | Pointer to [**HogQLFilters**](HogQLFilters.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "HogQLQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Client provided name of the query | [optional] 
**Query** | **string** |  | 
**Response** | Pointer to [**HogQLQueryResponse**](HogQLQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Values** | Pointer to **map[string]interface{}** | Constant values that can be referenced with the {placeholder} syntax in the query | [optional] 
**Variables** | Pointer to [**map[string]HogQLVariable**](HogQLVariable.md) | Variables to be substituted into the query | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewHogQLQuery

`func NewHogQLQuery(query string, ) *HogQLQuery`

NewHogQLQuery instantiates a new HogQLQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogQLQueryWithDefaults

`func NewHogQLQueryWithDefaults() *HogQLQuery`

NewHogQLQueryWithDefaults instantiates a new HogQLQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExplain

`func (o *HogQLQuery) GetExplain() bool`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *HogQLQuery) GetExplainOk() (*bool, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *HogQLQuery) SetExplain(v bool)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *HogQLQuery) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *HogQLQuery) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *HogQLQuery) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetFilters

`func (o *HogQLQuery) GetFilters() HogQLFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *HogQLQuery) GetFiltersOk() (*HogQLFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *HogQLQuery) SetFilters(v HogQLFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *HogQLQuery) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetKind

`func (o *HogQLQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HogQLQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HogQLQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HogQLQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *HogQLQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *HogQLQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *HogQLQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *HogQLQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetName

`func (o *HogQLQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HogQLQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HogQLQuery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HogQLQuery) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HogQLQuery) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HogQLQuery) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuery

`func (o *HogQLQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *HogQLQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *HogQLQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetResponse

`func (o *HogQLQuery) GetResponse() HogQLQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *HogQLQuery) GetResponseOk() (*HogQLQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *HogQLQuery) SetResponse(v HogQLQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *HogQLQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *HogQLQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HogQLQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HogQLQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HogQLQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetValues

`func (o *HogQLQuery) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HogQLQuery) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HogQLQuery) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *HogQLQuery) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *HogQLQuery) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *HogQLQuery) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetVariables

`func (o *HogQLQuery) GetVariables() map[string]HogQLVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *HogQLQuery) GetVariablesOk() (*map[string]HogQLVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *HogQLQuery) SetVariables(v map[string]HogQLVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *HogQLQuery) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *HogQLQuery) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *HogQLQuery) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetVersion

`func (o *HogQLQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HogQLQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HogQLQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HogQLQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *HogQLQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *HogQLQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


