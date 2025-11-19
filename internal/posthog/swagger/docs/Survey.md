# Survey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | [**SurveyType**](SurveyType.md) |  | 
**Schedule** | Pointer to **NullableString** |  | [optional] 
**LinkedFlag** | [**MinimalFeatureFlag**](MinimalFeatureFlag.md) |  | [readonly] 
**LinkedFlagId** | Pointer to **NullableInt32** |  | [optional] 
**TargetingFlag** | [**MinimalFeatureFlag**](MinimalFeatureFlag.md) |  | [readonly] 
**InternalTargetingFlag** | [**MinimalFeatureFlag**](MinimalFeatureFlag.md) |  | [readonly] 
**Questions** | Pointer to **interface{}** |          The &#x60;array&#x60; of questions included in the survey. Each question must conform to one of the defined question types: Basic, Link, Rating, or Multiple Choice.          Basic (open-ended question)         - &#x60;id&#x60;: The question ID         - &#x60;type&#x60;: &#x60;open&#x60;         - &#x60;question&#x60;: The text of the question.         - &#x60;description&#x60;: Optional description of the question.         - &#x60;descriptionContentType&#x60;: Content type of the description (&#x60;html&#x60; or &#x60;text&#x60;).         - &#x60;optional&#x60;: Whether the question is optional (&#x60;boolean&#x60;).         - &#x60;buttonText&#x60;: Text displayed on the submit button.         - &#x60;branching&#x60;: Branching logic for the question. See branching types below for details.          Link (a question with a link)         - &#x60;id&#x60;: The question ID         - &#x60;type&#x60;: &#x60;link&#x60;         - &#x60;question&#x60;: The text of the question.         - &#x60;description&#x60;: Optional description of the question.         - &#x60;descriptionContentType&#x60;: Content type of the description (&#x60;html&#x60; or &#x60;text&#x60;).         - &#x60;optional&#x60;: Whether the question is optional (&#x60;boolean&#x60;).         - &#x60;buttonText&#x60;: Text displayed on the submit button.         - &#x60;link&#x60;: The URL associated with the question.         - &#x60;branching&#x60;: Branching logic for the question. See branching types below for details.          Rating (a question with a rating scale)         - &#x60;id&#x60;: The question ID         - &#x60;type&#x60;: &#x60;rating&#x60;         - &#x60;question&#x60;: The text of the question.         - &#x60;description&#x60;: Optional description of the question.         - &#x60;descriptionContentType&#x60;: Content type of the description (&#x60;html&#x60; or &#x60;text&#x60;).         - &#x60;optional&#x60;: Whether the question is optional (&#x60;boolean&#x60;).         - &#x60;buttonText&#x60;: Text displayed on the submit button.         - &#x60;display&#x60;: Display style of the rating (&#x60;number&#x60; or &#x60;emoji&#x60;).         - &#x60;scale&#x60;: The scale of the rating (&#x60;number&#x60;).         - &#x60;lowerBoundLabel&#x60;: Label for the lower bound of the scale.         - &#x60;upperBoundLabel&#x60;: Label for the upper bound of the scale.         - &#x60;branching&#x60;: Branching logic for the question. See branching types below for details.          Multiple choice         - &#x60;id&#x60;: The question ID         - &#x60;type&#x60;: &#x60;single_choice&#x60; or &#x60;multiple_choice&#x60;         - &#x60;question&#x60;: The text of the question.         - &#x60;description&#x60;: Optional description of the question.         - &#x60;descriptionContentType&#x60;: Content type of the description (&#x60;html&#x60; or &#x60;text&#x60;).         - &#x60;optional&#x60;: Whether the question is optional (&#x60;boolean&#x60;).         - &#x60;buttonText&#x60;: Text displayed on the submit button.         - &#x60;choices&#x60;: An array of choices for the question.         - &#x60;shuffleOptions&#x60;: Whether to shuffle the order of the choices (&#x60;boolean&#x60;).         - &#x60;hasOpenChoice&#x60;: Whether the question allows an open-ended response (&#x60;boolean&#x60;).         - &#x60;branching&#x60;: Branching logic for the question. See branching types below for details.          Branching logic can be one of the following types:          Next question: Proceeds to the next question         &#x60;&#x60;&#x60;json         {             \&quot;type\&quot;: \&quot;next_question\&quot;         }         &#x60;&#x60;&#x60;          End: Ends the survey, optionally displaying a confirmation message.         &#x60;&#x60;&#x60;json         {             \&quot;type\&quot;: \&quot;end\&quot;         }         &#x60;&#x60;&#x60;          Response-based: Branches based on the response values. Available for the &#x60;rating&#x60; and &#x60;single_choice&#x60; question types.         &#x60;&#x60;&#x60;json         {             \&quot;type\&quot;: \&quot;response_based\&quot;,             \&quot;responseValues\&quot;: {                 \&quot;responseKey\&quot;: \&quot;value\&quot;             }         }         &#x60;&#x60;&#x60;          Specific question: Proceeds to a specific question by index.         &#x60;&#x60;&#x60;json         {             \&quot;type\&quot;: \&quot;specific_question\&quot;,             \&quot;index\&quot;: 2         }         &#x60;&#x60;&#x60;          | [optional] 
**Conditions** | **string** |  | [readonly] 
**Appearance** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**ResponsesLimit** | Pointer to **NullableInt32** |  | [optional] 
**FeatureFlagKeys** | **[]interface{}** |  | [readonly] 
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
**UserAccessLevel** | **NullableString** | The effective access level the user has for this object | [readonly] 

## Methods

### NewSurvey

`func NewSurvey(id string, name string, type_ SurveyType, linkedFlag MinimalFeatureFlag, targetingFlag MinimalFeatureFlag, internalTargetingFlag MinimalFeatureFlag, conditions string, createdAt time.Time, createdBy UserBasic, featureFlagKeys []interface{}, userAccessLevel NullableString, ) *Survey`

NewSurvey instantiates a new Survey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyWithDefaults

`func NewSurveyWithDefaults() *Survey`

NewSurveyWithDefaults instantiates a new Survey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Survey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Survey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Survey) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Survey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Survey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Survey) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Survey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Survey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Survey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Survey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Survey) GetType() SurveyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Survey) GetTypeOk() (*SurveyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Survey) SetType(v SurveyType)`

SetType sets Type field to given value.


### GetSchedule

`func (o *Survey) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Survey) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Survey) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Survey) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *Survey) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *Survey) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetLinkedFlag

`func (o *Survey) GetLinkedFlag() MinimalFeatureFlag`

GetLinkedFlag returns the LinkedFlag field if non-nil, zero value otherwise.

### GetLinkedFlagOk

`func (o *Survey) GetLinkedFlagOk() (*MinimalFeatureFlag, bool)`

GetLinkedFlagOk returns a tuple with the LinkedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFlag

`func (o *Survey) SetLinkedFlag(v MinimalFeatureFlag)`

SetLinkedFlag sets LinkedFlag field to given value.


### GetLinkedFlagId

`func (o *Survey) GetLinkedFlagId() int32`

GetLinkedFlagId returns the LinkedFlagId field if non-nil, zero value otherwise.

### GetLinkedFlagIdOk

`func (o *Survey) GetLinkedFlagIdOk() (*int32, bool)`

GetLinkedFlagIdOk returns a tuple with the LinkedFlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFlagId

`func (o *Survey) SetLinkedFlagId(v int32)`

SetLinkedFlagId sets LinkedFlagId field to given value.

### HasLinkedFlagId

`func (o *Survey) HasLinkedFlagId() bool`

HasLinkedFlagId returns a boolean if a field has been set.

### SetLinkedFlagIdNil

`func (o *Survey) SetLinkedFlagIdNil(b bool)`

 SetLinkedFlagIdNil sets the value for LinkedFlagId to be an explicit nil

### UnsetLinkedFlagId
`func (o *Survey) UnsetLinkedFlagId()`

UnsetLinkedFlagId ensures that no value is present for LinkedFlagId, not even an explicit nil
### GetTargetingFlag

`func (o *Survey) GetTargetingFlag() MinimalFeatureFlag`

GetTargetingFlag returns the TargetingFlag field if non-nil, zero value otherwise.

### GetTargetingFlagOk

`func (o *Survey) GetTargetingFlagOk() (*MinimalFeatureFlag, bool)`

GetTargetingFlagOk returns a tuple with the TargetingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingFlag

`func (o *Survey) SetTargetingFlag(v MinimalFeatureFlag)`

SetTargetingFlag sets TargetingFlag field to given value.


### GetInternalTargetingFlag

`func (o *Survey) GetInternalTargetingFlag() MinimalFeatureFlag`

GetInternalTargetingFlag returns the InternalTargetingFlag field if non-nil, zero value otherwise.

### GetInternalTargetingFlagOk

`func (o *Survey) GetInternalTargetingFlagOk() (*MinimalFeatureFlag, bool)`

GetInternalTargetingFlagOk returns a tuple with the InternalTargetingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTargetingFlag

`func (o *Survey) SetInternalTargetingFlag(v MinimalFeatureFlag)`

SetInternalTargetingFlag sets InternalTargetingFlag field to given value.


### GetQuestions

`func (o *Survey) GetQuestions() interface{}`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *Survey) GetQuestionsOk() (*interface{}, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *Survey) SetQuestions(v interface{})`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *Survey) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### SetQuestionsNil

`func (o *Survey) SetQuestionsNil(b bool)`

 SetQuestionsNil sets the value for Questions to be an explicit nil

### UnsetQuestions
`func (o *Survey) UnsetQuestions()`

UnsetQuestions ensures that no value is present for Questions, not even an explicit nil
### GetConditions

`func (o *Survey) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Survey) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Survey) SetConditions(v string)`

SetConditions sets Conditions field to given value.


### GetAppearance

`func (o *Survey) GetAppearance() interface{}`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *Survey) GetAppearanceOk() (*interface{}, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *Survey) SetAppearance(v interface{})`

SetAppearance sets Appearance field to given value.

### HasAppearance

`func (o *Survey) HasAppearance() bool`

HasAppearance returns a boolean if a field has been set.

### SetAppearanceNil

`func (o *Survey) SetAppearanceNil(b bool)`

 SetAppearanceNil sets the value for Appearance to be an explicit nil

### UnsetAppearance
`func (o *Survey) UnsetAppearance()`

UnsetAppearance ensures that no value is present for Appearance, not even an explicit nil
### GetCreatedAt

`func (o *Survey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Survey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Survey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Survey) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Survey) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Survey) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetStartDate

`func (o *Survey) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Survey) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Survey) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Survey) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *Survey) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Survey) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *Survey) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Survey) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Survey) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Survey) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *Survey) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Survey) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetArchived

`func (o *Survey) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Survey) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Survey) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Survey) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetResponsesLimit

`func (o *Survey) GetResponsesLimit() int32`

GetResponsesLimit returns the ResponsesLimit field if non-nil, zero value otherwise.

### GetResponsesLimitOk

`func (o *Survey) GetResponsesLimitOk() (*int32, bool)`

GetResponsesLimitOk returns a tuple with the ResponsesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsesLimit

`func (o *Survey) SetResponsesLimit(v int32)`

SetResponsesLimit sets ResponsesLimit field to given value.

### HasResponsesLimit

`func (o *Survey) HasResponsesLimit() bool`

HasResponsesLimit returns a boolean if a field has been set.

### SetResponsesLimitNil

`func (o *Survey) SetResponsesLimitNil(b bool)`

 SetResponsesLimitNil sets the value for ResponsesLimit to be an explicit nil

### UnsetResponsesLimit
`func (o *Survey) UnsetResponsesLimit()`

UnsetResponsesLimit ensures that no value is present for ResponsesLimit, not even an explicit nil
### GetFeatureFlagKeys

`func (o *Survey) GetFeatureFlagKeys() []interface{}`

GetFeatureFlagKeys returns the FeatureFlagKeys field if non-nil, zero value otherwise.

### GetFeatureFlagKeysOk

`func (o *Survey) GetFeatureFlagKeysOk() (*[]interface{}, bool)`

GetFeatureFlagKeysOk returns a tuple with the FeatureFlagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagKeys

`func (o *Survey) SetFeatureFlagKeys(v []interface{})`

SetFeatureFlagKeys sets FeatureFlagKeys field to given value.


### GetIterationCount

`func (o *Survey) GetIterationCount() int32`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *Survey) GetIterationCountOk() (*int32, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *Survey) SetIterationCount(v int32)`

SetIterationCount sets IterationCount field to given value.

### HasIterationCount

`func (o *Survey) HasIterationCount() bool`

HasIterationCount returns a boolean if a field has been set.

### SetIterationCountNil

`func (o *Survey) SetIterationCountNil(b bool)`

 SetIterationCountNil sets the value for IterationCount to be an explicit nil

### UnsetIterationCount
`func (o *Survey) UnsetIterationCount()`

UnsetIterationCount ensures that no value is present for IterationCount, not even an explicit nil
### GetIterationFrequencyDays

`func (o *Survey) GetIterationFrequencyDays() int32`

GetIterationFrequencyDays returns the IterationFrequencyDays field if non-nil, zero value otherwise.

### GetIterationFrequencyDaysOk

`func (o *Survey) GetIterationFrequencyDaysOk() (*int32, bool)`

GetIterationFrequencyDaysOk returns a tuple with the IterationFrequencyDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationFrequencyDays

`func (o *Survey) SetIterationFrequencyDays(v int32)`

SetIterationFrequencyDays sets IterationFrequencyDays field to given value.

### HasIterationFrequencyDays

`func (o *Survey) HasIterationFrequencyDays() bool`

HasIterationFrequencyDays returns a boolean if a field has been set.

### SetIterationFrequencyDaysNil

`func (o *Survey) SetIterationFrequencyDaysNil(b bool)`

 SetIterationFrequencyDaysNil sets the value for IterationFrequencyDays to be an explicit nil

### UnsetIterationFrequencyDays
`func (o *Survey) UnsetIterationFrequencyDays()`

UnsetIterationFrequencyDays ensures that no value is present for IterationFrequencyDays, not even an explicit nil
### GetIterationStartDates

`func (o *Survey) GetIterationStartDates() []*time.Time`

GetIterationStartDates returns the IterationStartDates field if non-nil, zero value otherwise.

### GetIterationStartDatesOk

`func (o *Survey) GetIterationStartDatesOk() (*[]*time.Time, bool)`

GetIterationStartDatesOk returns a tuple with the IterationStartDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationStartDates

`func (o *Survey) SetIterationStartDates(v []*time.Time)`

SetIterationStartDates sets IterationStartDates field to given value.

### HasIterationStartDates

`func (o *Survey) HasIterationStartDates() bool`

HasIterationStartDates returns a boolean if a field has been set.

### SetIterationStartDatesNil

`func (o *Survey) SetIterationStartDatesNil(b bool)`

 SetIterationStartDatesNil sets the value for IterationStartDates to be an explicit nil

### UnsetIterationStartDates
`func (o *Survey) UnsetIterationStartDates()`

UnsetIterationStartDates ensures that no value is present for IterationStartDates, not even an explicit nil
### GetCurrentIteration

`func (o *Survey) GetCurrentIteration() int32`

GetCurrentIteration returns the CurrentIteration field if non-nil, zero value otherwise.

### GetCurrentIterationOk

`func (o *Survey) GetCurrentIterationOk() (*int32, bool)`

GetCurrentIterationOk returns a tuple with the CurrentIteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIteration

`func (o *Survey) SetCurrentIteration(v int32)`

SetCurrentIteration sets CurrentIteration field to given value.

### HasCurrentIteration

`func (o *Survey) HasCurrentIteration() bool`

HasCurrentIteration returns a boolean if a field has been set.

### SetCurrentIterationNil

`func (o *Survey) SetCurrentIterationNil(b bool)`

 SetCurrentIterationNil sets the value for CurrentIteration to be an explicit nil

### UnsetCurrentIteration
`func (o *Survey) UnsetCurrentIteration()`

UnsetCurrentIteration ensures that no value is present for CurrentIteration, not even an explicit nil
### GetCurrentIterationStartDate

`func (o *Survey) GetCurrentIterationStartDate() time.Time`

GetCurrentIterationStartDate returns the CurrentIterationStartDate field if non-nil, zero value otherwise.

### GetCurrentIterationStartDateOk

`func (o *Survey) GetCurrentIterationStartDateOk() (*time.Time, bool)`

GetCurrentIterationStartDateOk returns a tuple with the CurrentIterationStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIterationStartDate

`func (o *Survey) SetCurrentIterationStartDate(v time.Time)`

SetCurrentIterationStartDate sets CurrentIterationStartDate field to given value.

### HasCurrentIterationStartDate

`func (o *Survey) HasCurrentIterationStartDate() bool`

HasCurrentIterationStartDate returns a boolean if a field has been set.

### SetCurrentIterationStartDateNil

`func (o *Survey) SetCurrentIterationStartDateNil(b bool)`

 SetCurrentIterationStartDateNil sets the value for CurrentIterationStartDate to be an explicit nil

### UnsetCurrentIterationStartDate
`func (o *Survey) UnsetCurrentIterationStartDate()`

UnsetCurrentIterationStartDate ensures that no value is present for CurrentIterationStartDate, not even an explicit nil
### GetResponseSamplingStartDate

`func (o *Survey) GetResponseSamplingStartDate() time.Time`

GetResponseSamplingStartDate returns the ResponseSamplingStartDate field if non-nil, zero value otherwise.

### GetResponseSamplingStartDateOk

`func (o *Survey) GetResponseSamplingStartDateOk() (*time.Time, bool)`

GetResponseSamplingStartDateOk returns a tuple with the ResponseSamplingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSamplingStartDate

`func (o *Survey) SetResponseSamplingStartDate(v time.Time)`

SetResponseSamplingStartDate sets ResponseSamplingStartDate field to given value.

### HasResponseSamplingStartDate

`func (o *Survey) HasResponseSamplingStartDate() bool`

HasResponseSamplingStartDate returns a boolean if a field has been set.

### SetResponseSamplingStartDateNil

`func (o *Survey) SetResponseSamplingStartDateNil(b bool)`

 SetResponseSamplingStartDateNil sets the value for ResponseSamplingStartDate to be an explicit nil

### UnsetResponseSamplingStartDate
`func (o *Survey) UnsetResponseSamplingStartDate()`

UnsetResponseSamplingStartDate ensures that no value is present for ResponseSamplingStartDate, not even an explicit nil
### GetResponseSamplingIntervalType

`func (o *Survey) GetResponseSamplingIntervalType() string`

GetResponseSamplingIntervalType returns the ResponseSamplingIntervalType field if non-nil, zero value otherwise.

### GetResponseSamplingIntervalTypeOk

`func (o *Survey) GetResponseSamplingIntervalTypeOk() (*string, bool)`

GetResponseSamplingIntervalTypeOk returns a tuple with the ResponseSamplingIntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSamplingIntervalType

`func (o *Survey) SetResponseSamplingIntervalType(v string)`

SetResponseSamplingIntervalType sets ResponseSamplingIntervalType field to given value.

### HasResponseSamplingIntervalType

`func (o *Survey) HasResponseSamplingIntervalType() bool`

HasResponseSamplingIntervalType returns a boolean if a field has been set.

### SetResponseSamplingIntervalTypeNil

`func (o *Survey) SetResponseSamplingIntervalTypeNil(b bool)`

 SetResponseSamplingIntervalTypeNil sets the value for ResponseSamplingIntervalType to be an explicit nil

### UnsetResponseSamplingIntervalType
`func (o *Survey) UnsetResponseSamplingIntervalType()`

UnsetResponseSamplingIntervalType ensures that no value is present for ResponseSamplingIntervalType, not even an explicit nil
### GetResponseSamplingInterval

`func (o *Survey) GetResponseSamplingInterval() int32`

GetResponseSamplingInterval returns the ResponseSamplingInterval field if non-nil, zero value otherwise.

### GetResponseSamplingIntervalOk

`func (o *Survey) GetResponseSamplingIntervalOk() (*int32, bool)`

GetResponseSamplingIntervalOk returns a tuple with the ResponseSamplingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSamplingInterval

`func (o *Survey) SetResponseSamplingInterval(v int32)`

SetResponseSamplingInterval sets ResponseSamplingInterval field to given value.

### HasResponseSamplingInterval

`func (o *Survey) HasResponseSamplingInterval() bool`

HasResponseSamplingInterval returns a boolean if a field has been set.

### SetResponseSamplingIntervalNil

`func (o *Survey) SetResponseSamplingIntervalNil(b bool)`

 SetResponseSamplingIntervalNil sets the value for ResponseSamplingInterval to be an explicit nil

### UnsetResponseSamplingInterval
`func (o *Survey) UnsetResponseSamplingInterval()`

UnsetResponseSamplingInterval ensures that no value is present for ResponseSamplingInterval, not even an explicit nil
### GetResponseSamplingLimit

`func (o *Survey) GetResponseSamplingLimit() int32`

GetResponseSamplingLimit returns the ResponseSamplingLimit field if non-nil, zero value otherwise.

### GetResponseSamplingLimitOk

`func (o *Survey) GetResponseSamplingLimitOk() (*int32, bool)`

GetResponseSamplingLimitOk returns a tuple with the ResponseSamplingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSamplingLimit

`func (o *Survey) SetResponseSamplingLimit(v int32)`

SetResponseSamplingLimit sets ResponseSamplingLimit field to given value.

### HasResponseSamplingLimit

`func (o *Survey) HasResponseSamplingLimit() bool`

HasResponseSamplingLimit returns a boolean if a field has been set.

### SetResponseSamplingLimitNil

`func (o *Survey) SetResponseSamplingLimitNil(b bool)`

 SetResponseSamplingLimitNil sets the value for ResponseSamplingLimit to be an explicit nil

### UnsetResponseSamplingLimit
`func (o *Survey) UnsetResponseSamplingLimit()`

UnsetResponseSamplingLimit ensures that no value is present for ResponseSamplingLimit, not even an explicit nil
### GetResponseSamplingDailyLimits

`func (o *Survey) GetResponseSamplingDailyLimits() interface{}`

GetResponseSamplingDailyLimits returns the ResponseSamplingDailyLimits field if non-nil, zero value otherwise.

### GetResponseSamplingDailyLimitsOk

`func (o *Survey) GetResponseSamplingDailyLimitsOk() (*interface{}, bool)`

GetResponseSamplingDailyLimitsOk returns a tuple with the ResponseSamplingDailyLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSamplingDailyLimits

`func (o *Survey) SetResponseSamplingDailyLimits(v interface{})`

SetResponseSamplingDailyLimits sets ResponseSamplingDailyLimits field to given value.

### HasResponseSamplingDailyLimits

`func (o *Survey) HasResponseSamplingDailyLimits() bool`

HasResponseSamplingDailyLimits returns a boolean if a field has been set.

### SetResponseSamplingDailyLimitsNil

`func (o *Survey) SetResponseSamplingDailyLimitsNil(b bool)`

 SetResponseSamplingDailyLimitsNil sets the value for ResponseSamplingDailyLimits to be an explicit nil

### UnsetResponseSamplingDailyLimits
`func (o *Survey) UnsetResponseSamplingDailyLimits()`

UnsetResponseSamplingDailyLimits ensures that no value is present for ResponseSamplingDailyLimits, not even an explicit nil
### GetEnablePartialResponses

`func (o *Survey) GetEnablePartialResponses() bool`

GetEnablePartialResponses returns the EnablePartialResponses field if non-nil, zero value otherwise.

### GetEnablePartialResponsesOk

`func (o *Survey) GetEnablePartialResponsesOk() (*bool, bool)`

GetEnablePartialResponsesOk returns a tuple with the EnablePartialResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartialResponses

`func (o *Survey) SetEnablePartialResponses(v bool)`

SetEnablePartialResponses sets EnablePartialResponses field to given value.

### HasEnablePartialResponses

`func (o *Survey) HasEnablePartialResponses() bool`

HasEnablePartialResponses returns a boolean if a field has been set.

### SetEnablePartialResponsesNil

`func (o *Survey) SetEnablePartialResponsesNil(b bool)`

 SetEnablePartialResponsesNil sets the value for EnablePartialResponses to be an explicit nil

### UnsetEnablePartialResponses
`func (o *Survey) UnsetEnablePartialResponses()`

UnsetEnablePartialResponses ensures that no value is present for EnablePartialResponses, not even an explicit nil
### GetUserAccessLevel

`func (o *Survey) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *Survey) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *Survey) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.


### SetUserAccessLevelNil

`func (o *Survey) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *Survey) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


