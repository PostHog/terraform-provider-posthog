# ErrorTrackingExternalReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUrl** | **string** |  | 
**Id** | **string** |  | 
**Integration** | [**ErrorTrackingExternalReferenceIntegration**](ErrorTrackingExternalReferenceIntegration.md) |  | 

## Methods

### NewErrorTrackingExternalReference

`func NewErrorTrackingExternalReference(externalUrl string, id string, integration ErrorTrackingExternalReferenceIntegration, ) *ErrorTrackingExternalReference`

NewErrorTrackingExternalReference instantiates a new ErrorTrackingExternalReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingExternalReferenceWithDefaults

`func NewErrorTrackingExternalReferenceWithDefaults() *ErrorTrackingExternalReference`

NewErrorTrackingExternalReferenceWithDefaults instantiates a new ErrorTrackingExternalReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUrl

`func (o *ErrorTrackingExternalReference) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *ErrorTrackingExternalReference) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *ErrorTrackingExternalReference) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.


### GetId

`func (o *ErrorTrackingExternalReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorTrackingExternalReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorTrackingExternalReference) SetId(v string)`

SetId sets Id field to given value.


### GetIntegration

`func (o *ErrorTrackingExternalReference) GetIntegration() ErrorTrackingExternalReferenceIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ErrorTrackingExternalReference) GetIntegrationOk() (*ErrorTrackingExternalReferenceIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ErrorTrackingExternalReference) SetIntegration(v ErrorTrackingExternalReferenceIntegration)`

SetIntegration sets Integration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


