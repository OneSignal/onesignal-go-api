/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.4.0
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// UserIdentityResponse struct for UserIdentityResponse
type UserIdentityResponse struct {
	Identity map[string]interface{} `json:"identity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentityResponse UserIdentityResponse

// NewUserIdentityResponse instantiates a new UserIdentityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentityResponse() *UserIdentityResponse {
	this := UserIdentityResponse{}
	return &this
}

// NewUserIdentityResponseWithDefaults instantiates a new UserIdentityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentityResponseWithDefaults() *UserIdentityResponse {
	this := UserIdentityResponse{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *UserIdentityResponse) GetIdentity() map[string]interface{} {
	if o == nil || o.Identity == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentityResponse) GetIdentityOk() (map[string]interface{}, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *UserIdentityResponse) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given map[string]interface{} and assigns it to the Identity field.
func (o *UserIdentityResponse) SetIdentity(v map[string]interface{}) {
	o.Identity = v
}

func (o UserIdentityResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserIdentityResponse) UnmarshalJSON(bytes []byte) (err error) {
	varUserIdentityResponse := _UserIdentityResponse{}

	if err = json.Unmarshal(bytes, &varUserIdentityResponse); err == nil {
		*o = UserIdentityResponse(varUserIdentityResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserIdentityResponse struct {
	value *UserIdentityResponse
	isSet bool
}

func (v NullableUserIdentityResponse) Get() *UserIdentityResponse {
	return v.value
}

func (v *NullableUserIdentityResponse) Set(val *UserIdentityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentityResponse(val *UserIdentityResponse) *NullableUserIdentityResponse {
	return &NullableUserIdentityResponse{value: val, isSet: true}
}

func (v NullableUserIdentityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


