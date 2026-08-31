# EstimateNotificationRecipientsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedSegments** | Pointer to **[]string** | The segment names you want to target. Users in these segments will receive a notification. This targeting parameter is only compatible with excluded_segments. Example: [\&quot;Active Users\&quot;, \&quot;Inactive Users\&quot;] &#x60;\&quot;All\&quot;&#x60; is a shorthand for every subscribed user: if the array includes the string &#x60;\&quot;All\&quot;&#x60; and the app has no segment actually named &#x60;All&#x60;, it targets all subscribers instead of a literal segment lookup.  | [optional] 
**ExcludedSegments** | Pointer to **[]string** | Segment that will be excluded when sending. Users in these segments will not receive a notification, even if they were included in included_segments. This targeting parameter is only compatible with included_segments. Example: [\&quot;Active Users\&quot;, \&quot;Inactive Users\&quot;]  | [optional] 
**AppId** | **string** | The OneSignal App ID for your app, which can be found in Keys &amp; IDs. | 
**Filters** | Pointer to [**[]FilterExpression**](FilterExpression.md) |  | [optional] 
**IncludeAliases** | Pointer to **map[string][]string** | Target specific users by aliases assigned via API. An alias can be an external_id, onesignal_id, or a custom alias. Accepts an object where keys are alias labels and values are arrays of alias IDs to include Example usage: { \&quot;external_id\&quot;: [\&quot;exId1\&quot;, \&quot;extId2\&quot;], \&quot;internal_label\&quot;: [\&quot;id1\&quot;, \&quot;id2\&quot;] } Keys must match API spellings exactly (for example the label for External ID is the string &#x60;external_id&#x60;; arbitrary keys such as camelCase variants are not aliases and may yield no recipients). Not compatible with any other targeting parameters. REQUIRED: REST API Key Authentication Limit of 2,000 entries per REST API call Note: If targeting push, email, or sms subscribers with same ids, use with target_channel to indicate you are sending a push or email or sms. | [optional] 
**TargetChannel** | Pointer to **string** | Which platforms to count recipients for. Selects the same default platforms Create notification would use for the channel. Individual platform flags (&#x60;isIos&#x60;, &#x60;isAndroid&#x60;, etc.) are not supported by this endpoint. | [optional] 

## Methods

### NewEstimateNotificationRecipientsRequest

`func NewEstimateNotificationRecipientsRequest(appId string, ) *EstimateNotificationRecipientsRequest`

NewEstimateNotificationRecipientsRequest instantiates a new EstimateNotificationRecipientsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateNotificationRecipientsRequestWithDefaults

`func NewEstimateNotificationRecipientsRequestWithDefaults() *EstimateNotificationRecipientsRequest`

NewEstimateNotificationRecipientsRequestWithDefaults instantiates a new EstimateNotificationRecipientsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedSegments

`func (o *EstimateNotificationRecipientsRequest) GetIncludedSegments() []string`

GetIncludedSegments returns the IncludedSegments field if non-nil, zero value otherwise.

### GetIncludedSegmentsOk

`func (o *EstimateNotificationRecipientsRequest) GetIncludedSegmentsOk() (*[]string, bool)`

GetIncludedSegmentsOk returns a tuple with the IncludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegments

`func (o *EstimateNotificationRecipientsRequest) SetIncludedSegments(v []string)`

SetIncludedSegments sets IncludedSegments field to given value.

### HasIncludedSegments

`func (o *EstimateNotificationRecipientsRequest) HasIncludedSegments() bool`

HasIncludedSegments returns a boolean if a field has been set.

### GetExcludedSegments

`func (o *EstimateNotificationRecipientsRequest) GetExcludedSegments() []string`

GetExcludedSegments returns the ExcludedSegments field if non-nil, zero value otherwise.

### GetExcludedSegmentsOk

`func (o *EstimateNotificationRecipientsRequest) GetExcludedSegmentsOk() (*[]string, bool)`

GetExcludedSegmentsOk returns a tuple with the ExcludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegments

`func (o *EstimateNotificationRecipientsRequest) SetExcludedSegments(v []string)`

SetExcludedSegments sets ExcludedSegments field to given value.

### HasExcludedSegments

`func (o *EstimateNotificationRecipientsRequest) HasExcludedSegments() bool`

HasExcludedSegments returns a boolean if a field has been set.

### GetAppId

`func (o *EstimateNotificationRecipientsRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *EstimateNotificationRecipientsRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *EstimateNotificationRecipientsRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetFilters

`func (o *EstimateNotificationRecipientsRequest) GetFilters() []FilterExpression`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *EstimateNotificationRecipientsRequest) GetFiltersOk() (*[]FilterExpression, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *EstimateNotificationRecipientsRequest) SetFilters(v []FilterExpression)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *EstimateNotificationRecipientsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *EstimateNotificationRecipientsRequest) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *EstimateNotificationRecipientsRequest) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetIncludeAliases

`func (o *EstimateNotificationRecipientsRequest) GetIncludeAliases() map[string][]string`

GetIncludeAliases returns the IncludeAliases field if non-nil, zero value otherwise.

### GetIncludeAliasesOk

`func (o *EstimateNotificationRecipientsRequest) GetIncludeAliasesOk() (*map[string][]string, bool)`

GetIncludeAliasesOk returns a tuple with the IncludeAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAliases

`func (o *EstimateNotificationRecipientsRequest) SetIncludeAliases(v map[string][]string)`

SetIncludeAliases sets IncludeAliases field to given value.

### HasIncludeAliases

`func (o *EstimateNotificationRecipientsRequest) HasIncludeAliases() bool`

HasIncludeAliases returns a boolean if a field has been set.

### SetIncludeAliasesNil

`func (o *EstimateNotificationRecipientsRequest) SetIncludeAliasesNil(b bool)`

 SetIncludeAliasesNil sets the value for IncludeAliases to be an explicit nil

### UnsetIncludeAliases
`func (o *EstimateNotificationRecipientsRequest) UnsetIncludeAliases()`

UnsetIncludeAliases ensures that no value is present for IncludeAliases, not even an explicit nil
### GetTargetChannel

`func (o *EstimateNotificationRecipientsRequest) GetTargetChannel() string`

GetTargetChannel returns the TargetChannel field if non-nil, zero value otherwise.

### GetTargetChannelOk

`func (o *EstimateNotificationRecipientsRequest) GetTargetChannelOk() (*string, bool)`

GetTargetChannelOk returns a tuple with the TargetChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChannel

`func (o *EstimateNotificationRecipientsRequest) SetTargetChannel(v string)`

SetTargetChannel sets TargetChannel field to given value.

### HasTargetChannel

`func (o *EstimateNotificationRecipientsRequest) HasTargetChannel() bool`

HasTargetChannel returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


