# JourneyReentryRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationSeconds** | Pointer to **NullableInt32** | Minimum seconds before a user can re-enter. Must be at least 600 (10 minutes). | [optional] 

## Methods

### NewJourneyReentryRules

`func NewJourneyReentryRules() *JourneyReentryRules`

NewJourneyReentryRules instantiates a new JourneyReentryRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyReentryRulesWithDefaults

`func NewJourneyReentryRulesWithDefaults() *JourneyReentryRules`

NewJourneyReentryRulesWithDefaults instantiates a new JourneyReentryRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationSeconds

`func (o *JourneyReentryRules) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *JourneyReentryRules) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *JourneyReentryRules) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *JourneyReentryRules) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *JourneyReentryRules) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *JourneyReentryRules) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


