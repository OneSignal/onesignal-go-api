# PlatformDeliveryDataEmailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Opened** | Pointer to **NullableInt32** | Number of times an email has been opened. | [optional] 
**UniqueOpens** | Pointer to **NullableInt32** | Number of unique recipients who have opened your email. | [optional] 
**Clicks** | Pointer to **NullableInt32** | Number of clicked links from your email. This can include the recipient clicking email links multiple times. | [optional] 
**UniqueClicks** | Pointer to **NullableInt32** | Number of unique clicks that your recipients have made on links from your email. | [optional] 
**Bounced** | Pointer to **NullableInt32** | Number of recipients who registered as a hard or soft bounce and didn&#39;t receive your email. | [optional] 
**ReportedSpam** | Pointer to **NullableInt32** | Number of recipients who reported this email as spam. | [optional] 
**Unsubscribed** | Pointer to **NullableInt32** | Number of recipients who opted out of your emails using the unsubscribe link in this email. | [optional] 

## Methods

### NewPlatformDeliveryDataEmailAllOf

`func NewPlatformDeliveryDataEmailAllOf() *PlatformDeliveryDataEmailAllOf`

NewPlatformDeliveryDataEmailAllOf instantiates a new PlatformDeliveryDataEmailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformDeliveryDataEmailAllOfWithDefaults

`func NewPlatformDeliveryDataEmailAllOfWithDefaults() *PlatformDeliveryDataEmailAllOf`

NewPlatformDeliveryDataEmailAllOfWithDefaults instantiates a new PlatformDeliveryDataEmailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpened

`func (o *PlatformDeliveryDataEmailAllOf) GetOpened() int32`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *PlatformDeliveryDataEmailAllOf) GetOpenedOk() (*int32, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *PlatformDeliveryDataEmailAllOf) SetOpened(v int32)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *PlatformDeliveryDataEmailAllOf) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### SetOpenedNil

`func (o *PlatformDeliveryDataEmailAllOf) SetOpenedNil(b bool)`

 SetOpenedNil sets the value for Opened to be an explicit nil

### UnsetOpened
`func (o *PlatformDeliveryDataEmailAllOf) UnsetOpened()`

UnsetOpened ensures that no value is present for Opened, not even an explicit nil
### GetUniqueOpens

`func (o *PlatformDeliveryDataEmailAllOf) GetUniqueOpens() int32`

GetUniqueOpens returns the UniqueOpens field if non-nil, zero value otherwise.

### GetUniqueOpensOk

`func (o *PlatformDeliveryDataEmailAllOf) GetUniqueOpensOk() (*int32, bool)`

GetUniqueOpensOk returns a tuple with the UniqueOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueOpens

`func (o *PlatformDeliveryDataEmailAllOf) SetUniqueOpens(v int32)`

SetUniqueOpens sets UniqueOpens field to given value.

### HasUniqueOpens

`func (o *PlatformDeliveryDataEmailAllOf) HasUniqueOpens() bool`

HasUniqueOpens returns a boolean if a field has been set.

### SetUniqueOpensNil

`func (o *PlatformDeliveryDataEmailAllOf) SetUniqueOpensNil(b bool)`

 SetUniqueOpensNil sets the value for UniqueOpens to be an explicit nil

### UnsetUniqueOpens
`func (o *PlatformDeliveryDataEmailAllOf) UnsetUniqueOpens()`

UnsetUniqueOpens ensures that no value is present for UniqueOpens, not even an explicit nil
### GetClicks

`func (o *PlatformDeliveryDataEmailAllOf) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *PlatformDeliveryDataEmailAllOf) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *PlatformDeliveryDataEmailAllOf) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *PlatformDeliveryDataEmailAllOf) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### SetClicksNil

`func (o *PlatformDeliveryDataEmailAllOf) SetClicksNil(b bool)`

 SetClicksNil sets the value for Clicks to be an explicit nil

### UnsetClicks
`func (o *PlatformDeliveryDataEmailAllOf) UnsetClicks()`

UnsetClicks ensures that no value is present for Clicks, not even an explicit nil
### GetUniqueClicks

`func (o *PlatformDeliveryDataEmailAllOf) GetUniqueClicks() int32`

GetUniqueClicks returns the UniqueClicks field if non-nil, zero value otherwise.

### GetUniqueClicksOk

`func (o *PlatformDeliveryDataEmailAllOf) GetUniqueClicksOk() (*int32, bool)`

GetUniqueClicksOk returns a tuple with the UniqueClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClicks

`func (o *PlatformDeliveryDataEmailAllOf) SetUniqueClicks(v int32)`

SetUniqueClicks sets UniqueClicks field to given value.

### HasUniqueClicks

`func (o *PlatformDeliveryDataEmailAllOf) HasUniqueClicks() bool`

HasUniqueClicks returns a boolean if a field has been set.

### SetUniqueClicksNil

`func (o *PlatformDeliveryDataEmailAllOf) SetUniqueClicksNil(b bool)`

 SetUniqueClicksNil sets the value for UniqueClicks to be an explicit nil

### UnsetUniqueClicks
`func (o *PlatformDeliveryDataEmailAllOf) UnsetUniqueClicks()`

UnsetUniqueClicks ensures that no value is present for UniqueClicks, not even an explicit nil
### GetBounced

`func (o *PlatformDeliveryDataEmailAllOf) GetBounced() int32`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *PlatformDeliveryDataEmailAllOf) GetBouncedOk() (*int32, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *PlatformDeliveryDataEmailAllOf) SetBounced(v int32)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *PlatformDeliveryDataEmailAllOf) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### SetBouncedNil

`func (o *PlatformDeliveryDataEmailAllOf) SetBouncedNil(b bool)`

 SetBouncedNil sets the value for Bounced to be an explicit nil

### UnsetBounced
`func (o *PlatformDeliveryDataEmailAllOf) UnsetBounced()`

UnsetBounced ensures that no value is present for Bounced, not even an explicit nil
### GetReportedSpam

`func (o *PlatformDeliveryDataEmailAllOf) GetReportedSpam() int32`

GetReportedSpam returns the ReportedSpam field if non-nil, zero value otherwise.

### GetReportedSpamOk

`func (o *PlatformDeliveryDataEmailAllOf) GetReportedSpamOk() (*int32, bool)`

GetReportedSpamOk returns a tuple with the ReportedSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedSpam

`func (o *PlatformDeliveryDataEmailAllOf) SetReportedSpam(v int32)`

SetReportedSpam sets ReportedSpam field to given value.

### HasReportedSpam

`func (o *PlatformDeliveryDataEmailAllOf) HasReportedSpam() bool`

HasReportedSpam returns a boolean if a field has been set.

### SetReportedSpamNil

`func (o *PlatformDeliveryDataEmailAllOf) SetReportedSpamNil(b bool)`

 SetReportedSpamNil sets the value for ReportedSpam to be an explicit nil

### UnsetReportedSpam
`func (o *PlatformDeliveryDataEmailAllOf) UnsetReportedSpam()`

UnsetReportedSpam ensures that no value is present for ReportedSpam, not even an explicit nil
### GetUnsubscribed

`func (o *PlatformDeliveryDataEmailAllOf) GetUnsubscribed() int32`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *PlatformDeliveryDataEmailAllOf) GetUnsubscribedOk() (*int32, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *PlatformDeliveryDataEmailAllOf) SetUnsubscribed(v int32)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *PlatformDeliveryDataEmailAllOf) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### SetUnsubscribedNil

`func (o *PlatformDeliveryDataEmailAllOf) SetUnsubscribedNil(b bool)`

 SetUnsubscribedNil sets the value for Unsubscribed to be an explicit nil

### UnsetUnsubscribed
`func (o *PlatformDeliveryDataEmailAllOf) UnsetUnsubscribed()`

UnsetUnsubscribed ensures that no value is present for Unsubscribed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


