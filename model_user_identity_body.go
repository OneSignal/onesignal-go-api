/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 5.2.0
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// UserIdentityBody struct for UserIdentityBody
type UserIdentityBody struct {
	Identity *map[string]string `json:"identity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentityBody UserIdentityBody

// NewUserIdentityBody instantiates a new UserIdentityBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentityBody() *UserIdentityBody {
	this := UserIdentityBody{}
	return &this
}

// NewUserIdentityBodyWithDefaults instantiates a new UserIdentityBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentityBodyWithDefaults() *UserIdentityBody {
	this := UserIdentityBody{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *UserIdentityBody) GetIdentity() map[string]string {
	if o == nil || o.Identity == nil {
		var ret map[string]string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentityBody) GetIdentityOk() (*map[string]string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *UserIdentityBody) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given map[string]string and assigns it to the Identity field.
func (o *UserIdentityBody) SetIdentity(v map[string]string) {
	o.Identity = &v
}

func (o UserIdentityBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserIdentityBody) UnmarshalJSON(bytes []byte) (err error) {
	varUserIdentityBody := _UserIdentityBody{}

	if err = json.Unmarshal(bytes, &varUserIdentityBody); err == nil {
		*o = UserIdentityBody(varUserIdentityBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserIdentityBody struct {
	value *UserIdentityBody
	isSet bool
}

func (v NullableUserIdentityBody) Get() *UserIdentityBody {
	return v.value
}

func (v *NullableUserIdentityBody) Set(val *UserIdentityBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentityBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentityBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentityBody(val *UserIdentityBody) *NullableUserIdentityBody {
	return &NullableUserIdentityBody{value: val, isSet: true}
}

func (v NullableUserIdentityBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentityBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


