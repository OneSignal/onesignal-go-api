# JourneyValidationErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]JourneyValidationError**](JourneyValidationError.md) |  | [optional] 

## Methods

### NewJourneyValidationErrorResponse

`func NewJourneyValidationErrorResponse() *JourneyValidationErrorResponse`

NewJourneyValidationErrorResponse instantiates a new JourneyValidationErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyValidationErrorResponseWithDefaults

`func NewJourneyValidationErrorResponseWithDefaults() *JourneyValidationErrorResponse`

NewJourneyValidationErrorResponseWithDefaults instantiates a new JourneyValidationErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *JourneyValidationErrorResponse) GetErrors() []JourneyValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *JourneyValidationErrorResponse) GetErrorsOk() (*[]JourneyValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *JourneyValidationErrorResponse) SetErrors(v []JourneyValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *JourneyValidationErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


