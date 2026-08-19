# JourneyListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Journey UUID. Read-only. | [optional] 
**AppId** | Pointer to **string** | UUID of the app the journey belongs to. Read-only. | [optional] 
**Name** | Pointer to **string** | Journey name, up to 300 characters. | [optional] 
**State** | Pointer to **string** | Journey state. New journeys are created as draft. processing is transient while activation is in progress. archived is a journey that has been stopped. Change it through the state field on Update journey. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 creation time. Read-only. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 last-update time. Read-only. | [optional] 
**StartedAt** | Pointer to **NullableString** | ISO 8601 time the journey was activated, or null. Read-only. | [optional] 
**ArchivedAt** | Pointer to **NullableString** | ISO 8601 time the journey was archived, or null. Read-only. | [optional] 
**CreatedSource** | Pointer to **NullableString** | Origin of the journey, for example public_api or dashboard. Read-only. | [optional] 
**Schedule** | Pointer to [**NullableJourneySchedule**](JourneySchedule.md) |  | [optional] 
**Audience** | Pointer to [**JourneyListAudience**](JourneyListAudience.md) |  | [optional] 
**ReentryRules** | Pointer to [**NullableJourneyReentryRules**](JourneyReentryRules.md) |  | [optional] 

## Methods

### NewJourneyListItem

`func NewJourneyListItem() *JourneyListItem`

NewJourneyListItem instantiates a new JourneyListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyListItemWithDefaults

`func NewJourneyListItemWithDefaults() *JourneyListItem`

NewJourneyListItemWithDefaults instantiates a new JourneyListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JourneyListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JourneyListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JourneyListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JourneyListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppId

`func (o *JourneyListItem) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *JourneyListItem) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *JourneyListItem) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *JourneyListItem) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetName

`func (o *JourneyListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JourneyListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JourneyListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JourneyListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *JourneyListItem) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JourneyListItem) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JourneyListItem) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *JourneyListItem) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreatedAt

`func (o *JourneyListItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JourneyListItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JourneyListItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *JourneyListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *JourneyListItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *JourneyListItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *JourneyListItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *JourneyListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *JourneyListItem) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *JourneyListItem) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *JourneyListItem) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *JourneyListItem) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *JourneyListItem) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *JourneyListItem) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetArchivedAt

`func (o *JourneyListItem) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *JourneyListItem) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *JourneyListItem) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *JourneyListItem) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### SetArchivedAtNil

`func (o *JourneyListItem) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *JourneyListItem) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetCreatedSource

`func (o *JourneyListItem) GetCreatedSource() string`

GetCreatedSource returns the CreatedSource field if non-nil, zero value otherwise.

### GetCreatedSourceOk

`func (o *JourneyListItem) GetCreatedSourceOk() (*string, bool)`

GetCreatedSourceOk returns a tuple with the CreatedSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedSource

`func (o *JourneyListItem) SetCreatedSource(v string)`

SetCreatedSource sets CreatedSource field to given value.

### HasCreatedSource

`func (o *JourneyListItem) HasCreatedSource() bool`

HasCreatedSource returns a boolean if a field has been set.

### SetCreatedSourceNil

`func (o *JourneyListItem) SetCreatedSourceNil(b bool)`

 SetCreatedSourceNil sets the value for CreatedSource to be an explicit nil

### UnsetCreatedSource
`func (o *JourneyListItem) UnsetCreatedSource()`

UnsetCreatedSource ensures that no value is present for CreatedSource, not even an explicit nil
### GetSchedule

`func (o *JourneyListItem) GetSchedule() JourneySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *JourneyListItem) GetScheduleOk() (*JourneySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *JourneyListItem) SetSchedule(v JourneySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *JourneyListItem) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *JourneyListItem) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *JourneyListItem) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetAudience

`func (o *JourneyListItem) GetAudience() JourneyListAudience`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *JourneyListItem) GetAudienceOk() (*JourneyListAudience, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *JourneyListItem) SetAudience(v JourneyListAudience)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *JourneyListItem) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetReentryRules

`func (o *JourneyListItem) GetReentryRules() JourneyReentryRules`

GetReentryRules returns the ReentryRules field if non-nil, zero value otherwise.

### GetReentryRulesOk

`func (o *JourneyListItem) GetReentryRulesOk() (*JourneyReentryRules, bool)`

GetReentryRulesOk returns a tuple with the ReentryRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReentryRules

`func (o *JourneyListItem) SetReentryRules(v JourneyReentryRules)`

SetReentryRules sets ReentryRules field to given value.

### HasReentryRules

`func (o *JourneyListItem) HasReentryRules() bool`

HasReentryRules returns a boolean if a field has been set.

### SetReentryRulesNil

`func (o *JourneyListItem) SetReentryRulesNil(b bool)`

 SetReentryRulesNil sets the value for ReentryRules to be an explicit nil

### UnsetReentryRules
`func (o *JourneyListItem) UnsetReentryRules()`

UnsetReentryRules ensures that no value is present for ReentryRules, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


