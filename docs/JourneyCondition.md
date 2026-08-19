# JourneyCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Condition kind. Selects which other fields apply. | 
**IncludedSegmentIds** | Pointer to **[]string** | segment_membership conditions: Segment UUIDs the user must belong to. | [optional] 
**ExcludedSegmentIds** | Pointer to **[]string** | segment_membership conditions: Segment UUIDs the user must not belong to. | [optional] 
**Action** | Pointer to **string** | on_notification_action conditions: the notification action to branch on. Which actions apply depends on the sending node&#39;s channel. | [optional] 
**SendingNodeId** | Pointer to **string** | on_notification_action conditions: id of the sending node this action refers to. Returned on reads; accepted on write. | [optional] 
**ClientNodeId** | Pointer to **string** | on_notification_action conditions: write-only alternative to sending_node_id. References the sending node by its client_node_id. | [optional] 
**Name** | Pointer to **string** | event_trigger conditions: event name, up to 255 characters. | [optional] 
**Attributes** | Pointer to [**[][]JourneyEventAttribute**]([]JourneyEventAttribute.md) | Event attribute matchers, as a list of condition groups. Send a single group whose conditions are AND&#39;d together. More than one group is rejected. | [optional] 
**EntryEventMatchAttributes** | Pointer to **[]map[string]interface{}** | event_trigger conditions: match incoming event properties against the journey&#39;s entry event. Only valid on event-triggered journeys. | [optional] 

## Methods

### NewJourneyCondition

`func NewJourneyCondition(kind string, ) *JourneyCondition`

NewJourneyCondition instantiates a new JourneyCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyConditionWithDefaults

`func NewJourneyConditionWithDefaults() *JourneyCondition`

NewJourneyConditionWithDefaults instantiates a new JourneyCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *JourneyCondition) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JourneyCondition) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JourneyCondition) SetKind(v string)`

SetKind sets Kind field to given value.


### GetIncludedSegmentIds

`func (o *JourneyCondition) GetIncludedSegmentIds() []string`

GetIncludedSegmentIds returns the IncludedSegmentIds field if non-nil, zero value otherwise.

### GetIncludedSegmentIdsOk

`func (o *JourneyCondition) GetIncludedSegmentIdsOk() (*[]string, bool)`

GetIncludedSegmentIdsOk returns a tuple with the IncludedSegmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegmentIds

`func (o *JourneyCondition) SetIncludedSegmentIds(v []string)`

SetIncludedSegmentIds sets IncludedSegmentIds field to given value.

### HasIncludedSegmentIds

`func (o *JourneyCondition) HasIncludedSegmentIds() bool`

HasIncludedSegmentIds returns a boolean if a field has been set.

### GetExcludedSegmentIds

`func (o *JourneyCondition) GetExcludedSegmentIds() []string`

GetExcludedSegmentIds returns the ExcludedSegmentIds field if non-nil, zero value otherwise.

### GetExcludedSegmentIdsOk

`func (o *JourneyCondition) GetExcludedSegmentIdsOk() (*[]string, bool)`

GetExcludedSegmentIdsOk returns a tuple with the ExcludedSegmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegmentIds

`func (o *JourneyCondition) SetExcludedSegmentIds(v []string)`

SetExcludedSegmentIds sets ExcludedSegmentIds field to given value.

### HasExcludedSegmentIds

`func (o *JourneyCondition) HasExcludedSegmentIds() bool`

HasExcludedSegmentIds returns a boolean if a field has been set.

### GetAction

`func (o *JourneyCondition) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *JourneyCondition) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *JourneyCondition) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *JourneyCondition) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSendingNodeId

`func (o *JourneyCondition) GetSendingNodeId() string`

GetSendingNodeId returns the SendingNodeId field if non-nil, zero value otherwise.

### GetSendingNodeIdOk

`func (o *JourneyCondition) GetSendingNodeIdOk() (*string, bool)`

GetSendingNodeIdOk returns a tuple with the SendingNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendingNodeId

`func (o *JourneyCondition) SetSendingNodeId(v string)`

SetSendingNodeId sets SendingNodeId field to given value.

### HasSendingNodeId

`func (o *JourneyCondition) HasSendingNodeId() bool`

HasSendingNodeId returns a boolean if a field has been set.

### GetClientNodeId

`func (o *JourneyCondition) GetClientNodeId() string`

GetClientNodeId returns the ClientNodeId field if non-nil, zero value otherwise.

### GetClientNodeIdOk

`func (o *JourneyCondition) GetClientNodeIdOk() (*string, bool)`

GetClientNodeIdOk returns a tuple with the ClientNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNodeId

`func (o *JourneyCondition) SetClientNodeId(v string)`

SetClientNodeId sets ClientNodeId field to given value.

### HasClientNodeId

`func (o *JourneyCondition) HasClientNodeId() bool`

HasClientNodeId returns a boolean if a field has been set.

### GetName

`func (o *JourneyCondition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JourneyCondition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JourneyCondition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JourneyCondition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAttributes

`func (o *JourneyCondition) GetAttributes() [][]JourneyEventAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *JourneyCondition) GetAttributesOk() (*[][]JourneyEventAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *JourneyCondition) SetAttributes(v [][]JourneyEventAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *JourneyCondition) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEntryEventMatchAttributes

`func (o *JourneyCondition) GetEntryEventMatchAttributes() []map[string]interface{}`

GetEntryEventMatchAttributes returns the EntryEventMatchAttributes field if non-nil, zero value otherwise.

### GetEntryEventMatchAttributesOk

`func (o *JourneyCondition) GetEntryEventMatchAttributesOk() (*[]map[string]interface{}, bool)`

GetEntryEventMatchAttributesOk returns a tuple with the EntryEventMatchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryEventMatchAttributes

`func (o *JourneyCondition) SetEntryEventMatchAttributes(v []map[string]interface{})`

SetEntryEventMatchAttributes sets EntryEventMatchAttributes field to given value.

### HasEntryEventMatchAttributes

`func (o *JourneyCondition) HasEntryEventMatchAttributes() bool`

HasEntryEventMatchAttributes returns a boolean if a field has been set.

### SetEntryEventMatchAttributesNil

`func (o *JourneyCondition) SetEntryEventMatchAttributesNil(b bool)`

 SetEntryEventMatchAttributesNil sets the value for EntryEventMatchAttributes to be an explicit nil

### UnsetEntryEventMatchAttributes
`func (o *JourneyCondition) UnsetEntryEventMatchAttributes()`

UnsetEntryEventMatchAttributes ensures that no value is present for EntryEventMatchAttributes, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


