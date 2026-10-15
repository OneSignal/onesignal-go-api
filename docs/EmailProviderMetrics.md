# EmailProviderMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** | The inbox provider. Known values include &#x60;apple&#x60;, &#x60;att&#x60;, &#x60;canadian&#x60;, &#x60;chinese&#x60;, &#x60;comcast&#x60;, &#x60;free_fr&#x60;, &#x60;french&#x60;, &#x60;german&#x60;, &#x60;gmail&#x60;, &#x60;google_workspace&#x60;, &#x60;googlers&#x60;, &#x60;inbox_lv&#x60;, &#x60;italian&#x60;, &#x60;japanese&#x60;, &#x60;outlook_365&#x60;, &#x60;outlook_eu&#x60;, &#x60;outlook_us&#x60;, &#x60;proton&#x60;, &#x60;rackspace&#x60;, &#x60;russian&#x60;, &#x60;uk&#x60;, &#x60;united_online&#x60;, &#x60;yahoo_eu&#x60;, &#x60;yahoo_us&#x60;, &#x60;other_eu&#x60;, &#x60;other_us&#x60;, and &#x60;other&#x60;. New providers may be added over time. | [optional] 
**Metrics** | Pointer to [**EmailMetrics**](EmailMetrics.md) |  | [optional] 

## Methods

### NewEmailProviderMetrics

`func NewEmailProviderMetrics() *EmailProviderMetrics`

NewEmailProviderMetrics instantiates a new EmailProviderMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailProviderMetricsWithDefaults

`func NewEmailProviderMetricsWithDefaults() *EmailProviderMetrics`

NewEmailProviderMetricsWithDefaults instantiates a new EmailProviderMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *EmailProviderMetrics) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EmailProviderMetrics) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EmailProviderMetrics) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EmailProviderMetrics) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetMetrics

`func (o *EmailProviderMetrics) GetMetrics() EmailMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *EmailProviderMetrics) GetMetricsOk() (*EmailMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *EmailProviderMetrics) SetMetrics(v EmailMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *EmailProviderMetrics) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


