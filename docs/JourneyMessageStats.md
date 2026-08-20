# JourneyMessageStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Totals** | Pointer to **map[string]float32** | All-time totals for this node, keyed by channel-specific stat name. | [optional] 

## Methods

### NewJourneyMessageStats

`func NewJourneyMessageStats() *JourneyMessageStats`

NewJourneyMessageStats instantiates a new JourneyMessageStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyMessageStatsWithDefaults

`func NewJourneyMessageStatsWithDefaults() *JourneyMessageStats`

NewJourneyMessageStatsWithDefaults instantiates a new JourneyMessageStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotals

`func (o *JourneyMessageStats) GetTotals() map[string]float32`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *JourneyMessageStats) GetTotalsOk() (*map[string]float32, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *JourneyMessageStats) SetTotals(v map[string]float32)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *JourneyMessageStats) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


