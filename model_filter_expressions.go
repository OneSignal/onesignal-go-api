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
	"fmt"
)

// FilterExpressions struct for FilterExpressions
type FilterExpressions struct {
	Filter *Filter
	Operator *Operator
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FilterExpressions) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Filter
	err = json.Unmarshal(data, &dst.Filter);
	if err == nil {
		jsonFilter, _ := json.Marshal(dst.Filter)
		if string(jsonFilter) == "{}" { // empty struct
			dst.Filter = nil
		} else {
			return nil // data stored in dst.Filter, return on the first match
		}
	} else {
		dst.Filter = nil
	}

	// try to unmarshal JSON data into Operator
	err = json.Unmarshal(data, &dst.Operator);
	if err == nil {
		jsonOperator, _ := json.Marshal(dst.Operator)
		if string(jsonOperator) == "{}" { // empty struct
			dst.Operator = nil
		} else {
			return nil // data stored in dst.Operator, return on the first match
		}
	} else {
		dst.Operator = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(FilterExpressions)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FilterExpressions) MarshalJSON() ([]byte, error) {
	if src.Filter != nil {
		return json.Marshal(&src.Filter)
	}

	if src.Operator != nil {
		return json.Marshal(&src.Operator)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFilterExpressions struct {
	value *FilterExpressions
	isSet bool
}

func (v NullableFilterExpressions) Get() *FilterExpressions {
	return v.value
}

func (v *NullableFilterExpressions) Set(val *FilterExpressions) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterExpressions) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterExpressions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterExpressions(val *FilterExpressions) *NullableFilterExpressions {
	return &NullableFilterExpressions{value: val, isSet: true}
}

func (v NullableFilterExpressions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterExpressions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


