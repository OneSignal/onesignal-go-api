# Notification200Errors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidExternalUserIds** | Pointer to **[]string** | Returned if using include_external_user_ids | [optional] 
**InvalidPlayerIds** | Pointer to **[]string** | Returned if using include_player_ids and some were valid and others were not. | [optional] 

## Methods

### NewNotification200Errors

`func NewNotification200Errors() *Notification200Errors`

NewNotification200Errors instantiates a new Notification200Errors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotification200ErrorsWithDefaults

`func NewNotification200ErrorsWithDefaults() *Notification200Errors`

NewNotification200ErrorsWithDefaults instantiates a new Notification200Errors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidExternalUserIds

`func (o *Notification200Errors) GetInvalidExternalUserIds() []string`

GetInvalidExternalUserIds returns the InvalidExternalUserIds field if non-nil, zero value otherwise.

### GetInvalidExternalUserIdsOk

`func (o *Notification200Errors) GetInvalidExternalUserIdsOk() (*[]string, bool)`

GetInvalidExternalUserIdsOk returns a tuple with the InvalidExternalUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidExternalUserIds

`func (o *Notification200Errors) SetInvalidExternalUserIds(v []string)`

SetInvalidExternalUserIds sets InvalidExternalUserIds field to given value.

### HasInvalidExternalUserIds

`func (o *Notification200Errors) HasInvalidExternalUserIds() bool`

HasInvalidExternalUserIds returns a boolean if a field has been set.

### GetInvalidPlayerIds

`func (o *Notification200Errors) GetInvalidPlayerIds() []string`

GetInvalidPlayerIds returns the InvalidPlayerIds field if non-nil, zero value otherwise.

### GetInvalidPlayerIdsOk

`func (o *Notification200Errors) GetInvalidPlayerIdsOk() (*[]string, bool)`

GetInvalidPlayerIdsOk returns a tuple with the InvalidPlayerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidPlayerIds

`func (o *Notification200Errors) SetInvalidPlayerIds(v []string)`

SetInvalidPlayerIds sets InvalidPlayerIds field to given value.

### HasInvalidPlayerIds

`func (o *Notification200Errors) HasInvalidPlayerIds() bool`

HasInvalidPlayerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


