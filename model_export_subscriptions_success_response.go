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

// ExportSubscriptionsSuccessResponse struct for ExportSubscriptionsSuccessResponse
type ExportSubscriptionsSuccessResponse struct {
	CsvFileUrl *string `json:"csv_file_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExportSubscriptionsSuccessResponse ExportSubscriptionsSuccessResponse

// NewExportSubscriptionsSuccessResponse instantiates a new ExportSubscriptionsSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportSubscriptionsSuccessResponse() *ExportSubscriptionsSuccessResponse {
	this := ExportSubscriptionsSuccessResponse{}
	return &this
}

// NewExportSubscriptionsSuccessResponseWithDefaults instantiates a new ExportSubscriptionsSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportSubscriptionsSuccessResponseWithDefaults() *ExportSubscriptionsSuccessResponse {
	this := ExportSubscriptionsSuccessResponse{}
	return &this
}

// GetCsvFileUrl returns the CsvFileUrl field value if set, zero value otherwise.
func (o *ExportSubscriptionsSuccessResponse) GetCsvFileUrl() string {
	if o == nil || o.CsvFileUrl == nil {
		var ret string
		return ret
	}
	return *o.CsvFileUrl
}

// GetCsvFileUrlOk returns a tuple with the CsvFileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportSubscriptionsSuccessResponse) GetCsvFileUrlOk() (*string, bool) {
	if o == nil || o.CsvFileUrl == nil {
		return nil, false
	}
	return o.CsvFileUrl, true
}

// HasCsvFileUrl returns a boolean if a field has been set.
func (o *ExportSubscriptionsSuccessResponse) HasCsvFileUrl() bool {
	if o != nil && o.CsvFileUrl != nil {
		return true
	}

	return false
}

// SetCsvFileUrl gets a reference to the given string and assigns it to the CsvFileUrl field.
func (o *ExportSubscriptionsSuccessResponse) SetCsvFileUrl(v string) {
	o.CsvFileUrl = &v
}

func (o ExportSubscriptionsSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsvFileUrl != nil {
		toSerialize["csv_file_url"] = o.CsvFileUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ExportSubscriptionsSuccessResponse) UnmarshalJSON(bytes []byte) (err error) {
	varExportSubscriptionsSuccessResponse := _ExportSubscriptionsSuccessResponse{}

	if err = json.Unmarshal(bytes, &varExportSubscriptionsSuccessResponse); err == nil {
		*o = ExportSubscriptionsSuccessResponse(varExportSubscriptionsSuccessResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csv_file_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExportSubscriptionsSuccessResponse struct {
	value *ExportSubscriptionsSuccessResponse
	isSet bool
}

func (v NullableExportSubscriptionsSuccessResponse) Get() *ExportSubscriptionsSuccessResponse {
	return v.value
}

func (v *NullableExportSubscriptionsSuccessResponse) Set(val *ExportSubscriptionsSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExportSubscriptionsSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExportSubscriptionsSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportSubscriptionsSuccessResponse(val *ExportSubscriptionsSuccessResponse) *NullableExportSubscriptionsSuccessResponse {
	return &NullableExportSubscriptionsSuccessResponse{value: val, isSet: true}
}

func (v NullableExportSubscriptionsSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportSubscriptionsSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


