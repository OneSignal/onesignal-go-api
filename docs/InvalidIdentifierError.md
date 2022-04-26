# InvalidIdentifierError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidExternalUserIds** | Pointer to **[]string** | Returned if using include_external_user_ids | [optional] 
**InvalidPlayerIds** | Pointer to **[]string** | Returned if using include_player_ids and some were valid and others were not. | [optional] 

## Methods

### NewInvalidIdentifierError

`func NewInvalidIdentifierError() *InvalidIdentifierError`

NewInvalidIdentifierError instantiates a new InvalidIdentifierError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidIdentifierErrorWithDefaults

`func NewInvalidIdentifierErrorWithDefaults() *InvalidIdentifierError`

NewInvalidIdentifierErrorWithDefaults instantiates a new InvalidIdentifierError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidExternalUserIds

`func (o *InvalidIdentifierError) GetInvalidExternalUserIds() []string`

GetInvalidExternalUserIds returns the InvalidExternalUserIds field if non-nil, zero value otherwise.

### GetInvalidExternalUserIdsOk

`func (o *InvalidIdentifierError) GetInvalidExternalUserIdsOk() (*[]string, bool)`

GetInvalidExternalUserIdsOk returns a tuple with the InvalidExternalUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidExternalUserIds

`func (o *InvalidIdentifierError) SetInvalidExternalUserIds(v []string)`

SetInvalidExternalUserIds sets InvalidExternalUserIds field to given value.

### HasInvalidExternalUserIds

`func (o *InvalidIdentifierError) HasInvalidExternalUserIds() bool`

HasInvalidExternalUserIds returns a boolean if a field has been set.

### GetInvalidPlayerIds

`func (o *InvalidIdentifierError) GetInvalidPlayerIds() []string`

GetInvalidPlayerIds returns the InvalidPlayerIds field if non-nil, zero value otherwise.

### GetInvalidPlayerIdsOk

`func (o *InvalidIdentifierError) GetInvalidPlayerIdsOk() (*[]string, bool)`

GetInvalidPlayerIdsOk returns a tuple with the InvalidPlayerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidPlayerIds

`func (o *InvalidIdentifierError) SetInvalidPlayerIds(v []string)`

SetInvalidPlayerIds sets InvalidPlayerIds field to given value.

### HasInvalidPlayerIds

`func (o *InvalidIdentifierError) HasInvalidPlayerIds() bool`

HasInvalidPlayerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


