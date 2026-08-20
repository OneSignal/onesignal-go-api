# UpdateJourneyNodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientNodeId** | Pointer to **string** | Optional client-assigned identifier, unique within the journey. Use it to reference this node from elsewhere in the same request. Persisted and returned on reads. | [optional] 
**Annotation** | Pointer to **string** | Optional free-text label, up to 255 characters. Stored and returned as-is with no effect on journey behavior. | [optional] 
**DurationSeconds** | Pointer to **NullableInt32** | wait nodes: seconds to hold the user. Minimum 60, maximum 31556952 (1 year). | [optional] 
**RelativeTo** | Pointer to **string** | time_window nodes: schedule_in_timezone uses the configured windows; last_active_time holds relative to the user&#39;s last active time. | [optional] 
**Windows** | Pointer to [**[]JourneyTimeWindow**](JourneyTimeWindow.md) | time_window nodes: one or more time windows. A window with no day_of_week applies to every day. Required when relative_to is schedule_in_timezone; omit when it is last_active_time. | [optional] 
**TimeZone** | Pointer to **string** | time_window nodes: IANA timezone identifier used when the user&#39;s timezone is unavailable. | [optional] 
**UseUserTimeZone** | Pointer to **NullableBool** | time_window nodes: when true, uses the user&#39;s timezone if available. | [optional] 
**TemplateId** | Pointer to **string** | send_push, send_email, and send_sms nodes: UUID of the template to send. | [optional] 
**IamId** | Pointer to **string** | send_iam nodes: UUID of the in-app message to send. | [optional] 
**UserTtlSeconds** | Pointer to **NullableInt32** | send_iam nodes: optional time-to-live for the in-app message, in seconds. | [optional] 
**WebhookId** | Pointer to **string** | send_webhook nodes: UUID of the webhook to send. | [optional] 
**Assignments** | Pointer to **map[string]string** | tag nodes: tag key-value pairs to assign. An empty string value removes the tag. Keys are limited to 255 characters and values to 1024. | [optional] 
**RandomizeOnEntry** | Pointer to **NullableBool** | split_range nodes: when true, assigns each user to a branch at random on entry. Defaults to false. | [optional] 
**Branches** | Pointer to [**[]JourneyBranch**](JourneyBranch.md) | Branching nodes: nested branches. split_range requires 2-20 weighted branches that sum to 100. yes_no requires exactly 2 branches. wait_until requires 1-10 condition branches. | [optional] 
**Expiration** | Pointer to [**NullableJourneyWaitUntilExpiration**](JourneyWaitUntilExpiration.md) |  | [optional] 
**ConcurrencyKey** | Pointer to **NullableString** | Optional optimistic-concurrency token. Pass the concurrency_key from a prior fetch to reject the update with 409 if the journey changed. Omit to skip the check. It is not merged onto the node. | [optional] 

## Methods

### NewUpdateJourneyNodeRequest

`func NewUpdateJourneyNodeRequest() *UpdateJourneyNodeRequest`

NewUpdateJourneyNodeRequest instantiates a new UpdateJourneyNodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateJourneyNodeRequestWithDefaults

`func NewUpdateJourneyNodeRequestWithDefaults() *UpdateJourneyNodeRequest`

NewUpdateJourneyNodeRequestWithDefaults instantiates a new UpdateJourneyNodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientNodeId

`func (o *UpdateJourneyNodeRequest) GetClientNodeId() string`

GetClientNodeId returns the ClientNodeId field if non-nil, zero value otherwise.

### GetClientNodeIdOk

`func (o *UpdateJourneyNodeRequest) GetClientNodeIdOk() (*string, bool)`

GetClientNodeIdOk returns a tuple with the ClientNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNodeId

`func (o *UpdateJourneyNodeRequest) SetClientNodeId(v string)`

SetClientNodeId sets ClientNodeId field to given value.

### HasClientNodeId

`func (o *UpdateJourneyNodeRequest) HasClientNodeId() bool`

HasClientNodeId returns a boolean if a field has been set.

### GetAnnotation

`func (o *UpdateJourneyNodeRequest) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *UpdateJourneyNodeRequest) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *UpdateJourneyNodeRequest) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *UpdateJourneyNodeRequest) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *UpdateJourneyNodeRequest) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *UpdateJourneyNodeRequest) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *UpdateJourneyNodeRequest) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *UpdateJourneyNodeRequest) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *UpdateJourneyNodeRequest) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *UpdateJourneyNodeRequest) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetRelativeTo

`func (o *UpdateJourneyNodeRequest) GetRelativeTo() string`

GetRelativeTo returns the RelativeTo field if non-nil, zero value otherwise.

### GetRelativeToOk

`func (o *UpdateJourneyNodeRequest) GetRelativeToOk() (*string, bool)`

GetRelativeToOk returns a tuple with the RelativeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTo

`func (o *UpdateJourneyNodeRequest) SetRelativeTo(v string)`

SetRelativeTo sets RelativeTo field to given value.

### HasRelativeTo

`func (o *UpdateJourneyNodeRequest) HasRelativeTo() bool`

HasRelativeTo returns a boolean if a field has been set.

### GetWindows

`func (o *UpdateJourneyNodeRequest) GetWindows() []JourneyTimeWindow`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *UpdateJourneyNodeRequest) GetWindowsOk() (*[]JourneyTimeWindow, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *UpdateJourneyNodeRequest) SetWindows(v []JourneyTimeWindow)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *UpdateJourneyNodeRequest) HasWindows() bool`

HasWindows returns a boolean if a field has been set.

### GetTimeZone

`func (o *UpdateJourneyNodeRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UpdateJourneyNodeRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UpdateJourneyNodeRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UpdateJourneyNodeRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUseUserTimeZone

`func (o *UpdateJourneyNodeRequest) GetUseUserTimeZone() bool`

GetUseUserTimeZone returns the UseUserTimeZone field if non-nil, zero value otherwise.

### GetUseUserTimeZoneOk

`func (o *UpdateJourneyNodeRequest) GetUseUserTimeZoneOk() (*bool, bool)`

GetUseUserTimeZoneOk returns a tuple with the UseUserTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUserTimeZone

`func (o *UpdateJourneyNodeRequest) SetUseUserTimeZone(v bool)`

SetUseUserTimeZone sets UseUserTimeZone field to given value.

### HasUseUserTimeZone

`func (o *UpdateJourneyNodeRequest) HasUseUserTimeZone() bool`

HasUseUserTimeZone returns a boolean if a field has been set.

### SetUseUserTimeZoneNil

`func (o *UpdateJourneyNodeRequest) SetUseUserTimeZoneNil(b bool)`

 SetUseUserTimeZoneNil sets the value for UseUserTimeZone to be an explicit nil

### UnsetUseUserTimeZone
`func (o *UpdateJourneyNodeRequest) UnsetUseUserTimeZone()`

UnsetUseUserTimeZone ensures that no value is present for UseUserTimeZone, not even an explicit nil
### GetTemplateId

`func (o *UpdateJourneyNodeRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *UpdateJourneyNodeRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *UpdateJourneyNodeRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *UpdateJourneyNodeRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetIamId

`func (o *UpdateJourneyNodeRequest) GetIamId() string`

GetIamId returns the IamId field if non-nil, zero value otherwise.

### GetIamIdOk

`func (o *UpdateJourneyNodeRequest) GetIamIdOk() (*string, bool)`

GetIamIdOk returns a tuple with the IamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamId

`func (o *UpdateJourneyNodeRequest) SetIamId(v string)`

SetIamId sets IamId field to given value.

### HasIamId

`func (o *UpdateJourneyNodeRequest) HasIamId() bool`

HasIamId returns a boolean if a field has been set.

### GetUserTtlSeconds

`func (o *UpdateJourneyNodeRequest) GetUserTtlSeconds() int32`

GetUserTtlSeconds returns the UserTtlSeconds field if non-nil, zero value otherwise.

### GetUserTtlSecondsOk

`func (o *UpdateJourneyNodeRequest) GetUserTtlSecondsOk() (*int32, bool)`

GetUserTtlSecondsOk returns a tuple with the UserTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtlSeconds

`func (o *UpdateJourneyNodeRequest) SetUserTtlSeconds(v int32)`

SetUserTtlSeconds sets UserTtlSeconds field to given value.

### HasUserTtlSeconds

`func (o *UpdateJourneyNodeRequest) HasUserTtlSeconds() bool`

HasUserTtlSeconds returns a boolean if a field has been set.

### SetUserTtlSecondsNil

`func (o *UpdateJourneyNodeRequest) SetUserTtlSecondsNil(b bool)`

 SetUserTtlSecondsNil sets the value for UserTtlSeconds to be an explicit nil

### UnsetUserTtlSeconds
`func (o *UpdateJourneyNodeRequest) UnsetUserTtlSeconds()`

UnsetUserTtlSeconds ensures that no value is present for UserTtlSeconds, not even an explicit nil
### GetWebhookId

`func (o *UpdateJourneyNodeRequest) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *UpdateJourneyNodeRequest) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *UpdateJourneyNodeRequest) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *UpdateJourneyNodeRequest) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetAssignments

`func (o *UpdateJourneyNodeRequest) GetAssignments() map[string]string`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *UpdateJourneyNodeRequest) GetAssignmentsOk() (*map[string]string, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *UpdateJourneyNodeRequest) SetAssignments(v map[string]string)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *UpdateJourneyNodeRequest) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetRandomizeOnEntry

`func (o *UpdateJourneyNodeRequest) GetRandomizeOnEntry() bool`

GetRandomizeOnEntry returns the RandomizeOnEntry field if non-nil, zero value otherwise.

### GetRandomizeOnEntryOk

`func (o *UpdateJourneyNodeRequest) GetRandomizeOnEntryOk() (*bool, bool)`

GetRandomizeOnEntryOk returns a tuple with the RandomizeOnEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizeOnEntry

`func (o *UpdateJourneyNodeRequest) SetRandomizeOnEntry(v bool)`

SetRandomizeOnEntry sets RandomizeOnEntry field to given value.

### HasRandomizeOnEntry

`func (o *UpdateJourneyNodeRequest) HasRandomizeOnEntry() bool`

HasRandomizeOnEntry returns a boolean if a field has been set.

### SetRandomizeOnEntryNil

`func (o *UpdateJourneyNodeRequest) SetRandomizeOnEntryNil(b bool)`

 SetRandomizeOnEntryNil sets the value for RandomizeOnEntry to be an explicit nil

### UnsetRandomizeOnEntry
`func (o *UpdateJourneyNodeRequest) UnsetRandomizeOnEntry()`

UnsetRandomizeOnEntry ensures that no value is present for RandomizeOnEntry, not even an explicit nil
### GetBranches

`func (o *UpdateJourneyNodeRequest) GetBranches() []JourneyBranch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *UpdateJourneyNodeRequest) GetBranchesOk() (*[]JourneyBranch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *UpdateJourneyNodeRequest) SetBranches(v []JourneyBranch)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *UpdateJourneyNodeRequest) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetExpiration

`func (o *UpdateJourneyNodeRequest) GetExpiration() JourneyWaitUntilExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *UpdateJourneyNodeRequest) GetExpirationOk() (*JourneyWaitUntilExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *UpdateJourneyNodeRequest) SetExpiration(v JourneyWaitUntilExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *UpdateJourneyNodeRequest) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *UpdateJourneyNodeRequest) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *UpdateJourneyNodeRequest) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetConcurrencyKey

`func (o *UpdateJourneyNodeRequest) GetConcurrencyKey() string`

GetConcurrencyKey returns the ConcurrencyKey field if non-nil, zero value otherwise.

### GetConcurrencyKeyOk

`func (o *UpdateJourneyNodeRequest) GetConcurrencyKeyOk() (*string, bool)`

GetConcurrencyKeyOk returns a tuple with the ConcurrencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyKey

`func (o *UpdateJourneyNodeRequest) SetConcurrencyKey(v string)`

SetConcurrencyKey sets ConcurrencyKey field to given value.

### HasConcurrencyKey

`func (o *UpdateJourneyNodeRequest) HasConcurrencyKey() bool`

HasConcurrencyKey returns a boolean if a field has been set.

### SetConcurrencyKeyNil

`func (o *UpdateJourneyNodeRequest) SetConcurrencyKeyNil(b bool)`

 SetConcurrencyKeyNil sets the value for ConcurrencyKey to be an explicit nil

### UnsetConcurrencyKey
`func (o *UpdateJourneyNodeRequest) UnsetConcurrencyKey()`

UnsetConcurrencyKey ensures that no value is present for ConcurrencyKey, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


