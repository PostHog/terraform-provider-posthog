# PatchedWebExperimentsAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**FeatureFlagKey** | Pointer to **string** |  | [optional] [readonly] 
**Variants** | Pointer to **interface{}** | Variants for the web experiment. Example:          {             \&quot;control\&quot;: {                 \&quot;transforms\&quot;: [                     {                         \&quot;text\&quot;: \&quot;Here comes Superman!\&quot;,                         \&quot;html\&quot;: \&quot;\&quot;,                         \&quot;selector\&quot;: \&quot;#page &gt; #body &gt; .header h1\&quot;                     }                 ],                 \&quot;conditions\&quot;: \&quot;None\&quot;,                 \&quot;rollout_percentage\&quot;: 50             },         } | [optional] 

## Methods

### NewPatchedWebExperimentsAPI

`func NewPatchedWebExperimentsAPI() *PatchedWebExperimentsAPI`

NewPatchedWebExperimentsAPI instantiates a new PatchedWebExperimentsAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWebExperimentsAPIWithDefaults

`func NewPatchedWebExperimentsAPIWithDefaults() *PatchedWebExperimentsAPI`

NewPatchedWebExperimentsAPIWithDefaults instantiates a new PatchedWebExperimentsAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWebExperimentsAPI) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWebExperimentsAPI) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWebExperimentsAPI) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWebExperimentsAPI) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedWebExperimentsAPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWebExperimentsAPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWebExperimentsAPI) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWebExperimentsAPI) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedWebExperimentsAPI) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedWebExperimentsAPI) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedWebExperimentsAPI) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedWebExperimentsAPI) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFeatureFlagKey

`func (o *PatchedWebExperimentsAPI) GetFeatureFlagKey() string`

GetFeatureFlagKey returns the FeatureFlagKey field if non-nil, zero value otherwise.

### GetFeatureFlagKeyOk

`func (o *PatchedWebExperimentsAPI) GetFeatureFlagKeyOk() (*string, bool)`

GetFeatureFlagKeyOk returns a tuple with the FeatureFlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagKey

`func (o *PatchedWebExperimentsAPI) SetFeatureFlagKey(v string)`

SetFeatureFlagKey sets FeatureFlagKey field to given value.

### HasFeatureFlagKey

`func (o *PatchedWebExperimentsAPI) HasFeatureFlagKey() bool`

HasFeatureFlagKey returns a boolean if a field has been set.

### GetVariants

`func (o *PatchedWebExperimentsAPI) GetVariants() interface{}`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *PatchedWebExperimentsAPI) GetVariantsOk() (*interface{}, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *PatchedWebExperimentsAPI) SetVariants(v interface{})`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *PatchedWebExperimentsAPI) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### SetVariantsNil

`func (o *PatchedWebExperimentsAPI) SetVariantsNil(b bool)`

 SetVariantsNil sets the value for Variants to be an explicit nil

### UnsetVariants
`func (o *PatchedWebExperimentsAPI) UnsetVariants()`

UnsetVariants ensures that no value is present for Variants, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


