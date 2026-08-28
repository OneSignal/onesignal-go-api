# EmailWarmUpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stages** | [**[]EmailWarmUpStage**](EmailWarmUpStage.md) | Required. The ordered stages that make up the campaign&#39;s sending schedule. | 
**Strategy** | Pointer to **NullableString** | How the stage schedule should be treated:   * &#x60;recommended&#x60; - (Default) OneSignal may adjust the provided stages based on past delivery volumes, scheduled Auto Warm Up emails, and the size of the current audience.   * &#x60;custom&#x60; - The stages provided are sent as-is.  | [optional] 

## Methods

### NewEmailWarmUpRequest

`func NewEmailWarmUpRequest(stages []EmailWarmUpStage, ) *EmailWarmUpRequest`

NewEmailWarmUpRequest instantiates a new EmailWarmUpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWarmUpRequestWithDefaults

`func NewEmailWarmUpRequestWithDefaults() *EmailWarmUpRequest`

NewEmailWarmUpRequestWithDefaults instantiates a new EmailWarmUpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStages

`func (o *EmailWarmUpRequest) GetStages() []EmailWarmUpStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *EmailWarmUpRequest) GetStagesOk() (*[]EmailWarmUpStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *EmailWarmUpRequest) SetStages(v []EmailWarmUpStage)`

SetStages sets Stages field to given value.


### GetStrategy

`func (o *EmailWarmUpRequest) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *EmailWarmUpRequest) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *EmailWarmUpRequest) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *EmailWarmUpRequest) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### SetStrategyNil

`func (o *EmailWarmUpRequest) SetStrategyNil(b bool)`

 SetStrategyNil sets the value for Strategy to be an explicit nil

### UnsetStrategy
`func (o *EmailWarmUpRequest) UnsetStrategy()`

UnsetStrategy ensures that no value is present for Strategy, not even an explicit nil

[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


