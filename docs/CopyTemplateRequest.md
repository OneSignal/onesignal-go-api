# CopyTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAppId** | **string** | Destination OneSignal App ID in UUID v4 format. | 

## Methods

### NewCopyTemplateRequest

`func NewCopyTemplateRequest(targetAppId string, ) *CopyTemplateRequest`

NewCopyTemplateRequest instantiates a new CopyTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyTemplateRequestWithDefaults

`func NewCopyTemplateRequestWithDefaults() *CopyTemplateRequest`

NewCopyTemplateRequestWithDefaults instantiates a new CopyTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAppId

`func (o *CopyTemplateRequest) GetTargetAppId() string`

GetTargetAppId returns the TargetAppId field if non-nil, zero value otherwise.

### GetTargetAppIdOk

`func (o *CopyTemplateRequest) GetTargetAppIdOk() (*string, bool)`

GetTargetAppIdOk returns a tuple with the TargetAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppId

`func (o *CopyTemplateRequest) SetTargetAppId(v string)`

SetTargetAppId sets TargetAppId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


