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

// GenericError struct for GenericError
type GenericError struct {
	Errors []GenericErrorErrorsInner `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenericError GenericError

// NewGenericError instantiates a new GenericError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericError() *GenericError {
	this := GenericError{}
	return &this
}

// NewGenericErrorWithDefaults instantiates a new GenericError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericErrorWithDefaults() *GenericError {
	this := GenericError{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GenericError) GetErrors() []GenericErrorErrorsInner {
	if o == nil || o.Errors == nil {
		var ret []GenericErrorErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericError) GetErrorsOk() ([]GenericErrorErrorsInner, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GenericError) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GenericErrorErrorsInner and assigns it to the Errors field.
func (o *GenericError) SetErrors(v []GenericErrorErrorsInner) {
	o.Errors = v
}

func (o GenericError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GenericError) UnmarshalJSON(bytes []byte) (err error) {
	varGenericError := _GenericError{}

	if err = json.Unmarshal(bytes, &varGenericError); err == nil {
		*o = GenericError(varGenericError)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenericError struct {
	value *GenericError
	isSet bool
}

func (v NullableGenericError) Get() *GenericError {
	return v.value
}

func (v *NullableGenericError) Set(val *GenericError) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericError) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericError(val *GenericError) *NullableGenericError {
	return &NullableGenericError{value: val, isSet: true}
}

func (v NullableGenericError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


