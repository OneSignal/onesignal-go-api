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

// Purchase struct for Purchase
type Purchase struct {
	// The unique identifier of the purchased item.
	Sku string `json:"sku"`
	// The amount, in USD, spent purchasing the item.
	Amount float32 `json:"amount"`
	// The 3-letter ISO 4217 currency code. Required for correct storage and conversion of amount.
	Iso string `json:"iso"`
	AdditionalProperties map[string]interface{}
}

type _Purchase Purchase

// NewPurchase instantiates a new Purchase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchase(sku string, amount float32, iso string) *Purchase {
	this := Purchase{}
	this.Sku = sku
	this.Amount = amount
	this.Iso = iso
	return &this
}

// NewPurchaseWithDefaults instantiates a new Purchase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseWithDefaults() *Purchase {
	this := Purchase{}
	return &this
}

// GetSku returns the Sku field value
func (o *Purchase) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *Purchase) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *Purchase) SetSku(v string) {
	o.Sku = v
}

// GetAmount returns the Amount field value
func (o *Purchase) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Purchase) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Purchase) SetAmount(v float32) {
	o.Amount = v
}

// GetIso returns the Iso field value
func (o *Purchase) GetIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iso
}

// GetIsoOk returns a tuple with the Iso field value
// and a boolean to check if the value has been set.
func (o *Purchase) GetIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iso, true
}

// SetIso sets field value
func (o *Purchase) SetIso(v string) {
	o.Iso = v
}

func (o Purchase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sku"] = o.Sku
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso"] = o.Iso
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Purchase) UnmarshalJSON(bytes []byte) (err error) {
	varPurchase := _Purchase{}

	if err = json.Unmarshal(bytes, &varPurchase); err == nil {
		*o = Purchase(varPurchase)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sku")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "iso")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePurchase struct {
	value *Purchase
	isSet bool
}

func (v NullablePurchase) Get() *Purchase {
	return v.value
}

func (v *NullablePurchase) Set(val *Purchase) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchase) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchase(val *Purchase) *NullablePurchase {
	return &NullablePurchase{value: val, isSet: true}
}

func (v NullablePurchase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


