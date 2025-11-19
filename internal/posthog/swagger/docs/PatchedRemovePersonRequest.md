# PatchedRemovePersonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonId** | Pointer to **string** | Person UUID to remove from the cohort | [optional] 

## Methods

### NewPatchedRemovePersonRequest

`func NewPatchedRemovePersonRequest() *PatchedRemovePersonRequest`

NewPatchedRemovePersonRequest instantiates a new PatchedRemovePersonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRemovePersonRequestWithDefaults

`func NewPatchedRemovePersonRequestWithDefaults() *PatchedRemovePersonRequest`

NewPatchedRemovePersonRequestWithDefaults instantiates a new PatchedRemovePersonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonId

`func (o *PatchedRemovePersonRequest) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *PatchedRemovePersonRequest) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *PatchedRemovePersonRequest) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *PatchedRemovePersonRequest) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


