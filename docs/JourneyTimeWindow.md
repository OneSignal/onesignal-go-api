# JourneyTimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to [**NullableJourneyTimePoint**](JourneyTimePoint.md) | When the window opens. | [optional] 
**End** | Pointer to [**NullableJourneyTimePoint**](JourneyTimePoint.md) | When the window closes. | [optional] 
**DayOfWeek** | Pointer to **int32** | Day of week, 1 &#x3D; Monday. Omit to apply the window to every day. | [optional] 

## Methods

### NewJourneyTimeWindow

`func NewJourneyTimeWindow() *JourneyTimeWindow`

NewJourneyTimeWindow instantiates a new JourneyTimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyTimeWindowWithDefaults

`func NewJourneyTimeWindowWithDefaults() *JourneyTimeWindow`

NewJourneyTimeWindowWithDefaults instantiates a new JourneyTimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *JourneyTimeWindow) GetStart() JourneyTimePoint`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *JourneyTimeWindow) GetStartOk() (*JourneyTimePoint, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *JourneyTimeWindow) SetStart(v JourneyTimePoint)`

SetStart sets Start field to given value.

### HasStart

`func (o *JourneyTimeWindow) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *JourneyTimeWindow) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *JourneyTimeWindow) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *JourneyTimeWindow) GetEnd() JourneyTimePoint`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *JourneyTimeWindow) GetEndOk() (*JourneyTimePoint, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *JourneyTimeWindow) SetEnd(v JourneyTimePoint)`

SetEnd sets End field to given value.

### HasEnd

`func (o *JourneyTimeWindow) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *JourneyTimeWindow) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *JourneyTimeWindow) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetDayOfWeek

`func (o *JourneyTimeWindow) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *JourneyTimeWindow) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *JourneyTimeWindow) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *JourneyTimeWindow) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


