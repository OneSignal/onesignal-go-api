# SubscriptionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**NotificationTypes** | Pointer to **float32** |  | [optional] 
**SessionTime** | Pointer to **float32** |  | [optional] 
**SessionCount** | Pointer to **float32** |  | [optional] 
**Sdk** | Pointer to **string** |  | [optional] 
**DeviceModel** | Pointer to **string** |  | [optional] 
**DeviceOs** | Pointer to **string** |  | [optional] 
**Rooted** | Pointer to **bool** |  | [optional] 
**TestType** | Pointer to **float32** |  | [optional] 
**AppVersion** | Pointer to **string** |  | [optional] 
**NetType** | Pointer to **float32** |  | [optional] 
**Carrier** | Pointer to **string** |  | [optional] 
**WebAuth** | Pointer to **string** |  | [optional] 
**WebP256** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionObject

`func NewSubscriptionObject() *SubscriptionObject`

NewSubscriptionObject instantiates a new SubscriptionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionObjectWithDefaults

`func NewSubscriptionObjectWithDefaults() *SubscriptionObject`

NewSubscriptionObjectWithDefaults instantiates a new SubscriptionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SubscriptionObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetToken

`func (o *SubscriptionObject) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SubscriptionObject) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SubscriptionObject) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SubscriptionObject) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetEnabled

`func (o *SubscriptionObject) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SubscriptionObject) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SubscriptionObject) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SubscriptionObject) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNotificationTypes

`func (o *SubscriptionObject) GetNotificationTypes() float32`

GetNotificationTypes returns the NotificationTypes field if non-nil, zero value otherwise.

### GetNotificationTypesOk

`func (o *SubscriptionObject) GetNotificationTypesOk() (*float32, bool)`

GetNotificationTypesOk returns a tuple with the NotificationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTypes

`func (o *SubscriptionObject) SetNotificationTypes(v float32)`

SetNotificationTypes sets NotificationTypes field to given value.

### HasNotificationTypes

`func (o *SubscriptionObject) HasNotificationTypes() bool`

HasNotificationTypes returns a boolean if a field has been set.

### GetSessionTime

`func (o *SubscriptionObject) GetSessionTime() float32`

GetSessionTime returns the SessionTime field if non-nil, zero value otherwise.

### GetSessionTimeOk

`func (o *SubscriptionObject) GetSessionTimeOk() (*float32, bool)`

GetSessionTimeOk returns a tuple with the SessionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTime

`func (o *SubscriptionObject) SetSessionTime(v float32)`

SetSessionTime sets SessionTime field to given value.

### HasSessionTime

`func (o *SubscriptionObject) HasSessionTime() bool`

HasSessionTime returns a boolean if a field has been set.

### GetSessionCount

`func (o *SubscriptionObject) GetSessionCount() float32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *SubscriptionObject) GetSessionCountOk() (*float32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *SubscriptionObject) SetSessionCount(v float32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *SubscriptionObject) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetSdk

`func (o *SubscriptionObject) GetSdk() string`

GetSdk returns the Sdk field if non-nil, zero value otherwise.

### GetSdkOk

`func (o *SubscriptionObject) GetSdkOk() (*string, bool)`

GetSdkOk returns a tuple with the Sdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdk

`func (o *SubscriptionObject) SetSdk(v string)`

SetSdk sets Sdk field to given value.

### HasSdk

`func (o *SubscriptionObject) HasSdk() bool`

HasSdk returns a boolean if a field has been set.

### GetDeviceModel

`func (o *SubscriptionObject) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *SubscriptionObject) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *SubscriptionObject) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *SubscriptionObject) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceOs

`func (o *SubscriptionObject) GetDeviceOs() string`

GetDeviceOs returns the DeviceOs field if non-nil, zero value otherwise.

### GetDeviceOsOk

`func (o *SubscriptionObject) GetDeviceOsOk() (*string, bool)`

GetDeviceOsOk returns a tuple with the DeviceOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOs

`func (o *SubscriptionObject) SetDeviceOs(v string)`

SetDeviceOs sets DeviceOs field to given value.

### HasDeviceOs

`func (o *SubscriptionObject) HasDeviceOs() bool`

HasDeviceOs returns a boolean if a field has been set.

### GetRooted

`func (o *SubscriptionObject) GetRooted() bool`

GetRooted returns the Rooted field if non-nil, zero value otherwise.

### GetRootedOk

`func (o *SubscriptionObject) GetRootedOk() (*bool, bool)`

GetRootedOk returns a tuple with the Rooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooted

`func (o *SubscriptionObject) SetRooted(v bool)`

SetRooted sets Rooted field to given value.

### HasRooted

`func (o *SubscriptionObject) HasRooted() bool`

HasRooted returns a boolean if a field has been set.

### GetTestType

`func (o *SubscriptionObject) GetTestType() float32`

GetTestType returns the TestType field if non-nil, zero value otherwise.

### GetTestTypeOk

`func (o *SubscriptionObject) GetTestTypeOk() (*float32, bool)`

GetTestTypeOk returns a tuple with the TestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestType

`func (o *SubscriptionObject) SetTestType(v float32)`

SetTestType sets TestType field to given value.

### HasTestType

`func (o *SubscriptionObject) HasTestType() bool`

HasTestType returns a boolean if a field has been set.

### GetAppVersion

`func (o *SubscriptionObject) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *SubscriptionObject) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *SubscriptionObject) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *SubscriptionObject) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetNetType

`func (o *SubscriptionObject) GetNetType() float32`

GetNetType returns the NetType field if non-nil, zero value otherwise.

### GetNetTypeOk

`func (o *SubscriptionObject) GetNetTypeOk() (*float32, bool)`

GetNetTypeOk returns a tuple with the NetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetType

`func (o *SubscriptionObject) SetNetType(v float32)`

SetNetType sets NetType field to given value.

### HasNetType

`func (o *SubscriptionObject) HasNetType() bool`

HasNetType returns a boolean if a field has been set.

### GetCarrier

`func (o *SubscriptionObject) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *SubscriptionObject) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *SubscriptionObject) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *SubscriptionObject) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetWebAuth

`func (o *SubscriptionObject) GetWebAuth() string`

GetWebAuth returns the WebAuth field if non-nil, zero value otherwise.

### GetWebAuthOk

`func (o *SubscriptionObject) GetWebAuthOk() (*string, bool)`

GetWebAuthOk returns a tuple with the WebAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuth

`func (o *SubscriptionObject) SetWebAuth(v string)`

SetWebAuth sets WebAuth field to given value.

### HasWebAuth

`func (o *SubscriptionObject) HasWebAuth() bool`

HasWebAuth returns a boolean if a field has been set.

### GetWebP256

`func (o *SubscriptionObject) GetWebP256() string`

GetWebP256 returns the WebP256 field if non-nil, zero value otherwise.

### GetWebP256Ok

`func (o *SubscriptionObject) GetWebP256Ok() (*string, bool)`

GetWebP256Ok returns a tuple with the WebP256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebP256

`func (o *SubscriptionObject) SetWebP256(v string)`

SetWebP256 sets WebP256 field to given value.

### HasWebP256

`func (o *SubscriptionObject) HasWebP256() bool`

HasWebP256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


