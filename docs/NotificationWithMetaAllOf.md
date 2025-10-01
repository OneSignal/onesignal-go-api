# NotificationWithMetaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remaining** | Pointer to **int32** | Number of notifications that have not been sent out yet. This can mean either our system is still processing the notification or you have delayed options set. | [optional] 
**Successful** | Pointer to **int32** | Number of notifications that were successfully delivered. | [optional] 
**Failed** | Pointer to **int32** | Number of notifications that could not be delivered due to those devices being unsubscribed. | [optional] 
**Errored** | Pointer to **int32** | Number of notifications that could not be delivered due to an error. You can find more information by viewing the notification in the dashboard. | [optional] 
**Converted** | Pointer to **int32** | Number of users who have clicked / tapped on your notification. | [optional] 
**QueuedAt** | Pointer to **int64** | Unix timestamp indicating when the notification was created. | [optional] 
**SendAfter** | Pointer to **NullableInt64** | Unix timestamp indicating when notification delivery should begin. | [optional] 
**CompletedAt** | Pointer to **NullableInt64** | Unix timestamp indicating when notification delivery completed. The delivery duration from start to finish can be calculated with completed_at - send_after. | [optional] 
**PlatformDeliveryStats** | Pointer to [**PlatformDeliveryData**](PlatformDeliveryData.md) |  | [optional] 
**Received** | Pointer to **NullableInt32** | Confirmed Deliveries number of devices that received the push notification. Paid Feature Only. Free accounts will see 0. | [optional] 
**ThrottleRatePerMinute** | Pointer to **NullableInt32** | number of push notifications sent per minute. Paid Feature Only. If throttling is not enabled for the app or the notification, and for free accounts, null is returned. Refer to Throttling for more details. | [optional] 
**Canceled** | Pointer to **bool** | Indicates whether the notification was canceled before it could be sent. | [optional] 

## Methods

### NewNotificationWithMetaAllOf

`func NewNotificationWithMetaAllOf() *NotificationWithMetaAllOf`

NewNotificationWithMetaAllOf instantiates a new NotificationWithMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithMetaAllOfWithDefaults

`func NewNotificationWithMetaAllOfWithDefaults() *NotificationWithMetaAllOf`

NewNotificationWithMetaAllOfWithDefaults instantiates a new NotificationWithMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemaining

`func (o *NotificationWithMetaAllOf) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *NotificationWithMetaAllOf) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *NotificationWithMetaAllOf) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *NotificationWithMetaAllOf) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetSuccessful

`func (o *NotificationWithMetaAllOf) GetSuccessful() int32`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *NotificationWithMetaAllOf) GetSuccessfulOk() (*int32, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *NotificationWithMetaAllOf) SetSuccessful(v int32)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *NotificationWithMetaAllOf) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### GetFailed

`func (o *NotificationWithMetaAllOf) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *NotificationWithMetaAllOf) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *NotificationWithMetaAllOf) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *NotificationWithMetaAllOf) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetErrored

`func (o *NotificationWithMetaAllOf) GetErrored() int32`

GetErrored returns the Errored field if non-nil, zero value otherwise.

### GetErroredOk

`func (o *NotificationWithMetaAllOf) GetErroredOk() (*int32, bool)`

GetErroredOk returns a tuple with the Errored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrored

`func (o *NotificationWithMetaAllOf) SetErrored(v int32)`

SetErrored sets Errored field to given value.

### HasErrored

`func (o *NotificationWithMetaAllOf) HasErrored() bool`

HasErrored returns a boolean if a field has been set.

### GetConverted

`func (o *NotificationWithMetaAllOf) GetConverted() int32`

GetConverted returns the Converted field if non-nil, zero value otherwise.

### GetConvertedOk

`func (o *NotificationWithMetaAllOf) GetConvertedOk() (*int32, bool)`

GetConvertedOk returns a tuple with the Converted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConverted

`func (o *NotificationWithMetaAllOf) SetConverted(v int32)`

SetConverted sets Converted field to given value.

### HasConverted

`func (o *NotificationWithMetaAllOf) HasConverted() bool`

HasConverted returns a boolean if a field has been set.

### GetQueuedAt

`func (o *NotificationWithMetaAllOf) GetQueuedAt() int64`

GetQueuedAt returns the QueuedAt field if non-nil, zero value otherwise.

### GetQueuedAtOk

`func (o *NotificationWithMetaAllOf) GetQueuedAtOk() (*int64, bool)`

GetQueuedAtOk returns a tuple with the QueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedAt

`func (o *NotificationWithMetaAllOf) SetQueuedAt(v int64)`

SetQueuedAt sets QueuedAt field to given value.

### HasQueuedAt

`func (o *NotificationWithMetaAllOf) HasQueuedAt() bool`

HasQueuedAt returns a boolean if a field has been set.

### GetSendAfter

`func (o *NotificationWithMetaAllOf) GetSendAfter() int64`

GetSendAfter returns the SendAfter field if non-nil, zero value otherwise.

### GetSendAfterOk

`func (o *NotificationWithMetaAllOf) GetSendAfterOk() (*int64, bool)`

GetSendAfterOk returns a tuple with the SendAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAfter

`func (o *NotificationWithMetaAllOf) SetSendAfter(v int64)`

SetSendAfter sets SendAfter field to given value.

### HasSendAfter

`func (o *NotificationWithMetaAllOf) HasSendAfter() bool`

HasSendAfter returns a boolean if a field has been set.

### SetSendAfterNil

`func (o *NotificationWithMetaAllOf) SetSendAfterNil(b bool)`

 SetSendAfterNil sets the value for SendAfter to be an explicit nil

### UnsetSendAfter
`func (o *NotificationWithMetaAllOf) UnsetSendAfter()`

UnsetSendAfter ensures that no value is present for SendAfter, not even an explicit nil
### GetCompletedAt

`func (o *NotificationWithMetaAllOf) GetCompletedAt() int64`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *NotificationWithMetaAllOf) GetCompletedAtOk() (*int64, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *NotificationWithMetaAllOf) SetCompletedAt(v int64)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *NotificationWithMetaAllOf) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *NotificationWithMetaAllOf) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *NotificationWithMetaAllOf) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetPlatformDeliveryStats

`func (o *NotificationWithMetaAllOf) GetPlatformDeliveryStats() PlatformDeliveryData`

GetPlatformDeliveryStats returns the PlatformDeliveryStats field if non-nil, zero value otherwise.

### GetPlatformDeliveryStatsOk

`func (o *NotificationWithMetaAllOf) GetPlatformDeliveryStatsOk() (*PlatformDeliveryData, bool)`

GetPlatformDeliveryStatsOk returns a tuple with the PlatformDeliveryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformDeliveryStats

`func (o *NotificationWithMetaAllOf) SetPlatformDeliveryStats(v PlatformDeliveryData)`

SetPlatformDeliveryStats sets PlatformDeliveryStats field to given value.

### HasPlatformDeliveryStats

`func (o *NotificationWithMetaAllOf) HasPlatformDeliveryStats() bool`

HasPlatformDeliveryStats returns a boolean if a field has been set.

### GetReceived

`func (o *NotificationWithMetaAllOf) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *NotificationWithMetaAllOf) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *NotificationWithMetaAllOf) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *NotificationWithMetaAllOf) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### SetReceivedNil

`func (o *NotificationWithMetaAllOf) SetReceivedNil(b bool)`

 SetReceivedNil sets the value for Received to be an explicit nil

### UnsetReceived
`func (o *NotificationWithMetaAllOf) UnsetReceived()`

UnsetReceived ensures that no value is present for Received, not even an explicit nil
### GetThrottleRatePerMinute

`func (o *NotificationWithMetaAllOf) GetThrottleRatePerMinute() int32`

GetThrottleRatePerMinute returns the ThrottleRatePerMinute field if non-nil, zero value otherwise.

### GetThrottleRatePerMinuteOk

`func (o *NotificationWithMetaAllOf) GetThrottleRatePerMinuteOk() (*int32, bool)`

GetThrottleRatePerMinuteOk returns a tuple with the ThrottleRatePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRatePerMinute

`func (o *NotificationWithMetaAllOf) SetThrottleRatePerMinute(v int32)`

SetThrottleRatePerMinute sets ThrottleRatePerMinute field to given value.

### HasThrottleRatePerMinute

`func (o *NotificationWithMetaAllOf) HasThrottleRatePerMinute() bool`

HasThrottleRatePerMinute returns a boolean if a field has been set.

### SetThrottleRatePerMinuteNil

`func (o *NotificationWithMetaAllOf) SetThrottleRatePerMinuteNil(b bool)`

 SetThrottleRatePerMinuteNil sets the value for ThrottleRatePerMinute to be an explicit nil

### UnsetThrottleRatePerMinute
`func (o *NotificationWithMetaAllOf) UnsetThrottleRatePerMinute()`

UnsetThrottleRatePerMinute ensures that no value is present for ThrottleRatePerMinute, not even an explicit nil
### GetCanceled

`func (o *NotificationWithMetaAllOf) GetCanceled() bool`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *NotificationWithMetaAllOf) GetCanceledOk() (*bool, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *NotificationWithMetaAllOf) SetCanceled(v bool)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *NotificationWithMetaAllOf) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


