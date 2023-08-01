# RateLimiterError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]GenericErrorErrorsInner**](GenericErrorErrorsInner.md) |  | [optional] 

## Methods

### NewRateLimiterError

`func NewRateLimiterError() *RateLimiterError`

NewRateLimiterError instantiates a new RateLimiterError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimiterErrorWithDefaults

`func NewRateLimiterErrorWithDefaults() *RateLimiterError`

NewRateLimiterErrorWithDefaults instantiates a new RateLimiterError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *RateLimiterError) GetErrors() []GenericErrorErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RateLimiterError) GetErrorsOk() (*[]GenericErrorErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RateLimiterError) SetErrors(v []GenericErrorErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *RateLimiterError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


