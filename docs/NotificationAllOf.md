# NotificationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **int32** |  | [optional] [readonly] 
**Aggregation** | Pointer to **string** |  | [optional] [readonly] 
**IsIos** | Pointer to **bool** | Indicates whether to send to all devices registered under your app&#39;s Apple iOS platform. | [optional] [default to true]
**IsAndroid** | Pointer to **bool** | Indicates whether to send to all devices registered under your app&#39;s Google Android platform. | [optional] 
**IsHuawei** | Pointer to **bool** | Indicates whether to send to all devices registered under your app&#39;s Huawei Android platform. | [optional] 
**IsAnyWeb** | Pointer to **bool** | Indicates whether to send to all subscribed web browser users, including Chrome, Firefox, and Safari. You may use this instead as a combined flag instead of separately enabling isChromeWeb, isFirefox, and isSafari, though the three options are equivalent to this one.  | [optional] 
**IsChromeWeb** | Pointer to **bool** | Indicates whether to send to all Google Chrome, Chrome on Android, and Mozilla Firefox users registered under your Chrome &amp; Firefox web push platform. | [optional] 
**IsFirefox** | Pointer to **bool** | Indicates whether to send to all Mozilla Firefox desktop users registered under your Firefox web push platform. | [optional] 
**IsSafari** | Pointer to **bool** | Does not support iOS Safari. Indicates whether to send to all Apple&#39;s Safari desktop users registered under your Safari web push platform. Read more iOS Safari | [optional] 
**IsWPWNS** | Pointer to **bool** | Indicates whether to send to all devices registered under your app&#39;s Windows platform. | [optional] 
**IsAdm** | Pointer to **bool** | Indicates whether to send to all devices registered under your app&#39;s Amazon Fire platform. | [optional] 
**IsChrome** | Pointer to **bool** | This flag is not used for web push Please see isChromeWeb for sending to web push users. This flag only applies to Google Chrome Apps &amp; Extensions. Indicates whether to send to all devices registered under your app&#39;s Google Chrome Apps &amp; Extension platform.  | [optional] 
**ChannelForExternalUserIds** | Pointer to **string** | Indicates if the message type when targeting with include_external_user_ids for cases where an email, sms, and/or push subscribers have the same external user id. Example: Use the string \&quot;push\&quot; to indicate you are sending a push notification or the string \&quot;email\&quot;for sending emails or \&quot;sms\&quot;for sending SMS.  | [optional] 
**AppId** | Pointer to **string** | Required: Your OneSignal Application ID, which can be found in Keys &amp; IDs. It is a UUID and looks similar to 8250eaf6-1a58-489e-b136-7c74a864b434.  | [optional] 
**ExternalId** | Pointer to **string** | Correlation and idempotency key. A request received with this parameter will first look for another notification with the same external_id. If one exists, a notification will not be sent, and result of the previous operation will instead be returned. Therefore, if you plan on using this feature, it&#39;s important to use a good source of randomness to generate the UUID passed here. This key is only idempotent for 30 days. After 30 days, the notification could be removed from our system and a notification with the same external_id will be sent again.   See Idempotent Notification Requests for more details writeOnly: true  | [optional] 
**Contents** | Pointer to [**NullableStringMap**](StringMap.md) |  | [optional] 
**Headings** | Pointer to [**NullableStringMap**](StringMap.md) |  | [optional] 
**Subtitle** | Pointer to [**NullableStringMap**](StringMap.md) |  | [optional] 
**Data** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: Huawei A custom map of data that is passed back to your app. Same as using Additional Data within the dashboard. Can use up to 2048 bytes of data. Example: {\&quot;abc\&quot;: 123, \&quot;foo\&quot;: \&quot;bar\&quot;, \&quot;event_performed\&quot;: true, \&quot;amount\&quot;: 12.1}  | [optional] 
**HuaweiMsgType** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Use \&quot;data\&quot; or \&quot;message\&quot; depending on the type of notification you are sending. More details in Data &amp; Background Notifications.  | [optional] 
**Url** | Pointer to **string** | Channel: Push Notifications Platform: All The URL to open in the browser when a user clicks on the notification. Note: iOS needs https or updated NSAppTransportSecurity in plist This field supports inline substitutions. Omit if including web_url or app_url Example: https://onesignal.com  | [optional] 
**WebUrl** | Pointer to **string** | Channel: Push Notifications Platform: All Browsers Same as url but only sent to web push platforms. Including Chrome, Firefox, Safari, Opera, etc. Example: https://onesignal.com  | [optional] 
**AppUrl** | Pointer to **string** | Channel: Push Notifications Platform: All Browsers Same as url but only sent to web push platforms. Including iOS, Android, macOS, Windows, ChromeApps, etc. Example: https://onesignal.com  | [optional] 
**IosAttachments** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: iOS 10+ Adds media attachments to notifications. Set as JSON object, key as a media id of your choice and the value as a valid local filename or URL. User must press and hold on the notification to view. Do not set mutable_content to download attachments. The OneSignal SDK does this automatically Example: {\&quot;id1\&quot;: \&quot;https://domain.com/image.jpg\&quot;}  | [optional] 
**TemplateId** | Pointer to **string** | Channel: Push Notifications Platform: All Use a template you setup on our dashboard. The template_id is the UUID found in the URL when viewing a template on our dashboard. Example: be4a8044-bbd6-11e4-a581-000c2940e62c  | [optional] 
**ContentAvailable** | Pointer to **bool** | Channel: Push Notifications Platform: iOS Sending true wakes your app from background to run custom native code (Apple interprets this as content-available&#x3D;1). Note: Not applicable if the app is in the \&quot;force-quit\&quot; state (i.e app was swiped away). Omit the contents field to prevent displaying a visible notification.  | [optional] 
**MutableContent** | Pointer to **bool** | Channel: Push Notifications Platform: iOS 10+ Always defaults to true and cannot be turned off. Allows tracking of notification receives and changing of the notification content in your app before it is displayed. Triggers didReceive(_:withContentHandler:) on your UNNotificationServiceExtension.  | [optional] 
**TargetContentIdentifier** | Pointer to **string** | Channel: Push Notifications Platform: iOS 13+ Use to target a specific experience in your App Clip, or to target your notification to a specific window in a multi-scene App.  | [optional] 
**BigPicture** | Pointer to **string** | Channel: Push Notifications Platform: Android Picture to display in the expanded view. Can be a drawable resource name or a URL.  | [optional] 
**HuaweiBigPicture** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Picture to display in the expanded view. Can be a drawable resource name or a URL.  | [optional] 
**AdmBigPicture** | Pointer to **string** | Channel: Push Notifications Platform: Amazon Picture to display in the expanded view. Can be a drawable resource name or a URL.  | [optional] 
**ChromeBigPicture** | Pointer to **string** | Channel: Push Notifications Platform: ChromeApp Large picture to display below the notification text. Must be a local URL.  | [optional] 
**ChromeWebImage** | Pointer to **string** | Channel: Push Notifications Platform: Chrome 56+ Sets the web push notification&#39;s large image to be shown below the notification&#39;s title and text. Please see Web Push Notification Icons.  | [optional] 
**Buttons** | Pointer to [**[]Button**](Button.md) | Channel: Push Notifications Platform: iOS 8.0+, Android 4.1+, and derivatives like Amazon Buttons to add to the notification. Icon only works for Android. Buttons show in reverse order of array position i.e. Last item in array shows as first button on device. Example: [{\&quot;id\&quot;: \&quot;id2\&quot;, \&quot;text\&quot;: \&quot;second button\&quot;, \&quot;icon\&quot;: \&quot;ic_menu_share\&quot;}, {\&quot;id\&quot;: \&quot;id1\&quot;, \&quot;text\&quot;: \&quot;first button\&quot;, \&quot;icon\&quot;: \&quot;ic_menu_send\&quot;}]  | [optional] 
**WebButtons** | Pointer to [**[]Button**](Button.md) | Channel: Push Notifications Platform: Chrome 48+ Add action buttons to the notification. The id field is required. Example: [{\&quot;id\&quot;: \&quot;like-button\&quot;, \&quot;text\&quot;: \&quot;Like\&quot;, \&quot;icon\&quot;: \&quot;http://i.imgur.com/N8SN8ZS.png\&quot;, \&quot;url\&quot;: \&quot;https://yoursite.com\&quot;}, {\&quot;id\&quot;: \&quot;read-more-button\&quot;, \&quot;text\&quot;: \&quot;Read more\&quot;, \&quot;icon\&quot;: \&quot;http://i.imgur.com/MIxJp1L.png\&quot;, \&quot;url\&quot;: \&quot;https://yoursite.com\&quot;}]  | [optional] 
**IosCategory** | Pointer to **string** | Channel: Push Notifications Platform: iOS Category APS payload, use with registerUserNotificationSettings:categories in your Objective-C / Swift code. Example: calendar category which contains actions like accept and decline iOS 10+ This will trigger your UNNotificationContentExtension whose ID matches this category.  | [optional] 
**AndroidChannelId** | Pointer to **string** | Channel: Push Notifications Platform: Android The Android Oreo Notification Category to send the notification under. See the Category documentation on creating one and getting it&#39;s id.  | [optional] 
**HuaweiChannelId** | Pointer to **string** | Channel: Push Notifications Platform: Huawei The Android Oreo Notification Category to send the notification under. See the Category documentation on creating one and getting it&#39;s id.  | [optional] 
**ExistingAndroidChannelId** | Pointer to **string** | Channel: Push Notifications Platform: Android Use this if you have client side Android Oreo Channels you have already defined in your app with code.  | [optional] 
**HuaweiExistingChannelId** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Use this if you have client side Android Oreo Channels you have already defined in your app with code.  | [optional] 
**AndroidBackgroundLayout** | Pointer to [**NotificationAllOfAndroidBackgroundLayout**](NotificationAllOfAndroidBackgroundLayout.md) |  | [optional] 
**SmallIcon** | Pointer to **string** | Channel: Push Notifications Platform: Android Icon shown in the status bar and on the top left of the notification. If not set a bell icon will be used or ic_stat_onesignal_default if you have set this resource name. See: How to create small icons  | [optional] 
**HuaweiSmallIcon** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Icon shown in the status bar and on the top left of the notification. Use an Android resource path (E.g. /drawable/small_icon). Defaults to your app icon if not set.  | [optional] 
**LargeIcon** | Pointer to **string** | Channel: Push Notifications Platform: Android Can be a drawable resource name or a URL. See: How to create large icons  | [optional] 
**HuaweiLargeIcon** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Can be a drawable resource name or a URL. See: How to create large icons  | [optional] 
**AdmSmallIcon** | Pointer to **string** | Channel: Push Notifications Platform: Amazon If not set a bell icon will be used or ic_stat_onesignal_default if you have set this resource name. See: How to create small icons  | [optional] 
**AdmLargeIcon** | Pointer to **string** | Channel: Push Notifications Platform: Amazon If blank the small_icon is used. Can be a drawable resource name or a URL. See: How to create large icons  | [optional] 
**ChromeWebIcon** | Pointer to **string** | Channel: Push Notifications Platform: Chrome Sets the web push notification&#39;s icon. An image URL linking to a valid image. Common image types are supported; GIF will not animate. We recommend 256x256 (at least 80x80) to display well on high DPI devices. Firefox will also use this icon, unless you specify firefox_icon.  | [optional] 
**ChromeWebBadge** | Pointer to **string** | Channel: Push Notifications Platform: Chrome Sets the web push notification icon for Android devices in the notification shade. Please see Web Push Notification Badge.  | [optional] 
**FirefoxIcon** | Pointer to **string** | Channel: Push Notifications Platform: Firefox Not recommended Few people need to set Firefox-specific icons. We recommend setting chrome_web_icon instead, which Firefox will also use. Sets the web push notification&#39;s icon for Firefox. An image URL linking to a valid image. Common image types are supported; GIF will not animate. We recommend 256x256 (at least 80x80) to display well on high DPI devices.  | [optional] 
**ChromeIcon** | Pointer to **string** | Channel: Push Notifications Platform: ChromeApp This flag is not used for web push For web push, please see chrome_web_icon instead. The local URL to an icon to use. If blank, the app icon will be used.  | [optional] 
**IosSound** | Pointer to **string** | Channel: Push Notifications Platform: iOS Sound file that is included in your app to play instead of the default device notification sound. Pass nil to disable vibration and sound for the notification. Example: \&quot;notification.wav\&quot;  | [optional] 
**AndroidSound** | Pointer to **string** | Channel: Push Notifications Platform: Android &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. Sound file that is included in your app to play instead of the default device notification sound. Pass nil to disable sound for the notification. NOTE: Leave off file extension for Android. Example: \&quot;notification\&quot;  | [optional] 
**HuaweiSound** | Pointer to **string** | Channel: Push Notifications Platform: Huawei &amp;#9888;&amp;#65039;Deprecated, this field ONLY works on EMUI 5 (Android 7 based) and older devices. Please also set Notification Categories / Channels noted above to support EMUI 8 (Android 8 based) devices. Sound file that is included in your app to play instead of the default device notification sound. NOTE: Leave off file extension for and include the full path.  Example: \&quot;/res/raw/notification\&quot;  | [optional] 
**AdmSound** | Pointer to **string** | Channel: Push Notifications Platform: Amazon &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. Sound file that is included in your app to play instead of the default device notification sound. Pass nil to disable sound for the notification. NOTE: Leave off file extension for Android. Example: \&quot;notification\&quot;  | [optional] 
**WpWnsSound** | Pointer to **string** | Channel: Push Notifications Platform: Windows Sound file that is included in your app to play instead of the default device notification sound. Example: \&quot;notification.wav\&quot;  | [optional] 
**AndroidLedColor** | Pointer to **string** | Channel: Push Notifications Platform: Android &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. Sets the devices LED notification light if the device has one. ARGB Hex format. Example(Blue): \&quot;FF0000FF\&quot;  | [optional] 
**HuaweiLedColor** | Pointer to **string** | Channel: Push Notifications Platform: Huawei &amp;#9888;&amp;#65039;Deprecated, this field ONLY works on EMUI 5 (Android 7 based) and older devices. Please also set Notification Categories / Channels noted above to support EMUI 8 (Android 8 based) devices. Sets the devices LED notification light if the device has one. RGB Hex format. Example(Blue): \&quot;0000FF\&quot;  | [optional] 
**AndroidAccentColor** | Pointer to **string** | Channel: Push Notifications Platform: Android Sets the background color of the notification circle to the left of the notification text. Only applies to apps targeting Android API level 21+ on Android 5.0+ devices. Example(Red): \&quot;FFFF0000\&quot;  | [optional] 
**HuaweiAccentColor** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Accent Color used on Action Buttons and Group overflow count. Uses RGB Hex value (E.g. #9900FF). Defaults to device&#39;s theme color if not set.  | [optional] 
**AndroidVisibility** | Pointer to **int32** | Channel: Push Notifications Platform: Android 5.0_ &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. 1 &#x3D; Public (default) (Shows the full message on the lock screen unless the user has disabled all notifications from showing on the lock screen. Please consider the user and mark private if the contents are.) 0 &#x3D; Private (Hides message contents on lock screen if the user set \&quot;Hide sensitive notification content\&quot; in the system settings) -1 &#x3D; Secret (Notification does not show on the lock screen at all)  | [optional] 
**HuaweiVisibility** | Pointer to **int32** | Channel: Push Notifications Platform: Huawei &amp;#9888;&amp;#65039;Deprecated, this field ONLY works on EMUI 5 (Android 7 based) and older devices. Please also set Notification Categories / Channels noted above to support EMUI 8 (Android 8 based) devices. 1 &#x3D; Public (default) (Shows the full message on the lock screen unless the user has disabled all notifications from showing on the lock screen. Please consider the user and mark private if the contents are.) 0 &#x3D; Private (Hides message contents on lock screen if the user set \&quot;Hide sensitive notification content\&quot; in the system settings) -1 &#x3D; Secret (Notification does not show on the lock screen at all)  | [optional] 
**IosBadgeType** | Pointer to **string** | Channel: Push Notifications Platform: iOS Describes whether to set or increase/decrease your app&#39;s iOS badge count by the ios_badgeCount specified count. Can specify None, SetTo, or Increase. &#x60;None&#x60; leaves the count unaffected. &#x60;SetTo&#x60; directly sets the badge count to the number specified in ios_badgeCount. &#x60;Increase&#x60; adds the number specified in ios_badgeCount to the total. Use a negative number to decrease the badge count.  | [optional] 
**IosBadgeCount** | Pointer to **int32** | Channel: Push Notifications Platform: iOS Used with ios_badgeType, describes the value to set or amount to increase/decrease your app&#39;s iOS badge count by. You can use a negative number to decrease the badge count when used with an ios_badgeType of Increase.  | [optional] 
**CollapseId** | Pointer to **string** | Channel: Push Notifications Platform: iOS 10+, Android Only one notification with the same id will be shown on the device. Use the same id to update an existing notification instead of showing a new one. Limit of 64 characters.  | [optional] 
**WebPushTopic** | Pointer to **string** | Channel: Push Notifications Platform: All Browsers Display multiple notifications at once with different topics.  | [optional] 
**ApnsAlert** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: iOS 10+ iOS can localize push notification messages on the client using special parameters such as loc-key. When using the Create Notification endpoint, you must include these parameters inside of a field called apns_alert. Please see Apple&#39;s guide on localizing push notifications to learn more.  | [optional] 
**SendAfter** | Pointer to **string** | Channel: All Schedule notification for future delivery. API defaults to UTC -1100 Examples: All examples are the exact same date &amp; time. \&quot;Thu Sep 24 2015 14:00:00 GMT-0700 (PDT)\&quot; \&quot;September 24th 2015, 2:00:00 pm UTC-07:00\&quot; \&quot;2015-09-24 14:00:00 GMT-0700\&quot; \&quot;Sept 24 2015 14:00:00 GMT-0700\&quot; \&quot;Thu Sep 24 2015 14:00:00 GMT-0700 (Pacific Daylight Time)\&quot; Note: SMS currently only supports send_after parameter.  | [optional] 
**DelayedOption** | Pointer to **string** | Channel: All Possible values are: timezone (Deliver at a specific time-of-day in each users own timezone) last-active Same as Intelligent Delivery . (Deliver at the same time of day as each user last used your app). If send_after is used, this takes effect after the send_after time has elapsed.  | [optional] 
**DeliveryTimeOfDay** | Pointer to **string** | Channel: All Use with delayed_option&#x3D;timezone. Examples: \&quot;9:00AM\&quot; \&quot;21:45\&quot; \&quot;9:45:30\&quot;  | [optional] 
**Ttl** | Pointer to **int32** | Channel: Push Notifications Platform: iOS, Android, Chrome, Firefox, Safari, ChromeWeb Time To Live - In seconds. The notification will be expired if the device does not come back online within this time. The default is 259,200 seconds (3 days). Max value to set is 2419200 seconds (28 days).  | [optional] 
**Priority** | Pointer to **int32** | Channel: Push Notifications Platform: Android, Chrome, ChromeWeb Delivery priority through the push server (example GCM/FCM). Pass 10 for high priority or any other integer for normal priority. Defaults to normal priority for Android and high for iOS. For Android 6.0+ devices setting priority to high will wake the device out of doze mode.  | [optional] 
**ApnsPushTypeOverride** | Pointer to **string** | Channel: Push Notifications Platform: iOS valid values: voip Set the value to voip for sending VoIP Notifications This field maps to the APNS header apns-push-type. Note: alert and background are automatically set by OneSignal  | [optional] 
**ThrottleRatePerMinute** | Pointer to **string** | Channel: All Apps with throttling enabled:   - the parameter value will be used to override the default application throttling value set from the dashboard settings.   - parameter value 0 indicates not to apply throttling to the notification.   - if the parameter is not passed then the default app throttling value will be applied to the notification. Apps with throttling disabled:   - this parameter can be used to throttle delivery for the notification even though throttling is not enabled at the application level. Refer to throttling for more details.  | [optional] 
**AndroidGroup** | Pointer to **string** | Channel: Push Notifications Platform: Android Notifications with the same group will be stacked together using Android&#39;s Notification Grouping feature.  | [optional] 
**AndroidGroupMessage** | Pointer to **string** | Channel: Push Notifications Platform: Android Note: This only works for Android 6 and older. Android 7+ allows full expansion of all message. Summary message to display when 2+ notifications are stacked together. Default is \&quot;# new messages\&quot;. Include $[notif_count] in your message and it will be replaced with the current number. Languages - The value of each key is the message that will be sent to users for that language. \&quot;en\&quot; (English) is required. The key of each hash is either a a 2 character language code or one of zh-Hans/zh-Hant for Simplified or Traditional Chinese. Read more: supported languages. Example: {\&quot;en\&quot;: \&quot;You have $[notif_count] new messages\&quot;}  | [optional] 
**AdmGroup** | Pointer to **string** | Channel: Push Notifications Platform: Amazon Notifications with the same group will be stacked together using Android&#39;s Notification Grouping feature.  | [optional] 
**AdmGroupMessage** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: Amazon Summary message to display when 2+ notifications are stacked together. Default is \&quot;# new messages\&quot;. Include $[notif_count] in your message and it will be replaced with the current number. \&quot;en\&quot; (English) is required. The key of each hash is either a a 2 character language code or one of zh-Hans/zh-Hant for Simplified or Traditional Chinese. The value of each key is the message that will be sent to users for that language. Example: {\&quot;en\&quot;: \&quot;You have $[notif_count] new messages\&quot;}  | [optional] 
**ThreadId** | Pointer to **string** | Channel: Push Notifications Platform: iOS 12+ This parameter is supported in iOS 12 and above. It allows you to group related notifications together. If two notifications have the same thread-id, they will both be added to the same group.  | [optional] 
**SummaryArg** | Pointer to **string** | Channel: Push Notifications Platform: iOS 12+ When using thread_id to create grouped notifications in iOS 12+, you can also control the summary. For example, a grouped notification can say \&quot;12 more notifications from John Doe\&quot;. The summary_arg lets you set the name of the person/thing the notifications are coming from, and will show up as \&quot;X more notifications from summary_arg\&quot;  | [optional] 
**SummaryArgCount** | Pointer to **int32** | Channel: Push Notifications Platform: iOS 12+ When using thread_id, you can also control the count of the number of notifications in the group. For example, if the group already has 12 notifications, and you send a new notification with summary_arg_count &#x3D; 2, the new total will be 14 and the summary will be \&quot;14 more notifications from summary_arg\&quot;  | [optional] 
**EmailSubect** | Pointer to **string** | Channel: Email Required.  The subject of the email.  | [optional] 
**EmailBody** | Pointer to **string** | Channel: Email Required unless template_id is set. HTML suported The body of the email you wish to send. Typically, customers include their own HTML templates here. Must include [unsubscribe_url] in an &lt;a&gt; tag somewhere in the email. Note: any malformed HTML content will be sent to users. Please double-check your HTML is valid.  | [optional] 
**EmailFromName** | Pointer to **string** | Channel: Email The name the email is from. If not specified, will default to \&quot;from name\&quot; set in the OneSignal Dashboard Email Settings.  | [optional] 
**EmailFromAddress** | Pointer to **string** | Channel: Email The email address the email is from. If not specified, will default to \&quot;from email\&quot; set in the OneSignal Dashboard Email Settings.  | [optional] 
**SmsFrom** | Pointer to **string** | Channel: SMS Phone Number used to send SMS. Should be a registered Twilio phone number in E.164 format.  | [optional] 
**SmsMediaUrls** | Pointer to **[]string** | Channel: SMS URLs for the media files to be attached to the SMS content. Limit: 10 media urls with a total max. size of 5MBs.  | [optional] 

## Methods

### NewNotificationAllOf

`func NewNotificationAllOf() *NotificationAllOf`

NewNotificationAllOf instantiates a new NotificationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationAllOfWithDefaults

`func NewNotificationAllOfWithDefaults() *NotificationAllOf`

NewNotificationAllOfWithDefaults instantiates a new NotificationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *NotificationAllOf) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NotificationAllOf) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NotificationAllOf) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *NotificationAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAggregation

`func (o *NotificationAllOf) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *NotificationAllOf) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *NotificationAllOf) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *NotificationAllOf) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetIsIos

`func (o *NotificationAllOf) GetIsIos() bool`

GetIsIos returns the IsIos field if non-nil, zero value otherwise.

### GetIsIosOk

`func (o *NotificationAllOf) GetIsIosOk() (*bool, bool)`

GetIsIosOk returns a tuple with the IsIos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIos

`func (o *NotificationAllOf) SetIsIos(v bool)`

SetIsIos sets IsIos field to given value.

### HasIsIos

`func (o *NotificationAllOf) HasIsIos() bool`

HasIsIos returns a boolean if a field has been set.

### GetIsAndroid

`func (o *NotificationAllOf) GetIsAndroid() bool`

GetIsAndroid returns the IsAndroid field if non-nil, zero value otherwise.

### GetIsAndroidOk

`func (o *NotificationAllOf) GetIsAndroidOk() (*bool, bool)`

GetIsAndroidOk returns a tuple with the IsAndroid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAndroid

`func (o *NotificationAllOf) SetIsAndroid(v bool)`

SetIsAndroid sets IsAndroid field to given value.

### HasIsAndroid

`func (o *NotificationAllOf) HasIsAndroid() bool`

HasIsAndroid returns a boolean if a field has been set.

### GetIsHuawei

`func (o *NotificationAllOf) GetIsHuawei() bool`

GetIsHuawei returns the IsHuawei field if non-nil, zero value otherwise.

### GetIsHuaweiOk

`func (o *NotificationAllOf) GetIsHuaweiOk() (*bool, bool)`

GetIsHuaweiOk returns a tuple with the IsHuawei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHuawei

`func (o *NotificationAllOf) SetIsHuawei(v bool)`

SetIsHuawei sets IsHuawei field to given value.

### HasIsHuawei

`func (o *NotificationAllOf) HasIsHuawei() bool`

HasIsHuawei returns a boolean if a field has been set.

### GetIsAnyWeb

`func (o *NotificationAllOf) GetIsAnyWeb() bool`

GetIsAnyWeb returns the IsAnyWeb field if non-nil, zero value otherwise.

### GetIsAnyWebOk

`func (o *NotificationAllOf) GetIsAnyWebOk() (*bool, bool)`

GetIsAnyWebOk returns a tuple with the IsAnyWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnyWeb

`func (o *NotificationAllOf) SetIsAnyWeb(v bool)`

SetIsAnyWeb sets IsAnyWeb field to given value.

### HasIsAnyWeb

`func (o *NotificationAllOf) HasIsAnyWeb() bool`

HasIsAnyWeb returns a boolean if a field has been set.

### GetIsChromeWeb

`func (o *NotificationAllOf) GetIsChromeWeb() bool`

GetIsChromeWeb returns the IsChromeWeb field if non-nil, zero value otherwise.

### GetIsChromeWebOk

`func (o *NotificationAllOf) GetIsChromeWebOk() (*bool, bool)`

GetIsChromeWebOk returns a tuple with the IsChromeWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChromeWeb

`func (o *NotificationAllOf) SetIsChromeWeb(v bool)`

SetIsChromeWeb sets IsChromeWeb field to given value.

### HasIsChromeWeb

`func (o *NotificationAllOf) HasIsChromeWeb() bool`

HasIsChromeWeb returns a boolean if a field has been set.

### GetIsFirefox

`func (o *NotificationAllOf) GetIsFirefox() bool`

GetIsFirefox returns the IsFirefox field if non-nil, zero value otherwise.

### GetIsFirefoxOk

`func (o *NotificationAllOf) GetIsFirefoxOk() (*bool, bool)`

GetIsFirefoxOk returns a tuple with the IsFirefox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirefox

`func (o *NotificationAllOf) SetIsFirefox(v bool)`

SetIsFirefox sets IsFirefox field to given value.

### HasIsFirefox

`func (o *NotificationAllOf) HasIsFirefox() bool`

HasIsFirefox returns a boolean if a field has been set.

### GetIsSafari

`func (o *NotificationAllOf) GetIsSafari() bool`

GetIsSafari returns the IsSafari field if non-nil, zero value otherwise.

### GetIsSafariOk

`func (o *NotificationAllOf) GetIsSafariOk() (*bool, bool)`

GetIsSafariOk returns a tuple with the IsSafari field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafari

`func (o *NotificationAllOf) SetIsSafari(v bool)`

SetIsSafari sets IsSafari field to given value.

### HasIsSafari

`func (o *NotificationAllOf) HasIsSafari() bool`

HasIsSafari returns a boolean if a field has been set.

### GetIsWPWNS

`func (o *NotificationAllOf) GetIsWPWNS() bool`

GetIsWPWNS returns the IsWPWNS field if non-nil, zero value otherwise.

### GetIsWPWNSOk

`func (o *NotificationAllOf) GetIsWPWNSOk() (*bool, bool)`

GetIsWPWNSOk returns a tuple with the IsWPWNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWPWNS

`func (o *NotificationAllOf) SetIsWPWNS(v bool)`

SetIsWPWNS sets IsWPWNS field to given value.

### HasIsWPWNS

`func (o *NotificationAllOf) HasIsWPWNS() bool`

HasIsWPWNS returns a boolean if a field has been set.

### GetIsAdm

`func (o *NotificationAllOf) GetIsAdm() bool`

GetIsAdm returns the IsAdm field if non-nil, zero value otherwise.

### GetIsAdmOk

`func (o *NotificationAllOf) GetIsAdmOk() (*bool, bool)`

GetIsAdmOk returns a tuple with the IsAdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdm

`func (o *NotificationAllOf) SetIsAdm(v bool)`

SetIsAdm sets IsAdm field to given value.

### HasIsAdm

`func (o *NotificationAllOf) HasIsAdm() bool`

HasIsAdm returns a boolean if a field has been set.

### GetIsChrome

`func (o *NotificationAllOf) GetIsChrome() bool`

GetIsChrome returns the IsChrome field if non-nil, zero value otherwise.

### GetIsChromeOk

`func (o *NotificationAllOf) GetIsChromeOk() (*bool, bool)`

GetIsChromeOk returns a tuple with the IsChrome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChrome

`func (o *NotificationAllOf) SetIsChrome(v bool)`

SetIsChrome sets IsChrome field to given value.

### HasIsChrome

`func (o *NotificationAllOf) HasIsChrome() bool`

HasIsChrome returns a boolean if a field has been set.

### GetChannelForExternalUserIds

`func (o *NotificationAllOf) GetChannelForExternalUserIds() string`

GetChannelForExternalUserIds returns the ChannelForExternalUserIds field if non-nil, zero value otherwise.

### GetChannelForExternalUserIdsOk

`func (o *NotificationAllOf) GetChannelForExternalUserIdsOk() (*string, bool)`

GetChannelForExternalUserIdsOk returns a tuple with the ChannelForExternalUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelForExternalUserIds

`func (o *NotificationAllOf) SetChannelForExternalUserIds(v string)`

SetChannelForExternalUserIds sets ChannelForExternalUserIds field to given value.

### HasChannelForExternalUserIds

`func (o *NotificationAllOf) HasChannelForExternalUserIds() bool`

HasChannelForExternalUserIds returns a boolean if a field has been set.

### GetAppId

`func (o *NotificationAllOf) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *NotificationAllOf) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *NotificationAllOf) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *NotificationAllOf) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetExternalId

`func (o *NotificationAllOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NotificationAllOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NotificationAllOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NotificationAllOf) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetContents

`func (o *NotificationAllOf) GetContents() StringMap`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *NotificationAllOf) GetContentsOk() (*StringMap, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *NotificationAllOf) SetContents(v StringMap)`

SetContents sets Contents field to given value.

### HasContents

`func (o *NotificationAllOf) HasContents() bool`

HasContents returns a boolean if a field has been set.

### SetContentsNil

`func (o *NotificationAllOf) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *NotificationAllOf) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetHeadings

`func (o *NotificationAllOf) GetHeadings() StringMap`

GetHeadings returns the Headings field if non-nil, zero value otherwise.

### GetHeadingsOk

`func (o *NotificationAllOf) GetHeadingsOk() (*StringMap, bool)`

GetHeadingsOk returns a tuple with the Headings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadings

`func (o *NotificationAllOf) SetHeadings(v StringMap)`

SetHeadings sets Headings field to given value.

### HasHeadings

`func (o *NotificationAllOf) HasHeadings() bool`

HasHeadings returns a boolean if a field has been set.

### SetHeadingsNil

`func (o *NotificationAllOf) SetHeadingsNil(b bool)`

 SetHeadingsNil sets the value for Headings to be an explicit nil

### UnsetHeadings
`func (o *NotificationAllOf) UnsetHeadings()`

UnsetHeadings ensures that no value is present for Headings, not even an explicit nil
### GetSubtitle

`func (o *NotificationAllOf) GetSubtitle() StringMap`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *NotificationAllOf) GetSubtitleOk() (*StringMap, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *NotificationAllOf) SetSubtitle(v StringMap)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *NotificationAllOf) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### SetSubtitleNil

`func (o *NotificationAllOf) SetSubtitleNil(b bool)`

 SetSubtitleNil sets the value for Subtitle to be an explicit nil

### UnsetSubtitle
`func (o *NotificationAllOf) UnsetSubtitle()`

UnsetSubtitle ensures that no value is present for Subtitle, not even an explicit nil
### GetData

`func (o *NotificationAllOf) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationAllOf) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationAllOf) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *NotificationAllOf) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHuaweiMsgType

`func (o *NotificationAllOf) GetHuaweiMsgType() string`

GetHuaweiMsgType returns the HuaweiMsgType field if non-nil, zero value otherwise.

### GetHuaweiMsgTypeOk

`func (o *NotificationAllOf) GetHuaweiMsgTypeOk() (*string, bool)`

GetHuaweiMsgTypeOk returns a tuple with the HuaweiMsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiMsgType

`func (o *NotificationAllOf) SetHuaweiMsgType(v string)`

SetHuaweiMsgType sets HuaweiMsgType field to given value.

### HasHuaweiMsgType

`func (o *NotificationAllOf) HasHuaweiMsgType() bool`

HasHuaweiMsgType returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWebUrl

`func (o *NotificationAllOf) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *NotificationAllOf) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *NotificationAllOf) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *NotificationAllOf) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetAppUrl

`func (o *NotificationAllOf) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *NotificationAllOf) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *NotificationAllOf) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.

### HasAppUrl

`func (o *NotificationAllOf) HasAppUrl() bool`

HasAppUrl returns a boolean if a field has been set.

### GetIosAttachments

`func (o *NotificationAllOf) GetIosAttachments() map[string]interface{}`

GetIosAttachments returns the IosAttachments field if non-nil, zero value otherwise.

### GetIosAttachmentsOk

`func (o *NotificationAllOf) GetIosAttachmentsOk() (*map[string]interface{}, bool)`

GetIosAttachmentsOk returns a tuple with the IosAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosAttachments

`func (o *NotificationAllOf) SetIosAttachments(v map[string]interface{})`

SetIosAttachments sets IosAttachments field to given value.

### HasIosAttachments

`func (o *NotificationAllOf) HasIosAttachments() bool`

HasIosAttachments returns a boolean if a field has been set.

### GetTemplateId

`func (o *NotificationAllOf) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *NotificationAllOf) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *NotificationAllOf) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *NotificationAllOf) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetContentAvailable

`func (o *NotificationAllOf) GetContentAvailable() bool`

GetContentAvailable returns the ContentAvailable field if non-nil, zero value otherwise.

### GetContentAvailableOk

`func (o *NotificationAllOf) GetContentAvailableOk() (*bool, bool)`

GetContentAvailableOk returns a tuple with the ContentAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAvailable

`func (o *NotificationAllOf) SetContentAvailable(v bool)`

SetContentAvailable sets ContentAvailable field to given value.

### HasContentAvailable

`func (o *NotificationAllOf) HasContentAvailable() bool`

HasContentAvailable returns a boolean if a field has been set.

### GetMutableContent

`func (o *NotificationAllOf) GetMutableContent() bool`

GetMutableContent returns the MutableContent field if non-nil, zero value otherwise.

### GetMutableContentOk

`func (o *NotificationAllOf) GetMutableContentOk() (*bool, bool)`

GetMutableContentOk returns a tuple with the MutableContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutableContent

`func (o *NotificationAllOf) SetMutableContent(v bool)`

SetMutableContent sets MutableContent field to given value.

### HasMutableContent

`func (o *NotificationAllOf) HasMutableContent() bool`

HasMutableContent returns a boolean if a field has been set.

### GetTargetContentIdentifier

`func (o *NotificationAllOf) GetTargetContentIdentifier() string`

GetTargetContentIdentifier returns the TargetContentIdentifier field if non-nil, zero value otherwise.

### GetTargetContentIdentifierOk

`func (o *NotificationAllOf) GetTargetContentIdentifierOk() (*string, bool)`

GetTargetContentIdentifierOk returns a tuple with the TargetContentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContentIdentifier

`func (o *NotificationAllOf) SetTargetContentIdentifier(v string)`

SetTargetContentIdentifier sets TargetContentIdentifier field to given value.

### HasTargetContentIdentifier

`func (o *NotificationAllOf) HasTargetContentIdentifier() bool`

HasTargetContentIdentifier returns a boolean if a field has been set.

### GetBigPicture

`func (o *NotificationAllOf) GetBigPicture() string`

GetBigPicture returns the BigPicture field if non-nil, zero value otherwise.

### GetBigPictureOk

`func (o *NotificationAllOf) GetBigPictureOk() (*string, bool)`

GetBigPictureOk returns a tuple with the BigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigPicture

`func (o *NotificationAllOf) SetBigPicture(v string)`

SetBigPicture sets BigPicture field to given value.

### HasBigPicture

`func (o *NotificationAllOf) HasBigPicture() bool`

HasBigPicture returns a boolean if a field has been set.

### GetHuaweiBigPicture

`func (o *NotificationAllOf) GetHuaweiBigPicture() string`

GetHuaweiBigPicture returns the HuaweiBigPicture field if non-nil, zero value otherwise.

### GetHuaweiBigPictureOk

`func (o *NotificationAllOf) GetHuaweiBigPictureOk() (*string, bool)`

GetHuaweiBigPictureOk returns a tuple with the HuaweiBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiBigPicture

`func (o *NotificationAllOf) SetHuaweiBigPicture(v string)`

SetHuaweiBigPicture sets HuaweiBigPicture field to given value.

### HasHuaweiBigPicture

`func (o *NotificationAllOf) HasHuaweiBigPicture() bool`

HasHuaweiBigPicture returns a boolean if a field has been set.

### GetAdmBigPicture

`func (o *NotificationAllOf) GetAdmBigPicture() string`

GetAdmBigPicture returns the AdmBigPicture field if non-nil, zero value otherwise.

### GetAdmBigPictureOk

`func (o *NotificationAllOf) GetAdmBigPictureOk() (*string, bool)`

GetAdmBigPictureOk returns a tuple with the AdmBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmBigPicture

`func (o *NotificationAllOf) SetAdmBigPicture(v string)`

SetAdmBigPicture sets AdmBigPicture field to given value.

### HasAdmBigPicture

`func (o *NotificationAllOf) HasAdmBigPicture() bool`

HasAdmBigPicture returns a boolean if a field has been set.

### GetChromeBigPicture

`func (o *NotificationAllOf) GetChromeBigPicture() string`

GetChromeBigPicture returns the ChromeBigPicture field if non-nil, zero value otherwise.

### GetChromeBigPictureOk

`func (o *NotificationAllOf) GetChromeBigPictureOk() (*string, bool)`

GetChromeBigPictureOk returns a tuple with the ChromeBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeBigPicture

`func (o *NotificationAllOf) SetChromeBigPicture(v string)`

SetChromeBigPicture sets ChromeBigPicture field to given value.

### HasChromeBigPicture

`func (o *NotificationAllOf) HasChromeBigPicture() bool`

HasChromeBigPicture returns a boolean if a field has been set.

### GetChromeWebImage

`func (o *NotificationAllOf) GetChromeWebImage() string`

GetChromeWebImage returns the ChromeWebImage field if non-nil, zero value otherwise.

### GetChromeWebImageOk

`func (o *NotificationAllOf) GetChromeWebImageOk() (*string, bool)`

GetChromeWebImageOk returns a tuple with the ChromeWebImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebImage

`func (o *NotificationAllOf) SetChromeWebImage(v string)`

SetChromeWebImage sets ChromeWebImage field to given value.

### HasChromeWebImage

`func (o *NotificationAllOf) HasChromeWebImage() bool`

HasChromeWebImage returns a boolean if a field has been set.

### GetButtons

`func (o *NotificationAllOf) GetButtons() []Button`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *NotificationAllOf) GetButtonsOk() (*[]Button, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *NotificationAllOf) SetButtons(v []Button)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *NotificationAllOf) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetWebButtons

`func (o *NotificationAllOf) GetWebButtons() []Button`

GetWebButtons returns the WebButtons field if non-nil, zero value otherwise.

### GetWebButtonsOk

`func (o *NotificationAllOf) GetWebButtonsOk() (*[]Button, bool)`

GetWebButtonsOk returns a tuple with the WebButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebButtons

`func (o *NotificationAllOf) SetWebButtons(v []Button)`

SetWebButtons sets WebButtons field to given value.

### HasWebButtons

`func (o *NotificationAllOf) HasWebButtons() bool`

HasWebButtons returns a boolean if a field has been set.

### GetIosCategory

`func (o *NotificationAllOf) GetIosCategory() string`

GetIosCategory returns the IosCategory field if non-nil, zero value otherwise.

### GetIosCategoryOk

`func (o *NotificationAllOf) GetIosCategoryOk() (*string, bool)`

GetIosCategoryOk returns a tuple with the IosCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosCategory

`func (o *NotificationAllOf) SetIosCategory(v string)`

SetIosCategory sets IosCategory field to given value.

### HasIosCategory

`func (o *NotificationAllOf) HasIosCategory() bool`

HasIosCategory returns a boolean if a field has been set.

### GetAndroidChannelId

`func (o *NotificationAllOf) GetAndroidChannelId() string`

GetAndroidChannelId returns the AndroidChannelId field if non-nil, zero value otherwise.

### GetAndroidChannelIdOk

`func (o *NotificationAllOf) GetAndroidChannelIdOk() (*string, bool)`

GetAndroidChannelIdOk returns a tuple with the AndroidChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidChannelId

`func (o *NotificationAllOf) SetAndroidChannelId(v string)`

SetAndroidChannelId sets AndroidChannelId field to given value.

### HasAndroidChannelId

`func (o *NotificationAllOf) HasAndroidChannelId() bool`

HasAndroidChannelId returns a boolean if a field has been set.

### GetHuaweiChannelId

`func (o *NotificationAllOf) GetHuaweiChannelId() string`

GetHuaweiChannelId returns the HuaweiChannelId field if non-nil, zero value otherwise.

### GetHuaweiChannelIdOk

`func (o *NotificationAllOf) GetHuaweiChannelIdOk() (*string, bool)`

GetHuaweiChannelIdOk returns a tuple with the HuaweiChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiChannelId

`func (o *NotificationAllOf) SetHuaweiChannelId(v string)`

SetHuaweiChannelId sets HuaweiChannelId field to given value.

### HasHuaweiChannelId

`func (o *NotificationAllOf) HasHuaweiChannelId() bool`

HasHuaweiChannelId returns a boolean if a field has been set.

### GetExistingAndroidChannelId

`func (o *NotificationAllOf) GetExistingAndroidChannelId() string`

GetExistingAndroidChannelId returns the ExistingAndroidChannelId field if non-nil, zero value otherwise.

### GetExistingAndroidChannelIdOk

`func (o *NotificationAllOf) GetExistingAndroidChannelIdOk() (*string, bool)`

GetExistingAndroidChannelIdOk returns a tuple with the ExistingAndroidChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingAndroidChannelId

`func (o *NotificationAllOf) SetExistingAndroidChannelId(v string)`

SetExistingAndroidChannelId sets ExistingAndroidChannelId field to given value.

### HasExistingAndroidChannelId

`func (o *NotificationAllOf) HasExistingAndroidChannelId() bool`

HasExistingAndroidChannelId returns a boolean if a field has been set.

### GetHuaweiExistingChannelId

`func (o *NotificationAllOf) GetHuaweiExistingChannelId() string`

GetHuaweiExistingChannelId returns the HuaweiExistingChannelId field if non-nil, zero value otherwise.

### GetHuaweiExistingChannelIdOk

`func (o *NotificationAllOf) GetHuaweiExistingChannelIdOk() (*string, bool)`

GetHuaweiExistingChannelIdOk returns a tuple with the HuaweiExistingChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiExistingChannelId

`func (o *NotificationAllOf) SetHuaweiExistingChannelId(v string)`

SetHuaweiExistingChannelId sets HuaweiExistingChannelId field to given value.

### HasHuaweiExistingChannelId

`func (o *NotificationAllOf) HasHuaweiExistingChannelId() bool`

HasHuaweiExistingChannelId returns a boolean if a field has been set.

### GetAndroidBackgroundLayout

`func (o *NotificationAllOf) GetAndroidBackgroundLayout() NotificationAllOfAndroidBackgroundLayout`

GetAndroidBackgroundLayout returns the AndroidBackgroundLayout field if non-nil, zero value otherwise.

### GetAndroidBackgroundLayoutOk

`func (o *NotificationAllOf) GetAndroidBackgroundLayoutOk() (*NotificationAllOfAndroidBackgroundLayout, bool)`

GetAndroidBackgroundLayoutOk returns a tuple with the AndroidBackgroundLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidBackgroundLayout

`func (o *NotificationAllOf) SetAndroidBackgroundLayout(v NotificationAllOfAndroidBackgroundLayout)`

SetAndroidBackgroundLayout sets AndroidBackgroundLayout field to given value.

### HasAndroidBackgroundLayout

`func (o *NotificationAllOf) HasAndroidBackgroundLayout() bool`

HasAndroidBackgroundLayout returns a boolean if a field has been set.

### GetSmallIcon

`func (o *NotificationAllOf) GetSmallIcon() string`

GetSmallIcon returns the SmallIcon field if non-nil, zero value otherwise.

### GetSmallIconOk

`func (o *NotificationAllOf) GetSmallIconOk() (*string, bool)`

GetSmallIconOk returns a tuple with the SmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallIcon

`func (o *NotificationAllOf) SetSmallIcon(v string)`

SetSmallIcon sets SmallIcon field to given value.

### HasSmallIcon

`func (o *NotificationAllOf) HasSmallIcon() bool`

HasSmallIcon returns a boolean if a field has been set.

### GetHuaweiSmallIcon

`func (o *NotificationAllOf) GetHuaweiSmallIcon() string`

GetHuaweiSmallIcon returns the HuaweiSmallIcon field if non-nil, zero value otherwise.

### GetHuaweiSmallIconOk

`func (o *NotificationAllOf) GetHuaweiSmallIconOk() (*string, bool)`

GetHuaweiSmallIconOk returns a tuple with the HuaweiSmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiSmallIcon

`func (o *NotificationAllOf) SetHuaweiSmallIcon(v string)`

SetHuaweiSmallIcon sets HuaweiSmallIcon field to given value.

### HasHuaweiSmallIcon

`func (o *NotificationAllOf) HasHuaweiSmallIcon() bool`

HasHuaweiSmallIcon returns a boolean if a field has been set.

### GetLargeIcon

`func (o *NotificationAllOf) GetLargeIcon() string`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *NotificationAllOf) GetLargeIconOk() (*string, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeIcon

`func (o *NotificationAllOf) SetLargeIcon(v string)`

SetLargeIcon sets LargeIcon field to given value.

### HasLargeIcon

`func (o *NotificationAllOf) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### GetHuaweiLargeIcon

`func (o *NotificationAllOf) GetHuaweiLargeIcon() string`

GetHuaweiLargeIcon returns the HuaweiLargeIcon field if non-nil, zero value otherwise.

### GetHuaweiLargeIconOk

`func (o *NotificationAllOf) GetHuaweiLargeIconOk() (*string, bool)`

GetHuaweiLargeIconOk returns a tuple with the HuaweiLargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiLargeIcon

`func (o *NotificationAllOf) SetHuaweiLargeIcon(v string)`

SetHuaweiLargeIcon sets HuaweiLargeIcon field to given value.

### HasHuaweiLargeIcon

`func (o *NotificationAllOf) HasHuaweiLargeIcon() bool`

HasHuaweiLargeIcon returns a boolean if a field has been set.

### GetAdmSmallIcon

`func (o *NotificationAllOf) GetAdmSmallIcon() string`

GetAdmSmallIcon returns the AdmSmallIcon field if non-nil, zero value otherwise.

### GetAdmSmallIconOk

`func (o *NotificationAllOf) GetAdmSmallIconOk() (*string, bool)`

GetAdmSmallIconOk returns a tuple with the AdmSmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmSmallIcon

`func (o *NotificationAllOf) SetAdmSmallIcon(v string)`

SetAdmSmallIcon sets AdmSmallIcon field to given value.

### HasAdmSmallIcon

`func (o *NotificationAllOf) HasAdmSmallIcon() bool`

HasAdmSmallIcon returns a boolean if a field has been set.

### GetAdmLargeIcon

`func (o *NotificationAllOf) GetAdmLargeIcon() string`

GetAdmLargeIcon returns the AdmLargeIcon field if non-nil, zero value otherwise.

### GetAdmLargeIconOk

`func (o *NotificationAllOf) GetAdmLargeIconOk() (*string, bool)`

GetAdmLargeIconOk returns a tuple with the AdmLargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmLargeIcon

`func (o *NotificationAllOf) SetAdmLargeIcon(v string)`

SetAdmLargeIcon sets AdmLargeIcon field to given value.

### HasAdmLargeIcon

`func (o *NotificationAllOf) HasAdmLargeIcon() bool`

HasAdmLargeIcon returns a boolean if a field has been set.

### GetChromeWebIcon

`func (o *NotificationAllOf) GetChromeWebIcon() string`

GetChromeWebIcon returns the ChromeWebIcon field if non-nil, zero value otherwise.

### GetChromeWebIconOk

`func (o *NotificationAllOf) GetChromeWebIconOk() (*string, bool)`

GetChromeWebIconOk returns a tuple with the ChromeWebIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebIcon

`func (o *NotificationAllOf) SetChromeWebIcon(v string)`

SetChromeWebIcon sets ChromeWebIcon field to given value.

### HasChromeWebIcon

`func (o *NotificationAllOf) HasChromeWebIcon() bool`

HasChromeWebIcon returns a boolean if a field has been set.

### GetChromeWebBadge

`func (o *NotificationAllOf) GetChromeWebBadge() string`

GetChromeWebBadge returns the ChromeWebBadge field if non-nil, zero value otherwise.

### GetChromeWebBadgeOk

`func (o *NotificationAllOf) GetChromeWebBadgeOk() (*string, bool)`

GetChromeWebBadgeOk returns a tuple with the ChromeWebBadge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebBadge

`func (o *NotificationAllOf) SetChromeWebBadge(v string)`

SetChromeWebBadge sets ChromeWebBadge field to given value.

### HasChromeWebBadge

`func (o *NotificationAllOf) HasChromeWebBadge() bool`

HasChromeWebBadge returns a boolean if a field has been set.

### GetFirefoxIcon

`func (o *NotificationAllOf) GetFirefoxIcon() string`

GetFirefoxIcon returns the FirefoxIcon field if non-nil, zero value otherwise.

### GetFirefoxIconOk

`func (o *NotificationAllOf) GetFirefoxIconOk() (*string, bool)`

GetFirefoxIconOk returns a tuple with the FirefoxIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirefoxIcon

`func (o *NotificationAllOf) SetFirefoxIcon(v string)`

SetFirefoxIcon sets FirefoxIcon field to given value.

### HasFirefoxIcon

`func (o *NotificationAllOf) HasFirefoxIcon() bool`

HasFirefoxIcon returns a boolean if a field has been set.

### GetChromeIcon

`func (o *NotificationAllOf) GetChromeIcon() string`

GetChromeIcon returns the ChromeIcon field if non-nil, zero value otherwise.

### GetChromeIconOk

`func (o *NotificationAllOf) GetChromeIconOk() (*string, bool)`

GetChromeIconOk returns a tuple with the ChromeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeIcon

`func (o *NotificationAllOf) SetChromeIcon(v string)`

SetChromeIcon sets ChromeIcon field to given value.

### HasChromeIcon

`func (o *NotificationAllOf) HasChromeIcon() bool`

HasChromeIcon returns a boolean if a field has been set.

### GetIosSound

`func (o *NotificationAllOf) GetIosSound() string`

GetIosSound returns the IosSound field if non-nil, zero value otherwise.

### GetIosSoundOk

`func (o *NotificationAllOf) GetIosSoundOk() (*string, bool)`

GetIosSoundOk returns a tuple with the IosSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosSound

`func (o *NotificationAllOf) SetIosSound(v string)`

SetIosSound sets IosSound field to given value.

### HasIosSound

`func (o *NotificationAllOf) HasIosSound() bool`

HasIosSound returns a boolean if a field has been set.

### GetAndroidSound

`func (o *NotificationAllOf) GetAndroidSound() string`

GetAndroidSound returns the AndroidSound field if non-nil, zero value otherwise.

### GetAndroidSoundOk

`func (o *NotificationAllOf) GetAndroidSoundOk() (*string, bool)`

GetAndroidSoundOk returns a tuple with the AndroidSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidSound

`func (o *NotificationAllOf) SetAndroidSound(v string)`

SetAndroidSound sets AndroidSound field to given value.

### HasAndroidSound

`func (o *NotificationAllOf) HasAndroidSound() bool`

HasAndroidSound returns a boolean if a field has been set.

### GetHuaweiSound

`func (o *NotificationAllOf) GetHuaweiSound() string`

GetHuaweiSound returns the HuaweiSound field if non-nil, zero value otherwise.

### GetHuaweiSoundOk

`func (o *NotificationAllOf) GetHuaweiSoundOk() (*string, bool)`

GetHuaweiSoundOk returns a tuple with the HuaweiSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiSound

`func (o *NotificationAllOf) SetHuaweiSound(v string)`

SetHuaweiSound sets HuaweiSound field to given value.

### HasHuaweiSound

`func (o *NotificationAllOf) HasHuaweiSound() bool`

HasHuaweiSound returns a boolean if a field has been set.

### GetAdmSound

`func (o *NotificationAllOf) GetAdmSound() string`

GetAdmSound returns the AdmSound field if non-nil, zero value otherwise.

### GetAdmSoundOk

`func (o *NotificationAllOf) GetAdmSoundOk() (*string, bool)`

GetAdmSoundOk returns a tuple with the AdmSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmSound

`func (o *NotificationAllOf) SetAdmSound(v string)`

SetAdmSound sets AdmSound field to given value.

### HasAdmSound

`func (o *NotificationAllOf) HasAdmSound() bool`

HasAdmSound returns a boolean if a field has been set.

### GetWpWnsSound

`func (o *NotificationAllOf) GetWpWnsSound() string`

GetWpWnsSound returns the WpWnsSound field if non-nil, zero value otherwise.

### GetWpWnsSoundOk

`func (o *NotificationAllOf) GetWpWnsSoundOk() (*string, bool)`

GetWpWnsSoundOk returns a tuple with the WpWnsSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpWnsSound

`func (o *NotificationAllOf) SetWpWnsSound(v string)`

SetWpWnsSound sets WpWnsSound field to given value.

### HasWpWnsSound

`func (o *NotificationAllOf) HasWpWnsSound() bool`

HasWpWnsSound returns a boolean if a field has been set.

### GetAndroidLedColor

`func (o *NotificationAllOf) GetAndroidLedColor() string`

GetAndroidLedColor returns the AndroidLedColor field if non-nil, zero value otherwise.

### GetAndroidLedColorOk

`func (o *NotificationAllOf) GetAndroidLedColorOk() (*string, bool)`

GetAndroidLedColorOk returns a tuple with the AndroidLedColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidLedColor

`func (o *NotificationAllOf) SetAndroidLedColor(v string)`

SetAndroidLedColor sets AndroidLedColor field to given value.

### HasAndroidLedColor

`func (o *NotificationAllOf) HasAndroidLedColor() bool`

HasAndroidLedColor returns a boolean if a field has been set.

### GetHuaweiLedColor

`func (o *NotificationAllOf) GetHuaweiLedColor() string`

GetHuaweiLedColor returns the HuaweiLedColor field if non-nil, zero value otherwise.

### GetHuaweiLedColorOk

`func (o *NotificationAllOf) GetHuaweiLedColorOk() (*string, bool)`

GetHuaweiLedColorOk returns a tuple with the HuaweiLedColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiLedColor

`func (o *NotificationAllOf) SetHuaweiLedColor(v string)`

SetHuaweiLedColor sets HuaweiLedColor field to given value.

### HasHuaweiLedColor

`func (o *NotificationAllOf) HasHuaweiLedColor() bool`

HasHuaweiLedColor returns a boolean if a field has been set.

### GetAndroidAccentColor

`func (o *NotificationAllOf) GetAndroidAccentColor() string`

GetAndroidAccentColor returns the AndroidAccentColor field if non-nil, zero value otherwise.

### GetAndroidAccentColorOk

`func (o *NotificationAllOf) GetAndroidAccentColorOk() (*string, bool)`

GetAndroidAccentColorOk returns a tuple with the AndroidAccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidAccentColor

`func (o *NotificationAllOf) SetAndroidAccentColor(v string)`

SetAndroidAccentColor sets AndroidAccentColor field to given value.

### HasAndroidAccentColor

`func (o *NotificationAllOf) HasAndroidAccentColor() bool`

HasAndroidAccentColor returns a boolean if a field has been set.

### GetHuaweiAccentColor

`func (o *NotificationAllOf) GetHuaweiAccentColor() string`

GetHuaweiAccentColor returns the HuaweiAccentColor field if non-nil, zero value otherwise.

### GetHuaweiAccentColorOk

`func (o *NotificationAllOf) GetHuaweiAccentColorOk() (*string, bool)`

GetHuaweiAccentColorOk returns a tuple with the HuaweiAccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiAccentColor

`func (o *NotificationAllOf) SetHuaweiAccentColor(v string)`

SetHuaweiAccentColor sets HuaweiAccentColor field to given value.

### HasHuaweiAccentColor

`func (o *NotificationAllOf) HasHuaweiAccentColor() bool`

HasHuaweiAccentColor returns a boolean if a field has been set.

### GetAndroidVisibility

`func (o *NotificationAllOf) GetAndroidVisibility() int32`

GetAndroidVisibility returns the AndroidVisibility field if non-nil, zero value otherwise.

### GetAndroidVisibilityOk

`func (o *NotificationAllOf) GetAndroidVisibilityOk() (*int32, bool)`

GetAndroidVisibilityOk returns a tuple with the AndroidVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidVisibility

`func (o *NotificationAllOf) SetAndroidVisibility(v int32)`

SetAndroidVisibility sets AndroidVisibility field to given value.

### HasAndroidVisibility

`func (o *NotificationAllOf) HasAndroidVisibility() bool`

HasAndroidVisibility returns a boolean if a field has been set.

### GetHuaweiVisibility

`func (o *NotificationAllOf) GetHuaweiVisibility() int32`

GetHuaweiVisibility returns the HuaweiVisibility field if non-nil, zero value otherwise.

### GetHuaweiVisibilityOk

`func (o *NotificationAllOf) GetHuaweiVisibilityOk() (*int32, bool)`

GetHuaweiVisibilityOk returns a tuple with the HuaweiVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiVisibility

`func (o *NotificationAllOf) SetHuaweiVisibility(v int32)`

SetHuaweiVisibility sets HuaweiVisibility field to given value.

### HasHuaweiVisibility

`func (o *NotificationAllOf) HasHuaweiVisibility() bool`

HasHuaweiVisibility returns a boolean if a field has been set.

### GetIosBadgeType

`func (o *NotificationAllOf) GetIosBadgeType() string`

GetIosBadgeType returns the IosBadgeType field if non-nil, zero value otherwise.

### GetIosBadgeTypeOk

`func (o *NotificationAllOf) GetIosBadgeTypeOk() (*string, bool)`

GetIosBadgeTypeOk returns a tuple with the IosBadgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosBadgeType

`func (o *NotificationAllOf) SetIosBadgeType(v string)`

SetIosBadgeType sets IosBadgeType field to given value.

### HasIosBadgeType

`func (o *NotificationAllOf) HasIosBadgeType() bool`

HasIosBadgeType returns a boolean if a field has been set.

### GetIosBadgeCount

`func (o *NotificationAllOf) GetIosBadgeCount() int32`

GetIosBadgeCount returns the IosBadgeCount field if non-nil, zero value otherwise.

### GetIosBadgeCountOk

`func (o *NotificationAllOf) GetIosBadgeCountOk() (*int32, bool)`

GetIosBadgeCountOk returns a tuple with the IosBadgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosBadgeCount

`func (o *NotificationAllOf) SetIosBadgeCount(v int32)`

SetIosBadgeCount sets IosBadgeCount field to given value.

### HasIosBadgeCount

`func (o *NotificationAllOf) HasIosBadgeCount() bool`

HasIosBadgeCount returns a boolean if a field has been set.

### GetCollapseId

`func (o *NotificationAllOf) GetCollapseId() string`

GetCollapseId returns the CollapseId field if non-nil, zero value otherwise.

### GetCollapseIdOk

`func (o *NotificationAllOf) GetCollapseIdOk() (*string, bool)`

GetCollapseIdOk returns a tuple with the CollapseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseId

`func (o *NotificationAllOf) SetCollapseId(v string)`

SetCollapseId sets CollapseId field to given value.

### HasCollapseId

`func (o *NotificationAllOf) HasCollapseId() bool`

HasCollapseId returns a boolean if a field has been set.

### GetWebPushTopic

`func (o *NotificationAllOf) GetWebPushTopic() string`

GetWebPushTopic returns the WebPushTopic field if non-nil, zero value otherwise.

### GetWebPushTopicOk

`func (o *NotificationAllOf) GetWebPushTopicOk() (*string, bool)`

GetWebPushTopicOk returns a tuple with the WebPushTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPushTopic

`func (o *NotificationAllOf) SetWebPushTopic(v string)`

SetWebPushTopic sets WebPushTopic field to given value.

### HasWebPushTopic

`func (o *NotificationAllOf) HasWebPushTopic() bool`

HasWebPushTopic returns a boolean if a field has been set.

### GetApnsAlert

`func (o *NotificationAllOf) GetApnsAlert() map[string]interface{}`

GetApnsAlert returns the ApnsAlert field if non-nil, zero value otherwise.

### GetApnsAlertOk

`func (o *NotificationAllOf) GetApnsAlertOk() (*map[string]interface{}, bool)`

GetApnsAlertOk returns a tuple with the ApnsAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsAlert

`func (o *NotificationAllOf) SetApnsAlert(v map[string]interface{})`

SetApnsAlert sets ApnsAlert field to given value.

### HasApnsAlert

`func (o *NotificationAllOf) HasApnsAlert() bool`

HasApnsAlert returns a boolean if a field has been set.

### GetSendAfter

`func (o *NotificationAllOf) GetSendAfter() string`

GetSendAfter returns the SendAfter field if non-nil, zero value otherwise.

### GetSendAfterOk

`func (o *NotificationAllOf) GetSendAfterOk() (*string, bool)`

GetSendAfterOk returns a tuple with the SendAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAfter

`func (o *NotificationAllOf) SetSendAfter(v string)`

SetSendAfter sets SendAfter field to given value.

### HasSendAfter

`func (o *NotificationAllOf) HasSendAfter() bool`

HasSendAfter returns a boolean if a field has been set.

### GetDelayedOption

`func (o *NotificationAllOf) GetDelayedOption() string`

GetDelayedOption returns the DelayedOption field if non-nil, zero value otherwise.

### GetDelayedOptionOk

`func (o *NotificationAllOf) GetDelayedOptionOk() (*string, bool)`

GetDelayedOptionOk returns a tuple with the DelayedOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedOption

`func (o *NotificationAllOf) SetDelayedOption(v string)`

SetDelayedOption sets DelayedOption field to given value.

### HasDelayedOption

`func (o *NotificationAllOf) HasDelayedOption() bool`

HasDelayedOption returns a boolean if a field has been set.

### GetDeliveryTimeOfDay

`func (o *NotificationAllOf) GetDeliveryTimeOfDay() string`

GetDeliveryTimeOfDay returns the DeliveryTimeOfDay field if non-nil, zero value otherwise.

### GetDeliveryTimeOfDayOk

`func (o *NotificationAllOf) GetDeliveryTimeOfDayOk() (*string, bool)`

GetDeliveryTimeOfDayOk returns a tuple with the DeliveryTimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeOfDay

`func (o *NotificationAllOf) SetDeliveryTimeOfDay(v string)`

SetDeliveryTimeOfDay sets DeliveryTimeOfDay field to given value.

### HasDeliveryTimeOfDay

`func (o *NotificationAllOf) HasDeliveryTimeOfDay() bool`

HasDeliveryTimeOfDay returns a boolean if a field has been set.

### GetTtl

`func (o *NotificationAllOf) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *NotificationAllOf) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *NotificationAllOf) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *NotificationAllOf) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetPriority

`func (o *NotificationAllOf) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NotificationAllOf) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NotificationAllOf) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NotificationAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetApnsPushTypeOverride

`func (o *NotificationAllOf) GetApnsPushTypeOverride() string`

GetApnsPushTypeOverride returns the ApnsPushTypeOverride field if non-nil, zero value otherwise.

### GetApnsPushTypeOverrideOk

`func (o *NotificationAllOf) GetApnsPushTypeOverrideOk() (*string, bool)`

GetApnsPushTypeOverrideOk returns a tuple with the ApnsPushTypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsPushTypeOverride

`func (o *NotificationAllOf) SetApnsPushTypeOverride(v string)`

SetApnsPushTypeOverride sets ApnsPushTypeOverride field to given value.

### HasApnsPushTypeOverride

`func (o *NotificationAllOf) HasApnsPushTypeOverride() bool`

HasApnsPushTypeOverride returns a boolean if a field has been set.

### GetThrottleRatePerMinute

`func (o *NotificationAllOf) GetThrottleRatePerMinute() string`

GetThrottleRatePerMinute returns the ThrottleRatePerMinute field if non-nil, zero value otherwise.

### GetThrottleRatePerMinuteOk

`func (o *NotificationAllOf) GetThrottleRatePerMinuteOk() (*string, bool)`

GetThrottleRatePerMinuteOk returns a tuple with the ThrottleRatePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRatePerMinute

`func (o *NotificationAllOf) SetThrottleRatePerMinute(v string)`

SetThrottleRatePerMinute sets ThrottleRatePerMinute field to given value.

### HasThrottleRatePerMinute

`func (o *NotificationAllOf) HasThrottleRatePerMinute() bool`

HasThrottleRatePerMinute returns a boolean if a field has been set.

### GetAndroidGroup

`func (o *NotificationAllOf) GetAndroidGroup() string`

GetAndroidGroup returns the AndroidGroup field if non-nil, zero value otherwise.

### GetAndroidGroupOk

`func (o *NotificationAllOf) GetAndroidGroupOk() (*string, bool)`

GetAndroidGroupOk returns a tuple with the AndroidGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidGroup

`func (o *NotificationAllOf) SetAndroidGroup(v string)`

SetAndroidGroup sets AndroidGroup field to given value.

### HasAndroidGroup

`func (o *NotificationAllOf) HasAndroidGroup() bool`

HasAndroidGroup returns a boolean if a field has been set.

### GetAndroidGroupMessage

`func (o *NotificationAllOf) GetAndroidGroupMessage() string`

GetAndroidGroupMessage returns the AndroidGroupMessage field if non-nil, zero value otherwise.

### GetAndroidGroupMessageOk

`func (o *NotificationAllOf) GetAndroidGroupMessageOk() (*string, bool)`

GetAndroidGroupMessageOk returns a tuple with the AndroidGroupMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidGroupMessage

`func (o *NotificationAllOf) SetAndroidGroupMessage(v string)`

SetAndroidGroupMessage sets AndroidGroupMessage field to given value.

### HasAndroidGroupMessage

`func (o *NotificationAllOf) HasAndroidGroupMessage() bool`

HasAndroidGroupMessage returns a boolean if a field has been set.

### GetAdmGroup

`func (o *NotificationAllOf) GetAdmGroup() string`

GetAdmGroup returns the AdmGroup field if non-nil, zero value otherwise.

### GetAdmGroupOk

`func (o *NotificationAllOf) GetAdmGroupOk() (*string, bool)`

GetAdmGroupOk returns a tuple with the AdmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmGroup

`func (o *NotificationAllOf) SetAdmGroup(v string)`

SetAdmGroup sets AdmGroup field to given value.

### HasAdmGroup

`func (o *NotificationAllOf) HasAdmGroup() bool`

HasAdmGroup returns a boolean if a field has been set.

### GetAdmGroupMessage

`func (o *NotificationAllOf) GetAdmGroupMessage() map[string]interface{}`

GetAdmGroupMessage returns the AdmGroupMessage field if non-nil, zero value otherwise.

### GetAdmGroupMessageOk

`func (o *NotificationAllOf) GetAdmGroupMessageOk() (*map[string]interface{}, bool)`

GetAdmGroupMessageOk returns a tuple with the AdmGroupMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmGroupMessage

`func (o *NotificationAllOf) SetAdmGroupMessage(v map[string]interface{})`

SetAdmGroupMessage sets AdmGroupMessage field to given value.

### HasAdmGroupMessage

`func (o *NotificationAllOf) HasAdmGroupMessage() bool`

HasAdmGroupMessage returns a boolean if a field has been set.

### GetThreadId

`func (o *NotificationAllOf) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *NotificationAllOf) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *NotificationAllOf) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *NotificationAllOf) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetSummaryArg

`func (o *NotificationAllOf) GetSummaryArg() string`

GetSummaryArg returns the SummaryArg field if non-nil, zero value otherwise.

### GetSummaryArgOk

`func (o *NotificationAllOf) GetSummaryArgOk() (*string, bool)`

GetSummaryArgOk returns a tuple with the SummaryArg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryArg

`func (o *NotificationAllOf) SetSummaryArg(v string)`

SetSummaryArg sets SummaryArg field to given value.

### HasSummaryArg

`func (o *NotificationAllOf) HasSummaryArg() bool`

HasSummaryArg returns a boolean if a field has been set.

### GetSummaryArgCount

`func (o *NotificationAllOf) GetSummaryArgCount() int32`

GetSummaryArgCount returns the SummaryArgCount field if non-nil, zero value otherwise.

### GetSummaryArgCountOk

`func (o *NotificationAllOf) GetSummaryArgCountOk() (*int32, bool)`

GetSummaryArgCountOk returns a tuple with the SummaryArgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryArgCount

`func (o *NotificationAllOf) SetSummaryArgCount(v int32)`

SetSummaryArgCount sets SummaryArgCount field to given value.

### HasSummaryArgCount

`func (o *NotificationAllOf) HasSummaryArgCount() bool`

HasSummaryArgCount returns a boolean if a field has been set.

### GetEmailSubect

`func (o *NotificationAllOf) GetEmailSubect() string`

GetEmailSubect returns the EmailSubect field if non-nil, zero value otherwise.

### GetEmailSubectOk

`func (o *NotificationAllOf) GetEmailSubectOk() (*string, bool)`

GetEmailSubectOk returns a tuple with the EmailSubect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubect

`func (o *NotificationAllOf) SetEmailSubect(v string)`

SetEmailSubect sets EmailSubect field to given value.

### HasEmailSubect

`func (o *NotificationAllOf) HasEmailSubect() bool`

HasEmailSubect returns a boolean if a field has been set.

### GetEmailBody

`func (o *NotificationAllOf) GetEmailBody() string`

GetEmailBody returns the EmailBody field if non-nil, zero value otherwise.

### GetEmailBodyOk

`func (o *NotificationAllOf) GetEmailBodyOk() (*string, bool)`

GetEmailBodyOk returns a tuple with the EmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailBody

`func (o *NotificationAllOf) SetEmailBody(v string)`

SetEmailBody sets EmailBody field to given value.

### HasEmailBody

`func (o *NotificationAllOf) HasEmailBody() bool`

HasEmailBody returns a boolean if a field has been set.

### GetEmailFromName

`func (o *NotificationAllOf) GetEmailFromName() string`

GetEmailFromName returns the EmailFromName field if non-nil, zero value otherwise.

### GetEmailFromNameOk

`func (o *NotificationAllOf) GetEmailFromNameOk() (*string, bool)`

GetEmailFromNameOk returns a tuple with the EmailFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromName

`func (o *NotificationAllOf) SetEmailFromName(v string)`

SetEmailFromName sets EmailFromName field to given value.

### HasEmailFromName

`func (o *NotificationAllOf) HasEmailFromName() bool`

HasEmailFromName returns a boolean if a field has been set.

### GetEmailFromAddress

`func (o *NotificationAllOf) GetEmailFromAddress() string`

GetEmailFromAddress returns the EmailFromAddress field if non-nil, zero value otherwise.

### GetEmailFromAddressOk

`func (o *NotificationAllOf) GetEmailFromAddressOk() (*string, bool)`

GetEmailFromAddressOk returns a tuple with the EmailFromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromAddress

`func (o *NotificationAllOf) SetEmailFromAddress(v string)`

SetEmailFromAddress sets EmailFromAddress field to given value.

### HasEmailFromAddress

`func (o *NotificationAllOf) HasEmailFromAddress() bool`

HasEmailFromAddress returns a boolean if a field has been set.

### GetSmsFrom

`func (o *NotificationAllOf) GetSmsFrom() string`

GetSmsFrom returns the SmsFrom field if non-nil, zero value otherwise.

### GetSmsFromOk

`func (o *NotificationAllOf) GetSmsFromOk() (*string, bool)`

GetSmsFromOk returns a tuple with the SmsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFrom

`func (o *NotificationAllOf) SetSmsFrom(v string)`

SetSmsFrom sets SmsFrom field to given value.

### HasSmsFrom

`func (o *NotificationAllOf) HasSmsFrom() bool`

HasSmsFrom returns a boolean if a field has been set.

### GetSmsMediaUrls

`func (o *NotificationAllOf) GetSmsMediaUrls() []string`

GetSmsMediaUrls returns the SmsMediaUrls field if non-nil, zero value otherwise.

### GetSmsMediaUrlsOk

`func (o *NotificationAllOf) GetSmsMediaUrlsOk() (*[]string, bool)`

GetSmsMediaUrlsOk returns a tuple with the SmsMediaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMediaUrls

`func (o *NotificationAllOf) SetSmsMediaUrls(v []string)`

SetSmsMediaUrls sets SmsMediaUrls field to given value.

### HasSmsMediaUrls

`func (o *NotificationAllOf) HasSmsMediaUrls() bool`

HasSmsMediaUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


