# EmailProviderMetricsTimeseriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]EmailProviderMetricsEvent**](EmailProviderMetricsEvent.md) | One entry for every time bucket in the reporting window, oldest first. Buckets with no email activity are included with all metrics set to &#x60;0&#x60;. Empty when the domain is not found for the app or is not currently active for sending. | [optional] 

## Methods

### NewEmailProviderMetricsTimeseriesResponse

`func NewEmailProviderMetricsTimeseriesResponse() *EmailProviderMetricsTimeseriesResponse`

NewEmailProviderMetricsTimeseriesResponse instantiates a new EmailProviderMetricsTimeseriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailProviderMetricsTimeseriesResponseWithDefaults

`func NewEmailProviderMetricsTimeseriesResponseWithDefaults() *EmailProviderMetricsTimeseriesResponse`

NewEmailProviderMetricsTimeseriesResponseWithDefaults instantiates a new EmailProviderMetricsTimeseriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EmailProviderMetricsTimeseriesResponse) GetEvents() []EmailProviderMetricsEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EmailProviderMetricsTimeseriesResponse) GetEventsOk() (*[]EmailProviderMetricsEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EmailProviderMetricsTimeseriesResponse) SetEvents(v []EmailProviderMetricsEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *EmailProviderMetricsTimeseriesResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


