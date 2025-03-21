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

// GenericErrorErrorsInner struct for GenericErrorErrorsInner
type GenericErrorErrorsInner struct {
	Code *string `json:"code,omitempty"`
	Title *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenericErrorErrorsInner GenericErrorErrorsInner

// NewGenericErrorErrorsInner instantiates a new GenericErrorErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericErrorErrorsInner() *GenericErrorErrorsInner {
	this := GenericErrorErrorsInner{}
	return &this
}

// NewGenericErrorErrorsInnerWithDefaults instantiates a new GenericErrorErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericErrorErrorsInnerWithDefaults() *GenericErrorErrorsInner {
	this := GenericErrorErrorsInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GenericErrorErrorsInner) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorErrorsInner) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GenericErrorErrorsInner) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GenericErrorErrorsInner) SetCode(v string) {
	o.Code = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GenericErrorErrorsInner) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorErrorsInner) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GenericErrorErrorsInner) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GenericErrorErrorsInner) SetTitle(v string) {
	o.Title = &v
}

func (o GenericErrorErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GenericErrorErrorsInner) UnmarshalJSON(bytes []byte) (err error) {
	varGenericErrorErrorsInner := _GenericErrorErrorsInner{}

	if err = json.Unmarshal(bytes, &varGenericErrorErrorsInner); err == nil {
		*o = GenericErrorErrorsInner(varGenericErrorErrorsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenericErrorErrorsInner struct {
	value *GenericErrorErrorsInner
	isSet bool
}

func (v NullableGenericErrorErrorsInner) Get() *GenericErrorErrorsInner {
	return v.value
}

func (v *NullableGenericErrorErrorsInner) Set(val *GenericErrorErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericErrorErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericErrorErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericErrorErrorsInner(val *GenericErrorErrorsInner) *NullableGenericErrorErrorsInner {
	return &NullableGenericErrorErrorsInner{value: val, isSet: true}
}

func (v NullableGenericErrorErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericErrorErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


