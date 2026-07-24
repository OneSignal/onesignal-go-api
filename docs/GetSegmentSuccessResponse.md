# GetSegmentSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriberCount** | Pointer to **int32** | The number of subscribers matching this segment. | [optional] 
**Payload** | Pointer to [**SegmentDetails**](SegmentDetails.md) |  | [optional] 

## Methods

### NewGetSegmentSuccessResponse

`func NewGetSegmentSuccessResponse() *GetSegmentSuccessResponse`

NewGetSegmentSuccessResponse instantiates a new GetSegmentSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSegmentSuccessResponseWithDefaults

`func NewGetSegmentSuccessResponseWithDefaults() *GetSegmentSuccessResponse`

NewGetSegmentSuccessResponseWithDefaults instantiates a new GetSegmentSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriberCount

`func (o *GetSegmentSuccessResponse) GetSubscriberCount() int32`

GetSubscriberCount returns the SubscriberCount field if non-nil, zero value otherwise.

### GetSubscriberCountOk

`func (o *GetSegmentSuccessResponse) GetSubscriberCountOk() (*int32, bool)`

GetSubscriberCountOk returns a tuple with the SubscriberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCount

`func (o *GetSegmentSuccessResponse) SetSubscriberCount(v int32)`

SetSubscriberCount sets SubscriberCount field to given value.

### HasSubscriberCount

`func (o *GetSegmentSuccessResponse) HasSubscriberCount() bool`

HasSubscriberCount returns a boolean if a field has been set.

### GetPayload

`func (o *GetSegmentSuccessResponse) GetPayload() SegmentDetails`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GetSegmentSuccessResponse) GetPayloadOk() (*SegmentDetails, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GetSegmentSuccessResponse) SetPayload(v SegmentDetails)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *GetSegmentSuccessResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


