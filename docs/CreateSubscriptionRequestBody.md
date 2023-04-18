# CreateSubscriptionRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | Pointer to [**SubscriptionObject**](SubscriptionObject.md) |  | [optional] 
**RetainPreviousOwner** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateSubscriptionRequestBody

`func NewCreateSubscriptionRequestBody() *CreateSubscriptionRequestBody`

NewCreateSubscriptionRequestBody instantiates a new CreateSubscriptionRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionRequestBodyWithDefaults

`func NewCreateSubscriptionRequestBodyWithDefaults() *CreateSubscriptionRequestBody`

NewCreateSubscriptionRequestBodyWithDefaults instantiates a new CreateSubscriptionRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *CreateSubscriptionRequestBody) GetSubscription() SubscriptionObject`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CreateSubscriptionRequestBody) GetSubscriptionOk() (*SubscriptionObject, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CreateSubscriptionRequestBody) SetSubscription(v SubscriptionObject)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *CreateSubscriptionRequestBody) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetRetainPreviousOwner

`func (o *CreateSubscriptionRequestBody) GetRetainPreviousOwner() bool`

GetRetainPreviousOwner returns the RetainPreviousOwner field if non-nil, zero value otherwise.

### GetRetainPreviousOwnerOk

`func (o *CreateSubscriptionRequestBody) GetRetainPreviousOwnerOk() (*bool, bool)`

GetRetainPreviousOwnerOk returns a tuple with the RetainPreviousOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousOwner

`func (o *CreateSubscriptionRequestBody) SetRetainPreviousOwner(v bool)`

SetRetainPreviousOwner sets RetainPreviousOwner field to given value.

### HasRetainPreviousOwner

`func (o *CreateSubscriptionRequestBody) HasRetainPreviousOwner() bool`

HasRetainPreviousOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


