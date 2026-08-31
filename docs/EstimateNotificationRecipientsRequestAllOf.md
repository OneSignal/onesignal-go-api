# EstimateNotificationRecipientsRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | The OneSignal App ID for your app, which can be found in Keys &amp; IDs. | [optional] 
**Filters** | Pointer to [**[]FilterExpression**](FilterExpression.md) |  | [optional] 
**IncludeAliases** | Pointer to **map[string][]string** | Target specific users by aliases assigned via API. An alias can be an external_id, onesignal_id, or a custom alias. Accepts an object where keys are alias labels and values are arrays of alias IDs to include Example usage: { \&quot;external_id\&quot;: [\&quot;exId1\&quot;, \&quot;extId2\&quot;], \&quot;internal_label\&quot;: [\&quot;id1\&quot;, \&quot;id2\&quot;] } Keys must match API spellings exactly (for example the label for External ID is the string &#x60;external_id&#x60;; arbitrary keys such as camelCase variants are not aliases and may yield no recipients). Not compatible with any other targeting parameters. REQUIRED: REST API Key Authentication Limit of 2,000 entries per REST API call Note: If targeting push, email, or sms subscribers with same ids, use with target_channel to indicate you are sending a push or email or sms. | [optional] 
**TargetChannel** | Pointer to **string** | Which platforms to count recipients for. Selects the same default platforms Create notification would use for the channel. Individual platform flags (&#x60;isIos&#x60;, &#x60;isAndroid&#x60;, etc.) are not supported by this endpoint. | [optional] 

## Methods

### NewEstimateNotificationRecipientsRequestAllOf

`func NewEstimateNotificationRecipientsRequestAllOf() *EstimateNotificationRecipientsRequestAllOf`

NewEstimateNotificationRecipientsRequestAllOf instantiates a new EstimateNotificationRecipientsRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateNotificationRecipientsRequestAllOfWithDefaults

`func NewEstimateNotificationRecipientsRequestAllOfWithDefaults() *EstimateNotificationRecipientsRequestAllOf`

NewEstimateNotificationRecipientsRequestAllOfWithDefaults instantiates a new EstimateNotificationRecipientsRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *EstimateNotificationRecipientsRequestAllOf) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *EstimateNotificationRecipientsRequestAllOf) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *EstimateNotificationRecipientsRequestAllOf) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *EstimateNotificationRecipientsRequestAllOf) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetFilters

`func (o *EstimateNotificationRecipientsRequestAllOf) GetFilters() []FilterExpression`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *EstimateNotificationRecipientsRequestAllOf) GetFiltersOk() (*[]FilterExpression, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *EstimateNotificationRecipientsRequestAllOf) SetFilters(v []FilterExpression)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *EstimateNotificationRecipientsRequestAllOf) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *EstimateNotificationRecipientsRequestAllOf) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *EstimateNotificationRecipientsRequestAllOf) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetIncludeAliases

`func (o *EstimateNotificationRecipientsRequestAllOf) GetIncludeAliases() map[string][]string`

GetIncludeAliases returns the IncludeAliases field if non-nil, zero value otherwise.

### GetIncludeAliasesOk

`func (o *EstimateNotificationRecipientsRequestAllOf) GetIncludeAliasesOk() (*map[string][]string, bool)`

GetIncludeAliasesOk returns a tuple with the IncludeAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAliases

`func (o *EstimateNotificationRecipientsRequestAllOf) SetIncludeAliases(v map[string][]string)`

SetIncludeAliases sets IncludeAliases field to given value.

### HasIncludeAliases

`func (o *EstimateNotificationRecipientsRequestAllOf) HasIncludeAliases() bool`

HasIncludeAliases returns a boolean if a field has been set.

### SetIncludeAliasesNil

`func (o *EstimateNotificationRecipientsRequestAllOf) SetIncludeAliasesNil(b bool)`

 SetIncludeAliasesNil sets the value for IncludeAliases to be an explicit nil

### UnsetIncludeAliases
`func (o *EstimateNotificationRecipientsRequestAllOf) UnsetIncludeAliases()`

UnsetIncludeAliases ensures that no value is present for IncludeAliases, not even an explicit nil
### GetTargetChannel

`func (o *EstimateNotificationRecipientsRequestAllOf) GetTargetChannel() string`

GetTargetChannel returns the TargetChannel field if non-nil, zero value otherwise.

### GetTargetChannelOk

`func (o *EstimateNotificationRecipientsRequestAllOf) GetTargetChannelOk() (*string, bool)`

GetTargetChannelOk returns a tuple with the TargetChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChannel

`func (o *EstimateNotificationRecipientsRequestAllOf) SetTargetChannel(v string)`

SetTargetChannel sets TargetChannel field to given value.

### HasTargetChannel

`func (o *EstimateNotificationRecipientsRequestAllOf) HasTargetChannel() bool`

HasTargetChannel returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


