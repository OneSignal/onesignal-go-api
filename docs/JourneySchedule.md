# JourneySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAt** | Pointer to **NullableString** | ISO 8601 start time. Use UTC (Z or +00:00). Must be at least 5 minutes in the future. | [optional] 
**StopAt** | Pointer to **NullableString** | ISO 8601 stop time. Use UTC (Z or +00:00). Must be in the future and later than start_at. | [optional] 
**Error** | Pointer to **NullableString** | Read-only. Present when a scheduling error occurred. | [optional] 

## Methods

### NewJourneySchedule

`func NewJourneySchedule() *JourneySchedule`

NewJourneySchedule instantiates a new JourneySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyScheduleWithDefaults

`func NewJourneyScheduleWithDefaults() *JourneySchedule`

NewJourneyScheduleWithDefaults instantiates a new JourneySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAt

`func (o *JourneySchedule) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *JourneySchedule) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *JourneySchedule) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *JourneySchedule) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAtNil

`func (o *JourneySchedule) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *JourneySchedule) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetStopAt

`func (o *JourneySchedule) GetStopAt() string`

GetStopAt returns the StopAt field if non-nil, zero value otherwise.

### GetStopAtOk

`func (o *JourneySchedule) GetStopAtOk() (*string, bool)`

GetStopAtOk returns a tuple with the StopAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopAt

`func (o *JourneySchedule) SetStopAt(v string)`

SetStopAt sets StopAt field to given value.

### HasStopAt

`func (o *JourneySchedule) HasStopAt() bool`

HasStopAt returns a boolean if a field has been set.

### SetStopAtNil

`func (o *JourneySchedule) SetStopAtNil(b bool)`

 SetStopAtNil sets the value for StopAt to be an explicit nil

### UnsetStopAt
`func (o *JourneySchedule) UnsetStopAt()`

UnsetStopAt ensures that no value is present for StopAt, not even an explicit nil
### GetError

`func (o *JourneySchedule) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *JourneySchedule) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *JourneySchedule) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *JourneySchedule) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *JourneySchedule) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *JourneySchedule) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


