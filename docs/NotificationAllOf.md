# NotificationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendAfter** | Pointer to **NullableTime** | Channel: All Schedule notification for future delivery. API defaults to UTC -1100 Examples: All examples are the exact same date &amp; time. \&quot;Thu Sep 24 2015 14:00:00 GMT-0700 (PDT)\&quot; \&quot;September 24th 2015, 2:00:00 pm UTC-07:00\&quot; \&quot;2015-09-24 14:00:00 GMT-0700\&quot; \&quot;Sept 24 2015 14:00:00 GMT-0700\&quot; \&quot;Thu Sep 24 2015 14:00:00 GMT-0700 (Pacific Daylight Time)\&quot; Note: SMS currently only supports send_after parameter.  | [optional] 

## Methods

### NewNotificationAllOf

`func NewNotificationAllOf() *NotificationAllOf`

NewNotificationAllOf instantiates a new NotificationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationAllOfWithDefaults

`func NewNotificationAllOfWithDefaults() *NotificationAllOf`

NewNotificationAllOfWithDefaults instantiates a new NotificationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendAfter

`func (o *NotificationAllOf) GetSendAfter() time.Time`

GetSendAfter returns the SendAfter field if non-nil, zero value otherwise.

### GetSendAfterOk

`func (o *NotificationAllOf) GetSendAfterOk() (*time.Time, bool)`

GetSendAfterOk returns a tuple with the SendAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAfter

`func (o *NotificationAllOf) SetSendAfter(v time.Time)`

SetSendAfter sets SendAfter field to given value.

### HasSendAfter

`func (o *NotificationAllOf) HasSendAfter() bool`

HasSendAfter returns a boolean if a field has been set.

### SetSendAfterNil

`func (o *NotificationAllOf) SetSendAfterNil(b bool)`

 SetSendAfterNil sets the value for SendAfter to be an explicit nil

### UnsetSendAfter
`func (o *NotificationAllOf) UnsetSendAfter()`

UnsetSendAfter ensures that no value is present for SendAfter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


