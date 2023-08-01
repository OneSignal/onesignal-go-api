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

// UpdateLiveActivitySuccessResponse struct for UpdateLiveActivitySuccessResponse
type UpdateLiveActivitySuccessResponse struct {
	NotificationId *string `json:"notification_id,omitempty"`
	Errors *Notification200Errors `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateLiveActivitySuccessResponse UpdateLiveActivitySuccessResponse

// NewUpdateLiveActivitySuccessResponse instantiates a new UpdateLiveActivitySuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLiveActivitySuccessResponse() *UpdateLiveActivitySuccessResponse {
	this := UpdateLiveActivitySuccessResponse{}
	return &this
}

// NewUpdateLiveActivitySuccessResponseWithDefaults instantiates a new UpdateLiveActivitySuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLiveActivitySuccessResponseWithDefaults() *UpdateLiveActivitySuccessResponse {
	this := UpdateLiveActivitySuccessResponse{}
	return &this
}

// GetNotificationId returns the NotificationId field value if set, zero value otherwise.
func (o *UpdateLiveActivitySuccessResponse) GetNotificationId() string {
	if o == nil || o.NotificationId == nil {
		var ret string
		return ret
	}
	return *o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLiveActivitySuccessResponse) GetNotificationIdOk() (*string, bool) {
	if o == nil || o.NotificationId == nil {
		return nil, false
	}
	return o.NotificationId, true
}

// HasNotificationId returns a boolean if a field has been set.
func (o *UpdateLiveActivitySuccessResponse) HasNotificationId() bool {
	if o != nil && o.NotificationId != nil {
		return true
	}

	return false
}

// SetNotificationId gets a reference to the given string and assigns it to the NotificationId field.
func (o *UpdateLiveActivitySuccessResponse) SetNotificationId(v string) {
	o.NotificationId = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *UpdateLiveActivitySuccessResponse) GetErrors() Notification200Errors {
	if o == nil || o.Errors == nil {
		var ret Notification200Errors
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLiveActivitySuccessResponse) GetErrorsOk() (*Notification200Errors, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *UpdateLiveActivitySuccessResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given Notification200Errors and assigns it to the Errors field.
func (o *UpdateLiveActivitySuccessResponse) SetErrors(v Notification200Errors) {
	o.Errors = &v
}

func (o UpdateLiveActivitySuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NotificationId != nil {
		toSerialize["notification_id"] = o.NotificationId
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateLiveActivitySuccessResponse) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateLiveActivitySuccessResponse := _UpdateLiveActivitySuccessResponse{}

	if err = json.Unmarshal(bytes, &varUpdateLiveActivitySuccessResponse); err == nil {
		*o = UpdateLiveActivitySuccessResponse(varUpdateLiveActivitySuccessResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "notification_id")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateLiveActivitySuccessResponse struct {
	value *UpdateLiveActivitySuccessResponse
	isSet bool
}

func (v NullableUpdateLiveActivitySuccessResponse) Get() *UpdateLiveActivitySuccessResponse {
	return v.value
}

func (v *NullableUpdateLiveActivitySuccessResponse) Set(val *UpdateLiveActivitySuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLiveActivitySuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLiveActivitySuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLiveActivitySuccessResponse(val *UpdateLiveActivitySuccessResponse) *NullableUpdateLiveActivitySuccessResponse {
	return &NullableUpdateLiveActivitySuccessResponse{value: val, isSet: true}
}

func (v NullableUpdateLiveActivitySuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLiveActivitySuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


