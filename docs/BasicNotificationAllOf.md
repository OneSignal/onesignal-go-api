# BasicNotificationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**AppId** | Pointer to **string** | Required: Your OneSignal Application ID, which can be found in Keys &amp; IDs. It is a UUID and looks similar to 8250eaf6-1a58-489e-b136-7c74a864b434.  | [optional] 
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

## Methods

### NewBasicNotificationAllOf

`func NewBasicNotificationAllOf() *BasicNotificationAllOf`

NewBasicNotificationAllOf instantiates a new BasicNotificationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicNotificationAllOfWithDefaults

`func NewBasicNotificationAllOfWithDefaults() *BasicNotificationAllOf`

NewBasicNotificationAllOfWithDefaults instantiates a new BasicNotificationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasicNotificationAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicNotificationAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicNotificationAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BasicNotificationAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *BasicNotificationAllOf) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BasicNotificationAllOf) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BasicNotificationAllOf) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *BasicNotificationAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetName

`func (o *BasicNotificationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicNotificationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicNotificationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasicNotificationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BasicNotificationAllOf) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BasicNotificationAllOf) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAggregation

`func (o *BasicNotificationAllOf) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *BasicNotificationAllOf) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *BasicNotificationAllOf) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *BasicNotificationAllOf) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetIsIos

`func (o *BasicNotificationAllOf) GetIsIos() bool`

GetIsIos returns the IsIos field if non-nil, zero value otherwise.

### GetIsIosOk

`func (o *BasicNotificationAllOf) GetIsIosOk() (*bool, bool)`

GetIsIosOk returns a tuple with the IsIos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIos

`func (o *BasicNotificationAllOf) SetIsIos(v bool)`

SetIsIos sets IsIos field to given value.

### HasIsIos

`func (o *BasicNotificationAllOf) HasIsIos() bool`

HasIsIos returns a boolean if a field has been set.

### SetIsIosNil

`func (o *BasicNotificationAllOf) SetIsIosNil(b bool)`

 SetIsIosNil sets the value for IsIos to be an explicit nil

### UnsetIsIos
`func (o *BasicNotificationAllOf) UnsetIsIos()`

UnsetIsIos ensures that no value is present for IsIos, not even an explicit nil
### GetIsAndroid

`func (o *BasicNotificationAllOf) GetIsAndroid() bool`

GetIsAndroid returns the IsAndroid field if non-nil, zero value otherwise.

### GetIsAndroidOk

`func (o *BasicNotificationAllOf) GetIsAndroidOk() (*bool, bool)`

GetIsAndroidOk returns a tuple with the IsAndroid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAndroid

`func (o *BasicNotificationAllOf) SetIsAndroid(v bool)`

SetIsAndroid sets IsAndroid field to given value.

### HasIsAndroid

`func (o *BasicNotificationAllOf) HasIsAndroid() bool`

HasIsAndroid returns a boolean if a field has been set.

### SetIsAndroidNil

`func (o *BasicNotificationAllOf) SetIsAndroidNil(b bool)`

 SetIsAndroidNil sets the value for IsAndroid to be an explicit nil

### UnsetIsAndroid
`func (o *BasicNotificationAllOf) UnsetIsAndroid()`

UnsetIsAndroid ensures that no value is present for IsAndroid, not even an explicit nil
### GetIsHuawei

`func (o *BasicNotificationAllOf) GetIsHuawei() bool`

GetIsHuawei returns the IsHuawei field if non-nil, zero value otherwise.

### GetIsHuaweiOk

`func (o *BasicNotificationAllOf) GetIsHuaweiOk() (*bool, bool)`

GetIsHuaweiOk returns a tuple with the IsHuawei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHuawei

`func (o *BasicNotificationAllOf) SetIsHuawei(v bool)`

SetIsHuawei sets IsHuawei field to given value.

### HasIsHuawei

`func (o *BasicNotificationAllOf) HasIsHuawei() bool`

HasIsHuawei returns a boolean if a field has been set.

### SetIsHuaweiNil

`func (o *BasicNotificationAllOf) SetIsHuaweiNil(b bool)`

 SetIsHuaweiNil sets the value for IsHuawei to be an explicit nil

### UnsetIsHuawei
`func (o *BasicNotificationAllOf) UnsetIsHuawei()`

UnsetIsHuawei ensures that no value is present for IsHuawei, not even an explicit nil
### GetIsAnyWeb

`func (o *BasicNotificationAllOf) GetIsAnyWeb() bool`

GetIsAnyWeb returns the IsAnyWeb field if non-nil, zero value otherwise.

### GetIsAnyWebOk

`func (o *BasicNotificationAllOf) GetIsAnyWebOk() (*bool, bool)`

GetIsAnyWebOk returns a tuple with the IsAnyWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnyWeb

`func (o *BasicNotificationAllOf) SetIsAnyWeb(v bool)`

SetIsAnyWeb sets IsAnyWeb field to given value.

### HasIsAnyWeb

`func (o *BasicNotificationAllOf) HasIsAnyWeb() bool`

HasIsAnyWeb returns a boolean if a field has been set.

### SetIsAnyWebNil

`func (o *BasicNotificationAllOf) SetIsAnyWebNil(b bool)`

 SetIsAnyWebNil sets the value for IsAnyWeb to be an explicit nil

### UnsetIsAnyWeb
`func (o *BasicNotificationAllOf) UnsetIsAnyWeb()`

UnsetIsAnyWeb ensures that no value is present for IsAnyWeb, not even an explicit nil
### GetIsChromeWeb

`func (o *BasicNotificationAllOf) GetIsChromeWeb() bool`

GetIsChromeWeb returns the IsChromeWeb field if non-nil, zero value otherwise.

### GetIsChromeWebOk

`func (o *BasicNotificationAllOf) GetIsChromeWebOk() (*bool, bool)`

GetIsChromeWebOk returns a tuple with the IsChromeWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChromeWeb

`func (o *BasicNotificationAllOf) SetIsChromeWeb(v bool)`

SetIsChromeWeb sets IsChromeWeb field to given value.

### HasIsChromeWeb

`func (o *BasicNotificationAllOf) HasIsChromeWeb() bool`

HasIsChromeWeb returns a boolean if a field has been set.

### SetIsChromeWebNil

`func (o *BasicNotificationAllOf) SetIsChromeWebNil(b bool)`

 SetIsChromeWebNil sets the value for IsChromeWeb to be an explicit nil

### UnsetIsChromeWeb
`func (o *BasicNotificationAllOf) UnsetIsChromeWeb()`

UnsetIsChromeWeb ensures that no value is present for IsChromeWeb, not even an explicit nil
### GetIsFirefox

`func (o *BasicNotificationAllOf) GetIsFirefox() bool`

GetIsFirefox returns the IsFirefox field if non-nil, zero value otherwise.

### GetIsFirefoxOk

`func (o *BasicNotificationAllOf) GetIsFirefoxOk() (*bool, bool)`

GetIsFirefoxOk returns a tuple with the IsFirefox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirefox

`func (o *BasicNotificationAllOf) SetIsFirefox(v bool)`

SetIsFirefox sets IsFirefox field to given value.

### HasIsFirefox

`func (o *BasicNotificationAllOf) HasIsFirefox() bool`

HasIsFirefox returns a boolean if a field has been set.

### SetIsFirefoxNil

`func (o *BasicNotificationAllOf) SetIsFirefoxNil(b bool)`

 SetIsFirefoxNil sets the value for IsFirefox to be an explicit nil

### UnsetIsFirefox
`func (o *BasicNotificationAllOf) UnsetIsFirefox()`

UnsetIsFirefox ensures that no value is present for IsFirefox, not even an explicit nil
### GetIsSafari

`func (o *BasicNotificationAllOf) GetIsSafari() bool`

GetIsSafari returns the IsSafari field if non-nil, zero value otherwise.

### GetIsSafariOk

`func (o *BasicNotificationAllOf) GetIsSafariOk() (*bool, bool)`

GetIsSafariOk returns a tuple with the IsSafari field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafari

`func (o *BasicNotificationAllOf) SetIsSafari(v bool)`

SetIsSafari sets IsSafari field to given value.

### HasIsSafari

`func (o *BasicNotificationAllOf) HasIsSafari() bool`

HasIsSafari returns a boolean if a field has been set.

### SetIsSafariNil

`func (o *BasicNotificationAllOf) SetIsSafariNil(b bool)`

 SetIsSafariNil sets the value for IsSafari to be an explicit nil

### UnsetIsSafari
`func (o *BasicNotificationAllOf) UnsetIsSafari()`

UnsetIsSafari ensures that no value is present for IsSafari, not even an explicit nil
### GetIsWPWNS

`func (o *BasicNotificationAllOf) GetIsWPWNS() bool`

GetIsWPWNS returns the IsWPWNS field if non-nil, zero value otherwise.

### GetIsWPWNSOk

`func (o *BasicNotificationAllOf) GetIsWPWNSOk() (*bool, bool)`

GetIsWPWNSOk returns a tuple with the IsWPWNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWPWNS

`func (o *BasicNotificationAllOf) SetIsWPWNS(v bool)`

SetIsWPWNS sets IsWPWNS field to given value.

### HasIsWPWNS

`func (o *BasicNotificationAllOf) HasIsWPWNS() bool`

HasIsWPWNS returns a boolean if a field has been set.

### SetIsWPWNSNil

`func (o *BasicNotificationAllOf) SetIsWPWNSNil(b bool)`

 SetIsWPWNSNil sets the value for IsWPWNS to be an explicit nil

### UnsetIsWPWNS
`func (o *BasicNotificationAllOf) UnsetIsWPWNS()`

UnsetIsWPWNS ensures that no value is present for IsWPWNS, not even an explicit nil
### GetIsAdm

`func (o *BasicNotificationAllOf) GetIsAdm() bool`

GetIsAdm returns the IsAdm field if non-nil, zero value otherwise.

### GetIsAdmOk

`func (o *BasicNotificationAllOf) GetIsAdmOk() (*bool, bool)`

GetIsAdmOk returns a tuple with the IsAdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdm

`func (o *BasicNotificationAllOf) SetIsAdm(v bool)`

SetIsAdm sets IsAdm field to given value.

### HasIsAdm

`func (o *BasicNotificationAllOf) HasIsAdm() bool`

HasIsAdm returns a boolean if a field has been set.

### SetIsAdmNil

`func (o *BasicNotificationAllOf) SetIsAdmNil(b bool)`

 SetIsAdmNil sets the value for IsAdm to be an explicit nil

### UnsetIsAdm
`func (o *BasicNotificationAllOf) UnsetIsAdm()`

UnsetIsAdm ensures that no value is present for IsAdm, not even an explicit nil
### GetIsChrome

`func (o *BasicNotificationAllOf) GetIsChrome() bool`

GetIsChrome returns the IsChrome field if non-nil, zero value otherwise.

### GetIsChromeOk

`func (o *BasicNotificationAllOf) GetIsChromeOk() (*bool, bool)`

GetIsChromeOk returns a tuple with the IsChrome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChrome

`func (o *BasicNotificationAllOf) SetIsChrome(v bool)`

SetIsChrome sets IsChrome field to given value.

### HasIsChrome

`func (o *BasicNotificationAllOf) HasIsChrome() bool`

HasIsChrome returns a boolean if a field has been set.

### SetIsChromeNil

`func (o *BasicNotificationAllOf) SetIsChromeNil(b bool)`

 SetIsChromeNil sets the value for IsChrome to be an explicit nil

### UnsetIsChrome
`func (o *BasicNotificationAllOf) UnsetIsChrome()`

UnsetIsChrome ensures that no value is present for IsChrome, not even an explicit nil
### GetChannelForExternalUserIds

`func (o *BasicNotificationAllOf) GetChannelForExternalUserIds() string`

GetChannelForExternalUserIds returns the ChannelForExternalUserIds field if non-nil, zero value otherwise.

### GetChannelForExternalUserIdsOk

`func (o *BasicNotificationAllOf) GetChannelForExternalUserIdsOk() (*string, bool)`

GetChannelForExternalUserIdsOk returns a tuple with the ChannelForExternalUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelForExternalUserIds

`func (o *BasicNotificationAllOf) SetChannelForExternalUserIds(v string)`

SetChannelForExternalUserIds sets ChannelForExternalUserIds field to given value.

### HasChannelForExternalUserIds

`func (o *BasicNotificationAllOf) HasChannelForExternalUserIds() bool`

HasChannelForExternalUserIds returns a boolean if a field has been set.

### GetAppId

`func (o *BasicNotificationAllOf) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *BasicNotificationAllOf) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *BasicNotificationAllOf) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *BasicNotificationAllOf) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetExternalId

`func (o *BasicNotificationAllOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BasicNotificationAllOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BasicNotificationAllOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BasicNotificationAllOf) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *BasicNotificationAllOf) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *BasicNotificationAllOf) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetIdempotencyKey

`func (o *BasicNotificationAllOf) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *BasicNotificationAllOf) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *BasicNotificationAllOf) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *BasicNotificationAllOf) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### SetIdempotencyKeyNil

`func (o *BasicNotificationAllOf) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *BasicNotificationAllOf) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
### GetContents

`func (o *BasicNotificationAllOf) GetContents() StringMap`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *BasicNotificationAllOf) GetContentsOk() (*StringMap, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *BasicNotificationAllOf) SetContents(v StringMap)`

SetContents sets Contents field to given value.

### HasContents

`func (o *BasicNotificationAllOf) HasContents() bool`

HasContents returns a boolean if a field has been set.

### SetContentsNil

`func (o *BasicNotificationAllOf) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *BasicNotificationAllOf) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetHeadings

`func (o *BasicNotificationAllOf) GetHeadings() StringMap`

GetHeadings returns the Headings field if non-nil, zero value otherwise.

### GetHeadingsOk

`func (o *BasicNotificationAllOf) GetHeadingsOk() (*StringMap, bool)`

GetHeadingsOk returns a tuple with the Headings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadings

`func (o *BasicNotificationAllOf) SetHeadings(v StringMap)`

SetHeadings sets Headings field to given value.

### HasHeadings

`func (o *BasicNotificationAllOf) HasHeadings() bool`

HasHeadings returns a boolean if a field has been set.

### SetHeadingsNil

`func (o *BasicNotificationAllOf) SetHeadingsNil(b bool)`

 SetHeadingsNil sets the value for Headings to be an explicit nil

### UnsetHeadings
`func (o *BasicNotificationAllOf) UnsetHeadings()`

UnsetHeadings ensures that no value is present for Headings, not even an explicit nil
### GetSubtitle

`func (o *BasicNotificationAllOf) GetSubtitle() StringMap`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *BasicNotificationAllOf) GetSubtitleOk() (*StringMap, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *BasicNotificationAllOf) SetSubtitle(v StringMap)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *BasicNotificationAllOf) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### SetSubtitleNil

`func (o *BasicNotificationAllOf) SetSubtitleNil(b bool)`

 SetSubtitleNil sets the value for Subtitle to be an explicit nil

### UnsetSubtitle
`func (o *BasicNotificationAllOf) UnsetSubtitle()`

UnsetSubtitle ensures that no value is present for Subtitle, not even an explicit nil
### GetData

`func (o *BasicNotificationAllOf) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BasicNotificationAllOf) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BasicNotificationAllOf) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *BasicNotificationAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BasicNotificationAllOf) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BasicNotificationAllOf) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetHuaweiMsgType

`func (o *BasicNotificationAllOf) GetHuaweiMsgType() string`

GetHuaweiMsgType returns the HuaweiMsgType field if non-nil, zero value otherwise.

### GetHuaweiMsgTypeOk

`func (o *BasicNotificationAllOf) GetHuaweiMsgTypeOk() (*string, bool)`

GetHuaweiMsgTypeOk returns a tuple with the HuaweiMsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiMsgType

`func (o *BasicNotificationAllOf) SetHuaweiMsgType(v string)`

SetHuaweiMsgType sets HuaweiMsgType field to given value.

### HasHuaweiMsgType

`func (o *BasicNotificationAllOf) HasHuaweiMsgType() bool`

HasHuaweiMsgType returns a boolean if a field has been set.

### SetHuaweiMsgTypeNil

`func (o *BasicNotificationAllOf) SetHuaweiMsgTypeNil(b bool)`

 SetHuaweiMsgTypeNil sets the value for HuaweiMsgType to be an explicit nil

### UnsetHuaweiMsgType
`func (o *BasicNotificationAllOf) UnsetHuaweiMsgType()`

UnsetHuaweiMsgType ensures that no value is present for HuaweiMsgType, not even an explicit nil
### GetUrl

`func (o *BasicNotificationAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BasicNotificationAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BasicNotificationAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BasicNotificationAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *BasicNotificationAllOf) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *BasicNotificationAllOf) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWebUrl

`func (o *BasicNotificationAllOf) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *BasicNotificationAllOf) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *BasicNotificationAllOf) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *BasicNotificationAllOf) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *BasicNotificationAllOf) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *BasicNotificationAllOf) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetAppUrl

`func (o *BasicNotificationAllOf) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *BasicNotificationAllOf) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *BasicNotificationAllOf) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.

### HasAppUrl

`func (o *BasicNotificationAllOf) HasAppUrl() bool`

HasAppUrl returns a boolean if a field has been set.

### SetAppUrlNil

`func (o *BasicNotificationAllOf) SetAppUrlNil(b bool)`

 SetAppUrlNil sets the value for AppUrl to be an explicit nil

### UnsetAppUrl
`func (o *BasicNotificationAllOf) UnsetAppUrl()`

UnsetAppUrl ensures that no value is present for AppUrl, not even an explicit nil
### GetIosAttachments

`func (o *BasicNotificationAllOf) GetIosAttachments() map[string]interface{}`

GetIosAttachments returns the IosAttachments field if non-nil, zero value otherwise.

### GetIosAttachmentsOk

`func (o *BasicNotificationAllOf) GetIosAttachmentsOk() (*map[string]interface{}, bool)`

GetIosAttachmentsOk returns a tuple with the IosAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosAttachments

`func (o *BasicNotificationAllOf) SetIosAttachments(v map[string]interface{})`

SetIosAttachments sets IosAttachments field to given value.

### HasIosAttachments

`func (o *BasicNotificationAllOf) HasIosAttachments() bool`

HasIosAttachments returns a boolean if a field has been set.

### SetIosAttachmentsNil

`func (o *BasicNotificationAllOf) SetIosAttachmentsNil(b bool)`

 SetIosAttachmentsNil sets the value for IosAttachments to be an explicit nil

### UnsetIosAttachments
`func (o *BasicNotificationAllOf) UnsetIosAttachments()`

UnsetIosAttachments ensures that no value is present for IosAttachments, not even an explicit nil
### GetTemplateId

`func (o *BasicNotificationAllOf) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *BasicNotificationAllOf) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *BasicNotificationAllOf) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *BasicNotificationAllOf) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *BasicNotificationAllOf) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *BasicNotificationAllOf) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetContentAvailable

`func (o *BasicNotificationAllOf) GetContentAvailable() bool`

GetContentAvailable returns the ContentAvailable field if non-nil, zero value otherwise.

### GetContentAvailableOk

`func (o *BasicNotificationAllOf) GetContentAvailableOk() (*bool, bool)`

GetContentAvailableOk returns a tuple with the ContentAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAvailable

`func (o *BasicNotificationAllOf) SetContentAvailable(v bool)`

SetContentAvailable sets ContentAvailable field to given value.

### HasContentAvailable

`func (o *BasicNotificationAllOf) HasContentAvailable() bool`

HasContentAvailable returns a boolean if a field has been set.

### SetContentAvailableNil

`func (o *BasicNotificationAllOf) SetContentAvailableNil(b bool)`

 SetContentAvailableNil sets the value for ContentAvailable to be an explicit nil

### UnsetContentAvailable
`func (o *BasicNotificationAllOf) UnsetContentAvailable()`

UnsetContentAvailable ensures that no value is present for ContentAvailable, not even an explicit nil
### GetMutableContent

`func (o *BasicNotificationAllOf) GetMutableContent() bool`

GetMutableContent returns the MutableContent field if non-nil, zero value otherwise.

### GetMutableContentOk

`func (o *BasicNotificationAllOf) GetMutableContentOk() (*bool, bool)`

GetMutableContentOk returns a tuple with the MutableContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutableContent

`func (o *BasicNotificationAllOf) SetMutableContent(v bool)`

SetMutableContent sets MutableContent field to given value.

### HasMutableContent

`func (o *BasicNotificationAllOf) HasMutableContent() bool`

HasMutableContent returns a boolean if a field has been set.

### GetTargetContentIdentifier

`func (o *BasicNotificationAllOf) GetTargetContentIdentifier() string`

GetTargetContentIdentifier returns the TargetContentIdentifier field if non-nil, zero value otherwise.

### GetTargetContentIdentifierOk

`func (o *BasicNotificationAllOf) GetTargetContentIdentifierOk() (*string, bool)`

GetTargetContentIdentifierOk returns a tuple with the TargetContentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContentIdentifier

`func (o *BasicNotificationAllOf) SetTargetContentIdentifier(v string)`

SetTargetContentIdentifier sets TargetContentIdentifier field to given value.

### HasTargetContentIdentifier

`func (o *BasicNotificationAllOf) HasTargetContentIdentifier() bool`

HasTargetContentIdentifier returns a boolean if a field has been set.

### SetTargetContentIdentifierNil

`func (o *BasicNotificationAllOf) SetTargetContentIdentifierNil(b bool)`

 SetTargetContentIdentifierNil sets the value for TargetContentIdentifier to be an explicit nil

### UnsetTargetContentIdentifier
`func (o *BasicNotificationAllOf) UnsetTargetContentIdentifier()`

UnsetTargetContentIdentifier ensures that no value is present for TargetContentIdentifier, not even an explicit nil
### GetBigPicture

`func (o *BasicNotificationAllOf) GetBigPicture() string`

GetBigPicture returns the BigPicture field if non-nil, zero value otherwise.

### GetBigPictureOk

`func (o *BasicNotificationAllOf) GetBigPictureOk() (*string, bool)`

GetBigPictureOk returns a tuple with the BigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigPicture

`func (o *BasicNotificationAllOf) SetBigPicture(v string)`

SetBigPicture sets BigPicture field to given value.

### HasBigPicture

`func (o *BasicNotificationAllOf) HasBigPicture() bool`

HasBigPicture returns a boolean if a field has been set.

### SetBigPictureNil

`func (o *BasicNotificationAllOf) SetBigPictureNil(b bool)`

 SetBigPictureNil sets the value for BigPicture to be an explicit nil

### UnsetBigPicture
`func (o *BasicNotificationAllOf) UnsetBigPicture()`

UnsetBigPicture ensures that no value is present for BigPicture, not even an explicit nil
### GetHuaweiBigPicture

`func (o *BasicNotificationAllOf) GetHuaweiBigPicture() string`

GetHuaweiBigPicture returns the HuaweiBigPicture field if non-nil, zero value otherwise.

### GetHuaweiBigPictureOk

`func (o *BasicNotificationAllOf) GetHuaweiBigPictureOk() (*string, bool)`

GetHuaweiBigPictureOk returns a tuple with the HuaweiBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiBigPicture

`func (o *BasicNotificationAllOf) SetHuaweiBigPicture(v string)`

SetHuaweiBigPicture sets HuaweiBigPicture field to given value.

### HasHuaweiBigPicture

`func (o *BasicNotificationAllOf) HasHuaweiBigPicture() bool`

HasHuaweiBigPicture returns a boolean if a field has been set.

### SetHuaweiBigPictureNil

`func (o *BasicNotificationAllOf) SetHuaweiBigPictureNil(b bool)`

 SetHuaweiBigPictureNil sets the value for HuaweiBigPicture to be an explicit nil

### UnsetHuaweiBigPicture
`func (o *BasicNotificationAllOf) UnsetHuaweiBigPicture()`

UnsetHuaweiBigPicture ensures that no value is present for HuaweiBigPicture, not even an explicit nil
### GetAdmBigPicture

`func (o *BasicNotificationAllOf) GetAdmBigPicture() string`

GetAdmBigPicture returns the AdmBigPicture field if non-nil, zero value otherwise.

### GetAdmBigPictureOk

`func (o *BasicNotificationAllOf) GetAdmBigPictureOk() (*string, bool)`

GetAdmBigPictureOk returns a tuple with the AdmBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmBigPicture

`func (o *BasicNotificationAllOf) SetAdmBigPicture(v string)`

SetAdmBigPicture sets AdmBigPicture field to given value.

### HasAdmBigPicture

`func (o *BasicNotificationAllOf) HasAdmBigPicture() bool`

HasAdmBigPicture returns a boolean if a field has been set.

### SetAdmBigPictureNil

`func (o *BasicNotificationAllOf) SetAdmBigPictureNil(b bool)`

 SetAdmBigPictureNil sets the value for AdmBigPicture to be an explicit nil

### UnsetAdmBigPicture
`func (o *BasicNotificationAllOf) UnsetAdmBigPicture()`

UnsetAdmBigPicture ensures that no value is present for AdmBigPicture, not even an explicit nil
### GetChromeBigPicture

`func (o *BasicNotificationAllOf) GetChromeBigPicture() string`

GetChromeBigPicture returns the ChromeBigPicture field if non-nil, zero value otherwise.

### GetChromeBigPictureOk

`func (o *BasicNotificationAllOf) GetChromeBigPictureOk() (*string, bool)`

GetChromeBigPictureOk returns a tuple with the ChromeBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeBigPicture

`func (o *BasicNotificationAllOf) SetChromeBigPicture(v string)`

SetChromeBigPicture sets ChromeBigPicture field to given value.

### HasChromeBigPicture

`func (o *BasicNotificationAllOf) HasChromeBigPicture() bool`

HasChromeBigPicture returns a boolean if a field has been set.

### SetChromeBigPictureNil

`func (o *BasicNotificationAllOf) SetChromeBigPictureNil(b bool)`

 SetChromeBigPictureNil sets the value for ChromeBigPicture to be an explicit nil

### UnsetChromeBigPicture
`func (o *BasicNotificationAllOf) UnsetChromeBigPicture()`

UnsetChromeBigPicture ensures that no value is present for ChromeBigPicture, not even an explicit nil
### GetChromeWebImage

`func (o *BasicNotificationAllOf) GetChromeWebImage() string`

GetChromeWebImage returns the ChromeWebImage field if non-nil, zero value otherwise.

### GetChromeWebImageOk

`func (o *BasicNotificationAllOf) GetChromeWebImageOk() (*string, bool)`

GetChromeWebImageOk returns a tuple with the ChromeWebImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebImage

`func (o *BasicNotificationAllOf) SetChromeWebImage(v string)`

SetChromeWebImage sets ChromeWebImage field to given value.

### HasChromeWebImage

`func (o *BasicNotificationAllOf) HasChromeWebImage() bool`

HasChromeWebImage returns a boolean if a field has been set.

### SetChromeWebImageNil

`func (o *BasicNotificationAllOf) SetChromeWebImageNil(b bool)`

 SetChromeWebImageNil sets the value for ChromeWebImage to be an explicit nil

### UnsetChromeWebImage
`func (o *BasicNotificationAllOf) UnsetChromeWebImage()`

UnsetChromeWebImage ensures that no value is present for ChromeWebImage, not even an explicit nil
### GetButtons

`func (o *BasicNotificationAllOf) GetButtons() []Button`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *BasicNotificationAllOf) GetButtonsOk() (*[]Button, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *BasicNotificationAllOf) SetButtons(v []Button)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *BasicNotificationAllOf) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### SetButtonsNil

`func (o *BasicNotificationAllOf) SetButtonsNil(b bool)`

 SetButtonsNil sets the value for Buttons to be an explicit nil

### UnsetButtons
`func (o *BasicNotificationAllOf) UnsetButtons()`

UnsetButtons ensures that no value is present for Buttons, not even an explicit nil
### GetWebButtons

`func (o *BasicNotificationAllOf) GetWebButtons() []Button`

GetWebButtons returns the WebButtons field if non-nil, zero value otherwise.

### GetWebButtonsOk

`func (o *BasicNotificationAllOf) GetWebButtonsOk() (*[]Button, bool)`

GetWebButtonsOk returns a tuple with the WebButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebButtons

`func (o *BasicNotificationAllOf) SetWebButtons(v []Button)`

SetWebButtons sets WebButtons field to given value.

### HasWebButtons

`func (o *BasicNotificationAllOf) HasWebButtons() bool`

HasWebButtons returns a boolean if a field has been set.

### SetWebButtonsNil

`func (o *BasicNotificationAllOf) SetWebButtonsNil(b bool)`

 SetWebButtonsNil sets the value for WebButtons to be an explicit nil

### UnsetWebButtons
`func (o *BasicNotificationAllOf) UnsetWebButtons()`

UnsetWebButtons ensures that no value is present for WebButtons, not even an explicit nil
### GetIosCategory

`func (o *BasicNotificationAllOf) GetIosCategory() string`

GetIosCategory returns the IosCategory field if non-nil, zero value otherwise.

### GetIosCategoryOk

`func (o *BasicNotificationAllOf) GetIosCategoryOk() (*string, bool)`

GetIosCategoryOk returns a tuple with the IosCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosCategory

`func (o *BasicNotificationAllOf) SetIosCategory(v string)`

SetIosCategory sets IosCategory field to given value.

### HasIosCategory

`func (o *BasicNotificationAllOf) HasIosCategory() bool`

HasIosCategory returns a boolean if a field has been set.

### SetIosCategoryNil

`func (o *BasicNotificationAllOf) SetIosCategoryNil(b bool)`

 SetIosCategoryNil sets the value for IosCategory to be an explicit nil

### UnsetIosCategory
`func (o *BasicNotificationAllOf) UnsetIosCategory()`

UnsetIosCategory ensures that no value is present for IosCategory, not even an explicit nil
### GetAndroidChannelId

`func (o *BasicNotificationAllOf) GetAndroidChannelId() string`

GetAndroidChannelId returns the AndroidChannelId field if non-nil, zero value otherwise.

### GetAndroidChannelIdOk

`func (o *BasicNotificationAllOf) GetAndroidChannelIdOk() (*string, bool)`

GetAndroidChannelIdOk returns a tuple with the AndroidChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidChannelId

`func (o *BasicNotificationAllOf) SetAndroidChannelId(v string)`

SetAndroidChannelId sets AndroidChannelId field to given value.

### HasAndroidChannelId

`func (o *BasicNotificationAllOf) HasAndroidChannelId() bool`

HasAndroidChannelId returns a boolean if a field has been set.

### GetHuaweiChannelId

`func (o *BasicNotificationAllOf) GetHuaweiChannelId() string`

GetHuaweiChannelId returns the HuaweiChannelId field if non-nil, zero value otherwise.

### GetHuaweiChannelIdOk

`func (o *BasicNotificationAllOf) GetHuaweiChannelIdOk() (*string, bool)`

GetHuaweiChannelIdOk returns a tuple with the HuaweiChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiChannelId

`func (o *BasicNotificationAllOf) SetHuaweiChannelId(v string)`

SetHuaweiChannelId sets HuaweiChannelId field to given value.

### HasHuaweiChannelId

`func (o *BasicNotificationAllOf) HasHuaweiChannelId() bool`

HasHuaweiChannelId returns a boolean if a field has been set.

### SetHuaweiChannelIdNil

`func (o *BasicNotificationAllOf) SetHuaweiChannelIdNil(b bool)`

 SetHuaweiChannelIdNil sets the value for HuaweiChannelId to be an explicit nil

### UnsetHuaweiChannelId
`func (o *BasicNotificationAllOf) UnsetHuaweiChannelId()`

UnsetHuaweiChannelId ensures that no value is present for HuaweiChannelId, not even an explicit nil
### GetExistingAndroidChannelId

`func (o *BasicNotificationAllOf) GetExistingAndroidChannelId() string`

GetExistingAndroidChannelId returns the ExistingAndroidChannelId field if non-nil, zero value otherwise.

### GetExistingAndroidChannelIdOk

`func (o *BasicNotificationAllOf) GetExistingAndroidChannelIdOk() (*string, bool)`

GetExistingAndroidChannelIdOk returns a tuple with the ExistingAndroidChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingAndroidChannelId

`func (o *BasicNotificationAllOf) SetExistingAndroidChannelId(v string)`

SetExistingAndroidChannelId sets ExistingAndroidChannelId field to given value.

### HasExistingAndroidChannelId

`func (o *BasicNotificationAllOf) HasExistingAndroidChannelId() bool`

HasExistingAndroidChannelId returns a boolean if a field has been set.

### GetHuaweiExistingChannelId

`func (o *BasicNotificationAllOf) GetHuaweiExistingChannelId() string`

GetHuaweiExistingChannelId returns the HuaweiExistingChannelId field if non-nil, zero value otherwise.

### GetHuaweiExistingChannelIdOk

`func (o *BasicNotificationAllOf) GetHuaweiExistingChannelIdOk() (*string, bool)`

GetHuaweiExistingChannelIdOk returns a tuple with the HuaweiExistingChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiExistingChannelId

`func (o *BasicNotificationAllOf) SetHuaweiExistingChannelId(v string)`

SetHuaweiExistingChannelId sets HuaweiExistingChannelId field to given value.

### HasHuaweiExistingChannelId

`func (o *BasicNotificationAllOf) HasHuaweiExistingChannelId() bool`

HasHuaweiExistingChannelId returns a boolean if a field has been set.

### SetHuaweiExistingChannelIdNil

`func (o *BasicNotificationAllOf) SetHuaweiExistingChannelIdNil(b bool)`

 SetHuaweiExistingChannelIdNil sets the value for HuaweiExistingChannelId to be an explicit nil

### UnsetHuaweiExistingChannelId
`func (o *BasicNotificationAllOf) UnsetHuaweiExistingChannelId()`

UnsetHuaweiExistingChannelId ensures that no value is present for HuaweiExistingChannelId, not even an explicit nil
### GetAndroidBackgroundLayout

`func (o *BasicNotificationAllOf) GetAndroidBackgroundLayout() BasicNotificationAllOfAndroidBackgroundLayout`

GetAndroidBackgroundLayout returns the AndroidBackgroundLayout field if non-nil, zero value otherwise.

### GetAndroidBackgroundLayoutOk

`func (o *BasicNotificationAllOf) GetAndroidBackgroundLayoutOk() (*BasicNotificationAllOfAndroidBackgroundLayout, bool)`

GetAndroidBackgroundLayoutOk returns a tuple with the AndroidBackgroundLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidBackgroundLayout

`func (o *BasicNotificationAllOf) SetAndroidBackgroundLayout(v BasicNotificationAllOfAndroidBackgroundLayout)`

SetAndroidBackgroundLayout sets AndroidBackgroundLayout field to given value.

### HasAndroidBackgroundLayout

`func (o *BasicNotificationAllOf) HasAndroidBackgroundLayout() bool`

HasAndroidBackgroundLayout returns a boolean if a field has been set.

### GetSmallIcon

`func (o *BasicNotificationAllOf) GetSmallIcon() string`

GetSmallIcon returns the SmallIcon field if non-nil, zero value otherwise.

### GetSmallIconOk

`func (o *BasicNotificationAllOf) GetSmallIconOk() (*string, bool)`

GetSmallIconOk returns a tuple with the SmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallIcon

`func (o *BasicNotificationAllOf) SetSmallIcon(v string)`

SetSmallIcon sets SmallIcon field to given value.

### HasSmallIcon

`func (o *BasicNotificationAllOf) HasSmallIcon() bool`

HasSmallIcon returns a boolean if a field has been set.

### SetSmallIconNil

`func (o *BasicNotificationAllOf) SetSmallIconNil(b bool)`

 SetSmallIconNil sets the value for SmallIcon to be an explicit nil

### UnsetSmallIcon
`func (o *BasicNotificationAllOf) UnsetSmallIcon()`

UnsetSmallIcon ensures that no value is present for SmallIcon, not even an explicit nil
### GetHuaweiSmallIcon

`func (o *BasicNotificationAllOf) GetHuaweiSmallIcon() string`

GetHuaweiSmallIcon returns the HuaweiSmallIcon field if non-nil, zero value otherwise.

### GetHuaweiSmallIconOk

`func (o *BasicNotificationAllOf) GetHuaweiSmallIconOk() (*string, bool)`

GetHuaweiSmallIconOk returns a tuple with the HuaweiSmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiSmallIcon

`func (o *BasicNotificationAllOf) SetHuaweiSmallIcon(v string)`

SetHuaweiSmallIcon sets HuaweiSmallIcon field to given value.

### HasHuaweiSmallIcon

`func (o *BasicNotificationAllOf) HasHuaweiSmallIcon() bool`

HasHuaweiSmallIcon returns a boolean if a field has been set.

### SetHuaweiSmallIconNil

`func (o *BasicNotificationAllOf) SetHuaweiSmallIconNil(b bool)`

 SetHuaweiSmallIconNil sets the value for HuaweiSmallIcon to be an explicit nil

### UnsetHuaweiSmallIcon
`func (o *BasicNotificationAllOf) UnsetHuaweiSmallIcon()`

UnsetHuaweiSmallIcon ensures that no value is present for HuaweiSmallIcon, not even an explicit nil
### GetLargeIcon

`func (o *BasicNotificationAllOf) GetLargeIcon() string`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *BasicNotificationAllOf) GetLargeIconOk() (*string, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeIcon

`func (o *BasicNotificationAllOf) SetLargeIcon(v string)`

SetLargeIcon sets LargeIcon field to given value.

### HasLargeIcon

`func (o *BasicNotificationAllOf) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### SetLargeIconNil

`func (o *BasicNotificationAllOf) SetLargeIconNil(b bool)`

 SetLargeIconNil sets the value for LargeIcon to be an explicit nil

### UnsetLargeIcon
`func (o *BasicNotificationAllOf) UnsetLargeIcon()`

UnsetLargeIcon ensures that no value is present for LargeIcon, not even an explicit nil
### GetHuaweiLargeIcon

`func (o *BasicNotificationAllOf) GetHuaweiLargeIcon() string`

GetHuaweiLargeIcon returns the HuaweiLargeIcon field if non-nil, zero value otherwise.

### GetHuaweiLargeIconOk

`func (o *BasicNotificationAllOf) GetHuaweiLargeIconOk() (*string, bool)`

GetHuaweiLargeIconOk returns a tuple with the HuaweiLargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiLargeIcon

`func (o *BasicNotificationAllOf) SetHuaweiLargeIcon(v string)`

SetHuaweiLargeIcon sets HuaweiLargeIcon field to given value.

### HasHuaweiLargeIcon

`func (o *BasicNotificationAllOf) HasHuaweiLargeIcon() bool`

HasHuaweiLargeIcon returns a boolean if a field has been set.

### SetHuaweiLargeIconNil

`func (o *BasicNotificationAllOf) SetHuaweiLargeIconNil(b bool)`

 SetHuaweiLargeIconNil sets the value for HuaweiLargeIcon to be an explicit nil

### UnsetHuaweiLargeIcon
`func (o *BasicNotificationAllOf) UnsetHuaweiLargeIcon()`

UnsetHuaweiLargeIcon ensures that no value is present for HuaweiLargeIcon, not even an explicit nil
### GetAdmSmallIcon

`func (o *BasicNotificationAllOf) GetAdmSmallIcon() string`

GetAdmSmallIcon returns the AdmSmallIcon field if non-nil, zero value otherwise.

### GetAdmSmallIconOk

`func (o *BasicNotificationAllOf) GetAdmSmallIconOk() (*string, bool)`

GetAdmSmallIconOk returns a tuple with the AdmSmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmSmallIcon

`func (o *BasicNotificationAllOf) SetAdmSmallIcon(v string)`

SetAdmSmallIcon sets AdmSmallIcon field to given value.

### HasAdmSmallIcon

`func (o *BasicNotificationAllOf) HasAdmSmallIcon() bool`

HasAdmSmallIcon returns a boolean if a field has been set.

### SetAdmSmallIconNil

`func (o *BasicNotificationAllOf) SetAdmSmallIconNil(b bool)`

 SetAdmSmallIconNil sets the value for AdmSmallIcon to be an explicit nil

### UnsetAdmSmallIcon
`func (o *BasicNotificationAllOf) UnsetAdmSmallIcon()`

UnsetAdmSmallIcon ensures that no value is present for AdmSmallIcon, not even an explicit nil
### GetAdmLargeIcon

`func (o *BasicNotificationAllOf) GetAdmLargeIcon() string`

GetAdmLargeIcon returns the AdmLargeIcon field if non-nil, zero value otherwise.

### GetAdmLargeIconOk

`func (o *BasicNotificationAllOf) GetAdmLargeIconOk() (*string, bool)`

GetAdmLargeIconOk returns a tuple with the AdmLargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmLargeIcon

`func (o *BasicNotificationAllOf) SetAdmLargeIcon(v string)`

SetAdmLargeIcon sets AdmLargeIcon field to given value.

### HasAdmLargeIcon

`func (o *BasicNotificationAllOf) HasAdmLargeIcon() bool`

HasAdmLargeIcon returns a boolean if a field has been set.

### SetAdmLargeIconNil

`func (o *BasicNotificationAllOf) SetAdmLargeIconNil(b bool)`

 SetAdmLargeIconNil sets the value for AdmLargeIcon to be an explicit nil

### UnsetAdmLargeIcon
`func (o *BasicNotificationAllOf) UnsetAdmLargeIcon()`

UnsetAdmLargeIcon ensures that no value is present for AdmLargeIcon, not even an explicit nil
### GetChromeWebIcon

`func (o *BasicNotificationAllOf) GetChromeWebIcon() string`

GetChromeWebIcon returns the ChromeWebIcon field if non-nil, zero value otherwise.

### GetChromeWebIconOk

`func (o *BasicNotificationAllOf) GetChromeWebIconOk() (*string, bool)`

GetChromeWebIconOk returns a tuple with the ChromeWebIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebIcon

`func (o *BasicNotificationAllOf) SetChromeWebIcon(v string)`

SetChromeWebIcon sets ChromeWebIcon field to given value.

### HasChromeWebIcon

`func (o *BasicNotificationAllOf) HasChromeWebIcon() bool`

HasChromeWebIcon returns a boolean if a field has been set.

### SetChromeWebIconNil

`func (o *BasicNotificationAllOf) SetChromeWebIconNil(b bool)`

 SetChromeWebIconNil sets the value for ChromeWebIcon to be an explicit nil

### UnsetChromeWebIcon
`func (o *BasicNotificationAllOf) UnsetChromeWebIcon()`

UnsetChromeWebIcon ensures that no value is present for ChromeWebIcon, not even an explicit nil
### GetChromeWebBadge

`func (o *BasicNotificationAllOf) GetChromeWebBadge() string`

GetChromeWebBadge returns the ChromeWebBadge field if non-nil, zero value otherwise.

### GetChromeWebBadgeOk

`func (o *BasicNotificationAllOf) GetChromeWebBadgeOk() (*string, bool)`

GetChromeWebBadgeOk returns a tuple with the ChromeWebBadge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebBadge

`func (o *BasicNotificationAllOf) SetChromeWebBadge(v string)`

SetChromeWebBadge sets ChromeWebBadge field to given value.

### HasChromeWebBadge

`func (o *BasicNotificationAllOf) HasChromeWebBadge() bool`

HasChromeWebBadge returns a boolean if a field has been set.

### SetChromeWebBadgeNil

`func (o *BasicNotificationAllOf) SetChromeWebBadgeNil(b bool)`

 SetChromeWebBadgeNil sets the value for ChromeWebBadge to be an explicit nil

### UnsetChromeWebBadge
`func (o *BasicNotificationAllOf) UnsetChromeWebBadge()`

UnsetChromeWebBadge ensures that no value is present for ChromeWebBadge, not even an explicit nil
### GetFirefoxIcon

`func (o *BasicNotificationAllOf) GetFirefoxIcon() string`

GetFirefoxIcon returns the FirefoxIcon field if non-nil, zero value otherwise.

### GetFirefoxIconOk

`func (o *BasicNotificationAllOf) GetFirefoxIconOk() (*string, bool)`

GetFirefoxIconOk returns a tuple with the FirefoxIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirefoxIcon

`func (o *BasicNotificationAllOf) SetFirefoxIcon(v string)`

SetFirefoxIcon sets FirefoxIcon field to given value.

### HasFirefoxIcon

`func (o *BasicNotificationAllOf) HasFirefoxIcon() bool`

HasFirefoxIcon returns a boolean if a field has been set.

### SetFirefoxIconNil

`func (o *BasicNotificationAllOf) SetFirefoxIconNil(b bool)`

 SetFirefoxIconNil sets the value for FirefoxIcon to be an explicit nil

### UnsetFirefoxIcon
`func (o *BasicNotificationAllOf) UnsetFirefoxIcon()`

UnsetFirefoxIcon ensures that no value is present for FirefoxIcon, not even an explicit nil
### GetChromeIcon

`func (o *BasicNotificationAllOf) GetChromeIcon() string`

GetChromeIcon returns the ChromeIcon field if non-nil, zero value otherwise.

### GetChromeIconOk

`func (o *BasicNotificationAllOf) GetChromeIconOk() (*string, bool)`

GetChromeIconOk returns a tuple with the ChromeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeIcon

`func (o *BasicNotificationAllOf) SetChromeIcon(v string)`

SetChromeIcon sets ChromeIcon field to given value.

### HasChromeIcon

`func (o *BasicNotificationAllOf) HasChromeIcon() bool`

HasChromeIcon returns a boolean if a field has been set.

### SetChromeIconNil

`func (o *BasicNotificationAllOf) SetChromeIconNil(b bool)`

 SetChromeIconNil sets the value for ChromeIcon to be an explicit nil

### UnsetChromeIcon
`func (o *BasicNotificationAllOf) UnsetChromeIcon()`

UnsetChromeIcon ensures that no value is present for ChromeIcon, not even an explicit nil
### GetIosSound

`func (o *BasicNotificationAllOf) GetIosSound() string`

GetIosSound returns the IosSound field if non-nil, zero value otherwise.

### GetIosSoundOk

`func (o *BasicNotificationAllOf) GetIosSoundOk() (*string, bool)`

GetIosSoundOk returns a tuple with the IosSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosSound

`func (o *BasicNotificationAllOf) SetIosSound(v string)`

SetIosSound sets IosSound field to given value.

### HasIosSound

`func (o *BasicNotificationAllOf) HasIosSound() bool`

HasIosSound returns a boolean if a field has been set.

### SetIosSoundNil

`func (o *BasicNotificationAllOf) SetIosSoundNil(b bool)`

 SetIosSoundNil sets the value for IosSound to be an explicit nil

### UnsetIosSound
`func (o *BasicNotificationAllOf) UnsetIosSound()`

UnsetIosSound ensures that no value is present for IosSound, not even an explicit nil
### GetAndroidSound

`func (o *BasicNotificationAllOf) GetAndroidSound() string`

GetAndroidSound returns the AndroidSound field if non-nil, zero value otherwise.

### GetAndroidSoundOk

`func (o *BasicNotificationAllOf) GetAndroidSoundOk() (*string, bool)`

GetAndroidSoundOk returns a tuple with the AndroidSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidSound

`func (o *BasicNotificationAllOf) SetAndroidSound(v string)`

SetAndroidSound sets AndroidSound field to given value.

### HasAndroidSound

`func (o *BasicNotificationAllOf) HasAndroidSound() bool`

HasAndroidSound returns a boolean if a field has been set.

### SetAndroidSoundNil

`func (o *BasicNotificationAllOf) SetAndroidSoundNil(b bool)`

 SetAndroidSoundNil sets the value for AndroidSound to be an explicit nil

### UnsetAndroidSound
`func (o *BasicNotificationAllOf) UnsetAndroidSound()`

UnsetAndroidSound ensures that no value is present for AndroidSound, not even an explicit nil
### GetHuaweiSound

`func (o *BasicNotificationAllOf) GetHuaweiSound() string`

GetHuaweiSound returns the HuaweiSound field if non-nil, zero value otherwise.

### GetHuaweiSoundOk

`func (o *BasicNotificationAllOf) GetHuaweiSoundOk() (*string, bool)`

GetHuaweiSoundOk returns a tuple with the HuaweiSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiSound

`func (o *BasicNotificationAllOf) SetHuaweiSound(v string)`

SetHuaweiSound sets HuaweiSound field to given value.

### HasHuaweiSound

`func (o *BasicNotificationAllOf) HasHuaweiSound() bool`

HasHuaweiSound returns a boolean if a field has been set.

### SetHuaweiSoundNil

`func (o *BasicNotificationAllOf) SetHuaweiSoundNil(b bool)`

 SetHuaweiSoundNil sets the value for HuaweiSound to be an explicit nil

### UnsetHuaweiSound
`func (o *BasicNotificationAllOf) UnsetHuaweiSound()`

UnsetHuaweiSound ensures that no value is present for HuaweiSound, not even an explicit nil
### GetAdmSound

`func (o *BasicNotificationAllOf) GetAdmSound() string`

GetAdmSound returns the AdmSound field if non-nil, zero value otherwise.

### GetAdmSoundOk

`func (o *BasicNotificationAllOf) GetAdmSoundOk() (*string, bool)`

GetAdmSoundOk returns a tuple with the AdmSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmSound

`func (o *BasicNotificationAllOf) SetAdmSound(v string)`

SetAdmSound sets AdmSound field to given value.

### HasAdmSound

`func (o *BasicNotificationAllOf) HasAdmSound() bool`

HasAdmSound returns a boolean if a field has been set.

### SetAdmSoundNil

`func (o *BasicNotificationAllOf) SetAdmSoundNil(b bool)`

 SetAdmSoundNil sets the value for AdmSound to be an explicit nil

### UnsetAdmSound
`func (o *BasicNotificationAllOf) UnsetAdmSound()`

UnsetAdmSound ensures that no value is present for AdmSound, not even an explicit nil
### GetWpWnsSound

`func (o *BasicNotificationAllOf) GetWpWnsSound() string`

GetWpWnsSound returns the WpWnsSound field if non-nil, zero value otherwise.

### GetWpWnsSoundOk

`func (o *BasicNotificationAllOf) GetWpWnsSoundOk() (*string, bool)`

GetWpWnsSoundOk returns a tuple with the WpWnsSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpWnsSound

`func (o *BasicNotificationAllOf) SetWpWnsSound(v string)`

SetWpWnsSound sets WpWnsSound field to given value.

### HasWpWnsSound

`func (o *BasicNotificationAllOf) HasWpWnsSound() bool`

HasWpWnsSound returns a boolean if a field has been set.

### SetWpWnsSoundNil

`func (o *BasicNotificationAllOf) SetWpWnsSoundNil(b bool)`

 SetWpWnsSoundNil sets the value for WpWnsSound to be an explicit nil

### UnsetWpWnsSound
`func (o *BasicNotificationAllOf) UnsetWpWnsSound()`

UnsetWpWnsSound ensures that no value is present for WpWnsSound, not even an explicit nil
### GetAndroidLedColor

`func (o *BasicNotificationAllOf) GetAndroidLedColor() string`

GetAndroidLedColor returns the AndroidLedColor field if non-nil, zero value otherwise.

### GetAndroidLedColorOk

`func (o *BasicNotificationAllOf) GetAndroidLedColorOk() (*string, bool)`

GetAndroidLedColorOk returns a tuple with the AndroidLedColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidLedColor

`func (o *BasicNotificationAllOf) SetAndroidLedColor(v string)`

SetAndroidLedColor sets AndroidLedColor field to given value.

### HasAndroidLedColor

`func (o *BasicNotificationAllOf) HasAndroidLedColor() bool`

HasAndroidLedColor returns a boolean if a field has been set.

### SetAndroidLedColorNil

`func (o *BasicNotificationAllOf) SetAndroidLedColorNil(b bool)`

 SetAndroidLedColorNil sets the value for AndroidLedColor to be an explicit nil

### UnsetAndroidLedColor
`func (o *BasicNotificationAllOf) UnsetAndroidLedColor()`

UnsetAndroidLedColor ensures that no value is present for AndroidLedColor, not even an explicit nil
### GetHuaweiLedColor

`func (o *BasicNotificationAllOf) GetHuaweiLedColor() string`

GetHuaweiLedColor returns the HuaweiLedColor field if non-nil, zero value otherwise.

### GetHuaweiLedColorOk

`func (o *BasicNotificationAllOf) GetHuaweiLedColorOk() (*string, bool)`

GetHuaweiLedColorOk returns a tuple with the HuaweiLedColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiLedColor

`func (o *BasicNotificationAllOf) SetHuaweiLedColor(v string)`

SetHuaweiLedColor sets HuaweiLedColor field to given value.

### HasHuaweiLedColor

`func (o *BasicNotificationAllOf) HasHuaweiLedColor() bool`

HasHuaweiLedColor returns a boolean if a field has been set.

### SetHuaweiLedColorNil

`func (o *BasicNotificationAllOf) SetHuaweiLedColorNil(b bool)`

 SetHuaweiLedColorNil sets the value for HuaweiLedColor to be an explicit nil

### UnsetHuaweiLedColor
`func (o *BasicNotificationAllOf) UnsetHuaweiLedColor()`

UnsetHuaweiLedColor ensures that no value is present for HuaweiLedColor, not even an explicit nil
### GetAndroidAccentColor

`func (o *BasicNotificationAllOf) GetAndroidAccentColor() string`

GetAndroidAccentColor returns the AndroidAccentColor field if non-nil, zero value otherwise.

### GetAndroidAccentColorOk

`func (o *BasicNotificationAllOf) GetAndroidAccentColorOk() (*string, bool)`

GetAndroidAccentColorOk returns a tuple with the AndroidAccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidAccentColor

`func (o *BasicNotificationAllOf) SetAndroidAccentColor(v string)`

SetAndroidAccentColor sets AndroidAccentColor field to given value.

### HasAndroidAccentColor

`func (o *BasicNotificationAllOf) HasAndroidAccentColor() bool`

HasAndroidAccentColor returns a boolean if a field has been set.

### SetAndroidAccentColorNil

`func (o *BasicNotificationAllOf) SetAndroidAccentColorNil(b bool)`

 SetAndroidAccentColorNil sets the value for AndroidAccentColor to be an explicit nil

### UnsetAndroidAccentColor
`func (o *BasicNotificationAllOf) UnsetAndroidAccentColor()`

UnsetAndroidAccentColor ensures that no value is present for AndroidAccentColor, not even an explicit nil
### GetHuaweiAccentColor

`func (o *BasicNotificationAllOf) GetHuaweiAccentColor() string`

GetHuaweiAccentColor returns the HuaweiAccentColor field if non-nil, zero value otherwise.

### GetHuaweiAccentColorOk

`func (o *BasicNotificationAllOf) GetHuaweiAccentColorOk() (*string, bool)`

GetHuaweiAccentColorOk returns a tuple with the HuaweiAccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiAccentColor

`func (o *BasicNotificationAllOf) SetHuaweiAccentColor(v string)`

SetHuaweiAccentColor sets HuaweiAccentColor field to given value.

### HasHuaweiAccentColor

`func (o *BasicNotificationAllOf) HasHuaweiAccentColor() bool`

HasHuaweiAccentColor returns a boolean if a field has been set.

### SetHuaweiAccentColorNil

`func (o *BasicNotificationAllOf) SetHuaweiAccentColorNil(b bool)`

 SetHuaweiAccentColorNil sets the value for HuaweiAccentColor to be an explicit nil

### UnsetHuaweiAccentColor
`func (o *BasicNotificationAllOf) UnsetHuaweiAccentColor()`

UnsetHuaweiAccentColor ensures that no value is present for HuaweiAccentColor, not even an explicit nil
### GetAndroidVisibility

`func (o *BasicNotificationAllOf) GetAndroidVisibility() int32`

GetAndroidVisibility returns the AndroidVisibility field if non-nil, zero value otherwise.

### GetAndroidVisibilityOk

`func (o *BasicNotificationAllOf) GetAndroidVisibilityOk() (*int32, bool)`

GetAndroidVisibilityOk returns a tuple with the AndroidVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidVisibility

`func (o *BasicNotificationAllOf) SetAndroidVisibility(v int32)`

SetAndroidVisibility sets AndroidVisibility field to given value.

### HasAndroidVisibility

`func (o *BasicNotificationAllOf) HasAndroidVisibility() bool`

HasAndroidVisibility returns a boolean if a field has been set.

### SetAndroidVisibilityNil

`func (o *BasicNotificationAllOf) SetAndroidVisibilityNil(b bool)`

 SetAndroidVisibilityNil sets the value for AndroidVisibility to be an explicit nil

### UnsetAndroidVisibility
`func (o *BasicNotificationAllOf) UnsetAndroidVisibility()`

UnsetAndroidVisibility ensures that no value is present for AndroidVisibility, not even an explicit nil
### GetHuaweiVisibility

`func (o *BasicNotificationAllOf) GetHuaweiVisibility() int32`

GetHuaweiVisibility returns the HuaweiVisibility field if non-nil, zero value otherwise.

### GetHuaweiVisibilityOk

`func (o *BasicNotificationAllOf) GetHuaweiVisibilityOk() (*int32, bool)`

GetHuaweiVisibilityOk returns a tuple with the HuaweiVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiVisibility

`func (o *BasicNotificationAllOf) SetHuaweiVisibility(v int32)`

SetHuaweiVisibility sets HuaweiVisibility field to given value.

### HasHuaweiVisibility

`func (o *BasicNotificationAllOf) HasHuaweiVisibility() bool`

HasHuaweiVisibility returns a boolean if a field has been set.

### SetHuaweiVisibilityNil

`func (o *BasicNotificationAllOf) SetHuaweiVisibilityNil(b bool)`

 SetHuaweiVisibilityNil sets the value for HuaweiVisibility to be an explicit nil

### UnsetHuaweiVisibility
`func (o *BasicNotificationAllOf) UnsetHuaweiVisibility()`

UnsetHuaweiVisibility ensures that no value is present for HuaweiVisibility, not even an explicit nil
### GetIosBadgeType

`func (o *BasicNotificationAllOf) GetIosBadgeType() string`

GetIosBadgeType returns the IosBadgeType field if non-nil, zero value otherwise.

### GetIosBadgeTypeOk

`func (o *BasicNotificationAllOf) GetIosBadgeTypeOk() (*string, bool)`

GetIosBadgeTypeOk returns a tuple with the IosBadgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosBadgeType

`func (o *BasicNotificationAllOf) SetIosBadgeType(v string)`

SetIosBadgeType sets IosBadgeType field to given value.

### HasIosBadgeType

`func (o *BasicNotificationAllOf) HasIosBadgeType() bool`

HasIosBadgeType returns a boolean if a field has been set.

### SetIosBadgeTypeNil

`func (o *BasicNotificationAllOf) SetIosBadgeTypeNil(b bool)`

 SetIosBadgeTypeNil sets the value for IosBadgeType to be an explicit nil

### UnsetIosBadgeType
`func (o *BasicNotificationAllOf) UnsetIosBadgeType()`

UnsetIosBadgeType ensures that no value is present for IosBadgeType, not even an explicit nil
### GetIosBadgeCount

`func (o *BasicNotificationAllOf) GetIosBadgeCount() int32`

GetIosBadgeCount returns the IosBadgeCount field if non-nil, zero value otherwise.

### GetIosBadgeCountOk

`func (o *BasicNotificationAllOf) GetIosBadgeCountOk() (*int32, bool)`

GetIosBadgeCountOk returns a tuple with the IosBadgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosBadgeCount

`func (o *BasicNotificationAllOf) SetIosBadgeCount(v int32)`

SetIosBadgeCount sets IosBadgeCount field to given value.

### HasIosBadgeCount

`func (o *BasicNotificationAllOf) HasIosBadgeCount() bool`

HasIosBadgeCount returns a boolean if a field has been set.

### SetIosBadgeCountNil

`func (o *BasicNotificationAllOf) SetIosBadgeCountNil(b bool)`

 SetIosBadgeCountNil sets the value for IosBadgeCount to be an explicit nil

### UnsetIosBadgeCount
`func (o *BasicNotificationAllOf) UnsetIosBadgeCount()`

UnsetIosBadgeCount ensures that no value is present for IosBadgeCount, not even an explicit nil
### GetCollapseId

`func (o *BasicNotificationAllOf) GetCollapseId() string`

GetCollapseId returns the CollapseId field if non-nil, zero value otherwise.

### GetCollapseIdOk

`func (o *BasicNotificationAllOf) GetCollapseIdOk() (*string, bool)`

GetCollapseIdOk returns a tuple with the CollapseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseId

`func (o *BasicNotificationAllOf) SetCollapseId(v string)`

SetCollapseId sets CollapseId field to given value.

### HasCollapseId

`func (o *BasicNotificationAllOf) HasCollapseId() bool`

HasCollapseId returns a boolean if a field has been set.

### GetWebPushTopic

`func (o *BasicNotificationAllOf) GetWebPushTopic() string`

GetWebPushTopic returns the WebPushTopic field if non-nil, zero value otherwise.

### GetWebPushTopicOk

`func (o *BasicNotificationAllOf) GetWebPushTopicOk() (*string, bool)`

GetWebPushTopicOk returns a tuple with the WebPushTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPushTopic

`func (o *BasicNotificationAllOf) SetWebPushTopic(v string)`

SetWebPushTopic sets WebPushTopic field to given value.

### HasWebPushTopic

`func (o *BasicNotificationAllOf) HasWebPushTopic() bool`

HasWebPushTopic returns a boolean if a field has been set.

### SetWebPushTopicNil

`func (o *BasicNotificationAllOf) SetWebPushTopicNil(b bool)`

 SetWebPushTopicNil sets the value for WebPushTopic to be an explicit nil

### UnsetWebPushTopic
`func (o *BasicNotificationAllOf) UnsetWebPushTopic()`

UnsetWebPushTopic ensures that no value is present for WebPushTopic, not even an explicit nil
### GetApnsAlert

`func (o *BasicNotificationAllOf) GetApnsAlert() map[string]interface{}`

GetApnsAlert returns the ApnsAlert field if non-nil, zero value otherwise.

### GetApnsAlertOk

`func (o *BasicNotificationAllOf) GetApnsAlertOk() (*map[string]interface{}, bool)`

GetApnsAlertOk returns a tuple with the ApnsAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsAlert

`func (o *BasicNotificationAllOf) SetApnsAlert(v map[string]interface{})`

SetApnsAlert sets ApnsAlert field to given value.

### HasApnsAlert

`func (o *BasicNotificationAllOf) HasApnsAlert() bool`

HasApnsAlert returns a boolean if a field has been set.

### SetApnsAlertNil

`func (o *BasicNotificationAllOf) SetApnsAlertNil(b bool)`

 SetApnsAlertNil sets the value for ApnsAlert to be an explicit nil

### UnsetApnsAlert
`func (o *BasicNotificationAllOf) UnsetApnsAlert()`

UnsetApnsAlert ensures that no value is present for ApnsAlert, not even an explicit nil
### GetDelayedOption

`func (o *BasicNotificationAllOf) GetDelayedOption() string`

GetDelayedOption returns the DelayedOption field if non-nil, zero value otherwise.

### GetDelayedOptionOk

`func (o *BasicNotificationAllOf) GetDelayedOptionOk() (*string, bool)`

GetDelayedOptionOk returns a tuple with the DelayedOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedOption

`func (o *BasicNotificationAllOf) SetDelayedOption(v string)`

SetDelayedOption sets DelayedOption field to given value.

### HasDelayedOption

`func (o *BasicNotificationAllOf) HasDelayedOption() bool`

HasDelayedOption returns a boolean if a field has been set.

### SetDelayedOptionNil

`func (o *BasicNotificationAllOf) SetDelayedOptionNil(b bool)`

 SetDelayedOptionNil sets the value for DelayedOption to be an explicit nil

### UnsetDelayedOption
`func (o *BasicNotificationAllOf) UnsetDelayedOption()`

UnsetDelayedOption ensures that no value is present for DelayedOption, not even an explicit nil
### GetDeliveryTimeOfDay

`func (o *BasicNotificationAllOf) GetDeliveryTimeOfDay() string`

GetDeliveryTimeOfDay returns the DeliveryTimeOfDay field if non-nil, zero value otherwise.

### GetDeliveryTimeOfDayOk

`func (o *BasicNotificationAllOf) GetDeliveryTimeOfDayOk() (*string, bool)`

GetDeliveryTimeOfDayOk returns a tuple with the DeliveryTimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeOfDay

`func (o *BasicNotificationAllOf) SetDeliveryTimeOfDay(v string)`

SetDeliveryTimeOfDay sets DeliveryTimeOfDay field to given value.

### HasDeliveryTimeOfDay

`func (o *BasicNotificationAllOf) HasDeliveryTimeOfDay() bool`

HasDeliveryTimeOfDay returns a boolean if a field has been set.

### SetDeliveryTimeOfDayNil

`func (o *BasicNotificationAllOf) SetDeliveryTimeOfDayNil(b bool)`

 SetDeliveryTimeOfDayNil sets the value for DeliveryTimeOfDay to be an explicit nil

### UnsetDeliveryTimeOfDay
`func (o *BasicNotificationAllOf) UnsetDeliveryTimeOfDay()`

UnsetDeliveryTimeOfDay ensures that no value is present for DeliveryTimeOfDay, not even an explicit nil
### GetTtl

`func (o *BasicNotificationAllOf) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *BasicNotificationAllOf) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *BasicNotificationAllOf) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *BasicNotificationAllOf) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *BasicNotificationAllOf) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *BasicNotificationAllOf) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetPriority

`func (o *BasicNotificationAllOf) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BasicNotificationAllOf) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BasicNotificationAllOf) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BasicNotificationAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *BasicNotificationAllOf) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *BasicNotificationAllOf) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetApnsPushTypeOverride

`func (o *BasicNotificationAllOf) GetApnsPushTypeOverride() string`

GetApnsPushTypeOverride returns the ApnsPushTypeOverride field if non-nil, zero value otherwise.

### GetApnsPushTypeOverrideOk

`func (o *BasicNotificationAllOf) GetApnsPushTypeOverrideOk() (*string, bool)`

GetApnsPushTypeOverrideOk returns a tuple with the ApnsPushTypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsPushTypeOverride

`func (o *BasicNotificationAllOf) SetApnsPushTypeOverride(v string)`

SetApnsPushTypeOverride sets ApnsPushTypeOverride field to given value.

### HasApnsPushTypeOverride

`func (o *BasicNotificationAllOf) HasApnsPushTypeOverride() bool`

HasApnsPushTypeOverride returns a boolean if a field has been set.

### GetThrottleRatePerMinute

`func (o *BasicNotificationAllOf) GetThrottleRatePerMinute() string`

GetThrottleRatePerMinute returns the ThrottleRatePerMinute field if non-nil, zero value otherwise.

### GetThrottleRatePerMinuteOk

`func (o *BasicNotificationAllOf) GetThrottleRatePerMinuteOk() (*string, bool)`

GetThrottleRatePerMinuteOk returns a tuple with the ThrottleRatePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRatePerMinute

`func (o *BasicNotificationAllOf) SetThrottleRatePerMinute(v string)`

SetThrottleRatePerMinute sets ThrottleRatePerMinute field to given value.

### HasThrottleRatePerMinute

`func (o *BasicNotificationAllOf) HasThrottleRatePerMinute() bool`

HasThrottleRatePerMinute returns a boolean if a field has been set.

### SetThrottleRatePerMinuteNil

`func (o *BasicNotificationAllOf) SetThrottleRatePerMinuteNil(b bool)`

 SetThrottleRatePerMinuteNil sets the value for ThrottleRatePerMinute to be an explicit nil

### UnsetThrottleRatePerMinute
`func (o *BasicNotificationAllOf) UnsetThrottleRatePerMinute()`

UnsetThrottleRatePerMinute ensures that no value is present for ThrottleRatePerMinute, not even an explicit nil
### GetAndroidGroup

`func (o *BasicNotificationAllOf) GetAndroidGroup() string`

GetAndroidGroup returns the AndroidGroup field if non-nil, zero value otherwise.

### GetAndroidGroupOk

`func (o *BasicNotificationAllOf) GetAndroidGroupOk() (*string, bool)`

GetAndroidGroupOk returns a tuple with the AndroidGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidGroup

`func (o *BasicNotificationAllOf) SetAndroidGroup(v string)`

SetAndroidGroup sets AndroidGroup field to given value.

### HasAndroidGroup

`func (o *BasicNotificationAllOf) HasAndroidGroup() bool`

HasAndroidGroup returns a boolean if a field has been set.

### SetAndroidGroupNil

`func (o *BasicNotificationAllOf) SetAndroidGroupNil(b bool)`

 SetAndroidGroupNil sets the value for AndroidGroup to be an explicit nil

### UnsetAndroidGroup
`func (o *BasicNotificationAllOf) UnsetAndroidGroup()`

UnsetAndroidGroup ensures that no value is present for AndroidGroup, not even an explicit nil
### GetAndroidGroupMessage

`func (o *BasicNotificationAllOf) GetAndroidGroupMessage() string`

GetAndroidGroupMessage returns the AndroidGroupMessage field if non-nil, zero value otherwise.

### GetAndroidGroupMessageOk

`func (o *BasicNotificationAllOf) GetAndroidGroupMessageOk() (*string, bool)`

GetAndroidGroupMessageOk returns a tuple with the AndroidGroupMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidGroupMessage

`func (o *BasicNotificationAllOf) SetAndroidGroupMessage(v string)`

SetAndroidGroupMessage sets AndroidGroupMessage field to given value.

### HasAndroidGroupMessage

`func (o *BasicNotificationAllOf) HasAndroidGroupMessage() bool`

HasAndroidGroupMessage returns a boolean if a field has been set.

### SetAndroidGroupMessageNil

`func (o *BasicNotificationAllOf) SetAndroidGroupMessageNil(b bool)`

 SetAndroidGroupMessageNil sets the value for AndroidGroupMessage to be an explicit nil

### UnsetAndroidGroupMessage
`func (o *BasicNotificationAllOf) UnsetAndroidGroupMessage()`

UnsetAndroidGroupMessage ensures that no value is present for AndroidGroupMessage, not even an explicit nil
### GetAdmGroup

`func (o *BasicNotificationAllOf) GetAdmGroup() string`

GetAdmGroup returns the AdmGroup field if non-nil, zero value otherwise.

### GetAdmGroupOk

`func (o *BasicNotificationAllOf) GetAdmGroupOk() (*string, bool)`

GetAdmGroupOk returns a tuple with the AdmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmGroup

`func (o *BasicNotificationAllOf) SetAdmGroup(v string)`

SetAdmGroup sets AdmGroup field to given value.

### HasAdmGroup

`func (o *BasicNotificationAllOf) HasAdmGroup() bool`

HasAdmGroup returns a boolean if a field has been set.

### SetAdmGroupNil

`func (o *BasicNotificationAllOf) SetAdmGroupNil(b bool)`

 SetAdmGroupNil sets the value for AdmGroup to be an explicit nil

### UnsetAdmGroup
`func (o *BasicNotificationAllOf) UnsetAdmGroup()`

UnsetAdmGroup ensures that no value is present for AdmGroup, not even an explicit nil
### GetAdmGroupMessage

`func (o *BasicNotificationAllOf) GetAdmGroupMessage() map[string]interface{}`

GetAdmGroupMessage returns the AdmGroupMessage field if non-nil, zero value otherwise.

### GetAdmGroupMessageOk

`func (o *BasicNotificationAllOf) GetAdmGroupMessageOk() (*map[string]interface{}, bool)`

GetAdmGroupMessageOk returns a tuple with the AdmGroupMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmGroupMessage

`func (o *BasicNotificationAllOf) SetAdmGroupMessage(v map[string]interface{})`

SetAdmGroupMessage sets AdmGroupMessage field to given value.

### HasAdmGroupMessage

`func (o *BasicNotificationAllOf) HasAdmGroupMessage() bool`

HasAdmGroupMessage returns a boolean if a field has been set.

### SetAdmGroupMessageNil

`func (o *BasicNotificationAllOf) SetAdmGroupMessageNil(b bool)`

 SetAdmGroupMessageNil sets the value for AdmGroupMessage to be an explicit nil

### UnsetAdmGroupMessage
`func (o *BasicNotificationAllOf) UnsetAdmGroupMessage()`

UnsetAdmGroupMessage ensures that no value is present for AdmGroupMessage, not even an explicit nil
### GetThreadId

`func (o *BasicNotificationAllOf) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *BasicNotificationAllOf) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *BasicNotificationAllOf) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *BasicNotificationAllOf) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### SetThreadIdNil

`func (o *BasicNotificationAllOf) SetThreadIdNil(b bool)`

 SetThreadIdNil sets the value for ThreadId to be an explicit nil

### UnsetThreadId
`func (o *BasicNotificationAllOf) UnsetThreadId()`

UnsetThreadId ensures that no value is present for ThreadId, not even an explicit nil
### GetSummaryArg

`func (o *BasicNotificationAllOf) GetSummaryArg() string`

GetSummaryArg returns the SummaryArg field if non-nil, zero value otherwise.

### GetSummaryArgOk

`func (o *BasicNotificationAllOf) GetSummaryArgOk() (*string, bool)`

GetSummaryArgOk returns a tuple with the SummaryArg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryArg

`func (o *BasicNotificationAllOf) SetSummaryArg(v string)`

SetSummaryArg sets SummaryArg field to given value.

### HasSummaryArg

`func (o *BasicNotificationAllOf) HasSummaryArg() bool`

HasSummaryArg returns a boolean if a field has been set.

### GetSummaryArgCount

`func (o *BasicNotificationAllOf) GetSummaryArgCount() int32`

GetSummaryArgCount returns the SummaryArgCount field if non-nil, zero value otherwise.

### GetSummaryArgCountOk

`func (o *BasicNotificationAllOf) GetSummaryArgCountOk() (*int32, bool)`

GetSummaryArgCountOk returns a tuple with the SummaryArgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryArgCount

`func (o *BasicNotificationAllOf) SetSummaryArgCount(v int32)`

SetSummaryArgCount sets SummaryArgCount field to given value.

### HasSummaryArgCount

`func (o *BasicNotificationAllOf) HasSummaryArgCount() bool`

HasSummaryArgCount returns a boolean if a field has been set.

### GetEmailSubject

`func (o *BasicNotificationAllOf) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *BasicNotificationAllOf) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *BasicNotificationAllOf) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *BasicNotificationAllOf) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### SetEmailSubjectNil

`func (o *BasicNotificationAllOf) SetEmailSubjectNil(b bool)`

 SetEmailSubjectNil sets the value for EmailSubject to be an explicit nil

### UnsetEmailSubject
`func (o *BasicNotificationAllOf) UnsetEmailSubject()`

UnsetEmailSubject ensures that no value is present for EmailSubject, not even an explicit nil
### GetEmailBody

`func (o *BasicNotificationAllOf) GetEmailBody() string`

GetEmailBody returns the EmailBody field if non-nil, zero value otherwise.

### GetEmailBodyOk

`func (o *BasicNotificationAllOf) GetEmailBodyOk() (*string, bool)`

GetEmailBodyOk returns a tuple with the EmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailBody

`func (o *BasicNotificationAllOf) SetEmailBody(v string)`

SetEmailBody sets EmailBody field to given value.

### HasEmailBody

`func (o *BasicNotificationAllOf) HasEmailBody() bool`

HasEmailBody returns a boolean if a field has been set.

### GetEmailFromName

`func (o *BasicNotificationAllOf) GetEmailFromName() string`

GetEmailFromName returns the EmailFromName field if non-nil, zero value otherwise.

### GetEmailFromNameOk

`func (o *BasicNotificationAllOf) GetEmailFromNameOk() (*string, bool)`

GetEmailFromNameOk returns a tuple with the EmailFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromName

`func (o *BasicNotificationAllOf) SetEmailFromName(v string)`

SetEmailFromName sets EmailFromName field to given value.

### HasEmailFromName

`func (o *BasicNotificationAllOf) HasEmailFromName() bool`

HasEmailFromName returns a boolean if a field has been set.

### SetEmailFromNameNil

`func (o *BasicNotificationAllOf) SetEmailFromNameNil(b bool)`

 SetEmailFromNameNil sets the value for EmailFromName to be an explicit nil

### UnsetEmailFromName
`func (o *BasicNotificationAllOf) UnsetEmailFromName()`

UnsetEmailFromName ensures that no value is present for EmailFromName, not even an explicit nil
### GetEmailFromAddress

`func (o *BasicNotificationAllOf) GetEmailFromAddress() string`

GetEmailFromAddress returns the EmailFromAddress field if non-nil, zero value otherwise.

### GetEmailFromAddressOk

`func (o *BasicNotificationAllOf) GetEmailFromAddressOk() (*string, bool)`

GetEmailFromAddressOk returns a tuple with the EmailFromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromAddress

`func (o *BasicNotificationAllOf) SetEmailFromAddress(v string)`

SetEmailFromAddress sets EmailFromAddress field to given value.

### HasEmailFromAddress

`func (o *BasicNotificationAllOf) HasEmailFromAddress() bool`

HasEmailFromAddress returns a boolean if a field has been set.

### SetEmailFromAddressNil

`func (o *BasicNotificationAllOf) SetEmailFromAddressNil(b bool)`

 SetEmailFromAddressNil sets the value for EmailFromAddress to be an explicit nil

### UnsetEmailFromAddress
`func (o *BasicNotificationAllOf) UnsetEmailFromAddress()`

UnsetEmailFromAddress ensures that no value is present for EmailFromAddress, not even an explicit nil
### GetEmailPreheader

`func (o *BasicNotificationAllOf) GetEmailPreheader() string`

GetEmailPreheader returns the EmailPreheader field if non-nil, zero value otherwise.

### GetEmailPreheaderOk

`func (o *BasicNotificationAllOf) GetEmailPreheaderOk() (*string, bool)`

GetEmailPreheaderOk returns a tuple with the EmailPreheader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPreheader

`func (o *BasicNotificationAllOf) SetEmailPreheader(v string)`

SetEmailPreheader sets EmailPreheader field to given value.

### HasEmailPreheader

`func (o *BasicNotificationAllOf) HasEmailPreheader() bool`

HasEmailPreheader returns a boolean if a field has been set.

### SetEmailPreheaderNil

`func (o *BasicNotificationAllOf) SetEmailPreheaderNil(b bool)`

 SetEmailPreheaderNil sets the value for EmailPreheader to be an explicit nil

### UnsetEmailPreheader
`func (o *BasicNotificationAllOf) UnsetEmailPreheader()`

UnsetEmailPreheader ensures that no value is present for EmailPreheader, not even an explicit nil
### GetIncludeUnsubscribed

`func (o *BasicNotificationAllOf) GetIncludeUnsubscribed() bool`

GetIncludeUnsubscribed returns the IncludeUnsubscribed field if non-nil, zero value otherwise.

### GetIncludeUnsubscribedOk

`func (o *BasicNotificationAllOf) GetIncludeUnsubscribedOk() (*bool, bool)`

GetIncludeUnsubscribedOk returns a tuple with the IncludeUnsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUnsubscribed

`func (o *BasicNotificationAllOf) SetIncludeUnsubscribed(v bool)`

SetIncludeUnsubscribed sets IncludeUnsubscribed field to given value.

### HasIncludeUnsubscribed

`func (o *BasicNotificationAllOf) HasIncludeUnsubscribed() bool`

HasIncludeUnsubscribed returns a boolean if a field has been set.

### GetSmsFrom

`func (o *BasicNotificationAllOf) GetSmsFrom() string`

GetSmsFrom returns the SmsFrom field if non-nil, zero value otherwise.

### GetSmsFromOk

`func (o *BasicNotificationAllOf) GetSmsFromOk() (*string, bool)`

GetSmsFromOk returns a tuple with the SmsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFrom

`func (o *BasicNotificationAllOf) SetSmsFrom(v string)`

SetSmsFrom sets SmsFrom field to given value.

### HasSmsFrom

`func (o *BasicNotificationAllOf) HasSmsFrom() bool`

HasSmsFrom returns a boolean if a field has been set.

### SetSmsFromNil

`func (o *BasicNotificationAllOf) SetSmsFromNil(b bool)`

 SetSmsFromNil sets the value for SmsFrom to be an explicit nil

### UnsetSmsFrom
`func (o *BasicNotificationAllOf) UnsetSmsFrom()`

UnsetSmsFrom ensures that no value is present for SmsFrom, not even an explicit nil
### GetSmsMediaUrls

`func (o *BasicNotificationAllOf) GetSmsMediaUrls() []string`

GetSmsMediaUrls returns the SmsMediaUrls field if non-nil, zero value otherwise.

### GetSmsMediaUrlsOk

`func (o *BasicNotificationAllOf) GetSmsMediaUrlsOk() (*[]string, bool)`

GetSmsMediaUrlsOk returns a tuple with the SmsMediaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMediaUrls

`func (o *BasicNotificationAllOf) SetSmsMediaUrls(v []string)`

SetSmsMediaUrls sets SmsMediaUrls field to given value.

### HasSmsMediaUrls

`func (o *BasicNotificationAllOf) HasSmsMediaUrls() bool`

HasSmsMediaUrls returns a boolean if a field has been set.

### SetSmsMediaUrlsNil

`func (o *BasicNotificationAllOf) SetSmsMediaUrlsNil(b bool)`

 SetSmsMediaUrlsNil sets the value for SmsMediaUrls to be an explicit nil

### UnsetSmsMediaUrls
`func (o *BasicNotificationAllOf) UnsetSmsMediaUrls()`

UnsetSmsMediaUrls ensures that no value is present for SmsMediaUrls, not even an explicit nil
### GetFilters

`func (o *BasicNotificationAllOf) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *BasicNotificationAllOf) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *BasicNotificationAllOf) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *BasicNotificationAllOf) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *BasicNotificationAllOf) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *BasicNotificationAllOf) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetCustomData

`func (o *BasicNotificationAllOf) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *BasicNotificationAllOf) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *BasicNotificationAllOf) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *BasicNotificationAllOf) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### SetCustomDataNil

`func (o *BasicNotificationAllOf) SetCustomDataNil(b bool)`

 SetCustomDataNil sets the value for CustomData to be an explicit nil

### UnsetCustomData
`func (o *BasicNotificationAllOf) UnsetCustomData()`

UnsetCustomData ensures that no value is present for CustomData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


