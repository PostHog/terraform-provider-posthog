# ExternalQueryError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**ExternalQueryErrorCode**](ExternalQueryErrorCode.md) |  | 
**Detail** | **string** |  | 

## Methods

### NewExternalQueryError

`func NewExternalQueryError(code ExternalQueryErrorCode, detail string, ) *ExternalQueryError`

NewExternalQueryError instantiates a new ExternalQueryError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalQueryErrorWithDefaults

`func NewExternalQueryErrorWithDefaults() *ExternalQueryError`

NewExternalQueryErrorWithDefaults instantiates a new ExternalQueryError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ExternalQueryError) GetCode() ExternalQueryErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExternalQueryError) GetCodeOk() (*ExternalQueryErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExternalQueryError) SetCode(v ExternalQueryErrorCode)`

SetCode sets Code field to given value.


### GetDetail

`func (o *ExternalQueryError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ExternalQueryError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ExternalQueryError) SetDetail(v string)`

SetDetail sets Detail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


