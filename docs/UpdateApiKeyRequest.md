# UpdateApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**IpAllowlistMode** | Pointer to **string** |  | [optional] 
**IpAllowlist** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateApiKeyRequest

`func NewUpdateApiKeyRequest() *UpdateApiKeyRequest`

NewUpdateApiKeyRequest instantiates a new UpdateApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApiKeyRequestWithDefaults

`func NewUpdateApiKeyRequestWithDefaults() *UpdateApiKeyRequest`

NewUpdateApiKeyRequestWithDefaults instantiates a new UpdateApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateApiKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApiKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApiKeyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateApiKeyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIpAllowlistMode

`func (o *UpdateApiKeyRequest) GetIpAllowlistMode() string`

GetIpAllowlistMode returns the IpAllowlistMode field if non-nil, zero value otherwise.

### GetIpAllowlistModeOk

`func (o *UpdateApiKeyRequest) GetIpAllowlistModeOk() (*string, bool)`

GetIpAllowlistModeOk returns a tuple with the IpAllowlistMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlistMode

`func (o *UpdateApiKeyRequest) SetIpAllowlistMode(v string)`

SetIpAllowlistMode sets IpAllowlistMode field to given value.

### HasIpAllowlistMode

`func (o *UpdateApiKeyRequest) HasIpAllowlistMode() bool`

HasIpAllowlistMode returns a boolean if a field has been set.

### GetIpAllowlist

`func (o *UpdateApiKeyRequest) GetIpAllowlist() []string`

GetIpAllowlist returns the IpAllowlist field if non-nil, zero value otherwise.

### GetIpAllowlistOk

`func (o *UpdateApiKeyRequest) GetIpAllowlistOk() (*[]string, bool)`

GetIpAllowlistOk returns a tuple with the IpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlist

`func (o *UpdateApiKeyRequest) SetIpAllowlist(v []string)`

SetIpAllowlist sets IpAllowlist field to given value.

### HasIpAllowlist

`func (o *UpdateApiKeyRequest) HasIpAllowlist() bool`

HasIpAllowlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


