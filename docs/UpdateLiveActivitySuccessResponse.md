# UpdateLiveActivitySuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationId** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**Notification200Errors**](Notification200Errors.md) |  | [optional] 

## Methods

### NewUpdateLiveActivitySuccessResponse

`func NewUpdateLiveActivitySuccessResponse() *UpdateLiveActivitySuccessResponse`

NewUpdateLiveActivitySuccessResponse instantiates a new UpdateLiveActivitySuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLiveActivitySuccessResponseWithDefaults

`func NewUpdateLiveActivitySuccessResponseWithDefaults() *UpdateLiveActivitySuccessResponse`

NewUpdateLiveActivitySuccessResponseWithDefaults instantiates a new UpdateLiveActivitySuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationId

`func (o *UpdateLiveActivitySuccessResponse) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *UpdateLiveActivitySuccessResponse) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *UpdateLiveActivitySuccessResponse) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *UpdateLiveActivitySuccessResponse) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetErrors

`func (o *UpdateLiveActivitySuccessResponse) GetErrors() Notification200Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UpdateLiveActivitySuccessResponse) GetErrorsOk() (*Notification200Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UpdateLiveActivitySuccessResponse) SetErrors(v Notification200Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UpdateLiveActivitySuccessResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


