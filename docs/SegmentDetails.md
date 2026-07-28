# SegmentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the segment (UUID v4). | [optional] 
**Name** | Pointer to **string** | The segment name. | [optional] 
**Description** | Pointer to **NullableString** | Human-readable description for the segment. &#x60;null&#x60; when unset. Maximum 255 characters. | [optional] 
**CreatedAt** | Pointer to **int32** | Unix timestamp when the segment was created. | [optional] 
**Source** | Pointer to **string** | The source of the segment. | [optional] 
**Filters** | Pointer to [**[]FilterExpression**](FilterExpression.md) | Array of filter and operator objects defining the segment criteria. Uses the same format as the Create Segment API, so filters can be directly used to recreate or update the segment. | [optional] 

## Methods

### NewSegmentDetails

`func NewSegmentDetails() *SegmentDetails`

NewSegmentDetails instantiates a new SegmentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentDetailsWithDefaults

`func NewSegmentDetailsWithDefaults() *SegmentDetails`

NewSegmentDetailsWithDefaults instantiates a new SegmentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SegmentDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SegmentDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SegmentDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SegmentDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SegmentDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SegmentDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SegmentDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SegmentDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SegmentDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SegmentDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatedAt

`func (o *SegmentDetails) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SegmentDetails) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SegmentDetails) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SegmentDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSource

`func (o *SegmentDetails) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SegmentDetails) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SegmentDetails) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SegmentDetails) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetFilters

`func (o *SegmentDetails) GetFilters() []FilterExpression`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SegmentDetails) GetFiltersOk() (*[]FilterExpression, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SegmentDetails) SetFilters(v []FilterExpression)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SegmentDetails) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


