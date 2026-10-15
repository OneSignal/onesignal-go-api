# EmailProviderMetricsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** | The start of the time bucket, aligned to the hour, day, or month in UTC. | [optional] 
**Metrics** | Pointer to [**EmailMetrics**](EmailMetrics.md) |  | [optional] 

## Methods

### NewEmailProviderMetricsEvent

`func NewEmailProviderMetricsEvent() *EmailProviderMetricsEvent`

NewEmailProviderMetricsEvent instantiates a new EmailProviderMetricsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailProviderMetricsEventWithDefaults

`func NewEmailProviderMetricsEventWithDefaults() *EmailProviderMetricsEvent`

NewEmailProviderMetricsEventWithDefaults instantiates a new EmailProviderMetricsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *EmailProviderMetricsEvent) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EmailProviderMetricsEvent) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EmailProviderMetricsEvent) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *EmailProviderMetricsEvent) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMetrics

`func (o *EmailProviderMetricsEvent) GetMetrics() EmailMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *EmailProviderMetricsEvent) GetMetricsOk() (*EmailMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *EmailProviderMetricsEvent) SetMetrics(v EmailMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *EmailProviderMetricsEvent) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


