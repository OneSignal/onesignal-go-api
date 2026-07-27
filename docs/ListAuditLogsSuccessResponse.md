# ListAuditLogsSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogs** | Pointer to [**[]AuditLogEvent**](AuditLogEvent.md) | Array of audit log events, ordered by occurred_at ascending. | [optional] 
**HasMore** | Pointer to **bool** | True if additional events exist beyond this page. Use next_cursor to fetch the next page. | [optional] 
**NextCursor** | Pointer to **string** | Opaque cursor to pass as cursor in the next request. Only present when has_more is true. | [optional] 

## Methods

### NewListAuditLogsSuccessResponse

`func NewListAuditLogsSuccessResponse() *ListAuditLogsSuccessResponse`

NewListAuditLogsSuccessResponse instantiates a new ListAuditLogsSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuditLogsSuccessResponseWithDefaults

`func NewListAuditLogsSuccessResponseWithDefaults() *ListAuditLogsSuccessResponse`

NewListAuditLogsSuccessResponseWithDefaults instantiates a new ListAuditLogsSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogs

`func (o *ListAuditLogsSuccessResponse) GetAuditLogs() []AuditLogEvent`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *ListAuditLogsSuccessResponse) GetAuditLogsOk() (*[]AuditLogEvent, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *ListAuditLogsSuccessResponse) SetAuditLogs(v []AuditLogEvent)`

SetAuditLogs sets AuditLogs field to given value.

### HasAuditLogs

`func (o *ListAuditLogsSuccessResponse) HasAuditLogs() bool`

HasAuditLogs returns a boolean if a field has been set.

### GetHasMore

`func (o *ListAuditLogsSuccessResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListAuditLogsSuccessResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListAuditLogsSuccessResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ListAuditLogsSuccessResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNextCursor

`func (o *ListAuditLogsSuccessResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListAuditLogsSuccessResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListAuditLogsSuccessResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListAuditLogsSuccessResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


