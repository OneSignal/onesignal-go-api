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

// GenericSuccessBoolResponse struct for GenericSuccessBoolResponse
type GenericSuccessBoolResponse struct {
	Success *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenericSuccessBoolResponse GenericSuccessBoolResponse

// NewGenericSuccessBoolResponse instantiates a new GenericSuccessBoolResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericSuccessBoolResponse() *GenericSuccessBoolResponse {
	this := GenericSuccessBoolResponse{}
	return &this
}

// NewGenericSuccessBoolResponseWithDefaults instantiates a new GenericSuccessBoolResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericSuccessBoolResponseWithDefaults() *GenericSuccessBoolResponse {
	this := GenericSuccessBoolResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GenericSuccessBoolResponse) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericSuccessBoolResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GenericSuccessBoolResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GenericSuccessBoolResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o GenericSuccessBoolResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GenericSuccessBoolResponse) UnmarshalJSON(bytes []byte) (err error) {
	varGenericSuccessBoolResponse := _GenericSuccessBoolResponse{}

	if err = json.Unmarshal(bytes, &varGenericSuccessBoolResponse); err == nil {
		*o = GenericSuccessBoolResponse(varGenericSuccessBoolResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenericSuccessBoolResponse struct {
	value *GenericSuccessBoolResponse
	isSet bool
}

func (v NullableGenericSuccessBoolResponse) Get() *GenericSuccessBoolResponse {
	return v.value
}

func (v *NullableGenericSuccessBoolResponse) Set(val *GenericSuccessBoolResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericSuccessBoolResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericSuccessBoolResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericSuccessBoolResponse(val *GenericSuccessBoolResponse) *NullableGenericSuccessBoolResponse {
	return &NullableGenericSuccessBoolResponse{value: val, isSet: true}
}

func (v NullableGenericSuccessBoolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericSuccessBoolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


