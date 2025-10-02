# UpdateTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Updated name of the template. | [optional] 
**Contents** | Pointer to [**LanguageStringMap**](LanguageStringMap.md) |  | [optional] 
**IsEmail** | Pointer to **bool** | Set true for an Email template. | [optional] 
**EmailSubject** | Pointer to **NullableString** | Subject of the email. | [optional] 
**EmailBody** | Pointer to **NullableString** | Body of the email (HTML supported). | [optional] 
**IsSMS** | Pointer to **bool** | Set true for an SMS template. | [optional] 
**DynamicContent** | Pointer to **NullableString** | JSON string for dynamic content personalization. | [optional] 

## Methods

### NewUpdateTemplateRequest

`func NewUpdateTemplateRequest() *UpdateTemplateRequest`

NewUpdateTemplateRequest instantiates a new UpdateTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTemplateRequestWithDefaults

`func NewUpdateTemplateRequestWithDefaults() *UpdateTemplateRequest`

NewUpdateTemplateRequestWithDefaults instantiates a new UpdateTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContents

`func (o *UpdateTemplateRequest) GetContents() LanguageStringMap`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *UpdateTemplateRequest) GetContentsOk() (*LanguageStringMap, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *UpdateTemplateRequest) SetContents(v LanguageStringMap)`

SetContents sets Contents field to given value.

### HasContents

`func (o *UpdateTemplateRequest) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetIsEmail

`func (o *UpdateTemplateRequest) GetIsEmail() bool`

GetIsEmail returns the IsEmail field if non-nil, zero value otherwise.

### GetIsEmailOk

`func (o *UpdateTemplateRequest) GetIsEmailOk() (*bool, bool)`

GetIsEmailOk returns a tuple with the IsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmail

`func (o *UpdateTemplateRequest) SetIsEmail(v bool)`

SetIsEmail sets IsEmail field to given value.

### HasIsEmail

`func (o *UpdateTemplateRequest) HasIsEmail() bool`

HasIsEmail returns a boolean if a field has been set.

### GetEmailSubject

`func (o *UpdateTemplateRequest) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *UpdateTemplateRequest) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *UpdateTemplateRequest) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *UpdateTemplateRequest) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### SetEmailSubjectNil

`func (o *UpdateTemplateRequest) SetEmailSubjectNil(b bool)`

 SetEmailSubjectNil sets the value for EmailSubject to be an explicit nil

### UnsetEmailSubject
`func (o *UpdateTemplateRequest) UnsetEmailSubject()`

UnsetEmailSubject ensures that no value is present for EmailSubject, not even an explicit nil
### GetEmailBody

`func (o *UpdateTemplateRequest) GetEmailBody() string`

GetEmailBody returns the EmailBody field if non-nil, zero value otherwise.

### GetEmailBodyOk

`func (o *UpdateTemplateRequest) GetEmailBodyOk() (*string, bool)`

GetEmailBodyOk returns a tuple with the EmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailBody

`func (o *UpdateTemplateRequest) SetEmailBody(v string)`

SetEmailBody sets EmailBody field to given value.

### HasEmailBody

`func (o *UpdateTemplateRequest) HasEmailBody() bool`

HasEmailBody returns a boolean if a field has been set.

### SetEmailBodyNil

`func (o *UpdateTemplateRequest) SetEmailBodyNil(b bool)`

 SetEmailBodyNil sets the value for EmailBody to be an explicit nil

### UnsetEmailBody
`func (o *UpdateTemplateRequest) UnsetEmailBody()`

UnsetEmailBody ensures that no value is present for EmailBody, not even an explicit nil
### GetIsSMS

`func (o *UpdateTemplateRequest) GetIsSMS() bool`

GetIsSMS returns the IsSMS field if non-nil, zero value otherwise.

### GetIsSMSOk

`func (o *UpdateTemplateRequest) GetIsSMSOk() (*bool, bool)`

GetIsSMSOk returns a tuple with the IsSMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSMS

`func (o *UpdateTemplateRequest) SetIsSMS(v bool)`

SetIsSMS sets IsSMS field to given value.

### HasIsSMS

`func (o *UpdateTemplateRequest) HasIsSMS() bool`

HasIsSMS returns a boolean if a field has been set.

### GetDynamicContent

`func (o *UpdateTemplateRequest) GetDynamicContent() string`

GetDynamicContent returns the DynamicContent field if non-nil, zero value otherwise.

### GetDynamicContentOk

`func (o *UpdateTemplateRequest) GetDynamicContentOk() (*string, bool)`

GetDynamicContentOk returns a tuple with the DynamicContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicContent

`func (o *UpdateTemplateRequest) SetDynamicContent(v string)`

SetDynamicContent sets DynamicContent field to given value.

### HasDynamicContent

`func (o *UpdateTemplateRequest) HasDynamicContent() bool`

HasDynamicContent returns a boolean if a field has been set.

### SetDynamicContentNil

`func (o *UpdateTemplateRequest) SetDynamicContentNil(b bool)`

 SetDynamicContentNil sets the value for DynamicContent to be an explicit nil

### UnsetDynamicContent
`func (o *UpdateTemplateRequest) UnsetDynamicContent()`

UnsetDynamicContent ensures that no value is present for DynamicContent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


