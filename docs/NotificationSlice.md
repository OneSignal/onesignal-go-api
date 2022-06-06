# NotificationSlice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Notifications** | Pointer to [**[]NotificationWithMeta**](NotificationWithMeta.md) |  | [optional] 

## Methods

### NewNotificationSlice

`func NewNotificationSlice() *NotificationSlice`

NewNotificationSlice instantiates a new NotificationSlice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSliceWithDefaults

`func NewNotificationSliceWithDefaults() *NotificationSlice`

NewNotificationSliceWithDefaults instantiates a new NotificationSlice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *NotificationSlice) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *NotificationSlice) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *NotificationSlice) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *NotificationSlice) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetOffset

`func (o *NotificationSlice) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *NotificationSlice) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *NotificationSlice) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *NotificationSlice) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *NotificationSlice) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NotificationSlice) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NotificationSlice) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *NotificationSlice) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNotifications

`func (o *NotificationSlice) GetNotifications() []NotificationWithMeta`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *NotificationSlice) GetNotificationsOk() (*[]NotificationWithMeta, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *NotificationSlice) SetNotifications(v []NotificationWithMeta)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *NotificationSlice) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


