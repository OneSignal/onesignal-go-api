# Player

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The device&#39;s OneSignal ID | [readonly] 
**InvalidIdentifier** | Pointer to **bool** | If true, this is the equivalent of a user being Unsubscribed | [optional] [readonly] 
**AppId** | Pointer to **string** |  |  
**DeviceType** | **int32** | Required The device&#39;s platform:   0 &#x3D; iOS   1 &#x3D; Android   2 &#x3D; Amazon   3 &#x3D; WindowsPhone (MPNS)   4 &#x3D; Chrome Apps / Extensions   5 &#x3D; Chrome Web Push   6 &#x3D; Windows (WNS)   7 &#x3D; Safari   8 &#x3D; Firefox   9 &#x3D; MacOS   10 &#x3D; Alexa   11 &#x3D; Email   13 &#x3D; For Huawei App Gallery Builds SDK Setup. Not for Huawei Devices using FCM   14 &#x3D; SMS  | 
**ExternalUserId** | Pointer to **NullableString** | a custom user ID | [optional] 
**ExternalUserIdAuthHash** | Pointer to **string** | Only required if you have enabled Identity Verification and device_type is NOT 11 email. | [optional] 
**EmailAuthHash** | Pointer to **string** | Email - Only required if you have enabled Identity Verification and device_type is email (11). | [optional] 
**Identifier** | Pointer to **NullableString** | Recommended: For Push Notifications, this is the Push Token Identifier from Google or Apple. For Apple Push identifiers, you must strip all non alphanumeric characters. Examples: iOS: 7abcd558f29d0b1f048083e2834ad8ea4b3d87d8ad9c088b33c132706ff445f0 Android: APA91bHbYHk7aq-Uam_2pyJ2qbZvqllyyh2wjfPRaw5gLEX2SUlQBRvOc6sck1sa7H7nGeLNlDco8lXj83HWWwzV... For Email Addresses, set the full email address email@email.com and make sure to set device_type to 11.  | [optional] 
**Language** | Pointer to **string** | Language code. Typically lower case two letters, except for Chinese where it must be one of zh-Hans or zh-Hant. Example: en  | [optional] 
**Timezone** | Pointer to **NullableInt32** | Number of seconds away from UTC. Example: -28800  | [optional] 
**GameVersion** | Pointer to **NullableString** | Version of your app. Example: 1.1  | [optional] 
**DeviceModel** | Pointer to **NullableString** | Device make and model. Example: iPhone5,1  | [optional] 
**DeviceOs** | Pointer to **NullableString** | Device operating system version. Example: 7.0.4  | [optional] 
**AdId** | Pointer to **NullableString** | The ad id for the device&#39;s platform: Android &#x3D; Advertising Id iOS &#x3D; identifierForVendor WP8.0 &#x3D; DeviceUniqueId WP8.1 &#x3D; AdvertisingId  | [optional] 
**Sdk** | Pointer to **NullableString** | Name and version of the sdk/plugin that&#39;s calling this API method (if any) | [optional] 
**SessionCount** | Pointer to **int32** | Number of times the user has played the game, defaults to 1 | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Custom tags for the player. Only support string and integer key value pairs. Does not support arrays or other nested objects. Setting a tag value to null or an empty string will remove the tag. Example: {\&quot;foo\&quot;:\&quot;bar\&quot;,\&quot;this\&quot;:\&quot;that\&quot;} Limitations: - 100 tags per call - Android SDK users: tags cannot be removed or changed via API if set through SDK sendTag methods. Recommended to only tag devices with 1 kilobyte of data Please consider using your own Database to save more than 1 kilobyte of data. See: Internal Database &amp; CRM  | [optional] 
**AmountSpent** | Pointer to **float32** | Amount the user has spent in USD, up to two decimal places | [optional] 
**CreatedAt** | Pointer to **int64** | Unixtime when the player joined the game | [optional] 
**Playtime** | Pointer to **int64** | Seconds player was running your app. | [optional] 
**BadgeCount** | Pointer to **int32** | Current iOS badge count displayed on the app icon NOTE: Not supported for apps created after June 2018, since badge count for apps created after this date are handled on the client.  | [optional] 
**LastActive** | Pointer to **int32** | Unixtime when the player was last active | [optional] 
**NotificationTypes** | Pointer to **int32** | 1 &#x3D; subscribed -2 &#x3D; unsubscribed iOS - These values are set each time the user opens the app from the SDK. Use the SDK function set Subscription instead. Android - You may set this but you can no longer use the SDK method setSubscription later in your app as it will create synchronization issues.  | [optional] 
**TestType** | Pointer to **NullableInt32** | This is used in deciding whether to use your iOS Sandbox or Production push certificate when sending a push when both have been uploaded. Set to the iOS provisioning profile that was used to build your app. 1 &#x3D; Development 2 &#x3D; Ad-Hoc Omit this field for App Store builds.  | [optional] 
**Long** | Pointer to **float32** | Longitude of the device, used for geotagging to segment on. | [optional] 
**Lat** | Pointer to **float32** | Latitude of the device, used for geotagging to segment on. | [optional] 
**Country** | Pointer to **string** | Country code in the ISO 3166-1 Alpha 2 format | [optional] 

## Methods

### NewPlayer

`func NewPlayer(id string, deviceType int32, appId *string) *Player`

NewPlayer instantiates a new Player object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerWithDefaults

`func NewPlayerWithDefaults() *Player`

NewPlayerWithDefaults instantiates a new Player object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Player) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Player) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Player) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Player) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvalidIdentifier

`func (o *Player) GetInvalidIdentifier() bool`

GetInvalidIdentifier returns the InvalidIdentifier field if non-nil, zero value otherwise.

### GetInvalidIdentifierOk

`func (o *Player) GetInvalidIdentifierOk() (*bool, bool)`

GetInvalidIdentifierOk returns a tuple with the InvalidIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidIdentifier

`func (o *Player) SetInvalidIdentifier(v bool)`

SetInvalidIdentifier sets InvalidIdentifier field to given value.

### HasInvalidIdentifier

`func (o *Player) HasInvalidIdentifier() bool`

HasInvalidIdentifier returns a boolean if a field has been set.

### GetAppId

`func (o *Player) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Player) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Player) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Player) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDeviceType

`func (o *Player) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *Player) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *Player) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.


### GetExternalUserId

`func (o *Player) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *Player) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *Player) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *Player) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### SetExternalUserIdNil

`func (o *Player) SetExternalUserIdNil(b bool)`

 SetExternalUserIdNil sets the value for ExternalUserId to be an explicit nil

### UnsetExternalUserId
`func (o *Player) UnsetExternalUserId()`

UnsetExternalUserId ensures that no value is present for ExternalUserId, not even an explicit nil
### GetExternalUserIdAuthHash

`func (o *Player) GetExternalUserIdAuthHash() string`

GetExternalUserIdAuthHash returns the ExternalUserIdAuthHash field if non-nil, zero value otherwise.

### GetExternalUserIdAuthHashOk

`func (o *Player) GetExternalUserIdAuthHashOk() (*string, bool)`

GetExternalUserIdAuthHashOk returns a tuple with the ExternalUserIdAuthHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserIdAuthHash

`func (o *Player) SetExternalUserIdAuthHash(v string)`

SetExternalUserIdAuthHash sets ExternalUserIdAuthHash field to given value.

### HasExternalUserIdAuthHash

`func (o *Player) HasExternalUserIdAuthHash() bool`

HasExternalUserIdAuthHash returns a boolean if a field has been set.

### GetEmailAuthHash

`func (o *Player) GetEmailAuthHash() string`

GetEmailAuthHash returns the EmailAuthHash field if non-nil, zero value otherwise.

### GetEmailAuthHashOk

`func (o *Player) GetEmailAuthHashOk() (*string, bool)`

GetEmailAuthHashOk returns a tuple with the EmailAuthHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAuthHash

`func (o *Player) SetEmailAuthHash(v string)`

SetEmailAuthHash sets EmailAuthHash field to given value.

### HasEmailAuthHash

`func (o *Player) HasEmailAuthHash() bool`

HasEmailAuthHash returns a boolean if a field has been set.

### GetIdentifier

`func (o *Player) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Player) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Player) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Player) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *Player) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *Player) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetLanguage

`func (o *Player) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Player) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Player) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Player) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetTimezone

`func (o *Player) GetTimezone() int32`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Player) GetTimezoneOk() (*int32, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Player) SetTimezone(v int32)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Player) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *Player) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *Player) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetGameVersion

`func (o *Player) GetGameVersion() string`

GetGameVersion returns the GameVersion field if non-nil, zero value otherwise.

### GetGameVersionOk

`func (o *Player) GetGameVersionOk() (*string, bool)`

GetGameVersionOk returns a tuple with the GameVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameVersion

`func (o *Player) SetGameVersion(v string)`

SetGameVersion sets GameVersion field to given value.

### HasGameVersion

`func (o *Player) HasGameVersion() bool`

HasGameVersion returns a boolean if a field has been set.

### SetGameVersionNil

`func (o *Player) SetGameVersionNil(b bool)`

 SetGameVersionNil sets the value for GameVersion to be an explicit nil

### UnsetGameVersion
`func (o *Player) UnsetGameVersion()`

UnsetGameVersion ensures that no value is present for GameVersion, not even an explicit nil
### GetDeviceModel

`func (o *Player) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *Player) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *Player) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *Player) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### SetDeviceModelNil

`func (o *Player) SetDeviceModelNil(b bool)`

 SetDeviceModelNil sets the value for DeviceModel to be an explicit nil

### UnsetDeviceModel
`func (o *Player) UnsetDeviceModel()`

UnsetDeviceModel ensures that no value is present for DeviceModel, not even an explicit nil
### GetDeviceOs

`func (o *Player) GetDeviceOs() string`

GetDeviceOs returns the DeviceOs field if non-nil, zero value otherwise.

### GetDeviceOsOk

`func (o *Player) GetDeviceOsOk() (*string, bool)`

GetDeviceOsOk returns a tuple with the DeviceOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOs

`func (o *Player) SetDeviceOs(v string)`

SetDeviceOs sets DeviceOs field to given value.

### HasDeviceOs

`func (o *Player) HasDeviceOs() bool`

HasDeviceOs returns a boolean if a field has been set.

### SetDeviceOsNil

`func (o *Player) SetDeviceOsNil(b bool)`

 SetDeviceOsNil sets the value for DeviceOs to be an explicit nil

### UnsetDeviceOs
`func (o *Player) UnsetDeviceOs()`

UnsetDeviceOs ensures that no value is present for DeviceOs, not even an explicit nil
### GetAdId

`func (o *Player) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *Player) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *Player) SetAdId(v string)`

SetAdId sets AdId field to given value.

### HasAdId

`func (o *Player) HasAdId() bool`

HasAdId returns a boolean if a field has been set.

### SetAdIdNil

`func (o *Player) SetAdIdNil(b bool)`

 SetAdIdNil sets the value for AdId to be an explicit nil

### UnsetAdId
`func (o *Player) UnsetAdId()`

UnsetAdId ensures that no value is present for AdId, not even an explicit nil
### GetSdk

`func (o *Player) GetSdk() string`

GetSdk returns the Sdk field if non-nil, zero value otherwise.

### GetSdkOk

`func (o *Player) GetSdkOk() (*string, bool)`

GetSdkOk returns a tuple with the Sdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdk

`func (o *Player) SetSdk(v string)`

SetSdk sets Sdk field to given value.

### HasSdk

`func (o *Player) HasSdk() bool`

HasSdk returns a boolean if a field has been set.

### SetSdkNil

`func (o *Player) SetSdkNil(b bool)`

 SetSdkNil sets the value for Sdk to be an explicit nil

### UnsetSdk
`func (o *Player) UnsetSdk()`

UnsetSdk ensures that no value is present for Sdk, not even an explicit nil
### GetSessionCount

`func (o *Player) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *Player) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *Player) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *Player) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetTags

`func (o *Player) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Player) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Player) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Player) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Player) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Player) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAmountSpent

`func (o *Player) GetAmountSpent() float32`

GetAmountSpent returns the AmountSpent field if non-nil, zero value otherwise.

### GetAmountSpentOk

`func (o *Player) GetAmountSpentOk() (*float32, bool)`

GetAmountSpentOk returns a tuple with the AmountSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSpent

`func (o *Player) SetAmountSpent(v float32)`

SetAmountSpent sets AmountSpent field to given value.

### HasAmountSpent

`func (o *Player) HasAmountSpent() bool`

HasAmountSpent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Player) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Player) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Player) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Player) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPlaytime

`func (o *Player) GetPlaytime() int64`

GetPlaytime returns the Playtime field if non-nil, zero value otherwise.

### GetPlaytimeOk

`func (o *Player) GetPlaytimeOk() (*int64, bool)`

GetPlaytimeOk returns a tuple with the Playtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaytime

`func (o *Player) SetPlaytime(v int64)`

SetPlaytime sets Playtime field to given value.

### HasPlaytime

`func (o *Player) HasPlaytime() bool`

HasPlaytime returns a boolean if a field has been set.

### GetBadgeCount

`func (o *Player) GetBadgeCount() int32`

GetBadgeCount returns the BadgeCount field if non-nil, zero value otherwise.

### GetBadgeCountOk

`func (o *Player) GetBadgeCountOk() (*int32, bool)`

GetBadgeCountOk returns a tuple with the BadgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeCount

`func (o *Player) SetBadgeCount(v int32)`

SetBadgeCount sets BadgeCount field to given value.

### HasBadgeCount

`func (o *Player) HasBadgeCount() bool`

HasBadgeCount returns a boolean if a field has been set.

### GetLastActive

`func (o *Player) GetLastActive() int32`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *Player) GetLastActiveOk() (*int32, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *Player) SetLastActive(v int32)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *Player) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetNotificationTypes

`func (o *Player) GetNotificationTypes() int32`

GetNotificationTypes returns the NotificationTypes field if non-nil, zero value otherwise.

### GetNotificationTypesOk

`func (o *Player) GetNotificationTypesOk() (*int32, bool)`

GetNotificationTypesOk returns a tuple with the NotificationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTypes

`func (o *Player) SetNotificationTypes(v int32)`

SetNotificationTypes sets NotificationTypes field to given value.

### HasNotificationTypes

`func (o *Player) HasNotificationTypes() bool`

HasNotificationTypes returns a boolean if a field has been set.

### GetTestType

`func (o *Player) GetTestType() int32`

GetTestType returns the TestType field if non-nil, zero value otherwise.

### GetTestTypeOk

`func (o *Player) GetTestTypeOk() (*int32, bool)`

GetTestTypeOk returns a tuple with the TestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestType

`func (o *Player) SetTestType(v int32)`

SetTestType sets TestType field to given value.

### HasTestType

`func (o *Player) HasTestType() bool`

HasTestType returns a boolean if a field has been set.

### SetTestTypeNil

`func (o *Player) SetTestTypeNil(b bool)`

 SetTestTypeNil sets the value for TestType to be an explicit nil

### UnsetTestType
`func (o *Player) UnsetTestType()`

UnsetTestType ensures that no value is present for TestType, not even an explicit nil
### GetLong

`func (o *Player) GetLong() float32`

GetLong returns the Long field if non-nil, zero value otherwise.

### GetLongOk

`func (o *Player) GetLongOk() (*float32, bool)`

GetLongOk returns a tuple with the Long field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLong

`func (o *Player) SetLong(v float32)`

SetLong sets Long field to given value.

### HasLong

`func (o *Player) HasLong() bool`

HasLong returns a boolean if a field has been set.

### GetLat

`func (o *Player) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *Player) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *Player) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *Player) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetCountry

`func (o *Player) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Player) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Player) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Player) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


