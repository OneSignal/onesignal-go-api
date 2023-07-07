# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to [**PropertiesObject**](PropertiesObject.md) |  | [optional] 
**RefreshDeviceMetadata** | Pointer to **bool** |  | [optional] [default to false]
**Deltas** | Pointer to [**PropertiesDeltas**](PropertiesDeltas.md) |  | [optional] 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest() *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *UpdateUserRequest) GetProperties() PropertiesObject`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UpdateUserRequest) GetPropertiesOk() (*PropertiesObject, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UpdateUserRequest) SetProperties(v PropertiesObject)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *UpdateUserRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRefreshDeviceMetadata

`func (o *UpdateUserRequest) GetRefreshDeviceMetadata() bool`

GetRefreshDeviceMetadata returns the RefreshDeviceMetadata field if non-nil, zero value otherwise.

### GetRefreshDeviceMetadataOk

`func (o *UpdateUserRequest) GetRefreshDeviceMetadataOk() (*bool, bool)`

GetRefreshDeviceMetadataOk returns a tuple with the RefreshDeviceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshDeviceMetadata

`func (o *UpdateUserRequest) SetRefreshDeviceMetadata(v bool)`

SetRefreshDeviceMetadata sets RefreshDeviceMetadata field to given value.

### HasRefreshDeviceMetadata

`func (o *UpdateUserRequest) HasRefreshDeviceMetadata() bool`

HasRefreshDeviceMetadata returns a boolean if a field has been set.

### GetDeltas

`func (o *UpdateUserRequest) GetDeltas() PropertiesDeltas`

GetDeltas returns the Deltas field if non-nil, zero value otherwise.

### GetDeltasOk

`func (o *UpdateUserRequest) GetDeltasOk() (*PropertiesDeltas, bool)`

GetDeltasOk returns a tuple with the Deltas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltas

`func (o *UpdateUserRequest) SetDeltas(v PropertiesDeltas)`

SetDeltas sets Deltas field to given value.

### HasDeltas

`func (o *UpdateUserRequest) HasDeltas() bool`

HasDeltas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


