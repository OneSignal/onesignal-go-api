# PlayerSlice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Players** | Pointer to [**[]Player**](Player.md) |  | [optional] 

## Methods

### NewPlayerSlice

`func NewPlayerSlice() *PlayerSlice`

NewPlayerSlice instantiates a new PlayerSlice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerSliceWithDefaults

`func NewPlayerSliceWithDefaults() *PlayerSlice`

NewPlayerSliceWithDefaults instantiates a new PlayerSlice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *PlayerSlice) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PlayerSlice) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PlayerSlice) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PlayerSlice) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetOffset

`func (o *PlayerSlice) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PlayerSlice) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PlayerSlice) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PlayerSlice) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PlayerSlice) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PlayerSlice) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PlayerSlice) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PlayerSlice) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPlayers

`func (o *PlayerSlice) GetPlayers() []Player`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *PlayerSlice) GetPlayersOk() (*[]Player, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *PlayerSlice) SetPlayers(v []Player)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *PlayerSlice) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


