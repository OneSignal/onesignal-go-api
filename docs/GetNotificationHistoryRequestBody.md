# GetNotificationHistoryRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to **string** | -&gt; \&quot;sent\&quot; - All the devices by player_id that were sent the specified notification_id.  Notifications targeting under 1000 recipients will not have \&quot;sent\&quot; events recorded, but will show \&quot;clicked\&quot; events. \&quot;clicked\&quot; - All the devices by &#x60;player_id&#x60; that clicked the specified notification_id. | [optional] 
**Email** | Pointer to **string** | The email address you would like the report sent. | [optional] 
**AppId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetNotificationHistoryRequestBody

`func NewGetNotificationHistoryRequestBody() *GetNotificationHistoryRequestBody`

NewGetNotificationHistoryRequestBody instantiates a new GetNotificationHistoryRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationHistoryRequestBodyWithDefaults

`func NewGetNotificationHistoryRequestBodyWithDefaults() *GetNotificationHistoryRequestBody`

NewGetNotificationHistoryRequestBodyWithDefaults instantiates a new GetNotificationHistoryRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *GetNotificationHistoryRequestBody) GetEvents() string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetNotificationHistoryRequestBody) GetEventsOk() (*string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetNotificationHistoryRequestBody) SetEvents(v string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetNotificationHistoryRequestBody) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetEmail

`func (o *GetNotificationHistoryRequestBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetNotificationHistoryRequestBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetNotificationHistoryRequestBody) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetNotificationHistoryRequestBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAppId

`func (o *GetNotificationHistoryRequestBody) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GetNotificationHistoryRequestBody) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GetNotificationHistoryRequestBody) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *GetNotificationHistoryRequestBody) HasAppId() bool`

HasAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


