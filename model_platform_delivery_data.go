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

// PlatformDeliveryData Hash of delivery statistics broken out by target device platform.
type PlatformDeliveryData struct {
	EdgeWebPush *DeliveryData `json:"edge_web_push,omitempty"`
	ChromeWebPush *DeliveryData `json:"chrome_web_push,omitempty"`
	FirefoxWebPush *DeliveryData `json:"firefox_web_push,omitempty"`
	SafariWebPush *DeliveryData `json:"safari_web_push,omitempty"`
	Android *DeliveryData `json:"android,omitempty"`
	Ios *DeliveryData `json:"ios,omitempty"`
	Sms NullableDeliveryData `json:"sms,omitempty"`
	Email NullableDeliveryData `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlatformDeliveryData PlatformDeliveryData

// NewPlatformDeliveryData instantiates a new PlatformDeliveryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformDeliveryData() *PlatformDeliveryData {
	this := PlatformDeliveryData{}
	return &this
}

// NewPlatformDeliveryDataWithDefaults instantiates a new PlatformDeliveryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformDeliveryDataWithDefaults() *PlatformDeliveryData {
	this := PlatformDeliveryData{}
	return &this
}

// GetEdgeWebPush returns the EdgeWebPush field value if set, zero value otherwise.
func (o *PlatformDeliveryData) GetEdgeWebPush() DeliveryData {
	if o == nil || o.EdgeWebPush == nil {
		var ret DeliveryData
		return ret
	}
	return *o.EdgeWebPush
}

// GetEdgeWebPushOk returns a tuple with the EdgeWebPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformDeliveryData) GetEdgeWebPushOk() (*DeliveryData, bool) {
	if o == nil || o.EdgeWebPush == nil {
		return nil, false
	}
	return o.EdgeWebPush, true
}

// HasEdgeWebPush returns a boolean if a field has been set.
func (o *PlatformDeliveryData) HasEdgeWebPush() bool {
	if o != nil && o.EdgeWebPush != nil {
		return true
	}

	return false
}

// SetEdgeWebPush gets a reference to the given DeliveryData and assigns it to the EdgeWebPush field.
func (o *PlatformDeliveryData) SetEdgeWebPush(v DeliveryData) {
	o.EdgeWebPush = &v
}

// GetChromeWebPush returns the ChromeWebPush field value if set, zero value otherwise.
func (o *PlatformDeliveryData) GetChromeWebPush() DeliveryData {
	if o == nil || o.ChromeWebPush == nil {
		var ret DeliveryData
		return ret
	}
	return *o.ChromeWebPush
}

// GetChromeWebPushOk returns a tuple with the ChromeWebPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformDeliveryData) GetChromeWebPushOk() (*DeliveryData, bool) {
	if o == nil || o.ChromeWebPush == nil {
		return nil, false
	}
	return o.ChromeWebPush, true
}

// HasChromeWebPush returns a boolean if a field has been set.
func (o *PlatformDeliveryData) HasChromeWebPush() bool {
	if o != nil && o.ChromeWebPush != nil {
		return true
	}

	return false
}

// SetChromeWebPush gets a reference to the given DeliveryData and assigns it to the ChromeWebPush field.
func (o *PlatformDeliveryData) SetChromeWebPush(v DeliveryData) {
	o.ChromeWebPush = &v
}

// GetFirefoxWebPush returns the FirefoxWebPush field value if set, zero value otherwise.
func (o *PlatformDeliveryData) GetFirefoxWebPush() DeliveryData {
	if o == nil || o.FirefoxWebPush == nil {
		var ret DeliveryData
		return ret
	}
	return *o.FirefoxWebPush
}

// GetFirefoxWebPushOk returns a tuple with the FirefoxWebPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformDeliveryData) GetFirefoxWebPushOk() (*DeliveryData, bool) {
	if o == nil || o.FirefoxWebPush == nil {
		return nil, false
	}
	return o.FirefoxWebPush, true
}

// HasFirefoxWebPush returns a boolean if a field has been set.
func (o *PlatformDeliveryData) HasFirefoxWebPush() bool {
	if o != nil && o.FirefoxWebPush != nil {
		return true
	}

	return false
}

// SetFirefoxWebPush gets a reference to the given DeliveryData and assigns it to the FirefoxWebPush field.
func (o *PlatformDeliveryData) SetFirefoxWebPush(v DeliveryData) {
	o.FirefoxWebPush = &v
}

// GetSafariWebPush returns the SafariWebPush field value if set, zero value otherwise.
func (o *PlatformDeliveryData) GetSafariWebPush() DeliveryData {
	if o == nil || o.SafariWebPush == nil {
		var ret DeliveryData
		return ret
	}
	return *o.SafariWebPush
}

// GetSafariWebPushOk returns a tuple with the SafariWebPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformDeliveryData) GetSafariWebPushOk() (*DeliveryData, bool) {
	if o == nil || o.SafariWebPush == nil {
		return nil, false
	}
	return o.SafariWebPush, true
}

// HasSafariWebPush returns a boolean if a field has been set.
func (o *PlatformDeliveryData) HasSafariWebPush() bool {
	if o != nil && o.SafariWebPush != nil {
		return true
	}

	return false
}

// SetSafariWebPush gets a reference to the given DeliveryData and assigns it to the SafariWebPush field.
func (o *PlatformDeliveryData) SetSafariWebPush(v DeliveryData) {
	o.SafariWebPush = &v
}

// GetAndroid returns the Android field value if set, zero value otherwise.
func (o *PlatformDeliveryData) GetAndroid() DeliveryData {
	if o == nil || o.Android == nil {
		var ret DeliveryData
		return ret
	}
	return *o.Android
}

// GetAndroidOk returns a tuple with the Android field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformDeliveryData) GetAndroidOk() (*DeliveryData, bool) {
	if o == nil || o.Android == nil {
		return nil, false
	}
	return o.Android, true
}

// HasAndroid returns a boolean if a field has been set.
func (o *PlatformDeliveryData) HasAndroid() bool {
	if o != nil && o.Android != nil {
		return true
	}

	return false
}

// SetAndroid gets a reference to the given DeliveryData and assigns it to the Android field.
func (o *PlatformDeliveryData) SetAndroid(v DeliveryData) {
	o.Android = &v
}

// GetIos returns the Ios field value if set, zero value otherwise.
func (o *PlatformDeliveryData) GetIos() DeliveryData {
	if o == nil || o.Ios == nil {
		var ret DeliveryData
		return ret
	}
	return *o.Ios
}

// GetIosOk returns a tuple with the Ios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformDeliveryData) GetIosOk() (*DeliveryData, bool) {
	if o == nil || o.Ios == nil {
		return nil, false
	}
	return o.Ios, true
}

// HasIos returns a boolean if a field has been set.
func (o *PlatformDeliveryData) HasIos() bool {
	if o != nil && o.Ios != nil {
		return true
	}

	return false
}

// SetIos gets a reference to the given DeliveryData and assigns it to the Ios field.
func (o *PlatformDeliveryData) SetIos(v DeliveryData) {
	o.Ios = &v
}

// GetSms returns the Sms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryData) GetSms() DeliveryData {
	if o == nil || o.Sms.Get() == nil {
		var ret DeliveryData
		return ret
	}
	return *o.Sms.Get()
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryData) GetSmsOk() (*DeliveryData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sms.Get(), o.Sms.IsSet()
}

// HasSms returns a boolean if a field has been set.
func (o *PlatformDeliveryData) HasSms() bool {
	if o != nil && o.Sms.IsSet() {
		return true
	}

	return false
}

// SetSms gets a reference to the given NullableDeliveryData and assigns it to the Sms field.
func (o *PlatformDeliveryData) SetSms(v DeliveryData) {
	o.Sms.Set(&v)
}
// SetSmsNil sets the value for Sms to be an explicit nil
func (o *PlatformDeliveryData) SetSmsNil() {
	o.Sms.Set(nil)
}

// UnsetSms ensures that no value is present for Sms, not even an explicit nil
func (o *PlatformDeliveryData) UnsetSms() {
	o.Sms.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryData) GetEmail() DeliveryData {
	if o == nil || o.Email.Get() == nil {
		var ret DeliveryData
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryData) GetEmailOk() (*DeliveryData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *PlatformDeliveryData) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableDeliveryData and assigns it to the Email field.
func (o *PlatformDeliveryData) SetEmail(v DeliveryData) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *PlatformDeliveryData) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *PlatformDeliveryData) UnsetEmail() {
	o.Email.Unset()
}

func (o PlatformDeliveryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EdgeWebPush != nil {
		toSerialize["edge_web_push"] = o.EdgeWebPush
	}
	if o.ChromeWebPush != nil {
		toSerialize["chrome_web_push"] = o.ChromeWebPush
	}
	if o.FirefoxWebPush != nil {
		toSerialize["firefox_web_push"] = o.FirefoxWebPush
	}
	if o.SafariWebPush != nil {
		toSerialize["safari_web_push"] = o.SafariWebPush
	}
	if o.Android != nil {
		toSerialize["android"] = o.Android
	}
	if o.Ios != nil {
		toSerialize["ios"] = o.Ios
	}
	if o.Sms.IsSet() {
		toSerialize["sms"] = o.Sms.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PlatformDeliveryData) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformDeliveryData := _PlatformDeliveryData{}

	if err = json.Unmarshal(bytes, &varPlatformDeliveryData); err == nil {
		*o = PlatformDeliveryData(varPlatformDeliveryData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "edge_web_push")
		delete(additionalProperties, "chrome_web_push")
		delete(additionalProperties, "firefox_web_push")
		delete(additionalProperties, "safari_web_push")
		delete(additionalProperties, "android")
		delete(additionalProperties, "ios")
		delete(additionalProperties, "sms")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlatformDeliveryData struct {
	value *PlatformDeliveryData
	isSet bool
}

func (v NullablePlatformDeliveryData) Get() *PlatformDeliveryData {
	return v.value
}

func (v *NullablePlatformDeliveryData) Set(val *PlatformDeliveryData) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformDeliveryData) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformDeliveryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformDeliveryData(val *PlatformDeliveryData) *NullablePlatformDeliveryData {
	return &NullablePlatformDeliveryData{value: val, isSet: true}
}

func (v NullablePlatformDeliveryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformDeliveryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


