# UpdateSegmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Required. The segment name. Maximum 128 characters. | 
**Description** | Pointer to **string** | Optional human-readable description for the segment. Maximum 255 characters. Pass an empty string to clear; omit to leave unchanged. | [optional] 
**Filters** | Pointer to [**[]FilterExpression**](FilterExpression.md) | Optional. When provided, replaces all existing filters. Filters define the segment based on user properties like tags, activity, or location using flexible AND/OR logic. Limited to 200 total entries, including fields and OR operators. | [optional] 

## Methods

### NewUpdateSegmentRequest

`func NewUpdateSegmentRequest(name string, ) *UpdateSegmentRequest`

NewUpdateSegmentRequest instantiates a new UpdateSegmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSegmentRequestWithDefaults

`func NewUpdateSegmentRequestWithDefaults() *UpdateSegmentRequest`

NewUpdateSegmentRequestWithDefaults instantiates a new UpdateSegmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSegmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSegmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSegmentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateSegmentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSegmentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSegmentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSegmentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilters

`func (o *UpdateSegmentRequest) GetFilters() []FilterExpression`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UpdateSegmentRequest) GetFiltersOk() (*[]FilterExpression, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UpdateSegmentRequest) SetFilters(v []FilterExpression)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UpdateSegmentRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)


