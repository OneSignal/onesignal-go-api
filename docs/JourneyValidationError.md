# JourneyValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Stable, kebab-case error identifier. Always &#x60;invalid-payload&#x60; for validation failures. | [optional] 
**Title** | Pointer to **string** | Human-readable message. Wording may change between releases. | [optional] 
**Meta** | Pointer to **map[string]interface{}** | Structured context. Always includes &#x60;attribute&#x60; (the field, or &#x60;base&#x60; for request-level errors); may also include &#x60;path&#x60;, &#x60;node_id&#x60;, or &#x60;client_node_id&#x60;. | [optional] 

## Methods

### NewJourneyValidationError

`func NewJourneyValidationError() *JourneyValidationError`

NewJourneyValidationError instantiates a new JourneyValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyValidationErrorWithDefaults

`func NewJourneyValidationErrorWithDefaults() *JourneyValidationError`

NewJourneyValidationErrorWithDefaults instantiates a new JourneyValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *JourneyValidationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *JourneyValidationError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *JourneyValidationError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *JourneyValidationError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTitle

`func (o *JourneyValidationError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *JourneyValidationError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *JourneyValidationError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *JourneyValidationError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMeta

`func (o *JourneyValidationError) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JourneyValidationError) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JourneyValidationError) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JourneyValidationError) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


