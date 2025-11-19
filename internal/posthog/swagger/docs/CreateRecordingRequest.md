# CreateRecordingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to [**CreateRecordingRequestPlatformEnum**](CreateRecordingRequestPlatformEnum.md) | Meeting platform being recorded  * &#x60;zoom&#x60; - zoom * &#x60;teams&#x60; - teams * &#x60;meet&#x60; - meet * &#x60;desktop_audio&#x60; - desktop_audio * &#x60;slack&#x60; - slack | [optional] [default to DESKTOP_AUDIO]

## Methods

### NewCreateRecordingRequest

`func NewCreateRecordingRequest() *CreateRecordingRequest`

NewCreateRecordingRequest instantiates a new CreateRecordingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecordingRequestWithDefaults

`func NewCreateRecordingRequestWithDefaults() *CreateRecordingRequest`

NewCreateRecordingRequestWithDefaults instantiates a new CreateRecordingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *CreateRecordingRequest) GetPlatform() CreateRecordingRequestPlatformEnum`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateRecordingRequest) GetPlatformOk() (*CreateRecordingRequestPlatformEnum, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateRecordingRequest) SetPlatform(v CreateRecordingRequestPlatformEnum)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CreateRecordingRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


