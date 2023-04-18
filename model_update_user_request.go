/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.2.1
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// UpdateUserRequest struct for UpdateUserRequest
type UpdateUserRequest struct {
	Properties *PropertiesObject `json:"properties,omitempty"`
	RefreshDeviceMetadata *bool `json:"refresh_device_metadata,omitempty"`
	Deltas *PropertiesDeltas `json:"deltas,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateUserRequest UpdateUserRequest

// NewUpdateUserRequest instantiates a new UpdateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequest() *UpdateUserRequest {
	this := UpdateUserRequest{}
	var refreshDeviceMetadata bool = false
	this.RefreshDeviceMetadata = &refreshDeviceMetadata
	return &this
}

// NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestWithDefaults() *UpdateUserRequest {
	this := UpdateUserRequest{}
	var refreshDeviceMetadata bool = false
	this.RefreshDeviceMetadata = &refreshDeviceMetadata
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetProperties() PropertiesObject {
	if o == nil || o.Properties == nil {
		var ret PropertiesObject
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetPropertiesOk() (*PropertiesObject, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given PropertiesObject and assigns it to the Properties field.
func (o *UpdateUserRequest) SetProperties(v PropertiesObject) {
	o.Properties = &v
}

// GetRefreshDeviceMetadata returns the RefreshDeviceMetadata field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetRefreshDeviceMetadata() bool {
	if o == nil || o.RefreshDeviceMetadata == nil {
		var ret bool
		return ret
	}
	return *o.RefreshDeviceMetadata
}

// GetRefreshDeviceMetadataOk returns a tuple with the RefreshDeviceMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetRefreshDeviceMetadataOk() (*bool, bool) {
	if o == nil || o.RefreshDeviceMetadata == nil {
		return nil, false
	}
	return o.RefreshDeviceMetadata, true
}

// HasRefreshDeviceMetadata returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasRefreshDeviceMetadata() bool {
	if o != nil && o.RefreshDeviceMetadata != nil {
		return true
	}

	return false
}

// SetRefreshDeviceMetadata gets a reference to the given bool and assigns it to the RefreshDeviceMetadata field.
func (o *UpdateUserRequest) SetRefreshDeviceMetadata(v bool) {
	o.RefreshDeviceMetadata = &v
}

// GetDeltas returns the Deltas field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetDeltas() PropertiesDeltas {
	if o == nil || o.Deltas == nil {
		var ret PropertiesDeltas
		return ret
	}
	return *o.Deltas
}

// GetDeltasOk returns a tuple with the Deltas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetDeltasOk() (*PropertiesDeltas, bool) {
	if o == nil || o.Deltas == nil {
		return nil, false
	}
	return o.Deltas, true
}

// HasDeltas returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasDeltas() bool {
	if o != nil && o.Deltas != nil {
		return true
	}

	return false
}

// SetDeltas gets a reference to the given PropertiesDeltas and assigns it to the Deltas field.
func (o *UpdateUserRequest) SetDeltas(v PropertiesDeltas) {
	o.Deltas = &v
}

func (o UpdateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.RefreshDeviceMetadata != nil {
		toSerialize["refresh_device_metadata"] = o.RefreshDeviceMetadata
	}
	if o.Deltas != nil {
		toSerialize["deltas"] = o.Deltas
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateUserRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateUserRequest := _UpdateUserRequest{}

	if err = json.Unmarshal(bytes, &varUpdateUserRequest); err == nil {
		*o = UpdateUserRequest(varUpdateUserRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "properties")
		delete(additionalProperties, "refresh_device_metadata")
		delete(additionalProperties, "deltas")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateUserRequest struct {
	value *UpdateUserRequest
	isSet bool
}

func (v NullableUpdateUserRequest) Get() *UpdateUserRequest {
	return v.value
}

func (v *NullableUpdateUserRequest) Set(val *UpdateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRequest(val *UpdateUserRequest) *NullableUpdateUserRequest {
	return &NullableUpdateUserRequest{value: val, isSet: true}
}

func (v NullableUpdateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

