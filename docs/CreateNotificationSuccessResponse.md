# CreateNotificationSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Notification identifier when the request created a notification. An empty string means no notification was created; read &#x60;errors&#x60; for details (HTTP may still be 200). | [optional] 
**ExternalId** | Pointer to **NullableString** | Optional correlation / idempotency-related value from the API response. This is not the end-user External ID used for targeting recipients (that lives under &#x60;include_aliases.external_id&#x60;). | [optional] 
**Errors** | Pointer to **interface{}** | Polymorphic field: may be an array of human-readable strings and/or an object (for example with &#x60;invalid_aliases&#x60;, &#x60;invalid_external_user_ids&#x60;, or &#x60;invalid_player_ids&#x60;) depending on the API response; HTTP may still be 200 with partial success. Typed SDKs model this loosely so both shapes deserialize. | [optional] 

## Methods

### NewCreateNotificationSuccessResponse

`func NewCreateNotificationSuccessResponse() *CreateNotificationSuccessResponse`

NewCreateNotificationSuccessResponse instantiates a new CreateNotificationSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationSuccessResponseWithDefaults

`func NewCreateNotificationSuccessResponseWithDefaults() *CreateNotificationSuccessResponse`

NewCreateNotificationSuccessResponseWithDefaults instantiates a new CreateNotificationSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNotificationSuccessResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNotificationSuccessResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNotificationSuccessResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNotificationSuccessResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateNotificationSuccessResponse) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateNotificationSuccessResponse) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateNotificationSuccessResponse) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateNotificationSuccessResponse) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CreateNotificationSuccessResponse) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CreateNotificationSuccessResponse) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetErrors

`func (o *CreateNotificationSuccessResponse) GetErrors() interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateNotificationSuccessResponse) GetErrorsOk() (*interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateNotificationSuccessResponse) SetErrors(v interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateNotificationSuccessResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *CreateNotificationSuccessResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *CreateNotificationSuccessResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


