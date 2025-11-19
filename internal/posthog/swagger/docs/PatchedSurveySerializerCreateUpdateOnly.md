# PatchedSurveySerializerCreateUpdateOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**SurveyType**](SurveyType.md) |  | [optional] 
**Schedule** | Pointer to **NullableString** |  | [optional] 
**LinkedFlag** | Pointer to [**MinimalFeatureFlag**](MinimalFeatureFlag.md) |  | [optional] [readonly] 
**LinkedFlagId** | Pointer to **NullableInt32** |  | [optional] 
**TargetingFlagId** | Pointer to **int32** |  | [optional] 
**TargetingFlag** | Pointer to [**MinimalFeatureFlag**](MinimalFeatureFlag.md) |  | [optional] [readonly] 
**InternalTargetingFlag** | Pointer to [**MinimalFeatureFlag**](MinimalFeatureFlag.md) |  | [optional] [readonly] 
**TargetingFlagFilters** | Pointer to **interface{}** |  | [optional] 
**RemoveTargetingFlag** | Pointer to **NullableBool** |  | [optional] 
**Questions** | Pointer to **interface{}** |          The &#x60;array&#x60; of questions included in the survey. Each question must conform to one of the defined question types: Basic, Link, Rating, or Multiple Choice.          Basic (open-ended question)         - &#x60;id&#x60;: The question ID         - &#x60;type&#x60;: &#x60;open&#x60;         - &#x60;question&#x60;: The text of the question.         - &#x60;description&#x60;: Optional description of the question.         - &#x60;descriptionContentType&#x60;: Content type of the description (&#x60;html&#x60; or &#x60;text&#x60;).         - &#x60;optional&#x60;: Whether the question is optional (&#x60;boolean&#x60;).         - &#x60;buttonText&#x60;: Text displayed on the submit button.         - &#x60;branching&#x60;: Branching logic for the question. See branching types below for details.          Link (a question with a link)         - &#x60;id&#x60;: The question ID         - &#x60;type&#x60;: &#x60;link&#x60;         - &#x60;question&#x60;: The text of the question.         - &#x60;description&#x60;: Optional description of the question.         - &#x60;descriptionContentType&#x60;: Content type of the description (&#x60;html&#x60; or &#x60;text&#x60;).         - &#x60;optional&#x60;: Whether the question is optional (&#x60;boolean&#x60;).         - &#x60;buttonText&#x60;: Text displayed on the submit button.         - &#x60;link&#x60;: The URL associated with the question.         - &#x60;branching&#x60;: Branching logic for the question. See branching types below for details.          Rating (a question with a rating scale)         - &#x60;id&#x60;: The question ID         - &#x60;type&#x60;: &#x60;rating&#x60;         - &#x60;question&#x60;: The text of the question.         - &#x60;description&#x60;: Optional description of the question.         - &#x60;descriptionContentType&#x60;: Content type of the description (&#x60;html&#x60; or &#x60;text&#x60;).         - &#x60;optional&#x60;: Whether the question is optional (&#x60;boolean&#x60;).         - &#x60;buttonText&#x60;: Text displayed on the submit button.         - &#x60;display&#x60;: Display style of the rating (&#x60;number&#x60; or &#x60;emoji&#x60;).         - &#x60;scale&#x60;: The scale of the rating (&#x60;number&#x60;).         - &#x60;lowerBoundLabel&#x60;: Label for the lower bound of the scale.         - &#x60;upperBoundLabel&#x60;: Label for the upper bound of the scale.         - &#x60;branching&#x60;: Branching logic for the question. See branching types below for details.          Multiple choice         - &#x60;id&#x60;: The question ID         - &#x60;type&#x60;: &#x60;single_choice&#x60; or &#x60;multiple_choice&#x60;         - &#x60;question&#x60;: The text of the question.         - &#x60;description&#x60;: Optional description of the question.         - &#x60;descriptionContentType&#x60;: Content type of the description (&#x60;html&#x60; or &#x60;text&#x60;).         - &#x60;optional&#x60;: Whether the question is optional (&#x60;boolean&#x60;).         - &#x60;buttonText&#x60;: Text displayed on the submit button.         - &#x60;choices&#x60;: An array of choices for the question.         - &#x60;shuffleOptions&#x60;: Whether to shuffle the order of the choices (&#x60;boolean&#x60;).         - &#x60;hasOpenChoice&#x60;: Whether the question allows an open-ended response (&#x60;boolean&#x60;).         - &#x60;branching&#x60;: Branching logic for the question. See branching types below for details.          Branching logic can be one of the following types:          Next question: Proceeds to the next question         &#x60;&#x60;&#x60;json         {             \&quot;type\&quot;: \&quot;next_question\&quot;         }         &#x60;&#x60;&#x60;          End: Ends the survey, optionally displaying a confirmation message.         &#x60;&#x60;&#x60;json         {             \&quot;type\&quot;: \&quot;end\&quot;         }         &#x60;&#x60;&#x60;          Response-based: Branches based on the response values. Available for the &#x60;rating&#x60; and &#x60;single_choice&#x60; question types.         &#x60;&#x60;&#x60;json         {             \&quot;type\&quot;: \&quot;response_based\&quot;,             \&quot;responseValues\&quot;: {                 \&quot;responseKey\&quot;: \&quot;value\&quot;             }         }         &#x60;&#x60;&#x60;          Specific question: Proceeds to a specific question by index.         &#x60;&#x60;&#x60;json         {             \&quot;type\&quot;: \&quot;specific_question\&quot;,             \&quot;index\&quot;: 2         }         &#x60;&#x60;&#x60;          | [optional] 
**Conditions** | Pointer to **interface{}** |  | [optional] 
**Appearance** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**ResponsesLimit** | Pointer to **NullableInt32** |  | [optional] 
**IterationCount** | Pointer to **NullableInt32** |  | [optional] 
**IterationFrequencyDays** | Pointer to **NullableInt32** |  | [optional] 
**IterationStartDates** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**CurrentIteration** | Pointer to **NullableInt32** |  | [optional] 
**CurrentIterationStartDate** | Pointer to **NullableTime** |  | [optional] 
**ResponseSamplingStartDate** | Pointer to **NullableTime** |  | [optional] 
**ResponseSamplingIntervalType** | Pointer to **NullableString** |  | [optional] 
**ResponseSamplingInterval** | Pointer to **NullableInt32** |  | [optional] 
**ResponseSamplingLimit** | Pointer to **NullableInt32** |  | [optional] 
**ResponseSamplingDailyLimits** | Pointer to **interface{}** |  | [optional] 
**EnablePartialResponses** | Pointer to **NullableBool** |  | [optional] 
**CreateInFolder** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedSurveySerializerCreateUpdateOnly

`func NewPatchedSurveySerializerCreateUpdateOnly() *PatchedSurveySerializerCreateUpdateOnly`

NewPatchedSurveySerializerCreateUpdateOnly instantiates a new PatchedSurveySerializerCreateUpdateOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSurveySerializerCreateUpdateOnlyWithDefaults

`func NewPatchedSurveySerializerCreateUpdateOnlyWithDefaults() *PatchedSurveySerializerCreateUpdateOnly`

NewPatchedSurveySerializerCreateUpdateOnlyWithDefaults instantiates a new PatchedSurveySerializerCreateUpdateOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetType() SurveyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetTypeOk() (*SurveyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetType(v SurveyType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSchedule

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetLinkedFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetLinkedFlag() MinimalFeatureFlag`

GetLinkedFlag returns the LinkedFlag field if non-nil, zero value otherwise.

### GetLinkedFlagOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetLinkedFlagOk() (*MinimalFeatureFlag, bool)`

GetLinkedFlagOk returns a tuple with the LinkedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetLinkedFlag(v MinimalFeatureFlag)`

SetLinkedFlag sets LinkedFlag field to given value.

### HasLinkedFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasLinkedFlag() bool`

HasLinkedFlag returns a boolean if a field has been set.

### GetLinkedFlagId

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetLinkedFlagId() int32`

GetLinkedFlagId returns the LinkedFlagId field if non-nil, zero value otherwise.

### GetLinkedFlagIdOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetLinkedFlagIdOk() (*int32, bool)`

GetLinkedFlagIdOk returns a tuple with the LinkedFlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFlagId

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetLinkedFlagId(v int32)`

SetLinkedFlagId sets LinkedFlagId field to given value.

### HasLinkedFlagId

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasLinkedFlagId() bool`

HasLinkedFlagId returns a boolean if a field has been set.

### SetLinkedFlagIdNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetLinkedFlagIdNil(b bool)`

 SetLinkedFlagIdNil sets the value for LinkedFlagId to be an explicit nil

### UnsetLinkedFlagId
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetLinkedFlagId()`

UnsetLinkedFlagId ensures that no value is present for LinkedFlagId, not even an explicit nil
### GetTargetingFlagId

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetTargetingFlagId() int32`

GetTargetingFlagId returns the TargetingFlagId field if non-nil, zero value otherwise.

### GetTargetingFlagIdOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetTargetingFlagIdOk() (*int32, bool)`

GetTargetingFlagIdOk returns a tuple with the TargetingFlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingFlagId

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetTargetingFlagId(v int32)`

SetTargetingFlagId sets TargetingFlagId field to given value.

### HasTargetingFlagId

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasTargetingFlagId() bool`

HasTargetingFlagId returns a boolean if a field has been set.

### GetTargetingFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetTargetingFlag() MinimalFeatureFlag`

GetTargetingFlag returns the TargetingFlag field if non-nil, zero value otherwise.

### GetTargetingFlagOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetTargetingFlagOk() (*MinimalFeatureFlag, bool)`

GetTargetingFlagOk returns a tuple with the TargetingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetTargetingFlag(v MinimalFeatureFlag)`

SetTargetingFlag sets TargetingFlag field to given value.

### HasTargetingFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasTargetingFlag() bool`

HasTargetingFlag returns a boolean if a field has been set.

### GetInternalTargetingFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetInternalTargetingFlag() MinimalFeatureFlag`

GetInternalTargetingFlag returns the InternalTargetingFlag field if non-nil, zero value otherwise.

### GetInternalTargetingFlagOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetInternalTargetingFlagOk() (*MinimalFeatureFlag, bool)`

GetInternalTargetingFlagOk returns a tuple with the InternalTargetingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTargetingFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetInternalTargetingFlag(v MinimalFeatureFlag)`

SetInternalTargetingFlag sets InternalTargetingFlag field to given value.

### HasInternalTargetingFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasInternalTargetingFlag() bool`

HasInternalTargetingFlag returns a boolean if a field has been set.

### GetTargetingFlagFilters

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetTargetingFlagFilters() interface{}`

GetTargetingFlagFilters returns the TargetingFlagFilters field if non-nil, zero value otherwise.

### GetTargetingFlagFiltersOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetTargetingFlagFiltersOk() (*interface{}, bool)`

GetTargetingFlagFiltersOk returns a tuple with the TargetingFlagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingFlagFilters

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetTargetingFlagFilters(v interface{})`

SetTargetingFlagFilters sets TargetingFlagFilters field to given value.

### HasTargetingFlagFilters

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasTargetingFlagFilters() bool`

HasTargetingFlagFilters returns a boolean if a field has been set.

### SetTargetingFlagFiltersNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetTargetingFlagFiltersNil(b bool)`

 SetTargetingFlagFiltersNil sets the value for TargetingFlagFilters to be an explicit nil

### UnsetTargetingFlagFilters
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetTargetingFlagFilters()`

UnsetTargetingFlagFilters ensures that no value is present for TargetingFlagFilters, not even an explicit nil
### GetRemoveTargetingFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetRemoveTargetingFlag() bool`

GetRemoveTargetingFlag returns the RemoveTargetingFlag field if non-nil, zero value otherwise.

### GetRemoveTargetingFlagOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetRemoveTargetingFlagOk() (*bool, bool)`

GetRemoveTargetingFlagOk returns a tuple with the RemoveTargetingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTargetingFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetRemoveTargetingFlag(v bool)`

SetRemoveTargetingFlag sets RemoveTargetingFlag field to given value.

### HasRemoveTargetingFlag

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasRemoveTargetingFlag() bool`

HasRemoveTargetingFlag returns a boolean if a field has been set.

### SetRemoveTargetingFlagNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetRemoveTargetingFlagNil(b bool)`

 SetRemoveTargetingFlagNil sets the value for RemoveTargetingFlag to be an explicit nil

### UnsetRemoveTargetingFlag
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetRemoveTargetingFlag()`

UnsetRemoveTargetingFlag ensures that no value is present for RemoveTargetingFlag, not even an explicit nil
### GetQuestions

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetQuestions() interface{}`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetQuestionsOk() (*interface{}, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetQuestions(v interface{})`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### SetQuestionsNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetQuestionsNil(b bool)`

 SetQuestionsNil sets the value for Questions to be an explicit nil

### UnsetQuestions
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetQuestions()`

UnsetQuestions ensures that no value is present for Questions, not even an explicit nil
### GetConditions

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetConditions() interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetConditionsOk() (*interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetConditions(v interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetAppearance

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetAppearance() interface{}`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetAppearanceOk() (*interface{}, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetAppearance(v interface{})`

SetAppearance sets Appearance field to given value.

### HasAppearance

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasAppearance() bool`

HasAppearance returns a boolean if a field has been set.

### SetAppearanceNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetAppearanceNil(b bool)`

 SetAppearanceNil sets the value for Appearance to be an explicit nil

### UnsetAppearance
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetAppearance()`

UnsetAppearance ensures that no value is present for Appearance, not even an explicit nil
### GetCreatedAt

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetStartDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetArchived

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetResponsesLimit

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponsesLimit() int32`

GetResponsesLimit returns the ResponsesLimit field if non-nil, zero value otherwise.

### GetResponsesLimitOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponsesLimitOk() (*int32, bool)`

GetResponsesLimitOk returns a tuple with the ResponsesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsesLimit

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponsesLimit(v int32)`

SetResponsesLimit sets ResponsesLimit field to given value.

### HasResponsesLimit

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasResponsesLimit() bool`

HasResponsesLimit returns a boolean if a field has been set.

### SetResponsesLimitNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponsesLimitNil(b bool)`

 SetResponsesLimitNil sets the value for ResponsesLimit to be an explicit nil

### UnsetResponsesLimit
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetResponsesLimit()`

UnsetResponsesLimit ensures that no value is present for ResponsesLimit, not even an explicit nil
### GetIterationCount

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetIterationCount() int32`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetIterationCountOk() (*int32, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetIterationCount(v int32)`

SetIterationCount sets IterationCount field to given value.

### HasIterationCount

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasIterationCount() bool`

HasIterationCount returns a boolean if a field has been set.

### SetIterationCountNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetIterationCountNil(b bool)`

 SetIterationCountNil sets the value for IterationCount to be an explicit nil

### UnsetIterationCount
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetIterationCount()`

UnsetIterationCount ensures that no value is present for IterationCount, not even an explicit nil
### GetIterationFrequencyDays

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetIterationFrequencyDays() int32`

GetIterationFrequencyDays returns the IterationFrequencyDays field if non-nil, zero value otherwise.

### GetIterationFrequencyDaysOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetIterationFrequencyDaysOk() (*int32, bool)`

GetIterationFrequencyDaysOk returns a tuple with the IterationFrequencyDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationFrequencyDays

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetIterationFrequencyDays(v int32)`

SetIterationFrequencyDays sets IterationFrequencyDays field to given value.

### HasIterationFrequencyDays

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasIterationFrequencyDays() bool`

HasIterationFrequencyDays returns a boolean if a field has been set.

### SetIterationFrequencyDaysNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetIterationFrequencyDaysNil(b bool)`

 SetIterationFrequencyDaysNil sets the value for IterationFrequencyDays to be an explicit nil

### UnsetIterationFrequencyDays
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetIterationFrequencyDays()`

UnsetIterationFrequencyDays ensures that no value is present for IterationFrequencyDays, not even an explicit nil
### GetIterationStartDates

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetIterationStartDates() []*time.Time`

GetIterationStartDates returns the IterationStartDates field if non-nil, zero value otherwise.

### GetIterationStartDatesOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetIterationStartDatesOk() (*[]*time.Time, bool)`

GetIterationStartDatesOk returns a tuple with the IterationStartDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationStartDates

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetIterationStartDates(v []*time.Time)`

SetIterationStartDates sets IterationStartDates field to given value.

### HasIterationStartDates

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasIterationStartDates() bool`

HasIterationStartDates returns a boolean if a field has been set.

### SetIterationStartDatesNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetIterationStartDatesNil(b bool)`

 SetIterationStartDatesNil sets the value for IterationStartDates to be an explicit nil

### UnsetIterationStartDates
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetIterationStartDates()`

UnsetIterationStartDates ensures that no value is present for IterationStartDates, not even an explicit nil
### GetCurrentIteration

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetCurrentIteration() int32`

GetCurrentIteration returns the CurrentIteration field if non-nil, zero value otherwise.

### GetCurrentIterationOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetCurrentIterationOk() (*int32, bool)`

GetCurrentIterationOk returns a tuple with the CurrentIteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIteration

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetCurrentIteration(v int32)`

SetCurrentIteration sets CurrentIteration field to given value.

### HasCurrentIteration

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasCurrentIteration() bool`

HasCurrentIteration returns a boolean if a field has been set.

### SetCurrentIterationNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetCurrentIterationNil(b bool)`

 SetCurrentIterationNil sets the value for CurrentIteration to be an explicit nil

### UnsetCurrentIteration
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetCurrentIteration()`

UnsetCurrentIteration ensures that no value is present for CurrentIteration, not even an explicit nil
### GetCurrentIterationStartDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetCurrentIterationStartDate() time.Time`

GetCurrentIterationStartDate returns the CurrentIterationStartDate field if non-nil, zero value otherwise.

### GetCurrentIterationStartDateOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetCurrentIterationStartDateOk() (*time.Time, bool)`

GetCurrentIterationStartDateOk returns a tuple with the CurrentIterationStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIterationStartDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetCurrentIterationStartDate(v time.Time)`

SetCurrentIterationStartDate sets CurrentIterationStartDate field to given value.

### HasCurrentIterationStartDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasCurrentIterationStartDate() bool`

HasCurrentIterationStartDate returns a boolean if a field has been set.

### SetCurrentIterationStartDateNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetCurrentIterationStartDateNil(b bool)`

 SetCurrentIterationStartDateNil sets the value for CurrentIterationStartDate to be an explicit nil

### UnsetCurrentIterationStartDate
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetCurrentIterationStartDate()`

UnsetCurrentIterationStartDate ensures that no value is present for CurrentIterationStartDate, not even an explicit nil
### GetResponseSamplingStartDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponseSamplingStartDate() time.Time`

GetResponseSamplingStartDate returns the ResponseSamplingStartDate field if non-nil, zero value otherwise.

### GetResponseSamplingStartDateOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponseSamplingStartDateOk() (*time.Time, bool)`

GetResponseSamplingStartDateOk returns a tuple with the ResponseSamplingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSamplingStartDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponseSamplingStartDate(v time.Time)`

SetResponseSamplingStartDate sets ResponseSamplingStartDate field to given value.

### HasResponseSamplingStartDate

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasResponseSamplingStartDate() bool`

HasResponseSamplingStartDate returns a boolean if a field has been set.

### SetResponseSamplingStartDateNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponseSamplingStartDateNil(b bool)`

 SetResponseSamplingStartDateNil sets the value for ResponseSamplingStartDate to be an explicit nil

### UnsetResponseSamplingStartDate
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetResponseSamplingStartDate()`

UnsetResponseSamplingStartDate ensures that no value is present for ResponseSamplingStartDate, not even an explicit nil
### GetResponseSamplingIntervalType

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponseSamplingIntervalType() string`

GetResponseSamplingIntervalType returns the ResponseSamplingIntervalType field if non-nil, zero value otherwise.

### GetResponseSamplingIntervalTypeOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponseSamplingIntervalTypeOk() (*string, bool)`

GetResponseSamplingIntervalTypeOk returns a tuple with the ResponseSamplingIntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSamplingIntervalType

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponseSamplingIntervalType(v string)`

SetResponseSamplingIntervalType sets ResponseSamplingIntervalType field to given value.

### HasResponseSamplingIntervalType

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasResponseSamplingIntervalType() bool`

HasResponseSamplingIntervalType returns a boolean if a field has been set.

### SetResponseSamplingIntervalTypeNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponseSamplingIntervalTypeNil(b bool)`

 SetResponseSamplingIntervalTypeNil sets the value for ResponseSamplingIntervalType to be an explicit nil

### UnsetResponseSamplingIntervalType
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetResponseSamplingIntervalType()`

UnsetResponseSamplingIntervalType ensures that no value is present for ResponseSamplingIntervalType, not even an explicit nil
### GetResponseSamplingInterval

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponseSamplingInterval() int32`

GetResponseSamplingInterval returns the ResponseSamplingInterval field if non-nil, zero value otherwise.

### GetResponseSamplingIntervalOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponseSamplingIntervalOk() (*int32, bool)`

GetResponseSamplingIntervalOk returns a tuple with the ResponseSamplingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSamplingInterval

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponseSamplingInterval(v int32)`

SetResponseSamplingInterval sets ResponseSamplingInterval field to given value.

### HasResponseSamplingInterval

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasResponseSamplingInterval() bool`

HasResponseSamplingInterval returns a boolean if a field has been set.

### SetResponseSamplingIntervalNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponseSamplingIntervalNil(b bool)`

 SetResponseSamplingIntervalNil sets the value for ResponseSamplingInterval to be an explicit nil

### UnsetResponseSamplingInterval
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetResponseSamplingInterval()`

UnsetResponseSamplingInterval ensures that no value is present for ResponseSamplingInterval, not even an explicit nil
### GetResponseSamplingLimit

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponseSamplingLimit() int32`

GetResponseSamplingLimit returns the ResponseSamplingLimit field if non-nil, zero value otherwise.

### GetResponseSamplingLimitOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponseSamplingLimitOk() (*int32, bool)`

GetResponseSamplingLimitOk returns a tuple with the ResponseSamplingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSamplingLimit

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponseSamplingLimit(v int32)`

SetResponseSamplingLimit sets ResponseSamplingLimit field to given value.

### HasResponseSamplingLimit

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasResponseSamplingLimit() bool`

HasResponseSamplingLimit returns a boolean if a field has been set.

### SetResponseSamplingLimitNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponseSamplingLimitNil(b bool)`

 SetResponseSamplingLimitNil sets the value for ResponseSamplingLimit to be an explicit nil

### UnsetResponseSamplingLimit
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetResponseSamplingLimit()`

UnsetResponseSamplingLimit ensures that no value is present for ResponseSamplingLimit, not even an explicit nil
### GetResponseSamplingDailyLimits

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponseSamplingDailyLimits() interface{}`

GetResponseSamplingDailyLimits returns the ResponseSamplingDailyLimits field if non-nil, zero value otherwise.

### GetResponseSamplingDailyLimitsOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetResponseSamplingDailyLimitsOk() (*interface{}, bool)`

GetResponseSamplingDailyLimitsOk returns a tuple with the ResponseSamplingDailyLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSamplingDailyLimits

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponseSamplingDailyLimits(v interface{})`

SetResponseSamplingDailyLimits sets ResponseSamplingDailyLimits field to given value.

### HasResponseSamplingDailyLimits

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasResponseSamplingDailyLimits() bool`

HasResponseSamplingDailyLimits returns a boolean if a field has been set.

### SetResponseSamplingDailyLimitsNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetResponseSamplingDailyLimitsNil(b bool)`

 SetResponseSamplingDailyLimitsNil sets the value for ResponseSamplingDailyLimits to be an explicit nil

### UnsetResponseSamplingDailyLimits
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetResponseSamplingDailyLimits()`

UnsetResponseSamplingDailyLimits ensures that no value is present for ResponseSamplingDailyLimits, not even an explicit nil
### GetEnablePartialResponses

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetEnablePartialResponses() bool`

GetEnablePartialResponses returns the EnablePartialResponses field if non-nil, zero value otherwise.

### GetEnablePartialResponsesOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetEnablePartialResponsesOk() (*bool, bool)`

GetEnablePartialResponsesOk returns a tuple with the EnablePartialResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartialResponses

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetEnablePartialResponses(v bool)`

SetEnablePartialResponses sets EnablePartialResponses field to given value.

### HasEnablePartialResponses

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasEnablePartialResponses() bool`

HasEnablePartialResponses returns a boolean if a field has been set.

### SetEnablePartialResponsesNil

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetEnablePartialResponsesNil(b bool)`

 SetEnablePartialResponsesNil sets the value for EnablePartialResponses to be an explicit nil

### UnsetEnablePartialResponses
`func (o *PatchedSurveySerializerCreateUpdateOnly) UnsetEnablePartialResponses()`

UnsetEnablePartialResponses ensures that no value is present for EnablePartialResponses, not even an explicit nil
### GetCreateInFolder

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *PatchedSurveySerializerCreateUpdateOnly) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *PatchedSurveySerializerCreateUpdateOnly) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *PatchedSurveySerializerCreateUpdateOnly) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


