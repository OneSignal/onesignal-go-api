# CreateJourneyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Journey name, up to 300 characters. | 
**Description** | Pointer to **NullableString** | Optional journey description, up to 1024 characters. | [optional] 
**Audience** | Pointer to [**JourneyAudience**](JourneyAudience.md) |  | [optional] 
**EarlyExit** | Pointer to [**NullableJourneyEarlyExit**](JourneyEarlyExit.md) |  | [optional] 
**ReentryRules** | Pointer to [**NullableJourneyReentryRules**](JourneyReentryRules.md) |  | [optional] 
**Schedule** | Pointer to [**NullableJourneySchedule**](JourneySchedule.md) |  | [optional] 
**Nodes** | Pointer to [**[]JourneyNode**](JourneyNode.md) | Ordered list of journey nodes. Server-assigned id fields are rejected on create. | [optional] 

## Methods

### NewCreateJourneyRequest

`func NewCreateJourneyRequest(name string, ) *CreateJourneyRequest`

NewCreateJourneyRequest instantiates a new CreateJourneyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJourneyRequestWithDefaults

`func NewCreateJourneyRequestWithDefaults() *CreateJourneyRequest`

NewCreateJourneyRequestWithDefaults instantiates a new CreateJourneyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateJourneyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateJourneyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateJourneyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateJourneyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateJourneyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateJourneyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateJourneyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateJourneyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateJourneyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAudience

`func (o *CreateJourneyRequest) GetAudience() JourneyAudience`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CreateJourneyRequest) GetAudienceOk() (*JourneyAudience, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CreateJourneyRequest) SetAudience(v JourneyAudience)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CreateJourneyRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetEarlyExit

`func (o *CreateJourneyRequest) GetEarlyExit() JourneyEarlyExit`

GetEarlyExit returns the EarlyExit field if non-nil, zero value otherwise.

### GetEarlyExitOk

`func (o *CreateJourneyRequest) GetEarlyExitOk() (*JourneyEarlyExit, bool)`

GetEarlyExitOk returns a tuple with the EarlyExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyExit

`func (o *CreateJourneyRequest) SetEarlyExit(v JourneyEarlyExit)`

SetEarlyExit sets EarlyExit field to given value.

### HasEarlyExit

`func (o *CreateJourneyRequest) HasEarlyExit() bool`

HasEarlyExit returns a boolean if a field has been set.

### SetEarlyExitNil

`func (o *CreateJourneyRequest) SetEarlyExitNil(b bool)`

 SetEarlyExitNil sets the value for EarlyExit to be an explicit nil

### UnsetEarlyExit
`func (o *CreateJourneyRequest) UnsetEarlyExit()`

UnsetEarlyExit ensures that no value is present for EarlyExit, not even an explicit nil
### GetReentryRules

`func (o *CreateJourneyRequest) GetReentryRules() JourneyReentryRules`

GetReentryRules returns the ReentryRules field if non-nil, zero value otherwise.

### GetReentryRulesOk

`func (o *CreateJourneyRequest) GetReentryRulesOk() (*JourneyReentryRules, bool)`

GetReentryRulesOk returns a tuple with the ReentryRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReentryRules

`func (o *CreateJourneyRequest) SetReentryRules(v JourneyReentryRules)`

SetReentryRules sets ReentryRules field to given value.

### HasReentryRules

`func (o *CreateJourneyRequest) HasReentryRules() bool`

HasReentryRules returns a boolean if a field has been set.

### SetReentryRulesNil

`func (o *CreateJourneyRequest) SetReentryRulesNil(b bool)`

 SetReentryRulesNil sets the value for ReentryRules to be an explicit nil

### UnsetReentryRules
`func (o *CreateJourneyRequest) UnsetReentryRules()`

UnsetReentryRules ensures that no value is present for ReentryRules, not even an explicit nil
### GetSchedule

`func (o *CreateJourneyRequest) GetSchedule() JourneySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateJourneyRequest) GetScheduleOk() (*JourneySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateJourneyRequest) SetSchedule(v JourneySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreateJourneyRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *CreateJourneyRequest) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *CreateJourneyRequest) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetNodes

`func (o *CreateJourneyRequest) GetNodes() []JourneyNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CreateJourneyRequest) GetNodesOk() (*[]JourneyNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CreateJourneyRequest) SetNodes(v []JourneyNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CreateJourneyRequest) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


