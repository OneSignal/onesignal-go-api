# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to [**PropertiesObject**](PropertiesObject.md) |  | [optional] 
**Identity** | Pointer to **map[string]interface{}** |  | [optional] 
**Subscriptions** | Pointer to [**[]SubscriptionObject**](SubscriptionObject.md) |  | [optional] 
**SubscriptionOptions** | Pointer to [**UserSubscriptionOptions**](UserSubscriptionOptions.md) |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *User) GetProperties() PropertiesObject`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *User) GetPropertiesOk() (*PropertiesObject, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *User) SetProperties(v PropertiesObject)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *User) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetIdentity

`func (o *User) GetIdentity() map[string]interface{}`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *User) GetIdentityOk() (*map[string]interface{}, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *User) SetIdentity(v map[string]interface{})`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *User) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetSubscriptions

`func (o *User) GetSubscriptions() []SubscriptionObject`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *User) GetSubscriptionsOk() (*[]SubscriptionObject, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *User) SetSubscriptions(v []SubscriptionObject)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *User) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetSubscriptionOptions

`func (o *User) GetSubscriptionOptions() UserSubscriptionOptions`

GetSubscriptionOptions returns the SubscriptionOptions field if non-nil, zero value otherwise.

### GetSubscriptionOptionsOk

`func (o *User) GetSubscriptionOptionsOk() (*UserSubscriptionOptions, bool)`

GetSubscriptionOptionsOk returns a tuple with the SubscriptionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOptions

`func (o *User) SetSubscriptionOptions(v UserSubscriptionOptions)`

SetSubscriptionOptions sets SubscriptionOptions field to given value.

### HasSubscriptionOptions

`func (o *User) HasSubscriptionOptions() bool`

HasSubscriptionOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


