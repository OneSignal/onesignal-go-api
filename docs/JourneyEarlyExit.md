# JourneyEarlyExit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**JourneyEarlyExitRules**](JourneyEarlyExitRules.md) |  | [optional] 
**TagOnEarlyExit** | Pointer to **map[string]string** | Tag key-value pairs applied when a user exits early. | [optional] 

## Methods

### NewJourneyEarlyExit

`func NewJourneyEarlyExit() *JourneyEarlyExit`

NewJourneyEarlyExit instantiates a new JourneyEarlyExit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyEarlyExitWithDefaults

`func NewJourneyEarlyExitWithDefaults() *JourneyEarlyExit`

NewJourneyEarlyExitWithDefaults instantiates a new JourneyEarlyExit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *JourneyEarlyExit) GetRules() JourneyEarlyExitRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *JourneyEarlyExit) GetRulesOk() (*JourneyEarlyExitRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *JourneyEarlyExit) SetRules(v JourneyEarlyExitRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *JourneyEarlyExit) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTagOnEarlyExit

`func (o *JourneyEarlyExit) GetTagOnEarlyExit() map[string]string`

GetTagOnEarlyExit returns the TagOnEarlyExit field if non-nil, zero value otherwise.

### GetTagOnEarlyExitOk

`func (o *JourneyEarlyExit) GetTagOnEarlyExitOk() (*map[string]string, bool)`

GetTagOnEarlyExitOk returns a tuple with the TagOnEarlyExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagOnEarlyExit

`func (o *JourneyEarlyExit) SetTagOnEarlyExit(v map[string]string)`

SetTagOnEarlyExit sets TagOnEarlyExit field to given value.

### HasTagOnEarlyExit

`func (o *JourneyEarlyExit) HasTagOnEarlyExit() bool`

HasTagOnEarlyExit returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


