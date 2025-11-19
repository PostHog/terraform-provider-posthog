# PatchedNotebook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ShortId** | Pointer to **string** |  | [optional] [readonly] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **interface{}** |  | [optional] 
**TextContent** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**LastModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastModifiedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**UserAccessLevel** | Pointer to **NullableString** | The effective access level the user has for this object | [optional] [readonly] 
**CreateInFolder** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedNotebook

`func NewPatchedNotebook() *PatchedNotebook`

NewPatchedNotebook instantiates a new PatchedNotebook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedNotebookWithDefaults

`func NewPatchedNotebookWithDefaults() *PatchedNotebook`

NewPatchedNotebookWithDefaults instantiates a new PatchedNotebook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedNotebook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedNotebook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedNotebook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedNotebook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShortId

`func (o *PatchedNotebook) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *PatchedNotebook) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *PatchedNotebook) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *PatchedNotebook) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetTitle

`func (o *PatchedNotebook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PatchedNotebook) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PatchedNotebook) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PatchedNotebook) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PatchedNotebook) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PatchedNotebook) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetContent

`func (o *PatchedNotebook) GetContent() interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PatchedNotebook) GetContentOk() (*interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PatchedNotebook) SetContent(v interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *PatchedNotebook) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *PatchedNotebook) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *PatchedNotebook) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetTextContent

`func (o *PatchedNotebook) GetTextContent() string`

GetTextContent returns the TextContent field if non-nil, zero value otherwise.

### GetTextContentOk

`func (o *PatchedNotebook) GetTextContentOk() (*string, bool)`

GetTextContentOk returns a tuple with the TextContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContent

`func (o *PatchedNotebook) SetTextContent(v string)`

SetTextContent sets TextContent field to given value.

### HasTextContent

`func (o *PatchedNotebook) HasTextContent() bool`

HasTextContent returns a boolean if a field has been set.

### SetTextContentNil

`func (o *PatchedNotebook) SetTextContentNil(b bool)`

 SetTextContentNil sets the value for TextContent to be an explicit nil

### UnsetTextContent
`func (o *PatchedNotebook) UnsetTextContent()`

UnsetTextContent ensures that no value is present for TextContent, not even an explicit nil
### GetVersion

`func (o *PatchedNotebook) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchedNotebook) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchedNotebook) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchedNotebook) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedNotebook) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedNotebook) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedNotebook) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedNotebook) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedNotebook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedNotebook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedNotebook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedNotebook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedNotebook) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedNotebook) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedNotebook) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedNotebook) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *PatchedNotebook) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *PatchedNotebook) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *PatchedNotebook) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *PatchedNotebook) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *PatchedNotebook) GetLastModifiedBy() UserBasic`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *PatchedNotebook) GetLastModifiedByOk() (*UserBasic, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *PatchedNotebook) SetLastModifiedBy(v UserBasic)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *PatchedNotebook) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetUserAccessLevel

`func (o *PatchedNotebook) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *PatchedNotebook) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *PatchedNotebook) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.

### HasUserAccessLevel

`func (o *PatchedNotebook) HasUserAccessLevel() bool`

HasUserAccessLevel returns a boolean if a field has been set.

### SetUserAccessLevelNil

`func (o *PatchedNotebook) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *PatchedNotebook) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil
### GetCreateInFolder

`func (o *PatchedNotebook) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *PatchedNotebook) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *PatchedNotebook) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *PatchedNotebook) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


