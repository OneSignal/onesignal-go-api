# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**NotificationTypes** | Pointer to **int32** |  | [optional] 
**SessionTime** | Pointer to **int32** |  | [optional] 
**SessionCount** | Pointer to **int32** |  | [optional] 
**Sdk** | Pointer to **string** |  | [optional] 
**DeviceModel** | Pointer to **string** |  | [optional] 
**DeviceOs** | Pointer to **string** |  | [optional] 
**Rooted** | Pointer to **bool** |  | [optional] 
**TestType** | Pointer to **int32** |  | [optional] 
**AppVersion** | Pointer to **string** |  | [optional] 
**NetType** | Pointer to **int32** |  | [optional] 
**Carrier** | Pointer to **string** |  | [optional] 
**WebAuth** | Pointer to **string** |  | [optional] 
**WebP256** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Subscription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Subscription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Subscription) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Subscription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetToken

`func (o *Subscription) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Subscription) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Subscription) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Subscription) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetEnabled

`func (o *Subscription) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Subscription) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Subscription) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Subscription) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNotificationTypes

`func (o *Subscription) GetNotificationTypes() int32`

GetNotificationTypes returns the NotificationTypes field if non-nil, zero value otherwise.

### GetNotificationTypesOk

`func (o *Subscription) GetNotificationTypesOk() (*int32, bool)`

GetNotificationTypesOk returns a tuple with the NotificationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTypes

`func (o *Subscription) SetNotificationTypes(v int32)`

SetNotificationTypes sets NotificationTypes field to given value.

### HasNotificationTypes

`func (o *Subscription) HasNotificationTypes() bool`

HasNotificationTypes returns a boolean if a field has been set.

### GetSessionTime

`func (o *Subscription) GetSessionTime() int32`

GetSessionTime returns the SessionTime field if non-nil, zero value otherwise.

### GetSessionTimeOk

`func (o *Subscription) GetSessionTimeOk() (*int32, bool)`

GetSessionTimeOk returns a tuple with the SessionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTime

`func (o *Subscription) SetSessionTime(v int32)`

SetSessionTime sets SessionTime field to given value.

### HasSessionTime

`func (o *Subscription) HasSessionTime() bool`

HasSessionTime returns a boolean if a field has been set.

### GetSessionCount

`func (o *Subscription) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *Subscription) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *Subscription) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *Subscription) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetSdk

`func (o *Subscription) GetSdk() string`

GetSdk returns the Sdk field if non-nil, zero value otherwise.

### GetSdkOk

`func (o *Subscription) GetSdkOk() (*string, bool)`

GetSdkOk returns a tuple with the Sdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdk

`func (o *Subscription) SetSdk(v string)`

SetSdk sets Sdk field to given value.

### HasSdk

`func (o *Subscription) HasSdk() bool`

HasSdk returns a boolean if a field has been set.

### GetDeviceModel

`func (o *Subscription) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *Subscription) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *Subscription) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *Subscription) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceOs

`func (o *Subscription) GetDeviceOs() string`

GetDeviceOs returns the DeviceOs field if non-nil, zero value otherwise.

### GetDeviceOsOk

`func (o *Subscription) GetDeviceOsOk() (*string, bool)`

GetDeviceOsOk returns a tuple with the DeviceOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOs

`func (o *Subscription) SetDeviceOs(v string)`

SetDeviceOs sets DeviceOs field to given value.

### HasDeviceOs

`func (o *Subscription) HasDeviceOs() bool`

HasDeviceOs returns a boolean if a field has been set.

### GetRooted

`func (o *Subscription) GetRooted() bool`

GetRooted returns the Rooted field if non-nil, zero value otherwise.

### GetRootedOk

`func (o *Subscription) GetRootedOk() (*bool, bool)`

GetRootedOk returns a tuple with the Rooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooted

`func (o *Subscription) SetRooted(v bool)`

SetRooted sets Rooted field to given value.

### HasRooted

`func (o *Subscription) HasRooted() bool`

HasRooted returns a boolean if a field has been set.

### GetTestType

`func (o *Subscription) GetTestType() int32`

GetTestType returns the TestType field if non-nil, zero value otherwise.

### GetTestTypeOk

`func (o *Subscription) GetTestTypeOk() (*int32, bool)`

GetTestTypeOk returns a tuple with the TestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestType

`func (o *Subscription) SetTestType(v int32)`

SetTestType sets TestType field to given value.

### HasTestType

`func (o *Subscription) HasTestType() bool`

HasTestType returns a boolean if a field has been set.

### GetAppVersion

`func (o *Subscription) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *Subscription) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *Subscription) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *Subscription) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetNetType

`func (o *Subscription) GetNetType() int32`

GetNetType returns the NetType field if non-nil, zero value otherwise.

### GetNetTypeOk

`func (o *Subscription) GetNetTypeOk() (*int32, bool)`

GetNetTypeOk returns a tuple with the NetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetType

`func (o *Subscription) SetNetType(v int32)`

SetNetType sets NetType field to given value.

### HasNetType

`func (o *Subscription) HasNetType() bool`

HasNetType returns a boolean if a field has been set.

### GetCarrier

`func (o *Subscription) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *Subscription) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *Subscription) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *Subscription) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetWebAuth

`func (o *Subscription) GetWebAuth() string`

GetWebAuth returns the WebAuth field if non-nil, zero value otherwise.

### GetWebAuthOk

`func (o *Subscription) GetWebAuthOk() (*string, bool)`

GetWebAuthOk returns a tuple with the WebAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuth

`func (o *Subscription) SetWebAuth(v string)`

SetWebAuth sets WebAuth field to given value.

### HasWebAuth

`func (o *Subscription) HasWebAuth() bool`

HasWebAuth returns a boolean if a field has been set.

### GetWebP256

`func (o *Subscription) GetWebP256() string`

GetWebP256 returns the WebP256 field if non-nil, zero value otherwise.

### GetWebP256Ok

`func (o *Subscription) GetWebP256Ok() (*string, bool)`

GetWebP256Ok returns a tuple with the WebP256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebP256

`func (o *Subscription) SetWebP256(v string)`

SetWebP256 sets WebP256 field to given value.

### HasWebP256

`func (o *Subscription) HasWebP256() bool`

HasWebP256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


