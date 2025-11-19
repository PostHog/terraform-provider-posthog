# WebExperimentsAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**FeatureFlagKey** | **string** |  | [readonly] 
**Variants** | **interface{}** | Variants for the web experiment. Example:          {             \&quot;control\&quot;: {                 \&quot;transforms\&quot;: [                     {                         \&quot;text\&quot;: \&quot;Here comes Superman!\&quot;,                         \&quot;html\&quot;: \&quot;\&quot;,                         \&quot;selector\&quot;: \&quot;#page &gt; #body &gt; .header h1\&quot;                     }                 ],                 \&quot;conditions\&quot;: \&quot;None\&quot;,                 \&quot;rollout_percentage\&quot;: 50             },         } | 

## Methods

### NewWebExperimentsAPI

`func NewWebExperimentsAPI(id int32, name string, featureFlagKey string, variants interface{}, ) *WebExperimentsAPI`

NewWebExperimentsAPI instantiates a new WebExperimentsAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebExperimentsAPIWithDefaults

`func NewWebExperimentsAPIWithDefaults() *WebExperimentsAPI`

NewWebExperimentsAPIWithDefaults instantiates a new WebExperimentsAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebExperimentsAPI) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebExperimentsAPI) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebExperimentsAPI) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *WebExperimentsAPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebExperimentsAPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebExperimentsAPI) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *WebExperimentsAPI) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebExperimentsAPI) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebExperimentsAPI) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebExperimentsAPI) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFeatureFlagKey

`func (o *WebExperimentsAPI) GetFeatureFlagKey() string`

GetFeatureFlagKey returns the FeatureFlagKey field if non-nil, zero value otherwise.

### GetFeatureFlagKeyOk

`func (o *WebExperimentsAPI) GetFeatureFlagKeyOk() (*string, bool)`

GetFeatureFlagKeyOk returns a tuple with the FeatureFlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagKey

`func (o *WebExperimentsAPI) SetFeatureFlagKey(v string)`

SetFeatureFlagKey sets FeatureFlagKey field to given value.


### GetVariants

`func (o *WebExperimentsAPI) GetVariants() interface{}`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *WebExperimentsAPI) GetVariantsOk() (*interface{}, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *WebExperimentsAPI) SetVariants(v interface{})`

SetVariants sets Variants field to given value.


### SetVariantsNil

`func (o *WebExperimentsAPI) SetVariantsNil(b bool)`

 SetVariantsNil sets the value for Variants to be an explicit nil

### UnsetVariants
`func (o *WebExperimentsAPI) UnsetVariants()`

UnsetVariants ensures that no value is present for Variants, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


