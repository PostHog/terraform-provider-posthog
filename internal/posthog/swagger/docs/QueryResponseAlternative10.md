# QueryResponseAlternative10

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncompleteList** | **bool** | Whether or not the suggestions returned are complete | 
**Suggestions** | [**[]AutocompleteCompletionItem**](AutocompleteCompletionItem.md) |  | 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 

## Methods

### NewQueryResponseAlternative10

`func NewQueryResponseAlternative10(incompleteList bool, suggestions []AutocompleteCompletionItem, ) *QueryResponseAlternative10`

NewQueryResponseAlternative10 instantiates a new QueryResponseAlternative10 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternative10WithDefaults

`func NewQueryResponseAlternative10WithDefaults() *QueryResponseAlternative10`

NewQueryResponseAlternative10WithDefaults instantiates a new QueryResponseAlternative10 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncompleteList

`func (o *QueryResponseAlternative10) GetIncompleteList() bool`

GetIncompleteList returns the IncompleteList field if non-nil, zero value otherwise.

### GetIncompleteListOk

`func (o *QueryResponseAlternative10) GetIncompleteListOk() (*bool, bool)`

GetIncompleteListOk returns a tuple with the IncompleteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteList

`func (o *QueryResponseAlternative10) SetIncompleteList(v bool)`

SetIncompleteList sets IncompleteList field to given value.


### GetSuggestions

`func (o *QueryResponseAlternative10) GetSuggestions() []AutocompleteCompletionItem`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *QueryResponseAlternative10) GetSuggestionsOk() (*[]AutocompleteCompletionItem, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *QueryResponseAlternative10) SetSuggestions(v []AutocompleteCompletionItem)`

SetSuggestions sets Suggestions field to given value.


### GetTimings

`func (o *QueryResponseAlternative10) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *QueryResponseAlternative10) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *QueryResponseAlternative10) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *QueryResponseAlternative10) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *QueryResponseAlternative10) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *QueryResponseAlternative10) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


