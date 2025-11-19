# SimilarIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Distance** | **float32** |  | 
**FirstSeen** | **string** |  | 
**Id** | **string** |  | 
**Library** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewSimilarIssue

`func NewSimilarIssue(description string, distance float32, firstSeen string, id string, name string, status string, ) *SimilarIssue`

NewSimilarIssue instantiates a new SimilarIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimilarIssueWithDefaults

`func NewSimilarIssueWithDefaults() *SimilarIssue`

NewSimilarIssueWithDefaults instantiates a new SimilarIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SimilarIssue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimilarIssue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimilarIssue) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDistance

`func (o *SimilarIssue) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *SimilarIssue) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *SimilarIssue) SetDistance(v float32)`

SetDistance sets Distance field to given value.


### GetFirstSeen

`func (o *SimilarIssue) GetFirstSeen() string`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *SimilarIssue) GetFirstSeenOk() (*string, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *SimilarIssue) SetFirstSeen(v string)`

SetFirstSeen sets FirstSeen field to given value.


### GetId

`func (o *SimilarIssue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimilarIssue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimilarIssue) SetId(v string)`

SetId sets Id field to given value.


### GetLibrary

`func (o *SimilarIssue) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *SimilarIssue) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *SimilarIssue) SetLibrary(v string)`

SetLibrary sets Library field to given value.

### HasLibrary

`func (o *SimilarIssue) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.

### SetLibraryNil

`func (o *SimilarIssue) SetLibraryNil(b bool)`

 SetLibraryNil sets the value for Library to be an explicit nil

### UnsetLibrary
`func (o *SimilarIssue) UnsetLibrary()`

UnsetLibrary ensures that no value is present for Library, not even an explicit nil
### GetName

`func (o *SimilarIssue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimilarIssue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimilarIssue) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *SimilarIssue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimilarIssue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimilarIssue) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


