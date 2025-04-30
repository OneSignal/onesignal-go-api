# ExportSubscriptionsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraFields** | Pointer to **[]string** | Additional fields that you wish to include. Currently supports location, country, rooted, notification_types, ip, external_user_id, web_auth, and web_p256. | [optional] 
**LastActiveSince** | Pointer to **string** | Export all devices with a last_active timestamp greater than this time.  Unixtime in seconds. | [optional] 
**SegmentName** | Pointer to **string** | Export all devices belonging to the segment. | [optional] 

## Methods

### NewExportSubscriptionsRequestBody

`func NewExportSubscriptionsRequestBody() *ExportSubscriptionsRequestBody`

NewExportSubscriptionsRequestBody instantiates a new ExportSubscriptionsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportSubscriptionsRequestBodyWithDefaults

`func NewExportSubscriptionsRequestBodyWithDefaults() *ExportSubscriptionsRequestBody`

NewExportSubscriptionsRequestBodyWithDefaults instantiates a new ExportSubscriptionsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraFields

`func (o *ExportSubscriptionsRequestBody) GetExtraFields() []string`

GetExtraFields returns the ExtraFields field if non-nil, zero value otherwise.

### GetExtraFieldsOk

`func (o *ExportSubscriptionsRequestBody) GetExtraFieldsOk() (*[]string, bool)`

GetExtraFieldsOk returns a tuple with the ExtraFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFields

`func (o *ExportSubscriptionsRequestBody) SetExtraFields(v []string)`

SetExtraFields sets ExtraFields field to given value.

### HasExtraFields

`func (o *ExportSubscriptionsRequestBody) HasExtraFields() bool`

HasExtraFields returns a boolean if a field has been set.

### GetLastActiveSince

`func (o *ExportSubscriptionsRequestBody) GetLastActiveSince() string`

GetLastActiveSince returns the LastActiveSince field if non-nil, zero value otherwise.

### GetLastActiveSinceOk

`func (o *ExportSubscriptionsRequestBody) GetLastActiveSinceOk() (*string, bool)`

GetLastActiveSinceOk returns a tuple with the LastActiveSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveSince

`func (o *ExportSubscriptionsRequestBody) SetLastActiveSince(v string)`

SetLastActiveSince sets LastActiveSince field to given value.

### HasLastActiveSince

`func (o *ExportSubscriptionsRequestBody) HasLastActiveSince() bool`

HasLastActiveSince returns a boolean if a field has been set.

### GetSegmentName

`func (o *ExportSubscriptionsRequestBody) GetSegmentName() string`

GetSegmentName returns the SegmentName field if non-nil, zero value otherwise.

### GetSegmentNameOk

`func (o *ExportSubscriptionsRequestBody) GetSegmentNameOk() (*string, bool)`

GetSegmentNameOk returns a tuple with the SegmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentName

`func (o *ExportSubscriptionsRequestBody) SetSegmentName(v string)`

SetSegmentName sets SegmentName field to given value.

### HasSegmentName

`func (o *ExportSubscriptionsRequestBody) HasSegmentName() bool`

HasSegmentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


