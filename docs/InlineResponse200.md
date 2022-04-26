# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Recipients** | **int32** | Estimated number of subscribers targetted by notification. | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**Notification200Errors**](Notification200Errors.md) |  | [optional] 

## Methods

### NewInlineResponse200

`func NewInlineResponse200(id string, recipients int32, ) *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200) SetId(v string)`

SetId sets Id field to given value.


### GetRecipients

`func (o *InlineResponse200) GetRecipients() int32`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineResponse200) GetRecipientsOk() (*int32, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineResponse200) SetRecipients(v int32)`

SetRecipients sets Recipients field to given value.


### GetExternalId

`func (o *InlineResponse200) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse200) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse200) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse200) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200) GetErrors() Notification200Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200) GetErrorsOk() (*Notification200Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200) SetErrors(v Notification200Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


