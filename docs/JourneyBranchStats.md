# JourneyBranchStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **int32** | Users who took this branch. | [optional] 

## Methods

### NewJourneyBranchStats

`func NewJourneyBranchStats() *JourneyBranchStats`

NewJourneyBranchStats instantiates a new JourneyBranchStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyBranchStatsWithDefaults

`func NewJourneyBranchStatsWithDefaults() *JourneyBranchStats`

NewJourneyBranchStatsWithDefaults instantiates a new JourneyBranchStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *JourneyBranchStats) GetCompleted() int32`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *JourneyBranchStats) GetCompletedOk() (*int32, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *JourneyBranchStats) SetCompleted(v int32)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *JourneyBranchStats) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


