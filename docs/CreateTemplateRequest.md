# CreateTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | Your OneSignal App ID in UUID v4 format. | 
**Name** | **string** | Name of the template. | 
**Contents** | [**LanguageStringMap**](LanguageStringMap.md) |  | 
**IsEmail** | Pointer to **bool** | Set true for an Email template. | [optional] 
**EmailSubject** | Pointer to **NullableString** | Subject of the email. | [optional] 
**EmailBody** | Pointer to **NullableString** | Body of the email (HTML supported). | [optional] 
**IsSMS** | Pointer to **bool** | Set true for an SMS template. | [optional] 
**DynamicContent** | Pointer to **NullableString** | JSON string for dynamic content personalization. | [optional] 

## Methods

### NewCreateTemplateRequest

`func NewCreateTemplateRequest(appId string, name string, contents LanguageStringMap, ) *CreateTemplateRequest`

NewCreateTemplateRequest instantiates a new CreateTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTemplateRequestWithDefaults

`func NewCreateTemplateRequestWithDefaults() *CreateTemplateRequest`

NewCreateTemplateRequestWithDefaults instantiates a new CreateTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *CreateTemplateRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateTemplateRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateTemplateRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetName

`func (o *CreateTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetContents

`func (o *CreateTemplateRequest) GetContents() LanguageStringMap`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *CreateTemplateRequest) GetContentsOk() (*LanguageStringMap, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *CreateTemplateRequest) SetContents(v LanguageStringMap)`

SetContents sets Contents field to given value.


### GetIsEmail

`func (o *CreateTemplateRequest) GetIsEmail() bool`

GetIsEmail returns the IsEmail field if non-nil, zero value otherwise.

### GetIsEmailOk

`func (o *CreateTemplateRequest) GetIsEmailOk() (*bool, bool)`

GetIsEmailOk returns a tuple with the IsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmail

`func (o *CreateTemplateRequest) SetIsEmail(v bool)`

SetIsEmail sets IsEmail field to given value.

### HasIsEmail

`func (o *CreateTemplateRequest) HasIsEmail() bool`

HasIsEmail returns a boolean if a field has been set.

### GetEmailSubject

`func (o *CreateTemplateRequest) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *CreateTemplateRequest) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *CreateTemplateRequest) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *CreateTemplateRequest) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### SetEmailSubjectNil

`func (o *CreateTemplateRequest) SetEmailSubjectNil(b bool)`

 SetEmailSubjectNil sets the value for EmailSubject to be an explicit nil

### UnsetEmailSubject
`func (o *CreateTemplateRequest) UnsetEmailSubject()`

UnsetEmailSubject ensures that no value is present for EmailSubject, not even an explicit nil
### GetEmailBody

`func (o *CreateTemplateRequest) GetEmailBody() string`

GetEmailBody returns the EmailBody field if non-nil, zero value otherwise.

### GetEmailBodyOk

`func (o *CreateTemplateRequest) GetEmailBodyOk() (*string, bool)`

GetEmailBodyOk returns a tuple with the EmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailBody

`func (o *CreateTemplateRequest) SetEmailBody(v string)`

SetEmailBody sets EmailBody field to given value.

### HasEmailBody

`func (o *CreateTemplateRequest) HasEmailBody() bool`

HasEmailBody returns a boolean if a field has been set.

### SetEmailBodyNil

`func (o *CreateTemplateRequest) SetEmailBodyNil(b bool)`

 SetEmailBodyNil sets the value for EmailBody to be an explicit nil

### UnsetEmailBody
`func (o *CreateTemplateRequest) UnsetEmailBody()`

UnsetEmailBody ensures that no value is present for EmailBody, not even an explicit nil
### GetIsSMS

`func (o *CreateTemplateRequest) GetIsSMS() bool`

GetIsSMS returns the IsSMS field if non-nil, zero value otherwise.

### GetIsSMSOk

`func (o *CreateTemplateRequest) GetIsSMSOk() (*bool, bool)`

GetIsSMSOk returns a tuple with the IsSMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSMS

`func (o *CreateTemplateRequest) SetIsSMS(v bool)`

SetIsSMS sets IsSMS field to given value.

### HasIsSMS

`func (o *CreateTemplateRequest) HasIsSMS() bool`

HasIsSMS returns a boolean if a field has been set.

### GetDynamicContent

`func (o *CreateTemplateRequest) GetDynamicContent() string`

GetDynamicContent returns the DynamicContent field if non-nil, zero value otherwise.

### GetDynamicContentOk

`func (o *CreateTemplateRequest) GetDynamicContentOk() (*string, bool)`

GetDynamicContentOk returns a tuple with the DynamicContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicContent

`func (o *CreateTemplateRequest) SetDynamicContent(v string)`

SetDynamicContent sets DynamicContent field to given value.

### HasDynamicContent

`func (o *CreateTemplateRequest) HasDynamicContent() bool`

HasDynamicContent returns a boolean if a field has been set.

### SetDynamicContentNil

`func (o *CreateTemplateRequest) SetDynamicContentNil(b bool)`

 SetDynamicContentNil sets the value for DynamicContent to be an explicit nil

### UnsetDynamicContent
`func (o *CreateTemplateRequest) UnsetDynamicContent()`

UnsetDynamicContent ensures that no value is present for DynamicContent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


