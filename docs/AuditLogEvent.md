# AuditLogEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action that was performed (e.g. notification.sent, segment.created, member.invited). | [optional] 
**Actor** | Pointer to [**AuditLogActor**](AuditLogActor.md) |  | [optional] 
**AppId** | Pointer to **string** | UUID of the app the event is associated with. Absent for org-level events. | [optional] 
**Context** | Pointer to [**AuditLogContext**](AuditLogContext.md) |  | [optional] 
**Id** | Pointer to **string** | UUID of the audit log event. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Additional event-specific data that does not fit into the standard fields. | [optional] 
**OccurredAt** | Pointer to **string** | RFC 3339 timestamp of when the event occurred (e.g. 2026-02-18T12:34:56Z). | [optional] 
**OrganizationId** | Pointer to **string** | UUID of the organization the event belongs to. | [optional] 
**Targets** | Pointer to [**[]AuditLogTarget**](AuditLogTarget.md) | The resources the action was performed on. May be empty for org-level events. | [optional] 
**Version** | Pointer to **int32** | Schema version of the event payload. | [optional] 

## Methods

### NewAuditLogEvent

`func NewAuditLogEvent() *AuditLogEvent`

NewAuditLogEvent instantiates a new AuditLogEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEventWithDefaults

`func NewAuditLogEventWithDefaults() *AuditLogEvent`

NewAuditLogEventWithDefaults instantiates a new AuditLogEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuditLogEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditLogEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditLogEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuditLogEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActor

`func (o *AuditLogEvent) GetActor() AuditLogActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AuditLogEvent) GetActorOk() (*AuditLogActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AuditLogEvent) SetActor(v AuditLogActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AuditLogEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAppId

`func (o *AuditLogEvent) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AuditLogEvent) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AuditLogEvent) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AuditLogEvent) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetContext

`func (o *AuditLogEvent) GetContext() AuditLogContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AuditLogEvent) GetContextOk() (*AuditLogContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AuditLogEvent) SetContext(v AuditLogContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *AuditLogEvent) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *AuditLogEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLogEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *AuditLogEvent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AuditLogEvent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AuditLogEvent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AuditLogEvent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOccurredAt

`func (o *AuditLogEvent) GetOccurredAt() string`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *AuditLogEvent) GetOccurredAtOk() (*string, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *AuditLogEvent) SetOccurredAt(v string)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *AuditLogEvent) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AuditLogEvent) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AuditLogEvent) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AuditLogEvent) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AuditLogEvent) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetTargets

`func (o *AuditLogEvent) GetTargets() []AuditLogTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AuditLogEvent) GetTargetsOk() (*[]AuditLogTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AuditLogEvent) SetTargets(v []AuditLogTarget)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *AuditLogEvent) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetVersion

`func (o *AuditLogEvent) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuditLogEvent) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuditLogEvent) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AuditLogEvent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


