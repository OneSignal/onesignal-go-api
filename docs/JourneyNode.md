# JourneyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Server-assigned node UUID. Returned on reads. Required on update to keep an existing node. Rejected on create with a 400 validation error. | [optional] 
**Kind** | **string** | Node kind. Selects which other fields apply. | 
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

## Methods

### NewJourneyNode

`func NewJourneyNode(kind string, ) *JourneyNode`

NewJourneyNode instantiates a new JourneyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyNodeWithDefaults

`func NewJourneyNodeWithDefaults() *JourneyNode`

NewJourneyNodeWithDefaults instantiates a new JourneyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JourneyNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JourneyNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JourneyNode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JourneyNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *JourneyNode) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JourneyNode) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JourneyNode) SetKind(v string)`

SetKind sets Kind field to given value.


### GetClientNodeId

`func (o *JourneyNode) GetClientNodeId() string`

GetClientNodeId returns the ClientNodeId field if non-nil, zero value otherwise.

### GetClientNodeIdOk

`func (o *JourneyNode) GetClientNodeIdOk() (*string, bool)`

GetClientNodeIdOk returns a tuple with the ClientNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNodeId

`func (o *JourneyNode) SetClientNodeId(v string)`

SetClientNodeId sets ClientNodeId field to given value.

### HasClientNodeId

`func (o *JourneyNode) HasClientNodeId() bool`

HasClientNodeId returns a boolean if a field has been set.

### GetAnnotation

`func (o *JourneyNode) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *JourneyNode) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *JourneyNode) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *JourneyNode) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *JourneyNode) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *JourneyNode) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *JourneyNode) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *JourneyNode) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *JourneyNode) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *JourneyNode) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetRelativeTo

`func (o *JourneyNode) GetRelativeTo() string`

GetRelativeTo returns the RelativeTo field if non-nil, zero value otherwise.

### GetRelativeToOk

`func (o *JourneyNode) GetRelativeToOk() (*string, bool)`

GetRelativeToOk returns a tuple with the RelativeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTo

`func (o *JourneyNode) SetRelativeTo(v string)`

SetRelativeTo sets RelativeTo field to given value.

### HasRelativeTo

`func (o *JourneyNode) HasRelativeTo() bool`

HasRelativeTo returns a boolean if a field has been set.

### GetWindows

`func (o *JourneyNode) GetWindows() []JourneyTimeWindow`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *JourneyNode) GetWindowsOk() (*[]JourneyTimeWindow, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *JourneyNode) SetWindows(v []JourneyTimeWindow)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *JourneyNode) HasWindows() bool`

HasWindows returns a boolean if a field has been set.

### GetTimeZone

`func (o *JourneyNode) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *JourneyNode) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *JourneyNode) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *JourneyNode) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUseUserTimeZone

`func (o *JourneyNode) GetUseUserTimeZone() bool`

GetUseUserTimeZone returns the UseUserTimeZone field if non-nil, zero value otherwise.

### GetUseUserTimeZoneOk

`func (o *JourneyNode) GetUseUserTimeZoneOk() (*bool, bool)`

GetUseUserTimeZoneOk returns a tuple with the UseUserTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUserTimeZone

`func (o *JourneyNode) SetUseUserTimeZone(v bool)`

SetUseUserTimeZone sets UseUserTimeZone field to given value.

### HasUseUserTimeZone

`func (o *JourneyNode) HasUseUserTimeZone() bool`

HasUseUserTimeZone returns a boolean if a field has been set.

### SetUseUserTimeZoneNil

`func (o *JourneyNode) SetUseUserTimeZoneNil(b bool)`

 SetUseUserTimeZoneNil sets the value for UseUserTimeZone to be an explicit nil

### UnsetUseUserTimeZone
`func (o *JourneyNode) UnsetUseUserTimeZone()`

UnsetUseUserTimeZone ensures that no value is present for UseUserTimeZone, not even an explicit nil
### GetTemplateId

`func (o *JourneyNode) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *JourneyNode) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *JourneyNode) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *JourneyNode) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetIamId

`func (o *JourneyNode) GetIamId() string`

GetIamId returns the IamId field if non-nil, zero value otherwise.

### GetIamIdOk

`func (o *JourneyNode) GetIamIdOk() (*string, bool)`

GetIamIdOk returns a tuple with the IamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamId

`func (o *JourneyNode) SetIamId(v string)`

SetIamId sets IamId field to given value.

### HasIamId

`func (o *JourneyNode) HasIamId() bool`

HasIamId returns a boolean if a field has been set.

### GetUserTtlSeconds

`func (o *JourneyNode) GetUserTtlSeconds() int32`

GetUserTtlSeconds returns the UserTtlSeconds field if non-nil, zero value otherwise.

### GetUserTtlSecondsOk

`func (o *JourneyNode) GetUserTtlSecondsOk() (*int32, bool)`

GetUserTtlSecondsOk returns a tuple with the UserTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtlSeconds

`func (o *JourneyNode) SetUserTtlSeconds(v int32)`

SetUserTtlSeconds sets UserTtlSeconds field to given value.

### HasUserTtlSeconds

`func (o *JourneyNode) HasUserTtlSeconds() bool`

HasUserTtlSeconds returns a boolean if a field has been set.

### SetUserTtlSecondsNil

`func (o *JourneyNode) SetUserTtlSecondsNil(b bool)`

 SetUserTtlSecondsNil sets the value for UserTtlSeconds to be an explicit nil

### UnsetUserTtlSeconds
`func (o *JourneyNode) UnsetUserTtlSeconds()`

UnsetUserTtlSeconds ensures that no value is present for UserTtlSeconds, not even an explicit nil
### GetWebhookId

`func (o *JourneyNode) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *JourneyNode) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *JourneyNode) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *JourneyNode) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetAssignments

`func (o *JourneyNode) GetAssignments() map[string]string`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *JourneyNode) GetAssignmentsOk() (*map[string]string, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *JourneyNode) SetAssignments(v map[string]string)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *JourneyNode) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetRandomizeOnEntry

`func (o *JourneyNode) GetRandomizeOnEntry() bool`

GetRandomizeOnEntry returns the RandomizeOnEntry field if non-nil, zero value otherwise.

### GetRandomizeOnEntryOk

`func (o *JourneyNode) GetRandomizeOnEntryOk() (*bool, bool)`

GetRandomizeOnEntryOk returns a tuple with the RandomizeOnEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizeOnEntry

`func (o *JourneyNode) SetRandomizeOnEntry(v bool)`

SetRandomizeOnEntry sets RandomizeOnEntry field to given value.

### HasRandomizeOnEntry

`func (o *JourneyNode) HasRandomizeOnEntry() bool`

HasRandomizeOnEntry returns a boolean if a field has been set.

### SetRandomizeOnEntryNil

`func (o *JourneyNode) SetRandomizeOnEntryNil(b bool)`

 SetRandomizeOnEntryNil sets the value for RandomizeOnEntry to be an explicit nil

### UnsetRandomizeOnEntry
`func (o *JourneyNode) UnsetRandomizeOnEntry()`

UnsetRandomizeOnEntry ensures that no value is present for RandomizeOnEntry, not even an explicit nil
### GetBranches

`func (o *JourneyNode) GetBranches() []JourneyBranch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *JourneyNode) GetBranchesOk() (*[]JourneyBranch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *JourneyNode) SetBranches(v []JourneyBranch)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *JourneyNode) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetExpiration

`func (o *JourneyNode) GetExpiration() JourneyWaitUntilExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *JourneyNode) GetExpirationOk() (*JourneyWaitUntilExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *JourneyNode) SetExpiration(v JourneyWaitUntilExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *JourneyNode) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *JourneyNode) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *JourneyNode) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


