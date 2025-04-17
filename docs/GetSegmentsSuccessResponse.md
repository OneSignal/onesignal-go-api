# GetSegmentsSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** | The number of Segments in the response. | [optional] 
**Offset** | Pointer to **int32** | Set with the offset query parameter. Default 0. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of Segments returned. Default 300. | [optional] 
**Segments** | Pointer to [**[]SegmentData**](SegmentData.md) | An array containing the Segment information. | [optional] 

## Methods

### NewGetSegmentsSuccessResponse

`func NewGetSegmentsSuccessResponse() *GetSegmentsSuccessResponse`

NewGetSegmentsSuccessResponse instantiates a new GetSegmentsSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSegmentsSuccessResponseWithDefaults

`func NewGetSegmentsSuccessResponseWithDefaults() *GetSegmentsSuccessResponse`

NewGetSegmentsSuccessResponseWithDefaults instantiates a new GetSegmentsSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *GetSegmentsSuccessResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetSegmentsSuccessResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetSegmentsSuccessResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetSegmentsSuccessResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetOffset

`func (o *GetSegmentsSuccessResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetSegmentsSuccessResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetSegmentsSuccessResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetSegmentsSuccessResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *GetSegmentsSuccessResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetSegmentsSuccessResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetSegmentsSuccessResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetSegmentsSuccessResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSegments

`func (o *GetSegmentsSuccessResponse) GetSegments() []SegmentData`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *GetSegmentsSuccessResponse) GetSegmentsOk() (*[]SegmentData, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *GetSegmentsSuccessResponse) SetSegments(v []SegmentData)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *GetSegmentsSuccessResponse) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


