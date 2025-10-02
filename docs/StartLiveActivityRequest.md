# StartLiveActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | An internal name to assist with your campaign organization. This does not get displayed in the message itself. | 
**Event** | **string** |  | [default to "start"]
**ActivityId** | **string** | Set a unique activity_id to track and manage the Live Activity. | 
**EventAttributes** | **map[string]interface{}** | Default/static data to initialize the Live Activity upon start. | 
**EventUpdates** | **map[string]interface{}** | Dynamic content used to update the running Live Activity at start. Must match the ContentState interface defined in your app. | 
**Contents** | [**LanguageStringMap**](LanguageStringMap.md) |  | 
**Headings** | [**LanguageStringMap**](LanguageStringMap.md) |  | 
**StaleDate** | Pointer to **int32** | Accepts Unix timestamp in seconds. When time reaches the configured stale date, the system considers the Live Activity out of date, and the ActivityState of the Live Activity changes to ActivityState.stale. | [optional] 
**Priority** | Pointer to **int32** | Delivery priority through the push provider (APNs). Pass 10 for higher priority notifications, or 5 for lower priority notifications. Lower priority notifications are sent based on the power considerations of the end user&#39;s device. If not set, defaults to 10. | [optional] 
**IosRelevanceScore** | Pointer to **NullableFloat32** | iOS 15+. A score to indicate how a notification should be displayed when grouped. Use a float between 0-1. | [optional] 
**IdempotencyKey** | Pointer to **NullableString** | Correlation and idempotency key. A request received with this parameter will first look for another notification with the same idempotency key. If one exists, a notification will not be sent, and result of the previous operation will instead be returned. Therefore, if you plan on using this feature, it&#39;s important to use a good source of randomness to generate the UUID passed here. This key is only idempotent for 30 days. After 30 days, the notification could be removed from our system and a notification with the same idempotency key will be sent again.   See Idempotent Notification Requests for more details writeOnly: true  | [optional] 
**IncludeAliases** | Pointer to **map[string][]string** | Target specific users by aliases assigned via API. An alias can be an external_id, onesignal_id, or a custom alias. Accepts an object where keys are alias labels and values are arrays of alias IDs to include Example usage: { \&quot;external_id\&quot;: [\&quot;exId1\&quot;, \&quot;extId2\&quot;], \&quot;internal_label\&quot;: [\&quot;id1\&quot;, \&quot;id2\&quot;] } Not compatible with any other targeting parameters. REQUIRED: REST API Key Authentication Limit of 2,000 entries per REST API call Note: If targeting push, email, or sms subscribers with same ids, use with target_channel to indicate you are sending a push or email or sms. | [optional] 
**IncludeSubscriptionIds** | Pointer to **[]string** | Specific subscription ids to target. Not compatible with other targeting parameters. | [optional] 
**IncludedSegments** | Pointer to **[]string** | Segment names to include. Only compatible with excluded_segments. | [optional] 
**ExcludedSegments** | Pointer to **[]string** | Segment names to exclude. Only compatible with included_segments. | [optional] 
**Filters** | Pointer to [**[]FilterExpression**](FilterExpression.md) |  | [optional] 

## Methods

### NewStartLiveActivityRequest

`func NewStartLiveActivityRequest(name string, event string, activityId string, eventAttributes map[string]interface{}, eventUpdates map[string]interface{}, contents LanguageStringMap, headings LanguageStringMap, ) *StartLiveActivityRequest`

NewStartLiveActivityRequest instantiates a new StartLiveActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartLiveActivityRequestWithDefaults

`func NewStartLiveActivityRequestWithDefaults() *StartLiveActivityRequest`

NewStartLiveActivityRequestWithDefaults instantiates a new StartLiveActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StartLiveActivityRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StartLiveActivityRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StartLiveActivityRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEvent

`func (o *StartLiveActivityRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *StartLiveActivityRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *StartLiveActivityRequest) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetActivityId

`func (o *StartLiveActivityRequest) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *StartLiveActivityRequest) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *StartLiveActivityRequest) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetEventAttributes

`func (o *StartLiveActivityRequest) GetEventAttributes() map[string]interface{}`

GetEventAttributes returns the EventAttributes field if non-nil, zero value otherwise.

### GetEventAttributesOk

`func (o *StartLiveActivityRequest) GetEventAttributesOk() (*map[string]interface{}, bool)`

GetEventAttributesOk returns a tuple with the EventAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAttributes

`func (o *StartLiveActivityRequest) SetEventAttributes(v map[string]interface{})`

SetEventAttributes sets EventAttributes field to given value.


### GetEventUpdates

`func (o *StartLiveActivityRequest) GetEventUpdates() map[string]interface{}`

GetEventUpdates returns the EventUpdates field if non-nil, zero value otherwise.

### GetEventUpdatesOk

`func (o *StartLiveActivityRequest) GetEventUpdatesOk() (*map[string]interface{}, bool)`

GetEventUpdatesOk returns a tuple with the EventUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUpdates

`func (o *StartLiveActivityRequest) SetEventUpdates(v map[string]interface{})`

SetEventUpdates sets EventUpdates field to given value.


### GetContents

`func (o *StartLiveActivityRequest) GetContents() LanguageStringMap`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *StartLiveActivityRequest) GetContentsOk() (*LanguageStringMap, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *StartLiveActivityRequest) SetContents(v LanguageStringMap)`

SetContents sets Contents field to given value.


### GetHeadings

`func (o *StartLiveActivityRequest) GetHeadings() LanguageStringMap`

GetHeadings returns the Headings field if non-nil, zero value otherwise.

### GetHeadingsOk

`func (o *StartLiveActivityRequest) GetHeadingsOk() (*LanguageStringMap, bool)`

GetHeadingsOk returns a tuple with the Headings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadings

`func (o *StartLiveActivityRequest) SetHeadings(v LanguageStringMap)`

SetHeadings sets Headings field to given value.


### GetStaleDate

`func (o *StartLiveActivityRequest) GetStaleDate() int32`

GetStaleDate returns the StaleDate field if non-nil, zero value otherwise.

### GetStaleDateOk

`func (o *StartLiveActivityRequest) GetStaleDateOk() (*int32, bool)`

GetStaleDateOk returns a tuple with the StaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleDate

`func (o *StartLiveActivityRequest) SetStaleDate(v int32)`

SetStaleDate sets StaleDate field to given value.

### HasStaleDate

`func (o *StartLiveActivityRequest) HasStaleDate() bool`

HasStaleDate returns a boolean if a field has been set.

### GetPriority

`func (o *StartLiveActivityRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *StartLiveActivityRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *StartLiveActivityRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *StartLiveActivityRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetIosRelevanceScore

`func (o *StartLiveActivityRequest) GetIosRelevanceScore() float32`

GetIosRelevanceScore returns the IosRelevanceScore field if non-nil, zero value otherwise.

### GetIosRelevanceScoreOk

`func (o *StartLiveActivityRequest) GetIosRelevanceScoreOk() (*float32, bool)`

GetIosRelevanceScoreOk returns a tuple with the IosRelevanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosRelevanceScore

`func (o *StartLiveActivityRequest) SetIosRelevanceScore(v float32)`

SetIosRelevanceScore sets IosRelevanceScore field to given value.

### HasIosRelevanceScore

`func (o *StartLiveActivityRequest) HasIosRelevanceScore() bool`

HasIosRelevanceScore returns a boolean if a field has been set.

### SetIosRelevanceScoreNil

`func (o *StartLiveActivityRequest) SetIosRelevanceScoreNil(b bool)`

 SetIosRelevanceScoreNil sets the value for IosRelevanceScore to be an explicit nil

### UnsetIosRelevanceScore
`func (o *StartLiveActivityRequest) UnsetIosRelevanceScore()`

UnsetIosRelevanceScore ensures that no value is present for IosRelevanceScore, not even an explicit nil
### GetIdempotencyKey

`func (o *StartLiveActivityRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *StartLiveActivityRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *StartLiveActivityRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *StartLiveActivityRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### SetIdempotencyKeyNil

`func (o *StartLiveActivityRequest) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *StartLiveActivityRequest) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
### GetIncludeAliases

`func (o *StartLiveActivityRequest) GetIncludeAliases() map[string][]string`

GetIncludeAliases returns the IncludeAliases field if non-nil, zero value otherwise.

### GetIncludeAliasesOk

`func (o *StartLiveActivityRequest) GetIncludeAliasesOk() (*map[string][]string, bool)`

GetIncludeAliasesOk returns a tuple with the IncludeAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAliases

`func (o *StartLiveActivityRequest) SetIncludeAliases(v map[string][]string)`

SetIncludeAliases sets IncludeAliases field to given value.

### HasIncludeAliases

`func (o *StartLiveActivityRequest) HasIncludeAliases() bool`

HasIncludeAliases returns a boolean if a field has been set.

### SetIncludeAliasesNil

`func (o *StartLiveActivityRequest) SetIncludeAliasesNil(b bool)`

 SetIncludeAliasesNil sets the value for IncludeAliases to be an explicit nil

### UnsetIncludeAliases
`func (o *StartLiveActivityRequest) UnsetIncludeAliases()`

UnsetIncludeAliases ensures that no value is present for IncludeAliases, not even an explicit nil
### GetIncludeSubscriptionIds

`func (o *StartLiveActivityRequest) GetIncludeSubscriptionIds() []string`

GetIncludeSubscriptionIds returns the IncludeSubscriptionIds field if non-nil, zero value otherwise.

### GetIncludeSubscriptionIdsOk

`func (o *StartLiveActivityRequest) GetIncludeSubscriptionIdsOk() (*[]string, bool)`

GetIncludeSubscriptionIdsOk returns a tuple with the IncludeSubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubscriptionIds

`func (o *StartLiveActivityRequest) SetIncludeSubscriptionIds(v []string)`

SetIncludeSubscriptionIds sets IncludeSubscriptionIds field to given value.

### HasIncludeSubscriptionIds

`func (o *StartLiveActivityRequest) HasIncludeSubscriptionIds() bool`

HasIncludeSubscriptionIds returns a boolean if a field has been set.

### SetIncludeSubscriptionIdsNil

`func (o *StartLiveActivityRequest) SetIncludeSubscriptionIdsNil(b bool)`

 SetIncludeSubscriptionIdsNil sets the value for IncludeSubscriptionIds to be an explicit nil

### UnsetIncludeSubscriptionIds
`func (o *StartLiveActivityRequest) UnsetIncludeSubscriptionIds()`

UnsetIncludeSubscriptionIds ensures that no value is present for IncludeSubscriptionIds, not even an explicit nil
### GetIncludedSegments

`func (o *StartLiveActivityRequest) GetIncludedSegments() []string`

GetIncludedSegments returns the IncludedSegments field if non-nil, zero value otherwise.

### GetIncludedSegmentsOk

`func (o *StartLiveActivityRequest) GetIncludedSegmentsOk() (*[]string, bool)`

GetIncludedSegmentsOk returns a tuple with the IncludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegments

`func (o *StartLiveActivityRequest) SetIncludedSegments(v []string)`

SetIncludedSegments sets IncludedSegments field to given value.

### HasIncludedSegments

`func (o *StartLiveActivityRequest) HasIncludedSegments() bool`

HasIncludedSegments returns a boolean if a field has been set.

### SetIncludedSegmentsNil

`func (o *StartLiveActivityRequest) SetIncludedSegmentsNil(b bool)`

 SetIncludedSegmentsNil sets the value for IncludedSegments to be an explicit nil

### UnsetIncludedSegments
`func (o *StartLiveActivityRequest) UnsetIncludedSegments()`

UnsetIncludedSegments ensures that no value is present for IncludedSegments, not even an explicit nil
### GetExcludedSegments

`func (o *StartLiveActivityRequest) GetExcludedSegments() []string`

GetExcludedSegments returns the ExcludedSegments field if non-nil, zero value otherwise.

### GetExcludedSegmentsOk

`func (o *StartLiveActivityRequest) GetExcludedSegmentsOk() (*[]string, bool)`

GetExcludedSegmentsOk returns a tuple with the ExcludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegments

`func (o *StartLiveActivityRequest) SetExcludedSegments(v []string)`

SetExcludedSegments sets ExcludedSegments field to given value.

### HasExcludedSegments

`func (o *StartLiveActivityRequest) HasExcludedSegments() bool`

HasExcludedSegments returns a boolean if a field has been set.

### SetExcludedSegmentsNil

`func (o *StartLiveActivityRequest) SetExcludedSegmentsNil(b bool)`

 SetExcludedSegmentsNil sets the value for ExcludedSegments to be an explicit nil

### UnsetExcludedSegments
`func (o *StartLiveActivityRequest) UnsetExcludedSegments()`

UnsetExcludedSegments ensures that no value is present for ExcludedSegments, not even an explicit nil
### GetFilters

`func (o *StartLiveActivityRequest) GetFilters() []FilterExpression`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *StartLiveActivityRequest) GetFiltersOk() (*[]FilterExpression, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *StartLiveActivityRequest) SetFilters(v []FilterExpression)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *StartLiveActivityRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *StartLiveActivityRequest) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *StartLiveActivityRequest) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


