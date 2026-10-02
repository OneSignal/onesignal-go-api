# DuplicateJourneyOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for the copy, up to 300 characters. If you omit it, the copy takes the name of the source plus \&quot; (Copy)\&quot;. | [optional] 
**Description** | Pointer to **NullableString** | Optional journey description, up to 1024 characters. If you omit it, the copy takes the description of the source. Send null to clear it. | [optional] 
**Audience** | Pointer to [**JourneyAudience**](JourneyAudience.md) |  | [optional] 
**EarlyExit** | Pointer to [**NullableJourneyEarlyExit**](JourneyEarlyExit.md) |  | [optional] 
**ReentryRules** | Pointer to [**NullableJourneyReentryRules**](JourneyReentryRules.md) |  | [optional] 
**Schedule** | Pointer to [**NullableJourneySchedule**](JourneySchedule.md) |  | [optional] 
**Nodes** | Pointer to [**[]JourneyNode**](JourneyNode.md) | Full ordered list of nodes. Replaces the copied graph. Server-assigned id fields are rejected. | [optional] 

## Methods

### NewDuplicateJourneyOverrides

`func NewDuplicateJourneyOverrides() *DuplicateJourneyOverrides`

NewDuplicateJourneyOverrides instantiates a new DuplicateJourneyOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicateJourneyOverridesWithDefaults

`func NewDuplicateJourneyOverridesWithDefaults() *DuplicateJourneyOverrides`

NewDuplicateJourneyOverridesWithDefaults instantiates a new DuplicateJourneyOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DuplicateJourneyOverrides) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DuplicateJourneyOverrides) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DuplicateJourneyOverrides) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DuplicateJourneyOverrides) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DuplicateJourneyOverrides) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DuplicateJourneyOverrides) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DuplicateJourneyOverrides) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DuplicateJourneyOverrides) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DuplicateJourneyOverrides) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DuplicateJourneyOverrides) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAudience

`func (o *DuplicateJourneyOverrides) GetAudience() JourneyAudience`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *DuplicateJourneyOverrides) GetAudienceOk() (*JourneyAudience, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *DuplicateJourneyOverrides) SetAudience(v JourneyAudience)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *DuplicateJourneyOverrides) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetEarlyExit

`func (o *DuplicateJourneyOverrides) GetEarlyExit() JourneyEarlyExit`

GetEarlyExit returns the EarlyExit field if non-nil, zero value otherwise.

### GetEarlyExitOk

`func (o *DuplicateJourneyOverrides) GetEarlyExitOk() (*JourneyEarlyExit, bool)`

GetEarlyExitOk returns a tuple with the EarlyExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyExit

`func (o *DuplicateJourneyOverrides) SetEarlyExit(v JourneyEarlyExit)`

SetEarlyExit sets EarlyExit field to given value.

### HasEarlyExit

`func (o *DuplicateJourneyOverrides) HasEarlyExit() bool`

HasEarlyExit returns a boolean if a field has been set.

### SetEarlyExitNil

`func (o *DuplicateJourneyOverrides) SetEarlyExitNil(b bool)`

 SetEarlyExitNil sets the value for EarlyExit to be an explicit nil

### UnsetEarlyExit
`func (o *DuplicateJourneyOverrides) UnsetEarlyExit()`

UnsetEarlyExit ensures that no value is present for EarlyExit, not even an explicit nil
### GetReentryRules

`func (o *DuplicateJourneyOverrides) GetReentryRules() JourneyReentryRules`

GetReentryRules returns the ReentryRules field if non-nil, zero value otherwise.

### GetReentryRulesOk

`func (o *DuplicateJourneyOverrides) GetReentryRulesOk() (*JourneyReentryRules, bool)`

GetReentryRulesOk returns a tuple with the ReentryRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReentryRules

`func (o *DuplicateJourneyOverrides) SetReentryRules(v JourneyReentryRules)`

SetReentryRules sets ReentryRules field to given value.

### HasReentryRules

`func (o *DuplicateJourneyOverrides) HasReentryRules() bool`

HasReentryRules returns a boolean if a field has been set.

### SetReentryRulesNil

`func (o *DuplicateJourneyOverrides) SetReentryRulesNil(b bool)`

 SetReentryRulesNil sets the value for ReentryRules to be an explicit nil

### UnsetReentryRules
`func (o *DuplicateJourneyOverrides) UnsetReentryRules()`

UnsetReentryRules ensures that no value is present for ReentryRules, not even an explicit nil
### GetSchedule

`func (o *DuplicateJourneyOverrides) GetSchedule() JourneySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DuplicateJourneyOverrides) GetScheduleOk() (*JourneySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DuplicateJourneyOverrides) SetSchedule(v JourneySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *DuplicateJourneyOverrides) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *DuplicateJourneyOverrides) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *DuplicateJourneyOverrides) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetNodes

`func (o *DuplicateJourneyOverrides) GetNodes() []JourneyNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *DuplicateJourneyOverrides) GetNodesOk() (*[]JourneyNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *DuplicateJourneyOverrides) SetNodes(v []JourneyNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *DuplicateJourneyOverrides) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


