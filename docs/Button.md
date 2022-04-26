# Button

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Text** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 

## Methods

### NewButton

`func NewButton(id string, ) *Button`

NewButton instantiates a new Button object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewButtonWithDefaults

`func NewButtonWithDefaults() *Button`

NewButtonWithDefaults instantiates a new Button object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Button) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Button) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Button) SetId(v string)`

SetId sets Id field to given value.


### GetText

`func (o *Button) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Button) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Button) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Button) HasText() bool`

HasText returns a boolean if a field has been set.

### GetIcon

`func (o *Button) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Button) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Button) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Button) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


