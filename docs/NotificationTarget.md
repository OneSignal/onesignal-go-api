# NotificationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedSegments** | Pointer to **[]string** | The segment names you want to target. Users in these segments will receive a notification. This targeting parameter is only compatible with excluded_segments. Example: [\&quot;Active Users\&quot;, \&quot;Inactive Users\&quot;]  | [optional] 
**ExcludedSegments** | Pointer to **[]string** | Segment that will be excluded when sending. Users in these segments will not receive a notification, even if they were included in included_segments. This targeting parameter is only compatible with included_segments. Example: [\&quot;Active Users\&quot;, \&quot;Inactive Users\&quot;]  | [optional] 
**IncludeSubscriptionIds** | Pointer to **[]string** | Specific subscription ids to send your notification to. _Does not require API Auth Key._ Not compatible with any other targeting parameters. Example: [\&quot;1dd608f2-c6a1-11e3-851d-000c2940e62c\&quot;] Limit of 2,000 entries per REST API call  | [optional] 
**IncludeEmailTokens** | Pointer to **[]string** | Recommended for Sending Emails - Target specific email addresses. If an email does not correspond to an existing user, a new user will be created. Example: nick@catfac.ts Limit of 2,000 entries per REST API call  | [optional] 
**IncludePhoneNumbers** | Pointer to **[]string** | Recommended for Sending SMS - Target specific phone numbers. The phone number should be in the E.164 format. Phone number should be an existing subscriber on OneSignal. Refer our docs to learn how to add phone numbers to OneSignal. Example phone number: +1999999999 Limit of 2,000 entries per REST API call  | [optional] 
**IncludeIosTokens** | Pointer to **[]string** | Not Recommended: Please consider using include_subscription_ids or include_aliases instead. Target using iOS device tokens. Warning: Only works with Production tokens. All non-alphanumeric characters must be removed from each token. If a token does not correspond to an existing user, a new user will be created. Example: ce777617da7f548fe7a9ab6febb56cf39fba6d38203... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeWpWnsUris** | Pointer to **[]string** | Not Recommended: Please consider using include_subscription_ids or include_aliases instead. Target using Windows URIs. If a token does not correspond to an existing user, a new user will be created. Example: http://s.notify.live.net/u/1/bn1/HmQAAACPaLDr-... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeAmazonRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_subscription_ids or include_aliases instead. Target using Amazon ADM registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: amzn1.adm-registration.v1.XpvSSUk0Rc3hTVVV... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeChromeRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_subscription_ids or include_aliases instead. Target using Chrome App registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeChromeWebRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_subscription_ids or include_aliases instead. Target using Chrome Web Push registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeAndroidRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_subscription_ids or include_aliases instead. Target using Android device registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeAliases** | Pointer to **map[string][]string** | Target specific users by aliases assigned via API. An alias can be an external_id, onesignal_id, or a custom alias. Accepts an object where keys are alias labels and values are arrays of alias IDs to include Example usage: { \&quot;external_id\&quot;: [\&quot;exId1\&quot;, \&quot;extId2\&quot;], \&quot;internal_label\&quot;: [\&quot;id1\&quot;, \&quot;id2\&quot;] } Not compatible with any other targeting parameters. REQUIRED: REST API Key Authentication Limit of 2,000 entries per REST API call Note: If targeting push, email, or sms subscribers with same ids, use with target_channel to indicate you are sending a push or email or sms. | [optional] 
**TargetChannel** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationTarget

`func NewNotificationTarget() *NotificationTarget`

NewNotificationTarget instantiates a new NotificationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationTargetWithDefaults

`func NewNotificationTargetWithDefaults() *NotificationTarget`

NewNotificationTargetWithDefaults instantiates a new NotificationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedSegments

`func (o *NotificationTarget) GetIncludedSegments() []string`

GetIncludedSegments returns the IncludedSegments field if non-nil, zero value otherwise.

### GetIncludedSegmentsOk

`func (o *NotificationTarget) GetIncludedSegmentsOk() (*[]string, bool)`

GetIncludedSegmentsOk returns a tuple with the IncludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegments

`func (o *NotificationTarget) SetIncludedSegments(v []string)`

SetIncludedSegments sets IncludedSegments field to given value.

### HasIncludedSegments

`func (o *NotificationTarget) HasIncludedSegments() bool`

HasIncludedSegments returns a boolean if a field has been set.

### GetExcludedSegments

`func (o *NotificationTarget) GetExcludedSegments() []string`

GetExcludedSegments returns the ExcludedSegments field if non-nil, zero value otherwise.

### GetExcludedSegmentsOk

`func (o *NotificationTarget) GetExcludedSegmentsOk() (*[]string, bool)`

GetExcludedSegmentsOk returns a tuple with the ExcludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegments

`func (o *NotificationTarget) SetExcludedSegments(v []string)`

SetExcludedSegments sets ExcludedSegments field to given value.

### HasExcludedSegments

`func (o *NotificationTarget) HasExcludedSegments() bool`

HasExcludedSegments returns a boolean if a field has been set.

### GetIncludeSubscriptionIds

`func (o *NotificationTarget) GetIncludeSubscriptionIds() []string`

GetIncludeSubscriptionIds returns the IncludeSubscriptionIds field if non-nil, zero value otherwise.

### GetIncludeSubscriptionIdsOk

`func (o *NotificationTarget) GetIncludeSubscriptionIdsOk() (*[]string, bool)`

GetIncludeSubscriptionIdsOk returns a tuple with the IncludeSubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubscriptionIds

`func (o *NotificationTarget) SetIncludeSubscriptionIds(v []string)`

SetIncludeSubscriptionIds sets IncludeSubscriptionIds field to given value.

### HasIncludeSubscriptionIds

`func (o *NotificationTarget) HasIncludeSubscriptionIds() bool`

HasIncludeSubscriptionIds returns a boolean if a field has been set.

### SetIncludeSubscriptionIdsNil

`func (o *NotificationTarget) SetIncludeSubscriptionIdsNil(b bool)`

 SetIncludeSubscriptionIdsNil sets the value for IncludeSubscriptionIds to be an explicit nil

### UnsetIncludeSubscriptionIds
`func (o *NotificationTarget) UnsetIncludeSubscriptionIds()`

UnsetIncludeSubscriptionIds ensures that no value is present for IncludeSubscriptionIds, not even an explicit nil
### GetIncludeEmailTokens

`func (o *NotificationTarget) GetIncludeEmailTokens() []string`

GetIncludeEmailTokens returns the IncludeEmailTokens field if non-nil, zero value otherwise.

### GetIncludeEmailTokensOk

`func (o *NotificationTarget) GetIncludeEmailTokensOk() (*[]string, bool)`

GetIncludeEmailTokensOk returns a tuple with the IncludeEmailTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEmailTokens

`func (o *NotificationTarget) SetIncludeEmailTokens(v []string)`

SetIncludeEmailTokens sets IncludeEmailTokens field to given value.

### HasIncludeEmailTokens

`func (o *NotificationTarget) HasIncludeEmailTokens() bool`

HasIncludeEmailTokens returns a boolean if a field has been set.

### GetIncludePhoneNumbers

`func (o *NotificationTarget) GetIncludePhoneNumbers() []string`

GetIncludePhoneNumbers returns the IncludePhoneNumbers field if non-nil, zero value otherwise.

### GetIncludePhoneNumbersOk

`func (o *NotificationTarget) GetIncludePhoneNumbersOk() (*[]string, bool)`

GetIncludePhoneNumbersOk returns a tuple with the IncludePhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePhoneNumbers

`func (o *NotificationTarget) SetIncludePhoneNumbers(v []string)`

SetIncludePhoneNumbers sets IncludePhoneNumbers field to given value.

### HasIncludePhoneNumbers

`func (o *NotificationTarget) HasIncludePhoneNumbers() bool`

HasIncludePhoneNumbers returns a boolean if a field has been set.

### GetIncludeIosTokens

`func (o *NotificationTarget) GetIncludeIosTokens() []string`

GetIncludeIosTokens returns the IncludeIosTokens field if non-nil, zero value otherwise.

### GetIncludeIosTokensOk

`func (o *NotificationTarget) GetIncludeIosTokensOk() (*[]string, bool)`

GetIncludeIosTokensOk returns a tuple with the IncludeIosTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIosTokens

`func (o *NotificationTarget) SetIncludeIosTokens(v []string)`

SetIncludeIosTokens sets IncludeIosTokens field to given value.

### HasIncludeIosTokens

`func (o *NotificationTarget) HasIncludeIosTokens() bool`

HasIncludeIosTokens returns a boolean if a field has been set.

### GetIncludeWpWnsUris

`func (o *NotificationTarget) GetIncludeWpWnsUris() []string`

GetIncludeWpWnsUris returns the IncludeWpWnsUris field if non-nil, zero value otherwise.

### GetIncludeWpWnsUrisOk

`func (o *NotificationTarget) GetIncludeWpWnsUrisOk() (*[]string, bool)`

GetIncludeWpWnsUrisOk returns a tuple with the IncludeWpWnsUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeWpWnsUris

`func (o *NotificationTarget) SetIncludeWpWnsUris(v []string)`

SetIncludeWpWnsUris sets IncludeWpWnsUris field to given value.

### HasIncludeWpWnsUris

`func (o *NotificationTarget) HasIncludeWpWnsUris() bool`

HasIncludeWpWnsUris returns a boolean if a field has been set.

### GetIncludeAmazonRegIds

`func (o *NotificationTarget) GetIncludeAmazonRegIds() []string`

GetIncludeAmazonRegIds returns the IncludeAmazonRegIds field if non-nil, zero value otherwise.

### GetIncludeAmazonRegIdsOk

`func (o *NotificationTarget) GetIncludeAmazonRegIdsOk() (*[]string, bool)`

GetIncludeAmazonRegIdsOk returns a tuple with the IncludeAmazonRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAmazonRegIds

`func (o *NotificationTarget) SetIncludeAmazonRegIds(v []string)`

SetIncludeAmazonRegIds sets IncludeAmazonRegIds field to given value.

### HasIncludeAmazonRegIds

`func (o *NotificationTarget) HasIncludeAmazonRegIds() bool`

HasIncludeAmazonRegIds returns a boolean if a field has been set.

### GetIncludeChromeRegIds

`func (o *NotificationTarget) GetIncludeChromeRegIds() []string`

GetIncludeChromeRegIds returns the IncludeChromeRegIds field if non-nil, zero value otherwise.

### GetIncludeChromeRegIdsOk

`func (o *NotificationTarget) GetIncludeChromeRegIdsOk() (*[]string, bool)`

GetIncludeChromeRegIdsOk returns a tuple with the IncludeChromeRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChromeRegIds

`func (o *NotificationTarget) SetIncludeChromeRegIds(v []string)`

SetIncludeChromeRegIds sets IncludeChromeRegIds field to given value.

### HasIncludeChromeRegIds

`func (o *NotificationTarget) HasIncludeChromeRegIds() bool`

HasIncludeChromeRegIds returns a boolean if a field has been set.

### GetIncludeChromeWebRegIds

`func (o *NotificationTarget) GetIncludeChromeWebRegIds() []string`

GetIncludeChromeWebRegIds returns the IncludeChromeWebRegIds field if non-nil, zero value otherwise.

### GetIncludeChromeWebRegIdsOk

`func (o *NotificationTarget) GetIncludeChromeWebRegIdsOk() (*[]string, bool)`

GetIncludeChromeWebRegIdsOk returns a tuple with the IncludeChromeWebRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChromeWebRegIds

`func (o *NotificationTarget) SetIncludeChromeWebRegIds(v []string)`

SetIncludeChromeWebRegIds sets IncludeChromeWebRegIds field to given value.

### HasIncludeChromeWebRegIds

`func (o *NotificationTarget) HasIncludeChromeWebRegIds() bool`

HasIncludeChromeWebRegIds returns a boolean if a field has been set.

### GetIncludeAndroidRegIds

`func (o *NotificationTarget) GetIncludeAndroidRegIds() []string`

GetIncludeAndroidRegIds returns the IncludeAndroidRegIds field if non-nil, zero value otherwise.

### GetIncludeAndroidRegIdsOk

`func (o *NotificationTarget) GetIncludeAndroidRegIdsOk() (*[]string, bool)`

GetIncludeAndroidRegIdsOk returns a tuple with the IncludeAndroidRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAndroidRegIds

`func (o *NotificationTarget) SetIncludeAndroidRegIds(v []string)`

SetIncludeAndroidRegIds sets IncludeAndroidRegIds field to given value.

### HasIncludeAndroidRegIds

`func (o *NotificationTarget) HasIncludeAndroidRegIds() bool`

HasIncludeAndroidRegIds returns a boolean if a field has been set.

### GetIncludeAliases

`func (o *NotificationTarget) GetIncludeAliases() map[string][]string`

GetIncludeAliases returns the IncludeAliases field if non-nil, zero value otherwise.

### GetIncludeAliasesOk

`func (o *NotificationTarget) GetIncludeAliasesOk() (*map[string][]string, bool)`

GetIncludeAliasesOk returns a tuple with the IncludeAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAliases

`func (o *NotificationTarget) SetIncludeAliases(v map[string][]string)`

SetIncludeAliases sets IncludeAliases field to given value.

### HasIncludeAliases

`func (o *NotificationTarget) HasIncludeAliases() bool`

HasIncludeAliases returns a boolean if a field has been set.

### SetIncludeAliasesNil

`func (o *NotificationTarget) SetIncludeAliasesNil(b bool)`

 SetIncludeAliasesNil sets the value for IncludeAliases to be an explicit nil

### UnsetIncludeAliases
`func (o *NotificationTarget) UnsetIncludeAliases()`

UnsetIncludeAliases ensures that no value is present for IncludeAliases, not even an explicit nil
### GetTargetChannel

`func (o *NotificationTarget) GetTargetChannel() string`

GetTargetChannel returns the TargetChannel field if non-nil, zero value otherwise.

### GetTargetChannelOk

`func (o *NotificationTarget) GetTargetChannelOk() (*string, bool)`

GetTargetChannelOk returns a tuple with the TargetChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChannel

`func (o *NotificationTarget) SetTargetChannel(v string)`

SetTargetChannel sets TargetChannel field to given value.

### HasTargetChannel

`func (o *NotificationTarget) HasTargetChannel() bool`

HasTargetChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


