# Segment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | UUID of the segment.  If left empty, it will be assigned automaticaly. | [optional] 
**Name** | **string** | Name of the segment.  You&#39;ll see this name on the Web UI. | 
**Filters** | [**[]FilterExpressions**](FilterExpressions.md) | Filter or operators the segment will have.  For a list of available filters with details, please see Send to Users Based on Filters. | 

## Methods

### NewSegment

`func NewSegment(name string, filters []FilterExpressions, ) *Segment`

NewSegment instantiates a new Segment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentWithDefaults

`func NewSegmentWithDefaults() *Segment`

NewSegmentWithDefaults instantiates a new Segment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Segment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Segment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Segment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Segment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Segment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Segment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Segment) SetName(v string)`

SetName sets Name field to given value.


### GetFilters

`func (o *Segment) GetFilters() []FilterExpressions`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Segment) GetFiltersOk() (*[]FilterExpressions, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Segment) SetFilters(v []FilterExpressions)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


