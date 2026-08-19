# JourneyStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | UUID of the journey these stats belong to. | [optional] 
**Started** | Pointer to **int32** | Users who entered the journey. | [optional] 
**Completed** | Pointer to **int32** | Users who reached the end of the journey normally. | [optional] 
**ExitedEarly** | Pointer to **int32** | Users who left the journey through an early exit rule. | [optional] 
**Nodes** | Pointer to [**map[string]JourneyNodeStats**](JourneyNodeStats.md) | Node stats keyed by node id. Includes every node in the graph, at any nesting depth. | [optional] 
**Branches** | Pointer to [**map[string]JourneyBranchStats**](JourneyBranchStats.md) | Branch stats keyed by branch id. Empty for a journey with no branching nodes. | [optional] 

## Methods

### NewJourneyStats

`func NewJourneyStats() *JourneyStats`

NewJourneyStats instantiates a new JourneyStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyStatsWithDefaults

`func NewJourneyStatsWithDefaults() *JourneyStats`

NewJourneyStatsWithDefaults instantiates a new JourneyStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JourneyStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JourneyStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JourneyStats) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JourneyStats) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStarted

`func (o *JourneyStats) GetStarted() int32`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *JourneyStats) GetStartedOk() (*int32, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *JourneyStats) SetStarted(v int32)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *JourneyStats) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetCompleted

`func (o *JourneyStats) GetCompleted() int32`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *JourneyStats) GetCompletedOk() (*int32, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *JourneyStats) SetCompleted(v int32)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *JourneyStats) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetExitedEarly

`func (o *JourneyStats) GetExitedEarly() int32`

GetExitedEarly returns the ExitedEarly field if non-nil, zero value otherwise.

### GetExitedEarlyOk

`func (o *JourneyStats) GetExitedEarlyOk() (*int32, bool)`

GetExitedEarlyOk returns a tuple with the ExitedEarly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitedEarly

`func (o *JourneyStats) SetExitedEarly(v int32)`

SetExitedEarly sets ExitedEarly field to given value.

### HasExitedEarly

`func (o *JourneyStats) HasExitedEarly() bool`

HasExitedEarly returns a boolean if a field has been set.

### GetNodes

`func (o *JourneyStats) GetNodes() map[string]JourneyNodeStats`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *JourneyStats) GetNodesOk() (*map[string]JourneyNodeStats, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *JourneyStats) SetNodes(v map[string]JourneyNodeStats)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *JourneyStats) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetBranches

`func (o *JourneyStats) GetBranches() map[string]JourneyBranchStats`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *JourneyStats) GetBranchesOk() (*map[string]JourneyBranchStats, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *JourneyStats) SetBranches(v map[string]JourneyBranchStats)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *JourneyStats) HasBranches() bool`

HasBranches returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


