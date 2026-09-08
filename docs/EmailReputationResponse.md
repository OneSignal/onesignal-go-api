# EmailReputationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Last24Hours** | Pointer to [**NullableEmailReputationWindow**](EmailReputationWindow.md) | Bounce and complaint rates for email delivered over the last 24 hours. | [optional] 
**Last7Days** | Pointer to [**NullableEmailReputationWindow**](EmailReputationWindow.md) | Bounce and complaint rates for email delivered over the last 7 days. | [optional] 
**Last30Days** | Pointer to [**NullableEmailReputationWindow**](EmailReputationWindow.md) | Bounce and complaint rates for email delivered over the last 30 days. | [optional] 

## Methods

### NewEmailReputationResponse

`func NewEmailReputationResponse() *EmailReputationResponse`

NewEmailReputationResponse instantiates a new EmailReputationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailReputationResponseWithDefaults

`func NewEmailReputationResponseWithDefaults() *EmailReputationResponse`

NewEmailReputationResponseWithDefaults instantiates a new EmailReputationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLast24Hours

`func (o *EmailReputationResponse) GetLast24Hours() EmailReputationWindow`

GetLast24Hours returns the Last24Hours field if non-nil, zero value otherwise.

### GetLast24HoursOk

`func (o *EmailReputationResponse) GetLast24HoursOk() (*EmailReputationWindow, bool)`

GetLast24HoursOk returns a tuple with the Last24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast24Hours

`func (o *EmailReputationResponse) SetLast24Hours(v EmailReputationWindow)`

SetLast24Hours sets Last24Hours field to given value.

### HasLast24Hours

`func (o *EmailReputationResponse) HasLast24Hours() bool`

HasLast24Hours returns a boolean if a field has been set.

### SetLast24HoursNil

`func (o *EmailReputationResponse) SetLast24HoursNil(b bool)`

 SetLast24HoursNil sets the value for Last24Hours to be an explicit nil

### UnsetLast24Hours
`func (o *EmailReputationResponse) UnsetLast24Hours()`

UnsetLast24Hours ensures that no value is present for Last24Hours, not even an explicit nil
### GetLast7Days

`func (o *EmailReputationResponse) GetLast7Days() EmailReputationWindow`

GetLast7Days returns the Last7Days field if non-nil, zero value otherwise.

### GetLast7DaysOk

`func (o *EmailReputationResponse) GetLast7DaysOk() (*EmailReputationWindow, bool)`

GetLast7DaysOk returns a tuple with the Last7Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast7Days

`func (o *EmailReputationResponse) SetLast7Days(v EmailReputationWindow)`

SetLast7Days sets Last7Days field to given value.

### HasLast7Days

`func (o *EmailReputationResponse) HasLast7Days() bool`

HasLast7Days returns a boolean if a field has been set.

### SetLast7DaysNil

`func (o *EmailReputationResponse) SetLast7DaysNil(b bool)`

 SetLast7DaysNil sets the value for Last7Days to be an explicit nil

### UnsetLast7Days
`func (o *EmailReputationResponse) UnsetLast7Days()`

UnsetLast7Days ensures that no value is present for Last7Days, not even an explicit nil
### GetLast30Days

`func (o *EmailReputationResponse) GetLast30Days() EmailReputationWindow`

GetLast30Days returns the Last30Days field if non-nil, zero value otherwise.

### GetLast30DaysOk

`func (o *EmailReputationResponse) GetLast30DaysOk() (*EmailReputationWindow, bool)`

GetLast30DaysOk returns a tuple with the Last30Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast30Days

`func (o *EmailReputationResponse) SetLast30Days(v EmailReputationWindow)`

SetLast30Days sets Last30Days field to given value.

### HasLast30Days

`func (o *EmailReputationResponse) HasLast30Days() bool`

HasLast30Days returns a boolean if a field has been set.

### SetLast30DaysNil

`func (o *EmailReputationResponse) SetLast30DaysNil(b bool)`

 SetLast30DaysNil sets the value for Last30Days to be an explicit nil

### UnsetLast30Days
`func (o *EmailReputationResponse) UnsetLast30Days()`

UnsetLast30Days ensures that no value is present for Last30Days, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


