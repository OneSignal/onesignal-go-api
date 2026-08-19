# JourneyNodeStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Node kind, repeated here so stats can be read without joining against the journey definition. | [optional] 
**Waiting** | Pointer to **int32** | Users currently held at this node. | [optional] 
**Completed** | Pointer to **int32** | Users who advanced past this node normally. | [optional] 
**ExitedEarly** | Pointer to **int32** | Users who left the journey from this node through an early exit rule. | [optional] 
**MessageStats** | Pointer to [**JourneyMessageStats**](JourneyMessageStats.md) |  | [optional] 

## Methods

### NewJourneyNodeStats

`func NewJourneyNodeStats() *JourneyNodeStats`

NewJourneyNodeStats instantiates a new JourneyNodeStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyNodeStatsWithDefaults

`func NewJourneyNodeStatsWithDefaults() *JourneyNodeStats`

NewJourneyNodeStatsWithDefaults instantiates a new JourneyNodeStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *JourneyNodeStats) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JourneyNodeStats) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JourneyNodeStats) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *JourneyNodeStats) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetWaiting

`func (o *JourneyNodeStats) GetWaiting() int32`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *JourneyNodeStats) GetWaitingOk() (*int32, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *JourneyNodeStats) SetWaiting(v int32)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *JourneyNodeStats) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.

### GetCompleted

`func (o *JourneyNodeStats) GetCompleted() int32`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *JourneyNodeStats) GetCompletedOk() (*int32, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *JourneyNodeStats) SetCompleted(v int32)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *JourneyNodeStats) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetExitedEarly

`func (o *JourneyNodeStats) GetExitedEarly() int32`

GetExitedEarly returns the ExitedEarly field if non-nil, zero value otherwise.

### GetExitedEarlyOk

`func (o *JourneyNodeStats) GetExitedEarlyOk() (*int32, bool)`

GetExitedEarlyOk returns a tuple with the ExitedEarly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitedEarly

`func (o *JourneyNodeStats) SetExitedEarly(v int32)`

SetExitedEarly sets ExitedEarly field to given value.

### HasExitedEarly

`func (o *JourneyNodeStats) HasExitedEarly() bool`

HasExitedEarly returns a boolean if a field has been set.

### GetMessageStats

`func (o *JourneyNodeStats) GetMessageStats() JourneyMessageStats`

GetMessageStats returns the MessageStats field if non-nil, zero value otherwise.

### GetMessageStatsOk

`func (o *JourneyNodeStats) GetMessageStatsOk() (*JourneyMessageStats, bool)`

GetMessageStatsOk returns a tuple with the MessageStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageStats

`func (o *JourneyNodeStats) SetMessageStats(v JourneyMessageStats)`

SetMessageStats sets MessageStats field to given value.

### HasMessageStats

`func (o *JourneyNodeStats) HasMessageStats() bool`

HasMessageStats returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


