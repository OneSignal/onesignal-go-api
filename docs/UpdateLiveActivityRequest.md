# UpdateLiveActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Type of live activity | 
**Event** | **string** |  | 
**EventUpdates** | **map[string]interface{}** |  | 
**DismissAt** | Pointer to **float32** | Timestamp; only allowed if event is \&quot;end\&quot; | [optional] 

## Methods

### NewUpdateLiveActivityRequest

`func NewUpdateLiveActivityRequest(name string, event string, eventUpdates map[string]interface{}, ) *UpdateLiveActivityRequest`

NewUpdateLiveActivityRequest instantiates a new UpdateLiveActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLiveActivityRequestWithDefaults

`func NewUpdateLiveActivityRequestWithDefaults() *UpdateLiveActivityRequest`

NewUpdateLiveActivityRequestWithDefaults instantiates a new UpdateLiveActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateLiveActivityRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLiveActivityRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLiveActivityRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEvent

`func (o *UpdateLiveActivityRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *UpdateLiveActivityRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *UpdateLiveActivityRequest) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetEventUpdates

`func (o *UpdateLiveActivityRequest) GetEventUpdates() map[string]interface{}`

GetEventUpdates returns the EventUpdates field if non-nil, zero value otherwise.

### GetEventUpdatesOk

`func (o *UpdateLiveActivityRequest) GetEventUpdatesOk() (*map[string]interface{}, bool)`

GetEventUpdatesOk returns a tuple with the EventUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUpdates

`func (o *UpdateLiveActivityRequest) SetEventUpdates(v map[string]interface{})`

SetEventUpdates sets EventUpdates field to given value.


### GetDismissAt

`func (o *UpdateLiveActivityRequest) GetDismissAt() float32`

GetDismissAt returns the DismissAt field if non-nil, zero value otherwise.

### GetDismissAtOk

`func (o *UpdateLiveActivityRequest) GetDismissAtOk() (*float32, bool)`

GetDismissAtOk returns a tuple with the DismissAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissAt

`func (o *UpdateLiveActivityRequest) SetDismissAt(v float32)`

SetDismissAt sets DismissAt field to given value.

### HasDismissAt

`func (o *UpdateLiveActivityRequest) HasDismissAt() bool`

HasDismissAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


