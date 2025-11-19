# PinnedSceneTab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Pathname** | Pointer to **string** |  | [optional] 
**Search** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**CustomTitle** | Pointer to **NullableString** |  | [optional] 
**IconType** | Pointer to **string** |  | [optional] 
**SceneId** | Pointer to **NullableString** |  | [optional] 
**SceneKey** | Pointer to **NullableString** |  | [optional] 
**SceneParams** | Pointer to **interface{}** |  | [optional] 
**Pinned** | Pointer to **bool** |  | [optional] 

## Methods

### NewPinnedSceneTab

`func NewPinnedSceneTab() *PinnedSceneTab`

NewPinnedSceneTab instantiates a new PinnedSceneTab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinnedSceneTabWithDefaults

`func NewPinnedSceneTabWithDefaults() *PinnedSceneTab`

NewPinnedSceneTabWithDefaults instantiates a new PinnedSceneTab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PinnedSceneTab) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PinnedSceneTab) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PinnedSceneTab) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PinnedSceneTab) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPathname

`func (o *PinnedSceneTab) GetPathname() string`

GetPathname returns the Pathname field if non-nil, zero value otherwise.

### GetPathnameOk

`func (o *PinnedSceneTab) GetPathnameOk() (*string, bool)`

GetPathnameOk returns a tuple with the Pathname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathname

`func (o *PinnedSceneTab) SetPathname(v string)`

SetPathname sets Pathname field to given value.

### HasPathname

`func (o *PinnedSceneTab) HasPathname() bool`

HasPathname returns a boolean if a field has been set.

### GetSearch

`func (o *PinnedSceneTab) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *PinnedSceneTab) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *PinnedSceneTab) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *PinnedSceneTab) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetHash

`func (o *PinnedSceneTab) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *PinnedSceneTab) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *PinnedSceneTab) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *PinnedSceneTab) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetTitle

`func (o *PinnedSceneTab) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PinnedSceneTab) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PinnedSceneTab) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PinnedSceneTab) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCustomTitle

`func (o *PinnedSceneTab) GetCustomTitle() string`

GetCustomTitle returns the CustomTitle field if non-nil, zero value otherwise.

### GetCustomTitleOk

`func (o *PinnedSceneTab) GetCustomTitleOk() (*string, bool)`

GetCustomTitleOk returns a tuple with the CustomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTitle

`func (o *PinnedSceneTab) SetCustomTitle(v string)`

SetCustomTitle sets CustomTitle field to given value.

### HasCustomTitle

`func (o *PinnedSceneTab) HasCustomTitle() bool`

HasCustomTitle returns a boolean if a field has been set.

### SetCustomTitleNil

`func (o *PinnedSceneTab) SetCustomTitleNil(b bool)`

 SetCustomTitleNil sets the value for CustomTitle to be an explicit nil

### UnsetCustomTitle
`func (o *PinnedSceneTab) UnsetCustomTitle()`

UnsetCustomTitle ensures that no value is present for CustomTitle, not even an explicit nil
### GetIconType

`func (o *PinnedSceneTab) GetIconType() string`

GetIconType returns the IconType field if non-nil, zero value otherwise.

### GetIconTypeOk

`func (o *PinnedSceneTab) GetIconTypeOk() (*string, bool)`

GetIconTypeOk returns a tuple with the IconType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconType

`func (o *PinnedSceneTab) SetIconType(v string)`

SetIconType sets IconType field to given value.

### HasIconType

`func (o *PinnedSceneTab) HasIconType() bool`

HasIconType returns a boolean if a field has been set.

### GetSceneId

`func (o *PinnedSceneTab) GetSceneId() string`

GetSceneId returns the SceneId field if non-nil, zero value otherwise.

### GetSceneIdOk

`func (o *PinnedSceneTab) GetSceneIdOk() (*string, bool)`

GetSceneIdOk returns a tuple with the SceneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneId

`func (o *PinnedSceneTab) SetSceneId(v string)`

SetSceneId sets SceneId field to given value.

### HasSceneId

`func (o *PinnedSceneTab) HasSceneId() bool`

HasSceneId returns a boolean if a field has been set.

### SetSceneIdNil

`func (o *PinnedSceneTab) SetSceneIdNil(b bool)`

 SetSceneIdNil sets the value for SceneId to be an explicit nil

### UnsetSceneId
`func (o *PinnedSceneTab) UnsetSceneId()`

UnsetSceneId ensures that no value is present for SceneId, not even an explicit nil
### GetSceneKey

`func (o *PinnedSceneTab) GetSceneKey() string`

GetSceneKey returns the SceneKey field if non-nil, zero value otherwise.

### GetSceneKeyOk

`func (o *PinnedSceneTab) GetSceneKeyOk() (*string, bool)`

GetSceneKeyOk returns a tuple with the SceneKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneKey

`func (o *PinnedSceneTab) SetSceneKey(v string)`

SetSceneKey sets SceneKey field to given value.

### HasSceneKey

`func (o *PinnedSceneTab) HasSceneKey() bool`

HasSceneKey returns a boolean if a field has been set.

### SetSceneKeyNil

`func (o *PinnedSceneTab) SetSceneKeyNil(b bool)`

 SetSceneKeyNil sets the value for SceneKey to be an explicit nil

### UnsetSceneKey
`func (o *PinnedSceneTab) UnsetSceneKey()`

UnsetSceneKey ensures that no value is present for SceneKey, not even an explicit nil
### GetSceneParams

`func (o *PinnedSceneTab) GetSceneParams() interface{}`

GetSceneParams returns the SceneParams field if non-nil, zero value otherwise.

### GetSceneParamsOk

`func (o *PinnedSceneTab) GetSceneParamsOk() (*interface{}, bool)`

GetSceneParamsOk returns a tuple with the SceneParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneParams

`func (o *PinnedSceneTab) SetSceneParams(v interface{})`

SetSceneParams sets SceneParams field to given value.

### HasSceneParams

`func (o *PinnedSceneTab) HasSceneParams() bool`

HasSceneParams returns a boolean if a field has been set.

### SetSceneParamsNil

`func (o *PinnedSceneTab) SetSceneParamsNil(b bool)`

 SetSceneParamsNil sets the value for SceneParams to be an explicit nil

### UnsetSceneParams
`func (o *PinnedSceneTab) UnsetSceneParams()`

UnsetSceneParams ensures that no value is present for SceneParams, not even an explicit nil
### GetPinned

`func (o *PinnedSceneTab) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *PinnedSceneTab) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *PinnedSceneTab) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *PinnedSceneTab) HasPinned() bool`

HasPinned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


