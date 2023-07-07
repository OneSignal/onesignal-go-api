# DeliveryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successful** | Pointer to **NullableInt32** | Number of messages delivered to push servers, mobile carriers, or email service providers. | [optional] 
**Failed** | Pointer to **NullableInt32** | Number of messages sent to unsubscribed devices. | [optional] 
**Errored** | Pointer to **NullableInt32** | Number of errors reported. | [optional] 
**Converted** | Pointer to **NullableInt32** | Number of messages that were clicked. | [optional] 
**Received** | Pointer to **NullableInt32** | Number of devices that received the message. | [optional] 

## Methods

### NewDeliveryData

`func NewDeliveryData() *DeliveryData`

NewDeliveryData instantiates a new DeliveryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryDataWithDefaults

`func NewDeliveryDataWithDefaults() *DeliveryData`

NewDeliveryDataWithDefaults instantiates a new DeliveryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessful

`func (o *DeliveryData) GetSuccessful() int32`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *DeliveryData) GetSuccessfulOk() (*int32, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *DeliveryData) SetSuccessful(v int32)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *DeliveryData) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### SetSuccessfulNil

`func (o *DeliveryData) SetSuccessfulNil(b bool)`

 SetSuccessfulNil sets the value for Successful to be an explicit nil

### UnsetSuccessful
`func (o *DeliveryData) UnsetSuccessful()`

UnsetSuccessful ensures that no value is present for Successful, not even an explicit nil
### GetFailed

`func (o *DeliveryData) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *DeliveryData) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *DeliveryData) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *DeliveryData) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### SetFailedNil

`func (o *DeliveryData) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *DeliveryData) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil
### GetErrored

`func (o *DeliveryData) GetErrored() int32`

GetErrored returns the Errored field if non-nil, zero value otherwise.

### GetErroredOk

`func (o *DeliveryData) GetErroredOk() (*int32, bool)`

GetErroredOk returns a tuple with the Errored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrored

`func (o *DeliveryData) SetErrored(v int32)`

SetErrored sets Errored field to given value.

### HasErrored

`func (o *DeliveryData) HasErrored() bool`

HasErrored returns a boolean if a field has been set.

### SetErroredNil

`func (o *DeliveryData) SetErroredNil(b bool)`

 SetErroredNil sets the value for Errored to be an explicit nil

### UnsetErrored
`func (o *DeliveryData) UnsetErrored()`

UnsetErrored ensures that no value is present for Errored, not even an explicit nil
### GetConverted

`func (o *DeliveryData) GetConverted() int32`

GetConverted returns the Converted field if non-nil, zero value otherwise.

### GetConvertedOk

`func (o *DeliveryData) GetConvertedOk() (*int32, bool)`

GetConvertedOk returns a tuple with the Converted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConverted

`func (o *DeliveryData) SetConverted(v int32)`

SetConverted sets Converted field to given value.

### HasConverted

`func (o *DeliveryData) HasConverted() bool`

HasConverted returns a boolean if a field has been set.

### SetConvertedNil

`func (o *DeliveryData) SetConvertedNil(b bool)`

 SetConvertedNil sets the value for Converted to be an explicit nil

### UnsetConverted
`func (o *DeliveryData) UnsetConverted()`

UnsetConverted ensures that no value is present for Converted, not even an explicit nil
### GetReceived

`func (o *DeliveryData) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *DeliveryData) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *DeliveryData) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *DeliveryData) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### SetReceivedNil

`func (o *DeliveryData) SetReceivedNil(b bool)`

 SetReceivedNil sets the value for Received to be an explicit nil

### UnsetReceived
`func (o *DeliveryData) UnsetReceived()`

UnsetReceived ensures that no value is present for Received, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


