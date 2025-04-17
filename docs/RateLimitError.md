# RateLimitError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **string** |  | [optional] 

## Methods

### NewRateLimitError

`func NewRateLimitError() *RateLimitError`

NewRateLimitError instantiates a new RateLimitError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitErrorWithDefaults

`func NewRateLimitErrorWithDefaults() *RateLimitError`

NewRateLimitErrorWithDefaults instantiates a new RateLimitError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *RateLimitError) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RateLimitError) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RateLimitError) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *RateLimitError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLimit

`func (o *RateLimitError) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RateLimitError) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RateLimitError) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RateLimitError) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


