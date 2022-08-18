# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | Pointer to **string** | The name of your app, as displayed on your apps list on the dashboard.  This can be renamed. | [optional] 
**Players** | Pointer to **int32** |  | [optional] [readonly] 
**MessageablePlayers** | Pointer to **int32** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**AndroidGcmSenderId** | Pointer to **string** | Android: Your Google Project number.  Also known as Sender ID. | [optional] 
**GcmKey** | Pointer to **NullableString** | Android: Your Google Push Messaging Auth Key | [optional] 
**ChromeWebOrigin** | Pointer to **NullableString** | Chrome (All Browsers except Safari) (Recommended): The URL to your website.  This field is required if you wish to enable web push and specify other web push parameters. | [optional] 
**ChromeKey** | Pointer to **NullableString** | Not for web push.  Your Google Push Messaging Auth Key if you use Chrome Apps / Extensions. | [optional] 
**ChromeWebDefaultNotificationIcon** | Pointer to **NullableString** | Chrome (All Browsers except Safari): Your default notification icon. Should be 256x256 pixels, min 80x80. | [optional] 
**ChromeWebSubDomain** | Pointer to **NullableString** | Chrome (All Browsers except Safari): A subdomain of your choice in order to support Web Push on non-HTTPS websites. This field must be set in order for the chrome_web_gcm_sender_id property to be processed. | [optional] 
**ApnsEnv** | Pointer to **NullableString** | iOS: Either sandbox or production | [optional] 
**ApnsP12** | Pointer to **string** | iOS: Your apple push notification p12 certificate file, converted to a string and Base64 encoded. | [optional] 
**ApnsP12Password** | Pointer to **string** | iOS: Required if using p12 certificate.  Password for the apns_p12 file. | [optional] 
**ApnsCertificates** | Pointer to **NullableString** |  | [optional] [readonly] 
**SafariApnsCertificates** | Pointer to **string** |  | [optional] [readonly] 
**SafariApnsP12** | Pointer to **string** | Safari: Your apple push notification p12 certificate file for Safari Push Notifications, converted to a string and Base64 encoded. | [optional] 
**SafariApnsP12Password** | Pointer to **string** | Safari: Password for safari_apns_p12 file | [optional] 
**SafariSiteOrigin** | Pointer to **NullableString** | Safari (Recommended): The hostname to your website including http(s):// | [optional] 
**SafariPushId** | Pointer to **NullableString** |  | [optional] [readonly] 
**SafariIcon1616** | Pointer to **string** |  | [optional] [readonly] 
**SafariIcon3232** | Pointer to **string** |  | [optional] [readonly] 
**SafariIcon6464** | Pointer to **string** |  | [optional] [readonly] 
**SafariIcon128128** | Pointer to **string** |  | [optional] [readonly] 
**SafariIcon256256** | Pointer to **string** | Safari: A url for a 256x256 png notification icon. This is the only Safari icon URL you need to provide. | [optional] 
**SiteName** | Pointer to **NullableString** | All Browsers (Recommended): The Site Name. Requires both chrome_web_origin and safari_site_origin to be set to add or update it. | [optional] 
**BasicAuthKey** | Pointer to **string** |  | [optional] [readonly] 
**OrganizationId** | Pointer to **string** | The Id of the Organization you would like to add this app to. | [optional] 
**AdditionalDataIsRootPayload** | Pointer to **bool** | iOS: Notification data (additional data) values will be added to the root of the apns payload when sent to the device.  Ignore if you&#39;re not using any other plugins, or not using OneSignal SDK methods to read the payload. | [optional] 

## Methods

### NewApp

`func NewApp(id string, ) *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *App) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *App) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *App) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *App) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *App) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlayers

`func (o *App) GetPlayers() int32`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *App) GetPlayersOk() (*int32, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *App) SetPlayers(v int32)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *App) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetMessageablePlayers

`func (o *App) GetMessageablePlayers() int32`

GetMessageablePlayers returns the MessageablePlayers field if non-nil, zero value otherwise.

### GetMessageablePlayersOk

`func (o *App) GetMessageablePlayersOk() (*int32, bool)`

GetMessageablePlayersOk returns a tuple with the MessageablePlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageablePlayers

`func (o *App) SetMessageablePlayers(v int32)`

SetMessageablePlayers sets MessageablePlayers field to given value.

### HasMessageablePlayers

`func (o *App) HasMessageablePlayers() bool`

HasMessageablePlayers returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *App) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *App) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *App) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *App) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *App) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *App) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *App) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *App) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAndroidGcmSenderId

`func (o *App) GetAndroidGcmSenderId() string`

GetAndroidGcmSenderId returns the AndroidGcmSenderId field if non-nil, zero value otherwise.

### GetAndroidGcmSenderIdOk

`func (o *App) GetAndroidGcmSenderIdOk() (*string, bool)`

GetAndroidGcmSenderIdOk returns a tuple with the AndroidGcmSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidGcmSenderId

`func (o *App) SetAndroidGcmSenderId(v string)`

SetAndroidGcmSenderId sets AndroidGcmSenderId field to given value.

### HasAndroidGcmSenderId

`func (o *App) HasAndroidGcmSenderId() bool`

HasAndroidGcmSenderId returns a boolean if a field has been set.

### GetGcmKey

`func (o *App) GetGcmKey() string`

GetGcmKey returns the GcmKey field if non-nil, zero value otherwise.

### GetGcmKeyOk

`func (o *App) GetGcmKeyOk() (*string, bool)`

GetGcmKeyOk returns a tuple with the GcmKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcmKey

`func (o *App) SetGcmKey(v string)`

SetGcmKey sets GcmKey field to given value.

### HasGcmKey

`func (o *App) HasGcmKey() bool`

HasGcmKey returns a boolean if a field has been set.

### SetGcmKeyNil

`func (o *App) SetGcmKeyNil(b bool)`

 SetGcmKeyNil sets the value for GcmKey to be an explicit nil

### UnsetGcmKey
`func (o *App) UnsetGcmKey()`

UnsetGcmKey ensures that no value is present for GcmKey, not even an explicit nil
### GetChromeWebOrigin

`func (o *App) GetChromeWebOrigin() string`

GetChromeWebOrigin returns the ChromeWebOrigin field if non-nil, zero value otherwise.

### GetChromeWebOriginOk

`func (o *App) GetChromeWebOriginOk() (*string, bool)`

GetChromeWebOriginOk returns a tuple with the ChromeWebOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebOrigin

`func (o *App) SetChromeWebOrigin(v string)`

SetChromeWebOrigin sets ChromeWebOrigin field to given value.

### HasChromeWebOrigin

`func (o *App) HasChromeWebOrigin() bool`

HasChromeWebOrigin returns a boolean if a field has been set.

### SetChromeWebOriginNil

`func (o *App) SetChromeWebOriginNil(b bool)`

 SetChromeWebOriginNil sets the value for ChromeWebOrigin to be an explicit nil

### UnsetChromeWebOrigin
`func (o *App) UnsetChromeWebOrigin()`

UnsetChromeWebOrigin ensures that no value is present for ChromeWebOrigin, not even an explicit nil
### GetChromeKey

`func (o *App) GetChromeKey() string`

GetChromeKey returns the ChromeKey field if non-nil, zero value otherwise.

### GetChromeKeyOk

`func (o *App) GetChromeKeyOk() (*string, bool)`

GetChromeKeyOk returns a tuple with the ChromeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeKey

`func (o *App) SetChromeKey(v string)`

SetChromeKey sets ChromeKey field to given value.

### HasChromeKey

`func (o *App) HasChromeKey() bool`

HasChromeKey returns a boolean if a field has been set.

### SetChromeKeyNil

`func (o *App) SetChromeKeyNil(b bool)`

 SetChromeKeyNil sets the value for ChromeKey to be an explicit nil

### UnsetChromeKey
`func (o *App) UnsetChromeKey()`

UnsetChromeKey ensures that no value is present for ChromeKey, not even an explicit nil
### GetChromeWebDefaultNotificationIcon

`func (o *App) GetChromeWebDefaultNotificationIcon() string`

GetChromeWebDefaultNotificationIcon returns the ChromeWebDefaultNotificationIcon field if non-nil, zero value otherwise.

### GetChromeWebDefaultNotificationIconOk

`func (o *App) GetChromeWebDefaultNotificationIconOk() (*string, bool)`

GetChromeWebDefaultNotificationIconOk returns a tuple with the ChromeWebDefaultNotificationIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebDefaultNotificationIcon

`func (o *App) SetChromeWebDefaultNotificationIcon(v string)`

SetChromeWebDefaultNotificationIcon sets ChromeWebDefaultNotificationIcon field to given value.

### HasChromeWebDefaultNotificationIcon

`func (o *App) HasChromeWebDefaultNotificationIcon() bool`

HasChromeWebDefaultNotificationIcon returns a boolean if a field has been set.

### SetChromeWebDefaultNotificationIconNil

`func (o *App) SetChromeWebDefaultNotificationIconNil(b bool)`

 SetChromeWebDefaultNotificationIconNil sets the value for ChromeWebDefaultNotificationIcon to be an explicit nil

### UnsetChromeWebDefaultNotificationIcon
`func (o *App) UnsetChromeWebDefaultNotificationIcon()`

UnsetChromeWebDefaultNotificationIcon ensures that no value is present for ChromeWebDefaultNotificationIcon, not even an explicit nil
### GetChromeWebSubDomain

`func (o *App) GetChromeWebSubDomain() string`

GetChromeWebSubDomain returns the ChromeWebSubDomain field if non-nil, zero value otherwise.

### GetChromeWebSubDomainOk

`func (o *App) GetChromeWebSubDomainOk() (*string, bool)`

GetChromeWebSubDomainOk returns a tuple with the ChromeWebSubDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebSubDomain

`func (o *App) SetChromeWebSubDomain(v string)`

SetChromeWebSubDomain sets ChromeWebSubDomain field to given value.

### HasChromeWebSubDomain

`func (o *App) HasChromeWebSubDomain() bool`

HasChromeWebSubDomain returns a boolean if a field has been set.

### SetChromeWebSubDomainNil

`func (o *App) SetChromeWebSubDomainNil(b bool)`

 SetChromeWebSubDomainNil sets the value for ChromeWebSubDomain to be an explicit nil

### UnsetChromeWebSubDomain
`func (o *App) UnsetChromeWebSubDomain()`

UnsetChromeWebSubDomain ensures that no value is present for ChromeWebSubDomain, not even an explicit nil
### GetApnsEnv

`func (o *App) GetApnsEnv() string`

GetApnsEnv returns the ApnsEnv field if non-nil, zero value otherwise.

### GetApnsEnvOk

`func (o *App) GetApnsEnvOk() (*string, bool)`

GetApnsEnvOk returns a tuple with the ApnsEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsEnv

`func (o *App) SetApnsEnv(v string)`

SetApnsEnv sets ApnsEnv field to given value.

### HasApnsEnv

`func (o *App) HasApnsEnv() bool`

HasApnsEnv returns a boolean if a field has been set.

### SetApnsEnvNil

`func (o *App) SetApnsEnvNil(b bool)`

 SetApnsEnvNil sets the value for ApnsEnv to be an explicit nil

### UnsetApnsEnv
`func (o *App) UnsetApnsEnv()`

UnsetApnsEnv ensures that no value is present for ApnsEnv, not even an explicit nil
### GetApnsP12

`func (o *App) GetApnsP12() string`

GetApnsP12 returns the ApnsP12 field if non-nil, zero value otherwise.

### GetApnsP12Ok

`func (o *App) GetApnsP12Ok() (*string, bool)`

GetApnsP12Ok returns a tuple with the ApnsP12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsP12

`func (o *App) SetApnsP12(v string)`

SetApnsP12 sets ApnsP12 field to given value.

### HasApnsP12

`func (o *App) HasApnsP12() bool`

HasApnsP12 returns a boolean if a field has been set.

### GetApnsP12Password

`func (o *App) GetApnsP12Password() string`

GetApnsP12Password returns the ApnsP12Password field if non-nil, zero value otherwise.

### GetApnsP12PasswordOk

`func (o *App) GetApnsP12PasswordOk() (*string, bool)`

GetApnsP12PasswordOk returns a tuple with the ApnsP12Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsP12Password

`func (o *App) SetApnsP12Password(v string)`

SetApnsP12Password sets ApnsP12Password field to given value.

### HasApnsP12Password

`func (o *App) HasApnsP12Password() bool`

HasApnsP12Password returns a boolean if a field has been set.

### GetApnsCertificates

`func (o *App) GetApnsCertificates() string`

GetApnsCertificates returns the ApnsCertificates field if non-nil, zero value otherwise.

### GetApnsCertificatesOk

`func (o *App) GetApnsCertificatesOk() (*string, bool)`

GetApnsCertificatesOk returns a tuple with the ApnsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsCertificates

`func (o *App) SetApnsCertificates(v string)`

SetApnsCertificates sets ApnsCertificates field to given value.

### HasApnsCertificates

`func (o *App) HasApnsCertificates() bool`

HasApnsCertificates returns a boolean if a field has been set.

### SetApnsCertificatesNil

`func (o *App) SetApnsCertificatesNil(b bool)`

 SetApnsCertificatesNil sets the value for ApnsCertificates to be an explicit nil

### UnsetApnsCertificates
`func (o *App) UnsetApnsCertificates()`

UnsetApnsCertificates ensures that no value is present for ApnsCertificates, not even an explicit nil
### GetSafariApnsCertificates

`func (o *App) GetSafariApnsCertificates() string`

GetSafariApnsCertificates returns the SafariApnsCertificates field if non-nil, zero value otherwise.

### GetSafariApnsCertificatesOk

`func (o *App) GetSafariApnsCertificatesOk() (*string, bool)`

GetSafariApnsCertificatesOk returns a tuple with the SafariApnsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariApnsCertificates

`func (o *App) SetSafariApnsCertificates(v string)`

SetSafariApnsCertificates sets SafariApnsCertificates field to given value.

### HasSafariApnsCertificates

`func (o *App) HasSafariApnsCertificates() bool`

HasSafariApnsCertificates returns a boolean if a field has been set.

### GetSafariApnsP12

`func (o *App) GetSafariApnsP12() string`

GetSafariApnsP12 returns the SafariApnsP12 field if non-nil, zero value otherwise.

### GetSafariApnsP12Ok

`func (o *App) GetSafariApnsP12Ok() (*string, bool)`

GetSafariApnsP12Ok returns a tuple with the SafariApnsP12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariApnsP12

`func (o *App) SetSafariApnsP12(v string)`

SetSafariApnsP12 sets SafariApnsP12 field to given value.

### HasSafariApnsP12

`func (o *App) HasSafariApnsP12() bool`

HasSafariApnsP12 returns a boolean if a field has been set.

### GetSafariApnsP12Password

`func (o *App) GetSafariApnsP12Password() string`

GetSafariApnsP12Password returns the SafariApnsP12Password field if non-nil, zero value otherwise.

### GetSafariApnsP12PasswordOk

`func (o *App) GetSafariApnsP12PasswordOk() (*string, bool)`

GetSafariApnsP12PasswordOk returns a tuple with the SafariApnsP12Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariApnsP12Password

`func (o *App) SetSafariApnsP12Password(v string)`

SetSafariApnsP12Password sets SafariApnsP12Password field to given value.

### HasSafariApnsP12Password

`func (o *App) HasSafariApnsP12Password() bool`

HasSafariApnsP12Password returns a boolean if a field has been set.

### GetSafariSiteOrigin

`func (o *App) GetSafariSiteOrigin() string`

GetSafariSiteOrigin returns the SafariSiteOrigin field if non-nil, zero value otherwise.

### GetSafariSiteOriginOk

`func (o *App) GetSafariSiteOriginOk() (*string, bool)`

GetSafariSiteOriginOk returns a tuple with the SafariSiteOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariSiteOrigin

`func (o *App) SetSafariSiteOrigin(v string)`

SetSafariSiteOrigin sets SafariSiteOrigin field to given value.

### HasSafariSiteOrigin

`func (o *App) HasSafariSiteOrigin() bool`

HasSafariSiteOrigin returns a boolean if a field has been set.

### SetSafariSiteOriginNil

`func (o *App) SetSafariSiteOriginNil(b bool)`

 SetSafariSiteOriginNil sets the value for SafariSiteOrigin to be an explicit nil

### UnsetSafariSiteOrigin
`func (o *App) UnsetSafariSiteOrigin()`

UnsetSafariSiteOrigin ensures that no value is present for SafariSiteOrigin, not even an explicit nil
### GetSafariPushId

`func (o *App) GetSafariPushId() string`

GetSafariPushId returns the SafariPushId field if non-nil, zero value otherwise.

### GetSafariPushIdOk

`func (o *App) GetSafariPushIdOk() (*string, bool)`

GetSafariPushIdOk returns a tuple with the SafariPushId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariPushId

`func (o *App) SetSafariPushId(v string)`

SetSafariPushId sets SafariPushId field to given value.

### HasSafariPushId

`func (o *App) HasSafariPushId() bool`

HasSafariPushId returns a boolean if a field has been set.

### SetSafariPushIdNil

`func (o *App) SetSafariPushIdNil(b bool)`

 SetSafariPushIdNil sets the value for SafariPushId to be an explicit nil

### UnsetSafariPushId
`func (o *App) UnsetSafariPushId()`

UnsetSafariPushId ensures that no value is present for SafariPushId, not even an explicit nil
### GetSafariIcon1616

`func (o *App) GetSafariIcon1616() string`

GetSafariIcon1616 returns the SafariIcon1616 field if non-nil, zero value otherwise.

### GetSafariIcon1616Ok

`func (o *App) GetSafariIcon1616Ok() (*string, bool)`

GetSafariIcon1616Ok returns a tuple with the SafariIcon1616 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariIcon1616

`func (o *App) SetSafariIcon1616(v string)`

SetSafariIcon1616 sets SafariIcon1616 field to given value.

### HasSafariIcon1616

`func (o *App) HasSafariIcon1616() bool`

HasSafariIcon1616 returns a boolean if a field has been set.

### GetSafariIcon3232

`func (o *App) GetSafariIcon3232() string`

GetSafariIcon3232 returns the SafariIcon3232 field if non-nil, zero value otherwise.

### GetSafariIcon3232Ok

`func (o *App) GetSafariIcon3232Ok() (*string, bool)`

GetSafariIcon3232Ok returns a tuple with the SafariIcon3232 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariIcon3232

`func (o *App) SetSafariIcon3232(v string)`

SetSafariIcon3232 sets SafariIcon3232 field to given value.

### HasSafariIcon3232

`func (o *App) HasSafariIcon3232() bool`

HasSafariIcon3232 returns a boolean if a field has been set.

### GetSafariIcon6464

`func (o *App) GetSafariIcon6464() string`

GetSafariIcon6464 returns the SafariIcon6464 field if non-nil, zero value otherwise.

### GetSafariIcon6464Ok

`func (o *App) GetSafariIcon6464Ok() (*string, bool)`

GetSafariIcon6464Ok returns a tuple with the SafariIcon6464 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariIcon6464

`func (o *App) SetSafariIcon6464(v string)`

SetSafariIcon6464 sets SafariIcon6464 field to given value.

### HasSafariIcon6464

`func (o *App) HasSafariIcon6464() bool`

HasSafariIcon6464 returns a boolean if a field has been set.

### GetSafariIcon128128

`func (o *App) GetSafariIcon128128() string`

GetSafariIcon128128 returns the SafariIcon128128 field if non-nil, zero value otherwise.

### GetSafariIcon128128Ok

`func (o *App) GetSafariIcon128128Ok() (*string, bool)`

GetSafariIcon128128Ok returns a tuple with the SafariIcon128128 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariIcon128128

`func (o *App) SetSafariIcon128128(v string)`

SetSafariIcon128128 sets SafariIcon128128 field to given value.

### HasSafariIcon128128

`func (o *App) HasSafariIcon128128() bool`

HasSafariIcon128128 returns a boolean if a field has been set.

### GetSafariIcon256256

`func (o *App) GetSafariIcon256256() string`

GetSafariIcon256256 returns the SafariIcon256256 field if non-nil, zero value otherwise.

### GetSafariIcon256256Ok

`func (o *App) GetSafariIcon256256Ok() (*string, bool)`

GetSafariIcon256256Ok returns a tuple with the SafariIcon256256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariIcon256256

`func (o *App) SetSafariIcon256256(v string)`

SetSafariIcon256256 sets SafariIcon256256 field to given value.

### HasSafariIcon256256

`func (o *App) HasSafariIcon256256() bool`

HasSafariIcon256256 returns a boolean if a field has been set.

### GetSiteName

`func (o *App) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *App) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *App) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *App) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *App) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *App) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetBasicAuthKey

`func (o *App) GetBasicAuthKey() string`

GetBasicAuthKey returns the BasicAuthKey field if non-nil, zero value otherwise.

### GetBasicAuthKeyOk

`func (o *App) GetBasicAuthKeyOk() (*string, bool)`

GetBasicAuthKeyOk returns a tuple with the BasicAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthKey

`func (o *App) SetBasicAuthKey(v string)`

SetBasicAuthKey sets BasicAuthKey field to given value.

### HasBasicAuthKey

`func (o *App) HasBasicAuthKey() bool`

HasBasicAuthKey returns a boolean if a field has been set.

### GetOrganizationId

`func (o *App) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *App) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *App) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *App) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAdditionalDataIsRootPayload

`func (o *App) GetAdditionalDataIsRootPayload() bool`

GetAdditionalDataIsRootPayload returns the AdditionalDataIsRootPayload field if non-nil, zero value otherwise.

### GetAdditionalDataIsRootPayloadOk

`func (o *App) GetAdditionalDataIsRootPayloadOk() (*bool, bool)`

GetAdditionalDataIsRootPayloadOk returns a tuple with the AdditionalDataIsRootPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDataIsRootPayload

`func (o *App) SetAdditionalDataIsRootPayload(v bool)`

SetAdditionalDataIsRootPayload sets AdditionalDataIsRootPayload field to given value.

### HasAdditionalDataIsRootPayload

`func (o *App) HasAdditionalDataIsRootPayload() bool`

HasAdditionalDataIsRootPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


