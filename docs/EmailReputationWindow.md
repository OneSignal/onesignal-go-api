# EmailReputationWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BounceRate** | Pointer to **float64** | The fraction of successfully delivered emails that hard or soft bounced during the window. | [optional] 
**ComplaintRate** | Pointer to **float64** | The fraction of successfully delivered emails that recipients reported as spam during the window. | [optional] 

## Methods

### NewEmailReputationWindow

`func NewEmailReputationWindow() *EmailReputationWindow`

NewEmailReputationWindow instantiates a new EmailReputationWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailReputationWindowWithDefaults

`func NewEmailReputationWindowWithDefaults() *EmailReputationWindow`

NewEmailReputationWindowWithDefaults instantiates a new EmailReputationWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBounceRate

`func (o *EmailReputationWindow) GetBounceRate() float64`

GetBounceRate returns the BounceRate field if non-nil, zero value otherwise.

### GetBounceRateOk

`func (o *EmailReputationWindow) GetBounceRateOk() (*float64, bool)`

GetBounceRateOk returns a tuple with the BounceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceRate

`func (o *EmailReputationWindow) SetBounceRate(v float64)`

SetBounceRate sets BounceRate field to given value.

### HasBounceRate

`func (o *EmailReputationWindow) HasBounceRate() bool`

HasBounceRate returns a boolean if a field has been set.

### GetComplaintRate

`func (o *EmailReputationWindow) GetComplaintRate() float64`

GetComplaintRate returns the ComplaintRate field if non-nil, zero value otherwise.

### GetComplaintRateOk

`func (o *EmailReputationWindow) GetComplaintRateOk() (*float64, bool)`

GetComplaintRateOk returns a tuple with the ComplaintRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplaintRate

`func (o *EmailReputationWindow) SetComplaintRate(v float64)`

SetComplaintRate sets ComplaintRate field to given value.

### HasComplaintRate

`func (o *EmailReputationWindow) HasComplaintRate() bool`

HasComplaintRate returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


