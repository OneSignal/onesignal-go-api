# JourneyTimePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to **int32** | Hour of day, 0-23. | [optional] 
**Minute** | Pointer to **int32** | Minute of hour, 0-59. Defaults to 0. | [optional] 

## Methods

### NewJourneyTimePoint

`func NewJourneyTimePoint() *JourneyTimePoint`

NewJourneyTimePoint instantiates a new JourneyTimePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyTimePointWithDefaults

`func NewJourneyTimePointWithDefaults() *JourneyTimePoint`

NewJourneyTimePointWithDefaults instantiates a new JourneyTimePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *JourneyTimePoint) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *JourneyTimePoint) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *JourneyTimePoint) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *JourneyTimePoint) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *JourneyTimePoint) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *JourneyTimePoint) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *JourneyTimePoint) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *JourneyTimePoint) HasMinute() bool`

HasMinute returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


