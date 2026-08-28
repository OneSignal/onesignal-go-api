# EmailWarmUpStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** | ISO 8601 timestamp for the start of this stage. Sending for this stage will not begin before this time. | 
**End** | **time.Time** | ISO 8601 timestamp for the end of this stage. This stage&#39;s quota is expected to be sent by this time. | 
**Quota** | **int32** | Number of emails to send during this stage. | 
**Acked** | Pointer to **bool** | Whether this stage has been picked up and acknowledged by the warm-up scheduler. Not accepted on create. This is only present when reading back a campaign. | [optional] [readonly] 

## Methods

### NewEmailWarmUpStage

`func NewEmailWarmUpStage(start time.Time, end time.Time, quota int32, ) *EmailWarmUpStage`

NewEmailWarmUpStage instantiates a new EmailWarmUpStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWarmUpStageWithDefaults

`func NewEmailWarmUpStageWithDefaults() *EmailWarmUpStage`

NewEmailWarmUpStageWithDefaults instantiates a new EmailWarmUpStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *EmailWarmUpStage) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EmailWarmUpStage) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EmailWarmUpStage) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *EmailWarmUpStage) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EmailWarmUpStage) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EmailWarmUpStage) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetQuota

`func (o *EmailWarmUpStage) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *EmailWarmUpStage) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *EmailWarmUpStage) SetQuota(v int32)`

SetQuota sets Quota field to given value.


### GetAcked

`func (o *EmailWarmUpStage) GetAcked() bool`

GetAcked returns the Acked field if non-nil, zero value otherwise.

### GetAckedOk

`func (o *EmailWarmUpStage) GetAckedOk() (*bool, bool)`

GetAckedOk returns a tuple with the Acked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcked

`func (o *EmailWarmUpStage) SetAcked(v bool)`

SetAcked sets Acked field to given value.

### HasAcked

`func (o *EmailWarmUpStage) HasAcked() bool`

HasAcked returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


