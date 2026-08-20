# JourneyEarlyExitRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnSegment** | Pointer to [**NullableJourneyEarlyExitRulesOnSegment**](JourneyEarlyExitRulesOnSegment.md) |  | [optional] 
**WhenNotInAudience** | Pointer to **NullableBool** | Exit when the user no longer matches the journey audience. Defaults to false. | [optional] 
**OnSession** | Pointer to **NullableBool** | Exit on a new session start. Defaults to false. | [optional] 
**OnEvent** | Pointer to [**NullableJourneyEarlyExitRulesOnEvent**](JourneyEarlyExitRulesOnEvent.md) |  | [optional] 

## Methods

### NewJourneyEarlyExitRules

`func NewJourneyEarlyExitRules() *JourneyEarlyExitRules`

NewJourneyEarlyExitRules instantiates a new JourneyEarlyExitRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyEarlyExitRulesWithDefaults

`func NewJourneyEarlyExitRulesWithDefaults() *JourneyEarlyExitRules`

NewJourneyEarlyExitRulesWithDefaults instantiates a new JourneyEarlyExitRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnSegment

`func (o *JourneyEarlyExitRules) GetOnSegment() JourneyEarlyExitRulesOnSegment`

GetOnSegment returns the OnSegment field if non-nil, zero value otherwise.

### GetOnSegmentOk

`func (o *JourneyEarlyExitRules) GetOnSegmentOk() (*JourneyEarlyExitRulesOnSegment, bool)`

GetOnSegmentOk returns a tuple with the OnSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSegment

`func (o *JourneyEarlyExitRules) SetOnSegment(v JourneyEarlyExitRulesOnSegment)`

SetOnSegment sets OnSegment field to given value.

### HasOnSegment

`func (o *JourneyEarlyExitRules) HasOnSegment() bool`

HasOnSegment returns a boolean if a field has been set.

### SetOnSegmentNil

`func (o *JourneyEarlyExitRules) SetOnSegmentNil(b bool)`

 SetOnSegmentNil sets the value for OnSegment to be an explicit nil

### UnsetOnSegment
`func (o *JourneyEarlyExitRules) UnsetOnSegment()`

UnsetOnSegment ensures that no value is present for OnSegment, not even an explicit nil
### GetWhenNotInAudience

`func (o *JourneyEarlyExitRules) GetWhenNotInAudience() bool`

GetWhenNotInAudience returns the WhenNotInAudience field if non-nil, zero value otherwise.

### GetWhenNotInAudienceOk

`func (o *JourneyEarlyExitRules) GetWhenNotInAudienceOk() (*bool, bool)`

GetWhenNotInAudienceOk returns a tuple with the WhenNotInAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenNotInAudience

`func (o *JourneyEarlyExitRules) SetWhenNotInAudience(v bool)`

SetWhenNotInAudience sets WhenNotInAudience field to given value.

### HasWhenNotInAudience

`func (o *JourneyEarlyExitRules) HasWhenNotInAudience() bool`

HasWhenNotInAudience returns a boolean if a field has been set.

### SetWhenNotInAudienceNil

`func (o *JourneyEarlyExitRules) SetWhenNotInAudienceNil(b bool)`

 SetWhenNotInAudienceNil sets the value for WhenNotInAudience to be an explicit nil

### UnsetWhenNotInAudience
`func (o *JourneyEarlyExitRules) UnsetWhenNotInAudience()`

UnsetWhenNotInAudience ensures that no value is present for WhenNotInAudience, not even an explicit nil
### GetOnSession

`func (o *JourneyEarlyExitRules) GetOnSession() bool`

GetOnSession returns the OnSession field if non-nil, zero value otherwise.

### GetOnSessionOk

`func (o *JourneyEarlyExitRules) GetOnSessionOk() (*bool, bool)`

GetOnSessionOk returns a tuple with the OnSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSession

`func (o *JourneyEarlyExitRules) SetOnSession(v bool)`

SetOnSession sets OnSession field to given value.

### HasOnSession

`func (o *JourneyEarlyExitRules) HasOnSession() bool`

HasOnSession returns a boolean if a field has been set.

### SetOnSessionNil

`func (o *JourneyEarlyExitRules) SetOnSessionNil(b bool)`

 SetOnSessionNil sets the value for OnSession to be an explicit nil

### UnsetOnSession
`func (o *JourneyEarlyExitRules) UnsetOnSession()`

UnsetOnSession ensures that no value is present for OnSession, not even an explicit nil
### GetOnEvent

`func (o *JourneyEarlyExitRules) GetOnEvent() JourneyEarlyExitRulesOnEvent`

GetOnEvent returns the OnEvent field if non-nil, zero value otherwise.

### GetOnEventOk

`func (o *JourneyEarlyExitRules) GetOnEventOk() (*JourneyEarlyExitRulesOnEvent, bool)`

GetOnEventOk returns a tuple with the OnEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnEvent

`func (o *JourneyEarlyExitRules) SetOnEvent(v JourneyEarlyExitRulesOnEvent)`

SetOnEvent sets OnEvent field to given value.

### HasOnEvent

`func (o *JourneyEarlyExitRules) HasOnEvent() bool`

HasOnEvent returns a boolean if a field has been set.

### SetOnEventNil

`func (o *JourneyEarlyExitRules) SetOnEventNil(b bool)`

 SetOnEventNil sets the value for OnEvent to be an explicit nil

### UnsetOnEvent
`func (o *JourneyEarlyExitRules) UnsetOnEvent()`

UnsetOnEvent ensures that no value is present for OnEvent, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


