# EmailMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | Pointer to **int64** | The number of emails sent. | [optional] 
**Delivered** | Pointer to **int64** | The number of emails successfully delivered. | [optional] 
**Opened** | Pointer to **int64** | The total number of email opens, including repeat opens by the same recipient. | [optional] 
**Clicked** | Pointer to **int64** | The total number of email link clicks, including repeat clicks by the same recipient. | [optional] 
**UniqueOpened** | Pointer to **int64** | The number of emails opened at least once. | [optional] 
**UniqueClicked** | Pointer to **int64** | The number of emails with at least one link click. | [optional] 
**Failed** | Pointer to **int64** | The number of emails that failed to send. | [optional] 
**Bounced** | Pointer to **int64** | The number of emails that bounced. | [optional] 
**Complained** | Pointer to **int64** | The number of emails that recipients reported as spam. | [optional] 

## Methods

### NewEmailMetrics

`func NewEmailMetrics() *EmailMetrics`

NewEmailMetrics instantiates a new EmailMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMetricsWithDefaults

`func NewEmailMetricsWithDefaults() *EmailMetrics`

NewEmailMetricsWithDefaults instantiates a new EmailMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *EmailMetrics) GetSent() int64`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *EmailMetrics) GetSentOk() (*int64, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *EmailMetrics) SetSent(v int64)`

SetSent sets Sent field to given value.

### HasSent

`func (o *EmailMetrics) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *EmailMetrics) GetDelivered() int64`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *EmailMetrics) GetDeliveredOk() (*int64, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *EmailMetrics) SetDelivered(v int64)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *EmailMetrics) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetOpened

`func (o *EmailMetrics) GetOpened() int64`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *EmailMetrics) GetOpenedOk() (*int64, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *EmailMetrics) SetOpened(v int64)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *EmailMetrics) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *EmailMetrics) GetClicked() int64`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *EmailMetrics) GetClickedOk() (*int64, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *EmailMetrics) SetClicked(v int64)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *EmailMetrics) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUniqueOpened

`func (o *EmailMetrics) GetUniqueOpened() int64`

GetUniqueOpened returns the UniqueOpened field if non-nil, zero value otherwise.

### GetUniqueOpenedOk

`func (o *EmailMetrics) GetUniqueOpenedOk() (*int64, bool)`

GetUniqueOpenedOk returns a tuple with the UniqueOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueOpened

`func (o *EmailMetrics) SetUniqueOpened(v int64)`

SetUniqueOpened sets UniqueOpened field to given value.

### HasUniqueOpened

`func (o *EmailMetrics) HasUniqueOpened() bool`

HasUniqueOpened returns a boolean if a field has been set.

### GetUniqueClicked

`func (o *EmailMetrics) GetUniqueClicked() int64`

GetUniqueClicked returns the UniqueClicked field if non-nil, zero value otherwise.

### GetUniqueClickedOk

`func (o *EmailMetrics) GetUniqueClickedOk() (*int64, bool)`

GetUniqueClickedOk returns a tuple with the UniqueClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClicked

`func (o *EmailMetrics) SetUniqueClicked(v int64)`

SetUniqueClicked sets UniqueClicked field to given value.

### HasUniqueClicked

`func (o *EmailMetrics) HasUniqueClicked() bool`

HasUniqueClicked returns a boolean if a field has been set.

### GetFailed

`func (o *EmailMetrics) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *EmailMetrics) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *EmailMetrics) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *EmailMetrics) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetBounced

`func (o *EmailMetrics) GetBounced() int64`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *EmailMetrics) GetBouncedOk() (*int64, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *EmailMetrics) SetBounced(v int64)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *EmailMetrics) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetComplained

`func (o *EmailMetrics) GetComplained() int64`

GetComplained returns the Complained field if non-nil, zero value otherwise.

### GetComplainedOk

`func (o *EmailMetrics) GetComplainedOk() (*int64, bool)`

GetComplainedOk returns a tuple with the Complained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplained

`func (o *EmailMetrics) SetComplained(v int64)`

SetComplained sets Complained field to given value.

### HasComplained

`func (o *EmailMetrics) HasComplained() bool`

HasComplained returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


