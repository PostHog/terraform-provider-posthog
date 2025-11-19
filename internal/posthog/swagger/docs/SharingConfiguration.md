# SharingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AccessToken** | **NullableString** |  | [readonly] 
**Settings** | Pointer to **interface{}** |  | [optional] 
**PasswordRequired** | Pointer to **bool** |  | [optional] 
**SharePasswords** | **string** |  | [readonly] 

## Methods

### NewSharingConfiguration

`func NewSharingConfiguration(createdAt time.Time, accessToken NullableString, sharePasswords string, ) *SharingConfiguration`

NewSharingConfiguration instantiates a new SharingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingConfigurationWithDefaults

`func NewSharingConfigurationWithDefaults() *SharingConfiguration`

NewSharingConfigurationWithDefaults instantiates a new SharingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SharingConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SharingConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SharingConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnabled

`func (o *SharingConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SharingConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SharingConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SharingConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAccessToken

`func (o *SharingConfiguration) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *SharingConfiguration) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *SharingConfiguration) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### SetAccessTokenNil

`func (o *SharingConfiguration) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *SharingConfiguration) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetSettings

`func (o *SharingConfiguration) GetSettings() interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SharingConfiguration) GetSettingsOk() (*interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SharingConfiguration) SetSettings(v interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SharingConfiguration) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *SharingConfiguration) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *SharingConfiguration) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetPasswordRequired

`func (o *SharingConfiguration) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *SharingConfiguration) GetPasswordRequiredOk() (*bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRequired

`func (o *SharingConfiguration) SetPasswordRequired(v bool)`

SetPasswordRequired sets PasswordRequired field to given value.

### HasPasswordRequired

`func (o *SharingConfiguration) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### GetSharePasswords

`func (o *SharingConfiguration) GetSharePasswords() string`

GetSharePasswords returns the SharePasswords field if non-nil, zero value otherwise.

### GetSharePasswordsOk

`func (o *SharingConfiguration) GetSharePasswordsOk() (*string, bool)`

GetSharePasswordsOk returns a tuple with the SharePasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePasswords

`func (o *SharingConfiguration) SetSharePasswords(v string)`

SetSharePasswords sets SharePasswords field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


