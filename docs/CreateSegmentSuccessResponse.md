# CreateSegmentSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | UUID of created segment | [optional] 

## Methods

### NewCreateSegmentSuccessResponse

`func NewCreateSegmentSuccessResponse() *CreateSegmentSuccessResponse`

NewCreateSegmentSuccessResponse instantiates a new CreateSegmentSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSegmentSuccessResponseWithDefaults

`func NewCreateSegmentSuccessResponseWithDefaults() *CreateSegmentSuccessResponse`

NewCreateSegmentSuccessResponseWithDefaults instantiates a new CreateSegmentSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateSegmentSuccessResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateSegmentSuccessResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateSegmentSuccessResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateSegmentSuccessResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetId

`func (o *CreateSegmentSuccessResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateSegmentSuccessResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateSegmentSuccessResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateSegmentSuccessResponse) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


