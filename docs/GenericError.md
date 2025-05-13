# GenericError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **interface{}** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Reference** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGenericError

`func NewGenericError() *GenericError`

NewGenericError instantiates a new GenericError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericErrorWithDefaults

`func NewGenericErrorWithDefaults() *GenericError`

NewGenericErrorWithDefaults instantiates a new GenericError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *GenericError) GetErrors() interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GenericError) GetErrorsOk() (*interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GenericError) SetErrors(v interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GenericError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *GenericError) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *GenericError) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetSuccess

`func (o *GenericError) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GenericError) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GenericError) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GenericError) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetReference

`func (o *GenericError) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GenericError) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GenericError) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GenericError) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GenericError) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GenericError) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


