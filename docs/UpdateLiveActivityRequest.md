# UpdateLiveActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | An internal name to assist with your campaign organization. This does not get displayed in the message itself. | 
**Event** | **string** |  | 
**EventUpdates** | **map[string]interface{}** | This must match the ContentState interface you have defined within your Live Activity in your app. | 
**Contents** | Pointer to [**LanguageStringMap**](LanguageStringMap.md) |  | [optional] 
**Headings** | Pointer to [**LanguageStringMap**](LanguageStringMap.md) |  | [optional] 
**Sound** | Pointer to **string** | Sound file that is included in your app to play instead of the default device notification sound. Omit to disable vibration and sound for the notification. | [optional] 
**StaleDate** | Pointer to **int32** | Accepts Unix timestamp in seconds. When time reaches the configured stale date, the system considers the Live Activity out of date, and the ActivityState of the Live Activity changes to ActivityState.stale. | [optional] 
**DismissalDate** | Pointer to **int32** | Accepts Unix timestamp in seconds; only allowed if event is \&quot;end\&quot; | [optional] 
**Priority** | Pointer to **int32** | Delivery priority through the the push provider (APNs). Pass 10 for higher priority notifications, or 5 for lower priority notifications. Lower priority notifications are sent based on the power considerations of the end user&#39;s device. If not set, defaults to 10. Some providers (APNs) allow for a limited budget of high priority notifications per hour, and if that budget is exceeded, the provider may throttle notification delivery. | [optional] 

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


### GetContents

`func (o *UpdateLiveActivityRequest) GetContents() LanguageStringMap`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *UpdateLiveActivityRequest) GetContentsOk() (*LanguageStringMap, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *UpdateLiveActivityRequest) SetContents(v LanguageStringMap)`

SetContents sets Contents field to given value.

### HasContents

`func (o *UpdateLiveActivityRequest) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetHeadings

`func (o *UpdateLiveActivityRequest) GetHeadings() LanguageStringMap`

GetHeadings returns the Headings field if non-nil, zero value otherwise.

### GetHeadingsOk

`func (o *UpdateLiveActivityRequest) GetHeadingsOk() (*LanguageStringMap, bool)`

GetHeadingsOk returns a tuple with the Headings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadings

`func (o *UpdateLiveActivityRequest) SetHeadings(v LanguageStringMap)`

SetHeadings sets Headings field to given value.

### HasHeadings

`func (o *UpdateLiveActivityRequest) HasHeadings() bool`

HasHeadings returns a boolean if a field has been set.

### GetSound

`func (o *UpdateLiveActivityRequest) GetSound() string`

GetSound returns the Sound field if non-nil, zero value otherwise.

### GetSoundOk

`func (o *UpdateLiveActivityRequest) GetSoundOk() (*string, bool)`

GetSoundOk returns a tuple with the Sound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSound

`func (o *UpdateLiveActivityRequest) SetSound(v string)`

SetSound sets Sound field to given value.

### HasSound

`func (o *UpdateLiveActivityRequest) HasSound() bool`

HasSound returns a boolean if a field has been set.

### GetStaleDate

`func (o *UpdateLiveActivityRequest) GetStaleDate() int32`

GetStaleDate returns the StaleDate field if non-nil, zero value otherwise.

### GetStaleDateOk

`func (o *UpdateLiveActivityRequest) GetStaleDateOk() (*int32, bool)`

GetStaleDateOk returns a tuple with the StaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleDate

`func (o *UpdateLiveActivityRequest) SetStaleDate(v int32)`

SetStaleDate sets StaleDate field to given value.

### HasStaleDate

`func (o *UpdateLiveActivityRequest) HasStaleDate() bool`

HasStaleDate returns a boolean if a field has been set.

### GetDismissalDate

`func (o *UpdateLiveActivityRequest) GetDismissalDate() int32`

GetDismissalDate returns the DismissalDate field if non-nil, zero value otherwise.

### GetDismissalDateOk

`func (o *UpdateLiveActivityRequest) GetDismissalDateOk() (*int32, bool)`

GetDismissalDateOk returns a tuple with the DismissalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissalDate

`func (o *UpdateLiveActivityRequest) SetDismissalDate(v int32)`

SetDismissalDate sets DismissalDate field to given value.

### HasDismissalDate

`func (o *UpdateLiveActivityRequest) HasDismissalDate() bool`

HasDismissalDate returns a boolean if a field has been set.

### GetPriority

`func (o *UpdateLiveActivityRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateLiveActivityRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateLiveActivityRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateLiveActivityRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


