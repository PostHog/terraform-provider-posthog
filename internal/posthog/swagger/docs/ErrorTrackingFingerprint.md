# ErrorTrackingFingerprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | **string** |  | 
**IssueId** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewErrorTrackingFingerprint

`func NewErrorTrackingFingerprint(fingerprint string, issueId string, createdAt time.Time, ) *ErrorTrackingFingerprint`

NewErrorTrackingFingerprint instantiates a new ErrorTrackingFingerprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingFingerprintWithDefaults

`func NewErrorTrackingFingerprintWithDefaults() *ErrorTrackingFingerprint`

NewErrorTrackingFingerprintWithDefaults instantiates a new ErrorTrackingFingerprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *ErrorTrackingFingerprint) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ErrorTrackingFingerprint) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ErrorTrackingFingerprint) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetIssueId

`func (o *ErrorTrackingFingerprint) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *ErrorTrackingFingerprint) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *ErrorTrackingFingerprint) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.


### GetCreatedAt

`func (o *ErrorTrackingFingerprint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ErrorTrackingFingerprint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ErrorTrackingFingerprint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


