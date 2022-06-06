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

// Operator struct for Operator
type Operator struct {
	// Strictly, this must be either `\"OR\"`, or `\"AND\"`.  It can be used to compose Filters as part of a Filters object.
	Operator *string `json:"operator,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Operator Operator

// NewOperator instantiates a new Operator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperator() *Operator {
	this := Operator{}
	return &this
}

// NewOperatorWithDefaults instantiates a new Operator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorWithDefaults() *Operator {
	this := Operator{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *Operator) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *Operator) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *Operator) SetOperator(v string) {
	o.Operator = &v
}

func (o Operator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Operator) UnmarshalJSON(bytes []byte) (err error) {
	varOperator := _Operator{}

	if err = json.Unmarshal(bytes, &varOperator); err == nil {
		*o = Operator(varOperator)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "operator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOperator struct {
	value *Operator
	isSet bool
}

func (v NullableOperator) Get() *Operator {
	return v.value
}

func (v *NullableOperator) Set(val *Operator) {
	v.value = val
	v.isSet = true
}

func (v NullableOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperator(val *Operator) *NullableOperator {
	return &NullableOperator{value: val, isSet: true}
}

func (v NullableOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


