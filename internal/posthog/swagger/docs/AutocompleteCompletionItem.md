# AutocompleteCompletionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **NullableString** | A human-readable string with additional information about this item, like type or symbol information. | [optional] 
**Documentation** | Pointer to **NullableString** | A human-readable string that represents a doc-comment. | [optional] 
**InsertText** | **string** | A string or snippet that should be inserted in a document when selecting this completion. | 
**Kind** | [**AutocompleteCompletionItemKind**](AutocompleteCompletionItemKind.md) |  | 
**Label** | **string** | The label of this completion item. By default this is also the text that is inserted when selecting this completion. | 

## Methods

### NewAutocompleteCompletionItem

`func NewAutocompleteCompletionItem(insertText string, kind AutocompleteCompletionItemKind, label string, ) *AutocompleteCompletionItem`

NewAutocompleteCompletionItem instantiates a new AutocompleteCompletionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteCompletionItemWithDefaults

`func NewAutocompleteCompletionItemWithDefaults() *AutocompleteCompletionItem`

NewAutocompleteCompletionItemWithDefaults instantiates a new AutocompleteCompletionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *AutocompleteCompletionItem) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AutocompleteCompletionItem) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AutocompleteCompletionItem) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AutocompleteCompletionItem) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *AutocompleteCompletionItem) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *AutocompleteCompletionItem) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetDocumentation

`func (o *AutocompleteCompletionItem) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *AutocompleteCompletionItem) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *AutocompleteCompletionItem) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *AutocompleteCompletionItem) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### SetDocumentationNil

`func (o *AutocompleteCompletionItem) SetDocumentationNil(b bool)`

 SetDocumentationNil sets the value for Documentation to be an explicit nil

### UnsetDocumentation
`func (o *AutocompleteCompletionItem) UnsetDocumentation()`

UnsetDocumentation ensures that no value is present for Documentation, not even an explicit nil
### GetInsertText

`func (o *AutocompleteCompletionItem) GetInsertText() string`

GetInsertText returns the InsertText field if non-nil, zero value otherwise.

### GetInsertTextOk

`func (o *AutocompleteCompletionItem) GetInsertTextOk() (*string, bool)`

GetInsertTextOk returns a tuple with the InsertText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertText

`func (o *AutocompleteCompletionItem) SetInsertText(v string)`

SetInsertText sets InsertText field to given value.


### GetKind

`func (o *AutocompleteCompletionItem) GetKind() AutocompleteCompletionItemKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AutocompleteCompletionItem) GetKindOk() (*AutocompleteCompletionItemKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AutocompleteCompletionItem) SetKind(v AutocompleteCompletionItemKind)`

SetKind sets Kind field to given value.


### GetLabel

`func (o *AutocompleteCompletionItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AutocompleteCompletionItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AutocompleteCompletionItem) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


