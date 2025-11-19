# CreateGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupTypeIndex** | **int32** |  | 
**GroupKey** | **string** |  | 
**GroupProperties** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCreateGroup

`func NewCreateGroup(groupTypeIndex int32, groupKey string, ) *CreateGroup`

NewCreateGroup instantiates a new CreateGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupWithDefaults

`func NewCreateGroupWithDefaults() *CreateGroup`

NewCreateGroupWithDefaults instantiates a new CreateGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupTypeIndex

`func (o *CreateGroup) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *CreateGroup) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *CreateGroup) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.


### GetGroupKey

`func (o *CreateGroup) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *CreateGroup) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *CreateGroup) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.


### GetGroupProperties

`func (o *CreateGroup) GetGroupProperties() interface{}`

GetGroupProperties returns the GroupProperties field if non-nil, zero value otherwise.

### GetGroupPropertiesOk

`func (o *CreateGroup) GetGroupPropertiesOk() (*interface{}, bool)`

GetGroupPropertiesOk returns a tuple with the GroupProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupProperties

`func (o *CreateGroup) SetGroupProperties(v interface{})`

SetGroupProperties sets GroupProperties field to given value.

### HasGroupProperties

`func (o *CreateGroup) HasGroupProperties() bool`

HasGroupProperties returns a boolean if a field has been set.

### SetGroupPropertiesNil

`func (o *CreateGroup) SetGroupPropertiesNil(b bool)`

 SetGroupPropertiesNil sets the value for GroupProperties to be an explicit nil

### UnsetGroupProperties
`func (o *CreateGroup) UnsetGroupProperties()`

UnsetGroupProperties ensures that no value is present for GroupProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


