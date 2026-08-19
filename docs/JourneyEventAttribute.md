# JourneyEventAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Event attribute key. | 
**Operator** | **string** | Comparison operator. | 
**Value** | Pointer to **string** | Value to compare against. Not required for exists and not_exists. | [optional] 

## Methods

### NewJourneyEventAttribute

`func NewJourneyEventAttribute(key string, operator string, ) *JourneyEventAttribute`

NewJourneyEventAttribute instantiates a new JourneyEventAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyEventAttributeWithDefaults

`func NewJourneyEventAttributeWithDefaults() *JourneyEventAttribute`

NewJourneyEventAttributeWithDefaults instantiates a new JourneyEventAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *JourneyEventAttribute) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *JourneyEventAttribute) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *JourneyEventAttribute) SetKey(v string)`

SetKey sets Key field to given value.


### GetOperator

`func (o *JourneyEventAttribute) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *JourneyEventAttribute) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *JourneyEventAttribute) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *JourneyEventAttribute) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JourneyEventAttribute) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JourneyEventAttribute) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *JourneyEventAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


