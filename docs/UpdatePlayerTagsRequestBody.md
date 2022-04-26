# UpdatePlayerTagsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string]interface{}** | Custom tags for the device record.  Only support string key value pairs.  Does not support arrays or other nested objects.  Example &#x60;{\&quot;foo\&quot;:\&quot;bar\&quot;,\&quot;this\&quot;:\&quot;that\&quot;}&#x60;. Limitations: - 100 tags per call - Android SDK users: tags cannot be removed or changed via API if set through SDK sendTag methods. Recommended to only tag devices with 1 kilobyte of ata Please consider using your own Database to save more than 1 kilobyte of data.  See: Internal Database &amp; CRM  | [optional] 

## Methods

### NewUpdatePlayerTagsRequestBody

`func NewUpdatePlayerTagsRequestBody() *UpdatePlayerTagsRequestBody`

NewUpdatePlayerTagsRequestBody instantiates a new UpdatePlayerTagsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePlayerTagsRequestBodyWithDefaults

`func NewUpdatePlayerTagsRequestBodyWithDefaults() *UpdatePlayerTagsRequestBody`

NewUpdatePlayerTagsRequestBodyWithDefaults instantiates a new UpdatePlayerTagsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *UpdatePlayerTagsRequestBody) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdatePlayerTagsRequestBody) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdatePlayerTagsRequestBody) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdatePlayerTagsRequestBody) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


