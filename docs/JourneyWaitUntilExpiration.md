# JourneyWaitUntilExpiration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationSeconds** | Pointer to **NullableInt32** | Seconds to wait before the timer fires. Minimum 60, maximum 31556952 (1 year). | [optional] 
**Exits** | Pointer to **NullableBool** | When true, the user exits the journey when the timer fires; when false, the user continues to convergence. | [optional] 

## Methods

### NewJourneyWaitUntilExpiration

`func NewJourneyWaitUntilExpiration() *JourneyWaitUntilExpiration`

NewJourneyWaitUntilExpiration instantiates a new JourneyWaitUntilExpiration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyWaitUntilExpirationWithDefaults

`func NewJourneyWaitUntilExpirationWithDefaults() *JourneyWaitUntilExpiration`

NewJourneyWaitUntilExpirationWithDefaults instantiates a new JourneyWaitUntilExpiration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationSeconds

`func (o *JourneyWaitUntilExpiration) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *JourneyWaitUntilExpiration) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *JourneyWaitUntilExpiration) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *JourneyWaitUntilExpiration) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *JourneyWaitUntilExpiration) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *JourneyWaitUntilExpiration) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetExits

`func (o *JourneyWaitUntilExpiration) GetExits() bool`

GetExits returns the Exits field if non-nil, zero value otherwise.

### GetExitsOk

`func (o *JourneyWaitUntilExpiration) GetExitsOk() (*bool, bool)`

GetExitsOk returns a tuple with the Exits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExits

`func (o *JourneyWaitUntilExpiration) SetExits(v bool)`

SetExits sets Exits field to given value.

### HasExits

`func (o *JourneyWaitUntilExpiration) HasExits() bool`

HasExits returns a boolean if a field has been set.

### SetExitsNil

`func (o *JourneyWaitUntilExpiration) SetExitsNil(b bool)`

 SetExitsNil sets the value for Exits to be an explicit nil

### UnsetExits
`func (o *JourneyWaitUntilExpiration) UnsetExits()`

UnsetExits ensures that no value is present for Exits, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


