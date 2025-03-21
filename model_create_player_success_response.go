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

// CreatePlayerSuccessResponse struct for CreatePlayerSuccessResponse
type CreatePlayerSuccessResponse struct {
	Success *bool `json:"success,omitempty"`
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreatePlayerSuccessResponse CreatePlayerSuccessResponse

// NewCreatePlayerSuccessResponse instantiates a new CreatePlayerSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePlayerSuccessResponse() *CreatePlayerSuccessResponse {
	this := CreatePlayerSuccessResponse{}
	return &this
}

// NewCreatePlayerSuccessResponseWithDefaults instantiates a new CreatePlayerSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePlayerSuccessResponseWithDefaults() *CreatePlayerSuccessResponse {
	this := CreatePlayerSuccessResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreatePlayerSuccessResponse) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlayerSuccessResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CreatePlayerSuccessResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreatePlayerSuccessResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreatePlayerSuccessResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlayerSuccessResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreatePlayerSuccessResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreatePlayerSuccessResponse) SetId(v string) {
	o.Id = &v
}

func (o CreatePlayerSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreatePlayerSuccessResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreatePlayerSuccessResponse := _CreatePlayerSuccessResponse{}

	if err = json.Unmarshal(bytes, &varCreatePlayerSuccessResponse); err == nil {
		*o = CreatePlayerSuccessResponse(varCreatePlayerSuccessResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreatePlayerSuccessResponse struct {
	value *CreatePlayerSuccessResponse
	isSet bool
}

func (v NullableCreatePlayerSuccessResponse) Get() *CreatePlayerSuccessResponse {
	return v.value
}

func (v *NullableCreatePlayerSuccessResponse) Set(val *CreatePlayerSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePlayerSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePlayerSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePlayerSuccessResponse(val *CreatePlayerSuccessResponse) *NullableCreatePlayerSuccessResponse {
	return &NullableCreatePlayerSuccessResponse{value: val, isSet: true}
}

func (v NullableCreatePlayerSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePlayerSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


