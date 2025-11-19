# QueryLogTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductKey** | Pointer to **NullableString** | Product responsible for this query. Use string, there&#39;s no need to churn the Schema when we add a new product * | [optional] 
**Scene** | Pointer to **NullableString** | Scene where this query is shown in the UI. Use string, there&#39;s no need to churn the Schema when we add a new Scene * | [optional] 

## Methods

### NewQueryLogTags

`func NewQueryLogTags() *QueryLogTags`

NewQueryLogTags instantiates a new QueryLogTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLogTagsWithDefaults

`func NewQueryLogTagsWithDefaults() *QueryLogTags`

NewQueryLogTagsWithDefaults instantiates a new QueryLogTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductKey

`func (o *QueryLogTags) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *QueryLogTags) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *QueryLogTags) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *QueryLogTags) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### SetProductKeyNil

`func (o *QueryLogTags) SetProductKeyNil(b bool)`

 SetProductKeyNil sets the value for ProductKey to be an explicit nil

### UnsetProductKey
`func (o *QueryLogTags) UnsetProductKey()`

UnsetProductKey ensures that no value is present for ProductKey, not even an explicit nil
### GetScene

`func (o *QueryLogTags) GetScene() string`

GetScene returns the Scene field if non-nil, zero value otherwise.

### GetSceneOk

`func (o *QueryLogTags) GetSceneOk() (*string, bool)`

GetSceneOk returns a tuple with the Scene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScene

`func (o *QueryLogTags) SetScene(v string)`

SetScene sets Scene field to given value.

### HasScene

`func (o *QueryLogTags) HasScene() bool`

HasScene returns a boolean if a field has been set.

### SetSceneNil

`func (o *QueryLogTags) SetSceneNil(b bool)`

 SetSceneNil sets the value for Scene to be an explicit nil

### UnsetScene
`func (o *QueryLogTags) UnsetScene()`

UnsetScene ensures that no value is present for Scene, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


