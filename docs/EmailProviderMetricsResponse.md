# EmailProviderMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]EmailProviderMetrics**](EmailProviderMetrics.md) | One entry per inbox provider that received email from the domain during the reporting window. Empty when the domain is not found for the app or is not currently active for sending. | [optional] 

## Methods

### NewEmailProviderMetricsResponse

`func NewEmailProviderMetricsResponse() *EmailProviderMetricsResponse`

NewEmailProviderMetricsResponse instantiates a new EmailProviderMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailProviderMetricsResponseWithDefaults

`func NewEmailProviderMetricsResponseWithDefaults() *EmailProviderMetricsResponse`

NewEmailProviderMetricsResponseWithDefaults instantiates a new EmailProviderMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *EmailProviderMetricsResponse) GetProviders() []EmailProviderMetrics`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *EmailProviderMetricsResponse) GetProvidersOk() (*[]EmailProviderMetrics, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *EmailProviderMetricsResponse) SetProviders(v []EmailProviderMetrics)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *EmailProviderMetricsResponse) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


