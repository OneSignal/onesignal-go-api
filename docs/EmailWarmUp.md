# EmailWarmUp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stages** | Pointer to [**[]EmailWarmUpStage**](EmailWarmUpStage.md) | The campaign&#39;s sending schedule, stage by stage. | [optional] 
**Strategy** | Pointer to **string** | How the stage schedule was produced:   * &#x60;recommended&#x60; - OneSignal generated (and may still adjust) the schedule based on past delivery volumes, scheduled Auto Warm Up emails, and the size of the current audience.   * &#x60;custom&#x60; - The stages were provided as-is in the create request.  | [optional] 
**Status** | Pointer to **string** | Current status of the campaign:   * &#x60;initializing&#x60; - The stages have been submitted and the schedule is being set up.   * &#x60;draft&#x60; - The campaign has been created but has not started sending.   * &#x60;active&#x60; - The campaign is currently working through its stages.   * &#x60;finished&#x60; - All stages have completed.   * &#x60;canceled&#x60; - The campaign was canceled before finishing.  | [optional] 
**IsLive** | Pointer to **bool** | Whether the campaign is currently live (actively sending). | [optional] 

## Methods

### NewEmailWarmUp

`func NewEmailWarmUp() *EmailWarmUp`

NewEmailWarmUp instantiates a new EmailWarmUp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWarmUpWithDefaults

`func NewEmailWarmUpWithDefaults() *EmailWarmUp`

NewEmailWarmUpWithDefaults instantiates a new EmailWarmUp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStages

`func (o *EmailWarmUp) GetStages() []EmailWarmUpStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *EmailWarmUp) GetStagesOk() (*[]EmailWarmUpStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *EmailWarmUp) SetStages(v []EmailWarmUpStage)`

SetStages sets Stages field to given value.

### HasStages

`func (o *EmailWarmUp) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetStrategy

`func (o *EmailWarmUp) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *EmailWarmUp) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *EmailWarmUp) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *EmailWarmUp) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetStatus

`func (o *EmailWarmUp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailWarmUp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailWarmUp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailWarmUp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsLive

`func (o *EmailWarmUp) GetIsLive() bool`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *EmailWarmUp) GetIsLiveOk() (*bool, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLive

`func (o *EmailWarmUp) SetIsLive(v bool)`

SetIsLive sets IsLive field to given value.

### HasIsLive

`func (o *EmailWarmUp) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


