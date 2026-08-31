# EstimateNotificationRecipientsSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The estimated audience size based on the user targeting method you&#39;ve set on the message, and the specific platforms the message is targeted to send to. | [optional] 
**UncappedCount** | Pointer to **NullableInt32** | The estimated audience size before the plan&#39;s web push subscriber cap is applied. Present only when &#x60;cap_applied&#x60; is &#x60;true&#x60;; &#x60;null&#x60; otherwise. | [optional] 
**CapApplied** | Pointer to **bool** | Whether &#x60;count&#x60; was reduced because the app is on a plan that caps the number of web push subscribers it can send to. | [optional] 
**MobileSuppressed** | Pointer to **bool** | The mobile equivalent of &#x60;cap_applied&#x60;. Whether mobile push deliveries will be dropped for this send because the org is over its plan&#39;s mobile push subscriber cap. &#x60;false&#x60; when the notification doesn&#39;t target any mobile push platforms. | [optional] 
**MobileExcludedCount** | Pointer to **int32** | How many mobile push recipients the &#x60;count&#x60; excludes due to the plan&#39;s mobile push subscriber cap. &#x60;0&#x60; when &#x60;mobile_suppressed&#x60; is &#x60;false&#x60;. | [optional] 

## Methods

### NewEstimateNotificationRecipientsSuccessResponse

`func NewEstimateNotificationRecipientsSuccessResponse() *EstimateNotificationRecipientsSuccessResponse`

NewEstimateNotificationRecipientsSuccessResponse instantiates a new EstimateNotificationRecipientsSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateNotificationRecipientsSuccessResponseWithDefaults

`func NewEstimateNotificationRecipientsSuccessResponseWithDefaults() *EstimateNotificationRecipientsSuccessResponse`

NewEstimateNotificationRecipientsSuccessResponseWithDefaults instantiates a new EstimateNotificationRecipientsSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *EstimateNotificationRecipientsSuccessResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EstimateNotificationRecipientsSuccessResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EstimateNotificationRecipientsSuccessResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *EstimateNotificationRecipientsSuccessResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetUncappedCount

`func (o *EstimateNotificationRecipientsSuccessResponse) GetUncappedCount() int32`

GetUncappedCount returns the UncappedCount field if non-nil, zero value otherwise.

### GetUncappedCountOk

`func (o *EstimateNotificationRecipientsSuccessResponse) GetUncappedCountOk() (*int32, bool)`

GetUncappedCountOk returns a tuple with the UncappedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncappedCount

`func (o *EstimateNotificationRecipientsSuccessResponse) SetUncappedCount(v int32)`

SetUncappedCount sets UncappedCount field to given value.

### HasUncappedCount

`func (o *EstimateNotificationRecipientsSuccessResponse) HasUncappedCount() bool`

HasUncappedCount returns a boolean if a field has been set.

### SetUncappedCountNil

`func (o *EstimateNotificationRecipientsSuccessResponse) SetUncappedCountNil(b bool)`

 SetUncappedCountNil sets the value for UncappedCount to be an explicit nil

### UnsetUncappedCount
`func (o *EstimateNotificationRecipientsSuccessResponse) UnsetUncappedCount()`

UnsetUncappedCount ensures that no value is present for UncappedCount, not even an explicit nil
### GetCapApplied

`func (o *EstimateNotificationRecipientsSuccessResponse) GetCapApplied() bool`

GetCapApplied returns the CapApplied field if non-nil, zero value otherwise.

### GetCapAppliedOk

`func (o *EstimateNotificationRecipientsSuccessResponse) GetCapAppliedOk() (*bool, bool)`

GetCapAppliedOk returns a tuple with the CapApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapApplied

`func (o *EstimateNotificationRecipientsSuccessResponse) SetCapApplied(v bool)`

SetCapApplied sets CapApplied field to given value.

### HasCapApplied

`func (o *EstimateNotificationRecipientsSuccessResponse) HasCapApplied() bool`

HasCapApplied returns a boolean if a field has been set.

### GetMobileSuppressed

`func (o *EstimateNotificationRecipientsSuccessResponse) GetMobileSuppressed() bool`

GetMobileSuppressed returns the MobileSuppressed field if non-nil, zero value otherwise.

### GetMobileSuppressedOk

`func (o *EstimateNotificationRecipientsSuccessResponse) GetMobileSuppressedOk() (*bool, bool)`

GetMobileSuppressedOk returns a tuple with the MobileSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileSuppressed

`func (o *EstimateNotificationRecipientsSuccessResponse) SetMobileSuppressed(v bool)`

SetMobileSuppressed sets MobileSuppressed field to given value.

### HasMobileSuppressed

`func (o *EstimateNotificationRecipientsSuccessResponse) HasMobileSuppressed() bool`

HasMobileSuppressed returns a boolean if a field has been set.

### GetMobileExcludedCount

`func (o *EstimateNotificationRecipientsSuccessResponse) GetMobileExcludedCount() int32`

GetMobileExcludedCount returns the MobileExcludedCount field if non-nil, zero value otherwise.

### GetMobileExcludedCountOk

`func (o *EstimateNotificationRecipientsSuccessResponse) GetMobileExcludedCountOk() (*int32, bool)`

GetMobileExcludedCountOk returns a tuple with the MobileExcludedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileExcludedCount

`func (o *EstimateNotificationRecipientsSuccessResponse) SetMobileExcludedCount(v int32)`

SetMobileExcludedCount sets MobileExcludedCount field to given value.

### HasMobileExcludedCount

`func (o *EstimateNotificationRecipientsSuccessResponse) HasMobileExcludedCount() bool`

HasMobileExcludedCount returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


