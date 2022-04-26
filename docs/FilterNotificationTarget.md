# FilterNotificationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSession** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot; or \&quot;&lt;\&quot; hours_ago &#x3D; number of hours before or after the users last session. Example: \&quot;1.1\&quot;  | [optional] 
**FirstSession** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot; or \&quot;&lt;\&quot; hours_ago &#x3D; number of hours before or after the users first session. Example: \&quot;1.1\&quot;  | [optional] 
**SessionCount** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot;, \&quot;&#x3D;\&quot; or \&quot;!&#x3D;\&quot; value &#x3D; number sessions. Example: \&quot;1\&quot;  | [optional] 
**SessionTime** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot;, \&quot;&#x3D;\&quot; or \&quot;!&#x3D;\&quot; value &#x3D; Time in seconds the user has been in your app. Example: \&quot;3600\&quot;  | [optional] 
**AmountSpent** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot;, or \&quot;&#x3D;\&quot; value &#x3D; Amount in USD a user has spent on IAP (In App Purchases). Example: \&quot;0.99\&quot;  | [optional] 
**BoughtSku** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot; or \&quot;&#x3D;\&quot; key &#x3D; SKU purchased in your app as an IAP (In App Purchases). Example: \&quot;com.domain.100coinpack\&quot; value &#x3D; value of SKU to compare to. Example: \&quot;0.99\&quot;  | [optional] 
**Tag** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot;, \&quot;&#x3D;\&quot;, \&quot;!&#x3D;\&quot;, \&quot;exists\&quot;, \&quot;not_exists\&quot;, \&quot;time_elapsed_gt\&quot; (paid plan only) or \&quot;time_elapsed_lt\&quot; (paid plan only) See Time Operators key &#x3D; Tag key to compare. value &#x3D; Tag value to compare. Not required for \&quot;exists\&quot; or \&quot;not_exists\&quot;. Example: See Formatting Filters  | [optional] 
**Language** | Pointer to **string** | relation &#x3D; \&quot;&#x3D;\&quot; or \&quot;!&#x3D;\&quot; value &#x3D; 2 character language code. Example: \&quot;en\&quot;. For a list of all language codes see Language &amp; Localization.  | [optional] 
**AppVersion** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot;, \&quot;&#x3D;\&quot; or \&quot;!&#x3D;\&quot; value &#x3D; app version. Example: \&quot;1.0.0\&quot;  | [optional] 
**Location** | Pointer to **string** | radius &#x3D; in meters lat &#x3D; latitude long &#x3D; longitude  | [optional] 
**Email** | Pointer to **string** | value &#x3D; email address Only for sending Push Notifications Use this for targeting push subscribers associated with an email set with all SDK setEmail methods To send emails to specific email addresses use include_email_tokens parameter  | [optional] 
**Country** | Pointer to **string** | relation &#x3D; \&quot;&#x3D;\&quot; value &#x3D; 2-digit Country code Example: \&quot;field\&quot;: \&quot;country\&quot;, \&quot;relation\&quot;: \&quot;&#x3D;\&quot;, \&quot;value\&quot;, \&quot;US\&quot;  | [optional] 

## Methods

### NewFilterNotificationTarget

`func NewFilterNotificationTarget() *FilterNotificationTarget`

NewFilterNotificationTarget instantiates a new FilterNotificationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterNotificationTargetWithDefaults

`func NewFilterNotificationTargetWithDefaults() *FilterNotificationTarget`

NewFilterNotificationTargetWithDefaults instantiates a new FilterNotificationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSession

`func (o *FilterNotificationTarget) GetLastSession() string`

GetLastSession returns the LastSession field if non-nil, zero value otherwise.

### GetLastSessionOk

`func (o *FilterNotificationTarget) GetLastSessionOk() (*string, bool)`

GetLastSessionOk returns a tuple with the LastSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSession

`func (o *FilterNotificationTarget) SetLastSession(v string)`

SetLastSession sets LastSession field to given value.

### HasLastSession

`func (o *FilterNotificationTarget) HasLastSession() bool`

HasLastSession returns a boolean if a field has been set.

### GetFirstSession

`func (o *FilterNotificationTarget) GetFirstSession() string`

GetFirstSession returns the FirstSession field if non-nil, zero value otherwise.

### GetFirstSessionOk

`func (o *FilterNotificationTarget) GetFirstSessionOk() (*string, bool)`

GetFirstSessionOk returns a tuple with the FirstSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSession

`func (o *FilterNotificationTarget) SetFirstSession(v string)`

SetFirstSession sets FirstSession field to given value.

### HasFirstSession

`func (o *FilterNotificationTarget) HasFirstSession() bool`

HasFirstSession returns a boolean if a field has been set.

### GetSessionCount

`func (o *FilterNotificationTarget) GetSessionCount() string`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *FilterNotificationTarget) GetSessionCountOk() (*string, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *FilterNotificationTarget) SetSessionCount(v string)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *FilterNotificationTarget) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetSessionTime

`func (o *FilterNotificationTarget) GetSessionTime() string`

GetSessionTime returns the SessionTime field if non-nil, zero value otherwise.

### GetSessionTimeOk

`func (o *FilterNotificationTarget) GetSessionTimeOk() (*string, bool)`

GetSessionTimeOk returns a tuple with the SessionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTime

`func (o *FilterNotificationTarget) SetSessionTime(v string)`

SetSessionTime sets SessionTime field to given value.

### HasSessionTime

`func (o *FilterNotificationTarget) HasSessionTime() bool`

HasSessionTime returns a boolean if a field has been set.

### GetAmountSpent

`func (o *FilterNotificationTarget) GetAmountSpent() string`

GetAmountSpent returns the AmountSpent field if non-nil, zero value otherwise.

### GetAmountSpentOk

`func (o *FilterNotificationTarget) GetAmountSpentOk() (*string, bool)`

GetAmountSpentOk returns a tuple with the AmountSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSpent

`func (o *FilterNotificationTarget) SetAmountSpent(v string)`

SetAmountSpent sets AmountSpent field to given value.

### HasAmountSpent

`func (o *FilterNotificationTarget) HasAmountSpent() bool`

HasAmountSpent returns a boolean if a field has been set.

### GetBoughtSku

`func (o *FilterNotificationTarget) GetBoughtSku() string`

GetBoughtSku returns the BoughtSku field if non-nil, zero value otherwise.

### GetBoughtSkuOk

`func (o *FilterNotificationTarget) GetBoughtSkuOk() (*string, bool)`

GetBoughtSkuOk returns a tuple with the BoughtSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoughtSku

`func (o *FilterNotificationTarget) SetBoughtSku(v string)`

SetBoughtSku sets BoughtSku field to given value.

### HasBoughtSku

`func (o *FilterNotificationTarget) HasBoughtSku() bool`

HasBoughtSku returns a boolean if a field has been set.

### GetTag

`func (o *FilterNotificationTarget) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *FilterNotificationTarget) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *FilterNotificationTarget) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *FilterNotificationTarget) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetLanguage

`func (o *FilterNotificationTarget) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *FilterNotificationTarget) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *FilterNotificationTarget) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *FilterNotificationTarget) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetAppVersion

`func (o *FilterNotificationTarget) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *FilterNotificationTarget) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *FilterNotificationTarget) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *FilterNotificationTarget) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetLocation

`func (o *FilterNotificationTarget) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FilterNotificationTarget) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FilterNotificationTarget) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FilterNotificationTarget) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetEmail

`func (o *FilterNotificationTarget) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FilterNotificationTarget) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FilterNotificationTarget) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FilterNotificationTarget) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCountry

`func (o *FilterNotificationTarget) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FilterNotificationTarget) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FilterNotificationTarget) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *FilterNotificationTarget) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


