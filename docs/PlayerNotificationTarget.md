# PlayerNotificationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludePlayerIds** | Pointer to **[]string** | Specific playerids to send your notification to. _Does not require API Auth Key. Do not combine with other targeting parameters. Not compatible with any other targeting parameters. Example: [\&quot;1dd608f2-c6a1-11e3-851d-000c2940e62c\&quot;] Limit of 2,000 entries per REST API call  | [optional] 
**IncludeExternalUserIds** | Pointer to **[]string** | Target specific devices by custom user IDs assigned via API. Not compatible with any other targeting parameters Example: [\&quot;custom-id-assigned-by-api\&quot;] REQUIRED: REST API Key Authentication Limit of 2,000 entries per REST API call. Note: If targeting push, email, or sms subscribers with same ids, use with channel_for_external_user_ids to indicate you are sending a push or email or sms.  | [optional] 
**IncludeEmailTokens** | Pointer to **[]string** | Recommended for Sending Emails - Target specific email addresses. If an email does not correspond to an existing user, a new user will be created. Example: nick@catfac.ts Limit of 2,000 entries per REST API call  | [optional] 
**IncludePhoneNumbers** | Pointer to **[]string** | Recommended for Sending SMS - Target specific phone numbers. The phone number should be in the E.164 format. Phone number should be an existing subscriber on OneSignal. Refer our docs to learn how to add phone numbers to OneSignal. Example phone number: +1999999999 Limit of 2,000 entries per REST API call  | [optional] 
**IncludeIosTokens** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using iOS device tokens. Warning: Only works with Production tokens. All non-alphanumeric characters must be removed from each token. If a token does not correspond to an existing user, a new user will be created. Example: ce777617da7f548fe7a9ab6febb56cf39fba6d38203... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeWpWnsUris** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Windows URIs. If a token does not correspond to an existing user, a new user will be created. Example: http://s.notify.live.net/u/1/bn1/HmQAAACPaLDr-... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeAmazonRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Amazon ADM registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: amzn1.adm-registration.v1.XpvSSUk0Rc3hTVVV... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeChromeRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Chrome App registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeChromeWebRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Chrome Web Push registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeAndroidRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Android device registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call  | [optional] 

## Methods

### NewPlayerNotificationTarget

`func NewPlayerNotificationTarget() *PlayerNotificationTarget`

NewPlayerNotificationTarget instantiates a new PlayerNotificationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerNotificationTargetWithDefaults

`func NewPlayerNotificationTargetWithDefaults() *PlayerNotificationTarget`

NewPlayerNotificationTargetWithDefaults instantiates a new PlayerNotificationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludePlayerIds

`func (o *PlayerNotificationTarget) GetIncludePlayerIds() []string`

GetIncludePlayerIds returns the IncludePlayerIds field if non-nil, zero value otherwise.

### GetIncludePlayerIdsOk

`func (o *PlayerNotificationTarget) GetIncludePlayerIdsOk() (*[]string, bool)`

GetIncludePlayerIdsOk returns a tuple with the IncludePlayerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePlayerIds

`func (o *PlayerNotificationTarget) SetIncludePlayerIds(v []string)`

SetIncludePlayerIds sets IncludePlayerIds field to given value.

### HasIncludePlayerIds

`func (o *PlayerNotificationTarget) HasIncludePlayerIds() bool`

HasIncludePlayerIds returns a boolean if a field has been set.

### SetIncludePlayerIdsNil

`func (o *PlayerNotificationTarget) SetIncludePlayerIdsNil(b bool)`

 SetIncludePlayerIdsNil sets the value for IncludePlayerIds to be an explicit nil

### UnsetIncludePlayerIds
`func (o *PlayerNotificationTarget) UnsetIncludePlayerIds()`

UnsetIncludePlayerIds ensures that no value is present for IncludePlayerIds, not even an explicit nil
### GetIncludeExternalUserIds

`func (o *PlayerNotificationTarget) GetIncludeExternalUserIds() []string`

GetIncludeExternalUserIds returns the IncludeExternalUserIds field if non-nil, zero value otherwise.

### GetIncludeExternalUserIdsOk

`func (o *PlayerNotificationTarget) GetIncludeExternalUserIdsOk() (*[]string, bool)`

GetIncludeExternalUserIdsOk returns a tuple with the IncludeExternalUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExternalUserIds

`func (o *PlayerNotificationTarget) SetIncludeExternalUserIds(v []string)`

SetIncludeExternalUserIds sets IncludeExternalUserIds field to given value.

### HasIncludeExternalUserIds

`func (o *PlayerNotificationTarget) HasIncludeExternalUserIds() bool`

HasIncludeExternalUserIds returns a boolean if a field has been set.

### SetIncludeExternalUserIdsNil

`func (o *PlayerNotificationTarget) SetIncludeExternalUserIdsNil(b bool)`

 SetIncludeExternalUserIdsNil sets the value for IncludeExternalUserIds to be an explicit nil

### UnsetIncludeExternalUserIds
`func (o *PlayerNotificationTarget) UnsetIncludeExternalUserIds()`

UnsetIncludeExternalUserIds ensures that no value is present for IncludeExternalUserIds, not even an explicit nil
### GetIncludeEmailTokens

`func (o *PlayerNotificationTarget) GetIncludeEmailTokens() []string`

GetIncludeEmailTokens returns the IncludeEmailTokens field if non-nil, zero value otherwise.

### GetIncludeEmailTokensOk

`func (o *PlayerNotificationTarget) GetIncludeEmailTokensOk() (*[]string, bool)`

GetIncludeEmailTokensOk returns a tuple with the IncludeEmailTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEmailTokens

`func (o *PlayerNotificationTarget) SetIncludeEmailTokens(v []string)`

SetIncludeEmailTokens sets IncludeEmailTokens field to given value.

### HasIncludeEmailTokens

`func (o *PlayerNotificationTarget) HasIncludeEmailTokens() bool`

HasIncludeEmailTokens returns a boolean if a field has been set.

### GetIncludePhoneNumbers

`func (o *PlayerNotificationTarget) GetIncludePhoneNumbers() []string`

GetIncludePhoneNumbers returns the IncludePhoneNumbers field if non-nil, zero value otherwise.

### GetIncludePhoneNumbersOk

`func (o *PlayerNotificationTarget) GetIncludePhoneNumbersOk() (*[]string, bool)`

GetIncludePhoneNumbersOk returns a tuple with the IncludePhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePhoneNumbers

`func (o *PlayerNotificationTarget) SetIncludePhoneNumbers(v []string)`

SetIncludePhoneNumbers sets IncludePhoneNumbers field to given value.

### HasIncludePhoneNumbers

`func (o *PlayerNotificationTarget) HasIncludePhoneNumbers() bool`

HasIncludePhoneNumbers returns a boolean if a field has been set.

### GetIncludeIosTokens

`func (o *PlayerNotificationTarget) GetIncludeIosTokens() []string`

GetIncludeIosTokens returns the IncludeIosTokens field if non-nil, zero value otherwise.

### GetIncludeIosTokensOk

`func (o *PlayerNotificationTarget) GetIncludeIosTokensOk() (*[]string, bool)`

GetIncludeIosTokensOk returns a tuple with the IncludeIosTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIosTokens

`func (o *PlayerNotificationTarget) SetIncludeIosTokens(v []string)`

SetIncludeIosTokens sets IncludeIosTokens field to given value.

### HasIncludeIosTokens

`func (o *PlayerNotificationTarget) HasIncludeIosTokens() bool`

HasIncludeIosTokens returns a boolean if a field has been set.

### GetIncludeWpWnsUris

`func (o *PlayerNotificationTarget) GetIncludeWpWnsUris() []string`

GetIncludeWpWnsUris returns the IncludeWpWnsUris field if non-nil, zero value otherwise.

### GetIncludeWpWnsUrisOk

`func (o *PlayerNotificationTarget) GetIncludeWpWnsUrisOk() (*[]string, bool)`

GetIncludeWpWnsUrisOk returns a tuple with the IncludeWpWnsUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeWpWnsUris

`func (o *PlayerNotificationTarget) SetIncludeWpWnsUris(v []string)`

SetIncludeWpWnsUris sets IncludeWpWnsUris field to given value.

### HasIncludeWpWnsUris

`func (o *PlayerNotificationTarget) HasIncludeWpWnsUris() bool`

HasIncludeWpWnsUris returns a boolean if a field has been set.

### GetIncludeAmazonRegIds

`func (o *PlayerNotificationTarget) GetIncludeAmazonRegIds() []string`

GetIncludeAmazonRegIds returns the IncludeAmazonRegIds field if non-nil, zero value otherwise.

### GetIncludeAmazonRegIdsOk

`func (o *PlayerNotificationTarget) GetIncludeAmazonRegIdsOk() (*[]string, bool)`

GetIncludeAmazonRegIdsOk returns a tuple with the IncludeAmazonRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAmazonRegIds

`func (o *PlayerNotificationTarget) SetIncludeAmazonRegIds(v []string)`

SetIncludeAmazonRegIds sets IncludeAmazonRegIds field to given value.

### HasIncludeAmazonRegIds

`func (o *PlayerNotificationTarget) HasIncludeAmazonRegIds() bool`

HasIncludeAmazonRegIds returns a boolean if a field has been set.

### GetIncludeChromeRegIds

`func (o *PlayerNotificationTarget) GetIncludeChromeRegIds() []string`

GetIncludeChromeRegIds returns the IncludeChromeRegIds field if non-nil, zero value otherwise.

### GetIncludeChromeRegIdsOk

`func (o *PlayerNotificationTarget) GetIncludeChromeRegIdsOk() (*[]string, bool)`

GetIncludeChromeRegIdsOk returns a tuple with the IncludeChromeRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChromeRegIds

`func (o *PlayerNotificationTarget) SetIncludeChromeRegIds(v []string)`

SetIncludeChromeRegIds sets IncludeChromeRegIds field to given value.

### HasIncludeChromeRegIds

`func (o *PlayerNotificationTarget) HasIncludeChromeRegIds() bool`

HasIncludeChromeRegIds returns a boolean if a field has been set.

### GetIncludeChromeWebRegIds

`func (o *PlayerNotificationTarget) GetIncludeChromeWebRegIds() []string`

GetIncludeChromeWebRegIds returns the IncludeChromeWebRegIds field if non-nil, zero value otherwise.

### GetIncludeChromeWebRegIdsOk

`func (o *PlayerNotificationTarget) GetIncludeChromeWebRegIdsOk() (*[]string, bool)`

GetIncludeChromeWebRegIdsOk returns a tuple with the IncludeChromeWebRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChromeWebRegIds

`func (o *PlayerNotificationTarget) SetIncludeChromeWebRegIds(v []string)`

SetIncludeChromeWebRegIds sets IncludeChromeWebRegIds field to given value.

### HasIncludeChromeWebRegIds

`func (o *PlayerNotificationTarget) HasIncludeChromeWebRegIds() bool`

HasIncludeChromeWebRegIds returns a boolean if a field has been set.

### GetIncludeAndroidRegIds

`func (o *PlayerNotificationTarget) GetIncludeAndroidRegIds() []string`

GetIncludeAndroidRegIds returns the IncludeAndroidRegIds field if non-nil, zero value otherwise.

### GetIncludeAndroidRegIdsOk

`func (o *PlayerNotificationTarget) GetIncludeAndroidRegIdsOk() (*[]string, bool)`

GetIncludeAndroidRegIdsOk returns a tuple with the IncludeAndroidRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAndroidRegIds

`func (o *PlayerNotificationTarget) SetIncludeAndroidRegIds(v []string)`

SetIncludeAndroidRegIds sets IncludeAndroidRegIds field to given value.

### HasIncludeAndroidRegIds

`func (o *PlayerNotificationTarget) HasIncludeAndroidRegIds() bool`

HasIncludeAndroidRegIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


