# Journey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Journey UUID. Read-only. | [optional] 
**AppId** | Pointer to **string** | UUID of the app the journey belongs to. Read-only. | [optional] 
**Name** | Pointer to **string** | Journey name, up to 300 characters. | [optional] 
**Description** | Pointer to **NullableString** | Journey description, up to 1024 characters. Defaults to an empty string. | [optional] 
**State** | Pointer to **string** | Journey state. New journeys are created as draft. processing is transient while activation is in progress. archived is a journey that has been stopped. Change it through the state field on Update journey. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 creation time. Read-only. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 last-update time. Read-only. | [optional] 
**StartedAt** | Pointer to **NullableString** | ISO 8601 time the journey was activated, or null. Read-only. May stay null briefly after you set state to active: activation is enqueued, and started_at populates once the journey finishes processing. | [optional] 
**ArchivedAt** | Pointer to **NullableString** | ISO 8601 time the journey was archived, or null. Read-only. | [optional] 
**CreatedSource** | Pointer to **NullableString** | Origin of the journey, for example public_api or dashboard. Read-only. | [optional] 
**Audience** | Pointer to [**JourneyAudience**](JourneyAudience.md) |  | [optional] 
**EarlyExit** | Pointer to [**NullableJourneyEarlyExit**](JourneyEarlyExit.md) |  | [optional] 
**ReentryRules** | Pointer to [**NullableJourneyReentryRules**](JourneyReentryRules.md) |  | [optional] 
**Schedule** | Pointer to [**NullableJourneySchedule**](JourneySchedule.md) |  | [optional] 
**Nodes** | Pointer to [**[]JourneyNode**](JourneyNode.md) | Ordered list of journey nodes. | [optional] 
**ConcurrencyKey** | Pointer to **string** | Opaque optimistic-concurrency token. Read-only. Pass it back on update to guard against overwriting a concurrent change (409). Send it back exactly as read; do not construct or parse it. | [optional] 

## Methods

### NewJourney

`func NewJourney() *Journey`

NewJourney instantiates a new Journey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyWithDefaults

`func NewJourneyWithDefaults() *Journey`

NewJourneyWithDefaults instantiates a new Journey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Journey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Journey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Journey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Journey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppId

`func (o *Journey) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Journey) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Journey) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Journey) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetName

`func (o *Journey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Journey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Journey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Journey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Journey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Journey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Journey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Journey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Journey) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Journey) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetState

`func (o *Journey) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Journey) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Journey) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Journey) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Journey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Journey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Journey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Journey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Journey) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Journey) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Journey) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Journey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *Journey) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Journey) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Journey) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Journey) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *Journey) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Journey) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetArchivedAt

`func (o *Journey) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Journey) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Journey) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Journey) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### SetArchivedAtNil

`func (o *Journey) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *Journey) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetCreatedSource

`func (o *Journey) GetCreatedSource() string`

GetCreatedSource returns the CreatedSource field if non-nil, zero value otherwise.

### GetCreatedSourceOk

`func (o *Journey) GetCreatedSourceOk() (*string, bool)`

GetCreatedSourceOk returns a tuple with the CreatedSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedSource

`func (o *Journey) SetCreatedSource(v string)`

SetCreatedSource sets CreatedSource field to given value.

### HasCreatedSource

`func (o *Journey) HasCreatedSource() bool`

HasCreatedSource returns a boolean if a field has been set.

### SetCreatedSourceNil

`func (o *Journey) SetCreatedSourceNil(b bool)`

 SetCreatedSourceNil sets the value for CreatedSource to be an explicit nil

### UnsetCreatedSource
`func (o *Journey) UnsetCreatedSource()`

UnsetCreatedSource ensures that no value is present for CreatedSource, not even an explicit nil
### GetAudience

`func (o *Journey) GetAudience() JourneyAudience`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *Journey) GetAudienceOk() (*JourneyAudience, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *Journey) SetAudience(v JourneyAudience)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *Journey) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetEarlyExit

`func (o *Journey) GetEarlyExit() JourneyEarlyExit`

GetEarlyExit returns the EarlyExit field if non-nil, zero value otherwise.

### GetEarlyExitOk

`func (o *Journey) GetEarlyExitOk() (*JourneyEarlyExit, bool)`

GetEarlyExitOk returns a tuple with the EarlyExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyExit

`func (o *Journey) SetEarlyExit(v JourneyEarlyExit)`

SetEarlyExit sets EarlyExit field to given value.

### HasEarlyExit

`func (o *Journey) HasEarlyExit() bool`

HasEarlyExit returns a boolean if a field has been set.

### SetEarlyExitNil

`func (o *Journey) SetEarlyExitNil(b bool)`

 SetEarlyExitNil sets the value for EarlyExit to be an explicit nil

### UnsetEarlyExit
`func (o *Journey) UnsetEarlyExit()`

UnsetEarlyExit ensures that no value is present for EarlyExit, not even an explicit nil
### GetReentryRules

`func (o *Journey) GetReentryRules() JourneyReentryRules`

GetReentryRules returns the ReentryRules field if non-nil, zero value otherwise.

### GetReentryRulesOk

`func (o *Journey) GetReentryRulesOk() (*JourneyReentryRules, bool)`

GetReentryRulesOk returns a tuple with the ReentryRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReentryRules

`func (o *Journey) SetReentryRules(v JourneyReentryRules)`

SetReentryRules sets ReentryRules field to given value.

### HasReentryRules

`func (o *Journey) HasReentryRules() bool`

HasReentryRules returns a boolean if a field has been set.

### SetReentryRulesNil

`func (o *Journey) SetReentryRulesNil(b bool)`

 SetReentryRulesNil sets the value for ReentryRules to be an explicit nil

### UnsetReentryRules
`func (o *Journey) UnsetReentryRules()`

UnsetReentryRules ensures that no value is present for ReentryRules, not even an explicit nil
### GetSchedule

`func (o *Journey) GetSchedule() JourneySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Journey) GetScheduleOk() (*JourneySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Journey) SetSchedule(v JourneySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Journey) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *Journey) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *Journey) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetNodes

`func (o *Journey) GetNodes() []JourneyNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Journey) GetNodesOk() (*[]JourneyNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Journey) SetNodes(v []JourneyNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Journey) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetConcurrencyKey

`func (o *Journey) GetConcurrencyKey() string`

GetConcurrencyKey returns the ConcurrencyKey field if non-nil, zero value otherwise.

### GetConcurrencyKeyOk

`func (o *Journey) GetConcurrencyKeyOk() (*string, bool)`

GetConcurrencyKeyOk returns a tuple with the ConcurrencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyKey

`func (o *Journey) SetConcurrencyKey(v string)`

SetConcurrencyKey sets ConcurrencyKey field to given value.

### HasConcurrencyKey

`func (o *Journey) HasConcurrencyKey() bool`

HasConcurrencyKey returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


