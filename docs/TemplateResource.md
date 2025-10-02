# TemplateResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Channel** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **map[string]interface{}** | Rendered content and channel/platform flags for the template. | [optional] 

## Methods

### NewTemplateResource

`func NewTemplateResource() *TemplateResource`

NewTemplateResource instantiates a new TemplateResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateResourceWithDefaults

`func NewTemplateResourceWithDefaults() *TemplateResource`

NewTemplateResourceWithDefaults instantiates a new TemplateResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TemplateResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TemplateResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TemplateResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TemplateResource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateResource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateResource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TemplateResource) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetChannel

`func (o *TemplateResource) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TemplateResource) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TemplateResource) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TemplateResource) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *TemplateResource) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *TemplateResource) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetContent

`func (o *TemplateResource) GetContent() map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TemplateResource) GetContentOk() (*map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TemplateResource) SetContent(v map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *TemplateResource) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


