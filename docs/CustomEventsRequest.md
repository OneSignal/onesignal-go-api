# CustomEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]CustomEvent**](CustomEvent.md) |  | 

## Methods

### NewCustomEventsRequest

`func NewCustomEventsRequest(events []CustomEvent, ) *CustomEventsRequest`

NewCustomEventsRequest instantiates a new CustomEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEventsRequestWithDefaults

`func NewCustomEventsRequestWithDefaults() *CustomEventsRequest`

NewCustomEventsRequestWithDefaults instantiates a new CustomEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *CustomEventsRequest) GetEvents() []CustomEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CustomEventsRequest) GetEventsOk() (*[]CustomEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CustomEventsRequest) SetEvents(v []CustomEvent)`

SetEvents sets Events field to given value.



[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


