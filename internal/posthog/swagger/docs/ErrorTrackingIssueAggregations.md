# ErrorTrackingIssueAggregations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Occurrences** | **float32** |  | 
**Sessions** | **float32** |  | 
**Users** | **float32** |  | 
**VolumeRange** | Pointer to **[]float32** |  | [optional] 
**VolumeBuckets** | [**[]VolumeBucket**](VolumeBucket.md) |  | 

## Methods

### NewErrorTrackingIssueAggregations

`func NewErrorTrackingIssueAggregations(occurrences float32, sessions float32, users float32, volumeBuckets []VolumeBucket, ) *ErrorTrackingIssueAggregations`

NewErrorTrackingIssueAggregations instantiates a new ErrorTrackingIssueAggregations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingIssueAggregationsWithDefaults

`func NewErrorTrackingIssueAggregationsWithDefaults() *ErrorTrackingIssueAggregations`

NewErrorTrackingIssueAggregationsWithDefaults instantiates a new ErrorTrackingIssueAggregations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurrences

`func (o *ErrorTrackingIssueAggregations) GetOccurrences() float32`

GetOccurrences returns the Occurrences field if non-nil, zero value otherwise.

### GetOccurrencesOk

`func (o *ErrorTrackingIssueAggregations) GetOccurrencesOk() (*float32, bool)`

GetOccurrencesOk returns a tuple with the Occurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrences

`func (o *ErrorTrackingIssueAggregations) SetOccurrences(v float32)`

SetOccurrences sets Occurrences field to given value.


### GetSessions

`func (o *ErrorTrackingIssueAggregations) GetSessions() float32`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ErrorTrackingIssueAggregations) GetSessionsOk() (*float32, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ErrorTrackingIssueAggregations) SetSessions(v float32)`

SetSessions sets Sessions field to given value.


### GetUsers

`func (o *ErrorTrackingIssueAggregations) GetUsers() float32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ErrorTrackingIssueAggregations) GetUsersOk() (*float32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ErrorTrackingIssueAggregations) SetUsers(v float32)`

SetUsers sets Users field to given value.


### GetVolumeRange

`func (o *ErrorTrackingIssueAggregations) GetVolumeRange() []float32`

GetVolumeRange returns the VolumeRange field if non-nil, zero value otherwise.

### GetVolumeRangeOk

`func (o *ErrorTrackingIssueAggregations) GetVolumeRangeOk() (*[]float32, bool)`

GetVolumeRangeOk returns a tuple with the VolumeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRange

`func (o *ErrorTrackingIssueAggregations) SetVolumeRange(v []float32)`

SetVolumeRange sets VolumeRange field to given value.

### HasVolumeRange

`func (o *ErrorTrackingIssueAggregations) HasVolumeRange() bool`

HasVolumeRange returns a boolean if a field has been set.

### SetVolumeRangeNil

`func (o *ErrorTrackingIssueAggregations) SetVolumeRangeNil(b bool)`

 SetVolumeRangeNil sets the value for VolumeRange to be an explicit nil

### UnsetVolumeRange
`func (o *ErrorTrackingIssueAggregations) UnsetVolumeRange()`

UnsetVolumeRange ensures that no value is present for VolumeRange, not even an explicit nil
### GetVolumeBuckets

`func (o *ErrorTrackingIssueAggregations) GetVolumeBuckets() []VolumeBucket`

GetVolumeBuckets returns the VolumeBuckets field if non-nil, zero value otherwise.

### GetVolumeBucketsOk

`func (o *ErrorTrackingIssueAggregations) GetVolumeBucketsOk() (*[]VolumeBucket, bool)`

GetVolumeBucketsOk returns a tuple with the VolumeBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBuckets

`func (o *ErrorTrackingIssueAggregations) SetVolumeBuckets(v []VolumeBucket)`

SetVolumeBuckets sets VolumeBuckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


