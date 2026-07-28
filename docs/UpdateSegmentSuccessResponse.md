# UpdateSegmentSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | true if the segment was updated successfully, false otherwise. | [optional] 
**Id** | Pointer to **string** | UUID of the updated segment. | [optional] 

## Methods

### NewUpdateSegmentSuccessResponse

`func NewUpdateSegmentSuccessResponse() *UpdateSegmentSuccessResponse`

NewUpdateSegmentSuccessResponse instantiates a new UpdateSegmentSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSegmentSuccessResponseWithDefaults

`func NewUpdateSegmentSuccessResponseWithDefaults() *UpdateSegmentSuccessResponse`

NewUpdateSegmentSuccessResponseWithDefaults instantiates a new UpdateSegmentSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UpdateSegmentSuccessResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateSegmentSuccessResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateSegmentSuccessResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateSegmentSuccessResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetId

`func (o *UpdateSegmentSuccessResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSegmentSuccessResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSegmentSuccessResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateSegmentSuccessResponse) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


