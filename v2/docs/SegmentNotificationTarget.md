# SegmentNotificationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedSegments** | Pointer to **[]string** | The segment names you want to target. Users in these segments will receive a notification. This targeting parameter is only compatible with excluded_segments. Example: [\&quot;Active Users\&quot;, \&quot;Inactive Users\&quot;]  | [optional] 
**ExcludedSegments** | Pointer to **[]string** | Segment that will be excluded when sending. Users in these segments will not receive a notification, even if they were included in included_segments. This targeting parameter is only compatible with included_segments. Example: [\&quot;Active Users\&quot;, \&quot;Inactive Users\&quot;]  | [optional] 

## Methods

### NewSegmentNotificationTarget

`func NewSegmentNotificationTarget() *SegmentNotificationTarget`

NewSegmentNotificationTarget instantiates a new SegmentNotificationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentNotificationTargetWithDefaults

`func NewSegmentNotificationTargetWithDefaults() *SegmentNotificationTarget`

NewSegmentNotificationTargetWithDefaults instantiates a new SegmentNotificationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedSegments

`func (o *SegmentNotificationTarget) GetIncludedSegments() []string`

GetIncludedSegments returns the IncludedSegments field if non-nil, zero value otherwise.

### GetIncludedSegmentsOk

`func (o *SegmentNotificationTarget) GetIncludedSegmentsOk() (*[]string, bool)`

GetIncludedSegmentsOk returns a tuple with the IncludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegments

`func (o *SegmentNotificationTarget) SetIncludedSegments(v []string)`

SetIncludedSegments sets IncludedSegments field to given value.

### HasIncludedSegments

`func (o *SegmentNotificationTarget) HasIncludedSegments() bool`

HasIncludedSegments returns a boolean if a field has been set.

### GetExcludedSegments

`func (o *SegmentNotificationTarget) GetExcludedSegments() []string`

GetExcludedSegments returns the ExcludedSegments field if non-nil, zero value otherwise.

### GetExcludedSegmentsOk

`func (o *SegmentNotificationTarget) GetExcludedSegmentsOk() (*[]string, bool)`

GetExcludedSegmentsOk returns a tuple with the ExcludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegments

`func (o *SegmentNotificationTarget) SetExcludedSegments(v []string)`

SetExcludedSegments sets ExcludedSegments field to given value.

### HasExcludedSegments

`func (o *SegmentNotificationTarget) HasExcludedSegments() bool`

HasExcludedSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


