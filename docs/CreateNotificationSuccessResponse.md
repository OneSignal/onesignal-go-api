# CreateNotificationSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Recipients** | **int32** | Estimated number of subscribers targetted by notification. | 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Errors** | Pointer to [**Notification200Errors**](Notification200Errors.md) |  | [optional] 

## Methods

### NewCreateNotificationSuccessResponse

`func NewCreateNotificationSuccessResponse(id string, recipients int32, ) *CreateNotificationSuccessResponse`

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


### GetRecipients

`func (o *CreateNotificationSuccessResponse) GetRecipients() int32`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateNotificationSuccessResponse) GetRecipientsOk() (*int32, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateNotificationSuccessResponse) SetRecipients(v int32)`

SetRecipients sets Recipients field to given value.


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

`func (o *CreateNotificationSuccessResponse) GetErrors() Notification200Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateNotificationSuccessResponse) GetErrorsOk() (*Notification200Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateNotificationSuccessResponse) SetErrors(v Notification200Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateNotificationSuccessResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


