# UserIdentityBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUserIdentityBody

`func NewUserIdentityBody() *UserIdentityBody`

NewUserIdentityBody instantiates a new UserIdentityBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentityBodyWithDefaults

`func NewUserIdentityBodyWithDefaults() *UserIdentityBody`

NewUserIdentityBodyWithDefaults instantiates a new UserIdentityBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *UserIdentityBody) GetIdentity() map[string]string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *UserIdentityBody) GetIdentityOk() (*map[string]string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *UserIdentityBody) SetIdentity(v map[string]string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *UserIdentityBody) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


