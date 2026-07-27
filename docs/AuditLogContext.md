# AuditLogContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Country code derived from the request IP. | [optional] 
**Ip** | Pointer to **string** | IP address the request originated from. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Additional context-specific data. | [optional] 
**UserAgent** | Pointer to **string** | User agent of the client that made the request. | [optional] 

## Methods

### NewAuditLogContext

`func NewAuditLogContext() *AuditLogContext`

NewAuditLogContext instantiates a new AuditLogContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogContextWithDefaults

`func NewAuditLogContextWithDefaults() *AuditLogContext`

NewAuditLogContextWithDefaults instantiates a new AuditLogContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *AuditLogContext) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AuditLogContext) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AuditLogContext) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AuditLogContext) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetIp

`func (o *AuditLogContext) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AuditLogContext) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AuditLogContext) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AuditLogContext) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMetadata

`func (o *AuditLogContext) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AuditLogContext) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AuditLogContext) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AuditLogContext) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUserAgent

`func (o *AuditLogContext) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AuditLogContext) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AuditLogContext) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AuditLogContext) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


