# GetNotificationRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to **string** | -&gt; \&quot;sent\&quot; - All the devices by player_id that were sent the specified notification_id.  Notifications targeting under 1000 recipients will not have \&quot;sent\&quot; events recorded, but will show \&quot;clicked\&quot; events. \&quot;clicked\&quot; - All the devices by &#x60;player_id&#x60; that clicked the specified notification_id. | [optional] 
**Email** | Pointer to **string** | The email address you would like the report sent. | [optional] 
**AppId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetNotificationRequestBody

`func NewGetNotificationRequestBody() *GetNotificationRequestBody`

NewGetNotificationRequestBody instantiates a new GetNotificationRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationRequestBodyWithDefaults

`func NewGetNotificationRequestBodyWithDefaults() *GetNotificationRequestBody`

NewGetNotificationRequestBodyWithDefaults instantiates a new GetNotificationRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *GetNotificationRequestBody) GetEvents() string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetNotificationRequestBody) GetEventsOk() (*string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetNotificationRequestBody) SetEvents(v string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetNotificationRequestBody) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetEmail

`func (o *GetNotificationRequestBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetNotificationRequestBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetNotificationRequestBody) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetNotificationRequestBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAppId

`func (o *GetNotificationRequestBody) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GetNotificationRequestBody) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GetNotificationRequestBody) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *GetNotificationRequestBody) HasAppId() bool`

HasAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


