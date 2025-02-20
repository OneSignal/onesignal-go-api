# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedSegments** | Pointer to **[]string** | The segment names you want to target. Users in these segments will receive a notification. This targeting parameter is only compatible with excluded_segments. Example: [\&quot;Active Users\&quot;, \&quot;Inactive Users\&quot;]  | [optional] 
**ExcludedSegments** | Pointer to **[]string** | Segment that will be excluded when sending. Users in these segments will not receive a notification, even if they were included in included_segments. This targeting parameter is only compatible with included_segments. Example: [\&quot;Active Users\&quot;, \&quot;Inactive Users\&quot;]  | [optional] 
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
**IncludeAliases** | Pointer to [**NullablePlayerNotificationTargetIncludeAliases**](PlayerNotificationTargetIncludeAliases.md) |  | [optional] 
**TargetChannel** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** | Required for SMS Messages. An identifier for tracking message within the OneSignal dashboard or export analytics. Not shown to end user. | [optional] 
**Aggregation** | Pointer to **string** |  | [optional] [readonly] 
**IsIos** | Pointer to **NullableBool** | Indicates whether to send to all devices registered under your app&#39;s Apple iOS platform. | [optional] 
**IsAndroid** | Pointer to **NullableBool** | Indicates whether to send to all devices registered under your app&#39;s Google Android platform. | [optional] 
**IsHuawei** | Pointer to **NullableBool** | Indicates whether to send to all devices registered under your app&#39;s Huawei Android platform. | [optional] 
**IsAnyWeb** | Pointer to **NullableBool** | Indicates whether to send to all subscribed web browser users, including Chrome, Firefox, and Safari. You may use this instead as a combined flag instead of separately enabling isChromeWeb, isFirefox, and isSafari, though the three options are equivalent to this one.  | [optional] 
**IsChromeWeb** | Pointer to **NullableBool** | Indicates whether to send to all Google Chrome, Chrome on Android, and Mozilla Firefox users registered under your Chrome &amp; Firefox web push platform. | [optional] 
**IsFirefox** | Pointer to **NullableBool** | Indicates whether to send to all Mozilla Firefox desktop users registered under your Firefox web push platform. | [optional] 
**IsSafari** | Pointer to **NullableBool** | Does not support iOS Safari. Indicates whether to send to all Apple&#39;s Safari desktop users registered under your Safari web push platform. Read more iOS Safari | [optional] 
**IsWPWNS** | Pointer to **NullableBool** | Indicates whether to send to all devices registered under your app&#39;s Windows platform. | [optional] 
**IsAdm** | Pointer to **NullableBool** | Indicates whether to send to all devices registered under your app&#39;s Amazon Fire platform. | [optional] 
**IsChrome** | Pointer to **NullableBool** | This flag is not used for web push Please see isChromeWeb for sending to web push users. This flag only applies to Google Chrome Apps &amp; Extensions. Indicates whether to send to all devices registered under your app&#39;s Google Chrome Apps &amp; Extension platform.  | [optional] 
**ChannelForExternalUserIds** | Pointer to **string** | Indicates if the message type when targeting with include_external_user_ids for cases where an email, sms, and/or push subscribers have the same external user id. Example: Use the string \&quot;push\&quot; to indicate you are sending a push notification or the string \&quot;email\&quot;for sending emails or \&quot;sms\&quot;for sending SMS.  | [optional] 
**AppId** | **string** | Required: Your OneSignal Application ID, which can be found in Keys &amp; IDs. It is a UUID and looks similar to 8250eaf6-1a58-489e-b136-7c74a864b434.  | 
**ExternalId** | Pointer to **NullableString** | [DEPRECATED] Correlation and idempotency key. A request received with this parameter will first look for another notification with the same external_id. If one exists, a notification will not be sent, and result of the previous operation will instead be returned. Therefore, if you plan on using this feature, it&#39;s important to use a good source of randomness to generate the UUID passed here. This key is only idempotent for 30 days. After 30 days, the notification could be removed from our system and a notification with the same external_id will be sent again.   See Idempotent Notification Requests for more details writeOnly: true  | [optional] 
**IdempotencyKey** | Pointer to **NullableString** | Correlation and idempotency key. A request received with this parameter will first look for another notification with the same idempotency key. If one exists, a notification will not be sent, and result of the previous operation will instead be returned. Therefore, if you plan on using this feature, it&#39;s important to use a good source of randomness to generate the UUID passed here. This key is only idempotent for 30 days. After 30 days, the notification could be removed from our system and a notification with the same idempotency key will be sent again.   See Idempotent Notification Requests for more details writeOnly: true  | [optional] 
**Contents** | Pointer to [**NullableStringMap**](StringMap.md) |  | [optional] 
**Headings** | Pointer to [**NullableStringMap**](StringMap.md) |  | [optional] 
**Subtitle** | Pointer to [**NullableStringMap**](StringMap.md) |  | [optional] 
**Data** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: Huawei A custom map of data that is passed back to your app. Same as using Additional Data within the dashboard. Can use up to 2048 bytes of data. Example: {\&quot;abc\&quot;: 123, \&quot;foo\&quot;: \&quot;bar\&quot;, \&quot;event_performed\&quot;: true, \&quot;amount\&quot;: 12.1}  | [optional] 
**HuaweiMsgType** | Pointer to **NullableString** | Channel: Push Notifications Platform: Huawei Use \&quot;data\&quot; or \&quot;message\&quot; depending on the type of notification you are sending. More details in Data &amp; Background Notifications.  | [optional] 
**Url** | Pointer to **NullableString** | Channel: Push Notifications Platform: All The URL to open in the browser when a user clicks on the notification. Note: iOS needs https or updated NSAppTransportSecurity in plist This field supports inline substitutions. Omit if including web_url or app_url Example: https://onesignal.com  | [optional] 
**WebUrl** | Pointer to **NullableString** | Channel: Push Notifications Platform: All Browsers Same as url but only sent to web push platforms. Including Chrome, Firefox, Safari, Opera, etc. Example: https://onesignal.com  | [optional] 
**AppUrl** | Pointer to **NullableString** | Channel: Push Notifications Platform: All Browsers Same as url but only sent to web push platforms. Including iOS, Android, macOS, Windows, ChromeApps, etc. Example: https://onesignal.com  | [optional] 
**IosAttachments** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: iOS 10+ Adds media attachments to notifications. Set as JSON object, key as a media id of your choice and the value as a valid local filename or URL. User must press and hold on the notification to view. Do not set mutable_content to download attachments. The OneSignal SDK does this automatically Example: {\&quot;id1\&quot;: \&quot;https://domain.com/image.jpg\&quot;}  | [optional] 
**TemplateId** | Pointer to **NullableString** | Channel: Push Notifications Platform: All Use a template you setup on our dashboard. The template_id is the UUID found in the URL when viewing a template on our dashboard. Example: be4a8044-bbd6-11e4-a581-000c2940e62c  | [optional] 
**ContentAvailable** | Pointer to **NullableBool** | Channel: Push Notifications Platform: iOS Sending true wakes your app from background to run custom native code (Apple interprets this as content-available&#x3D;1). Note: Not applicable if the app is in the \&quot;force-quit\&quot; state (i.e app was swiped away). Omit the contents field to prevent displaying a visible notification.  | [optional] 
**MutableContent** | Pointer to **bool** | Channel: Push Notifications Platform: iOS 10+ Always defaults to true and cannot be turned off. Allows tracking of notification receives and changing of the notification content in your app before it is displayed. Triggers didReceive(_:withContentHandler:) on your UNNotificationServiceExtension.  | [optional] 
**TargetContentIdentifier** | Pointer to **NullableString** | Channel: Push Notifications Platform: iOS 13+ Use to target a specific experience in your App Clip, or to target your notification to a specific window in a multi-scene App.  | [optional] 
**BigPicture** | Pointer to **NullableString** | Channel: Push Notifications Platform: Android Picture to display in the expanded view. Can be a drawable resource name or a URL.  | [optional] 
**HuaweiBigPicture** | Pointer to **NullableString** | Channel: Push Notifications Platform: Huawei Picture to display in the expanded view. Can be a drawable resource name or a URL.  | [optional] 
**AdmBigPicture** | Pointer to **NullableString** | Channel: Push Notifications Platform: Amazon Picture to display in the expanded view. Can be a drawable resource name or a URL.  | [optional] 
**ChromeBigPicture** | Pointer to **NullableString** | Channel: Push Notifications Platform: ChromeApp Large picture to display below the notification text. Must be a local URL.  | [optional] 
**ChromeWebImage** | Pointer to **NullableString** | Channel: Push Notifications Platform: Chrome 56+ Sets the web push notification&#39;s large image to be shown below the notification&#39;s title and text. Please see Web Push Notification Icons.  | [optional] 
**Buttons** | Pointer to [**[]Button**](Button.md) | Channel: Push Notifications Platform: iOS 8.0+, Android 4.1+, and derivatives like Amazon Buttons to add to the notification. Icon only works for Android. Buttons show in reverse order of array position i.e. Last item in array shows as first button on device. Example: [{\&quot;id\&quot;: \&quot;id2\&quot;, \&quot;text\&quot;: \&quot;second button\&quot;, \&quot;icon\&quot;: \&quot;ic_menu_share\&quot;}, {\&quot;id\&quot;: \&quot;id1\&quot;, \&quot;text\&quot;: \&quot;first button\&quot;, \&quot;icon\&quot;: \&quot;ic_menu_send\&quot;}]  | [optional] 
**WebButtons** | Pointer to [**[]Button**](Button.md) | Channel: Push Notifications Platform: Chrome 48+ Add action buttons to the notification. The id field is required. Example: [{\&quot;id\&quot;: \&quot;like-button\&quot;, \&quot;text\&quot;: \&quot;Like\&quot;, \&quot;icon\&quot;: \&quot;http://i.imgur.com/N8SN8ZS.png\&quot;, \&quot;url\&quot;: \&quot;https://yoursite.com\&quot;}, {\&quot;id\&quot;: \&quot;read-more-button\&quot;, \&quot;text\&quot;: \&quot;Read more\&quot;, \&quot;icon\&quot;: \&quot;http://i.imgur.com/MIxJp1L.png\&quot;, \&quot;url\&quot;: \&quot;https://yoursite.com\&quot;}]  | [optional] 
**IosCategory** | Pointer to **NullableString** | Channel: Push Notifications Platform: iOS Category APS payload, use with registerUserNotificationSettings:categories in your Objective-C / Swift code. Example: calendar category which contains actions like accept and decline iOS 10+ This will trigger your UNNotificationContentExtension whose ID matches this category.  | [optional] 
**AndroidChannelId** | Pointer to **string** | Channel: Push Notifications Platform: Android The Android Oreo Notification Category to send the notification under. See the Category documentation on creating one and getting it&#39;s id.  | [optional] 
**HuaweiChannelId** | Pointer to **NullableString** | Channel: Push Notifications Platform: Huawei The Android Oreo Notification Category to send the notification under. See the Category documentation on creating one and getting it&#39;s id.  | [optional] 
**ExistingAndroidChannelId** | Pointer to **string** | Channel: Push Notifications Platform: Android Use this if you have client side Android Oreo Channels you have already defined in your app with code.  | [optional] 
**HuaweiExistingChannelId** | Pointer to **NullableString** | Channel: Push Notifications Platform: Huawei Use this if you have client side Android Oreo Channels you have already defined in your app with code.  | [optional] 
**AndroidBackgroundLayout** | Pointer to [**BasicNotificationAllOfAndroidBackgroundLayout**](BasicNotificationAllOfAndroidBackgroundLayout.md) |  | [optional] 
**SmallIcon** | Pointer to **NullableString** | Channel: Push Notifications Platform: Android Icon shown in the status bar and on the top left of the notification. If not set a bell icon will be used or ic_stat_onesignal_default if you have set this resource name. See: How to create small icons  | [optional] 
**HuaweiSmallIcon** | Pointer to **NullableString** | Channel: Push Notifications Platform: Huawei Icon shown in the status bar and on the top left of the notification. Use an Android resource path (E.g. /drawable/small_icon). Defaults to your app icon if not set.  | [optional] 
**LargeIcon** | Pointer to **NullableString** | Channel: Push Notifications Platform: Android Can be a drawable resource name or a URL. See: How to create large icons  | [optional] 
**HuaweiLargeIcon** | Pointer to **NullableString** | Channel: Push Notifications Platform: Huawei Can be a drawable resource name or a URL. See: How to create large icons  | [optional] 
**AdmSmallIcon** | Pointer to **NullableString** | Channel: Push Notifications Platform: Amazon If not set a bell icon will be used or ic_stat_onesignal_default if you have set this resource name. See: How to create small icons  | [optional] 
**AdmLargeIcon** | Pointer to **NullableString** | Channel: Push Notifications Platform: Amazon If blank the small_icon is used. Can be a drawable resource name or a URL. See: How to create large icons  | [optional] 
**ChromeWebIcon** | Pointer to **NullableString** | Channel: Push Notifications Platform: Chrome Sets the web push notification&#39;s icon. An image URL linking to a valid image. Common image types are supported; GIF will not animate. We recommend 256x256 (at least 80x80) to display well on high DPI devices. Firefox will also use this icon, unless you specify firefox_icon.  | [optional] 
**ChromeWebBadge** | Pointer to **NullableString** | Channel: Push Notifications Platform: Chrome Sets the web push notification icon for Android devices in the notification shade. Please see Web Push Notification Badge.  | [optional] 
**FirefoxIcon** | Pointer to **NullableString** | Channel: Push Notifications Platform: Firefox Not recommended Few people need to set Firefox-specific icons. We recommend setting chrome_web_icon instead, which Firefox will also use. Sets the web push notification&#39;s icon for Firefox. An image URL linking to a valid image. Common image types are supported; GIF will not animate. We recommend 256x256 (at least 80x80) to display well on high DPI devices.  | [optional] 
**ChromeIcon** | Pointer to **NullableString** | Channel: Push Notifications Platform: ChromeApp This flag is not used for web push For web push, please see chrome_web_icon instead. The local URL to an icon to use. If blank, the app icon will be used.  | [optional] 
**IosSound** | Pointer to **NullableString** | Channel: Push Notifications Platform: iOS Sound file that is included in your app to play instead of the default device notification sound. Pass nil to disable vibration and sound for the notification. Example: \&quot;notification.wav\&quot;  | [optional] 
**AndroidSound** | Pointer to **NullableString** | Channel: Push Notifications Platform: Android &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. Sound file that is included in your app to play instead of the default device notification sound. Pass nil to disable sound for the notification. NOTE: Leave off file extension for Android. Example: \&quot;notification\&quot;  | [optional] 
**HuaweiSound** | Pointer to **NullableString** | Channel: Push Notifications Platform: Huawei &amp;#9888;&amp;#65039;Deprecated, this field ONLY works on EMUI 5 (Android 7 based) and older devices. Please also set Notification Categories / Channels noted above to support EMUI 8 (Android 8 based) devices. Sound file that is included in your app to play instead of the default device notification sound. NOTE: Leave off file extension for and include the full path.  Example: \&quot;/res/raw/notification\&quot;  | [optional] 
**AdmSound** | Pointer to **NullableString** | Channel: Push Notifications Platform: Amazon &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. Sound file that is included in your app to play instead of the default device notification sound. Pass nil to disable sound for the notification. NOTE: Leave off file extension for Android. Example: \&quot;notification\&quot;  | [optional] 
**WpWnsSound** | Pointer to **NullableString** | Channel: Push Notifications Platform: Windows Sound file that is included in your app to play instead of the default device notification sound. Example: \&quot;notification.wav\&quot;  | [optional] 
**AndroidLedColor** | Pointer to **NullableString** | Channel: Push Notifications Platform: Android &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. Sets the devices LED notification light if the device has one. ARGB Hex format. Example(Blue): \&quot;FF0000FF\&quot;  | [optional] 
**HuaweiLedColor** | Pointer to **NullableString** | Channel: Push Notifications Platform: Huawei &amp;#9888;&amp;#65039;Deprecated, this field ONLY works on EMUI 5 (Android 7 based) and older devices. Please also set Notification Categories / Channels noted above to support EMUI 8 (Android 8 based) devices. Sets the devices LED notification light if the device has one. RGB Hex format. Example(Blue): \&quot;0000FF\&quot;  | [optional] 
**AndroidAccentColor** | Pointer to **NullableString** | Channel: Push Notifications Platform: Android Sets the background color of the notification circle to the left of the notification text. Only applies to apps targeting Android API level 21+ on Android 5.0+ devices. Example(Red): \&quot;FFFF0000\&quot;  | [optional] 
**HuaweiAccentColor** | Pointer to **NullableString** | Channel: Push Notifications Platform: Huawei Accent Color used on Action Buttons and Group overflow count. Uses RGB Hex value (E.g. #9900FF). Defaults to device&#39;s theme color if not set.  | [optional] 
**AndroidVisibility** | Pointer to **NullableInt32** | Channel: Push Notifications Platform: Android 5.0_ &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. 1 &#x3D; Public (default) (Shows the full message on the lock screen unless the user has disabled all notifications from showing on the lock screen. Please consider the user and mark private if the contents are.) 0 &#x3D; Private (Hides message contents on lock screen if the user set \&quot;Hide sensitive notification content\&quot; in the system settings) -1 &#x3D; Secret (Notification does not show on the lock screen at all)  | [optional] 
**HuaweiVisibility** | Pointer to **NullableInt32** | Channel: Push Notifications Platform: Huawei &amp;#9888;&amp;#65039;Deprecated, this field ONLY works on EMUI 5 (Android 7 based) and older devices. Please also set Notification Categories / Channels noted above to support EMUI 8 (Android 8 based) devices. 1 &#x3D; Public (default) (Shows the full message on the lock screen unless the user has disabled all notifications from showing on the lock screen. Please consider the user and mark private if the contents are.) 0 &#x3D; Private (Hides message contents on lock screen if the user set \&quot;Hide sensitive notification content\&quot; in the system settings) -1 &#x3D; Secret (Notification does not show on the lock screen at all)  | [optional] 
**IosBadgeType** | Pointer to **NullableString** | Channel: Push Notifications Platform: iOS Describes whether to set or increase/decrease your app&#39;s iOS badge count by the ios_badgeCount specified count. Can specify None, SetTo, or Increase. &#x60;None&#x60; leaves the count unaffected. &#x60;SetTo&#x60; directly sets the badge count to the number specified in ios_badgeCount. &#x60;Increase&#x60; adds the number specified in ios_badgeCount to the total. Use a negative number to decrease the badge count.  | [optional] 
**IosBadgeCount** | Pointer to **NullableInt32** | Channel: Push Notifications Platform: iOS Used with ios_badgeType, describes the value to set or amount to increase/decrease your app&#39;s iOS badge count by. You can use a negative number to decrease the badge count when used with an ios_badgeType of Increase.  | [optional] 
**CollapseId** | Pointer to **string** | Channel: Push Notifications Platform: iOS 10+, Android Only one notification with the same id will be shown on the device. Use the same id to update an existing notification instead of showing a new one. Limit of 64 characters.  | [optional] 
**WebPushTopic** | Pointer to **NullableString** | Channel: Push Notifications Platform: All Browsers Display multiple notifications at once with different topics.  | [optional] 
**ApnsAlert** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: iOS 10+ iOS can localize push notification messages on the client using special parameters such as loc-key. When using the Create Notification endpoint, you must include these parameters inside of a field called apns_alert. Please see Apple&#39;s guide on localizing push notifications to learn more.  | [optional] 
**DelayedOption** | Pointer to **NullableString** | Channel: All Possible values are: timezone (Deliver at a specific time-of-day in each users own timezone) last-active Same as Intelligent Delivery . (Deliver at the same time of day as each user last used your app). If send_after is used, this takes effect after the send_after time has elapsed.  | [optional] 
**DeliveryTimeOfDay** | Pointer to **NullableString** | Channel: All Use with delayed_option&#x3D;timezone. Examples: \&quot;9:00AM\&quot; \&quot;21:45\&quot; \&quot;9:45:30\&quot;  | [optional] 
**Ttl** | Pointer to **NullableInt32** | Channel: Push Notifications Platform: iOS, Android, Chrome, Firefox, Safari, ChromeWeb Time To Live - In seconds. The notification will be expired if the device does not come back online within this time. The default is 259,200 seconds (3 days). Max value to set is 2419200 seconds (28 days).  | [optional] 
**Priority** | Pointer to **NullableInt32** | Channel: Push Notifications Platform: Android, Chrome, ChromeWeb Delivery priority through the push server (example GCM/FCM). Pass 10 for high priority or any other integer for normal priority. Defaults to normal priority for Android and high for iOS. For Android 6.0+ devices setting priority to high will wake the device out of doze mode.  | [optional] 
**ApnsPushTypeOverride** | Pointer to **string** | Channel: Push Notifications Platform: iOS valid values: voip Set the value to voip for sending VoIP Notifications This field maps to the APNS header apns-push-type. Note: alert and background are automatically set by OneSignal  | [optional] 
**ThrottleRatePerMinute** | Pointer to **NullableString** | Channel: All Apps with throttling enabled:   - the parameter value will be used to override the default application throttling value set from the dashboard settings.   - parameter value 0 indicates not to apply throttling to the notification.   - if the parameter is not passed then the default app throttling value will be applied to the notification. Apps with throttling disabled:   - this parameter can be used to throttle delivery for the notification even though throttling is not enabled at the application level. Refer to throttling for more details.  | [optional] 
**AndroidGroup** | Pointer to **NullableString** | Channel: Push Notifications Platform: Android Notifications with the same group will be stacked together using Android&#39;s Notification Grouping feature.  | [optional] 
**AndroidGroupMessage** | Pointer to **NullableString** | Channel: Push Notifications Platform: Android Note: This only works for Android 6 and older. Android 7+ allows full expansion of all message. Summary message to display when 2+ notifications are stacked together. Default is \&quot;# new messages\&quot;. Include $[notif_count] in your message and it will be replaced with the current number. Languages - The value of each key is the message that will be sent to users for that language. \&quot;en\&quot; (English) is required. The key of each hash is either a a 2 character language code or one of zh-Hans/zh-Hant for Simplified or Traditional Chinese. Read more: supported languages. Example: {\&quot;en\&quot;: \&quot;You have $[notif_count] new messages\&quot;}  | [optional] 
**AdmGroup** | Pointer to **NullableString** | Channel: Push Notifications Platform: Amazon Notifications with the same group will be stacked together using Android&#39;s Notification Grouping feature.  | [optional] 
**AdmGroupMessage** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: Amazon Summary message to display when 2+ notifications are stacked together. Default is \&quot;# new messages\&quot;. Include $[notif_count] in your message and it will be replaced with the current number. \&quot;en\&quot; (English) is required. The key of each hash is either a a 2 character language code or one of zh-Hans/zh-Hant for Simplified or Traditional Chinese. The value of each key is the message that will be sent to users for that language. Example: {\&quot;en\&quot;: \&quot;You have $[notif_count] new messages\&quot;}  | [optional] 
**ThreadId** | Pointer to **NullableString** | Channel: Push Notifications Platform: iOS 12+ This parameter is supported in iOS 12 and above. It allows you to group related notifications together. If two notifications have the same thread-id, they will both be added to the same group.  | [optional] 
**SummaryArg** | Pointer to **string** | Channel: Push Notifications Platform: iOS 12+ When using thread_id to create grouped notifications in iOS 12+, you can also control the summary. For example, a grouped notification can say \&quot;12 more notifications from John Doe\&quot;. The summary_arg lets you set the name of the person/thing the notifications are coming from, and will show up as \&quot;X more notifications from summary_arg\&quot;  | [optional] 
**SummaryArgCount** | Pointer to **int32** | Channel: Push Notifications Platform: iOS 12+ When using thread_id, you can also control the count of the number of notifications in the group. For example, if the group already has 12 notifications, and you send a new notification with summary_arg_count &#x3D; 2, the new total will be 14 and the summary will be \&quot;14 more notifications from summary_arg\&quot;  | [optional] 
**EmailSubject** | Pointer to **NullableString** | Channel: Email Required.  The subject of the email.  | [optional] 
**EmailBody** | Pointer to **string** | Channel: Email Required unless template_id is set. HTML suported The body of the email you wish to send. Typically, customers include their own HTML templates here. Must include [unsubscribe_url] in an &lt;a&gt; tag somewhere in the email. Note: any malformed HTML content will be sent to users. Please double-check your HTML is valid.  | [optional] 
**EmailFromName** | Pointer to **NullableString** | Channel: Email The name the email is from. If not specified, will default to \&quot;from name\&quot; set in the OneSignal Dashboard Email Settings.  | [optional] 
**EmailFromAddress** | Pointer to **NullableString** | Channel: Email The email address the email is from. If not specified, will default to \&quot;from email\&quot; set in the OneSignal Dashboard Email Settings.  | [optional] 
**EmailPreheader** | Pointer to **NullableString** | Channel: Email The preheader text of the email. Preheader is the preview text displayed immediately after an email subject that provides additional context about the email content. If not specified, will default to null.  | [optional] 
**IncludeUnsubscribed** | Pointer to **bool** | Channel: Email Default is &#x60;false&#x60;. This field is used to send transactional notifications. If set to &#x60;true&#x60;, this notification will also be sent to unsubscribed emails. If a &#x60;template_id&#x60; is provided, the &#x60;include_unsubscribed&#x60; value from the template will be inherited. If you are using a third-party ESP, this field requires the ESP&#39;s list of unsubscribed emails to be cleared. | [optional] 
**SmsFrom** | Pointer to **NullableString** | Channel: SMS Phone Number used to send SMS. Should be a registered Twilio phone number in E.164 format.  | [optional] 
**SmsMediaUrls** | Pointer to **[]string** | Channel: SMS URLs for the media files to be attached to the SMS content. Limit: 10 media urls with a total max. size of 5MBs.  | [optional] 
**Filters** | Pointer to [**[]Filter**](Filter.md) |  | [optional] 
**CustomData** | Pointer to **map[string]interface{}** | Channel: All JSON object that can be used as a source of message personalization data for fields that support tag variable substitution. Push, SMS: Can accept up to 2048 bytes of valid JSON. Email: Can accept up to 10000 bytes of valid JSON. Example: {\&quot;order_id\&quot;: 123, \&quot;currency\&quot;: \&quot;USD\&quot;, \&quot;amount\&quot;: 25}  | [optional] 
**SendAfter** | Pointer to **NullableTime** | Channel: All Schedule notification for future delivery. API defaults to UTC -1100 Examples: All examples are the exact same date &amp; time. \&quot;Thu Sep 24 2015 14:00:00 GMT-0700 (PDT)\&quot; \&quot;September 24th 2015, 2:00:00 pm UTC-07:00\&quot; \&quot;2015-09-24 14:00:00 GMT-0700\&quot; \&quot;Sept 24 2015 14:00:00 GMT-0700\&quot; \&quot;Thu Sep 24 2015 14:00:00 GMT-0700 (Pacific Daylight Time)\&quot; Note: SMS currently only supports send_after parameter.  | [optional] 

## Methods

### NewNotification

`func NewNotification(appId string, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedSegments

`func (o *Notification) GetIncludedSegments() []string`

GetIncludedSegments returns the IncludedSegments field if non-nil, zero value otherwise.

### GetIncludedSegmentsOk

`func (o *Notification) GetIncludedSegmentsOk() (*[]string, bool)`

GetIncludedSegmentsOk returns a tuple with the IncludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegments

`func (o *Notification) SetIncludedSegments(v []string)`

SetIncludedSegments sets IncludedSegments field to given value.

### HasIncludedSegments

`func (o *Notification) HasIncludedSegments() bool`

HasIncludedSegments returns a boolean if a field has been set.

### GetExcludedSegments

`func (o *Notification) GetExcludedSegments() []string`

GetExcludedSegments returns the ExcludedSegments field if non-nil, zero value otherwise.

### GetExcludedSegmentsOk

`func (o *Notification) GetExcludedSegmentsOk() (*[]string, bool)`

GetExcludedSegmentsOk returns a tuple with the ExcludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegments

`func (o *Notification) SetExcludedSegments(v []string)`

SetExcludedSegments sets ExcludedSegments field to given value.

### HasExcludedSegments

`func (o *Notification) HasExcludedSegments() bool`

HasExcludedSegments returns a boolean if a field has been set.

### GetIncludePlayerIds

`func (o *Notification) GetIncludePlayerIds() []string`

GetIncludePlayerIds returns the IncludePlayerIds field if non-nil, zero value otherwise.

### GetIncludePlayerIdsOk

`func (o *Notification) GetIncludePlayerIdsOk() (*[]string, bool)`

GetIncludePlayerIdsOk returns a tuple with the IncludePlayerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePlayerIds

`func (o *Notification) SetIncludePlayerIds(v []string)`

SetIncludePlayerIds sets IncludePlayerIds field to given value.

### HasIncludePlayerIds

`func (o *Notification) HasIncludePlayerIds() bool`

HasIncludePlayerIds returns a boolean if a field has been set.

### SetIncludePlayerIdsNil

`func (o *Notification) SetIncludePlayerIdsNil(b bool)`

 SetIncludePlayerIdsNil sets the value for IncludePlayerIds to be an explicit nil

### UnsetIncludePlayerIds
`func (o *Notification) UnsetIncludePlayerIds()`

UnsetIncludePlayerIds ensures that no value is present for IncludePlayerIds, not even an explicit nil
### GetIncludeExternalUserIds

`func (o *Notification) GetIncludeExternalUserIds() []string`

GetIncludeExternalUserIds returns the IncludeExternalUserIds field if non-nil, zero value otherwise.

### GetIncludeExternalUserIdsOk

`func (o *Notification) GetIncludeExternalUserIdsOk() (*[]string, bool)`

GetIncludeExternalUserIdsOk returns a tuple with the IncludeExternalUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExternalUserIds

`func (o *Notification) SetIncludeExternalUserIds(v []string)`

SetIncludeExternalUserIds sets IncludeExternalUserIds field to given value.

### HasIncludeExternalUserIds

`func (o *Notification) HasIncludeExternalUserIds() bool`

HasIncludeExternalUserIds returns a boolean if a field has been set.

### SetIncludeExternalUserIdsNil

`func (o *Notification) SetIncludeExternalUserIdsNil(b bool)`

 SetIncludeExternalUserIdsNil sets the value for IncludeExternalUserIds to be an explicit nil

### UnsetIncludeExternalUserIds
`func (o *Notification) UnsetIncludeExternalUserIds()`

UnsetIncludeExternalUserIds ensures that no value is present for IncludeExternalUserIds, not even an explicit nil
### GetIncludeEmailTokens

`func (o *Notification) GetIncludeEmailTokens() []string`

GetIncludeEmailTokens returns the IncludeEmailTokens field if non-nil, zero value otherwise.

### GetIncludeEmailTokensOk

`func (o *Notification) GetIncludeEmailTokensOk() (*[]string, bool)`

GetIncludeEmailTokensOk returns a tuple with the IncludeEmailTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEmailTokens

`func (o *Notification) SetIncludeEmailTokens(v []string)`

SetIncludeEmailTokens sets IncludeEmailTokens field to given value.

### HasIncludeEmailTokens

`func (o *Notification) HasIncludeEmailTokens() bool`

HasIncludeEmailTokens returns a boolean if a field has been set.

### GetIncludePhoneNumbers

`func (o *Notification) GetIncludePhoneNumbers() []string`

GetIncludePhoneNumbers returns the IncludePhoneNumbers field if non-nil, zero value otherwise.

### GetIncludePhoneNumbersOk

`func (o *Notification) GetIncludePhoneNumbersOk() (*[]string, bool)`

GetIncludePhoneNumbersOk returns a tuple with the IncludePhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePhoneNumbers

`func (o *Notification) SetIncludePhoneNumbers(v []string)`

SetIncludePhoneNumbers sets IncludePhoneNumbers field to given value.

### HasIncludePhoneNumbers

`func (o *Notification) HasIncludePhoneNumbers() bool`

HasIncludePhoneNumbers returns a boolean if a field has been set.

### GetIncludeIosTokens

`func (o *Notification) GetIncludeIosTokens() []string`

GetIncludeIosTokens returns the IncludeIosTokens field if non-nil, zero value otherwise.

### GetIncludeIosTokensOk

`func (o *Notification) GetIncludeIosTokensOk() (*[]string, bool)`

GetIncludeIosTokensOk returns a tuple with the IncludeIosTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIosTokens

`func (o *Notification) SetIncludeIosTokens(v []string)`

SetIncludeIosTokens sets IncludeIosTokens field to given value.

### HasIncludeIosTokens

`func (o *Notification) HasIncludeIosTokens() bool`

HasIncludeIosTokens returns a boolean if a field has been set.

### GetIncludeWpWnsUris

`func (o *Notification) GetIncludeWpWnsUris() []string`

GetIncludeWpWnsUris returns the IncludeWpWnsUris field if non-nil, zero value otherwise.

### GetIncludeWpWnsUrisOk

`func (o *Notification) GetIncludeWpWnsUrisOk() (*[]string, bool)`

GetIncludeWpWnsUrisOk returns a tuple with the IncludeWpWnsUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeWpWnsUris

`func (o *Notification) SetIncludeWpWnsUris(v []string)`

SetIncludeWpWnsUris sets IncludeWpWnsUris field to given value.

### HasIncludeWpWnsUris

`func (o *Notification) HasIncludeWpWnsUris() bool`

HasIncludeWpWnsUris returns a boolean if a field has been set.

### GetIncludeAmazonRegIds

`func (o *Notification) GetIncludeAmazonRegIds() []string`

GetIncludeAmazonRegIds returns the IncludeAmazonRegIds field if non-nil, zero value otherwise.

### GetIncludeAmazonRegIdsOk

`func (o *Notification) GetIncludeAmazonRegIdsOk() (*[]string, bool)`

GetIncludeAmazonRegIdsOk returns a tuple with the IncludeAmazonRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAmazonRegIds

`func (o *Notification) SetIncludeAmazonRegIds(v []string)`

SetIncludeAmazonRegIds sets IncludeAmazonRegIds field to given value.

### HasIncludeAmazonRegIds

`func (o *Notification) HasIncludeAmazonRegIds() bool`

HasIncludeAmazonRegIds returns a boolean if a field has been set.

### GetIncludeChromeRegIds

`func (o *Notification) GetIncludeChromeRegIds() []string`

GetIncludeChromeRegIds returns the IncludeChromeRegIds field if non-nil, zero value otherwise.

### GetIncludeChromeRegIdsOk

`func (o *Notification) GetIncludeChromeRegIdsOk() (*[]string, bool)`

GetIncludeChromeRegIdsOk returns a tuple with the IncludeChromeRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChromeRegIds

`func (o *Notification) SetIncludeChromeRegIds(v []string)`

SetIncludeChromeRegIds sets IncludeChromeRegIds field to given value.

### HasIncludeChromeRegIds

`func (o *Notification) HasIncludeChromeRegIds() bool`

HasIncludeChromeRegIds returns a boolean if a field has been set.

### GetIncludeChromeWebRegIds

`func (o *Notification) GetIncludeChromeWebRegIds() []string`

GetIncludeChromeWebRegIds returns the IncludeChromeWebRegIds field if non-nil, zero value otherwise.

### GetIncludeChromeWebRegIdsOk

`func (o *Notification) GetIncludeChromeWebRegIdsOk() (*[]string, bool)`

GetIncludeChromeWebRegIdsOk returns a tuple with the IncludeChromeWebRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChromeWebRegIds

`func (o *Notification) SetIncludeChromeWebRegIds(v []string)`

SetIncludeChromeWebRegIds sets IncludeChromeWebRegIds field to given value.

### HasIncludeChromeWebRegIds

`func (o *Notification) HasIncludeChromeWebRegIds() bool`

HasIncludeChromeWebRegIds returns a boolean if a field has been set.

### GetIncludeAndroidRegIds

`func (o *Notification) GetIncludeAndroidRegIds() []string`

GetIncludeAndroidRegIds returns the IncludeAndroidRegIds field if non-nil, zero value otherwise.

### GetIncludeAndroidRegIdsOk

`func (o *Notification) GetIncludeAndroidRegIdsOk() (*[]string, bool)`

GetIncludeAndroidRegIdsOk returns a tuple with the IncludeAndroidRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAndroidRegIds

`func (o *Notification) SetIncludeAndroidRegIds(v []string)`

SetIncludeAndroidRegIds sets IncludeAndroidRegIds field to given value.

### HasIncludeAndroidRegIds

`func (o *Notification) HasIncludeAndroidRegIds() bool`

HasIncludeAndroidRegIds returns a boolean if a field has been set.

### GetIncludeAliases

`func (o *Notification) GetIncludeAliases() PlayerNotificationTargetIncludeAliases`

GetIncludeAliases returns the IncludeAliases field if non-nil, zero value otherwise.

### GetIncludeAliasesOk

`func (o *Notification) GetIncludeAliasesOk() (*PlayerNotificationTargetIncludeAliases, bool)`

GetIncludeAliasesOk returns a tuple with the IncludeAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAliases

`func (o *Notification) SetIncludeAliases(v PlayerNotificationTargetIncludeAliases)`

SetIncludeAliases sets IncludeAliases field to given value.

### HasIncludeAliases

`func (o *Notification) HasIncludeAliases() bool`

HasIncludeAliases returns a boolean if a field has been set.

### SetIncludeAliasesNil

`func (o *Notification) SetIncludeAliasesNil(b bool)`

 SetIncludeAliasesNil sets the value for IncludeAliases to be an explicit nil

### UnsetIncludeAliases
`func (o *Notification) UnsetIncludeAliases()`

UnsetIncludeAliases ensures that no value is present for IncludeAliases, not even an explicit nil
### GetTargetChannel

`func (o *Notification) GetTargetChannel() string`

GetTargetChannel returns the TargetChannel field if non-nil, zero value otherwise.

### GetTargetChannelOk

`func (o *Notification) GetTargetChannelOk() (*string, bool)`

GetTargetChannelOk returns a tuple with the TargetChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChannel

`func (o *Notification) SetTargetChannel(v string)`

SetTargetChannel sets TargetChannel field to given value.

### HasTargetChannel

`func (o *Notification) HasTargetChannel() bool`

HasTargetChannel returns a boolean if a field has been set.

### GetId

`func (o *Notification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Notification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *Notification) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Notification) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Notification) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Notification) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetName

`func (o *Notification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Notification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Notification) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Notification) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Notification) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Notification) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAggregation

`func (o *Notification) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *Notification) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *Notification) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *Notification) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetIsIos

`func (o *Notification) GetIsIos() bool`

GetIsIos returns the IsIos field if non-nil, zero value otherwise.

### GetIsIosOk

`func (o *Notification) GetIsIosOk() (*bool, bool)`

GetIsIosOk returns a tuple with the IsIos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIos

`func (o *Notification) SetIsIos(v bool)`

SetIsIos sets IsIos field to given value.

### HasIsIos

`func (o *Notification) HasIsIos() bool`

HasIsIos returns a boolean if a field has been set.

### SetIsIosNil

`func (o *Notification) SetIsIosNil(b bool)`

 SetIsIosNil sets the value for IsIos to be an explicit nil

### UnsetIsIos
`func (o *Notification) UnsetIsIos()`

UnsetIsIos ensures that no value is present for IsIos, not even an explicit nil
### GetIsAndroid

`func (o *Notification) GetIsAndroid() bool`

GetIsAndroid returns the IsAndroid field if non-nil, zero value otherwise.

### GetIsAndroidOk

`func (o *Notification) GetIsAndroidOk() (*bool, bool)`

GetIsAndroidOk returns a tuple with the IsAndroid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAndroid

`func (o *Notification) SetIsAndroid(v bool)`

SetIsAndroid sets IsAndroid field to given value.

### HasIsAndroid

`func (o *Notification) HasIsAndroid() bool`

HasIsAndroid returns a boolean if a field has been set.

### SetIsAndroidNil

`func (o *Notification) SetIsAndroidNil(b bool)`

 SetIsAndroidNil sets the value for IsAndroid to be an explicit nil

### UnsetIsAndroid
`func (o *Notification) UnsetIsAndroid()`

UnsetIsAndroid ensures that no value is present for IsAndroid, not even an explicit nil
### GetIsHuawei

`func (o *Notification) GetIsHuawei() bool`

GetIsHuawei returns the IsHuawei field if non-nil, zero value otherwise.

### GetIsHuaweiOk

`func (o *Notification) GetIsHuaweiOk() (*bool, bool)`

GetIsHuaweiOk returns a tuple with the IsHuawei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHuawei

`func (o *Notification) SetIsHuawei(v bool)`

SetIsHuawei sets IsHuawei field to given value.

### HasIsHuawei

`func (o *Notification) HasIsHuawei() bool`

HasIsHuawei returns a boolean if a field has been set.

### SetIsHuaweiNil

`func (o *Notification) SetIsHuaweiNil(b bool)`

 SetIsHuaweiNil sets the value for IsHuawei to be an explicit nil

### UnsetIsHuawei
`func (o *Notification) UnsetIsHuawei()`

UnsetIsHuawei ensures that no value is present for IsHuawei, not even an explicit nil
### GetIsAnyWeb

`func (o *Notification) GetIsAnyWeb() bool`

GetIsAnyWeb returns the IsAnyWeb field if non-nil, zero value otherwise.

### GetIsAnyWebOk

`func (o *Notification) GetIsAnyWebOk() (*bool, bool)`

GetIsAnyWebOk returns a tuple with the IsAnyWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnyWeb

`func (o *Notification) SetIsAnyWeb(v bool)`

SetIsAnyWeb sets IsAnyWeb field to given value.

### HasIsAnyWeb

`func (o *Notification) HasIsAnyWeb() bool`

HasIsAnyWeb returns a boolean if a field has been set.

### SetIsAnyWebNil

`func (o *Notification) SetIsAnyWebNil(b bool)`

 SetIsAnyWebNil sets the value for IsAnyWeb to be an explicit nil

### UnsetIsAnyWeb
`func (o *Notification) UnsetIsAnyWeb()`

UnsetIsAnyWeb ensures that no value is present for IsAnyWeb, not even an explicit nil
### GetIsChromeWeb

`func (o *Notification) GetIsChromeWeb() bool`

GetIsChromeWeb returns the IsChromeWeb field if non-nil, zero value otherwise.

### GetIsChromeWebOk

`func (o *Notification) GetIsChromeWebOk() (*bool, bool)`

GetIsChromeWebOk returns a tuple with the IsChromeWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChromeWeb

`func (o *Notification) SetIsChromeWeb(v bool)`

SetIsChromeWeb sets IsChromeWeb field to given value.

### HasIsChromeWeb

`func (o *Notification) HasIsChromeWeb() bool`

HasIsChromeWeb returns a boolean if a field has been set.

### SetIsChromeWebNil

`func (o *Notification) SetIsChromeWebNil(b bool)`

 SetIsChromeWebNil sets the value for IsChromeWeb to be an explicit nil

### UnsetIsChromeWeb
`func (o *Notification) UnsetIsChromeWeb()`

UnsetIsChromeWeb ensures that no value is present for IsChromeWeb, not even an explicit nil
### GetIsFirefox

`func (o *Notification) GetIsFirefox() bool`

GetIsFirefox returns the IsFirefox field if non-nil, zero value otherwise.

### GetIsFirefoxOk

`func (o *Notification) GetIsFirefoxOk() (*bool, bool)`

GetIsFirefoxOk returns a tuple with the IsFirefox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirefox

`func (o *Notification) SetIsFirefox(v bool)`

SetIsFirefox sets IsFirefox field to given value.

### HasIsFirefox

`func (o *Notification) HasIsFirefox() bool`

HasIsFirefox returns a boolean if a field has been set.

### SetIsFirefoxNil

`func (o *Notification) SetIsFirefoxNil(b bool)`

 SetIsFirefoxNil sets the value for IsFirefox to be an explicit nil

### UnsetIsFirefox
`func (o *Notification) UnsetIsFirefox()`

UnsetIsFirefox ensures that no value is present for IsFirefox, not even an explicit nil
### GetIsSafari

`func (o *Notification) GetIsSafari() bool`

GetIsSafari returns the IsSafari field if non-nil, zero value otherwise.

### GetIsSafariOk

`func (o *Notification) GetIsSafariOk() (*bool, bool)`

GetIsSafariOk returns a tuple with the IsSafari field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafari

`func (o *Notification) SetIsSafari(v bool)`

SetIsSafari sets IsSafari field to given value.

### HasIsSafari

`func (o *Notification) HasIsSafari() bool`

HasIsSafari returns a boolean if a field has been set.

### SetIsSafariNil

`func (o *Notification) SetIsSafariNil(b bool)`

 SetIsSafariNil sets the value for IsSafari to be an explicit nil

### UnsetIsSafari
`func (o *Notification) UnsetIsSafari()`

UnsetIsSafari ensures that no value is present for IsSafari, not even an explicit nil
### GetIsWPWNS

`func (o *Notification) GetIsWPWNS() bool`

GetIsWPWNS returns the IsWPWNS field if non-nil, zero value otherwise.

### GetIsWPWNSOk

`func (o *Notification) GetIsWPWNSOk() (*bool, bool)`

GetIsWPWNSOk returns a tuple with the IsWPWNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWPWNS

`func (o *Notification) SetIsWPWNS(v bool)`

SetIsWPWNS sets IsWPWNS field to given value.

### HasIsWPWNS

`func (o *Notification) HasIsWPWNS() bool`

HasIsWPWNS returns a boolean if a field has been set.

### SetIsWPWNSNil

`func (o *Notification) SetIsWPWNSNil(b bool)`

 SetIsWPWNSNil sets the value for IsWPWNS to be an explicit nil

### UnsetIsWPWNS
`func (o *Notification) UnsetIsWPWNS()`

UnsetIsWPWNS ensures that no value is present for IsWPWNS, not even an explicit nil
### GetIsAdm

`func (o *Notification) GetIsAdm() bool`

GetIsAdm returns the IsAdm field if non-nil, zero value otherwise.

### GetIsAdmOk

`func (o *Notification) GetIsAdmOk() (*bool, bool)`

GetIsAdmOk returns a tuple with the IsAdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdm

`func (o *Notification) SetIsAdm(v bool)`

SetIsAdm sets IsAdm field to given value.

### HasIsAdm

`func (o *Notification) HasIsAdm() bool`

HasIsAdm returns a boolean if a field has been set.

### SetIsAdmNil

`func (o *Notification) SetIsAdmNil(b bool)`

 SetIsAdmNil sets the value for IsAdm to be an explicit nil

### UnsetIsAdm
`func (o *Notification) UnsetIsAdm()`

UnsetIsAdm ensures that no value is present for IsAdm, not even an explicit nil
### GetIsChrome

`func (o *Notification) GetIsChrome() bool`

GetIsChrome returns the IsChrome field if non-nil, zero value otherwise.

### GetIsChromeOk

`func (o *Notification) GetIsChromeOk() (*bool, bool)`

GetIsChromeOk returns a tuple with the IsChrome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChrome

`func (o *Notification) SetIsChrome(v bool)`

SetIsChrome sets IsChrome field to given value.

### HasIsChrome

`func (o *Notification) HasIsChrome() bool`

HasIsChrome returns a boolean if a field has been set.

### SetIsChromeNil

`func (o *Notification) SetIsChromeNil(b bool)`

 SetIsChromeNil sets the value for IsChrome to be an explicit nil

### UnsetIsChrome
`func (o *Notification) UnsetIsChrome()`

UnsetIsChrome ensures that no value is present for IsChrome, not even an explicit nil
### GetChannelForExternalUserIds

`func (o *Notification) GetChannelForExternalUserIds() string`

GetChannelForExternalUserIds returns the ChannelForExternalUserIds field if non-nil, zero value otherwise.

### GetChannelForExternalUserIdsOk

`func (o *Notification) GetChannelForExternalUserIdsOk() (*string, bool)`

GetChannelForExternalUserIdsOk returns a tuple with the ChannelForExternalUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelForExternalUserIds

`func (o *Notification) SetChannelForExternalUserIds(v string)`

SetChannelForExternalUserIds sets ChannelForExternalUserIds field to given value.

### HasChannelForExternalUserIds

`func (o *Notification) HasChannelForExternalUserIds() bool`

HasChannelForExternalUserIds returns a boolean if a field has been set.

### GetAppId

`func (o *Notification) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Notification) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Notification) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetExternalId

`func (o *Notification) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Notification) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Notification) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Notification) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *Notification) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Notification) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetIdempotencyKey

`func (o *Notification) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *Notification) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *Notification) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *Notification) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### SetIdempotencyKeyNil

`func (o *Notification) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *Notification) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
### GetContents

`func (o *Notification) GetContents() StringMap`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *Notification) GetContentsOk() (*StringMap, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *Notification) SetContents(v StringMap)`

SetContents sets Contents field to given value.

### HasContents

`func (o *Notification) HasContents() bool`

HasContents returns a boolean if a field has been set.

### SetContentsNil

`func (o *Notification) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *Notification) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetHeadings

`func (o *Notification) GetHeadings() StringMap`

GetHeadings returns the Headings field if non-nil, zero value otherwise.

### GetHeadingsOk

`func (o *Notification) GetHeadingsOk() (*StringMap, bool)`

GetHeadingsOk returns a tuple with the Headings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadings

`func (o *Notification) SetHeadings(v StringMap)`

SetHeadings sets Headings field to given value.

### HasHeadings

`func (o *Notification) HasHeadings() bool`

HasHeadings returns a boolean if a field has been set.

### SetHeadingsNil

`func (o *Notification) SetHeadingsNil(b bool)`

 SetHeadingsNil sets the value for Headings to be an explicit nil

### UnsetHeadings
`func (o *Notification) UnsetHeadings()`

UnsetHeadings ensures that no value is present for Headings, not even an explicit nil
### GetSubtitle

`func (o *Notification) GetSubtitle() StringMap`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *Notification) GetSubtitleOk() (*StringMap, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *Notification) SetSubtitle(v StringMap)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *Notification) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### SetSubtitleNil

`func (o *Notification) SetSubtitleNil(b bool)`

 SetSubtitleNil sets the value for Subtitle to be an explicit nil

### UnsetSubtitle
`func (o *Notification) UnsetSubtitle()`

UnsetSubtitle ensures that no value is present for Subtitle, not even an explicit nil
### GetData

`func (o *Notification) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Notification) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Notification) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Notification) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Notification) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Notification) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetHuaweiMsgType

`func (o *Notification) GetHuaweiMsgType() string`

GetHuaweiMsgType returns the HuaweiMsgType field if non-nil, zero value otherwise.

### GetHuaweiMsgTypeOk

`func (o *Notification) GetHuaweiMsgTypeOk() (*string, bool)`

GetHuaweiMsgTypeOk returns a tuple with the HuaweiMsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiMsgType

`func (o *Notification) SetHuaweiMsgType(v string)`

SetHuaweiMsgType sets HuaweiMsgType field to given value.

### HasHuaweiMsgType

`func (o *Notification) HasHuaweiMsgType() bool`

HasHuaweiMsgType returns a boolean if a field has been set.

### SetHuaweiMsgTypeNil

`func (o *Notification) SetHuaweiMsgTypeNil(b bool)`

 SetHuaweiMsgTypeNil sets the value for HuaweiMsgType to be an explicit nil

### UnsetHuaweiMsgType
`func (o *Notification) UnsetHuaweiMsgType()`

UnsetHuaweiMsgType ensures that no value is present for HuaweiMsgType, not even an explicit nil
### GetUrl

`func (o *Notification) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Notification) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Notification) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Notification) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *Notification) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Notification) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWebUrl

`func (o *Notification) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *Notification) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *Notification) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *Notification) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *Notification) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *Notification) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetAppUrl

`func (o *Notification) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *Notification) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *Notification) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.

### HasAppUrl

`func (o *Notification) HasAppUrl() bool`

HasAppUrl returns a boolean if a field has been set.

### SetAppUrlNil

`func (o *Notification) SetAppUrlNil(b bool)`

 SetAppUrlNil sets the value for AppUrl to be an explicit nil

### UnsetAppUrl
`func (o *Notification) UnsetAppUrl()`

UnsetAppUrl ensures that no value is present for AppUrl, not even an explicit nil
### GetIosAttachments

`func (o *Notification) GetIosAttachments() map[string]interface{}`

GetIosAttachments returns the IosAttachments field if non-nil, zero value otherwise.

### GetIosAttachmentsOk

`func (o *Notification) GetIosAttachmentsOk() (*map[string]interface{}, bool)`

GetIosAttachmentsOk returns a tuple with the IosAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosAttachments

`func (o *Notification) SetIosAttachments(v map[string]interface{})`

SetIosAttachments sets IosAttachments field to given value.

### HasIosAttachments

`func (o *Notification) HasIosAttachments() bool`

HasIosAttachments returns a boolean if a field has been set.

### SetIosAttachmentsNil

`func (o *Notification) SetIosAttachmentsNil(b bool)`

 SetIosAttachmentsNil sets the value for IosAttachments to be an explicit nil

### UnsetIosAttachments
`func (o *Notification) UnsetIosAttachments()`

UnsetIosAttachments ensures that no value is present for IosAttachments, not even an explicit nil
### GetTemplateId

`func (o *Notification) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Notification) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Notification) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Notification) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *Notification) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *Notification) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetContentAvailable

`func (o *Notification) GetContentAvailable() bool`

GetContentAvailable returns the ContentAvailable field if non-nil, zero value otherwise.

### GetContentAvailableOk

`func (o *Notification) GetContentAvailableOk() (*bool, bool)`

GetContentAvailableOk returns a tuple with the ContentAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAvailable

`func (o *Notification) SetContentAvailable(v bool)`

SetContentAvailable sets ContentAvailable field to given value.

### HasContentAvailable

`func (o *Notification) HasContentAvailable() bool`

HasContentAvailable returns a boolean if a field has been set.

### SetContentAvailableNil

`func (o *Notification) SetContentAvailableNil(b bool)`

 SetContentAvailableNil sets the value for ContentAvailable to be an explicit nil

### UnsetContentAvailable
`func (o *Notification) UnsetContentAvailable()`

UnsetContentAvailable ensures that no value is present for ContentAvailable, not even an explicit nil
### GetMutableContent

`func (o *Notification) GetMutableContent() bool`

GetMutableContent returns the MutableContent field if non-nil, zero value otherwise.

### GetMutableContentOk

`func (o *Notification) GetMutableContentOk() (*bool, bool)`

GetMutableContentOk returns a tuple with the MutableContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutableContent

`func (o *Notification) SetMutableContent(v bool)`

SetMutableContent sets MutableContent field to given value.

### HasMutableContent

`func (o *Notification) HasMutableContent() bool`

HasMutableContent returns a boolean if a field has been set.

### GetTargetContentIdentifier

`func (o *Notification) GetTargetContentIdentifier() string`

GetTargetContentIdentifier returns the TargetContentIdentifier field if non-nil, zero value otherwise.

### GetTargetContentIdentifierOk

`func (o *Notification) GetTargetContentIdentifierOk() (*string, bool)`

GetTargetContentIdentifierOk returns a tuple with the TargetContentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContentIdentifier

`func (o *Notification) SetTargetContentIdentifier(v string)`

SetTargetContentIdentifier sets TargetContentIdentifier field to given value.

### HasTargetContentIdentifier

`func (o *Notification) HasTargetContentIdentifier() bool`

HasTargetContentIdentifier returns a boolean if a field has been set.

### SetTargetContentIdentifierNil

`func (o *Notification) SetTargetContentIdentifierNil(b bool)`

 SetTargetContentIdentifierNil sets the value for TargetContentIdentifier to be an explicit nil

### UnsetTargetContentIdentifier
`func (o *Notification) UnsetTargetContentIdentifier()`

UnsetTargetContentIdentifier ensures that no value is present for TargetContentIdentifier, not even an explicit nil
### GetBigPicture

`func (o *Notification) GetBigPicture() string`

GetBigPicture returns the BigPicture field if non-nil, zero value otherwise.

### GetBigPictureOk

`func (o *Notification) GetBigPictureOk() (*string, bool)`

GetBigPictureOk returns a tuple with the BigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigPicture

`func (o *Notification) SetBigPicture(v string)`

SetBigPicture sets BigPicture field to given value.

### HasBigPicture

`func (o *Notification) HasBigPicture() bool`

HasBigPicture returns a boolean if a field has been set.

### SetBigPictureNil

`func (o *Notification) SetBigPictureNil(b bool)`

 SetBigPictureNil sets the value for BigPicture to be an explicit nil

### UnsetBigPicture
`func (o *Notification) UnsetBigPicture()`

UnsetBigPicture ensures that no value is present for BigPicture, not even an explicit nil
### GetHuaweiBigPicture

`func (o *Notification) GetHuaweiBigPicture() string`

GetHuaweiBigPicture returns the HuaweiBigPicture field if non-nil, zero value otherwise.

### GetHuaweiBigPictureOk

`func (o *Notification) GetHuaweiBigPictureOk() (*string, bool)`

GetHuaweiBigPictureOk returns a tuple with the HuaweiBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiBigPicture

`func (o *Notification) SetHuaweiBigPicture(v string)`

SetHuaweiBigPicture sets HuaweiBigPicture field to given value.

### HasHuaweiBigPicture

`func (o *Notification) HasHuaweiBigPicture() bool`

HasHuaweiBigPicture returns a boolean if a field has been set.

### SetHuaweiBigPictureNil

`func (o *Notification) SetHuaweiBigPictureNil(b bool)`

 SetHuaweiBigPictureNil sets the value for HuaweiBigPicture to be an explicit nil

### UnsetHuaweiBigPicture
`func (o *Notification) UnsetHuaweiBigPicture()`

UnsetHuaweiBigPicture ensures that no value is present for HuaweiBigPicture, not even an explicit nil
### GetAdmBigPicture

`func (o *Notification) GetAdmBigPicture() string`

GetAdmBigPicture returns the AdmBigPicture field if non-nil, zero value otherwise.

### GetAdmBigPictureOk

`func (o *Notification) GetAdmBigPictureOk() (*string, bool)`

GetAdmBigPictureOk returns a tuple with the AdmBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmBigPicture

`func (o *Notification) SetAdmBigPicture(v string)`

SetAdmBigPicture sets AdmBigPicture field to given value.

### HasAdmBigPicture

`func (o *Notification) HasAdmBigPicture() bool`

HasAdmBigPicture returns a boolean if a field has been set.

### SetAdmBigPictureNil

`func (o *Notification) SetAdmBigPictureNil(b bool)`

 SetAdmBigPictureNil sets the value for AdmBigPicture to be an explicit nil

### UnsetAdmBigPicture
`func (o *Notification) UnsetAdmBigPicture()`

UnsetAdmBigPicture ensures that no value is present for AdmBigPicture, not even an explicit nil
### GetChromeBigPicture

`func (o *Notification) GetChromeBigPicture() string`

GetChromeBigPicture returns the ChromeBigPicture field if non-nil, zero value otherwise.

### GetChromeBigPictureOk

`func (o *Notification) GetChromeBigPictureOk() (*string, bool)`

GetChromeBigPictureOk returns a tuple with the ChromeBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeBigPicture

`func (o *Notification) SetChromeBigPicture(v string)`

SetChromeBigPicture sets ChromeBigPicture field to given value.

### HasChromeBigPicture

`func (o *Notification) HasChromeBigPicture() bool`

HasChromeBigPicture returns a boolean if a field has been set.

### SetChromeBigPictureNil

`func (o *Notification) SetChromeBigPictureNil(b bool)`

 SetChromeBigPictureNil sets the value for ChromeBigPicture to be an explicit nil

### UnsetChromeBigPicture
`func (o *Notification) UnsetChromeBigPicture()`

UnsetChromeBigPicture ensures that no value is present for ChromeBigPicture, not even an explicit nil
### GetChromeWebImage

`func (o *Notification) GetChromeWebImage() string`

GetChromeWebImage returns the ChromeWebImage field if non-nil, zero value otherwise.

### GetChromeWebImageOk

`func (o *Notification) GetChromeWebImageOk() (*string, bool)`

GetChromeWebImageOk returns a tuple with the ChromeWebImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebImage

`func (o *Notification) SetChromeWebImage(v string)`

SetChromeWebImage sets ChromeWebImage field to given value.

### HasChromeWebImage

`func (o *Notification) HasChromeWebImage() bool`

HasChromeWebImage returns a boolean if a field has been set.

### SetChromeWebImageNil

`func (o *Notification) SetChromeWebImageNil(b bool)`

 SetChromeWebImageNil sets the value for ChromeWebImage to be an explicit nil

### UnsetChromeWebImage
`func (o *Notification) UnsetChromeWebImage()`

UnsetChromeWebImage ensures that no value is present for ChromeWebImage, not even an explicit nil
### GetButtons

`func (o *Notification) GetButtons() []Button`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *Notification) GetButtonsOk() (*[]Button, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *Notification) SetButtons(v []Button)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *Notification) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### SetButtonsNil

`func (o *Notification) SetButtonsNil(b bool)`

 SetButtonsNil sets the value for Buttons to be an explicit nil

### UnsetButtons
`func (o *Notification) UnsetButtons()`

UnsetButtons ensures that no value is present for Buttons, not even an explicit nil
### GetWebButtons

`func (o *Notification) GetWebButtons() []Button`

GetWebButtons returns the WebButtons field if non-nil, zero value otherwise.

### GetWebButtonsOk

`func (o *Notification) GetWebButtonsOk() (*[]Button, bool)`

GetWebButtonsOk returns a tuple with the WebButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebButtons

`func (o *Notification) SetWebButtons(v []Button)`

SetWebButtons sets WebButtons field to given value.

### HasWebButtons

`func (o *Notification) HasWebButtons() bool`

HasWebButtons returns a boolean if a field has been set.

### SetWebButtonsNil

`func (o *Notification) SetWebButtonsNil(b bool)`

 SetWebButtonsNil sets the value for WebButtons to be an explicit nil

### UnsetWebButtons
`func (o *Notification) UnsetWebButtons()`

UnsetWebButtons ensures that no value is present for WebButtons, not even an explicit nil
### GetIosCategory

`func (o *Notification) GetIosCategory() string`

GetIosCategory returns the IosCategory field if non-nil, zero value otherwise.

### GetIosCategoryOk

`func (o *Notification) GetIosCategoryOk() (*string, bool)`

GetIosCategoryOk returns a tuple with the IosCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosCategory

`func (o *Notification) SetIosCategory(v string)`

SetIosCategory sets IosCategory field to given value.

### HasIosCategory

`func (o *Notification) HasIosCategory() bool`

HasIosCategory returns a boolean if a field has been set.

### SetIosCategoryNil

`func (o *Notification) SetIosCategoryNil(b bool)`

 SetIosCategoryNil sets the value for IosCategory to be an explicit nil

### UnsetIosCategory
`func (o *Notification) UnsetIosCategory()`

UnsetIosCategory ensures that no value is present for IosCategory, not even an explicit nil
### GetAndroidChannelId

`func (o *Notification) GetAndroidChannelId() string`

GetAndroidChannelId returns the AndroidChannelId field if non-nil, zero value otherwise.

### GetAndroidChannelIdOk

`func (o *Notification) GetAndroidChannelIdOk() (*string, bool)`

GetAndroidChannelIdOk returns a tuple with the AndroidChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidChannelId

`func (o *Notification) SetAndroidChannelId(v string)`

SetAndroidChannelId sets AndroidChannelId field to given value.

### HasAndroidChannelId

`func (o *Notification) HasAndroidChannelId() bool`

HasAndroidChannelId returns a boolean if a field has been set.

### GetHuaweiChannelId

`func (o *Notification) GetHuaweiChannelId() string`

GetHuaweiChannelId returns the HuaweiChannelId field if non-nil, zero value otherwise.

### GetHuaweiChannelIdOk

`func (o *Notification) GetHuaweiChannelIdOk() (*string, bool)`

GetHuaweiChannelIdOk returns a tuple with the HuaweiChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiChannelId

`func (o *Notification) SetHuaweiChannelId(v string)`

SetHuaweiChannelId sets HuaweiChannelId field to given value.

### HasHuaweiChannelId

`func (o *Notification) HasHuaweiChannelId() bool`

HasHuaweiChannelId returns a boolean if a field has been set.

### SetHuaweiChannelIdNil

`func (o *Notification) SetHuaweiChannelIdNil(b bool)`

 SetHuaweiChannelIdNil sets the value for HuaweiChannelId to be an explicit nil

### UnsetHuaweiChannelId
`func (o *Notification) UnsetHuaweiChannelId()`

UnsetHuaweiChannelId ensures that no value is present for HuaweiChannelId, not even an explicit nil
### GetExistingAndroidChannelId

`func (o *Notification) GetExistingAndroidChannelId() string`

GetExistingAndroidChannelId returns the ExistingAndroidChannelId field if non-nil, zero value otherwise.

### GetExistingAndroidChannelIdOk

`func (o *Notification) GetExistingAndroidChannelIdOk() (*string, bool)`

GetExistingAndroidChannelIdOk returns a tuple with the ExistingAndroidChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingAndroidChannelId

`func (o *Notification) SetExistingAndroidChannelId(v string)`

SetExistingAndroidChannelId sets ExistingAndroidChannelId field to given value.

### HasExistingAndroidChannelId

`func (o *Notification) HasExistingAndroidChannelId() bool`

HasExistingAndroidChannelId returns a boolean if a field has been set.

### GetHuaweiExistingChannelId

`func (o *Notification) GetHuaweiExistingChannelId() string`

GetHuaweiExistingChannelId returns the HuaweiExistingChannelId field if non-nil, zero value otherwise.

### GetHuaweiExistingChannelIdOk

`func (o *Notification) GetHuaweiExistingChannelIdOk() (*string, bool)`

GetHuaweiExistingChannelIdOk returns a tuple with the HuaweiExistingChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiExistingChannelId

`func (o *Notification) SetHuaweiExistingChannelId(v string)`

SetHuaweiExistingChannelId sets HuaweiExistingChannelId field to given value.

### HasHuaweiExistingChannelId

`func (o *Notification) HasHuaweiExistingChannelId() bool`

HasHuaweiExistingChannelId returns a boolean if a field has been set.

### SetHuaweiExistingChannelIdNil

`func (o *Notification) SetHuaweiExistingChannelIdNil(b bool)`

 SetHuaweiExistingChannelIdNil sets the value for HuaweiExistingChannelId to be an explicit nil

### UnsetHuaweiExistingChannelId
`func (o *Notification) UnsetHuaweiExistingChannelId()`

UnsetHuaweiExistingChannelId ensures that no value is present for HuaweiExistingChannelId, not even an explicit nil
### GetAndroidBackgroundLayout

`func (o *Notification) GetAndroidBackgroundLayout() BasicNotificationAllOfAndroidBackgroundLayout`

GetAndroidBackgroundLayout returns the AndroidBackgroundLayout field if non-nil, zero value otherwise.

### GetAndroidBackgroundLayoutOk

`func (o *Notification) GetAndroidBackgroundLayoutOk() (*BasicNotificationAllOfAndroidBackgroundLayout, bool)`

GetAndroidBackgroundLayoutOk returns a tuple with the AndroidBackgroundLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidBackgroundLayout

`func (o *Notification) SetAndroidBackgroundLayout(v BasicNotificationAllOfAndroidBackgroundLayout)`

SetAndroidBackgroundLayout sets AndroidBackgroundLayout field to given value.

### HasAndroidBackgroundLayout

`func (o *Notification) HasAndroidBackgroundLayout() bool`

HasAndroidBackgroundLayout returns a boolean if a field has been set.

### GetSmallIcon

`func (o *Notification) GetSmallIcon() string`

GetSmallIcon returns the SmallIcon field if non-nil, zero value otherwise.

### GetSmallIconOk

`func (o *Notification) GetSmallIconOk() (*string, bool)`

GetSmallIconOk returns a tuple with the SmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallIcon

`func (o *Notification) SetSmallIcon(v string)`

SetSmallIcon sets SmallIcon field to given value.

### HasSmallIcon

`func (o *Notification) HasSmallIcon() bool`

HasSmallIcon returns a boolean if a field has been set.

### SetSmallIconNil

`func (o *Notification) SetSmallIconNil(b bool)`

 SetSmallIconNil sets the value for SmallIcon to be an explicit nil

### UnsetSmallIcon
`func (o *Notification) UnsetSmallIcon()`

UnsetSmallIcon ensures that no value is present for SmallIcon, not even an explicit nil
### GetHuaweiSmallIcon

`func (o *Notification) GetHuaweiSmallIcon() string`

GetHuaweiSmallIcon returns the HuaweiSmallIcon field if non-nil, zero value otherwise.

### GetHuaweiSmallIconOk

`func (o *Notification) GetHuaweiSmallIconOk() (*string, bool)`

GetHuaweiSmallIconOk returns a tuple with the HuaweiSmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiSmallIcon

`func (o *Notification) SetHuaweiSmallIcon(v string)`

SetHuaweiSmallIcon sets HuaweiSmallIcon field to given value.

### HasHuaweiSmallIcon

`func (o *Notification) HasHuaweiSmallIcon() bool`

HasHuaweiSmallIcon returns a boolean if a field has been set.

### SetHuaweiSmallIconNil

`func (o *Notification) SetHuaweiSmallIconNil(b bool)`

 SetHuaweiSmallIconNil sets the value for HuaweiSmallIcon to be an explicit nil

### UnsetHuaweiSmallIcon
`func (o *Notification) UnsetHuaweiSmallIcon()`

UnsetHuaweiSmallIcon ensures that no value is present for HuaweiSmallIcon, not even an explicit nil
### GetLargeIcon

`func (o *Notification) GetLargeIcon() string`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *Notification) GetLargeIconOk() (*string, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeIcon

`func (o *Notification) SetLargeIcon(v string)`

SetLargeIcon sets LargeIcon field to given value.

### HasLargeIcon

`func (o *Notification) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### SetLargeIconNil

`func (o *Notification) SetLargeIconNil(b bool)`

 SetLargeIconNil sets the value for LargeIcon to be an explicit nil

### UnsetLargeIcon
`func (o *Notification) UnsetLargeIcon()`

UnsetLargeIcon ensures that no value is present for LargeIcon, not even an explicit nil
### GetHuaweiLargeIcon

`func (o *Notification) GetHuaweiLargeIcon() string`

GetHuaweiLargeIcon returns the HuaweiLargeIcon field if non-nil, zero value otherwise.

### GetHuaweiLargeIconOk

`func (o *Notification) GetHuaweiLargeIconOk() (*string, bool)`

GetHuaweiLargeIconOk returns a tuple with the HuaweiLargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiLargeIcon

`func (o *Notification) SetHuaweiLargeIcon(v string)`

SetHuaweiLargeIcon sets HuaweiLargeIcon field to given value.

### HasHuaweiLargeIcon

`func (o *Notification) HasHuaweiLargeIcon() bool`

HasHuaweiLargeIcon returns a boolean if a field has been set.

### SetHuaweiLargeIconNil

`func (o *Notification) SetHuaweiLargeIconNil(b bool)`

 SetHuaweiLargeIconNil sets the value for HuaweiLargeIcon to be an explicit nil

### UnsetHuaweiLargeIcon
`func (o *Notification) UnsetHuaweiLargeIcon()`

UnsetHuaweiLargeIcon ensures that no value is present for HuaweiLargeIcon, not even an explicit nil
### GetAdmSmallIcon

`func (o *Notification) GetAdmSmallIcon() string`

GetAdmSmallIcon returns the AdmSmallIcon field if non-nil, zero value otherwise.

### GetAdmSmallIconOk

`func (o *Notification) GetAdmSmallIconOk() (*string, bool)`

GetAdmSmallIconOk returns a tuple with the AdmSmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmSmallIcon

`func (o *Notification) SetAdmSmallIcon(v string)`

SetAdmSmallIcon sets AdmSmallIcon field to given value.

### HasAdmSmallIcon

`func (o *Notification) HasAdmSmallIcon() bool`

HasAdmSmallIcon returns a boolean if a field has been set.

### SetAdmSmallIconNil

`func (o *Notification) SetAdmSmallIconNil(b bool)`

 SetAdmSmallIconNil sets the value for AdmSmallIcon to be an explicit nil

### UnsetAdmSmallIcon
`func (o *Notification) UnsetAdmSmallIcon()`

UnsetAdmSmallIcon ensures that no value is present for AdmSmallIcon, not even an explicit nil
### GetAdmLargeIcon

`func (o *Notification) GetAdmLargeIcon() string`

GetAdmLargeIcon returns the AdmLargeIcon field if non-nil, zero value otherwise.

### GetAdmLargeIconOk

`func (o *Notification) GetAdmLargeIconOk() (*string, bool)`

GetAdmLargeIconOk returns a tuple with the AdmLargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmLargeIcon

`func (o *Notification) SetAdmLargeIcon(v string)`

SetAdmLargeIcon sets AdmLargeIcon field to given value.

### HasAdmLargeIcon

`func (o *Notification) HasAdmLargeIcon() bool`

HasAdmLargeIcon returns a boolean if a field has been set.

### SetAdmLargeIconNil

`func (o *Notification) SetAdmLargeIconNil(b bool)`

 SetAdmLargeIconNil sets the value for AdmLargeIcon to be an explicit nil

### UnsetAdmLargeIcon
`func (o *Notification) UnsetAdmLargeIcon()`

UnsetAdmLargeIcon ensures that no value is present for AdmLargeIcon, not even an explicit nil
### GetChromeWebIcon

`func (o *Notification) GetChromeWebIcon() string`

GetChromeWebIcon returns the ChromeWebIcon field if non-nil, zero value otherwise.

### GetChromeWebIconOk

`func (o *Notification) GetChromeWebIconOk() (*string, bool)`

GetChromeWebIconOk returns a tuple with the ChromeWebIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebIcon

`func (o *Notification) SetChromeWebIcon(v string)`

SetChromeWebIcon sets ChromeWebIcon field to given value.

### HasChromeWebIcon

`func (o *Notification) HasChromeWebIcon() bool`

HasChromeWebIcon returns a boolean if a field has been set.

### SetChromeWebIconNil

`func (o *Notification) SetChromeWebIconNil(b bool)`

 SetChromeWebIconNil sets the value for ChromeWebIcon to be an explicit nil

### UnsetChromeWebIcon
`func (o *Notification) UnsetChromeWebIcon()`

UnsetChromeWebIcon ensures that no value is present for ChromeWebIcon, not even an explicit nil
### GetChromeWebBadge

`func (o *Notification) GetChromeWebBadge() string`

GetChromeWebBadge returns the ChromeWebBadge field if non-nil, zero value otherwise.

### GetChromeWebBadgeOk

`func (o *Notification) GetChromeWebBadgeOk() (*string, bool)`

GetChromeWebBadgeOk returns a tuple with the ChromeWebBadge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebBadge

`func (o *Notification) SetChromeWebBadge(v string)`

SetChromeWebBadge sets ChromeWebBadge field to given value.

### HasChromeWebBadge

`func (o *Notification) HasChromeWebBadge() bool`

HasChromeWebBadge returns a boolean if a field has been set.

### SetChromeWebBadgeNil

`func (o *Notification) SetChromeWebBadgeNil(b bool)`

 SetChromeWebBadgeNil sets the value for ChromeWebBadge to be an explicit nil

### UnsetChromeWebBadge
`func (o *Notification) UnsetChromeWebBadge()`

UnsetChromeWebBadge ensures that no value is present for ChromeWebBadge, not even an explicit nil
### GetFirefoxIcon

`func (o *Notification) GetFirefoxIcon() string`

GetFirefoxIcon returns the FirefoxIcon field if non-nil, zero value otherwise.

### GetFirefoxIconOk

`func (o *Notification) GetFirefoxIconOk() (*string, bool)`

GetFirefoxIconOk returns a tuple with the FirefoxIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirefoxIcon

`func (o *Notification) SetFirefoxIcon(v string)`

SetFirefoxIcon sets FirefoxIcon field to given value.

### HasFirefoxIcon

`func (o *Notification) HasFirefoxIcon() bool`

HasFirefoxIcon returns a boolean if a field has been set.

### SetFirefoxIconNil

`func (o *Notification) SetFirefoxIconNil(b bool)`

 SetFirefoxIconNil sets the value for FirefoxIcon to be an explicit nil

### UnsetFirefoxIcon
`func (o *Notification) UnsetFirefoxIcon()`

UnsetFirefoxIcon ensures that no value is present for FirefoxIcon, not even an explicit nil
### GetChromeIcon

`func (o *Notification) GetChromeIcon() string`

GetChromeIcon returns the ChromeIcon field if non-nil, zero value otherwise.

### GetChromeIconOk

`func (o *Notification) GetChromeIconOk() (*string, bool)`

GetChromeIconOk returns a tuple with the ChromeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeIcon

`func (o *Notification) SetChromeIcon(v string)`

SetChromeIcon sets ChromeIcon field to given value.

### HasChromeIcon

`func (o *Notification) HasChromeIcon() bool`

HasChromeIcon returns a boolean if a field has been set.

### SetChromeIconNil

`func (o *Notification) SetChromeIconNil(b bool)`

 SetChromeIconNil sets the value for ChromeIcon to be an explicit nil

### UnsetChromeIcon
`func (o *Notification) UnsetChromeIcon()`

UnsetChromeIcon ensures that no value is present for ChromeIcon, not even an explicit nil
### GetIosSound

`func (o *Notification) GetIosSound() string`

GetIosSound returns the IosSound field if non-nil, zero value otherwise.

### GetIosSoundOk

`func (o *Notification) GetIosSoundOk() (*string, bool)`

GetIosSoundOk returns a tuple with the IosSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosSound

`func (o *Notification) SetIosSound(v string)`

SetIosSound sets IosSound field to given value.

### HasIosSound

`func (o *Notification) HasIosSound() bool`

HasIosSound returns a boolean if a field has been set.

### SetIosSoundNil

`func (o *Notification) SetIosSoundNil(b bool)`

 SetIosSoundNil sets the value for IosSound to be an explicit nil

### UnsetIosSound
`func (o *Notification) UnsetIosSound()`

UnsetIosSound ensures that no value is present for IosSound, not even an explicit nil
### GetAndroidSound

`func (o *Notification) GetAndroidSound() string`

GetAndroidSound returns the AndroidSound field if non-nil, zero value otherwise.

### GetAndroidSoundOk

`func (o *Notification) GetAndroidSoundOk() (*string, bool)`

GetAndroidSoundOk returns a tuple with the AndroidSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidSound

`func (o *Notification) SetAndroidSound(v string)`

SetAndroidSound sets AndroidSound field to given value.

### HasAndroidSound

`func (o *Notification) HasAndroidSound() bool`

HasAndroidSound returns a boolean if a field has been set.

### SetAndroidSoundNil

`func (o *Notification) SetAndroidSoundNil(b bool)`

 SetAndroidSoundNil sets the value for AndroidSound to be an explicit nil

### UnsetAndroidSound
`func (o *Notification) UnsetAndroidSound()`

UnsetAndroidSound ensures that no value is present for AndroidSound, not even an explicit nil
### GetHuaweiSound

`func (o *Notification) GetHuaweiSound() string`

GetHuaweiSound returns the HuaweiSound field if non-nil, zero value otherwise.

### GetHuaweiSoundOk

`func (o *Notification) GetHuaweiSoundOk() (*string, bool)`

GetHuaweiSoundOk returns a tuple with the HuaweiSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiSound

`func (o *Notification) SetHuaweiSound(v string)`

SetHuaweiSound sets HuaweiSound field to given value.

### HasHuaweiSound

`func (o *Notification) HasHuaweiSound() bool`

HasHuaweiSound returns a boolean if a field has been set.

### SetHuaweiSoundNil

`func (o *Notification) SetHuaweiSoundNil(b bool)`

 SetHuaweiSoundNil sets the value for HuaweiSound to be an explicit nil

### UnsetHuaweiSound
`func (o *Notification) UnsetHuaweiSound()`

UnsetHuaweiSound ensures that no value is present for HuaweiSound, not even an explicit nil
### GetAdmSound

`func (o *Notification) GetAdmSound() string`

GetAdmSound returns the AdmSound field if non-nil, zero value otherwise.

### GetAdmSoundOk

`func (o *Notification) GetAdmSoundOk() (*string, bool)`

GetAdmSoundOk returns a tuple with the AdmSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmSound

`func (o *Notification) SetAdmSound(v string)`

SetAdmSound sets AdmSound field to given value.

### HasAdmSound

`func (o *Notification) HasAdmSound() bool`

HasAdmSound returns a boolean if a field has been set.

### SetAdmSoundNil

`func (o *Notification) SetAdmSoundNil(b bool)`

 SetAdmSoundNil sets the value for AdmSound to be an explicit nil

### UnsetAdmSound
`func (o *Notification) UnsetAdmSound()`

UnsetAdmSound ensures that no value is present for AdmSound, not even an explicit nil
### GetWpWnsSound

`func (o *Notification) GetWpWnsSound() string`

GetWpWnsSound returns the WpWnsSound field if non-nil, zero value otherwise.

### GetWpWnsSoundOk

`func (o *Notification) GetWpWnsSoundOk() (*string, bool)`

GetWpWnsSoundOk returns a tuple with the WpWnsSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpWnsSound

`func (o *Notification) SetWpWnsSound(v string)`

SetWpWnsSound sets WpWnsSound field to given value.

### HasWpWnsSound

`func (o *Notification) HasWpWnsSound() bool`

HasWpWnsSound returns a boolean if a field has been set.

### SetWpWnsSoundNil

`func (o *Notification) SetWpWnsSoundNil(b bool)`

 SetWpWnsSoundNil sets the value for WpWnsSound to be an explicit nil

### UnsetWpWnsSound
`func (o *Notification) UnsetWpWnsSound()`

UnsetWpWnsSound ensures that no value is present for WpWnsSound, not even an explicit nil
### GetAndroidLedColor

`func (o *Notification) GetAndroidLedColor() string`

GetAndroidLedColor returns the AndroidLedColor field if non-nil, zero value otherwise.

### GetAndroidLedColorOk

`func (o *Notification) GetAndroidLedColorOk() (*string, bool)`

GetAndroidLedColorOk returns a tuple with the AndroidLedColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidLedColor

`func (o *Notification) SetAndroidLedColor(v string)`

SetAndroidLedColor sets AndroidLedColor field to given value.

### HasAndroidLedColor

`func (o *Notification) HasAndroidLedColor() bool`

HasAndroidLedColor returns a boolean if a field has been set.

### SetAndroidLedColorNil

`func (o *Notification) SetAndroidLedColorNil(b bool)`

 SetAndroidLedColorNil sets the value for AndroidLedColor to be an explicit nil

### UnsetAndroidLedColor
`func (o *Notification) UnsetAndroidLedColor()`

UnsetAndroidLedColor ensures that no value is present for AndroidLedColor, not even an explicit nil
### GetHuaweiLedColor

`func (o *Notification) GetHuaweiLedColor() string`

GetHuaweiLedColor returns the HuaweiLedColor field if non-nil, zero value otherwise.

### GetHuaweiLedColorOk

`func (o *Notification) GetHuaweiLedColorOk() (*string, bool)`

GetHuaweiLedColorOk returns a tuple with the HuaweiLedColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiLedColor

`func (o *Notification) SetHuaweiLedColor(v string)`

SetHuaweiLedColor sets HuaweiLedColor field to given value.

### HasHuaweiLedColor

`func (o *Notification) HasHuaweiLedColor() bool`

HasHuaweiLedColor returns a boolean if a field has been set.

### SetHuaweiLedColorNil

`func (o *Notification) SetHuaweiLedColorNil(b bool)`

 SetHuaweiLedColorNil sets the value for HuaweiLedColor to be an explicit nil

### UnsetHuaweiLedColor
`func (o *Notification) UnsetHuaweiLedColor()`

UnsetHuaweiLedColor ensures that no value is present for HuaweiLedColor, not even an explicit nil
### GetAndroidAccentColor

`func (o *Notification) GetAndroidAccentColor() string`

GetAndroidAccentColor returns the AndroidAccentColor field if non-nil, zero value otherwise.

### GetAndroidAccentColorOk

`func (o *Notification) GetAndroidAccentColorOk() (*string, bool)`

GetAndroidAccentColorOk returns a tuple with the AndroidAccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidAccentColor

`func (o *Notification) SetAndroidAccentColor(v string)`

SetAndroidAccentColor sets AndroidAccentColor field to given value.

### HasAndroidAccentColor

`func (o *Notification) HasAndroidAccentColor() bool`

HasAndroidAccentColor returns a boolean if a field has been set.

### SetAndroidAccentColorNil

`func (o *Notification) SetAndroidAccentColorNil(b bool)`

 SetAndroidAccentColorNil sets the value for AndroidAccentColor to be an explicit nil

### UnsetAndroidAccentColor
`func (o *Notification) UnsetAndroidAccentColor()`

UnsetAndroidAccentColor ensures that no value is present for AndroidAccentColor, not even an explicit nil
### GetHuaweiAccentColor

`func (o *Notification) GetHuaweiAccentColor() string`

GetHuaweiAccentColor returns the HuaweiAccentColor field if non-nil, zero value otherwise.

### GetHuaweiAccentColorOk

`func (o *Notification) GetHuaweiAccentColorOk() (*string, bool)`

GetHuaweiAccentColorOk returns a tuple with the HuaweiAccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiAccentColor

`func (o *Notification) SetHuaweiAccentColor(v string)`

SetHuaweiAccentColor sets HuaweiAccentColor field to given value.

### HasHuaweiAccentColor

`func (o *Notification) HasHuaweiAccentColor() bool`

HasHuaweiAccentColor returns a boolean if a field has been set.

### SetHuaweiAccentColorNil

`func (o *Notification) SetHuaweiAccentColorNil(b bool)`

 SetHuaweiAccentColorNil sets the value for HuaweiAccentColor to be an explicit nil

### UnsetHuaweiAccentColor
`func (o *Notification) UnsetHuaweiAccentColor()`

UnsetHuaweiAccentColor ensures that no value is present for HuaweiAccentColor, not even an explicit nil
### GetAndroidVisibility

`func (o *Notification) GetAndroidVisibility() int32`

GetAndroidVisibility returns the AndroidVisibility field if non-nil, zero value otherwise.

### GetAndroidVisibilityOk

`func (o *Notification) GetAndroidVisibilityOk() (*int32, bool)`

GetAndroidVisibilityOk returns a tuple with the AndroidVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidVisibility

`func (o *Notification) SetAndroidVisibility(v int32)`

SetAndroidVisibility sets AndroidVisibility field to given value.

### HasAndroidVisibility

`func (o *Notification) HasAndroidVisibility() bool`

HasAndroidVisibility returns a boolean if a field has been set.

### SetAndroidVisibilityNil

`func (o *Notification) SetAndroidVisibilityNil(b bool)`

 SetAndroidVisibilityNil sets the value for AndroidVisibility to be an explicit nil

### UnsetAndroidVisibility
`func (o *Notification) UnsetAndroidVisibility()`

UnsetAndroidVisibility ensures that no value is present for AndroidVisibility, not even an explicit nil
### GetHuaweiVisibility

`func (o *Notification) GetHuaweiVisibility() int32`

GetHuaweiVisibility returns the HuaweiVisibility field if non-nil, zero value otherwise.

### GetHuaweiVisibilityOk

`func (o *Notification) GetHuaweiVisibilityOk() (*int32, bool)`

GetHuaweiVisibilityOk returns a tuple with the HuaweiVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiVisibility

`func (o *Notification) SetHuaweiVisibility(v int32)`

SetHuaweiVisibility sets HuaweiVisibility field to given value.

### HasHuaweiVisibility

`func (o *Notification) HasHuaweiVisibility() bool`

HasHuaweiVisibility returns a boolean if a field has been set.

### SetHuaweiVisibilityNil

`func (o *Notification) SetHuaweiVisibilityNil(b bool)`

 SetHuaweiVisibilityNil sets the value for HuaweiVisibility to be an explicit nil

### UnsetHuaweiVisibility
`func (o *Notification) UnsetHuaweiVisibility()`

UnsetHuaweiVisibility ensures that no value is present for HuaweiVisibility, not even an explicit nil
### GetIosBadgeType

`func (o *Notification) GetIosBadgeType() string`

GetIosBadgeType returns the IosBadgeType field if non-nil, zero value otherwise.

### GetIosBadgeTypeOk

`func (o *Notification) GetIosBadgeTypeOk() (*string, bool)`

GetIosBadgeTypeOk returns a tuple with the IosBadgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosBadgeType

`func (o *Notification) SetIosBadgeType(v string)`

SetIosBadgeType sets IosBadgeType field to given value.

### HasIosBadgeType

`func (o *Notification) HasIosBadgeType() bool`

HasIosBadgeType returns a boolean if a field has been set.

### SetIosBadgeTypeNil

`func (o *Notification) SetIosBadgeTypeNil(b bool)`

 SetIosBadgeTypeNil sets the value for IosBadgeType to be an explicit nil

### UnsetIosBadgeType
`func (o *Notification) UnsetIosBadgeType()`

UnsetIosBadgeType ensures that no value is present for IosBadgeType, not even an explicit nil
### GetIosBadgeCount

`func (o *Notification) GetIosBadgeCount() int32`

GetIosBadgeCount returns the IosBadgeCount field if non-nil, zero value otherwise.

### GetIosBadgeCountOk

`func (o *Notification) GetIosBadgeCountOk() (*int32, bool)`

GetIosBadgeCountOk returns a tuple with the IosBadgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosBadgeCount

`func (o *Notification) SetIosBadgeCount(v int32)`

SetIosBadgeCount sets IosBadgeCount field to given value.

### HasIosBadgeCount

`func (o *Notification) HasIosBadgeCount() bool`

HasIosBadgeCount returns a boolean if a field has been set.

### SetIosBadgeCountNil

`func (o *Notification) SetIosBadgeCountNil(b bool)`

 SetIosBadgeCountNil sets the value for IosBadgeCount to be an explicit nil

### UnsetIosBadgeCount
`func (o *Notification) UnsetIosBadgeCount()`

UnsetIosBadgeCount ensures that no value is present for IosBadgeCount, not even an explicit nil
### GetCollapseId

`func (o *Notification) GetCollapseId() string`

GetCollapseId returns the CollapseId field if non-nil, zero value otherwise.

### GetCollapseIdOk

`func (o *Notification) GetCollapseIdOk() (*string, bool)`

GetCollapseIdOk returns a tuple with the CollapseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseId

`func (o *Notification) SetCollapseId(v string)`

SetCollapseId sets CollapseId field to given value.

### HasCollapseId

`func (o *Notification) HasCollapseId() bool`

HasCollapseId returns a boolean if a field has been set.

### GetWebPushTopic

`func (o *Notification) GetWebPushTopic() string`

GetWebPushTopic returns the WebPushTopic field if non-nil, zero value otherwise.

### GetWebPushTopicOk

`func (o *Notification) GetWebPushTopicOk() (*string, bool)`

GetWebPushTopicOk returns a tuple with the WebPushTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPushTopic

`func (o *Notification) SetWebPushTopic(v string)`

SetWebPushTopic sets WebPushTopic field to given value.

### HasWebPushTopic

`func (o *Notification) HasWebPushTopic() bool`

HasWebPushTopic returns a boolean if a field has been set.

### SetWebPushTopicNil

`func (o *Notification) SetWebPushTopicNil(b bool)`

 SetWebPushTopicNil sets the value for WebPushTopic to be an explicit nil

### UnsetWebPushTopic
`func (o *Notification) UnsetWebPushTopic()`

UnsetWebPushTopic ensures that no value is present for WebPushTopic, not even an explicit nil
### GetApnsAlert

`func (o *Notification) GetApnsAlert() map[string]interface{}`

GetApnsAlert returns the ApnsAlert field if non-nil, zero value otherwise.

### GetApnsAlertOk

`func (o *Notification) GetApnsAlertOk() (*map[string]interface{}, bool)`

GetApnsAlertOk returns a tuple with the ApnsAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsAlert

`func (o *Notification) SetApnsAlert(v map[string]interface{})`

SetApnsAlert sets ApnsAlert field to given value.

### HasApnsAlert

`func (o *Notification) HasApnsAlert() bool`

HasApnsAlert returns a boolean if a field has been set.

### SetApnsAlertNil

`func (o *Notification) SetApnsAlertNil(b bool)`

 SetApnsAlertNil sets the value for ApnsAlert to be an explicit nil

### UnsetApnsAlert
`func (o *Notification) UnsetApnsAlert()`

UnsetApnsAlert ensures that no value is present for ApnsAlert, not even an explicit nil
### GetDelayedOption

`func (o *Notification) GetDelayedOption() string`

GetDelayedOption returns the DelayedOption field if non-nil, zero value otherwise.

### GetDelayedOptionOk

`func (o *Notification) GetDelayedOptionOk() (*string, bool)`

GetDelayedOptionOk returns a tuple with the DelayedOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedOption

`func (o *Notification) SetDelayedOption(v string)`

SetDelayedOption sets DelayedOption field to given value.

### HasDelayedOption

`func (o *Notification) HasDelayedOption() bool`

HasDelayedOption returns a boolean if a field has been set.

### SetDelayedOptionNil

`func (o *Notification) SetDelayedOptionNil(b bool)`

 SetDelayedOptionNil sets the value for DelayedOption to be an explicit nil

### UnsetDelayedOption
`func (o *Notification) UnsetDelayedOption()`

UnsetDelayedOption ensures that no value is present for DelayedOption, not even an explicit nil
### GetDeliveryTimeOfDay

`func (o *Notification) GetDeliveryTimeOfDay() string`

GetDeliveryTimeOfDay returns the DeliveryTimeOfDay field if non-nil, zero value otherwise.

### GetDeliveryTimeOfDayOk

`func (o *Notification) GetDeliveryTimeOfDayOk() (*string, bool)`

GetDeliveryTimeOfDayOk returns a tuple with the DeliveryTimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeOfDay

`func (o *Notification) SetDeliveryTimeOfDay(v string)`

SetDeliveryTimeOfDay sets DeliveryTimeOfDay field to given value.

### HasDeliveryTimeOfDay

`func (o *Notification) HasDeliveryTimeOfDay() bool`

HasDeliveryTimeOfDay returns a boolean if a field has been set.

### SetDeliveryTimeOfDayNil

`func (o *Notification) SetDeliveryTimeOfDayNil(b bool)`

 SetDeliveryTimeOfDayNil sets the value for DeliveryTimeOfDay to be an explicit nil

### UnsetDeliveryTimeOfDay
`func (o *Notification) UnsetDeliveryTimeOfDay()`

UnsetDeliveryTimeOfDay ensures that no value is present for DeliveryTimeOfDay, not even an explicit nil
### GetTtl

`func (o *Notification) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Notification) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Notification) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Notification) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *Notification) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *Notification) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetPriority

`func (o *Notification) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Notification) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Notification) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Notification) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *Notification) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *Notification) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetApnsPushTypeOverride

`func (o *Notification) GetApnsPushTypeOverride() string`

GetApnsPushTypeOverride returns the ApnsPushTypeOverride field if non-nil, zero value otherwise.

### GetApnsPushTypeOverrideOk

`func (o *Notification) GetApnsPushTypeOverrideOk() (*string, bool)`

GetApnsPushTypeOverrideOk returns a tuple with the ApnsPushTypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsPushTypeOverride

`func (o *Notification) SetApnsPushTypeOverride(v string)`

SetApnsPushTypeOverride sets ApnsPushTypeOverride field to given value.

### HasApnsPushTypeOverride

`func (o *Notification) HasApnsPushTypeOverride() bool`

HasApnsPushTypeOverride returns a boolean if a field has been set.

### GetThrottleRatePerMinute

`func (o *Notification) GetThrottleRatePerMinute() string`

GetThrottleRatePerMinute returns the ThrottleRatePerMinute field if non-nil, zero value otherwise.

### GetThrottleRatePerMinuteOk

`func (o *Notification) GetThrottleRatePerMinuteOk() (*string, bool)`

GetThrottleRatePerMinuteOk returns a tuple with the ThrottleRatePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRatePerMinute

`func (o *Notification) SetThrottleRatePerMinute(v string)`

SetThrottleRatePerMinute sets ThrottleRatePerMinute field to given value.

### HasThrottleRatePerMinute

`func (o *Notification) HasThrottleRatePerMinute() bool`

HasThrottleRatePerMinute returns a boolean if a field has been set.

### SetThrottleRatePerMinuteNil

`func (o *Notification) SetThrottleRatePerMinuteNil(b bool)`

 SetThrottleRatePerMinuteNil sets the value for ThrottleRatePerMinute to be an explicit nil

### UnsetThrottleRatePerMinute
`func (o *Notification) UnsetThrottleRatePerMinute()`

UnsetThrottleRatePerMinute ensures that no value is present for ThrottleRatePerMinute, not even an explicit nil
### GetAndroidGroup

`func (o *Notification) GetAndroidGroup() string`

GetAndroidGroup returns the AndroidGroup field if non-nil, zero value otherwise.

### GetAndroidGroupOk

`func (o *Notification) GetAndroidGroupOk() (*string, bool)`

GetAndroidGroupOk returns a tuple with the AndroidGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidGroup

`func (o *Notification) SetAndroidGroup(v string)`

SetAndroidGroup sets AndroidGroup field to given value.

### HasAndroidGroup

`func (o *Notification) HasAndroidGroup() bool`

HasAndroidGroup returns a boolean if a field has been set.

### SetAndroidGroupNil

`func (o *Notification) SetAndroidGroupNil(b bool)`

 SetAndroidGroupNil sets the value for AndroidGroup to be an explicit nil

### UnsetAndroidGroup
`func (o *Notification) UnsetAndroidGroup()`

UnsetAndroidGroup ensures that no value is present for AndroidGroup, not even an explicit nil
### GetAndroidGroupMessage

`func (o *Notification) GetAndroidGroupMessage() string`

GetAndroidGroupMessage returns the AndroidGroupMessage field if non-nil, zero value otherwise.

### GetAndroidGroupMessageOk

`func (o *Notification) GetAndroidGroupMessageOk() (*string, bool)`

GetAndroidGroupMessageOk returns a tuple with the AndroidGroupMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidGroupMessage

`func (o *Notification) SetAndroidGroupMessage(v string)`

SetAndroidGroupMessage sets AndroidGroupMessage field to given value.

### HasAndroidGroupMessage

`func (o *Notification) HasAndroidGroupMessage() bool`

HasAndroidGroupMessage returns a boolean if a field has been set.

### SetAndroidGroupMessageNil

`func (o *Notification) SetAndroidGroupMessageNil(b bool)`

 SetAndroidGroupMessageNil sets the value for AndroidGroupMessage to be an explicit nil

### UnsetAndroidGroupMessage
`func (o *Notification) UnsetAndroidGroupMessage()`

UnsetAndroidGroupMessage ensures that no value is present for AndroidGroupMessage, not even an explicit nil
### GetAdmGroup

`func (o *Notification) GetAdmGroup() string`

GetAdmGroup returns the AdmGroup field if non-nil, zero value otherwise.

### GetAdmGroupOk

`func (o *Notification) GetAdmGroupOk() (*string, bool)`

GetAdmGroupOk returns a tuple with the AdmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmGroup

`func (o *Notification) SetAdmGroup(v string)`

SetAdmGroup sets AdmGroup field to given value.

### HasAdmGroup

`func (o *Notification) HasAdmGroup() bool`

HasAdmGroup returns a boolean if a field has been set.

### SetAdmGroupNil

`func (o *Notification) SetAdmGroupNil(b bool)`

 SetAdmGroupNil sets the value for AdmGroup to be an explicit nil

### UnsetAdmGroup
`func (o *Notification) UnsetAdmGroup()`

UnsetAdmGroup ensures that no value is present for AdmGroup, not even an explicit nil
### GetAdmGroupMessage

`func (o *Notification) GetAdmGroupMessage() map[string]interface{}`

GetAdmGroupMessage returns the AdmGroupMessage field if non-nil, zero value otherwise.

### GetAdmGroupMessageOk

`func (o *Notification) GetAdmGroupMessageOk() (*map[string]interface{}, bool)`

GetAdmGroupMessageOk returns a tuple with the AdmGroupMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmGroupMessage

`func (o *Notification) SetAdmGroupMessage(v map[string]interface{})`

SetAdmGroupMessage sets AdmGroupMessage field to given value.

### HasAdmGroupMessage

`func (o *Notification) HasAdmGroupMessage() bool`

HasAdmGroupMessage returns a boolean if a field has been set.

### SetAdmGroupMessageNil

`func (o *Notification) SetAdmGroupMessageNil(b bool)`

 SetAdmGroupMessageNil sets the value for AdmGroupMessage to be an explicit nil

### UnsetAdmGroupMessage
`func (o *Notification) UnsetAdmGroupMessage()`

UnsetAdmGroupMessage ensures that no value is present for AdmGroupMessage, not even an explicit nil
### GetThreadId

`func (o *Notification) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *Notification) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *Notification) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *Notification) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### SetThreadIdNil

`func (o *Notification) SetThreadIdNil(b bool)`

 SetThreadIdNil sets the value for ThreadId to be an explicit nil

### UnsetThreadId
`func (o *Notification) UnsetThreadId()`

UnsetThreadId ensures that no value is present for ThreadId, not even an explicit nil
### GetSummaryArg

`func (o *Notification) GetSummaryArg() string`

GetSummaryArg returns the SummaryArg field if non-nil, zero value otherwise.

### GetSummaryArgOk

`func (o *Notification) GetSummaryArgOk() (*string, bool)`

GetSummaryArgOk returns a tuple with the SummaryArg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryArg

`func (o *Notification) SetSummaryArg(v string)`

SetSummaryArg sets SummaryArg field to given value.

### HasSummaryArg

`func (o *Notification) HasSummaryArg() bool`

HasSummaryArg returns a boolean if a field has been set.

### GetSummaryArgCount

`func (o *Notification) GetSummaryArgCount() int32`

GetSummaryArgCount returns the SummaryArgCount field if non-nil, zero value otherwise.

### GetSummaryArgCountOk

`func (o *Notification) GetSummaryArgCountOk() (*int32, bool)`

GetSummaryArgCountOk returns a tuple with the SummaryArgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryArgCount

`func (o *Notification) SetSummaryArgCount(v int32)`

SetSummaryArgCount sets SummaryArgCount field to given value.

### HasSummaryArgCount

`func (o *Notification) HasSummaryArgCount() bool`

HasSummaryArgCount returns a boolean if a field has been set.

### GetEmailSubject

`func (o *Notification) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *Notification) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *Notification) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *Notification) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### SetEmailSubjectNil

`func (o *Notification) SetEmailSubjectNil(b bool)`

 SetEmailSubjectNil sets the value for EmailSubject to be an explicit nil

### UnsetEmailSubject
`func (o *Notification) UnsetEmailSubject()`

UnsetEmailSubject ensures that no value is present for EmailSubject, not even an explicit nil
### GetEmailBody

`func (o *Notification) GetEmailBody() string`

GetEmailBody returns the EmailBody field if non-nil, zero value otherwise.

### GetEmailBodyOk

`func (o *Notification) GetEmailBodyOk() (*string, bool)`

GetEmailBodyOk returns a tuple with the EmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailBody

`func (o *Notification) SetEmailBody(v string)`

SetEmailBody sets EmailBody field to given value.

### HasEmailBody

`func (o *Notification) HasEmailBody() bool`

HasEmailBody returns a boolean if a field has been set.

### GetEmailFromName

`func (o *Notification) GetEmailFromName() string`

GetEmailFromName returns the EmailFromName field if non-nil, zero value otherwise.

### GetEmailFromNameOk

`func (o *Notification) GetEmailFromNameOk() (*string, bool)`

GetEmailFromNameOk returns a tuple with the EmailFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromName

`func (o *Notification) SetEmailFromName(v string)`

SetEmailFromName sets EmailFromName field to given value.

### HasEmailFromName

`func (o *Notification) HasEmailFromName() bool`

HasEmailFromName returns a boolean if a field has been set.

### SetEmailFromNameNil

`func (o *Notification) SetEmailFromNameNil(b bool)`

 SetEmailFromNameNil sets the value for EmailFromName to be an explicit nil

### UnsetEmailFromName
`func (o *Notification) UnsetEmailFromName()`

UnsetEmailFromName ensures that no value is present for EmailFromName, not even an explicit nil
### GetEmailFromAddress

`func (o *Notification) GetEmailFromAddress() string`

GetEmailFromAddress returns the EmailFromAddress field if non-nil, zero value otherwise.

### GetEmailFromAddressOk

`func (o *Notification) GetEmailFromAddressOk() (*string, bool)`

GetEmailFromAddressOk returns a tuple with the EmailFromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromAddress

`func (o *Notification) SetEmailFromAddress(v string)`

SetEmailFromAddress sets EmailFromAddress field to given value.

### HasEmailFromAddress

`func (o *Notification) HasEmailFromAddress() bool`

HasEmailFromAddress returns a boolean if a field has been set.

### SetEmailFromAddressNil

`func (o *Notification) SetEmailFromAddressNil(b bool)`

 SetEmailFromAddressNil sets the value for EmailFromAddress to be an explicit nil

### UnsetEmailFromAddress
`func (o *Notification) UnsetEmailFromAddress()`

UnsetEmailFromAddress ensures that no value is present for EmailFromAddress, not even an explicit nil
### GetEmailPreheader

`func (o *Notification) GetEmailPreheader() string`

GetEmailPreheader returns the EmailPreheader field if non-nil, zero value otherwise.

### GetEmailPreheaderOk

`func (o *Notification) GetEmailPreheaderOk() (*string, bool)`

GetEmailPreheaderOk returns a tuple with the EmailPreheader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPreheader

`func (o *Notification) SetEmailPreheader(v string)`

SetEmailPreheader sets EmailPreheader field to given value.

### HasEmailPreheader

`func (o *Notification) HasEmailPreheader() bool`

HasEmailPreheader returns a boolean if a field has been set.

### SetEmailPreheaderNil

`func (o *Notification) SetEmailPreheaderNil(b bool)`

 SetEmailPreheaderNil sets the value for EmailPreheader to be an explicit nil

### UnsetEmailPreheader
`func (o *Notification) UnsetEmailPreheader()`

UnsetEmailPreheader ensures that no value is present for EmailPreheader, not even an explicit nil
### GetIncludeUnsubscribed

`func (o *Notification) GetIncludeUnsubscribed() bool`

GetIncludeUnsubscribed returns the IncludeUnsubscribed field if non-nil, zero value otherwise.

### GetIncludeUnsubscribedOk

`func (o *Notification) GetIncludeUnsubscribedOk() (*bool, bool)`

GetIncludeUnsubscribedOk returns a tuple with the IncludeUnsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUnsubscribed

`func (o *Notification) SetIncludeUnsubscribed(v bool)`

SetIncludeUnsubscribed sets IncludeUnsubscribed field to given value.

### HasIncludeUnsubscribed

`func (o *Notification) HasIncludeUnsubscribed() bool`

HasIncludeUnsubscribed returns a boolean if a field has been set.

### GetSmsFrom

`func (o *Notification) GetSmsFrom() string`

GetSmsFrom returns the SmsFrom field if non-nil, zero value otherwise.

### GetSmsFromOk

`func (o *Notification) GetSmsFromOk() (*string, bool)`

GetSmsFromOk returns a tuple with the SmsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFrom

`func (o *Notification) SetSmsFrom(v string)`

SetSmsFrom sets SmsFrom field to given value.

### HasSmsFrom

`func (o *Notification) HasSmsFrom() bool`

HasSmsFrom returns a boolean if a field has been set.

### SetSmsFromNil

`func (o *Notification) SetSmsFromNil(b bool)`

 SetSmsFromNil sets the value for SmsFrom to be an explicit nil

### UnsetSmsFrom
`func (o *Notification) UnsetSmsFrom()`

UnsetSmsFrom ensures that no value is present for SmsFrom, not even an explicit nil
### GetSmsMediaUrls

`func (o *Notification) GetSmsMediaUrls() []string`

GetSmsMediaUrls returns the SmsMediaUrls field if non-nil, zero value otherwise.

### GetSmsMediaUrlsOk

`func (o *Notification) GetSmsMediaUrlsOk() (*[]string, bool)`

GetSmsMediaUrlsOk returns a tuple with the SmsMediaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMediaUrls

`func (o *Notification) SetSmsMediaUrls(v []string)`

SetSmsMediaUrls sets SmsMediaUrls field to given value.

### HasSmsMediaUrls

`func (o *Notification) HasSmsMediaUrls() bool`

HasSmsMediaUrls returns a boolean if a field has been set.

### SetSmsMediaUrlsNil

`func (o *Notification) SetSmsMediaUrlsNil(b bool)`

 SetSmsMediaUrlsNil sets the value for SmsMediaUrls to be an explicit nil

### UnsetSmsMediaUrls
`func (o *Notification) UnsetSmsMediaUrls()`

UnsetSmsMediaUrls ensures that no value is present for SmsMediaUrls, not even an explicit nil
### GetFilters

`func (o *Notification) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Notification) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Notification) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Notification) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *Notification) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *Notification) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetCustomData

`func (o *Notification) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *Notification) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *Notification) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *Notification) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### SetCustomDataNil

`func (o *Notification) SetCustomDataNil(b bool)`

 SetCustomDataNil sets the value for CustomData to be an explicit nil

### UnsetCustomData
`func (o *Notification) UnsetCustomData()`

UnsetCustomData ensures that no value is present for CustomData, not even an explicit nil
### GetSendAfter

`func (o *Notification) GetSendAfter() time.Time`

GetSendAfter returns the SendAfter field if non-nil, zero value otherwise.

### GetSendAfterOk

`func (o *Notification) GetSendAfterOk() (*time.Time, bool)`

GetSendAfterOk returns a tuple with the SendAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAfter

`func (o *Notification) SetSendAfter(v time.Time)`

SetSendAfter sets SendAfter field to given value.

### HasSendAfter

`func (o *Notification) HasSendAfter() bool`

HasSendAfter returns a boolean if a field has been set.

### SetSendAfterNil

`func (o *Notification) SetSendAfterNil(b bool)`

 SetSendAfterNil sets the value for SendAfter to be an explicit nil

### UnsetSendAfter
`func (o *Notification) UnsetSendAfter()`

UnsetSendAfter ensures that no value is present for SendAfter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


