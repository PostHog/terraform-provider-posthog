# HogQLNotice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **NullableInt32** |  | [optional] 
**Fix** | Pointer to **NullableString** |  | [optional] 
**Message** | **string** |  | 
**Start** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewHogQLNotice

`func NewHogQLNotice(message string, ) *HogQLNotice`

NewHogQLNotice instantiates a new HogQLNotice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogQLNoticeWithDefaults

`func NewHogQLNoticeWithDefaults() *HogQLNotice`

NewHogQLNoticeWithDefaults instantiates a new HogQLNotice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *HogQLNotice) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *HogQLNotice) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *HogQLNotice) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *HogQLNotice) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *HogQLNotice) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *HogQLNotice) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetFix

`func (o *HogQLNotice) GetFix() string`

GetFix returns the Fix field if non-nil, zero value otherwise.

### GetFixOk

`func (o *HogQLNotice) GetFixOk() (*string, bool)`

GetFixOk returns a tuple with the Fix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFix

`func (o *HogQLNotice) SetFix(v string)`

SetFix sets Fix field to given value.

### HasFix

`func (o *HogQLNotice) HasFix() bool`

HasFix returns a boolean if a field has been set.

### SetFixNil

`func (o *HogQLNotice) SetFixNil(b bool)`

 SetFixNil sets the value for Fix to be an explicit nil

### UnsetFix
`func (o *HogQLNotice) UnsetFix()`

UnsetFix ensures that no value is present for Fix, not even an explicit nil
### GetMessage

`func (o *HogQLNotice) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HogQLNotice) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HogQLNotice) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStart

`func (o *HogQLNotice) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *HogQLNotice) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *HogQLNotice) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *HogQLNotice) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *HogQLNotice) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *HogQLNotice) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


