/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.0.1
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// InlineResponse2007 struct for InlineResponse2007
type InlineResponse2007 struct {
	Success *string `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineResponse2007 InlineResponse2007

// NewInlineResponse2007 instantiates a new InlineResponse2007 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2007() *InlineResponse2007 {
	this := InlineResponse2007{}
	return &this
}

// NewInlineResponse2007WithDefaults instantiates a new InlineResponse2007 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2007WithDefaults() *InlineResponse2007 {
	this := InlineResponse2007{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InlineResponse2007) GetSuccess() string {
	if o == nil || o.Success == nil {
		var ret string
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetSuccessOk() (*string, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InlineResponse2007) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given string and assigns it to the Success field.
func (o *InlineResponse2007) SetSuccess(v string) {
	o.Success = &v
}

func (o InlineResponse2007) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineResponse2007) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2007 := _InlineResponse2007{}

	if err = json.Unmarshal(bytes, &varInlineResponse2007); err == nil {
		*o = InlineResponse2007(varInlineResponse2007)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineResponse2007 struct {
	value *InlineResponse2007
	isSet bool
}

func (v NullableInlineResponse2007) Get() *InlineResponse2007 {
	return v.value
}

func (v *NullableInlineResponse2007) Set(val *InlineResponse2007) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2007) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2007) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2007(val *InlineResponse2007) *NullableInlineResponse2007 {
	return &NullableInlineResponse2007{value: val, isSet: true}
}

func (v NullableInlineResponse2007) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2007) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

