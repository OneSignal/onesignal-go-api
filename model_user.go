/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.2.2
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// User struct for User
type User struct {
	Properties *PropertiesObject `json:"properties,omitempty"`
	Identity map[string]interface{} `json:"identity,omitempty"`
	Subscriptions []SubscriptionObject `json:"subscriptions,omitempty"`
	SubscriptionOptions *UserSubscriptionOptions `json:"subscription_options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *User) GetProperties() PropertiesObject {
	if o == nil || o.Properties == nil {
		var ret PropertiesObject
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPropertiesOk() (*PropertiesObject, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *User) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given PropertiesObject and assigns it to the Properties field.
func (o *User) SetProperties(v PropertiesObject) {
	o.Properties = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *User) GetIdentity() map[string]interface{} {
	if o == nil || o.Identity == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdentityOk() (map[string]interface{}, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *User) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given map[string]interface{} and assigns it to the Identity field.
func (o *User) SetIdentity(v map[string]interface{}) {
	o.Identity = v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *User) GetSubscriptions() []SubscriptionObject {
	if o == nil || o.Subscriptions == nil {
		var ret []SubscriptionObject
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSubscriptionsOk() ([]SubscriptionObject, bool) {
	if o == nil || o.Subscriptions == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *User) HasSubscriptions() bool {
	if o != nil && o.Subscriptions != nil {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []SubscriptionObject and assigns it to the Subscriptions field.
func (o *User) SetSubscriptions(v []SubscriptionObject) {
	o.Subscriptions = v
}

// GetSubscriptionOptions returns the SubscriptionOptions field value if set, zero value otherwise.
func (o *User) GetSubscriptionOptions() UserSubscriptionOptions {
	if o == nil || o.SubscriptionOptions == nil {
		var ret UserSubscriptionOptions
		return ret
	}
	return *o.SubscriptionOptions
}

// GetSubscriptionOptionsOk returns a tuple with the SubscriptionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSubscriptionOptionsOk() (*UserSubscriptionOptions, bool) {
	if o == nil || o.SubscriptionOptions == nil {
		return nil, false
	}
	return o.SubscriptionOptions, true
}

// HasSubscriptionOptions returns a boolean if a field has been set.
func (o *User) HasSubscriptionOptions() bool {
	if o != nil && o.SubscriptionOptions != nil {
		return true
	}

	return false
}

// SetSubscriptionOptions gets a reference to the given UserSubscriptionOptions and assigns it to the SubscriptionOptions field.
func (o *User) SetSubscriptionOptions(v UserSubscriptionOptions) {
	o.SubscriptionOptions = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.SubscriptionOptions != nil {
		toSerialize["subscription_options"] = o.SubscriptionOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *User) UnmarshalJSON(bytes []byte) (err error) {
	varUser := _User{}

	if err = json.Unmarshal(bytes, &varUser); err == nil {
		*o = User(varUser)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "properties")
		delete(additionalProperties, "identity")
		delete(additionalProperties, "subscriptions")
		delete(additionalProperties, "subscription_options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


