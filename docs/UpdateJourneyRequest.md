# UpdateJourneyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Journey name. | [optional] 
**Description** | Pointer to **NullableString** | Journey description. Send null to clear it. | [optional] 
**Audience** | Pointer to [**JourneyAudience**](JourneyAudience.md) |  | [optional] 
**EarlyExit** | Pointer to [**NullableJourneyEarlyExit**](JourneyEarlyExit.md) |  | [optional] 
**ReentryRules** | Pointer to [**NullableJourneyReentryRules**](JourneyReentryRules.md) |  | [optional] 
**Schedule** | Pointer to [**NullableJourneySchedule**](JourneySchedule.md) |  | [optional] 
**Nodes** | Pointer to [**[]JourneyNode**](JourneyNode.md) | Full ordered list of nodes, which replaces the existing graph wholesale. Preserve each node&#39;s server-assigned id from a prior fetch to keep in-flight users on that node; omit id to add a new node. | [optional] 
**State** | Pointer to **string** | Target state. Set active to activate a draft journey, or scheduled together with a future schedule.start_at to activate it later. Set archived to stop a running journey; archiving is permanent. Only scheduled and processing journeys can return to draft. | [optional] 
**ConcurrencyKey** | Pointer to **NullableString** | Optional optimistic-concurrency token. Pass the concurrency_key from a prior fetch to reject the update with 409 if the journey changed. Omit to skip the check. | [optional] 

## Methods

### NewUpdateJourneyRequest

`func NewUpdateJourneyRequest() *UpdateJourneyRequest`

NewUpdateJourneyRequest instantiates a new UpdateJourneyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateJourneyRequestWithDefaults

`func NewUpdateJourneyRequestWithDefaults() *UpdateJourneyRequest`

NewUpdateJourneyRequestWithDefaults instantiates a new UpdateJourneyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateJourneyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateJourneyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateJourneyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateJourneyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateJourneyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateJourneyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateJourneyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateJourneyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateJourneyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateJourneyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAudience

`func (o *UpdateJourneyRequest) GetAudience() JourneyAudience`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *UpdateJourneyRequest) GetAudienceOk() (*JourneyAudience, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *UpdateJourneyRequest) SetAudience(v JourneyAudience)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *UpdateJourneyRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetEarlyExit

`func (o *UpdateJourneyRequest) GetEarlyExit() JourneyEarlyExit`

GetEarlyExit returns the EarlyExit field if non-nil, zero value otherwise.

### GetEarlyExitOk

`func (o *UpdateJourneyRequest) GetEarlyExitOk() (*JourneyEarlyExit, bool)`

GetEarlyExitOk returns a tuple with the EarlyExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyExit

`func (o *UpdateJourneyRequest) SetEarlyExit(v JourneyEarlyExit)`

SetEarlyExit sets EarlyExit field to given value.

### HasEarlyExit

`func (o *UpdateJourneyRequest) HasEarlyExit() bool`

HasEarlyExit returns a boolean if a field has been set.

### SetEarlyExitNil

`func (o *UpdateJourneyRequest) SetEarlyExitNil(b bool)`

 SetEarlyExitNil sets the value for EarlyExit to be an explicit nil

### UnsetEarlyExit
`func (o *UpdateJourneyRequest) UnsetEarlyExit()`

UnsetEarlyExit ensures that no value is present for EarlyExit, not even an explicit nil
### GetReentryRules

`func (o *UpdateJourneyRequest) GetReentryRules() JourneyReentryRules`

GetReentryRules returns the ReentryRules field if non-nil, zero value otherwise.

### GetReentryRulesOk

`func (o *UpdateJourneyRequest) GetReentryRulesOk() (*JourneyReentryRules, bool)`

GetReentryRulesOk returns a tuple with the ReentryRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReentryRules

`func (o *UpdateJourneyRequest) SetReentryRules(v JourneyReentryRules)`

SetReentryRules sets ReentryRules field to given value.

### HasReentryRules

`func (o *UpdateJourneyRequest) HasReentryRules() bool`

HasReentryRules returns a boolean if a field has been set.

### SetReentryRulesNil

`func (o *UpdateJourneyRequest) SetReentryRulesNil(b bool)`

 SetReentryRulesNil sets the value for ReentryRules to be an explicit nil

### UnsetReentryRules
`func (o *UpdateJourneyRequest) UnsetReentryRules()`

UnsetReentryRules ensures that no value is present for ReentryRules, not even an explicit nil
### GetSchedule

`func (o *UpdateJourneyRequest) GetSchedule() JourneySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateJourneyRequest) GetScheduleOk() (*JourneySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateJourneyRequest) SetSchedule(v JourneySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpdateJourneyRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *UpdateJourneyRequest) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *UpdateJourneyRequest) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetNodes

`func (o *UpdateJourneyRequest) GetNodes() []JourneyNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *UpdateJourneyRequest) GetNodesOk() (*[]JourneyNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *UpdateJourneyRequest) SetNodes(v []JourneyNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *UpdateJourneyRequest) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetState

`func (o *UpdateJourneyRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateJourneyRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateJourneyRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateJourneyRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetConcurrencyKey

`func (o *UpdateJourneyRequest) GetConcurrencyKey() string`

GetConcurrencyKey returns the ConcurrencyKey field if non-nil, zero value otherwise.

### GetConcurrencyKeyOk

`func (o *UpdateJourneyRequest) GetConcurrencyKeyOk() (*string, bool)`

GetConcurrencyKeyOk returns a tuple with the ConcurrencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyKey

`func (o *UpdateJourneyRequest) SetConcurrencyKey(v string)`

SetConcurrencyKey sets ConcurrencyKey field to given value.

### HasConcurrencyKey

`func (o *UpdateJourneyRequest) HasConcurrencyKey() bool`

HasConcurrencyKey returns a boolean if a field has been set.

### SetConcurrencyKeyNil

`func (o *UpdateJourneyRequest) SetConcurrencyKeyNil(b bool)`

 SetConcurrencyKeyNil sets the value for ConcurrencyKey to be an explicit nil

### UnsetConcurrencyKey
`func (o *UpdateJourneyRequest) UnsetConcurrencyKey()`

UnsetConcurrencyKey ensures that no value is present for ConcurrencyKey, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


