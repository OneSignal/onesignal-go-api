# ExportPlayersRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraFields** | Pointer to **[]string** | Additional fields that you wish to include. Currently supports location, country, rooted, notification_types, ip, external_user_id, web_auth, and web_p256. | [optional] 
**LastActiveSince** | Pointer to **int32** | Export all devices with a last_active timestamp greater than this time.  Unixtime in seconds. | [optional] 
**SegmentName** | Pointer to **string** | Export al ldevices belonging to the segment. | [optional] 

## Methods

### NewExportPlayersRequestBody

`func NewExportPlayersRequestBody() *ExportPlayersRequestBody`

NewExportPlayersRequestBody instantiates a new ExportPlayersRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportPlayersRequestBodyWithDefaults

`func NewExportPlayersRequestBodyWithDefaults() *ExportPlayersRequestBody`

NewExportPlayersRequestBodyWithDefaults instantiates a new ExportPlayersRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraFields

`func (o *ExportPlayersRequestBody) GetExtraFields() []string`

GetExtraFields returns the ExtraFields field if non-nil, zero value otherwise.

### GetExtraFieldsOk

`func (o *ExportPlayersRequestBody) GetExtraFieldsOk() (*[]string, bool)`

GetExtraFieldsOk returns a tuple with the ExtraFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFields

`func (o *ExportPlayersRequestBody) SetExtraFields(v []string)`

SetExtraFields sets ExtraFields field to given value.

### HasExtraFields

`func (o *ExportPlayersRequestBody) HasExtraFields() bool`

HasExtraFields returns a boolean if a field has been set.

### GetLastActiveSince

`func (o *ExportPlayersRequestBody) GetLastActiveSince() int32`

GetLastActiveSince returns the LastActiveSince field if non-nil, zero value otherwise.

### GetLastActiveSinceOk

`func (o *ExportPlayersRequestBody) GetLastActiveSinceOk() (*int32, bool)`

GetLastActiveSinceOk returns a tuple with the LastActiveSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveSince

`func (o *ExportPlayersRequestBody) SetLastActiveSince(v int32)`

SetLastActiveSince sets LastActiveSince field to given value.

### HasLastActiveSince

`func (o *ExportPlayersRequestBody) HasLastActiveSince() bool`

HasLastActiveSince returns a boolean if a field has been set.

### GetSegmentName

`func (o *ExportPlayersRequestBody) GetSegmentName() string`

GetSegmentName returns the SegmentName field if non-nil, zero value otherwise.

### GetSegmentNameOk

`func (o *ExportPlayersRequestBody) GetSegmentNameOk() (*string, bool)`

GetSegmentNameOk returns a tuple with the SegmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentName

`func (o *ExportPlayersRequestBody) SetSegmentName(v string)`

SetSegmentName sets SegmentName field to given value.

### HasSegmentName

`func (o *ExportPlayersRequestBody) HasSegmentName() bool`

HasSegmentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


